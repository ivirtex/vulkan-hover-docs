# VkOpacityMicromapFormatEXT(3) Manual Page

## Name

VkOpacityMicromapFormatEXT - Format enum for opacity micromaps



## <a href="#_c_specification" class="anchor"></a>C Specification

Formats which **can** be set in
[VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html)::`format` and
[VkMicromapTriangleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTriangleEXT.html)::`format` for
micromap builds, are:

``` c
// Provided by VK_EXT_opacity_micromap
typedef enum VkOpacityMicromapFormatEXT {
    VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT = 1,
    VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT = 2,
} VkOpacityMicromapFormatEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT` indicates that the given
  micromap format has one bit per subtriangle encoding either fully
  opaque or fully transparent.

- `VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT` indicates that the given
  micromap format has two bits per subtriangle encoding four modes which
  can be interpreted as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-opacity-micromap"
  target="_blank" rel="noopener">ray traversal</a>.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>For compactness, these values are stored as 16-bit in some
structures.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOpacityMicromapFormatEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
