# VkSemaphoreWaitInfo(3) Manual Page

## Name

VkSemaphoreWaitInfo - Structure containing information about the
semaphore wait condition



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSemaphoreWaitInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkSemaphoreWaitInfo {
    VkStructureType         sType;
    const void*             pNext;
    VkSemaphoreWaitFlags    flags;
    uint32_t                semaphoreCount;
    const VkSemaphore*      pSemaphores;
    const uint64_t*         pValues;
} VkSemaphoreWaitInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_timeline_semaphore
typedef VkSemaphoreWaitInfo VkSemaphoreWaitInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkSemaphoreWaitFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitFlagBits.html) specifying
  additional parameters for the semaphore wait operation.

- `semaphoreCount` is the number of semaphores to wait on.

- `pSemaphores` is a pointer to an array of `semaphoreCount` semaphore
  handles to wait on.

- `pValues` is a pointer to an array of `semaphoreCount` timeline
  semaphore values.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSemaphoreWaitInfo-pSemaphores-03256"
  id="VUID-VkSemaphoreWaitInfo-pSemaphores-03256"></a>
  VUID-VkSemaphoreWaitInfo-pSemaphores-03256  
  All of the elements of `pSemaphores` **must** reference a semaphore
  that was created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE`

Valid Usage (Implicit)

- <a href="#VUID-VkSemaphoreWaitInfo-sType-sType"
  id="VUID-VkSemaphoreWaitInfo-sType-sType"></a>
  VUID-VkSemaphoreWaitInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SEMAPHORE_WAIT_INFO`

- <a href="#VUID-VkSemaphoreWaitInfo-pNext-pNext"
  id="VUID-VkSemaphoreWaitInfo-pNext-pNext"></a>
  VUID-VkSemaphoreWaitInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkSemaphoreWaitInfo-flags-parameter"
  id="VUID-VkSemaphoreWaitInfo-flags-parameter"></a>
  VUID-VkSemaphoreWaitInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkSemaphoreWaitFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitFlagBits.html) values

- <a href="#VUID-VkSemaphoreWaitInfo-pSemaphores-parameter"
  id="VUID-VkSemaphoreWaitInfo-pSemaphores-parameter"></a>
  VUID-VkSemaphoreWaitInfo-pSemaphores-parameter  
  `pSemaphores` **must** be a valid pointer to an array of
  `semaphoreCount` valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handles

- <a href="#VUID-VkSemaphoreWaitInfo-pValues-parameter"
  id="VUID-VkSemaphoreWaitInfo-pValues-parameter"></a>
  VUID-VkSemaphoreWaitInfo-pValues-parameter  
  `pValues` **must** be a valid pointer to an array of `semaphoreCount`
  `uint64_t` values

- <a href="#VUID-VkSemaphoreWaitInfo-semaphoreCount-arraylength"
  id="VUID-VkSemaphoreWaitInfo-semaphoreCount-arraylength"></a>
  VUID-VkSemaphoreWaitInfo-semaphoreCount-arraylength  
  `semaphoreCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_timeline_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_timeline_semaphore.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkSemaphoreWaitFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkWaitSemaphores](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWaitSemaphores.html),
[vkWaitSemaphoresKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWaitSemaphoresKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSemaphoreWaitInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
