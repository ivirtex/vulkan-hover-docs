# VK\_EXT\_external\_memory\_metal(3) Manual Page

## Name

VK\_EXT\_external\_memory\_metal - device extension



## [](#_registered_extension_number)Registered Extension Number

603

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Aitor Camacho Larrondo [\[GitHub\]aitor-lunarg](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_external_memory_metal%5D%20%40aitor-lunarg%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_external_memory_metal%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_external\_memory\_metal](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_external_memory_metal.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-07-18

**IP Status**

No known IP claims.

**Contributors**

- Aitor Camacho Larrondo, LunarG Inc.

## [](#_description)Description

An application may wish to reference device memory in multiple Vulkan device instances, in multiple processes, and/or in Metal API. This extension enables an application to export and import Metal handles from Vulkan memory objects such that the underlying resources can be referenced outside the scope of the Vulkan device instance that created them.

## [](#_new_commands)New Commands

- [vkGetMemoryMetalHandleEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryMetalHandleEXT.html)
- [vkGetMemoryMetalHandlePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryMetalHandlePropertiesEXT.html)

## [](#_new_structures)New Structures

- [VkMemoryGetMetalHandleInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetMetalHandleInfoEXT.html)
- [VkMemoryMetalHandlePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMetalHandlePropertiesEXT.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkImportMemoryMetalHandleInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryMetalHandleInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_EXTERNAL_MEMORY_METAL_EXTENSION_NAME`
- `VK_EXT_EXTERNAL_MEMORY_METAL_SPEC_VERSION`
- Extending [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html):
  
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_MTLBUFFER_BIT_EXT`
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_MTLHEAP_BIT_EXT`
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_MTLTEXTURE_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMPORT_MEMORY_METAL_HANDLE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_MEMORY_GET_METAL_HANDLE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_MEMORY_METAL_HANDLE_PROPERTIES_EXT`

## [](#_version_history)Version History

- Revision 1, 2024-07-18 (Aitor Camacho Larrondo)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_external_memory_metal).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0