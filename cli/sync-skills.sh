#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"
rm -rf embedded/skills
mkdir -p embedded/skills
cp -r ../.agents/skills/* embedded/skills/
