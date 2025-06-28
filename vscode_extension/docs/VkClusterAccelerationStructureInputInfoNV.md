# VkClusterAccelerationStructureInputInfoNV(3) Manual Page

## Name

VkClusterAccelerationStructureInputInfoNV - Structure describing a cluster acceleration structure



## [](#_c_specification)C Specification

The [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkClusterAccelerationStructureInputInfoNV {
    VkStructureType                            sType;
    void*                                      pNext;
    uint32_t                                   maxAccelerationStructureCount;
    VkBuildAccelerationStructureFlagsKHR       flags;
    VkClusterAccelerationStructureOpTypeNV     opType;
    VkClusterAccelerationStructureOpModeNV     opMode;
    VkClusterAccelerationStructureOpInputNV    opInput;
} VkClusterAccelerationStructureInputInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxAccelerationStructureCount` is the maximum number of acceleration structures that will be provided to the multi indirect operation.
- `flags` is a bitmask of [VkBuildAccelerationStructureFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagsKHR.html) specifying flags for the multi indirect operation.
- `opType` is a [VkClusterAccelerationStructureOpTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpTypeNV.html) value specifying the type of operation to perform.
- `opMode` is a [VkClusterAccelerationStructureOpModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpModeNV.html) value specifying the mode of operation.
- `opInput` is a [VkClusterAccelerationStructureOpInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpInputNV.html) value specifying the descriptions of the operation.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkClusterAccelerationStructureInputInfoNV-sType-sType)VUID-VkClusterAccelerationStructureInputInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CLUSTER_ACCELERATION_STRUCTURE_INPUT_INFO_NV`
- [](#VUID-VkClusterAccelerationStructureInputInfoNV-pNext-pNext)VUID-VkClusterAccelerationStructureInputInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkClusterAccelerationStructureInputInfoNV-flags-parameter)VUID-VkClusterAccelerationStructureInputInfoNV-flags-parameter  
  `flags` **must** be a valid combination of [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagBitsKHR.html) values
- [](#VUID-VkClusterAccelerationStructureInputInfoNV-opType-parameter)VUID-VkClusterAccelerationStructureInputInfoNV-opType-parameter  
  `opType` **must** be a valid [VkClusterAccelerationStructureOpTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpTypeNV.html) value
- [](#VUID-VkClusterAccelerationStructureInputInfoNV-opMode-parameter)VUID-VkClusterAccelerationStructureInputInfoNV-opMode-parameter  
  `opMode` **must** be a valid [VkClusterAccelerationStructureOpModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpModeNV.html) value
- [](#VUID-VkClusterAccelerationStructureInputInfoNV-pClustersBottomLevel-parameter)VUID-VkClusterAccelerationStructureInputInfoNV-pClustersBottomLevel-parameter  
  If `opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_CLUSTERS_BOTTOM_LEVEL_NV`, the `pClustersBottomLevel` member of `opInput` **must** be a valid pointer to a [VkClusterAccelerationStructureClustersBottomLevelInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClustersBottomLevelInputNV.html) structure
- [](#VUID-VkClusterAccelerationStructureInputInfoNV-pTriangleClusters-parameter)VUID-VkClusterAccelerationStructureInputInfoNV-pTriangleClusters-parameter  
  If `opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_TRIANGLE_CLUSTER_NV`, `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_TRIANGLE_CLUSTER_TEMPLATE_NV`, `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_INSTANTIATE_TRIANGLE_CLUSTER_NV`, or `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_GET_CLUSTER_TEMPLATE_INDICES_NV`, the `pTriangleClusters` member of `opInput` **must** be a valid pointer to a [VkClusterAccelerationStructureTriangleClusterInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTriangleClusterInputNV.html) structure
- [](#VUID-VkClusterAccelerationStructureInputInfoNV-pMoveObjects-parameter)VUID-VkClusterAccelerationStructureInputInfoNV-pMoveObjects-parameter  
  If `opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_MOVE_OBJECTS_NV`, the `pMoveObjects` member of `opInput` **must** be a valid pointer to a [VkClusterAccelerationStructureMoveObjectsInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureMoveObjectsInputNV.html) structure

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkBuildAccelerationStructureFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagsKHR.html), [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html), [VkClusterAccelerationStructureOpInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpInputNV.html), [VkClusterAccelerationStructureOpModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpModeNV.html), [VkClusterAccelerationStructureOpTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpTypeNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetClusterAccelerationStructureBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetClusterAccelerationStructureBuildSizesNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureInputInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0