# vkUnmapMemory(3) Manual Page

## Name

vkUnmapMemory - Unmap a previously mapped memory object



## <a href="#_c_specification" class="anchor"></a>C Specification

To unmap a memory object once host access to it is no longer needed by
the application, call:

``` c
// Provided by VK_VERSION_1_0
void vkUnmapMemory(
    VkDevice                                    device,
    VkDeviceMemory                              memory);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the memory.

- `memory` is the memory object to be unmapped.

## <a href="#_description" class="anchor"></a>Description

Calling `vkUnmapMemory` is equivalent to calling
[vkUnmapMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUnmapMemory2KHR.html) with an empty `pNext` chain
and the flags parameter set to zero.

Valid Usage

- <a href="#VUID-vkUnmapMemory-memory-00689"
  id="VUID-vkUnmapMemory-memory-00689"></a>
  VUID-vkUnmapMemory-memory-00689  
  `memory` **must** be currently host mapped

Valid Usage (Implicit)

- <a href="#VUID-vkUnmapMemory-device-parameter"
  id="VUID-vkUnmapMemory-device-parameter"></a>
  VUID-vkUnmapMemory-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkUnmapMemory-memory-parameter"
  id="VUID-vkUnmapMemory-memory-parameter"></a>
  VUID-vkUnmapMemory-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

- <a href="#VUID-vkUnmapMemory-memory-parent"
  id="VUID-vkUnmapMemory-memory-parent"></a>
  VUID-vkUnmapMemory-memory-parent  
  `memory` **must** have been created, allocated, or retrieved from
  `device`

Host Synchronization

- Host access to `memory` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkUnmapMemory"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
