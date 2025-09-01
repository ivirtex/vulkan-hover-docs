# VK\_QCOM\_tile\_memory\_heap(3) Manual Page

## Name

VK\_QCOM\_tile\_memory\_heap - device extension



## [](#_registered_extension_number)Registered Extension Number

548

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_memory\_requirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_memory_requirements2.html)  
     and  
     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_QCOM\_tile\_properties

## [](#_contact)Contact

- Patrick Boyle [\[GitHub\]pboyleQCOM](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_tile_memory_heap%5D%20%40pboyleQCOM%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_tile_memory_heap%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_QCOM\_tile\_memory\_heap](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_QCOM_tile_memory_heap.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-05-05

**Interactions and External Dependencies**

- Interacts with `VK_QCOM_tile_properties`
- Interacts with `VK_QCOM_tile_shading`

**Contributors**

- Patrick Boyle, Qualcomm Technologies, Inc.
- Matthew Netsch, Qualcomm Technologies, Inc.
- Srihari Babu Alla, Qualcomm Technologies, Inc.
- Kevin Matlage, Qualcomm Technologies, Inc.

## [](#_description)Description

This extension adds a new memory heap which allows applications to allocate and manage tile memory. A tile memory heap is denoted by the new `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property. Memory contents within this heap behave differently than other heaps and only persist its memory contents within a command buffer submission batch boundary. This boundary may be extended to a queue submit boundary by querying `queueSubmitBoundary` in the new `VkPhysicalDeviceTileMemoryHeapPropertiesQCOM` structure.

Tile memory from this heap can be bound to VkImages or VkBuffers. The following new usage flags `VK_IMAGE_USAGE_TILE_MEMORY_BIT_QCOM`, `VK_BUFFER_USAGE_TILE_MEMORY_BIT_QCOM`, `VK_BUFFER_USAGE_2_TILE_MEMORY_BIT_QCOM` were added for this. A new extended structure is added to get memory requirements for tile memory `VkTileMemoryRequirementsQCOM`.

A new command is added to bind tile memory to a command buffer in order to access and persist tile memory contents while executing commands [vkCmdBindTileMemoryQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindTileMemoryQCOM.html).

This extension can be used in combination with [VK\_QCOM\_tile\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_properties.html) with the new structure [VkTileMemorySizeInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileMemorySizeInfoQCOM.html).

## [](#_issues)Issues

None.

## [](#_new_commands)New Commands

- [vkCmdBindTileMemoryQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindTileMemoryQCOM.html)

## [](#_new_structures)New Structures

- Extending [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html):
  
  - [VkTileMemoryBindInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileMemoryBindInfoQCOM.html)
- Extending [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html):
  
  - [VkTileMemoryRequirementsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileMemoryRequirementsQCOM.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceTileMemoryHeapFeaturesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTileMemoryHeapFeaturesQCOM.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceTileMemoryHeapPropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTileMemoryHeapPropertiesQCOM.html)

If [VK\_QCOM\_tile\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_properties.html) is supported:

- Extending [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html), [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html), [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html):
  
  - [VkTileMemorySizeInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileMemorySizeInfoQCOM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_QCOM_TILE_MEMORY_HEAP_EXTENSION_NAME`
- `VK_QCOM_TILE_MEMORY_HEAP_SPEC_VERSION`
- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html):
  
  - `VK_BUFFER_USAGE_TILE_MEMORY_BIT_QCOM`
- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_TILE_MEMORY_BIT_QCOM`
- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html):
  
  - `VK_IMAGE_USAGE_TILE_MEMORY_BIT_QCOM`
- Extending [VkMemoryHeapFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHeapFlagBits.html):
  
  - `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TILE_MEMORY_HEAP_FEATURES_QCOM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TILE_MEMORY_HEAP_PROPERTIES_QCOM`
  - `VK_STRUCTURE_TYPE_TILE_MEMORY_BIND_INFO_QCOM`
  - `VK_STRUCTURE_TYPE_TILE_MEMORY_REQUIREMENTS_QCOM`

If [VK\_QCOM\_tile\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_properties.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_TILE_MEMORY_SIZE_INFO_QCOM`

## [](#_version_history)Version History

- Revision 1, 2025-03-26 (Patrick Boyle)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QCOM_tile_memory_heap).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0