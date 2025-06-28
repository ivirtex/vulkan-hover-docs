# vkUnmapMemory2(3) Manual Page

## Name

vkUnmapMemory2 - Unmap a previously mapped memory object



## [](#_c_specification)C Specification

Alternatively, to unmap a memory object once host access to it is no longer needed by the application, call:

```c++
// Provided by VK_VERSION_1_4
VkResult vkUnmapMemory2(
    VkDevice                                    device,
    const VkMemoryUnmapInfo*                    pMemoryUnmapInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_map_memory2
VkResult vkUnmapMemory2KHR(
    VkDevice                                    device,
    const VkMemoryUnmapInfo*                    pMemoryUnmapInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the memory.
- `pMemoryUnmapInfo` is a pointer to a [VkMemoryUnmapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapInfo.html) structure describing parameters of the unmap.

## [](#_description)Description

This function behaves identically to [vkUnmapMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUnmapMemory.html) except that it gets its parameters via an extensible structure pointer rather than directly as function arguments.

Valid Usage (Implicit)

- [](#VUID-vkUnmapMemory2-device-parameter)VUID-vkUnmapMemory2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkUnmapMemory2-pMemoryUnmapInfo-parameter)VUID-vkUnmapMemory2-pMemoryUnmapInfo-parameter  
  `pMemoryUnmapInfo` **must** be a valid pointer to a valid [VkMemoryUnmapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapInfo.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_MEMORY_MAP_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_map\_memory2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_map_memory2.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMemoryUnmapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkUnmapMemory2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0