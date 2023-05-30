terraform {
    required_providers {
      swapi = {
          source = "example.com/hharvey/swapi-provider"
          version = "~> 0.0.1"
        }
    }
  }

provider "swapi" {}

data "swapi_people" "data" {}


output "WHAT" {
    value = data.swapi_people.data
  }
