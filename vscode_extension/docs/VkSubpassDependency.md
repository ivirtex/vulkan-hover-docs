# VkSubpassDependency(3) Manual Page

## Name

VkSubpassDependency - Structure specifying a subpass dependency



## [](#_c_specification)C Specification

The `VkSubpassDependency` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.2](#versions-1.2). See [Deprecated Functionality](#deprecation-renderpass2) for more information.

```c++
// Provided by VK_VERSION_1_0
typedef struct VkSubpassDependency {
    uint32_t                srcSubpass;
    uint32_t                dstSubpass;
    VkPipelineStageFlags    srcStageMask;
    VkPipelineStageFlags    dstStageMask;
    VkAccessFlags           srcAccessMask;
    VkAccessFlags           dstAccessMask;
    VkDependencyFlags       dependencyFlags;
} VkSubpassDependency;
```

## [](#_members)Members

- `srcSubpass` is the subpass index of the first subpass in the dependency, or `VK_SUBPASS_EXTERNAL`.
- `dstSubpass` is the subpass index of the second subpass in the dependency, or `VK_SUBPASS_EXTERNAL`.
- `srcStageMask` is a bitmask of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) specifying the [source stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks). If set to `VK_PIPELINE_STAGE_ALL_COMMANDS_BIT`, it is equivalent to setting it to `VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT`.
- `dstStageMask` is a bitmask of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) specifying the [destination stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks) If set to `VK_PIPELINE_STAGE_ALL_COMMANDS_BIT`, it is equivalent to setting it to `VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT`.
- `srcAccessMask` is a bitmask of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) specifying a [source access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks).
- `dstAccessMask` is a bitmask of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) specifying a [destination access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks).
- `dependencyFlags` is a bitmask of [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlagBits.html).

## [](#_description)Description

If `srcSubpass` is equal to `dstSubpass` then the [VkSubpassDependency](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency.html) does not directly define a [dependency](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies). Instead, it enables pipeline barriers to be used in a render pass instance within the identified subpass, where the scopes of one pipeline barrier **must** be a subset of those described by one subpass dependency. Subpass dependencies specified in this way that include [framebuffer-space stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-framebuffer-regions) in the `srcStageMask` **must** only include [framebuffer-space stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-framebuffer-regions) in `dstStageMask`, and **must** include `VK_DEPENDENCY_BY_REGION_BIT`. When a subpass dependency is specified in this way for a subpass that has more than one view in its view mask, its `dependencyFlags` **must** include `VK_DEPENDENCY_VIEW_LOCAL_BIT`.

If `srcSubpass` and `dstSubpass` are not equal, when a render pass instance which includes a subpass dependency is submitted to a queue, it defines a [dependency](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) between the subpasses identified by `srcSubpass` and `dstSubpass`.

If `srcSubpass` is equal to `VK_SUBPASS_EXTERNAL`, the first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes commands that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order) than the [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRenderPass.html) used to begin the render pass instance. Otherwise, the first set of commands includes all commands submitted as part of the subpass instance identified by `srcSubpass` and any [load](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations), [store](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-store-operations), or [multisample resolve](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-resolve-operations) operations on attachments used in `srcSubpass`. In either case, the first synchronization scope is limited to operations on the pipeline stages determined by the [source stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks) specified by `srcStageMask`.

If `dstSubpass` is equal to `VK_SUBPASS_EXTERNAL`, the second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes commands that occur later in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order) than the [vkCmdEndRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndRenderPass.html) used to end the render pass instance. Otherwise, the second set of commands includes all commands submitted as part of the subpass instance identified by `dstSubpass` and any [load](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations), [store](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-store-operations), and [multisample resolve](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-resolve-operations) operations on attachments used in `dstSubpass`. In either case, the second synchronization scope is limited to operations on the pipeline stages determined by the [destination stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks) specified by `dstStageMask`.

The first [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) is limited to accesses in the pipeline stages determined by the [source stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks) specified by `srcStageMask`. It is also limited to access types in the [source access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks) specified by `srcAccessMask`.

The second [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) is limited to accesses in the pipeline stages determined by the [destination stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks) specified by `dstStageMask`. It is also limited to access types in the [destination access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks) specified by `dstAccessMask`.

The [availability and visibility operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-available-and-visible) defined by a subpass dependency affect the execution of [image layout transitions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-layout-transitions) within the render pass.

Note

For non-attachment resources, the memory dependency expressed by subpass dependency is nearly identical to that of a [VkMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier.html) (with matching `srcAccessMask` and `dstAccessMask` parameters) submitted as a part of a [vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier.html) (with matching `srcStageMask` and `dstStageMask` parameters). The only difference being that its scopes are limited to the identified subpasses rather than potentially affecting everything before and after.

For attachments however, subpass dependencies work more like a [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier.html) defined similarly to the [VkMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier.html) above, the queue family indices set to `VK_QUEUE_FAMILY_IGNORED`, and layouts as follows:

- The equivalent to `oldLayout` is the attachment’s layout according to the subpass description for `srcSubpass`.
- The equivalent to `newLayout` is the attachment’s layout according to the subpass description for `dstSubpass`.

Valid Usage

- [](#VUID-VkSubpassDependency-srcStageMask-04090)VUID-VkSubpassDependency-srcStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`
- [](#VUID-VkSubpassDependency-srcStageMask-04091)VUID-VkSubpassDependency-srcStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-VkSubpassDependency-srcStageMask-04092)VUID-VkSubpassDependency-srcStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-VkSubpassDependency-srcStageMask-04093)VUID-VkSubpassDependency-srcStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-VkSubpassDependency-srcStageMask-04094)VUID-VkSubpassDependency-srcStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-VkSubpassDependency-srcStageMask-04095)VUID-VkSubpassDependency-srcStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`
- [](#VUID-VkSubpassDependency-srcStageMask-04096)VUID-VkSubpassDependency-srcStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`
- [](#VUID-VkSubpassDependency-srcStageMask-07318)VUID-VkSubpassDependency-srcStageMask-07318  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkSubpassDependency-srcStageMask-03937)VUID-VkSubpassDependency-srcStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `srcStageMask` **must** not be `0`
- [](#VUID-VkSubpassDependency-srcStageMask-07949)VUID-VkSubpassDependency-srcStageMask-07949  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkSubpassDependency-srcStageMask-10754)VUID-VkSubpassDependency-srcStageMask-10754  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`

<!--THE END-->

- [](#VUID-VkSubpassDependency-dstStageMask-04090)VUID-VkSubpassDependency-dstStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`
- [](#VUID-VkSubpassDependency-dstStageMask-04091)VUID-VkSubpassDependency-dstStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-VkSubpassDependency-dstStageMask-04092)VUID-VkSubpassDependency-dstStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-VkSubpassDependency-dstStageMask-04093)VUID-VkSubpassDependency-dstStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-VkSubpassDependency-dstStageMask-04094)VUID-VkSubpassDependency-dstStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-VkSubpassDependency-dstStageMask-04095)VUID-VkSubpassDependency-dstStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`
- [](#VUID-VkSubpassDependency-dstStageMask-04096)VUID-VkSubpassDependency-dstStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`
- [](#VUID-VkSubpassDependency-dstStageMask-07318)VUID-VkSubpassDependency-dstStageMask-07318  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkSubpassDependency-dstStageMask-03937)VUID-VkSubpassDependency-dstStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `dstStageMask` **must** not be `0`
- [](#VUID-VkSubpassDependency-dstStageMask-07949)VUID-VkSubpassDependency-dstStageMask-07949  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkSubpassDependency-dstStageMask-10754)VUID-VkSubpassDependency-dstStageMask-10754  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-VkSubpassDependency-srcSubpass-00864)VUID-VkSubpassDependency-srcSubpass-00864  
  `srcSubpass` **must** be less than or equal to `dstSubpass`, unless one of them is `VK_SUBPASS_EXTERNAL`, to avoid cyclic dependencies and ensure a valid execution order
- [](#VUID-VkSubpassDependency-srcSubpass-00865)VUID-VkSubpassDependency-srcSubpass-00865  
  `srcSubpass` and `dstSubpass` **must** not both be equal to `VK_SUBPASS_EXTERNAL`
- [](#VUID-VkSubpassDependency-srcSubpass-06809)VUID-VkSubpassDependency-srcSubpass-06809  
  If `srcSubpass` is equal to `dstSubpass` and `srcStageMask` includes a [framebuffer-space stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-framebuffer-regions), `dstStageMask` **must** only contain [framebuffer-space stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-framebuffer-regions)
- [](#VUID-VkSubpassDependency-srcAccessMask-00868)VUID-VkSubpassDependency-srcAccessMask-00868  
  Any access flag included in `srcAccessMask` **must** be supported by one of the pipeline stages in `srcStageMask`, as specified in the [table of supported access types](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types-supported)
- [](#VUID-VkSubpassDependency-dstAccessMask-00869)VUID-VkSubpassDependency-dstAccessMask-00869  
  Any access flag included in `dstAccessMask` **must** be supported by one of the pipeline stages in `dstStageMask`, as specified in the [table of supported access types](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types-supported)
- [](#VUID-VkSubpassDependency-srcSubpass-02243)VUID-VkSubpassDependency-srcSubpass-02243  
  If `srcSubpass` equals `dstSubpass`, and `srcStageMask` and `dstStageMask` both include a [framebuffer-space stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-framebuffer-regions), then `dependencyFlags` **must** include `VK_DEPENDENCY_BY_REGION_BIT`
- [](#VUID-VkSubpassDependency-dependencyFlags-02520)VUID-VkSubpassDependency-dependencyFlags-02520  
  If `dependencyFlags` includes `VK_DEPENDENCY_VIEW_LOCAL_BIT`, `srcSubpass` **must** not be equal to `VK_SUBPASS_EXTERNAL`
- [](#VUID-VkSubpassDependency-dependencyFlags-02521)VUID-VkSubpassDependency-dependencyFlags-02521  
  If `dependencyFlags` includes `VK_DEPENDENCY_VIEW_LOCAL_BIT`, `dstSubpass` **must** not be equal to `VK_SUBPASS_EXTERNAL`
- [](#VUID-VkSubpassDependency-srcSubpass-00872)VUID-VkSubpassDependency-srcSubpass-00872  
  If `srcSubpass` equals `dstSubpass` and that subpass has more than one bit set in the view mask, then `dependencyFlags` **must** include `VK_DEPENDENCY_VIEW_LOCAL_BIT`
- [](#VUID-VkSubpassDependency-dependencyFlags-10203)VUID-VkSubpassDependency-dependencyFlags-10203  
  `dependencyFlags` **must** not include `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-VkSubpassDependency-srcStageMask-parameter)VUID-VkSubpassDependency-srcStageMask-parameter  
  `srcStageMask` **must** be a valid combination of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) values
- [](#VUID-VkSubpassDependency-dstStageMask-parameter)VUID-VkSubpassDependency-dstStageMask-parameter  
  `dstStageMask` **must** be a valid combination of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) values
- [](#VUID-VkSubpassDependency-srcAccessMask-parameter)VUID-VkSubpassDependency-srcAccessMask-parameter  
  `srcAccessMask` **must** be a valid combination of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) values
- [](#VUID-VkSubpassDependency-dstAccessMask-parameter)VUID-VkSubpassDependency-dstAccessMask-parameter  
  `dstAccessMask` **must** be a valid combination of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) values
- [](#VUID-VkSubpassDependency-dependencyFlags-parameter)VUID-VkSubpassDependency-dependencyFlags-parameter  
  `dependencyFlags` **must** be a valid combination of [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAccessFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags.html), [VkDependencyFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlags.html), [VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags.html), [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubpassDependency)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0