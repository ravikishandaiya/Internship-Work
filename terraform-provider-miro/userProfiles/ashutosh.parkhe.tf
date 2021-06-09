resource "miro_user" "user1" {
    email   = "ashutosh.parkhe@clevertap.com"
    role    = "member"
}

output "user1" {
    value = miro_user.user1
}