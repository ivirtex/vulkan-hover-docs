# VK\_KHR\_shader\_relaxed\_extended\_instruction(3) Manual Page

## Name

VK\_KHR\_shader\_relaxed\_extended\_instruction - device extension



## [](#_registered_extension_number)Registered Extension Number

559

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_relaxed\_extended\_instruction](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_relaxed_extended_instruction.html)

## [](#_contact)Contact

- Nathan Gauër [\[GitHub\]Keenuts](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_relaxed_extended_instruction%5D%20%40Keenuts%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_relaxed_extended_instruction%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_shader\_relaxed\_extended\_instruction](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_relaxed_extended_instruction.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-01-24

**IP Status**

No known IP claims.

**Contributors**

- Alan Baker, Google LLC
- Nathan Gauër, Google LLC

## [](#_description)Description

This extension allows the use of the `SPV_KHR_relaxed_extended_instruction` extension in SPIR-V shader modules.

It adds a new SPIR-V instruction, which allows some usage of forward references in non-semantic instruction sets. This extensions interacts with the `SPV_KHR_non_semantic_info` extension, hence with `VK_KHR_shader_non_semantic_info`.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderRelaxedExtendedInstructionFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderRelaxedExtendedInstructionFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SHADER_RELAXED_EXTENDED_INSTRUCTION_EXTENSION_NAME`
- `VK_KHR_SHADER_RELAXED_EXTENDED_INSTRUCTION_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_RELAXED_EXTENDED_INSTRUCTION_FEATURES_KHR`

## [](#_version_history)Version History

- Revision 1, 2024-01-24 (Nathan Gauër)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_shader_relaxed_extended_instruction)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0