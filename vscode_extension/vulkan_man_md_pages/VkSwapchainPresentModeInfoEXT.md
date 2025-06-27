# VkSwapchainPresentModeInfoEXT(3) Manual Page

## Name

VkSwapchainPresentModeInfoEXT - Presentation modes for a
vkQueuePresentKHR operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSwapchainPresentModeInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_swapchain_maintenance1
typedef struct VkSwapchainPresentModeInfoEXT {
    VkStructureType            sType;
    const void*                pNext;
    uint32_t                   swapchainCount;
    const VkPresentModeKHR*    pPresentModes;
} VkSwapchainPresentModeInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `swapchainCount` is the number of swapchains being presented to by
  this command.

- `pPresentModes` is a list of presentation modes with `swapchainCount`
  entries.

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain of [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html)
includes a `VkSwapchainPresentModeInfoEXT` structure, then that
structure defines the presentation modes used for the current and
subsequent presentation operations.

When the application changes present modes with
[VkSwapchainPresentModeInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModeInfoEXT.html),
images that have already been queued for presentation will continue to
be presented according to the previous present mode. The current image
being queued for presentation and subsequent images will be presented
according to the new present mode. The behavior during the transition
between the two modes is defined as follows.

- Transition from `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR` to
  `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR`: the presentation engine
  updates the shared presentable image according to the behavior of
  `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR`.

- Transition from `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` to
  `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`: the presentation
  engine **may** update the shared presentable image or defer that to
  its regular refresh cycle, according to the behavior of
  `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`.

- Transition between `VK_PRESENT_MODE_FIFO_KHR` and
  `VK_PRESENT_MODE_FIFO_RELAXED_KHR`: Images continue to be appended to
  the same FIFO queue, and the behavior with respect to waiting for
  vertical blanking period will follow the new mode for current and
  subsequent images.

- Transition from `VK_PRESENT_MODE_IMMEDIATE_KHR` to
  `VK_PRESENT_MODE_FIFO_KHR` or `VK_PRESENT_MODE_FIFO_RELAXED_KHR`: As
  all prior present requests in the `VK_PRESENT_MODE_IMMEDIATE_KHR` mode
  are applied immediately, there are no outstanding present operations
  in this mode, and current and subsequent images are appended to the
  FIFO queue and presented according to the new mode.

- Transition from `VK_PRESENT_MODE_MAILBOX_KHR` to
  `VK_PRESENT_MODE_FIFO_KHR` or `VK_PRESENT_MODE_FIFO_RELAXED_KHR`:
  Presentation in both modes require waiting for the next vertical
  blanking period, with `VK_PRESENT_MODE_MAILBOX_KHR` allowing the
  pending present operation to be replaced by a new one. In this case,
  the current present operation will replace the pending present
  operation and is applied according to the new mode.

- Transition from `VK_PRESENT_MODE_FIFO_KHR` or
  `VK_PRESENT_MODE_FIFO_RELAXED_KHR` to `VK_PRESENT_MODE_IMMEDIATE_KHR`
  or `VK_PRESENT_MODE_MAILBOX_KHR`: If the FIFO queue is empty,
  presentation is done according to the behavior of the new mode. If
  there are present operations in the FIFO queue, once the last present
  operation is performed based on the respective vertical blanking
  period, the current and subsequent updates are applied according to
  the new mode.

- The behavior during transition between any other present modes, if
  possible, is implementation defined.

Valid Usage

- <a href="#VUID-VkSwapchainPresentModeInfoEXT-swapchainCount-07760"
  id="VUID-VkSwapchainPresentModeInfoEXT-swapchainCount-07760"></a>
  VUID-VkSwapchainPresentModeInfoEXT-swapchainCount-07760  
  `swapchainCount` **must** be equal to
  [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html)::`swapchainCount`

- <a href="#VUID-VkSwapchainPresentModeInfoEXT-pPresentModes-07761"
  id="VUID-VkSwapchainPresentModeInfoEXT-pPresentModes-07761"></a>
  VUID-VkSwapchainPresentModeInfoEXT-pPresentModes-07761  
  Each entry in `pPresentModes` must be a presentation mode specified in
  [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModesCreateInfoEXT.html)::`pPresentModes`
  when creating the entry’s corresponding swapchain

Valid Usage (Implicit)

- <a href="#VUID-VkSwapchainPresentModeInfoEXT-sType-sType"
  id="VUID-VkSwapchainPresentModeInfoEXT-sType-sType"></a>
  VUID-VkSwapchainPresentModeInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_MODE_INFO_EXT`

- <a href="#VUID-VkSwapchainPresentModeInfoEXT-pPresentModes-parameter"
  id="VUID-VkSwapchainPresentModeInfoEXT-pPresentModes-parameter"></a>
  VUID-VkSwapchainPresentModeInfoEXT-pPresentModes-parameter  
  `pPresentModes` **must** be a valid pointer to an array of
  `swapchainCount` valid [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html)
  values

- <a href="#VUID-VkSwapchainPresentModeInfoEXT-swapchainCount-arraylength"
  id="VUID-VkSwapchainPresentModeInfoEXT-swapchainCount-arraylength"></a>
  VUID-VkSwapchainPresentModeInfoEXT-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_swapchain_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_swapchain_maintenance1.html),
[VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSwapchainPresentModeInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
