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

// CUDA handle for adduniaxialanisotropy kernel
var adduniaxialanisotropy_code cu.Function

// Stores the arguments for adduniaxialanisotropy kernel invocation
type adduniaxialanisotropy_args_t struct {
	arg_Bx      unsafe.Pointer
	arg_By      unsafe.Pointer
	arg_Bz      unsafe.Pointer
	arg_mx      unsafe.Pointer
	arg_my      unsafe.Pointer
	arg_mz      unsafe.Pointer
	arg_K1LUT   unsafe.Pointer
	arg_K2LUT   unsafe.Pointer
	arg_uxLUT   unsafe.Pointer
	arg_uyLUT   unsafe.Pointer
	arg_uzLUT   unsafe.Pointer
	arg_regions unsafe.Pointer
	arg_N       int
	argptr      [13]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for adduniaxialanisotropy kernel invocation
var adduniaxialanisotropy_args adduniaxialanisotropy_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	adduniaxialanisotropy_args.argptr[0] = unsafe.Pointer(&adduniaxialanisotropy_args.arg_Bx)
	adduniaxialanisotropy_args.argptr[1] = unsafe.Pointer(&adduniaxialanisotropy_args.arg_By)
	adduniaxialanisotropy_args.argptr[2] = unsafe.Pointer(&adduniaxialanisotropy_args.arg_Bz)
	adduniaxialanisotropy_args.argptr[3] = unsafe.Pointer(&adduniaxialanisotropy_args.arg_mx)
	adduniaxialanisotropy_args.argptr[4] = unsafe.Pointer(&adduniaxialanisotropy_args.arg_my)
	adduniaxialanisotropy_args.argptr[5] = unsafe.Pointer(&adduniaxialanisotropy_args.arg_mz)
	adduniaxialanisotropy_args.argptr[6] = unsafe.Pointer(&adduniaxialanisotropy_args.arg_K1LUT)
	adduniaxialanisotropy_args.argptr[7] = unsafe.Pointer(&adduniaxialanisotropy_args.arg_K2LUT)
	adduniaxialanisotropy_args.argptr[8] = unsafe.Pointer(&adduniaxialanisotropy_args.arg_uxLUT)
	adduniaxialanisotropy_args.argptr[9] = unsafe.Pointer(&adduniaxialanisotropy_args.arg_uyLUT)
	adduniaxialanisotropy_args.argptr[10] = unsafe.Pointer(&adduniaxialanisotropy_args.arg_uzLUT)
	adduniaxialanisotropy_args.argptr[11] = unsafe.Pointer(&adduniaxialanisotropy_args.arg_regions)
	adduniaxialanisotropy_args.argptr[12] = unsafe.Pointer(&adduniaxialanisotropy_args.arg_N)
}

// Wrapper for adduniaxialanisotropy CUDA kernel, asynchronous.
func k_adduniaxialanisotropy_async(Bx unsafe.Pointer, By unsafe.Pointer, Bz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, K1LUT unsafe.Pointer, K2LUT unsafe.Pointer, uxLUT unsafe.Pointer, uyLUT unsafe.Pointer, uzLUT unsafe.Pointer, regions unsafe.Pointer, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("adduniaxialanisotropy")
	}

	adduniaxialanisotropy_args.Lock()
	defer adduniaxialanisotropy_args.Unlock()

	if adduniaxialanisotropy_code == 0 {
		adduniaxialanisotropy_code = fatbinLoad(adduniaxialanisotropy_map, "adduniaxialanisotropy")
	}

	adduniaxialanisotropy_args.arg_Bx = Bx
	adduniaxialanisotropy_args.arg_By = By
	adduniaxialanisotropy_args.arg_Bz = Bz
	adduniaxialanisotropy_args.arg_mx = mx
	adduniaxialanisotropy_args.arg_my = my
	adduniaxialanisotropy_args.arg_mz = mz
	adduniaxialanisotropy_args.arg_K1LUT = K1LUT
	adduniaxialanisotropy_args.arg_K2LUT = K2LUT
	adduniaxialanisotropy_args.arg_uxLUT = uxLUT
	adduniaxialanisotropy_args.arg_uyLUT = uyLUT
	adduniaxialanisotropy_args.arg_uzLUT = uzLUT
	adduniaxialanisotropy_args.arg_regions = regions
	adduniaxialanisotropy_args.arg_N = N

	args := adduniaxialanisotropy_args.argptr[:]
	cu.LaunchKernel(adduniaxialanisotropy_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("adduniaxialanisotropy")
	}
}

// maps compute capability on PTX code for adduniaxialanisotropy kernel.
var adduniaxialanisotropy_map = map[int]string{0: "",
	20: adduniaxialanisotropy_ptx_20,
	30: adduniaxialanisotropy_ptx_30,
	35: adduniaxialanisotropy_ptx_35,
	50: adduniaxialanisotropy_ptx_50,
	52: adduniaxialanisotropy_ptx_52}

// adduniaxialanisotropy PTX code for various compute capabilities.
const (
	adduniaxialanisotropy_ptx_20 = `
.version 4.1
.target sm_20
.address_size 64


.visible .entry adduniaxialanisotropy(
	.param .u64 adduniaxialanisotropy_param_0,
	.param .u64 adduniaxialanisotropy_param_1,
	.param .u64 adduniaxialanisotropy_param_2,
	.param .u64 adduniaxialanisotropy_param_3,
	.param .u64 adduniaxialanisotropy_param_4,
	.param .u64 adduniaxialanisotropy_param_5,
	.param .u64 adduniaxialanisotropy_param_6,
	.param .u64 adduniaxialanisotropy_param_7,
	.param .u64 adduniaxialanisotropy_param_8,
	.param .u64 adduniaxialanisotropy_param_9,
	.param .u64 adduniaxialanisotropy_param_10,
	.param .u64 adduniaxialanisotropy_param_11,
	.param .u32 adduniaxialanisotropy_param_12
)
{
	.reg .pred 	%p<3>;
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<11>;
	.reg .f32 	%f<41>;
	.reg .s64 	%rd<43>;


	ld.param.u64 	%rd3, [adduniaxialanisotropy_param_0];
	ld.param.u64 	%rd4, [adduniaxialanisotropy_param_1];
	ld.param.u64 	%rd5, [adduniaxialanisotropy_param_2];
	ld.param.u64 	%rd6, [adduniaxialanisotropy_param_3];
	ld.param.u64 	%rd7, [adduniaxialanisotropy_param_4];
	ld.param.u64 	%rd8, [adduniaxialanisotropy_param_5];
	ld.param.u64 	%rd9, [adduniaxialanisotropy_param_6];
	ld.param.u64 	%rd10, [adduniaxialanisotropy_param_7];
	ld.param.u64 	%rd11, [adduniaxialanisotropy_param_8];
	ld.param.u64 	%rd12, [adduniaxialanisotropy_param_9];
	ld.param.u64 	%rd13, [adduniaxialanisotropy_param_10];
	ld.param.u64 	%rd14, [adduniaxialanisotropy_param_11];
	ld.param.u32 	%r2, [adduniaxialanisotropy_param_12];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_5;

	cvta.to.global.u64 	%rd15, %rd14;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd16, %rd15, %rd1;
	ld.global.u8 	%rs1, [%rd16];
	cvt.u64.u16	%rd17, %rs1;
	and.b64  	%rd2, %rd17, 255;
	cvta.to.global.u64 	%rd18, %rd11;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd19, %r10, 4;
	add.s64 	%rd20, %rd18, %rd19;
	cvta.to.global.u64 	%rd21, %rd12;
	add.s64 	%rd22, %rd21, %rd19;
	cvta.to.global.u64 	%rd23, %rd13;
	add.s64 	%rd24, %rd23, %rd19;
	ld.global.f32 	%f1, [%rd20];
	ld.global.f32 	%f2, [%rd22];
	mul.f32 	%f7, %f2, %f2;
	fma.rn.f32 	%f8, %f1, %f1, %f7;
	ld.global.f32 	%f3, [%rd24];
	fma.rn.f32 	%f9, %f3, %f3, %f8;
	sqrt.rn.f32 	%f4, %f9;
	setp.neu.f32	%p2, %f4, 0f00000000;
	@%p2 bra 	BB0_3;

	mov.f32 	%f40, 0f00000000;
	bra.uni 	BB0_4;

BB0_3:
	rcp.rn.f32 	%f40, %f4;

BB0_4:
	cvta.to.global.u64 	%rd25, %rd5;
	cvta.to.global.u64 	%rd26, %rd4;
	cvta.to.global.u64 	%rd27, %rd3;
	cvta.to.global.u64 	%rd28, %rd8;
	cvta.to.global.u64 	%rd29, %rd9;
	shl.b64 	%rd30, %rd2, 2;
	add.s64 	%rd31, %rd29, %rd30;
	cvta.to.global.u64 	%rd32, %rd10;
	add.s64 	%rd33, %rd32, %rd30;
	cvta.to.global.u64 	%rd34, %rd6;
	shl.b64 	%rd35, %rd1, 2;
	add.s64 	%rd36, %rd34, %rd35;
	cvta.to.global.u64 	%rd37, %rd7;
	add.s64 	%rd38, %rd37, %rd35;
	ld.global.f32 	%f11, [%rd36];
	mul.f32 	%f12, %f40, %f1;
	ld.global.f32 	%f13, [%rd38];
	mul.f32 	%f14, %f40, %f2;
	mul.f32 	%f15, %f13, %f14;
	fma.rn.f32 	%f16, %f11, %f12, %f15;
	add.s64 	%rd39, %rd28, %rd35;
	ld.global.f32 	%f17, [%rd39];
	mul.f32 	%f18, %f40, %f3;
	fma.rn.f32 	%f19, %f17, %f18, %f16;
	ld.global.f32 	%f20, [%rd31];
	add.f32 	%f21, %f20, %f20;
	mul.f32 	%f22, %f21, %f19;
	ld.global.f32 	%f23, [%rd33];
	mul.f32 	%f24, %f23, 0f40800000;
	mul.f32 	%f25, %f19, %f19;
	mul.f32 	%f26, %f25, %f19;
	mul.f32 	%f27, %f24, %f26;
	mul.f32 	%f28, %f27, %f12;
	mul.f32 	%f29, %f27, %f14;
	mul.f32 	%f30, %f27, %f18;
	fma.rn.f32 	%f31, %f22, %f12, %f28;
	fma.rn.f32 	%f32, %f22, %f14, %f29;
	fma.rn.f32 	%f33, %f22, %f18, %f30;
	add.s64 	%rd40, %rd27, %rd35;
	ld.global.f32 	%f34, [%rd40];
	add.f32 	%f35, %f34, %f31;
	st.global.f32 	[%rd40], %f35;
	add.s64 	%rd41, %rd26, %rd35;
	ld.global.f32 	%f36, [%rd41];
	add.f32 	%f37, %f36, %f32;
	st.global.f32 	[%rd41], %f37;
	add.s64 	%rd42, %rd25, %rd35;
	ld.global.f32 	%f38, [%rd42];
	add.f32 	%f39, %f38, %f33;
	st.global.f32 	[%rd42], %f39;

BB0_5:
	ret;
}


`
	adduniaxialanisotropy_ptx_30 = `
.version 4.1
.target sm_30
.address_size 64


.visible .entry adduniaxialanisotropy(
	.param .u64 adduniaxialanisotropy_param_0,
	.param .u64 adduniaxialanisotropy_param_1,
	.param .u64 adduniaxialanisotropy_param_2,
	.param .u64 adduniaxialanisotropy_param_3,
	.param .u64 adduniaxialanisotropy_param_4,
	.param .u64 adduniaxialanisotropy_param_5,
	.param .u64 adduniaxialanisotropy_param_6,
	.param .u64 adduniaxialanisotropy_param_7,
	.param .u64 adduniaxialanisotropy_param_8,
	.param .u64 adduniaxialanisotropy_param_9,
	.param .u64 adduniaxialanisotropy_param_10,
	.param .u64 adduniaxialanisotropy_param_11,
	.param .u32 adduniaxialanisotropy_param_12
)
{
	.reg .pred 	%p<3>;
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<11>;
	.reg .f32 	%f<41>;
	.reg .s64 	%rd<43>;


	ld.param.u64 	%rd3, [adduniaxialanisotropy_param_0];
	ld.param.u64 	%rd4, [adduniaxialanisotropy_param_1];
	ld.param.u64 	%rd5, [adduniaxialanisotropy_param_2];
	ld.param.u64 	%rd6, [adduniaxialanisotropy_param_3];
	ld.param.u64 	%rd7, [adduniaxialanisotropy_param_4];
	ld.param.u64 	%rd8, [adduniaxialanisotropy_param_5];
	ld.param.u64 	%rd9, [adduniaxialanisotropy_param_6];
	ld.param.u64 	%rd10, [adduniaxialanisotropy_param_7];
	ld.param.u64 	%rd11, [adduniaxialanisotropy_param_8];
	ld.param.u64 	%rd12, [adduniaxialanisotropy_param_9];
	ld.param.u64 	%rd13, [adduniaxialanisotropy_param_10];
	ld.param.u64 	%rd14, [adduniaxialanisotropy_param_11];
	ld.param.u32 	%r2, [adduniaxialanisotropy_param_12];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_5;

	cvta.to.global.u64 	%rd15, %rd14;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd16, %rd15, %rd1;
	ld.global.u8 	%rs1, [%rd16];
	cvt.u64.u16	%rd17, %rs1;
	and.b64  	%rd2, %rd17, 255;
	cvta.to.global.u64 	%rd18, %rd11;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd19, %r10, 4;
	add.s64 	%rd20, %rd18, %rd19;
	cvta.to.global.u64 	%rd21, %rd12;
	add.s64 	%rd22, %rd21, %rd19;
	cvta.to.global.u64 	%rd23, %rd13;
	add.s64 	%rd24, %rd23, %rd19;
	ld.global.f32 	%f1, [%rd20];
	ld.global.f32 	%f2, [%rd22];
	mul.f32 	%f7, %f2, %f2;
	fma.rn.f32 	%f8, %f1, %f1, %f7;
	ld.global.f32 	%f3, [%rd24];
	fma.rn.f32 	%f9, %f3, %f3, %f8;
	sqrt.rn.f32 	%f4, %f9;
	setp.neu.f32	%p2, %f4, 0f00000000;
	@%p2 bra 	BB0_3;

	mov.f32 	%f40, 0f00000000;
	bra.uni 	BB0_4;

BB0_3:
	rcp.rn.f32 	%f40, %f4;

BB0_4:
	cvta.to.global.u64 	%rd25, %rd5;
	cvta.to.global.u64 	%rd26, %rd4;
	cvta.to.global.u64 	%rd27, %rd3;
	cvta.to.global.u64 	%rd28, %rd8;
	cvta.to.global.u64 	%rd29, %rd9;
	shl.b64 	%rd30, %rd2, 2;
	add.s64 	%rd31, %rd29, %rd30;
	cvta.to.global.u64 	%rd32, %rd10;
	add.s64 	%rd33, %rd32, %rd30;
	cvta.to.global.u64 	%rd34, %rd6;
	shl.b64 	%rd35, %rd1, 2;
	add.s64 	%rd36, %rd34, %rd35;
	cvta.to.global.u64 	%rd37, %rd7;
	add.s64 	%rd38, %rd37, %rd35;
	ld.global.f32 	%f11, [%rd36];
	mul.f32 	%f12, %f40, %f1;
	ld.global.f32 	%f13, [%rd38];
	mul.f32 	%f14, %f40, %f2;
	mul.f32 	%f15, %f13, %f14;
	fma.rn.f32 	%f16, %f11, %f12, %f15;
	add.s64 	%rd39, %rd28, %rd35;
	ld.global.f32 	%f17, [%rd39];
	mul.f32 	%f18, %f40, %f3;
	fma.rn.f32 	%f19, %f17, %f18, %f16;
	ld.global.f32 	%f20, [%rd31];
	add.f32 	%f21, %f20, %f20;
	mul.f32 	%f22, %f21, %f19;
	ld.global.f32 	%f23, [%rd33];
	mul.f32 	%f24, %f23, 0f40800000;
	mul.f32 	%f25, %f19, %f19;
	mul.f32 	%f26, %f25, %f19;
	mul.f32 	%f27, %f24, %f26;
	mul.f32 	%f28, %f27, %f12;
	mul.f32 	%f29, %f27, %f14;
	mul.f32 	%f30, %f27, %f18;
	fma.rn.f32 	%f31, %f22, %f12, %f28;
	fma.rn.f32 	%f32, %f22, %f14, %f29;
	fma.rn.f32 	%f33, %f22, %f18, %f30;
	add.s64 	%rd40, %rd27, %rd35;
	ld.global.f32 	%f34, [%rd40];
	add.f32 	%f35, %f34, %f31;
	st.global.f32 	[%rd40], %f35;
	add.s64 	%rd41, %rd26, %rd35;
	ld.global.f32 	%f36, [%rd41];
	add.f32 	%f37, %f36, %f32;
	st.global.f32 	[%rd41], %f37;
	add.s64 	%rd42, %rd25, %rd35;
	ld.global.f32 	%f38, [%rd42];
	add.f32 	%f39, %f38, %f33;
	st.global.f32 	[%rd42], %f39;

BB0_5:
	ret;
}


`
	adduniaxialanisotropy_ptx_35 = `
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

.visible .entry adduniaxialanisotropy(
	.param .u64 adduniaxialanisotropy_param_0,
	.param .u64 adduniaxialanisotropy_param_1,
	.param .u64 adduniaxialanisotropy_param_2,
	.param .u64 adduniaxialanisotropy_param_3,
	.param .u64 adduniaxialanisotropy_param_4,
	.param .u64 adduniaxialanisotropy_param_5,
	.param .u64 adduniaxialanisotropy_param_6,
	.param .u64 adduniaxialanisotropy_param_7,
	.param .u64 adduniaxialanisotropy_param_8,
	.param .u64 adduniaxialanisotropy_param_9,
	.param .u64 adduniaxialanisotropy_param_10,
	.param .u64 adduniaxialanisotropy_param_11,
	.param .u32 adduniaxialanisotropy_param_12
)
{
	.reg .pred 	%p<3>;
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<11>;
	.reg .f32 	%f<41>;
	.reg .s64 	%rd<43>;


	ld.param.u64 	%rd3, [adduniaxialanisotropy_param_0];
	ld.param.u64 	%rd4, [adduniaxialanisotropy_param_1];
	ld.param.u64 	%rd5, [adduniaxialanisotropy_param_2];
	ld.param.u64 	%rd6, [adduniaxialanisotropy_param_3];
	ld.param.u64 	%rd7, [adduniaxialanisotropy_param_4];
	ld.param.u64 	%rd8, [adduniaxialanisotropy_param_5];
	ld.param.u64 	%rd9, [adduniaxialanisotropy_param_6];
	ld.param.u64 	%rd10, [adduniaxialanisotropy_param_7];
	ld.param.u64 	%rd11, [adduniaxialanisotropy_param_8];
	ld.param.u64 	%rd12, [adduniaxialanisotropy_param_9];
	ld.param.u64 	%rd13, [adduniaxialanisotropy_param_10];
	ld.param.u64 	%rd14, [adduniaxialanisotropy_param_11];
	ld.param.u32 	%r2, [adduniaxialanisotropy_param_12];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB5_5;

	cvta.to.global.u64 	%rd15, %rd14;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd16, %rd15, %rd1;
	ld.global.nc.u8 	%rs1, [%rd16];
	cvt.u64.u16	%rd17, %rs1;
	and.b64  	%rd2, %rd17, 255;
	cvta.to.global.u64 	%rd18, %rd11;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd19, %r10, 4;
	add.s64 	%rd20, %rd18, %rd19;
	cvta.to.global.u64 	%rd21, %rd12;
	add.s64 	%rd22, %rd21, %rd19;
	cvta.to.global.u64 	%rd23, %rd13;
	add.s64 	%rd24, %rd23, %rd19;
	ld.global.nc.f32 	%f1, [%rd20];
	ld.global.nc.f32 	%f2, [%rd22];
	mul.f32 	%f7, %f2, %f2;
	fma.rn.f32 	%f8, %f1, %f1, %f7;
	ld.global.nc.f32 	%f3, [%rd24];
	fma.rn.f32 	%f9, %f3, %f3, %f8;
	sqrt.rn.f32 	%f4, %f9;
	setp.neu.f32	%p2, %f4, 0f00000000;
	@%p2 bra 	BB5_3;

	mov.f32 	%f40, 0f00000000;
	bra.uni 	BB5_4;

BB5_3:
	rcp.rn.f32 	%f40, %f4;

BB5_4:
	cvta.to.global.u64 	%rd25, %rd5;
	cvta.to.global.u64 	%rd26, %rd4;
	cvta.to.global.u64 	%rd27, %rd3;
	cvta.to.global.u64 	%rd28, %rd8;
	cvta.to.global.u64 	%rd29, %rd7;
	cvta.to.global.u64 	%rd30, %rd6;
	cvta.to.global.u64 	%rd31, %rd10;
	cvta.to.global.u64 	%rd32, %rd9;
	shl.b64 	%rd33, %rd1, 2;
	add.s64 	%rd34, %rd30, %rd33;
	ld.global.nc.f32 	%f11, [%rd34];
	mul.f32 	%f12, %f40, %f1;
	add.s64 	%rd35, %rd29, %rd33;
	ld.global.nc.f32 	%f13, [%rd35];
	mul.f32 	%f14, %f40, %f2;
	mul.f32 	%f15, %f13, %f14;
	fma.rn.f32 	%f16, %f11, %f12, %f15;
	add.s64 	%rd36, %rd28, %rd33;
	ld.global.nc.f32 	%f17, [%rd36];
	mul.f32 	%f18, %f40, %f3;
	fma.rn.f32 	%f19, %f17, %f18, %f16;
	shl.b64 	%rd37, %rd2, 2;
	add.s64 	%rd38, %rd32, %rd37;
	ld.global.nc.f32 	%f20, [%rd38];
	add.f32 	%f21, %f20, %f20;
	mul.f32 	%f22, %f21, %f19;
	add.s64 	%rd39, %rd31, %rd37;
	ld.global.nc.f32 	%f23, [%rd39];
	mul.f32 	%f24, %f23, 0f40800000;
	mul.f32 	%f25, %f19, %f19;
	mul.f32 	%f26, %f25, %f19;
	mul.f32 	%f27, %f24, %f26;
	mul.f32 	%f28, %f27, %f12;
	mul.f32 	%f29, %f27, %f14;
	mul.f32 	%f30, %f27, %f18;
	fma.rn.f32 	%f31, %f22, %f12, %f28;
	fma.rn.f32 	%f32, %f22, %f14, %f29;
	fma.rn.f32 	%f33, %f22, %f18, %f30;
	add.s64 	%rd40, %rd27, %rd33;
	ld.global.f32 	%f34, [%rd40];
	add.f32 	%f35, %f34, %f31;
	st.global.f32 	[%rd40], %f35;
	add.s64 	%rd41, %rd26, %rd33;
	ld.global.f32 	%f36, [%rd41];
	add.f32 	%f37, %f36, %f32;
	st.global.f32 	[%rd41], %f37;
	add.s64 	%rd42, %rd25, %rd33;
	ld.global.f32 	%f38, [%rd42];
	add.f32 	%f39, %f38, %f33;
	st.global.f32 	[%rd42], %f39;

BB5_5:
	ret;
}


`
	adduniaxialanisotropy_ptx_50 = `
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

.visible .entry adduniaxialanisotropy(
	.param .u64 adduniaxialanisotropy_param_0,
	.param .u64 adduniaxialanisotropy_param_1,
	.param .u64 adduniaxialanisotropy_param_2,
	.param .u64 adduniaxialanisotropy_param_3,
	.param .u64 adduniaxialanisotropy_param_4,
	.param .u64 adduniaxialanisotropy_param_5,
	.param .u64 adduniaxialanisotropy_param_6,
	.param .u64 adduniaxialanisotropy_param_7,
	.param .u64 adduniaxialanisotropy_param_8,
	.param .u64 adduniaxialanisotropy_param_9,
	.param .u64 adduniaxialanisotropy_param_10,
	.param .u64 adduniaxialanisotropy_param_11,
	.param .u32 adduniaxialanisotropy_param_12
)
{
	.reg .pred 	%p<3>;
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<11>;
	.reg .f32 	%f<41>;
	.reg .s64 	%rd<43>;


	ld.param.u64 	%rd3, [adduniaxialanisotropy_param_0];
	ld.param.u64 	%rd4, [adduniaxialanisotropy_param_1];
	ld.param.u64 	%rd5, [adduniaxialanisotropy_param_2];
	ld.param.u64 	%rd6, [adduniaxialanisotropy_param_3];
	ld.param.u64 	%rd7, [adduniaxialanisotropy_param_4];
	ld.param.u64 	%rd8, [adduniaxialanisotropy_param_5];
	ld.param.u64 	%rd9, [adduniaxialanisotropy_param_6];
	ld.param.u64 	%rd10, [adduniaxialanisotropy_param_7];
	ld.param.u64 	%rd11, [adduniaxialanisotropy_param_8];
	ld.param.u64 	%rd12, [adduniaxialanisotropy_param_9];
	ld.param.u64 	%rd13, [adduniaxialanisotropy_param_10];
	ld.param.u64 	%rd14, [adduniaxialanisotropy_param_11];
	ld.param.u32 	%r2, [adduniaxialanisotropy_param_12];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB5_5;

	cvta.to.global.u64 	%rd15, %rd14;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd16, %rd15, %rd1;
	ld.global.nc.u8 	%rs1, [%rd16];
	cvt.u64.u16	%rd17, %rs1;
	and.b64  	%rd2, %rd17, 255;
	cvta.to.global.u64 	%rd18, %rd11;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd19, %r10, 4;
	add.s64 	%rd20, %rd18, %rd19;
	cvta.to.global.u64 	%rd21, %rd12;
	add.s64 	%rd22, %rd21, %rd19;
	cvta.to.global.u64 	%rd23, %rd13;
	add.s64 	%rd24, %rd23, %rd19;
	ld.global.nc.f32 	%f1, [%rd20];
	ld.global.nc.f32 	%f2, [%rd22];
	mul.f32 	%f7, %f2, %f2;
	fma.rn.f32 	%f8, %f1, %f1, %f7;
	ld.global.nc.f32 	%f3, [%rd24];
	fma.rn.f32 	%f9, %f3, %f3, %f8;
	sqrt.rn.f32 	%f4, %f9;
	setp.neu.f32	%p2, %f4, 0f00000000;
	@%p2 bra 	BB5_3;

	mov.f32 	%f40, 0f00000000;
	bra.uni 	BB5_4;

BB5_3:
	rcp.rn.f32 	%f40, %f4;

BB5_4:
	cvta.to.global.u64 	%rd25, %rd5;
	cvta.to.global.u64 	%rd26, %rd4;
	cvta.to.global.u64 	%rd27, %rd3;
	cvta.to.global.u64 	%rd28, %rd8;
	cvta.to.global.u64 	%rd29, %rd7;
	cvta.to.global.u64 	%rd30, %rd6;
	cvta.to.global.u64 	%rd31, %rd10;
	cvta.to.global.u64 	%rd32, %rd9;
	shl.b64 	%rd33, %rd1, 2;
	add.s64 	%rd34, %rd30, %rd33;
	ld.global.nc.f32 	%f11, [%rd34];
	mul.f32 	%f12, %f40, %f1;
	add.s64 	%rd35, %rd29, %rd33;
	ld.global.nc.f32 	%f13, [%rd35];
	mul.f32 	%f14, %f40, %f2;
	mul.f32 	%f15, %f13, %f14;
	fma.rn.f32 	%f16, %f11, %f12, %f15;
	add.s64 	%rd36, %rd28, %rd33;
	ld.global.nc.f32 	%f17, [%rd36];
	mul.f32 	%f18, %f40, %f3;
	fma.rn.f32 	%f19, %f17, %f18, %f16;
	shl.b64 	%rd37, %rd2, 2;
	add.s64 	%rd38, %rd32, %rd37;
	ld.global.nc.f32 	%f20, [%rd38];
	add.f32 	%f21, %f20, %f20;
	mul.f32 	%f22, %f21, %f19;
	add.s64 	%rd39, %rd31, %rd37;
	ld.global.nc.f32 	%f23, [%rd39];
	mul.f32 	%f24, %f23, 0f40800000;
	mul.f32 	%f25, %f19, %f19;
	mul.f32 	%f26, %f25, %f19;
	mul.f32 	%f27, %f24, %f26;
	mul.f32 	%f28, %f27, %f12;
	mul.f32 	%f29, %f27, %f14;
	mul.f32 	%f30, %f27, %f18;
	fma.rn.f32 	%f31, %f22, %f12, %f28;
	fma.rn.f32 	%f32, %f22, %f14, %f29;
	fma.rn.f32 	%f33, %f22, %f18, %f30;
	add.s64 	%rd40, %rd27, %rd33;
	ld.global.f32 	%f34, [%rd40];
	add.f32 	%f35, %f34, %f31;
	st.global.f32 	[%rd40], %f35;
	add.s64 	%rd41, %rd26, %rd33;
	ld.global.f32 	%f36, [%rd41];
	add.f32 	%f37, %f36, %f32;
	st.global.f32 	[%rd41], %f37;
	add.s64 	%rd42, %rd25, %rd33;
	ld.global.f32 	%f38, [%rd42];
	add.f32 	%f39, %f38, %f33;
	st.global.f32 	[%rd42], %f39;

BB5_5:
	ret;
}


`
	adduniaxialanisotropy_ptx_52 = `
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

	// .globl	adduniaxialanisotropy
.visible .entry adduniaxialanisotropy(
	.param .u64 adduniaxialanisotropy_param_0,
	.param .u64 adduniaxialanisotropy_param_1,
	.param .u64 adduniaxialanisotropy_param_2,
	.param .u64 adduniaxialanisotropy_param_3,
	.param .u64 adduniaxialanisotropy_param_4,
	.param .u64 adduniaxialanisotropy_param_5,
	.param .u64 adduniaxialanisotropy_param_6,
	.param .u64 adduniaxialanisotropy_param_7,
	.param .u64 adduniaxialanisotropy_param_8,
	.param .u64 adduniaxialanisotropy_param_9,
	.param .u64 adduniaxialanisotropy_param_10,
	.param .u64 adduniaxialanisotropy_param_11,
	.param .u32 adduniaxialanisotropy_param_12
)
{
	.reg .pred 	%p<3>;
	.reg .s16 	%rs<2>;
	.reg .f32 	%f<41>;
	.reg .s32 	%r<11>;
	.reg .s64 	%rd<43>;


	ld.param.u64 	%rd3, [adduniaxialanisotropy_param_0];
	ld.param.u64 	%rd4, [adduniaxialanisotropy_param_1];
	ld.param.u64 	%rd5, [adduniaxialanisotropy_param_2];
	ld.param.u64 	%rd6, [adduniaxialanisotropy_param_3];
	ld.param.u64 	%rd7, [adduniaxialanisotropy_param_4];
	ld.param.u64 	%rd8, [adduniaxialanisotropy_param_5];
	ld.param.u64 	%rd9, [adduniaxialanisotropy_param_6];
	ld.param.u64 	%rd10, [adduniaxialanisotropy_param_7];
	ld.param.u64 	%rd11, [adduniaxialanisotropy_param_8];
	ld.param.u64 	%rd12, [adduniaxialanisotropy_param_9];
	ld.param.u64 	%rd13, [adduniaxialanisotropy_param_10];
	ld.param.u64 	%rd14, [adduniaxialanisotropy_param_11];
	ld.param.u32 	%r2, [adduniaxialanisotropy_param_12];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB6_4;

	cvta.to.global.u64 	%rd15, %rd14;
	cvt.s64.s32	%rd1, %r1;
	add.s64 	%rd16, %rd15, %rd1;
	ld.global.nc.u8 	%rs1, [%rd16];
	cvt.u64.u16	%rd17, %rs1;
	and.b64  	%rd2, %rd17, 255;
	cvta.to.global.u64 	%rd18, %rd11;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd19, %r10, 4;
	add.s64 	%rd20, %rd18, %rd19;
	cvta.to.global.u64 	%rd21, %rd12;
	add.s64 	%rd22, %rd21, %rd19;
	cvta.to.global.u64 	%rd23, %rd13;
	add.s64 	%rd24, %rd23, %rd19;
	ld.global.nc.f32 	%f1, [%rd20];
	ld.global.nc.f32 	%f2, [%rd22];
	mul.f32 	%f8, %f2, %f2;
	fma.rn.f32 	%f9, %f1, %f1, %f8;
	ld.global.nc.f32 	%f3, [%rd24];
	fma.rn.f32 	%f10, %f3, %f3, %f9;
	sqrt.rn.f32 	%f4, %f10;
	mov.f32 	%f40, 0f00000000;
	setp.eq.f32	%p2, %f4, 0f00000000;
	@%p2 bra 	BB6_3;

	rcp.rn.f32 	%f40, %f4;

BB6_3:
	cvta.to.global.u64 	%rd25, %rd5;
	cvta.to.global.u64 	%rd26, %rd4;
	cvta.to.global.u64 	%rd27, %rd3;
	cvta.to.global.u64 	%rd28, %rd8;
	cvta.to.global.u64 	%rd29, %rd7;
	cvta.to.global.u64 	%rd30, %rd6;
	cvta.to.global.u64 	%rd31, %rd10;
	cvta.to.global.u64 	%rd32, %rd9;
	shl.b64 	%rd33, %rd1, 2;
	add.s64 	%rd34, %rd30, %rd33;
	ld.global.nc.f32 	%f11, [%rd34];
	mul.f32 	%f12, %f1, %f40;
	add.s64 	%rd35, %rd29, %rd33;
	ld.global.nc.f32 	%f13, [%rd35];
	mul.f32 	%f14, %f2, %f40;
	mul.f32 	%f15, %f14, %f13;
	fma.rn.f32 	%f16, %f12, %f11, %f15;
	add.s64 	%rd36, %rd28, %rd33;
	ld.global.nc.f32 	%f17, [%rd36];
	mul.f32 	%f18, %f3, %f40;
	fma.rn.f32 	%f19, %f18, %f17, %f16;
	shl.b64 	%rd37, %rd2, 2;
	add.s64 	%rd38, %rd32, %rd37;
	ld.global.nc.f32 	%f20, [%rd38];
	add.f32 	%f21, %f20, %f20;
	mul.f32 	%f22, %f21, %f19;
	add.s64 	%rd39, %rd31, %rd37;
	ld.global.nc.f32 	%f23, [%rd39];
	mul.f32 	%f24, %f23, 0f40800000;
	mul.f32 	%f25, %f19, %f19;
	mul.f32 	%f26, %f19, %f25;
	mul.f32 	%f27, %f24, %f26;
	mul.f32 	%f28, %f12, %f27;
	mul.f32 	%f29, %f14, %f27;
	mul.f32 	%f30, %f18, %f27;
	fma.rn.f32 	%f31, %f12, %f22, %f28;
	fma.rn.f32 	%f32, %f14, %f22, %f29;
	fma.rn.f32 	%f33, %f18, %f22, %f30;
	add.s64 	%rd40, %rd27, %rd33;
	ld.global.f32 	%f34, [%rd40];
	add.f32 	%f35, %f34, %f31;
	st.global.f32 	[%rd40], %f35;
	add.s64 	%rd41, %rd26, %rd33;
	ld.global.f32 	%f36, [%rd41];
	add.f32 	%f37, %f36, %f32;
	st.global.f32 	[%rd41], %f37;
	add.s64 	%rd42, %rd25, %rd33;
	ld.global.f32 	%f38, [%rd42];
	add.f32 	%f39, %f38, %f33;
	st.global.f32 	[%rd42], %f39;

BB6_4:
	ret;
}


`
)
