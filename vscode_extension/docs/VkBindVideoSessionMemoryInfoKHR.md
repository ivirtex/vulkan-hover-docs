# VkBindVideoSessionMemoryInfoKHR(3) Manual Page

## Name

VkBindVideoSessionMemoryInfoKHR - Structure specifying memory bindings for a video session object



## [](#_c_specification)C Specification

The `VkBindVideoSessionMemoryInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkBindVideoSessionMemoryInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           memoryBindIndex;
    VkDeviceMemory     memory;
    VkDeviceSize       memoryOffset;
    VkDeviceSize       memorySize;
} VkBindVideoSessionMemoryInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memoryBindIndex` is the memory binding index to bind memory to.
- `memory` is the allocated device memory to be bound to the video session’s memory binding with index `memoryBindIndex`.
- `memoryOffset` is the start offset of the region of `memory` which is to be bound.
- `memorySize` is the size in bytes of the region of `memory`, starting from `memoryOffset` bytes, to be bound.

## [](#_description)Description

Valid Usage

- [](#VUID-VkBindVideoSessionMemoryInfoKHR-memoryOffset-07201)VUID-VkBindVideoSessionMemoryInfoKHR-memoryOffset-07201  
  `memoryOffset` **must** be less than the size of `memory`
- [](#VUID-VkBindVideoSessionMemoryInfoKHR-memorySize-07202)VUID-VkBindVideoSessionMemoryInfoKHR-memorySize-07202  
  `memorySize` **must** be less than or equal to the size of `memory` minus `memoryOffset`

Valid Usage (Implicit)

- [](#VUID-VkBindVideoSessionMemoryInfoKHR-sType-sType)VUID-VkBindVideoSessionMemoryInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_VIDEO_SESSION_MEMORY_INFO_KHR`
- [](#VUID-VkBindVideoSessionMemoryInfoKHR-pNext-pNext)VUID-VkBindVideoSessionMemoryInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkBindVideoSessionMemoryInfoKHR-memory-parameter)VUID-VkBindVideoSessionMemoryInfoKHR-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkBindVideoSessionMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindVideoSessionMemoryKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBindVideoSessionMemoryInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0