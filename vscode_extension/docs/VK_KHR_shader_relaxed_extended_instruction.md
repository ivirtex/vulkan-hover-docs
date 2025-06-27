# VK_KHR_shader_relaxed_extended_instruction(3) Manual Page

## Name

VK_KHR_shader_relaxed_extended_instruction - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

559

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_relaxed_extended_instruction](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_relaxed_extended_instruction.html)

## <a href="#_contact" class="anchor"></a>Contact

- Nathan Gauër <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_relaxed_extended_instruction%5D%20@Keenuts%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_relaxed_extended_instruction%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>Keenuts</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_shader_relaxed_extended_instruction](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_relaxed_extended_instruction.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2024-01-24

**IP Status**  
No known IP claims.

**Contributors**  
- Alan Baker, Google LLC

- Nathan Gauër, Google LLC

## <a href="#_description" class="anchor"></a>Description

This extension allows the use of the
`SPV_KHR_relaxed_extended_instruction` extension in SPIR-V shader
modules.

It adds a new SPIR-V instruction, which allows some usage of forward
references in non-semantic instruction sets. This extensions interacts
with the `SPV_KHR_non_semantic_info` extension, hence with
`VK_KHR_shader_non_semantic_info`.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderRelaxedExtendedInstructionFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderRelaxedExtendedInstructionFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_RELAXED_EXTENDED_INSTRUCTION_EXTENSION_NAME`

- `VK_KHR_SHADER_RELAXED_EXTENDED_INSTRUCTION_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_RELAXED_EXTENDED_INSTRUCTION_FEATURES_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2024-01-24 (Nathan Gauër)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_relaxed_extended_instruction"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
