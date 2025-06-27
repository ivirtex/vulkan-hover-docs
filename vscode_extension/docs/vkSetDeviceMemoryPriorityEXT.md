# vkSetDeviceMemoryPriorityEXT(3) Manual Page

## Name

vkSetDeviceMemoryPriorityEXT - Change a memory allocation priority



## <a href="#_c_specification" class="anchor"></a>C Specification

To modify the priority of an existing memory allocation, call:

``` c
// Provided by VK_EXT_pageable_device_local_memory
void vkSetDeviceMemoryPriorityEXT(
    VkDevice                                    device,
    VkDeviceMemory                              memory,
    float                                       priority);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the memory.

- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object to which
  the new priority will be applied.

- `priority` is a floating-point value between `0` and `1`, indicating
  the priority of the allocation relative to other memory allocations.
  Larger values are higher priority. The granularity of the priorities
  is implementation-dependent.

## <a href="#_description" class="anchor"></a>Description

Memory allocations with higher priority **may** be more likely to stay
in device-local memory when the system is under memory pressure.

Valid Usage

- <a href="#VUID-vkSetDeviceMemoryPriorityEXT-priority-06258"
  id="VUID-vkSetDeviceMemoryPriorityEXT-priority-06258"></a>
  VUID-vkSetDeviceMemoryPriorityEXT-priority-06258  
  `priority` **must** be between `0` and `1`, inclusive

Valid Usage (Implicit)

- <a href="#VUID-vkSetDeviceMemoryPriorityEXT-device-parameter"
  id="VUID-vkSetDeviceMemoryPriorityEXT-device-parameter"></a>
  VUID-vkSetDeviceMemoryPriorityEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkSetDeviceMemoryPriorityEXT-memory-parameter"
  id="VUID-vkSetDeviceMemoryPriorityEXT-memory-parameter"></a>
  VUID-vkSetDeviceMemoryPriorityEXT-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

- <a href="#VUID-vkSetDeviceMemoryPriorityEXT-memory-parent"
  id="VUID-vkSetDeviceMemoryPriorityEXT-memory-parent"></a>
  VUID-vkSetDeviceMemoryPriorityEXT-memory-parent  
  `memory` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_pageable_device_local_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pageable_device_local_memory.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkSetDeviceMemoryPriorityEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
