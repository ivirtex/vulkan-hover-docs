# vkGetImageViewAddressNVX(3) Manual Page

## Name

vkGetImageViewAddressNVX - Get the device address of an image view



## [](#_c_specification)C Specification

To get the device address for an image view, call:

```c++
// Provided by VK_NVX_image_view_handle
VkResult vkGetImageViewAddressNVX(
    VkDevice                                    device,
    VkImageView                                 imageView,
    VkImageViewAddressPropertiesNVX*            pProperties);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the image view.
- `imageView` is a handle to the image view.
- `pProperties` contains the device address and size when the call returns.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetImageViewAddressNVX-device-parameter)VUID-vkGetImageViewAddressNVX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetImageViewAddressNVX-imageView-parameter)VUID-vkGetImageViewAddressNVX-imageView-parameter  
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle
- [](#VUID-vkGetImageViewAddressNVX-pProperties-parameter)VUID-vkGetImageViewAddressNVX-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a [VkImageViewAddressPropertiesNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewAddressPropertiesNVX.html) structure
- [](#VUID-vkGetImageViewAddressNVX-imageView-parent)VUID-vkGetImageViewAddressNVX-imageView-parent  
  `imageView` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NVX\_image\_view\_handle](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_image_view_handle.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkImageViewAddressPropertiesNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewAddressPropertiesNVX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetImageViewAddressNVX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0