# findami



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

```
ami-0b6355270fae3a72e   Linux/UNIX   x86_64   gp3   al2022-ami-2022.0.20211118.0-kernel-5.10-x86_64
ami-0e357443256755720   Linux/UNIX   x86_64   gp3   al2022-ami-minimal-2022.0.20211118.0-kernel-5.10-x86_64
```


how to set [filters](https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-images.html)
