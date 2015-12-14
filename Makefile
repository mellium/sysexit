PROJECT=mel

.PHONEY: build
build: $(PROJECT)

.PHONEY: run
run:
	GO15VENDOREXPERIMENT=1 go run ./cmd/$(PROJECT)/main.go

$(PROJECT): **/*.go
	GO15VENDOREXPERIMENT=1 go build ./cmd/$(PROJECT)


