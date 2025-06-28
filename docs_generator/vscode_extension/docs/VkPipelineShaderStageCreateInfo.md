# VkPipelineShaderStageCreateInfo(3) Manual Page

## Name

VkPipelineShaderStageCreateInfo - Structure specifying parameters of a newly created pipeline shader stage



## [](#_c_specification)C Specification

The `VkPipelineShaderStageCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPipelineShaderStageCreateInfo {
    VkStructureType                     sType;
    const void*                         pNext;
    VkPipelineShaderStageCreateFlags    flags;
    VkShaderStageFlagBits               stage;
    VkShaderModule                      module;
    const char*                         pName;
    const VkSpecializationInfo*         pSpecializationInfo;
} VkPipelineShaderStageCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkPipelineShaderStageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateFlagBits.html) specifying how the pipeline shader stage will be generated.
- `stage` is a [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) value specifying a single pipeline stage.
- `module` is optionally a [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html) object containing the shader code for this stage. The implementation **must** not access this object outside of the duration of the command this structure is passed to.
- `pName` is a pointer to a null-terminated UTF-8 string specifying the entry point name of the shader for this stage.
- `pSpecializationInfo` is a pointer to a [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSpecializationInfo.html) structure, as described in [Specialization Constants](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-specialization-constants), or `NULL`.

## [](#_description)Description

If `module` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the shader code used by the pipeline is defined by `module`. If `module` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the shader code is defined by the chained [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) if present.

If the [`shaderModuleIdentifier`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderModuleIdentifier) feature is enabled, applications **can** omit shader code for `stage` and instead provide a module identifier. This is done by including a [VkPipelineShaderStageModuleIdentifierCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageModuleIdentifierCreateInfoEXT.html) structure with `identifierSize` not equal to 0 in the `pNext` chain. A shader stage created in this way is equivalent to one created using a shader module with the same identifier. The identifier allows an implementation to look up a pipeline without consuming a valid SPIR-V module. If a pipeline is not found, pipeline compilation is not possible and the implementation **must** fail as specified by `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT`.

When an identifier is used in lieu of a shader module, implementations **may** fail pipeline compilation with `VK_PIPELINE_COMPILE_REQUIRED` for any reason.

Note

The rationale for the relaxed requirement on implementations to return a pipeline with [VkPipelineShaderStageModuleIdentifierCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageModuleIdentifierCreateInfoEXT.html) is that layers or tools may intercept pipeline creation calls and require the full SPIR-V context to operate correctly. ICDs are not expected to fail pipeline compilation if the pipeline exists in a cache somewhere.

Applications **can** use identifiers when creating pipelines with `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`. When creating such pipelines, `VK_SUCCESS` **may** be returned, but subsequently fail when referencing the pipeline in a [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLibraryCreateInfoKHR.html) struct. Applications **must** allow pipeline compilation to fail during link steps with `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` as it **may** not be possible to determine if a pipeline **can** be created from identifiers until the link step.

Valid Usage

- [](#VUID-VkPipelineShaderStageCreateInfo-stage-00704)VUID-VkPipelineShaderStageCreateInfo-stage-00704  
  If the [`geometryShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-geometryShader) feature is not enabled, `stage` **must** not be `VK_SHADER_STAGE_GEOMETRY_BIT`
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-00705)VUID-VkPipelineShaderStageCreateInfo-stage-00705  
  If the [`tessellationShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tessellationShader) feature is not enabled, `stage` **must** not be `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` or `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-02091)VUID-VkPipelineShaderStageCreateInfo-stage-02091  
  If the [`meshShaders`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-meshShader) feature is not enabled, `stage` **must** not be `VK_SHADER_STAGE_MESH_BIT_EXT`
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-02092)VUID-VkPipelineShaderStageCreateInfo-stage-02092  
  If the [`taskShaders`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-taskShader) feature is not enabled, `stage` **must** not be `VK_SHADER_STAGE_TASK_BIT_EXT`
- [](#VUID-VkPipelineShaderStageCreateInfo-clustercullingShader-07813)VUID-VkPipelineShaderStageCreateInfo-clustercullingShader-07813  
  If the [`clustercullingShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-clustercullingShader) feature is not enabled, `stage` **must** not be `VK_SHADER_STAGE_CLUSTER_CULLING_BIT_HUAWEI`
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-00706)VUID-VkPipelineShaderStageCreateInfo-stage-00706  
  `stage` **must** not be `VK_SHADER_STAGE_ALL_GRAPHICS`, or `VK_SHADER_STAGE_ALL`
- [](#VUID-VkPipelineShaderStageCreateInfo-pName-00707)VUID-VkPipelineShaderStageCreateInfo-pName-00707  
  `pName` **must** be the name of an `OpEntryPoint` in `module` with an execution model that matches `stage`
- [](#VUID-VkPipelineShaderStageCreateInfo-maxClipDistances-00708)VUID-VkPipelineShaderStageCreateInfo-maxClipDistances-00708  
  If the identified entry point includes any variable in its interface that is declared with the `ClipDistance` `BuiltIn` decoration, that variable **must** not have an array size greater than `VkPhysicalDeviceLimits`::`maxClipDistances`
- [](#VUID-VkPipelineShaderStageCreateInfo-maxCullDistances-00709)VUID-VkPipelineShaderStageCreateInfo-maxCullDistances-00709  
  If the identified entry point includes any variable in its interface that is declared with the `CullDistance` `BuiltIn` decoration, that variable **must** not have an array size greater than `VkPhysicalDeviceLimits`::`maxCullDistances`
- [](#VUID-VkPipelineShaderStageCreateInfo-maxCombinedClipAndCullDistances-00710)VUID-VkPipelineShaderStageCreateInfo-maxCombinedClipAndCullDistances-00710  
  If the identified entry point includes variables in its interface that are declared with the `ClipDistance` `BuiltIn` decoration and variables in its interface that are declared with the `CullDistance` `BuiltIn` decoration, those variables **must** not have array sizes which sum to more than `VkPhysicalDeviceLimits`::`maxCombinedClipAndCullDistances`
- [](#VUID-VkPipelineShaderStageCreateInfo-maxSampleMaskWords-00711)VUID-VkPipelineShaderStageCreateInfo-maxSampleMaskWords-00711  
  If the identified entry point includes any variable in its interface that is declared with the `SampleMask` `BuiltIn` decoration, that variable **must** not have an array size greater than `VkPhysicalDeviceLimits`::`maxSampleMaskWords`
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-00713)VUID-VkPipelineShaderStageCreateInfo-stage-00713  
  If `stage` is `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` or `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, and the identified entry point has an `OpExecutionMode` instruction specifying a patch size with `OutputVertices`, the patch size **must** be greater than `0` and less than or equal to `VkPhysicalDeviceLimits`::`maxTessellationPatchSize`
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-00714)VUID-VkPipelineShaderStageCreateInfo-stage-00714  
  If `stage` is `VK_SHADER_STAGE_GEOMETRY_BIT`, the identified entry point **must** have an `OpExecutionMode` instruction specifying a maximum output vertex count that is greater than `0` and less than or equal to `VkPhysicalDeviceLimits`::`maxGeometryOutputVertices`
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-00715)VUID-VkPipelineShaderStageCreateInfo-stage-00715  
  If `stage` is `VK_SHADER_STAGE_GEOMETRY_BIT`, the identified entry point **must** have an `OpExecutionMode` instruction specifying an invocation count that is greater than `0` and less than or equal to `VkPhysicalDeviceLimits`::`maxGeometryShaderInvocations`
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-02596)VUID-VkPipelineShaderStageCreateInfo-stage-02596  
  If `stage` is either `VK_SHADER_STAGE_VERTEX_BIT`, `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`, `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, or `VK_SHADER_STAGE_GEOMETRY_BIT`, and the identified entry point writes to `Layer` for any primitive, it **must** write the same value to `Layer` for all vertices of a given primitive
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-02597)VUID-VkPipelineShaderStageCreateInfo-stage-02597  
  If `stage` is either `VK_SHADER_STAGE_VERTEX_BIT`, `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`, `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, or `VK_SHADER_STAGE_GEOMETRY_BIT`, and the identified entry point writes to `ViewportIndex` for any primitive, it **must** write the same value to `ViewportIndex` for all vertices of a given primitive
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-06685)VUID-VkPipelineShaderStageCreateInfo-stage-06685  
  If `stage` is `VK_SHADER_STAGE_FRAGMENT_BIT`, and the identified entry point writes to `FragDepth` in any execution path, all execution paths that are not exclusive to helper invocations **must** either discard the fragment, or write or initialize the value of `FragDepth`
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-06686)VUID-VkPipelineShaderStageCreateInfo-stage-06686  
  If `stage` is `VK_SHADER_STAGE_FRAGMENT_BIT`, and the identified entry point writes to `FragStencilRefEXT` in any execution path, all execution paths that are not exclusive to helper invocations **must** either discard the fragment, or write or initialize the value of `FragStencilRefEXT`
- [](#VUID-VkPipelineShaderStageCreateInfo-flags-02784)VUID-VkPipelineShaderStageCreateInfo-flags-02784  
  If `flags` has the `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` flag set, the [`subgroupSizeControl`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-subgroupSizeControl) feature **must** be enabled
- [](#VUID-VkPipelineShaderStageCreateInfo-flags-02785)VUID-VkPipelineShaderStageCreateInfo-flags-02785  
  If `flags` has the `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` flag set, the [`computeFullSubgroups`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-computeFullSubgroups) feature **must** be enabled
- [](#VUID-VkPipelineShaderStageCreateInfo-flags-08988)VUID-VkPipelineShaderStageCreateInfo-flags-08988  
  If `flags` includes `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT`, `stage` **must** be one of `VK_SHADER_STAGE_MESH_BIT_EXT`, `VK_SHADER_STAGE_TASK_BIT_EXT`, or `VK_SHADER_STAGE_COMPUTE_BIT`
- [](#VUID-VkPipelineShaderStageCreateInfo-pNext-02754)VUID-VkPipelineShaderStageCreateInfo-pNext-02754  
  If a [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html) structure is included in the `pNext` chain, `flags` **must** not have the `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` flag set
- [](#VUID-VkPipelineShaderStageCreateInfo-pNext-02755)VUID-VkPipelineShaderStageCreateInfo-pNext-02755  
  If a [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html) structure is included in the `pNext` chain, the [`subgroupSizeControl`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-subgroupSizeControl) feature **must** be enabled, and `stage` **must** be a valid bit specified in [`requiredSubgroupSizeStages`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-requiredSubgroupSizeStages)
- [](#VUID-VkPipelineShaderStageCreateInfo-pNext-02756)VUID-VkPipelineShaderStageCreateInfo-pNext-02756  
  If a [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html) structure is included in the `pNext` chain and `stage` is `VK_SHADER_STAGE_COMPUTE_BIT`, `VK_SHADER_STAGE_MESH_BIT_EXT`, or `VK_SHADER_STAGE_TASK_BIT_EXT`, the local workgroup size of the shader **must** be less than or equal to the product of [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)::`requiredSubgroupSize` and [`maxComputeWorkgroupSubgroups`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxComputeWorkgroupSubgroups)
- [](#VUID-VkPipelineShaderStageCreateInfo-pNext-02757)VUID-VkPipelineShaderStageCreateInfo-pNext-02757  
  If a [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html) structure is included in the `pNext` chain, and `flags` has the `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` flag set, the local workgroup size in the X dimension of the pipeline **must** be a multiple of [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)::`requiredSubgroupSize`
- [](#VUID-VkPipelineShaderStageCreateInfo-flags-02758)VUID-VkPipelineShaderStageCreateInfo-flags-02758  
  If `flags` has both the `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` and `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` flags set, the local workgroup size in the X dimension of the pipeline **must** be a multiple of [`maxSubgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxSubgroupSize)
- [](#VUID-VkPipelineShaderStageCreateInfo-flags-02759)VUID-VkPipelineShaderStageCreateInfo-flags-02759  
  If `flags` has the `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` flag set and `flags` does not have the `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` flag set and no [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html) structure is included in the `pNext` chain, the local workgroup size in the X dimension of the pipeline **must** be a multiple of [`subgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-subgroupSize)
- [](#VUID-VkPipelineShaderStageCreateInfo-module-08987)VUID-VkPipelineShaderStageCreateInfo-module-08987  
  If `module` uses the `OpTypeCooperativeMatrixKHR` instruction with a `Scope` equal to `Subgroup`, then the local workgroup size in the X dimension of the pipeline **must** be a multiple of the [effective subgroup size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-sgs).
- [](#VUID-VkPipelineShaderStageCreateInfo-module-10169)VUID-VkPipelineShaderStageCreateInfo-module-10169  
  If `module` uses the `OpTypeCooperativeMatrixKHR` instruction with a `Scope` equal to `Workgroup`, then the local workgroup size in the X dimension of the pipeline **must** be a multiple of the [effective subgroup size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-sgs) and the total local workgroup size **must** be a power of two multiple of the [effective subgroup size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-sgs) and **must** be less than or equal to [cooperativeMatrixWorkgroupScopeMaxWorkgroupSize](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-cooperativeMatrixWorkgroupScopeMaxWorkgroupSize)
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-08771)VUID-VkPipelineShaderStageCreateInfo-stage-08771  
  If a shader module identifier is not specified for this `stage`, `module` **must** be a valid [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html) , or the `pNext` chain of the parent `Vk*CreateInfo` structure **must** set [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` to a value greater than `0`, if none of the following features are enabled:
  
  - [`graphicsPipelineLibrary`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-graphicsPipelineLibrary)
  - [`maintenance5`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance5)
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-06845)VUID-VkPipelineShaderStageCreateInfo-stage-06845  
  If a shader module identifier is not specified for this `stage`, `module` **must** be a valid [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html), or there **must** be a valid [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structure in the `pNext` chain , or the `pNext` chain of the parent `Vk*CreateInfo` structure **must** set [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` to a value greater than `0`,
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-06844)VUID-VkPipelineShaderStageCreateInfo-stage-06844  
  If a shader module identifier is specified for this `stage`, the `pNext` chain **must** not include a [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structure
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-06848)VUID-VkPipelineShaderStageCreateInfo-stage-06848  
  If a shader module identifier is specified for this `stage`, `module` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkPipelineShaderStageCreateInfo-pSpecializationInfo-06849)VUID-VkPipelineShaderStageCreateInfo-pSpecializationInfo-06849  
  If a shader module identifier is not specified, the shader code used by the pipeline **must** be valid as described by the [Khronos SPIR-V Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirv-spec) after applying the specializations provided in `pSpecializationInfo`, if any, and then converting all specialization constants into fixed constants

Valid Usage (Implicit)

- [](#VUID-VkPipelineShaderStageCreateInfo-sType-sType)VUID-VkPipelineShaderStageCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO`
- [](#VUID-VkPipelineShaderStageCreateInfo-pNext-pNext)VUID-VkPipelineShaderStageCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkDebugUtilsObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsObjectNameInfoEXT.html), [VkPipelineRobustnessCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessCreateInfo.html), [VkPipelineShaderStageModuleIdentifierCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageModuleIdentifierCreateInfoEXT.html), [VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html), [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html), [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html), or [VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html)
- [](#VUID-VkPipelineShaderStageCreateInfo-sType-unique)VUID-VkPipelineShaderStageCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPipelineShaderStageCreateInfo-flags-parameter)VUID-VkPipelineShaderStageCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkPipelineShaderStageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateFlagBits.html) values
- [](#VUID-VkPipelineShaderStageCreateInfo-stage-parameter)VUID-VkPipelineShaderStageCreateInfo-stage-parameter  
  `stage` **must** be a valid [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) value
- [](#VUID-VkPipelineShaderStageCreateInfo-module-parameter)VUID-VkPipelineShaderStageCreateInfo-module-parameter  
  If `module` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `module` **must** be a valid [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html) handle
- [](#VUID-VkPipelineShaderStageCreateInfo-pName-parameter)VUID-VkPipelineShaderStageCreateInfo-pName-parameter  
  `pName` **must** be a null-terminated UTF-8 string
- [](#VUID-VkPipelineShaderStageCreateInfo-pSpecializationInfo-parameter)VUID-VkPipelineShaderStageCreateInfo-pSpecializationInfo-parameter  
  If `pSpecializationInfo` is not `NULL`, `pSpecializationInfo` **must** be a valid pointer to a valid [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSpecializationInfo.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), [VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkGraphicsShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsShaderGroupCreateInfoNV.html), [VkPipelineShaderStageCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateFlags.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html), [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html), [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html), [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html), [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSpecializationInfo.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineShaderStageCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0