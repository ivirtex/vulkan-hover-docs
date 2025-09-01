# VkSamplerCustomBorderColorCreateInfoEXT(3) Manual Page

## Name

VkSamplerCustomBorderColorCreateInfoEXT - Structure specifying custom border color



## [](#_c_specification)C Specification

In addition to the predefined border color values, applications **can** provide a custom border color value by including the `VkSamplerCustomBorderColorCreateInfoEXT` structure in the [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)::`pNext` chain.

The `VkSamplerCustomBorderColorCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_custom_border_color
typedef struct VkSamplerCustomBorderColorCreateInfoEXT {
    VkStructureType      sType;
    const void*          pNext;
    VkClearColorValue    customBorderColor;
    VkFormat             format;
} VkSamplerCustomBorderColorCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `customBorderColor` is a [VkClearColorValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearColorValue.html) representing the desired custom sampler border color.
- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) representing the format of the sampled image view(s). This field may be `VK_FORMAT_UNDEFINED` if the [`customBorderColorWithoutFormat`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-customBorderColorWithoutFormat) feature is enabled.

## [](#_description)Description

Note

If `format` is a depth/stencil format, the aspect is determined by the value of [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)::`borderColor`. If [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)::`borderColor` is `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT`, the depth aspect is considered. If [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)::`borderColor` is `VK_BORDER_COLOR_INT_CUSTOM_EXT`, the stencil aspect is considered.

If `format` is `VK_FORMAT_UNDEFINED`, the [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)::`borderColor` is `VK_BORDER_COLOR_INT_CUSTOM_EXT`, and the sampler is used with an image with a stencil format, then the implementation **must** source the custom border color from either the first or second components of [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)::`borderColor` and **should** source it from the first component.

Valid Usage

- [](#VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-07605)VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-07605  
  If `format` is not `VK_FORMAT_UNDEFINED` and `format` is not a depth/stencil format then the [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)::`borderColor` type **must** match the sampled type of the provided `format`, as shown in the *SPIR-V Type* column of the [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-numericformat](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-numericformat) table
- [](#VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-04014)VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-04014  
  If the [`customBorderColorWithoutFormat`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-customBorderColorWithoutFormat) feature is not enabled then `format` **must** not be `VK_FORMAT_UNDEFINED`
- [](#VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-04015)VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-04015  
  If the sampler is used to sample an image view of `VK_FORMAT_B4G4R4A4_UNORM_PACK16`, `VK_FORMAT_B5G6R5_UNORM_PACK16`, `VK_FORMAT_A1B5G5R5_UNORM_PACK16`, or `VK_FORMAT_B5G5R5A1_UNORM_PACK16` format then `format` **must** not be `VK_FORMAT_UNDEFINED`

Valid Usage (Implicit)

- [](#VUID-VkSamplerCustomBorderColorCreateInfoEXT-sType-sType)VUID-VkSamplerCustomBorderColorCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLER_CUSTOM_BORDER_COLOR_CREATE_INFO_EXT`
- [](#VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-parameter)VUID-VkSamplerCustomBorderColorCreateInfoEXT-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value

## [](#_see_also)See Also

[VK\_EXT\_custom\_border\_color](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_custom_border_color.html), [VkClearColorValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearColorValue.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerCustomBorderColorCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0