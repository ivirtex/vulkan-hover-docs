# vkGetAccelerationStructureMemoryRequirementsNV(3) Manual Page

## Name

vkGetAccelerationStructureMemoryRequirementsNV - Get acceleration structure memory requirements



## [](#_c_specification)C Specification

An acceleration structure has memory requirements for the structure object itself, scratch space for the build, and scratch space for the update.

Scratch space is allocated as a `VkBuffer`, so for `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BUILD_SCRATCH_NV` and `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV` the `pMemoryRequirements->alignment` and `pMemoryRequirements->memoryTypeBits` values returned by this call **must** be filled with zero, and **should** be ignored by the application.

To query the memory requirements, call:

```c++
// Provided by VK_NV_ray_tracing
void vkGetAccelerationStructureMemoryRequirementsNV(
    VkDevice                                    device,
    const VkAccelerationStructureMemoryRequirementsInfoNV* pInfo,
    VkMemoryRequirements2KHR*                   pMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device on which the acceleration structure was created.
- `pInfo` is a pointer to a [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html) structure specifying the acceleration structure to get memory requirements for.
- `pMemoryRequirements` is a pointer to a [VkMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2KHR.html) structure in which the requested acceleration structure memory requirements are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetAccelerationStructureMemoryRequirementsNV-device-parameter)VUID-vkGetAccelerationStructureMemoryRequirementsNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetAccelerationStructureMemoryRequirementsNV-pInfo-parameter)VUID-vkGetAccelerationStructureMemoryRequirementsNV-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html) structure
- [](#VUID-vkGetAccelerationStructureMemoryRequirementsNV-pMemoryRequirements-parameter)VUID-vkGetAccelerationStructureMemoryRequirementsNV-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a [VkMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2KHR.html) structure

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetAccelerationStructureMemoryRequirementsNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0