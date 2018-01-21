# closeby

Find places within the 100km range using the orthodromic distance.

## Install

**Prerequisite**
- Docker
- make

### Linux

```shell
$ make linux
$ ./builds/linux/closeby

USAGE
  ./builds/linux/closeby -c $FILELOCATION

FLAGS
  -c               Location of the list of customers.
  -lant 53.339428  Latitude of the location.
  -long -6.257664  Longitude of the location.
```

### MacOS/Darwin

```shell
$ make darwin
$ ./builds/darwin/closeby

USAGE
  ./builds/darwin/closeby -c $FILELOCATION

FLAGS
  -c               Location of the list of customers.
  -lant 53.339428  Latitude of the location.
  -long -6.257664  Longitude of the location.

```

### Windows

```shell
$ make windows
$ ./builds/windows/closeby.exe

USAGE
  ./builds/windows/closeby.exe -c $FILELOCATION

FLAGS
  -c               Location of the list of customers.
  -lant 53.339428  Latitude of the location.
  -long -6.257664  Longitude of the location.
```
