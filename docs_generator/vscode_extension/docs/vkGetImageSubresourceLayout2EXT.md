# vkGetImageSubresourceLayout2(3) Manual Page

## Name

vkGetImageSubresourceLayout2 - Retrieve information about an image subresource



## [](#_c_specification)C Specification

To query the memory layout of an image subresource, call:

```c++
// Provided by VK_VERSION_1_4
void vkGetImageSubresourceLayout2(
    VkDevice                                    device,
    VkImage                                     image,
    const VkImageSubresource2*                  pSubresource,
    VkSubresourceLayout2*                       pLayout);
```

or the equivalent command

```c++
// Provided by VK_KHR_maintenance5
void vkGetImageSubresourceLayout2KHR(
    VkDevice                                    device,
    VkImage                                     image,
    const VkImageSubresource2*                  pSubresource,
    VkSubresourceLayout2*                       pLayout);
```

or the equivalent command

```c++
// Provided by VK_EXT_host_image_copy, VK_EXT_image_compression_control
void vkGetImageSubresourceLayout2EXT(
    VkDevice                                    device,
    VkImage                                     image,
    const VkImageSubresource2*                  pSubresource,
    VkSubresourceLayout2*                       pLayout);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the image.
- `image` is the image whose layout is being queried.
- `pSubresource` is a pointer to a [VkImageSubresource2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource2.html) structure selecting a specific image for the image subresource.
- `pLayout` is a pointer to a [VkSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2.html) structure in which the layout is returned.

## [](#_description)Description

`vkGetImageSubresourceLayout2` behaves similarly to [vkGetImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout.html), with the ability to specify extended inputs via chained input structures, and to return extended information via chained output structures.

It is legal to call `vkGetImageSubresourceLayout2` with an `image` created with `tiling` equal to `VK_IMAGE_TILING_OPTIMAL`, but the members of [VkSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2.html)::`subresourceLayout` will have undefined values in this case.

Note

Structures chained from [VkImageSubresource2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource2.html)::`pNext` will also be updated when `tiling` is equal to `VK_IMAGE_TILING_OPTIMAL`.

Valid Usage

- [](#VUID-vkGetImageSubresourceLayout2-aspectMask-00997)VUID-vkGetImageSubresourceLayout2-aspectMask-00997  
  The `aspectMask` member of `pSubresource` **must** only have a single bit set
- [](#VUID-vkGetImageSubresourceLayout2-mipLevel-01716)VUID-vkGetImageSubresourceLayout2-mipLevel-01716  
  The `mipLevel` member of `pSubresource` **must** be less than the `mipLevels` specified in `image`
- [](#VUID-vkGetImageSubresourceLayout2-arrayLayer-01717)VUID-vkGetImageSubresourceLayout2-arrayLayer-01717  
  The `arrayLayer` member of `pSubresource` **must** be less than the `arrayLayers` specified in `image`
- [](#VUID-vkGetImageSubresourceLayout2-format-08886)VUID-vkGetImageSubresourceLayout2-format-08886  
  If `format` of the `image` is a color format that is not a [multi-planar format](#formats-multiplanar), and `tiling` of the `image` is `VK_IMAGE_TILING_LINEAR` or `VK_IMAGE_TILING_OPTIMAL`, the `aspectMask` member of `pSubresource` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-vkGetImageSubresourceLayout2-format-04462)VUID-vkGetImageSubresourceLayout2-format-04462  
  If `format` of the `image` has a depth component, the `aspectMask` member of `pSubresource` **must** contain `VK_IMAGE_ASPECT_DEPTH_BIT`
- [](#VUID-vkGetImageSubresourceLayout2-format-04463)VUID-vkGetImageSubresourceLayout2-format-04463  
  If `format` of the `image` has a stencil component, the `aspectMask` member of `pSubresource` **must** contain `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-vkGetImageSubresourceLayout2-format-04464)VUID-vkGetImageSubresourceLayout2-format-04464  
  If `format` of the `image` does not contain a stencil or depth component, the `aspectMask` member of `pSubresource` **must** not contain `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-vkGetImageSubresourceLayout2-tiling-08717)VUID-vkGetImageSubresourceLayout2-tiling-08717  
  If the `tiling` of the `image` is `VK_IMAGE_TILING_LINEAR` and has a [multi-planar format](#formats-multiplanar), then the `aspectMask` member of `pSubresource` **must** be a single valid [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit
- [](#VUID-vkGetImageSubresourceLayout2-image-09434)VUID-vkGetImageSubresourceLayout2-image-09434  
  If `image` was created with the `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID` external memory handle type, then `image` **must** be bound to memory
- [](#VUID-vkGetImageSubresourceLayout2-tiling-09435)VUID-vkGetImageSubresourceLayout2-tiling-09435  
  If the `tiling` of the `image` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then the `aspectMask` member of `pSubresource` **must** be `VK_IMAGE_ASPECT_MEMORY_PLANE_i_BIT_EXT` and the index *i* **must** be less than the [VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrmFormatModifierPropertiesEXT.html)::`drmFormatModifierPlaneCount` associated with the image’s `format` and [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierPropertiesEXT.html)::`drmFormatModifier`

Valid Usage (Implicit)

- [](#VUID-vkGetImageSubresourceLayout2-device-parameter)VUID-vkGetImageSubresourceLayout2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetImageSubresourceLayout2-image-parameter)VUID-vkGetImageSubresourceLayout2-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-vkGetImageSubresourceLayout2-pSubresource-parameter)VUID-vkGetImageSubresourceLayout2-pSubresource-parameter  
  `pSubresource` **must** be a valid pointer to a valid [VkImageSubresource2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource2.html) structure
- [](#VUID-vkGetImageSubresourceLayout2-pLayout-parameter)VUID-vkGetImageSubresourceLayout2-pLayout-parameter  
  `pLayout` **must** be a valid pointer to a [VkSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2.html) structure
- [](#VUID-vkGetImageSubresourceLayout2-image-parent)VUID-vkGetImageSubresourceLayout2-image-parent  
  `image` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_EXT\_image\_compression\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_compression_control.html), [VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageSubresource2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource2.html), [VkSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetImageSubresourceLayout2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0