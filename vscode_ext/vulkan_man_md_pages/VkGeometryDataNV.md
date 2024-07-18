# VkGeometryDataNV(3) Manual Page

## Name

VkGeometryDataNV - Structure specifying geometry in a bottom-level
acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkGeometryDataNV` structure specifies geometry in a bottom-level
acceleration structure and is defined as:

``` c
// Provided by VK_NV_ray_tracing
typedef struct VkGeometryDataNV {
    VkGeometryTrianglesNV    triangles;
    VkGeometryAABBNV         aabbs;
} VkGeometryDataNV;
```

## <a href="#_members" class="anchor"></a>Members

- `triangles` contains triangle data if
  [VkGeometryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryNV.html)::`geometryType` is
  `VK_GEOMETRY_TYPE_TRIANGLES_NV`.

- `aabbs` contains axis-aligned bounding box data if
  [VkGeometryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryNV.html)::`geometryType` is
  `VK_GEOMETRY_TYPE_AABBS_NV`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkGeometryDataNV-triangles-parameter"
  id="VUID-VkGeometryDataNV-triangles-parameter"></a>
  VUID-VkGeometryDataNV-triangles-parameter  
  `triangles` **must** be a valid
  [VkGeometryTrianglesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryTrianglesNV.html) structure

- <a href="#VUID-VkGeometryDataNV-aabbs-parameter"
  id="VUID-VkGeometryDataNV-aabbs-parameter"></a>
  VUID-VkGeometryDataNV-aabbs-parameter  
  `aabbs` **must** be a valid [VkGeometryAABBNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryAABBNV.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkGeometryAABBNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryAABBNV.html),
[VkGeometryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryNV.html),
[VkGeometryTrianglesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryTrianglesNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGeometryDataNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
