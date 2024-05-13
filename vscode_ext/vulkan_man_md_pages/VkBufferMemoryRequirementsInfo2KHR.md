# VkBufferMemoryRequirementsInfo2(3) Manual Page

## Name

VkBufferMemoryRequirementsInfo2 - (None)



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBufferMemoryRequirementsInfo2` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkBufferMemoryRequirementsInfo2 {
    VkStructureType    sType;
    const void*        pNext;
    VkBuffer           buffer;
} VkBufferMemoryRequirementsInfo2;
```

or the equivalent

``` c
// Provided by VK_KHR_get_memory_requirements2
typedef VkBufferMemoryRequirementsInfo2 VkBufferMemoryRequirementsInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `buffer` is the buffer to query.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkBufferMemoryRequirementsInfo2-sType-sType"
  id="VUID-VkBufferMemoryRequirementsInfo2-sType-sType"></a>
  VUID-VkBufferMemoryRequirementsInfo2-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2`

- <a href="#VUID-VkBufferMemoryRequirementsInfo2-pNext-pNext"
  id="VUID-VkBufferMemoryRequirementsInfo2-pNext-pNext"></a>
  VUID-VkBufferMemoryRequirementsInfo2-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkBufferMemoryRequirementsInfo2-buffer-parameter"
  id="VUID-VkBufferMemoryRequirementsInfo2-buffer-parameter"></a>
  VUID-VkBufferMemoryRequirementsInfo2-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetBufferMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements2.html),
[vkGetBufferMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferMemoryRequirementsInfo2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
