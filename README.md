### Getting Started

```
$ go version # should be >= 1.11
$ cd $GOPATH/src
$ git clone <this_repo>
$ cd this_repo_name

$ export GO111MODULE=on
$ echo $GO111MODULE # should be 'on'

$ go mod tidy
$ go run *.go # first run might take a while
```

Be sure to rename `config.yaml.example` to `config.yaml` and replace the configs with correct values.
