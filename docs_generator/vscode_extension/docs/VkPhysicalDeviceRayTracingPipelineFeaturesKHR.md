# VkPhysicalDeviceRayTracingPipelineFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceRayTracingPipelineFeaturesKHR - Structure describing the ray tracing features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceRayTracingPipelineFeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_ray_tracing_pipeline
typedef struct VkPhysicalDeviceRayTracingPipelineFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           rayTracingPipeline;
    VkBool32           rayTracingPipelineShaderGroupHandleCaptureReplay;
    VkBool32           rayTracingPipelineShaderGroupHandleCaptureReplayMixed;
    VkBool32           rayTracingPipelineTraceRaysIndirect;
    VkBool32           rayTraversalPrimitiveCulling;
} VkPhysicalDeviceRayTracingPipelineFeaturesKHR;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`rayTracingPipeline` indicates whether the implementation supports the ray tracing pipeline functionality. See [Ray Tracing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-tracing).
- []()`rayTracingPipelineShaderGroupHandleCaptureReplay` indicates whether the implementation supports saving and reusing shader group handles, e.g. for trace capture and replay.
- []()`rayTracingPipelineShaderGroupHandleCaptureReplayMixed` indicates whether the implementation supports reuse of shader group handles being arbitrarily mixed with creation of non-reused shader group handles. If this is `VK_FALSE`, all reused shader group handles **must** be specified before any non-reused handles **may** be created.
- []()`rayTracingPipelineTraceRaysIndirect` indicates whether the implementation supports indirect ray tracing commands, e.g. [vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysIndirectKHR.html).
- []()`rayTraversalPrimitiveCulling` indicates whether the implementation supports [primitive culling during ray traversal](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-traversal-culling-primitive).

## [](#_description)Description

If the `VkPhysicalDeviceRayTracingPipelineFeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceRayTracingPipelineFeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage

- [](#VUID-VkPhysicalDeviceRayTracingPipelineFeaturesKHR-rayTracingPipelineShaderGroupHandleCaptureReplayMixed-03575)VUID-VkPhysicalDeviceRayTracingPipelineFeaturesKHR-rayTracingPipelineShaderGroupHandleCaptureReplayMixed-03575  
  If `rayTracingPipelineShaderGroupHandleCaptureReplayMixed` is `VK_TRUE`, `rayTracingPipelineShaderGroupHandleCaptureReplay` **must** also be `VK_TRUE`

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceRayTracingPipelineFeaturesKHR-sType-sType)VUID-VkPhysicalDeviceRayTracingPipelineFeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PIPELINE_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceRayTracingPipelineFeaturesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0