terraform {
  required_providers {
    hashicups = {
      version = "0.2"
      source  = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "hashicups" {}

module "psl" {
  source = "./coffee"

  coffee_name = "Terraspresso"
}

output "psl" {
  value = module.psl.coffee
}
