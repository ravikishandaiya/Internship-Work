provider "zoom" {
    address = "https://api.zoom.us/v2/users/"
    token = "api_token"
}

resource "provider_user" "user1" {
    first_name = "Ravi"
    last_name = "Daiya"
    email = "abc@gmail.com"
}

/*
resource "provider_user" "user2" {
    first_name = "A"
    last_name = "B"
    email = "ravi@gmail.com"
}
*/
