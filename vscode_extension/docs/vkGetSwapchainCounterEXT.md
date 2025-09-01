# vkGetSwapchainCounterEXT(3) Manual Page

## Name

vkGetSwapchainCounterEXT - Query the current value of a surface counter



## [](#_c_specification)C Specification

The requested counters become active when the first presentation command for the associated swapchain is processed by the presentation engine. To query the value of an active counter, use:

```c++
// Provided by VK_EXT_display_control
VkResult vkGetSwapchainCounterEXT(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    VkSurfaceCounterFlagBitsEXT                 counter,
    uint64_t*                                   pCounterValue);
```

## [](#_parameters)Parameters

- `device` is the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) associated with `swapchain`.
- `swapchain` is the swapchain from which to query the counter value.
- `counter` is a [VkSurfaceCounterFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCounterFlagBitsEXT.html) value specifying the counter to query.
- `pCounterValue` will return the current value of the counter.

## [](#_description)Description

If a counter is not available because the swapchain is out of date, the implementation **may** return `VK_ERROR_OUT_OF_DATE_KHR`.

Valid Usage

- [](#VUID-vkGetSwapchainCounterEXT-swapchain-01245)VUID-vkGetSwapchainCounterEXT-swapchain-01245  
  One or more present commands on `swapchain` **must** have been processed by the presentation engine

Valid Usage (Implicit)

- [](#VUID-vkGetSwapchainCounterEXT-device-parameter)VUID-vkGetSwapchainCounterEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetSwapchainCounterEXT-swapchain-parameter)VUID-vkGetSwapchainCounterEXT-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkGetSwapchainCounterEXT-counter-parameter)VUID-vkGetSwapchainCounterEXT-counter-parameter  
  `counter` **must** be a valid [VkSurfaceCounterFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCounterFlagBitsEXT.html) value
- [](#VUID-vkGetSwapchainCounterEXT-pCounterValue-parameter)VUID-vkGetSwapchainCounterEXT-pCounterValue-parameter  
  `pCounterValue` **must** be a valid pointer to a `uint64_t` value
- [](#VUID-vkGetSwapchainCounterEXT-swapchain-parent)VUID-vkGetSwapchainCounterEXT-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_DEVICE_LOST`
- `VK_ERROR_OUT_OF_DATE_KHR`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_display\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_display_control.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSurfaceCounterFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCounterFlagBitsEXT.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetSwapchainCounterEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0