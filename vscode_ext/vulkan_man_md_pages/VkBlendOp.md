# VkBlendOp(3) Manual Page

## Name

VkBlendOp - Framebuffer blending operations



## <a href="#_c_specification" class="anchor"></a>C Specification

Once the source and destination blend factors have been selected, they
along with the source and destination components are passed to the
blending operations. RGB and alpha components **can** use different
operations. Possible values of [VkBlendOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html), specifying
the operations, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkBlendOp {
    VK_BLEND_OP_ADD = 0,
    VK_BLEND_OP_SUBTRACT = 1,
    VK_BLEND_OP_REVERSE_SUBTRACT = 2,
    VK_BLEND_OP_MIN = 3,
    VK_BLEND_OP_MAX = 4,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_ZERO_EXT = 1000148000,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_SRC_EXT = 1000148001,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_DST_EXT = 1000148002,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_SRC_OVER_EXT = 1000148003,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_DST_OVER_EXT = 1000148004,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_SRC_IN_EXT = 1000148005,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_DST_IN_EXT = 1000148006,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_SRC_OUT_EXT = 1000148007,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_DST_OUT_EXT = 1000148008,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_SRC_ATOP_EXT = 1000148009,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_DST_ATOP_EXT = 1000148010,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_XOR_EXT = 1000148011,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_MULTIPLY_EXT = 1000148012,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_SCREEN_EXT = 1000148013,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_OVERLAY_EXT = 1000148014,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_DARKEN_EXT = 1000148015,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_LIGHTEN_EXT = 1000148016,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_COLORDODGE_EXT = 1000148017,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_COLORBURN_EXT = 1000148018,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_HARDLIGHT_EXT = 1000148019,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_SOFTLIGHT_EXT = 1000148020,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_DIFFERENCE_EXT = 1000148021,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_EXCLUSION_EXT = 1000148022,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_INVERT_EXT = 1000148023,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_INVERT_RGB_EXT = 1000148024,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_LINEARDODGE_EXT = 1000148025,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_LINEARBURN_EXT = 1000148026,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_VIVIDLIGHT_EXT = 1000148027,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_LINEARLIGHT_EXT = 1000148028,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_PINLIGHT_EXT = 1000148029,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_HARDMIX_EXT = 1000148030,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_HSL_HUE_EXT = 1000148031,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_HSL_SATURATION_EXT = 1000148032,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_HSL_COLOR_EXT = 1000148033,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_HSL_LUMINOSITY_EXT = 1000148034,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_PLUS_EXT = 1000148035,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_PLUS_CLAMPED_EXT = 1000148036,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_PLUS_CLAMPED_ALPHA_EXT = 1000148037,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_PLUS_DARKER_EXT = 1000148038,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_MINUS_EXT = 1000148039,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_MINUS_CLAMPED_EXT = 1000148040,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_CONTRAST_EXT = 1000148041,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_INVERT_OVG_EXT = 1000148042,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_RED_EXT = 1000148043,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_GREEN_EXT = 1000148044,
  // Provided by VK_EXT_blend_operation_advanced
    VK_BLEND_OP_BLUE_EXT = 1000148045,
} VkBlendOp;
```

## <a href="#_description" class="anchor"></a>Description

The semantics of the basic blend operations are described in the table
below:

<table class="tableblock frame-all grid-all stretch">
<caption>Table 1. Basic Blend Operations</caption>
<colgroup>
<col style="width: 45%" />
<col style="width: 30%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr class="header">
<th class="tableblock halign-left valign-top"><a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html">VkBlendOp</a></th>
<th class="tableblock halign-left valign-top">RGB Components</th>
<th class="tableblock halign-left valign-top">Alpha Component</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td
class="tableblock halign-left valign-top"><p><code>VK_BLEND_OP_ADD</code></p></td>
<td class="tableblock halign-left valign-top"><p>R = R<sub>s0</sub> ×
S<sub>r</sub> + R<sub>d</sub> × D<sub>r</sub><br />
G = G<sub>s0</sub> × S<sub>g</sub> + G<sub>d</sub> × D<sub>g</sub><br />
B = B<sub>s0</sub> × S<sub>b</sub> + B<sub>d</sub> ×
D<sub>b</sub></p></td>
<td class="tableblock halign-left valign-top"><p>A = A<sub>s0</sub> ×
S<sub>a</sub> + A<sub>d</sub> × D<sub>a</sub></p></td>
</tr>
<tr class="even">
<td
class="tableblock halign-left valign-top"><p><code>VK_BLEND_OP_SUBTRACT</code></p></td>
<td class="tableblock halign-left valign-top"><p>R = R<sub>s0</sub> ×
S<sub>r</sub> - R<sub>d</sub> × D<sub>r</sub><br />
G = G<sub>s0</sub> × S<sub>g</sub> - G<sub>d</sub> × D<sub>g</sub><br />
B = B<sub>s0</sub> × S<sub>b</sub> - B<sub>d</sub> ×
D<sub>b</sub></p></td>
<td class="tableblock halign-left valign-top"><p>A = A<sub>s0</sub> ×
S<sub>a</sub> - A<sub>d</sub> × D<sub>a</sub></p></td>
</tr>
<tr class="odd">
<td
class="tableblock halign-left valign-top"><p><code>VK_BLEND_OP_REVERSE_SUBTRACT</code></p></td>
<td class="tableblock halign-left valign-top"><p>R = R<sub>d</sub> ×
D<sub>r</sub> - R<sub>s0</sub> × S<sub>r</sub><br />
G = G<sub>d</sub> × D<sub>g</sub> - G<sub>s0</sub> × S<sub>g</sub><br />
B = B<sub>d</sub> × D<sub>b</sub> - B<sub>s0</sub> ×
S<sub>b</sub></p></td>
<td class="tableblock halign-left valign-top"><p>A = A<sub>d</sub> ×
D<sub>a</sub> - A<sub>s0</sub> × S<sub>a</sub></p></td>
</tr>
<tr class="even">
<td
class="tableblock halign-left valign-top"><p><code>VK_BLEND_OP_MIN</code></p></td>
<td class="tableblock halign-left valign-top"><p>R =
min(R<sub>s0</sub>,R<sub>d</sub>)<br />
G = min(G<sub>s0</sub>,G<sub>d</sub>)<br />
B = min(B<sub>s0</sub>,B<sub>d</sub>)</p></td>
<td class="tableblock halign-left valign-top"><p>A =
min(A<sub>s0</sub>,A<sub>d</sub>)</p></td>
</tr>
<tr class="odd">
<td
class="tableblock halign-left valign-top"><p><code>VK_BLEND_OP_MAX</code></p></td>
<td class="tableblock halign-left valign-top"><p>R =
max(R<sub>s0</sub>,R<sub>d</sub>)<br />
G = max(G<sub>s0</sub>,G<sub>d</sub>)<br />
B = max(B<sub>s0</sub>,B<sub>d</sub>)</p></td>
<td class="tableblock halign-left valign-top"><p>A =
max(A<sub>s0</sub>,A<sub>d</sub>)</p></td>
</tr>
</tbody>
</table>

Table 1. Basic Blend Operations

In this table, the following conventions are used:

- R<sub>s0</sub>, G<sub>s0</sub>, B<sub>s0</sub> and A<sub>s0</sub>
  represent the first source color R, G, B, and A components,
  respectively.

- R<sub>d</sub>, G<sub>d</sub>, B<sub>d</sub> and A<sub>d</sub>
  represent the R, G, B, and A components of the destination color. That
  is, the color currently in the corresponding color attachment for this
  fragment/sample.

- S<sub>r</sub>, S<sub>g</sub>, S<sub>b</sub> and S<sub>a</sub>
  represent the source blend factor R, G, B, and A components,
  respectively.

- D<sub>r</sub>, D<sub>g</sub>, D<sub>b</sub> and D<sub>a</sub>
  represent the destination blend factor R, G, B, and A components,
  respectively.

The blending operation produces a new set of values R, G, B and A, which
are written to the framebuffer attachment. If blending is not enabled
for this attachment, then R, G, B and A are assigned R<sub>s0</sub>,
G<sub>s0</sub>, B<sub>s0</sub> and A<sub>s0</sub>, respectively.

If the color attachment is fixed-point, the components of the source and
destination values and blend factors are each clamped to \[0,1\] or
\[-1,1\] respectively for an unsigned normalized or signed normalized
color attachment prior to evaluating the blend operations. If the color
attachment is floating-point, no clamping occurs.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendAdvancedEXT.html),
[VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendEquationEXT.html),
[VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBlendOp"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
