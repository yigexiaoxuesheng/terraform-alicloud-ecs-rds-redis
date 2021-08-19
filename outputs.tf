output "this_vpc_id" {
  value = alicloud_vpc.default.id
}

output "this_vswitch_id" {
  value = alicloud_vswitch.default.id
}

output "this_security_group_id" {
  value = alicloud_security_group.default.id
}

output "this_ecs_id" {
  value = alicloud_instance.default.id
}

output "this_rds_id" {
  value = alicloud_db_instance.default.id
}

output "this_redis_id" {
  value = alicloud_kvstore_instance.default.id
}

output "this_zone_id" {
  value = data.alicloud_zones.default.zones[0].id
}
