# VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR - Structure describing
the ray tracing maintenance features that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR` structure is
defined as:

``` c
// Provided by VK_KHR_ray_tracing_maintenance1
typedef struct VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           rayTracingMaintenance1;
    VkBool32           rayTracingPipelineTraceRaysIndirect2;
} VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-rayTracingMaintenance1"></span>
  `rayTracingMaintenance1` indicates that the implementation supports
  the following:

  - The `CullMaskKHR` SPIR-V builtin using the `SPV_KHR_ray_cull_mask`
    SPIR-V extension.

  - Additional acceleration structure property queries:
    `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`
    and `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR`.

  - A new access flag `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR`.

  - A new pipeline stage flag bit
    `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`

- <span id="features-rayTracingPipelineTraceRaysIndirect2"></span>
  `rayTracingPipelineTraceRaysIndirect2` indicates whether the
  implementation supports the extended indirect ray tracing command
  [vkCmdTraceRaysIndirect2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirect2KHR.html).

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR` **can** also be used
in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_MAINTENANCE_1_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_maintenance1.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
