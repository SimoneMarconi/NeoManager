# NeoManager

NeoManager is a version manager for NeoVim, allowing you to easily install, switch, and manage different versions of NeoVim.

## Features

- **Easy Initialization**: Quickly set up NeoManager with a single command.
- **Version Installation**: Install specific versions of NeoVim effortlessly.
- **Version Switching**: Change between installed versions of NeoVim with ease.

## Installation

To get started with NeoManager, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/neomanager.git
    cd neomanager
    ```

2. Add NeoManager to your PATH:
    ```sh
    export PATH=$PATH:/path/to/neomanager
    ```

## Usage

### Initialize NeoManager

Before you can use NeoManager, you need to initialize it. Run the following command:

    
    NeoManager init
    

This command sets up the necessary environment for NeoManager to manage your NeoVim versions.

### Install a NeoVim Version

To install a specific version of NeoVim, use the `install` command followed by the version number. For example, to install version `0.10.0`:

    
    NeoManager install 10
    

### Change NeoVim Version

To switch to a different installed version of NeoVim, use the `change` command followed by the version number. For example, to change to version `0.10.0`:

    
    NeoManager change 10

