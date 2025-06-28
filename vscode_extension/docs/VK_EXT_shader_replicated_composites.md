# VK\_EXT\_shader\_replicated\_composites(3) Manual Page

## Name

VK\_EXT\_shader\_replicated\_composites - device extension



## [](#_registered_extension_number)Registered Extension Number

565

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_EXT\_replicated\_composites](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_replicated_composites.html)

## [](#_contact)Contact

- Kevin Petit [\[GitHub\]kpet](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_replicated_composites%5D%20%40kpet%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_replicated_composites%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_shader\_replicated\_composites](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_shader_replicated_composites.adoc)

**Last Modified Date**

2024-02-08

**IP Status**

No known IP claims.

**Contributors**

- Kévin Petit, Arm Ltd.
- Jeff Bolz, NVIDIA
- Piers Daniell, NVIDIA

This extension adds support for creating composites from a single value in SPIR-V modules, as defined by SPV\_EXT\_replicated\_composites.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SHADER_REPLICATED_COMPOSITES_EXTENSION_NAME`
- `VK_EXT_SHADER_REPLICATED_COMPOSITES_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_REPLICATED_COMPOSITES_FEATURES_EXT`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [ReplicatedCompositesEXT](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ReplicatedCompositesEXT)

## [](#_version_history)Version History

- Revision 1, 2024-02-08 (Kévin Petit)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_shader_replicated_composites)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0