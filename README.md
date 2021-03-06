# yuml2
A yuml.me client that is written in Go (golang). It was inspired by a 
python client for yuml, https://github.com/wandernauta/yuml .
This client compiles to a static binary with no additional
dependencies (runtimes or libraries) which makes it ideal to
integrate in to your build processes. It also supports using
private yuml services.

[![GoDoc](https://godoc.org/github.com/daniel-garcia/yuml2/go-yuml?status.png)](https://godoc.org/github.com/daniel-garcia/yuml2/go-yuml)

Installation
------------

```bash
go get github.com/daniel-garcia/yuml2
```

Or you can download binaries for your platform:

  * [Mac OS X, 386](https://circle-artifacts.com/gh/daniel-garcia/yuml2/10/artifacts/0/tmp/circle-artifacts.DxLWOl2/darwin/386/yuml2)
  * [Mac OS X, amd64](https://circle-artifacts.com/gh/daniel-garcia/yuml2/10/artifacts/0/tmp/circle-artifacts.DxLWOl2/darwin/amd64/yuml2)
  * [Linux, 386](https://circle-artifacts.com/gh/daniel-garcia/yuml2/10/artifacts/0/tmp/circle-artifacts.DxLWOl2/linux/386/yuml2)
  * [Linux, amd64](https://circle-artifacts.com/gh/daniel-garcia/yuml2/10/artifacts/0/tmp/circle-artifacts.DxLWOl2/linux/amd64/yuml2)
  * [Windows, 386](https://circle-artifacts.com/gh/daniel-garcia/yuml2/10/artifacts/0/tmp/circle-artifacts.DxLWOl2/windows/386/yuml2.exe)
  * [Windows, amd64](https://circle-artifacts.com/gh/daniel-garcia/yuml2/10/artifacts/0/tmp/circle-artifacts.DxLWOl2/windows/amd64/yuml2.exe)


Usage
-----

```
Usage: yuml2 [options] INPUT_FILE OUTPUT_FILE
  -direction="LR": text direction (LR, RL, TD)
  -format="": format of output (png, pdf, jpg, svg)
  -scale=0: percentage to scale output
  -style="scruffy": style of image (scruffy, nofunky, plain)
  -t="class": type of diagram (class, activity, usecase)
  -u="http://yuml.me/diagram": base url for service [YUML_ENDPOINT]
```

You can set the YUML_ENDPOINT environment variable to point
to your own yuml server. The dash "-" character can be used to
read from stdin or write to stdout.

Examples
--------
You can make class diagrams:
```
echo '[Customer]1-0..*[Address]' | yuml2 - example.png
```
![Customer to Address](http://yuml.me/diagram/scruffy/class/[Customer]-%3E[Billing%20Address])

You can also create use case digrams:
```
curl --silent https://raw.githubusercontent.com/daniel-garcia/yuml2/master/sample_usecase.yuml | yuml2 - sample_usecase.png
```
![Sample Use Case](http://yuml.me/diagram/scruffy/usecase/%5BCustomer%5D-(Sign%20In),%20%5BCustomer%5D-(Buy%20Products),%20(Buy%20Products)%3E(Browse%20Products),%20(Buy%20Products)%3E(Checkout),%20(Checkout)%3C(Add%20New%20Credit%20Card).png)


For more examples, visit http://yuml.me

