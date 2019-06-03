<!-- Code generated from the comments of the Config struct in builder/azure/arm/config.go; DO NOT EDIT MANUALLY -->

-   `capture_name_prefix` (string) - Capture

-   `capture_container_name` (string) - Capture Container Name
-   `shared_image_gallery` (SharedImageGallery) - Use a Shared Gallery
image
as the source for this build. VHD targets are incompatible with this build
type - the target must be a Managed Image.

-   `image_version` (string) - Specify a specific version of an OS to boot from.
Defaults to latest. There may be a difference in versions available
across regions due to image synchronization latency. To ensure a consistent
version across regions set this value to one that is available in all
regions where you are deploying.

-   `image_url` (string) - Specify a custom VHD to use. If this value is set, do
not set image_publisher, image_offer, image_sku, or image_version.

-   `custom_managed_image_resource_group_name` (string) - Specify the source
managed image's resource group used to use. If this value is set, do not
set image_publisher, image_offer, image_sku, or image_version. If this
value is set, the value custom_managed_image_name must also be set. See
documentation
to learn more about managed images.

-   `custom_managed_image_name` (string) - Specify the source managed image's
name to use. If this value is set, do not set image_publisher,
image_offer, image_sku, or image_version. If this value is set, the
value custom_managed_image_resource_group_name must also be set. See
documentation
to learn more about managed images.

-   `location` (string) - Location
-   `vm_size` (string) - Size of the VM used for building. This can be changed
when you deploy a VM from your VHD. See
pricing
information. Defaults to Standard_A1.

-   `managed_image_resource_group_name` (string) - Managed Image Resource Group Name
-   `managed_image_name` (string) - Managed Image Name
-   `managed_image_storage_account_type` (string) - Specify the storage account
type for a managed image. Valid values are Standard_LRS and Premium_LRS.
The default is Standard_LRS.

-   `managed_image_os_disk_snapshot_name` (string) - If
managed_image_os_disk_snapshot_name is set, a snapshot of the OS disk
is created with the same name as this value before the VM is captured.

-   `managed_image_data_disk_snapshot_prefix` (string) - If
managed_image_data_disk_snapshot_prefix is set, snapshot of the data
disk(s) is created with the same prefix as this value before the VM is
captured.

-   `managed_image_zone_resilient` (bool) - Store the image in zone-resilient storage. You need to create it
in a region that supports availability zones.

-   `azure_tags` (map[string]*string) - the user can define up to 15
tags. Tag names cannot exceed 512 characters, and tag values cannot exceed
256 characters. Tags are applied to every resource deployed by a Packer
build, i.e. Resource Group, VM, NIC, VNET, Public IP, KeyVault, etc.

-   `resource_group_name` (string) - Resource Group Name
-   `storage_account` (string) - Storage Account
-   `temp_compute_name` (string) - temporary name assigned to the VM. If this
value is not set, a random value will be assigned. Knowing the resource
group and VM name allows one to execute commands to update the VM during a
Packer build, e.g. attach a resource disk to the VM.

-   `temp_resource_group_name` (string) - Temp Resource Group Name
-   `build_resource_group_name` (string) - Build Resource Group Name
-   `private_virtual_network_with_public_ip` (bool) - This value allows you to
set a virtual_network_name and obtain a public IP. If this value is not
set and virtual_network_name is defined Packer is only allowed to be
executed from a host on the same subnet / virtual network.

-   `virtual_network_name` (string) - Use a pre-existing virtual network for the
VM. This option enables private communication with the VM, no public IP
address is used or provisioned (unless you set
private_virtual_network_with_public_ip).

-   `virtual_network_subnet_name` (string) - If virtual_network_name is set,
this value may also be set. If virtual_network_name is set, and this
value is not set the builder attempts to determine the subnet to use with
the virtual network. If the subnet cannot be found, or it cannot be
disambiguated, this value should be set.

-   `virtual_network_resource_group_name` (string) - If virtual_network_name is
set, this value may also be set. If virtual_network_name is set, and
this value is not set the builder attempts to determine the resource group
containing the virtual network. If the resource group cannot be found, or
it cannot be disambiguated, this value should be set.

-   `custom_data_file` (string) - Specify a file containing custom data to inject
into the cloud-init process. The contents of the file are read and injected
into the ARM template. The custom data will be passed to cloud-init for
processing at the time of provisioning. See
documentation
to learn more about custom data, and how it can be used to influence the
provisioning process.

-   `plan_info` (PlanInformation) - Used for creating images from Marketplace images.
Please refer to Deploy an image with Marketplace
terms for more details. Not
all Marketplace images support programmatic deployment, and support is
controlled by the image publisher.

-   `os_type` (string) - If either Linux or Windows is specified Packer will
automatically configure authentication credentials for the provisioned
machine. For Linux this configures an SSH authorized key. For Windows
this configures a WinRM certificate.

-   `os_disk_size_gb` (int32) - Specify the size of the OS disk in GB
(gigabytes). Values of zero or less than zero are ignored.

-   `disk_additional_size` ([]int32) - The size(s) of any additional
hard disks for the VM in gigabytes. If this is not specified then the VM
will only contain an OS disk. The number of additional disks and maximum
size of a disk depends on the configuration of your VM. See
Windows
or
Linux
for more information.

-   `disk_caching_type` (string) - Specify the disk caching type. Valid values
are None, ReadOnly, and ReadWrite. The default value is ReadWrite.

-   `async_resourcegroup_delete` (bool) - If you want packer to delete the
temporary resource group asynchronously set this value. It's a boolean
value and defaults to false. Important Setting this true means that
your builds are faster, however any failed deletes are not reported.
