#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

shopt -s globstar

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)" # Directory where this script exists.
__root="$(cd "$(dirname "${__dir}")" && pwd)"         # Root directory of project.


cd "$__root"
rm -rf \
	"$__root/internal/example/library/bare" \
	"$__root/internal/example/library/ent" \
	"$__root/internal/example/library/schema" \
	"$__root/internal/example/library/"*.pb.go \
	"$__root/internal/example/proto/example/library/entity.svc.proto"
