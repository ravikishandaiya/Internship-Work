resource "miro_user" "user4" {
    email   = "swapnil@clevertap.com"
    role    = "member"
}

output "use4" {
    value = miro_user.user4
}