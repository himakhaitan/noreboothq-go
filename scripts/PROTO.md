# Proto Generation Script

This bash script automates the process of generating Go code from Protocol Buffer (`.proto`) files using the `protoc` compiler and the Go protobuf plugins. It ensures required tools are installed, finds `.proto` files, cleans previous outputs, and generates Go code with gRPC support.

## Script Overview

The script performs the following key steps:

1. Ensure required tools are installed (`protoc`, `protoc-gen-go`, `protoc-gen-go-grpc`)
2. Collect `.proto` files to generate code from (defaults to all `.proto` files in `idl/` folder)
3. Clean the output directory (`proto/`)
4. Run `protoc` with appropriate flags to generate Go and gRPC code
5. Report success or failure

## Script Breakdown

### 1. Constants and Environment Setup

```bash
PROTO_SRC_DIR="idl"
PROTO_OUT_DIR="proto"

PROTOC_GEN_GO_VERSION="v1.36.6"
PROTOC_GEN_GO_GRPC_VERSION="v1.5.1"

export PATH="$(go env GOPATH)/bin:$PATH"
```

- `PROTO_SRC_DIR` is the root directory containing your `.proto` files.
- `PROTO_OUT_DIR` is where the generated Go files will be saved.
- The versions of the Go protobuf plugins to install or check are defined.
- Adds the Go bin directory to `PATH` to ensure installed tools are accessible.

### 2. Color-coded Logging Functions

```bash
COLOR_RESET='\033[0m'
COLOR_INFO='\033[1;34m'    # Bright Blue
COLOR_SUCCESS='\033[1;32m' # Bright Green
COLOR_ERROR='\033[1;31m'   # Bright Red
COLOR_WARN='\033[1;33m'    # Bright Yellow

info() { echo -e "${COLOR_INFO}info${COLOR_RESET}\t: $1"; }
success() { echo -e "${COLOR_SUCCESS}success${COLOR_RESET}\t: $1"; }
error() { echo -e "${COLOR_ERROR}error${COLOR_RESET}\t: $1"; }
warn() { echo -e "${COLOR_WARN}warn${COLOR_RESET}\t: $1"; }
```

These helper functions print messages with color coding to differentiate info, success, errors, and warnings, improving readability in the terminal.

### 3. Tool Checks and Installation

```bash
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
```

- Verifies if `protoc` is installed; if not, stops execution with instructions to install it.
- Checks for `protoc-gen-go` and `protoc-gen-go-grpc` plugins.
- If missing, installs them at the specified versions using `go install`.
- Prints appropriate success or error messages for each check.

### 4. Proto Files Collection and Validation

```bash
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
  ...
}
```

- If no command line arguments are passed to the script, it recursively finds all `.proto` files under `idl/`.
- If arguments are passed, it treats each argument as a file or directory:
    - For directories, recursively finds `.proto` files inside.
    - For files, adds them directly if valid.
- Invalid paths are skipped with a warning.
- If no proto files are found, the script exits with an error.

### 5. Cleaning Output Directory

```bash
 info "Cleaning previous generated files in '${PROTO_OUT_DIR}'..."
  rm -rf "${PROTO_OUT_DIR}"
  mkdir -p "${PROTO_OUT_DIR}"
  info "Output directory '${PROTO_OUT_DIR}' ready."
```

- Deletes the previous generated files and directory to avoid stale code.
- Recreates the output directory to ensure it's ready to receive new generated files.

### 6. Generating Go Code from Proto Files

```bash
 info "Generating Go code from proto files..."
  protoc \
    --proto_path="${PROTO_SRC_DIR}" \
    --go_out="${PROTO_OUT_DIR}" \
    --go_opt=paths=source_relative \
    --go-grpc_out="${PROTO_OUT_DIR}" \
    --go-grpc_opt=paths=source_relative \
    "${proto_targets[@]}"
```

- Calls the `protoc` compiler with flags:
    - `--proto_path`: Root directory for `.proto` imports
    - `--go_out`: Output directory for generated Go protobuf messages
    - `--go_opt=paths=source_relative`: Makes generated files have paths relative to the source `.proto` file (helps keep directory structure clean)
    - `--go-grpc_out`: Output directory for generated gRPC service code
    - `--go-grpc_opt=paths=source_relative`: Same path option for gRPC code
- Passes all `.proto` files found to the compiler.

### 7. Main Function and Execution

```bash
main() {
  info "Starting proto generation script..."
  ensure_tools
  generate_protos "$@"
  success "Proto generation complete."
}

main "$@"
```

- Entry point of the script.
- Prints starting message.
- Ensures tools are installed.
- Runs the proto generation function passing any CLI args.
- Prints success message on completion.

## Summary

- The script is designed to be idempotent: you can run it multiple times without leftover artifacts.
- It automates tool installation if missing, helping new developers get started easily.
- It supports flexible input: generate for all .proto files or a custom subset.
- It produces Go protobuf and gRPC code with paths relative to source, ensuring clean project structure.
- Provides color-coded informative logs to understand progress and errors clearly.