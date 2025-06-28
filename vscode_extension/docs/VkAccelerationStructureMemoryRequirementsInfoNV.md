# VkAccelerationStructureMemoryRequirementsInfoNV(3) Manual Page

## Name

VkAccelerationStructureMemoryRequirementsInfoNV - Structure specifying acceleration to query for memory requirements



## [](#_c_specification)C Specification

The `VkAccelerationStructureMemoryRequirementsInfoNV` structure is defined as:

```c++
// Provided by VK_NV_ray_tracing
typedef struct VkAccelerationStructureMemoryRequirementsInfoNV {
    VkStructureType                                    sType;
    const void*                                        pNext;
    VkAccelerationStructureMemoryRequirementsTypeNV    type;
    VkAccelerationStructureNV                          accelerationStructure;
} VkAccelerationStructureMemoryRequirementsInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `type` selects the type of memory requirement being queried. `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV` returns the memory requirements for the object itself. `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BUILD_SCRATCH_NV` returns the memory requirements for the scratch memory when doing a build. `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV` returns the memory requirements for the scratch memory when doing an update.
- `accelerationStructure` is the acceleration structure to be queried for memory requirements.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureMemoryRequirementsInfoNV-sType-sType)VUID-VkAccelerationStructureMemoryRequirementsInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_INFO_NV`
- [](#VUID-VkAccelerationStructureMemoryRequirementsInfoNV-pNext-pNext)VUID-VkAccelerationStructureMemoryRequirementsInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkAccelerationStructureMemoryRequirementsInfoNV-type-parameter)VUID-VkAccelerationStructureMemoryRequirementsInfoNV-type-parameter  
  `type` **must** be a valid [VkAccelerationStructureMemoryRequirementsTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMemoryRequirementsTypeNV.html) value
- [](#VUID-VkAccelerationStructureMemoryRequirementsInfoNV-accelerationStructure-parameter)VUID-VkAccelerationStructureMemoryRequirementsInfoNV-accelerationStructure-parameter  
  `accelerationStructure` **must** be a valid [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) handle

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureMemoryRequirementsTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMemoryRequirementsTypeNV.html), [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureMemoryRequirementsInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0