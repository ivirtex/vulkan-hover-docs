# VkHdrMetadataEXT(3) Manual Page

## Name

VkHdrMetadataEXT - Specify HDR metadata



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkHdrMetadataEXT` structure is defined as:

``` c
// Provided by VK_EXT_hdr_metadata
typedef struct VkHdrMetadataEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkXYColorEXT       displayPrimaryRed;
    VkXYColorEXT       displayPrimaryGreen;
    VkXYColorEXT       displayPrimaryBlue;
    VkXYColorEXT       whitePoint;
    float              maxLuminance;
    float              minLuminance;
    float              maxContentLightLevel;
    float              maxFrameAverageLightLevel;
} VkHdrMetadataEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `displayPrimaryRed` is a [VkXYColorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXYColorEXT.html) structure
  specifying the red primary of the display used to optimize the content

- `displayPrimaryGreen` is a [VkXYColorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXYColorEXT.html) structure
  specifying the green primary of the display used to optimize the
  content

- `displayPrimaryBlue` is a [VkXYColorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXYColorEXT.html) structure
  specifying the blue primary of the display used to optimize the
  content

- `whitePoint` is a [VkXYColorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXYColorEXT.html) structure
  specifying the white-point of the display used to optimize the content

- `maxLuminance` is the maximum luminance of the display used to
  optimize the content in nits

- `minLuminance` is the minimum luminance of the display used to
  optimize the content in nits

- `maxContentLightLevel` is the value in nits of the desired luminance
  for the brightest pixels in the displayed image.

- `maxFrameAverageLightLevel` is the value in nits of the average
  luminance of the frame which has the brightest average luminance
  anywhere in the content.

## <a href="#_description" class="anchor"></a>Description

If any of the above values are unknown, they **can** be set to 0.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The meta-data provided here is intended to be used as defined in the
SMPTE 2086, CTA 861.3 and CIE 15:2004 specifications. The validity and
use of this data is outside the scope of Vulkan.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkHdrMetadataEXT-sType-sType"
  id="VUID-VkHdrMetadataEXT-sType-sType"></a>
  VUID-VkHdrMetadataEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_HDR_METADATA_EXT`

- <a href="#VUID-VkHdrMetadataEXT-pNext-pNext"
  id="VUID-VkHdrMetadataEXT-pNext-pNext"></a>
  VUID-VkHdrMetadataEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_hdr_metadata](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_hdr_metadata.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkXYColorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXYColorEXT.html),
[vkSetHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetHdrMetadataEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkHdrMetadataEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
