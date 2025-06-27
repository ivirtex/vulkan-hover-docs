# VkDrmFormatModifierPropertiesEXT(3) Manual Page

## Name

VkDrmFormatModifierPropertiesEXT - Structure specifying properties of a
format when combined with a DRM format modifier



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html)
structure describes properties of a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) when that
format is combined with a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-drm-format-modifier"
target="_blank" rel="noopener">Linux DRM format modifier</a>. These
properties, like those of
[VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html), are independent of any
particular image.

The
[VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html)
structure is defined as:

``` c
// Provided by VK_EXT_image_drm_format_modifier
typedef struct VkDrmFormatModifierPropertiesEXT {
    uint64_t                drmFormatModifier;
    uint32_t                drmFormatModifierPlaneCount;
    VkFormatFeatureFlags    drmFormatModifierTilingFeatures;
} VkDrmFormatModifierPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `drmFormatModifier` is a *Linux DRM format modifier*.

- `drmFormatModifierPlaneCount` is the number of *memory planes* in any
  image created with `format` and `drmFormatModifier`. An image’s
  *memory planecount* is distinct from its *format planecount*, as
  explained below.

- `drmFormatModifierTilingFeatures` is a bitmask of
  [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html) that are
  supported by any image created with `format` and `drmFormatModifier`.

## <a href="#_description" class="anchor"></a>Description

The returned `drmFormatModifierTilingFeatures` **must** contain at least
one bit.

The implementation **must** not return `DRM_FORMAT_MOD_INVALID` in
`drmFormatModifier`.

An image’s *memory planecount* (as returned by
`drmFormatModifierPlaneCount`) is distinct from its *format planecount*
(in the sense of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
target="_blank" rel="noopener">multi-planar</a>
Y′C<sub>B</sub>C<sub>R</sub> formats). In
[VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlags.html), each
`VK_IMAGE_ASPECT_MEMORY_PLANE`*`_i_`*`BIT_EXT` represents a *memory
plane* and each `VK_IMAGE_ASPECT_PLANE`*`_i_`*`BIT` a *format plane*.

An image’s set of *format planes* is an ordered partition of the image’s
**content** into separable groups of format components. The ordered
partition is encoded in the name of each [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html). For
example, `VK_FORMAT_G8_B8R8_2PLANE_420_UNORM` contains two *format
planes*; the first plane contains the green component and the second
plane contains the blue component and red component. If the format name
does not contain `PLANE`, then the format contains a single plane; for
example, `VK_FORMAT_R8G8B8A8_UNORM`. Some commands, such as
[vkCmdCopyBufferToImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage.html), do not operate on
all format components in the image, but instead operate only on the
*format planes* explicitly chosen by the application and operate on each
*format plane* independently.

An image’s set of *memory planes* is an ordered partition of the image’s
**memory** rather than the image’s **content**. Each *memory plane* is a
contiguous range of memory. The union of an image’s *memory planes* is
not necessarily contiguous.

If an image is <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-linear-resource"
target="_blank" rel="noopener">linear</a>, then the partition is the
same for *memory planes* and for *format planes*. Therefore, if the
returned `drmFormatModifier` is `DRM_FORMAT_MOD_LINEAR`, then
`drmFormatModifierPlaneCount` **must** equal the *format planecount*,
and `drmFormatModifierTilingFeatures` **must** be identical to the
[VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html)::`linearTilingFeatures`
returned in the same `pNext` chain.

If an image is <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-linear-resource"
target="_blank" rel="noopener">non-linear</a>, then the partition of the
image’s **memory** into *memory planes* is implementation-specific and
**may** be unrelated to the partition of the image’s **content** into
*format planes*. For example, consider an image whose `format` is
`VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM`, `tiling` is
`VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, whose `drmFormatModifier` is
not `DRM_FORMAT_MOD_LINEAR`, and `flags` lacks
`VK_IMAGE_CREATE_DISJOINT_BIT`. The image has 3 *format planes*, and
commands such [vkCmdCopyBufferToImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage.html) act
on each *format plane* independently as if the data of each *format
plane* were separable from the data of the other planes. In a
straightforward implementation, the implementation **may** store the
image’s content in 3 adjacent *memory planes* where each *memory plane*
corresponds exactly to a *format plane*. However, the implementation
**may** also store the image’s content in a single *memory plane* where
all format components are combined using an implementation-private
block-compressed format; or the implementation **may** store the image’s
content in a collection of 7 adjacent *memory planes* using an
implementation-private sharding technique. Because the image is
non-linear and non-disjoint, the implementation has much freedom when
choosing the image’s placement in memory.

The *memory planecount* applies to function parameters and structures
only when the API specifies an explicit requirement on
`drmFormatModifierPlaneCount`. In all other cases, the *memory
planecount* is ignored.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_drm_format_modifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_drm_format_modifier.html),
[VkDrmFormatModifierPropertiesListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesListEXT.html),
[VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDrmFormatModifierPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
