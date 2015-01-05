package main

/*
#cgo pkg-config: fftw3f

#include <math.h>
#include <stdint.h>

#include <fftw3.h>

int num_magnitudes(int n, fftwf_plan p, char *bytes, float *input, fftwf_complex *output)
{
    for (int i = 0; i < n; i++) {
        input[i] = (float)(*(int32_t*)(bytes + i*4)) / UINT32_MAX;
    }

    fftwf_execute_dft_r2c(p, input, output);

    int result = 0;
    float n2 = n * n;
    for (int i = 0; i <= n / 2; i++) {
        float re = output[i][0];
        float im = output[i][1];
        float magnitude = 10.0f * log10f((re * re + im * im) / n2);

        if (magnitude > -48.0f) {
            result++;
        }
    }

    return result;
}
*/
import "C"

import (
	"fmt"
	"io"
	"os"
	"unsafe"
)

func main() {
	const n = 2048

	input := C.fftwf_alloc_real(n)
	output := C.fftwf_alloc_complex(n/2 + 1)
	p := C.fftwf_plan_dft_r2c_1d(n, input, output, C.FFTW_ESTIMATE)

	buf := make([]byte, n*4)
	result := 0

	for {
		_, err := io.ReadFull(os.Stdin, buf)
		if err != nil {
			break
		}

		result += int(C.num_magnitudes(n, p, (*C.char)(unsafe.Pointer(&buf[0])), input, output))
	}

	fmt.Printf("%d\n", result)

	C.fftwf_destroy_plan(p)
	C.fftwf_free(unsafe.Pointer(output))
	C.fftwf_free(unsafe.Pointer(input))
}
