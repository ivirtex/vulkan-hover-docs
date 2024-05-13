# VkImageCompressionFlagBitsEXT(3) Manual Page

## Name

VkImageCompressionFlagBitsEXT - Bitmask specifying image compression
controls



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)::`flags`,
specifying compression controls for an image, are:

``` c
// Provided by VK_EXT_image_compression_control
typedef enum VkImageCompressionFlagBitsEXT {
    VK_IMAGE_COMPRESSION_DEFAULT_EXT = 0,
    VK_IMAGE_COMPRESSION_FIXED_RATE_DEFAULT_EXT = 0x00000001,
    VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT = 0x00000002,
    VK_IMAGE_COMPRESSION_DISABLED_EXT = 0x00000004,
} VkImageCompressionFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_IMAGE_COMPRESSION_DEFAULT_EXT` specifies that the default image
  compression setting is used. Implementations **must** not apply
  fixed-rate compression.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_DEFAULT_EXT` specifies that the
  implementation **may** choose any supported fixed-rate compression
  setting in an implementation-defined manner based on the properties of
  the image.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT` specifies that
  fixed-rate compression **may** be used and that the allowed
  compression rates are specified by
  [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)::`pFixedRateFlags`.

- `VK_IMAGE_COMPRESSION_DISABLED_EXT` specifies that all lossless and
  fixed-rate compression **should** be disabled.

If
[VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)::`flags`
is `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT`, then the
`i`<sup>th</sup> member of the `pFixedRateFlags` array specifies the
allowed compression rates for the image’s `i`<sup>th</sup> plane.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>If <code>VK_IMAGE_COMPRESSION_DISABLED_EXT</code> is included in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html">VkImageCompressionControlEXT</a>::<code>flags</code>,
both lossless and fixed-rate compression will be disabled. This is
likely to have a negative impact on performance and is only intended to
be used for debugging purposes.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_compression_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_compression_control.html),
[VkImageCompressionFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageCompressionFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
