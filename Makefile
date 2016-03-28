help:
	@echo "build           - builds the things"
	@echo "deps            - installs the deps (with glide)"
	@echo "clean           - deletes the binary"

build:
	GO15VENDOREXPERIMENT=1 go build -o hap main.go

deps:
	glide install

clean:
	rm -rf hap
