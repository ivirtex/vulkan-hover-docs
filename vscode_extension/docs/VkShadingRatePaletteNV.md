# VkShadingRatePaletteNV(3) Manual Page

## Name

VkShadingRatePaletteNV - Structure specifying a single shading rate palette



## [](#_c_specification)C Specification

The `VkShadingRatePaletteNV` structure specifies to contents of a single shading rate image palette and is defined as:

```c++
// Provided by VK_NV_shading_rate_image
typedef struct VkShadingRatePaletteNV {
    uint32_t                              shadingRatePaletteEntryCount;
    const VkShadingRatePaletteEntryNV*    pShadingRatePaletteEntries;
} VkShadingRatePaletteNV;
```

## [](#_members)Members

- `shadingRatePaletteEntryCount` specifies the number of entries in the shading rate image palette.
- `pShadingRatePaletteEntries` is a pointer to an array of [VkShadingRatePaletteEntryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShadingRatePaletteEntryNV.html) enums defining the shading rate for each palette entry.

## [](#_description)Description

Valid Usage

- [](#VUID-VkShadingRatePaletteNV-shadingRatePaletteEntryCount-02071)VUID-VkShadingRatePaletteNV-shadingRatePaletteEntryCount-02071  
  `shadingRatePaletteEntryCount` **must** be between `1` and `VkPhysicalDeviceShadingRateImagePropertiesNV`::`shadingRatePaletteSize`, inclusive

Valid Usage (Implicit)

- [](#VUID-VkShadingRatePaletteNV-pShadingRatePaletteEntries-parameter)VUID-VkShadingRatePaletteNV-pShadingRatePaletteEntries-parameter  
  `pShadingRatePaletteEntries` **must** be a valid pointer to an array of `shadingRatePaletteEntryCount` valid [VkShadingRatePaletteEntryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShadingRatePaletteEntryNV.html) values
- [](#VUID-VkShadingRatePaletteNV-shadingRatePaletteEntryCount-arraylength)VUID-VkShadingRatePaletteNV-shadingRatePaletteEntryCount-arraylength  
  `shadingRatePaletteEntryCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_NV\_shading\_rate\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shading_rate_image.html), [VkPipelineViewportShadingRateImageStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportShadingRateImageStateCreateInfoNV.html), [VkShadingRatePaletteEntryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShadingRatePaletteEntryNV.html), [vkCmdSetViewportShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetViewportShadingRatePaletteNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkShadingRatePaletteNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0