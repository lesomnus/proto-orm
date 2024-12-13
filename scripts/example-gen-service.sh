#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

shopt -s globstar

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)" # Directory where this script exists.
__root="$(cd "$(dirname "${__dir}")" && pwd)"         # Root directory of project.

NAME="orm-service"
APP="protoc-gen-$NAME"

cd "$__root"

go build -o "/tmp/$APP" "./cmd/$APP"

MODULE_NAME=github.com/lesomnus/proto-orm/_example
PROTO_ROOT="${__root}/_example/proto"
OUTPUT_DIR="${__root}/_example/proto/example"
cd "${PROTO_ROOT}"

rm -rf "$OUTPUT_DIR"/**/*.svc.proto

protoc \
	--"plugin=protoc-gen-$NAME"="/tmp/$APP" \
	--proto_path="$PROTO_ROOT" \
	--proto_path="$__root/proto" \
	\
	--"$NAME"_out="$OUTPUT_DIR" \
	--"$NAME"_opt=module="$MODULE_NAME" \
	--"$NAME"_opt=naming="{{ .Name }}.g.proto" \
	\
	"$PROTO_ROOT"/**/*.proto
