# gofork

`gofork` is tool, which is used to help you to easy to setup golang project.

[![asciicast](https://asciinema.org/a/221108.svg)](https://asciinema.org/a/221108)

## Get started

When you try to contribute golang project, basically the following steps will
be used.

```
$ mkdir -p $GOPATH/src/github.com/<upstream-user>/<project>
$ git clone https://github.com/<your-account>/<project> \
    $GOPATH/src/github.com/<upstream-user>/<project>
$ cd $GOPATH/src/github.com/<upstream-user>/<project>
$ git remote add upstream https://github.com/<upstream-user>/<project>.git
```

In order to combine the steps, the `gofork` shows up.

You can use the following command to setup a golang project which you will
contribute to. It is very easier.

```
$ gofork clone \
  --upstream https://github.com/<upstream-user>/<project> \
  github.com/<upstream-user>/<project> \
  https://github.com/<your-account>/<project>
```

## Install

Requirement is >= golang1.11 which supports go modules.

```
$ git clone https://github.com/fuweid/gofork
$ cd gofork
$ make
$ sudo make install # by default, /usr/bin/gofork
```

## TODO

* [ ] allow to set `no-pushing` for upstream remote
* [ ] allow to set upstream remote url without clone
