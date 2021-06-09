resource "miro_user" "user6" {
    email   = "vaidehi@clevertap.com"
    role    = "member"
}

output "user6" {
    value = miro_user.user6
}