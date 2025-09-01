# vkAcquireDrmDisplayEXT(3) Manual Page

## Name

vkAcquireDrmDisplayEXT - Acquire access to a VkDisplayKHR using DRM



## [](#_c_specification)C Specification

To acquire permission to directly a display in Vulkan from the Direct Rendering Manager (DRM) interface, call:

```c++
// Provided by VK_EXT_acquire_drm_display
VkResult vkAcquireDrmDisplayEXT(
    VkPhysicalDevice                            physicalDevice,
    int32_t                                     drmFd,
    VkDisplayKHR                                display);
```

## [](#_parameters)Parameters

- `physicalDevice` The physical device the display is on.
- `drmFd` DRM primary file descriptor.
- `display` The display the caller wishes Vulkan to control.

## [](#_description)Description

All permissions necessary to control the display are granted to the Vulkan instance associated with the provided `physicalDevice` until the display is either released or the connector is unplugged. The provided `drmFd` **must** correspond to the one owned by the `physicalDevice`. If not, the error code `VK_ERROR_UNKNOWN` **must** be returned. The DRM FD must have DRM mast⁠er permissions. If any error is encountered during the acquisition of the display, the call **must** return the error code `VK_ERROR_INITIALIZATION_FAILED`.

The provided DRM fd should not be closed before the display is released, attempting to do it may result in undefined behavior.

Valid Usage (Implicit)

- [](#VUID-vkAcquireDrmDisplayEXT-physicalDevice-parameter)VUID-vkAcquireDrmDisplayEXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkAcquireDrmDisplayEXT-display-parameter)VUID-vkAcquireDrmDisplayEXT-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) handle
- [](#VUID-vkAcquireDrmDisplayEXT-display-parent)VUID-vkAcquireDrmDisplayEXT-display-parent  
  `display` **must** have been created, allocated, or retrieved from `physicalDevice`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_acquire\_drm\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_acquire_drm_display.html), [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkAcquireDrmDisplayEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0