# SampleMask(3) Manual Page

## Name

SampleMask - Coverage mask for a fragment shader invocation



## <a href="#_description" class="anchor"></a>Description

`SampleMask`  
Decorating a variable with the `SampleMask` built-in decoration will
make any variable contain the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-shader-samplemask"
target="_blank" rel="noopener">sample mask</a> for the current fragment
shader invocation.

A variable in the `Input` storage class decorated with `SampleMask` will
contain a bitmask of the set of samples covered by the primitive
generating the fragment during rasterization. It has a sample bit set if
and only if the sample is considered covered for this fragment shader
invocation. `SampleMask`\[\] is an array of integers. Bits are mapped to
samples in a manner where bit B of mask M (`SampleMask[M]`) corresponds
to sample 32 × M + B.

A variable in the `Output` storage class decorated with `SampleMask` is
an array of integers forming a bit array in a manner similar to an input
variable decorated with `SampleMask`, but where each bit represents
coverage as computed by the shader. This computed `SampleMask` is
combined with the generated coverage mask in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-covg"
target="_blank" rel="noopener">multisample coverage</a> operation.

Variables decorated with `SampleMask` **must** be either an unsized
array, or explicitly sized to be no larger than the
implementation-dependent maximum sample-mask (as an array of 32-bit
elements), determined by the maximum number of samples.

If a fragment shader entry point’s interface includes an output variable
decorated with `SampleMask`, the sample mask will be undefined for any
array elements of any fragment shader invocations that fail to assign a
value. If a fragment shader entry point’s interface does not include an
output variable decorated with `SampleMask`, the sample mask has no
effect on the processing of a fragment.

Valid Usage

- <a href="#VUID-SampleMask-SampleMask-04357"
  id="VUID-SampleMask-SampleMask-04357"></a>
  VUID-SampleMask-SampleMask-04357  
  The `SampleMask` decoration **must** be used only within the
  `Fragment` `Execution` `Model`

- <a href="#VUID-SampleMask-SampleMask-04358"
  id="VUID-SampleMask-SampleMask-04358"></a>
  VUID-SampleMask-SampleMask-04358  
  The variable decorated with `SampleMask` **must** be declared using
  the `Input` or `Output` `Storage` `Class`

- <a href="#VUID-SampleMask-SampleMask-04359"
  id="VUID-SampleMask-SampleMask-04359"></a>
  VUID-SampleMask-SampleMask-04359  
  The variable decorated with `SampleMask` **must** be declared as an
  array of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#SampleMask"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
