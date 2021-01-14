# GOgle-drive
A minimal, CLI based Google Drive client for Linux written in the Go programming language. Main goal is to lower power consumption while supporting full sync.

## Commands
* **gog start**: launches the gogle-drive service
* **gog auth**: (re)-authenticates and (re)-sets credentials
* **gog quit**: clean-exit the gogle-drive service

## IPC Commands
The IPC-interfaces can be used to derive state information. That can be used for populating a system bar (eg. polybar), setting a watcher or simply informational purpose. Note that it is not able to set
* **gog ipc interval**: Returns the scan interval in milliseconds.
* **gog ipc size**: Total size of the drive-folder in KB.
* **gog ipc state**: State can be IDLE (waiting for next scan), SCANNING (checking for local changes), PULLING (checking for remote changes), DLOAD (downloading remote files), ULOAD (uploading local changes), RESOLV (resolving conflicts)

## Configuration

## Installation 

## Setup
