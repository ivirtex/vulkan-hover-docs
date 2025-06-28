# vkGetLatencyTimingsNV(3) Manual Page

## Name

vkGetLatencyTimingsNV - Get latency marker results



## [](#_c_specification)C Specification

To get an array containing the newest collected latency data, call:

```c++
// Provided by VK_NV_low_latency2
void vkGetLatencyTimingsNV(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    VkGetLatencyMarkerInfoNV*                   pLatencyMarkerInfo);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapchain`.
- `swapchain` is the swapchain to return data from.
- `pLatencyMarkerInfo` is a pointer to a [VkGetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGetLatencyMarkerInfoNV.html) structure specifying the parameters for returning latency information.

## [](#_description)Description

The timings returned by `vkGetLatencyTimingsNV` contain the timestamps requested from [vkSetLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetLatencyMarkerNV.html) and additional implementation-specific markers defined in [VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencyTimingsFrameReportNV.html).

Valid Usage (Implicit)

- [](#VUID-vkGetLatencyTimingsNV-device-parameter)VUID-vkGetLatencyTimingsNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetLatencyTimingsNV-swapchain-parameter)VUID-vkGetLatencyTimingsNV-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkGetLatencyTimingsNV-pLatencyMarkerInfo-parameter)VUID-vkGetLatencyTimingsNV-pLatencyMarkerInfo-parameter  
  `pLatencyMarkerInfo` **must** be a valid pointer to a [VkGetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGetLatencyMarkerInfoNV.html) structure
- [](#VUID-vkGetLatencyTimingsNV-swapchain-parent)VUID-vkGetLatencyTimingsNV-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkGetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGetLatencyMarkerInfoNV.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetLatencyTimingsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0