
default: run

build:
	go build -o streamadd.out src/streamadd.go

run: build
	cat /dev/urandom | head -n 10 | ./streamadd.out

run-wild: build
	cat /dev/urandom | ./streamadd.out
