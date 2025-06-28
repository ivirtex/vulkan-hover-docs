# Required\_Limits(3) Manual Page

## Name

Required\_Limits - Vulkan required limit tables



## [](#_description)Description

The following table specifies the **required** minimum/maximum for all Vulkan graphics implementations. Where a limit corresponds to a fine-grained device feature which is **optional**, the feature name is listed with two **required** limits, one when the feature is supported and one when it is not supported. If an implementation supports a feature, the limits reported are the same whether or not the feature is enabled.

Table 1. Required Limit Types    Type Limit Feature

`uint32_t`

`maxImageDimension1D`

\-

`uint32_t`

`maxImageDimension2D`

\-

`uint32_t`

`maxImageDimension3D`

\-

`uint32_t`

`maxImageDimensionCube`

\-

`uint32_t`

`maxImageArrayLayers`

\-

`uint32_t`

`maxTexelBufferElements`

\-

`uint32_t`

`maxUniformBufferRange`

\-

`uint32_t`

`maxStorageBufferRange`

\-

`uint32_t`

`maxPushConstantsSize`

\-

`uint32_t`

`maxMemoryAllocationCount`

\-

`uint32_t`

`maxSamplerAllocationCount`

\-

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`bufferImageGranularity`

\-

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`sparseAddressSpaceSize`

`sparseBinding`

`uint32_t`

`maxBoundDescriptorSets`

\-

`uint32_t`

`maxPerStageDescriptorSamplers`

\-

`uint32_t`

`maxPerStageDescriptorUniformBuffers`

\-

`uint32_t`

`maxPerStageDescriptorStorageBuffers`

\-

`uint32_t`

`maxPerStageDescriptorSampledImages`

\-

`uint32_t`

`maxPerStageDescriptorStorageImages`

\-

`uint32_t`

`maxPerStageDescriptorInputAttachments`

\-

`uint32_t`

`maxPerStageResources`

\-

`uint32_t`

`maxDescriptorSetSamplers`

\-

`uint32_t`

`maxDescriptorSetUniformBuffers`

\-

`uint32_t`

`maxDescriptorSetUniformBuffersDynamic`

\-

`uint32_t`

`maxDescriptorSetStorageBuffers`

\-

`uint32_t`

`maxDescriptorSetStorageBuffersDynamic`

\-

`uint32_t`

`maxDescriptorSetSampledImages`

\-

`uint32_t`

`maxDescriptorSetStorageImages`

\-

`uint32_t`

`maxDescriptorSetInputAttachments`

\-

`uint32_t`

`maxVertexInputAttributes`

\-

`uint32_t`

`maxVertexInputBindings`

\-

`uint32_t`

`maxVertexInputAttributeOffset`

\-

`uint32_t`

`maxVertexInputBindingStride`

\-

`uint32_t`

`maxVertexOutputComponents`

\-

`uint32_t`

`maxTessellationGenerationLevel`

`tessellationShader`

`uint32_t`

`maxTessellationPatchSize`

`tessellationShader`

`uint32_t`

`maxTessellationControlPerVertexInputComponents`

`tessellationShader`

`uint32_t`

`maxTessellationControlPerVertexOutputComponents`

`tessellationShader`

`uint32_t`

`maxTessellationControlPerPatchOutputComponents`

`tessellationShader`

`uint32_t`

`maxTessellationControlTotalOutputComponents`

`tessellationShader`

`uint32_t`

`maxTessellationEvaluationInputComponents`

`tessellationShader`

`uint32_t`

`maxTessellationEvaluationOutputComponents`

`tessellationShader`

`uint32_t`

`maxGeometryShaderInvocations`

`geometryShader`

`uint32_t`

`maxGeometryInputComponents`

`geometryShader`

`uint32_t`

`maxGeometryOutputComponents`

`geometryShader`

`uint32_t`

`maxGeometryOutputVertices`

`geometryShader`

`uint32_t`

`maxGeometryTotalOutputComponents`

`geometryShader`

`uint32_t`

`maxFragmentInputComponents`

\-

`uint32_t`

`maxFragmentOutputAttachments`

\-

`uint32_t`

`maxFragmentDualSrcAttachments`

`dualSrcBlend`

`uint32_t`

`maxFragmentCombinedOutputResources`

\-

`uint32_t`

`maxComputeSharedMemorySize`

\-

3 × `uint32_t`

`maxComputeWorkGroupCount`

\-

`uint32_t`

`maxComputeWorkGroupInvocations`

\-

3 × `uint32_t`

`maxComputeWorkGroupSize`

\-

`uint32_t`

`subPixelPrecisionBits`

\-

`uint32_t`

`subTexelPrecisionBits`

\-

`uint32_t`

`mipmapPrecisionBits`

\-

`uint32_t`

`maxDrawIndexedIndexValue`

`fullDrawIndexUint32`

`uint32_t`

`maxDrawIndirectCount`

`multiDrawIndirect`

`float`

`maxSamplerLodBias`

\-

`float`

`maxSamplerAnisotropy`

`samplerAnisotropy`

`uint32_t`

`maxViewports`

`multiViewport`

2 × `uint32_t`

`maxViewportDimensions`

\-

2 × `float`

`viewportBoundsRange`

\-

`uint32_t`

`viewportSubPixelBits`

\-

`size_t`

`minMemoryMapAlignment`

\-

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`minTexelBufferOffsetAlignment`

\-

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`minUniformBufferOffsetAlignment`

\-

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`minStorageBufferOffsetAlignment`

\-

`int32_t`

`minTexelOffset`

\-

`uint32_t`

`maxTexelOffset`

\-

`int32_t`

`minTexelGatherOffset`

`shaderImageGatherExtended`

`uint32_t`

`maxTexelGatherOffset`

`shaderImageGatherExtended`

`float`

`minInterpolationOffset`

`sampleRateShading`

`float`

`maxInterpolationOffset`

`sampleRateShading`

`uint32_t`

`subPixelInterpolationOffsetBits`

`sampleRateShading`

`uint32_t`

`maxFramebufferWidth`

\-

`uint32_t`

`maxFramebufferHeight`

\-

`uint32_t`

`maxFramebufferLayers`

\-

[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html)

`framebufferColorSampleCounts`

\-

[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html)

`framebufferIntegerColorSampleCounts`

\-

[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html)

`framebufferDepthSampleCounts`

\-

[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html)

`framebufferStencilSampleCounts`

\-

[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html)

`framebufferNoAttachmentsSampleCounts`

\-

`uint32_t`

`maxColorAttachments`

\-

[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html)

`sampledImageColorSampleCounts`

\-

[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html)

`sampledImageIntegerSampleCounts`

\-

[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html)

`sampledImageDepthSampleCounts`

\-

[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html)

`sampledImageStencilSampleCounts`

\-

[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html)

`storageImageSampleCounts`

`shaderStorageImageMultisample`

`uint32_t`

`maxSampleMaskWords`

\-

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`timestampComputeAndGraphics`

\-

`float`

`timestampPeriod`

\-

`uint32_t`

`maxClipDistances`

`shaderClipDistance`

`uint32_t`

`maxCullDistances`

`shaderCullDistance`

`uint32_t`

`maxCombinedClipAndCullDistances`

`shaderCullDistance`

`uint32_t`

`discreteQueuePriorities`

\-

2 × `float`

`pointSizeRange`

`largePoints`

2 × `float`

`lineWidthRange`

`wideLines`

`float`

`pointSizeGranularity`

`largePoints`

`float`

`lineWidthGranularity`

`wideLines`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`strictLines`

\-

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`standardSampleLocations`

\-

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`optimalBufferCopyOffsetAlignment`

\-

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`optimalBufferCopyRowPitchAlignment`

\-

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`nonCoherentAtomSize`

\-

`uint32_t`

`maxDiscardRectangles`

`VK_EXT_discard_rectangles`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`filterMinmaxSingleComponentFormats`

`samplerFilterMinmax` `VK_EXT_sampler_filter_minmax`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`filterMinmaxImageComponentMapping`

`samplerFilterMinmax` `VK_EXT_sampler_filter_minmax`

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`maxBufferSize`

`maintenance4`

`float`

`primitiveOverestimationSize`

`VK_EXT_conservative_rasterization`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`maxExtraPrimitiveOverestimationSize`

`VK_EXT_conservative_rasterization`

`float`

`extraPrimitiveOverestimationSizeGranularity`

`VK_EXT_conservative_rasterization`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`degenerateTriangleRasterized`

`VK_EXT_conservative_rasterization`

`float`

`degenerateLinesRasterized`

`VK_EXT_conservative_rasterization`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`fullyCoveredFragmentShaderInputVariable`

`VK_EXT_conservative_rasterization`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`conservativeRasterizationPostDepthCoverage`

`VK_EXT_conservative_rasterization`

`uint32_t`

`maxUpdateAfterBindDescriptorsInAllPools`

`descriptorIndexing`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`shaderUniformBufferArrayNonUniformIndexingNative`

\-

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`shaderSampledImageArrayNonUniformIndexingNative`

\-

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`shaderStorageBufferArrayNonUniformIndexingNative`

\-

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`shaderStorageImageArrayNonUniformIndexingNative`

\-

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`shaderInputAttachmentArrayNonUniformIndexingNative`

\-

`uint32_t`

`maxPerStageDescriptorUpdateAfterBindSamplers`

`descriptorIndexing`

`uint32_t`

`maxPerStageDescriptorUpdateAfterBindUniformBuffers`

`descriptorIndexing`

`uint32_t`

`maxPerStageDescriptorUpdateAfterBindStorageBuffers`

`descriptorIndexing`

`uint32_t`

`maxPerStageDescriptorUpdateAfterBindSampledImages`

`descriptorIndexing`

`uint32_t`

`maxPerStageDescriptorUpdateAfterBindStorageImages`

`descriptorIndexing`

`uint32_t`

`maxPerStageDescriptorUpdateAfterBindInputAttachments`

`descriptorIndexing`

`uint32_t`

`maxPerStageUpdateAfterBindResources`

`descriptorIndexing`

`uint32_t`

`maxDescriptorSetUpdateAfterBindSamplers`

`descriptorIndexing`

`uint32_t`

`maxDescriptorSetUpdateAfterBindUniformBuffers`

`descriptorIndexing`

`uint32_t`

`maxDescriptorSetUpdateAfterBindUniformBuffersDynamic`

`descriptorIndexing`

`uint32_t`

`maxDescriptorSetUpdateAfterBindStorageBuffers`

`descriptorIndexing`

`uint32_t`

`maxDescriptorSetUpdateAfterBindStorageBuffersDynamic`

`descriptorIndexing`

`uint32_t`

`maxDescriptorSetUpdateAfterBindSampledImages`

`descriptorIndexing`

`uint32_t`

`maxDescriptorSetUpdateAfterBindStorageImages`

`descriptorIndexing`

`uint32_t`

`maxDescriptorSetUpdateAfterBindInputAttachments`

`descriptorIndexing`

`uint32_t`

`maxInlineUniformBlockSize`

`inlineUniformBlock`

`uint32_t`

`maxPerStageDescriptorInlineUniformBlocks`

`inlineUniformBlock`

`uint32_t`

`maxPerStageDescriptorUpdateAfterBindInlineUniformBlocks`

`inlineUniformBlock`

`uint32_t`

`maxDescriptorSetInlineUniformBlocks`

`inlineUniformBlock`

`uint32_t`

`maxDescriptorSetUpdateAfterBindInlineUniformBlocks`

`inlineUniformBlock`

`uint32_t`

`maxInlineUniformTotalSize`

`inlineUniformBlock`

`uint32_t`

`maxVertexAttribDivisor`

Vulkan 1.4, [VK\_KHR\_vertex\_attribute\_divisor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_vertex_attribute_divisor.html), [VK\_EXT\_vertex\_attribute\_divisor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_vertex_attribute_divisor.html)

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxDrawMeshTasksCount`

`VK_NV_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxTaskWorkGroupInvocations`

`VK_NV_mesh_shader`

3 × `uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxTaskWorkGroupSize`

`VK_NV_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxTaskTotalMemorySize`

`VK_NV_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxTaskOutputCount`

`VK_NV_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshWorkGroupInvocations`

`VK_NV_mesh_shader`

3 × `uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshWorkGroupSize`

`VK_NV_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshTotalMemorySize`

`VK_NV_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshOutputVertices`

`VK_NV_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshOutputPrimitives`

`VK_NV_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshMultiviewViewCount`

`VK_NV_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`meshOutputPerVertexGranularity`

`VK_NV_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`meshOutputPerPrimitiveGranularity`

`VK_NV_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskWorkGroupTotalCount`

`VK_EXT_mesh_shader`

3 × `uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskWorkGroupCount`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskWorkGroupInvocations`

`VK_EXT_mesh_shader`

3 × `uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskWorkGroupSize`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskPayloadSize`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskSharedMemorySize`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskPayloadAndSharedMemorySize`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupTotalCount`

`VK_EXT_mesh_shader`

3 × `uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupCount`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupInvocations`

`VK_EXT_mesh_shader`

3 × `uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupSize`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshSharedMemorySize`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshPayloadAndSharedMemorySize`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshOutputMemorySize`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshPayloadAndOutputMemorySize`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshOutputComponents`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshOutputVertices`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshOutputPrimitives`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshOutputLayers`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshMultiviewViewCount`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`meshOutputPerVertexGranularity`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`meshOutputPerPrimitiveGranularity`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxPreferredTaskWorkGroupInvocations`

`VK_EXT_mesh_shader`

`uint32_t`

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxPreferredMeshWorkGroupInvocations`

`VK_EXT_mesh_shader`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`prefersLocalInvocationVertexOutput`

`VK_EXT_mesh_shader`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`prefersLocalInvocationPrimitiveOutput`

`VK_EXT_mesh_shader`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`prefersCompactVertexOutput`

`VK_EXT_mesh_shader`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`prefersCompactPrimitiveOutput`

`VK_EXT_mesh_shader`

`uint32_t`

`maxTransformFeedbackStreams`

`VK_EXT_transform_feedback`

`uint32_t`

`maxTransformFeedbackBuffers`

`VK_EXT_transform_feedback`

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`maxTransformFeedbackBufferSize`

`VK_EXT_transform_feedback`

`uint32_t`

`maxTransformFeedbackStreamDataSize`

`VK_EXT_transform_feedback`

`uint32_t`

`maxTransformFeedbackBufferDataSize`

`VK_EXT_transform_feedback`

`uint32_t`

`maxTransformFeedbackBufferDataStride`

`VK_EXT_transform_feedback`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`transformFeedbackQueries`

`VK_EXT_transform_feedback`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`transformFeedbackStreamsLinesTriangles`

`VK_EXT_transform_feedback`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`transformFeedbackRasterizationStreamSelect`

`VK_EXT_transform_feedback`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`transformFeedbackDraw`

`VK_EXT_transform_feedback`

[VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html)

`minFragmentDensityTexelSize`

`fragmentDensityMap`

[VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html)

`maxFragmentDensityTexelSize`

`fragmentDensityMap`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`fragmentDensityInvocations`

`fragmentDensityMap`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`subsampledLoads`

`VK_EXT_fragment_density_map2`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`subsampledCoarseReconstructionEarlyAccess`

`VK_EXT_fragment_density_map2`

`uint32_t`

`maxSubsampledArrayLayers`

`VK_EXT_fragment_density_map2`

`uint32_t`

`maxDescriptorSetSubsampledSamplers`

`VK_EXT_fragment_density_map2`

[VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html)

`fragmentDensityOffsetGranularity`

`fragmentDensityMapOffset`

`uint32_t`

`maxFragmentDensityMapLayers`

`fragmentDensityMapLayered`

`uint32_t`

`maxGeometryCount`

`VK_NV_ray_tracing`, `VK_KHR_acceleration_structure`

`uint32_t`

`maxInstanceCount`

`VK_NV_ray_tracing`, `VK_KHR_acceleration_structure`

`uint32_t`

`maxVerticesPerCluster`

`clusterAccelerationStructure`

`uint32_t`

`maxTrianglesPerCluster`

`clusterAccelerationStructure`

`uint32_t`

`clusterScratchByteAlignment`

`clusterAccelerationStructure`

`uint32_t`

`clusterByteAlignment`

`clusterAccelerationStructure`

`uint32_t`

`clusterTemplateByteAlignment`

`clusterAccelerationStructure`

`uint32_t`

`clusterBottomLevelByteAlignment`

`clusterAccelerationStructure`

`uint32_t`

`clusterTemplateBoundsByteAlignment`

`clusterAccelerationStructure`

`uint32_t`

`maxClusterGeometryIndex`

`clusterAccelerationStructure`

`uint32_t`

`shaderGroupHandleSize`

`VK_NV_ray_tracing`, `VK_KHR_ray_tracing_pipeline`

`uint32_t`

`maxShaderGroupStride`

`VK_NV_ray_tracing`, `VK_KHR_ray_tracing_pipeline`

`uint32_t`

`shaderGroupBaseAlignment`

`VK_NV_ray_tracing`, `VK_KHR_ray_tracing_pipeline`

`uint32_t`

`maxRecursionDepth`

`VK_NV_ray_tracing`

`uint32_t`

`maxTriangleCount`

`VK_NV_ray_tracing`

`uint32_t`

`maxPerStageDescriptorAccelerationStructures`

`VK_KHR_acceleration_structure`

`uint32_t`

`maxPerStageDescriptorUpdateAfterBindAccelerationStructures`

`VK_KHR_acceleration_structure`

`uint32_t`

`maxDescriptorSetAccelerationStructures`

`VK_NV_ray_tracing`, `VK_KHR_acceleration_structure`

`uint32_t`

`maxDescriptorSetUpdateAfterBindAccelerationStructures`

`VK_KHR_acceleration_structure`

`uint32_t`

`minAccelerationStructureScratchOffsetAlignment`

`VK_KHR_acceleration_structure`

`uint32_t`

`maxRayRecursionDepth`

`VK_KHR_ray_tracing_pipeline`

`uint32_t`

`shaderGroupHandleCaptureReplaySize`

`VK_KHR_ray_tracing_pipeline`

`uint32_t`

`maxRayDispatchInvocationCount`

`VK_KHR_ray_tracing_pipeline`

`uint32_t`

`shaderGroupHandleAlignment`

`VK_KHR_ray_tracing_pipeline`

`uint32_t`

`maxRayHitAttributeSize`

`VK_KHR_ray_tracing_pipeline`

`uint32_t`

`maxPartitionCount`

`partitionedAccelerationStructure`

`uint64_t`

`maxTimelineSemaphoreValueDifference`

`timelineSemaphore`

`uint32_t`

`lineSubPixelPrecisionBits`

Vulkan 1.4, [VK\_KHR\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_line_rasterization.html), [VK\_EXT\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_line_rasterization.html)

`uint32_t`

`maxCustomBorderColorSamplers`

`VK_EXT_custom_border_color`

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`robustStorageBufferAccessSizeAlignment`

`VK_EXT_robustness2`, `VK_KHR_robustness2`

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`robustUniformBufferAccessSizeAlignment`

`VK_EXT_robustness2`, `VK_KHR_robustness2`

2 × `uint32_t`

`minFragmentShadingRateAttachmentTexelSize`

`attachmentFragmentShadingRate`

2 × `uint32_t`

`maxFragmentShadingRateAttachmentTexelSize`

`attachmentFragmentShadingRate`

`uint32_t`

`maxFragmentShadingRateAttachmentTexelSizeAspectRatio`

`attachmentFragmentShadingRate`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`primitiveFragmentShadingRateWithMultipleViewports`

`primitiveFragmentShadingRate`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`layeredShadingRateAttachments`

`attachmentFragmentShadingRate`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`fragmentShadingRateNonTrivialCombinerOps`

`pipelineFragmentShadingRate`

2 × `uint32_t`

`maxFragmentSize`

`pipelineFragmentShadingRate`

`uint32_t`

`maxFragmentSizeAspectRatio`

`pipelineFragmentShadingRate`

`uint32_t`

`maxFragmentShadingRateCoverageSamples`

`pipelineFragmentShadingRate`

[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html)

`maxFragmentShadingRateRasterizationSamples`

`pipelineFragmentShadingRate`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`fragmentShadingRateWithShaderDepthStencilWrites`

`pipelineFragmentShadingRate`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`fragmentShadingRateWithSampleMask`

`pipelineFragmentShadingRate`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`fragmentShadingRateWithShaderSampleMask`

`pipelineFragmentShadingRate`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`fragmentShadingRateWithConservativeRasterization`

`pipelineFragmentShadingRate`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`fragmentShadingRateWithFragmentShaderInterlock`

`pipelineFragmentShadingRate`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`fragmentShadingRateWithCustomSampleLocations`

`pipelineFragmentShadingRate`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`fragmentShadingRateStrictMultiplyCombiner`

`pipelineFragmentShadingRate`

[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html)

`maxFragmentShadingRateInvocationCount`

`supersampleFragmentShadingRates`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`combinedImageSamplerDescriptorSingleArray`

`VK_EXT_descriptor_buffer`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`bufferlessPushDescriptors`

`VK_EXT_descriptor_buffer`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`allowSamplerImageViewPostSubmitCreation`

`VK_EXT_descriptor_buffer`

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`descriptorBufferOffsetAlignment`

`VK_EXT_descriptor_buffer`

`uint32_t`

`maxDescriptorBufferBindings`

`VK_EXT_descriptor_buffer`

`uint32_t`

`maxResourceDescriptorBufferBindings`

`VK_EXT_descriptor_buffer`

`uint32_t`

`maxSamplerDescriptorBufferBindings`

`VK_EXT_descriptor_buffer`

`uint32_t`

`maxEmbeddedImmutableSamplerBindings`

`VK_EXT_descriptor_buffer`

`uint32_t`

`maxEmbeddedImmutableSamplers`

`VK_EXT_descriptor_buffer`

`size_t`

`bufferCaptureReplayDescriptorDataSize`

`VK_EXT_descriptor_buffer`

`size_t`

`imageCaptureReplayDescriptorDataSize`

`VK_EXT_descriptor_buffer`

`size_t`

`imageViewCaptureReplayDescriptorDataSize`

`VK_EXT_descriptor_buffer`

`size_t`

`samplerCaptureReplayDescriptorDataSize`

`VK_EXT_descriptor_buffer`

`size_t`

`accelerationStructureCaptureReplayDescriptorDataSize`

`VK_EXT_descriptor_buffer`

`size_t`

`samplerDescriptorSize`

`VK_EXT_descriptor_buffer`

`size_t`

`combinedImageSamplerDescriptorSize`

`VK_EXT_descriptor_buffer`

`size_t`

`sampledImageDescriptorSize`

`VK_EXT_descriptor_buffer`

`size_t`

`storageImageDescriptorSize`

`VK_EXT_descriptor_buffer`

`size_t`

`uniformTexelBufferDescriptorSize`

`VK_EXT_descriptor_buffer`

`size_t`

`robustUniformTexelBufferDescriptorSize`

`VK_EXT_descriptor_buffer`

`size_t`

`storageTexelBufferDescriptorSize`

`VK_EXT_descriptor_buffer`

`size_t`

`robustStorageTexelBufferDescriptorSize`

`VK_EXT_descriptor_buffer`

`size_t`

`uniformBufferDescriptorSize`

`VK_EXT_descriptor_buffer`

`size_t`

`robustUniformBufferDescriptorSize`

`VK_EXT_descriptor_buffer`

`size_t`

`storageBufferDescriptorSize`

`VK_EXT_descriptor_buffer`

`size_t`

`robustStorageBufferDescriptorSize`

`VK_EXT_descriptor_buffer`

`size_t`

`inputAttachmentDescriptorSize`

`VK_EXT_descriptor_buffer`

`size_t`

`accelerationStructureDescriptorSize`

`VK_EXT_descriptor_buffer`

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`maxSamplerDescriptorBufferRange`

`VK_EXT_descriptor_buffer`

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`maxResourceDescriptorBufferRange`

`VK_EXT_descriptor_buffer`

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`samplerDescriptorBufferAddressSpaceSize`

`VK_EXT_descriptor_buffer`

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`resourceDescriptorBufferAddressSpaceSize`

`VK_EXT_descriptor_buffer`

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`descriptorBufferAddressSpaceSize`

`VK_EXT_descriptor_buffer`

`size_t`

`combinedImageSamplerDensityMapDescriptorSize`

`VK_EXT_descriptor_buffer`

`uint32_t`

`maxSubpassShadingWorkgroupSizeAspectRatio`

`subpassShading`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`graphicsPipelineLibraryFastLinking`

`graphicsPipelineLibrary`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`graphicsPipelineLibraryIndependentInterpolationDecoration`

`graphicsPipelineLibrary`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`triStripVertexOrderIndependentOfProvokingVertex`

\-

`uint32_t`

`maxWeightFilterPhases`

`textureSampleWeighted`

2 × `uint32_t`

`maxWeightFilterDimension`

`textureSampleWeighted`

2 × `uint32_t`

`maxBlockMatchRegion`

`textureBlockMatch`

2 × `uint32_t`

`maxBoxFilterBlockSize`

`textureBoxFilter`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`dynamicPrimitiveTopologyUnrestricted`

`VK_EXT_extended_dynamic_state3`

`uint32_t`

`maxOpacity2StateSubdivisionLevel`

`VK_EXT_opacity_micromap`

`uint32_t`

`maxOpacity4StateSubdivisionLevel`

`VK_EXT_opacity_micromap`

`uint64_t`

`maxDecompressionIndirectCount`

`VK_NV_memory_decompression`

3 × `uint32_t`

`maxWorkGroupCount`

`VK_HUAWEI_cluster_culling_shader`

3 × `uint32_t`

`maxWorkGroupSize`

`VK_HUAWEI_cluster_culling_shader`

`uint32_t`

`maxOutputClusterCount`

`VK_HUAWEI_cluster_culling_shader`

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`indirectBufferOffsetAlignment`

`VK_HUAWEI_cluster_culling_shader`

`uint32_t`

`maxExecutionGraphDepth`

`shaderEnqueue`

`uint32_t`

`maxExecutionGraphShaderOutputNodes`

`shaderEnqueue`

`uint32_t`

`maxExecutionGraphShaderPayloadSize`

`shaderEnqueue`

`uint32_t`

`maxExecutionGraphShaderPayloadCount`

`shaderEnqueue`

`uint32_t`

`executionGraphDispatchAddressAlignment`

`shaderEnqueue`

`uint32_t`

`maxExecutionGraphVertexBufferBindings`

`shaderEnqueue`

3 × `uint32_t`

`maxExecutionGraphWorkgroupCount`

`shaderEnqueue`

`uint32_t`

`maxExecutionGraphWorkgroups`

`shaderEnqueue`

`uint32_t`

`maxIndirectShaderObjectCount`

`shaderObject`

[VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

`extendedSparseAddressSpaceSize`

`sparseBinding`, `extendedSparseAddressSpace`

`uint32_t`

`supportedImageAlignmentMask`

`imageAlignmentControl`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`separateDepthStencilAttachmentAccess`

[`maintenance7`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance7)

`uint32_t`

`maxDescriptorSetTotalUniformBuffersDynamic`

`maintenance7`

`uint32_t`

`maxDescriptorSetTotalStorageBuffersDynamic`

`maintenance7`

`uint32_t`

`maxDescriptorSetTotalBuffersDynamic`

`maintenance7`

`uint32_t`

`maxDescriptorSetUpdateAfterBindTotalUniformBuffersDynamic`

`maintenance7`

`uint32_t`

`maxDescriptorSetUpdateAfterBindTotalStorageBuffersDynamic`

`maintenance7`

`uint32_t`

`maxDescriptorSetUpdateAfterBindTotalBuffersDynamic`

`maintenance7`

`uint32_t`

`cooperativeMatrixWorkgroupScopeMaxWorkgroupSize`

`cooperativeMatrixWorkgroupScope`

`uint32_t`

`cooperativeMatrixFlexibleDimensionsMaxDimension`

`cooperativeMatrixFlexibleDimensions`

`uint32_t`

`cooperativeMatrixWorkgroupScopeReservedSharedMemory`

`cooperativeMatrixWorkgroupScope`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`shaderSignedZeroInfNanPreserveFloat16`

`shaderFloat16`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`cooperativeVectorTrainingFloat16Accumulation`

\-

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`cooperativeVectorTrainingFloat32Accumulation`

\-

`uint32_t`

`maxApronSize`

`tileShadingApron`

[VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html)

`preferNonCoherent`

`tileShading`

2 × `uint32_t`

`tileGranularity`

`tileShading`

2 × `uint32_t`

`maxTileShadingRate`

`tileShadingDispatchTile`

Table 2. Required Limits     Limit Unsupported Limit Supported Limit Limit Type1

`maxImageDimension1D`

\-

4096 (Vulkan Core)  
8192 (Vulkan Roadmap 2022, Vulkan 1.4)

min

`maxImageDimension2D`

\-

4096 (Vulkan Core)  
8192 (Vulkan Roadmap 2022, Vulkan 1.4)

min

`maxImageDimension3D`

\-

256 (Vulkan Core)  
512 (Vulkan 1.4)

min

`maxImageDimensionCube`

\-

4096 (Vulkan Core)  
8192 (Vulkan Roadmap 2022, Vulkan 1.4)

min

`maxImageArrayLayers`

\-

256 (Vulkan Core)  
2048 (Vulkan Roadmap 2022, Vulkan 1.4)

min

`maxTexelBufferElements`

\-

65536

min

`maxUniformBufferRange`

\-

16384 (Vulkan Core)  
65536 (Vulkan Roadmap 2022, Vulkan 1.4)

min

`maxStorageBufferRange`

\-

227

min

`maxPushConstantsSize`

\-

128 (Vulkan Core)  
256 (Vulkan 1.4)

min

`maxMemoryAllocationCount`

\-

4096

min

`maxSamplerAllocationCount`

\-

4000

min

`bufferImageGranularity`

\-

131072 (Vulkan Core)  
4096 (Vulkan Roadmap 2022, Vulkan 1.4)

max

`sparseAddressSpaceSize`

0

231

min

`maxBoundDescriptorSets`

\-

4 (Vulkan Core)  
7 (Vulkan Roadmap 2024, Vulkan 1.4)

min

`maxPerStageDescriptorSamplers`

\-

16

min

`maxPerStageDescriptorUniformBuffers`

\-

12 (Vulkan Core)  
15 (Vulkan Roadmap 2022, Vulkan 1.4)

min

`maxPerStageDescriptorStorageBuffers`

\-

4 (Vulkan Core)  
30 (Vulkan Roadmap 2022)

min

`maxPerStageDescriptorSampledImages`

\-

16 (Vulkan Core)  
200 (Vulkan Roadmap 2022)

min

`maxPerStageDescriptorStorageImages`

\-

4 (Vulkan Core)  
144 (Vulkan Roadmap 2022)

min

`maxPerStageDescriptorInputAttachments`

\-

4

min

`maxPerStageResources`

\-

128 2 (Vulkan Core)  
200 (Vulkan Roadmap 2022, Vulkan 1.4)

min

`maxDescriptorSetSamplers`

\-

96 8 (Vulkan Core)  
576 (Vulkan Roadmap 2022)

min, *n* × PerStage

`maxDescriptorSetUniformBuffers`

\-

72 8 (Vulkan Core)  
90 (Vulkan Roadmap 2022, Vulkan 1.4)

min, *n* × PerStage

`maxDescriptorSetUniformBuffersDynamic`

\-

8

min

`maxDescriptorSetStorageBuffers`

\-

24 8 (Vulkan Core)  
96 (Vulkan Roadmap 2022, Vulkan 1.4)

min, *n* × PerStage

`maxDescriptorSetStorageBuffersDynamic`

\-

4

min

`maxDescriptorSetTotalUniformBuffersDynamic`

\-

`maxDescriptorSetUniformBuffersDynamic`

min

`maxDescriptorSetTotalStorageBuffersDynamic`

\-

`maxDescriptorSetStorageBuffersDynamic`

min

`maxDescriptorSetTotalBuffersDynamic`

\-

`maxDescriptorSetUniformBuffersDynamic` + `maxDescriptorSetStorageBuffersDynamic`

min

`maxDescriptorSetSampledImages`

\-

96 8 (Vulkan Core)  
1800 (Vulkan Roadmap 2022)

min, *n* × PerStage

`maxDescriptorSetStorageImages`

\-

24 8 (Vulkan Core)  
144 (Vulkan Roadmap 2022, Vulkan 1.4)

min, *n* × PerStage

`maxDescriptorSetInputAttachments`

\-

4

min

`maxVertexInputAttributes`

\-

16

min

`maxVertexInputBindings`

\-

16 10

min

`maxVertexInputAttributeOffset`

\-

2047

min

`maxVertexInputBindingStride`

\-

2048

min

`maxVertexOutputComponents`

\-

64

min

`maxTessellationGenerationLevel`

0

64

min

`maxTessellationPatchSize`

0

32

min

`maxTessellationControlPerVertexInputComponents`

0

64

min

`maxTessellationControlPerVertexOutputComponents`

0

64

min

`maxTessellationControlPerPatchOutputComponents`

0

120

min

`maxTessellationControlTotalOutputComponents`

0

2048

min

`maxTessellationEvaluationInputComponents`

0

64

min

`maxTessellationEvaluationOutputComponents`

0

64

min

`maxGeometryShaderInvocations`

0

32

min

`maxGeometryInputComponents`

0

64

min

`maxGeometryOutputComponents`

0

64

min

`maxGeometryOutputVertices`

0

256

min

`maxGeometryTotalOutputComponents`

0

1024

min

`maxFragmentInputComponents`

\-

64

min

`maxFragmentOutputAttachments`

\-

4

min

`maxFragmentDualSrcAttachments`

0

1

min

`maxFragmentCombinedOutputResources`

\-

4 (Vulkan Core)  
16 (Vulkan Roadmap 2022, Vulkan 1.4)

min

`maxComputeSharedMemorySize`

\-

16384

min

`maxComputeWorkGroupCount`

\-

(65535,65535,65535)

min

`maxComputeWorkGroupInvocations`

\-

128 (Vulkan Core)  
256 (Vulkan Roadmap 2022, Vulkan 1.4)

min

`maxComputeWorkGroupSize`

\-

(128,128,64) (Vulkan Core)  
(256,256,64) (Vulkan Roadmap 2022, Vulkan 1.4)

min

`subgroupSize`

\-

1/4 (Vulkan Core)  
4 (Vulkan Roadmap 2022)

min

`subgroupSupportedStages`

\-

`VK_SHADER_STAGE_COMPUTE_BIT` (Vulkan Core)  
`VK_SHADER_STAGE_COMPUTE_BIT` |  
`VK_SHADER_STAGE_FRAGMENT_BIT` (Vulkan Roadmap 2022)

bitfield

`subgroupSupportedOperations`

\-

`VK_SUBGROUP_FEATURE_BASIC_BIT` (Vulkan Core)  
`VK_SUBGROUP_FEATURE_BASIC_BIT` |  
`VK_SUBGROUP_FEATURE_VOTE_BIT` |  
`VK_SUBGROUP_FEATURE_ARITHMETIC_BIT` |  
`VK_SUBGROUP_FEATURE_BALLOT_BIT` |  
`VK_SUBGROUP_FEATURE_SHUFFLE_BIT` |  
`VK_SUBGROUP_FEATURE_SHUFFLE_RELATIVE_BIT` |  
`VK_SUBGROUP_FEATURE_QUAD_BIT` (Vulkan Roadmap 2022)

bitfield

`shaderSignedZeroInfNanPreserveFloat16`

\-

\- (Vulkan Core)  
`VK_TRUE` (Vulkan Roadmap 2022, Vulkan 1.4)

Boolean

`shaderSignedZeroInfNanPreserveFloat32`

\-

\- (Vulkan Core)  
`VK_TRUE` (Vulkan Roadmap 2022, Vulkan 1.4)

Boolean

`shaderRoundingModeRTEFloat16`

\-

`VK_FALSE` (Vulkan Core)  
`VK_TRUE` (Vulkan Roadmap 2024)

Boolean

`shaderRoundingModeRTEFloat32`

\-

`VK_FALSE` (Vulkan Core)  
`VK_TRUE` (Vulkan Roadmap 2024)

Boolean

`maxSubgroupSize`

\-

\- (Vulkan Core)  
4 (Vulkan Roadmap 2022)

min

`subPixelPrecisionBits`

\-

4

min

`subTexelPrecisionBits`

\-

4 (Vulkan Core)  
8 (Vulkan Roadmap 2022, Vulkan 1.4)

min

`mipmapPrecisionBits`

\-

4 (Vulkan Core)  
6 (Vulkan Roadmap 2022, Vulkan 1.4)

min

`maxDrawIndexedIndexValue`

224-1

232-1

min

`maxDrawIndirectCount`

1

216-1

min

`maxSamplerLodBias`

\-

2 (Vulkan Core)  
14 (Vulkan Roadmap 2022, Vulkan 1.4)

min

`maxSamplerAnisotropy`

1

16

min

`maxViewports`

1

16

min

`maxViewportDimensions` 3

\-

(4096,4096) (Vulkan Core)  
(7680,7680) (Vulkan 1.4)

min

`viewportBoundsRange` 4

\-

(-8192,8191) (Vulkan Core)  
(-15360,15359) (Vulkan 1.4)

(max,min)

`viewportSubPixelBits`

\-

0

min

`minMemoryMapAlignment`

\-

64

min

`minTexelBufferOffsetAlignment`

\-

256

max

`minUniformBufferOffsetAlignment`

\-

256

max

`minStorageBufferOffsetAlignment`

\-

256

max

`minTexelOffset`

\-

-8

max

`maxTexelOffset`

\-

7

min

`minTexelGatherOffset`

0

-8

max

`maxTexelGatherOffset`

0

7

min

`minInterpolationOffset`

0.0

-0.5 5

max

`maxInterpolationOffset`

0.0

0.5 - (1 ULP) 5

min

`subPixelInterpolationOffsetBits`

0

4 5

min

`maxFramebufferWidth`

\-

4096 (Vulkan Core)  
7680 (Vulkan 1.4)

min

`maxFramebufferHeight`

\-

4096 (Vulkan Core)  
7680 (Vulkan 1.4)

min

`maxFramebufferLayers`

\-

256

min

`framebufferColorSampleCounts`

\-

(`VK_SAMPLE_COUNT_1_BIT` | `VK_SAMPLE_COUNT_4_BIT`)

min

`framebufferIntegerColorSampleCounts`

\-

(`VK_SAMPLE_COUNT_1_BIT`)

min

`framebufferDepthSampleCounts`

\-

(`VK_SAMPLE_COUNT_1_BIT` | `VK_SAMPLE_COUNT_4_BIT`)

min

`framebufferStencilSampleCounts`

\-

(`VK_SAMPLE_COUNT_1_BIT` | `VK_SAMPLE_COUNT_4_BIT`)

min

`framebufferNoAttachmentsSampleCounts`

\-

(`VK_SAMPLE_COUNT_1_BIT` | `VK_SAMPLE_COUNT_4_BIT`)

min

`maxColorAttachments`

\-

4 (Vulkan Core)  
7 (Vulkan Roadmap 2022)  
8 (Vulkan Roadmap 2024, Vulkan 1.4)

min

`sampledImageColorSampleCounts`

\-

(`VK_SAMPLE_COUNT_1_BIT` | `VK_SAMPLE_COUNT_4_BIT`)

min

`sampledImageIntegerSampleCounts`

\-

`VK_SAMPLE_COUNT_1_BIT`

min

`sampledImageDepthSampleCounts`

\-

(`VK_SAMPLE_COUNT_1_BIT` | `VK_SAMPLE_COUNT_4_BIT`)

min

`sampledImageStencilSampleCounts`

\-

(`VK_SAMPLE_COUNT_1_BIT` | `VK_SAMPLE_COUNT_4_BIT`)

min

`storageImageSampleCounts`

`VK_SAMPLE_COUNT_1_BIT`

(`VK_SAMPLE_COUNT_1_BIT` | `VK_SAMPLE_COUNT_4_BIT`)

min

`maxSampleMaskWords`

\-

1

min

`timestampComputeAndGraphics`

\-

\- (Vulkan Core)  
`VK_TRUE` (Vulkan Roadmap 2024, Vulkan 1.4)

Boolean

`timestampPeriod`

\-

\-

duration

`maxClipDistances`

0

8

min

`maxCullDistances`

0

8

min

`maxCombinedClipAndCullDistances`

0

8

min

`discreteQueuePriorities`

\-

2

min

`pointSizeRange`

(1.0,1.0)

(1.0,64.0 - ULP) 6 (Vulkan Core)  
(1.0,256.0 - `pointSizeGranularity`) (Vulkan 1.4)

(max,min)

`lineWidthRange`

(1.0,1.0)

(1.0,8.0 - ULP) 7

(max,min)

`pointSizeGranularity`

0.0

1.0 6 (Vulkan Core)  
0.125 (Vulkan Roadmap 2022, Vulkan 1.4)

max, fixed point increment

`lineWidthGranularity`

0.0

1.0 7 (Vulkan Core)  
0.5 (Vulkan Roadmap 2022, Vulkan 1.4)

max, fixed point increment

`strictLines`

\-

\-

implementation-dependent

`standardSampleLocations`

\-

\- (Vulkan Core)  
`VK_TRUE` (Vulkan Roadmap 2022, Vulkan 1.4)

Boolean

`optimalBufferCopyOffsetAlignment`

\-

\-

recommendation

`optimalBufferCopyRowPitchAlignment`

\-

\-

recommendation

`nonCoherentAtomSize`

\-

256

max

`maxPushDescriptors`

\-

32

min

`maxMultiviewViewCount`

\-

6

min

`maxMultiviewInstanceIndex`

\-

227-1

min

`maxDiscardRectangles`

0

4

min

`sampleLocationSampleCounts`

\-

`VK_SAMPLE_COUNT_4_BIT`

min

`maxSampleLocationGridSize`

\-

(1,1)

min

`sampleLocationCoordinateRange`

\-

(0.0, 0.9375)

(max,min)

`sampleLocationSubPixelBits`

\-

4

min

`variableSampleLocations`

\-

`VK_FALSE`

implementation-dependent

`nativeUnalignedPerformance`

\-

`VK_FALSE`

implementation-dependent

`minImportedHostPointerAlignment`

\-

65536

max

`perViewPositionAllComponents`

\-

\-

implementation-dependent

`filterMinmaxSingleComponentFormats`

\-

\-

implementation-dependent

`filterMinmaxImageComponentMapping`

\-

\-

implementation-dependent

`advancedBlendMaxColorAttachments`

\-

1

min

`advancedBlendIndependentBlend`

\-

`VK_FALSE`

implementation-dependent

`advancedBlendNonPremultipliedSrcColor`

\-

`VK_FALSE`

implementation-dependent

`advancedBlendNonPremultipliedDstColor`

\-

`VK_FALSE`

implementation-dependent

`advancedBlendCorrelatedOverlap`

\-

`VK_FALSE`

implementation-dependent

`advancedBlendAllOperations`

\-

`VK_FALSE`

implementation-dependent

`maxPerSetDescriptors`

\-

1024

min

`maxMemoryAllocationSize`

\-

230

min

`maxBufferSize`

\-

230

min

`primitiveOverestimationSize`

\-

0.0

min

`maxExtraPrimitiveOverestimationSize`

\-

0.0

min

`extraPrimitiveOverestimationSizeGranularity`

\-

0.0

min

`primitiveUnderestimation`

\-

`VK_FALSE`

implementation-dependent

`conservativePointAndLineRasterization`

\-

`VK_FALSE`

implementation-dependent

`degenerateTrianglesRasterized`

\-

`VK_FALSE`

implementation-dependent

`degenerateLinesRasterized`

\-

`VK_FALSE`

implementation-dependent

`fullyCoveredFragmentShaderInputVariable`

\-

`VK_FALSE`

implementation-dependent

`conservativeRasterizationPostDepthCoverage`

\-

`VK_FALSE`

implementation-dependent

`maxUpdateAfterBindDescriptorsInAllPools`

0

500000

min

`shaderUniformBufferArrayNonUniformIndexingNative`

\-

`VK_FALSE`

implementation-dependent

`shaderSampledImageArrayNonUniformIndexingNative`

\-

`VK_FALSE`

implementation-dependent

`shaderStorageBufferArrayNonUniformIndexingNative`

\-

`VK_FALSE`

implementation-dependent

`shaderStorageImageArrayNonUniformIndexingNative`

\-

`VK_FALSE`

implementation-dependent

`shaderInputAttachmentArrayNonUniformIndexingNative`

\-

`VK_FALSE`

implementation-dependent

`maxPerStageDescriptorUpdateAfterBindSamplers`

0 9

500000 9

min

`maxPerStageDescriptorUpdateAfterBindUniformBuffers`

0 9

12 9

min

`maxPerStageDescriptorUpdateAfterBindStorageBuffers`

0 9

500000 9

min

`maxPerStageDescriptorUpdateAfterBindSampledImages`

0 9

500000 9

min

`maxPerStageDescriptorUpdateAfterBindStorageImages`

0 9

500000 9

min

`maxPerStageDescriptorUpdateAfterBindInputAttachments`

0 9

4 9 (Vulkan Core)  
7 (Vulkan Roadmap 2022)

min

`maxPerStageUpdateAfterBindResources`

0 9

500000 9

min

`maxDescriptorSetUpdateAfterBindSamplers`

0 9

500000 9

min

`maxDescriptorSetUpdateAfterBindUniformBuffers`

0 9

72 8 9

min, *n* × PerStage

`maxDescriptorSetUpdateAfterBindUniformBuffersDynamic`

0 9

8 9

min

`maxDescriptorSetUpdateAfterBindStorageBuffers`

0 9

500000 9

min

`maxDescriptorSetUpdateAfterBindStorageBuffersDynamic`

0 9

4 9

min

`maxDescriptorSetUpdateAfterBindTotalUniformBuffersDynamic`

0 9

`maxDescriptorSetUpdateAfterBindUniformBuffersDynamic`

min

`maxDescriptorSetUpdateAfterBindTotalStorageBuffersDynamic`

0 9

`maxDescriptorSetUpdateAfterBindStorageBuffersDynamic`

min

`maxDescriptorSetUpdateAfterBindTotalBuffersDynamic`

0 9

`maxDescriptorSetUpdateAfterBindUniformBuffersDynamic` + `maxDescriptorSetUpdateAfterBindStorageBuffersDynamic`

min

`maxDescriptorSetUpdateAfterBindSampledImages`

0 9

500000 9

min

`maxDescriptorSetUpdateAfterBindStorageImages`

0 9

500000 9

min

`maxDescriptorSetUpdateAfterBindInputAttachments`

0 9

4 9

min

`maxInlineUniformBlockSize`

\-

256

min

`maxPerStageDescriptorInlineUniformBlocks`

\-

4

min

`maxPerStageDescriptorUpdateAfterBindInlineUniformBlocks`

\-

4

min

`maxDescriptorSetInlineUniformBlocks`

\-

4

min

`maxDescriptorSetUpdateAfterBindInlineUniformBlocks`

\-

4

min

`maxInlineUniformTotalSize`

\-

256

min

`maxVertexAttribDivisor`

\-

216-1

min

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxDrawMeshTasksCount`

\-

216-1

min

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxTaskWorkGroupInvocations`

\-

32

min

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxTaskWorkGroupSize`

\-

(32,1,1)

min

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxTaskTotalMemorySize`

\-

16384

min

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxTaskOutputCount`

\-

216-1

min

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshWorkGroupInvocations`

\-

32

min

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshWorkGroupSize`

\-

(32,1,1)

min

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshTotalMemorySize`

\-

16384

min

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshOutputVertices`

\-

256

min

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshOutputPrimitives`

\-

256

min

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshMultiviewViewCount`

\-

1

min

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`meshOutputPerVertexGranularity`

\-

\-

implementation-dependent

[VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`meshOutputPerPrimitiveGranularity`

\-

\-

implementation-dependent

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskWorkGroupTotalCount`

\-

2^22

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskWorkGroupCount`

\-

(65535,65535,65535)

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskWorkGroupInvocations`

\-

128

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskWorkGroupSize`

\-

(128,128,128)

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskPayloadSize`

\-

16384

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskSharedMemorySize`

\-

32768

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskPayloadAndSharedMemorySize`

\-

32768

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupTotalCount`

\-

2^22

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupCount`

\-

(65535,65535,65535)

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupInvocations`

\-

128

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupSize`

\-

(128,128,128)

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshSharedMemorySize`

\-

28672

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshPayloadAndSharedMemorySize`

\-

28672

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshOutputMemorySize`

\-

32768

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshPayloadAndOutputMemorySize`

\-

48128

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshOutputComponents`

\-

128

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshOutputVertices`

\-

256

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshOutputPrimitives`

\-

256

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshOutputLayers`

\-

8

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshMultiviewViewCount`

\-

1

min

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`meshOutputPerVertexGranularity`

0

32

max

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`meshOutputPerPrimitiveGranularity`

0

32

max

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxPreferredTaskWorkGroupInvocations`

\-

\-

implementation-dependent

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxPreferredMeshWorkGroupInvocations`

\-

\-

implementation-dependent

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`prefersLocalInvocationVertexOutput`

\-

\-

implementation-dependent

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`prefersLocalInvocationPrimitiveOutput`

\-

\-

implementation-dependent

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`prefersCompactVertexOutput`

\-

\-

implementation-dependent

[VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`prefersCompactPrimitiveOutput`

\-

\-

implementation-dependent

`maxTransformFeedbackStreams`

\-

1

min

`maxTransformFeedbackBuffers`

\-

1

min

`maxTransformFeedbackBufferSize`

\-

227

min

`maxTransformFeedbackStreamDataSize`

\-

512

min

`maxTransformFeedbackBufferDataSize`

\-

512

min

`maxTransformFeedbackBufferDataStride`

\-

512

min

`transformFeedbackQueries`

\-

`VK_FALSE`

implementation-dependent

`transformFeedbackStreamsLinesTriangles`

\-

`VK_FALSE`

implementation-dependent

`transformFeedbackRasterizationStreamSelect`

\-

`VK_FALSE`

implementation-dependent

`transformFeedbackDraw`

\-

`VK_FALSE`

implementation-dependent

`minFragmentDensityTexelSize`

\-

(1,1)

min

`maxFragmentDensityTexelSize`

\-

(1,1)

min

`fragmentDensityInvocations`

\-

\-

implementation-dependent

`subsampledLoads`

`VK_TRUE`

`VK_FALSE`

implementation-dependent

`subsampledCoarseReconstructionEarlyAccess`

`VK_FALSE`

`VK_FALSE`

implementation-dependent

`maxSubsampledArrayLayers`

2

2

min

`maxDescriptorSetSubsampledSamplers`

1

1

min

`fragmentDensityOffsetGranularity`

\-

(1024,1024)

max

`maxFragmentDensityMapLayers`

\-

(2)

max

[VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html)::`shaderGroupHandleSize`

\-

16

min

[VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html)::`maxRecursionDepth`

\-

31

min

[VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`shaderGroupHandleSize`

\-

32

exact

[VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`maxRayRecursionDepth`

\-

1

min

`maxShaderGroupStride`

\-

4096

min

`shaderGroupBaseAlignment`

\-

64

max

`maxGeometryCount`

\-

224-1

min

`maxInstanceCount`

\-

224-1

min

`maxTriangleCount`

\-

229-1

min

`maxPrimitiveCount`

\-

229-1

min

`maxPerStageDescriptorAccelerationStructures`

\-

16

min

`maxPerStageDescriptorUpdateAfterBindAccelerationStructures`

\-

500000 9

min

`maxVerticesPerCluster`

\-

256

min

`maxTrianglesPerCluster`

\-

256

min

`clusterScratchByteAlignment`

\-

256

max

`clusterByteAlignment`

\-

128

max

`clusterTemplateByteAlignment`

\-

32

max

`clusterBottomLevelByteAlignment`

\-

256

max

`clusterTemplateBoundsByteAlignment`

\-

32

max

`maxClusterGeometryIndex`

\-

224-1

min

`maxDescriptorSetAccelerationStructures`

\-

16

min

`maxDescriptorSetUpdateAfterBindAccelerationStructures`

\-

500000 9

min

`minAccelerationStructureScratchOffsetAlignment`

\-

256

max

`shaderGroupHandleCaptureReplaySize`

\-

64

max

`maxRayDispatchInvocationCount`

\-

230

min

`shaderGroupHandleAlignment`

\-

32

max

`maxRayHitAttributeSize`

\-

32

min

`maxPartitionCount`

\-

224-1

min

`maxTimelineSemaphoreValueDifference`

\-

231-1

min

`lineSubPixelPrecisionBits`

\-

4

min

`maxGraphicsShaderGroupCount`

\-

212

min

`maxIndirectCommandsStreamCount` + (for NV extension)

\-

212

min

`maxIndirectCommandsStreamStride`

\-

2048

min

`minIndirectCommandsBufferOffsetAlignment`

\-

256

max

`minSequencesCountBufferOffsetAlignment`

\-

256

max

`minSequencesIndexBufferOffsetAlignment`

\-

256

max

`maxIndirectSequenceCount`

\-

220

min

`maxIndirectCommandsTokenCount`

\-

16

min

`maxIndirectCommandsTokenOffset`

\-

2047

min

`maxIndirectPipelineCount`

\-

212

min

`deviceGeneratedCommandsTransformFeedback`

\-

false

implementation-dependent

`deviceGeneratedCommandsMultiDrawIndirectCount`

\-

false

implementation-dependent

`maxIndirectShaderObjectCount`

0

212

implementation-dependent

`maxIndirectCommandsIndirectStride`

\-

2048

min

`supportedIndirectCommandsInputModes`

\-

`VK_INDIRECT_COMMANDS_INPUT_MODE_VULKAN_INDEX_BUFFER_EXT`

min

`supportedIndirectCommandsShaderStages`

\-

(`VK_SHADER_STAGE_COMPUTE_BIT` | `VK_SHADER_STAGE_VERTEX_BIT` | `VK_SHADER_STAGE_FRAGMENT_BIT`)

min

`supportedIndirectCommandsShaderStagesPipelineBinding`

\-

0

min

`supportedIndirectCommandsShaderStagesShaderBinding`

\-

0

min

`maxCustomBorderColorSamplers`

\-

32

min

`robustStorageBufferAccessSizeAlignment`

\-

4

max

`robustUniformBufferAccessSizeAlignment`

\-

256

max

`minFragmentShadingRateAttachmentTexelSize`

(0,0)

(32,32)

max

`maxFragmentShadingRateAttachmentTexelSize`

(0,0)

(8,8)

min

`maxFragmentShadingRateAttachmentTexelSizeAspectRatio`

0

1

min

`primitiveFragmentShadingRateWithMultipleViewports`

`VK_FALSE`

`VK_FALSE`

implementation-dependent

`layeredShadingRateAttachments`

`VK_FALSE`

`VK_FALSE`

implementation-dependent

`fragmentShadingRateNonTrivialCombinerOps`

\-

`VK_FALSE`

implementation-dependent

`maxFragmentSize`

\-

(2,2)

min

`maxFragmentSizeAspectRatio`

\-

2

min

`maxFragmentShadingRateCoverageSamples`

\-

16

min

`maxFragmentShadingRateRasterizationSamples`

\-

`VK_SAMPLE_COUNT_4_BIT`

min

`fragmentShadingRateWithShaderDepthStencilWrites`

\-

`VK_FALSE`

implementation-dependent

`fragmentShadingRateWithSampleMask`

\-

`VK_FALSE`

implementation-dependent

`fragmentShadingRateWithShaderSampleMask`

\-

`VK_FALSE`

implementation-dependent

`fragmentShadingRateWithConservativeRasterization`

\-

`VK_FALSE`

implementation-dependent

`fragmentShadingRateWithFragmentShaderInterlock`

\-

`VK_FALSE`

implementation-dependent

`fragmentShadingRateWithCustomSampleLocations`

\-

`VK_FALSE`

implementation-dependent

`fragmentShadingRateStrictMultiplyCombiner`

\-

`VK_FALSE`

implementation-dependent

`maxFragmentShadingRateInvocationCount`

\-

`VK_SAMPLE_COUNT_4_BIT`

min

`combinedImageSamplerDescriptorSingleArray`

\-

`VK_FALSE`

implementation-dependent

`bufferlessPushDescriptors`

\-

`VK_FALSE`

implementation-dependent

`allowSamplerImageViewPostSubmitCreation`

\-

`VK_FALSE`

implementation-dependent

`descriptorBufferOffsetAlignment`

\-

256

max

`maxDescriptorBufferBindings`

\-

3

min

`maxResourceDescriptorBufferBindings`

\-

1

min

`maxSamplerDescriptorBufferBindings`

\-

1

min

`maxEmbeddedImmutableSamplerBindings`

\-

1

min

`maxEmbeddedImmutableSamplers`

\-

2032

min

`bufferCaptureReplayDescriptorDataSize`

\-

64

max

`imageCaptureReplayDescriptorDataSize`

\-

64

max

`imageViewCaptureReplayDescriptorDataSize`

\-

64

max

`samplerCaptureReplayDescriptorDataSize`

\-

64

max

`accelerationStructureCaptureReplayDescriptorDataSize`

\-

64

max

`samplerDescriptorSize`

\-

256

max

`combinedImageSamplerDescriptorSize`

\-

256

max

`sampledImageDescriptorSize`

\-

256

max

`storageImageDescriptorSize`

\-

256

max

`uniformTexelBufferDescriptorSize`

\-

256

max

`robustUniformTexelBufferDescriptorSize`

\-

256

max

`storageTexelBufferDescriptorSize`

\-

256

max

`robustStorageTexelBufferDescriptorSize`

\-

256

max

`uniformBufferDescriptorSize`

\-

256

max

`robustUniformBufferDescriptorSize`

\-

256

max

`storageBufferDescriptorSize`

\-

256

max

`robustStorageBufferDescriptorSize`

\-

256

max

`inputAttachmentDescriptorSize`

\-

256

max

`accelerationStructureDescriptorSize`

\-

256

max

`maxSamplerDescriptorBufferRange`

\-

211 × `samplerDescriptorSize`

min

`maxResourceDescriptorBufferRange`

\-

(220 - 215) × `maxResourceDescriptorSize` 12

min

`samplerDescriptorBufferAddressSpaceSize`

\-

227

min

`resourceDescriptorBufferAddressSpaceSize`

\-

227

min

`descriptorBufferAddressSpaceSize`

\-

227

min

`combinedImageSamplerDensityMapDescriptorSize`

\-

256

max

`maxSubpassShadingWorkgroupSizeAspectRatio`

0

1

min

`maxMultiDrawCount`

\-

1024

min

`maxCommandBufferNestingLevel`

\-

1

min

`graphicsPipelineLibraryFastLinking`

\-

`VK_FALSE`

implementation-dependent

`graphicsPipelineLibraryIndependentInterpolationDecoration`

\-

`VK_FALSE`

implementation-dependent

`triStripVertexOrderIndependentOfProvokingVertex`

\-

`VK_FALSE`

implementation-dependent

`maxWeightFilterPhases`

\-

1024

min

`maxWeightFilterDimension`

\-

(64,64)

min

`maxBlockMatchRegion`

\-

(64,64)

min

`maxBoxFilterBlockSize`

\-

(64,64)

min

`dynamicPrimitiveTopologyUnrestricted`

\-

\-

implementation-dependent

`maxOpacity2StateSubdivisionLevel`

\-

3

min

`maxOpacity4StateSubdivisionLevel`

\-

3

min

`maxDecompressionIndirectCount`

1

216-1

min

`maxWorkGroupCount`

\-

(65536,1,1)

min

`maxWorkGroupSize`

\-

(32,1,1)

min

`maxOutputClusterCount`

\-

1024

min

`indirectBufferOffsetAlignment`

\-

\-

implementation-dependent

`maxExecutionGraphDepth`

\-

32

min

`maxExecutionGraphShaderOutputNodes`

\-

256

min

`maxExecutionGraphShaderPayloadSize`

\-

32768

min

`maxExecutionGraphShaderPayloadCount`

\-

256

min

`executionGraphDispatchAddressAlignment`

\-

4

max

`maxExecutionGraphVertexBufferBindings`

\-

1024

min

`maxExecutionGraphWorkgroupCount`

\-

(65535,65535,65535)

min

`maxExecutionGraphWorkgroups`

\-

224-1

min

`extendedSparseAddressSpaceSize`

0

`sparseAddressSpaceSize`

min

`renderPassStripeGranularity`

\-

(64,64)

max

`maxRenderPassStripes`

\-

32

min

`minPlacedMemoryMapAlignment`

\-

65536

max

`supportedImageAlignmentMask`

\-

1

min

`separateDepthStencilAttachmentAccess`

`VK_FALSE`

\-

implementation-dependent

`cooperativeMatrixWorkgroupScopeMaxWorkgroupSize`

\-

subgroupSize × 2

min

`cooperativeMatrixFlexibleDimensionsMaxDimension`

\-

256

min

`cooperativeMatrixWorkgroupScopeReservedSharedMemory`

\-

`maxComputeSharedMemorySize` / 2

max

`maxCooperativeVectorComponents`

\-

128

min

`maxApronSize`

\-

1

min

`preferNonCoherent`

\-

\-

implementation-dependent

`tileGranularity`

\-

(16,16)

min

`maxTileShadingRate`

\-

(8,8)

min

`maxTensorDimensionCount`

\-

4

min

`maxTensorElements`

\-

65536

min

`maxPerDimensionTensorElements`

\-

65536

min

`maxTensorStride`

\-

65536

min

`maxDescriptorSetStorageTensors`

\-

16

min

`maxPerStageDescriptorSetStorageTensors`

\-

16

min

`maxDescriptorSetUpdateAfterBindStorageTensors`

0

500000

min

`maxPerStageDescriptorUpdateAfterBindStorageTensors`

0

500000

min

`shaderTensorSupportedStages`

\-

`VK_SHADER_STAGE_COMPUTE_BIT`

bitfield

1

The **Limit Type** column specifies the limit is either the minimum limit all implementations **must** support, the maximum limit all implementations **must** support, or the exact value all implementations **must** support. For bitmasks a minimum limit is the least bits all implementations **must** set, but they **may** have additional bits set beyond this minimum.

2

The `maxPerStageResources` **must** be at least the smallest of the following:

- the sum of the `maxPerStageDescriptorUniformBuffers`, `maxPerStageDescriptorStorageBuffers`, `maxPerStageDescriptorSampledImages`, `maxPerStageDescriptorStorageImages`, `maxPerStageDescriptorInputAttachments`, `maxColorAttachments` limits, or
- 128\.
  
  It **may** not be possible to reach this limit in every stage.

3

See [`maxViewportDimensions`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxViewportDimensions) for the **required** relationship to other limits.

4

See [`viewportBoundsRange`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-viewportboundsrange) for the **required** relationship to other limits.

5

The values `minInterpolationOffset` and `maxInterpolationOffset` describe the closed interval of supported interpolation offsets: \[`minInterpolationOffset`, `maxInterpolationOffset`]. The ULP is determined by `subPixelInterpolationOffsetBits`. If `subPixelInterpolationOffsetBits` is 4, this provides increments of (1/24) = 0.0625, and thus the range of supported interpolation offsets would be \[-0.5, 0.4375].

6

The point size ULP is determined by `pointSizeGranularity`. If the `pointSizeGranularity` is 0.125, the range of supported point sizes **must** be at least \[1.0, 63.875].

7

The line width ULP is determined by `lineWidthGranularity`. If the `lineWidthGranularity` is 0.0625, the range of supported line widths **must** be at least \[1.0, 7.9375].

8

The minimum `maxDescriptorSet*` limit is *n* times the corresponding *specification* minimum `maxPerStageDescriptor*` limit, where *n* is the number of shader stages supported by the [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html). If all shader stages are supported, *n* = 6 (vertex, tessellation control, tessellation evaluation, geometry, fragment, compute).

9

The `UpdateAfterBind` descriptor limits **must** each be greater than or equal to the corresponding `non`-UpdateAfterBind limit.

10

If the `VK_KHR_portability_subset` extension is enabled, the required minimum value of `maxVertexInputBindings` is `8`.

12

`maxResourceDescriptorSize` is defined as the maximum value of `storageImageDescriptorSize`, `sampledImageDescriptorSize`, `robustUniformTexelBufferDescriptorSize`, `robustStorageTexelBufferDescriptorSize`, `robustUniformBufferDescriptorSize`, `robustStorageBufferDescriptorSize`, `inputAttachmentDescriptorSize`, and `accelerationStructureDescriptorSize`.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-minmax)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0