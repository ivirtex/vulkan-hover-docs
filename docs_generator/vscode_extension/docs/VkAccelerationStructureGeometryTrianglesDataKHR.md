# VkAccelerationStructureGeometryTrianglesDataKHR(3) Manual Page

## Name

VkAccelerationStructureGeometryTrianglesDataKHR - Structure specifying a triangle geometry in a bottom-level acceleration structure



## [](#_c_specification)C Specification

The `VkAccelerationStructureGeometryTrianglesDataKHR` structure is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureGeometryTrianglesDataKHR {
    VkStructureType                  sType;
    const void*                      pNext;
    VkFormat                         vertexFormat;
    VkDeviceOrHostAddressConstKHR    vertexData;
    VkDeviceSize                     vertexStride;
    uint32_t                         maxVertex;
    VkIndexType                      indexType;
    VkDeviceOrHostAddressConstKHR    indexData;
    VkDeviceOrHostAddressConstKHR    transformData;
} VkAccelerationStructureGeometryTrianglesDataKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `vertexFormat` is the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each vertex element.
- `vertexData` is a device or host address to memory containing vertex data for this geometry.
- `vertexStride` is the stride in bytes between each vertex.
- `maxVertex` is the number of vertices in `vertexData` minus one.
- `indexType` is the [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) of each index element.
- `indexData` is a device or host address to memory containing index data for this geometry.
- `transformData` is a device or host address to memory containing an optional reference to a [VkTransformMatrixKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTransformMatrixKHR.html) structure describing a transformation from the space in which the vertices in this geometry are described to the space in which the acceleration structure is defined.

## [](#_description)Description

Note

Unlike the stride for vertex buffers in [VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription.html) for graphics pipelines which must not exceed `maxVertexInputBindingStride`, `vertexStride` for acceleration structure geometry is instead restricted to being a 32-bit value.

Valid Usage

- [](#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexStride-03735)VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexStride-03735  
  `vertexStride` **must** be a multiple of the size in bytes of the smallest component of `vertexFormat`
- [](#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexStride-03819)VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexStride-03819  
  `vertexStride` **must** be less than or equal to 232-1
- [](#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexFormat-03797)VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexFormat-03797  
  The [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-view-format-features) of `vertexFormat` **must** contain `VK_FORMAT_FEATURE_ACCELERATION_STRUCTURE_VERTEX_BUFFER_BIT_KHR`
- [](#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-indexType-03798)VUID-VkAccelerationStructureGeometryTrianglesDataKHR-indexType-03798  
  `indexType` **must** be `VK_INDEX_TYPE_UINT16`, `VK_INDEX_TYPE_UINT32`, or `VK_INDEX_TYPE_NONE_KHR`

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-sType-sType)VUID-VkAccelerationStructureGeometryTrianglesDataKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_TRIANGLES_DATA_KHR`
- [](#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-pNext-pNext)VUID-VkAccelerationStructureGeometryTrianglesDataKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkAccelerationStructureGeometryMotionTrianglesDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryMotionTrianglesDataNV.html), [VkAccelerationStructureTrianglesDisplacementMicromapNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesDisplacementMicromapNV.html), or [VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html)
- [](#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-sType-unique)VUID-VkAccelerationStructureGeometryTrianglesDataKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexFormat-parameter)VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexFormat-parameter  
  `vertexFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-indexType-parameter)VUID-VkAccelerationStructureGeometryTrianglesDataKHR-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureGeometryDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryDataKHR.html), [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureGeometryTrianglesDataKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0