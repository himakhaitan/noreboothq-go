# 🖥️ NoRebootHQ

**A Hands On Microservices Project in Go**

NoRebootHQ is a hands-on backend project where you'll build a production-grade system to manage, version, and deploy application configurations in real-time — without restarting services.



This is a starter repo that gives you the foundational project setup and structure you'll use to build the full platform using:

> Go · gRPC · ProtoBuf · PostgreSQL · Redis · Bazel · Docker and more

Ready to go beyond tutorials and build a system that companies actually need?

<p>
<img src="https://badgen.net/github/issues/himakhaitan/noreboothq-go?style=flat-square&scale=1.4">
&nbsp;
<img alt="node-current" src="https://badgen.net/github/stars/himakhaitan/noreboothq-go?style=flat-square&scale=1.4">&nbsp;
<img alt="APM" src="https://badgen.net/github/forks/himakhaitan/noreboothq-go?style=flat-square&scale=1.4">&nbsp;
<img alt="node-current" src="https://badgen.net/github/closed-issues/himakhaitan/noreboothq-go?style=flat-square&scale=1.4">
</p>

## 📁 Project Overview

NoRebootHQ is a Dynamic Configuration Management System that lets you manage, version, and roll out config changes and feature flags without restarting services.

### 🧠 Why This Matters

In modern backend architectures, especially those powered by microservices, runtime behavior is increasingly controlled through dynamic configurations. These include:

- Feature flags
- Rate limits
- Access rules
- Environment-specific settings
- A/B testing controls

Without a proper system in place, updating these values often means triggering service restarts, risking downtime, or relying on error-prone manual deployments.

Platforms like Facebook, LinkedIn, Uber, and Airbnb have long adopted internal systems that handle dynamic configuration safely and efficiently:

- Facebook uses Gatekeeper for dynamic feature control
- Uber developed Flipr, a tool for dynamic configuration management, including feature flags and incremental rollouts.
- Airbnb uses SmartConfig to manage and deploy configurations without service restarts.

These tools offer versioning, rollback, scoping, and runtime propagation — all of which increase system flexibility while minimizing risk.

### 🧩 What You’ll Build

In this project, you'll architect and scaffold the foundation of a complete dynamic config system. With NoRebootHQ, you'll gain experience working with:

- A modular microservices-based architecture using Go
- gRPC APIs to create, fetch, and activate configurations
- PostgreSQL + GORM for persistent, queryable storage
- Redis for high-speed config reads and caching
- Bazel and Docker to structure and containerize the project
- Version control and change activation mechanisms

All designed to simulate the type of tooling used by platform teams in tech-forward engineering organizations.

## 📦 What's in This Repo?

This is a sample starter repo for the NoRebootHQ project. It includes basic setup, code structure, and select features to give you a clear understanding of how the full project is organized and built.

### ✅ Included

- Project layout built with Go
- Starter setup for PostgreSQL
- gRPC setup with proto definitions
- A partially implemented auth-service to demonstrate architecture and coding standards
- Comments, README (at all folder levels) and folder structure to help you navigate and extend

This is intended to give you a glimpse into the project’s architecture and best practices — not a full codebase.

### 🚫 What’s Not Included

This repo does not include the complete codebase — and it never will.

Instead of giving away a finished project to copy, the full experience includes:

- 📘 Step-by-step implementation guides
- 🧠 Detailed explanations of architectural decisions
- 🧩 System design breakdowns for every service
- 💬 Community and 1:1 support via Telegram + TopMate

This is designed to help you build it yourself, with full context and confidence — instead of copying code you can’t explain or use effectively in interviews.

## 🚀 Getting Started

This is not a full-fledged project — it’s a guided sample repo to help you understand how a scalable backend system is structured and organized.

### 📂 How to Use This Repo

Go through each folder in the repo — especially:

- `services/auth` : explore how the auth service is structured
- `idl/` : check how service contracts are defined using Protocal Buffers
- `scripts/` :  automated workflows — like protobuf codegen, build steps, etc.
- `shared/` : common utilities used across services
(e.g., YAML config parsing, structured logging, and env loading)

Each folder contains inline comments or README snippets to help you understand:

- What the folder is for and why it exists
- How it contributes to clean architecture and service modularity
- How you can adapt or extend it for your own projects

> 💡 This is meant to guide how you build, not give you finished code to run.

## 📘 Build the Full Project with Me

Want to go from exploring this repo to actually building the full project step-by-step?

The complete digital product includes:

- ✅ Getting Started Guide
- 🧠 System Design Deep Dive
- 🧩 Full Codebase Walkthrough (with why behind every decision)
- 🛠️ Step-by-Step Implementation Instructions
- 🎯 Interview Prep Resources
- Telegram Community
- 📅 Optional Live Cohort Sessions (with throughout support and mentorship)

Join the bootcamp or access the digital product on Topmate

🔗 Join via Topmate

## ❓ Questions or Support

If you have any questions, feel free to reach out!

👉 Message me directly on [Topmate](https://topmate.io/himakhaitan/1461713/pay)
👉 Or DM me on [LinkedIn](https://www.linkedin.com/in/himakhaitan/)

I'm always happy to help curious and committed learners.

## 📄 License & Usage

This sample repo is provided strictly for educational use.

- You are welcome to complete the project for personal learning, academic use, or interview preparation.
- However, you may not sell, redistribute, or create professional products derived from this code without permission.
- This is a learning reference — intended to help you build and understand, not for commercial reuse or resale.