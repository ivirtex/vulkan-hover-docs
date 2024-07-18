# VkPhysicalDevice4444FormatsFeaturesEXT(3) Manual Page

## Name

VkPhysicalDevice4444FormatsFeaturesEXT - Structure describing additional
4444 formats supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDevice4444FormatsFeaturesEXT` structure is defined as:

``` c
// Provided by VK_EXT_4444_formats
typedef struct VkPhysicalDevice4444FormatsFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           formatA4R4G4B4;
    VkBool32           formatA4B4G4R4;
} VkPhysicalDevice4444FormatsFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-formatA4R4G4B4"></span> `formatA4R4G4B4` indicates
  that the implementation **must** support using a
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `VK_FORMAT_A4R4G4B4_UNORM_PACK16_EXT`
  with at least the following
  [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html):

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`

  - `VK_FORMAT_FEATURE_BLIT_SRC_BIT`

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`

- <span id="features-formatA4B4G4R4"></span> `formatA4B4G4R4` indicates
  that the implementation **must** support using a
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `VK_FORMAT_A4B4G4R4_UNORM_PACK16_EXT`
  with at least the following
  [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html):

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`

  - `VK_FORMAT_FEATURE_BLIT_SRC_BIT`

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDevice4444FormatsFeaturesEXT` structure is included in
the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDevice4444FormatsFeaturesEXT` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDevice4444FormatsFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDevice4444FormatsFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDevice4444FormatsFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_4444_FORMATS_FEATURES_EXT`

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Although the formats defined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_4444_formats.html"><code>VK_EXT_4444_formats</code></a>
extension were promoted to Vulkan 1.3 as optional formats, the <a
href="VkPhysicalDevice4444FormatsFeaturesEXT.html">VkPhysicalDevice4444FormatsFeaturesEXT</a>
structure was not promoted to Vulkan 1.3.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_4444_formats](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_4444_formats.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDevice4444FormatsFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
