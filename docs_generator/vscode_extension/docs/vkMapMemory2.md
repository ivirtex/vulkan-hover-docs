# vkMapMemory2(3) Manual Page

## Name

vkMapMemory2 - Map a memory object into application address space



## [](#_c_specification)C Specification

Alternatively, to retrieve a host virtual address pointer to a region of a mappable memory object, call:

```c++
// Provided by VK_VERSION_1_4
VkResult vkMapMemory2(
    VkDevice                                    device,
    const VkMemoryMapInfo*                      pMemoryMapInfo,
    void**                                      ppData);
```

or the equivalent command

```c++
// Provided by VK_KHR_map_memory2
VkResult vkMapMemory2KHR(
    VkDevice                                    device,
    const VkMemoryMapInfo*                      pMemoryMapInfo,
    void**                                      ppData);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the memory.
- `pMemoryMapInfo` is a pointer to a [VkMemoryMapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapInfo.html) structure describing parameters of the map.
- `ppData` is a pointer to a `void *` variable in which is returned a host-accessible pointer to the beginning of the mapped range. This pointer minus [VkMemoryMapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapInfo.html)::`offset` **must** be aligned to at least [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`minMemoryMapAlignment`.

## [](#_description)Description

This function behaves identically to [vkMapMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory.html) except that it gets its parameters via an extensible structure pointer rather than directly as function arguments.

Valid Usage (Implicit)

- [](#VUID-vkMapMemory2-device-parameter)VUID-vkMapMemory2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkMapMemory2-pMemoryMapInfo-parameter)VUID-vkMapMemory2-pMemoryMapInfo-parameter  
  `pMemoryMapInfo` **must** be a valid pointer to a valid [VkMemoryMapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapInfo.html) structure
- [](#VUID-vkMapMemory2-ppData-parameter)VUID-vkMapMemory2-ppData-parameter  
  `ppData` **must** be a valid pointer to a pointer value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_MEMORY_MAP_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_map\_memory2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_map_memory2.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMemoryMapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkMapMemory2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0