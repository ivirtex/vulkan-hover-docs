# VkBlendFactor(3) Manual Page

## Name

VkBlendFactor - Framebuffer blending factors



## <a href="#_c_specification" class="anchor"></a>C Specification

The source and destination color and alpha blending factors are selected
from the enum:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkBlendFactor {
    VK_BLEND_FACTOR_ZERO = 0,
    VK_BLEND_FACTOR_ONE = 1,
    VK_BLEND_FACTOR_SRC_COLOR = 2,
    VK_BLEND_FACTOR_ONE_MINUS_SRC_COLOR = 3,
    VK_BLEND_FACTOR_DST_COLOR = 4,
    VK_BLEND_FACTOR_ONE_MINUS_DST_COLOR = 5,
    VK_BLEND_FACTOR_SRC_ALPHA = 6,
    VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA = 7,
    VK_BLEND_FACTOR_DST_ALPHA = 8,
    VK_BLEND_FACTOR_ONE_MINUS_DST_ALPHA = 9,
    VK_BLEND_FACTOR_CONSTANT_COLOR = 10,
    VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR = 11,
    VK_BLEND_FACTOR_CONSTANT_ALPHA = 12,
    VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA = 13,
    VK_BLEND_FACTOR_SRC_ALPHA_SATURATE = 14,
    VK_BLEND_FACTOR_SRC1_COLOR = 15,
    VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR = 16,
    VK_BLEND_FACTOR_SRC1_ALPHA = 17,
    VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA = 18,
} VkBlendFactor;
```

## <a href="#_description" class="anchor"></a>Description

The semantics of the enum values are described in the table below:

| [VkBlendFactor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html) | RGB Blend Factors (S<sub>r</sub>,S<sub>g</sub>,S<sub>b</sub>) or (D<sub>r</sub>,D<sub>g</sub>,D<sub>b</sub>) | Alpha Blend Factor (S<sub>a</sub> or D<sub>a</sub>) |
|----|----|----|
| `VK_BLEND_FACTOR_ZERO` | (0,0,0) | 0 |
| `VK_BLEND_FACTOR_ONE` | (1,1,1) | 1 |
| `VK_BLEND_FACTOR_SRC_COLOR` | (R<sub>s0</sub>,G<sub>s0</sub>,B<sub>s0</sub>) | A<sub>s0</sub> |
| `VK_BLEND_FACTOR_ONE_MINUS_SRC_COLOR` | (1-R<sub>s0</sub>,1-G<sub>s0</sub>,1-B<sub>s0</sub>) | 1-A<sub>s0</sub> |
| `VK_BLEND_FACTOR_DST_COLOR` | (R<sub>d</sub>,G<sub>d</sub>,B<sub>d</sub>) | A<sub>d</sub> |
| `VK_BLEND_FACTOR_ONE_MINUS_DST_COLOR` | (1-R<sub>d</sub>,1-G<sub>d</sub>,1-B<sub>d</sub>) | 1-A<sub>d</sub> |
| `VK_BLEND_FACTOR_SRC_ALPHA` | (A<sub>s0</sub>,A<sub>s0</sub>,A<sub>s0</sub>) | A<sub>s0</sub> |
| `VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA` | (1-A<sub>s0</sub>,1-A<sub>s0</sub>,1-A<sub>s0</sub>) | 1-A<sub>s0</sub> |
| `VK_BLEND_FACTOR_DST_ALPHA` | (A<sub>d</sub>,A<sub>d</sub>,A<sub>d</sub>) | A<sub>d</sub> |
| `VK_BLEND_FACTOR_ONE_MINUS_DST_ALPHA` | (1-A<sub>d</sub>,1-A<sub>d</sub>,1-A<sub>d</sub>) | 1-A<sub>d</sub> |
| `VK_BLEND_FACTOR_CONSTANT_COLOR` | (R<sub>c</sub>,G<sub>c</sub>,B<sub>c</sub>) | A<sub>c</sub> |
| `VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR` | (1-R<sub>c</sub>,1-G<sub>c</sub>,1-B<sub>c</sub>) | 1-A<sub>c</sub> |
| `VK_BLEND_FACTOR_CONSTANT_ALPHA` | (A<sub>c</sub>,A<sub>c</sub>,A<sub>c</sub>) | A<sub>c</sub> |
| `VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA` | (1-A<sub>c</sub>,1-A<sub>c</sub>,1-A<sub>c</sub>) | 1-A<sub>c</sub> |
| `VK_BLEND_FACTOR_SRC_ALPHA_SATURATE` | (f,f,f); f = min(A<sub>s0</sub>,1-A<sub>d</sub>) | 1 |
| `VK_BLEND_FACTOR_SRC1_COLOR` | (R<sub>s1</sub>,G<sub>s1</sub>,B<sub>s1</sub>) | A<sub>s1</sub> |
| `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR` | (1-R<sub>s1</sub>,1-G<sub>s1</sub>,1-B<sub>s1</sub>) | 1-A<sub>s1</sub> |
| `VK_BLEND_FACTOR_SRC1_ALPHA` | (A<sub>s1</sub>,A<sub>s1</sub>,A<sub>s1</sub>) | A<sub>s1</sub> |
| `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA` | (1-A<sub>s1</sub>,1-A<sub>s1</sub>,1-A<sub>s1</sub>) | 1-A<sub>s1</sub> |

Table 1. Blend Factors

In this table, the following conventions are used:

- R<sub>s0</sub>,G<sub>s0</sub>,B<sub>s0</sub> and A<sub>s0</sub>
  represent the first source color R, G, B, and A components,
  respectively, for the fragment output location corresponding to the
  color attachment being blended.

- R<sub>s1</sub>,G<sub>s1</sub>,B<sub>s1</sub> and A<sub>s1</sub>
  represent the second source color R, G, B, and A components,
  respectively, used in dual source blending modes, for the fragment
  output location corresponding to the color attachment being blended.

- R<sub>d</sub>,G<sub>d</sub>,B<sub>d</sub> and A<sub>d</sub> represent
  the R, G, B, and A components of the destination color. That is, the
  color currently in the corresponding color attachment for this
  fragment/sample.

- R<sub>c</sub>,G<sub>c</sub>,B<sub>c</sub> and A<sub>c</sub> represent
  the blend constant R, G, B, and A components, respectively.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendEquationEXT.html),
[VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBlendFactor"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
