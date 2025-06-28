# VkSwapchainPresentModeInfoEXT(3) Manual Page

## Name

VkSwapchainPresentModeInfoEXT - Presentation modes for a vkQueuePresentKHR operation



## [](#_c_specification)C Specification

The `VkSwapchainPresentModeInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_swapchain_maintenance1
typedef struct VkSwapchainPresentModeInfoEXT {
    VkStructureType            sType;
    const void*                pNext;
    uint32_t                   swapchainCount;
    const VkPresentModeKHR*    pPresentModes;
} VkSwapchainPresentModeInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `swapchainCount` is the number of swapchains being presented to by this command.
- `pPresentModes` is a list of presentation modes with `swapchainCount` entries.

## [](#_description)Description

If the `pNext` chain of [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html) includes a `VkSwapchainPresentModeInfoEXT` structure, then that structure defines the presentation modes used for the current and subsequent presentation operations.

When the application changes present modes with [VkSwapchainPresentModeInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModeInfoEXT.html), images that have already been queued for presentation will continue to be presented according to the previous present mode. The current image being queued for presentation and subsequent images will be presented according to the new present mode. The behavior during the transition between the two modes is defined as follows.

- Transition from `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR` to `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR`: the presentation engine updates the shared presentable image according to the behavior of `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR`.
- Transition from `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` to `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`: the presentation engine **may** update the shared presentable image or defer that to its regular refresh cycle, according to the behavior of `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`.
- Transition between `VK_PRESENT_MODE_FIFO_KHR` and `VK_PRESENT_MODE_FIFO_RELAXED_KHR`: Images continue to be appended to the same FIFO queue, and the behavior with respect to waiting for vertical blanking period will follow the new mode for current and subsequent images.
- Transition from `VK_PRESENT_MODE_IMMEDIATE_KHR` to `VK_PRESENT_MODE_FIFO_KHR` or `VK_PRESENT_MODE_FIFO_RELAXED_KHR` or `VK_PRESENT_MODE_FIFO_LATEST_READY_EXT` : As all prior present requests in the `VK_PRESENT_MODE_IMMEDIATE_KHR` mode are applied immediately, there are no outstanding present operations in this mode, and current and subsequent images are appended to the FIFO queue and presented according to the new mode.
- Transition from `VK_PRESENT_MODE_MAILBOX_KHR` to `VK_PRESENT_MODE_FIFO_KHR` or `VK_PRESENT_MODE_FIFO_RELAXED_KHR` or `VK_PRESENT_MODE_FIFO_LATEST_READY_EXT` : Presentation in FIFO modes require waiting for the next vertical blanking period, with `VK_PRESENT_MODE_MAILBOX_KHR` allowing the pending present operation to be replaced by a new one. In this case, the current present operation will replace the pending present operation and is applied according to the new mode.
- Transition from `VK_PRESENT_MODE_FIFO_KHR` or `VK_PRESENT_MODE_FIFO_RELAXED_KHR` or `VK_PRESENT_MODE_FIFO_LATEST_READY_EXT` to `VK_PRESENT_MODE_IMMEDIATE_KHR` or `VK_PRESENT_MODE_MAILBOX_KHR`: If the FIFO queue is empty, presentation is done according to the behavior of the new mode. If there are present operations in the FIFO queue, once the last present operation is performed based on the respective vertical blanking period, the current and subsequent updates are applied according to the new mode.
- Transition between `VK_PRESENT_MODE_FIFO_KHR` or `VK_PRESENT_MODE_FIFO_RELAXED_KHR`, and `VK_PRESENT_MODE_FIFO_LATEST_READY_EXT`: Images continue to be appended to the same FIFO queue, and the behavior with respect to waiting for vertical blanking period and dequeuing requests will follow the new mode for current and subsequent images.
- The behavior during transition between any other present modes, if possible, is implementation defined.

Valid Usage

- [](#VUID-VkSwapchainPresentModeInfoEXT-swapchainCount-07760)VUID-VkSwapchainPresentModeInfoEXT-swapchainCount-07760  
  `swapchainCount` **must** be equal to [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html)::`swapchainCount`
- [](#VUID-VkSwapchainPresentModeInfoEXT-pPresentModes-07761)VUID-VkSwapchainPresentModeInfoEXT-pPresentModes-07761  
  Each entry in `pPresentModes` **must** be a presentation mode specified in [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html)::`pPresentModes` when creating the entry’s corresponding swapchain

Valid Usage (Implicit)

- [](#VUID-VkSwapchainPresentModeInfoEXT-sType-sType)VUID-VkSwapchainPresentModeInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_MODE_INFO_EXT`
- [](#VUID-VkSwapchainPresentModeInfoEXT-pPresentModes-parameter)VUID-VkSwapchainPresentModeInfoEXT-pPresentModes-parameter  
  `pPresentModes` **must** be a valid pointer to an array of `swapchainCount` valid [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) values
- [](#VUID-VkSwapchainPresentModeInfoEXT-swapchainCount-arraylength)VUID-VkSwapchainPresentModeInfoEXT-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_swapchain_maintenance1.html), [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSwapchainPresentModeInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0