# VkSubresourceLayout(3) Manual Page

## Name

VkSubresourceLayout - Structure specifying subresource layout



## <a href="#_c_specification" class="anchor"></a>C Specification

Information about the layout of the image subresource is returned in a
`VkSubresourceLayout` structure:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkSubresourceLayout {
    VkDeviceSize    offset;
    VkDeviceSize    size;
    VkDeviceSize    rowPitch;
    VkDeviceSize    arrayPitch;
    VkDeviceSize    depthPitch;
} VkSubresourceLayout;
```

## <a href="#_members" class="anchor"></a>Members

- `offset` is the byte offset from the start of the image or the plane
  where the image subresource begins.

- `size` is the size in bytes of the image subresource. `size` includes
  any extra memory that is required based on `rowPitch`.

- `rowPitch` describes the number of bytes between each row of texels in
  an image.

- `arrayPitch` describes the number of bytes between each array layer of
  an image.

- `depthPitch` describes the number of bytes between each slice of 3D
  image.

## <a href="#_description" class="anchor"></a>Description

If the image is <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-linear-resource"
target="_blank" rel="noopener">linear</a>, then `rowPitch`, `arrayPitch`
and `depthPitch` describe the layout of the image subresource in linear
memory. For uncompressed formats, `rowPitch` is the number of bytes
between texels with the same x coordinate in adjacent rows (y
coordinates differ by one). `arrayPitch` is the number of bytes between
texels with the same x and y coordinate in adjacent array layers of the
image (array layer values differ by one). `depthPitch` is the number of
bytes between texels with the same x and y coordinate in adjacent slices
of a 3D image (z coordinates differ by one). Expressed as an addressing
formula, the starting byte of a texel in the image subresource has
address:

``` c
// (x,y,z,layer) are in texel coordinates
address(x,y,z,layer) = layer*arrayPitch + z*depthPitch + y*rowPitch + x*elementSize + offset
```

For compressed formats, the `rowPitch` is the number of bytes between
compressed texel blocks in adjacent rows. `arrayPitch` is the number of
bytes between compressed texel blocks in adjacent array layers.
`depthPitch` is the number of bytes between compressed texel blocks in
adjacent slices of a 3D image.

``` c
// (x,y,z,layer) are in compressed texel block coordinates
address(x,y,z,layer) = layer*arrayPitch + z*depthPitch + y*rowPitch + x*compressedTexelBlockByteSize + offset;
```

The value of `arrayPitch` is undefined for images that were not created
as arrays. `depthPitch` is defined only for 3D images.

If the image has a *single-plane* color format and its tiling is
`VK_IMAGE_TILING_LINEAR` , then the `aspectMask` member of
`VkImageSubresource` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`.

If the image has a depth/stencil format and its tiling is
`VK_IMAGE_TILING_LINEAR` , then `aspectMask` **must** be either
`VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`. On
implementations that store depth and stencil aspects separately,
querying each of these image subresource layouts will return a different
`offset` and `size` representing the region of memory used for that
aspect. On implementations that store depth and stencil aspects
interleaved, the same `offset` and `size` are returned and represent the
interleaved memory allocation.

If the image has a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
target="_blank" rel="noopener">multi-planar format</a> and its tiling is
`VK_IMAGE_TILING_LINEAR` , then the `aspectMask` member of
`VkImageSubresource` **must** be `VK_IMAGE_ASPECT_PLANE_0_BIT`,
`VK_IMAGE_ASPECT_PLANE_1_BIT`, or (for 3-plane formats only)
`VK_IMAGE_ASPECT_PLANE_2_BIT`. Querying each of these image subresource
layouts will return a different `offset` and `size` representing the
region of memory used for that plane. If the image is *disjoint*, then
the `offset` is relative to the base address of the plane. If the image
is *non-disjoint*, then the `offset` is relative to the base address of
the image.

If the image’s tiling is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then
the `aspectMask` member of `VkImageSubresource` **must** be one of
`VK_IMAGE_ASPECT_MEMORY_PLANE`*`_i_`*`BIT_EXT`, where the maximum
allowed plane index *i* is defined by the
[VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html)::`drmFormatModifierPlaneCount`
associated with the image’s
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`format` and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-drm-format-modifier"
target="_blank" rel="noopener">modifier</a>. The memory range used by
the subresource is described by `offset` and `size`. If the image is
*disjoint*, then the `offset` is relative to the base address of the
*memory plane*. If the image is *non-disjoint*, then the `offset` is
relative to the base address of the image. If the image is <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-linear-resource"
target="_blank" rel="noopener">non-linear</a>, then `rowPitch`,
`arrayPitch`, and `depthPitch` have an implementation-dependent meaning.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html),
[VkSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2KHR.html),
[vkGetImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubresourceLayout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
