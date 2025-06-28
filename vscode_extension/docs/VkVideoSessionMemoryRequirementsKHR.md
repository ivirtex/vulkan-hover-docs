# VkVideoSessionMemoryRequirementsKHR(3) Manual Page

## Name

VkVideoSessionMemoryRequirementsKHR - Structure describing video session memory requirements



## [](#_c_specification)C Specification

The `VkVideoSessionMemoryRequirementsKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkVideoSessionMemoryRequirementsKHR {
    VkStructureType         sType;
    void*                   pNext;
    uint32_t                memoryBindIndex;
    VkMemoryRequirements    memoryRequirements;
} VkVideoSessionMemoryRequirementsKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memoryBindIndex` is the index of the memory binding.
- `memoryRequirements` is a [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure in which the requested memory binding requirements for the binding index specified by `memoryBindIndex` are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoSessionMemoryRequirementsKHR-sType-sType)VUID-VkVideoSessionMemoryRequirementsKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_SESSION_MEMORY_REQUIREMENTS_KHR`
- [](#VUID-VkVideoSessionMemoryRequirementsKHR-pNext-pNext)VUID-VkVideoSessionMemoryRequirementsKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetVideoSessionMemoryRequirementsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoSessionMemoryRequirementsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0