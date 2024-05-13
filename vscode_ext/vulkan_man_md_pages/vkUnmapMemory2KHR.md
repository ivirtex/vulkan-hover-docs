# vkUnmapMemory2KHR(3) Manual Page

## Name

vkUnmapMemory2KHR - Unmap a previously mapped memory object



## <a href="#_c_specification" class="anchor"></a>C Specification

Alternatively, to unmap a memory object once host access to it is no
longer needed by the application, call:

``` c
// Provided by VK_KHR_map_memory2
VkResult vkUnmapMemory2KHR(
    VkDevice                                    device,
    const VkMemoryUnmapInfoKHR*                 pMemoryUnmapInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the memory.

- `pMemoryUnmapInfo` is a pointer to a
  [VkMemoryUnmapInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryUnmapInfoKHR.html) structure describing
  parameters of the unmap.

## <a href="#_description" class="anchor"></a>Description

This function behaves identically to [vkUnmapMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUnmapMemory.html)
except that it gets its parameters via an extensible structure pointer
rather than directly as function arguments.

Valid Usage (Implicit)

- <a href="#VUID-vkUnmapMemory2KHR-device-parameter"
  id="VUID-vkUnmapMemory2KHR-device-parameter"></a>
  VUID-vkUnmapMemory2KHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkUnmapMemory2KHR-pMemoryUnmapInfo-parameter"
  id="VUID-vkUnmapMemory2KHR-pMemoryUnmapInfo-parameter"></a>
  VUID-vkUnmapMemory2KHR-pMemoryUnmapInfo-parameter  
  `pMemoryUnmapInfo` **must** be a valid pointer to a valid
  [VkMemoryUnmapInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryUnmapInfoKHR.html) structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_MEMORY_MAP_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_map_memory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_map_memory2.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMemoryUnmapInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryUnmapInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkUnmapMemory2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
