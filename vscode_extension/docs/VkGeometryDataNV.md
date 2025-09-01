# VkGeometryDataNV(3) Manual Page

## Name

VkGeometryDataNV - Structure specifying geometry in a bottom-level acceleration structure



## [](#_c_specification)C Specification

The `VkGeometryDataNV` structure specifies geometry in a bottom-level acceleration structure and is defined as:

```c++
// Provided by VK_NV_ray_tracing
typedef struct VkGeometryDataNV {
    VkGeometryTrianglesNV    triangles;
    VkGeometryAABBNV         aabbs;
} VkGeometryDataNV;
```

## [](#_members)Members

- `triangles` contains triangle data if [VkGeometryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryNV.html)::`geometryType` is `VK_GEOMETRY_TYPE_TRIANGLES_NV`.
- `aabbs` contains axis-aligned bounding box data if [VkGeometryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryNV.html)::`geometryType` is `VK_GEOMETRY_TYPE_AABBS_NV`.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkGeometryDataNV-triangles-parameter)VUID-VkGeometryDataNV-triangles-parameter  
  `triangles` **must** be a valid [VkGeometryTrianglesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTrianglesNV.html) structure
- [](#VUID-VkGeometryDataNV-aabbs-parameter)VUID-VkGeometryDataNV-aabbs-parameter  
  `aabbs` **must** be a valid [VkGeometryAABBNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryAABBNV.html) structure

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkGeometryAABBNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryAABBNV.html), [VkGeometryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryNV.html), [VkGeometryTrianglesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTrianglesNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGeometryDataNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0