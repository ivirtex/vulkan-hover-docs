# VkBindVideoSessionMemoryInfoKHR(3) Manual Page

## Name

VkBindVideoSessionMemoryInfoKHR - Structure specifying memory bindings
for a video session object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBindVideoSessionMemoryInfoKHR` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memoryBindIndex` is the memory binding index to bind memory to.

- `memory` is the allocated device memory to be bound to the video
  session’s memory binding with index `memoryBindIndex`.

- `memoryOffset` is the start offset of the region of `memory` which is
  to be bound.

- `memorySize` is the size in bytes of the region of `memory`, starting
  from `memoryOffset` bytes, to be bound.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBindVideoSessionMemoryInfoKHR-memoryOffset-07201"
  id="VUID-VkBindVideoSessionMemoryInfoKHR-memoryOffset-07201"></a>
  VUID-VkBindVideoSessionMemoryInfoKHR-memoryOffset-07201  
  `memoryOffset` **must** be less than the size of `memory`

- <a href="#VUID-VkBindVideoSessionMemoryInfoKHR-memorySize-07202"
  id="VUID-VkBindVideoSessionMemoryInfoKHR-memorySize-07202"></a>
  VUID-VkBindVideoSessionMemoryInfoKHR-memorySize-07202  
  `memorySize` **must** be less than or equal to the size of `memory`
  minus `memoryOffset`

Valid Usage (Implicit)

- <a href="#VUID-VkBindVideoSessionMemoryInfoKHR-sType-sType"
  id="VUID-VkBindVideoSessionMemoryInfoKHR-sType-sType"></a>
  VUID-VkBindVideoSessionMemoryInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BIND_VIDEO_SESSION_MEMORY_INFO_KHR`

- <a href="#VUID-VkBindVideoSessionMemoryInfoKHR-pNext-pNext"
  id="VUID-VkBindVideoSessionMemoryInfoKHR-pNext-pNext"></a>
  VUID-VkBindVideoSessionMemoryInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkBindVideoSessionMemoryInfoKHR-memory-parameter"
  id="VUID-VkBindVideoSessionMemoryInfoKHR-memory-parameter"></a>
  VUID-VkBindVideoSessionMemoryInfoKHR-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkBindVideoSessionMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindVideoSessionMemoryKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindVideoSessionMemoryInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
