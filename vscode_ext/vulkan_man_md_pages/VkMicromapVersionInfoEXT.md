# VkMicromapVersionInfoEXT(3) Manual Page

## Name

VkMicromapVersionInfoEXT - Micromap version information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMicromapVersionInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_opacity_micromap
typedef struct VkMicromapVersionInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    const uint8_t*     pVersionData;
} VkMicromapVersionInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pVersionData` is a pointer to the version header of a micromap as
  defined in
  [vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMicromapToMemoryEXT.html)

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
<p><code>pVersionData</code> is a <em>pointer</em> to an array of
2×<code>VK_UUID_SIZE</code> <code>uint8_t</code> values instead of two
<code>VK_UUID_SIZE</code> arrays as the expected use case for this
member is to be pointed at the header of a previously serialized
micromap (via <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMicromapToMemoryEXT.html">vkCmdCopyMicromapToMemoryEXT</a>
or <a
href="vkCopyMicromapToMemoryEXT.html">vkCopyMicromapToMemoryEXT</a>)
that is loaded in memory. Using arrays would necessitate extra memory
copies of the UUIDs.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkMicromapVersionInfoEXT-sType-sType"
  id="VUID-VkMicromapVersionInfoEXT-sType-sType"></a>
  VUID-VkMicromapVersionInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MICROMAP_VERSION_INFO_EXT`

- <a href="#VUID-VkMicromapVersionInfoEXT-pNext-pNext"
  id="VUID-VkMicromapVersionInfoEXT-pNext-pNext"></a>
  VUID-VkMicromapVersionInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkMicromapVersionInfoEXT-pVersionData-parameter"
  id="VUID-VkMicromapVersionInfoEXT-pVersionData-parameter"></a>
  VUID-VkMicromapVersionInfoEXT-pVersionData-parameter  
  `pVersionData` **must** be a valid pointer to an array of
  2×VK_UUID_SIZE `uint8_t` values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDeviceMicromapCompatibilityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMicromapCompatibilityEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMicromapVersionInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
