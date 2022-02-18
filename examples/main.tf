terraform {
  required_providers {
    roles = {
        version = "0.3"
        source = "deepakrohan.com/edu/roles"
    }
  }
}


provider "roles" {}

module "psl" {
    source = "./permission"

}

output "psl" {
    value = "module.psl.permission"
}