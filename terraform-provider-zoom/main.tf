terraform {
  required_providers{
      zoom = {
          version ="0.1.0"
          source = "hashicorp.com/edu/zoom"
      }
  }
}

provider "zoom" {
    //address = "https://api.zoom.us/v2/users/"
    token = "token"
  
  }

resource "zoom_user" "user2" {
    //firstname = "Ravi"
    //lastname = "Daiya"
    email ="ravikishandaiya@gmail.com"  
}

/*
resource "provider_user" "user2" {
    first_name = "A"
    last_name = "B"
    email ="abc@gmail.com"  
} 
*/
