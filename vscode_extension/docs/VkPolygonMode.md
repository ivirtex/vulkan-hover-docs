# VkPolygonMode(3) Manual Page

## Name

VkPolygonMode - Control polygon rasterization mode



## [](#_c_specification)C Specification

Possible values of the [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html)::`polygonMode` property of the currently active pipeline, specifying the method of rasterization for polygons, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkPolygonMode {
    VK_POLYGON_MODE_FILL = 0,
    VK_POLYGON_MODE_LINE = 1,
    VK_POLYGON_MODE_POINT = 2,
  // Provided by VK_NV_fill_rectangle
    VK_POLYGON_MODE_FILL_RECTANGLE_NV = 1000153000,
} VkPolygonMode;
```

## [](#_description)Description

- `VK_POLYGON_MODE_POINT` specifies that polygon vertices are drawn as points.
- `VK_POLYGON_MODE_LINE` specifies that polygon edges are drawn as line segments.
- `VK_POLYGON_MODE_FILL` specifies that polygons are rendered using the polygon rasterization rules in this section.
- `VK_POLYGON_MODE_FILL_RECTANGLE_NV` specifies that polygons are rendered using polygon rasterization rules, modified to consider a sample within the primitive if the sample location is inside the axis-aligned bounding box of the triangle after projection. Note that the barycentric weights used in attribute interpolation **can** extend outside the range \[0,1] when these primitives are shaded. Special treatment is given to a sample position on the boundary edge of the bounding box. In such a case, if two rectangles lie on either side of a common edge (with identical endpoints) on which a sample position lies, then exactly one of the triangles **must** produce a fragment that covers that sample during rasterization.
  
  Polygons rendered in `VK_POLYGON_MODE_FILL_RECTANGLE_NV` mode **may** be clipped by the frustum or by user clip planes. If clipping is applied, the triangle is culled rather than clipped.
  
  Area calculation and facingness are determined for `VK_POLYGON_MODE_FILL_RECTANGLE_NV` mode using the triangle’s vertices.

These modes affect only the final rasterization of polygons: in particular, a polygon’s vertices are shaded and the polygon is clipped and possibly culled before these modes are applied.

If `VkPhysicalDeviceMaintenance5Properties`::`polygonModePointSize` is `VK_TRUE`, the point size of the final rasterization of polygons is taken from `PointSize` when [polygon mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-polygonmode) is `VK_POLYGON_MODE_POINT`.

Otherwise, if `VkPhysicalDeviceMaintenance5Properties`::`polygonModePointSize` is `VK_FALSE`, the point size of the final rasterization of polygons is 1.0 when [polygon mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-polygonmode) is `VK_POLYGON_MODE_POINT`.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html), [vkCmdSetPolygonModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPolygonModeEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPolygonMode)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0