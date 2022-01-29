RELEASES=bin/srv-yy-darwin-amd64 \
	 bin/srv-yy-linux-amd64

all: $(RELEASES)

bin/srv-yy-%: GOOS=$(firstword $(subst -, ,$*))
bin/srv-yy-%: GOARCH=$(subst .exe,,$(word 2,$(subst -, ,$*)))
bin/srv-yy-%: $(wildcard *.go)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build \
	     -ldflags "-X main.osarch=$(GOOS)/$(GOARCH) -X main.gitVersion=`git rev-parse HEAD` -X 'main.buildDate=`date`' -s -w" \
	     -buildmode=exe \
	     -tags release \
	     -o $@

clean:
	rm -rf bin
