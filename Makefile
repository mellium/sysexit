PROJECT=honey

.PHONEY: build
build: $(PROJECT)

.PHONEY: run
run:
	GO15VENDOREXPERIMENT=1 go run $(PROJECT).go

$(PROJECT): *.go **/*.go
	GO15VENDOREXPERIMENT=1 go build


