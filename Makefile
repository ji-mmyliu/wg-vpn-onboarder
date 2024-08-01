collectstatic:
	go-bindata -o wg-config-templates/templates.go wg-config-templates/...
	sed 's/package main/package templates/g' wg-config-templates/templates.go > wg-config-templates/templates_temp.go
	mv wg-config-templates/templates_temp.go wg-config-templates/templates.go

setupstatic:
	go install github.com/go-bindata/go-bindata/...@latest

build:
	go build -o wgv .

test:
	go test -v ./...
