# VkWriteDescriptorSetAccelerationStructureKHR(3) Manual Page

## Name

VkWriteDescriptorSetAccelerationStructureKHR - Structure specifying acceleration structure descriptor information



## [](#_c_specification)C Specification

The `VkWriteDescriptorSetAccelerationStructureKHR` structure is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkWriteDescriptorSetAccelerationStructureKHR {
    VkStructureType                      sType;
    const void*                          pNext;
    uint32_t                             accelerationStructureCount;
    const VkAccelerationStructureKHR*    pAccelerationStructures;
} VkWriteDescriptorSetAccelerationStructureKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `accelerationStructureCount` is the number of elements in `pAccelerationStructures`.
- `pAccelerationStructures` is a pointer to an array of [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) structures specifying the acceleration structures to update.

## [](#_description)Description

Valid Usage

- [](#VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-03579)VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-03579  
  Each acceleration structure in `pAccelerationStructures` **must** have been created with a `type` of `VK_ACCELERATION_STRUCTURE_TYPE_TOP_LEVEL_KHR` or `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`
- [](#VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-03580)VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-03580  
  If the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, each element of `pAccelerationStructures` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)

Valid Usage (Implicit)

- [](#VUID-VkWriteDescriptorSetAccelerationStructureKHR-sType-sType)VUID-VkWriteDescriptorSetAccelerationStructureKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_ACCELERATION_STRUCTURE_KHR`
- [](#VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-parameter)VUID-VkWriteDescriptorSetAccelerationStructureKHR-pAccelerationStructures-parameter  
  `pAccelerationStructures` **must** be a valid pointer to an array of `accelerationStructureCount` valid or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handles
- [](#VUID-VkWriteDescriptorSetAccelerationStructureKHR-accelerationStructureCount-arraylength)VUID-VkWriteDescriptorSetAccelerationStructureKHR-accelerationStructureCount-arraylength  
  `accelerationStructureCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkWriteDescriptorSetAccelerationStructureKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0