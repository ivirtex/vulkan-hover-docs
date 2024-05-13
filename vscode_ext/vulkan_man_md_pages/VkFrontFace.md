# VkFrontFace(3) Manual Page

## Name

VkFrontFace - Interpret polygon front-facing orientation



## <a href="#_c_specification" class="anchor"></a>C Specification

The first step of polygon rasterization is to determine whether the
triangle is *back-facing* or *front-facing*. This determination is made
based on the sign of the (clipped or unclipped) polygon’s area computed
in framebuffer coordinates. One way to compute this area is:

a=−21​i=0∑n−1​xfi​yfi⊕1​−xfi⊕1​yfi​

where xfi​ and yfi​ are the x and y framebuffer coordinates of the ith
vertex of the n-vertex polygon (vertices are numbered starting at zero
for the purposes of this computation) and i ⊕ 1 is (i + 1) mod n.

The interpretation of the sign of a is determined by the
[VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)::`frontFace`
property of the currently active pipeline. Possible values are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkFrontFace {
    VK_FRONT_FACE_COUNTER_CLOCKWISE = 0,
    VK_FRONT_FACE_CLOCKWISE = 1,
} VkFrontFace;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_FRONT_FACE_COUNTER_CLOCKWISE` specifies that a triangle with
  positive area is considered front-facing.

- `VK_FRONT_FACE_CLOCKWISE` specifies that a triangle with negative area
  is considered front-facing.

Any triangle which is not front-facing is back-facing, including
zero-area triangles.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html),
[vkCmdSetFrontFace](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFrontFace.html),
[vkCmdSetFrontFaceEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFrontFaceEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFrontFace"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
