# VK\_EXT\_map\_memory\_placed(3) Manual Page

## Name

VK\_EXT\_map\_memory\_placed - device extension



## [](#_registered_extension_number)Registered Extension Number

273

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_map\_memory2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_map_memory2.html)  
or  
[Vulkan Version 1.4](#versions-1.4)

## [](#_contact)Contact

- Faith Ekstrand [\[GitHub\]gfxstrand](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_map_memory_placed%5D%20%40gfxstrand%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_map_memory_placed%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_map\_memory\_placed](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_map_memory_placed.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-03-21

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- Depends on apitext:VK\_KHR\_map\_memory2
- Interacts with apitext:VK\_EXT\_external\_memory\_host

**Contributors**

- Faith Ekstrand, Collabora
- Tobias Hector, AMD
- James Jones, NVIDIA
- Georg Lehmann, Valve
- Derek Lesho, Codeweavers

## [](#_description)Description

This extension allows an application to request that [vkMapMemory2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory2KHR.html) attempt to place the memory map at a particular virtual address.

## [](#_new_structures)New Structures

- Extending [VkMemoryMapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapInfo.html):
  
  - [VkMemoryMapPlacedInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapPlacedInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMapMemoryPlacedFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMapMemoryPlacedFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceMapMemoryPlacedPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMapMemoryPlacedPropertiesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_MAP_MEMORY_PLACED_EXTENSION_NAME`
- `VK_EXT_MAP_MEMORY_PLACED_SPEC_VERSION`
- Extending [VkMemoryMapFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapFlagBits.html):
  
  - `VK_MEMORY_MAP_PLACED_BIT_EXT`
- Extending [VkMemoryUnmapFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapFlagBits.html):
  
  - `VK_MEMORY_UNMAP_RESERVE_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_MEMORY_MAP_PLACED_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAP_MEMORY_PLACED_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAP_MEMORY_PLACED_PROPERTIES_EXT`

## [](#_version_history)Version History

- Revision 1, 2024-01-14 (Faith Ekstrand)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_map_memory_placed).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0