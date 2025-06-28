# VkWriteDescriptorSetAccelerationStructureNV(3) Manual Page

## Name

VkWriteDescriptorSetAccelerationStructureNV - Structure specifying acceleration structure descriptor information



## [](#_c_specification)C Specification

The `VkWriteDescriptorSetAccelerationStructureNV` structure is defined as:

```c++
// Provided by VK_NV_ray_tracing
typedef struct VkWriteDescriptorSetAccelerationStructureNV {
    VkStructureType                     sType;
    const void*                         pNext;
    uint32_t                            accelerationStructureCount;
    const VkAccelerationStructureNV*    pAccelerationStructures;
} VkWriteDescriptorSetAccelerationStructureNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `accelerationStructureCount` is the number of elements in `pAccelerationStructures`.
- `pAccelerationStructures` is a pointer to an array of [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) structures specifying the acceleration structures to update.

## [](#_description)Description

Valid Usage

- [](#VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-03748)VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-03748  
  Each acceleration structure in `pAccelerationStructures` **must** have been created with `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR`
- [](#VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-03749)VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-03749  
  If the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, each member of `pAccelerationStructures` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)

Valid Usage (Implicit)

- [](#VUID-VkWriteDescriptorSetAccelerationStructureNV-sType-sType)VUID-VkWriteDescriptorSetAccelerationStructureNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_ACCELERATION_STRUCTURE_NV`
- [](#VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-parameter)VUID-VkWriteDescriptorSetAccelerationStructureNV-pAccelerationStructures-parameter  
  `pAccelerationStructures` **must** be a valid pointer to an array of `accelerationStructureCount` valid or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) handles
- [](#VUID-VkWriteDescriptorSetAccelerationStructureNV-accelerationStructureCount-arraylength)VUID-VkWriteDescriptorSetAccelerationStructureNV-accelerationStructureCount-arraylength  
  `accelerationStructureCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkWriteDescriptorSetAccelerationStructureNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0