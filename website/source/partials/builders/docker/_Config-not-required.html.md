<!-- Code generated from the comments of the Config struct in builder/docker/config.go; DO NOT EDIT MANUALLY -->

-   `container_dir` (string) - The directory inside container to mount temp
directory from host server for work file
provisioner. This defaults to
c:/packer-files on windows and /packer-files on other systems.

-   `exec_user` (string) - Username (UID) to run remote commands with. You can
also set the group name/ID if you want: (UID or UID:GID).
You may need this if you get permission errors trying to run the shell or
other provisioners.

-   `export_path` (string) - Export Path
-   `privileged` (bool) - If true, run the docker container with the
--privileged flag. This defaults to false if not set.

-   `run_command` ([]string) - An array of arguments to pass to
docker run in order to run the container. By default this is set to
["-d", "-i", "-t", "--entrypoint=/bin/sh", "--", "{{.Image}}"] if you are
using a linux container, and
["-d", "-i", "-t", "--entrypoint=powershell", "--", "{{.Image}}"] if you
are running a windows container. {{.Image}} is a template variable that
corresponds to the image template option. Passing the entrypoint option
this way will make it the default entrypoint of the resulting image, so
running docker run -it --rm  will start the docker image from the
/bin/sh shell interpreter; you could run a script or another shell by
running docker run -it --rm  -c /bin/bash. If your docker image
embeds a binary intended to be run often, you should consider changing the
default entrypoint to point to it.

-   `fix_upload_owner` (bool) - If true, files uploaded to the container
will be owned by the user the container is running as. If false, the owner
will depend on the version of docker installed in the system. Defaults to
true.

-   `windows_container` (bool) - If "true", tells Packer that you are building a
Windows container running on a windows host. This is necessary for building
Windows containers, because our normal docker bindings do not work for them.

-   `login_password` (string) - The password to use to authenticate to login.

-   `login_server` (string) - The server address to login to.

-   `login_username` (string) - The username to use to authenticate to login.

-   `ecr_login` (bool) - Defaults to false. If true, the builder will login
in order to pull the image from Amazon EC2 Container Registry
(ECR). The builder only logs in for the
duration of the pull. If true login_server is required and login,
login_username, and login_password will be ignored. For more
information see the section on ECR.
