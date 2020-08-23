#!/bin/bash

GO=@@GO@@
GAZELLE_PATH=@@GAZELLE@@
RUNFILES=$(pwd)
GO_RUNTIME="$RUNFILES"/"$GO"
GAZELLE="$RUNFILES"/"$GAZELLE_PATH"

cd "$BUILD_WORKSPACE_DIRECTORY"

"$GO_RUNTIME" mod tidy 2>&1 | grep -v bazel
"$GO_RUNTIME" mod vendor 2>&1 | grep -v bazel
find vendor -type f -name BUILD.bazel -delete
find vendor -type f -name BUILD -delete
"$GAZELLE" update
