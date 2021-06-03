resource "miro_user" "user5" {
    email   = "kapil@clevertap.com"
    role    = "member"
}

output "user5" {
    value = miro_user.user5
}