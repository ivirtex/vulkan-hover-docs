# vkReleaseSwapchainImagesKHR(3) Manual Page

## Name

vkReleaseSwapchainImagesKHR - Release previously acquired but unused images



## [](#_c_specification)C Specification

To release images previously acquired through [vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImage2KHR.html) or [vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImageKHR.html), call:

```c++
// Provided by VK_KHR_swapchain_maintenance1
VkResult vkReleaseSwapchainImagesKHR(
    VkDevice                                    device,
    const VkReleaseSwapchainImagesInfoKHR*      pReleaseInfo);
```

or the equivalent

```c++
// Provided by VK_EXT_swapchain_maintenance1
VkResult vkReleaseSwapchainImagesEXT(
    VkDevice                                    device,
    const VkReleaseSwapchainImagesInfoKHR*      pReleaseInfo);
```

## [](#_parameters)Parameters

- `device` is the device associated with [VkReleaseSwapchainImagesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkReleaseSwapchainImagesInfoKHR.html)::`swapchain`.
- `pReleaseInfo` is a pointer to a [VkReleaseSwapchainImagesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkReleaseSwapchainImagesInfoKHR.html) structure containing parameters of the release.

## [](#_description)Description

Only images that are not in use by the device **can** be released.

Releasing images is a read-only operation that will not affect the content of the released images. Upon reacquiring the image, the image contents and its layout will be the same as they were prior to releasing it. However, if a mechanism other than Vulkan is used to modify the platform window associated with the swapchain, the content of all presentable images in the swapchain becomes undefined.

Note

This functionality is useful during swapchain recreation, where acquired images from the old swapchain can be released instead of presented.

Valid Usage

- [](#VUID-vkReleaseSwapchainImagesKHR-swapchainMaintenance1-10159)VUID-vkReleaseSwapchainImagesKHR-swapchainMaintenance1-10159  
  Feature [`swapchainMaintenance1`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-swapchainMaintenance1) **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkReleaseSwapchainImagesKHR-device-parameter)VUID-vkReleaseSwapchainImagesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkReleaseSwapchainImagesKHR-pReleaseInfo-parameter)VUID-vkReleaseSwapchainImagesKHR-pReleaseInfo-parameter  
  `pReleaseInfo` **must** be a valid pointer to a valid [VkReleaseSwapchainImagesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkReleaseSwapchainImagesInfoKHR.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_SURFACE_LOST_KHR`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_swapchain_maintenance1.html), [VK\_KHR\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain_maintenance1.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkReleaseSwapchainImagesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkReleaseSwapchainImagesInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkReleaseSwapchainImagesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0