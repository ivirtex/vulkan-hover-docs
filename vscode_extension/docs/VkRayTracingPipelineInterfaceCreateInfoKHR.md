# VkRayTracingPipelineInterfaceCreateInfoKHR(3) Manual Page

## Name

VkRayTracingPipelineInterfaceCreateInfoKHR - Structure specifying additional interface information when using libraries



## [](#_c_specification)C Specification

The `VkRayTracingPipelineInterfaceCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_ray_tracing_pipeline
typedef struct VkRayTracingPipelineInterfaceCreateInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           maxPipelineRayPayloadSize;
    uint32_t           maxPipelineRayHitAttributeSize;
} VkRayTracingPipelineInterfaceCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxPipelineRayPayloadSize` is the maximum payload size in bytes used by any shader in the pipeline.
- `maxPipelineRayHitAttributeSize` is the maximum attribute structure size in bytes used by any shader in the pipeline.

## [](#_description)Description

`maxPipelineRayPayloadSize` is calculated as the maximum size of the block (in bytes) declared in the `RayPayloadKHR` or `IncomingRayPayloadKHR` storage classes. `maxPipelineRayHitAttributeSize` is calculated as the maximum size of any block (in bytes) declared in the `HitAttributeKHR` storage class. As variables in these storage classes do not have explicit offsets, the size should be calculated as if each variable has a [scalar alignment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-alignment-requirements) equal to the largest scalar alignment of any of the block’s members.

Note

There is no explicit upper limit for `maxPipelineRayPayloadSize`, but in practice it should be kept as small as possible. Similar to invocation local memory, it must be allocated for each shader invocation and for devices which support many simultaneous invocations, this storage can rapidly be exhausted, resulting in failure.

Valid Usage

- [](#VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-maxPipelineRayHitAttributeSize-03605)VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-maxPipelineRayHitAttributeSize-03605  
  `maxPipelineRayHitAttributeSize` **must** be less than or equal to [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`maxRayHitAttributeSize`

Valid Usage (Implicit)

- [](#VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-sType-sType)VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_INTERFACE_CREATE_INFO_KHR`
- [](#VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-pNext-pNext)VUID-VkRayTracingPipelineInterfaceCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRayTracingPipelineInterfaceCreateInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0