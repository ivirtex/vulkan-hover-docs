# VK\_NV\_cooperative\_matrix(3) Manual Page

## Name

VK\_NV\_cooperative\_matrix - device extension



## [](#_registered_extension_number)Registered Extension Number

250

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_cooperative\_matrix](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_cooperative_matrix.html)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_cooperative_matrix%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_cooperative_matrix%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-02-05

**Interactions and External Dependencies**

- This extension provides API support for [`GL_NV_cooperative_matrix`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_cooperative_matrix.txt)

**Contributors**

- Jeff Bolz, NVIDIA
- Markus Tavenrath, NVIDIA
- Daniel Koch, NVIDIA

## [](#_description)Description

This extension adds support for using cooperative matrix types in SPIR-V. Cooperative matrix types are medium-sized matrices that are primarily supported in compute shaders, where the storage for the matrix is spread across all invocations in some scope (usually a subgroup) and those invocations cooperate to efficiently perform matrix multiplies.

Cooperative matrix types are defined by the [`SPV_NV_cooperative_matrix`](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_cooperative_matrix.html) SPIR-V extension and can be used with the [`GL_NV_cooperative_matrix`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_cooperative_matrix.txt) GLSL extension.

This extension includes support for enumerating the matrix types and dimensions that are supported by the implementation.

## [](#_new_commands)New Commands

- [vkGetPhysicalDeviceCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCooperativeMatrixPropertiesNV.html)

## [](#_new_structures)New Structures

- [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCooperativeMatrixPropertiesNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceCooperativeMatrixFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCooperativeMatrixFeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCooperativeMatrixPropertiesNV.html)

## [](#_new_enums)New Enums

- [VkComponentTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeNV.html)
- [VkScopeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScopeNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_COOPERATIVE_MATRIX_EXTENSION_NAME`
- `VK_NV_COOPERATIVE_MATRIX_SPEC_VERSION`
- Extending [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html):
  
  - `VK_COMPONENT_TYPE_FLOAT16_NV`
  - `VK_COMPONENT_TYPE_FLOAT32_NV`
  - `VK_COMPONENT_TYPE_FLOAT64_NV`
  - `VK_COMPONENT_TYPE_SINT16_NV`
  - `VK_COMPONENT_TYPE_SINT32_NV`
  - `VK_COMPONENT_TYPE_SINT64_NV`
  - `VK_COMPONENT_TYPE_SINT8_NV`
  - `VK_COMPONENT_TYPE_UINT16_NV`
  - `VK_COMPONENT_TYPE_UINT32_NV`
  - `VK_COMPONENT_TYPE_UINT64_NV`
  - `VK_COMPONENT_TYPE_UINT8_NV`
- Extending [VkScopeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScopeKHR.html):
  
  - `VK_SCOPE_DEVICE_NV`
  - `VK_SCOPE_QUEUE_FAMILY_NV`
  - `VK_SCOPE_SUBGROUP_NV`
  - `VK_SCOPE_WORKGROUP_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_COOPERATIVE_MATRIX_PROPERTIES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_PROPERTIES_NV`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`CooperativeMatrixNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-CooperativeMatrixNV)

## [](#_issues)Issues

(1) What matrix properties will be supported in practice?

**RESOLVED**: In NVIDIA’s initial implementation, we will support:

- AType = BType = fp16 CType = DType = fp16 MxNxK = 16x8x16 scope = Subgroup
- AType = BType = fp16 CType = DType = fp16 MxNxK = 16x8x8 scope = Subgroup
- AType = BType = fp16 CType = DType = fp32 MxNxK = 16x8x16 scope = Subgroup
- AType = BType = fp16 CType = DType = fp32 MxNxK = 16x8x8 scope = Subgroup

## [](#_version_history)Version History

- Revision 1, 2019-02-05 (Jeff Bolz)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_cooperative_matrix).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0