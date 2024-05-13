# vkGetImageViewAddressNVX(3) Manual Page

## Name

vkGetImageViewAddressNVX - Get the device address of an image view



## <a href="#_c_specification" class="anchor"></a>C Specification

To get the device address for an image view, call:

``` c
// Provided by VK_NVX_image_view_handle
VkResult vkGetImageViewAddressNVX(
    VkDevice                                    device,
    VkImageView                                 imageView,
    VkImageViewAddressPropertiesNVX*            pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the image view.

- `imageView` is a handle to the image view.

- `pProperties` contains the device address and size when the call
  returns.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetImageViewAddressNVX-device-parameter"
  id="VUID-vkGetImageViewAddressNVX-device-parameter"></a>
  VUID-vkGetImageViewAddressNVX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetImageViewAddressNVX-imageView-parameter"
  id="VUID-vkGetImageViewAddressNVX-imageView-parameter"></a>
  VUID-vkGetImageViewAddressNVX-imageView-parameter  
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle

- <a href="#VUID-vkGetImageViewAddressNVX-pProperties-parameter"
  id="VUID-vkGetImageViewAddressNVX-pProperties-parameter"></a>
  VUID-vkGetImageViewAddressNVX-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a
  [VkImageViewAddressPropertiesNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewAddressPropertiesNVX.html)
  structure

- <a href="#VUID-vkGetImageViewAddressNVX-imageView-parent"
  id="VUID-vkGetImageViewAddressNVX-imageView-parent"></a>
  VUID-vkGetImageViewAddressNVX-imageView-parent  
  `imageView` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NVX_image_view_handle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NVX_image_view_handle.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html),
[VkImageViewAddressPropertiesNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewAddressPropertiesNVX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetImageViewAddressNVX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
