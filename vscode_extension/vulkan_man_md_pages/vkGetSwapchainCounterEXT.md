# vkGetSwapchainCounterEXT(3) Manual Page

## Name

vkGetSwapchainCounterEXT - Query the current value of a surface counter



## <a href="#_c_specification" class="anchor"></a>C Specification

The requested counters become active when the first presentation command
for the associated swapchain is processed by the presentation engine. To
query the value of an active counter, use:

``` c
// Provided by VK_EXT_display_control
VkResult vkGetSwapchainCounterEXT(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    VkSurfaceCounterFlagBitsEXT                 counter,
    uint64_t*                                   pCounterValue);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) associated with `swapchain`.

- `swapchain` is the swapchain from which to query the counter value.

- `counter` is a
  [VkSurfaceCounterFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCounterFlagBitsEXT.html) value
  specifying the counter to query.

- `pCounterValue` will return the current value of the counter.

## <a href="#_description" class="anchor"></a>Description

If a counter is not available because the swapchain is out of date, the
implementation **may** return `VK_ERROR_OUT_OF_DATE_KHR`.

Valid Usage

- <a href="#VUID-vkGetSwapchainCounterEXT-swapchain-01245"
  id="VUID-vkGetSwapchainCounterEXT-swapchain-01245"></a>
  VUID-vkGetSwapchainCounterEXT-swapchain-01245  
  One or more present commands on `swapchain` **must** have been
  processed by the presentation engine

Valid Usage (Implicit)

- <a href="#VUID-vkGetSwapchainCounterEXT-device-parameter"
  id="VUID-vkGetSwapchainCounterEXT-device-parameter"></a>
  VUID-vkGetSwapchainCounterEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetSwapchainCounterEXT-swapchain-parameter"
  id="VUID-vkGetSwapchainCounterEXT-swapchain-parameter"></a>
  VUID-vkGetSwapchainCounterEXT-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-vkGetSwapchainCounterEXT-counter-parameter"
  id="VUID-vkGetSwapchainCounterEXT-counter-parameter"></a>
  VUID-vkGetSwapchainCounterEXT-counter-parameter  
  `counter` **must** be a valid
  [VkSurfaceCounterFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCounterFlagBitsEXT.html) value

- <a href="#VUID-vkGetSwapchainCounterEXT-pCounterValue-parameter"
  id="VUID-vkGetSwapchainCounterEXT-pCounterValue-parameter"></a>
  VUID-vkGetSwapchainCounterEXT-pCounterValue-parameter  
  `pCounterValue` **must** be a valid pointer to a `uint64_t` value

- <a href="#VUID-vkGetSwapchainCounterEXT-swapchain-parent"
  id="VUID-vkGetSwapchainCounterEXT-swapchain-parent"></a>
  VUID-vkGetSwapchainCounterEXT-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_DEVICE_LOST`

- `VK_ERROR_OUT_OF_DATE_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_control.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkSurfaceCounterFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCounterFlagBitsEXT.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetSwapchainCounterEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
