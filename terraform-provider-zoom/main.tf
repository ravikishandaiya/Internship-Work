terraform {
  required_providers {
    zoom = {
      version = "0.1.0"
      source  = "hashicorp.com/edu/zoom"
    }
  }
}

provider "zoom" {
  //token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTk2NzY3OTQsImlhdCI6MTYxOTA3MTk5MX0.Hj_3kz19mxRBOf9nPsknekW2I5kRTBfzXpngJAFW9rI"
}

resource "zoom_user" "example" { 
  email = "ravikishandaiya@gmail.com"
  firstname = "Ravi"
  lastname = "Kishan"
}
