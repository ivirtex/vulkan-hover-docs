# VkMemoryWin32HandlePropertiesKHR(3) Manual Page

## Name

VkMemoryWin32HandlePropertiesKHR - Properties of External Memory Windows Handles



## [](#_c_specification)C Specification

The `VkMemoryWin32HandlePropertiesKHR` structure returned is defined as:

```c++
// Provided by VK_KHR_external_memory_win32
typedef struct VkMemoryWin32HandlePropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           memoryTypeBits;
} VkMemoryWin32HandlePropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memoryTypeBits` is a bitmask containing one bit set for every memory type which the specified windows handle **can** be imported as.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkMemoryWin32HandlePropertiesKHR-sType-sType)VUID-VkMemoryWin32HandlePropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR`
- [](#VUID-VkMemoryWin32HandlePropertiesKHR-pNext-pNext)VUID-VkMemoryWin32HandlePropertiesKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_external\_memory\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_win32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetMemoryWin32HandlePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryWin32HandlePropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryWin32HandlePropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0