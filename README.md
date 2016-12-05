# Fly server

***

[![Travis Widget]][Travis] [![Report Card]][Report] [![GoDoc]][DocFly]

[Travis]: https://travis-ci.org/ivan-iver/fly
[Travis Widget]: https://travis-ci.org/ivan-iver/fly.svg?branch=master

[Report Card]: https://goreportcard.com/badge/github.com/ivan-iver/fly
[Report]: https://goreportcard.com/report/github.com/ivan-iver/fly

[GoDoc]: https://godoc.org/github.com/ivan-iver/fly?status.svg
[DocFly]: https://godoc.org/github.com/ivan-iver/fly

Is a lightweight server with markdown support. Fly is ideal to serve static HTML content on the network using a predefined layout template.

**WorkInProgress**

This project is under construction and could be not work correctly right now. If you wish to test it you can: 

* Clone it

  ```
  $ mkdir -p $GOPATH/src/github.com/ivan-iver/
  $ git clone http://github.com/ivan-iver/fly $GOPATH/src/github.com/ivan-iver/fly
  $ cd $GOPATH/src/github.com/ivan-iver/fly
  ```
  
* Compile it

  ```
  $ make build
  ```
  
* Run it

  ```
  $ cd bin/
  $ ./fly
  ```

## Content

 * [Why fly](#why)
 * [How to install](#install)
 * [Run](#run)
 * [Uninstall](#uninstall)
 * [Copyright and license](#license)

 
<a name="why"></a>

***

### There are many web server. Why use fly?

Fly is designed to use as easy and fast tool for:

 * Custom presentations.
 * Documentation sites.


<a name="install"></a>

***

### INSTALL

Add Fly and its package dependencies to your go `src` directory. If you are new on go, read about [how to install](https://golang.org/doc/install) `go`.

```
$ go get -v github.com/ivan-iver/fly
```

Once the get completes, you should find your new fly (or fly.exe from windows) executable sitting inside $GOPATH/bin/.

To update fly's dependencies, use go get with the -u option.

```
$ go get -u github.com/ivan-iver/fly
```

<a name="run"></a>

***

### RUN

After the get completes, you can run `fly`:

```
$ fly
```

<a name="uninstall"></a>

***

### UNINSTALL

One way is delete the archive files and executable binaries that go install (or go get) produces for a package something like that:

```
$ rm -rf $GOPATH/src/github.com/ivan-iver/fly
$ rm -rf $GOPATH/pkg/darwin_amd64/github.com/ivan-iver/
$ rm -rf $GOPATH/bin/fly
```

Or use `go clean -i github.com/ivan-iver/fly...` . 

Be sure to include `...` on the importpath (github.com/ivan-iver/fly), since it appears that if a package includes an executable go clean -i will only remove that and not archive files for subpackages,

See `go help clean`, `go clean` has an `-n` flag for a dry run that prints what will be run without executing it.

<a name="license"></a>

***

## Copyright and license

***

Copyright (c) 2014-2015 Iván Jaimes. See [LICENSE](LICENSE) for details.
