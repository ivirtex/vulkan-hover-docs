# VkBufferMemoryRequirementsInfo2(3) Manual Page

## Name

VkBufferMemoryRequirementsInfo2 - (None)



## [](#_c_specification)C Specification

The `VkBufferMemoryRequirementsInfo2` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkBufferMemoryRequirementsInfo2 {
    VkStructureType    sType;
    const void*        pNext;
    VkBuffer           buffer;
} VkBufferMemoryRequirementsInfo2;
```

or the equivalent

```c++
// Provided by VK_KHR_get_memory_requirements2
typedef VkBufferMemoryRequirementsInfo2 VkBufferMemoryRequirementsInfo2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `buffer` is the buffer to query.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkBufferMemoryRequirementsInfo2-sType-sType)VUID-VkBufferMemoryRequirementsInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2`
- [](#VUID-VkBufferMemoryRequirementsInfo2-pNext-pNext)VUID-VkBufferMemoryRequirementsInfo2-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkBufferMemoryRequirementsInfo2-buffer-parameter)VUID-VkBufferMemoryRequirementsInfo2-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetBufferMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferMemoryRequirements2.html), [vkGetBufferMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferMemoryRequirements2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferMemoryRequirementsInfo2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0