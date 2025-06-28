# VkCommandBufferAllocateInfo(3) Manual Page

## Name

VkCommandBufferAllocateInfo - Structure specifying the allocation parameters for command buffer object



## [](#_c_specification)C Specification

The `VkCommandBufferAllocateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkCommandBufferAllocateInfo {
    VkStructureType         sType;
    const void*             pNext;
    VkCommandPool           commandPool;
    VkCommandBufferLevel    level;
    uint32_t                commandBufferCount;
} VkCommandBufferAllocateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `commandPool` is the command pool from which the command buffers are allocated.
- `level` is a [VkCommandBufferLevel](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferLevel.html) value specifying the command buffer level.
- `commandBufferCount` is the number of command buffers to allocate from the pool.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkCommandBufferAllocateInfo-sType-sType)VUID-VkCommandBufferAllocateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO`
- [](#VUID-VkCommandBufferAllocateInfo-pNext-pNext)VUID-VkCommandBufferAllocateInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCommandBufferAllocateInfo-commandPool-parameter)VUID-VkCommandBufferAllocateInfo-commandPool-parameter  
  `commandPool` **must** be a valid [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) handle
- [](#VUID-VkCommandBufferAllocateInfo-level-parameter)VUID-VkCommandBufferAllocateInfo-level-parameter  
  `level` **must** be a valid [VkCommandBufferLevel](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferLevel.html) value

Host Synchronization

- Host access to `commandPool` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBufferLevel](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferLevel.html), [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkAllocateCommandBuffers](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateCommandBuffers.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCommandBufferAllocateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0