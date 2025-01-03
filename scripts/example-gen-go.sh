#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

shopt -s globstar

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)" # Directory where this script exists.
__root="$(cd "$(dirname "${__dir}")" && pwd)"         # Root directory of project.

NAME="orm-go"
APP="protoc-gen-$NAME"

cd "$__root"
"$__root"/scripts/example-gen-service.sh

go build -o "/tmp/$APP" "./cmd/$APP"

MODULE_NAME=github.com/lesomnus/proto-orm
PROTO_ROOT="${__root}/internal/example/proto"
OUTPUT_DIR="${__root}"
cd "${PROTO_ROOT}"

protoc \
	--"plugin=protoc-gen-$NAME"="/tmp/$APP" \
	--proto_path="$PROTO_ROOT" \
	--proto_path="$__root/proto" \
	\
	--go_out="${OUTPUT_DIR}" \
	--go_opt=module="${MODULE_NAME}" \
	--go_opt=default_api_level=API_OPAQUE \
	\
	--go-grpc_out="${OUTPUT_DIR}" \
	--go-grpc_opt=module="${MODULE_NAME}" \
	\
	--"$NAME"_out="$OUTPUT_DIR" \
	--"$NAME"_opt=module="$MODULE_NAME" \
	--"$NAME"_opt=naming="library.g.go" \
	\
	"$PROTO_ROOT"/**/*.proto
