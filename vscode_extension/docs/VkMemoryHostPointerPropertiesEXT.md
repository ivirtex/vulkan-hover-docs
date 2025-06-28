# VkMemoryHostPointerPropertiesEXT(3) Manual Page

## Name

VkMemoryHostPointerPropertiesEXT - Properties of external memory host pointer



## [](#_c_specification)C Specification

The `VkMemoryHostPointerPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_external_memory_host
typedef struct VkMemoryHostPointerPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           memoryTypeBits;
} VkMemoryHostPointerPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memoryTypeBits` is a bitmask containing one bit set for every memory type which the specified host pointer **can** be imported as.

## [](#_description)Description

The value returned by `memoryTypeBits` **should** only include bits that identify memory types which are host visible. Implementations **may** include bits that identify memory types which are not host visible. Behavior for imported pointers of such types is defined by [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#host-memory-import-non-visible-type)[VkImportMemoryHostPointerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryHostPointerInfoEXT.html).

Valid Usage (Implicit)

- [](#VUID-VkMemoryHostPointerPropertiesEXT-sType-sType)VUID-VkMemoryHostPointerPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_HOST_POINTER_PROPERTIES_EXT`
- [](#VUID-VkMemoryHostPointerPropertiesEXT-pNext-pNext)VUID-VkMemoryHostPointerPropertiesEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_EXT\_external\_memory\_host](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_external_memory_host.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetMemoryHostPointerPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryHostPointerPropertiesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryHostPointerPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0