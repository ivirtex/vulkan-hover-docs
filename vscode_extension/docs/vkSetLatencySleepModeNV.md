# vkSetLatencySleepModeNV(3) Manual Page

## Name

vkSetLatencySleepModeNV - Enable or Disable low latency mode on a swapchain



## [](#_c_specification)C Specification

To enable or disable low latency mode on a swapchain, call:

```c++
// Provided by VK_NV_low_latency2
VkResult vkSetLatencySleepModeNV(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    const VkLatencySleepModeInfoNV*             pSleepModeInfo);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapchain`.
- `swapchain` is the swapchain to enable or disable low latency mode on.
- `pSleepModeInfo` is `NULL` or a pointer to a [VkLatencySleepModeInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySleepModeInfoNV.html) structure specifying the parameters of the latency sleep mode.

## [](#_description)Description

If `pSleepModeInfo` is `NULL`, `vkSetLatencySleepModeNV` will disable low latency mode, low latency boost, and set the minimum present interval previously specified by [VkLatencySleepModeInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySleepModeInfoNV.html) to zero on `swapchain`. As an exception to the normal rules for objects which are externally synchronized, the swapchain passed to `vkSetLatencySleepModeNV` **may** be simultaneously used by other threads in calls to functions other than [vkDestroySwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroySwapchainKHR.html). Access to the swapchain data associated with this extension **must** be atomic within the implementation.

Valid Usage (Implicit)

- [](#VUID-vkSetLatencySleepModeNV-device-parameter)VUID-vkSetLatencySleepModeNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkSetLatencySleepModeNV-swapchain-parameter)VUID-vkSetLatencySleepModeNV-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkSetLatencySleepModeNV-pSleepModeInfo-parameter)VUID-vkSetLatencySleepModeNV-pSleepModeInfo-parameter  
  `pSleepModeInfo` **must** be a valid pointer to a valid [VkLatencySleepModeInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySleepModeInfoNV.html) structure
- [](#VUID-vkSetLatencySleepModeNV-swapchain-parent)VUID-vkSetLatencySleepModeNV-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INITIALIZATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkLatencySleepModeInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySleepModeInfoNV.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkSetLatencySleepModeNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0