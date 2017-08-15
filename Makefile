VERSION!=git describe --tags --dirty
COMMIT!=git rev-parse --short HEAD 2>/dev/null
LOCAL_DEPS!=go list -f '{{ join .Imports "\n" }}' . | grep '^mellium.im/mel' | cut -c16-

LDFLAGS =-X main.commit=$(COMMIT)
LDFLAGS+=-X main.version=$(VERSION)
LDFLAGS+=-X main.defConfigFile=config.toml

mel: *.go $(LOCAL_DEPS) vendor
	go build -o $@ \
           -ldflags "$(LDFLAGS)"

vendor: Gopkg.lock Gopkg.toml
	dep ensure

deps.svg: Gopkg.lock
	(   echo "digraph G {"; \
	go list -f '{{range .Imports}}{{printf "\t%q -> %q;\n" $$.ImportPath .}}{{end}}' \
		$$(go list -f '{{join .Deps " "}}' .) .; \
	echo "}"; \
	) | dot -Tsvg -o $@
