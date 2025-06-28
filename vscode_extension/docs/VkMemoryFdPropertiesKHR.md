# VkMemoryFdPropertiesKHR(3) Manual Page

## Name

VkMemoryFdPropertiesKHR - Properties of External Memory File Descriptors



## [](#_c_specification)C Specification

The `VkMemoryFdPropertiesKHR` structure returned is defined as:

```c++
// Provided by VK_KHR_external_memory_fd
typedef struct VkMemoryFdPropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           memoryTypeBits;
} VkMemoryFdPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memoryTypeBits` is a bitmask containing one bit set for every memory type which the specified file descriptor **can** be imported as.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkMemoryFdPropertiesKHR-sType-sType)VUID-VkMemoryFdPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_FD_PROPERTIES_KHR`
- [](#VUID-VkMemoryFdPropertiesKHR-pNext-pNext)VUID-VkMemoryFdPropertiesKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_external\_memory\_fd](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_fd.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetMemoryFdPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryFdPropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryFdPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0