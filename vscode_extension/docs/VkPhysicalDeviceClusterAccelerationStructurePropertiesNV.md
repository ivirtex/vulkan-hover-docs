# VkPhysicalDeviceClusterAccelerationStructurePropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceClusterAccelerationStructurePropertiesNV - Structure describing properties supported by a cluster acceleration structure implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceClusterAccelerationStructurePropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkPhysicalDeviceClusterAccelerationStructurePropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxVerticesPerCluster;
    uint32_t           maxTrianglesPerCluster;
    uint32_t           clusterScratchByteAlignment;
    uint32_t           clusterByteAlignment;
    uint32_t           clusterTemplateByteAlignment;
    uint32_t           clusterBottomLevelByteAlignment;
    uint32_t           clusterTemplateBoundsByteAlignment;
    uint32_t           maxClusterGeometryIndex;
} VkPhysicalDeviceClusterAccelerationStructurePropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maxVerticesPerCluster` indicates the maximum number of unique vertices that **can** be specified in the index buffer for a cluster.
- []()`maxTrianglesPerCluster` indicates the maximum number of triangles in a cluster.
- []()`clusterScratchByteAlignment` indicates the alignment required for scratch memory used in building or moving cluster acceleration structures.
- []()`clusterByteAlignment` indicates the alignment of buffers when building cluster acceleration structures.
- []()`clusterTemplateByteAlignment` indicates the alignment of buffers when building cluster templates.
- []()`clusterBottomLevelByteAlignment` indicates the alignment of buffers when building bottom level acceleration structures.
- []()`clusterTemplateBoundsByteAlignment` indicates the alignment of [VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV.html)::pname::instantiationBoundingBoxLimit.
- []()`maxClusterGeometryIndex` indicates the maximum geometry index possible for a triangle in an cluster acceleration structures.

## [](#_description)Description

If the `VkPhysicalDeviceClusterAccelerationStructurePropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceClusterAccelerationStructurePropertiesNV-sType-sType)VUID-VkPhysicalDeviceClusterAccelerationStructurePropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CLUSTER_ACCELERATION_STRUCTURE_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceClusterAccelerationStructurePropertiesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0