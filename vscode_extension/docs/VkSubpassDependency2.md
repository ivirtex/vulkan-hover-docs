# VkSubpassDependency2(3) Manual Page

## Name

VkSubpassDependency2 - Structure specifying a subpass dependency



## [](#_c_specification)C Specification

The `VkSubpassDependency2` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_2
typedef struct VkSubpassDependency2 {
    VkStructureType         sType;
    const void*             pNext;
    uint32_t                srcSubpass;
    uint32_t                dstSubpass;
    VkPipelineStageFlags    srcStageMask;
    VkPipelineStageFlags    dstStageMask;
    VkAccessFlags           srcAccessMask;
    VkAccessFlags           dstAccessMask;
    VkDependencyFlags       dependencyFlags;
    int32_t                 viewOffset;
} VkSubpassDependency2;
```

or the equivalent

```c++
// Provided by VK_KHR_create_renderpass2
typedef VkSubpassDependency2 VkSubpassDependency2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcSubpass` is the subpass index of the first subpass in the dependency, or `VK_SUBPASS_EXTERNAL`.
- `dstSubpass` is the subpass index of the second subpass in the dependency, or `VK_SUBPASS_EXTERNAL`.
- `srcStageMask` is a bitmask of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) specifying the [source stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks). If set to `VK_PIPELINE_STAGE_ALL_COMMANDS_BIT`, it is equivalent to setting it to `VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT`.
- `dstStageMask` is a bitmask of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) specifying the [destination stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks) If set to `VK_PIPELINE_STAGE_ALL_COMMANDS_BIT`, it is equivalent to setting it to `VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT`.
- `srcAccessMask` is a bitmask of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) specifying a [source access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks).
- `dstAccessMask` is a bitmask of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) specifying a [destination access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks).
- `dependencyFlags` is a bitmask of [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlagBits.html).
- `viewOffset` controls which views in the source subpass the views in the destination subpass depend on.

## [](#_description)Description

Parameters defined by this structure with the same name as those in [VkSubpassDependency](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency.html) have the identical effect to those parameters.

`viewOffset` has the same effect for the described subpass dependency as [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html)::`pViewOffsets` has on each corresponding subpass dependency.

If a [VkMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier2.html) is included in the `pNext` chain, `srcStageMask`, `dstStageMask`, `srcAccessMask`, and `dstAccessMask` parameters are ignored. The synchronization and access scopes instead are defined by the parameters of [VkMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier2.html). If either `srcStageMask` or `dstStageMask` are set to `VK_PIPELINE_STAGE_ALL_COMMANDS_BIT`, it is equivalent to setting `VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT`.

Valid Usage

- [](#VUID-VkSubpassDependency2-srcStageMask-04090)VUID-VkSubpassDependency2-srcStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`
- [](#VUID-VkSubpassDependency2-srcStageMask-04091)VUID-VkSubpassDependency2-srcStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-VkSubpassDependency2-srcStageMask-04092)VUID-VkSubpassDependency2-srcStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-VkSubpassDependency2-srcStageMask-04093)VUID-VkSubpassDependency2-srcStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-VkSubpassDependency2-srcStageMask-04094)VUID-VkSubpassDependency2-srcStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-VkSubpassDependency2-srcStageMask-04095)VUID-VkSubpassDependency2-srcStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`
- [](#VUID-VkSubpassDependency2-srcStageMask-04096)VUID-VkSubpassDependency2-srcStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`
- [](#VUID-VkSubpassDependency2-srcStageMask-07318)VUID-VkSubpassDependency2-srcStageMask-07318  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkSubpassDependency2-srcStageMask-03937)VUID-VkSubpassDependency2-srcStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `srcStageMask` **must** not be `0`
- [](#VUID-VkSubpassDependency2-srcStageMask-07949)VUID-VkSubpassDependency2-srcStageMask-07949  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkSubpassDependency2-srcStageMask-10754)VUID-VkSubpassDependency2-srcStageMask-10754  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`

<!--THE END-->

- [](#VUID-VkSubpassDependency2-dstStageMask-04090)VUID-VkSubpassDependency2-dstStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`
- [](#VUID-VkSubpassDependency2-dstStageMask-04091)VUID-VkSubpassDependency2-dstStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-VkSubpassDependency2-dstStageMask-04092)VUID-VkSubpassDependency2-dstStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-VkSubpassDependency2-dstStageMask-04093)VUID-VkSubpassDependency2-dstStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-VkSubpassDependency2-dstStageMask-04094)VUID-VkSubpassDependency2-dstStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-VkSubpassDependency2-dstStageMask-04095)VUID-VkSubpassDependency2-dstStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`
- [](#VUID-VkSubpassDependency2-dstStageMask-04096)VUID-VkSubpassDependency2-dstStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`
- [](#VUID-VkSubpassDependency2-dstStageMask-07318)VUID-VkSubpassDependency2-dstStageMask-07318  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkSubpassDependency2-dstStageMask-03937)VUID-VkSubpassDependency2-dstStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `dstStageMask` **must** not be `0`
- [](#VUID-VkSubpassDependency2-dstStageMask-07949)VUID-VkSubpassDependency2-dstStageMask-07949  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkSubpassDependency2-dstStageMask-10754)VUID-VkSubpassDependency2-dstStageMask-10754  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-VkSubpassDependency2-srcSubpass-03084)VUID-VkSubpassDependency2-srcSubpass-03084  
  `srcSubpass` **must** be less than or equal to `dstSubpass`, unless one of them is `VK_SUBPASS_EXTERNAL`, to avoid cyclic dependencies and ensure a valid execution order
- [](#VUID-VkSubpassDependency2-srcSubpass-03085)VUID-VkSubpassDependency2-srcSubpass-03085  
  `srcSubpass` and `dstSubpass` **must** not both be equal to `VK_SUBPASS_EXTERNAL`
- [](#VUID-VkSubpassDependency2-srcSubpass-06810)VUID-VkSubpassDependency2-srcSubpass-06810  
  If `srcSubpass` is equal to `dstSubpass` and `srcStageMask` includes a [framebuffer-space stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-framebuffer-regions), `dstStageMask` **must** only contain [framebuffer-space stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-framebuffer-regions)
- [](#VUID-VkSubpassDependency2-srcAccessMask-03088)VUID-VkSubpassDependency2-srcAccessMask-03088  
  Any access flag included in `srcAccessMask` **must** be supported by one of the pipeline stages in `srcStageMask`, as specified in the [table of supported access types](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types-supported)
- [](#VUID-VkSubpassDependency2-dstAccessMask-03089)VUID-VkSubpassDependency2-dstAccessMask-03089  
  Any access flag included in `dstAccessMask` **must** be supported by one of the pipeline stages in `dstStageMask`, as specified in the [table of supported access types](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types-supported)
- [](#VUID-VkSubpassDependency2-dependencyFlags-03090)VUID-VkSubpassDependency2-dependencyFlags-03090  
  If `dependencyFlags` includes `VK_DEPENDENCY_VIEW_LOCAL_BIT`, `srcSubpass` **must** not be equal to `VK_SUBPASS_EXTERNAL`
- [](#VUID-VkSubpassDependency2-dependencyFlags-03091)VUID-VkSubpassDependency2-dependencyFlags-03091  
  If `dependencyFlags` includes `VK_DEPENDENCY_VIEW_LOCAL_BIT`, `dstSubpass` **must** not be equal to `VK_SUBPASS_EXTERNAL`
- [](#VUID-VkSubpassDependency2-srcSubpass-02245)VUID-VkSubpassDependency2-srcSubpass-02245  
  If `srcSubpass` equals `dstSubpass`, and `srcStageMask` and `dstStageMask` both include a [framebuffer-space stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-framebuffer-regions), then `dependencyFlags` **must** include `VK_DEPENDENCY_BY_REGION_BIT`
- [](#VUID-VkSubpassDependency2-viewOffset-02530)VUID-VkSubpassDependency2-viewOffset-02530  
  If `viewOffset` is not equal to `0`, `srcSubpass` **must** not be equal to `dstSubpass`
- [](#VUID-VkSubpassDependency2-dependencyFlags-03092)VUID-VkSubpassDependency2-dependencyFlags-03092  
  If `dependencyFlags` does not include `VK_DEPENDENCY_VIEW_LOCAL_BIT`, `viewOffset` **must** be `0`
- [](#VUID-VkSubpassDependency2-dependencyFlags-10204)VUID-VkSubpassDependency2-dependencyFlags-10204  
  `dependencyFlags` **must** not include `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-VkSubpassDependency2-sType-sType)VUID-VkSubpassDependency2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2`
- [](#VUID-VkSubpassDependency2-pNext-pNext)VUID-VkSubpassDependency2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier2.html) or [VkMemoryBarrierAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrierAccessFlags3KHR.html)
- [](#VUID-VkSubpassDependency2-sType-unique)VUID-VkSubpassDependency2-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkSubpassDependency2-srcStageMask-parameter)VUID-VkSubpassDependency2-srcStageMask-parameter  
  `srcStageMask` **must** be a valid combination of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) values
- [](#VUID-VkSubpassDependency2-dstStageMask-parameter)VUID-VkSubpassDependency2-dstStageMask-parameter  
  `dstStageMask` **must** be a valid combination of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) values
- [](#VUID-VkSubpassDependency2-srcAccessMask-parameter)VUID-VkSubpassDependency2-srcAccessMask-parameter  
  `srcAccessMask` **must** be a valid combination of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) values
- [](#VUID-VkSubpassDependency2-dstAccessMask-parameter)VUID-VkSubpassDependency2-dstAccessMask-parameter  
  `dstAccessMask` **must** be a valid combination of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) values
- [](#VUID-VkSubpassDependency2-dependencyFlags-parameter)VUID-VkSubpassDependency2-dependencyFlags-parameter  
  `dependencyFlags` **must** be a valid combination of [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlagBits.html) values

## [](#_see_also)See Also

[VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkAccessFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags.html), [VkDependencyFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlags.html), [VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags.html), [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubpassDependency2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0