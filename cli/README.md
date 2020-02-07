# Introducing the xi-iot Command Line Interface (CLI)

**Note**: Nutanix is actively developing the xi-iot CLI. As a result, the software published here is release candidate (RC) quality until further notice. 
Use this software at your own risk and do not use it in a production environment or deployment. The latest release candidate is shown below.

The xi-iot CLI is a command line tool that helps you manage your Xi IoT resources from your terminal or shell. With minimal 
configuration, you can start using Xi IoT functionality equal to that provided by the browser-based Xi IoT Management 
Console from the command prompt in your favorite terminal program.

## Shell Support
**Linux and macOS** You can use common shell programs such as bash and zsh.
Note that autocompletion for zsh is experimental in this release.

# Installing xi-iot CLI
## MacOS
Download the xi-iot CLI for MacOS:

**64-bit**
[xi-iot-v1.0.0-rc3-darwin_amd64.tar.gz](https://xi-iot-cli.s3-us-west-2.amazonaws.com/xi-iot-v1.0.0-rc3-darwin_amd64.tar.gz)

**32-bit**
[xi-iot-v1.0.0-rc3-darwin_386.tar.gz](https://xi-iot-cli.s3-us-west-2.amazonaws.com/xi-iot-v1.0.0-rc3-darwin_386.tar.gz)

Optionally, from your terminal or shell, use `wget`.
```
$ wget https://xi-iot-cli.s3-us-west-2.amazonaws.com/xi-iot-v1.0.0-rc3-darwin_amd64.tar.gz
$ mkdir xi-iot && tar zxvf xi-iot-v1.0.0-rc3-darwin_amd64.tar.gz -C xi-iot
$ sudo xi-iot/install [*optional_path*]
```

**Note** The xi-iot binary default installation directory is `/usr/local/bin/`. 
To install the xi-iot binary in a custom location, specify the custom directory or path as follows, where
my/custom/dir is the installation path:
```
$ sudo xi-iot/install /my/custom/dir
```

## Linux
Download the xi-iot CLI for Linux:

**64-bit**
[xi-iot-v1.0.0-rc3-linux_amd64.tar.gz](https://xi-iot-cli.s3-us-west-2.amazonaws.com/xi-iot-v1.0.0-rc3-linux_amd64.tar.gz)

**32-bit**
[xi-iot-v1.0.0-rc3-linux_386.tar.gz](https://xi-iot-cli.s3-us-west-2.amazonaws.com/xi-iot-v1.0.0-rc3-linux_386.tar.gz)

Optionally, from your terminal or shell, use `wget`.
```
$ wget https://xi-iot-cli.s3-us-west-2.amazonaws.com/xi-iot-v1.0.0-rc3-linux_amd64.tar.gz
$ mkdir xi-iot && tar zxvf xi-iot-v1.0.0-rc3-linux_amd64.tar.gz -C xi-iot
$ sudo xi-iot/install [*optional_path*]
```

**Note** The xi-iot binary default installation directory is `/usr/local/bin/`. 
To install the xi-iot binary in a custom location, specify the custom directory or path as follows, where
my/custom/dir is the installation path:
```
$ sudo xi-iot/install /my/custom/dir
```

# Configuring xi-iot CLI
Configure the xi-iot CLI to interact with Xi IoT API, including your user role and credentials.
A xi-iot CLI user operates under a specific context. A user email address, password, and context name defines the context.
The user operates according to the role assigned to their user name and email for Xi IoT users. 

For example, create a context named local_user_ctx_1 for an existing local user user1@contoso.com with a password of nutanix/4u.
(A local user is a user that does not have My Nutanix portal credentials.) 

**Note**: You can only specify an existing local user email name and password in this example. 
If you do not have a local user, create one according to the [Creating a User topic](https://portal.nutanix.com/#/page/docs/details?targetId=Xi-IoT-Infra-Admin-Guide:edg-iot-add-users-t.html) at the Nutanix Support Portal.

```
xi-iot config create-context local_user_ctx_1 --email user1@contoso.com --password nutanix/4u
```


# Autocompletion for bash/zsh
The xi-iot CLI supports autocompletion with bash on MacOS and Linux. Using this command for zsh is experimental and not formally supported.
To configure auto-completion for your favorite shell, run this command and follow the instructions.
```
xi-iot completion --help
```

# Using samples
Under `samples/` directory, there are samples for various Xi IoT entities including Data Sources, Data Pipelines, Applications, etc. These samples can be used as a reference when creating your own Xi IoT entities. Each entity has a README which explains the details of each individual attribute of that entity.
