# VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT - Structure describing
whether the implementation supports additional 2-plane 444
Y′C<sub>B</sub>C<sub>R</sub> formats



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_ycbcr_2plane_444_formats
typedef struct VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           ycbcr2plane444Formats;
} VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-ycbcr2plane444Formats"></span>
  `ycbcr2plane444Formats` indicates that the implementation supports the
  following 2-plane 444 Y′C<sub>B</sub>C<sub>R</sub> formats:

  - `VK_FORMAT_G8_B8R8_2PLANE_444_UNORM`

  - `VK_FORMAT_G10X6_B10X6R10X6_2PLANE_444_UNORM_3PACK16`

  - `VK_FORMAT_G12X4_B12X4R12X4_2PLANE_444_UNORM_3PACK16`

  - `VK_FORMAT_G16_B16R16_2PLANE_444_UNORM`

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT` **can** also be used
in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_YCBCR_2_PLANE_444_FORMATS_FEATURES_EXT`

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Although the formats defined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_ycbcr_2plane_444_formats.html"><code>VK_EXT_ycbcr_2plane_444_formats</code></a>
were promoted to Vulkan 1.3 as optional formats, the <a
href="VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT.html">VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT</a>
structure was not promoted to Vulkan 1.3.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_ycbcr_2plane_444_formats](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_ycbcr_2plane_444_formats.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
