### Steps to build Provider

OS: Windows

#### Step 1.
```go mod init``` command makes current directory the root of the module.
run 
```$ go mod init terraform-provider-zoom```

#### Step 2.
Then, run ```go mod vendor``` to create a vendor directory that contains all the provider's dependencies.
run
```$ go mod vendor```

#### Step 3.
Then, run ```go build``` to get binary provider file.
```$ go build -o terraform-provider-zoom.exe```

#### Step 4.
Make direcory in following manner.
```C:\Users\ravik\AppData\Roaming\terraform.d\plugins\hashicorp.com\edu\[provider_name]\[version]\[OS_arch]```

In my case
```$ mkdir C:\Users\ravik\AppData\Roaming\terraform.d\plugins\hashicorp.com\edu\zoom\0.1.0\windows_amd64```

And put this (```terraform-provider-zoom.exe```) binary file in this folder.

#### 5.
Put respective data in the ```main.tf``` file and 
run 
```$ terraform init```
```$ terraform plan```
```$ terraform apply```
By these command respective operation will be done. 

Run
```$ terraform destroy```
It frees resorces and does undo apply. 

HaPpY LeaRninG.
