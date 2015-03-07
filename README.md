# yuml
A yuml.me client that is written in Go.

Inspired by a python client for yuml, https://github.com/wandernauta/yuml

Installation
------------

```bash
go get github.com/daniel-garcia/yuml
```

Usage
-----

```
Usage: ./yuml [INPUT_FILE] [OUTPUT_FILE]
  -direction="LR": text direction (LR, RL, TD)
  -format="": format of output (png, pdf, jpg, svg
  -scale=0: percentage to scale output
  -style="scruffy": style of image (scruffy, nofunky, plain)
  -t="class": type of diagram (class, activity, usecase)
  -u="http://yuml.me/diagram": base url for service
```

