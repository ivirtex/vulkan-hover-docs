# VkImageCompressionControlEXT(3) Manual Page

## Name

VkImageCompressionControlEXT - Specify image compression properties



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` list of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
includes a `VkImageCompressionControlEXT` structure, then that structure
describes compression controls for this image.

The `VkImageCompressionControlEXT` structure is defined as:

``` c
// Provided by VK_EXT_image_compression_control
typedef struct VkImageCompressionControlEXT {
    VkStructureType                         sType;
    const void*                             pNext;
    VkImageCompressionFlagsEXT              flags;
    uint32_t                                compressionControlPlaneCount;
    VkImageCompressionFixedRateFlagsEXT*    pFixedRateFlags;
} VkImageCompressionControlEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkImageCompressionFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFlagBitsEXT.html)
  describing compression controls for the image.

- `compressionControlPlaneCount` is the number of entries in the
  `pFixedRateFlags` array.

- `pFixedRateFlags` is `NULL` or a pointer to an array of
  [VkImageCompressionFixedRateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFixedRateFlagsEXT.html)
  bitfields describing allowed fixed-rate compression rates of each
  image plane. It is ignored if `flags` does not include
  `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT`.

## <a href="#_description" class="anchor"></a>Description

If enabled, fixed-rate compression is done in an implementation-defined
manner and **may** be applied at block granularity. In that case, a
write to an individual texel **may** modify the value of other texels in
the same block.

Valid Usage

- <a href="#VUID-VkImageCompressionControlEXT-flags-06747"
  id="VUID-VkImageCompressionControlEXT-flags-06747"></a>
  VUID-VkImageCompressionControlEXT-flags-06747  
  `flags` **must** be one of `VK_IMAGE_COMPRESSION_DEFAULT_EXT`,
  `VK_IMAGE_COMPRESSION_FIXED_RATE_DEFAULT_EXT`,
  `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT`, or
  `VK_IMAGE_COMPRESSION_DISABLED_EXT`

- <a href="#VUID-VkImageCompressionControlEXT-flags-06748"
  id="VUID-VkImageCompressionControlEXT-flags-06748"></a>
  VUID-VkImageCompressionControlEXT-flags-06748  
  If `flags` includes `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT`,
  `pFixedRateFlags` **must** not be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-VkImageCompressionControlEXT-sType-sType"
  id="VUID-VkImageCompressionControlEXT-sType-sType"></a>
  VUID-VkImageCompressionControlEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_COMPRESSION_CONTROL_EXT`

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Some combinations of compression properties may not be supported. For
example, some implementations may not support different fixed-rate
compression rates per plane of a multi-planar format and will not be
able to enable fixed-rate compression for any plane if the requested
rates differ.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_compression_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_compression_control.html),
[VkImageCompressionFixedRateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFixedRateFlagsEXT.html),
[VkImageCompressionFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageCompressionControlEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
