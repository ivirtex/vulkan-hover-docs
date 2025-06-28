# VK\_NV\_cooperative\_vector(3) Manual Page

## Name

VK\_NV\_cooperative\_vector - device extension



## [](#_registered_extension_number)Registered Extension Number

492

## [](#_revision)Revision

4

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_cooperative\_vector](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_cooperative_vector.html)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_cooperative_vector%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_cooperative_vector%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_NV\_cooperative\_vector](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_NV_cooperative_vector.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-05-23

**Interactions and External Dependencies**

- This extension requires [`SPV_NV_cooperative_vector`](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_cooperative_vector.html)
- This extension provides API support for [`GL_NV_cooperative_vector`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_cooperative_vector.txt)

**Contributors**

- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension adds support for using cooperative vector types in SPIR-V. Unlike cooperative matrix types, a variable with a cooperative vector type is logically stored in the invocation it belongs to, but they can cooperate behind the scenes when performing matrix-vector multiplies. Cooperative vectors do not require a fully occupied subgroup or uniform control flow like cooperative matrices, although these do increase the likelihood of being on the fast path. And unlike normal vector types, they have arbitrary length and support a relatively limited set of operations. These types are intended to help accelerate the evaluation of small neural networks, where each invocation is performing its own independent evaluation of the network.

Cooperative vector types are defined by the [`SPV_NV_cooperative_vector`](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_cooperative_vector.html) SPIR-V extension and can be used with the [`GL_NV_cooperative_vector`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_cooperative_vector.txt) GLSL extension.

This extension includes support for enumerating the combinations of types that are supported by the implementation, and for converting matrix data to and from an optimized opaque layout.

## [](#_new_commands)New Commands

- [vkCmdConvertCooperativeVectorMatrixNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdConvertCooperativeVectorMatrixNV.html)
- [vkConvertCooperativeVectorMatrixNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkConvertCooperativeVectorMatrixNV.html)
- [vkGetPhysicalDeviceCooperativeVectorPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCooperativeVectorPropertiesNV.html)

## [](#_new_structures)New Structures

- [VkConvertCooperativeVectorMatrixInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConvertCooperativeVectorMatrixInfoNV.html)
- [VkCooperativeVectorPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCooperativeVectorPropertiesNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceCooperativeVectorFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCooperativeVectorFeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceCooperativeVectorPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCooperativeVectorPropertiesNV.html)

## [](#_new_unions)New Unions

- [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html)
- [VkDeviceOrHostAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressKHR.html)

## [](#_new_enums)New Enums

- [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html)
- [VkCooperativeVectorMatrixLayoutNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCooperativeVectorMatrixLayoutNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_COOPERATIVE_VECTOR_EXTENSION_NAME`
- `VK_NV_COOPERATIVE_VECTOR_SPEC_VERSION`
- Extending [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html):
  
  - `VK_COMPONENT_TYPE_FLOAT_E4M3_NV`
  - `VK_COMPONENT_TYPE_FLOAT_E5M2_NV`
  - `VK_COMPONENT_TYPE_SINT8_PACKED_NV`
  - `VK_COMPONENT_TYPE_UINT8_PACKED_NV`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_CONVERT_COOPERATIVE_VECTOR_MATRIX_BIT_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_CONVERT_COOPERATIVE_VECTOR_MATRIX_INFO_NV`
  - `VK_STRUCTURE_TYPE_COOPERATIVE_VECTOR_PROPERTIES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_VECTOR_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_VECTOR_PROPERTIES_NV`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [CooperativeVectorNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-CooperativeVectorNV)

## [](#_issues)Issues

## [](#_version_history)Version History

- Revision 4, 2024-05-23 (Jeff Bolz)
  
  - Add maxCooperativeVectorComponents
- Revision 3, 2024-05-23 (Jeff Bolz)
  
  - Add training functions
- Revision 2, 2024-02-10 (Jeff Bolz)
  
  - Add device-side matrix conversion
- Revision 1, 2023-12-13 (Jeff Bolz)
  
  - Initial revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_cooperative_vector)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0