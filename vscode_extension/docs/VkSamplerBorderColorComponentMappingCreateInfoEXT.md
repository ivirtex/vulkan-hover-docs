# VkSamplerBorderColorComponentMappingCreateInfoEXT(3) Manual Page

## Name

VkSamplerBorderColorComponentMappingCreateInfoEXT - Structure specifying the component mapping of the border color



## [](#_c_specification)C Specification

If the sampler is created with `VK_BORDER_COLOR_FLOAT_OPAQUE_BLACK`, `VK_BORDER_COLOR_INT_OPAQUE_BLACK`, `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT`, or `VK_BORDER_COLOR_INT_CUSTOM_EXT` `borderColor`, and that sampler will be combined with an image view that does not have an [identity swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings), and [VkPhysicalDeviceBorderColorSwizzleFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBorderColorSwizzleFeaturesEXT.html)::`borderColorSwizzleFromImage` is not enabled, then it is necessary to specify the component mapping of the border color, by including the `VkSamplerBorderColorComponentMappingCreateInfoEXT` structure in the [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)::`pNext` chain, to get defined results.

The `VkSamplerBorderColorComponentMappingCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_border_color_swizzle
typedef struct VkSamplerBorderColorComponentMappingCreateInfoEXT {
    VkStructureType       sType;
    const void*           pNext;
    VkComponentMapping    components;
    VkBool32              srgb;
} VkSamplerBorderColorComponentMappingCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `components` is a [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html) structure specifying a remapping of the border color components.
- `srgb` indicates that the sampler will be combined with an image view that has an image format which is sRGB encoded.

## [](#_description)Description

The [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html) `components` member describes a remapping from components of the border color to components of the vector returned by shader image instructions when the border color is used.

Valid Usage

- [](#VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-borderColorSwizzle-06437)VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-borderColorSwizzle-06437  
  The [`borderColorSwizzle`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-borderColorSwizzle) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-sType-sType)VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLER_BORDER_COLOR_COMPONENT_MAPPING_CREATE_INFO_EXT`
- [](#VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-components-parameter)VUID-VkSamplerBorderColorComponentMappingCreateInfoEXT-components-parameter  
  `components` **must** be a valid [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html) structure

## [](#_see_also)See Also

[VK\_EXT\_border\_color\_swizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_border_color_swizzle.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerBorderColorComponentMappingCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0