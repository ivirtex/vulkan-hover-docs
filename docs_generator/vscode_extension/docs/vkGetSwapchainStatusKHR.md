# vkGetSwapchainStatusKHR(3) Manual Page

## Name

vkGetSwapchainStatusKHR - Get a swapchain's status



## [](#_c_specification)C Specification

In order to query a swapchain’s status when rendering to a shared presentable image, call:

```c++
// Provided by VK_KHR_shared_presentable_image
VkResult vkGetSwapchainStatusKHR(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapchain`.
- `swapchain` is the swapchain to query.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetSwapchainStatusKHR-device-parameter)VUID-vkGetSwapchainStatusKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetSwapchainStatusKHR-swapchain-parameter)VUID-vkGetSwapchainStatusKHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkGetSwapchainStatusKHR-swapchain-parent)VUID-vkGetSwapchainStatusKHR-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_SUBOPTIMAL_KHR`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_DEVICE_LOST`
- `VK_ERROR_OUT_OF_DATE_KHR`
- `VK_ERROR_SURFACE_LOST_KHR`
- `VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT`

## [](#_see_also)See Also

[VK\_KHR\_shared\_presentable\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shared_presentable_image.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetSwapchainStatusKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0