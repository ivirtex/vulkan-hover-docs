# vkSetLatencySleepModeNV(3) Manual Page

## Name

vkSetLatencySleepModeNV - Enable or Disable low latency mode on a
swapchain



## <a href="#_c_specification" class="anchor"></a>C Specification

To enable or disable low latency mode on a swapchain, call:

``` c
// Provided by VK_NV_low_latency2
VkResult vkSetLatencySleepModeNV(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    const VkLatencySleepModeInfoNV*             pSleepModeInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapchain`.

- `swapchain` is the swapchain to enable or disable low latency mode on.

- `pSleepModeInfo` is `NULL` or a pointer to a
  [VkLatencySleepModeInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySleepModeInfoNV.html) structure
  specifying the parameters of the latency sleep mode.

## <a href="#_description" class="anchor"></a>Description

If `pSleepModeInfo` is `NULL`, `vkSetLatencySleepModeNV` will disable
low latency mode, low latency boost, and set the minimum present
interval previously specified by
[VkLatencySleepModeInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySleepModeInfoNV.html) to zero on
`swapchain`. As an exception to the normal rules for objects which are
externally synchronized, the swapchain passed to
`vkSetLatencySleepModeNV` **may** be simultaneously used by other
threads in calls to functions other than
[vkDestroySwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySwapchainKHR.html). Access to the
swapchain data associated with this extension **must** be atomic within
the implementation.

Valid Usage (Implicit)

- <a href="#VUID-vkSetLatencySleepModeNV-device-parameter"
  id="VUID-vkSetLatencySleepModeNV-device-parameter"></a>
  VUID-vkSetLatencySleepModeNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkSetLatencySleepModeNV-swapchain-parameter"
  id="VUID-vkSetLatencySleepModeNV-swapchain-parameter"></a>
  VUID-vkSetLatencySleepModeNV-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-vkSetLatencySleepModeNV-pSleepModeInfo-parameter"
  id="VUID-vkSetLatencySleepModeNV-pSleepModeInfo-parameter"></a>
  VUID-vkSetLatencySleepModeNV-pSleepModeInfo-parameter  
  `pSleepModeInfo` **must** be a valid pointer to a valid
  [VkLatencySleepModeInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySleepModeInfoNV.html) structure

- <a href="#VUID-vkSetLatencySleepModeNV-swapchain-parent"
  id="VUID-vkSetLatencySleepModeNV-swapchain-parent"></a>
  VUID-vkSetLatencySleepModeNV-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkLatencySleepModeInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySleepModeInfoNV.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkSetLatencySleepModeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
