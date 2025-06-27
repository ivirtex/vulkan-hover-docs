# VkSamplerBorderColorComponentMappingCreateInfoEXT(3) Manual Page

## Name

VkSamplerBorderColorComponentMappingCreateInfoEXT - Structure specifying
the component mapping of the border color



## <a href="#_c_specification" class="anchor"></a>C Specification

If the sampler is created with `VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK`,
`VK_BORDER_COLOR_INT_OPAQUE_BLACK`, `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT`,
or `VK_BORDER_COLOR_INT_CUSTOM_EXT` `borderColor`, and that sampler will
be combined with an image view that does not have an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
target="_blank" rel="noopener">identity swizzle</a>, and
[VkPhysicalDeviceBorderColorSwizzleFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBorderColorSwizzleFeaturesEXT.html)::`borderColorSwizzleFromImage`
is not enabled, then it is necessary to specify the component mapping of
the border color, by including the
`VkSamplerBorderColorComponentMappingCreateInfoEXT` structure in the
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)::`pNext` chain, to get
defined results.

The `VkSamplerBorderColorComponentMappingCreateInfoEXT` structure is
defined as:

``` c
// Provided by VK_EXT_border_color_swizzle
typedef struct VkSamplerBorderColorComponentMappingCreateInfoEXT {
    VkStructureType       sType;
    const void*           pNext;
    VkComponentMapping    components;
    VkBool32              srgb;
} VkSamplerBorderColorComponentMappingCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `components` is a [VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html)
  structure specifying a remapping of the border color components.

- `srgb` indicates that the sampler will be combined with an image view
  that has an image format which is sRGB encoded.

## <a href="#_description" class="anchor"></a>Description

The [VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html) `components` member
describes a remapping from components of the border color to components
of the vector returned by shader image instructions when the border
color is used.

Valid Usage

- <a
  href="#VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-borderColorSwizzle-06437"
  id="VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-borderColorSwizzle-06437"></a>
  VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-borderColorSwizzle-06437  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-borderColorSwizzle"
  target="_blank" rel="noopener"><code>borderColorSwizzle</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a
  href="#VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-sType-sType"
  id="VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-sType-sType"></a>
  VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SAMPLER_BORDER_COLOR_COMPONENT_MAPPING_CREATE_INFO_EXT`

- <a
  href="#VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-components-parameter"
  id="VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-components-parameter"></a>
  VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-components-parameter  
  `components` **must** be a valid
  [VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_border_color_swizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_border_color_swizzle.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerBorderColorComponentMappingCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
