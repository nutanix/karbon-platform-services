# Introducing the kpsctl Command Line Interface (CLI)

**Note**: Nutanix is actively developing the kpsctl CLI. As a result, the software published here is release candidate (RC) quality until further notice. 
Use this software at your own risk and do not use it in a production environment or deployment. The latest release candidate is shown below.

The kpsctl CLI is a command line tool that helps you manage your Karbon Platform Services for IoT resources from your terminal or shell. With minimal 
configuration, you can start using Karbon Platform Services for IoT functionality equal to that provided by the browser-based Karbon Platform Services for IoT Management 
Console from the command prompt in your favorite terminal program.

## Shell Support
**Linux and macOS** You can use common shell programs such as bash and zsh.
Note that autocompletion for zsh is experimental in this release.

# Installing kpsctl CLI
## MacOS
Download the kpsctl CLI for MacOS:

**64-bit**
[kpsctl-v1.0.0-rc5-darwin_amd64.tar.gz](https://kpsctl-cli.s3-us-west-2.amazonaws.com/kpsctl-v1.0.0-rc5-darwin_amd64.tar.gz)

**32-bit**
[kpsctl-v1.0.0-rc5-darwin_386.tar.gz](https://kpsctl-cli.s3-us-west-2.amazonaws.com/kpsctl-v1.0.0-rc5-darwin_386.tar.gz)

Optionally, from your terminal or shell, use `wget`.
```
$ wget https://kpsctl-cli.s3-us-west-2.amazonaws.com/kpsctl-v1.0.0-rc5-darwin_amd64.tar.gz
$ mkdir kpsctl && tar zxvf kpsctl-v1.0.0-rc5-darwin_amd64.tar.gz -C kpsctl
$ sudo kpsctl/install [*optional_path*]
```

**Note** The kpsctl binary default installation directory is `/usr/local/bin/`. 
To install the kpsctl binary in a custom location, specify the custom directory or path as follows, where
my/custom/dir is the installation path:
```
$ sudo kpsctl/install /my/custom/dir
```

## Linux
Download the kpsctl CLI for Linux:

**64-bit**
[kpsctl-v1.0.0-rc5-linux_amd64.tar.gz](https://kpsctl-cli.s3-us-west-2.amazonaws.com/kpsctl-v1.0.0-rc5-linux_amd64.tar.gz)

**32-bit**
[kpsctl-v1.0.0-rc5-linux_386.tar.gz](https://kpsctl-cli.s3-us-west-2.amazonaws.com/kpsctl-v1.0.0-rc5-linux_386.tar.gz)

Optionally, from your terminal or shell, use `wget`.
```
$ wget https://kpsctl-cli.s3-us-west-2.amazonaws.com/kpsctl-v1.0.0-rc5-linux_amd64.tar.gz
$ mkdir kpsctl && tar zxvf kpsctl-v1.0.0-rc5-linux_amd64.tar.gz -C kpsctl
$ sudo kpsctl/install [*optional_path*]
```

**Note** The kpsctl binary default installation directory is `/usr/local/bin/`. 
To install the kpsctl binary in a custom location, specify the custom directory or path as follows, where
my/custom/dir is the installation path:
```
$ sudo kpsctl/install /my/custom/dir
```

## Windows

Download the kpsctl CLI for Microsoft Windows:

**64-bit**
[kpsctl-v1.0.0-rc5-windows_amd64.zip](https://kpsctl-cli.s3-us-west-2.amazonaws.com/kpsctl-v1.0.0-rc5-windows_amd64.zip)

**32-bit**
[kpsctl-v1.0.0-rc5-windows_386.zip](https://kpsctl-cli.s3-us-west-2.amazonaws.com/kpsctl-v1.0.0-rc5-windows_386.zip)

1. Extract the .EXE file from the zip file.
2. Rename the extracted .EXE file to kpsctl.exe
2. Open a command windows, navigate to the .EXE file, and run it.


# Configuring kpsctl CLI
Configure the kpsctl CLI to interact with Karbon Platform Services for IoT API, including your user role and credentials.
A kpsctl CLI user operates under a specific context. A user email address, password, and context name defines the context.
The user operates according to the role assigned to their user name and email for Karbon Platform Services for IoT users. 

For example, create a context named local_user_ctx_1 for an existing local user user1@contoso.com with a password of nutanix/4u.
(A local user is a user that does not have My Nutanix portal credentials.) 

**Note**: You can only specify an existing local user email name and password in this example. 
If you do not have a local user, create one according to the [Creating a User topic](https://portal.nutanix.com/#/page/docs/details?targetId=kpsctl-Infra-Admin-Guide:edg-iot-add-users-t.html) at the Nutanix Support Portal.

```
kpsctl config create-context local_user_ctx_1 --email user1@contoso.com --password nutanix/4u
```


# Autocompletion for bash/zsh
The kpsctl CLI supports autocompletion with bash on MacOS and Linux. Using this command for zsh is experimental and not formally supported.
To configure auto-completion for your favorite shell, run this command and follow the instructions.
```
kpsctl completion --help
```

# Using samples
Under `samples/` directory, there are samples for various Karbon Platform Services for IoT entities including Data Sources, Data Pipelines, Applications, etc. These samples can be used as a reference when creating your own Karbon Platform Services for IoT entities. Each entity has a README which explains the details of each individual attribute of that entity.
