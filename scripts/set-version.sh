#!/usr/bin/env bash
set -euo pipefail

# A script to update version strings in specified files based on the latest Git tag.
# Used in CI/CD pipelines to ensure version consistency.

replace_version() {
  local file=$1
  local search=$2
  local template=$3
  local version=$4

  if [[ ! -f $file ]]; then
    echo "Error: file not found: $file" >&2
    return 1
  fi
  if ! grep -Fq "$search" "$file"; then
    echo "Error: target text not found in $file" >&2
    return 1
  fi

  local temp
  temp=$(mktemp)
  trap 'rm -f "$temp"' RETURN

  local replaced=0
  while IFS= read -r line; do
    if (( replaced == 0 )) && [[ $line == "$search" ]]; then
      printf '%s\n' "${template//%/$version}" >>"$temp"
      replaced=1
    else
      printf '%s\n' "$line" >>"$temp"
    fi
  done <"$file"

  mv "$temp" "$file"
  trap - RETURN
  echo "Updated $file to $version"
}

main() {
  if [[ $# -ne 1 ]]; then
    echo "Usage: ${0##*/} <version>" >&2
    exit 1
  fi

  local version=$1
  if [[ $version != v* ]]; then
    echo "Error: version must start with 'v'." >&2
    exit 1
  fi

  replace_version "./internal/config/version.go" 'const Version = "v0.0.0-dev"' 'const Version = "%"' "$version"
}

main "$@"
