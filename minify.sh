#!/bin/zsh

# Delete folders and their contents
rm -rf client/_components
rm -rf client/node_modules

# Delete files with .svelte extension
find . -name "*.svelte" -type f -delete

# Delete specific files
rm -f client/build.js
rm -f client/package-lock.json
rm -f client/package.json
rm -f client/_layout.html
rm -f client/pnpm-lock.yaml
rm -f client/tailwind.config.js
rm -f client/README.md
rm -f client/svelte.config.js