# ViewportMaskNV(3) Manual Page

## Name

ViewportMaskNV - Mask of the viewports used



## [](#_description)Description

`ViewportMaskNV`

Decorating a variable with the `ViewportMaskNV` built-in decoration will make that variable contain the viewport mask.

In a mesh, vertex, tessellation evaluation, or geometry shader, the variable decorated with `ViewportMaskNV` can be written to with the mask of which viewports the primitive produced by that shader will directed.

The `ViewportMaskNV` variable **must** be an array that has ⌈(`VkPhysicalDeviceLimits`::`maxViewports` / 32)⌉ elements. When a shader writes to this variable, bit B of element M controls whether a primitive is emitted to viewport 32 × M + B. The viewports indicated by the mask are used to select the viewport transform, scissor rectangle, and exclusive scissor rectangle that a primitive will be transformed by.

The last active [*pre-rasterization shader stage*](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization) (in pipeline order) controls the `ViewportMaskNV` that is used. Outputs in previous shader stages are not used, even if the last stage fails to write the `ViewportMaskNV`. When `ViewportMaskNV` is written by the final [pre-rasterization shader stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization), any variable decorated with `ViewportIndex` in the fragment shader will have the index of the viewport that was used in generating that fragment.

If a [pre-rasterization shader stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization) shader entry point’s interface includes a variable decorated with `ViewportMaskNV`, it **must** write the same value to `ViewportMaskNV` for all output vertices of a given primitive.

Valid Usage

- [](#VUID-ViewportMaskNV-ViewportMaskNV-04409)VUID-ViewportMaskNV-ViewportMaskNV-04409  
  The `ViewportMaskNV` decoration **must** be used only within the `Vertex`, `MeshNV`, `TessellationEvaluation`, or `Geometry` `Execution` `Model`
- [](#VUID-ViewportMaskNV-ViewportMaskNV-04410)VUID-ViewportMaskNV-ViewportMaskNV-04410  
  The variable decorated with `ViewportMaskNV` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-ViewportMaskNV-ViewportMaskNV-04411)VUID-ViewportMaskNV-ViewportMaskNV-04411  
  The variable decorated with `ViewportMaskNV` **must** be declared as an array of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ViewportMaskNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0