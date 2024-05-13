# VkDisplacementMicromapFormatNV(3) Manual Page

## Name

VkDisplacementMicromapFormatNV - Format enum for displacement micromaps



## <a href="#_c_specification" class="anchor"></a>C Specification

Formats which **can** be set in
[VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html)::`format` and
[VkMicromapTriangleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTriangleEXT.html)::`format` for
micromap builds, are:

``` c
// Provided by VK_NV_displacement_micromap
typedef enum VkDisplacementMicromapFormatNV {
    VK_DISPLACEMENT_MICROMAP_FORMAT_64_TRIANGLES_64_BYTES_NV = 1,
    VK_DISPLACEMENT_MICROMAP_FORMAT_256_TRIANGLES_128_BYTES_NV = 2,
    VK_DISPLACEMENT_MICROMAP_FORMAT_1024_TRIANGLES_128_BYTES_NV = 3,
} VkDisplacementMicromapFormatNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DISPLACEMENT_MICROMAP_FORMAT_64_TRIANGLES_64_BYTES_NV` indicates
  that the given micromap format encodes 64 micro-triangles worth of
  displacements in 64 bytes as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#displacement-micromap-encoding"
  target="_blank" rel="noopener">Displacement Micromap Encoding</a>.

- `VK_DISPLACEMENT_MICROMAP_FORMAT_256_TRIANGLES_128_BYTES_NV` indicates
  that the given micromap format encodes 256 micro-triangles worth of
  displacements in 128 bytes as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#displacement-micromap-encoding"
  target="_blank" rel="noopener">Displacement Micromap Encoding</a>.

- `VK_DISPLACEMENT_MICROMAP_FORMAT_1024_TRIANGLES_128_BYTES_NV`
  indicates that the given micromap format encodes 1024 micro-triangles
  worth of displacements in 128 bytes as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#displacement-micromap-encoding"
  target="_blank" rel="noopener">Displacement Micromap Encoding</a>.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>For compactness, these values are stored as 16-bit in some
structures.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_displacement_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_displacement_micromap.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplacementMicromapFormatNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
