# VK_EXT_pageable_device_local_memory(3) Manual Page

## Name

VK_EXT_pageable_device_local_memory - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

413

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_EXT_memory_priority](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_memory_priority.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_pageable_device_local_memory%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_pageable_device_local_memory%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

Vulkan is frequently implemented on multi-user and multi-process
operating systems where the device-local memory can be shared by more
than one process. On such systems the size of the device-local memory
available to the application may not be the full size of the memory heap
at all times. In order for these operating systems to support multiple
applications the device-local memory is virtualized and paging is used
to move memory between device-local and host-local memory heaps,
transparent to the application.

The current Vulkan specification does not expose this behavior well and
may cause applications to make suboptimal memory choices when allocating
memory. For example, in a system with multiple applications running, the
application may think that device-local memory is full and revert to
making performance-sensitive allocations from host-local memory. In
reality the memory heap might not have been full, it just appeared to be
at the time memory consumption was queried, and a device-local
allocation would have succeeded. A well designed operating system that
implements pageable device-local memory will try to have all memory
allocations for the foreground application paged into device-local
memory, and paged out for other applications as needed when not in use.

When this extension is exposed by the Vulkan implementation it indicates
to the application that the operating system implements pageable
device-local memory and the application should adjust its memory
allocation strategy accordingly. The extension also exposes a new
[vkSetDeviceMemoryPriorityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetDeviceMemoryPriorityEXT.html)
function to allow the application to dynamically adjust the priority of
existing memory allocations based on its current needs. This will help
the operating system page out lower priority memory allocations before
higher priority allocations when needed. It will also help the operating
system decide which memory allocations to page back into device-local
memory first.

To take best advantage of pageable device-local memory the application
must create the Vulkan device with the
[VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT.html)::`pageableDeviceLocalMemory`
feature enabled. When enabled the Vulkan implementation will allow
device-local memory allocations to be paged in and out by the operating
system, and **may** not return VK_ERROR_OUT_OF_DEVICE_MEMORY even if
device-local memory appears to be full, but will instead page this, or
other allocations, out to make room. The Vulkan implementation will also
ensure that host-local memory allocations will never be promoted to
device-local memory by the operating system, or consume device-local
memory.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkSetDeviceMemoryPriorityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetDeviceMemoryPriorityEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_PAGEABLE_DEVICE_LOCAL_MEMORY_EXTENSION_NAME`

- `VK_EXT_PAGEABLE_DEVICE_LOCAL_MEMORY_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PAGEABLE_DEVICE_LOCAL_MEMORY_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-08-24 (Piers Daniell)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_pageable_device_local_memory"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
