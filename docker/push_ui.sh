#!/bin/sh

# Check if the current branch is 'main'
current_branch=$(git rev-parse --abbrev-ref HEAD)
if [ "$current_branch" != "main" ]; then
  echo "Not on main branch. Exiting."
  exit 1
fi

# Check for pending changes
if ! git diff-index --quiet HEAD --; then
  echo "There are pending changes. Exiting."
  exit 1
fi

./docker/build_frontend.sh
pushd bin/frontend || exit 1
commit_hash=$(git rev-parse --short HEAD)
git init
git add .
git commit -m "$commit_hash"
git checkout -B deployed-frontend
git remote add origin git@github.com:mvndaai/known-anywhere.git
git push -f origin deployed-frontend
popd