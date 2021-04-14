provider "aws" {
    access_key = "ACCESS KEY"
    secret_key = "SECRET KEY"
    region = "us-east-2"
}

resource "aws_instance" "example" {
    ami = "ami-08962a4068733a2b6"
    instance_type = "t2.micro"
}