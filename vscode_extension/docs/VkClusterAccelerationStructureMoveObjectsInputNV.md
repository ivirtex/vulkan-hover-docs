# VkClusterAccelerationStructureMoveObjectsInputNV(3) Manual Page

## Name

VkClusterAccelerationStructureMoveObjectsInputNV - Parameters describing move information for an acceleration structure



## [](#_c_specification)C Specification

The [VkClusterAccelerationStructureMoveObjectsInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureMoveObjectsInputNV.html) structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkClusterAccelerationStructureMoveObjectsInputNV {
    VkStructureType                         sType;
    void*                                   pNext;
    VkClusterAccelerationStructureTypeNV    type;
    VkBool32                                noMoveOverlap;
    VkDeviceSize                            maxMovedBytes;
} VkClusterAccelerationStructureMoveObjectsInputNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `type` is a [VkClusterAccelerationStructureTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTypeNV.html) value identifying the type of cluster acceleration structure.
- `noMoveOverlap` specifies if the source and destination cluster acceleration structures overlap in memory for the move operation. If set to `VK_TRUE`, the source cluster acceleration structure remains valid after the move and move operation acts like a copy.
- `maxMovedBytes` is the maximum number of bytes that **may** be moved in this operation.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkClusterAccelerationStructureMoveObjectsInputNV-sType-sType)VUID-VkClusterAccelerationStructureMoveObjectsInputNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CLUSTER_ACCELERATION_STRUCTURE_MOVE_OBJECTS_INPUT_NV`
- [](#VUID-VkClusterAccelerationStructureMoveObjectsInputNV-pNext-pNext)VUID-VkClusterAccelerationStructureMoveObjectsInputNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkClusterAccelerationStructureMoveObjectsInputNV-type-parameter)VUID-VkClusterAccelerationStructureMoveObjectsInputNV-type-parameter  
  `type` **must** be a valid [VkClusterAccelerationStructureTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTypeNV.html) value

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkClusterAccelerationStructureOpInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpInputNV.html), [VkClusterAccelerationStructureTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTypeNV.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureMoveObjectsInputNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0