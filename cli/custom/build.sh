LOCAL_BUILD=${LOCAL_BUILD:-0}
if [[ ${LOCAL_BUILD} -eq 0 ]]; then
  . ./.workspace/.canaveralrc
    export GOPATH=~/.go_workspace
    export GOROOT=/tmp/go
else
  export SCRIPT_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
  echo "SCRIPT_DIR=${SCRIPT_DIR}"
  export PROJECT_ROOT_FOLDER=$(cd ${SCRIPT_DIR}/.. && pwd)
fi

. ./custom/common.sh
function main {
    case $1 in
    setup) setup_gospace ;;
    build) build ;;
    *) setup_gospace && do_all ;;
    esac
}

main $@
