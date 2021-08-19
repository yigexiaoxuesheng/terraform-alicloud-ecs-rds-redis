module "ecs-rds" {
  source              = "../"
  name                = "tf-ecs-rds"
  instance_type       = "ecs.n4.large"
  rds_instance_type   = "rds.mysql.s2.large"
}
