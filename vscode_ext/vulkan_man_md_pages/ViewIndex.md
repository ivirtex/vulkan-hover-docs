# ViewIndex(3) Manual Page

## Name

ViewIndex - View index of a shader invocation



## <a href="#_description" class="anchor"></a>Description

`ViewIndex`  
The `ViewIndex` decoration **can** be applied to a shader input which
will be filled with the index of the view that is being processed by the
current shader invocation.

If multiview is enabled in the render pass, this value will be one of
the bits set in the view mask of the subpass the pipeline is compiled
against. If multiview is not enabled in the render pass, this value will
be zero.

Valid Usage

- <a href="#VUID-ViewIndex-ViewIndex-04401"
  id="VUID-ViewIndex-ViewIndex-04401"></a>
  VUID-ViewIndex-ViewIndex-04401  
  The `ViewIndex` decoration **must** be used only within the `MeshEXT`,
  `Vertex`, `Geometry`, `TessellationControl`, `TessellationEvaluation`
  or `Fragment` `Execution` `Model`

- <a href="#VUID-ViewIndex-ViewIndex-04402"
  id="VUID-ViewIndex-ViewIndex-04402"></a>
  VUID-ViewIndex-ViewIndex-04402  
  The variable decorated with `ViewIndex` **must** be declared using the
  `Input` `Storage` `Class`

- <a href="#VUID-ViewIndex-ViewIndex-04403"
  id="VUID-ViewIndex-ViewIndex-04403"></a>
  VUID-ViewIndex-ViewIndex-04403  
  The variable decorated with `ViewIndex` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ViewIndex"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
