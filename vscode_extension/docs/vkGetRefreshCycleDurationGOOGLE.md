# vkGetRefreshCycleDurationGOOGLE(3) Manual Page

## Name

vkGetRefreshCycleDurationGOOGLE - Obtain the RC duration of the PE's display



## [](#_c_specification)C Specification

To query the duration of a refresh cycle (RC) for the presentation engine’s display, call:

```c++
// Provided by VK_GOOGLE_display_timing
VkResult vkGetRefreshCycleDurationGOOGLE(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    VkRefreshCycleDurationGOOGLE*               pDisplayTimingProperties);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapchain`.
- `swapchain` is the swapchain to obtain the refresh duration for.
- `pDisplayTimingProperties` is a pointer to a `VkRefreshCycleDurationGOOGLE` structure.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetRefreshCycleDurationGOOGLE-device-parameter)VUID-vkGetRefreshCycleDurationGOOGLE-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetRefreshCycleDurationGOOGLE-swapchain-parameter)VUID-vkGetRefreshCycleDurationGOOGLE-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkGetRefreshCycleDurationGOOGLE-pDisplayTimingProperties-parameter)VUID-vkGetRefreshCycleDurationGOOGLE-pDisplayTimingProperties-parameter  
  `pDisplayTimingProperties` **must** be a valid pointer to a [VkRefreshCycleDurationGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRefreshCycleDurationGOOGLE.html) structure
- [](#VUID-vkGetRefreshCycleDurationGOOGLE-swapchain-parent)VUID-vkGetRefreshCycleDurationGOOGLE-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_DEVICE_LOST`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_SURFACE_LOST_KHR`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_GOOGLE\_display\_timing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_GOOGLE_display_timing.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkRefreshCycleDurationGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRefreshCycleDurationGOOGLE.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetRefreshCycleDurationGOOGLE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0