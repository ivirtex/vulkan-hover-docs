# PrimitiveShadingRateKHR(3) Manual Page

## Name

PrimitiveShadingRateKHR - Primitive contribution to fragment shading
rate



## <a href="#_description" class="anchor"></a>Description

`PrimitiveShadingRateKHR`  
Decorating a variable with the `PrimitiveShadingRateKHR` built-in
decoration will make that variable contain the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-primitive"
target="_blank" rel="noopener">primitive fragment shading rate</a>.

The value written to the variable decorated with
`PrimitiveShadingRateKHR` by the last <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> in the
pipeline is used as the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-primitive"
target="_blank" rel="noopener">primitive fragment shading rate</a>.
Outputs in previous shader stages are ignored.

If the last active <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> shader
entry point’s interface does not include a variable decorated with
`PrimitiveShadingRateKHR`, then it is as if the shader specified a
fragment shading rate value of 0, indicating a horizontal and vertical
rate of 1 pixel.

If a shader has `PrimitiveShadingRateKHR` in the output interface and
there is an execution path through the shader that does not write to it,
its value is undefined for executions of the shader that take that path.

Valid Usage

- <a href="#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04484"
  id="VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04484"></a>
  VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04484  
  The `PrimitiveShadingRateKHR` decoration **must** be used only within
  the `MeshEXT`, `MeshNV`, `Vertex`, or `Geometry` `Execution` `Model`

- <a href="#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04485"
  id="VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04485"></a>
  VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04485  
  The variable decorated with `PrimitiveShadingRateKHR` **must** be
  declared using the `Output` `Storage` `Class`

- <a href="#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04486"
  id="VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04486"></a>
  VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04486  
  The variable decorated with `PrimitiveShadingRateKHR` **must** be
  declared as a scalar 32-bit integer value

- <a href="#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04487"
  id="VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04487"></a>
  VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04487  
  The value written to `PrimitiveShadingRateKHR` **must** include no
  more than one of `Vertical2Pixels` and `Vertical4Pixels`

- <a href="#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04488"
  id="VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04488"></a>
  VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04488  
  The value written to `PrimitiveShadingRateKHR` **must** include no
  more than one of `Horizontal2Pixels` and `Horizontal4Pixels`

- <a href="#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04489"
  id="VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04489"></a>
  VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04489  
  The value written to `PrimitiveShadingRateKHR` **must** not have any
  bits set other than those defined by **Fragment Shading Rate Flags**
  enumerants in the SPIR-V specification

- <a href="#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-07059"
  id="VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-07059"></a>
  VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-07059  
  The variable decorated with `PrimitiveShadingRateKHR` within the
  `MeshEXT` `Execution` `Model` **must** also be decorated with the
  `PerPrimitiveEXT` decoration

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PrimitiveShadingRateKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
