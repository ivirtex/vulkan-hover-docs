# VkRayTracingPipelineCreateInfoKHR(3) Manual Page

## Name

VkRayTracingPipelineCreateInfoKHR - Structure specifying parameters of a newly created ray tracing pipeline



## [](#_c_specification)C Specification

The `VkRayTracingPipelineCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_ray_tracing_pipeline
typedef struct VkRayTracingPipelineCreateInfoKHR {
    VkStructureType                                      sType;
    const void*                                          pNext;
    VkPipelineCreateFlags                                flags;
    uint32_t                                             stageCount;
    const VkPipelineShaderStageCreateInfo*               pStages;
    uint32_t                                             groupCount;
    const VkRayTracingShaderGroupCreateInfoKHR*          pGroups;
    uint32_t                                             maxPipelineRayRecursionDepth;
    const VkPipelineLibraryCreateInfoKHR*                pLibraryInfo;
    const VkRayTracingPipelineInterfaceCreateInfoKHR*    pLibraryInterface;
    const VkPipelineDynamicStateCreateInfo*              pDynamicState;
    VkPipelineLayout                                     layout;
    VkPipeline                                           basePipelineHandle;
    int32_t                                              basePipelineIndex;
} VkRayTracingPipelineCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html) specifying how the pipeline will be generated.
- `stageCount` is the number of entries in the `pStages` array.
- `pStages` is a pointer to an array of `stageCount` [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html) structures describing the set of the shader stages to be included in the ray tracing pipeline.
- `groupCount` is the number of entries in the `pGroups` array.
- `pGroups` is a pointer to an array of `groupCount` [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupCreateInfoKHR.html) structures describing the set of the shader stages to be included in each shader group in the ray tracing pipeline.
- `maxPipelineRayRecursionDepth` is the [maximum recursion depth](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-tracing-recursion-depth) of shaders executed by this pipeline.
- `pLibraryInfo` is a pointer to a [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLibraryCreateInfoKHR.html) structure defining pipeline libraries to include.
- `pLibraryInterface` is a pointer to a [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html) structure defining additional information when using pipeline libraries.
- `pDynamicState` is a pointer to a [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html) structure, and is used to indicate which properties of the pipeline state object are dynamic and **can** be changed independently of the pipeline state. This **can** be `NULL`, which means no state in the pipeline is considered dynamic.
- `layout` is the description of binding locations used by both the pipeline and descriptor sets used with the pipeline. If [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is greater than or equal to Vulkan 1.3 or [VK\_KHR\_maintenance4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance4.html) is enabled `layout` **must** not be accessed outside of the duration of the command this structure is passed to.
- `basePipelineHandle` is a pipeline to derive from.
- `basePipelineIndex` is an index into the `pCreateInfos` parameter to use as a pipeline to derive from.

## [](#_description)Description

The parameters `basePipelineHandle` and `basePipelineIndex` are described in more detail in [Pipeline Derivatives](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-pipeline-derivatives).

When `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` is specified, this pipeline defines a *pipeline library* which **cannot** be bound as a ray tracing pipeline directly. Instead, pipeline libraries define common shaders and shader groups which **can** be included in future pipeline creation.

If pipeline libraries are included in `pLibraryInfo`, shaders defined in those libraries are treated as if they were defined as additional entries in `pStages`, appended in the order they appear in the `pLibraries` array and in the `pStages` array when those libraries were defined.

When referencing shader groups in order to obtain a shader group handle, groups defined in those libraries are treated as if they were defined as additional entries in `pGroups`, appended in the order they appear in the `pLibraries` array and in the `pGroups` array when those libraries were defined. The shaders these groups reference are set when the pipeline library is created, referencing those specified in the pipeline library, not in the pipeline that includes it.

The default stack size for a pipeline if `VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR` is not provided is computed as described in [Ray Tracing Pipeline Stack](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-tracing-pipeline-stack).

If the `pNext` chain includes a [VkPipelineCreateFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfo.html) structure, [VkPipelineCreateFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfo.html)::`flags` from that structure is used instead of `flags` from this structure.

If the `pNext` chain includes a [VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV.html) structure, then that structure controls whether cluster acceleration structures are allowed in this ray tracing pipeline.

Valid Usage

- [](#VUID-VkRayTracingPipelineCreateInfoKHR-None-09497)VUID-VkRayTracingPipelineCreateInfoKHR-None-09497  
  If the `pNext` chain does not include a [VkPipelineCreateFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfo.html) structure, `flags` **must** be a valid combination of [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html) values
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-07984)VUID-VkRayTracingPipelineCreateInfoKHR-flags-07984  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and `basePipelineIndex` is -1, `basePipelineHandle` **must** be a valid ray tracing `VkPipeline` handle
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-07985)VUID-VkRayTracingPipelineCreateInfoKHR-flags-07985  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and `basePipelineHandle` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `basePipelineIndex` **must** be a valid index into the calling command’s `pCreateInfos` parameter
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-07986)VUID-VkRayTracingPipelineCreateInfoKHR-flags-07986  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, `basePipelineIndex` **must** be -1 or `basePipelineHandle` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-layout-07987)VUID-VkRayTracingPipelineCreateInfoKHR-layout-07987  
  If a push constant block is declared in a shader, a push constant range in `layout` **must** match the shader stage
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-layout-10069)VUID-VkRayTracingPipelineCreateInfoKHR-layout-10069  
  If a push constant block is declared in a shader, the block must be contained inside the push constant range in `layout` that matches the stage
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-layout-07988)VUID-VkRayTracingPipelineCreateInfoKHR-layout-07988  
  If a [resource variable](#interfaces-resources) is declared in a shader, the corresponding descriptor set in `layout` **must** match the shader stage
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-layout-07990)VUID-VkRayTracingPipelineCreateInfoKHR-layout-07990  
  If a [resource variable](#interfaces-resources) is declared in a shader, and the descriptor type is not `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, the corresponding descriptor set in `layout` **must** match the descriptor type
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-layout-07991)VUID-VkRayTracingPipelineCreateInfoKHR-layout-07991  
  If a [resource variable](#interfaces-resources) is declared in a shader as an array, the corresponding descriptor set in `layout` **must** match the descriptor count
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-None-10391)VUID-VkRayTracingPipelineCreateInfoKHR-None-10391  
  If a [resource variables](#interfaces-resources) is declared in a shader as an array of descriptors, then the descriptor type of that variable **must** not be `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`

<!--THE END-->

- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pStages-03426)VUID-VkRayTracingPipelineCreateInfoKHR-pStages-03426  
  The shader code for the entry points identified by `pStages`, and the rest of the state identified by this structure **must** adhere to the pipeline linking rules described in the [Shader Interfaces](#interfaces) chapter
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-layout-03428)VUID-VkRayTracingPipelineCreateInfoKHR-layout-03428  
  The number of resources in `layout` accessible to each shader stage that is used by the pipeline **must** be less than or equal to [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`maxPerStageResources`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-02904)VUID-VkRayTracingPipelineCreateInfoKHR-flags-02904  
  `flags` **must** not include `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pipelineCreationCacheControl-02905)VUID-VkRayTracingPipelineCreateInfoKHR-pipelineCreationCacheControl-02905  
  If the [`pipelineCreationCacheControl`](#features-pipelineCreationCacheControl) feature is not enabled, `flags` **must** not include `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` or `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-stage-03425)VUID-VkRayTracingPipelineCreateInfoKHR-stage-03425  
  If `flags` does not include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`, the `stage` member of at least one element of `pStages`, including those implicitly added by `pLibraryInfo`, **must** be `VK_SHADER_STAGE_RAYGEN_BIT_KHR`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-maxPipelineRayRecursionDepth-03589)VUID-VkRayTracingPipelineCreateInfoKHR-maxPipelineRayRecursionDepth-03589  
  `maxPipelineRayRecursionDepth` **must** be less than or equal to [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`maxRayRecursionDepth`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-03465)VUID-VkRayTracingPipelineCreateInfoKHR-flags-03465  
  If `flags` includes `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`, `pLibraryInterface` **must** not be `NULL`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03590)VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03590  
  If `pLibraryInfo` is not `NULL` and its `libraryCount` member is greater than `0`, `pLibraryInterface` **must** not be `NULL`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraries-03591)VUID-VkRayTracingPipelineCreateInfoKHR-pLibraries-03591  
  Each element of `pLibraryInfo->pLibraries` **must** have been created with the value of `maxPipelineRayRecursionDepth` equal to that in this pipeline
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03592)VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03592  
  If `pLibraryInfo` is not `NULL`, each element of its `pLibraries` member **must** have been created with a `layout` that is compatible with the `layout` in this pipeline
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03593)VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03593  
  If `pLibraryInfo` is not `NULL`, each element of its `pLibraries` member **must** have been created with values of the `maxPipelineRayPayloadSize` and `maxPipelineRayHitAttributeSize` members of `pLibraryInterface` equal to those in this pipeline
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-03594)VUID-VkRayTracingPipelineCreateInfoKHR-flags-03594  
  If `flags` includes `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`, each element of `pLibraryInfo->pLibraries` **must** have been created with the `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR` bit set
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-04718)VUID-VkRayTracingPipelineCreateInfoKHR-flags-04718  
  If `flags` includes `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`, each element of `pLibraryInfo->pLibraries` **must** have been created with the `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR` bit set
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-04719)VUID-VkRayTracingPipelineCreateInfoKHR-flags-04719  
  If `flags` includes `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`, each element of `pLibraryInfo->pLibraries` **must** have been created with the `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR` bit set
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-04720)VUID-VkRayTracingPipelineCreateInfoKHR-flags-04720  
  If `flags` includes `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`, each element of `pLibraryInfo->pLibraries` **must** have been created with the `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR` bit set
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-04721)VUID-VkRayTracingPipelineCreateInfoKHR-flags-04721  
  If `flags` includes `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`, each element of `pLibraryInfo->pLibraries` **must** have been created with the `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR` bit set
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-04722)VUID-VkRayTracingPipelineCreateInfoKHR-flags-04722  
  If `flags` includes `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`, each element of `pLibraryInfo->pLibraries` **must** have been created with the `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR` bit set
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-04723)VUID-VkRayTracingPipelineCreateInfoKHR-flags-04723  
  If `flags` includes `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`, each element of `pLibraryInfo->pLibraries` **must** have been created with the `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR` bit set
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03595)VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03595  
  If the `VK_KHR_pipeline_library` extension is not enabled, `pLibraryInfo` and `pLibraryInterface` **must** be `NULL`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-03470)VUID-VkRayTracingPipelineCreateInfoKHR-flags-03470  
  If `flags` includes `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`, for any element of `pGroups` with a `type` of `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR` or `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR`, the `anyHitShader` of that element **must** not be `VK_SHADER_UNUSED_KHR`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-03471)VUID-VkRayTracingPipelineCreateInfoKHR-flags-03471  
  If `flags` includes `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`, for any element of `pGroups` with a `type` of `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR` or `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR`, the `closestHitShader` of that element **must** not be `VK_SHADER_UNUSED_KHR`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-rayTraversalPrimitiveCulling-03596)VUID-VkRayTracingPipelineCreateInfoKHR-rayTraversalPrimitiveCulling-03596  
  If the [`rayTraversalPrimitiveCulling`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-rayTraversalPrimitiveCulling) feature is not enabled, `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-rayTraversalPrimitiveCulling-03597)VUID-VkRayTracingPipelineCreateInfoKHR-rayTraversalPrimitiveCulling-03597  
  If the [`rayTraversalPrimitiveCulling`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-rayTraversalPrimitiveCulling) feature is not enabled, `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-06546)VUID-VkRayTracingPipelineCreateInfoKHR-flags-06546  
  `flags` **must** not include both `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR` and `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-03598)VUID-VkRayTracingPipelineCreateInfoKHR-flags-03598  
  If `flags` includes `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`, [`rayTracingPipelineShaderGroupHandleCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-rayTracingPipelineShaderGroupHandleCaptureReplay) **must** be enabled
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-rayTracingPipelineShaderGroupHandleCaptureReplay-03599)VUID-VkRayTracingPipelineCreateInfoKHR-rayTracingPipelineShaderGroupHandleCaptureReplay-03599  
  If [VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html)::`rayTracingPipelineShaderGroupHandleCaptureReplay` is `VK_TRUE` and the `pShaderGroupCaptureReplayHandle` member of any element of `pGroups` is not `NULL`, `flags` **must** include `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-07999)VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-07999  
  If `pLibraryInfo` is `NULL` or its `libraryCount` is `0`, `stageCount` **must** not be `0`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-08700)VUID-VkRayTracingPipelineCreateInfoKHR-flags-08700  
  If `flags` does not include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` and either `pLibraryInfo` is `NULL` or its `libraryCount` is `0`, `groupCount` **must** not be `0`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pDynamicStates-03602)VUID-VkRayTracingPipelineCreateInfoKHR-pDynamicStates-03602  
  Any element of the `pDynamicStates` member of `pDynamicState` **must** be `VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pipelineStageCreationFeedbackCount-06652)VUID-VkRayTracingPipelineCreateInfoKHR-pipelineStageCreationFeedbackCount-06652  
  If [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackCreateInfo.html)::`pipelineStageCreationFeedbackCount` is not `0`, it **must** be equal to `stageCount`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-stage-06899)VUID-VkRayTracingPipelineCreateInfoKHR-stage-06899  
  The `stage` value in all `pStages` elements **must** be one of `VK_SHADER_STAGE_RAYGEN_BIT_KHR`, `VK_SHADER_STAGE_ANY_HIT_BIT_KHR`, `VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR`, `VK_SHADER_STAGE_MISS_BIT_KHR`, `VK_SHADER_STAGE_INTERSECTION_BIT_KHR`, or `VK_SHADER_STAGE_CALLABLE_BIT_KHR`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-07403)VUID-VkRayTracingPipelineCreateInfoKHR-flags-07403  
  If `flags` includes `VK_PIPELINE_CREATE_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT`, each element of `pLibraryInfo->pLibraries` **must** have been created with the `VK_PIPELINE_CREATE_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT` bit set
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-08701)VUID-VkRayTracingPipelineCreateInfoKHR-flags-08701  
  If `flags` includes `VK_PIPELINE_CREATE_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV`, each element of `pLibraryInfo->pLibraries` **must** have been created with the `VK_PIPELINE_CREATE_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV` bit set
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-flags-10392)VUID-VkRayTracingPipelineCreateInfoKHR-flags-10392  
  If the `pNext` chain includes a [VkPipelineCreateFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfo.html) structure, `flags` **must** not include both `VK_PIPELINE_CREATE_2_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT` and `VK_PIPELINE_CREATE_2_DISALLOW_OPACITY_MICROMAP_BIT_ARM`

Valid Usage (Implicit)

- [](#VUID-VkRayTracingPipelineCreateInfoKHR-sType-sType)VUID-VkRayTracingPipelineCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CREATE_INFO_KHR`
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pNext-pNext)VUID-VkRayTracingPipelineCreateInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html), [VkPipelineCreateFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfo.html), [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackCreateInfo.html), [VkPipelineRobustnessCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessCreateInfo.html), or [VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV.html)
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-sType-unique)VUID-VkRayTracingPipelineCreateInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pStages-parameter)VUID-VkRayTracingPipelineCreateInfoKHR-pStages-parameter  
  If `stageCount` is not `0`, `pStages` **must** be a valid pointer to an array of `stageCount` valid [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html) structures
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pGroups-parameter)VUID-VkRayTracingPipelineCreateInfoKHR-pGroups-parameter  
  If `groupCount` is not `0`, `pGroups` **must** be a valid pointer to an array of `groupCount` valid [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupCreateInfoKHR.html) structures
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-parameter)VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-parameter  
  If `pLibraryInfo` is not `NULL`, `pLibraryInfo` **must** be a valid pointer to a valid [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLibraryCreateInfoKHR.html) structure
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInterface-parameter)VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInterface-parameter  
  If `pLibraryInterface` is not `NULL`, `pLibraryInterface` **must** be a valid pointer to a valid [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html) structure
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-pDynamicState-parameter)VUID-VkRayTracingPipelineCreateInfoKHR-pDynamicState-parameter  
  If `pDynamicState` is not `NULL`, `pDynamicState` **must** be a valid pointer to a valid [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html) structure
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-layout-parameter)VUID-VkRayTracingPipelineCreateInfoKHR-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkRayTracingPipelineCreateInfoKHR-commonparent)VUID-VkRayTracingPipelineCreateInfoKHR-commonparent  
  Both of `basePipelineHandle`, and `layout` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags.html), [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLibraryCreateInfoKHR.html), [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html), [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupCreateInfoKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRayTracingPipelinesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRayTracingPipelineCreateInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0