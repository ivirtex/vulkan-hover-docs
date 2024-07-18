# vkWaitForPresentKHR(3) Manual Page

## Name

vkWaitForPresentKHR - Wait for presentation



## <a href="#_c_specification" class="anchor"></a>C Specification

When the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-presentWait"
target="_blank" rel="noopener"><code>presentWait</code></a> feature is
enabled, an application **can** wait for an image to be presented to the
user by first specifying a presentId for the target presentation by
adding a `VkPresentIdKHR` structure to the `pNext` chain of the
[VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html) structure and then waiting for
that presentation to complete by calling:

``` c
// Provided by VK_KHR_present_wait
VkResult vkWaitForPresentKHR(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    uint64_t                                    presentId,
    uint64_t                                    timeout);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapchain`.

- `swapchain` is the non-retired swapchain on which an image was queued
  for presentation.

- `presentId` is the presentation presentId to wait for.

- `timeout` is the timeout period in units of nanoseconds. `timeout` is
  adjusted to the closest value allowed by the implementation-dependent
  timeout accuracy, which **may** be substantially longer than one
  nanosecond, and **may** be longer than the requested period.

## <a href="#_description" class="anchor"></a>Description

`vkWaitForPresentKHR` waits for the presentId associated with
`swapchain` to be increased in value so that it is at least equal to
`presentId`.

For `VK_PRESENT_MODE_MAILBOX_KHR` (or other present mode where images
may be replaced in the presentation queue) any wait of this type
associated with such an image **must** be signaled no later than a wait
associated with the replacing image would be signaled.

When the presentation has completed, the presentId associated with the
related `pSwapchains` entry will be increased in value so that it is at
least equal to the value provided in the `VkPresentIdKHR` structure.

There is no requirement for any precise timing relationship between the
presentation of the image to the user and the update of the presentId
value, but implementations **should** make this as close as possible to
the presentation of the first pixel in the next image being presented to
the user.

The call to `vkWaitForPresentKHR` will block until either the presentId
associated with `swapchain` is greater than or equal to `presentId`, or
`timeout` nanoseconds passes. When the swapchain becomes OUT_OF_DATE,
the call will either return `VK_SUCCESS` (if the image was delivered to
the presentation engine and may have been presented to the user) or will
return early with status `VK_ERROR_OUT_OF_DATE_KHR` (if the image was
not presented to the user).

As an exception to the normal rules for objects which are externally
synchronized, the `swapchain` passed to `vkWaitForPresentKHR` **may** be
simultaneously used by other threads in calls to functions other than
[vkDestroySwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySwapchainKHR.html). Access to the
swapchain data associated with this extension **must** be atomic within
the implementation.

Valid Usage

- <a href="#VUID-vkWaitForPresentKHR-swapchain-04997"
  id="VUID-vkWaitForPresentKHR-swapchain-04997"></a>
  VUID-vkWaitForPresentKHR-swapchain-04997  
  `swapchain` **must** not be in the retired state

- <a href="#VUID-vkWaitForPresentKHR-presentWait-06234"
  id="VUID-vkWaitForPresentKHR-presentWait-06234"></a>
  VUID-vkWaitForPresentKHR-presentWait-06234  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-presentWait"
  target="_blank" rel="noopener"><code>presentWait</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkWaitForPresentKHR-device-parameter"
  id="VUID-vkWaitForPresentKHR-device-parameter"></a>
  VUID-vkWaitForPresentKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkWaitForPresentKHR-swapchain-parameter"
  id="VUID-vkWaitForPresentKHR-swapchain-parameter"></a>
  VUID-vkWaitForPresentKHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-vkWaitForPresentKHR-swapchain-parent"
  id="VUID-vkWaitForPresentKHR-swapchain-parent"></a>
  VUID-vkWaitForPresentKHR-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from
  `device`

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_TIMEOUT`

- `VK_SUBOPTIMAL_KHR`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

- `VK_ERROR_OUT_OF_DATE_KHR`

- `VK_ERROR_SURFACE_LOST_KHR`

- `VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_present_wait](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_present_wait.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkWaitForPresentKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
