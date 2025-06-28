# VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV(3) Manual Page

## Name

VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV - Structure controlling if cluster acceleration structures are allowed



## [](#_c_specification)C Specification

The [VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV.html) structure is defined as:

```c++
// Provided by VK_KHR_ray_tracing_pipeline with VK_NV_cluster_acceleration_structure
typedef struct VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           allowClusterAccelerationStructure;
} VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `allowClusterAccelerationStructure` controls if cluster acceleration structures are allowed in the ray tracing pipeline.

## [](#_description)Description

If no cluster acceleration structures are present in the ray tracing pipeline, `allowClusterAccelerationStructure` **should** not be used to prevent traversal penalty on some implementations.

Valid Usage

- [](#VUID-VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV-clusterAccelerationStructure-10576)VUID-VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV-clusterAccelerationStructure-10576  
  The [`VkPhysicalDeviceClusterAccelerationStructureFeaturesNV`::`clusterAccelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-clusterAccelerationStructure) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV-sType-sType)VUID-VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CLUSTER_ACCELERATION_STRUCTURE_CREATE_INFO_NV`

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0