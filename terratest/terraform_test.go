package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the simple Terraform module in examples/terraform-basic-example using Terratest.
// Make sure you have the dep binary, https://github.com/golang/dep
// Run 'dep ensure' before run test cases.

func TestTerraformBasicExampleNew(t *testing.T) {
	t.Parallel()
	name := "tf-ecs-rds-redis"
	description := "tf-ecs-rds-redis-description"
	availableDiskCategory := "cloud_efficiency"
	availableResourceCreation := "VSwitch"
	vpcCidrBlock := "192.168.0.0/16"
	vswitchCidrBlock := "192.168.1.0/24"
	instanceType := "ecs.n4.large"
	systemDiskCategory := "cloud_efficiency"
	systemDiskName := "system_disk"
	systemDiskDescription := "system_disk_description"
	imageId := "ubuntu_18_04_64_20G_alibase_20190624.vhd"
	internetMaxBandwidthOut := 10
	ecsSize := 1200
	category := "cloud_efficiency"
	engine := "MySQL"
	engineVersion := "5.6"
	rdsInstanceType := "rds.mysql.s2.large"
	instanceStorage := "30"
	instanceChargeType := "Postpaid"
	monitoringPeriod := "60"
	securityIps := []string{"127.0.0.1"}
	redisInstanceType := "Redis"
	redisEngineVersion := "4.0"
	redisAppendonly := "yes"
	redisLazyfreeLazyEviction := "yes"
	redisResourceGroupId := "rg-123456"
	redisInstanceClass := "redis.master.large.default"

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "./basic/",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":                         name,
			"description":                  description,
			"available_disk_category":      availableDiskCategory,
			"available_resource_creation":  availableResourceCreation,
			"vpc_cidr_block":               vpcCidrBlock,
			"vswitch_cidr_block":           vswitchCidrBlock,
			"instance_type":                instanceType,
			"system_disk_category":         systemDiskCategory,
			"system_disk_name":             systemDiskName,
			"system_disk_description":      systemDiskDescription,
			"image_id":                     imageId,
			"internet_max_bandwidth_out":   internetMaxBandwidthOut,
			"ecs_size":                     ecsSize,
			"category":                     category,
			"engine":                       engine,
			"engine_version":               engineVersion,
			"rds_instance_type":            rdsInstanceType,
			"instance_storage":             instanceStorage,
			"instance_charge_type":         instanceChargeType,
			"monitoring_period":            monitoringPeriod,
			"security_ips":                 securityIps,
			"redis_instance_type":          redisInstanceType,
			"redis_engine_version":         redisEngineVersion,
			"redis_appendonly":             redisAppendonly,
			"redis_lazyfree_lazy_eviction": redisLazyfreeLazyEviction,
			"redis_resource_group_id":      redisResourceGroupId,
			"redis_instance_class":         redisInstanceClass,
		},

		// Disable colors in Terraform commands, so it's easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	thisEcsName := terraform.Output(t, terraformOptions, "this_ecs_name")
	thisEcsInstanceType := terraform.Output(t, terraformOptions, "this_ecs_instance_type")
	thisRdsName := terraform.Output(t, terraformOptions, "this_rds_name")
	thisRdsInstanceType := terraform.Output(t, terraformOptions, "this_rds_instance_type")
	thisRedisInstanceClass := terraform.Output(t, terraformOptions, "this_redis_instance_type")

	// Verify we're getting back the outputs we expect
	assert.Equal(t, thisEcsName, name)
	assert.Equal(t, thisEcsInstanceType, instanceType)
	assert.Equal(t, thisRdsName, name)
	assert.Equal(t, thisRdsInstanceType, rdsInstanceType)
	assert.Equal(t, thisRedisInstanceClass, redisInstanceClass)
}
