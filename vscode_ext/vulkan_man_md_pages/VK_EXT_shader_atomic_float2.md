# VK_EXT_shader_atomic_float2(3) Manual Page

## Name

VK_EXT_shader_atomic_float2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

274

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_EXT_shader_atomic_float](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_atomic_float.html)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_EXT_shader_atomic_float16_add](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_shader_atomic_float16_add.html)

- [SPV_EXT_shader_atomic_float_min_max](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_shader_atomic_float_min_max.html)

## <a href="#_contact" class="anchor"></a>Contact

- Faith Ekstrand <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_atomic_float2%5D%20@gfxstrand%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_atomic_float2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>gfxstrand</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-08-14

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GLSL_EXT_shader_atomic_float2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_shader_atomic_float2.txt)

**Contributors**  
- Faith Ekstrand, Intel

## <a href="#_description" class="anchor"></a>Description

This extension allows a shader to perform 16-bit floating-point atomic
operations on buffer and workgroup memory as well as floating-point
atomic minimum and maximum operations on buffer, workgroup, and image
memory. It advertises the SPIR-V `AtomicFloat16AddEXT` capability which
allows atomic add operations on 16-bit floating-point numbers and the
SPIR-V `AtomicFloat16MinMaxEXT`, `AtomicFloat32MinMaxEXT` and
`AtomicFloat64MinMaxEXT` capabilities which allow atomic minimum and
maximum operations on floating-point numbers. The supported operations
include `OpAtomicFAddEXT`, `OpAtomicFMinEXT` and `OpAtomicFMaxEXT`.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SHADER_ATOMIC_FLOAT_2_EXTENSION_NAME`

- `VK_EXT_SHADER_ATOMIC_FLOAT_2_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_FLOAT_2_FEATURES_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should this extension add support for 16-bit image atomics?

**RESOLVED**: No. While Vulkan supports creating storage images with
`VK_FORMAT_R16_SFLOAT` and doing load and store on them, the data in the
shader has a 32-bit representation. Vulkan currently has no facility for
even basic reading or writing such images using 16-bit float values in
the shader. Adding such functionality would be required before 16-bit
image atomics would make sense and is outside the scope of this
extension.

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-AtomicFloat16AddEXT"
  target="_blank" rel="noopener"><code>AtomicFloat32MinMaxEXT</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-AtomicFloat16MinMaxEXT"
  target="_blank" rel="noopener"><code>AtomicFloat32MinMaxEXT</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-AtomicFloat32MinMaxEXT"
  target="_blank" rel="noopener"><code>AtomicFloat32MinMaxEXT</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-AtomicFloat64MinMaxEXT"
  target="_blank" rel="noopener"><code>AtomicFloat64MinMaxEXT</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-08-14 (Faith Ekstrand)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_shader_atomic_float2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
