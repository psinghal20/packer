<!-- Code generated from the comments of the AMIBlockDevices struct in builder/amazon/common/block_device.go; DO NOT EDIT MANUALLY -->

-   `ami_block_device_mappings` ([]BlockDevice) - Add one or
more block device
mappings
to the AMI. These will be attached when booting a new instance from your
AMI. If this field is populated, and you are building from an existing source image,
the block device mappings in the source image will be overwritten. This means you
must have a block device mapping entry for your root volume, root_volume_size,
and root_device_name. `Your options here may vary depending on the type of VM
you use. The block device mappings allow for the following configuration:
