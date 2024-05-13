# VkHdrMetadataEXT(3) Manual Page

## Name

VkHdrMetadataEXT - Specify Hdr metadata



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
  specifying the reference monitor’s red primary in chromaticity
  coordinates

- `displayPrimaryGreen` is a [VkXYColorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXYColorEXT.html) structure
  specifying the reference monitor’s green primary in chromaticity
  coordinates

- `displayPrimaryBlue` is a [VkXYColorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXYColorEXT.html) structure
  specifying the reference monitor’s blue primary in chromaticity
  coordinates

- `whitePoint` is a [VkXYColorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXYColorEXT.html) structure
  specifying the reference monitor’s white-point in chromaticity
  coordinates

- `maxLuminance` is the maximum luminance of the reference monitor in
  nits

- `minLuminance` is the minimum luminance of the reference monitor in
  nits

- `maxContentLightLevel` is content’s maximum luminance in nits

- `maxFrameAverageLightLevel` is the maximum frame average light level
  in nits

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkHdrMetadataEXT-sType-sType"
  id="VUID-VkHdrMetadataEXT-sType-sType"></a>
  VUID-VkHdrMetadataEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_HDR_METADATA_EXT`

- <a href="#VUID-VkHdrMetadataEXT-pNext-pNext"
  id="VUID-VkHdrMetadataEXT-pNext-pNext"></a>
  VUID-VkHdrMetadataEXT-pNext-pNext  
  `pNext` **must** be `NULL`

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The validity and use of this data is outside the scope of
Vulkan.</p></td>
</tr>
</tbody>
</table>

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

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
