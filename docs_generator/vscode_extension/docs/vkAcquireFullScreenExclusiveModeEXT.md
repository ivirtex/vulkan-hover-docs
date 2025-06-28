# vkAcquireFullScreenExclusiveModeEXT(3) Manual Page

## Name

vkAcquireFullScreenExclusiveModeEXT - Acquire full-screen exclusive mode for a swapchain



## [](#_c_specification)C Specification

To acquire exclusive full-screen access for a swapchain, call:

```c++
// Provided by VK_EXT_full_screen_exclusive
VkResult vkAcquireFullScreenExclusiveModeEXT(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapchain`.
- `swapchain` is the swapchain to acquire exclusive full-screen access for.

## [](#_description)Description

Valid Usage

- [](#VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02674)VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02674  
  `swapchain` **must** not be in the retired state
- [](#VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02675)VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02675  
  `swapchain` **must** be a swapchain created with a [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html) structure, with `fullScreenExclusive` set to `VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT`
- [](#VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02676)VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-02676  
  `swapchain` **must** not currently have exclusive full-screen access

A return value of `VK_SUCCESS` indicates that the `swapchain` successfully acquired exclusive full-screen access. The swapchain will retain this exclusivity until either the application releases exclusive full-screen access with [vkReleaseFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseFullScreenExclusiveModeEXT.html), destroys the swapchain, or if any of the swapchain commands return `VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT` indicating that the mode was lost because of platform-specific changes.

If the swapchain was unable to acquire exclusive full-screen access to the display then `VK_ERROR_INITIALIZATION_FAILED` is returned. An application **can** attempt to acquire exclusive full-screen access again for the same swapchain even if this command fails, or if `VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT` has been returned by a swapchain command.

Valid Usage (Implicit)

- [](#VUID-vkAcquireFullScreenExclusiveModeEXT-device-parameter)VUID-vkAcquireFullScreenExclusiveModeEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-parameter)VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-parent)VUID-vkAcquireFullScreenExclusiveModeEXT-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_SURFACE_LOST_KHR`

## [](#_see_also)See Also

[VK\_EXT\_full\_screen\_exclusive](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_full_screen_exclusive.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkAcquireFullScreenExclusiveModeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0