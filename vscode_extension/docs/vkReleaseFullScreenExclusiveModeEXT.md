# vkReleaseFullScreenExclusiveModeEXT(3) Manual Page

## Name

vkReleaseFullScreenExclusiveModeEXT - Release full-screen exclusive mode from a swapchain



## [](#_c_specification)C Specification

To release exclusive full-screen access from a swapchain, call:

```c++
// Provided by VK_EXT_full_screen_exclusive
VkResult vkReleaseFullScreenExclusiveModeEXT(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapchain`.
- `swapchain` is the swapchain to release exclusive full-screen access from.

## [](#_description)Description

Note

Applications will not be able to present to `swapchain` after this call until exclusive full-screen access is reacquired. This is usually useful to handle when an application is minimized or otherwise intends to stop presenting for a time.

Valid Usage

- [](#VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-02677)VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-02677  
  `swapchain` **must** not be in the retired state
- [](#VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-02678)VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-02678  
  `swapchain` **must** be a swapchain created with a [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html) structure, with `fullScreenExclusive` set to `VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT`

Valid Usage (Implicit)

- [](#VUID-vkReleaseFullScreenExclusiveModeEXT-device-parameter)VUID-vkReleaseFullScreenExclusiveModeEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-parameter)VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-parent)VUID-vkReleaseFullScreenExclusiveModeEXT-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_SURFACE_LOST_KHR`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_full\_screen\_exclusive](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_full_screen_exclusive.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkReleaseFullScreenExclusiveModeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0