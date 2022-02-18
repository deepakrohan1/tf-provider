terraform {
  required_providers {
    roles = {
        version = "0.3"
        source = "deepakrohan.com/edu/roles"
    }
  }
}

data "roles_permissions" "all" {}


output "all_permissions" {
  value = "data.roles_permissions.all.permissions"
}