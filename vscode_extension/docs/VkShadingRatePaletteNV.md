# VkShadingRatePaletteNV(3) Manual Page

## Name

VkShadingRatePaletteNV - Structure specifying a single shading rate
palette



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkShadingRatePaletteNV` structure specifies to contents of a single
shading rate image palette and is defined as:

``` c
// Provided by VK_NV_shading_rate_image
typedef struct VkShadingRatePaletteNV {
    uint32_t                              shadingRatePaletteEntryCount;
    const VkShadingRatePaletteEntryNV*    pShadingRatePaletteEntries;
} VkShadingRatePaletteNV;
```

## <a href="#_members" class="anchor"></a>Members

- `shadingRatePaletteEntryCount` specifies the number of entries in the
  shading rate image palette.

- `pShadingRatePaletteEntries` is a pointer to an array of
  [VkShadingRatePaletteEntryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShadingRatePaletteEntryNV.html) enums
  defining the shading rate for each palette entry.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkShadingRatePaletteNV-shadingRatePaletteEntryCount-02071"
  id="VUID-VkShadingRatePaletteNV-shadingRatePaletteEntryCount-02071"></a>
  VUID-VkShadingRatePaletteNV-shadingRatePaletteEntryCount-02071  
  `shadingRatePaletteEntryCount` **must** be between `1` and
  `VkPhysicalDeviceShadingRateImagePropertiesNV`::`shadingRatePaletteSize`,
  inclusive

Valid Usage (Implicit)

- <a
  href="#VUID-VkShadingRatePaletteNV-pShadingRatePaletteEntries-parameter"
  id="VUID-VkShadingRatePaletteNV-pShadingRatePaletteEntries-parameter"></a>
  VUID-VkShadingRatePaletteNV-pShadingRatePaletteEntries-parameter  
  `pShadingRatePaletteEntries` **must** be a valid pointer to an array
  of `shadingRatePaletteEntryCount` valid
  [VkShadingRatePaletteEntryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShadingRatePaletteEntryNV.html) values

- <a
  href="#VUID-VkShadingRatePaletteNV-shadingRatePaletteEntryCount-arraylength"
  id="VUID-VkShadingRatePaletteNV-shadingRatePaletteEntryCount-arraylength"></a>
  VUID-VkShadingRatePaletteNV-shadingRatePaletteEntryCount-arraylength  
  `shadingRatePaletteEntryCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_shading_rate_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shading_rate_image.html),
[VkPipelineViewportShadingRateImageStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportShadingRateImageStateCreateInfoNV.html),
[VkShadingRatePaletteEntryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShadingRatePaletteEntryNV.html),
[vkCmdSetViewportShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportShadingRatePaletteNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkShadingRatePaletteNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
