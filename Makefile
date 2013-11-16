VERSION=v0.0.2

default: dist

dist: clean
	mkdir -p pkg
	mkdir -p pkg/linux-amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pkg/linux-amd64/nagioseasier

	mkdir -p pkg/darwin-amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o pkg/darwin-amd64/nagios-easier

	cd pkg && tar -cf nagioseasier_${VERSION}_darwin-amd64.tar.gz darwin-amd64 && cd ..
	cd pkg && tar -cf nagioseasier_${VERSION}_linux-amd64.tar.gz linux-amd64 && cd ..

clean:
	rm -rf pkg
	rm -f nagioseasier
