# vkCmdPipelineBarrier(3) Manual Page

## Name

vkCmdPipelineBarrier - Insert a memory dependency



## [](#_c_specification)C Specification

To record a pipeline barrier, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdPipelineBarrier(
    VkCommandBuffer                             commandBuffer,
    VkPipelineStageFlags                        srcStageMask,
    VkPipelineStageFlags                        dstStageMask,
    VkDependencyFlags                           dependencyFlags,
    uint32_t                                    memoryBarrierCount,
    const VkMemoryBarrier*                      pMemoryBarriers,
    uint32_t                                    bufferMemoryBarrierCount,
    const VkBufferMemoryBarrier*                pBufferMemoryBarriers,
    uint32_t                                    imageMemoryBarrierCount,
    const VkImageMemoryBarrier*                 pImageMemoryBarriers);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `srcStageMask` is a bitmask of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) specifying the [source stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks).
- `dstStageMask` is a bitmask of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) specifying the [destination stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks).
- `dependencyFlags` is a bitmask of [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlagBits.html) specifying how execution and memory dependencies are formed.
- `memoryBarrierCount` is the length of the `pMemoryBarriers` array.
- `pMemoryBarriers` is a pointer to an array of [VkMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier.html) structures.
- `bufferMemoryBarrierCount` is the length of the `pBufferMemoryBarriers` array.
- `pBufferMemoryBarriers` is a pointer to an array of [VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier.html) structures.
- `imageMemoryBarrierCount` is the length of the `pImageMemoryBarriers` array.
- `pImageMemoryBarriers` is a pointer to an array of [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier.html) structures.

## [](#_description)Description

`vkCmdPipelineBarrier` operates almost identically to [vkCmdPipelineBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier2.html), except that the scopes and barriers are defined as direct parameters rather than being defined by a [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html).

When [vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier.html) is submitted to a queue, it defines a memory dependency between commands that were submitted to the same queue before it, and those submitted to the same queue after it.

If [vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier.html) was recorded outside a render pass instance, the first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order). If [vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier.html) was recorded inside a render pass instance, the first synchronization scope includes only commands that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order) within the same subpass. In either case, the first synchronization scope is limited to operations on the pipeline stages determined by the [source stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks) specified by `srcStageMask`.

If [vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier.html) was recorded outside a render pass instance, the second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands that occur later in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order). If [vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier.html) was recorded inside a render pass instance, the second synchronization scope includes only commands that occur later in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order) within the same subpass. In either case, the second synchronization scope is limited to operations on the pipeline stages determined by the [destination stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks) specified by `dstStageMask`.

The first [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) is limited to accesses in the pipeline stages determined by the [source stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks) specified by `srcStageMask`. Within that, the first access scope only includes the first access scopes defined by elements of the `pMemoryBarriers`, `pBufferMemoryBarriers` and `pImageMemoryBarriers` arrays, which each define a set of [memory barriers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-memory-barriers). If no memory barriers are specified, then the first access scope includes no accesses.

The second [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) is limited to accesses in the pipeline stages determined by the [destination stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks) specified by `dstStageMask`. Within that, the second access scope only includes the second access scopes defined by elements of the `pMemoryBarriers`, `pBufferMemoryBarriers` and `pImageMemoryBarriers` arrays, which each define a set of [memory barriers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-memory-barriers). If no memory barriers are specified, then the second access scope includes no accesses.

If `dependencyFlags` includes `VK_DEPENDENCY_BY_REGION_BIT`, then any dependency between [framebuffer-space](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-framebuffer-regions) pipeline stages is [framebuffer-local](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-framebuffer-regions) - otherwise it is [framebuffer-global](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-framebuffer-regions).

Valid Usage

- [](#VUID-vkCmdPipelineBarrier-srcStageMask-04090)VUID-vkCmdPipelineBarrier-srcStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-04091)VUID-vkCmdPipelineBarrier-srcStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-04092)VUID-vkCmdPipelineBarrier-srcStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-04093)VUID-vkCmdPipelineBarrier-srcStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-04094)VUID-vkCmdPipelineBarrier-srcStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-04095)VUID-vkCmdPipelineBarrier-srcStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-04096)VUID-vkCmdPipelineBarrier-srcStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-07318)VUID-vkCmdPipelineBarrier-srcStageMask-07318  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-03937)VUID-vkCmdPipelineBarrier-srcStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `srcStageMask` **must** not be `0`
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-07949)VUID-vkCmdPipelineBarrier-srcStageMask-07949  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-10754)VUID-vkCmdPipelineBarrier-srcStageMask-10754  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `srcStageMask` **must** not contain `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`

<!--THE END-->

- [](#VUID-vkCmdPipelineBarrier-srcAccessMask-06257)VUID-vkCmdPipelineBarrier-srcAccessMask-06257  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and a memory barrier `srcAccessMask` includes `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `srcStageMask` **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages except `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

<!--THE END-->

- [](#VUID-vkCmdPipelineBarrier-dstStageMask-04090)VUID-vkCmdPipelineBarrier-dstStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`
- [](#VUID-vkCmdPipelineBarrier-dstStageMask-04091)VUID-vkCmdPipelineBarrier-dstStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-vkCmdPipelineBarrier-dstStageMask-04092)VUID-vkCmdPipelineBarrier-dstStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-vkCmdPipelineBarrier-dstStageMask-04093)VUID-vkCmdPipelineBarrier-dstStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-vkCmdPipelineBarrier-dstStageMask-04094)VUID-vkCmdPipelineBarrier-dstStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-vkCmdPipelineBarrier-dstStageMask-04095)VUID-vkCmdPipelineBarrier-dstStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`
- [](#VUID-vkCmdPipelineBarrier-dstStageMask-04096)VUID-vkCmdPipelineBarrier-dstStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`
- [](#VUID-vkCmdPipelineBarrier-dstStageMask-07318)VUID-vkCmdPipelineBarrier-dstStageMask-07318  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-vkCmdPipelineBarrier-dstStageMask-03937)VUID-vkCmdPipelineBarrier-dstStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `dstStageMask` **must** not be `0`
- [](#VUID-vkCmdPipelineBarrier-dstStageMask-07949)VUID-vkCmdPipelineBarrier-dstStageMask-07949  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-vkCmdPipelineBarrier-dstStageMask-10754)VUID-vkCmdPipelineBarrier-dstStageMask-10754  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `dstStageMask` **must** not contain `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`

<!--THE END-->

- [](#VUID-vkCmdPipelineBarrier-dstAccessMask-06257)VUID-vkCmdPipelineBarrier-dstAccessMask-06257  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and a memory barrier `dstAccessMask` includes `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `dstStageMask` **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT` stages except `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

<!--THE END-->

- [](#VUID-vkCmdPipelineBarrier-srcAccessMask-02815)VUID-vkCmdPipelineBarrier-srcAccessMask-02815  
  The `srcAccessMask` member of each element of `pMemoryBarriers` **must** only include access flags that are supported by one or more of the pipeline stages in `srcStageMask`, as specified in the [table of supported access types](#synchronization-access-types-supported)
- [](#VUID-vkCmdPipelineBarrier-dstAccessMask-02816)VUID-vkCmdPipelineBarrier-dstAccessMask-02816  
  The `dstAccessMask` member of each element of `pMemoryBarriers` **must** only include access flags that are supported by one or more of the pipeline stages in `dstStageMask`, as specified in the [table of supported access types](#synchronization-access-types-supported)
- [](#VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-02817)VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-02817  
  For any element of `pBufferMemoryBarriers`, if its `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members are equal, or if its `srcQueueFamilyIndex` is the queue family index that was used to create the command pool that `commandBuffer` was allocated from, then its `srcAccessMask` member **must** only contain access flags that are supported by one or more of the pipeline stages in `srcStageMask`, as specified in the [table of supported access types](#synchronization-access-types-supported)
- [](#VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-02818)VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-02818  
  For any element of `pBufferMemoryBarriers`, if its `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members are equal, or if its `dstQueueFamilyIndex` is the queue family index that was used to create the command pool that `commandBuffer` was allocated from, then its `dstAccessMask` member **must** only contain access flags that are supported by one or more of the pipeline stages in `dstStageMask`, as specified in the [table of supported access types](#synchronization-access-types-supported)
- [](#VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-02819)VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-02819  
  For any element of `pImageMemoryBarriers`, if its `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members are equal, or if its `srcQueueFamilyIndex` is the queue family index that was used to create the command pool that `commandBuffer` was allocated from, then its `srcAccessMask` member **must** only contain access flags that are supported by one or more of the pipeline stages in `srcStageMask`, as specified in the [table of supported access types](#synchronization-access-types-supported)
- [](#VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-02820)VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-02820  
  For any element of `pImageMemoryBarriers`, if its `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members are equal, or if its `dstQueueFamilyIndex` is the queue family index that was used to create the command pool that `commandBuffer` was allocated from, then its `dstAccessMask` member **must** only contain access flags that are supported by one or more of the pipeline stages in `dstStageMask`, as specified in the [table of supported access types](#synchronization-access-types-supported)

<!--THE END-->

- [](#VUID-vkCmdPipelineBarrier-None-07889)VUID-vkCmdPipelineBarrier-None-07889  
  If `vkCmdPipelineBarrier` is called within a render pass instance using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) object, the render pass **must** have been created with at least one subpass dependency that expresses a dependency from the current subpass to itself, does not include `VK_DEPENDENCY_BY_REGION_BIT` if this command does not, does not include `VK_DEPENDENCY_VIEW_LOCAL_BIT` if this command does not, and has [synchronization scopes](#synchronization-dependencies-scopes) and [access scopes](#synchronization-dependencies-access-scopes) that are all supersets of the scopes defined in this command
- [](#VUID-vkCmdPipelineBarrier-bufferMemoryBarrierCount-01178)VUID-vkCmdPipelineBarrier-bufferMemoryBarrierCount-01178  
  If `vkCmdPipelineBarrier` is called within a render pass instance using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) object, it **must** not include any buffer memory barriers
- [](#VUID-vkCmdPipelineBarrier-image-04073)VUID-vkCmdPipelineBarrier-image-04073  
  If `vkCmdPipelineBarrier` is called within a render pass instance using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) object, the `image` member of any image memory barrier included in this command **must** be an attachment used in the current subpass both as an input attachment, and as either a color, color resolve, or depth/stencil attachment
- [](#VUID-vkCmdPipelineBarrier-image-09373)VUID-vkCmdPipelineBarrier-image-09373  
  If `vkCmdPipelineBarrier` is called within a render pass instance using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) object, and the `image` member of any image memory barrier is a color resolve attachment, the corresponding color attachment **must** be `VK_ATTACHMENT_UNUSED`
- [](#VUID-vkCmdPipelineBarrier-image-09374)VUID-vkCmdPipelineBarrier-image-09374  
  If `vkCmdPipelineBarrier` is called within a render pass instance using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) object, and the `image` member of any image memory barrier is a color resolve attachment, it **must** have been created with a non-zero [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html)::`externalFormat` value
- [](#VUID-vkCmdPipelineBarrier-oldLayout-01181)VUID-vkCmdPipelineBarrier-oldLayout-01181  
  If `vkCmdPipelineBarrier` is called within a render pass instance, the `oldLayout` and `newLayout` members of any image memory barrier included in this command **must** be equal
- [](#VUID-vkCmdPipelineBarrier-srcQueueFamilyIndex-01182)VUID-vkCmdPipelineBarrier-srcQueueFamilyIndex-01182  
  If `vkCmdPipelineBarrier` is called within a render pass instance, the `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members of any memory barrier included in this command **must** be equal
- [](#VUID-vkCmdPipelineBarrier-None-07890)VUID-vkCmdPipelineBarrier-None-07890  
  If `vkCmdPipelineBarrier` is called within a render pass instance, and the source stage masks of any memory barriers include [framebuffer-space stages](#synchronization-framebuffer-regions), destination stage masks of all memory barriers **must** only include [framebuffer-space stages](#synchronization-framebuffer-regions)
- [](#VUID-vkCmdPipelineBarrier-dependencyFlags-07891)VUID-vkCmdPipelineBarrier-dependencyFlags-07891  
  If `vkCmdPipelineBarrier` is called within a render pass instance, and the source stage masks of any memory barriers include [framebuffer-space stages](#synchronization-framebuffer-regions), then `dependencyFlags` **must** include `VK_DEPENDENCY_BY_REGION_BIT`
- [](#VUID-vkCmdPipelineBarrier-None-07892)VUID-vkCmdPipelineBarrier-None-07892  
  If `vkCmdPipelineBarrier` is called within a render pass instance, the source and destination stage masks of any memory barriers **must** only include graphics pipeline stages
- [](#VUID-vkCmdPipelineBarrier-dependencyFlags-01186)VUID-vkCmdPipelineBarrier-dependencyFlags-01186  
  If `vkCmdPipelineBarrier` is called outside of a render pass instance, the dependency flags **must** not include `VK_DEPENDENCY_VIEW_LOCAL_BIT`
- [](#VUID-vkCmdPipelineBarrier-None-07893)VUID-vkCmdPipelineBarrier-None-07893  
  If `vkCmdPipelineBarrier` is called inside a render pass instance, and there is more than one view in the current subpass, dependency flags **must** include `VK_DEPENDENCY_VIEW_LOCAL_BIT`
- [](#VUID-vkCmdPipelineBarrier-None-09553)VUID-vkCmdPipelineBarrier-None-09553  
  If none of the [`shaderTileImageColorReadAccess`](#features-shaderTileImageColorReadAccess), [`shaderTileImageStencilReadAccess`](#features-shaderTileImageStencilReadAccess), or [`shaderTileImageDepthReadAccess`](#features-shaderTileImageDepthReadAccess) features are enabled, and the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, `vkCmdPipelineBarrier` **must** not be called within a render pass instance started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html)
- [](#VUID-vkCmdPipelineBarrier-None-09554)VUID-vkCmdPipelineBarrier-None-09554  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, and `vkCmdPipelineBarrier` is called within a render pass instance started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html), there **must** be no buffer or image memory barriers specified by this command
- [](#VUID-vkCmdPipelineBarrier-None-09586)VUID-vkCmdPipelineBarrier-None-09586  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, and `vkCmdPipelineBarrier` is called within a render pass instance started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html), memory barriers specified by this command **must** only include `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT`, `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT`, `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT`, or `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT` in their access masks
- [](#VUID-vkCmdPipelineBarrier-image-09555)VUID-vkCmdPipelineBarrier-image-09555  
  If `vkCmdPipelineBarrier` is called within a render pass instance started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html), and the `image` member of any image memory barrier is used as an attachment in the current render pass instance, it **must** be in the `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ` or `VK_IMAGE_LAYOUT_GENERAL` layout
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-09556)VUID-vkCmdPipelineBarrier-srcStageMask-09556  
  If `vkCmdPipelineBarrier` is called within a render pass instance started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html), this command **must** only specify [framebuffer-space stages](#synchronization-framebuffer-regions) in `srcStageMask` and `dstStageMask`
- [](#VUID-vkCmdPipelineBarrier-oldLayout-10758)VUID-vkCmdPipelineBarrier-oldLayout-10758  
  If called within a render pass instance using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) object, the `oldLayout` member of any image memory barrier included in this command **must** be equal to the layout that the corresponding attachment uses during the subpass
- [](#VUID-vkCmdPipelineBarrier-oldLayout-10759)VUID-vkCmdPipelineBarrier-oldLayout-10759  
  If called within a render pass instance started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html), the `oldLayout` member of any image memory barrier included in this command **must** be equal to the layout that the corresponding attachment uses during the render pass instance
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-06461)VUID-vkCmdPipelineBarrier-srcStageMask-06461  
  Any pipeline stage included in `srcStageMask` **must** be supported by the capabilities of the queue family specified by the `queueFamilyIndex` member of the [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) structure that was used to create the `VkCommandPool` that `commandBuffer` was allocated from, as specified in the [table of supported pipeline stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-supported)
- [](#VUID-vkCmdPipelineBarrier-dstStageMask-06462)VUID-vkCmdPipelineBarrier-dstStageMask-06462  
  Any pipeline stage included in `dstStageMask` **must** be supported by the capabilities of the queue family specified by the `queueFamilyIndex` member of the [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) structure that was used to create the `VkCommandPool` that `commandBuffer` was allocated from, as specified in the [table of supported pipeline stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-supported)
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-09633)VUID-vkCmdPipelineBarrier-srcStageMask-09633  
  If either `srcStageMask` or `dstStageMask` includes `VK_PIPELINE_STAGE_HOST_BIT`, for any element of `pImageMemoryBarriers`, `srcQueueFamilyIndex` and `dstQueueFamilyIndex` **must** be equal
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-09634)VUID-vkCmdPipelineBarrier-srcStageMask-09634  
  If either `srcStageMask` or `dstStageMask` includes `VK_PIPELINE_STAGE_HOST_BIT`, for any element of `pBufferMemoryBarriers`, `srcQueueFamilyIndex` and `dstQueueFamilyIndex` **must** be equal
- [](#VUID-vkCmdPipelineBarrier-srcQueueFamilyIndex-10388)VUID-vkCmdPipelineBarrier-srcQueueFamilyIndex-10388  
  If a buffer or image memory barrier specifies a [queue family ownership transfer operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers), either the `srcQueueFamilyIndex` or `dstQueueFamilyIndex` member and the queue family index that was used to create the command pool that `commandBuffer` was allocated from **must** be equal
- [](#VUID-vkCmdPipelineBarrier-maintenance8-10206)VUID-vkCmdPipelineBarrier-maintenance8-10206  
  If the [`maintenance8`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance8) feature is not enabled, `dependencyFlags` **must** not include `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-vkCmdPipelineBarrier-commandBuffer-parameter)VUID-vkCmdPipelineBarrier-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdPipelineBarrier-srcStageMask-parameter)VUID-vkCmdPipelineBarrier-srcStageMask-parameter  
  `srcStageMask` **must** be a valid combination of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) values
- [](#VUID-vkCmdPipelineBarrier-dstStageMask-parameter)VUID-vkCmdPipelineBarrier-dstStageMask-parameter  
  `dstStageMask` **must** be a valid combination of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) values
- [](#VUID-vkCmdPipelineBarrier-dependencyFlags-parameter)VUID-vkCmdPipelineBarrier-dependencyFlags-parameter  
  `dependencyFlags` **must** be a valid combination of [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlagBits.html) values
- [](#VUID-vkCmdPipelineBarrier-pMemoryBarriers-parameter)VUID-vkCmdPipelineBarrier-pMemoryBarriers-parameter  
  If `memoryBarrierCount` is not `0`, `pMemoryBarriers` **must** be a valid pointer to an array of `memoryBarrierCount` valid [VkMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier.html) structures
- [](#VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-parameter)VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-parameter  
  If `bufferMemoryBarrierCount` is not `0`, `pBufferMemoryBarriers` **must** be a valid pointer to an array of `bufferMemoryBarrierCount` valid [VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier.html) structures
- [](#VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-parameter)VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-parameter  
  If `imageMemoryBarrierCount` is not `0`, `pImageMemoryBarriers` **must** be a valid pointer to an array of `imageMemoryBarrierCount` valid [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier.html) structures
- [](#VUID-vkCmdPipelineBarrier-commandBuffer-recording)VUID-vkCmdPipelineBarrier-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdPipelineBarrier-commandBuffer-cmdpool)VUID-vkCmdPipelineBarrier-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, VK\_QUEUE\_GRAPHICS\_BIT, VK\_QUEUE\_TRANSFER\_BIT, VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR, or VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR operations

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Both

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT  
VK\_QUEUE\_TRANSFER\_BIT  
VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR  
VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR

Synchronization

Conditional Rendering

vkCmdPipelineBarrier is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDependencyFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlags.html), [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier.html), [VkMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier.html), [VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdPipelineBarrier).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0