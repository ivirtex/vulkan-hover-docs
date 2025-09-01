# vkGetDrmDisplayEXT(3) Manual Page

## Name

vkGetDrmDisplayEXT - Query the VkDisplayKHR corresponding to a DRM connector ID



## [](#_c_specification)C Specification

Before acquiring a display from the DRM interface, the caller may want to select a specific `VkDisplayKHR` handle by identifying it using a `connectorId`. To do so, call:

```c++
// Provided by VK_EXT_acquire_drm_display
VkResult vkGetDrmDisplayEXT(
    VkPhysicalDevice                            physicalDevice,
    int32_t                                     drmFd,
    uint32_t                                    connectorId,
    VkDisplayKHR*                               display);
```

## [](#_parameters)Parameters

- `physicalDevice` The physical device to query the display from.
- `drmFd` DRM primary file descriptor.
- `connectorId` Identifier of the specified DRM connector.
- `display` The corresponding [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) handle will be returned here.

## [](#_description)Description

If there is no [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) corresponding to the `connectorId` on the `physicalDevice`, the returning `display` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html). The provided `drmFd` **must** correspond to the one owned by the `physicalDevice`. If not, the error code `VK_ERROR_UNKNOWN` **must** be returned. Mast⁠er permissions are not required, because the file descriptor is just used for information gathering purposes. The given `connectorId` **must** be a resource owned by the provided `drmFd`. If not, the error code `VK_ERROR_UNKNOWN` **must** be returned. If any error is encountered during the identification of the display, the call **must** return the error code `VK_ERROR_INITIALIZATION_FAILED`.

Valid Usage (Implicit)

- [](#VUID-vkGetDrmDisplayEXT-physicalDevice-parameter)VUID-vkGetDrmDisplayEXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetDrmDisplayEXT-display-parameter)VUID-vkGetDrmDisplayEXT-display-parameter  
  `display` **must** be a valid pointer to a [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_acquire\_drm\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_acquire_drm_display.html), [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDrmDisplayEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0