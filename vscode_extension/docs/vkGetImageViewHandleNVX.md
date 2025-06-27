# vkGetImageViewHandleNVX(3) Manual Page

## Name

vkGetImageViewHandleNVX - Get the handle for an image view for a
specific descriptor type



## <a href="#_c_specification" class="anchor"></a>C Specification

To get the handle for an image view, call:

``` c
// Provided by VK_NVX_image_view_handle
uint32_t vkGetImageViewHandleNVX(
    VkDevice                                    device,
    const VkImageViewHandleInfoNVX*             pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the image view.

- `pInfo` describes the image view to query and type of handle.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetImageViewHandleNVX-device-parameter"
  id="VUID-vkGetImageViewHandleNVX-device-parameter"></a>
  VUID-vkGetImageViewHandleNVX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetImageViewHandleNVX-pInfo-parameter"
  id="VUID-vkGetImageViewHandleNVX-pInfo-parameter"></a>
  VUID-vkGetImageViewHandleNVX-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkImageViewHandleInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewHandleInfoNVX.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NVX_image_view_handle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NVX_image_view_handle.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkImageViewHandleInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewHandleInfoNVX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetImageViewHandleNVX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
