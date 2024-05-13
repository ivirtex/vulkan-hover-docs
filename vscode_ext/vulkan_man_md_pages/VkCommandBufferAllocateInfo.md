# VkCommandBufferAllocateInfo(3) Manual Page

## Name

VkCommandBufferAllocateInfo - Structure specifying the allocation
parameters for command buffer object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCommandBufferAllocateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkCommandBufferAllocateInfo {
    VkStructureType         sType;
    const void*             pNext;
    VkCommandPool           commandPool;
    VkCommandBufferLevel    level;
    uint32_t                commandBufferCount;
} VkCommandBufferAllocateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `commandPool` is the command pool from which the command buffers are
  allocated.

- `level` is a [VkCommandBufferLevel](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferLevel.html) value
  specifying the command buffer level.

- `commandBufferCount` is the number of command buffers to allocate from
  the pool.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkCommandBufferAllocateInfo-sType-sType"
  id="VUID-VkCommandBufferAllocateInfo-sType-sType"></a>
  VUID-VkCommandBufferAllocateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO`

- <a href="#VUID-VkCommandBufferAllocateInfo-pNext-pNext"
  id="VUID-VkCommandBufferAllocateInfo-pNext-pNext"></a>
  VUID-VkCommandBufferAllocateInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCommandBufferAllocateInfo-commandPool-parameter"
  id="VUID-VkCommandBufferAllocateInfo-commandPool-parameter"></a>
  VUID-VkCommandBufferAllocateInfo-commandPool-parameter  
  `commandPool` **must** be a valid [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html)
  handle

- <a href="#VUID-VkCommandBufferAllocateInfo-level-parameter"
  id="VUID-VkCommandBufferAllocateInfo-level-parameter"></a>
  VUID-VkCommandBufferAllocateInfo-level-parameter  
  `level` **must** be a valid
  [VkCommandBufferLevel](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferLevel.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBufferLevel](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferLevel.html),
[VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkAllocateCommandBuffers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateCommandBuffers.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCommandBufferAllocateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
