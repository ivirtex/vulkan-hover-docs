# vkWaitForPresentKHR(3) Manual Page

## Name

vkWaitForPresentKHR - Wait for presentation



## [](#_c_specification)C Specification

When the [`presentWait`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-presentWait) feature is enabled, an application **can** wait for an image to be presented to the user by first specifying a presentId for the target presentation by adding a `VkPresentIdKHR` structure to the `pNext` chain of the [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html) structure and then waiting for that presentation to complete by calling:

```c++
// Provided by VK_KHR_present_wait
VkResult vkWaitForPresentKHR(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    uint64_t                                    presentId,
    uint64_t                                    timeout);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapchain`.
- `swapchain` is the non-retired swapchain on which an image was queued for presentation.
- `presentId` is the presentation presentId to wait for.
- `timeout` is the timeout period in units of nanoseconds. `timeout` is adjusted to the closest value allowed by the implementation-dependent timeout accuracy, which **may** be substantially longer than one nanosecond, and **may** be longer than the requested period.

## [](#_description)Description

`vkWaitForPresentKHR` waits for the presentId associated with `swapchain` to be increased in value so that it is at least equal to `presentId`.

For `VK_PRESENT_MODE_MAILBOX_KHR` (or other present mode where images may be replaced in the presentation queue) any wait of this type associated with such an image **must** be signaled no later than a wait associated with the replacing image would be signaled.

When the presentation has completed, the presentId associated with the related `pSwapchains` entry will be increased in value so that it is at least equal to the value provided in the `VkPresentIdKHR` structure.

There is no requirement for any precise timing relationship between the presentation of the image to the user and the update of the presentId value, but implementations **should** make this as close as possible to the presentation of the first pixel in the next image being presented to the user.

The call to `vkWaitForPresentKHR` will block until either the presentId associated with `swapchain` is greater than or equal to `presentId`, or `timeout` nanoseconds passes. When the swapchain becomes OUT\_OF\_DATE, the call will either return `VK_SUCCESS` (if the image was delivered to the presentation engine and may have been presented to the user) or will return early with status `VK_ERROR_OUT_OF_DATE_KHR` (if the image could not be presented to the user).

As an exception to the normal rules for objects which are externally synchronized, the `swapchain` passed to `vkWaitForPresentKHR` **may** be simultaneously used by other threads in calls to functions other than [vkDestroySwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroySwapchainKHR.html). Access to the swapchain data associated with this extension **must** be atomic within the implementation.

Valid Usage

- [](#VUID-vkWaitForPresentKHR-swapchain-04997)VUID-vkWaitForPresentKHR-swapchain-04997  
  `swapchain` **must** not be in the retired state
- [](#VUID-vkWaitForPresentKHR-presentWait-06234)VUID-vkWaitForPresentKHR-presentWait-06234  
  The [`presentWait`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-presentWait) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkWaitForPresentKHR-device-parameter)VUID-vkWaitForPresentKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkWaitForPresentKHR-swapchain-parameter)VUID-vkWaitForPresentKHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkWaitForPresentKHR-swapchain-parent)VUID-vkWaitForPresentKHR-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUBOPTIMAL_KHR`
- `VK_SUCCESS`
- `VK_TIMEOUT`

On failure, this command returns

- `VK_ERROR_DEVICE_LOST`
- `VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT`
- `VK_ERROR_OUT_OF_DATE_KHR`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_SURFACE_LOST_KHR`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_present\_wait](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_present_wait.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkWaitForPresentKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0