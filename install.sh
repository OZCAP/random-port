#!/bin/bash

PACKAGE_NAME="random-port"

# Define the install directory, adjust if necessary.
INSTALL_DIR="/usr/local/bin"

# Ensure the Go language is installed and $GOPATH/bin is in your PATH.
if ! command -v go &> /dev/null; then
  echo "Go is not installed. Please install it to proceed."
  exit 1
fi

# Build the project
echo "Building the $PACKAGE_NAME binary..."
go build -o $PACKAGE_NAME
if [ $? -ne 0 ]; then
  echo "Failed to build the $PACKAGE_NAME binary."
  exit 1
fi

# Move the binary to the install directory
echo "Installing the $PACKAGE_NAME binary to $INSTALL_DIR..."
if sudo mv $PACKAGE_NAME "$INSTALL_DIR"; then
  echo "The $PACKAGE_NAME binary has been installed successfully."
else
  echo "Failed to move the $PACKAGE_NAME binary to $INSTALL_DIR."
  exit 1
fi

# Ensure the binary is executable
sudo chmod +x "$INSTALL_DIR/$PACKAGE_NAME"

# Check if the install directory is in the PATH
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
  echo "NOTE: $INSTALL_DIR is not in your PATH. Please add it to your PATH to run '$PACKAGE_NAME' from anywhere."
fi
