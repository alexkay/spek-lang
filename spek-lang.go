package main

/*
#cgo pkg-config: fftw3f
#include <fftw3.h>
*/
import "C"

import (
	"fmt"
	"io"
	"math"
	"os"
	"unsafe"
)

func main() {
	const n = 2048

	input := C.fftwf_alloc_real(n)
	floats := (*[1 << 30]float32)(unsafe.Pointer(input))[:n:n]

	output := C.fftwf_alloc_complex(n/2 + 1)
	complex := (*[1 << 30]float32)(unsafe.Pointer(output))[:(n/2+1)*2 : (n/2+1)*2]
	p := C.fftwf_plan_dft_r2c_1d(n, input, output, C.FFTW_ESTIMATE)
	const n2 = n * n

	buf := make([]byte, n*4)
	ints := (*[1 << 30]int32)(unsafe.Pointer(&buf[0]))[:n:n]
	result := 0

	for {
		_, err := io.ReadFull(os.Stdin, buf)
		if err != nil {
			break
		}
		for i := 0; i < n; i++ {
			floats[i] = float32(ints[i]) / math.MaxUint32
		}

		C.fftwf_execute_dft_r2c(p, input, output)

		for i := 0; i <= n/2; i++ {
			re := complex[i*2]
			im := complex[i*2+1]
			magnitude := float32(10.0 * math.Log10(float64((re*re+im*im)/n2)))

			if magnitude > -48.0 {
				result++
			}
		}
	}

	fmt.Printf("%d\n", result)

	C.fftwf_destroy_plan(p)
	C.fftwf_free(unsafe.Pointer(output))
	C.fftwf_free(unsafe.Pointer(input))
}
