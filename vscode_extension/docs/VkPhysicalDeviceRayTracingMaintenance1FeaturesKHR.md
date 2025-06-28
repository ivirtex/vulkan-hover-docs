# VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR - Structure describing the ray tracing maintenance features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_ray_tracing_maintenance1
typedef struct VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           rayTracingMaintenance1;
    VkBool32           rayTracingPipelineTraceRaysIndirect2;
} VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`rayTracingMaintenance1` indicates that the implementation supports the following:
  
  - The `CullMaskKHR` SPIR-V builtin using the `SPV_KHR_ray_cull_mask` SPIR-V extension.
  - Additional acceleration structure property queries: `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR` and `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR`.
  - A new access flag `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR`.
  - A new pipeline stage flag bit `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`
- []()`rayTracingPipelineTraceRaysIndirect2` indicates whether the implementation supports the extended indirect ray tracing command [vkCmdTraceRaysIndirect2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysIndirect2KHR.html).

## [](#_description)Description

If the `VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR-sType-sType)VUID-VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_MAINTENANCE_1_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_maintenance1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0