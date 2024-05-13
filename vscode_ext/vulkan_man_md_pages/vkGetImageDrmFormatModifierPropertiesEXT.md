# vkGetImageDrmFormatModifierPropertiesEXT(3) Manual Page

## Name

vkGetImageDrmFormatModifierPropertiesEXT - Returns an image's DRM format
modifier



## <a href="#_c_specification" class="anchor"></a>C Specification

If an image was created with `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`,
then the image has a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-drm-format-modifier"
target="_blank" rel="noopener">Linux DRM format modifier</a>. To query
the *modifier*, call:

``` c
// Provided by VK_EXT_image_drm_format_modifier
VkResult vkGetImageDrmFormatModifierPropertiesEXT(
    VkDevice                                    device,
    VkImage                                     image,
    VkImageDrmFormatModifierPropertiesEXT*      pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the image.

- `image` is the queried image.

- `pProperties` is a pointer to a
  [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierPropertiesEXT.html)
  structure in which properties of the image’s *DRM format modifier* are
  returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-02272"
  id="VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-02272"></a>
  VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-02272  
  `image` **must** have been created with
  <a href="VkImageCreateInfo.html" target="_blank"
  rel="noopener"><code>tiling</code></a> equal to
  `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetImageDrmFormatModifierPropertiesEXT-device-parameter"
  id="VUID-vkGetImageDrmFormatModifierPropertiesEXT-device-parameter"></a>
  VUID-vkGetImageDrmFormatModifierPropertiesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-parameter"
  id="VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-parameter"></a>
  VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a
  href="#VUID-vkGetImageDrmFormatModifierPropertiesEXT-pProperties-parameter"
  id="VUID-vkGetImageDrmFormatModifierPropertiesEXT-pProperties-parameter"></a>
  VUID-vkGetImageDrmFormatModifierPropertiesEXT-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a
  [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierPropertiesEXT.html)
  structure

- <a href="#VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-parent"
  id="VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-parent"></a>
  VUID-vkGetImageDrmFormatModifierPropertiesEXT-image-parent  
  `image` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_drm_format_modifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_drm_format_modifier.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierPropertiesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetImageDrmFormatModifierPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
