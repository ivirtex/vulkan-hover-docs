# VK\_KHR\_shader\_bfloat16(3) Manual Page

## Name

VK\_KHR\_shader\_bfloat16 - device extension



## [](#_registered_extension_number)Registered Extension Number

142

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_KHR\_cooperative\_matrix

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_bfloat16](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_bfloat16.html)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_bfloat16%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_bfloat16%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_shader\_bfloat16](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_bfloat16.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-04-09

**IP Status**

No known IP claims.

**Contributors**

- Tobias Hector, AMD
- Stu Smith, AMD
- Jeff Bolz, Nvidia
- Kévin Petit, Arm
- David Neto, Google
- Graeme Leese, Broadcom
- Ruihao Zhang, Qualcomm
- Mark Sheppard, Imagination
- Ben Ashbaugh, Intel
- Dmitry Sidorov, Intel
- Victor Mustya, Intel

## [](#_description)Description

This extension enables support for bfloat16 (“brain float”) operations in shaders as defined in `SPV_KHR_bfloat16`.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderBfloat16FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderBfloat16FeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SHADER_BFLOAT16_EXTENSION_NAME`
- `VK_KHR_SHADER_BFLOAT16_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_BFLOAT16_FEATURES_KHR`

If [VK\_KHR\_cooperative\_matrix](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_cooperative_matrix.html) is supported:

- Extending [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html):
  
  - `VK_COMPONENT_TYPE_BFLOAT16_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [BFloat16TypeKHR](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-BFloat16TypeKHR)
- [BFloat16DotProductKHR](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-BFloat16DotProductKHR)
- [BFloat16CooperativeMatrixKHR](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-BFloat16CooperativeMatrixKHR)

## [](#_version_history)Version History

- Revision 1, 2024-04-09 (Stu Smith)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_shader_bfloat16)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0