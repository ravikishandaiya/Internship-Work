resource "miro_user" "user9" {
    email   = "aayushi@clevertap.com"
    role    = "member"
}

output "user9" {
    value = miro_user.user9
}