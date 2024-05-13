# VkImageFormatProperties(3) Manual Page

## Name

VkImageFormatProperties - Structure specifying an image format
properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageFormatProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkImageFormatProperties {
    VkExtent3D            maxExtent;
    uint32_t              maxMipLevels;
    uint32_t              maxArrayLayers;
    VkSampleCountFlags    sampleCounts;
    VkDeviceSize          maxResourceSize;
} VkImageFormatProperties;
```

## <a href="#_members" class="anchor"></a>Members

- `maxExtent` are the maximum image dimensions. See the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extentperimagetype"
  target="_blank" rel="noopener">Allowed Extent Values</a> section below
  for how these values are constrained by `type`.

- `maxMipLevels` is the maximum number of mipmap levels. `maxMipLevels`
  **must** be equal to the number of levels in the complete mipmap chain
  based on the `maxExtent.width`, `maxExtent.height`, and
  `maxExtent.depth`, except when one of the following conditions is
  true, in which case it **may** instead be `1`:

  - `vkGetPhysicalDeviceImageFormatProperties`::`tiling` was
    `VK_IMAGE_TILING_LINEAR`

  - [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`tiling`
    was `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`

  - the
    [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`pNext`
    chain included a
    [VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalImageFormatInfo.html)
    structure with a handle type included in the `handleTypes` member
    for which mipmap image support is not required

  - image `format` is one of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
    target="_blank" rel="noopener">formats that require a sampler
    Y′C<sub>B</sub>C<sub>R</sub> conversion</a>

  - `flags` contains `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

- `maxArrayLayers` is the maximum number of array layers.
  `maxArrayLayers` **must** be no less than
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxImageArrayLayers`,
  except when one of the following conditions is true, in which case it
  **may** instead be `1`:

  - `tiling` is `VK_IMAGE_TILING_LINEAR`

  - `tiling` is `VK_IMAGE_TILING_OPTIMAL` and `type` is
    `VK_IMAGE_TYPE_3D`

  - `format` is one of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
    target="_blank" rel="noopener">formats that require a sampler
    Y′C<sub>B</sub>C<sub>R</sub> conversion</a>

- If `tiling` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then
  `maxArrayLayers` **must** not be 0.

- `sampleCounts` is a bitmask of
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) specifying all the
  supported sample counts for this image as described <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-supported-sample-counts"
  target="_blank" rel="noopener">below</a>.

- `maxResourceSize` is an upper bound on the total image size in bytes,
  inclusive of all image subresources. Implementations **may** have an
  address space limit on total size of a resource, which is advertised
  by this property. `maxResourceSize` **must** be at least
  2<sup>31</sup>.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>There is no mechanism to query the size of an image before creating
it, to compare that size against <code>maxResourceSize</code>. If an
application attempts to create an image that exceeds this limit, the
creation will fail and <a href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html">vkCreateImage</a>
will return <code>VK_ERROR_OUT_OF_DEVICE_MEMORY</code>. While the
advertised limit <strong>must</strong> be at least 2<sup>31</sup>, it
<strong>may</strong> not be possible to create an image that approaches
that size, particularly for <code>VK_IMAGE_TYPE_1D</code>.</p></td>
</tr>
</tbody>
</table>

If the combination of parameters to
`vkGetPhysicalDeviceImageFormatProperties` is not supported by the
implementation for use in [vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html), then all
members of `VkImageFormatProperties` will be filled with zero.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Filling <code>VkImageFormatProperties</code> with zero for
unsupported formats is an exception to the usual rule that output
structures have undefined contents on error. This exception was
unintentional, but is preserved for backwards compatibility.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html),
[VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatPropertiesNV.html),
[VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html),
[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlags.html),
[vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageFormatProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
