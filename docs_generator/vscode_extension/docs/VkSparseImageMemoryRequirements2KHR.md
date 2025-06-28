# VkSparseImageMemoryRequirements2(3) Manual Page

## Name

VkSparseImageMemoryRequirements2 - (None)



## [](#_c_specification)C Specification

The `VkSparseImageMemoryRequirements2` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkSparseImageMemoryRequirements2 {
    VkStructureType                    sType;
    void*                              pNext;
    VkSparseImageMemoryRequirements    memoryRequirements;
} VkSparseImageMemoryRequirements2;
```

or the equivalent

```c++
// Provided by VK_KHR_get_memory_requirements2
typedef VkSparseImageMemoryRequirements2 VkSparseImageMemoryRequirements2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memoryRequirements` is a [VkSparseImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryRequirements.html) structure describing the memory requirements of the sparse image.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkSparseImageMemoryRequirements2-sType-sType)VUID-VkSparseImageMemoryRequirements2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2`
- [](#VUID-VkSparseImageMemoryRequirements2-pNext-pNext)VUID-VkSparseImageMemoryRequirements2-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkSparseImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryRequirements.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDeviceImageSparseMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSparseMemoryRequirements.html), [vkGetDeviceImageSparseMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSparseMemoryRequirementsKHR.html), [vkGetImageSparseMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSparseMemoryRequirements2.html), [vkGetImageSparseMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSparseMemoryRequirements2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSparseImageMemoryRequirements2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0