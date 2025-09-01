# VkDependencyInfo(3) Manual Page

## Name

VkDependencyInfo - Structure specifying dependency information for a synchronization command



## [](#_c_specification)C Specification

The `VkDependencyInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkDependencyInfo {
    VkStructureType                  sType;
    const void*                      pNext;
    VkDependencyFlags                dependencyFlags;
    uint32_t                         memoryBarrierCount;
    const VkMemoryBarrier2*          pMemoryBarriers;
    uint32_t                         bufferMemoryBarrierCount;
    const VkBufferMemoryBarrier2*    pBufferMemoryBarriers;
    uint32_t                         imageMemoryBarrierCount;
    const VkImageMemoryBarrier2*     pImageMemoryBarriers;
} VkDependencyInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_synchronization2
typedef VkDependencyInfo VkDependencyInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `dependencyFlags` is a bitmask of [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlagBits.html) specifying how execution and memory dependencies are formed.
- `memoryBarrierCount` is the length of the `pMemoryBarriers` array.
- `pMemoryBarriers` is a pointer to an array of [VkMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier2.html) structures defining memory dependencies between any memory accesses.
- `bufferMemoryBarrierCount` is the length of the `pBufferMemoryBarriers` array.
- `pBufferMemoryBarriers` is a pointer to an array of [VkBufferMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier2.html) structures defining memory dependencies between buffer ranges.
- `imageMemoryBarrierCount` is the length of the `pImageMemoryBarriers` array.
- `pImageMemoryBarriers` is a pointer to an array of [VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier2.html) structures defining memory dependencies between image subresources.

## [](#_description)Description

This structure defines a set of [memory dependencies](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-memory), as well as [queue family ownership transfer operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers) and [image layout transitions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions).

Each member of `pMemoryBarriers`, `pBufferMemoryBarriers`, and `pImageMemoryBarriers` defines a separate [memory dependency](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-memory).

Valid Usage

- [](#VUID-VkDependencyInfo-pMemoryBarriers-10605)VUID-VkDependencyInfo-pMemoryBarriers-10605  
  For each element of `pMemoryBarriers`, the `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkDependencyInfo-pMemoryBarriers-10606)VUID-VkDependencyInfo-pMemoryBarriers-10606  
  For each element of `pMemoryBarriers`, `pNext` **must** be either `NULL` or a pointer to a valid instance of [VkMemoryBarrierAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrierAccessFlags3KHR.html)
- [](#VUID-VkDependencyInfo-pNext-09754)VUID-VkDependencyInfo-pNext-09754  
  If a [VkTensorDependencyInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDependencyInfoARM.html) structure is included in the `pNext` chain, a [VkTensorMemoryBarrierARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorMemoryBarrierARM.html) structure **must** not be included in the `pNext` chain

Valid Usage (Implicit)

- [](#VUID-VkDependencyInfo-sType-sType)VUID-VkDependencyInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEPENDENCY_INFO`
- [](#VUID-VkDependencyInfo-pNext-pNext)VUID-VkDependencyInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkTensorDependencyInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDependencyInfoARM.html) or [VkTensorMemoryBarrierARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorMemoryBarrierARM.html)
- [](#VUID-VkDependencyInfo-sType-unique)VUID-VkDependencyInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkDependencyInfo-dependencyFlags-parameter)VUID-VkDependencyInfo-dependencyFlags-parameter  
  `dependencyFlags` **must** be a valid combination of [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlagBits.html) values
- [](#VUID-VkDependencyInfo-pMemoryBarriers-parameter)VUID-VkDependencyInfo-pMemoryBarriers-parameter  
  If `memoryBarrierCount` is not `0`, `pMemoryBarriers` **must** be a valid pointer to an array of `memoryBarrierCount` valid [VkMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier2.html) structures
- [](#VUID-VkDependencyInfo-pBufferMemoryBarriers-parameter)VUID-VkDependencyInfo-pBufferMemoryBarriers-parameter  
  If `bufferMemoryBarrierCount` is not `0`, `pBufferMemoryBarriers` **must** be a valid pointer to an array of `bufferMemoryBarrierCount` valid [VkBufferMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier2.html) structures
- [](#VUID-VkDependencyInfo-pImageMemoryBarriers-parameter)VUID-VkDependencyInfo-pImageMemoryBarriers-parameter  
  If `imageMemoryBarrierCount` is not `0`, `pImageMemoryBarriers` **must** be a valid pointer to an array of `imageMemoryBarrierCount` valid [VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier2.html) structures

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkBufferMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier2.html), [VkDependencyFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlags.html), [VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier2.html), [VkMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdPipelineBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier2.html), [vkCmdPipelineBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier2.html), [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2.html), [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2.html), [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents2.html), [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDependencyInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0