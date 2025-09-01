# VkPhysicalDeviceRayTracingPipelinePropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceRayTracingPipelinePropertiesKHR - Properties of the physical device for ray tracing



## [](#_c_specification)C Specification

The `VkPhysicalDeviceRayTracingPipelinePropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_ray_tracing_pipeline
typedef struct VkPhysicalDeviceRayTracingPipelinePropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           shaderGroupHandleSize;
    uint32_t           maxRayRecursionDepth;
    uint32_t           maxShaderGroupStride;
    uint32_t           shaderGroupBaseAlignment;
    uint32_t           shaderGroupHandleCaptureReplaySize;
    uint32_t           maxRayDispatchInvocationCount;
    uint32_t           shaderGroupHandleAlignment;
    uint32_t           maxRayHitAttributeSize;
} VkPhysicalDeviceRayTracingPipelinePropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `shaderGroupHandleSize` is the size in bytes of the shader header.
- []()`maxRayRecursionDepth` is the maximum number of levels of ray recursion allowed in a trace command.
- `maxShaderGroupStride` is the maximum stride in bytes allowed between shader groups in the shader binding table.
- `shaderGroupBaseAlignment` is the **required** alignment in bytes for the base of the shader binding table.
- `shaderGroupHandleCaptureReplaySize` is the number of bytes for the information required to do capture and replay for shader group handles.
- `maxRayDispatchInvocationCount` is the maximum number of ray generation shader invocations which **may** be produced by a single [vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysIndirectKHR.html) or [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysKHR.html) command.
- `shaderGroupHandleAlignment` is the **required** alignment in bytes for each entry in a shader binding table. The value **must** be a power of two.
- `maxRayHitAttributeSize` is the maximum size in bytes for a ray attribute structure

## [](#_description)Description

If the `VkPhysicalDeviceRayTracingPipelinePropertiesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Limits specified by this structure **must** match those specified with the same name in [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceRayTracingPipelinePropertiesKHR-sType-sType)VUID-VkPhysicalDeviceRayTracingPipelinePropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_PIPELINE_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceRayTracingPipelinePropertiesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0