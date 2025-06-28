# VkPartitionedAccelerationStructureFlagsNV(3) Manual Page

## Name

VkPartitionedAccelerationStructureFlagsNV - Structure describing additional flags for PTLAS



## [](#_c_specification)C Specification

The [VkPartitionedAccelerationStructureFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureFlagsNV.html) structure is defined as:

```c++
// Provided by VK_NV_partitioned_acceleration_structure
typedef struct VkPartitionedAccelerationStructureFlagsNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           enablePartitionTranslation;
} VkPartitionedAccelerationStructureFlagsNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `enablePartitionTranslation` specifies if a [partition translation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ptlas-partition-translation) **may** be applied with [VkPartitionedAccelerationStructureWritePartitionTranslationDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureWritePartitionTranslationDataNV.html).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPartitionedAccelerationStructureFlagsNV-sType-sType)VUID-VkPartitionedAccelerationStructureFlagsNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PARTITIONED_ACCELERATION_STRUCTURE_FLAGS_NV`

## [](#_see_also)See Also

[VK\_NV\_partitioned\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_partitioned_acceleration_structure.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPartitionedAccelerationStructureFlagsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0