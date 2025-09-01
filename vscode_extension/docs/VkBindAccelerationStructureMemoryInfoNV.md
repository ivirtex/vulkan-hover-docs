# VkBindAccelerationStructureMemoryInfoNV(3) Manual Page

## Name

VkBindAccelerationStructureMemoryInfoNV - Structure specifying acceleration structure memory binding



## [](#_c_specification)C Specification

The `VkBindAccelerationStructureMemoryInfoNV` structure is defined as:

```c++
// Provided by VK_NV_ray_tracing
typedef struct VkBindAccelerationStructureMemoryInfoNV {
    VkStructureType              sType;
    const void*                  pNext;
    VkAccelerationStructureNV    accelerationStructure;
    VkDeviceMemory               memory;
    VkDeviceSize                 memoryOffset;
    uint32_t                     deviceIndexCount;
    const uint32_t*              pDeviceIndices;
} VkBindAccelerationStructureMemoryInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `accelerationStructure` is the acceleration structure to be attached to memory.
- `memory` is a `VkDeviceMemory` object describing the device memory to attach.
- `memoryOffset` is the start offset of the region of memory that is to be bound to the acceleration structure. The number of bytes returned in the [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html)::`size` member in `memory`, starting from `memoryOffset` bytes, will be bound to the specified acceleration structure.
- `deviceIndexCount` is the number of elements in `pDeviceIndices`.
- `pDeviceIndices` is a pointer to an array of device indices.

## [](#_description)Description

Valid Usage

- [](#VUID-VkBindAccelerationStructureMemoryInfoNV-accelerationStructure-03620)VUID-VkBindAccelerationStructureMemoryInfoNV-accelerationStructure-03620  
  `accelerationStructure` **must** not already be backed by a memory object
- [](#VUID-VkBindAccelerationStructureMemoryInfoNV-memoryOffset-03621)VUID-VkBindAccelerationStructureMemoryInfoNV-memoryOffset-03621  
  `memoryOffset` **must** be less than the size of `memory`
- [](#VUID-VkBindAccelerationStructureMemoryInfoNV-memory-03622)VUID-VkBindAccelerationStructureMemoryInfoNV-memory-03622  
  `memory` **must** have been allocated using one of the memory types allowed in the `memoryTypeBits` member of the [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure returned from a call to [vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html) with `accelerationStructure` and `type` of `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV`
- [](#VUID-VkBindAccelerationStructureMemoryInfoNV-memoryOffset-03623)VUID-VkBindAccelerationStructureMemoryInfoNV-memoryOffset-03623  
  `memoryOffset` **must** be an integer multiple of the `alignment` member of the [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure returned from a call to [vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html) with `accelerationStructure` and `type` of `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV`
- [](#VUID-VkBindAccelerationStructureMemoryInfoNV-size-03624)VUID-VkBindAccelerationStructureMemoryInfoNV-size-03624  
  The `size` member of the `VkMemoryRequirements` structure returned from a call to [vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html) with `accelerationStructure` and `type` of `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV` **must** be less than or equal to the size of `memory` minus `memoryOffset`

Valid Usage (Implicit)

- [](#VUID-VkBindAccelerationStructureMemoryInfoNV-sType-sType)VUID-VkBindAccelerationStructureMemoryInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_ACCELERATION_STRUCTURE_MEMORY_INFO_NV`
- [](#VUID-VkBindAccelerationStructureMemoryInfoNV-pNext-pNext)VUID-VkBindAccelerationStructureMemoryInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkBindAccelerationStructureMemoryInfoNV-accelerationStructure-parameter)VUID-VkBindAccelerationStructureMemoryInfoNV-accelerationStructure-parameter  
  `accelerationStructure` **must** be a valid [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) handle
- [](#VUID-VkBindAccelerationStructureMemoryInfoNV-memory-parameter)VUID-VkBindAccelerationStructureMemoryInfoNV-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-VkBindAccelerationStructureMemoryInfoNV-pDeviceIndices-parameter)VUID-VkBindAccelerationStructureMemoryInfoNV-pDeviceIndices-parameter  
  If `deviceIndexCount` is not `0`, `pDeviceIndices` **must** be a valid pointer to an array of `deviceIndexCount` `uint32_t` values
- [](#VUID-VkBindAccelerationStructureMemoryInfoNV-commonparent)VUID-VkBindAccelerationStructureMemoryInfoNV-commonparent  
  Both of `accelerationStructure`, and `memory` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindAccelerationStructureMemoryNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBindAccelerationStructureMemoryInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0