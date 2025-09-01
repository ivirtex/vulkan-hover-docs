# PointSize(3) Manual Page

## Name

PointSize - Size of a point primitive



## [](#_description)Description

`PointSize`

Decorating a variable with the `PointSize` built-in decoration will make that variable contain the size of point primitives or the final rasterization of polygons if [polygon mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-polygonmode) is `VK_POLYGON_MODE_POINT` when `VkPhysicalDeviceMaintenance5Properties`::`polygonModePointSize` is set to `VK_TRUE` . The value written to the variable decorated with `PointSize` by the last [pre-rasterization shader stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization) in the pipeline is used as the framebuffer-space size of points produced by rasterization. If the [`maintenance5`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance5) feature is enabled and a value is not written to a variable decorated with `PointSize`, a value of 1.0 is used as the size of points.

Note

When `PointSize` decorates a variable in the `Input` `Storage` `Class`, it contains the data written to the output variable decorated with `PointSize` from the previous shader stage.

Valid Usage

- [](#VUID-PointSize-PointSize-04314)VUID-PointSize-PointSize-04314  
  The `PointSize` decoration **must** be used only within the `MeshEXT`, `MeshNV`, `Vertex`, `TessellationControl`, `TessellationEvaluation`, or `Geometry` `Execution` `Model`
- [](#VUID-PointSize-PointSize-04315)VUID-PointSize-PointSize-04315  
  The variable decorated with `PointSize` within the `MeshEXT`, `MeshNV`, or `Vertex` `Execution` `Model` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-PointSize-PointSize-04316)VUID-PointSize-PointSize-04316  
  The variable decorated with `PointSize` within the `TessellationControl`, `TessellationEvaluation`, or `Geometry` `Execution` `Model` **must** not be declared using a `Storage` `Class` other than `Input` or `Output`
- [](#VUID-PointSize-PointSize-04317)VUID-PointSize-PointSize-04317  
  The variable decorated with `PointSize` **must** be declared as a scalar 32-bit floating-point value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PointSize).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0