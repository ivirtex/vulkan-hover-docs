# vkMapMemory2KHR(3) Manual Page

## Name

vkMapMemory2KHR - Map a memory object into application address space



## <a href="#_c_specification" class="anchor"></a>C Specification

Alternatively, to retrieve a host virtual address pointer to a region of
a mappable memory object, call:

``` c
// Provided by VK_KHR_map_memory2
VkResult vkMapMemory2KHR(
    VkDevice                                    device,
    const VkMemoryMapInfoKHR*                   pMemoryMapInfo,
    void**                                      ppData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the memory.

- `pMemoryMapInfo` is a pointer to a
  [VkMemoryMapInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapInfoKHR.html) structure describing
  parameters of the map.

- `ppData` is a pointer to a `void *` variable in which is returned a
  host-accessible pointer to the beginning of the mapped range. This
  pointer minus [VkMemoryMapInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapInfoKHR.html)::`offset`
  **must** be aligned to at least
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`minMemoryMapAlignment`.

## <a href="#_description" class="anchor"></a>Description

This function behaves identically to [vkMapMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMapMemory.html)
except that it gets its parameters via an extensible structure pointer
rather than directly as function arguments.

Valid Usage (Implicit)

- <a href="#VUID-vkMapMemory2KHR-device-parameter"
  id="VUID-vkMapMemory2KHR-device-parameter"></a>
  VUID-vkMapMemory2KHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkMapMemory2KHR-pMemoryMapInfo-parameter"
  id="VUID-vkMapMemory2KHR-pMemoryMapInfo-parameter"></a>
  VUID-vkMapMemory2KHR-pMemoryMapInfo-parameter  
  `pMemoryMapInfo` **must** be a valid pointer to a valid
  [VkMemoryMapInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapInfoKHR.html) structure

- <a href="#VUID-vkMapMemory2KHR-ppData-parameter"
  id="VUID-vkMapMemory2KHR-ppData-parameter"></a>
  VUID-vkMapMemory2KHR-ppData-parameter  
  `ppData` **must** be a valid pointer to a pointer value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_MEMORY_MAP_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_map_memory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_map_memory2.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkMemoryMapInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkMapMemory2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
