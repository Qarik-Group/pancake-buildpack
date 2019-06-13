#!/bin/bash

set -eu

version=$(cat pancake/version)
sha=$(sha256sum pancake/cf-pancake-linux-amd64.tar.xz | awk '{print $1}')

git clone git pushme

pushd pushme

cat > manifest.yml <<YAML
---
language: pancake
default_versions:
- name: cf-pancake
  version: ${version}
dependency_deprecation_dates: []

include_files:
  - README.md
  - VERSION
  - bin/supply
  - manifest.yml
pre_package: scripts/build.sh

dependencies:
- name: cf-pancake
  version: ${version}
  uri: https://github.com/starkandwayne/cf-pancake/releases/download/v${version}/cf-pancake-linux-amd64.tar.xz
  sha256: ${sha}
  cf_stacks:
  - cflinuxfs2
  - cflinuxfs3
YAML

 echo "${version}" > VERSION

if [[ "$(git status -s)X" != "X" ]]; then
  set +e
  if [[ -z $(git config --global user.email) ]]; then
    git config --global user.email "drnic+bot@starkandwayne.com"
  fi
  if [[ -z $(git config --global user.name) ]]; then
    git config --global user.name "CI Bot"
  fi

  set -e
  echo ">> Running git operations as $(git config --global user.name) <$(git config --global user.email)>"
  echo ">> Getting back to master (from detached-head)"
  git merge --no-edit master
  git status
  git --no-pager diff
  git add manifest.yml VERSION
  git commit -m "Updated cf-pancake to v${version}"
else
  echo ">> No update needed"
fi

popd
