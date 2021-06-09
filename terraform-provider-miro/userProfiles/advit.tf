resource "miro_user" "user3" {
    email   = "advit@clevertap.com"
    role    = "member"
}

output "user3" {
    value = miro_user.user3
}