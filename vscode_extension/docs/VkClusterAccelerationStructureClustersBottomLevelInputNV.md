# VkClusterAccelerationStructureClustersBottomLevelInputNV(3) Manual Page

## Name

VkClusterAccelerationStructureClustersBottomLevelInputNV - Parameters describing bottom level acceleration structure



## [](#_c_specification)C Specification

The [VkClusterAccelerationStructureClustersBottomLevelInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClustersBottomLevelInputNV.html) structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkClusterAccelerationStructureClustersBottomLevelInputNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxTotalClusterCount;
    uint32_t           maxClusterCountPerAccelerationStructure;
} VkClusterAccelerationStructureClustersBottomLevelInputNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxTotalClusterCount` is the total number of clusters acceleration structures that will be built or moved across all input arguments.
- `maxClusterCountPerAccelerationStructure` is the maximum number of clusters acceleration structures that will be built or moved per input argument.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkClusterAccelerationStructureClustersBottomLevelInputNV-sType-sType)VUID-VkClusterAccelerationStructureClustersBottomLevelInputNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CLUSTER_ACCELERATION_STRUCTURE_CLUSTERS_BOTTOM_LEVEL_INPUT_NV`
- [](#VUID-VkClusterAccelerationStructureClustersBottomLevelInputNV-pNext-pNext)VUID-VkClusterAccelerationStructureClustersBottomLevelInputNV-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkClusterAccelerationStructureOpInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpInputNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureClustersBottomLevelInputNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0