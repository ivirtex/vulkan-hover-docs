# vkSetDeviceMemoryPriorityEXT(3) Manual Page

## Name

vkSetDeviceMemoryPriorityEXT - Change a memory allocation priority



## [](#_c_specification)C Specification

To modify the priority of an existing memory allocation, call:

```c++
// Provided by VK_EXT_pageable_device_local_memory
void vkSetDeviceMemoryPriorityEXT(
    VkDevice                                    device,
    VkDeviceMemory                              memory,
    float                                       priority);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the memory.
- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object to which the new priority will be applied.
- `priority` is a floating-point value between `0` and `1`, indicating the priority of the allocation relative to other memory allocations. Larger values are higher priority. The granularity of the priorities is implementation-dependent.

## [](#_description)Description

Memory allocations with higher priority **may** be more likely to stay in device-local memory when the system is under memory pressure.

Valid Usage

- [](#VUID-vkSetDeviceMemoryPriorityEXT-priority-06258)VUID-vkSetDeviceMemoryPriorityEXT-priority-06258  
  `priority` **must** be between `0` and `1`, inclusive

Valid Usage (Implicit)

- [](#VUID-vkSetDeviceMemoryPriorityEXT-device-parameter)VUID-vkSetDeviceMemoryPriorityEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkSetDeviceMemoryPriorityEXT-memory-parameter)VUID-vkSetDeviceMemoryPriorityEXT-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-vkSetDeviceMemoryPriorityEXT-memory-parent)VUID-vkSetDeviceMemoryPriorityEXT-memory-parent  
  `memory` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_EXT\_pageable\_device\_local\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pageable_device_local_memory.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkSetDeviceMemoryPriorityEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0