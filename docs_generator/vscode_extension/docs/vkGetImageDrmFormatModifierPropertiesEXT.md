# vkGetImageDrmFormatModifierPropertiesEXT(3) Manual Page

## Name

vkGetImageDrmFormatModifierPropertiesEXT - Returns an image's DRM format modifier



## [](#_c_specification)C Specification

If an image was created with `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then the image has a [Linux DRM format modifier](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-drm-format-modifier). To query the *modifier*, call:

```c++
// Provided by VK_EXT_image_drm_format_modifier
VkResult vkGetImageDrmFormatModifierPropertiesEXT(
    VkDevice                                    device,
    VkImage                                     image,
    VkImageDrmFormatModifierPropertiesEXT*      pProperties);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the image.
- `image` is the queried image.
- `pProperties` is a pointer to a [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierPropertiesEXT.html) structure in which properties of the image’s *DRM format modifier* are returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-02272)VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-02272  
  `image` **must** have been created with [`tiling`](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) equal to `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`

Valid Usage (Implicit)

- [](#VUID-vkGetImageDrmFormatModifierPropertiesEXT-device-parameter)VUID-vkGetImageDrmFormatModifierPropertiesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-parameter)VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-vkGetImageDrmFormatModifierPropertiesEXT-pProperties-parameter)VUID-vkGetImageDrmFormatModifierPropertiesEXT-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierPropertiesEXT.html) structure
- [](#VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-parent)VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-parent  
  `image` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## [](#_see_also)See Also

[VK\_EXT\_image\_drm\_format\_modifier](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_drm_format_modifier.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierPropertiesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetImageDrmFormatModifierPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0