CC = cc
CFLAGS = -O2 -std=c99 -static $(shell pkg-config --cflags fftw3f)
LDFLAGS = $(shell pkg-config --static --libs fftw3f)

all: samples.bin spek-c

clean:
	rm -f spek-c

samples.bin:
	dd if=/dev/urandom of=samples.bin bs=256 count=2480625

spek-c: spek-lang.c
	$(CC) $(CFLAGS) -o spek-c spek-lang.c $(LDFLAGS)

.PHONY: clean
