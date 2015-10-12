PROJECT=honey

.PHONEY: build
build: $(PROJECT)

.PHONEY: run
run:
	go run honey.go

$(PROJECT): *.go **/*.go
	go build
