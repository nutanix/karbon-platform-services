#!/bin/bash

# Invoked if the packaging is set to custom
pushd "${PROJECT_ROOT_FOLDER}"
. ./.workspace/.canaveralrc
. ./custom/common.sh
popd

echo "noop"
