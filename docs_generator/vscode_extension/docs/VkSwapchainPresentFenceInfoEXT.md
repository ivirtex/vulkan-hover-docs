# VkSwapchainPresentFenceInfoEXT(3) Manual Page

## Name

VkSwapchainPresentFenceInfoEXT - Fences associated with a vkQueuePresentKHR operation



## [](#_c_specification)C Specification

The `VkSwapchainPresentFenceInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_swapchain_maintenance1
typedef struct VkSwapchainPresentFenceInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           swapchainCount;
    const VkFence*     pFences;
} VkSwapchainPresentFenceInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `swapchainCount` is the number of swapchains being presented to by this command.
- `pFences` is a list of fences with `swapchainCount` entries. Each entry **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or the handle of a fence to signal when the relevant operations on the associated swapchain have completed.

## [](#_description)Description

The set of *queue operations* defined by queuing an image for presentation, as well as operations performed by the presentation engine access the payloads of objects associated with the presentation operation. The associated objects include:

- The swapchain image, its implicitly bound memory, and any other resources bound to that memory.
- The wait semaphores specified when queuing the image for presentation.

The application **can** provide a fence that the implementation will signal when all such queue operations have completed and the presentation engine has taken a reference to the payload of any objects it accesses as part of the present operation. For all binary wait semaphores imported by the presentation engine using the equivalent of reference transference, as described in [Importing Semaphore Payloads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-importing), this fence **must** not signal until all such semaphore payloads have been reset by the presentation engine.

The application **can** destroy the wait semaphores associated with a given presentation operation when at least one of the associated fences is signaled, and **can** destroy the swapchain when the fences associated with all past presentation requests referring to that swapchain have signaled.

Fences associated with presentations to the same swapchain on the same [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) **must** be signaled in the same order as the present operations.

To specify a fence for each swapchain in a present operation, include the `VkSwapchainPresentFenceInfoEXT` structure in the `pNext` chain of the [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html) structure.

Valid Usage

- [](#VUID-VkSwapchainPresentFenceInfoEXT-swapchainCount-07757)VUID-VkSwapchainPresentFenceInfoEXT-swapchainCount-07757  
  `swapchainCount` **must** be equal to [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html)::`swapchainCount`
- [](#VUID-VkSwapchainPresentFenceInfoEXT-pFences-07758)VUID-VkSwapchainPresentFenceInfoEXT-pFences-07758  
  Each element of `pFences` that is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) **must** be unsignaled
- [](#VUID-VkSwapchainPresentFenceInfoEXT-pFences-07759)VUID-VkSwapchainPresentFenceInfoEXT-pFences-07759  
  Each element of `pFences` that is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) **must** not be associated with any other queue command that has not yet completed execution on that queue

Valid Usage (Implicit)

- [](#VUID-VkSwapchainPresentFenceInfoEXT-sType-sType)VUID-VkSwapchainPresentFenceInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_FENCE_INFO_EXT`
- [](#VUID-VkSwapchainPresentFenceInfoEXT-pFences-parameter)VUID-VkSwapchainPresentFenceInfoEXT-pFences-parameter  
  `pFences` **must** be a valid pointer to an array of `swapchainCount` valid or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handles
- [](#VUID-VkSwapchainPresentFenceInfoEXT-swapchainCount-arraylength)VUID-VkSwapchainPresentFenceInfoEXT-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_swapchain_maintenance1.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSwapchainPresentFenceInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0