# VK\_NV\_shader\_atomic\_float16\_vector(3) Manual Page

## Name

VK\_NV\_shader\_atomic\_float16\_vector - device extension



## [](#_registered_extension_number)Registered Extension Number

564

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_shader\_atomic\_fp16\_vector](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_shader_atomic_fp16_vector.html)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_shader_atomic_float16_vector%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_shader_atomic_float16_vector%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-02-03

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_NV_shader_atomic_fp16_vector`](https://registry.khronos.org/OpenGL/extensions/NV/NV_shader_atomic_fp16_vector.txt)

**Contributors**

- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension allows a shader to perform atomic add, min, max, and exchange operations on 2- and 4-component vectors of float16. Buffer, workgroup, and image storage classes are all supported.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_SHADER_ATOMIC_FLOAT16_VECTOR_EXTENSION_NAME`
- `VK_NV_SHADER_ATOMIC_FLOAT16_VECTOR_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_FLOAT16_VECTOR_FEATURES_NV`

## [](#_issues)Issues

None.

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`AtomicFloat16VectorNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-AtomicFloat16VectorNV)

## [](#_version_history)Version History

- Revision 1, 2024-02-03 (Jeff Bolz)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_shader_atomic_float16_vector)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0