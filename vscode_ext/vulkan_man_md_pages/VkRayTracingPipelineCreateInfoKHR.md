# VkRayTracingPipelineCreateInfoKHR(3) Manual Page

## Name

VkRayTracingPipelineCreateInfoKHR - Structure specifying parameters of a
newly created ray tracing pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRayTracingPipelineCreateInfoKHR` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html) specifying
  how the pipeline will be generated.

- `stageCount` is the number of entries in the `pStages` array.

- `pStages` is a pointer to an array of `stageCount`
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  structures describing the set of the shader stages to be included in
  the ray tracing pipeline.

- `groupCount` is the number of entries in the `pGroups` array.

- `pGroups` is a pointer to an array of `groupCount`
  [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoKHR.html)
  structures describing the set of the shader stages to be included in
  each shader group in the ray tracing pipeline.

- `maxPipelineRayRecursionDepth` is the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-tracing-recursion-depth"
  target="_blank" rel="noopener">maximum recursion depth</a> of shaders
  executed by this pipeline.

- `pLibraryInfo` is a pointer to a
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)
  structure defining pipeline libraries to include.

- `pLibraryInterface` is a pointer to a
  [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html)
  structure defining additional information when using pipeline
  libraries.

- `pDynamicState` is a pointer to a
  [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)
  structure, and is used to indicate which properties of the pipeline
  state object are dynamic and **can** be changed independently of the
  pipeline state. This **can** be `NULL`, which means no state in the
  pipeline is considered dynamic.

- `layout` is the description of binding locations used by both the
  pipeline and descriptor sets used with the pipeline.

- `basePipelineHandle` is a pipeline to derive from.

- `basePipelineIndex` is an index into the `pCreateInfos` parameter to
  use as a pipeline to derive from.

## <a href="#_description" class="anchor"></a>Description

The parameters `basePipelineHandle` and `basePipelineIndex` are
described in more detail in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-pipeline-derivatives"
target="_blank" rel="noopener">Pipeline Derivatives</a>.

When `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` is specified, this pipeline
defines a *pipeline library* which **cannot** be bound as a ray tracing
pipeline directly. Instead, pipeline libraries define common shaders and
shader groups which **can** be included in future pipeline creation.

If pipeline libraries are included in `pLibraryInfo`, shaders defined in
those libraries are treated as if they were defined as additional
entries in `pStages`, appended in the order they appear in the
`pLibraries` array and in the `pStages` array when those libraries were
defined.

When referencing shader groups in order to obtain a shader group handle,
groups defined in those libraries are treated as if they were defined as
additional entries in `pGroups`, appended in the order they appear in
the `pLibraries` array and in the `pGroups` array when those libraries
were defined. The shaders these groups reference are set when the
pipeline library is created, referencing those specified in the pipeline
library, not in the pipeline that includes it.

The default stack size for a pipeline if
`VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR` is not provided
is computed as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-tracing-pipeline-stack"
target="_blank" rel="noopener">Ray Tracing Pipeline Stack</a>.

If a
[VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)
structure is present in the `pNext` chain,
[VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)::`flags`
from that structure is used instead of `flags` from this structure.

Valid Usage

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-None-09497"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-None-09497"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-None-09497  
  If the `pNext` chain does not include a
  [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)
  structure, `flags` must be a valid combination of
  [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html) values

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-07984"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-07984"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-07984  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and
  `basePipelineIndex` is -1, `basePipelineHandle` **must** be a valid
  ray tracing `VkPipeline` handle

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-07985"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-07985"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-07985  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and
  `basePipelineHandle` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `basePipelineIndex` **must** be a valid index into the calling
  command’s `pCreateInfos` parameter

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-07986"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-07986"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-07986  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag,
  `basePipelineIndex` **must** be -1 or `basePipelineHandle` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-layout-07987"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-layout-07987"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-layout-07987  
  If a push constant block is declared in a shader, a push constant
  range in `layout` **must** match both the shader stage and range

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-layout-07988"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-layout-07988"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-layout-07988  
  If a [resource variables](#interfaces-resources) is declared in a
  shader, a descriptor slot in `layout` **must** match the shader stage

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-layout-07990"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-layout-07990"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-layout-07990  
  If a [resource variables](#interfaces-resources) is declared in a
  shader, and the descriptor type is not
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, a descriptor slot in `layout`
  **must** match the descriptor type

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-layout-07991"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-layout-07991"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-layout-07991  
  If a [resource variables](#interfaces-resources) is declared in a
  shader as an array, a descriptor slot in `layout` **must** match the
  descriptor count

<!-- -->

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-pStages-03426"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pStages-03426"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pStages-03426  
  The shader code for the entry points identified by `pStages`, and the
  rest of the state identified by this structure **must** adhere to the
  pipeline linking rules described in the [Shader
  Interfaces](#interfaces) chapter

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-layout-03428"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-layout-03428"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-layout-03428  
  The number of resources in `layout` accessible to each shader stage
  that is used by the pipeline **must** be less than or equal to
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxPerStageResources`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-02904"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-02904"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-02904  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`

- <a
  href="#VUID-VkRayTracingPipelineCreateInfoKHR-pipelineCreationCacheControl-02905"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pipelineCreationCacheControl-02905"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pipelineCreationCacheControl-02905  
  If the
  [`pipelineCreationCacheControl`](#features-pipelineCreationCacheControl)
  feature is not enabled, `flags` **must** not include
  `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` or
  `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-stage-03425"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-stage-03425"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-stage-03425  
  If `flags` does not include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`, the
  `stage` member of at least one element of `pStages`, including those
  implicitly added by `pLibraryInfo`, **must** be
  `VK_SHADER_STAGE_RAYGEN_BIT_KHR`

- <a
  href="#VUID-VkRayTracingPipelineCreateInfoKHR-maxPipelineRayRecursionDepth-03589"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-maxPipelineRayRecursionDepth-03589"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-maxPipelineRayRecursionDepth-03589  
  `maxPipelineRayRecursionDepth` **must** be less than or equal to
  [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`maxRayRecursionDepth`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-03465"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-03465"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-03465  
  If `flags` includes `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`,
  `pLibraryInterface` **must** not be `NULL`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03590"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03590"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03590  
  If `pLibraryInfo` is not `NULL` and its `libraryCount` member is
  greater than `0`, `pLibraryInterface` **must** not be `NULL`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraries-03591"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pLibraries-03591"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pLibraries-03591  
  Each element of `pLibraryInfo->pLibraries` **must** have been created
  with the value of `maxPipelineRayRecursionDepth` equal to that in this
  pipeline

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03592"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03592"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03592  
  If `pLibraryInfo` is not `NULL`, each element of its `pLibraries`
  member **must** have been created with a `layout` that is compatible
  with the `layout` in this pipeline

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03593"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03593"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03593  
  If `pLibraryInfo` is not `NULL`, each element of its `pLibraries`
  member **must** have been created with values of the
  `maxPipelineRayPayloadSize` and `maxPipelineRayHitAttributeSize`
  members of `pLibraryInterface` equal to those in this pipeline

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-03594"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-03594"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-03594  
  If `flags` includes
  `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`,
  each element of `pLibraryInfo->pLibraries` **must** have been created
  with the
  `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`
  bit set

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-04718"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-04718"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-04718  
  If `flags` includes
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`, each element of
  `pLibraryInfo->pLibraries` **must** have been created with the
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR` bit set

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-04719"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-04719"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-04719  
  If `flags` includes
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`, each element
  of `pLibraryInfo->pLibraries` **must** have been created with the
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR` bit set

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-04720"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-04720"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-04720  
  If `flags` includes
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`, each
  element of `pLibraryInfo->pLibraries` **must** have been created with
  the `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`
  bit set

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-04721"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-04721"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-04721  
  If `flags` includes
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`,
  each element of `pLibraryInfo->pLibraries` **must** have been created
  with the
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`
  bit set

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-04722"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-04722"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-04722  
  If `flags` includes
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`,
  each element of `pLibraryInfo->pLibraries` **must** have been created
  with the
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`
  bit set

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-04723"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-04723"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-04723  
  If `flags` includes
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`, each
  element of `pLibraryInfo->pLibraries` **must** have been created with
  the `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR` bit
  set

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03595"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03595"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-03595  
  If the [`VK_KHR_pipeline_library`](VK_KHR_pipeline_library.html)
  extension is not enabled, `pLibraryInfo` and `pLibraryInterface`
  **must** be `NULL`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-03470"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-03470"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-03470  
  If `flags` includes
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`, for
  any element of `pGroups` with a `type` of
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR` or
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR`, the
  `anyHitShader` of that element **must** not be `VK_SHADER_UNUSED_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-03471"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-03471"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-03471  
  If `flags` includes
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`,
  for any element of `pGroups` with a `type` of
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR` or
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR`, the
  `closestHitShader` of that element **must** not be
  `VK_SHADER_UNUSED_KHR`

- <a
  href="#VUID-VkRayTracingPipelineCreateInfoKHR-rayTraversalPrimitiveCulling-03596"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-rayTraversalPrimitiveCulling-03596"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-rayTraversalPrimitiveCulling-03596  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-rayTraversalPrimitiveCulling"
  target="_blank"
  rel="noopener"><code>rayTraversalPrimitiveCulling</code></a> feature
  is not enabled, `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`

- <a
  href="#VUID-VkRayTracingPipelineCreateInfoKHR-rayTraversalPrimitiveCulling-03597"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-rayTraversalPrimitiveCulling-03597"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-rayTraversalPrimitiveCulling-03597  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-rayTraversalPrimitiveCulling"
  target="_blank"
  rel="noopener"><code>rayTraversalPrimitiveCulling</code></a> feature
  is not enabled, `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-06546"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-06546"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-06546  
  `flags` **must** not include both
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR` and
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-03598"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-03598"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-03598  
  If `flags` includes
  `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`,
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-rayTracingPipelineShaderGroupHandleCaptureReplay"
  target="_blank"
  rel="noopener"><code>rayTracingPipelineShaderGroupHandleCaptureReplay</code></a>
  **must** be enabled

- <a
  href="#VUID-VkRayTracingPipelineCreateInfoKHR-rayTracingPipelineShaderGroupHandleCaptureReplay-03599"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-rayTracingPipelineShaderGroupHandleCaptureReplay-03599"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-rayTracingPipelineShaderGroupHandleCaptureReplay-03599  
  If
  [VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html)::`rayTracingPipelineShaderGroupHandleCaptureReplay`
  is `VK_TRUE` and the `pShaderGroupCaptureReplayHandle` member of any
  element of `pGroups` is not `NULL`, `flags` **must** include
  `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-07999"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-07999"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-07999  
  If `pLibraryInfo` is `NULL` or its `libraryCount` is `0`, `stageCount`
  **must** not be `0`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-08700"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-08700"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-08700  
  If `flags` does not include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` and
  either `pLibraryInfo` is `NULL` or its `libraryCount` is `0`,
  `groupCount` **must** not be `0`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-pDynamicStates-03602"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pDynamicStates-03602"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pDynamicStates-03602  
  Any element of the `pDynamicStates` member of `pDynamicState` **must**
  be `VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR`

- <a
  href="#VUID-VkRayTracingPipelineCreateInfoKHR-pipelineStageCreationFeedbackCount-06652"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pipelineStageCreationFeedbackCount-06652"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pipelineStageCreationFeedbackCount-06652  
  If
  [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackCreateInfo.html)::`pipelineStageCreationFeedbackCount`
  is not `0`, it **must** be equal to `stageCount`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-stage-06899"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-stage-06899"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-stage-06899  
  The `stage` value in all `pStages` elements **must** be one of
  `VK_SHADER_STAGE_RAYGEN_BIT_KHR`, `VK_SHADER_STAGE_ANY_HIT_BIT_KHR`,
  `VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR`, `VK_SHADER_STAGE_MISS_BIT_KHR`,
  `VK_SHADER_STAGE_INTERSECTION_BIT_KHR`, or
  `VK_SHADER_STAGE_CALLABLE_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-07403"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-07403"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-07403  
  If `flags` includes
  `VK_PIPELINE_CREATE_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT`, each
  element of `pLibraryInfo->pLibraries` **must** have been created with
  the `VK_PIPELINE_CREATE_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT` bit set

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-flags-08701"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-flags-08701"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-flags-08701  
  If `flags` includes
  `VK_PIPELINE_CREATE_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV`, each
  element of `pLibraryInfo->pLibraries` **must** have been created with
  the `VK_PIPELINE_CREATE_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV` bit
  set

Valid Usage (Implicit)

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-sType-sType"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-sType-sType"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CREATE_INFO_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-pNext-pNext"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pNext-pNext"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html),
  [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackCreateInfo.html),
  or
  [VkPipelineRobustnessCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessCreateInfoEXT.html)

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-sType-unique"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-sType-unique"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-pStages-parameter"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pStages-parameter"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pStages-parameter  
  If `stageCount` is not `0`, `pStages` **must** be a valid pointer to
  an array of `stageCount` valid
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  structures

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-pGroups-parameter"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pGroups-parameter"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pGroups-parameter  
  If `groupCount` is not `0`, `pGroups` **must** be a valid pointer to
  an array of `groupCount` valid
  [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoKHR.html)
  structures

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-parameter"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-parameter"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInfo-parameter  
  If `pLibraryInfo` is not `NULL`, `pLibraryInfo` **must** be a valid
  pointer to a valid
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)
  structure

- <a
  href="#VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInterface-parameter"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInterface-parameter"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pLibraryInterface-parameter  
  If `pLibraryInterface` is not `NULL`, `pLibraryInterface` **must** be
  a valid pointer to a valid
  [VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html)
  structure

- <a
  href="#VUID-VkRayTracingPipelineCreateInfoKHR-pDynamicState-parameter"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-pDynamicState-parameter"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-pDynamicState-parameter  
  If `pDynamicState` is not `NULL`, `pDynamicState` **must** be a valid
  pointer to a valid
  [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)
  structure

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-layout-parameter"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-layout-parameter"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)
  handle

- <a href="#VUID-VkRayTracingPipelineCreateInfoKHR-commonparent"
  id="VUID-VkRayTracingPipelineCreateInfoKHR-commonparent"></a>
  VUID-VkRayTracingPipelineCreateInfoKHR-commonparent  
  Both of `basePipelineHandle`, and `layout` that are valid handles of
  non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags.html),
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html),
[VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
[VkRayTracingPipelineInterfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineInterfaceCreateInfoKHR.html),
[VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRayTracingPipelineCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
