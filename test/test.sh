#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/.."

# Load test environment variables
if [ -f ".env.test" ]; then
  export $(grep -v '^#' .env.test | xargs)
else
  echo ".env.test file not found!"
  exit 1
fi

# sleep for connection db
sleep 3

# Run Go integration tests
echo "Running integration tests..."
go test -v ./test/integration/...

echo "Tests completed."
