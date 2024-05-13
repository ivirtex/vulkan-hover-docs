# VkPipelineShaderStageCreateInfo(3) Manual Page

## Name

VkPipelineShaderStageCreateInfo - Structure specifying parameters of a
newly created pipeline shader stage



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineShaderStageCreateInfo` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkPipelineShaderStageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateFlagBits.html)
  specifying how the pipeline shader stage will be generated.

- `stage` is a [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) value
  specifying a single pipeline stage.

- `module` is optionally a [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html) object
  containing the shader code for this stage.

- `pName` is a pointer to a null-terminated UTF-8 string specifying the
  entry point name of the shader for this stage.

- `pSpecializationInfo` is a pointer to a
  [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSpecializationInfo.html) structure, as
  described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-specialization-constants"
  target="_blank" rel="noopener">Specialization Constants</a>, or
  `NULL`.

## <a href="#_description" class="anchor"></a>Description

If `module` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the shader
code used by the pipeline is defined by `module`. If `module` is
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the shader code is defined by the
chained [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) if
present.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderModuleIdentifier"
target="_blank" rel="noopener"><code>shaderModuleIdentifier</code></a>
feature is enabled, applications **can** omit shader code for `stage`
and instead provide a module identifier. This is done by including a
[VkPipelineShaderStageModuleIdentifierCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageModuleIdentifierCreateInfoEXT.html)
struct with `identifierSize` not equal to 0 in the `pNext` chain. A
shader stage created in this way is equivalent to one created using a
shader module with the same identifier. The identifier allows an
implementation to look up a pipeline without consuming a valid SPIR-V
module. If a pipeline is not found, pipeline compilation is not possible
and the implementation **must** fail as specified by
`VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT`.

When an identifier is used in lieu of a shader module, implementations
**may** fail pipeline compilation with `VK_PIPELINE_COMPILE_REQUIRED`
for any reason.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The rationale for the relaxed requirement on implementations to
return a pipeline with <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageModuleIdentifierCreateInfoEXT.html">VkPipelineShaderStageModuleIdentifierCreateInfoEXT</a>
is that layers or tools may intercept pipeline creation calls and
require the full SPIR-V context to operate correctly. ICDs are not
expected to fail pipeline compilation if the pipeline exists in a cache
somewhere.</p></td>
</tr>
</tbody>
</table>

Applications **can** use identifiers when creating pipelines with
`VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`. When creating such pipelines,
`VK_SUCCESS` **may** be returned, but subsequently fail when referencing
the pipeline in a
[VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)
struct. Applications **must** allow pipeline compilation to fail during
link steps with
`VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` as it **may**
not be possible to determine if a pipeline **can** be created from
identifiers until the link step.

Valid Usage

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-00704"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-00704"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-00704  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-geometryShader"
  target="_blank" rel="noopener"><code>geometryShader</code></a> feature
  is not enabled, `stage` **must** not be `VK_SHADER_STAGE_GEOMETRY_BIT`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-00705"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-00705"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-00705  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-tessellationShader"
  target="_blank" rel="noopener"><code>tessellationShader</code></a>
  feature is not enabled, `stage` **must** not be
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` or
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-02091"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-02091"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-02091  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-meshShader"
  target="_blank" rel="noopener"><code>meshShaders</code></a> feature is
  not enabled, `stage` **must** not be `VK_SHADER_STAGE_MESH_BIT_EXT`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-02092"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-02092"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-02092  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-taskShader"
  target="_blank" rel="noopener"><code>taskShaders</code></a> feature is
  not enabled, `stage` **must** not be `VK_SHADER_STAGE_TASK_BIT_EXT`

- <a
  href="#VUID-VkPipelineShaderStageCreateInfo-clustercullingShader-07813"
  id="VUID-VkPipelineShaderStageCreateInfo-clustercullingShader-07813"></a>
  VUID-VkPipelineShaderStageCreateInfo-clustercullingShader-07813  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-clustercullingShader"
  target="_blank" rel="noopener"><code>clustercullingShader</code></a>
  feature is not enabled, `stage` **must** not be
  `VK_SHADER_STAGE_CLUSTER_CULLING_BIT_HUAWEI`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-00706"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-00706"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-00706  
  `stage` **must** not be `VK_SHADER_STAGE_ALL_GRAPHICS`, or
  `VK_SHADER_STAGE_ALL`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-pName-00707"
  id="VUID-VkPipelineShaderStageCreateInfo-pName-00707"></a>
  VUID-VkPipelineShaderStageCreateInfo-pName-00707  
  `pName` **must** be the name of an `OpEntryPoint` in `module` with an
  execution model that matches `stage`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-maxClipDistances-00708"
  id="VUID-VkPipelineShaderStageCreateInfo-maxClipDistances-00708"></a>
  VUID-VkPipelineShaderStageCreateInfo-maxClipDistances-00708  
  If the identified entry point includes any variable in its interface
  that is declared with the `ClipDistance` `BuiltIn` decoration, that
  variable **must** not have an array size greater than
  `VkPhysicalDeviceLimits`::`maxClipDistances`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-maxCullDistances-00709"
  id="VUID-VkPipelineShaderStageCreateInfo-maxCullDistances-00709"></a>
  VUID-VkPipelineShaderStageCreateInfo-maxCullDistances-00709  
  If the identified entry point includes any variable in its interface
  that is declared with the `CullDistance` `BuiltIn` decoration, that
  variable **must** not have an array size greater than
  `VkPhysicalDeviceLimits`::`maxCullDistances`

- <a
  href="#VUID-VkPipelineShaderStageCreateInfo-maxCombinedClipAndCullDistances-00710"
  id="VUID-VkPipelineShaderStageCreateInfo-maxCombinedClipAndCullDistances-00710"></a>
  VUID-VkPipelineShaderStageCreateInfo-maxCombinedClipAndCullDistances-00710  
  If the identified entry point includes variables in its interface that
  are declared with the `ClipDistance` `BuiltIn` decoration and
  variables in its interface that are declared with the `CullDistance`
  `BuiltIn` decoration, those variables **must** not have array sizes
  which sum to more than
  `VkPhysicalDeviceLimits`::`maxCombinedClipAndCullDistances`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-maxSampleMaskWords-00711"
  id="VUID-VkPipelineShaderStageCreateInfo-maxSampleMaskWords-00711"></a>
  VUID-VkPipelineShaderStageCreateInfo-maxSampleMaskWords-00711  
  If the identified entry point includes any variable in its interface
  that is declared with the `SampleMask` `BuiltIn` decoration, that
  variable **must** not have an array size greater than
  `VkPhysicalDeviceLimits`::`maxSampleMaskWords`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-00713"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-00713"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-00713  
  If `stage` is `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` or
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, and the identified
  entry point has an `OpExecutionMode` instruction specifying a patch
  size with `OutputVertices`, the patch size **must** be greater than
  `0` and less than or equal to
  `VkPhysicalDeviceLimits`::`maxTessellationPatchSize`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-00714"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-00714"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-00714  
  If `stage` is `VK_SHADER_STAGE_GEOMETRY_BIT`, the identified entry
  point **must** have an `OpExecutionMode` instruction specifying a
  maximum output vertex count that is greater than `0` and less than or
  equal to `VkPhysicalDeviceLimits`::`maxGeometryOutputVertices`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-00715"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-00715"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-00715  
  If `stage` is `VK_SHADER_STAGE_GEOMETRY_BIT`, the identified entry
  point **must** have an `OpExecutionMode` instruction specifying an
  invocation count that is greater than `0` and less than or equal to
  `VkPhysicalDeviceLimits`::`maxGeometryShaderInvocations`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-02596"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-02596"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-02596  
  If `stage` is either `VK_SHADER_STAGE_VERTEX_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, or
  `VK_SHADER_STAGE_GEOMETRY_BIT`, and the identified entry point writes
  to `Layer` for any primitive, it **must** write the same value to
  `Layer` for all vertices of a given primitive

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-02597"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-02597"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-02597  
  If `stage` is either `VK_SHADER_STAGE_VERTEX_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, or
  `VK_SHADER_STAGE_GEOMETRY_BIT`, and the identified entry point writes
  to `ViewportIndex` for any primitive, it **must** write the same value
  to `ViewportIndex` for all vertices of a given primitive

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-06685"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-06685"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-06685  
  If `stage` is `VK_SHADER_STAGE_FRAGMENT_BIT`, and the identified entry
  point writes to `FragDepth` in any execution path, all execution paths
  that are not exclusive to helper invocations **must** either discard
  the fragment, or write or initialize the value of `FragDepth`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-06686"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-06686"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-06686  
  If `stage` is `VK_SHADER_STAGE_FRAGMENT_BIT`, and the identified entry
  point writes to `FragStencilRefEXT` in any execution path, all
  execution paths that are not exclusive to helper invocations **must**
  either discard the fragment, or write or initialize the value of
  `FragStencilRefEXT`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-flags-02784"
  id="VUID-VkPipelineShaderStageCreateInfo-flags-02784"></a>
  VUID-VkPipelineShaderStageCreateInfo-flags-02784  
  If `flags` has the
  `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` flag
  set, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-subgroupSizeControl"
  target="_blank" rel="noopener"><code>subgroupSizeControl</code></a>
  feature **must** be enabled

- <a href="#VUID-VkPipelineShaderStageCreateInfo-flags-02785"
  id="VUID-VkPipelineShaderStageCreateInfo-flags-02785"></a>
  VUID-VkPipelineShaderStageCreateInfo-flags-02785  
  If `flags` has the
  `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` flag set,
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-computeFullSubgroups"
  target="_blank" rel="noopener"><code>computeFullSubgroups</code></a>
  feature **must** be enabled

- <a href="#VUID-VkPipelineShaderStageCreateInfo-flags-08988"
  id="VUID-VkPipelineShaderStageCreateInfo-flags-08988"></a>
  VUID-VkPipelineShaderStageCreateInfo-flags-08988  
  If `flags` includes
  `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT`, `stage`
  **must** be one of `VK_SHADER_STAGE_MESH_BIT_EXT`,
  `VK_SHADER_STAGE_TASK_BIT_EXT`, or `VK_SHADER_STAGE_COMPUTE_BIT`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-pNext-02754"
  id="VUID-VkPipelineShaderStageCreateInfo-pNext-02754"></a>
  VUID-VkPipelineShaderStageCreateInfo-pNext-02754  
  If a
  [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)
  structure is included in the `pNext` chain, `flags` **must** not have
  the `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT`
  flag set

- <a href="#VUID-VkPipelineShaderStageCreateInfo-pNext-02755"
  id="VUID-VkPipelineShaderStageCreateInfo-pNext-02755"></a>
  VUID-VkPipelineShaderStageCreateInfo-pNext-02755  
  If a
  [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)
  structure is included in the `pNext` chain, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-subgroupSizeControl"
  target="_blank" rel="noopener"><code>subgroupSizeControl</code></a>
  feature **must** be enabled, and `stage` **must** be a valid bit
  specified in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-requiredSubgroupSizeStages"
  target="_blank"
  rel="noopener"><code>requiredSubgroupSizeStages</code></a>

- <a href="#VUID-VkPipelineShaderStageCreateInfo-pNext-02756"
  id="VUID-VkPipelineShaderStageCreateInfo-pNext-02756"></a>
  VUID-VkPipelineShaderStageCreateInfo-pNext-02756  
  If a
  [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)
  structure is included in the `pNext` chain and `stage` is
  `VK_SHADER_STAGE_COMPUTE_BIT`, `VK_SHADER_STAGE_MESH_BIT_EXT`, or
  `VK_SHADER_STAGE_TASK_BIT_EXT`, the local workgroup size of the shader
  **must** be less than or equal to the product of
  [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)::`requiredSubgroupSize`
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxComputeWorkgroupSubgroups"
  target="_blank"
  rel="noopener"><code>maxComputeWorkgroupSubgroups</code></a>

- <a href="#VUID-VkPipelineShaderStageCreateInfo-pNext-02757"
  id="VUID-VkPipelineShaderStageCreateInfo-pNext-02757"></a>
  VUID-VkPipelineShaderStageCreateInfo-pNext-02757  
  If a
  [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)
  structure is included in the `pNext` chain, and `flags` has the
  `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` flag set,
  the local workgroup size in the X dimension of the pipeline **must**
  be a multiple of
  [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)::`requiredSubgroupSize`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-flags-02758"
  id="VUID-VkPipelineShaderStageCreateInfo-flags-02758"></a>
  VUID-VkPipelineShaderStageCreateInfo-flags-02758  
  If `flags` has both the
  `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` and
  `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT`
  flags set, the local workgroup size in the X dimension of the pipeline
  **must** be a multiple of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxSubgroupSize"
  target="_blank" rel="noopener"><code>maxSubgroupSize</code></a>

- <a href="#VUID-VkPipelineShaderStageCreateInfo-flags-02759"
  id="VUID-VkPipelineShaderStageCreateInfo-flags-02759"></a>
  VUID-VkPipelineShaderStageCreateInfo-flags-02759  
  If `flags` has the
  `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` flag set
  and `flags` does not have the
  `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` flag
  set and no
  [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)
  structure is included in the `pNext` chain, the local workgroup size
  in the X dimension of the pipeline **must** be a multiple of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-subgroup-size"
  target="_blank" rel="noopener"><code>subgroupSize</code></a>

- <a href="#VUID-VkPipelineShaderStageCreateInfo-module-08987"
  id="VUID-VkPipelineShaderStageCreateInfo-module-08987"></a>
  VUID-VkPipelineShaderStageCreateInfo-module-08987  
  If `module` uses the `OpTypeCooperativeMatrixKHR` instruction with a
  `Scope` equal to `Subgroup`, then the local workgroup size in the X
  dimension of the pipeline **must** be a multiple of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-subgroup-size"
  target="_blank" rel="noopener"><code>subgroupSize</code></a>

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-08771"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-08771"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-08771  
  If a shader module identifier is not specified for this `stage`,
  `module` **must** be a valid [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html) if
  none of the following features are enabled:

  - <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-graphicsPipelineLibrary"
    target="_blank" rel="noopener"><code>graphicsPipelineLibrary</code></a>

  - <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance5"
    target="_blank" rel="noopener"><code>maintenance5</code></a>

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-06845"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-06845"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-06845  
  If a shader module identifier is not specified for this `stage`,
  `module` **must** be a valid [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html), or
  there **must** be a valid
  [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) structure in
  the `pNext` chain

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-06844"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-06844"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-06844  
  If a shader module identifier is specified for this `stage`, a
  [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) structure
  **must** not be present in the `pNext` chain

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-06848"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-06848"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-06848  
  If a shader module identifier is specified for this `stage`, `module`
  **must** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a
  href="#VUID-VkPipelineShaderStageCreateInfo-pSpecializationInfo-06849"
  id="VUID-VkPipelineShaderStageCreateInfo-pSpecializationInfo-06849"></a>
  VUID-VkPipelineShaderStageCreateInfo-pSpecializationInfo-06849  
  If a shader module identifier is not specified, the shader code used
  by the pipeline **must** be valid as described by the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirv-spec"
  target="_blank" rel="noopener">Khronos SPIR-V Specification</a> after
  applying the specializations provided in `pSpecializationInfo`, if
  any, and then converting all specialization constants into fixed
  constants

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineShaderStageCreateInfo-sType-sType"
  id="VUID-VkPipelineShaderStageCreateInfo-sType-sType"></a>
  VUID-VkPipelineShaderStageCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO`

- <a href="#VUID-VkPipelineShaderStageCreateInfo-pNext-pNext"
  id="VUID-VkPipelineShaderStageCreateInfo-pNext-pNext"></a>
  VUID-VkPipelineShaderStageCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDebugUtilsObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectNameInfoEXT.html),
  [VkPipelineRobustnessCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessCreateInfoEXT.html),
  [VkPipelineShaderStageModuleIdentifierCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageModuleIdentifierCreateInfoEXT.html),
  [VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html),
  [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html),
  [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html), or
  [VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html)

- <a href="#VUID-VkPipelineShaderStageCreateInfo-sType-unique"
  id="VUID-VkPipelineShaderStageCreateInfo-sType-unique"></a>
  VUID-VkPipelineShaderStageCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkPipelineShaderStageCreateInfo-flags-parameter"
  id="VUID-VkPipelineShaderStageCreateInfo-flags-parameter"></a>
  VUID-VkPipelineShaderStageCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkPipelineShaderStageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateFlagBits.html)
  values

- <a href="#VUID-VkPipelineShaderStageCreateInfo-stage-parameter"
  id="VUID-VkPipelineShaderStageCreateInfo-stage-parameter"></a>
  VUID-VkPipelineShaderStageCreateInfo-stage-parameter  
  `stage` **must** be a valid
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) value

- <a href="#VUID-VkPipelineShaderStageCreateInfo-module-parameter"
  id="VUID-VkPipelineShaderStageCreateInfo-module-parameter"></a>
  VUID-VkPipelineShaderStageCreateInfo-module-parameter  
  If `module` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `module`
  **must** be a valid [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html) handle

- <a href="#VUID-VkPipelineShaderStageCreateInfo-pName-parameter"
  id="VUID-VkPipelineShaderStageCreateInfo-pName-parameter"></a>
  VUID-VkPipelineShaderStageCreateInfo-pName-parameter  
  `pName` **must** be a null-terminated UTF-8 string

- <a
  href="#VUID-VkPipelineShaderStageCreateInfo-pSpecializationInfo-parameter"
  id="VUID-VkPipelineShaderStageCreateInfo-pSpecializationInfo-parameter"></a>
  VUID-VkPipelineShaderStageCreateInfo-pSpecializationInfo-parameter  
  If `pSpecializationInfo` is not `NULL`, `pSpecializationInfo` **must**
  be a valid pointer to a valid
  [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSpecializationInfo.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html),
[VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html),
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkGraphicsShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsShaderGroupCreateInfoNV.html),
[VkPipelineShaderStageCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateFlags.html),
[VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html),
[VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html),
[VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html),
[VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html),
[VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSpecializationInfo.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineShaderStageCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
