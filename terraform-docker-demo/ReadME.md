# NGINX server using docker & Terraform

* Assuming terraform is installed &
* Working on windows


Download & Install Docker Desktop: https://docs.docker.com/docker-for-windows/install/

Verify the installation: run "docker -help" if isn't working then ensure that your PATH variable contains the directory where Docker was installed. In my case these directories are "C:\Program Files\Docker\Docker\resources\bin"  & "C:\ProgramData\DockerDesktop\version-bin"

Let's provision an NGINX server using docker and Terraform.

## Step 1: 
Run Docker Desktop (GUI desktop app or can search in start)

Let's create a differnt directory so that terraform config files doesn't conflict. 
run
    mkdir terraform-docker-demo
    cd terraform-docker-demo

create a main.tf file same as in respective directory at this Github.

## Step 2:
Run following commands:
    terraform init
    terraform plan
    terraform apply
    
And you are done.

If you haven't made changes (means port no) in the tf file then you can check it at
"http://localhost:8000/"

You can check this IMAGE STATUS on docker desktop or you can run "docker ps" for all the created Images.

## Step 3:
By "terraform destroy" you can stop this instance. (make sure you run this command in that same directory.)


