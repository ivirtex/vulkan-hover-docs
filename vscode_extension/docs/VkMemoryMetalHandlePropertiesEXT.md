# VkMemoryMetalHandlePropertiesEXT(3) Manual Page

## Name

VkMemoryMetalHandlePropertiesEXT - Properties of External Memory Metal Handles



## [](#_c_specification)C Specification

The `VkMemoryMetalHandlePropertiesEXT` structure returned is defined as:

```c++
// Provided by VK_EXT_external_memory_metal
typedef struct VkMemoryMetalHandlePropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           memoryTypeBits;
} VkMemoryMetalHandlePropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memoryTypeBits` is a bitmask containing one bit set for every memory type which the specified Metal handle **can** be imported as.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkMemoryMetalHandlePropertiesEXT-sType-sType)VUID-VkMemoryMetalHandlePropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_METAL_HANDLE_PROPERTIES_EXT`
- [](#VUID-VkMemoryMetalHandlePropertiesEXT-pNext-pNext)VUID-VkMemoryMetalHandlePropertiesEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_EXT\_external\_memory\_metal](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_external_memory_metal.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetMemoryMetalHandlePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryMetalHandlePropertiesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryMetalHandlePropertiesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0