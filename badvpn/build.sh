#!/bin/bash
set -eu

# Build badvpn with UDPGW
# Usage: ./build.sh

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
SRC_DIR="$SCRIPT_DIR/src"

if [ ! -d "$SRC_DIR" ]; then
  echo "Cloning badvpn repository into $SRC_DIR..."
  git clone https://github.com/ambrop72/badvpn.git "$SRC_DIR"
fi

cd "$SRC_DIR"
mkdir -p build
cd build

cmake .. -DBUILD_NOTHING_BY_DEFAULT=1 -DBUILD_UDPGW=1
make -j$(nproc)

echo "Build completed. Binary is at $SRC_DIR/build/badvpn-udpgw"