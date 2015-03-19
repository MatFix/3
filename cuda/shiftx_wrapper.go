package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"github.com/mumax/3/timer"
	"sync"
	"unsafe"
)

// CUDA handle for shiftx kernel
var shiftx_code cu.Function

// Stores the arguments for shiftx kernel invocation
type shiftx_args_t struct {
	arg_dst    unsafe.Pointer
	arg_src    unsafe.Pointer
	arg_Nx     int
	arg_Ny     int
	arg_Nz     int
	arg_shx    int
	arg_clampL float32
	arg_clampR float32
	argptr     [8]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for shiftx kernel invocation
var shiftx_args shiftx_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	shiftx_args.argptr[0] = unsafe.Pointer(&shiftx_args.arg_dst)
	shiftx_args.argptr[1] = unsafe.Pointer(&shiftx_args.arg_src)
	shiftx_args.argptr[2] = unsafe.Pointer(&shiftx_args.arg_Nx)
	shiftx_args.argptr[3] = unsafe.Pointer(&shiftx_args.arg_Ny)
	shiftx_args.argptr[4] = unsafe.Pointer(&shiftx_args.arg_Nz)
	shiftx_args.argptr[5] = unsafe.Pointer(&shiftx_args.arg_shx)
	shiftx_args.argptr[6] = unsafe.Pointer(&shiftx_args.arg_clampL)
	shiftx_args.argptr[7] = unsafe.Pointer(&shiftx_args.arg_clampR)
}

// Wrapper for shiftx CUDA kernel, asynchronous.
func k_shiftx_async(dst unsafe.Pointer, src unsafe.Pointer, Nx int, Ny int, Nz int, shx int, clampL float32, clampR float32, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("shiftx")
	}

	shiftx_args.Lock()
	defer shiftx_args.Unlock()

	if shiftx_code == 0 {
		shiftx_code = fatbinLoad(shiftx_map, "shiftx")
	}

	shiftx_args.arg_dst = dst
	shiftx_args.arg_src = src
	shiftx_args.arg_Nx = Nx
	shiftx_args.arg_Ny = Ny
	shiftx_args.arg_Nz = Nz
	shiftx_args.arg_shx = shx
	shiftx_args.arg_clampL = clampL
	shiftx_args.arg_clampR = clampR

	args := shiftx_args.argptr[:]
	cu.LaunchKernel(shiftx_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("shiftx")
	}
}

// maps compute capability on PTX code for shiftx kernel.
var shiftx_map = map[int]string{0: "",
	20: shiftx_ptx_20,
	30: shiftx_ptx_30,
	35: shiftx_ptx_35,
	50: shiftx_ptx_50,
	52: shiftx_ptx_52}

// shiftx PTX code for various compute capabilities.
const (
	shiftx_ptx_20 = `
.version 4.1
.target sm_20
.address_size 64


.visible .entry shiftx(
	.param .u64 shiftx_param_0,
	.param .u64 shiftx_param_1,
	.param .u32 shiftx_param_2,
	.param .u32 shiftx_param_3,
	.param .u32 shiftx_param_4,
	.param .u32 shiftx_param_5,
	.param .f32 shiftx_param_6,
	.param .f32 shiftx_param_7
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<22>;
	.reg .f32 	%f<6>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftx_param_0];
	ld.param.u64 	%rd2, [shiftx_param_1];
	ld.param.u32 	%r5, [shiftx_param_2];
	ld.param.u32 	%r6, [shiftx_param_3];
	ld.param.u32 	%r8, [shiftx_param_4];
	ld.param.u32 	%r7, [shiftx_param_5];
	ld.param.f32 	%f3, [shiftx_param_6];
	ld.param.f32 	%f4, [shiftx_param_7];
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r11, %tid.x;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	mov.u32 	%r12, %ntid.y;
	mov.u32 	%r13, %ctaid.y;
	mov.u32 	%r14, %tid.y;
	mad.lo.s32 	%r2, %r12, %r13, %r14;
	mov.u32 	%r15, %ntid.z;
	mov.u32 	%r16, %ctaid.z;
	mov.u32 	%r17, %tid.z;
	mad.lo.s32 	%r3, %r15, %r16, %r17;
	setp.lt.s32	%p1, %r1, %r5;
	setp.lt.s32	%p2, %r2, %r6;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r8;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_5;
	bra.uni 	BB0_1;

BB0_1:
	sub.s32 	%r4, %r1, %r7;
	setp.lt.s32	%p6, %r4, 0;
	mov.f32 	%f5, %f3;
	@%p6 bra 	BB0_4;

	setp.ge.s32	%p7, %r4, %r5;
	mov.f32 	%f5, %f4;
	@%p7 bra 	BB0_4;

	cvta.to.global.u64 	%rd3, %rd2;
	mad.lo.s32 	%r18, %r3, %r6, %r2;
	mad.lo.s32 	%r19, %r18, %r5, %r4;
	mul.wide.s32 	%rd4, %r19, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.f32 	%f5, [%rd5];

BB0_4:
	cvta.to.global.u64 	%rd6, %rd1;
	mad.lo.s32 	%r20, %r3, %r6, %r2;
	mad.lo.s32 	%r21, %r20, %r5, %r1;
	mul.wide.s32 	%rd7, %r21, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f5;

BB0_5:
	ret;
}


`
	shiftx_ptx_30 = `
.version 4.1
.target sm_30
.address_size 64


.visible .entry shiftx(
	.param .u64 shiftx_param_0,
	.param .u64 shiftx_param_1,
	.param .u32 shiftx_param_2,
	.param .u32 shiftx_param_3,
	.param .u32 shiftx_param_4,
	.param .u32 shiftx_param_5,
	.param .f32 shiftx_param_6,
	.param .f32 shiftx_param_7
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<22>;
	.reg .f32 	%f<6>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftx_param_0];
	ld.param.u64 	%rd2, [shiftx_param_1];
	ld.param.u32 	%r5, [shiftx_param_2];
	ld.param.u32 	%r6, [shiftx_param_3];
	ld.param.u32 	%r8, [shiftx_param_4];
	ld.param.u32 	%r7, [shiftx_param_5];
	ld.param.f32 	%f3, [shiftx_param_6];
	ld.param.f32 	%f4, [shiftx_param_7];
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r11, %tid.x;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	mov.u32 	%r12, %ntid.y;
	mov.u32 	%r13, %ctaid.y;
	mov.u32 	%r14, %tid.y;
	mad.lo.s32 	%r2, %r12, %r13, %r14;
	mov.u32 	%r15, %ntid.z;
	mov.u32 	%r16, %ctaid.z;
	mov.u32 	%r17, %tid.z;
	mad.lo.s32 	%r3, %r15, %r16, %r17;
	setp.lt.s32	%p1, %r1, %r5;
	setp.lt.s32	%p2, %r2, %r6;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r8;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_5;
	bra.uni 	BB0_1;

BB0_1:
	sub.s32 	%r4, %r1, %r7;
	setp.lt.s32	%p6, %r4, 0;
	mov.f32 	%f5, %f3;
	@%p6 bra 	BB0_4;

	setp.ge.s32	%p7, %r4, %r5;
	mov.f32 	%f5, %f4;
	@%p7 bra 	BB0_4;

	cvta.to.global.u64 	%rd3, %rd2;
	mad.lo.s32 	%r18, %r3, %r6, %r2;
	mad.lo.s32 	%r19, %r18, %r5, %r4;
	mul.wide.s32 	%rd4, %r19, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.f32 	%f5, [%rd5];

BB0_4:
	cvta.to.global.u64 	%rd6, %rd1;
	mad.lo.s32 	%r20, %r3, %r6, %r2;
	mad.lo.s32 	%r21, %r20, %r5, %r1;
	mul.wide.s32 	%rd7, %r21, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f5;

BB0_5:
	ret;
}


`
	shiftx_ptx_35 = `
.version 4.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.visible .entry shiftx(
	.param .u64 shiftx_param_0,
	.param .u64 shiftx_param_1,
	.param .u32 shiftx_param_2,
	.param .u32 shiftx_param_3,
	.param .u32 shiftx_param_4,
	.param .u32 shiftx_param_5,
	.param .f32 shiftx_param_6,
	.param .f32 shiftx_param_7
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<22>;
	.reg .f32 	%f<6>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftx_param_0];
	ld.param.u64 	%rd2, [shiftx_param_1];
	ld.param.u32 	%r5, [shiftx_param_2];
	ld.param.u32 	%r6, [shiftx_param_3];
	ld.param.u32 	%r8, [shiftx_param_4];
	ld.param.u32 	%r7, [shiftx_param_5];
	ld.param.f32 	%f3, [shiftx_param_6];
	ld.param.f32 	%f4, [shiftx_param_7];
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r11, %tid.x;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	mov.u32 	%r12, %ntid.y;
	mov.u32 	%r13, %ctaid.y;
	mov.u32 	%r14, %tid.y;
	mad.lo.s32 	%r2, %r12, %r13, %r14;
	mov.u32 	%r15, %ntid.z;
	mov.u32 	%r16, %ctaid.z;
	mov.u32 	%r17, %tid.z;
	mad.lo.s32 	%r3, %r15, %r16, %r17;
	setp.lt.s32	%p1, %r1, %r5;
	setp.lt.s32	%p2, %r2, %r6;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r8;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB5_5;
	bra.uni 	BB5_1;

BB5_1:
	sub.s32 	%r4, %r1, %r7;
	setp.lt.s32	%p6, %r4, 0;
	mov.f32 	%f5, %f3;
	@%p6 bra 	BB5_4;

	setp.ge.s32	%p7, %r4, %r5;
	mov.f32 	%f5, %f4;
	@%p7 bra 	BB5_4;

	cvta.to.global.u64 	%rd3, %rd2;
	mad.lo.s32 	%r18, %r3, %r6, %r2;
	mad.lo.s32 	%r19, %r18, %r5, %r4;
	mul.wide.s32 	%rd4, %r19, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.f32 	%f5, [%rd5];

BB5_4:
	cvta.to.global.u64 	%rd6, %rd1;
	mad.lo.s32 	%r20, %r3, %r6, %r2;
	mad.lo.s32 	%r21, %r20, %r5, %r1;
	mul.wide.s32 	%rd7, %r21, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f5;

BB5_5:
	ret;
}


`
	shiftx_ptx_50 = `
.version 4.1
.target sm_50
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.visible .entry shiftx(
	.param .u64 shiftx_param_0,
	.param .u64 shiftx_param_1,
	.param .u32 shiftx_param_2,
	.param .u32 shiftx_param_3,
	.param .u32 shiftx_param_4,
	.param .u32 shiftx_param_5,
	.param .f32 shiftx_param_6,
	.param .f32 shiftx_param_7
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<22>;
	.reg .f32 	%f<6>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftx_param_0];
	ld.param.u64 	%rd2, [shiftx_param_1];
	ld.param.u32 	%r5, [shiftx_param_2];
	ld.param.u32 	%r6, [shiftx_param_3];
	ld.param.u32 	%r8, [shiftx_param_4];
	ld.param.u32 	%r7, [shiftx_param_5];
	ld.param.f32 	%f3, [shiftx_param_6];
	ld.param.f32 	%f4, [shiftx_param_7];
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r11, %tid.x;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	mov.u32 	%r12, %ntid.y;
	mov.u32 	%r13, %ctaid.y;
	mov.u32 	%r14, %tid.y;
	mad.lo.s32 	%r2, %r12, %r13, %r14;
	mov.u32 	%r15, %ntid.z;
	mov.u32 	%r16, %ctaid.z;
	mov.u32 	%r17, %tid.z;
	mad.lo.s32 	%r3, %r15, %r16, %r17;
	setp.lt.s32	%p1, %r1, %r5;
	setp.lt.s32	%p2, %r2, %r6;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r8;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB5_5;
	bra.uni 	BB5_1;

BB5_1:
	sub.s32 	%r4, %r1, %r7;
	setp.lt.s32	%p6, %r4, 0;
	mov.f32 	%f5, %f3;
	@%p6 bra 	BB5_4;

	setp.ge.s32	%p7, %r4, %r5;
	mov.f32 	%f5, %f4;
	@%p7 bra 	BB5_4;

	cvta.to.global.u64 	%rd3, %rd2;
	mad.lo.s32 	%r18, %r3, %r6, %r2;
	mad.lo.s32 	%r19, %r18, %r5, %r4;
	mul.wide.s32 	%rd4, %r19, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.f32 	%f5, [%rd5];

BB5_4:
	cvta.to.global.u64 	%rd6, %rd1;
	mad.lo.s32 	%r20, %r3, %r6, %r2;
	mad.lo.s32 	%r21, %r20, %r5, %r1;
	mul.wide.s32 	%rd7, %r21, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f5;

BB5_5:
	ret;
}


`
	shiftx_ptx_52 = `
.version 4.2
.target sm_52
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	shiftx
.visible .entry shiftx(
	.param .u64 shiftx_param_0,
	.param .u64 shiftx_param_1,
	.param .u32 shiftx_param_2,
	.param .u32 shiftx_param_3,
	.param .u32 shiftx_param_4,
	.param .u32 shiftx_param_5,
	.param .f32 shiftx_param_6,
	.param .f32 shiftx_param_7
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<6>;
	.reg .s32 	%r<22>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftx_param_0];
	ld.param.u64 	%rd2, [shiftx_param_1];
	ld.param.u32 	%r5, [shiftx_param_2];
	ld.param.u32 	%r6, [shiftx_param_3];
	ld.param.u32 	%r8, [shiftx_param_4];
	ld.param.u32 	%r7, [shiftx_param_5];
	ld.param.f32 	%f3, [shiftx_param_6];
	ld.param.f32 	%f4, [shiftx_param_7];
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r11, %tid.x;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	mov.u32 	%r12, %ntid.y;
	mov.u32 	%r13, %ctaid.y;
	mov.u32 	%r14, %tid.y;
	mad.lo.s32 	%r2, %r12, %r13, %r14;
	mov.u32 	%r15, %ntid.z;
	mov.u32 	%r16, %ctaid.z;
	mov.u32 	%r17, %tid.z;
	mad.lo.s32 	%r3, %r15, %r16, %r17;
	setp.lt.s32	%p1, %r1, %r5;
	setp.lt.s32	%p2, %r2, %r6;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r8;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB6_5;
	bra.uni 	BB6_1;

BB6_1:
	sub.s32 	%r4, %r1, %r7;
	setp.lt.s32	%p6, %r4, 0;
	mov.f32 	%f5, %f3;
	@%p6 bra 	BB6_4;

	setp.ge.s32	%p7, %r4, %r5;
	mov.f32 	%f5, %f4;
	@%p7 bra 	BB6_4;

	cvta.to.global.u64 	%rd3, %rd2;
	mad.lo.s32 	%r18, %r3, %r6, %r2;
	mad.lo.s32 	%r19, %r18, %r5, %r4;
	mul.wide.s32 	%rd4, %r19, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.f32 	%f5, [%rd5];

BB6_4:
	cvta.to.global.u64 	%rd6, %rd1;
	mad.lo.s32 	%r20, %r3, %r6, %r2;
	mad.lo.s32 	%r21, %r20, %r5, %r1;
	mul.wide.s32 	%rd7, %r21, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f5;

BB6_5:
	ret;
}


`
)
