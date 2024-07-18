# VkImageSparseMemoryRequirementsInfo2(3) Manual Page

## Name

VkImageSparseMemoryRequirementsInfo2 - (None)



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageSparseMemoryRequirementsInfo2` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkImageSparseMemoryRequirementsInfo2 {
    VkStructureType    sType;
    const void*        pNext;
    VkImage            image;
} VkImageSparseMemoryRequirementsInfo2;
```

or the equivalent

``` c
// Provided by VK_KHR_get_memory_requirements2
typedef VkImageSparseMemoryRequirementsInfo2 VkImageSparseMemoryRequirementsInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `image` is the image to query.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkImageSparseMemoryRequirementsInfo2-sType-sType"
  id="VUID-VkImageSparseMemoryRequirementsInfo2-sType-sType"></a>
  VUID-VkImageSparseMemoryRequirementsInfo2-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2`

- <a href="#VUID-VkImageSparseMemoryRequirementsInfo2-pNext-pNext"
  id="VUID-VkImageSparseMemoryRequirementsInfo2-pNext-pNext"></a>
  VUID-VkImageSparseMemoryRequirementsInfo2-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkImageSparseMemoryRequirementsInfo2-image-parameter"
  id="VUID-VkImageSparseMemoryRequirementsInfo2-image-parameter"></a>
  VUID-VkImageSparseMemoryRequirementsInfo2-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetImageSparseMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSparseMemoryRequirements2.html),
[vkGetImageSparseMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSparseMemoryRequirements2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageSparseMemoryRequirementsInfo2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
