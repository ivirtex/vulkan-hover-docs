# VkTraceRaysIndirectCommand2KHR(3) Manual Page

## Name

VkTraceRaysIndirectCommand2KHR - Structure specifying the parameters of
an indirect trace ray command with indirect shader binding tables



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkTraceRaysIndirectCommand2KHR` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `raygenShaderRecordAddress` is a
  [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html) of the ray generation shader
  binding table record used by this command.

- `raygenShaderRecordSize` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) number
  of bytes corresponding to the ray generation shader binding table
  record at base address `raygenShaderRecordAddress`.

- `missShaderBindingTableAddress` is a
  [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html) of the first record in the
  miss shader binding table used by this command.

- `missShaderBindingTableSize` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)
  number of bytes corresponding to the total size of the miss shader
  binding table at `missShaderBindingTableAddress` that may be accessed
  by this command.

- `missShaderBindingTableStride` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)
  number of bytes between records of the miss shader binding table.

- `hitShaderBindingTableAddress` is a
  [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html) of the first record in the hit
  shader binding table used by this command.

- `hitShaderBindingTableSize` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)
  number of bytes corresponding to the total size of the hit shader
  binding table at `hitShaderBindingTableAddress` that may be accessed
  by this command.

- `hitShaderBindingTableStride` is a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)
  number of bytes between records of the hit shader binding table.

- `callableShaderBindingTableAddress` is a
  [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html) of the first record in the
  callable shader binding table used by this command.

- `callableShaderBindingTableSize` is a
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) number of bytes corresponding to the
  total size of the callable shader binding table at
  `callableShaderBindingTableAddress` that may be accessed by this
  command.

- `callableShaderBindingTableStride` is a
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) number of bytes between records of
  the callable shader binding table.

- `width` is the width of the ray trace query dimensions.

- `height` is height of the ray trace query dimensions.

- `depth` is depth of the ray trace query dimensions.

## <a href="#_description" class="anchor"></a>Description

The members of `VkTraceRaysIndirectCommand2KHR` have the same meaning as
the similarly named parameters of
[vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysKHR.html).

Indirect shader binding table buffer parameters must satisfy the same
memory alignment and binding requirements as their counterparts in
[vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirectKHR.html) and
[vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysKHR.html).

Valid Usage

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pRayGenShaderBindingTable-03680"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pRayGenShaderBindingTable-03680"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pRayGenShaderBindingTable-03680  
  If the buffer from which `raygenShaderRecordAddress` was queried is
  non-sparse then it **must** be bound completely and contiguously to a
  single `VkDeviceMemory` object

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pRayGenShaderBindingTable-03681"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pRayGenShaderBindingTable-03681"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pRayGenShaderBindingTable-03681  
  The buffer from which the `raygenShaderRecordAddress` is queried
  **must** have been created with the
  `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pRayGenShaderBindingTable-03682"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pRayGenShaderBindingTable-03682"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pRayGenShaderBindingTable-03682  
  `raygenShaderRecordAddress` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pMissShaderBindingTable-03683"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pMissShaderBindingTable-03683"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pMissShaderBindingTable-03683  
  If the buffer from which `missShaderBindingTableAddress` was queried
  is non-sparse then it **must** be bound completely and contiguously to
  a single `VkDeviceMemory` object

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pMissShaderBindingTable-03684"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pMissShaderBindingTable-03684"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pMissShaderBindingTable-03684  
  The buffer from which the `missShaderBindingTableAddress` is queried
  **must** have been created with the
  `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pMissShaderBindingTable-03685"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pMissShaderBindingTable-03685"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pMissShaderBindingTable-03685  
  `missShaderBindingTableAddress` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-stride-03686"
  id="VUID-VkTraceRaysIndirectCommand2KHR-stride-03686"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-stride-03686  
  `missShaderBindingTableStride` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupHandleAlignment`

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-stride-04029"
  id="VUID-VkTraceRaysIndirectCommand2KHR-stride-04029"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-stride-04029  
  `missShaderBindingTableStride` **must** be less than or equal to
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxShaderGroupStride`

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-03687"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-03687"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-03687  
  If the buffer from which `hitShaderBindingTableAddress` was queried is
  non-sparse then it **must** be bound completely and contiguously to a
  single `VkDeviceMemory` object

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-03688"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-03688"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-03688  
  The buffer from which the `hitShaderBindingTableAddress` is queried
  **must** have been created with the
  `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-03689"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-03689"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-03689  
  `hitShaderBindingTableAddress` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-stride-03690"
  id="VUID-VkTraceRaysIndirectCommand2KHR-stride-03690"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-stride-03690  
  `hitShaderBindingTableStride` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupHandleAlignment`

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-stride-04035"
  id="VUID-VkTraceRaysIndirectCommand2KHR-stride-04035"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-stride-04035  
  `hitShaderBindingTableStride` **must** be less than or equal to
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxShaderGroupStride`

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pCallableShaderBindingTable-03691"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pCallableShaderBindingTable-03691"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pCallableShaderBindingTable-03691  
  If the buffer from which `callableShaderBindingTableAddress` was
  queried is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pCallableShaderBindingTable-03692"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pCallableShaderBindingTable-03692"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pCallableShaderBindingTable-03692  
  The buffer from which the `callableShaderBindingTableAddress` is
  queried **must** have been created with the
  `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pCallableShaderBindingTable-03693"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pCallableShaderBindingTable-03693"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pCallableShaderBindingTable-03693  
  `callableShaderBindingTableAddress` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-stride-03694"
  id="VUID-VkTraceRaysIndirectCommand2KHR-stride-03694"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-stride-03694  
  `callableShaderBindingTableStride` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupHandleAlignment`

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-stride-04041"
  id="VUID-VkTraceRaysIndirectCommand2KHR-stride-04041"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-stride-04041  
  `callableShaderBindingTableStride` **must** be less than or equal to
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxShaderGroupStride`

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-flags-03696"
  id="VUID-VkTraceRaysIndirectCommand2KHR-flags-03696"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-flags-03696  
  If the currently bound ray tracing pipeline was created with `flags`
  that included
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`,
  `hitShaderBindingTableAddress` **must** not be zero

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-flags-03697"
  id="VUID-VkTraceRaysIndirectCommand2KHR-flags-03697"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-flags-03697  
  If the currently bound ray tracing pipeline was created with `flags`
  that included
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`,
  `hitShaderBindingTableAddress` **must** not be zero

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-flags-03511"
  id="VUID-VkTraceRaysIndirectCommand2KHR-flags-03511"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-flags-03511  
  If the currently bound ray tracing pipeline was created with `flags`
  that included
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`, the
  shader group handle identified by `missShaderBindingTableAddress`
  **must** not be set to zero

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-flags-03512"
  id="VUID-VkTraceRaysIndirectCommand2KHR-flags-03512"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-flags-03512  
  If the currently bound ray tracing pipeline was created with `flags`
  that included
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`,
  entries in the table identified by `hitShaderBindingTableAddress`
  accessed as a result of this command in order to execute an any-hit
  shader **must** not be set to zero

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-flags-03513"
  id="VUID-VkTraceRaysIndirectCommand2KHR-flags-03513"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-flags-03513  
  If the currently bound ray tracing pipeline was created with `flags`
  that included
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`,
  entries in the table identified by `hitShaderBindingTableAddress`
  accessed as a result of this command in order to execute a closest hit
  shader **must** not be set to zero

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-flags-03514"
  id="VUID-VkTraceRaysIndirectCommand2KHR-flags-03514"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-flags-03514  
  If the currently bound ray tracing pipeline was created with `flags`
  that included
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`,
  entries in the table identified by `hitShaderBindingTableAddress`
  accessed as a result of this command in order to execute an
  intersection shader **must** not be set to zero

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-04735"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-04735"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-04735  
  Any non-zero hit shader group entries in the table identified by
  `hitShaderBindingTableAddress` accessed by this call from a geometry
  with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR` **must**
  have been created with
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR`

- <a
  href="#VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-04736"
  id="VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-04736"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-pHitShaderBindingTable-04736  
  Any non-zero hit shader group entries in the table identified by
  `hitShaderBindingTableAddress` accessed by this call from a geometry
  with a `geometryType` of `VK_GEOMETRY_TYPE_AABBS_KHR` **must** have
  been created with
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR`

<!-- -->

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-width-03638"
  id="VUID-VkTraceRaysIndirectCommand2KHR-width-03638"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-width-03638  
  `width` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[0\] ×
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[0\]

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-height-03639"
  id="VUID-VkTraceRaysIndirectCommand2KHR-height-03639"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-height-03639  
  `height` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[1\] ×
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[1\]

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-depth-03640"
  id="VUID-VkTraceRaysIndirectCommand2KHR-depth-03640"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-depth-03640  
  `depth` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[2\] ×
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[2\]

- <a href="#VUID-VkTraceRaysIndirectCommand2KHR-width-03641"
  id="VUID-VkTraceRaysIndirectCommand2KHR-width-03641"></a>
  VUID-VkTraceRaysIndirectCommand2KHR-width-03641  
  `width` × `height` × `depth` **must** be less than or equal to
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxRayDispatchInvocationCount`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_maintenance1.html),
[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkTraceRaysIndirectCommand2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
