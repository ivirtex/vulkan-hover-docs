# VkPipelineViewportShadingRateImageStateCreateInfoNV(3) Manual Page

## Name

VkPipelineViewportShadingRateImageStateCreateInfoNV - Structure specifying parameters controlling shading rate image usage



## [](#_c_specification)C Specification

If the `pNext` chain of [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html) includes a `VkPipelineViewportShadingRateImageStateCreateInfoNV` structure, then that structure includes parameters controlling the shading rate.

The `VkPipelineViewportShadingRateImageStateCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_shading_rate_image
typedef struct VkPipelineViewportShadingRateImageStateCreateInfoNV {
    VkStructureType                  sType;
    const void*                      pNext;
    VkBool32                         shadingRateImageEnable;
    uint32_t                         viewportCount;
    const VkShadingRatePaletteNV*    pShadingRatePalettes;
} VkPipelineViewportShadingRateImageStateCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `shadingRateImageEnable` specifies whether shading rate image and palettes are used during rasterization.
- `viewportCount` specifies the number of per-viewport palettes used to translate values stored in shading rate images.
- `pShadingRatePalettes` is a pointer to an array of [VkShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShadingRatePaletteNV.html) structures defining the palette for each viewport. If the shading rate palette state is dynamic, this member is ignored.

## [](#_description)Description

If this structure is not present, `shadingRateImageEnable` is considered to be `VK_FALSE`, and the shading rate image and palettes are not used.

Valid Usage

- [](#VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-viewportCount-02054)VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-viewportCount-02054  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `viewportCount` **must** be `0` or `1`
- [](#VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-viewportCount-02055)VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-viewportCount-02055  
  `viewportCount` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxViewports`
- [](#VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-shadingRateImageEnable-02056)VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-shadingRateImageEnable-02056  
  If `shadingRateImageEnable` is `VK_TRUE`, `viewportCount` **must** be greater or equal to the `viewportCount` member of [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html)

Valid Usage (Implicit)

- [](#VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-sType-sType)VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_SHADING_RATE_IMAGE_STATE_CREATE_INFO_NV`

## [](#_see_also)See Also

[VK\_NV\_shading\_rate\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shading_rate_image.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShadingRatePaletteNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineViewportShadingRateImageStateCreateInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0