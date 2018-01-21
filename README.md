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

## Run

First, make sure you have built the binary for your OS with the instructions
above.

```shell
$ ./builds/darwin/closeby -c ~/Downloads/gistfile1.tx

        CLOSE BY CUSTOMERS

        4       Ian Kehoe
        5       Nora Dempsey
        6       Theresa Enright
        8       Eoin Ahearn
        11      Richard Finnegan
        12      Christina McArdle
        13      Olive Ahearn
        15      Michael Ahearn
        17      Patricia Cahill
        23      Eoin Gallagher
        24      Rose Enright
        26      Stephen McArdle
        29      Oliver Ahearn
        30      Nick Enright
        31      Alan Behan
        39      Lisa Ahearn
```

### Flags available

- `lat` The latitude of the location you want to measure the distance from.
- `long` The longitude of the location you want to measure the distance from.
- `c` The location of the file of the list of customers in JSON format.

### Expected file content

This binary expects the file to contain a string of JSON per line with the
information of the customer, like the example below.

```json
{"latitude": "52.986375", "user_id": 12, "name": "Christina McArdle", "longitude": "-6.043701"}
{"latitude": "51.92893", "user_id": 1, "name": "Alice Cahill", "longitude": "-10.27699"}
```
