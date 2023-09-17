terraform {
    required_providers {
      chucknorris = {
          source = "example.com/hharvey/chucknorris-provider"
          version = "~> 0.0.1"
        }
    }
  }

provider "chucknorris" {}

data "chucknorris_joke" "pillow" {
    joke_id = "m5prryVlQgWTx7mC2L8Rrw"
}

data "chucknorris_joke" "carpet" {
    joke_id = "XBEz9TVCSVe9e7J14vDhBw"
}


output "pillow" {
    value = data.chucknorris_joke.pillow
}

output "carpet" {
    value = data.chucknorris_joke.carpet
}
