# VK_EXT_map_memory_placed(3) Manual Page

## Name

VK_EXT_map_memory_placed - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

273

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_map_memory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_map_memory2.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Faith Ekstrand <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_map_memory_placed%5D%20@gfxstrand%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_map_memory_placed%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>gfxstrand</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_map_memory_placed](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_map_memory_placed.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-03-21

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- Depends on apitext:VK_KHR_map_memory2

- Interacts with apitext:VK_EXT_external_memory_host

**Contributors**  
- Faith Ekstrand, Collabora

- Tobias Hector, AMD

- James Jones, NVIDIA

- Georg Lehmann, Valve

- Derek Lesho, Codeweavers

## <a href="#_description" class="anchor"></a>Description

This extension allows a client to request that
[vkMapMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMapMemory2KHR.html) attempt to place the memory map
at a particular virtual address.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkMemoryMapInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapInfoKHR.html):

  - [VkMemoryMapPlacedInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapPlacedInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceMapMemoryPlacedFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMapMemoryPlacedFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceMapMemoryPlacedPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMapMemoryPlacedPropertiesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_MAP_MEMORY_PLACED_EXTENSION_NAME`

- `VK_EXT_MAP_MEMORY_PLACED_SPEC_VERSION`

- Extending [VkMemoryMapFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapFlagBits.html):

  - `VK_MEMORY_MAP_PLACED_BIT_EXT`

- Extending [VkMemoryUnmapFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryUnmapFlagBitsKHR.html):

  - `VK_MEMORY_UNMAP_RESERVE_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_MEMORY_MAP_PLACED_INFO_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAP_MEMORY_PLACED_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAP_MEMORY_PLACED_PROPERTIES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 0, 2024-01-14 (Faith Ekstrand)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_map_memory_placed"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
