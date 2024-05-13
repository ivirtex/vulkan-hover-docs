# VkDependencyInfo(3) Manual Page

## Name

VkDependencyInfo - Structure specifying dependency information for a
synchronization command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDependencyInfo` structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_synchronization2
typedef VkDependencyInfo VkDependencyInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `dependencyFlags` is a bitmask of
  [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html) specifying how
  execution and memory dependencies are formed.

- `memoryBarrierCount` is the length of the `pMemoryBarriers` array.

- `pMemoryBarriers` is a pointer to an array of
  [VkMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier2.html) structures defining memory
  dependencies between any memory accesses.

- `bufferMemoryBarrierCount` is the length of the
  `pBufferMemoryBarriers` array.

- `pBufferMemoryBarriers` is a pointer to an array of
  [VkBufferMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier2.html) structures
  defining memory dependencies between buffer ranges.

- `imageMemoryBarrierCount` is the length of the `pImageMemoryBarriers`
  array.

- `pImageMemoryBarriers` is a pointer to an array of
  [VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier2.html) structures
  defining memory dependencies between image subresources.

## <a href="#_description" class="anchor"></a>Description

This structure defines a set of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-memory"
target="_blank" rel="noopener">memory dependencies</a>, as well as <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
target="_blank" rel="noopener">queue family ownership transfer
operations</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
target="_blank" rel="noopener">image layout transitions</a>.

Each member of `pMemoryBarriers`, `pBufferMemoryBarriers`, and
`pImageMemoryBarriers` defines a separate <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-memory"
target="_blank" rel="noopener">memory dependency</a>.

Valid Usage (Implicit)

- <a href="#VUID-VkDependencyInfo-sType-sType"
  id="VUID-VkDependencyInfo-sType-sType"></a>
  VUID-VkDependencyInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEPENDENCY_INFO`

- <a href="#VUID-VkDependencyInfo-pNext-pNext"
  id="VUID-VkDependencyInfo-pNext-pNext"></a>
  VUID-VkDependencyInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDependencyInfo-dependencyFlags-parameter"
  id="VUID-VkDependencyInfo-dependencyFlags-parameter"></a>
  VUID-VkDependencyInfo-dependencyFlags-parameter  
  `dependencyFlags` **must** be a valid combination of
  [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html) values

- <a href="#VUID-VkDependencyInfo-pMemoryBarriers-parameter"
  id="VUID-VkDependencyInfo-pMemoryBarriers-parameter"></a>
  VUID-VkDependencyInfo-pMemoryBarriers-parameter  
  If `memoryBarrierCount` is not `0`, `pMemoryBarriers` **must** be a
  valid pointer to an array of `memoryBarrierCount` valid
  [VkMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier2.html) structures

- <a href="#VUID-VkDependencyInfo-pBufferMemoryBarriers-parameter"
  id="VUID-VkDependencyInfo-pBufferMemoryBarriers-parameter"></a>
  VUID-VkDependencyInfo-pBufferMemoryBarriers-parameter  
  If `bufferMemoryBarrierCount` is not `0`, `pBufferMemoryBarriers`
  **must** be a valid pointer to an array of `bufferMemoryBarrierCount`
  valid [VkBufferMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier2.html) structures

- <a href="#VUID-VkDependencyInfo-pImageMemoryBarriers-parameter"
  id="VUID-VkDependencyInfo-pImageMemoryBarriers-parameter"></a>
  VUID-VkDependencyInfo-pImageMemoryBarriers-parameter  
  If `imageMemoryBarrierCount` is not `0`, `pImageMemoryBarriers`
  **must** be a valid pointer to an array of `imageMemoryBarrierCount`
  valid [VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier2.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkBufferMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier2.html),
[VkDependencyFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlags.html),
[VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier2.html),
[VkMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdPipelineBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier2.html),
[vkCmdPipelineBarrier2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier2KHR.html),
[vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2.html),
[vkCmdSetEvent2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2KHR.html),
[vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2.html),
[vkCmdWaitEvents2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDependencyInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
