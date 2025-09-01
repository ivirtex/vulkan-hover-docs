# vkWaitForPresent2KHR(3) Manual Page

## Name

vkWaitForPresent2KHR - Wait for presentation



## [](#_c_specification)C Specification

When the `VkSurfaceCapabilitiesPresentWait2KHR` surface capability is present for a given surface, an application **can** wait for an image to be presented to the user by first specifying a `presentId` for the target presentation by adding a `VkPresentId2KHR` structure to the `pNext` chain of the [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html) structure and then waiting for that presentation to complete by calling:

```c++
// Provided by VK_KHR_present_wait2
VkResult vkWaitForPresent2KHR(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    const VkPresentWait2InfoKHR*                pPresentWait2Info);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapchain`.
- `swapchain` is the non-retired swapchain on which an image was queued for presentation.
- `pPresentWait2Info` is a pointer to a [VkPresentWait2InfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentWait2InfoKHR.html) structure specifying the parameters of the wait.

## [](#_description)Description

`vkWaitForPresent2KHR` waits for the presentation engine to have begun presentation of the presentation request associated with the [VkPresentWait2InfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentWait2InfoKHR.html)::`presentId` on `swapchain`, or for [VkPresentWait2InfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentWait2InfoKHR.html)::`timeout` to have expired.

The wait request will complete when the timeout expires, or after the corresponding presentation request has either taken effect within the presentation engine or has been replaced without presentation.

The timing relationship between the presentation of the image to the user and the wait request completing is implementation-dependent due to variations in window system implementations.

If the `swapchain` becomes `VK_ERROR_OUT_OF_DATE_KHR` either before or during this call, the call **may** either return `VK_SUCCESS` (if the image was delivered to the presentation engine and **may** have been presented to the user) or return early with status `VK_ERROR_OUT_OF_DATE_KHR` (if the image could not be presented to the user).

As an exception to the normal rules for objects which are externally synchronized, the `swapchain` passed to `vkWaitForPresent2KHR` **may** be simultaneously used by other threads in calls to functions other than [vkDestroySwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroySwapchainKHR.html). Access to the swapchain data associated with this extension **must** be atomic within the implementation.

Valid Usage

- [](#VUID-vkWaitForPresent2KHR-presentWait2-10814)VUID-vkWaitForPresent2KHR-presentWait2-10814  
  The [`presentWait2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-presentWait2) feature **must** be enabled
- [](#VUID-vkWaitForPresent2KHR-None-10815)VUID-vkWaitForPresent2KHR-None-10815  
  The [VkSurfaceCapabilitiesPresentWait2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesPresentWait2KHR.html) surface capability **must** be present for the underlying surface
- [](#VUID-vkWaitForPresent2KHR-None-10816)VUID-vkWaitForPresent2KHR-None-10816  
  The swapchain must have been created with `VK_SWAPCHAIN_CREATE_PRESENT_WAIT_2_BIT_KHR` bit set in the `VkSwapchainCreateFlagBitsKHR` field
- [](#VUID-vkWaitForPresent2KHR-presentId-10817)VUID-vkWaitForPresent2KHR-presentId-10817  
  The `presentId` value **must** have been associated with a [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) request on the `swapchain` which returned a non-error value

Valid Usage (Implicit)

- [](#VUID-vkWaitForPresent2KHR-device-parameter)VUID-vkWaitForPresent2KHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkWaitForPresent2KHR-swapchain-parameter)VUID-vkWaitForPresent2KHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkWaitForPresent2KHR-pPresentWait2Info-parameter)VUID-vkWaitForPresent2KHR-pPresentWait2Info-parameter  
  `pPresentWait2Info` **must** be a valid pointer to a valid [VkPresentWait2InfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentWait2InfoKHR.html) structure
- [](#VUID-vkWaitForPresent2KHR-swapchain-parent)VUID-vkWaitForPresent2KHR-swapchain-parent  
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

[VK\_KHR\_present\_wait2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_present_wait2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPresentWait2InfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentWait2InfoKHR.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkWaitForPresent2KHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0