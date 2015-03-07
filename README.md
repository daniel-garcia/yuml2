# yuml
A yuml.me client that is written in Go. It was inspired by a 
python client for yuml, https://github.com/wandernauta/yuml .
This client compiles to a static binary with no additional
dependencies (runtimes or libraries) which makes it ideal to
integrate in to your build processes. It also supports using
private yuml services.

Installation
------------

```bash
go get github.com/daniel-garcia/yuml
```

Usage
-----

```
Usage: ./yuml [options] INPUT_FILE OUTPUT_FILE
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

```
echo '[Customer]1-0..*[Address]' | yuml - example.png
```
![Customer to Address](http://yuml.me/diagram/scruffy/class/[Customer]-%3E[Billing%20Address])




