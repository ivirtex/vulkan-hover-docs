# VK\_EXT\_memory\_priority(3) Manual Page

## Name

VK\_EXT\_memory\_priority - device extension



## [](#_registered_extension_number)Registered Extension Number

239

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_memory_priority%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_memory_priority%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-10-08

**Contributors**

- Jeff Bolz, NVIDIA
- Jeff Juliano, NVIDIA

## [](#_description)Description

This extension adds a `priority` value specified at memory allocation time. On some systems with both device-local and non-device-local memory heaps, the implementation may transparently move memory from one heap to another when a heap becomes full (for example, when the total memory used across all processes exceeds the size of the heap). In such a case, this priority value may be used to determine which allocations are more likely to remain in device-local memory.

## [](#_new_structures)New Structures

- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkMemoryPriorityAllocateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryPriorityAllocateInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMemoryPriorityFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryPriorityFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_MEMORY_PRIORITY_EXTENSION_NAME`
- `VK_EXT_MEMORY_PRIORITY_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_MEMORY_PRIORITY_ALLOCATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PRIORITY_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2018-10-08 (Jeff Bolz)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_memory_priority)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0