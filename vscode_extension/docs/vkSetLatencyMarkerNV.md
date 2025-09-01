# vkSetLatencyMarkerNV(3) Manual Page

## Name

vkSetLatencyMarkerNV - Pass in marker for timing info



## [](#_c_specification)C Specification

An application **can** provide timestamps at various stages of its frame generation work by calling:

```c++
// Provided by VK_NV_low_latency2
void vkSetLatencyMarkerNV(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    const VkSetLatencyMarkerInfoNV*             pLatencyMarkerInfo);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapchain`.
- `swapchain` is the swapchain to capture timestamps on.
- `pSetLatencyMarkerInfo` is a pointer to a [VkSetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSetLatencyMarkerInfoNV.html) structure specifying the parameters of the marker to set.

## [](#_description)Description

At the beginning and end of simulation and render threads and beginning and end of [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) calls, `vkSetLatencyMarkerNV` **can** be called to provide timestamps for the application’s reference. These timestamps are returned with a call to [vkGetLatencyTimingsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetLatencyTimingsNV.html) alongside driver provided timestamps at various points of interest with regards to latency within the application. As an exception to the normal rules for objects which are externally synchronized, the swapchain passed to `vkSetLatencyMarkerNV` **may** be simultaneously used by other threads in calls to functions other than [vkDestroySwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroySwapchainKHR.html). Access to the swapchain data associated with this extension **must** be atomic within the implementation.

Valid Usage (Implicit)

- [](#VUID-vkSetLatencyMarkerNV-device-parameter)VUID-vkSetLatencyMarkerNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkSetLatencyMarkerNV-swapchain-parameter)VUID-vkSetLatencyMarkerNV-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkSetLatencyMarkerNV-pLatencyMarkerInfo-parameter)VUID-vkSetLatencyMarkerNV-pLatencyMarkerInfo-parameter  
  `pLatencyMarkerInfo` **must** be a valid pointer to a valid [VkSetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSetLatencyMarkerInfoNV.html) structure
- [](#VUID-vkSetLatencyMarkerNV-swapchain-parent)VUID-vkSetLatencyMarkerNV-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSetLatencyMarkerInfoNV.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkSetLatencyMarkerNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0