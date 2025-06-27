# VkBlendOverlapEXT(3) Manual Page

## Name

VkBlendOverlapEXT - Enumerant specifying the blend overlap parameter



## <a href="#_c_specification" class="anchor"></a>C Specification

The weighting functions p<sub>0</sub>, p<sub>1</sub>, and p<sub>2</sub>
are defined in table <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blend-advanced-overlap-modes"
target="_blank" rel="noopener">Advanced Blend Overlap Modes</a>. In
these functions, the A components of the source and destination colors
are taken to indicate the portion of the pixel covered by the fragment
(source) and the fragments previously accumulated in the pixel
(destination). The functions p<sub>0</sub>, p<sub>1</sub>, and
p<sub>2</sub> approximate the relative portion of the pixel covered by
the intersection of the source and destination, covered only by the
source, and covered only by the destination, respectively.

Possible values of
[VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`blendOverlap`,
specifying the blend overlap functions, are:

``` c
// Provided by VK_EXT_blend_operation_advanced
typedef enum VkBlendOverlapEXT {
    VK_BLEND_OVERLAP_UNCORRELATED_EXT = 0,
    VK_BLEND_OVERLAP_DISJOINT_EXT = 1,
    VK_BLEND_OVERLAP_CONJOINT_EXT = 2,
} VkBlendOverlapEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_BLEND_OVERLAP_UNCORRELATED_EXT` specifies that there is no
  correlation between the source and destination coverage.

- `VK_BLEND_OVERLAP_CONJOINT_EXT` specifies that the source and
  destination coverage are considered to have maximal overlap.

- `VK_BLEND_OVERLAP_DISJOINT_EXT` specifies that the source and
  destination coverage are considered to have minimal overlap.

| Overlap Mode | Weighting Equations |
|----|----|
| `VK_BLEND_OVERLAP_UNCORRELATED_EXT` | p0​(As​,Ad​)p1​(As​,Ad​)p2​(As​,Ad​)​=As​Ad​=As​(1−Ad​)=Ad​(1−As​)​ |
| `VK_BLEND_OVERLAP_CONJOINT_EXT` | p0​(As​,Ad​)p1​(As​,Ad​)p2​(As​,Ad​)​=min(As​,Ad​)=max(As​−Ad​,0)=max(Ad​−As​,0)​ |
| `VK_BLEND_OVERLAP_DISJOINT_EXT` | p0​(As​,Ad​)p1​(As​,Ad​)p2​(As​,Ad​)​=max(As​+Ad​−1,0)=min(As​,1−Ad​)=min(Ad​,1−As​)​ |

Table 1. Advanced Blend Overlap Modes
{#framebuffer-blend-advanced-overlap-modes}

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_blend_operation_advanced](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_blend_operation_advanced.html),
[VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendAdvancedEXT.html),
[VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBlendOverlapEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
