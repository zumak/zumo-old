IMPORT_PATH:=github.com/zumak/zumo
BIN_NAME:=$(notdir $(IMPORT_PATH))

default: $(BIN_NAME)

GIT_COMMIT_ID:=$(shell git rev-parse --short HEAD)
VERSION:=$(GIT_COMMIT_ID)-$(shell date +"%Y%m%d.%H%M%S")

# if gopath not set, make inside current dir
ifeq ($(GOPATH),)
	GOPATH=$(PWD)/.GOPATH
endif

GO_SOURCES = $(shell find . -name ".GOPATH" -prune -o -type f -name '*.go' -print)
JS_SOURCES = $(shell find app/js -type f -name '*.js' -print)
HTML_SOURCES = $(shell find app/html -type f -name '*.html' -print)
CSS_SOURCES = $(shell find app/less -type f -name "*.less" -print)
WEB_LIBS = $(shell find app/lib -type f -type f -print)
_CSS_ENTRY_ = $(shell find app/less/entry -type f -name "*.less" -print)

DISTS += $(HTML_SOURCES:app/html/%=dist/html/%)
DISTS += $(JS_SOURCES:app/js/%=dist/js/%)
DISTS += $(_CSS_ENTRY_:app/less/entry/%.less=dist/css/%.css)
DISTS += $(WEB_LIBS:app/lib/%=dist/lib/%)

# Automatic runner
DIRS = $(shell find . \
	   -name ".git" -prune -o \
	   -name ".GOPATH" -prune -o \
	   -name "vendor" -prune -o \
	   -type d -print)

.sources:
	@echo $(DIRS) makefile \
		$(GO_SOURCES) \
		$(JS_SOURCES) \
		$(HTML_SOURCES) \
		$(CSS_SOURCES) \
		$(WEB_LIBS)| tr " " "\n"
run: $(BIN_NAME)
	#go test ./backend/...
	./$(BIN_NAME) #--config ../zumo-config.yaml
auto-run:
	while true; do \
		make .sources | entr -rd make run ; \
		echo "hit ^C again to quit" && sleep 1 ; \
	done
reset:
	ps -e | grep make | grep -v grep | awk '{print $$1}' | xargs kill

## Binary build
$(BIN_NAME).bin: $(GOPATH)/src/$(IMPORT_PATH) $(GO_SOURCES)
	go get -v -d $(IMPORT_PATH)            # can replace with glide
	go build -v \
		-ldflags "-X main.VERSION=$(VERSION)" \
		-ldflags "-extldflags -static" \
		-o $(BIN_NAME).bin ./cli
	@echo Build DONE

$(BIN_NAME): $(BIN_NAME).bin $(DISTS)
	cp $(BIN_NAME).bin $(BIN_NAME).tmp
	rice append -v --exec $(BIN_NAME).tmp \
		-i $(IMPORT_PATH)/http-server
	mv $(BIN_NAME).tmp $(BIN_NAME)
	@echo Embed resources DONE

## Web dist
dist/css/%.css: $(CSS_SOURCES)
	lessc app/less/entry/$*.less $@
dist/%: app/%
	@mkdir -p $(basename $@)
	cp $< $@

tools:
	npm install -g less
	go get github.com/GeertJohan/go.rice/rice
clean:
	rm -rf dist/ vendor/ $(BIN_NAME) $(BIN_NAME).bin $(BIN_NAME).tmp
	go clean

$(GOPATH)/src/$(IMPORT_PATH):
	@echo "make symbolic link on $(GOPATH)/src/$(IMPORT_PATH)..."
	@mkdir -p $(dir $(GOPATH)/src/$(IMPORT_PATH))
	ln -s $(PWD) $(GOPATH)/src/$(IMPORT_PATH)

## Multi platform
deploy: build/linux/amd64/$(BIN_NAME)
deploy: build/linux/arm/$(BIN_NAME)
deploy: build/windows/amd64/$(BIN_NAME)
#deploy: build/windows/arm/$(BIN_NAME)
# make hook.mk file for your hook (example. following lines)
deploy:
	# TODO scp or upload binary
	# TODO call hook to deploy(ex. docker command)

build/%/$(BIN_NAME): export GOOS=$(subst /,,$(dir $*))
build/%/$(BIN_NAME): export GOARCH=$(notdir $*)
build/%/$(BIN_NAME):
	@echo --------------------------BUILD $$GOOS $$GOARCH-----------------------------
	make clean
	make $(BIN_NAME)
	mkdir -p $(@D)
	mv $(BIN_NAME) $@

.PHONY: .sources run auto-run reset tools clean
