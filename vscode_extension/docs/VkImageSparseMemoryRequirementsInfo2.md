# VkImageSparseMemoryRequirementsInfo2(3) Manual Page

## Name

VkImageSparseMemoryRequirementsInfo2 - (None)



## [](#_c_specification)C Specification

The `VkImageSparseMemoryRequirementsInfo2` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkImageSparseMemoryRequirementsInfo2 {
    VkStructureType    sType;
    const void*        pNext;
    VkImage            image;
} VkImageSparseMemoryRequirementsInfo2;
```

or the equivalent

```c++
// Provided by VK_KHR_get_memory_requirements2
typedef VkImageSparseMemoryRequirementsInfo2 VkImageSparseMemoryRequirementsInfo2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `image` is the image to query.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkImageSparseMemoryRequirementsInfo2-sType-sType)VUID-VkImageSparseMemoryRequirementsInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2`
- [](#VUID-VkImageSparseMemoryRequirementsInfo2-pNext-pNext)VUID-VkImageSparseMemoryRequirementsInfo2-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImageSparseMemoryRequirementsInfo2-image-parameter)VUID-VkImageSparseMemoryRequirementsInfo2-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetImageSparseMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSparseMemoryRequirements2.html), [vkGetImageSparseMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSparseMemoryRequirements2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageSparseMemoryRequirementsInfo2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0