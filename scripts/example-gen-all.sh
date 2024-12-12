#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)" # Directory where this script exists.
__root="$(cd "$(dirname "${__dir}")" && pwd)"         # Root directory of project.


cd "$__root"

./scripts/example-clean.sh
./scripts/example-gen-service.sh
./scripts/example-gen-ent.sh
./scripts/example-gen-ent-grpc.sh

go generate ./_example/library
