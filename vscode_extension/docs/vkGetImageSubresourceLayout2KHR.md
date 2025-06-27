# vkGetImageSubresourceLayout2KHR(3) Manual Page

## Name

vkGetImageSubresourceLayout2KHR - Retrieve information about an image
subresource



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the memory layout of an image subresource, call:

``` c
// Provided by VK_KHR_maintenance5
void vkGetImageSubresourceLayout2KHR(
    VkDevice                                    device,
    VkImage                                     image,
    const VkImageSubresource2KHR*               pSubresource,
    VkSubresourceLayout2KHR*                    pLayout);
```

or the equivalent command

``` c
// Provided by VK_EXT_host_image_copy, VK_EXT_image_compression_control
void vkGetImageSubresourceLayout2EXT(
    VkDevice                                    device,
    VkImage                                     image,
    const VkImageSubresource2KHR*               pSubresource,
    VkSubresourceLayout2KHR*                    pLayout);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the image.

- `image` is the image whose layout is being queried.

- `pSubresource` is a pointer to a
  [VkImageSubresource2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource2KHR.html) structure
  selecting a specific image for the image subresource.

- `pLayout` is a pointer to a
  [VkSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2KHR.html) structure in
  which the layout is returned.

## <a href="#_description" class="anchor"></a>Description

`vkGetImageSubresourceLayout2KHR` behaves similarly to
[vkGetImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout.html), with
the ability to specify extended inputs via chained input structures, and
to return extended information via chained output structures.

It is legal to call `vkGetImageSubresourceLayout2KHR` with an `image`
created with `tiling` equal to `VK_IMAGE_TILING_OPTIMAL`, but the
members of
[VkSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2KHR.html)::`subresourceLayout`
will have undefined values in this case.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Structures chained from <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource2KHR.html">VkImageSubresource2KHR</a>::<code>pNext</code>
will also be updated when <code>tiling</code> is equal to
<code>VK_IMAGE_TILING_OPTIMAL</code>.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-aspectMask-00997"
  id="VUID-vkGetImageSubresourceLayout2KHR-aspectMask-00997"></a>
  VUID-vkGetImageSubresourceLayout2KHR-aspectMask-00997  
  The `aspectMask` member of `pSubresource` **must** only have a single
  bit set

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-mipLevel-01716"
  id="VUID-vkGetImageSubresourceLayout2KHR-mipLevel-01716"></a>
  VUID-vkGetImageSubresourceLayout2KHR-mipLevel-01716  
  The `mipLevel` member of `pSubresource` **must** be less than the
  `mipLevels` specified in `image`

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-arrayLayer-01717"
  id="VUID-vkGetImageSubresourceLayout2KHR-arrayLayer-01717"></a>
  VUID-vkGetImageSubresourceLayout2KHR-arrayLayer-01717  
  The `arrayLayer` member of `pSubresource` **must** be less than the
  `arrayLayers` specified in `image`

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-format-08886"
  id="VUID-vkGetImageSubresourceLayout2KHR-format-08886"></a>
  VUID-vkGetImageSubresourceLayout2KHR-format-08886  
  If `format` of the `image` is a color format that is not a
  [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), and `tiling` of
  the `image` is `VK_IMAGE_TILING_LINEAR` or `VK_IMAGE_TILING_OPTIMAL`,
  the `aspectMask` member of `pSubresource` **must** be
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-format-04462"
  id="VUID-vkGetImageSubresourceLayout2KHR-format-04462"></a>
  VUID-vkGetImageSubresourceLayout2KHR-format-04462  
  If `format` of the `image` has a depth component, the `aspectMask`
  member of `pSubresource` **must** contain `VK_IMAGE_ASPECT_DEPTH_BIT`

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-format-04463"
  id="VUID-vkGetImageSubresourceLayout2KHR-format-04463"></a>
  VUID-vkGetImageSubresourceLayout2KHR-format-04463  
  If `format` of the `image` has a stencil component, the `aspectMask`
  member of `pSubresource` **must** contain
  `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-format-04464"
  id="VUID-vkGetImageSubresourceLayout2KHR-format-04464"></a>
  VUID-vkGetImageSubresourceLayout2KHR-format-04464  
  If `format` of the `image` does not contain a stencil or depth
  component, the `aspectMask` member of `pSubresource` **must** not
  contain `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-tiling-08717"
  id="VUID-vkGetImageSubresourceLayout2KHR-tiling-08717"></a>
  VUID-vkGetImageSubresourceLayout2KHR-tiling-08717  
  If the `tiling` of the `image` is `VK_IMAGE_TILING_LINEAR` and has a
  [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then the
  `aspectMask` member of `pSubresource` **must** be a single valid
  [multi-planar aspect mask](#formats-planes-image-aspect) bit

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-image-09434"
  id="VUID-vkGetImageSubresourceLayout2KHR-image-09434"></a>
  VUID-vkGetImageSubresourceLayout2KHR-image-09434  
  If `image` was created with the
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`
  external memory handle type, then `image` **must** be bound to memory

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-tiling-09435"
  id="VUID-vkGetImageSubresourceLayout2KHR-tiling-09435"></a>
  VUID-vkGetImageSubresourceLayout2KHR-tiling-09435  
  If the `tiling` of the `image` is
  `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then the `aspectMask`
  member of `pSubresource` **must** be
  `VK_IMAGE_ASPECT_MEMORY_PLANE`*`_i_`*`BIT_EXT` and the index *i*
  **must** be less than the
  [VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html)::`drmFormatModifierPlaneCount`
  associated with the image’s `format` and
  [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierPropertiesEXT.html)::`drmFormatModifier`

Valid Usage (Implicit)

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-device-parameter"
  id="VUID-vkGetImageSubresourceLayout2KHR-device-parameter"></a>
  VUID-vkGetImageSubresourceLayout2KHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-image-parameter"
  id="VUID-vkGetImageSubresourceLayout2KHR-image-parameter"></a>
  VUID-vkGetImageSubresourceLayout2KHR-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-pSubresource-parameter"
  id="VUID-vkGetImageSubresourceLayout2KHR-pSubresource-parameter"></a>
  VUID-vkGetImageSubresourceLayout2KHR-pSubresource-parameter  
  `pSubresource` **must** be a valid pointer to a valid
  [VkImageSubresource2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource2KHR.html) structure

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-pLayout-parameter"
  id="VUID-vkGetImageSubresourceLayout2KHR-pLayout-parameter"></a>
  VUID-vkGetImageSubresourceLayout2KHR-pLayout-parameter  
  `pLayout` **must** be a valid pointer to a
  [VkSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2KHR.html) structure

- <a href="#VUID-vkGetImageSubresourceLayout2KHR-image-parent"
  id="VUID-vkGetImageSubresourceLayout2KHR-image-parent"></a>
  VUID-vkGetImageSubresourceLayout2KHR-image-parent  
  `image` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageSubresource2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource2KHR.html),
[VkSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetImageSubresourceLayout2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
