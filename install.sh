#!/bin/bash

PACKAGE_NAME="random-port"
REPO_URL="https://github.com/OZCAP/random-port.git"

# Define the install directory, adjust if necessary.
INSTALL_DIR="/usr/local/bin"

# Temporary directory for cloning and building
BUILD_DIR=$(mktemp -d)

# Ensure Git is installed
if ! command -v git &>/dev/null; then
	echo "Git is not installed. Please install it to proceed."
	exit 1
fi

# Clone the repository
echo "Cloning latest release..."
git clone $REPO_URL $BUILD_DIR
if [ $? -ne 0 ]; then
	echo "Failed to clone the repository."
	exit 1
fi

# Change to the repository directory
cd $BUILD_DIR

# Ensure the Go language is installed and $GOPATH/bin is in your PATH.
if ! command -v go &>/dev/null; then
	echo "Go is not installed. Please install it to proceed."
	exit 1
fi

# Build the project
echo "Building binary..."
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

# Cleanup the temporary build directory
echo "Cleaning up..."
rm -rf $BUILD_DIR
echo "Installation completed."
