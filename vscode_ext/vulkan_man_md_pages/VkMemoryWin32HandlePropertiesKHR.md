# VkMemoryWin32HandlePropertiesKHR(3) Manual Page

## Name

VkMemoryWin32HandlePropertiesKHR - Properties of External Memory Windows
Handles



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryWin32HandlePropertiesKHR` structure returned is defined as:

``` c
// Provided by VK_KHR_external_memory_win32
typedef struct VkMemoryWin32HandlePropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           memoryTypeBits;
} VkMemoryWin32HandlePropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memoryTypeBits` is a bitmask containing one bit set for every memory
  type which the specified windows handle **can** be imported as.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryWin32HandlePropertiesKHR-sType-sType"
  id="VUID-VkMemoryWin32HandlePropertiesKHR-sType-sType"></a>
  VUID-VkMemoryWin32HandlePropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR`

- <a href="#VUID-VkMemoryWin32HandlePropertiesKHR-pNext-pNext"
  id="VUID-VkMemoryWin32HandlePropertiesKHR-pNext-pNext"></a>
  VUID-VkMemoryWin32HandlePropertiesKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_memory_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_win32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetMemoryWin32HandlePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryWin32HandlePropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryWin32HandlePropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
