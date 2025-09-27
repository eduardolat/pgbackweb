#!/usr/bin/env bash
set -euo pipefail

# A script to update version strings in specified files based on the latest Git tag.
# Used in CI/CD pipelines to ensure version consistency.

main() {
  local version
  if ! version=$(git describe --tags --abbrev=0 2>/dev/null); then
    echo "Error: no git tags found." >&2
    exit 1
  fi
  [[ $version == v* ]] || version="v$version"

  local file="./internal/config/version.go"
  local search='const Version = "v0.0.0-dev"'
  local replacement="const Version = \"$version\""

  if [[ ! -f $file ]]; then
    echo "Set version error: file not found: $file" >&2
    exit 1
  fi
  if ! grep -Fq "$search" "$file"; then
    echo "Set version error: target text not found in $file" >&2
    exit 1
  fi

  local temp
  temp=$(mktemp)
  trap 'rm -f "$temp"' EXIT

  local replaced=0
  while IFS= read -r line; do
    if (( replaced == 0 )) && [[ $line == "$search" ]]; then
      printf '%s\n' "$replacement" >>"$temp"
      replaced=1
    else
      printf '%s\n' "$line" >>"$temp"
    fi
  done <"$file"

  mv "$temp" "$file"
  trap - EXIT
  echo "Updated $file to $version"
}

main "$@"
