# VkWriteDescriptorSetPartitionedAccelerationStructureNV(3) Manual Page

## Name

VkWriteDescriptorSetPartitionedAccelerationStructureNV - Structure specifying descriptor for PTLAS



## [](#_c_specification)C Specification

If the `descriptorType` member of [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html) is `VK_DESCRIPTOR_TYPE_PARTITIONED_ACCELERATION_STRUCTURE_NV`, then the data to write to the descriptor set is specified through a `VkWriteDescriptorSetPartitionedAccelerationStructureNV` structure included in the `pNext` chain of `VkWriteDescriptorSet`.

The `VkWriteDescriptorSetPartitionedAccelerationStructureNV` structure is defined as:

```c++
// Provided by VK_NV_partitioned_acceleration_structure
typedef struct VkWriteDescriptorSetPartitionedAccelerationStructureNV {
    VkStructureType           sType;
    void*                     pNext;
    uint32_t                  accelerationStructureCount;
    const VkDeviceAddress*    pAccelerationStructures;
} VkWriteDescriptorSetPartitionedAccelerationStructureNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `accelerationStructureCount` is the number of elements in `pAccelerationStructures`.
- `pAccelerationStructures` is a pointer to an array of `accelerationStructureCount` device addresses pointing to previously built PTLAS.

## [](#_description)Description

Valid Usage

- [](#VUID-VkWriteDescriptorSetPartitionedAccelerationStructureNV-accelerationStructureCount-10511)VUID-VkWriteDescriptorSetPartitionedAccelerationStructureNV-accelerationStructureCount-10511  
  `accelerationStructureCount` **must** be equal to `descriptorCount` in the extended structure
- [](#VUID-VkWriteDescriptorSetPartitionedAccelerationStructureNV-pAccelerationStructures-10512)VUID-VkWriteDescriptorSetPartitionedAccelerationStructureNV-pAccelerationStructures-10512  
  Each entry in `pAccelerationStructures` **must** be a valid address of a PTLAS

Valid Usage (Implicit)

- [](#VUID-VkWriteDescriptorSetPartitionedAccelerationStructureNV-sType-sType)VUID-VkWriteDescriptorSetPartitionedAccelerationStructureNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_PARTITIONED_ACCELERATION_STRUCTURE_NV`
- [](#VUID-VkWriteDescriptorSetPartitionedAccelerationStructureNV-pAccelerationStructures-parameter)VUID-VkWriteDescriptorSetPartitionedAccelerationStructureNV-pAccelerationStructures-parameter  
  `pAccelerationStructures` **must** be a valid pointer to an array of `accelerationStructureCount` or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) values
- [](#VUID-VkWriteDescriptorSetPartitionedAccelerationStructureNV-accelerationStructureCount-arraylength)VUID-VkWriteDescriptorSetPartitionedAccelerationStructureNV-accelerationStructureCount-arraylength  
  `accelerationStructureCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_NV\_partitioned\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_partitioned_acceleration_structure.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkWriteDescriptorSetPartitionedAccelerationStructureNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0