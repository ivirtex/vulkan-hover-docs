# VkSparseImageMemoryRequirements2(3) Manual Page

## Name

VkSparseImageMemoryRequirements2 - (None)



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSparseImageMemoryRequirements2` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkSparseImageMemoryRequirements2 {
    VkStructureType                    sType;
    void*                              pNext;
    VkSparseImageMemoryRequirements    memoryRequirements;
} VkSparseImageMemoryRequirements2;
```

or the equivalent

``` c
// Provided by VK_KHR_get_memory_requirements2
typedef VkSparseImageMemoryRequirements2 VkSparseImageMemoryRequirements2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memoryRequirements` is a
  [VkSparseImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements.html)
  structure describing the memory requirements of the sparse image.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkSparseImageMemoryRequirements2-sType-sType"
  id="VUID-VkSparseImageMemoryRequirements2-sType-sType"></a>
  VUID-VkSparseImageMemoryRequirements2-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2`

- <a href="#VUID-VkSparseImageMemoryRequirements2-pNext-pNext"
  id="VUID-VkSparseImageMemoryRequirements2-pNext-pNext"></a>
  VUID-VkSparseImageMemoryRequirements2-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkSparseImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDeviceImageSparseMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageSparseMemoryRequirements.html),
[vkGetDeviceImageSparseMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageSparseMemoryRequirementsKHR.html),
[vkGetImageSparseMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSparseMemoryRequirements2.html),
[vkGetImageSparseMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSparseMemoryRequirements2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSparseImageMemoryRequirements2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
