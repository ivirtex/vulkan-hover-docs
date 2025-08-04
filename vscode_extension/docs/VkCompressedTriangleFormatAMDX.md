# VkCompressedTriangleFormatAMDX(3) Manual Page

## Name

VkCompressedTriangleFormatAMDX - Available compressed triangle formats



## [](#_c_specification)C Specification

The `VkCompressedTriangleFormatAMDX` enumeration is defined as:

```c++
// Provided by VK_AMDX_dense_geometry_format
typedef enum VkCompressedTriangleFormatAMDX {
    VK_COMPRESSED_TRIANGLE_FORMAT_DGF1_AMDX = 0,
} VkCompressedTriangleFormatAMDX;
```

## [](#_description)Description

- `VK_COMPRESSED_TRIANGLE_FORMAT_DGF1_AMDX` specifies that the compressed triangle data is in [Dense Geometry Format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dense-geometry-format), version 1, consisting of an array of 128B DGF blocks.

## [](#_see_also)See Also

[VK\_AMDX\_dense\_geometry\_format](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_dense_geometry_format.html), [VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCompressedTriangleFormatAMDX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0