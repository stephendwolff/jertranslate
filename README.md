# README

Golang application to retrieve multiple translations of lines of text, through chains of different languages.

All stored in local MySQL database 


## Set up Google API key

```bash
$ export GOOGLE_APPLICATION_CREDENTIALS="<path to secret key>.json"
```
ie:
```bash
$ export GOOGLE_APPLICATION_CREDENTIALS="`pwd`/Mr Gee-cdeb2e5979ac.json"
```


## Set up $GOPATH
```bash
$  export GOPATH=`pwd`
```

## Build (using makefile)
```bash
$ make var check all
```


## Build (for server using makefile)
```bash
$ make var check deploy
```