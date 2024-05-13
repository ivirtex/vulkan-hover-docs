# vkGetLatencyTimingsNV(3) Manual Page

## Name

vkGetLatencyTimingsNV - Get latency marker results



## <a href="#_c_specification" class="anchor"></a>C Specification

To get an array containing the newest collected latency data, call:

``` c
// Provided by VK_NV_low_latency2
void vkGetLatencyTimingsNV(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    VkGetLatencyMarkerInfoNV*                   pLatencyMarkerInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapchain`.

- `swapchain` is the swapchain to return data from.

- `pGetLatencyMarkerInfo` is a pointer to a
  [VkGetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGetLatencyMarkerInfoNV.html) structure
  specifying the parameters for returning latency information.

## <a href="#_description" class="anchor"></a>Description

The timings returned by `vkGetLatencyTimingsNV` contain the timestamps
requested from [vkSetLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetLatencyMarkerNV.html) and
additional implementation-specific markers defined in
[VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencyTimingsFrameReportNV.html).

Valid Usage (Implicit)

- <a href="#VUID-vkGetLatencyTimingsNV-device-parameter"
  id="VUID-vkGetLatencyTimingsNV-device-parameter"></a>
  VUID-vkGetLatencyTimingsNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetLatencyTimingsNV-swapchain-parameter"
  id="VUID-vkGetLatencyTimingsNV-swapchain-parameter"></a>
  VUID-vkGetLatencyTimingsNV-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-vkGetLatencyTimingsNV-pLatencyMarkerInfo-parameter"
  id="VUID-vkGetLatencyTimingsNV-pLatencyMarkerInfo-parameter"></a>
  VUID-vkGetLatencyTimingsNV-pLatencyMarkerInfo-parameter  
  `pLatencyMarkerInfo` **must** be a valid pointer to a
  [VkGetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGetLatencyMarkerInfoNV.html) structure

- <a href="#VUID-vkGetLatencyTimingsNV-swapchain-parent"
  id="VUID-vkGetLatencyTimingsNV-swapchain-parent"></a>
  VUID-vkGetLatencyTimingsNV-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkGetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGetLatencyMarkerInfoNV.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetLatencyTimingsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
