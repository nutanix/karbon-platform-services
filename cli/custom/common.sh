#!/bin/bash
LOCAL_BUILD=${LOCAL_BUILD:-0}

CLOUDMGMT_API_SWAGGER_NAME="cloudmgmtapi_swagger.json"
EDGE_API_SWAGGER_NAME="controller_swagger.json"
DEVTOOLS_API_SWAGGER_NAME="devtools_swagger.json"

# golang specific variables
export GOPATH=~/.go_workspace
export SRC=$GOPATH/src
export PKG=$GOPATH/pkg
export BIN=$GOPATH/bin
export GO_VER="1.12.1"
export GOINSTALL=/tmp/go-$GO_VER
export GOROOT=$GOINSTALL/go
export PATH=$GOROOT/bin:$PATH
export SHERLOCK_SRC=${PROJECT_ROOT_FOLDER}/services
export XI_IOT_CLI_SRC=$SRC/xi-iot-cli
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     MACHINE=linux;;
    Darwin*)    MACHINE=darwin;;
    CYGWIN*)    MACHINE=Cygwin;;
    MINGW*)     MACHINE=MinGw;;
    *)          MACHINE="UNKNOWN:${unameOut}"
esac

# OS and ARCH of host
export GOARCH=amd64
export GOOS=${MACHINE}

function op_url {
  local OP=$1
  shift
  local URL=$1
  shift

  echo "Calling: ${URL} ..."
  # get output, append HTTP status code in separate line, discard error message
  set +e
  OUT=$( curl -k -LSfsw '\n%{http_code}' -X "${OP}" "${URL}" ) 2>/dev/null
  EXIT_CODE=$?
  set -e

  # get exit code
  echo "Exit Status: ${EXIT_CODE}"

  STATUS_CODE=`echo "${OUT}" | tail -n1`
  RETURN_BODY=''

  if [[ ${EXIT_CODE} -ne 0 ]] ; then
    if [[ ${EXIT_CODE} -ne 22 ]] ; then
      echo "Error ${EXIT_CODE}"

      # print HTTP error
      echo "HTTP Error: $(echo "$OUT" | tail -n1 )"
      test ${EXIT_CODE} -eq 0
    fi
  else
    RETURN_BODY=`echo "$OUT" | head -n1`
  fi
}

function fetch_latest {
    local SERVICE_NAME=$1
    shift
    local ARTIFACT=$1
    shift

    DOWNLOAD_KEY_URL="http://canaveral-artifacts.corp.nutanix.com:8080/artifacts/build-artifacts/presigned-download/edgecomputing/${SERVICE_NAME}/latest.json"

    op_url "GET" ${DOWNLOAD_KEY_URL}

    if [[ ${STATUS_CODE} -eq 200 ]] ; then
    echo "Success, HTTP status is: ${STATUS_CODE}"
    else
    echo "HTTP Status Code was not expected: ${STATUS_CODE}, expecting 200 ..."
    test 0 -eq 1
    fi

    DOWNLOAD_LINK=`echo ${RETURN_BODY} | jq -r ".url"`
    wget --no-check-certificate -c -O "./latest.json" "${DOWNLOAD_LINK}" -nv

    LATEST_BUILD_NUM=`cat latest.json`

    # Remove latest.json file
    rm latest.json

    op_url "GET" "http://canaveral-artifacts.corp.nutanix.com:8080/artifacts/build-artifacts/presigned-download/edgecomputing/${LATEST_BUILD_NUM}/${ARTIFACT}"

    if [[ ${STATUS_CODE} -eq 200 ]] ; then
    echo "Success, HTTP status is: ${STATUS_CODE}"
    #echo "Response is:"
    #echo "${RETURN_BODY}"
    else
    echo "HTTP Status Code was not expected: ${STATUS_CODE}, expecting 200 ..."
    test 0 -eq 1
    fi

    ARTIFACT_DOWNLOAD_LINK=`echo ${RETURN_BODY} | jq -r ".url"`

    echo "ARTIFACT DOWNLOAD LINK ${ARTIFACT_DOWNLOAD_LINK}"

    wget --no-check-certificate -c -O "./${ARTIFACT}" "${ARTIFACT_DOWNLOAD_LINK}" -nv
}

function fetch_artifact {
    local ARTIFACT=$1
    shift
    local BUILD_NUM=$1
    shift

    op_url "GET" "http://canaveral-artifacts.corp.nutanix.com:8080/artifacts/build-artifacts/presigned-download/edgecomputing/${BUILD_NUM}/${ARTIFACT}"

    if [[ ${STATUS_CODE} -eq 200 ]] ; then
    echo "Success, HTTP status is: ${STATUS_CODE}"
    #echo "Response is:"
    #echo "${RETURN_BODY}"
    else
    echo "HTTP Status Code was not expected: ${STATUS_CODE}, expecting 200 ..."
    test 0 -eq 1
    fi

    ARTIFACT_DOWNLOAD_LINK=`echo ${RETURN_BODY} | jq -r ".url"`

    echo "ARTIFACT DOWNLOAD LINK ${ARTIFACT_DOWNLOAD_LINK}"

    wget --no-check-certificate -c -O "./${ARTIFACT}" "${ARTIFACT_DOWNLOAD_LINK}" -nv
}

# set up gospace
function setup_gospace {
    mkdir -p $SRC
    # remove potential old symlink in gospace
    rm -rf $XI_IOT_CLI_SRC

    # create symlink in gospace
    ln -s ${PROJECT_ROOT_FOLDER}/services $XI_IOT_CLI_SRC
}

function install_goswagger {
    pushd $XI_IOT_CLI_SRC/build 
    if [ ! -f ./swagger ]; then
        #latestv=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r .tag_name)
        #echo "Install swagger " $latestv
        latestv=0.15.0
        curl -o ./swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/$latestv/swagger_$(echo `uname`|tr '[:upper:]' '[:lower:]')_amd64
        chmod +x ./swagger
    else
        echo "swagger already installed"
    fi
    popd
}

function fetch_dependencies {
    pushd $XI_IOT_CLI_SRC

    #delete existing swagger file


    rm -rf $CLOUDMGMT_API_SWAGGER_NAME $EDGE_API_SWAGGER_NAME $DEVTOOLS_API_SWAGGER_NAME
    #fetch_latest "sherlock-cloudmgmt" $CLOUDMGMT_API_SWAGGER_NAME
    fetch_artifact $CLOUDMGMT_API_SWAGGER_NAME 5351
    fetch_artifact $DEVTOOLS_API_SWAGGER_NAME 5088
    fetch_artifact $EDGE_API_SWAGGER_NAME 1683
    popd
}

function generate_swagger {
    pushd $XI_IOT_CLI_SRC
    DEVTOOLS_SWAGGER_GEN=$XI_IOT_CLI_SRC/generated/devtools_swagger
    EDGE_SWAGGER_GEN=$XI_IOT_CLI_SRC/generated/edge_swagger
    SWAGGER_GEN=$XI_IOT_CLI_SRC/generated/swagger
    mkdir -p $DEVTOOLS_SWAGGER_GEN
    mkdir -p $EDGE_SWAGGER_GEN
    mkdir -p $SWAGGER_GEN
    echo "generate swagger client stubs"
    ./build/swagger generate client -t $DEVTOOLS_SWAGGER_GEN -f ./$DEVTOOLS_API_SWAGGER_NAME -A sherlock
    ./build/swagger generate client -t $EDGE_SWAGGER_GEN -f ./$EDGE_API_SWAGGER_NAME -A sherlock
    ./build/swagger generate client -t $SWAGGER_GEN -f ./$CLOUDMGMT_API_SWAGGER_NAME -A sherlock
    echo "done generating swagger client stubs"
    popd
}

function install_golang {
    if [ -d $GOROOT ]; then
        echo "golang is already installed at $GOINSTALL"
        return
    fi

    local GO_TARBALL="go${GO_VER}.${GOOS}-${GOARCH}.tar.gz"

    echo "Downloading go bundle ${GO_TARBALL}"
    curl -4 -o "/tmp/${GO_TARBALL}" "https://storage.googleapis.com/golang/${GO_TARBALL}"
    echo "Installing go version $GO_VER in $GOINSTALL"
    mkdir -p $GOINSTALL
    tar -xvzf /tmp/${GO_TARBALL} -C $GOINSTALL 2>/dev/null 1>&2
}

function setup {
    # pushd $XI_IOT_CLI_SRC
    setup_gospace
    install_golang
    mkdir -p $XI_IOT_CLI_SRC/build
    pushd $XI_IOT_CLI_SRC/build
    install_goswagger
    generate_swagger
    popd # build
}

function build {
    pushd $XI_IOT_CLI_SRC
    echo "$PWD"
    go build -ldflags '-w -s' -a -installsuffix cgo -o ./build/xi-iot ./xi-iot/main.go
    popd
}

function install_cli {
    # pushd $SHERLOCK_SRC
    go install $XI_IOT_CLI_SRC/xi-iot/
    # popd
}

function top_level_commands {
    if ! $XI_IOT_CLI_SRC/build/xi-iot help; then
        echo "failed to run help command"
    fi
    declare -a commands=("create" "get" "delete")
    for i in "${commands[@]}"
    do
        $XI_IOT_CLI_SRC/build/xi-iot $i --help
        if [ $? != 0 ]; then
            exit 1
        fi
    done
}

function get_commands {
    declare -a sub_commands=("datasrc" "datapipeline" "application" "category" "edge" "function" "project" "runtime")
    for i in "${sub_commands[@]}"
    do
        $XI_IOT_CLI_SRC/build/xi-iot get $i --help
        if [ $? != 0 ]; then
            exit 1
        fi
    done
}

function create_delete_sub_commands {
    declare -a sub_commands=("datasrc" "datapipeline" "application" "category")
    for i in "${sub_commands[@]}"
    do
        $XI_IOT_CLI_SRC/build/xi-iot create $i --help
        if [ $? != 0 ]; then
            exit 1
        fi
        $XI_IOT_CLI_SRC/build/xi-iot delete $i --help
        if [ $? != 0 ]; then
            exit 1
        fi
    done
}

function basic_smoke_tests {
    top_level_commands
    get_commands
    create_delete_sub_commands
}

function run_tests {
    set -x
    basic_smoke_tests

    # E2E tests
    pushd $XI_IOT_CLI_SRC/tests
    source e2e.sh
    BIN_PATH=$XI_IOT_CLI_SRC/build/xi-iot

    if [ "$MACHINE" = "darwin" ]; then
        NAME_SALT=$(date |md5 | head -c8)
    elif [ "$MACHINE" = "linux" ]; then
        NAME_SALT=$(date |md5sum | head -c8)
    fi

    if ! setup $BIN_PATH $NAME_SALT; then
        echo "failed setup step"
        popd
        exit 1
    fi

    if ! teardown $BIN_PATH $NAME_SALT; then
        echo "failed teardown step"
        popd
        exit 1
    fi

    if ! create $BIN_PATH $NAME_SALT; then
        echo "failed create step"
        # try our best to clean up, no error handling for this teardown
        # because we are going to exit 1 anyway.
        teardown $BIN_PATH $NAME_SALT
        popd
        exit 1
    fi

    if ! get $BIN_PATH $NAME_SALT; then
        echo "failed get step"
        # try our best to clean up, no error handling for this teardown
        # because we are going to exit 1 anyway.
        teardown $BIN_PATH $NAME_SALT
        popd
        exit 1
    fi

    if ! teardown $BIN_PATH $NAME_SALT; then
        echo "failed teardown step"
        popd
        exit 1
    fi

    if ! debug $BIN_PATH $NAME_SALT; then
        echo "failed debug step"
        teardown_debug $BIN_PATH $NAME_SALT
        popd
        exit 1
    fi

    if ! logging $BIN_PATH $NAME_SALT; then
        echo "failed logging step"
        teardown_logging $BIN_PATH $NAME_SALT
        popd
        exit 1
    fi
}

function clean {
    pushd $XI_IOT_CLI_SRC
    rm -rf generated
    rm -rf swagger
    rm -rf build
    rm -rf $GOROOT
    popd
}

# Usage: add_git_tag <TAG_NAME> <COMMIT_HASH>
function add_git_tag {
    pushd $XI_IOT_CLI_SRC
    if ! (git tag  -a $1 $2 -m "Releasing $1" && git push --tags); then
	popd
	return 1
    fi
    popd
}

# Usage: remove_git_tag <TAG_NAME>
function remove_git_tag {
    pushd $XI_IOT_CLI_SRC
    if ! (git tag -d $1 && git push origin :refs/tags/$1); then
        popd
	return 1
    fi
    popd
}

# Usage: gox_build <TAG_NAME>
function gox_build {
    pushd $XI_IOT_CLI_SRC
    if ! go get -u github.com/mitchellh/gox; then
	popd
	return 1
    fi

    echo "generating binaries for multiple platforms using gox"
    if ! gox -os="linux darwin" -arch="amd64 386" -output="./build/dist/xi-iot-$1-{{.OS}}_{{.Arch}}/xi-iot" ./xi-iot/; then
	popd
	return 1
    fi

    # copy install script
    for os in linux darwin; do
	for arch in amd64 386; do
	    cp $XI_IOT_CLI_SRC/documentation/install.sh "./build/dist/xi-iot-$1-${os}_${arch}/install"
	done
    done

    echo "packaging binaries to tarballs"
    pushd build/dist
    if ! find . -type d -maxdepth 1 -mindepth 1 -exec tar -czvf {}.tar.gz -C {} . \;; then
	popd
	return 1
    fi
    popd
    popd
}

# Usage: check_git_tag <TAG>
function check_git_tag {
    pushd $XI_IOT_CLI_SRC
    matching_tags=$(git ls-remote --tags origin | grep $1)
    num_matching_tags=$(git ls-remote --tags origin | grep $1 | wc -l | xargs)
    if [ $num_matching_tags != 0 ] ; then
	echo "git tag: $1 exists. existing tags: $matching_tags. please remove these tags"
	return 1
    fi
    popd
}

# Usage: build_and_release <RELEASE_TAG>
function build_and_release {
    if ! check_git_tag $1;  then
	popd
	return 1
    fi

    if ! gox_build $1; then
	echo "failed to build and package binaries"
	popd
	return 1
    fi

    if ! aws s3 cp $XI_IOT_CLI_SRC/build/dist/ s3://xi-iot-cli --recursive --exclude "*" --include "*.tar.gz" --acl public-read; then
	echo "failed to copy tarballs to s3"
	popd
	return 1
    fi

    if ! (add_git_tag $1 HEAD -m "Releasing $1" && git push --tags); then
	echo "failed to tag commit"
	popd
	return 1
    fi
}

function do_all {
    clean
    fetch_dependencies
    setup
    build
}
