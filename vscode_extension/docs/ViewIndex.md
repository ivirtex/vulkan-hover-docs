# ViewIndex(3) Manual Page

## Name

ViewIndex - View index of a shader invocation



## [](#_description)Description

`ViewIndex`

The `ViewIndex` decoration **can** be applied to a shader input which will be filled with the index of the view that is being processed by the current shader invocation.

If multiview is enabled in the render pass, this value will be the index of one of the bits set in the view mask of the subpass the pipeline is compiled against. If multiview is not enabled in the render pass, this value will be zero.

Valid Usage

- [](#VUID-ViewIndex-ViewIndex-04401)VUID-ViewIndex-ViewIndex-04401  
  The `ViewIndex` decoration **must** be used only within the `MeshEXT`, `Vertex`, `Geometry`, `TessellationControl`, `TessellationEvaluation` or `Fragment` `Execution` `Model`
- [](#VUID-ViewIndex-ViewIndex-04402)VUID-ViewIndex-ViewIndex-04402  
  The variable decorated with `ViewIndex` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-ViewIndex-ViewIndex-04403)VUID-ViewIndex-ViewIndex-04403  
  The variable decorated with `ViewIndex` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ViewIndex)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0