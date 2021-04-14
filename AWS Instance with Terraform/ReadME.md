# terraform-aws-instance

* Resource: https://learn.hashicorp.com/tutorials/terraform/aws-build?in=terraform/aws-get-started


### #Prerequisites:
* The Terraform CLI installed: 
  https://learn.hashicorp.com/tutorials/terraform/install-cli?in=terraform/aws-get-started

* The AWS CLI installed: 
  https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html

* An AWS account: 
  https://aws.amazon.com/free/?all-free-tier.sort-by=item.additionalFields.SortRank&all-free-tier.sort-order=asc

(If any isn't working after installation ensure that your PATH variable contains the directory where the application is installed.)

### let's start AWS instance with tf (terraform)

## Step 1
create a seperate directory for this.

mkdir learn-terraform-aws-instance
cd learn-terraform-aws-instance

## Step 2
#### Method 1:

Either you can run "aws configure" or can put your credentials in the tf file.

So if you run "aws configure"
It will ask for these 4 things.
AWS Access Key ID 
AWS Secret Access Key
get above credential at aws account.
Default region name
Default output format
these two can be skipped, latter we will initialize them in tf file.

And so you already have put your credentials via running "aws configure"
if so then create a main.tf file as following


```
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }
}
provider "aws" {
  profile = "default"
  region  = "us-west-2"
}
resource "aws_instance" "app_server" {
  ami           = "ami-830c94e3"
  instance_type = "t2.micro"
tags = {
    Name = "ExampleAppServerInstance"
  }
}
```


#### Method 2:

Make as tf config file same as main.tf from this repository if you are wishing to initialize credential in the tf file and make sure you are putting credential in the file (access_key & secret key).


## Step 3:

run terraform commands to execute this tf config file
```
$ terraform init
$ terraform plan
$ terraform apply

```

Hope everything worked right.

To stop/ realse resources run

```
$ terraform destroy

```



# Most Important

## Make sure your ami belong to the same region which you had put in your config file.
