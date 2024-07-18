# VkPipelineViewportShadingRateImageStateCreateInfoNV(3) Manual Page

## Name

VkPipelineViewportShadingRateImageStateCreateInfoNV - Structure
specifying parameters controlling shading rate image usage



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)
includes a `VkPipelineViewportShadingRateImageStateCreateInfoNV`
structure, then that structure includes parameters controlling the
shading rate.

The `VkPipelineViewportShadingRateImageStateCreateInfoNV` structure is
defined as:

``` c
// Provided by VK_NV_shading_rate_image
typedef struct VkPipelineViewportShadingRateImageStateCreateInfoNV {
    VkStructureType                  sType;
    const void*                      pNext;
    VkBool32                         shadingRateImageEnable;
    uint32_t                         viewportCount;
    const VkShadingRatePaletteNV*    pShadingRatePalettes;
} VkPipelineViewportShadingRateImageStateCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `shadingRateImageEnable` specifies whether shading rate image and
  palettes are used during rasterization.

- `viewportCount` specifies the number of per-viewport palettes used to
  translate values stored in shading rate images.

- `pShadingRatePalettes` is a pointer to an array of
  [VkShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShadingRatePaletteNV.html) structures
  defining the palette for each viewport. If the shading rate palette
  state is dynamic, this member is ignored.

## <a href="#_description" class="anchor"></a>Description

If this structure is not present, `shadingRateImageEnable` is considered
to be `VK_FALSE`, and the shading rate image and palettes are not used.

Valid Usage

- <a
  href="#VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-viewportCount-02054"
  id="VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-viewportCount-02054"></a>
  VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-viewportCount-02054  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `viewportCount` **must** be `0` or `1`

- <a
  href="#VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-viewportCount-02055"
  id="VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-viewportCount-02055"></a>
  VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-viewportCount-02055  
  `viewportCount` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxViewports`

- <a
  href="#VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-shadingRateImageEnable-02056"
  id="VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-shadingRateImageEnable-02056"></a>
  VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-shadingRateImageEnable-02056  
  If `shadingRateImageEnable` is `VK_TRUE`, `viewportCount` **must** be
  greater or equal to the `viewportCount` member of
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-sType-sType"
  id="VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-sType-sType"></a>
  VUID-VkPipelineViewportShadingRateImageStateCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_SHADING_RATE_IMAGE_STATE_CREATE_INFO_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_shading_rate_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shading_rate_image.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShadingRatePaletteNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineViewportShadingRateImageStateCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
