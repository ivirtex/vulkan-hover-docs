# VkShaderCreateInfoEXT(3) Manual Page

## Name

VkShaderCreateInfoEXT - Structure specifying parameters of a newly
created shader



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkShaderCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_shader_object
typedef struct VkShaderCreateInfoEXT {
    VkStructureType                 sType;
    const void*                     pNext;
    VkShaderCreateFlagsEXT          flags;
    VkShaderStageFlagBits           stage;
    VkShaderStageFlags              nextStage;
    VkShaderCodeTypeEXT             codeType;
    size_t                          codeSize;
    const void*                     pCode;
    const char*                     pName;
    uint32_t                        setLayoutCount;
    const VkDescriptorSetLayout*    pSetLayouts;
    uint32_t                        pushConstantRangeCount;
    const VkPushConstantRange*      pPushConstantRanges;
    const VkSpecializationInfo*     pSpecializationInfo;
} VkShaderCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateFlagBitsEXT.html) describing
  additional parameters of the shader.

- `stage` is a [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) value
  specifying a single shader stage.

- `nextStage` is a bitmask of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) specifying zero or
  stages which **may** be used as a logically next bound stage when
  drawing with the shader bound.

- `codeType` is a [VkShaderCodeTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCodeTypeEXT.html) value
  specifying the type of the shader code pointed to be `pCode`.

- `codeSize` is the size in bytes of the shader code pointed to be
  `pCode`.

- `pCode` is a pointer to the shader code to use to create the shader.

- `pName` is a pointer to a null-terminated UTF-8 string specifying the
  entry point name of the shader for this stage.

- `setLayoutCount` is the number of descriptor set layouts pointed to by
  `pSetLayouts`.

- `pSetLayouts` is a pointer to an array of
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) objects used by
  the shader stage.

- `pushConstantRangeCount` is the number of push constant ranges pointed
  to by `pPushConstantRanges`.

- `pPushConstantRanges` is a pointer to an array of
  [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantRange.html) structures used by the
  shader stage.

- `pSpecializationInfo` is a pointer to a
  [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSpecializationInfo.html) structure, as
  described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-specialization-constants"
  target="_blank" rel="noopener">Specialization Constants</a>, or
  `NULL`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkShaderCreateInfoEXT-codeSize-08735"
  id="VUID-VkShaderCreateInfoEXT-codeSize-08735"></a>
  VUID-VkShaderCreateInfoEXT-codeSize-08735  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `codeSize` **must**
  be a multiple of 4

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08736"
  id="VUID-VkShaderCreateInfoEXT-pCode-08736"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08736  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pCode` **must**
  point to valid SPIR-V code, formatted and packed as described by the
  [Khronos SPIR-V Specification](#spirv-spec)

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08737"
  id="VUID-VkShaderCreateInfoEXT-pCode-08737"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08737  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pCode` **must**
  adhere to the validation rules described by the [Validation Rules
  within a Module](#spirvenv-module-validation) section of the [SPIR-V
  Environment](#spirvenv-capabilities) appendix

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08738"
  id="VUID-VkShaderCreateInfoEXT-pCode-08738"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08738  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pCode` **must**
  declare the `Shader` capability for SPIR-V code

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08739"
  id="VUID-VkShaderCreateInfoEXT-pCode-08739"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08739  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pCode` **must** not
  declare any capability that is not supported by the API, as described
  by the [Capabilities](#spirvenv-module-validation) section of the
  [SPIR-V Environment](#spirvenv-capabilities) appendix

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08740"
  id="VUID-VkShaderCreateInfoEXT-pCode-08740"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08740  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `pCode` declares
  any of the capabilities listed in the [SPIR-V
  Environment](#spirvenv-capabilities-table) appendix, one of the
  corresponding requirements **must** be satisfied

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08741"
  id="VUID-VkShaderCreateInfoEXT-pCode-08741"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08741  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pCode` **must** not
  declare any SPIR-V extension that is not supported by the API, as
  described by the [Extension](#spirvenv-extensions) section of the
  [SPIR-V Environment](#spirvenv-capabilities) appendix

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08742"
  id="VUID-VkShaderCreateInfoEXT-pCode-08742"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08742  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `pCode` declares
  any of the SPIR-V extensions listed in the [SPIR-V
  Environment](#spirvenv-extensions-table) appendix, one of the
  corresponding requirements **must** be satisfied

- <a href="#VUID-VkShaderCreateInfoEXT-flags-08412"
  id="VUID-VkShaderCreateInfoEXT-flags-08412"></a>
  VUID-VkShaderCreateInfoEXT-flags-08412  
  If `stage` is not `VK_SHADER_STAGE_TASK_BIT_EXT`,
  `VK_SHADER_STAGE_MESH_BIT_EXT`, `VK_SHADER_STAGE_VERTEX_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`,
  `VK_SHADER_STAGE_GEOMETRY_BIT`, or `VK_SHADER_STAGE_FRAGMENT_BIT`,
  `flags` **must** not include `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`

- <a href="#VUID-VkShaderCreateInfoEXT-flags-08486"
  id="VUID-VkShaderCreateInfoEXT-flags-08486"></a>
  VUID-VkShaderCreateInfoEXT-flags-08486  
  If `stage` is not `VK_SHADER_STAGE_FRAGMENT_BIT`, `flags` **must** not
  include `VK_SHADER_CREATE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_EXT`

- <a href="#VUID-VkShaderCreateInfoEXT-flags-08487"
  id="VUID-VkShaderCreateInfoEXT-flags-08487"></a>
  VUID-VkShaderCreateInfoEXT-flags-08487  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> feature
  is not enabled, `flags` **must** not include
  `VK_SHADER_CREATE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_EXT`

- <a href="#VUID-VkShaderCreateInfoEXT-flags-08488"
  id="VUID-VkShaderCreateInfoEXT-flags-08488"></a>
  VUID-VkShaderCreateInfoEXT-flags-08488  
  If `stage` is not `VK_SHADER_STAGE_FRAGMENT_BIT`, `flags` **must** not
  include `VK_SHADER_CREATE_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`

- <a href="#VUID-VkShaderCreateInfoEXT-flags-08489"
  id="VUID-VkShaderCreateInfoEXT-flags-08489"></a>
  VUID-VkShaderCreateInfoEXT-flags-08489  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentDensityMap"
  target="_blank" rel="noopener"><code>fragmentDensityMap</code></a>
  feature is not enabled, `flags` **must** not include
  `VK_SHADER_CREATE_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`

- <a href="#VUID-VkShaderCreateInfoEXT-flags-09404"
  id="VUID-VkShaderCreateInfoEXT-flags-09404"></a>
  VUID-VkShaderCreateInfoEXT-flags-09404  
  If `flags` includes
  `VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-subgroupSizeControl"
  target="_blank" rel="noopener"><code>subgroupSizeControl</code></a>
  feature **must** be enabled

- <a href="#VUID-VkShaderCreateInfoEXT-flags-09405"
  id="VUID-VkShaderCreateInfoEXT-flags-09405"></a>
  VUID-VkShaderCreateInfoEXT-flags-09405  
  If `flags` includes `VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT`,
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-computeFullSubgroups"
  target="_blank" rel="noopener"><code>computeFullSubgroups</code></a>
  feature **must** be enabled

- <a href="#VUID-VkShaderCreateInfoEXT-flags-08992"
  id="VUID-VkShaderCreateInfoEXT-flags-08992"></a>
  VUID-VkShaderCreateInfoEXT-flags-08992  
  If `flags` includes `VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT`,
  `stage` **must** be one of `VK_SHADER_STAGE_MESH_BIT_EXT`,
  `VK_SHADER_STAGE_TASK_BIT_EXT`, or `VK_SHADER_STAGE_COMPUTE_BIT`

- <a href="#VUID-VkShaderCreateInfoEXT-flags-08485"
  id="VUID-VkShaderCreateInfoEXT-flags-08485"></a>
  VUID-VkShaderCreateInfoEXT-flags-08485  
  If `stage` is not `VK_SHADER_STAGE_COMPUTE_BIT`, `flags` **must** not
  include `VK_SHADER_CREATE_DISPATCH_BASE_BIT_EXT`

- <a href="#VUID-VkShaderCreateInfoEXT-flags-08414"
  id="VUID-VkShaderCreateInfoEXT-flags-08414"></a>
  VUID-VkShaderCreateInfoEXT-flags-08414  
  If `stage` is not `VK_SHADER_STAGE_MESH_BIT_EXT`, `flags` **must** not
  include `VK_SHADER_CREATE_NO_TASK_SHADER_BIT_EXT`

- <a href="#VUID-VkShaderCreateInfoEXT-flags-08416"
  id="VUID-VkShaderCreateInfoEXT-flags-08416"></a>
  VUID-VkShaderCreateInfoEXT-flags-08416  
  If `flags` includes both
  `VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT` and
  `VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT`, the local workgroup
  size in the X dimension of the shader **must** be a multiple of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxSubgroupSize"
  target="_blank" rel="noopener"><code>maxSubgroupSize</code></a>

- <a href="#VUID-VkShaderCreateInfoEXT-flags-08417"
  id="VUID-VkShaderCreateInfoEXT-flags-08417"></a>
  VUID-VkShaderCreateInfoEXT-flags-08417  
  If `flags` includes `VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT`
  but not `VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT` and no
  [VkShaderRequiredSubgroupSizeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderRequiredSubgroupSizeCreateInfoEXT.html)
  structure is included in the `pNext` chain, the local workgroup size
  in the X dimension of the shader **must** be a multiple of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-subgroup-size"
  target="_blank" rel="noopener"><code>subgroupSize</code></a>

- <a href="#VUID-VkShaderCreateInfoEXT-stage-08418"
  id="VUID-VkShaderCreateInfoEXT-stage-08418"></a>
  VUID-VkShaderCreateInfoEXT-stage-08418  
  `stage` **must** not be `VK_SHADER_STAGE_ALL_GRAPHICS` or
  `VK_SHADER_STAGE_ALL`

- <a href="#VUID-VkShaderCreateInfoEXT-stage-08419"
  id="VUID-VkShaderCreateInfoEXT-stage-08419"></a>
  VUID-VkShaderCreateInfoEXT-stage-08419  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-tessellationShader"
  target="_blank" rel="noopener"><code>tessellationShader</code></a>
  feature is not enabled, `stage` **must** not be
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` or
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`

- <a href="#VUID-VkShaderCreateInfoEXT-stage-08420"
  id="VUID-VkShaderCreateInfoEXT-stage-08420"></a>
  VUID-VkShaderCreateInfoEXT-stage-08420  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-geometryShader"
  target="_blank" rel="noopener"><code>geometryShader</code></a> feature
  is not enabled, `stage` **must** not be `VK_SHADER_STAGE_GEOMETRY_BIT`

- <a href="#VUID-VkShaderCreateInfoEXT-stage-08421"
  id="VUID-VkShaderCreateInfoEXT-stage-08421"></a>
  VUID-VkShaderCreateInfoEXT-stage-08421  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-taskShader"
  target="_blank" rel="noopener"><code>taskShader</code></a> feature is
  not enabled, `stage` **must** not be `VK_SHADER_STAGE_TASK_BIT_EXT`

- <a href="#VUID-VkShaderCreateInfoEXT-stage-08422"
  id="VUID-VkShaderCreateInfoEXT-stage-08422"></a>
  VUID-VkShaderCreateInfoEXT-stage-08422  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-meshShader"
  target="_blank" rel="noopener"><code>meshShader</code></a> feature is
  not enabled, `stage` **must** not be `VK_SHADER_STAGE_MESH_BIT_EXT`

- <a href="#VUID-VkShaderCreateInfoEXT-stage-08425"
  id="VUID-VkShaderCreateInfoEXT-stage-08425"></a>
  VUID-VkShaderCreateInfoEXT-stage-08425  
  `stage` **must** not be `VK_SHADER_STAGE_SUBPASS_SHADING_BIT_HUAWEI`

- <a href="#VUID-VkShaderCreateInfoEXT-stage-08426"
  id="VUID-VkShaderCreateInfoEXT-stage-08426"></a>
  VUID-VkShaderCreateInfoEXT-stage-08426  
  `stage` **must** not be `VK_SHADER_STAGE_CLUSTER_CULLING_BIT_HUAWEI`

- <a href="#VUID-VkShaderCreateInfoEXT-nextStage-08427"
  id="VUID-VkShaderCreateInfoEXT-nextStage-08427"></a>
  VUID-VkShaderCreateInfoEXT-nextStage-08427  
  If `stage` is `VK_SHADER_STAGE_VERTEX_BIT`, `nextStage` **must** not
  include any bits other than
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`,
  `VK_SHADER_STAGE_GEOMETRY_BIT`, and `VK_SHADER_STAGE_FRAGMENT_BIT`

- <a href="#VUID-VkShaderCreateInfoEXT-nextStage-08428"
  id="VUID-VkShaderCreateInfoEXT-nextStage-08428"></a>
  VUID-VkShaderCreateInfoEXT-nextStage-08428  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-tessellationShader"
  target="_blank" rel="noopener"><code>tessellationShader</code></a>
  feature is not enabled, `nextStage` **must** not include
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` or
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`

- <a href="#VUID-VkShaderCreateInfoEXT-nextStage-08429"
  id="VUID-VkShaderCreateInfoEXT-nextStage-08429"></a>
  VUID-VkShaderCreateInfoEXT-nextStage-08429  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-geometryShader"
  target="_blank" rel="noopener"><code>geometryShader</code></a> feature
  is not enabled, `nextStage` **must** not include
  `VK_SHADER_STAGE_GEOMETRY_BIT`

- <a href="#VUID-VkShaderCreateInfoEXT-nextStage-08430"
  id="VUID-VkShaderCreateInfoEXT-nextStage-08430"></a>
  VUID-VkShaderCreateInfoEXT-nextStage-08430  
  If `stage` is `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`, `nextStage`
  **must** not include any bits other than
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`

- <a href="#VUID-VkShaderCreateInfoEXT-nextStage-08431"
  id="VUID-VkShaderCreateInfoEXT-nextStage-08431"></a>
  VUID-VkShaderCreateInfoEXT-nextStage-08431  
  If `stage` is `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`,
  `nextStage` **must** not include any bits other than
  `VK_SHADER_STAGE_GEOMETRY_BIT` and `VK_SHADER_STAGE_FRAGMENT_BIT`

- <a href="#VUID-VkShaderCreateInfoEXT-nextStage-08433"
  id="VUID-VkShaderCreateInfoEXT-nextStage-08433"></a>
  VUID-VkShaderCreateInfoEXT-nextStage-08433  
  If `stage` is `VK_SHADER_STAGE_GEOMETRY_BIT`, `nextStage` **must** not
  include any bits other than `VK_SHADER_STAGE_FRAGMENT_BIT`

- <a href="#VUID-VkShaderCreateInfoEXT-nextStage-08434"
  id="VUID-VkShaderCreateInfoEXT-nextStage-08434"></a>
  VUID-VkShaderCreateInfoEXT-nextStage-08434  
  If `stage` is `VK_SHADER_STAGE_FRAGMENT_BIT` or
  `VK_SHADER_STAGE_COMPUTE_BIT`, `nextStage` **must** be 0

- <a href="#VUID-VkShaderCreateInfoEXT-nextStage-08435"
  id="VUID-VkShaderCreateInfoEXT-nextStage-08435"></a>
  VUID-VkShaderCreateInfoEXT-nextStage-08435  
  If `stage` is `VK_SHADER_STAGE_TASK_BIT_EXT`, `nextStage` **must** not
  include any bits other than `VK_SHADER_STAGE_MESH_BIT_EXT`

- <a href="#VUID-VkShaderCreateInfoEXT-nextStage-08436"
  id="VUID-VkShaderCreateInfoEXT-nextStage-08436"></a>
  VUID-VkShaderCreateInfoEXT-nextStage-08436  
  If `stage` is `VK_SHADER_STAGE_MESH_BIT_EXT`, `nextStage` **must** not
  include any bits other than `VK_SHADER_STAGE_FRAGMENT_BIT`

- <a href="#VUID-VkShaderCreateInfoEXT-pName-08440"
  id="VUID-VkShaderCreateInfoEXT-pName-08440"></a>
  VUID-VkShaderCreateInfoEXT-pName-08440  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pName` **must** be
  the name of an `OpEntryPoint` in `pCode` with an execution model that
  matches `stage`

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08492"
  id="VUID-VkShaderCreateInfoEXT-pCode-08492"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08492  
  If `codeType` is `VK_SHADER_CODE_TYPE_BINARY_EXT`, `pCode` **must** be
  aligned to `16` bytes

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08493"
  id="VUID-VkShaderCreateInfoEXT-pCode-08493"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08493  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pCode` **must** be
  aligned to `4` bytes

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08448"
  id="VUID-VkShaderCreateInfoEXT-pCode-08448"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08448  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and the identified
  entry point includes any variable in its interface that is declared
  with the `ClipDistance` `BuiltIn` decoration, that variable **must**
  not have an array size greater than
  `VkPhysicalDeviceLimits`::`maxClipDistances`

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08449"
  id="VUID-VkShaderCreateInfoEXT-pCode-08449"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08449  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and the identified
  entry point includes any variable in its interface that is declared
  with the `CullDistance` `BuiltIn` decoration, that variable **must**
  not have an array size greater than
  `VkPhysicalDeviceLimits`::`maxCullDistances`

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08450"
  id="VUID-VkShaderCreateInfoEXT-pCode-08450"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08450  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and the identified
  entry point includes variables in its interface that are declared with
  the `ClipDistance` `BuiltIn` decoration and variables in its interface
  that are declared with the `CullDistance` `BuiltIn` decoration, those
  variables **must** not have array sizes which sum to more than
  `VkPhysicalDeviceLimits`::`maxCombinedClipAndCullDistances`

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08451"
  id="VUID-VkShaderCreateInfoEXT-pCode-08451"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08451  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and the identified
  entry point includes any variable in its interface that is declared
  with the `SampleMask` `BuiltIn` decoration, that variable **must** not
  have an array size greater than
  `VkPhysicalDeviceLimits`::`maxSampleMaskWords`

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08452"
  id="VUID-VkShaderCreateInfoEXT-pCode-08452"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08452  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is
  `VK_SHADER_STAGE_VERTEX_BIT`, the identified entry point **must** not
  include any input variable in its interface that is decorated with
  `CullDistance`

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08453"
  id="VUID-VkShaderCreateInfoEXT-pCode-08453"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08453  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` or
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, and the identified
  entry point has an `OpExecutionMode` instruction specifying a patch
  size with `OutputVertices`, the patch size **must** be greater than
  `0` and less than or equal to
  `VkPhysicalDeviceLimits`::`maxTessellationPatchSize`

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08454"
  id="VUID-VkShaderCreateInfoEXT-pCode-08454"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08454  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is
  `VK_SHADER_STAGE_GEOMETRY_BIT`, the identified entry point **must**
  have an `OpExecutionMode` instruction specifying a maximum output
  vertex count that is greater than `0` and less than or equal to
  `VkPhysicalDeviceLimits`::`maxGeometryOutputVertices`

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08455"
  id="VUID-VkShaderCreateInfoEXT-pCode-08455"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08455  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is
  `VK_SHADER_STAGE_GEOMETRY_BIT`, the identified entry point **must**
  have an `OpExecutionMode` instruction specifying an invocation count
  that is greater than `0` and less than or equal to
  `VkPhysicalDeviceLimits`::`maxGeometryShaderInvocations`

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08456"
  id="VUID-VkShaderCreateInfoEXT-pCode-08456"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08456  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader stage</a>, and
  the identified entry point writes to `Layer` for any primitive, it
  **must** write the same value to `Layer` for all vertices of a given
  primitive

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08457"
  id="VUID-VkShaderCreateInfoEXT-pCode-08457"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08457  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader stage</a>, and
  the identified entry point writes to `ViewportIndex` for any
  primitive, it **must** write the same value to `ViewportIndex` for all
  vertices of a given primitive

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08458"
  id="VUID-VkShaderCreateInfoEXT-pCode-08458"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08458  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is
  `VK_SHADER_STAGE_FRAGMENT_BIT`, the identified entry point **must**
  not include any output variables in its interface decorated with
  `CullDistance`

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08459"
  id="VUID-VkShaderCreateInfoEXT-pCode-08459"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08459  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is
  `VK_SHADER_STAGE_FRAGMENT_BIT`, and the identified entry point writes
  to `FragDepth` in any execution path, all execution paths that are not
  exclusive to helper invocations **must** either discard the fragment,
  or write or initialize the value of `FragDepth`

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-08460"
  id="VUID-VkShaderCreateInfoEXT-pCode-08460"></a>
  VUID-VkShaderCreateInfoEXT-pCode-08460  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, the shader code in
  `pCode` **must** be valid as described by the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirv-spec"
  target="_blank" rel="noopener">Khronos SPIR-V Specification</a> after
  applying the specializations provided in `pSpecializationInfo`, if
  any, and then converting all specialization constants into fixed
  constants

- <a href="#VUID-VkShaderCreateInfoEXT-codeType-08872"
  id="VUID-VkShaderCreateInfoEXT-codeType-08872"></a>
  VUID-VkShaderCreateInfoEXT-codeType-08872  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, `pCode` **must**
  contain an `OpExecutionMode` instruction specifying the type of
  subdivision

- <a href="#VUID-VkShaderCreateInfoEXT-codeType-08873"
  id="VUID-VkShaderCreateInfoEXT-codeType-08873"></a>
  VUID-VkShaderCreateInfoEXT-codeType-08873  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, `pCode` **must**
  contain an `OpExecutionMode` instruction specifying the orientation of
  triangles generated by the tessellator

- <a href="#VUID-VkShaderCreateInfoEXT-codeType-08874"
  id="VUID-VkShaderCreateInfoEXT-codeType-08874"></a>
  VUID-VkShaderCreateInfoEXT-codeType-08874  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, `pCode` **must**
  contain an `OpExecutionMode` instruction specifying the spacing of
  segments on the edges of tessellated primitives

- <a href="#VUID-VkShaderCreateInfoEXT-codeType-08875"
  id="VUID-VkShaderCreateInfoEXT-codeType-08875"></a>
  VUID-VkShaderCreateInfoEXT-codeType-08875  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, `pCode` **must**
  contain an `OpExecutionMode` instruction specifying the output patch
  size

Valid Usage (Implicit)

- <a href="#VUID-VkShaderCreateInfoEXT-sType-sType"
  id="VUID-VkShaderCreateInfoEXT-sType-sType"></a>
  VUID-VkShaderCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SHADER_CREATE_INFO_EXT`

- <a href="#VUID-VkShaderCreateInfoEXT-pNext-pNext"
  id="VUID-VkShaderCreateInfoEXT-pNext-pNext"></a>
  VUID-VkShaderCreateInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)

- <a href="#VUID-VkShaderCreateInfoEXT-sType-unique"
  id="VUID-VkShaderCreateInfoEXT-sType-unique"></a>
  VUID-VkShaderCreateInfoEXT-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkShaderCreateInfoEXT-flags-parameter"
  id="VUID-VkShaderCreateInfoEXT-flags-parameter"></a>
  VUID-VkShaderCreateInfoEXT-flags-parameter  
  `flags` **must** be a valid combination of
  [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateFlagBitsEXT.html) values

- <a href="#VUID-VkShaderCreateInfoEXT-stage-parameter"
  id="VUID-VkShaderCreateInfoEXT-stage-parameter"></a>
  VUID-VkShaderCreateInfoEXT-stage-parameter  
  `stage` **must** be a valid
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) value

- <a href="#VUID-VkShaderCreateInfoEXT-nextStage-parameter"
  id="VUID-VkShaderCreateInfoEXT-nextStage-parameter"></a>
  VUID-VkShaderCreateInfoEXT-nextStage-parameter  
  `nextStage` **must** be a valid combination of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) values

- <a href="#VUID-VkShaderCreateInfoEXT-codeType-parameter"
  id="VUID-VkShaderCreateInfoEXT-codeType-parameter"></a>
  VUID-VkShaderCreateInfoEXT-codeType-parameter  
  `codeType` **must** be a valid
  [VkShaderCodeTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCodeTypeEXT.html) value

- <a href="#VUID-VkShaderCreateInfoEXT-pCode-parameter"
  id="VUID-VkShaderCreateInfoEXT-pCode-parameter"></a>
  VUID-VkShaderCreateInfoEXT-pCode-parameter  
  `pCode` **must** be a valid pointer to an array of `codeSize` bytes

- <a href="#VUID-VkShaderCreateInfoEXT-pName-parameter"
  id="VUID-VkShaderCreateInfoEXT-pName-parameter"></a>
  VUID-VkShaderCreateInfoEXT-pName-parameter  
  If `pName` is not `NULL`, `pName` **must** be a null-terminated UTF-8
  string

- <a href="#VUID-VkShaderCreateInfoEXT-pSetLayouts-parameter"
  id="VUID-VkShaderCreateInfoEXT-pSetLayouts-parameter"></a>
  VUID-VkShaderCreateInfoEXT-pSetLayouts-parameter  
  If `setLayoutCount` is not `0`, and `pSetLayouts` is not `NULL`,
  `pSetLayouts` **must** be a valid pointer to an array of
  `setLayoutCount` valid
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) handles

- <a href="#VUID-VkShaderCreateInfoEXT-pPushConstantRanges-parameter"
  id="VUID-VkShaderCreateInfoEXT-pPushConstantRanges-parameter"></a>
  VUID-VkShaderCreateInfoEXT-pPushConstantRanges-parameter  
  If `pushConstantRangeCount` is not `0`, and `pPushConstantRanges` is
  not `NULL`, `pPushConstantRanges` **must** be a valid pointer to an
  array of `pushConstantRangeCount` valid
  [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantRange.html) structures

- <a href="#VUID-VkShaderCreateInfoEXT-pSpecializationInfo-parameter"
  id="VUID-VkShaderCreateInfoEXT-pSpecializationInfo-parameter"></a>
  VUID-VkShaderCreateInfoEXT-pSpecializationInfo-parameter  
  If `pSpecializationInfo` is not `NULL`, `pSpecializationInfo` **must**
  be a valid pointer to a valid
  [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSpecializationInfo.html) structure

- <a href="#VUID-VkShaderCreateInfoEXT-codeSize-arraylength"
  id="VUID-VkShaderCreateInfoEXT-codeSize-arraylength"></a>
  VUID-VkShaderCreateInfoEXT-codeSize-arraylength  
  `codeSize` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html),
[VkPushConstantRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantRange.html),
[VkShaderCodeTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCodeTypeEXT.html),
[VkShaderCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateFlagsEXT.html),
[VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html),
[VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSpecializationInfo.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateShadersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateShadersEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkShaderCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
