# VK\_EXT\_pageable\_device\_local\_memory(3) Manual Page

## Name

VK\_EXT\_pageable\_device\_local\_memory - device extension



## [](#_registered_extension_number)Registered Extension Number

413

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_memory\_priority](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_memory_priority.html)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_pageable_device_local_memory%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_pageable_device_local_memory%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-08-24

**Contributors**

- Hans-Kristian Arntzen, Valve
- Axel Gneiting, id Software
- Billy Khan, id Software
- Daniel Koch, NVIDIA
- Chris Lentini, NVIDIA
- Joshua Schnarr, NVIDIA
- Stu Smith, AMD

## [](#_description)Description

Vulkan is frequently implemented on multi-user and multi-process operating systems where the device-local memory can be shared by more than one process. On such systems the size of the device-local memory available to the application may not be the full size of the memory heap at all times. In order for these operating systems to support multiple applications the device-local memory is virtualized and paging is used to move memory between device-local and host-local memory heaps, transparent to the application.

The current Vulkan specification does not expose this behavior well and may cause applications to make suboptimal memory choices when allocating memory. For example, in a system with multiple applications running, the application may think that device-local memory is full and revert to making performance-sensitive allocations from host-local memory. In reality the memory heap might not have been full, it just appeared to be at the time memory consumption was queried, and a device-local allocation would have succeeded. A well designed operating system that implements pageable device-local memory will try to have all memory allocations for the foreground application paged into device-local memory, and paged out for other applications as needed when not in use.

When this extension is exposed by the Vulkan implementation it indicates to the application that the operating system implements pageable device-local memory and the application should adjust its memory allocation strategy accordingly. The extension also exposes a new [vkSetDeviceMemoryPriorityEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetDeviceMemoryPriorityEXT.html) function to allow the application to dynamically adjust the priority of existing memory allocations based on its current needs. This will help the operating system page out lower priority memory allocations before higher priority allocations when needed. It will also help the operating system decide which memory allocations to page back into device-local memory first.

To take best advantage of pageable device-local memory the application must create the Vulkan device with the [VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT.html)::`pageableDeviceLocalMemory` feature enabled. When enabled the Vulkan implementation will allow device-local memory allocations to be paged in and out by the operating system, and **may** not return VK\_ERROR\_OUT\_OF\_DEVICE\_MEMORY even if device-local memory appears to be full, but will instead page this, or other allocations, out to make room. The Vulkan implementation will also ensure that host-local memory allocations will never be promoted to device-local memory by the operating system, or consume device-local memory.

## [](#_new_commands)New Commands

- [vkSetDeviceMemoryPriorityEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetDeviceMemoryPriorityEXT.html)

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_PAGEABLE_DEVICE_LOCAL_MEMORY_EXTENSION_NAME`
- `VK_EXT_PAGEABLE_DEVICE_LOCAL_MEMORY_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PAGEABLE_DEVICE_LOCAL_MEMORY_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2021-08-24 (Piers Daniell)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_pageable_device_local_memory)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0