# vkGetRefreshCycleDurationGOOGLE(3) Manual Page

## Name

vkGetRefreshCycleDurationGOOGLE - Obtain the RC duration of the PE's
display



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the duration of a refresh cycle (RC) for the presentation
engine’s display, call:

``` c
// Provided by VK_GOOGLE_display_timing
VkResult vkGetRefreshCycleDurationGOOGLE(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    VkRefreshCycleDurationGOOGLE*               pDisplayTimingProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapchain`.

- `swapchain` is the swapchain to obtain the refresh duration for.

- `pDisplayTimingProperties` is a pointer to a
  `VkRefreshCycleDurationGOOGLE` structure.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetRefreshCycleDurationGOOGLE-device-parameter"
  id="VUID-vkGetRefreshCycleDurationGOOGLE-device-parameter"></a>
  VUID-vkGetRefreshCycleDurationGOOGLE-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetRefreshCycleDurationGOOGLE-swapchain-parameter"
  id="VUID-vkGetRefreshCycleDurationGOOGLE-swapchain-parameter"></a>
  VUID-vkGetRefreshCycleDurationGOOGLE-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a
  href="#VUID-vkGetRefreshCycleDurationGOOGLE-pDisplayTimingProperties-parameter"
  id="VUID-vkGetRefreshCycleDurationGOOGLE-pDisplayTimingProperties-parameter"></a>
  VUID-vkGetRefreshCycleDurationGOOGLE-pDisplayTimingProperties-parameter  
  `pDisplayTimingProperties` **must** be a valid pointer to a
  [VkRefreshCycleDurationGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRefreshCycleDurationGOOGLE.html)
  structure

- <a href="#VUID-vkGetRefreshCycleDurationGOOGLE-swapchain-parent"
  id="VUID-vkGetRefreshCycleDurationGOOGLE-swapchain-parent"></a>
  VUID-vkGetRefreshCycleDurationGOOGLE-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from
  `device`

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_DEVICE_LOST`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_GOOGLE_display_timing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_GOOGLE_display_timing.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkRefreshCycleDurationGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRefreshCycleDurationGOOGLE.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetRefreshCycleDurationGOOGLE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
