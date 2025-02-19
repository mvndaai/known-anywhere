#!/bin/sh

pushd frontend
bun install
bun run build
popd
