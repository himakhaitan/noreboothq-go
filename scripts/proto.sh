#!/usr/bin/env bash

set -euo pipefail

# Constants
PROTO_SRC_DIR="idl"
PROTO_OUT_DIR="proto"

# Tool versions
PROTOC_GEN_GO_VERSION="v1.36.6"
PROTOC_GEN_GO_GRPC_VERSION="v1.5.1"

export PATH="$(go env GOPATH)/bin:$PATH"

# Color codes
COLOR_RESET='\033[0m'
COLOR_INFO='\033[1;34m'    # Bright Blue
COLOR_SUCCESS='\033[1;32m' # Bright Green
COLOR_ERROR='\033[1;31m'   # Bright Red
COLOR_WARN='\033[1;33m'    # Bright Yellow

info() {
  echo -e "${COLOR_INFO}info${COLOR_RESET}\t: $1"
}

success() {
  echo -e "${COLOR_SUCCESS}success${COLOR_RESET}\t: $1"
}

error() {
  echo -e "${COLOR_ERROR}error${COLOR_RESET}\t: $1"
}

warn() {
  echo -e "${COLOR_WARN}warn${COLOR_RESET}\t: $1"
}

ensure_tools() {
  info "Checking for required tools..."

  if ! command -v protoc &> /dev/null; then
    error "'protoc' compiler is not installed."
    info "Please install it from https://grpc.io/docs/protoc-installation/"
    exit 1
  else
    success "Found 'protoc' compiler."
  fi

  if ! command -v protoc-gen-go &> /dev/null; then
    info "Installing 'protoc-gen-go' plugin (version ${PROTOC_GEN_GO_VERSION})..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOC_GEN_GO_VERSION}
    success "Installed 'protoc-gen-go'."
  else
    success "'protoc-gen-go' plugin already installed."
  fi

  if ! command -v protoc-gen-go-grpc &> /dev/null; then
    info "Installing 'protoc-gen-go-grpc' plugin (version ${PROTOC_GEN_GO_GRPC_VERSION})..."
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@${PROTOC_GEN_GO_GRPC_VERSION}
    success "Installed 'protoc-gen-go-grpc'."
  else
    success "'protoc-gen-go-grpc' plugin already installed."
  fi
}

generate_protos() {
  local proto_targets=()

  if [[ $# -eq 0 ]]; then
    info "No proto files or directories specified as arguments."
    info "Searching for all '.proto' files under '${PROTO_SRC_DIR}'..."
    while IFS= read -r proto; do
      proto_targets+=("$proto")
    done < <(find "${PROTO_SRC_DIR}" -name '*.proto')
  else
    info "Processing input arguments to find proto files..."
    for target in "$@"; do
      if [[ -d "$target" ]]; then
        info "Found directory: '$target'. Searching for .proto files inside..."
        mapfile -t dir_protos < <(find "$target" -name '*.proto')
        proto_targets+=("${dir_protos[@]}")
        info "Found ${#dir_protos[@]} proto files in '$target'."
      elif [[ -f "$target" ]]; then
        info "Found proto file: '$target'. Adding to generation list."
        proto_targets+=("$target")
      else
        warn "'$target' is not a valid file or directory, skipping."
      fi
    done
  fi

  if [[ ${#proto_targets[@]} -eq 0 ]]; then
    error "No proto files found to generate. Aborting."
    exit 1
  fi

  info "Cleaning previous generated files in '${PROTO_OUT_DIR}'..."
  rm -rf "${PROTO_OUT_DIR}"
  mkdir -p "${PROTO_OUT_DIR}"
  info "Output directory '${PROTO_OUT_DIR}' ready."

  info "Generating Go code from proto files..."
  protoc \
    --proto_path="${PROTO_SRC_DIR}" \
    --go_out="${PROTO_OUT_DIR}" \
    --go_opt=paths=source_relative \
    --go-grpc_out="${PROTO_OUT_DIR}" \
    --go-grpc_opt=paths=source_relative \
    "${proto_targets[@]}"

  success "Proto generation succeeded for ${#proto_targets[@]} files."
}

main() {
  info "Starting proto generation script..."
  ensure_tools
  generate_protos "$@"
  success "Proto generation complete."
}

main "$@"