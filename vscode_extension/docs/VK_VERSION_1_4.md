# VK\_VERSION\_1\_4(3) Manual Page

## Name

VK\_VERSION\_1\_4 - Vulkan version 1.4



## [](#_description)Description

Vulkan Version 1.4 [promoted](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-promotion) a number of key extensions into the core API:

- [VK\_KHR\_dynamic\_rendering\_local\_read](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering_local_read.html)
- [VK\_KHR\_global\_priority](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_global_priority.html)
- [VK\_KHR\_index\_type\_uint8](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_index_type_uint8.html)
- [VK\_KHR\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_line_rasterization.html)
- [VK\_KHR\_load\_store\_op\_none](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_load_store_op_none.html)
- [VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html)
- [VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html)
- [VK\_KHR\_map\_memory2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_map_memory2.html)
- [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html)
- [VK\_KHR\_shader\_expect\_assume](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_expect_assume.html)
- [VK\_KHR\_shader\_float\_controls2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_float_controls2.html)
- [VK\_KHR\_shader\_subgroup\_rotate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_subgroup_rotate.html)
- [VK\_KHR\_vertex\_attribute\_divisor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_vertex_attribute_divisor.html)
- [VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html)
- [VK\_EXT\_pipeline\_protected\_access](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_protected_access.html)
- [VK\_EXT\_pipeline\_robustness](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_robustness.html)

All differences in behavior between these extensions and the corresponding Vulkan 1.4 functionality are summarized below.

Differences Relative to `VK_KHR_dynamic_rendering_local_read`

If the [VK\_KHR\_dynamic\_rendering\_local\_read](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering_local_read.html) extension is not supported, Vulkan 1.4 implementations **must** support local read only for storage resources and single sampled color attachments.

Support for reading depth/stencil attachments and multi-sampled attachments are respectively gated behind the new boolean `dynamicRenderingLocalReadDepthStencilAttachments` and `dynamicRenderingLocalReadMultisampledAttachments` properties.

- If `dynamicRenderingLocalReadDepthStencilAttachments` is `VK_FALSE`, implementations do not support depth/stencil attachment access within dynamic rendering.
- If `dynamicRenderingLocalReadMultisampledAttachments` is `VK_FALSE`, implementations do not support multisampled attachment access within dynamic rendering.
- If both properties are `VK_TRUE`, the full functionality of the extension is supported.

Differences Relative to `VK_EXT_host_image_copy`

If the [VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html) extension is not supported, support for it is optional in Vulkan 1.4.

- An implementation that has a `VK_QUEUE_GRAPHICS_BIT` queue must support either:
  
  - the [`hostImageCopy`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-hostImageCopy) feature; or
  - an additional queue that supports `VK_QUEUE_TRANSFER_BIT`.

Differences Relative to `VK_KHR_push_descriptor`

[VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html) did not include a feature bit, so a new feature bit has been added to [VkPhysicalDeviceVulkan14Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan14Features.html) to gate its functionality: [`pushDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pushDescriptor). Enabling this new feature has the same effect as enabling the extension.

Differences Relative to `VK_EXT_pipeline_protected_access`

[VK\_EXT\_pipeline\_protected\_access](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_protected_access.html) is only useful when the [`protectedMemory`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-protectedMemory) feature is supported. As the [`protectedMemory`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-protectedMemory) feature is optional in core Vulkan, the [`pipelineProtectedAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineProtectedAccess) feature is only required when the [`protectedMemory`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-protectedMemory) feature is supported.

Differences Relative to `VK_KHR_line_rasterization`

The [`bresenhamLines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bresenhamLines) feature is required, rather than just any one of the line style features.

Differences Relative to `VK_KHR_shader_subgroup_rotate`

The [`shaderSubgroupRotateClustered`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderSubgroupRotateClustered) feature is required in addition to [`shaderSubgroupRotate`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderSubgroupRotate).

Additional Vulkan 1.4 Feature Support

In addition to the promoted extensions described above, Vulkan 1.4 added required support for:

- All queues supporting `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT` **must** also advertise `VK_QUEUE_TRANSFER_BIT`.
- Clustered subgroup operations **must** be advertised in Vulkan 1.4 via setting both `VK_SUBGROUP_FEATURE_CLUSTERED_BIT` and `VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT` (as an interaction with the promoted [VK\_KHR\_shader\_subgroup\_rotate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_subgroup_rotate.html) functionality) in [`supportedOperations`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-subgroupSupportedOperations).
- The following features that were optional in earlier versions:
  
  - [`fullDrawIndexUint32`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-fullDrawIndexUint32)
  - [`imageCubeArray`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-imageCubeArray)
  - [`independentBlend`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-independentBlend)
  - [`sampleRateShading`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-sampleRateShading)
  - [`drawIndirectFirstInstance`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-drawIndirectFirstInstance)
  - [`depthClamp`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-depthClamp)
  - [`depthBiasClamp`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-depthBiasClamp)
  - [`samplerAnisotropy`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-samplerAnisotropy)
  - [`fragmentStoresAndAtomics`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-fragmentStoresAndAtomics)
  - [`shaderStorageImageExtendedFormats`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderStorageImageExtendedFormats)
  - [`shaderUniformBufferArrayDynamicIndexing`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderUniformBufferArrayDynamicIndexing)
  - [`shaderSampledImageArrayDynamicIndexing`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderSampledImageArrayDynamicIndexing)
  - [`shaderStorageBufferArrayDynamicIndexing`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderStorageBufferArrayDynamicIndexing)
  - [`shaderStorageImageArrayDynamicIndexing`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderStorageImageArrayDynamicIndexing)
  - [`shaderImageGatherExtended`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderImageGatherExtended)
  - [`shaderInt16`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderInt16)
  - [`largePoints`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-largePoints)
  - [`samplerYcbcrConversion`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-samplerYcbcrConversion)
  - [`storageBuffer16BitAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-storageBuffer16BitAccess)
  - [`variablePointers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-variablePointers)
  - [`variablePointersStorageBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-variablePointersStorageBuffer)
  - [`samplerMirrorClampToEdge`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-samplerMirrorClampToEdge)
  - [`scalarBlockLayout`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-scalarBlockLayout)
  - [`shaderUniformTexelBufferArrayDynamicIndexing`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderUniformTexelBufferArrayDynamicIndexing)
  - [`shaderStorageTexelBufferArrayDynamicIndexing`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderStorageTexelBufferArrayDynamicIndexing)
  - [`shaderInt8`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderInt8)
  - [`storageBuffer8BitAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-storageBuffer8BitAccess)

Updated Vulkan 1.4 Limit Support

Vulkan 1.4 also requires support for the following updated limits:

- [`maxImageDimension1D`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxImageDimension1D) is increased from 4096 to 8192
- [`maxImageDimension2D`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxImageDimension2D) is increased from 4096 to 8192
- [`maxImageDimension3D`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxImageDimension3D) is increased from 256 to 512
- [`maxImageDimensionCube`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxImageDimensionCube) is increased from 4096 to 8192
- [`maxImageArrayLayers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxImageArrayLayers) is increased from 256 to 2048
- [`maxUniformBufferRange`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxUniformBufferRange) is increased from 16384 to 65536
- [`maxPushConstantsSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxPushConstantsSize) is increased from 128 to 256
- [`bufferImageGranularity`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-bufferImageGranularity) is decreased from 131072 to 4096
- [`maxBoundDescriptorSets`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxBoundDescriptorSets) is increased from 4 to 7
- [`maxPerStageDescriptorUniformBuffers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxPerStageDescriptorUniformBuffers) is increased from 12 to 15
- [`maxPerStageResources`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxPerStageResources) is increased from 128 to 200
- [`maxDescriptorSetUniformBuffers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxDescriptorSetUniformBuffers) is increased from 72 to 90
- [`maxDescriptorSetStorageBuffers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxDescriptorSetStorageBuffers) is increased from 24 to 96
- [`maxDescriptorSetStorageImages`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxDescriptorSetStorageImages) is increased from 24 to 144
- [`maxFragmentCombinedOutputResources`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentCombinedOutputResources) is increased from 4 to 16
- [`maxComputeWorkGroupInvocations`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxComputeWorkGroupInvocations) is increased from 128 to 256
- [`maxComputeWorkGroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxComputeWorkGroupSize) is increased from (128,128,64) to (256,256,64)
- [`shaderSignedZeroInfNanPreserveFloat16`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-shaderSignedZeroInfNanPreserveFloat16) is changed from unspecified to `VK_TRUE`
- [`shaderSignedZeroInfNanPreserveFloat32`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-shaderSignedZeroInfNanPreserveFloat32) is changed from unspecified to `VK_TRUE`
- [`subTexelPrecisionBits`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-subTexelPrecisionBits) is increased from 4 to 8
- [`mipmapPrecisionBits`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-mipmapPrecisionBits) is increased from 4 to 6
- [`maxSamplerLodBias`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxSamplerLodBias) is increased from 2 to 14
- [`maxViewportDimensions`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxViewportDimensions) is increased from (4096,4096) to (7680,7680)
- [`viewportBoundsRange`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-viewportboundsrange) is increased from (-8192,8191) to (-15360,15359)
- [`maxFramebufferWidth`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFramebufferWidth) is increased from 4096 to 7680
- [`maxFramebufferHeight`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFramebufferHeight) is increased from 4096 to 7680
- [`maxColorAttachments`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxColorAttachments) is increased from 7 to 8
- [`timestampComputeAndGraphics`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-timestampComputeAndGraphics) is changed from unspecified to `VK_TRUE`
- [`pointSizeRange`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-pointSizeRange) is increased from (1.0,64.0 - ULP) to (1.0,256.0 - `pointSizeGranularity`)
- [`pointSizeGranularity`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-pointSizeGranularity) is decreased from 1.0 to 0.125
- [`lineWidthGranularity`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-lineWidthGranularity) is decreased from 1.0 to 0.5
- [`maxPushDescriptors`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxPushDescriptors) is increased from 16 to 32
- [`standardSampleLocations`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-standardSampleLocations) is changed from unspecified to `VK_TRUE`

### [](#_new_macros)New Macros

- [VK\_API\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_API_VERSION_1_4.html)

### [](#_new_commands)New Commands

- [vkCmdBindDescriptorSets2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets2.html)
- [vkCmdBindIndexBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer2.html)
- [vkCmdPushConstants2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushConstants2.html)
- [vkCmdPushDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSet.html)
- [vkCmdPushDescriptorSet2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSet2.html)
- [vkCmdPushDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSetWithTemplate.html)
- [vkCmdPushDescriptorSetWithTemplate2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSetWithTemplate2.html)
- [vkCmdSetLineStipple](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetLineStipple.html)
- [vkCmdSetRenderingAttachmentLocations](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRenderingAttachmentLocations.html)
- [vkCmdSetRenderingInputAttachmentIndices](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRenderingInputAttachmentIndices.html)
- [vkCopyImageToImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyImageToImage.html)
- [vkCopyImageToMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyImageToMemory.html)
- [vkCopyMemoryToImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToImage.html)
- [vkGetDeviceImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSubresourceLayout.html)
- [vkGetImageSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2.html)
- [vkGetRenderingAreaGranularity](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRenderingAreaGranularity.html)
- [vkMapMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory2.html)
- [vkTransitionImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkTransitionImageLayout.html)
- [vkUnmapMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUnmapMemory2.html)

### [](#_new_structures)New Structures

- [VkBindDescriptorSetsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDescriptorSetsInfo.html)
- [VkCopyImageToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToImageInfo.html)
- [VkCopyImageToMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToMemoryInfo.html)
- [VkCopyMemoryToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageInfo.html)
- [VkDeviceImageSubresourceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageSubresourceInfo.html)
- [VkHostImageLayoutTransitionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageLayoutTransitionInfo.html)
- [VkImageSubresource2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource2.html)
- [VkImageToMemoryCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageToMemoryCopy.html)
- [VkMemoryMapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapInfo.html)
- [VkMemoryToImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryToImageCopy.html)
- [VkMemoryUnmapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapInfo.html)
- [VkPushConstantsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantsInfo.html)
- [VkPushDescriptorSetInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushDescriptorSetInfo.html)
- [VkPushDescriptorSetWithTemplateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushDescriptorSetWithTemplateInfo.html)
- [VkRenderingAreaInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAreaInfo.html)
- [VkSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2.html)
- [VkVertexInputBindingDivisorDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDivisorDescription.html)
- Extending [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryInfo.html), [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html):
  
  - [VkBindMemoryStatus](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindMemoryStatus.html)
- Extending [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferViewCreateInfo.html), [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html), [VkPhysicalDeviceExternalBufferInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalBufferInfo.html), [VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferBindingInfoEXT.html):
  
  - [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html)
- Extending [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html):
  
  - [VkPipelineCreateFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfo.html)
- Extending [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html):
  
  - [VkDeviceQueueGlobalPriorityCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueGlobalPriorityCreateInfo.html)
- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html):
  
  - [VkRenderingAttachmentLocationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentLocationInfo.html)
  - [VkRenderingInputAttachmentIndexInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInputAttachmentIndexInfo.html)
- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html):
  
  - [VkPipelineRobustnessCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessCreateInfo.html)
- Extending [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html):
  
  - [VkHostImageCopyDevicePerformanceQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageCopyDevicePerformanceQuery.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDynamicRenderingLocalReadFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDynamicRenderingLocalReadFeatures.html)
  - [VkPhysicalDeviceGlobalPriorityQueryFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceGlobalPriorityQueryFeatures.html)
  - [VkPhysicalDeviceHostImageCopyFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceHostImageCopyFeatures.html)
  - [VkPhysicalDeviceIndexTypeUint8Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceIndexTypeUint8Features.html)
  - [VkPhysicalDeviceLineRasterizationFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLineRasterizationFeatures.html)
  - [VkPhysicalDeviceMaintenance5Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance5Features.html)
  - [VkPhysicalDeviceMaintenance6Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance6Features.html)
  - [VkPhysicalDevicePipelineProtectedAccessFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineProtectedAccessFeatures.html)
  - [VkPhysicalDevicePipelineRobustnessFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineRobustnessFeatures.html)
  - [VkPhysicalDeviceShaderExpectAssumeFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderExpectAssumeFeatures.html)
  - [VkPhysicalDeviceShaderFloatControls2Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderFloatControls2Features.html)
  - [VkPhysicalDeviceShaderSubgroupRotateFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderSubgroupRotateFeatures.html)
  - [VkPhysicalDeviceVertexAttributeDivisorFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVertexAttributeDivisorFeatures.html)
  - [VkPhysicalDeviceVulkan14Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan14Features.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceHostImageCopyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceHostImageCopyProperties.html)
  - [VkPhysicalDeviceLineRasterizationProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLineRasterizationProperties.html)
  - [VkPhysicalDeviceMaintenance5Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance5Properties.html)
  - [VkPhysicalDeviceMaintenance6Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance6Properties.html)
  - [VkPhysicalDevicePipelineRobustnessProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineRobustnessProperties.html)
  - [VkPhysicalDevicePushDescriptorProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePushDescriptorProperties.html)
  - [VkPhysicalDeviceVertexAttributeDivisorProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVertexAttributeDivisorProperties.html)
  - [VkPhysicalDeviceVulkan14Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan14Properties.html)
- Extending [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html):
  
  - [VkPipelineRasterizationLineStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationLineStateCreateInfo.html)
- Extending [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputStateCreateInfo.html):
  
  - [VkPipelineVertexInputDivisorStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputDivisorStateCreateInfo.html)
- Extending [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2.html):
  
  - [VkQueueFamilyGlobalPriorityProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyGlobalPriorityProperties.html)
- Extending [VkSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2.html):
  
  - [VkSubresourceHostMemcpySize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceHostMemcpySize.html)

### [](#_new_enums)New Enums

- [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html)
- [VkHostImageCopyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageCopyFlagBits.html)
- [VkLineRasterizationMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLineRasterizationMode.html)
- [VkMemoryUnmapFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapFlagBits.html)
- [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html)
- [VkPipelineRobustnessBufferBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessBufferBehavior.html)
- [VkPipelineRobustnessImageBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessImageBehavior.html)
- [VkQueueGlobalPriority](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueGlobalPriority.html)

### [](#_new_bitmasks)New Bitmasks

- [VkBufferUsageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2.html)
- [VkHostImageCopyFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageCopyFlags.html)
- [VkMemoryUnmapFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapFlags.html)
- [VkPipelineCreateFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2.html)

### [](#_new_enum_constants)New Enum Constants

- `VK_MAX_GLOBAL_PRIORITY_SIZE`
- Extending [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html):
  
  - `VK_ATTACHMENT_LOAD_OP_NONE`
- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_SHADER_DEVICE_ADDRESS_BIT`
- Extending [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateFlagBits.html):
  
  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT`
- Extending [VkDescriptorUpdateTemplateType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateType.html):
  
  - `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS`
- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_LINE_STIPPLE`
- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html):
  
  - `VK_FORMAT_A1B5G5R5_UNORM_PACK16`
  - `VK_FORMAT_A8_UNORM`
- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_HOST_IMAGE_TRANSFER_BIT`
- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html):
  
  - `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ`
- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html):
  
  - `VK_IMAGE_USAGE_HOST_TRANSFER_BIT`
- Extending [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html):
  
  - `VK_INDEX_TYPE_UINT8`
- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html):
  
  - `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT`
  - `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_NOT_PERMITTED`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BIND_DESCRIPTOR_SETS_INFO`
  - `VK_STRUCTURE_TYPE_BIND_MEMORY_STATUS`
  - `VK_STRUCTURE_TYPE_BUFFER_USAGE_FLAGS_2_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_COPY_IMAGE_TO_IMAGE_INFO`
  - `VK_STRUCTURE_TYPE_COPY_IMAGE_TO_MEMORY_INFO`
  - `VK_STRUCTURE_TYPE_COPY_MEMORY_TO_IMAGE_INFO`
  - `VK_STRUCTURE_TYPE_DEVICE_IMAGE_SUBRESOURCE_INFO`
  - `VK_STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_HOST_IMAGE_COPY_DEVICE_PERFORMANCE_QUERY`
  - `VK_STRUCTURE_TYPE_HOST_IMAGE_LAYOUT_TRANSITION_INFO`
  - `VK_STRUCTURE_TYPE_IMAGE_SUBRESOURCE_2`
  - `VK_STRUCTURE_TYPE_IMAGE_TO_MEMORY_COPY`
  - `VK_STRUCTURE_TYPE_MEMORY_MAP_INFO`
  - `VK_STRUCTURE_TYPE_MEMORY_TO_IMAGE_COPY`
  - `VK_STRUCTURE_TYPE_MEMORY_UNMAP_INFO`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_LOCAL_READ_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GLOBAL_PRIORITY_QUERY_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_IMAGE_COPY_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_IMAGE_COPY_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INDEX_TYPE_UINT8_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_5_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_5_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_6_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_6_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_PROTECTED_ACCESS_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_ROBUSTNESS_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_ROBUSTNESS_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PUSH_DESCRIPTOR_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_EXPECT_ASSUME_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT_CONTROLS_2_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_ROTATE_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_4_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_4_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PIPELINE_CREATE_FLAGS_2_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_LINE_STATE_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_PIPELINE_ROBUSTNESS_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_PUSH_CONSTANTS_INFO`
  - `VK_STRUCTURE_TYPE_PUSH_DESCRIPTOR_SET_INFO`
  - `VK_STRUCTURE_TYPE_PUSH_DESCRIPTOR_SET_WITH_TEMPLATE_INFO`
  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_GLOBAL_PRIORITY_PROPERTIES`
  - `VK_STRUCTURE_TYPE_RENDERING_AREA_INFO`
  - `VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_LOCATION_INFO`
  - `VK_STRUCTURE_TYPE_RENDERING_INPUT_ATTACHMENT_INDEX_INFO`
  - `VK_STRUCTURE_TYPE_SUBRESOURCE_HOST_MEMCPY_SIZE`
  - `VK_STRUCTURE_TYPE_SUBRESOURCE_LAYOUT_2`
- Extending [VkSubgroupFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubgroupFeatureFlagBits.html):
  
  - `VK_SUBGROUP_FEATURE_ROTATE_BIT`
  - `VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0