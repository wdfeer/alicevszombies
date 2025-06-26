#!/bin/bash
set -e

APP_NAME="alicevszombies"
BUILD_DIR="release"
ASSETS_DIR="assets"
LICENSE_FILE="LICENSE"

# Clean up previous release
rm -rf "$BUILD_DIR"
mkdir -p "$BUILD_DIR"

echo "Building Windows binary..."
CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 \
    go build -o "$BUILD_DIR/${APP_NAME}.exe" -ldflags "-s -w"

echo "Building Linux binary..."
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
    go build -o "$BUILD_DIR/${APP_NAME}" -ldflags "-s -w"

echo "Copying assets and license..."
cp -r "$ASSETS_DIR" "$BUILD_DIR/"
cp "$LICENSE_FILE" "$BUILD_DIR/"

echo "Creating release zip..."
cd "$BUILD_DIR"
zip -r "${APP_NAME}_release.zip" "${APP_NAME}" "${APP_NAME}.exe" "$ASSETS_DIR" "$LICENSE_FILE"

echo "Release package created at: $BUILD_DIR/${APP_NAME}_release.zip"
