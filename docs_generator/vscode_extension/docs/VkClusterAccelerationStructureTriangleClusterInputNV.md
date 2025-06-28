# VkClusterAccelerationStructureTriangleClusterInputNV(3) Manual Page

## Name

VkClusterAccelerationStructureTriangleClusterInputNV - Parameters describing a cluster acceleration structure



## [](#_c_specification)C Specification

The [VkClusterAccelerationStructureTriangleClusterInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTriangleClusterInputNV.html) structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkClusterAccelerationStructureTriangleClusterInputNV {
    VkStructureType    sType;
    void*              pNext;
    VkFormat           vertexFormat;
    uint32_t           maxGeometryIndexValue;
    uint32_t           maxClusterUniqueGeometryCount;
    uint32_t           maxClusterTriangleCount;
    uint32_t           maxClusterVertexCount;
    uint32_t           maxTotalTriangleCount;
    uint32_t           maxTotalVertexCount;
    uint32_t           minPositionTruncateBitCount;
} VkClusterAccelerationStructureTriangleClusterInputNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `vertexFormat` is the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each vertex element.
- `maxGeometryIndexValue` is the maximum geometry index value for any constructed geometry.
- `maxClusterUniqueGeometryCount` is the maximum number of unique values of the geometry index for each cluster or cluster template.
- []()`maxClusterTriangleCount` is the maximum number of triangles in a cluster or cluster template.
- `maxClusterVertexCount` is the maximum number of unique vertices in the cluster’s index buffer.
- `maxTotalTriangleCount` is the sum of all triangles across all clusters or cluster templates.
- `maxTotalVertexCount` is the maximum number of vertices across all clusters or cluster templates.
- []()`minPositionTruncateBitCount` is the least value specified in cluster build in [VkClusterAccelerationStructureBuildTriangleClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterInfoNV.html)::`positionTruncateBitCount` or cluster template build in [VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV.html)::`positionTruncateBitCount`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkClusterAccelerationStructureTriangleClusterInputNV-vertexFormat-10439)VUID-VkClusterAccelerationStructureTriangleClusterInputNV-vertexFormat-10439  
  The [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-view-format-features) of `vertexFormat` **must** contain `VK_FORMAT_FEATURE_ACCELERATION_STRUCTURE_VERTEX_BUFFER_BIT_KHR`
- [](#VUID-VkClusterAccelerationStructureTriangleClusterInputNV-maxClusterTriangleCount-10440)VUID-VkClusterAccelerationStructureTriangleClusterInputNV-maxClusterTriangleCount-10440  
  `maxClusterTriangleCount` **must** be less than or equal to [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`maxTrianglesPerCluster`
- [](#VUID-VkClusterAccelerationStructureTriangleClusterInputNV-maxClusterVertexCount-10441)VUID-VkClusterAccelerationStructureTriangleClusterInputNV-maxClusterVertexCount-10441  
  `maxClusterVertexCount` **must** be less than or equal to [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`maxVerticesPerCluster`
- [](#VUID-VkClusterAccelerationStructureTriangleClusterInputNV-minPositionTruncateBitCount-10442)VUID-VkClusterAccelerationStructureTriangleClusterInputNV-minPositionTruncateBitCount-10442  
  `minPositionTruncateBitCount` **must** be less than or equal to `32`

Valid Usage (Implicit)

- [](#VUID-VkClusterAccelerationStructureTriangleClusterInputNV-sType-sType)VUID-VkClusterAccelerationStructureTriangleClusterInputNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CLUSTER_ACCELERATION_STRUCTURE_TRIANGLE_CLUSTER_INPUT_NV`
- [](#VUID-VkClusterAccelerationStructureTriangleClusterInputNV-pNext-pNext)VUID-VkClusterAccelerationStructureTriangleClusterInputNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkClusterAccelerationStructureTriangleClusterInputNV-vertexFormat-parameter)VUID-VkClusterAccelerationStructureTriangleClusterInputNV-vertexFormat-parameter  
  `vertexFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkClusterAccelerationStructureOpInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpInputNV.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureTriangleClusterInputNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0