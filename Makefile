GO = $(shell which go 2>/dev/null)
DOCKER = $(shell which docker 2>/dev/null)

APP := papa-dew-ai
APP_VERSION 		:= v0.1.0
GCR_URL:= asia.gcr.io/refined-oven-274115

ifeq ($(GO),)
$(warning "go is not in your system PATH")
else
$(info "go found")
endif

.PHONY: all 

all: clean build

build:
	$(GO) build -o bin/$(APP) *.go
test:
	$(GO) test -count=1 ./... -v
clean:
	$(RM) -r bin
