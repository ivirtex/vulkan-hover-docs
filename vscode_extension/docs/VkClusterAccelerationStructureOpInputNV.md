# VkClusterAccelerationStructureOpInputNV(3) Manual Page

## Name

VkClusterAccelerationStructureOpInputNV - Union specifying cluster acceleration structure description



## [](#_c_specification)C Specification

The `VkClusterAccelerationStructureOpInputNV` union is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef union VkClusterAccelerationStructureOpInputNV {
    VkClusterAccelerationStructureClustersBottomLevelInputNV*    pClustersBottomLevel;
    VkClusterAccelerationStructureTriangleClusterInputNV*        pTriangleClusters;
    VkClusterAccelerationStructureMoveObjectsInputNV*            pMoveObjects;
} VkClusterAccelerationStructureOpInputNV;
```

## [](#_members)Members

- `pClustersBottomLevel` is a [VkClusterAccelerationStructureClustersBottomLevelInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClustersBottomLevelInputNV.html) structure specifying an upper threshold on parameters to build multiple bottom level acceleration structures from multiple cluster level acceleration structures.
- `pTriangleClusters` is a [VkClusterAccelerationStructureTriangleClusterInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTriangleClusterInputNV.html) structure specifying an upper threshold on parameters to build a regular or templated cluster acceleration structure.
- `pMoveObjects` is a [VkClusterAccelerationStructureMoveObjectsInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureMoveObjectsInputNV.html) structure specifying an upper threshold on the number of bytes moved and the type of acceleration structure being moved.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkClusterAccelerationStructureClustersBottomLevelInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClustersBottomLevelInputNV.html), [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html), [VkClusterAccelerationStructureMoveObjectsInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureMoveObjectsInputNV.html), [VkClusterAccelerationStructureTriangleClusterInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTriangleClusterInputNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureOpInputNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0