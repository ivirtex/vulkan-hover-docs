# VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX(3) Manual Page

## Name

VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX - Structure specifying acceleration structure DGF compressed triangle data



## [](#_c_specification)C Specification

If a `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure is included in the `pNext` chain of a [VkAccelerationStructureGeometryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryKHR.html) structure whose `geometryType` member is `VK_GEOMETRY_TYPE_DENSE_GEOMETRY_FORMAT_TRIANGLES_AMDX`, then that structure defines triangle geometry using compressed data.

The `VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX` structure is defined as:

```c++
// Provided by VK_AMDX_dense_geometry_format
typedef struct VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX {
    VkStructureType                   sType;
    const void*                       pNext;
    VkDeviceOrHostAddressConstKHR     compressedData;
    VkDeviceSize                      dataSize;
    uint32_t                          numTriangles;
    uint32_t                          numVertices;
    uint32_t                          maxPrimitiveIndex;
    uint32_t                          maxGeometryIndex;
    VkCompressedTriangleFormatAMDX    format;
} VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `compressedData` specifies the base address of the compressed data.
- `dataSize` specifies the size of the compressed data.
- `numTriangles` specifies the total number of triangles encoded in the compressed data.
- `numVertices` specifies the number of vertices in the compressed data.
- `maxPrimitiveIndex` specifies the maximum primitive index encoded in the compressed data.
- `maxGeometryIndex` specifies the maximum geometry index encoded in the compressed data.
- `format` specifies the [VkCompressedTriangleFormatAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompressedTriangleFormatAMDX.html) format of the compressed data.

## [](#_description)Description

If `format` is `VK_COMPRESSED_TRIANGLE_FORMAT_DGF1_AMDX`, `numVertices` specifies the sum of vertex counts across all blocks.

Valid Usage

- [](#VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-compressedData-10885)VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-compressedData-10885  
  The buffer from which `compressedData.deviceAddress` is queried **must** have been created with the `VK_BUFFER_USAGE_2_COMPRESSED_DATA_DGF1_BIT_AMDX` usage flag
- [](#VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-denseGeometryFormat-10886)VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-denseGeometryFormat-10886  
  The [`VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX`::`denseGeometryFormat`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-denseGeometryFormat) feature **must** be enabled
- [](#VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-format-10887)VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-format-10887  
  If `format` is VK\_COMPRESSED\_TRIANGLE\_FORMAT\_DGF1\_AMDX, then `compressedData.address` **must** be aligned to 128 bytes
- [](#VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-format-10888)VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-format-10888  
  If `format` is VK\_COMPRESSED\_TRIANGLE\_FORMAT\_DGF1\_AMDX, then `dataSize` **must** be aligned to 128 bytes
- [](#VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-pNext-10890)VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-pNext-10890  
  `pNext` **must** be `NULL` or a pointer to a valid [VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html) structure
- [](#VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-pNext-10891)VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-pNext-10891  
  If `pNext` is a pointer to a valid [VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html) structure, the [`micromap`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-micromap) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-sType-sType)VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_DENSE_GEOMETRY_FORMAT_TRIANGLES_DATA_AMDX`
- [](#VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-compressedData-parameter)VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-compressedData-parameter  
  `compressedData` **must** be a valid [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html) union
- [](#VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-format-parameter)VUID-VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX-format-parameter  
  `format` **must** be a valid [VkCompressedTriangleFormatAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompressedTriangleFormatAMDX.html) value

## [](#_see_also)See Also

[VK\_AMDX\_dense\_geometry\_format](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_dense_geometry_format.html), [VkCompressedTriangleFormatAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompressedTriangleFormatAMDX.html), [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0