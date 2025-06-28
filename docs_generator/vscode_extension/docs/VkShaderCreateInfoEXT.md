# VkShaderCreateInfoEXT(3) Manual Page

## Name

VkShaderCreateInfoEXT - Structure specifying parameters of a newly created shader



## [](#_c_specification)C Specification

The `VkShaderCreateInfoEXT` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateFlagBitsEXT.html) describing additional parameters of the shader.
- `stage` is a [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) value specifying a single shader stage.
- `nextStage` is a bitmask of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) specifying which stages **may** be used as a logically next bound stage when drawing with the shader bound. A value of zero indicates this shader stage **must** be the last one.
- `codeType` is a [VkShaderCodeTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCodeTypeEXT.html) value specifying the type of the shader code pointed to be `pCode`.
- `codeSize` is the size in bytes of the shader code pointed to be `pCode`.
- `pCode` is a pointer to the shader code to use to create the shader.
- `pName` is a pointer to a null-terminated UTF-8 string specifying the entry point name of the shader for this stage.
- `setLayoutCount` is the number of descriptor set layouts pointed to by `pSetLayouts`.
- `pSetLayouts` is a pointer to an array of [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) objects used by the shader stage. The implementation **must** not access these objects outside of the duration of the command this structure is passed to.
- `pushConstantRangeCount` is the number of push constant ranges pointed to by `pPushConstantRanges`.
- `pPushConstantRanges` is a pointer to an array of [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantRange.html) structures used by the shader stage.
- `pSpecializationInfo` is a pointer to a [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSpecializationInfo.html) structure, as described in [Specialization Constants](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-specialization-constants), or `NULL`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkShaderCreateInfoEXT-codeSize-08735)VUID-VkShaderCreateInfoEXT-codeSize-08735  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `codeSize` **must** be a multiple of 4
- [](#VUID-VkShaderCreateInfoEXT-pCode-08736)VUID-VkShaderCreateInfoEXT-pCode-08736  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pCode` **must** point to valid SPIR-V code, formatted and packed as described by the [Khronos SPIR-V Specification](#spirv-spec)
- [](#VUID-VkShaderCreateInfoEXT-pCode-08737)VUID-VkShaderCreateInfoEXT-pCode-08737  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pCode` **must** adhere to the validation rules described by the [Validation Rules within a Module](#spirvenv-module-validation) section of the [SPIR-V Environment](#spirvenv-capabilities) appendix
- [](#VUID-VkShaderCreateInfoEXT-pCode-08738)VUID-VkShaderCreateInfoEXT-pCode-08738  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pCode` **must** declare the `Shader` capability for SPIR-V code
- [](#VUID-VkShaderCreateInfoEXT-pCode-08739)VUID-VkShaderCreateInfoEXT-pCode-08739  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pCode` **must** not declare any capability that is not supported by the API, as described by the [Capabilities](#spirvenv-module-validation) section of the [SPIR-V Environment](#spirvenv-capabilities) appendix
- [](#VUID-VkShaderCreateInfoEXT-pCode-08740)VUID-VkShaderCreateInfoEXT-pCode-08740  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `pCode` declares any of the capabilities listed in the [SPIR-V Environment](#spirvenv-capabilities-table) appendix, one of the corresponding requirements **must** be satisfied
- [](#VUID-VkShaderCreateInfoEXT-pCode-08741)VUID-VkShaderCreateInfoEXT-pCode-08741  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pCode` **must** not declare any SPIR-V extension that is not supported by the API, as described by the [Extension](#spirvenv-extensions) section of the [SPIR-V Environment](#spirvenv-capabilities) appendix
- [](#VUID-VkShaderCreateInfoEXT-pCode-08742)VUID-VkShaderCreateInfoEXT-pCode-08742  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `pCode` declares any of the SPIR-V extensions listed in the [SPIR-V Environment](#spirvenv-extensions-table) appendix, one of the corresponding requirements **must** be satisfied
- [](#VUID-VkShaderCreateInfoEXT-flags-08412)VUID-VkShaderCreateInfoEXT-flags-08412  
  If `stage` is not `VK_SHADER_STAGE_TASK_BIT_EXT`, `VK_SHADER_STAGE_MESH_BIT_EXT`, `VK_SHADER_STAGE_VERTEX_BIT`, `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`, `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, `VK_SHADER_STAGE_GEOMETRY_BIT`, or `VK_SHADER_STAGE_FRAGMENT_BIT`, `flags` **must** not include `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`
- [](#VUID-VkShaderCreateInfoEXT-flags-08486)VUID-VkShaderCreateInfoEXT-flags-08486  
  If `stage` is not `VK_SHADER_STAGE_FRAGMENT_BIT`, `flags` **must** not include `VK_SHADER_CREATE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_EXT`
- [](#VUID-VkShaderCreateInfoEXT-flags-08487)VUID-VkShaderCreateInfoEXT-flags-08487  
  If the [`attachmentFragmentShadingRate`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-attachmentFragmentShadingRate) feature is not enabled, `flags` **must** not include `VK_SHADER_CREATE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_EXT`
- [](#VUID-VkShaderCreateInfoEXT-flags-08488)VUID-VkShaderCreateInfoEXT-flags-08488  
  If `stage` is not `VK_SHADER_STAGE_FRAGMENT_BIT`, `flags` **must** not include `VK_SHADER_CREATE_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`
- [](#VUID-VkShaderCreateInfoEXT-flags-08489)VUID-VkShaderCreateInfoEXT-flags-08489  
  If the [`fragmentDensityMap`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-fragmentDensityMap) feature is not enabled, `flags` **must** not include `VK_SHADER_CREATE_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`
- [](#VUID-VkShaderCreateInfoEXT-flags-09404)VUID-VkShaderCreateInfoEXT-flags-09404  
  If `flags` includes `VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT`, the [`subgroupSizeControl`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-subgroupSizeControl) feature **must** be enabled
- [](#VUID-VkShaderCreateInfoEXT-flags-09405)VUID-VkShaderCreateInfoEXT-flags-09405  
  If `flags` includes `VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT`, the [`computeFullSubgroups`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-computeFullSubgroups) feature **must** be enabled
- [](#VUID-VkShaderCreateInfoEXT-flags-11005)VUID-VkShaderCreateInfoEXT-flags-11005  
  If `flags` includes `VK_SHADER_CREATE_INDIRECT_BINDABLE_BIT_EXT`, then the [`VkPhysicalDeviceDeviceGeneratedCommandsFeaturesEXT`::`deviceGeneratedCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedCommands) feature **must** be enabled
- [](#VUID-VkShaderCreateInfoEXT-flags-11006)VUID-VkShaderCreateInfoEXT-flags-11006  
  If `flags` includes `VK_SHADER_CREATE_INDIRECT_BINDABLE_BIT_EXT`, then the identified entry point **must** not specify `Xfb` execution mode
- [](#VUID-VkShaderCreateInfoEXT-flags-08992)VUID-VkShaderCreateInfoEXT-flags-08992  
  If `flags` includes `VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT`, `stage` **must** be one of `VK_SHADER_STAGE_MESH_BIT_EXT`, `VK_SHADER_STAGE_TASK_BIT_EXT`, or `VK_SHADER_STAGE_COMPUTE_BIT`
- [](#VUID-VkShaderCreateInfoEXT-flags-08485)VUID-VkShaderCreateInfoEXT-flags-08485  
  If `stage` is not `VK_SHADER_STAGE_COMPUTE_BIT`, `flags` **must** not include `VK_SHADER_CREATE_DISPATCH_BASE_BIT_EXT`
- [](#VUID-VkShaderCreateInfoEXT-flags-08414)VUID-VkShaderCreateInfoEXT-flags-08414  
  If `stage` is not `VK_SHADER_STAGE_MESH_BIT_EXT`, `flags` **must** not include `VK_SHADER_CREATE_NO_TASK_SHADER_BIT_EXT`
- [](#VUID-VkShaderCreateInfoEXT-flags-08416)VUID-VkShaderCreateInfoEXT-flags-08416  
  If `flags` includes both `VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT` and `VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT`, the local workgroup size in the X dimension of the shader **must** be a multiple of [`maxSubgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxSubgroupSize)
- [](#VUID-VkShaderCreateInfoEXT-flags-08417)VUID-VkShaderCreateInfoEXT-flags-08417  
  If `flags` includes `VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT` but not `VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT` and no [VkShaderRequiredSubgroupSizeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderRequiredSubgroupSizeCreateInfoEXT.html) structure is included in the `pNext` chain, the local workgroup size in the X dimension of the shader **must** be a multiple of [`subgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-subgroupSize)
- [](#VUID-VkShaderCreateInfoEXT-stage-08418)VUID-VkShaderCreateInfoEXT-stage-08418  
  `stage` **must** not be `VK_SHADER_STAGE_ALL_GRAPHICS` or `VK_SHADER_STAGE_ALL`
- [](#VUID-VkShaderCreateInfoEXT-stage-08419)VUID-VkShaderCreateInfoEXT-stage-08419  
  If the [`tessellationShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tessellationShader) feature is not enabled, `stage` **must** not be `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` or `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`
- [](#VUID-VkShaderCreateInfoEXT-stage-08420)VUID-VkShaderCreateInfoEXT-stage-08420  
  If the [`geometryShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-geometryShader) feature is not enabled, `stage` **must** not be `VK_SHADER_STAGE_GEOMETRY_BIT`
- [](#VUID-VkShaderCreateInfoEXT-stage-08421)VUID-VkShaderCreateInfoEXT-stage-08421  
  If the [`taskShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-taskShader) feature is not enabled, `stage` **must** not be `VK_SHADER_STAGE_TASK_BIT_EXT`
- [](#VUID-VkShaderCreateInfoEXT-stage-08422)VUID-VkShaderCreateInfoEXT-stage-08422  
  If the [`meshShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-meshShader) feature is not enabled, `stage` **must** not be `VK_SHADER_STAGE_MESH_BIT_EXT`
- [](#VUID-VkShaderCreateInfoEXT-stage-08425)VUID-VkShaderCreateInfoEXT-stage-08425  
  `stage` **must** not be `VK_SHADER_STAGE_SUBPASS_SHADING_BIT_HUAWEI`
- [](#VUID-VkShaderCreateInfoEXT-stage-08426)VUID-VkShaderCreateInfoEXT-stage-08426  
  `stage` **must** not be `VK_SHADER_STAGE_CLUSTER_CULLING_BIT_HUAWEI`
- [](#VUID-VkShaderCreateInfoEXT-nextStage-08427)VUID-VkShaderCreateInfoEXT-nextStage-08427  
  If `stage` is `VK_SHADER_STAGE_VERTEX_BIT`, `nextStage` **must** not include any bits other than `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`, `VK_SHADER_STAGE_GEOMETRY_BIT`, and `VK_SHADER_STAGE_FRAGMENT_BIT`
- [](#VUID-VkShaderCreateInfoEXT-nextStage-08428)VUID-VkShaderCreateInfoEXT-nextStage-08428  
  If the [`tessellationShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tessellationShader) feature is not enabled, `nextStage` **must** not include `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` or `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`
- [](#VUID-VkShaderCreateInfoEXT-nextStage-08429)VUID-VkShaderCreateInfoEXT-nextStage-08429  
  If the [`geometryShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-geometryShader) feature is not enabled, `nextStage` **must** not include `VK_SHADER_STAGE_GEOMETRY_BIT`
- [](#VUID-VkShaderCreateInfoEXT-nextStage-08430)VUID-VkShaderCreateInfoEXT-nextStage-08430  
  If `stage` is `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`, `nextStage` **must** not include any bits other than `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`
- [](#VUID-VkShaderCreateInfoEXT-nextStage-08431)VUID-VkShaderCreateInfoEXT-nextStage-08431  
  If `stage` is `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, `nextStage` **must** not include any bits other than `VK_SHADER_STAGE_GEOMETRY_BIT` and `VK_SHADER_STAGE_FRAGMENT_BIT`
- [](#VUID-VkShaderCreateInfoEXT-nextStage-08433)VUID-VkShaderCreateInfoEXT-nextStage-08433  
  If `stage` is `VK_SHADER_STAGE_GEOMETRY_BIT`, `nextStage` **must** not include any bits other than `VK_SHADER_STAGE_FRAGMENT_BIT`
- [](#VUID-VkShaderCreateInfoEXT-nextStage-08434)VUID-VkShaderCreateInfoEXT-nextStage-08434  
  If `stage` is `VK_SHADER_STAGE_FRAGMENT_BIT` or `VK_SHADER_STAGE_COMPUTE_BIT`, `nextStage` **must** be 0
- [](#VUID-VkShaderCreateInfoEXT-nextStage-08435)VUID-VkShaderCreateInfoEXT-nextStage-08435  
  If `stage` is `VK_SHADER_STAGE_TASK_BIT_EXT`, `nextStage` **must** not include any bits other than `VK_SHADER_STAGE_MESH_BIT_EXT`
- [](#VUID-VkShaderCreateInfoEXT-nextStage-08436)VUID-VkShaderCreateInfoEXT-nextStage-08436  
  If `stage` is `VK_SHADER_STAGE_MESH_BIT_EXT`, `nextStage` **must** not include any bits other than `VK_SHADER_STAGE_FRAGMENT_BIT`
- [](#VUID-VkShaderCreateInfoEXT-pName-08440)VUID-VkShaderCreateInfoEXT-pName-08440  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pName` **must** be the name of an `OpEntryPoint` in `pCode` with an execution model that matches `stage`
- [](#VUID-VkShaderCreateInfoEXT-pCode-08492)VUID-VkShaderCreateInfoEXT-pCode-08492  
  If `codeType` is `VK_SHADER_CODE_TYPE_BINARY_EXT`, `pCode` **must** be aligned to `16` bytes
- [](#VUID-VkShaderCreateInfoEXT-pCode-08493)VUID-VkShaderCreateInfoEXT-pCode-08493  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, `pCode` **must** be aligned to `4` bytes
- [](#VUID-VkShaderCreateInfoEXT-pCode-08448)VUID-VkShaderCreateInfoEXT-pCode-08448  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and the identified entry point includes any variable in its interface that is declared with the `ClipDistance` `BuiltIn` decoration, that variable **must** not have an array size greater than `VkPhysicalDeviceLimits`::`maxClipDistances`
- [](#VUID-VkShaderCreateInfoEXT-pCode-08449)VUID-VkShaderCreateInfoEXT-pCode-08449  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and the identified entry point includes any variable in its interface that is declared with the `CullDistance` `BuiltIn` decoration, that variable **must** not have an array size greater than `VkPhysicalDeviceLimits`::`maxCullDistances`
- [](#VUID-VkShaderCreateInfoEXT-pCode-08450)VUID-VkShaderCreateInfoEXT-pCode-08450  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and the identified entry point includes variables in its interface that are declared with the `ClipDistance` `BuiltIn` decoration and variables in its interface that are declared with the `CullDistance` `BuiltIn` decoration, those variables **must** not have array sizes which sum to more than `VkPhysicalDeviceLimits`::`maxCombinedClipAndCullDistances`
- [](#VUID-VkShaderCreateInfoEXT-pCode-08451)VUID-VkShaderCreateInfoEXT-pCode-08451  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and the identified entry point includes any variable in its interface that is declared with the `SampleMask` `BuiltIn` decoration, that variable **must** not have an array size greater than `VkPhysicalDeviceLimits`::`maxSampleMaskWords`
- [](#VUID-VkShaderCreateInfoEXT-pCode-08453)VUID-VkShaderCreateInfoEXT-pCode-08453  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` or `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, and the identified entry point has an `OpExecutionMode` instruction specifying a patch size with `OutputVertices`, the patch size **must** be greater than `0` and less than or equal to `VkPhysicalDeviceLimits`::`maxTessellationPatchSize`
- [](#VUID-VkShaderCreateInfoEXT-pCode-08454)VUID-VkShaderCreateInfoEXT-pCode-08454  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is `VK_SHADER_STAGE_GEOMETRY_BIT`, the identified entry point **must** have an `OpExecutionMode` instruction specifying a maximum output vertex count that is greater than `0` and less than or equal to `VkPhysicalDeviceLimits`::`maxGeometryOutputVertices`
- [](#VUID-VkShaderCreateInfoEXT-pCode-08455)VUID-VkShaderCreateInfoEXT-pCode-08455  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is `VK_SHADER_STAGE_GEOMETRY_BIT`, the identified entry point **must** have an `OpExecutionMode` instruction specifying an invocation count that is greater than `0` and less than or equal to `VkPhysicalDeviceLimits`::`maxGeometryShaderInvocations`
- [](#VUID-VkShaderCreateInfoEXT-pCode-08456)VUID-VkShaderCreateInfoEXT-pCode-08456  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is a [pre-rasterization shader stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization), and the identified entry point writes to `Layer` for any primitive, it **must** write the same value to `Layer` for all vertices of a given primitive
- [](#VUID-VkShaderCreateInfoEXT-pCode-08457)VUID-VkShaderCreateInfoEXT-pCode-08457  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is a [pre-rasterization shader stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization), and the identified entry point writes to `ViewportIndex` for any primitive, it **must** write the same value to `ViewportIndex` for all vertices of a given primitive
- [](#VUID-VkShaderCreateInfoEXT-pCode-08459)VUID-VkShaderCreateInfoEXT-pCode-08459  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is `VK_SHADER_STAGE_FRAGMENT_BIT`, and the identified entry point writes to `FragDepth` in any execution path, all execution paths that are not exclusive to helper invocations **must** either discard the fragment, or write or initialize the value of `FragDepth`
- [](#VUID-VkShaderCreateInfoEXT-pCode-08460)VUID-VkShaderCreateInfoEXT-pCode-08460  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, the shader code in `pCode` **must** be valid as described by the [Khronos SPIR-V Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirv-spec) after applying the specializations provided in `pSpecializationInfo`, if any, and then converting all specialization constants into fixed constants
- [](#VUID-VkShaderCreateInfoEXT-codeType-08872)VUID-VkShaderCreateInfoEXT-codeType-08872  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, `pCode` **must** contain an `OpExecutionMode` instruction specifying the type of subdivision
- [](#VUID-VkShaderCreateInfoEXT-codeType-08873)VUID-VkShaderCreateInfoEXT-codeType-08873  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, `pCode` **must** contain an `OpExecutionMode` instruction specifying the orientation of triangles generated by the tessellator
- [](#VUID-VkShaderCreateInfoEXT-codeType-08874)VUID-VkShaderCreateInfoEXT-codeType-08874  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, `pCode` **must** contain an `OpExecutionMode` instruction specifying the spacing of segments on the edges of tessellated primitives
- [](#VUID-VkShaderCreateInfoEXT-codeType-08875)VUID-VkShaderCreateInfoEXT-codeType-08875  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and `stage` is `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, `pCode` **must** contain an `OpExecutionMode` instruction specifying the output patch size
- [](#VUID-VkShaderCreateInfoEXT-pPushConstantRanges-10063)VUID-VkShaderCreateInfoEXT-pPushConstantRanges-10063  
  Any two elements of `pPushConstantRanges` **must** not include the same stage in `stageFlags`
- [](#VUID-VkShaderCreateInfoEXT-codeType-10064)VUID-VkShaderCreateInfoEXT-codeType-10064  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and if a push constant block is declared in a shader, then an element of `pPushConstantRanges`::`stageFlags` **must** match `stage`
- [](#VUID-VkShaderCreateInfoEXT-codeType-10065)VUID-VkShaderCreateInfoEXT-codeType-10065  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and if a push constant block is declared in a shader, the block must be contained inside the element of `pPushConstantRanges` that matches the stage
- [](#VUID-VkShaderCreateInfoEXT-codeType-10383)VUID-VkShaderCreateInfoEXT-codeType-10383  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and a [resource variable](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-resources) is declared in a shader, the corresponding descriptor set in `pSetLayouts` **must** match the shader stage
- [](#VUID-VkShaderCreateInfoEXT-codeType-10384)VUID-VkShaderCreateInfoEXT-codeType-10384  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and a [resource variable](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-resources) is declared in a shader, and the descriptor type is not `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, the corresponding descriptor set in `pSetLayouts` **must** match the descriptor type
- [](#VUID-VkShaderCreateInfoEXT-codeType-10385)VUID-VkShaderCreateInfoEXT-codeType-10385  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and a [resource variable](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-resources) is declared in a shader as an array, the corresponding descriptor set in `pSetLayouts` **must** match the descriptor count
- [](#VUID-VkShaderCreateInfoEXT-codeType-10386)VUID-VkShaderCreateInfoEXT-codeType-10386  
  If `codeType` is `VK_SHADER_CODE_TYPE_SPIRV_EXT`, and a [resource variable](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-resources) is declared in a shader as an array of descriptors, then the descriptor type of that variable **must** not be `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`

Valid Usage (Implicit)

- [](#VUID-VkShaderCreateInfoEXT-sType-sType)VUID-VkShaderCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SHADER_CREATE_INFO_EXT`
- [](#VUID-VkShaderCreateInfoEXT-pNext-pNext)VUID-VkShaderCreateInfoEXT-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html) or [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeaturesEXT.html)
- [](#VUID-VkShaderCreateInfoEXT-sType-unique)VUID-VkShaderCreateInfoEXT-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkShaderCreateInfoEXT-flags-parameter)VUID-VkShaderCreateInfoEXT-flags-parameter  
  `flags` **must** be a valid combination of [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateFlagBitsEXT.html) values
- [](#VUID-VkShaderCreateInfoEXT-stage-parameter)VUID-VkShaderCreateInfoEXT-stage-parameter  
  `stage` **must** be a valid [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) value
- [](#VUID-VkShaderCreateInfoEXT-nextStage-parameter)VUID-VkShaderCreateInfoEXT-nextStage-parameter  
  `nextStage` **must** be a valid combination of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values
- [](#VUID-VkShaderCreateInfoEXT-codeType-parameter)VUID-VkShaderCreateInfoEXT-codeType-parameter  
  `codeType` **must** be a valid [VkShaderCodeTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCodeTypeEXT.html) value
- [](#VUID-VkShaderCreateInfoEXT-pCode-parameter)VUID-VkShaderCreateInfoEXT-pCode-parameter  
  `pCode` **must** be a valid pointer to an array of `codeSize` bytes
- [](#VUID-VkShaderCreateInfoEXT-pName-parameter)VUID-VkShaderCreateInfoEXT-pName-parameter  
  If `pName` is not `NULL`, `pName` **must** be a null-terminated UTF-8 string
- [](#VUID-VkShaderCreateInfoEXT-pSetLayouts-parameter)VUID-VkShaderCreateInfoEXT-pSetLayouts-parameter  
  If `setLayoutCount` is not `0`, and `pSetLayouts` is not `NULL`, `pSetLayouts` **must** be a valid pointer to an array of `setLayoutCount` valid [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) handles
- [](#VUID-VkShaderCreateInfoEXT-pPushConstantRanges-parameter)VUID-VkShaderCreateInfoEXT-pPushConstantRanges-parameter  
  If `pushConstantRangeCount` is not `0`, and `pPushConstantRanges` is not `NULL`, `pPushConstantRanges` **must** be a valid pointer to an array of `pushConstantRangeCount` valid [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantRange.html) structures
- [](#VUID-VkShaderCreateInfoEXT-pSpecializationInfo-parameter)VUID-VkShaderCreateInfoEXT-pSpecializationInfo-parameter  
  If `pSpecializationInfo` is not `NULL`, `pSpecializationInfo` **must** be a valid pointer to a valid [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSpecializationInfo.html) structure
- [](#VUID-VkShaderCreateInfoEXT-codeSize-arraylength)VUID-VkShaderCreateInfoEXT-codeSize-arraylength  
  `codeSize` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html), [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantRange.html), [VkShaderCodeTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCodeTypeEXT.html), [VkShaderCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateFlagsEXT.html), [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSpecializationInfo.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateShadersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateShadersEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkShaderCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0