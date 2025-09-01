# VkPresentInfoKHR(3) Manual Page

## Name

VkPresentInfoKHR - Structure describing parameters of a queue presentation



## [](#_c_specification)C Specification

The `VkPresentInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_swapchain
typedef struct VkPresentInfoKHR {
    VkStructureType          sType;
    const void*              pNext;
    uint32_t                 waitSemaphoreCount;
    const VkSemaphore*       pWaitSemaphores;
    uint32_t                 swapchainCount;
    const VkSwapchainKHR*    pSwapchains;
    const uint32_t*          pImageIndices;
    VkResult*                pResults;
} VkPresentInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `waitSemaphoreCount` is the number of semaphores to wait for before issuing the present request. The number **may** be zero.
- `pWaitSemaphores` is `NULL` or a pointer to an array of [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) objects with `waitSemaphoreCount` entries, and specifies the semaphores to wait for before issuing the present request.
- `swapchainCount` is the number of swapchains being presented to by this command.
- `pSwapchains` is a pointer to an array of [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) objects with `swapchainCount` entries.
- `pImageIndices` is a pointer to an array of indices into the array of each swapchain’s presentable images, with `swapchainCount` entries. Each entry in this array identifies the image to present on the corresponding entry in the `pSwapchains` array.
- `pResults` is a pointer to an array of [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html) typed elements with `swapchainCount` entries. Applications that do not need per-swapchain results **can** use `NULL` for `pResults`. If non-`NULL`, each entry in `pResults` will be set to the [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html) for presenting the swapchain corresponding to the same index in `pSwapchains`.

## [](#_description)Description

Before an application **can** present an image, the image’s layout **must** be transitioned to the `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR` layout, or for a shared presentable image the `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR` layout.

Note

When transitioning the image to the appropriate layout, there is no need to delay subsequent processing, or perform any visibility operations (as [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) performs automatic visibility operations). To achieve this, the `dstAccessMask` member of the [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier.html) **should** be `0`, and the `dstStageMask` parameter **should** be `VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT`.

Valid Usage

- [](#VUID-VkPresentInfoKHR-pSwapchain-09231)VUID-VkPresentInfoKHR-pSwapchain-09231  
  Elements of `pSwapchain` **must** be unique
- [](#VUID-VkPresentInfoKHR-pImageIndices-01430)VUID-VkPresentInfoKHR-pImageIndices-01430  
  Each element of `pImageIndices` **must** be the index of a presentable image acquired from the swapchain specified by the corresponding element of the `pSwapchains` array, and the presented image subresource **must** be in the `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR` or `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR` layout at the time the operation is executed on a `VkDevice`
- [](#VUID-VkPresentInfoKHR-pNext-06235)VUID-VkPresentInfoKHR-pNext-06235  
  If a [VkPresentIdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentIdKHR.html) structure is included in the `pNext` chain, and the [`presentId`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-presentId) feature is not enabled, each `presentIds` entry in that structure **must** be NULL
- [](#VUID-VkPresentInfoKHR-swapchainMaintenance1-10158)VUID-VkPresentInfoKHR-swapchainMaintenance1-10158  
  If the [`swapchainMaintenance1`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-swapchainMaintenance1) feature is not enabled, then the `pNext` chain **must** not include a [VkSwapchainPresentFenceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentFenceInfoKHR.html) structure
- [](#VUID-VkPresentInfoKHR-pSwapchains-09199)VUID-VkPresentInfoKHR-pSwapchains-09199  
  If any element of the `pSwapchains` array has been created with [VkSwapchainPresentModesCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoKHR.html), all of the elements of this array **must** be created with [VkSwapchainPresentModesCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoKHR.html)
- [](#VUID-VkPresentInfoKHR-pNext-09759)VUID-VkPresentInfoKHR-pNext-09759  
  If the `pNext` chain of this structure includes a [VkFrameBoundaryTensorsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryTensorsARM.html) structure then it **must** also include a [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryEXT.html) structure.
- [](#VUID-VkPresentInfoKHR-pNext-10821)VUID-VkPresentInfoKHR-pNext-10821  
  If a [VkPresentId2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentId2KHR.html) structure is included in the `pNext` chain, and the [`presentId2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-presentId2) feature is not enabled, each `presentIds` entry in that structure **must** be zero
- [](#VUID-VkPresentInfoKHR-presentId2Supported-10822)VUID-VkPresentInfoKHR-presentId2Supported-10822  
  If a [VkPresentId2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentId2KHR.html) structure is included and contains non-zero presentIds, `presentId2Supported` **must** be `VK_TRUE` in the [VkSurfaceCapabilitiesPresentId2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesPresentId2KHR.html) structure returned by [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html) for the `surface`

Valid Usage (Implicit)

- [](#VUID-VkPresentInfoKHR-sType-sType)VUID-VkPresentInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PRESENT_INFO_KHR`
- [](#VUID-VkPresentInfoKHR-pNext-pNext)VUID-VkPresentInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkDeviceGroupPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentInfoKHR.html), [VkDisplayPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPresentInfoKHR.html), [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryEXT.html), [VkFrameBoundaryTensorsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryTensorsARM.html), [VkPresentFrameTokenGGP](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentFrameTokenGGP.html), [VkPresentId2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentId2KHR.html), [VkPresentIdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentIdKHR.html), [VkPresentRegionsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentRegionsKHR.html), [VkPresentTimesInfoGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentTimesInfoGOOGLE.html), [VkSetPresentConfigNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSetPresentConfigNV.html), [VkSwapchainPresentFenceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentFenceInfoKHR.html), or [VkSwapchainPresentModeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModeInfoKHR.html)
- [](#VUID-VkPresentInfoKHR-sType-unique)VUID-VkPresentInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPresentInfoKHR-pWaitSemaphores-parameter)VUID-VkPresentInfoKHR-pWaitSemaphores-parameter  
  If `waitSemaphoreCount` is not `0`, `pWaitSemaphores` **must** be a valid pointer to an array of `waitSemaphoreCount` valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handles
- [](#VUID-VkPresentInfoKHR-pSwapchains-parameter)VUID-VkPresentInfoKHR-pSwapchains-parameter  
  `pSwapchains` **must** be a valid pointer to an array of `swapchainCount` valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handles
- [](#VUID-VkPresentInfoKHR-pImageIndices-parameter)VUID-VkPresentInfoKHR-pImageIndices-parameter  
  `pImageIndices` **must** be a valid pointer to an array of `swapchainCount` `uint32_t` values
- [](#VUID-VkPresentInfoKHR-pResults-parameter)VUID-VkPresentInfoKHR-pResults-parameter  
  If `pResults` is not `NULL`, `pResults` **must** be a valid pointer to an array of `swapchainCount` [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html) values
- [](#VUID-VkPresentInfoKHR-swapchainCount-arraylength)VUID-VkPresentInfoKHR-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`
- [](#VUID-VkPresentInfoKHR-commonparent)VUID-VkPresentInfoKHR-commonparent  
  Both of the elements of `pSwapchains`, and the elements of `pWaitSemaphores` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to each member of `pWaitSemaphores` **must** be externally synchronized
- Host access to each member of `pSwapchains` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html), [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html), [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPresentInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0