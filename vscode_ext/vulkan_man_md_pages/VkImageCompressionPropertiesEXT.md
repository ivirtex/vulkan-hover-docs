# VkImageCompressionPropertiesEXT(3) Manual Page

## Name

VkImageCompressionPropertiesEXT - Compression properties of an image



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the compression properties of an image, add a
[VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionPropertiesEXT.html)
structure to the `pNext` chain of the
[VkSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout2EXT.html) structure in a
call to
[vkGetImageSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout2KHR.html)
or
[vkGetImageSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout2EXT.html).

To determine the compression rates that are supported for a given image
format, add a
[VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionPropertiesEXT.html)
structure to the `pNext` chain of the
[VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html) structure in a
call to
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html).

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Since fixed-rate compression is disabled by default, the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionPropertiesEXT.html">VkImageCompressionPropertiesEXT</a>
structure passed to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html">vkGetPhysicalDeviceImageFormatProperties2</a>
will not indicate any fixed-rate compression support unless a <a
href="VkImageCompressionControlEXT.html">VkImageCompressionControlEXT</a>
structure is also included in the <code>pNext</code> chain of the <a
href="VkPhysicalDeviceImageFormatInfo2.html">VkPhysicalDeviceImageFormatInfo2</a>
structure passed to the same command.</p></td>
</tr>
</tbody>
</table>

The `VkImageCompressionPropertiesEXT` structure is defined as:

``` c
// Provided by VK_EXT_image_compression_control
typedef struct VkImageCompressionPropertiesEXT {
    VkStructureType                        sType;
    void*                                  pNext;
    VkImageCompressionFlagsEXT             imageCompressionFlags;
    VkImageCompressionFixedRateFlagsEXT    imageCompressionFixedRateFlags;
} VkImageCompressionPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `imageCompressionFlags` returns a value describing the compression
  controls that apply to the image. The value will be either
  `VK_IMAGE_COMPRESSION_DEFAULT_EXT` to indicate no fixed-rate
  compression, `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT` to
  indicate fixed-rate compression, or
  `VK_IMAGE_COMPRESSION_DISABLED_EXT` to indicate no compression.

- `imageCompressionFixedRateFlags` returns a
  [VkImageCompressionFixedRateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFixedRateFlagsEXT.html)
  value describing the compression rates that apply to the specified
  aspect of the image.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkImageCompressionPropertiesEXT-sType-sType"
  id="VUID-VkImageCompressionPropertiesEXT-sType-sType"></a>
  VUID-VkImageCompressionPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_COMPRESSION_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_compression_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_compression_control.html),
[VkImageCompressionFixedRateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFixedRateFlagsEXT.html),
[VkImageCompressionFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageCompressionPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
