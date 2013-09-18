
extern "C" __global__ void
regionselect(float* __restrict__  dst, float* __restrict__ src, int8_t* regions, int8_t region, int N) {

    int i = ( blockIdx.y*gridDim.x + blockIdx.x ) * blockDim.x + threadIdx.x;
    if (i < N) {
        dst[i] = (regions[i] == region? src[i]: 0.0f);
    }
}

