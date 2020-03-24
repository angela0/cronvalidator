all: cronnext cronvalidate

cronnext:
	go build -o . github.com/angela0/cronvalidator/cmd/cronnext

cronvalidate:
	go build -o . github.com/angela0/cronvalidator/cmd/cronvalidate

install:
	go install github.com/angela0/cronvalidator/cmd/cronnext
	go install github.com/angela0/cronvalidator/cmd/cronvalidate

clean:
	rm -f cronvalidate cronnext
