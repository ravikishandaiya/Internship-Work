terraform {
  required_providers {
    zoom = {
      version = "0.1.0"
      source  = "hashicorp.com/edu/zoom"
    }
  }
}

provider "zoom" {
  token = "JWT_Token"
}

resource "zoom_user" "example" { 
  email = "ravikishandaiya@gmail.com"
  firstname = "Ravi"
  lastname = "Kishan"
}
