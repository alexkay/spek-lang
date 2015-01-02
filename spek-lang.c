#include <math.h>
#include <stdint.h>
#include <unistd.h>

#include <fftw3.h>

int main()
{
    const int n = 2048;

    float *input = fftwf_alloc_real(n);
    int32_t *ints = (int32_t*)input;
    fftwf_complex *output = fftwf_alloc_complex(n / 2 + 1);
    fftwf_plan p = fftwf_plan_dft_r2c_1d(n, input, output, FFTW_ESTIMATE);
    const float n2 = n * n;

    int result = 0;

    while (read(STDIN_FILENO, input, n * sizeof(float)) == n * sizeof(float)) {
        for (int i = 0; i < n; i++) {
            input[i] = ((float)(int32_t)ints[i]) / UINT32_MAX;
        }

        fftwf_execute_dft_r2c(p, input, output);

        for (int i = 0; i <= n / 2; i++) {
            float re = output[i][0];
            float im = output[i][1];
            float magnitude = 10.0f * log10f((re * re + im * im) / n2);

            if (magnitude > -48.0f) {
                result++;
            }
        }
    }

    printf("%d\n", result);

    fftwf_destroy_plan(p);
    fftwf_free(output);
    fftwf_free(input);

    return 0;
}
