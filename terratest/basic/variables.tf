variable "name" {
  description = "The specification of module name."
  type        = string
  default     = "tf-ecs-rds-redis"
}

variable "description" {
  description = "The specification of module description."
  type        = string
  default     = "tf-ecs-rds-redis-description"
}

variable "available_disk_category" {
  description = "The specification of available disk category."
  type        = string
  default     = "cloud_efficiency"
}

variable "available_resource_creation" {
  description = "The specification of available resource creation."
  type        = string
  default     = "KVStore"
}

variable "vpc_cidr_block" {
  description = "The specification of the vpc cidr block."
  type        = string
  default     = "192.168.0.0/16"
}

variable "vswitch_cidr_block" {
  description = "The specification of the vswitch cidr block."
  type        = string
  default     = "192.168.1.0/24"
}

variable "instance_type" {
  description = "The specification of the instance type."
  type        = string
  default     = "ecs.n4.large"
}

variable "system_disk_category" {
  description = "The specification of the system disk category."
  type        = string
  default     = "cloud_efficiency"
}

variable "system_disk_name" {
  description = "The specification of the system disk name."
  type        = string
  default     = "system_disk"
}

variable "system_disk_description" {
  description = "The specification of the system disk description."
  type        = string
  default     = "system_disk_description"
}

variable "image_id" {
  description = "The specification of the image id."
  type        = string
  default     = "ubuntu_18_04_64_20G_alibase_20190624.vhd"
}

variable "internet_max_bandwidth_out" {
  description = "The specification of the internet max bandwidth out."
  type        = number
  default     = 10
}

variable "ecs_size" {
  description = "The specification of the ecs size."
  type        = number
  default     = 1200
}

variable "category" {
  description = "The specification of the category."
  type        = string
  default     = "cloud_efficiency"
}

variable "engine" {
  description = "The specification of the engine."
  type        = string
  default     = "MySQL"
}

variable "engine_version" {
  description = "The specification of the engine version."
  type        = string
  default     = "5.6"
}

variable "rds_instance_type" {
  description = "The specification of the rds instance type."
  type        = string
  default     = "rds.mysql.s2.large"
}

variable "instance_storage" {
  description = "The specification of the instance storage."
  type        = string
  default     = "30"
}

variable "instance_charge_type" {
  description = "The specification of the instance charge type."
  type        = string
  default     = "Postpaid"
}

variable "monitoring_period" {
  description = "The specification of the monitoring period."
  type        = string
  default     = "60"
}

variable "security_ips" {
  description = "The specification of the security ips."
  type        = list(string)
  default     = ["127.0.0.1"]
}

variable "redis_instance_type" {
  description = "The specification of the redis instance type."
  type        = string
  default     = "Redis"
}

variable "redis_engine_version" {
  description = "The specification of the redis engine version."
  type        = string
  default     = "4.0"
}

variable "redis_appendonly" {
  description = "The specification of the redis appendonly."
  type        = string
  default     = "yes"
}

variable "redis_lazyfree_lazy_eviction" {
  description = "The specification of the redis lazyfree-lazy-eviction."
  type        = string
  default     = "yes"
}

variable "redis_resource_group_id" {
  description = "The specification of the redis resource group id."
  type        = string
  default     = "rg-123456"
}

variable "redis_instance_class" {
  description = "The specification of the redis resource instance class."
  type        = string
  default     = "redis.master.large.default"
}
