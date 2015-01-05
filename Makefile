CC = cc
CFLAGS = -O2 -std=c99 -static $(shell pkg-config --cflags fftw3f)
LDFLAGS = $(shell pkg-config --static --libs fftw3f)

all: samples.bin spek-c spek-go

clean:
	rm -f samples.bin spek-c spek-go

samples.bin:
	dd if=/dev/urandom of=samples.bin bs=256 count=2480625

spek-c: spek-lang.c
	$(CC) $(CFLAGS) -o spek-c spek-lang.c $(LDFLAGS)

spek-go: spek-lang.go
	go fmt spek-lang.go && go build --ldflags '-extldflags "-static"' -o spek-go spek-lang.go

run:
	time ./spek-c < samples.bin
	time ./spek-go < samples.bin

.PHONY: clean run
