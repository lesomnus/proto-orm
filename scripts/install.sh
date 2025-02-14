#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)" # Directory where this script exists.
__root="$(cd "$(dirname "${__dir}")" && pwd)"         # Root directory of project.


cd "$__root"

go install ./cmd/protoc-gen-orm-service
go install ./cmd/protoc-gen-orm-go
go install ./cmd/protoc-gen-orm-ent
go install ./cmd/protoc-gen-orm-ent-grpc
