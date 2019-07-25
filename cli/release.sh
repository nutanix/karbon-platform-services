export PROJECT_ROOT_FOLDER=$PWD
LOCAL_BUILD=1 . custom/common.sh 
LOCAL_BUILD=1 build_and_release $1
