# VkTraceRaysIndirectCommand2KHR(3) Manual Page

## Name

VkTraceRaysIndirectCommand2KHR - Structure specifying the parameters of an indirect trace ray command with indirect shader binding tables



## [](#_c_specification)C Specification

The `VkTraceRaysIndirectCommand2KHR` structure is defined as:

```c++
// Provided by VK_KHR_ray_tracing_maintenance1 with VK_KHR_ray_tracing_pipeline
typedef struct VkTraceRaysIndirectCommand2KHR {
    VkDeviceAddress    raygenShaderRecordAddress;
    VkDeviceSize       raygenShaderRecordSize;
    VkDeviceAddress    missShaderBindingTableAddress;
    VkDeviceSize       missShaderBindingTableSize;
    VkDeviceSize       missShaderBindingTableStride;
    VkDeviceAddress    hitShaderBindingTableAddress;
    VkDeviceSize       hitShaderBindingTableSize;
    VkDeviceSize       hitShaderBindingTableStride;
    VkDeviceAddress    callableShaderBindingTableAddress;
    VkDeviceSize       callableShaderBindingTableSize;
    VkDeviceSize       callableShaderBindingTableStride;
    uint32_t           width;
    uint32_t           height;
    uint32_t           depth;
} VkTraceRaysIndirectCommand2KHR;
```

## [](#_members)Members

- `raygenShaderRecordAddress` is a [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) of the ray generation shader binding table record used by this command.
- `raygenShaderRecordSize` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) number of bytes corresponding to the ray generation shader binding table record at base address `raygenShaderRecordAddress`.
- `missShaderBindingTableAddress` is a [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) of the first record in the miss shader binding table used by this command.
- `missShaderBindingTableSize` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) number of bytes corresponding to the total size of the miss shader binding table at `missShaderBindingTableAddress` that may be accessed by this command.
- `missShaderBindingTableStride` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) number of bytes between records of the miss shader binding table.
- `hitShaderBindingTableAddress` is a [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) of the first record in the hit shader binding table used by this command.
- `hitShaderBindingTableSize` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) number of bytes corresponding to the total size of the hit shader binding table at `hitShaderBindingTableAddress` that may be accessed by this command.
- `hitShaderBindingTableStride` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) number of bytes between records of the hit shader binding table.
- `callableShaderBindingTableAddress` is a [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) of the first record in the callable shader binding table used by this command.
- `callableShaderBindingTableSize` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) number of bytes corresponding to the total size of the callable shader binding table at `callableShaderBindingTableAddress` that may be accessed by this command.
- `callableShaderBindingTableStride` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) number of bytes between records of the callable shader binding table.
- `width` is the width of the ray trace query dimensions.
- `height` is height of the ray trace query dimensions.
- `depth` is depth of the ray trace query dimensions.

## [](#_description)Description

The members of `VkTraceRaysIndirectCommand2KHR` have the same meaning as the similarly named parameters of [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysKHR.html).

Indirect shader binding table buffer parameters **must** satisfy the same memory alignment and binding requirements as their counterparts in [vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysIndirectKHR.html) and [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysKHR.html).

Valid Usage

- [](#VUID-VkTraceRaysIndirectCommand2KHR-pRayGenShaderBindingTable-03681)VUID-VkTraceRaysIndirectCommand2KHR-pRayGenShaderBindingTable-03681  
  The buffer from which the `raygenShaderRecordAddress` is queried **must** have been created with the `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag
- [](#VUID-VkTraceRaysIndirectCommand2KHR-pRayGenShaderBindingTable-03682)VUID-VkTraceRaysIndirectCommand2KHR-pRayGenShaderBindingTable-03682  
  `raygenShaderRecordAddress` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`
- [](#VUID-VkTraceRaysIndirectCommand2KHR-pMissShaderBindingTable-03684)VUID-VkTraceRaysIndirectCommand2KHR-pMissShaderBindingTable-03684  
  The buffer from which the `missShaderBindingTableAddress` is queried **must** have been created with the `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag
- [](#VUID-VkTraceRaysIndirectCommand2KHR-pMissShaderBindingTable-03685)VUID-VkTraceRaysIndirectCommand2KHR-pMissShaderBindingTable-03685  
  `missShaderBindingTableAddress` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`
- [](#VUID-VkTraceRaysIndirectCommand2KHR-stride-03686)VUID-VkTraceRaysIndirectCommand2KHR-stride-03686  
  `missShaderBindingTableStride` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupHandleAlignment`
- [](#VUID-VkTraceRaysIndirectCommand2KHR-stride-04029)VUID-VkTraceRaysIndirectCommand2KHR-stride-04029  
  `missShaderBindingTableStride` **must** be less than or equal to `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxShaderGroupStride`
- [](#VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-03688)VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-03688  
  The buffer from which the `hitShaderBindingTableAddress` is queried **must** have been created with the `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag
- [](#VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-03689)VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-03689  
  `hitShaderBindingTableAddress` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`
- [](#VUID-VkTraceRaysIndirectCommand2KHR-stride-03690)VUID-VkTraceRaysIndirectCommand2KHR-stride-03690  
  `hitShaderBindingTableStride` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupHandleAlignment`
- [](#VUID-VkTraceRaysIndirectCommand2KHR-stride-04035)VUID-VkTraceRaysIndirectCommand2KHR-stride-04035  
  `hitShaderBindingTableStride` **must** be less than or equal to `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxShaderGroupStride`
- [](#VUID-VkTraceRaysIndirectCommand2KHR-pCallableShaderBindingTable-03692)VUID-VkTraceRaysIndirectCommand2KHR-pCallableShaderBindingTable-03692  
  The buffer from which the `callableShaderBindingTableAddress` is queried **must** have been created with the `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag
- [](#VUID-VkTraceRaysIndirectCommand2KHR-pCallableShaderBindingTable-03693)VUID-VkTraceRaysIndirectCommand2KHR-pCallableShaderBindingTable-03693  
  `callableShaderBindingTableAddress` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`
- [](#VUID-VkTraceRaysIndirectCommand2KHR-stride-03694)VUID-VkTraceRaysIndirectCommand2KHR-stride-03694  
  `callableShaderBindingTableStride` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupHandleAlignment`
- [](#VUID-VkTraceRaysIndirectCommand2KHR-stride-04041)VUID-VkTraceRaysIndirectCommand2KHR-stride-04041  
  `callableShaderBindingTableStride` **must** be less than or equal to `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxShaderGroupStride`
- [](#VUID-VkTraceRaysIndirectCommand2KHR-flags-03696)VUID-VkTraceRaysIndirectCommand2KHR-flags-03696  
  If the bound ray tracing pipeline was created with `flags` that included `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`, `hitShaderBindingTableAddress` **must** not be zero
- [](#VUID-VkTraceRaysIndirectCommand2KHR-flags-03697)VUID-VkTraceRaysIndirectCommand2KHR-flags-03697  
  If the bound ray tracing pipeline was created with `flags` that included `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`, `hitShaderBindingTableAddress` **must** not be zero
- [](#VUID-VkTraceRaysIndirectCommand2KHR-flags-03511)VUID-VkTraceRaysIndirectCommand2KHR-flags-03511  
  If the bound ray tracing pipeline was created with `flags` that included `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`, the shader group handle identified by `missShaderBindingTableAddress` **must** not be zero
- [](#VUID-VkTraceRaysIndirectCommand2KHR-flags-03512)VUID-VkTraceRaysIndirectCommand2KHR-flags-03512  
  If the bound ray tracing pipeline was created with `flags` that included `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`, entries in the table identified by `hitShaderBindingTableAddress` accessed as a result of this command in order to execute an any-hit shader **must** not be zero
- [](#VUID-VkTraceRaysIndirectCommand2KHR-flags-03513)VUID-VkTraceRaysIndirectCommand2KHR-flags-03513  
  If the bound ray tracing pipeline was created with `flags` that included `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`, entries in the table identified by `hitShaderBindingTableAddress` accessed as a result of this command in order to execute a closest hit shader **must** not be zero
- [](#VUID-VkTraceRaysIndirectCommand2KHR-flags-03514)VUID-VkTraceRaysIndirectCommand2KHR-flags-03514  
  If the bound ray tracing pipeline was created with `flags` that included `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`, entries in the table identified by `hitShaderBindingTableAddress` accessed as a result of this command in order to execute an intersection shader **must** not be zero
- [](#VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-04735)VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-04735  
  Any non-zero hit shader group entries in the table identified by `hitShaderBindingTableAddress` accessed by this call from a geometry with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR` **must** have been created with `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR`
- [](#VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-04736)VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-04736  
  Any non-zero hit shader group entries in the table identified by `hitShaderBindingTableAddress` accessed by this call from a geometry with a `geometryType` of `VK_GEOMETRY_TYPE_AABBS_KHR` **must** have been created with `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR`

<!--THE END-->

- [](#VUID-VkTraceRaysIndirectCommand2KHR-width-03638)VUID-VkTraceRaysIndirectCommand2KHR-width-03638  
  `width` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[0] × `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[0]
- [](#VUID-VkTraceRaysIndirectCommand2KHR-height-03639)VUID-VkTraceRaysIndirectCommand2KHR-height-03639  
  `height` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[1] × `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[1]
- [](#VUID-VkTraceRaysIndirectCommand2KHR-depth-03640)VUID-VkTraceRaysIndirectCommand2KHR-depth-03640  
  `depth` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[2] × `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[2]
- [](#VUID-VkTraceRaysIndirectCommand2KHR-width-03641)VUID-VkTraceRaysIndirectCommand2KHR-width-03641  
  `width` × `height` × `depth` **must** be less than or equal to `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxRayDispatchInvocationCount`

Valid Usage (Implicit)

- [](#VUID-VkTraceRaysIndirectCommand2KHR-raygenShaderRecordAddress-parameter)VUID-VkTraceRaysIndirectCommand2KHR-raygenShaderRecordAddress-parameter  
  `raygenShaderRecordAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkTraceRaysIndirectCommand2KHR-missShaderBindingTableAddress-parameter)VUID-VkTraceRaysIndirectCommand2KHR-missShaderBindingTableAddress-parameter  
  `missShaderBindingTableAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkTraceRaysIndirectCommand2KHR-hitShaderBindingTableAddress-parameter)VUID-VkTraceRaysIndirectCommand2KHR-hitShaderBindingTableAddress-parameter  
  `hitShaderBindingTableAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkTraceRaysIndirectCommand2KHR-callableShaderBindingTableAddress-parameter)VUID-VkTraceRaysIndirectCommand2KHR-callableShaderBindingTableAddress-parameter  
  `callableShaderBindingTableAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_maintenance1.html), [VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTraceRaysIndirectCommand2KHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0