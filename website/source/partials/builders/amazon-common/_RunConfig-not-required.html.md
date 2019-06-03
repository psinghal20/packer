<!-- Code generated from the comments of the RunConfig struct in builder/amazon/common/run_config.go; DO NOT EDIT MANUALLY -->

-   `associate_public_ip_address` (bool) - If using a non-default VPC,
public IP addresses are not provided by default. If this is true, your
new instance will get a Public IP. default: false

-   `availability_zone` (string) - Destination availability zone to launch
instance in. Leave this empty to allow Amazon to auto-assign.

-   `block_duration_minutes` (int64) - Requires spot_price to be set. The
required duration for the Spot Instances (also known as Spot blocks). This
value must be a multiple of 60 (60, 120, 180, 240, 300, or 360). You can't
specify an Availability Zone group or a launch group if you specify a
duration.

-   `disable_stop_instance` (bool) - Packer normally stops the build instance after all provisioners have
run. For Windows instances, it is sometimes desirable to [run
Sysprep](http://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ami-create-standard.html)
which will stop the instance for you. If this is set to `true`, Packer
*will not* stop the instance but will assume that you will send the stop
signal yourself through your final provisioner. You can do this with a
[windows-shell
provisioner](https://www.packer.io/docs/provisioners/windows-shell.html).
Note that Packer will still wait for the instance to be stopped, and
failing to send the stop signal yourself, when you have set this flag to
`true`, will cause a timeout.
Example of a valid shutdown command:

``` json
{
  "type": "windows-shell",
  "inline": ["\"c:\\Program Files\\Amazon\\Ec2ConfigService\\ec2config.exe\" -sysprep"]
}
```

-   `ebs_optimized` (bool) - Mark instance as [EBS
Optimized](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html).
Default `false`.

-   `enable_t2_unlimited` (bool) - Enabling T2 Unlimited allows the source
instance to burst additional CPU beyond its available CPU
Credits
for as long as the demand exists. This is in contrast to the standard
configuration that only allows an instance to consume up to its available
CPU Credits. See the AWS documentation for T2
Unlimited
and the T2 Unlimited Pricing section of the Amazon EC2 On-Demand
Pricing document for more
information. By default this option is disabled and Packer will set up a
T2
Standard
instance instead.

-   `iam_instance_profile` (string) - The name of an IAM instance
profile
to launch the EC2 instance with.

-   `shutdown_behavior` (string) - Automatically terminate instances on
shutdown in case Packer exits ungracefully. Possible values are stop and
terminate. Defaults to stop.

-   `security_group_filter` (SecurityGroupFilterOptions) - Filters used to populate the
`security_group_ids` field. Example:

``` json
{
  "security_group_filter": {
    "filters": {
      "tag:Class": "packer"
    }
  }
}
```

-   `run_tags` (map[string]string) - Tags to apply to the instance
that is launched to create the AMI. These tags are not applied to the
resulting AMI unless they're duplicated in tags. This is a template
engine, see Build template
data for more information.

-   `security_group_id` (string) - The ID (not the name) of the security
group to assign to the instance. By default this is not set and Packer will
automatically create a new temporary security group to allow SSH access.
Note that if this is specified, you must be sure the security group allows
access to the ssh_port given below.

-   `security_group_ids` ([]string) - A list of security groups as
described above. Note that if this is specified, you must omit the
security_group_id.

-   `source_ami_filter` (AmiFilterOptions) - Filters used to populate the source_ami
field. Example:

-   `spot_instance_types` ([]string) - a list of acceptable instance
types to run your build on. We will request a spot instance using the max
price of spot_price and the allocation strategy of "lowest price".
Your instance will be launched on an instance type of the lowest available
price that you have in your list.  This is used in place of instance_type.
You may only set either spot_instance_types or instance_type, not both.
This feature exists to help prevent situations where a Packer build fails
because a particular availability zone does not have capacity for the
specific instance_type requested in instance_type.

-   `spot_price` (string) - The maximum hourly price to pay for a spot instance
to create the AMI. Spot instances are a type of instance that EC2 starts
when the current spot price is less than the maximum price you specify.
Spot price will be updated based on available spot instance capacity and
current spot instance requests. It may save you some costs. You can set
this to auto for Packer to automatically discover the best spot price or
to "0" to use an on demand instance (default).

-   `spot_price_auto_product` (string) - Required if spot_price is set to
auto. This tells Packer what sort of AMI you're launching to find the
best spot price. This must be one of: Linux/UNIX, SUSE Linux,
Windows, Linux/UNIX (Amazon VPC), SUSE Linux (Amazon VPC),
Windows (Amazon VPC)

-   `spot_tags` (map[string]string) - Requires spot_price to be
set. This tells Packer to apply tags to the spot request that is issued.

-   `subnet_filter` (SubnetFilterOptions) - Filters used to populate the subnet_id field.
Example:

-   `subnet_id` (string) - If using VPC, the ID of the subnet, such as
subnet-12345def, where Packer will launch the EC2 instance. This field is
required if you are using an non-default VPC.

-   `temporary_key_pair_name` (string) - The name of the temporary key pair to
generate. By default, Packer generates a name that looks like
packer_<UUID>, where <UUID> is a 36 character unique identifier.

-   `temporary_security_group_source_cidrs` ([]string) - A list of IPv4
CIDR blocks to be authorized access to the instance, when packer is creating a temporary security group.

-   `user_data` (string) - User data to apply when launching the instance. Note
that you need to be careful about escaping characters due to the templates
being JSON. It is often more convenient to use user_data_file, instead.
Packer will not automatically wait for a user script to finish before
shutting down the instance this must be handled in a provisioner.

-   `user_data_file` (string) - Path to a file that will be used for the user
data when launching the instance.

-   `vpc_filter` (VpcFilterOptions) - Filters used to populate the vpc_id field.
vpc_id take precedence over this.
Example:

-   `vpc_id` (string) - If launching into a VPC subnet, Packer needs the VPC ID
in order to create a temporary security group within the VPC. Requires
subnet_id to be set. If this field is left blank, Packer will try to get
the VPC ID from the subnet_id.

-   `windows_password_timeout` (time.Duration) - The timeout for waiting for a Windows
password for Windows instances. Defaults to 20 minutes. Example value:
10m
