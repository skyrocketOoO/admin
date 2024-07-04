#!/bin/bash

function build() {
  rm -rf document out
  buf format -w
  buf lint
  buf generate
}

function main() {
  local call="${1:-}"

  if [[ -z $(typeset -F "${call}") ]]; then
    echo "No find command: ${call}"
    return
  fi

  shift
  ${call} "$@"
}

main "$@"
