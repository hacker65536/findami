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


	- architecture : i386 | x86_64 | arm64
	- block-device-mapping.delete-on-termination: true | false
	- block-device-mapping.device-name: "/dev/sdh" | "xvdh"
	- block-device-mapping.snapshot-id: snap-xxxxxxxxxxxxxxxxx
	- block-device-mapping.volume-size: 8
	- block-device-mapping.volume-type: io1 | io2 | gp2 | gp3 | sc1 | st1 | standard
	- block-device-mapping.encrypted: true | false
	- description: "The description of the image"
	- ena-support: true | false
	- hypervisor: ovm | xen
	- image-id: ami-xxxxxx
	- image-type: machine | kernel | ramdisk
	- is-public: true | false
	- kernel-id: xxxxx 
	- manifest-location: manifest 
	- name: name
	- owner-alias: amazon | aws-marketplace
	- owner-id: xxxxx
	- platform: only windows 
	- product-code: code 
	- product-code.type: marketplace 
	- ramdisk-id: The RAM disk ID
	- root-device-name: /dev/sda1
	- root-device-type: ebs | instance-store
	- state: available | pending | failed
	- state-reason-code: xxx 
	- state-reason-message: reason
	- sriov-net-support: simple 
	- tag:  key/value
	- tag-key: key 
	- virtualization-type: paravirtual | hvm
