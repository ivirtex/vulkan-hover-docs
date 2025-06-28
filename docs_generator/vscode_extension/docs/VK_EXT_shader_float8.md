# VK\_EXT\_shader\_float8(3) Manual Page

## Name

VK\_EXT\_shader\_float8 - device extension



## [](#_registered_extension_number)Registered Extension Number

568

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_api_interactions)API Interactions

- Interacts with VK\_KHR\_cooperative\_matrix

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_EXT\_float8](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_float8.html)

## [](#_contact)Contact

- Kevin Petit [\[GitHub\]kpet](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_float8%5D%20%40kpet%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_float8%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_shader\_float8](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_shader_float8.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-04-16

**IP Status**

No known IP claims.

**Contributors**

- Kévin Petit, Arm Ltd.
- Stu Smith, AMD
- Jeff Bolz, NVIDIA
- Craig Graham, Samsung

## [](#_description)Description

This extension enables support for 8-bit floating-point data types as defined in SPV\_EXT\_float8.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderFloat8FeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderFloat8FeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SHADER_FLOAT8_EXTENSION_NAME`
- `VK_EXT_SHADER_FLOAT8_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT8_FEATURES_EXT`

If [VK\_KHR\_cooperative\_matrix](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_cooperative_matrix.html) is supported:

- Extending [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html):
  
  - `VK_COMPONENT_TYPE_FLOAT8_E4M3_EXT`
  - `VK_COMPONENT_TYPE_FLOAT8_E5M2_EXT`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [Float8EXT](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-Float8EXT)
- [Float8CooperativeMatrixEXT](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-Float8CooperativeMatrixEXT)

## [](#_issues)Issues

1\) Resolve interactions with the changes VK\_KHR\_shader\_float16 makes to rules for denorm flushing (always allowed by default for all FP formats). How to describe the requirement to preserve subnormals?

\+ **RESOLVED**: Subnormals are always preserved when converting FP8 values to IEEE 754 binary 16. In all other cases, subnormals may be flushed to zero.

\+

## [](#_version_history)Version History

- Revision 1, 2025-04-16 (Kévin Petit)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_shader_float8)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0