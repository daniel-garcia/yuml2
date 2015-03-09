

PDIR=/go/src/github.com/daniel-garcia/yuml2
DOCKER_IMAGE=golang:1.3-cross

CIRCLE_ARTIFACTS?=output
OUTPUT=$(CIRCLE_ARTIFACTS)

binaries: $(OUTPUT)/windows/386/yuml2.exe $(OUTPUT)/windows/amd64/yuml2.exe $(OUTPUT)/linux/386/yuml2 $(OUTPUT)/linux/amd64/yuml2 $(OUTPUT)/darwin/386/yuml2 $(OUTPUT)/darwin/amd64/yuml2


$(OUTPUT)/windows/386/yuml2.exe:
	mkdir -p $(OUTPUT)/windows/386
	docker run --rm -v "$(PWD)":$(PDIR) -w $(PDIR) -e GOOS=windows -e GOARCH=386 $(DOCKER_IMAGE) go build -o $@ -v

$(OUTPUT)/windows/amd64/yuml2.exe:
	mkdir -p $(OUTPUT)/windows/amd64
	docker run --rm -v "$(PWD)":$(PDIR) -w $(PDIR) -e GOOS=windows -e GOARCH=amd64 $(DOCKER_IMAGE) go build -o $@ -v

$(OUTPUT)/linux/386/yuml2:
	mkdir -p $(OUTPUT)/linux/386
	docker run --rm -v "$(PWD)":$(PDIR) -w $(PDIR) -e GOOS=linux -e GOARCH=386 $(DOCKER_IMAGE) go build -o $@ -v

$(OUTPUT)/linux/amd64/yuml2:
	mkdir -p $(OUTPUT)/linux/amd64
	docker run --rm -v "$(PWD)":$(PDIR) -w $(PDIR) -e GOOS=linux -e GOARCH=amd64 $(DOCKER_IMAGE) go build -o $@ -v


$(OUTPUT)/darwin/386/yuml2:
	mkdir -p $(OUTPUT)/darwin/386
	docker run --rm -v "$(PWD)":$(PDIR) -w $(PDIR) -e GOOS=darwin -e GOARCH=386 $(DOCKER_IMAGE) go build -o $@ -v

$(OUTPUT)/darwin/amd64/yuml2:
	mkdir -p $(OUTPUT)/darwin/amd64
	docker run --rm -v "$(PWD)":$(PDIR) -w $(PDIR) -e GOOS=darwin -e GOARCH=amd64 $(DOCKER_IMAGE) go build -o $@ -v




clean:
	rm -Rf $(OUTPUT)
	go clean
