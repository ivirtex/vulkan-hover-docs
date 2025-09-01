# vkUnmapMemory(3) Manual Page

## Name

vkUnmapMemory - Unmap a previously mapped memory object



## [](#_c_specification)C Specification

To unmap a memory object once host access to it is no longer needed by the application, call:

```c++
// Provided by VK_VERSION_1_0
void vkUnmapMemory(
    VkDevice                                    device,
    VkDeviceMemory                              memory);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the memory.
- `memory` is the memory object to be unmapped.

## [](#_description)Description

Calling `vkUnmapMemory` is equivalent to calling [vkUnmapMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUnmapMemory2.html) with an empty `pNext` chain and `flags` set to zero.

Valid Usage

- [](#VUID-vkUnmapMemory-memory-00689)VUID-vkUnmapMemory-memory-00689  
  `memory` **must** be currently host mapped

Valid Usage (Implicit)

- [](#VUID-vkUnmapMemory-device-parameter)VUID-vkUnmapMemory-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkUnmapMemory-memory-parameter)VUID-vkUnmapMemory-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-vkUnmapMemory-memory-parent)VUID-vkUnmapMemory-memory-parent  
  `memory` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `memory` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkUnmapMemory).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0