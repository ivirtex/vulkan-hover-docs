# VK\_EXT\_shader\_atomic\_float2(3) Manual Page

## Name

VK\_EXT\_shader\_atomic\_float2 - device extension



## [](#_registered_extension_number)Registered Extension Number

274

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_shader\_atomic\_float](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_atomic_float.html)

## [](#_api_interactions)API Interactions

- Interacts with VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT::sparseImageFloat32AtomicMinMax

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_EXT\_shader\_atomic\_float16\_add](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_shader_atomic_float16_add.html)
- [SPV\_EXT\_shader\_atomic\_float\_min\_max](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_shader_atomic_float_min_max.html)

## [](#_contact)Contact

- Faith Ekstrand [\[GitHub\]gfxstrand](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_atomic_float2%5D%20%40gfxstrand%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_atomic_float2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-08-14

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GLSL_EXT_shader_atomic_float2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_shader_atomic_float2.txt)

**Contributors**

- Faith Ekstrand, Intel

## [](#_description)Description

This extension allows a shader to perform 16-bit floating-point atomic operations on buffer and workgroup memory as well as floating-point atomic minimum and maximum operations on buffer, workgroup, and image memory. It advertises the SPIR-V `AtomicFloat16AddEXT` capability which allows atomic add operations on 16-bit floating-point numbers and the SPIR-V `AtomicFloat16MinMaxEXT`, `AtomicFloat32MinMaxEXT` and `AtomicFloat64MinMaxEXT` capabilities which allow atomic minimum and maximum operations on floating-point numbers. The supported operations include `OpAtomicFAddEXT`, `OpAtomicFMinEXT` and `OpAtomicFMaxEXT`.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SHADER_ATOMIC_FLOAT_2_EXTENSION_NAME`
- `VK_EXT_SHADER_ATOMIC_FLOAT_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_FLOAT_2_FEATURES_EXT`

## [](#_issues)Issues

1\) Should this extension add support for 16-bit image atomics?

**RESOLVED**: No. While Vulkan supports creating storage images with `VK_FORMAT_R16_SFLOAT` and doing load and store on them, the data in the shader has a 32-bit representation. Vulkan currently has no facility for even basic reading or writing such images using 16-bit float values in the shader. Adding such functionality would be required before 16-bit image atomics would make sense and is outside the scope of this extension.

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`AtomicFloat32MinMaxEXT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-AtomicFloat16AddEXT)
- [`AtomicFloat32MinMaxEXT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-AtomicFloat16MinMaxEXT)
- [`AtomicFloat32MinMaxEXT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-AtomicFloat32MinMaxEXT)
- [`AtomicFloat64MinMaxEXT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-AtomicFloat64MinMaxEXT)

## [](#_version_history)Version History

- Revision 1, 2020-08-14 (Faith Ekstrand)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_shader_atomic_float2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0