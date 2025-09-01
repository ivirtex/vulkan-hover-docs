# VkBlendOverlapEXT(3) Manual Page

## Name

VkBlendOverlapEXT - Enumerant specifying the blend overlap parameter



## [](#_c_specification)C Specification

The weighting functions p0, p1, and p2 are defined in table [Advanced Blend Overlap Modes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blend-advanced-overlap-modes). In these functions, the A components of the source and destination colors are taken to indicate the portion of the pixel covered by the fragment (source) and the fragments previously accumulated in the pixel (destination). The functions p0, p1, and p2 approximate the relative portion of the pixel covered by the intersection of the source and destination, covered only by the source, and covered only by the destination, respectively.

Possible values of [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`blendOverlap`, specifying the blend overlap functions, are:

```c++
// Provided by VK_EXT_blend_operation_advanced
typedef enum VkBlendOverlapEXT {
    VK_BLEND_OVERLAP_UNCORRELATED_EXT = 0,
    VK_BLEND_OVERLAP_DISJOINT_EXT = 1,
    VK_BLEND_OVERLAP_CONJOINT_EXT = 2,
} VkBlendOverlapEXT;
```

## [](#_description)Description

- `VK_BLEND_OVERLAP_UNCORRELATED_EXT` specifies that there is no correlation between the source and destination coverage.
- `VK_BLEND_OVERLAP_CONJOINT_EXT` specifies that the source and destination coverage are considered to have maximal overlap.
- `VK_BLEND_OVERLAP_DISJOINT_EXT` specifies that the source and destination coverage are considered to have minimal overlap.

Table 1. Advanced Blend Overlap Modes   Overlap Mode Weighting Equations

`VK_BLEND_OVERLAP_UNCORRELATED_EXT`

p0​(As​,Ad​)p1​(As​,Ad​)p2​(As​,Ad​)​=As​Ad​=As​(1−Ad​)=Ad​(1−As​)​

`VK_BLEND_OVERLAP_CONJOINT_EXT`

p0​(As​,Ad​)p1​(As​,Ad​)p2​(As​,Ad​)​=min(As​,Ad​)=max(As​−Ad​,0)=max(Ad​−As​,0)​

`VK_BLEND_OVERLAP_DISJOINT_EXT`

p0​(As​,Ad​)p1​(As​,Ad​)p2​(As​,Ad​)​=max(As​+Ad​−1,0)=min(As​,1−Ad​)=min(Ad​,1−As​)​

## [](#_see_also)See Also

[VK\_EXT\_blend\_operation\_advanced](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_blend_operation_advanced.html), [VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorBlendAdvancedEXT.html), [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBlendOverlapEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0