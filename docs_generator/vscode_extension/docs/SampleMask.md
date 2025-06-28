# SampleMask(3) Manual Page

## Name

SampleMask - Coverage mask for a fragment shader invocation



## [](#_description)Description

`SampleMask`

Decorating a variable with the `SampleMask` built-in decoration will make any variable contain the [sample mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-shader-samplemask) for the current fragment shader invocation.

A variable in the `Input` storage class decorated with `SampleMask` will contain a bitmask of the set of samples covered by the primitive generating the fragment during rasterization. It has a sample bit set if and only if the sample is considered covered for this fragment shader invocation. `SampleMask`\[] is an array of integers. Bits are mapped to samples in a manner where bit B of mask M (`SampleMask[M]`) corresponds to sample 32 × M + B.

A variable in the `Output` storage class decorated with `SampleMask` is an array of integers forming a bit array in a manner similar to an input variable decorated with `SampleMask`, but where each bit represents coverage as computed by the shader. This computed `SampleMask` is combined with the generated coverage mask in the [multisample coverage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-covg) operation.

Variables decorated with `SampleMask` **must** be either an unsized array, or explicitly sized to be no larger than the implementation-dependent maximum sample-mask (as an array of 32-bit elements), determined by the maximum number of samples.

If a fragment shader entry point’s interface includes an output variable decorated with `SampleMask`, the sample mask will be undefined for any array elements of any fragment shader invocations that fail to assign a value. If a fragment shader entry point’s interface does not include an output variable decorated with `SampleMask`, the sample mask has no effect on the processing of a fragment.

Valid Usage

- [](#VUID-SampleMask-SampleMask-04357)VUID-SampleMask-SampleMask-04357  
  The `SampleMask` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-SampleMask-SampleMask-04358)VUID-SampleMask-SampleMask-04358  
  The variable decorated with `SampleMask` **must** be declared using the `Input` or `Output` `Storage` `Class`
- [](#VUID-SampleMask-SampleMask-04359)VUID-SampleMask-SampleMask-04359  
  The variable decorated with `SampleMask` **must** be declared as an array of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#SampleMask)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0