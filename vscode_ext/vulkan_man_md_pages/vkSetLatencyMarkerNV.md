# vkSetLatencyMarkerNV(3) Manual Page

## Name

vkSetLatencyMarkerNV - Pass in marker for timing info



## <a href="#_c_specification" class="anchor"></a>C Specification

An application **can** provide timestamps at various stages of its frame
generation work by calling:

``` c
// Provided by VK_NV_low_latency2
void vkSetLatencyMarkerNV(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    const VkSetLatencyMarkerInfoNV*             pLatencyMarkerInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapchain`.

- `swapchain` is the swapchain to capture timestamps on.

- `pSetLatencyMarkerInfo` is a pointer to a
  [VkSetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSetLatencyMarkerInfoNV.html) structure
  specifying the parameters of the marker to set.

## <a href="#_description" class="anchor"></a>Description

At the beginning and end of simulation and render threads and beginning
and end of [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html) calls,
`vkSetLatencyMarkerNV` **can** be called to provide timestamps for the
application’s reference. These timestamps are returned with a call to
[vkGetLatencyTimingsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetLatencyTimingsNV.html) alongside driver
provided timestamps at various points of interest with regards to
latency within the application. As an exception to the normal rules for
objects which are externally synchronized, the swapchain passed to
`vkSetLatencyMarkerNV` **may** be simultaneously used by other threads
in calls to functions other than
[vkDestroySwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySwapchainKHR.html). Access to the
swapchain data associated with this extension **must** be atomic within
the implementation.

Valid Usage (Implicit)

- <a href="#VUID-vkSetLatencyMarkerNV-device-parameter"
  id="VUID-vkSetLatencyMarkerNV-device-parameter"></a>
  VUID-vkSetLatencyMarkerNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkSetLatencyMarkerNV-swapchain-parameter"
  id="VUID-vkSetLatencyMarkerNV-swapchain-parameter"></a>
  VUID-vkSetLatencyMarkerNV-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-vkSetLatencyMarkerNV-pLatencyMarkerInfo-parameter"
  id="VUID-vkSetLatencyMarkerNV-pLatencyMarkerInfo-parameter"></a>
  VUID-vkSetLatencyMarkerNV-pLatencyMarkerInfo-parameter  
  `pLatencyMarkerInfo` **must** be a valid pointer to a valid
  [VkSetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSetLatencyMarkerInfoNV.html) structure

- <a href="#VUID-vkSetLatencyMarkerNV-swapchain-parent"
  id="VUID-vkSetLatencyMarkerNV-swapchain-parent"></a>
  VUID-vkSetLatencyMarkerNV-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkSetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSetLatencyMarkerInfoNV.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkSetLatencyMarkerNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
