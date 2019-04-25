# How to update the Go version

System: Debian/Ubuntu/Fedora. Might work for others as well.

### 1. Uninstall the exisiting version

As mentioned [here](https://golang.org/doc/install#install), to update a go version you will first need to uninstall the original version.

To uninstall, delete the `/usr/local/go` directory by:

```
$ sudo rm -rf /usr/local/go
```

### 2. Install the new version

Go to the [downloads](https://golang.org/dl/) page and download the binary release suitable for your system.

### 3. Extract the archive file

To extract the archive file:

```
$ sudo tar -C /usr/local -xzf /home/nikhita/Downloads/go1.8.1.linux-amd64.tar.gz
```

### 4. Make sure that your PATH contains `/usr/local/go/bin`

```
$ echo $PATH | grep "/usr/local/go/bin"
```