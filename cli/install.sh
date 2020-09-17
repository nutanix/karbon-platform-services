TARGET_DIR=${1:-/usr/local/bin}
LOC="$( cd "$( dirname "$0" )" && pwd )"
echo "copying CLI from $LOC to $TARGET_DIR/kps"
if ! cp $LOC/kps $TARGET_DIR/kps; then
    echo "failed to copy CLI binary to  $TARGET_DIR/kps"
fi
echo "successfully copied CLI binary to $TARGET_DIR/kps. Please make sure $TARGET_DIR is added to your \$PATH"
