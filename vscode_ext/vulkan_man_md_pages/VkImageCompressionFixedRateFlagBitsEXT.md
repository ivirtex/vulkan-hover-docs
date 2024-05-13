# VkImageCompressionFixedRateFlagBitsEXT(3) Manual Page

## Name

VkImageCompressionFixedRateFlagBitsEXT - Bitmask specifying fixed rate
image compression rates



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)::`pFixedRateFlags`,
specifying allowed compression rates for an image plane, are:

``` c
// Provided by VK_EXT_image_compression_control
typedef enum VkImageCompressionFixedRateFlagBitsEXT {
    VK_IMAGE_COMPRESSION_FIXED_RATE_NONE_EXT = 0,
    VK_IMAGE_COMPRESSION_FIXED_RATE_1BPC_BIT_EXT = 0x00000001,
    VK_IMAGE_COMPRESSION_FIXED_RATE_2BPC_BIT_EXT = 0x00000002,
    VK_IMAGE_COMPRESSION_FIXED_RATE_3BPC_BIT_EXT = 0x00000004,
    VK_IMAGE_COMPRESSION_FIXED_RATE_4BPC_BIT_EXT = 0x00000008,
    VK_IMAGE_COMPRESSION_FIXED_RATE_5BPC_BIT_EXT = 0x00000010,
    VK_IMAGE_COMPRESSION_FIXED_RATE_6BPC_BIT_EXT = 0x00000020,
    VK_IMAGE_COMPRESSION_FIXED_RATE_7BPC_BIT_EXT = 0x00000040,
    VK_IMAGE_COMPRESSION_FIXED_RATE_8BPC_BIT_EXT = 0x00000080,
    VK_IMAGE_COMPRESSION_FIXED_RATE_9BPC_BIT_EXT = 0x00000100,
    VK_IMAGE_COMPRESSION_FIXED_RATE_10BPC_BIT_EXT = 0x00000200,
    VK_IMAGE_COMPRESSION_FIXED_RATE_11BPC_BIT_EXT = 0x00000400,
    VK_IMAGE_COMPRESSION_FIXED_RATE_12BPC_BIT_EXT = 0x00000800,
    VK_IMAGE_COMPRESSION_FIXED_RATE_13BPC_BIT_EXT = 0x00001000,
    VK_IMAGE_COMPRESSION_FIXED_RATE_14BPC_BIT_EXT = 0x00002000,
    VK_IMAGE_COMPRESSION_FIXED_RATE_15BPC_BIT_EXT = 0x00004000,
    VK_IMAGE_COMPRESSION_FIXED_RATE_16BPC_BIT_EXT = 0x00008000,
    VK_IMAGE_COMPRESSION_FIXED_RATE_17BPC_BIT_EXT = 0x00010000,
    VK_IMAGE_COMPRESSION_FIXED_RATE_18BPC_BIT_EXT = 0x00020000,
    VK_IMAGE_COMPRESSION_FIXED_RATE_19BPC_BIT_EXT = 0x00040000,
    VK_IMAGE_COMPRESSION_FIXED_RATE_20BPC_BIT_EXT = 0x00080000,
    VK_IMAGE_COMPRESSION_FIXED_RATE_21BPC_BIT_EXT = 0x00100000,
    VK_IMAGE_COMPRESSION_FIXED_RATE_22BPC_BIT_EXT = 0x00200000,
    VK_IMAGE_COMPRESSION_FIXED_RATE_23BPC_BIT_EXT = 0x00400000,
    VK_IMAGE_COMPRESSION_FIXED_RATE_24BPC_BIT_EXT = 0x00800000,
} VkImageCompressionFixedRateFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_IMAGE_COMPRESSION_FIXED_RATE_NONE_EXT` specifies that fixed-rate
  compression **must** not be used.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_1BPC_BIT_EXT` specifies that
  fixed-rate compression with a bitrate of \[1,2) bits per component
  **may** be used.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_2BPC_BIT_EXT` specifies that
  fixed-rate compression with a bitrate of \[2,3) bits per component
  **may** be used.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_3BPC_BIT_EXT` specifies that
  fixed-rate compression with a bitrate of \[3,4) bits per component
  **may** be used.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_4BPC_BIT_EXT` specifies that
  fixed-rate compression with a bitrate of \[4,5) bits per component
  **may** be used.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_5BPC_BIT_EXT` specifies that
  fixed-rate compression with a bitrate of \[5,6) bits per component
  **may** be used.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_6BPC_BIT_EXT` specifies that
  fixed-rate compression with a bitrate of \[6,7) bits per component
  **may** be used.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_7BPC_BIT_EXT` specifies that
  fixed-rate compression with a bitrate of \[7,8) bits per component
  **may** be used.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_8BPC_BIT_EXT` specifies that
  fixed-rate compression with a bitrate of \[8,9) bits per component
  **may** be used.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_9BPC_BIT_EXT` specifies that
  fixed-rate compression with a bitrate of \[9,10) bits per component
  **may** be used.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_10BPC_BIT_EXT` specifies that
  fixed-rate compression with a bitrate of \[10,11) bits per component
  **may** be used.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_11BPC_BIT_EXT` specifies that
  fixed-rate compression with a bitrate of \[11,12) bits per component
  **may** be used.

- `VK_IMAGE_COMPRESSION_FIXED_RATE_12BPC_BIT_EXT` specifies that
  fixed-rate compression with a bitrate of at least 12 bits per
  component **may** be used.

If the format has a different bit rate for different components,
[VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)::`pFixedRateFlags`
describes the rate of the component with the largest number of bits
assigned to it, scaled pro rata. For example, to request that a
`VK_FORMAT_A2R10G10B10_UNORM_PACK32` format be stored at a rate of 8
bits per pixel, use `VK_IMAGE_COMPRESSION_FIXED_RATE_2BPC_BIT_EXT` (10
bits for the largest component, stored at quarter the original size, 2.5
bits, rounded down).

If `flags` includes `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT`, and
multiple bits are set in
[VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html)::`pFixedRateFlags`
for a plane, implementations **should** apply the lowest allowed bitrate
that is supported.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The choice of “bits per component” terminology was chosen so that the
same compression rate describes the same degree of compression applied
to formats that differ only in the number of components. For example,
<code>VK_FORMAT_R8G8_UNORM</code> compressed to half its original size
is a rate of 4 bits per component, 8 bits per pixel.
<code>VK_FORMAT_R8G8B8A8_UNORM</code> compressed to half <em>its</em>
original size is 4 bits per component, 16 bits per pixel. Both of these
cases can be requested with
<code>VK_IMAGE_COMPRESSION_FIXED_RATE_4BPC_BIT_EXT</code>.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_compression_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_compression_control.html),
[VkImageCompressionFixedRateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFixedRateFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageCompressionFixedRateFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
