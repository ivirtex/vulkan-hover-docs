# VkMemoryFdPropertiesKHR(3) Manual Page

## Name

VkMemoryFdPropertiesKHR - Properties of External Memory File Descriptors



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryFdPropertiesKHR` structure returned is defined as:

``` c
// Provided by VK_KHR_external_memory_fd
typedef struct VkMemoryFdPropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           memoryTypeBits;
} VkMemoryFdPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memoryTypeBits` is a bitmask containing one bit set for every memory
  type which the specified file descriptor **can** be imported as.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryFdPropertiesKHR-sType-sType"
  id="VUID-VkMemoryFdPropertiesKHR-sType-sType"></a>
  VUID-VkMemoryFdPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_FD_PROPERTIES_KHR`

- <a href="#VUID-VkMemoryFdPropertiesKHR-pNext-pNext"
  id="VUID-VkMemoryFdPropertiesKHR-pNext-pNext"></a>
  VUID-VkMemoryFdPropertiesKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_memory_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_fd.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetMemoryFdPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryFdPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryFdPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
