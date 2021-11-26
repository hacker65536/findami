# findami

search ami with config file.

search ami for 'Amazon Linux 2022'

`.findami.yml`
```yaml
AMIFilters:
  - Name: architecture
    Values: x86_64
  - Name: block-device-mapping.volume-type
    Values: gp3
  - Name: owner-alias
    Values: amazon
  - Name: is-public
    Values: "true"
  - Name: virtualization-type
    Values: hvm
  - Name: root-device-type
    Values: ebs
  - Name: sriov-net-support
    Values: simple
  - Name: ena-support
    Values: "true"
  - Name: state
    Values: available
  - Name: name
    Values: "*al2022*"
```

```console
$ findami 
ami-0b6355270fae3a72e   Linux/UNIX   x86_64   gp3   al2022-ami-2022.0.20211118.0-kernel-5.10-x86_64
ami-0e357443256755720   Linux/UNIX   x86_64   gp3   al2022-ami-minimal-2022.0.20211118.0-kernel-5.10-x86_64
```

how to set [filters](https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-images.html)

## install

```
go install github.com/hacker65536/findami@latest
```


This way same as `aws ec2 describe-images --image-id ami-0b6355270fae3a72e`
```console
$ findami -v ami-0b6355270fae3a72e | jq .
{
  "Images": [
    {
      "Architecture": "x86_64",
      "BlockDeviceMappings": [
        {
          "DeviceName": "/dev/xvda",
          "Ebs": {
            "DeleteOnTermination": true,
            "Encrypted": false,
            "Iops": null,
            "KmsKeyId": null,
            "OutpostArn": null,
            "SnapshotId": "snap-0f4d9678f534201ea",
            "Throughput": 125,
            "VolumeSize": 8,
            "VolumeType": "gp3"
          },
          "NoDevice": null,
          "VirtualName": null
        }
      ],
      "BootMode": "",
      "CreationDate": "2021-11-19T01:49:55.000Z",
      "DeprecationTime": null,
      "Description": "Amazon Linux 2022 AMI 2022.0.20211118.0 x86_64 HVM kernel-5.10",
      "EnaSupport": true,
      "Hypervisor": "xen",
      "ImageId": "ami-0b6355270fae3a72e",
      "ImageLocation": "amazon/al2022-ami-2022.0.20211118.0-kernel-5.10-x86_64",
      "ImageOwnerAlias": "amazon",
      "ImageType": "machine",
      "KernelId": null,
      "Name": "al2022-ami-2022.0.20211118.0-kernel-5.10-x86_64",
      "OwnerId": "137112412989",
      "Platform": "",
      "PlatformDetails": "Linux/UNIX",
      "ProductCodes": null,
      "Public": true,
      "RamdiskId": null,
      "RootDeviceName": "/dev/xvda",
      "RootDeviceType": "ebs",
      "SriovNetSupport": "simple",
      "State": "available",
      "StateReason": null,
      "Tags": null,
      "UsageOperation": "RunInstances",
      "VirtualizationType": "hvm"
    }
  ],
  "ResultMetadata": {}
}
```


### ubuntu

https://ubuntu.com/server/docs/cloud-images/amazon-ec2

`owner-id: 679593333241`

### centos

https://wiki.centos.org/Cloud/AWS

centos8 `product-code: 47k9ia2igxpcce2bzo8u3kj03`  
centos7 `product-code: cvugziknvmxgqna9noibqnnsy`  
centos6 `product-code: ckx0h8ljio731afm2k92jtg62`  
