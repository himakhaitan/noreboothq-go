# `scripts`

## 📋 Overview

This folder contains automation scripts that simplify common development tasks such as code generation, build automation, and environment setup.

Currently, it includes a script to generate Go gRPC code from `.proto` files.

Use this folder as a reference to create and maintain scripts that help automate repetitive or complex workflows in your project.

## 🧩 Responsibilities

- Automate generation of Go gRPC code from Protocol Buffer definitions (`proto.sh`)
- Provide a centralized place for any scripts that improve developer efficiency and consistency

## 📂 Contents

- `proto.sh` — Generates Go gRPC code from `.proto `files using protoc and relevant plugins

## 🔍 Sample Use Cases

Here are common types of scripts that such a folder might contain, to help you understand its potential scope:

| Script Name     | Purpose                                            |
| --------------- | -------------------------------------------------- |
| `proto.sh`      | Generate Go code from `.proto` files               |
| `migrate_db.sh` | Run database migration scripts on your Postgres DB |
| `build.sh`      | Build project binaries or Bazel targets            |
| `clean.sh`      | Clean build artifacts and generated files          |
| `env_load.sh`   | Load environment variables from `.env` files       |
| `test.sh`       | Run automated test suites                          |
| `deploy.sh`     | Automate deployment tasks to staging or production |

Feel free to create or modify scripts based on your project needs.

## ⚙️ How to Use / Run

Make sure scripts have execute permissions:

```bash
chmod +x scripts/*.sh
```

Run the protobuf generation script as an example:

```bash
./scripts/proto.sh
```

## 🔧 Important Details

- You should run `proto.sh` whenever `.proto` files change to keep generated code updated.
- Scripts should be version-controlled and updated to keep pace with tooling and project changes.
- This folder acts as a single source of automation for your project workflows.

## 📝 Notes / Tips

- Review scripts before running to understand their effects.
- Customize scripts to fit your workflow and tooling preferences.
- Adding scripts here improves reproducibility and reduces manual errors.

## 🔗 Related Modules / Folders

- `idl/` — Source Protocol Buffer files consumed by proto.sh
- `proto/` - Output of `proto.sh` i.e. generated Go code
- `services/` — Services relying on generated protobuf code

### 📚 References

- [Protocal Buffers Doc](https://protobuf.dev/)
- [Bash Scripting Basics](https://tldp.org/LDP/Bash-Beginners-Guide/html/)