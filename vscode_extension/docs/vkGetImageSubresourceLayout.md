# vkGetImageSubresourceLayout(3) Manual Page

## Name

vkGetImageSubresourceLayout - Retrieve information about an image
subresource



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the memory layout of an image subresource, call:

``` c
// Provided by VK_VERSION_1_0
void vkGetImageSubresourceLayout(
    VkDevice                                    device,
    VkImage                                     image,
    const VkImageSubresource*                   pSubresource,
    VkSubresourceLayout*                        pLayout);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the image.

- `image` is the image whose layout is being queried.

- `pSubresource` is a pointer to a
  [VkImageSubresource](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource.html) structure selecting a
  specific image subresource from the image.

- `pLayout` is a pointer to a
  [VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout.html) structure in which the
  layout is returned.

## <a href="#_description" class="anchor"></a>Description

If the image is <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-linear-resource"
target="_blank" rel="noopener">linear</a>, then the returned layout is
valid for <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-device-hostaccess"
target="_blank" rel="noopener">host access</a>.

If the image’s tiling is `VK_IMAGE_TILING_LINEAR` and its format is a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
target="_blank" rel="noopener">multi-planar format</a>, then
`vkGetImageSubresourceLayout` describes one *format plane* of the image.
If the image’s tiling is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then
`vkGetImageSubresourceLayout` describes one *memory plane* of the image.
If the image’s tiling is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT` and
the image is <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-linear-resource"
target="_blank" rel="noopener">non-linear</a>, then the returned layout
has an implementation-dependent meaning; the vendor of the image’s <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-drm-format-modifier"
target="_blank" rel="noopener">DRM format modifier</a> **may** provide
documentation that explains how to interpret the returned layout.

`vkGetImageSubresourceLayout` is invariant for the lifetime of a single
image. However, the subresource layout of images in Android hardware
buffer or QNX Screen buffer external memory is not known until the image
has been bound to memory, so applications **must** not call
[vkGetImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout.html) for such
an image before it has been bound.

Valid Usage

- <a href="#VUID-vkGetImageSubresourceLayout-image-07790"
  id="VUID-vkGetImageSubresourceLayout-image-07790"></a>
  VUID-vkGetImageSubresourceLayout-image-07790  
  `image` **must** have been created with `tiling` equal to
  `VK_IMAGE_TILING_LINEAR` or `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`

<!-- -->

- <a href="#VUID-vkGetImageSubresourceLayout-aspectMask-00997"
  id="VUID-vkGetImageSubresourceLayout-aspectMask-00997"></a>
  VUID-vkGetImageSubresourceLayout-aspectMask-00997  
  The `aspectMask` member of `pSubresource` **must** only have a single
  bit set

- <a href="#VUID-vkGetImageSubresourceLayout-mipLevel-01716"
  id="VUID-vkGetImageSubresourceLayout-mipLevel-01716"></a>
  VUID-vkGetImageSubresourceLayout-mipLevel-01716  
  The `mipLevel` member of `pSubresource` **must** be less than the
  `mipLevels` specified in `image`

- <a href="#VUID-vkGetImageSubresourceLayout-arrayLayer-01717"
  id="VUID-vkGetImageSubresourceLayout-arrayLayer-01717"></a>
  VUID-vkGetImageSubresourceLayout-arrayLayer-01717  
  The `arrayLayer` member of `pSubresource` **must** be less than the
  `arrayLayers` specified in `image`

- <a href="#VUID-vkGetImageSubresourceLayout-format-08886"
  id="VUID-vkGetImageSubresourceLayout-format-08886"></a>
  VUID-vkGetImageSubresourceLayout-format-08886  
  If `format` of the `image` is a color format that is not a
  [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), and `tiling` of
  the `image` is `VK_IMAGE_TILING_LINEAR` or `VK_IMAGE_TILING_OPTIMAL`,
  the `aspectMask` member of `pSubresource` **must** be
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-vkGetImageSubresourceLayout-format-04462"
  id="VUID-vkGetImageSubresourceLayout-format-04462"></a>
  VUID-vkGetImageSubresourceLayout-format-04462  
  If `format` of the `image` has a depth component, the `aspectMask`
  member of `pSubresource` **must** contain `VK_IMAGE_ASPECT_DEPTH_BIT`

- <a href="#VUID-vkGetImageSubresourceLayout-format-04463"
  id="VUID-vkGetImageSubresourceLayout-format-04463"></a>
  VUID-vkGetImageSubresourceLayout-format-04463  
  If `format` of the `image` has a stencil component, the `aspectMask`
  member of `pSubresource` **must** contain
  `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-vkGetImageSubresourceLayout-format-04464"
  id="VUID-vkGetImageSubresourceLayout-format-04464"></a>
  VUID-vkGetImageSubresourceLayout-format-04464  
  If `format` of the `image` does not contain a stencil or depth
  component, the `aspectMask` member of `pSubresource` **must** not
  contain `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-vkGetImageSubresourceLayout-tiling-08717"
  id="VUID-vkGetImageSubresourceLayout-tiling-08717"></a>
  VUID-vkGetImageSubresourceLayout-tiling-08717  
  If the `tiling` of the `image` is `VK_IMAGE_TILING_LINEAR` and has a
  [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then the
  `aspectMask` member of `pSubresource` **must** be a single valid
  [multi-planar aspect mask](#formats-planes-image-aspect) bit

- <a href="#VUID-vkGetImageSubresourceLayout-image-09432"
  id="VUID-vkGetImageSubresourceLayout-image-09432"></a>
  VUID-vkGetImageSubresourceLayout-image-09432  
  If `image` was created with the
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`
  external memory handle type, then `image` **must** be bound to memory

- <a href="#VUID-vkGetImageSubresourceLayout-tiling-09433"
  id="VUID-vkGetImageSubresourceLayout-tiling-09433"></a>
  VUID-vkGetImageSubresourceLayout-tiling-09433  
  If the `tiling` of the `image` is
  `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then the `aspectMask`
  member of `pSubresource` **must** be
  `VK_IMAGE_ASPECT_MEMORY_PLANE`*`_i_`*`BIT_EXT` and the index *i*
  **must** be less than the
  [VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html)::`drmFormatModifierPlaneCount`
  associated with the image’s `format` and
  [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierPropertiesEXT.html)::`drmFormatModifier`

Valid Usage (Implicit)

- <a href="#VUID-vkGetImageSubresourceLayout-device-parameter"
  id="VUID-vkGetImageSubresourceLayout-device-parameter"></a>
  VUID-vkGetImageSubresourceLayout-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetImageSubresourceLayout-image-parameter"
  id="VUID-vkGetImageSubresourceLayout-image-parameter"></a>
  VUID-vkGetImageSubresourceLayout-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkGetImageSubresourceLayout-pSubresource-parameter"
  id="VUID-vkGetImageSubresourceLayout-pSubresource-parameter"></a>
  VUID-vkGetImageSubresourceLayout-pSubresource-parameter  
  `pSubresource` **must** be a valid pointer to a valid
  [VkImageSubresource](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource.html) structure

- <a href="#VUID-vkGetImageSubresourceLayout-pLayout-parameter"
  id="VUID-vkGetImageSubresourceLayout-pLayout-parameter"></a>
  VUID-vkGetImageSubresourceLayout-pLayout-parameter  
  `pLayout` **must** be a valid pointer to a
  [VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout.html) structure

- <a href="#VUID-vkGetImageSubresourceLayout-image-parent"
  id="VUID-vkGetImageSubresourceLayout-image-parent"></a>
  VUID-vkGetImageSubresourceLayout-image-parent  
  `image` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkImageSubresource](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource.html),
[VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetImageSubresourceLayout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
