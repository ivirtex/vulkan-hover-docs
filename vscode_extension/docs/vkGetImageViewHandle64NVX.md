# vkGetImageViewHandle64NVX(3) Manual Page

## Name

vkGetImageViewHandle64NVX - Get the 64-bit handle for an image view for a specific descriptor type



## [](#_c_specification)C Specification

To get the 64-bit handle for an image view, call:

```c++
// Provided by VK_NVX_image_view_handle
uint64_t vkGetImageViewHandle64NVX(
    VkDevice                                    device,
    const VkImageViewHandleInfoNVX*             pInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the image view.
- `pInfo` describes the image view to query and type of handle.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetImageViewHandle64NVX-device-parameter)VUID-vkGetImageViewHandle64NVX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetImageViewHandle64NVX-pInfo-parameter)VUID-vkGetImageViewHandle64NVX-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkImageViewHandleInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewHandleInfoNVX.html) structure

## [](#_see_also)See Also

[VK\_NVX\_image\_view\_handle](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_image_view_handle.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImageViewHandleInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewHandleInfoNVX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetImageViewHandle64NVX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0