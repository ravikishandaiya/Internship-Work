resource "miro_user" "user7" {
    email   = "geetaa@clevertap.com"
    role    = "admin"
}

output "user7" {
    value = miro_user.user7
}