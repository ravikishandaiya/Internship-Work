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
    token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTk3NzQ4NjAsImlhdCI6MTYxOTE3MDA1Nn0.EgucHFpdFPjoHKMc4dSepij0NohuMGBnFUSxdcbZc3w"
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