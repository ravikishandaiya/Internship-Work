resource "miro_user" "user8" {
    email   = "saurabh.malkar@clevertap.com"
    role    = "member"
}

output "user8" {
    value = miro_user.user8
}