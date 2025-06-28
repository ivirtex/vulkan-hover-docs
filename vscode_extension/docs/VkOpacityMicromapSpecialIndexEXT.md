# VkOpacityMicromapSpecialIndexEXT(3) Manual Page

## Name

VkOpacityMicromapSpecialIndexEXT - Enum for special indices in the opacity micromap



## [](#_c_specification)C Specification

The `VkOpacityMicromapSpecialIndexEXT` enumeration is defined as:

```c++
// Provided by VK_EXT_opacity_micromap
typedef enum VkOpacityMicromapSpecialIndexEXT {
    VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_TRANSPARENT_EXT = -1,
    VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_OPAQUE_EXT = -2,
    VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_UNKNOWN_TRANSPARENT_EXT = -3,
    VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_UNKNOWN_OPAQUE_EXT = -4,
  // Provided by VK_EXT_opacity_micromap with VK_NV_cluster_acceleration_structure
    VK_OPACITY_MICROMAP_SPECIAL_INDEX_CLUSTER_GEOMETRY_DISABLE_OPACITY_MICROMAP_NV = -5,
} VkOpacityMicromapSpecialIndexEXT;
```

## [](#_description)Description

- `VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_TRANSPARENT_EXT` specifies that the entire triangle is fully transparent.
- `VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_OPAQUE_EXT` specifies that the entire triangle is fully opaque.
- `VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_UNKNOWN_TRANSPARENT_EXT` specifies that the entire triangle is unknown-transparent.
- `VK_OPACITY_MICROMAP_SPECIAL_INDEX_FULLY_UNKNOWN_OPAQUE_EXT` specifies that the entire triangle is unknown-opaque.
- `VK_OPACITY_MICROMAP_SPECIAL_INDEX_CLUSTER_GEOMETRY_DISABLE_OPACITY_MICROMAP_NV` specifies that [Opacity Micromap](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-opacity-micromap) will be disabled for this triangle and opacity value will be picked from [VkClusterAccelerationStructureBuildTriangleClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterInfoNV.html)::`baseGeometryIndexAndGeometryFlags` instead. Note that this special index is only valid for [Cluster Geometry](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#cluster-geometry).

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOpacityMicromapSpecialIndexEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0