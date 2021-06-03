resource "miro_user" "user2" {
    email   = "mihir@clevertap.com"
    role    = "member"
}

output "user2" {
    value = miro_user.user2
}