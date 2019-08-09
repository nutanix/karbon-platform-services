TARGET_DIR=${1:-/usr/local/bin}
LOC="$( cd "$( dirname "$0" )" && pwd )"
echo "copying CLI from $LOC to $TARGET_DIR/xi-iot"
if ! cp $LOC/xi-iot $TARGET_DIR/xi-iot; then
    echo "failed to copy CLI binary to  $TARGET_DIR/xi-iot"
fi
echo "successfully copied CLI binary to $TARGET_DIR/xi-iot. Please make sure $TARGET_DIR is added to your \$PATH"
