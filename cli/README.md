# xi-iot CLI

CLI for managing Xi IoT resources.

# Getting Started
## Build the Xi IoT binary locally
`> <PROJECT_ROOT_DIR>/build_local.sh`

## Install the binary to your GOPATH
```
> <PROJECT_ROOT_DIR>/install.sh
```

If Golang is installed, do this step. To use this binary without specifying the full path, run:
```
> export PATH=$PATH:$GOPATH/bin
``` 
Or add the binary to your .bashrc or .zshrc file. 

# Basic set up commands
## Configuring auto completion
Configure this CLI for auto completion with either `zsh` or `bash`. To leverage auto-completion, you may have to install necessary
plug-ins for the corresponding shell.

Get more help with setting up auto-completion.
```
✗ xi-iot completion --help
```

## Configure the CLI
You can configure the CLI to use different users, tenants, and cloud management. The experience is similar to `kubectl`, where 
you can set up and switch between different contexts. For more information on using contexts, see: 
```
✗ xi-iot configure --help
```

# Some more help
A few create commands require yaml as the input. For example, you can reference data sources, app
sample yamls, in a sample directory on this repo.

# List of commands fully supported
1. `get app/category/datasrc/datapipelines/projects/edges`
2. `create app/datasrc1`
3. `delete datasrc/category/app`

Give it a shot and feel free to open feature requests or issues!

# Packaging & Releases
We use git tags to release a new version of CLI. In order to release a new version of the CLI at the current commit commit, you can run the command:
```
./release <RELEASE_VERSION>
```
 where <RELEASE_VERSION> is the version (example: v1.0.0-rc1) that you would like to release and follows semantic versioning scheme.

This script will build binaries for the supported platforms(linux, mac) and upload the tarball to s3. For more details, see documentation/quickstart.md.

Sample path for tarball: https://xi-iot-cli.s3-us-west-2.amazonaws.com/xi-iot-v1.0.0-rc1-darwin_[amd64/386].tar.gz
