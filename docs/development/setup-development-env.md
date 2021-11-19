# Setup development environment

This document helps you get started developing Dumbi CMD.

## Git

1. Install [Git](https://git-scm.com/downloads)

   > On Windows, the build environment depends on Git BASH that comes as a part of [Git for Windows](https://gitforwindows.org).

## Golang dev environment

1. Download and install [Go 1.17 or later](https://golang.org/doc/install#tarball).

   - Make sure that your `GOPATH` and `PATH` environment variables are configured so that go modules can be run. For example:

     **Linux and MacOS**

     ```bash
     export GOPATH=~/go
     export PATH=$PATH:$GOPATH/bin
     ```

     **Windows**

     ```cmd
     set GOPATH=%USERPROFILE%\go
     set PATH=%PATH%;%GOPATH%\bin
     ```

## Installing Make

Make needs to be installed as appropriate for your platform:

### Linux

1. Install the `build-essential` package:

   ```bash
   sudo apt-get install build-essential
   ```

### MacOS

1. In Xcode Preferences, go to the _Downloads_ tab.
2. Under _Components_ push the _Install_ button next to _Command Line Tools_.
3. After the tools are installed, ensure that the tools are switched to the new versions:

   ```bash
   sudo xcode-select -switch /Applications/Xcode.app/Contents/Developer
   ```

4. When completed, you should see `make` and other command line developer tools in `/usr/bin`.

### Windows

1. Install [MinGW w64 11.2.0](https://chocolatey.org/packages/mingw/11.2.0) with [Chocolatey](https://chocolatey.org/install):

   ```cmd
   choco install mingw --version 11.2.0
   ```

   > Versions of MinGW newer than 11.2.0 will not contain `mingw32-make.exe`, so the `--version` must be specified.

2. Create a link to `mingw32-make.exe` so that it can be invoked as `make`:

   ```cmd
   mklink C:\ProgramData\chocolatey\bin\make.exe C:\ProgramData\chocolatey\bin\mingw32-make.exe
   ```

## Pre-Commit Hooks

1. Install [pre-commit](https://pre-commit.com/#install)
2. Install the git hook scripts: `pre-commit install`
3. Run:
   ```shell
   pre-commit run --all-files  # run all hooks on all files
   pre-commit run <HOOK_ID> --all-files # run one hook on all files
   pre-commit run --files <PATH_TO_FILE>  # run all hooks on a file
   pre-commit run <HOOK_ID> --files <PATH_TO_FILE> # run one hook on a file
   ```
4. Skip:
   ```shell
   # Commit without hooks
   git commit -m <MESSAGE> --no-verify
   ```
5. Update:
   ```shell
   # Autoupdate
   pre-commit autoupdate
   ```
