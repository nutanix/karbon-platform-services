# Deploying the Service Domain image on AWS

## About the Service Domain Image File

The Service Domain bare metal image is available as a raw disk image provided by Nutanix for installation on approved hardware.

Download the Service Domain VM image file from the Nutanix Support portal [Downloads](https://portal.nutanix.com/page/downloads?product=karbonplatformservices) page.

There is a raw image file available for download:

  * sherlock-aws_<version>.raw

## Deploying/Onboarding AWS Service Domain

1. Install and configure aws CLI.

	* Note: In the steps below, a new trust policy will be applied to an S3 bucket storing the Service Domain image. It is recommended to create a new bucket to avoid the potential for policy error or unintended configuration on buckets used for other purposes. 
2. Create an S3 bucket to import the Service Domain raw image:
```
aws s3 mb s3://raw-image-bkt
```
3. Copy the raw image to the S3 bucket
```
aws s3 cp sherlock-aws_<verson>.raw s3://raw-image-test
```
4. Create a new trust policy configuration file called _trust-policy.json_:
```yaml
{
   "Version": "2012-10-17",
   "Statement": [
      {
         "Effect": "Allow",
         "Principal": { "Service": "vmie.amazonaws.com" },
         "Action": "sts:AssumeRole",
         "Condition": {
            "StringEquals":{
               "sts:Externalid": "vmimport"
            }
         }
      }
   ]
}
```
5. Create a role named vmimport and grant VM Import/Export access to it.
```
aws iam create-role --role-name vmimport --assume-role-policy-document file://trust-policy.json
```
6. Create a new role policy configuration file called role-policy.json like below, changing <S3 bucket> to the name of S3 bucket created in step 2.
```yaml
{
   "Version":"2012-10-17",
   "Statement":[
      {
         "Effect": "Allow",
         "Action": [
            "s3:GetBucketLocation",
            "s3:GetObject",
            "s3:ListBucket" 
         ],
         "Resource": [
            "arn:aws:s3:::<S3 bucket>",
            "arn:aws:s3:::<S3 bucket>/*"
         ]
      },
      {
         "Effect": "Allow",
         "Action": [
            "s3:GetBucketLocation",
            "s3:GetObject",
            "s3:ListBucket",
            "s3:PutObject",
            "s3:GetBucketAcl"
         ],
         "Resource": [
            "arn:aws:s3:::<S3 bucket>",
            "arn:aws:s3:::<S3 bucket>/*"
         ]
      },
      {
         "Effect": "Allow",
         "Action": [
            "ec2:ModifySnapshotAttribute",
            "ec2:CopySnapshot",
            "ec2:RegisterImage",
            "ec2:Describe*"
         ],
         "Resource": "*"
      }
   ]
}
```

7. Attach the role policy to the vmimport role
```
aws iam put-role-policy --role-name vmimport --policy-name vmimport --policy-document file://role-policy.json
```

8. Create a file called _container.json_:
```yaml
{
     "Description": "Karbon Platform Services Raw Image",
     "Format": "RAW",
     "UserBucket": {
        "S3Bucket": "<S3 bucket>",
        "S3Key": "sherlock-aws_<version>.raw"
    }
 }

```
9. Run the following command to import the snapshot as an AMI
`aws ec2 import-snapshot --description "example" --disk-container "file://container.json"`

	* **Note**: you can query your snapshot task progress with the following command
`aws ec2 describe-import-snapshot-tasks --import-task-ids <TaskId>`

10. Once completed your snapshot will be viewable in your EC2 console.
11. Select the snapshot and choose Create Image. Fill out as follows:
	* **Architecture**: x86_64
	* **Root device name**: /dev/xvdf
	* Everything else can be set to default
12. Navigate back to the EC2 console and under **Images** select **AMIs**
13. Select your recently created AMI and hit **Launch**.
	* For the Instance Type select an instance with a minimum of 4 vCPU and 16GiB of memory.
	* You can skip to **Add Storage**. Create an EBS Volume with a minimum size of 100 GIB and proceed to **Review**.
	* Before launching, use an existing key pair for which you already possess the private key, or select **Create** a new key pair and download the .pem file corresponding to the new key pair.
14. Run the following command to get the Private IP address of your instance.
```
aws ec2 describe-instances --instance-id <instance-id> --query 'Reservations[].Instances[].[PrivateIpAddress]' --output text | sed '$!N;s/\n/ /'
```
15. Select **Connect** in your EC2 console and follow the instructions to SSH into your instance.
16. Once on your EC2 instance run the following commands to grab your Serial Number and Gateway/Subnet:
	* Serial Number: ```cat /config/serial_number.txt```
	* Gateway/Subnet: ```route -n```
17. Finally navigate to KPS and onboard your edge with the values you just received.
