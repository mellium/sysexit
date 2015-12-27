PROJECT=mel

.PHONEY: build
build: $(PROJECT)

.PHONEY: run
run:
	GO15VENDOREXPERIMENT=1 go run ./cmd/$(PROJECT)/main.go

$(PROJECT): *.go
	GO15VENDOREXPERIMENT=1 go build ./cmd/$(PROJECT)

.PHONEY: test
test:
	GO15VENDOREXPERIMENT=1 go test $$(go list ./... | grep -v '/vendor/')

deps.svg:
	(   echo "digraph G {"; \
	go list -f '{{range .Imports}}{{printf "\t%q -> %q;\n" $$.ImportPath .}}{{end}}' \
		$$(go list -f '{{join .Deps " "}}' .) .; \
	echo "}"; \
	) | dot -Tsvg -o $@
