# VkSwapchainPresentFenceInfoEXT(3) Manual Page

## Name

VkSwapchainPresentFenceInfoEXT - Fences associated with a
vkQueuePresentKHR operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSwapchainPresentFenceInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_swapchain_maintenance1
typedef struct VkSwapchainPresentFenceInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           swapchainCount;
    const VkFence*     pFences;
} VkSwapchainPresentFenceInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `swapchainCount` is the number of swapchains being presented to by
  this command.

- `pFences` is a list of fences with `swapchainCount` entries. Each
  entry **must** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or the handle
  of a fence to signal when the relevant operations on the associated
  swapchain have completed.

## <a href="#_description" class="anchor"></a>Description

The set of *queue operations* defined by queuing an image for
presentation, as well as operations performed by the presentation engine
access the payloads of objects associated with the presentation
operation. The associated objects include:

- The swapchain image, its implicitly bound memory, and any other
  resources bound to that memory.

- The wait semaphores specified when queuing the image for presentation.

The application **can** provide a fence that the implementation will
signal when all such queue operations have completed and the
presentation engine has taken a reference to the payload of any objects
it accesses as part of the present operation. For all binary wait
semaphores imported by the presentation engine using the equivalent of
reference transference, as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-importing"
target="_blank" rel="noopener">Importing Semaphore Payloads</a>, this
fence **must** not signal until all such semaphore payloads have been
reset by the presentation engine.

The application **can** destroy the wait semaphores associated with a
given presentation operation when at least one of the associated fences
is signaled, and **can** destroy the swapchain when the fences
associated with all past presentation requests referring to that
swapchain have signaled.

Fences associated with presentations to the same swapchain on the same
[VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) **must** be signaled in the same order as the
present operations.

To specify a fence for each swapchain in a present operation, include
the `VkSwapchainPresentFenceInfoEXT` structure in the `pNext` chain of
the [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html) structure.

Valid Usage

- <a href="#VUID-VkSwapchainPresentFenceInfoEXT-swapchainCount-07757"
  id="VUID-VkSwapchainPresentFenceInfoEXT-swapchainCount-07757"></a>
  VUID-VkSwapchainPresentFenceInfoEXT-swapchainCount-07757  
  `swapchainCount` **must** be equal to
  [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html)::`swapchainCount`

- <a href="#VUID-VkSwapchainPresentFenceInfoEXT-pFences-07758"
  id="VUID-VkSwapchainPresentFenceInfoEXT-pFences-07758"></a>
  VUID-VkSwapchainPresentFenceInfoEXT-pFences-07758  
  Each element of `pFences` **must** be unsignaled

- <a href="#VUID-VkSwapchainPresentFenceInfoEXT-pFences-07759"
  id="VUID-VkSwapchainPresentFenceInfoEXT-pFences-07759"></a>
  VUID-VkSwapchainPresentFenceInfoEXT-pFences-07759  
  Each element of `pFences` **must** not be associated with any other
  queue command that has not yet completed execution on that queue

Valid Usage (Implicit)

- <a href="#VUID-VkSwapchainPresentFenceInfoEXT-sType-sType"
  id="VUID-VkSwapchainPresentFenceInfoEXT-sType-sType"></a>
  VUID-VkSwapchainPresentFenceInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_FENCE_INFO_EXT`

- <a href="#VUID-VkSwapchainPresentFenceInfoEXT-pFences-parameter"
  id="VUID-VkSwapchainPresentFenceInfoEXT-pFences-parameter"></a>
  VUID-VkSwapchainPresentFenceInfoEXT-pFences-parameter  
  `pFences` **must** be a valid pointer to an array of `swapchainCount`
  valid [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) handles

- <a
  href="#VUID-VkSwapchainPresentFenceInfoEXT-swapchainCount-arraylength"
  id="VUID-VkSwapchainPresentFenceInfoEXT-swapchainCount-arraylength"></a>
  VUID-VkSwapchainPresentFenceInfoEXT-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_swapchain_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_swapchain_maintenance1.html),
[VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSwapchainPresentFenceInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
