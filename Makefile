PROJECT=honey

.PHONEY: build
build: $(PROJECT)

.PHONEY: run
run:
	go run $(PROJECT).go

$(PROJECT): *.go **/*.go
	go build


