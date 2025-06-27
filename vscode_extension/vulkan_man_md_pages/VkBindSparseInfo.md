# VkBindSparseInfo(3) Manual Page

## Name

VkBindSparseInfo - Structure specifying a sparse binding operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBindSparseInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkBindSparseInfo {
    VkStructureType                             sType;
    const void*                                 pNext;
    uint32_t                                    waitSemaphoreCount;
    const VkSemaphore*                          pWaitSemaphores;
    uint32_t                                    bufferBindCount;
    const VkSparseBufferMemoryBindInfo*         pBufferBinds;
    uint32_t                                    imageOpaqueBindCount;
    const VkSparseImageOpaqueMemoryBindInfo*    pImageOpaqueBinds;
    uint32_t                                    imageBindCount;
    const VkSparseImageMemoryBindInfo*          pImageBinds;
    uint32_t                                    signalSemaphoreCount;
    const VkSemaphore*                          pSignalSemaphores;
} VkBindSparseInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `waitSemaphoreCount` is the number of semaphores upon which to wait
  before executing the sparse binding operations for the batch.

- `pWaitSemaphores` is a pointer to an array of semaphores upon which to
  wait on before the sparse binding operations for this batch begin
  execution. If semaphores to wait on are provided, they define a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-waiting"
  target="_blank" rel="noopener">semaphore wait operation</a>.

- `bufferBindCount` is the number of sparse buffer bindings to perform
  in the batch.

- `pBufferBinds` is a pointer to an array of
  [VkSparseBufferMemoryBindInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseBufferMemoryBindInfo.html)
  structures.

- `imageOpaqueBindCount` is the number of opaque sparse image bindings
  to perform.

- `pImageOpaqueBinds` is a pointer to an array of
  [VkSparseImageOpaqueMemoryBindInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageOpaqueMemoryBindInfo.html)
  structures, indicating opaque sparse image bindings to perform.

- `imageBindCount` is the number of sparse image bindings to perform.

- `pImageBinds` is a pointer to an array of
  [VkSparseImageMemoryBindInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryBindInfo.html)
  structures, indicating sparse image bindings to perform.

- `signalSemaphoreCount` is the number of semaphores to be signaled once
  the sparse binding operations specified by the structure have
  completed execution.

- `pSignalSemaphores` is a pointer to an array of semaphores which will
  be signaled when the sparse binding operations for this batch have
  completed execution. If semaphores to be signaled are provided, they
  define a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
  target="_blank" rel="noopener">semaphore signal operation</a>.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBindSparseInfo-pWaitSemaphores-03246"
  id="VUID-VkBindSparseInfo-pWaitSemaphores-03246"></a>
  VUID-VkBindSparseInfo-pWaitSemaphores-03246  
  If any element of `pWaitSemaphores` or `pSignalSemaphores` was created
  with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE` then the `pNext` chain **must** include a
  [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html)
  structure

- <a href="#VUID-VkBindSparseInfo-pNext-03247"
  id="VUID-VkBindSparseInfo-pNext-03247"></a>
  VUID-VkBindSparseInfo-pNext-03247  
  If the `pNext` chain of this structure includes a
  [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html)
  structure and any element of `pWaitSemaphores` was created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE` then its `waitSemaphoreValueCount` member
  **must** equal `waitSemaphoreCount`

- <a href="#VUID-VkBindSparseInfo-pNext-03248"
  id="VUID-VkBindSparseInfo-pNext-03248"></a>
  VUID-VkBindSparseInfo-pNext-03248  
  If the `pNext` chain of this structure includes a
  [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html)
  structure and any element of `pSignalSemaphores` was created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE` then its `signalSemaphoreValueCount`
  member **must** equal `signalSemaphoreCount`

- <a href="#VUID-VkBindSparseInfo-pSignalSemaphores-03249"
  id="VUID-VkBindSparseInfo-pSignalSemaphores-03249"></a>
  VUID-VkBindSparseInfo-pSignalSemaphores-03249  
  For each element of `pSignalSemaphores` created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE` the corresponding element of
  [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html)::`pSignalSemaphoreValues`
  **must** have a value greater than the current value of the semaphore
  when the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
  target="_blank" rel="noopener">semaphore signal operation</a> is
  executed

- <a href="#VUID-VkBindSparseInfo-pWaitSemaphores-03250"
  id="VUID-VkBindSparseInfo-pWaitSemaphores-03250"></a>
  VUID-VkBindSparseInfo-pWaitSemaphores-03250  
  For each element of `pWaitSemaphores` created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE` the corresponding element of
  [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html)::`pWaitSemaphoreValues`
  **must** have a value which does not differ from the current value of
  the semaphore or from the value of any outstanding semaphore wait or
  signal operation on that semaphore by more than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxTimelineSemaphoreValueDifference"
  target="_blank"
  rel="noopener"><code>maxTimelineSemaphoreValueDifference</code></a>

- <a href="#VUID-VkBindSparseInfo-pSignalSemaphores-03251"
  id="VUID-VkBindSparseInfo-pSignalSemaphores-03251"></a>
  VUID-VkBindSparseInfo-pSignalSemaphores-03251  
  For each element of `pSignalSemaphores` created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE` the corresponding element of
  [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html)::`pSignalSemaphoreValues`
  **must** have a value which does not differ from the current value of
  the semaphore or from the value of any outstanding semaphore wait or
  signal operation on that semaphore by more than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxTimelineSemaphoreValueDifference"
  target="_blank"
  rel="noopener"><code>maxTimelineSemaphoreValueDifference</code></a>

Valid Usage (Implicit)

- <a href="#VUID-VkBindSparseInfo-sType-sType"
  id="VUID-VkBindSparseInfo-sType-sType"></a>
  VUID-VkBindSparseInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_SPARSE_INFO`

- <a href="#VUID-VkBindSparseInfo-pNext-pNext"
  id="VUID-VkBindSparseInfo-pNext-pNext"></a>
  VUID-VkBindSparseInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDeviceGroupBindSparseInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupBindSparseInfo.html),
  [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrameBoundaryEXT.html), or
  [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html)

- <a href="#VUID-VkBindSparseInfo-sType-unique"
  id="VUID-VkBindSparseInfo-sType-unique"></a>
  VUID-VkBindSparseInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkBindSparseInfo-pWaitSemaphores-parameter"
  id="VUID-VkBindSparseInfo-pWaitSemaphores-parameter"></a>
  VUID-VkBindSparseInfo-pWaitSemaphores-parameter  
  If `waitSemaphoreCount` is not `0`, `pWaitSemaphores` **must** be a
  valid pointer to an array of `waitSemaphoreCount` valid
  [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handles

- <a href="#VUID-VkBindSparseInfo-pBufferBinds-parameter"
  id="VUID-VkBindSparseInfo-pBufferBinds-parameter"></a>
  VUID-VkBindSparseInfo-pBufferBinds-parameter  
  If `bufferBindCount` is not `0`, `pBufferBinds` **must** be a valid
  pointer to an array of `bufferBindCount` valid
  [VkSparseBufferMemoryBindInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseBufferMemoryBindInfo.html)
  structures

- <a href="#VUID-VkBindSparseInfo-pImageOpaqueBinds-parameter"
  id="VUID-VkBindSparseInfo-pImageOpaqueBinds-parameter"></a>
  VUID-VkBindSparseInfo-pImageOpaqueBinds-parameter  
  If `imageOpaqueBindCount` is not `0`, `pImageOpaqueBinds` **must** be
  a valid pointer to an array of `imageOpaqueBindCount` valid
  [VkSparseImageOpaqueMemoryBindInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageOpaqueMemoryBindInfo.html)
  structures

- <a href="#VUID-VkBindSparseInfo-pImageBinds-parameter"
  id="VUID-VkBindSparseInfo-pImageBinds-parameter"></a>
  VUID-VkBindSparseInfo-pImageBinds-parameter  
  If `imageBindCount` is not `0`, `pImageBinds` **must** be a valid
  pointer to an array of `imageBindCount` valid
  [VkSparseImageMemoryBindInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryBindInfo.html)
  structures

- <a href="#VUID-VkBindSparseInfo-pSignalSemaphores-parameter"
  id="VUID-VkBindSparseInfo-pSignalSemaphores-parameter"></a>
  VUID-VkBindSparseInfo-pSignalSemaphores-parameter  
  If `signalSemaphoreCount` is not `0`, `pSignalSemaphores` **must** be
  a valid pointer to an array of `signalSemaphoreCount` valid
  [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handles

- <a href="#VUID-VkBindSparseInfo-commonparent"
  id="VUID-VkBindSparseInfo-commonparent"></a>
  VUID-VkBindSparseInfo-commonparent  
  Both of the elements of `pSignalSemaphores`, and the elements of
  `pWaitSemaphores` that are valid handles of non-ignored parameters
  **must** have been created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkSparseBufferMemoryBindInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseBufferMemoryBindInfo.html),
[VkSparseImageMemoryBindInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryBindInfo.html),
[VkSparseImageOpaqueMemoryBindInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageOpaqueMemoryBindInfo.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkQueueBindSparse](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueBindSparse.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindSparseInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
