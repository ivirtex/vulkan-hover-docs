# vkGetSwapchainImagesKHR(3) Manual Page

## Name

vkGetSwapchainImagesKHR - Obtain the array of presentable images associated with a swapchain



## [](#_c_specification)C Specification

To obtain the array of presentable images associated with a swapchain, call:

```c++
// Provided by VK_KHR_swapchain
VkResult vkGetSwapchainImagesKHR(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    uint32_t*                                   pSwapchainImageCount,
    VkImage*                                    pSwapchainImages);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapchain`.
- `swapchain` is the swapchain to query.
- `pSwapchainImageCount` is a pointer to an integer related to the number of presentable images available or queried, as described below.
- `pSwapchainImages` is either `NULL` or a pointer to an array of `VkImage` handles.

## [](#_description)Description

If `pSwapchainImages` is `NULL`, then the number of presentable images for `swapchain` is returned in `pSwapchainImageCount`. Otherwise, `pSwapchainImageCount` **must** point to a variable set by the application to the number of elements in the `pSwapchainImages` array, and on return the variable is overwritten with the number of structures actually written to `pSwapchainImages`. If the value of `pSwapchainImageCount` is less than the number of presentable images for `swapchain`, at most `pSwapchainImageCount` structures will be written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available presentable images were returned.

Valid Usage (Implicit)

- [](#VUID-vkGetSwapchainImagesKHR-device-parameter)VUID-vkGetSwapchainImagesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetSwapchainImagesKHR-swapchain-parameter)VUID-vkGetSwapchainImagesKHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkGetSwapchainImagesKHR-pSwapchainImageCount-parameter)VUID-vkGetSwapchainImagesKHR-pSwapchainImageCount-parameter  
  `pSwapchainImageCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetSwapchainImagesKHR-pSwapchainImages-parameter)VUID-vkGetSwapchainImagesKHR-pSwapchainImages-parameter  
  If the value referenced by `pSwapchainImageCount` is not `0`, and `pSwapchainImages` is not `NULL`, `pSwapchainImages` **must** be a valid pointer to an array of `pSwapchainImageCount` [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handles
- [](#VUID-vkGetSwapchainImagesKHR-swapchain-parent)VUID-vkGetSwapchainImagesKHR-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetSwapchainImagesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0