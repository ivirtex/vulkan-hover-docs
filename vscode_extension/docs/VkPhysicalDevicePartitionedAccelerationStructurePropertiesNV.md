# VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV(3) Manual Page

## Name

VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV - Structure describing properties supported by a partitioned acceleration structure implementation



## [](#_c_specification)C Specification

The `VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_partitioned_acceleration_structure
typedef struct VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxPartitionCount;
} VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maxPartitionCount` indicates the maximum number of partitions allowed in a partitioned acceleration structure.

## [](#_description)Description

If the `VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV-sType-sType)VUID-VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PARTITIONED_ACCELERATION_STRUCTURE_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_partitioned\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_partitioned_acceleration_structure.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0