collectstatic:
	go-bindata -o templates/templates.go templates/...
	sed -i 's/package main/package templates/g' templates/templates.go

build:
	go build .
	mv wg-vpn-onboarder wgv
