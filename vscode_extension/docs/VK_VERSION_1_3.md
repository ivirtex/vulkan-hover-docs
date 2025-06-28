# VK\_VERSION\_1\_3(3) Manual Page

## Name

VK\_VERSION\_1\_3 - Vulkan version 1.3



## [](#_description)Description

Vulkan Version 1.3 [promoted](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-promotion) a number of key extensions into the core API:

- [VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html)
- [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html)
- [VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html)
- [VK\_KHR\_maintenance4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance4.html)
- [VK\_KHR\_shader\_integer\_dot\_product](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_integer_dot_product.html)
- [VK\_KHR\_shader\_non\_semantic\_info](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_non_semantic_info.html)
- [VK\_KHR\_shader\_terminate\_invocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_terminate_invocation.html)
- [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html)
- [VK\_KHR\_zero\_initialize\_workgroup\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_zero_initialize_workgroup_memory.html)
- [VK\_EXT\_4444\_formats](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_4444_formats.html)
- [VK\_EXT\_extended\_dynamic\_state](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state.html)
- [VK\_EXT\_extended\_dynamic\_state2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state2.html)
- [VK\_EXT\_image\_robustness](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_robustness.html)
- [VK\_EXT\_inline\_uniform\_block](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_inline_uniform_block.html)
- [VK\_EXT\_pipeline\_creation\_cache\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_creation_cache_control.html)
- [VK\_EXT\_pipeline\_creation\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_creation_feedback.html)
- [VK\_EXT\_private\_data](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_private_data.html)
- [VK\_EXT\_shader\_demote\_to\_helper\_invocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_demote_to_helper_invocation.html)
- [VK\_EXT\_subgroup\_size\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_subgroup_size_control.html)
- [VK\_EXT\_texel\_buffer\_alignment](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_texel_buffer_alignment.html)
- [VK\_EXT\_texture\_compression\_astc\_hdr](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_texture_compression_astc_hdr.html)
- [VK\_EXT\_tooling\_info](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_tooling_info.html)
- [VK\_EXT\_ycbcr\_2plane\_444\_formats](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_ycbcr_2plane_444_formats.html)

All differences in behavior between these extensions and the corresponding Vulkan 1.3 functionality are summarized below.

Differences Relative to `VK_EXT_4444_formats`

If the `VK_EXT_4444_formats` extension is not supported, support for all formats defined by it are optional in Vulkan 1.3. There are no members in the [VkPhysicalDeviceVulkan13Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan13Features.html) structure corresponding to the [VkPhysicalDevice4444FormatsFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice4444FormatsFeaturesEXT.html) structure.

Differences Relative to `VK_EXT_extended_dynamic_state`

All dynamic state enumerants and commands defined by `VK_EXT_extended_dynamic_state` are required in Vulkan 1.3. There are no members in the [VkPhysicalDeviceVulkan13Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan13Features.html) structure corresponding to the [VkPhysicalDeviceExtendedDynamicStateFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExtendedDynamicStateFeaturesEXT.html) structure.

Differences Relative to `VK_EXT_extended_dynamic_state2`

The optional dynamic state enumerants and commands defined by `VK_EXT_extended_dynamic_state2` for patch control points and logic op are not promoted in Vulkan 1.3. There are no members in the [VkPhysicalDeviceVulkan13Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan13Features.html) structure corresponding to the [VkPhysicalDeviceExtendedDynamicState2FeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExtendedDynamicState2FeaturesEXT.html) structure.

Differences Relative to `VK_EXT_texel_buffer_alignment`

The more specific alignment requirements defined by [VkPhysicalDeviceTexelBufferAlignmentProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTexelBufferAlignmentProperties.html) are required in Vulkan 1.3. There are no members in the [VkPhysicalDeviceVulkan13Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan13Features.html) structure corresponding to the [VkPhysicalDeviceTexelBufferAlignmentFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTexelBufferAlignmentFeaturesEXT.html) structure. The `texelBufferAlignment` feature is enabled if using a Vulkan 1.3 instance.

Differences Relative to `VK_EXT_texture_compression_astc_hdr`

If the `VK_EXT_texture_compression_astc_hdr` extension is not supported, support for all formats defined by it are optional in Vulkan 1.3. The [`textureCompressionASTC_HDR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-textureCompressionASTC_HDR) member of [VkPhysicalDeviceVulkan13Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan13Features.html) indicates whether a Vulkan 1.3 implementation supports these formats.

Differences Relative to `VK_EXT_ycbcr_2plane_444_formats`

If the `VK_EXT_ycbcr_2plane_444_formats` extension is not supported, support for all formats defined by it are optional in Vulkan 1.3. There are no members in the [VkPhysicalDeviceVulkan13Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan13Features.html) structure corresponding to the [VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT.html) structure.

Additional Vulkan 1.3 Feature Support

In addition to the promoted extensions described above, Vulkan 1.3 added required support for:

- SPIR-V version 1.6
  
  - SPIR-V 1.6 deprecates (but does not remove) the `WorkgroupSize` decoration.
- The [`bufferDeviceAddress`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddress) feature which indicates support for accessing memory in shaders as storage buffers via [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html).
- The [`vulkanMemoryModel`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-vulkanMemoryModel) and [`vulkanMemoryModelDeviceScope`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-vulkanMemoryModelDeviceScope) features, which indicate support for the corresponding Vulkan Memory Model capabilities.
- The [`maxInlineUniformTotalSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxInlineUniformTotalSize) limit is added to provide the total size of all inline uniform block bindings in a pipeline layout.

### [](#_new_macros)New Macros

- [VK\_API\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_API_VERSION_1_3.html)

### [](#_new_base_types)New Base Types

- [VkFlags64](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFlags64.html)

### [](#_new_object_types)New Object Types

- [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html)

### [](#_new_commands)New Commands

- [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html)
- [vkCmdBindVertexBuffers2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindVertexBuffers2.html)
- [vkCmdBlitImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage2.html)
- [vkCmdCopyBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBuffer2.html)
- [vkCmdCopyBufferToImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBufferToImage2.html)
- [vkCmdCopyImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImage2.html)
- [vkCmdCopyImageToBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImageToBuffer2.html)
- [vkCmdEndRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndRendering.html)
- [vkCmdPipelineBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier2.html)
- [vkCmdResetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResetEvent2.html)
- [vkCmdResolveImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResolveImage2.html)
- [vkCmdSetCullMode](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCullMode.html)
- [vkCmdSetDepthBiasEnable](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthBiasEnable.html)
- [vkCmdSetDepthBoundsTestEnable](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthBoundsTestEnable.html)
- [vkCmdSetDepthCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthCompareOp.html)
- [vkCmdSetDepthTestEnable](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthTestEnable.html)
- [vkCmdSetDepthWriteEnable](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthWriteEnable.html)
- [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2.html)
- [vkCmdSetFrontFace](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetFrontFace.html)
- [vkCmdSetPrimitiveRestartEnable](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPrimitiveRestartEnable.html)
- [vkCmdSetPrimitiveTopology](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPrimitiveTopology.html)
- [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRasterizerDiscardEnable.html)
- [vkCmdSetScissorWithCount](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetScissorWithCount.html)
- [vkCmdSetStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetStencilOp.html)
- [vkCmdSetStencilTestEnable](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetStencilTestEnable.html)
- [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetViewportWithCount.html)
- [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents2.html)
- [vkCmdWriteTimestamp2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteTimestamp2.html)
- [vkCreatePrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePrivateDataSlot.html)
- [vkDestroyPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyPrivateDataSlot.html)
- [vkGetDeviceBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceBufferMemoryRequirements.html)
- [vkGetDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageMemoryRequirements.html)
- [vkGetDeviceImageSparseMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSparseMemoryRequirements.html)
- [vkGetPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceToolProperties.html)
- [vkGetPrivateData](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPrivateData.html)
- [vkQueueSubmit2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit2.html)
- [vkSetPrivateData](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetPrivateData.html)

### [](#_new_structures)New Structures

- [VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlitImageInfo2.html)
- [VkBufferCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCopy2.html)
- [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html)
- [VkBufferMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier2.html)
- [VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferSubmitInfo.html)
- [VkCopyBufferInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyBufferInfo2.html)
- [VkCopyBufferToImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyBufferToImageInfo2.html)
- [VkCopyImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageInfo2.html)
- [VkCopyImageToBufferInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToBufferInfo2.html)
- [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html)
- [VkDeviceBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceBufferMemoryRequirements.html)
- [VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageMemoryRequirements.html)
- [VkImageBlit2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageBlit2.html)
- [VkImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy2.html)
- [VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier2.html)
- [VkImageResolve2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageResolve2.html)
- [VkPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceToolProperties.html)
- [VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedback.html)
- [VkPrivateDataSlotCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlotCreateInfo.html)
- [VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentInfo.html)
- [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)
- [VkResolveImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveImageInfo2.html)
- [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSubmitInfo.html)
- [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo2.html)
- Extending [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html):
  
  - [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceRenderingInfo.html)
- Extending [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html):
  
  - [VkDescriptorPoolInlineUniformBlockCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolInlineUniformBlockCreateInfo.html)
- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkDevicePrivateDataCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevicePrivateDataCreateInfo.html)
- Extending [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties2.html):
  
  - [VkFormatProperties3](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties3.html)
- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html):
  
  - [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRenderingCreateInfo.html)
- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html), [VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html), [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html):
  
  - [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackCreateInfo.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDynamicRenderingFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDynamicRenderingFeatures.html)
  - [VkPhysicalDeviceImageRobustnessFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageRobustnessFeatures.html)
  - [VkPhysicalDeviceInlineUniformBlockFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceInlineUniformBlockFeatures.html)
  - [VkPhysicalDeviceMaintenance4Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance4Features.html)
  - [VkPhysicalDevicePipelineCreationCacheControlFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineCreationCacheControlFeatures.html)
  - [VkPhysicalDevicePrivateDataFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePrivateDataFeatures.html)
  - [VkPhysicalDeviceShaderDemoteToHelperInvocationFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderDemoteToHelperInvocationFeatures.html)
  - [VkPhysicalDeviceShaderIntegerDotProductFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderIntegerDotProductFeatures.html)
  - [VkPhysicalDeviceShaderTerminateInvocationFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderTerminateInvocationFeatures.html)
  - [VkPhysicalDeviceSubgroupSizeControlFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupSizeControlFeatures.html)
  - [VkPhysicalDeviceSynchronization2Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSynchronization2Features.html)
  - [VkPhysicalDeviceTextureCompressionASTCHDRFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTextureCompressionASTCHDRFeatures.html)
  - [VkPhysicalDeviceVulkan13Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan13Features.html)
  - [VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceInlineUniformBlockProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceInlineUniformBlockProperties.html)
  - [VkPhysicalDeviceMaintenance4Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance4Properties.html)
  - [VkPhysicalDeviceShaderIntegerDotProductProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderIntegerDotProductProperties.html)
  - [VkPhysicalDeviceSubgroupSizeControlProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupSizeControlProperties.html)
  - [VkPhysicalDeviceTexelBufferAlignmentProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTexelBufferAlignmentProperties.html)
  - [VkPhysicalDeviceVulkan13Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan13Properties.html)
- Extending [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateInfoEXT.html):
  
  - [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)
- Extending [VkSubpassDependency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency2.html):
  
  - [VkMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier2.html)
- Extending [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html):
  
  - [VkWriteDescriptorSetInlineUniformBlock](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetInlineUniformBlock.html)

### [](#_new_enums)New Enums

- [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html)
- [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html)
- [VkPipelineCreationFeedbackFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackFlagBits.html)
- [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html)
- [VkRenderingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFlagBits.html)
- [VkSubmitFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitFlagBits.html)
- [VkToolPurposeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkToolPurposeFlagBits.html)

### [](#_new_bitmasks)New Bitmasks

- [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2.html)
- [VkFormatFeatureFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlags2.html)
- [VkPipelineCreationFeedbackFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackFlags.html)
- [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html)
- [VkPrivateDataSlotCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlotCreateFlags.html)
- [VkRenderingFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFlags.html)
- [VkSubmitFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitFlags.html)
- [VkToolPurposeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkToolPurposeFlags.html)

### [](#_new_enum_constants)New Enum Constants

- Extending [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html):
  
  - `VK_ACCESS_NONE`
- Extending [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html):
  
  - `VK_ATTACHMENT_STORE_OP_NONE`
- Extending [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html):
  
  - `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`
- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_CULL_MODE`
  - `VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE`
  - `VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE`
  - `VK_DYNAMIC_STATE_DEPTH_COMPARE_OP`
  - `VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE`
  - `VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE`
  - `VK_DYNAMIC_STATE_FRONT_FACE`
  - `VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE`
  - `VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY`
  - `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE`
  - `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT`
  - `VK_DYNAMIC_STATE_STENCIL_OP`
  - `VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE`
  - `VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE`
  - `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT`
- Extending [VkEventCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateFlagBits.html):
  
  - `VK_EVENT_CREATE_DEVICE_ONLY_BIT`
- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html):
  
  - `VK_FORMAT_A4B4G4R4_UNORM_PACK16`
  - `VK_FORMAT_A4R4G4B4_UNORM_PACK16`
  - `VK_FORMAT_ASTC_10x10_SFLOAT_BLOCK`
  - `VK_FORMAT_ASTC_10x5_SFLOAT_BLOCK`
  - `VK_FORMAT_ASTC_10x6_SFLOAT_BLOCK`
  - `VK_FORMAT_ASTC_10x8_SFLOAT_BLOCK`
  - `VK_FORMAT_ASTC_12x10_SFLOAT_BLOCK`
  - `VK_FORMAT_ASTC_12x12_SFLOAT_BLOCK`
  - `VK_FORMAT_ASTC_4x4_SFLOAT_BLOCK`
  - `VK_FORMAT_ASTC_5x4_SFLOAT_BLOCK`
  - `VK_FORMAT_ASTC_5x5_SFLOAT_BLOCK`
  - `VK_FORMAT_ASTC_6x5_SFLOAT_BLOCK`
  - `VK_FORMAT_ASTC_6x6_SFLOAT_BLOCK`
  - `VK_FORMAT_ASTC_8x5_SFLOAT_BLOCK`
  - `VK_FORMAT_ASTC_8x6_SFLOAT_BLOCK`
  - `VK_FORMAT_ASTC_8x8_SFLOAT_BLOCK`
  - `VK_FORMAT_G10X6_B10X6R10X6_2PLANE_444_UNORM_3PACK16`
  - `VK_FORMAT_G12X4_B12X4R12X4_2PLANE_444_UNORM_3PACK16`
  - `VK_FORMAT_G16_B16R16_2PLANE_444_UNORM`
  - `VK_FORMAT_G8_B8R8_2PLANE_444_UNORM`
- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_FILTER_CUBIC_BIT`
- Extending [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html):
  
  - `VK_IMAGE_ASPECT_NONE`
- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html):
  
  - `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL`
  - `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_PRIVATE_DATA_SLOT`
- Extending [VkPipelineCacheCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheCreateFlagBits.html):
  
  - `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`
- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html):
  
  - `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT`
  - `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT`
- Extending [VkPipelineShaderStageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateFlagBits.html):
  
  - `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT`
  - `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT`
- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html):
  
  - `VK_PIPELINE_STAGE_NONE`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_PIPELINE_COMPILE_REQUIRED`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BLIT_IMAGE_INFO_2`
  - `VK_STRUCTURE_TYPE_BUFFER_COPY_2`
  - `VK_STRUCTURE_TYPE_BUFFER_IMAGE_COPY_2`
  - `VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER_2`
  - `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_RENDERING_INFO`
  - `VK_STRUCTURE_TYPE_COMMAND_BUFFER_SUBMIT_INFO`
  - `VK_STRUCTURE_TYPE_COPY_BUFFER_INFO_2`
  - `VK_STRUCTURE_TYPE_COPY_BUFFER_TO_IMAGE_INFO_2`
  - `VK_STRUCTURE_TYPE_COPY_IMAGE_INFO_2`
  - `VK_STRUCTURE_TYPE_COPY_IMAGE_TO_BUFFER_INFO_2`
  - `VK_STRUCTURE_TYPE_DEPENDENCY_INFO`
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_INLINE_UNIFORM_BLOCK_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_DEVICE_BUFFER_MEMORY_REQUIREMENTS`
  - `VK_STRUCTURE_TYPE_DEVICE_IMAGE_MEMORY_REQUIREMENTS`
  - `VK_STRUCTURE_TYPE_DEVICE_PRIVATE_DATA_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_3`
  - `VK_STRUCTURE_TYPE_IMAGE_BLIT_2`
  - `VK_STRUCTURE_TYPE_IMAGE_COPY_2`
  - `VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER_2`
  - `VK_STRUCTURE_TYPE_IMAGE_RESOLVE_2`
  - `VK_STRUCTURE_TYPE_MEMORY_BARRIER_2`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ROBUSTNESS_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_CREATION_CACHE_CONTROL_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRIVATE_DATA_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DEMOTE_TO_HELPER_INVOCATION_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_DOT_PRODUCT_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_DOT_PRODUCT_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_TERMINATE_INVOCATION_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SYNCHRONIZATION_2_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXEL_BUFFER_ALIGNMENT_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TEXTURE_COMPRESSION_ASTC_HDR_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TOOL_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_3_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_3_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ZERO_INITIALIZE_WORKGROUP_MEMORY_FEATURES`
  - `VK_STRUCTURE_TYPE_PIPELINE_CREATION_FEEDBACK_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_PIPELINE_RENDERING_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_REQUIRED_SUBGROUP_SIZE_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_PRIVATE_DATA_SLOT_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_INFO`
  - `VK_STRUCTURE_TYPE_RENDERING_INFO`
  - `VK_STRUCTURE_TYPE_RESOLVE_IMAGE_INFO_2`
  - `VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO`
  - `VK_STRUCTURE_TYPE_SUBMIT_INFO_2`
  - `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_INLINE_UNIFORM_BLOCK`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0