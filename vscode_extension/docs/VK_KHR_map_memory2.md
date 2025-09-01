# VK\_KHR\_map\_memory2(3) Manual Page

## Name

VK\_KHR\_map\_memory2 - device extension



## [](#_registered_extension_number)Registered Extension Number

272

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Faith Ekstrand [\[GitHub\]gfxstrand](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_map_memory2%5D%20%40gfxstrand%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_map_memory2%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_map\_memory2](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_map_memory2.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-03-14

**Interactions and External Dependencies**

- None

**Contributors**

- Faith Ekstrand, Collabora
- Tobias Hector, AMD

## [](#_description)Description

This extension provides extensible versions of the Vulkan memory map and unmap commands. The new commands are functionally identical to the core commands, except that their parameters are specified using extensible structures that can be used to pass extension-specific information.

## [](#_new_commands)New Commands

- [vkMapMemory2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory2KHR.html)
- [vkUnmapMemory2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUnmapMemory2KHR.html)

## [](#_new_structures)New Structures

- [VkMemoryMapInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapInfoKHR.html)
- [VkMemoryUnmapInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapInfoKHR.html)

## [](#_new_enums)New Enums

- [VkMemoryUnmapFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkMemoryUnmapFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_MAP_MEMORY_2_EXTENSION_NAME`
- `VK_KHR_MAP_MEMORY_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_MEMORY_MAP_INFO_KHR`
  - `VK_STRUCTURE_TYPE_MEMORY_UNMAP_INFO_KHR`

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4 with the KHR suffix omitted. The original type, enum and command names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 0, 2022-08-03 (Faith Ekstrand)
  
  - Internal revisions
- Revision 1, 2023-03-14
  
  - Public release

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_map_memory2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0