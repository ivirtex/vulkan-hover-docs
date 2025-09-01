# VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV(3) Manual Page

## Name

VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV - Parameters describing build operation for a bottom level cluster acceleration structure



## [](#_c_specification)C Specification

The [VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV {
    uint32_t           clusterReferencesCount;
    uint32_t           clusterReferencesStride;
    VkDeviceAddress    clusterReferences;
} VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV;
```

## [](#_members)Members

- `clusterReferencesCount` is the number of clusters this bottom level acceleration structure will be built from.
- `clusterReferencesStride` is the stride in `clusterReferences`.
- `clusterReferences` is the device memory containing the address of the clusters.

## [](#_description)Description

Valid Usage

- [](#VUID-VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV-clusterReferences-10484)VUID-VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV-clusterReferences-10484  
  All cluster references in `clusterReferences` **must** be unique
- [](#VUID-VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV-clusterReferences-10485)VUID-VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV-clusterReferences-10485  
  `clusterReferences` **must** have at least `clusterReferencesCount` values
- [](#VUID-VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV-clusterReferencesStride-10486)VUID-VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV-clusterReferencesStride-10486  
  `clusterReferencesStride` **must** be greater than or equal to 8

Valid Usage (Implicit)

- [](#VUID-VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV-clusterReferences-parameter)VUID-VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV-clusterReferences-parameter  
  `clusterReferences` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0