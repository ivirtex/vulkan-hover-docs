# VkVideoSessionMemoryRequirementsKHR(3) Manual Page

## Name

VkVideoSessionMemoryRequirementsKHR - Structure describing video session
memory requirements



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoSessionMemoryRequirementsKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_queue
typedef struct VkVideoSessionMemoryRequirementsKHR {
    VkStructureType         sType;
    void*                   pNext;
    uint32_t                memoryBindIndex;
    VkMemoryRequirements    memoryRequirements;
} VkVideoSessionMemoryRequirementsKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memoryBindIndex` is the index of the memory binding.

- `memoryRequirements` is a
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure in which
  the requested memory binding requirements for the binding index
  specified by `memoryBindIndex` are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoSessionMemoryRequirementsKHR-sType-sType"
  id="VUID-VkVideoSessionMemoryRequirementsKHR-sType-sType"></a>
  VUID-VkVideoSessionMemoryRequirementsKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_SESSION_MEMORY_REQUIREMENTS_KHR`

- <a href="#VUID-VkVideoSessionMemoryRequirementsKHR-pNext-pNext"
  id="VUID-VkVideoSessionMemoryRequirementsKHR-pNext-pNext"></a>
  VUID-VkVideoSessionMemoryRequirementsKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetVideoSessionMemoryRequirementsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoSessionMemoryRequirementsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
