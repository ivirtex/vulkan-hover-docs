# VK\_KHR\_cooperative\_matrix(3) Manual Page

## Name

VK\_KHR\_cooperative\_matrix - device extension



## [](#_registered_extension_number)Registered Extension Number

507

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_cooperative\_matrix](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_cooperative_matrix.html)

## [](#_contact)Contact

- Kevin Petit [\[GitHub\]kpet](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_cooperative_matrix%5D%20%40kpet%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_cooperative_matrix%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_cooperative\_matrix](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_cooperative_matrix.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-05-03

**Interactions and External Dependencies**

- This extension provides API support for [`GLSL_KHR_cooperative_matrix`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/khr/GLSL_KHR_cooperative_matrix.txt)

**Contributors**

- Jeff Bolz, NVIDIA
- Markus Tavenrath, NVIDIA
- Daniel Koch, NVIDIA
- Kevin Petit, Arm Ltd.
- Boris Zanin, AMD

## [](#_description)Description

This extension adds support for using cooperative matrix types in SPIR-V. Cooperative matrix types are medium-sized matrices that are primarily supported in compute shaders, where the storage for the matrix is spread across all invocations in some scope (usually a subgroup) and those invocations cooperate to efficiently perform matrix multiplies.

Cooperative matrix types are defined by the [`SPV_KHR_cooperative_matrix`](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_cooperative_matrix.html) SPIR-V extension and can be used with the [`GLSL_KHR_cooperative_matrix`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/khr/GLSL_KHR_cooperative_matrix.txt) GLSL extension.

This extension includes support for enumerating the matrix types and dimensions that are supported by the implementation.

## [](#_new_commands)New Commands

- [vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR.html)

## [](#_new_structures)New Structures

- [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCooperativeMatrixPropertiesKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceCooperativeMatrixFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCooperativeMatrixFeaturesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCooperativeMatrixPropertiesKHR.html)

## [](#_new_enums)New Enums

- [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html)
- [VkScopeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScopeKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_COOPERATIVE_MATRIX_EXTENSION_NAME`
- `VK_KHR_COOPERATIVE_MATRIX_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_COOPERATIVE_MATRIX_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_PROPERTIES_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [CooperativeMatrixKHR](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-CooperativeMatrixKHR)

## [](#_issues)Issues

## [](#_version_history)Version History

- Revision 2, 2023-05-03 (Kevin Petit)
  
  - First KHR revision
- Revision 1, 2019-02-05 (Jeff Bolz)
  
  - NVIDIA vendor extension

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_cooperative_matrix).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0