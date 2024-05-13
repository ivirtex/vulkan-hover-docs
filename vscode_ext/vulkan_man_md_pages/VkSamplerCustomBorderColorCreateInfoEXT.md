# VkSamplerCustomBorderColorCreateInfoEXT(3) Manual Page

## Name

VkSamplerCustomBorderColorCreateInfoEXT - Structure specifying custom
border color



## <a href="#_c_specification" class="anchor"></a>C Specification

In addition to the predefined border color values, applications **can**
provide a custom border color value by including the
`VkSamplerCustomBorderColorCreateInfoEXT` structure in the
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)::`pNext` chain.

The `VkSamplerCustomBorderColorCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_custom_border_color
typedef struct VkSamplerCustomBorderColorCreateInfoEXT {
    VkStructureType      sType;
    const void*          pNext;
    VkClearColorValue    customBorderColor;
    VkFormat             format;
} VkSamplerCustomBorderColorCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `customBorderColor` is a [VkClearColorValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearColorValue.html)
  representing the desired custom sampler border color.

- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) representing the format of the
  sampled image view(s). This field may be `VK_FORMAT_UNDEFINED` if the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-customBorderColorWithoutFormat"
  target="_blank"
  rel="noopener"><code>customBorderColorWithoutFormat</code></a> feature
  is enabled.

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
<p>If <code>format</code> is a depth/stencil format, the aspect is
determined by the value of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html">VkSamplerCreateInfo</a>::<code>borderColor</code>.
If <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html">VkSamplerCreateInfo</a>::<code>borderColor</code>
is <code>VK_BORDER_COLOR_FLOAT_CUSTOM_EXT</code>, the depth aspect is
considered. If <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html">VkSamplerCreateInfo</a>::<code>borderColor</code>
is <code>VK_BORDER_COLOR_INT_CUSTOM_EXT</code>, the stencil aspect is
considered.</p>
<p>If <code>format</code> is <code>VK_FORMAT_UNDEFINED</code>, the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html">VkSamplerCreateInfo</a>::<code>borderColor</code>
is <code>VK_BORDER_COLOR_INT_CUSTOM_EXT</code>, and the sampler is used
with an image with a stencil format, then the implementation
<strong>must</strong> source the custom border color from either the
first or second components of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html">VkSamplerCreateInfo</a>::<code>customBorderColor</code>
and <strong>should</strong> source it from the first component.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-07605"
  id="VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-07605"></a>
  VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-07605  
  If `format` is not `VK_FORMAT_UNDEFINED` and `format` is not a
  depth/stencil format then the
  [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)::`borderColor` type
  **must** match the sampled type of the provided `format`, as shown in
  the *SPIR-V Type* column of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-numericformat"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-numericformat</a>
  table

- <a href="#VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-04014"
  id="VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-04014"></a>
  VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-04014  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-customBorderColorWithoutFormat"
  target="_blank"
  rel="noopener"><code>customBorderColorWithoutFormat</code></a> feature
  is not enabled then `format` **must** not be `VK_FORMAT_UNDEFINED`

- <a href="#VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-04015"
  id="VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-04015"></a>
  VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-04015  
  If the sampler is used to sample an image view of
  `VK_FORMAT_B4G4R4A4_UNORM_PACK16`, `VK_FORMAT_B5G6R5_UNORM_PACK16`,
  `VK_FORMAT_A1B5G5R5_UNORM_PACK16_KHR`, or
  `VK_FORMAT_B5G5R5A1_UNORM_PACK16` format then `format` **must** not be
  `VK_FORMAT_UNDEFINED`

Valid Usage (Implicit)

- <a href="#VUID-VkSamplerCustomBorderColorCreateInfoEXT-sType-sType"
  id="VUID-VkSamplerCustomBorderColorCreateInfoEXT-sType-sType"></a>
  VUID-VkSamplerCustomBorderColorCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SAMPLER_CUSTOM_BORDER_COLOR_CREATE_INFO_EXT`

- <a href="#VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-parameter"
  id="VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-parameter"></a>
  VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_custom_border_color](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_custom_border_color.html),
[VkClearColorValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearColorValue.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerCustomBorderColorCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
