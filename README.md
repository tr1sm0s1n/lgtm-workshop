# LGTM Workshop

Welcome to the LGTM workshop repository!

[![Build](https://github.com/tr1sm0s1n/lgtm-workshop/actions/workflows/release.yml/badge.svg)](https://github.com/tr1sm0s1n/lgtm-workshop/actions/workflows/release.yml)
[![Release](https://img.shields.io/github/v/release/tr1sm0s1n/lgtm-workshop)](https://github.com/tr1sm0s1n/lgtm-workshop/releases)
[![CI](https://github.com/tr1sm0s1n/lgtm-workshop/actions/workflows/ci.yml/badge.svg)](https://github.com/tr1sm0s1n/lgtm-workshop/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE.md)

## Workshop Goal

Learn professional development practices by:

- Writing tests before implementation (TDD)
- Using CI/CD pipelines (GitHub Actions)
- Following open-source conventions (commits, PRs, issues)

## Quick Start

### Prerequisites

- Git with SSH configured
- **Go 1.24+** OR **Node.js 24+** (choose your language)
- GitHub account
- Code editor

### Setup

```bash
# Fork this repository first (click Fork button above)

# Clone your fork
git clone git@github.com:YOUR_USERNAME/lgtm-workshop.git
cd lgtm-workshop

# For Go
go mod download
go test ./...

# For JavaScript
npm install
npm test
```
