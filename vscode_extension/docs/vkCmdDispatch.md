# vkCmdDispatch(3) Manual Page

## Name

vkCmdDispatch - Dispatch compute work items



## [](#_c_specification)C Specification

To record a dispatch, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdDispatch(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    groupCountX,
    uint32_t                                    groupCountY,
    uint32_t                                    groupCountZ);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `groupCountX` is the number of local workgroups to dispatch in the X dimension.
- `groupCountY` is the number of local workgroups to dispatch in the Y dimension.
- `groupCountZ` is the number of local workgroups to dispatch in the Z dimension.

## [](#_description)Description

When the command is executed, a global workgroup consisting of `groupCountX` × `groupCountY` × `groupCountZ` local workgroups is assembled.

Valid Usage

- [](#VUID-vkCmdDispatch-magFilter-04553)VUID-vkCmdDispatch-magFilter-04553  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `magFilter` or `minFilter` equal to `VK_FILTER_LINEAR`, `reductionMode` equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`, and `compareEnable` equal to `VK_FALSE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`
- [](#VUID-vkCmdDispatch-magFilter-09598)VUID-vkCmdDispatch-magFilter-09598  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `magFilter` or `minFilter` equal to `VK_FILTER_LINEAR` and `reductionMode` equal to either `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT`
- [](#VUID-vkCmdDispatch-mipmapMode-04770)VUID-vkCmdDispatch-mipmapMode-04770  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `mipmapMode` equal to `VK_SAMPLER_MIPMAP_MODE_LINEAR`, `reductionMode` equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`, and `compareEnable` equal to `VK_FALSE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`
- [](#VUID-vkCmdDispatch-mipmapMode-09599)VUID-vkCmdDispatch-mipmapMode-09599  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `mipmapMode` equal to `VK_SAMPLER_MIPMAP_MODE_LINEAR` and `reductionMode` equal to either `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT`
- [](#VUID-vkCmdDispatch-unnormalizedCoordinates-09635)VUID-vkCmdDispatch-unnormalizedCoordinates-09635  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s `levelCount` and `layerCount` **must** be 1
- [](#VUID-vkCmdDispatch-None-08609)VUID-vkCmdDispatch-None-08609  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s `viewType` **must** be `VK_IMAGE_VIEW_TYPE_1D` or `VK_IMAGE_VIEW_TYPE_2D`
- [](#VUID-vkCmdDispatch-None-08610)VUID-vkCmdDispatch-None-08610  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the sampler **must** not be used with any of the SPIR-V `OpImageSample*` or `OpImageSparseSample*` instructions with `ImplicitLod`, `Dref` or `Proj` in their name
- [](#VUID-vkCmdDispatch-None-08611)VUID-vkCmdDispatch-None-08611  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the sampler **must** not be used with any of the SPIR-V `OpImageSample*` or `OpImageSparseSample*` instructions that includes a LOD bias or any offset values
- [](#VUID-vkCmdDispatch-None-06479)VUID-vkCmdDispatch-None-06479  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is sampled with [depth comparison](#textures-depth-compare-operation), the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_DEPTH_COMPARISON_BIT`
- [](#VUID-vkCmdDispatch-None-02691)VUID-vkCmdDispatch-None-02691  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is accessed using atomic operations as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT`
- [](#VUID-vkCmdDispatch-None-07888)VUID-vkCmdDispatch-None-07888  
  If a `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor is accessed using atomic operations as a result of this command, then the storage texel buffer’s [format features](#resources-buffer-view-format-features) **must** contain `VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT`
- [](#VUID-vkCmdDispatch-None-02692)VUID-vkCmdDispatch-None-02692  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is sampled with `VK_FILTER_CUBIC_EXT` as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT`
- [](#VUID-vkCmdDispatch-None-02693)VUID-vkCmdDispatch-None-02693  
  If the [VK\_EXT\_filter\_cubic](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_filter_cubic.html) extension is not enabled and any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is sampled with `VK_FILTER_CUBIC_EXT` as a result of this command, it **must** not have a [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html) of `VK_IMAGE_VIEW_TYPE_3D`, `VK_IMAGE_VIEW_TYPE_CUBE`, or `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY`
- [](#VUID-vkCmdDispatch-filterCubic-02694)VUID-vkCmdDispatch-filterCubic-02694  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with `VK_FILTER_CUBIC_EXT` as a result of this command **must** have a [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html) and format that supports cubic filtering, as specified by [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)::`filterCubic` returned by [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
- [](#VUID-vkCmdDispatch-filterCubicMinmax-02695)VUID-vkCmdDispatch-filterCubicMinmax-02695  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with `VK_FILTER_CUBIC_EXT` with a reduction mode of either `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` as a result of this command **must** have a [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html) and format that supports cubic filtering together with minmax filtering, as specified by [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)::`filterCubicMinmax` returned by [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
- [](#VUID-vkCmdDispatch-cubicRangeClamp-09212)VUID-vkCmdDispatch-cubicRangeClamp-09212  
  If the [`cubicRangeClamp`](#features-cubicRangeClamp) feature is not enabled, then any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with `VK_FILTER_CUBIC_EXT` as a result of this command **must** not have a [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeCreateInfo.html)::`reductionMode` equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM`
- [](#VUID-vkCmdDispatch-reductionMode-09213)VUID-vkCmdDispatch-reductionMode-09213  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with a [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeCreateInfo.html)::`reductionMode` equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM` as a result of this command **must** sample with `VK_FILTER_CUBIC_EXT`
- [](#VUID-vkCmdDispatch-selectableCubicWeights-09214)VUID-vkCmdDispatch-selectableCubicWeights-09214  
  If the [`selectableCubicWeights`](#features-selectableCubicWeights) feature is not enabled, then any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with `VK_FILTER_CUBIC_EXT` as a result of this command **must** have [VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html)::`cubicWeights` equal to `VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM`
- [](#VUID-vkCmdDispatch-flags-02696)VUID-vkCmdDispatch-flags-02696  
  Any [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) created with a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`flags` containing `VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV` sampled as a result of this command **must** only be sampled using a [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerAddressMode.html) of `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE`
- [](#VUID-vkCmdDispatch-OpTypeImage-07027)VUID-vkCmdDispatch-OpTypeImage-07027  
  For any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being written as a storage image where the image format field of the `OpTypeImage` is `Unknown`, the view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT`
- [](#VUID-vkCmdDispatch-OpTypeImage-07028)VUID-vkCmdDispatch-OpTypeImage-07028  
  For any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being read as a storage image where the image format field of the `OpTypeImage` is `Unknown`, the view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT`
- [](#VUID-vkCmdDispatch-OpTypeImage-07029)VUID-vkCmdDispatch-OpTypeImage-07029  
  For any [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) being written as a storage texel buffer where the image format field of the `OpTypeImage` is `Unknown`, the view’s [buffer features](#VkFormatProperties3) **must** contain `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT`
- [](#VUID-vkCmdDispatch-OpTypeImage-07030)VUID-vkCmdDispatch-OpTypeImage-07030  
  Any [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) being read as a storage texel buffer where the image format field of the `OpTypeImage` is `Unknown` then the view’s [buffer features](#VkFormatProperties3) **must** contain `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT`
- [](#VUID-vkCmdDispatch-None-08600)VUID-vkCmdDispatch-None-08600  
  For each set *n* that is statically used by [a bound shader](#shaders-binding), a descriptor set **must** have been bound to *n* at the same pipeline bind point, with a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) that is compatible for set *n*, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) used to create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) or the [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) array used to create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) , as described in [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)
- [](#VUID-vkCmdDispatch-None-08601)VUID-vkCmdDispatch-None-08601  
  For each push constant that is statically used by [a bound shader](#shaders-binding), a push constant value **must** have been set for the same pipeline bind point, with a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) that is compatible for push constants, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) used to create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) or the [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) array used to create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) , as described in [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)
- [](#VUID-vkCmdDispatch-None-10068)VUID-vkCmdDispatch-None-10068  
  For each array of resources that is used by [a bound shader](#shaders-binding), the indices used to access members of the array **must** be less than the descriptor count for the identified binding in the descriptor sets used by this command
- [](#VUID-vkCmdDispatch-maintenance4-08602)VUID-vkCmdDispatch-maintenance4-08602  
  If the [`maintenance4`](#features-maintenance4) feature is not enabled, then for each push constant that is statically used by [a bound shader](#shaders-binding), a push constant value **must** have been set for the same pipeline bind point, with a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) that is compatible for push constants, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) used to create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) or the [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) and [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantRange.html) arrays used to create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) , as described in [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)
- [](#VUID-vkCmdDispatch-None-08114)VUID-vkCmdDispatch-None-08114  
  Descriptors in each bound descriptor set, specified via [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets.html), **must** be valid as described by [descriptor validity](#descriptor-validity) if they are statically used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point used by this command and the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) was not created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdDispatch-None-08115)VUID-vkCmdDispatch-None-08115  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point were specified via [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets.html), the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) **must** have been created without `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdDispatch-None-08116)VUID-vkCmdDispatch-None-08116  
  Descriptors in bound descriptor buffers, specified via [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), **must** be valid if they are dynamically used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point used by this command and the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) was created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdDispatch-None-08604)VUID-vkCmdDispatch-None-08604  
  Descriptors in bound descriptor buffers, specified via [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), **must** be valid if they are dynamically used by any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) bound to a stage corresponding to the pipeline bind point used by this command
- [](#VUID-vkCmdDispatch-None-08117)VUID-vkCmdDispatch-None-08117  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point were specified via [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) **must** have been created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdDispatch-None-08119)VUID-vkCmdDispatch-None-08119  
  If a descriptor is dynamically used with a [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, the descriptor memory **must** be resident
- [](#VUID-vkCmdDispatch-None-08605)VUID-vkCmdDispatch-None-08605  
  If a descriptor is dynamically used with a [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) created with a `VkDescriptorSetLayout` that was created with `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, the descriptor memory **must** be resident
- [](#VUID-vkCmdDispatch-None-08606)VUID-vkCmdDispatch-None-08606  
  If the [`shaderObject`](#features-shaderObject) feature is not enabled, a valid pipeline **must** be bound to the pipeline bind point used by this command
- [](#VUID-vkCmdDispatch-None-08608)VUID-vkCmdDispatch-None-08608  
  If a pipeline is bound to the pipeline bind point used by this command, there **must** not have been any calls to dynamic state setting commands for any state specified statically in the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) object bound to the pipeline bind point used by this command, since that pipeline was bound
- [](#VUID-vkCmdDispatch-uniformBuffers-06935)VUID-vkCmdDispatch-uniformBuffers-06935  
  If any stage of the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) object bound to the pipeline bind point used by this command accesses a uniform buffer, and that stage was created without enabling either `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS` or `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2` for `uniformBuffers`, and the [`robustBufferAccess`](#features-robustBufferAccess) feature is not enabled, that stage **must** not access values outside of the range of the buffer as specified in the descriptor set bound to the same pipeline bind point
- [](#VUID-vkCmdDispatch-None-08612)VUID-vkCmdDispatch-None-08612  
  If the [`robustBufferAccess`](#features-robustBufferAccess) feature is not enabled, and any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) bound to a stage corresponding to the pipeline bind point used by this command accesses a uniform buffer, it **must** not access values outside of the range of the buffer as specified in the descriptor set bound to the same pipeline bind point
- [](#VUID-vkCmdDispatch-storageBuffers-06936)VUID-vkCmdDispatch-storageBuffers-06936  
  If any stage of the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) object bound to the pipeline bind point used by this command accesses a storage buffer, and that stage was created without enabling either `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS` or `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2` for `storageBuffers`, and the [`robustBufferAccess`](#features-robustBufferAccess) feature is not enabled, that stage **must** not access values outside of the range of the buffer as specified in the descriptor set bound to the same pipeline bind point
- [](#VUID-vkCmdDispatch-None-08613)VUID-vkCmdDispatch-None-08613  
  If the [`robustBufferAccess`](#features-robustBufferAccess) feature is not enabled, and any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) bound to a stage corresponding to the pipeline bind point used by this command accesses a storage buffer, it **must** not access values outside of the range of the buffer as specified in the descriptor set bound to the same pipeline bind point
- [](#VUID-vkCmdDispatch-commandBuffer-02707)VUID-vkCmdDispatch-commandBuffer-02707  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, any resource accessed by [bound shaders](#shaders-binding) **must** not be a protected resource
- [](#VUID-vkCmdDispatch-viewType-07752)VUID-vkCmdDispatch-viewType-07752  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is accessed as a result of this command, then the image view’s `viewType` **must** match the `Dim` operand of the `OpTypeImage` as described in [\[spirvenv-image-dimensions\]](#spirvenv-image-dimensions)
- [](#VUID-vkCmdDispatch-format-07753)VUID-vkCmdDispatch-format-07753  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) or [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) is accessed as a result of this command, then the [numeric type](#formats-numericformat) of the view’s `format` and the `Sampled` `Type` operand of the `OpTypeImage` **must** match
- [](#VUID-vkCmdDispatch-OpImageWrite-08795)VUID-vkCmdDispatch-OpImageWrite-08795  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) created with a format other than `VK_FORMAT_A8_UNORM` is accessed using `OpImageWrite` as a result of this command, then the `Type` of the `Texel` operand of that instruction **must** have at least as many components as the image view’s format
- [](#VUID-vkCmdDispatch-OpImageWrite-08796)VUID-vkCmdDispatch-OpImageWrite-08796  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) created with the format `VK_FORMAT_A8_UNORM` is accessed using `OpImageWrite` as a result of this command, then the `Type` of the `Texel` operand of that instruction **must** have four components
- [](#VUID-vkCmdDispatch-OpImageWrite-04469)VUID-vkCmdDispatch-OpImageWrite-04469  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) is accessed using `OpImageWrite` as a result of this command, then the `Type` of the `Texel` operand of that instruction **must** have at least as many components as the buffer view’s format
- [](#VUID-vkCmdDispatch-SampledType-04470)VUID-vkCmdDispatch-SampledType-04470  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that has a 64-bit component width is accessed as a result of this command, the `SampledType` of the `OpTypeImage` operand of that instruction **must** have a `Width` of 64
- [](#VUID-vkCmdDispatch-SampledType-04471)VUID-vkCmdDispatch-SampledType-04471  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that has a component width less than 64-bit is accessed as a result of this command, the `SampledType` of the `OpTypeImage` operand of that instruction **must** have a `Width` of 32
- [](#VUID-vkCmdDispatch-SampledType-04472)VUID-vkCmdDispatch-SampledType-04472  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that has a 64-bit component width is accessed as a result of this command, the `SampledType` of the `OpTypeImage` operand of that instruction **must** have a `Width` of 64
- [](#VUID-vkCmdDispatch-SampledType-04473)VUID-vkCmdDispatch-SampledType-04473  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that has a component width less than 64-bit is accessed as a result of this command, the `SampledType` of the `OpTypeImage` operand of that instruction **must** have a `Width` of 32
- [](#VUID-vkCmdDispatch-sparseImageInt64Atomics-04474)VUID-vkCmdDispatch-sparseImageInt64Atomics-04474  
  If the [`sparseImageInt64Atomics`](#features-sparseImageInt64Atomics) feature is not enabled, [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) objects created with the `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` flag **must** not be accessed by atomic instructions through an `OpTypeImage` with a `SampledType` with a `Width` of 64 by this command
- [](#VUID-vkCmdDispatch-sparseImageInt64Atomics-04475)VUID-vkCmdDispatch-sparseImageInt64Atomics-04475  
  If the [`sparseImageInt64Atomics`](#features-sparseImageInt64Atomics) feature is not enabled, [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) objects created with the `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT` flag **must** not be accessed by atomic instructions through an `OpTypeImage` with a `SampledType` with a `Width` of 64 by this command
- [](#VUID-vkCmdDispatch-OpImageWeightedSampleQCOM-06971)VUID-vkCmdDispatch-OpImageWeightedSampleQCOM-06971  
  If `OpImageWeightedSampleQCOM` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_WEIGHT_SAMPLED_IMAGE_BIT_QCOM`
- [](#VUID-vkCmdDispatch-OpImageWeightedSampleQCOM-06972)VUID-vkCmdDispatch-OpImageWeightedSampleQCOM-06972  
  If `OpImageWeightedSampleQCOM` uses a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a sample weight image as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_WEIGHT_IMAGE_BIT_QCOM`
- [](#VUID-vkCmdDispatch-OpImageBoxFilterQCOM-06973)VUID-vkCmdDispatch-OpImageBoxFilterQCOM-06973  
  If `OpImageBoxFilterQCOM` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_BOX_FILTER_SAMPLED_BIT_QCOM`
- [](#VUID-vkCmdDispatch-OpImageBlockMatchSSDQCOM-06974)VUID-vkCmdDispatch-OpImageBlockMatchSSDQCOM-06974  
  If `OpImageBlockMatchSSDQCOM` is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`
- [](#VUID-vkCmdDispatch-OpImageBlockMatchSADQCOM-06975)VUID-vkCmdDispatch-OpImageBlockMatchSADQCOM-06975  
  If `OpImageBlockMatchSADQCOM` is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`
- [](#VUID-vkCmdDispatch-OpImageBlockMatchSADQCOM-06976)VUID-vkCmdDispatch-OpImageBlockMatchSADQCOM-06976  
  If `OpImageBlockMatchSADQCOM` or OpImageBlockMatchSSDQCOM is used to read from a reference image as result of this command, then the specified reference coordinates **must** not fail [integer texel coordinate validation](#textures-integer-coordinate-validation)
- [](#VUID-vkCmdDispatch-OpImageWeightedSampleQCOM-06977)VUID-vkCmdDispatch-OpImageWeightedSampleQCOM-06977  
  If `OpImageWeightedSampleQCOM`, `OpImageBoxFilterQCOM`, `OpImageBlockMatchWindowSSDQCOM`, `OpImageBlockMatchWindowSADQCOM`, `OpImageBlockMatchGatherSSDQCOM`, `OpImageBlockMatchGatherSADQCOM`, `OpImageBlockMatchSSDQCOM`, or `OpImageBlockMatchSADQCOM` uses a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) as a result of this command, then the sampler **must** have been created with `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`
- [](#VUID-vkCmdDispatch-OpImageWeightedSampleQCOM-06978)VUID-vkCmdDispatch-OpImageWeightedSampleQCOM-06978  
  If any command other than `OpImageWeightedSampleQCOM`, `OpImageBoxFilterQCOM`, `OpImageBlockMatchWindowSSDQCOM`, `OpImageBlockMatchWindowSADQCOM`, `OpImageBlockMatchGatherSSDQCOM`, `OpImageBlockMatchGatherSADQCOM`, `OpImageBlockMatchSSDQCOM`, or `OpImageBlockMatchSADQCOM` uses a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) as a result of this command, then the sampler **must** not have been created with `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`
- [](#VUID-vkCmdDispatch-OpImageBlockMatchWindow-09215)VUID-vkCmdDispatch-OpImageBlockMatchWindow-09215  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM` instruction is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`
- [](#VUID-vkCmdDispatch-OpImageBlockMatchWindow-09216)VUID-vkCmdDispatch-OpImageBlockMatchWindow-09216  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM` instruction is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s format **must** be a single-component format
- [](#VUID-vkCmdDispatch-OpImageBlockMatchWindow-09217)VUID-vkCmdDispatch-OpImageBlockMatchWindow-09217  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM` read from a reference image as result of this command, then the specified reference coordinates **must** not fail [integer texel coordinate validation](#textures-integer-coordinate-validation)
- [](#VUID-vkCmdDispatch-None-07288)VUID-vkCmdDispatch-None-07288  
  Any shader invocation executed by this command **must** [terminate](#shaders-termination)
- [](#VUID-vkCmdDispatch-None-09600)VUID-vkCmdDispatch-None-09600  
  If a descriptor with type equal to any of `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM`, `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`, `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` is accessed as a result of this command, all image subresources identified by that descriptor **must** be in the image layout identified when the descriptor was written
- [](#VUID-vkCmdDispatch-commandBuffer-10746)VUID-vkCmdDispatch-commandBuffer-10746  
  The `VkDeviceMemory` object allocated from a `VkMemoryHeap` with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property that is bound to a resource accessed as a result of this command **must** be the active bound [bound tile memory object](#memory-bind-tile-memory) in `commandBuffer`
- [](#VUID-vkCmdDispatch-None-10678)VUID-vkCmdDispatch-None-10678  
  If this command is recorded inside a [tile shading render pass](#renderpass-tile-shading) instance, the stages corresponding to the pipeline bind point used by this command **must** only include `VK_SHADER_STAGE_VERTEX_BIT`, `VK_SHADER_STAGE_FRAGMENT_BIT`, and/or `VK_SHADER_STAGE_COMPUTE_BIT`
- [](#VUID-vkCmdDispatch-None-10679)VUID-vkCmdDispatch-None-10679  
  If this command is recorded where [per-tile execution model](#renderpass-per-tile-execution-model) is enabled, there **must** be no access to any image while the image was be transitioned to the `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` layout
- [](#VUID-vkCmdDispatch-pDescription-09900)VUID-vkCmdDispatch-pDescription-09900  
  If a `VK_DESCRIPTOR_TYPE_TENSOR_ARM` descriptor is accessed as a result of this command, then the underlying [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) object **must** have been created with a [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)::`pDescription` whose `usage` member contained `VK_TENSOR_USAGE_SHADER_BIT_ARM`
- [](#VUID-vkCmdDispatch-dimensionCount-09905)VUID-vkCmdDispatch-dimensionCount-09905  
  If a `VK_DESCRIPTOR_TYPE_TENSOR_ARM` descriptor is accessed as a result of this command, then the `Rank` of the `OpTypeTensorARM` of the tensor resource variable **must** be equal to the `dimensionCount` provided via [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)::`pDescription` when creating the underlying [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) object
- [](#VUID-vkCmdDispatch-OpTypeTensorARM-09906)VUID-vkCmdDispatch-OpTypeTensorARM-09906  
  If a `VK_DESCRIPTOR_TYPE_TENSOR_ARM` descriptor is accessed as a result of this command, then the element type of the `OpTypeTensorARM` of the tensor resource variable **must** be [compatible](#spirvenv-tensor-formats) with the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of the [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html) used for the access
- [](#VUID-vkCmdDispatch-None-10672)VUID-vkCmdDispatch-None-10672  
  If the [per-tile execution model](#renderpass-per-tile-execution-model) is not enabled, this command **must** be called outside of a render pass instance
- [](#VUID-vkCmdDispatch-aspectMask-10673)VUID-vkCmdDispatch-aspectMask-10673  
  If this command is recorded where [per-tile execution model](#renderpass-per-tile-execution-model) is enabled, and if the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) object bound to the pipeline bind point used by this command writes to a variable of storage class `Storage` `Class` `TileAttachmentQCOM`, the corresponding [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) using **must** not have been created with an `aspectMask` that contains `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-vkCmdDispatch-None-10674)VUID-vkCmdDispatch-None-10674  
  If the [per-tile execution model](#renderpass-per-tile-execution-model) is enabled, the [tileShadingPerTileDispatch](#features-tileShadingPerTileDispatch) feature **must** be enabled
- [](#VUID-vkCmdDispatch-None-10675)VUID-vkCmdDispatch-None-10675  
  Memory backing image subresources used as [tile attachments](#renderpass-tile-shading-attachment-access) in the current render pass **must** not be written in any way other than as a tile attachment by this command
- [](#VUID-vkCmdDispatch-None-10676)VUID-vkCmdDispatch-None-10676  
  If any recorded command in the current subpass will write to an image subresource as a [tile attachment](#renderpass-tile-shading-attachment-access), this command **must** not read from the memory backing that image subresource in any other way than as a tile attachment
- [](#VUID-vkCmdDispatch-None-10743)VUID-vkCmdDispatch-None-10743  
  If there is no bound compute pipeline, a valid `VkShaderEXT` **must** be bound to the `VK_SHADER_STAGE_COMPUTE_BIT` stage

<!--THE END-->

- [](#VUID-vkCmdDispatch-commandBuffer-02712)VUID-vkCmdDispatch-commandBuffer-02712  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, any resource written to by the `VkPipeline` object bound to the pipeline bind point used by this command **must** not be an unprotected resource
- [](#VUID-vkCmdDispatch-commandBuffer-02713)VUID-vkCmdDispatch-commandBuffer-02713  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, pipeline stages other than the framebuffer-space and compute stages in the `VkPipeline` object bound to the pipeline bind point used by this command **must** not write to any resource
- [](#VUID-vkCmdDispatch-commandBuffer-04617)VUID-vkCmdDispatch-commandBuffer-04617  
  If any of the shader stages of the `VkPipeline` bound to the pipeline bind point used by this command uses the [`RayQueryKHR`](#spirvenv-capabilities-table-RayQueryKHR) capability, then `commandBuffer` **must** not be a protected command buffer
- [](#VUID-vkCmdDispatch-groupCountX-00386)VUID-vkCmdDispatch-groupCountX-00386  
  `groupCountX` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[0]
- [](#VUID-vkCmdDispatch-groupCountY-00387)VUID-vkCmdDispatch-groupCountY-00387  
  `groupCountY` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[1]
- [](#VUID-vkCmdDispatch-groupCountZ-00388)VUID-vkCmdDispatch-groupCountZ-00388  
  `groupCountZ` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[2]

Valid Usage (Implicit)

- [](#VUID-vkCmdDispatch-commandBuffer-parameter)VUID-vkCmdDispatch-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdDispatch-commandBuffer-recording)VUID-vkCmdDispatch-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdDispatch-commandBuffer-cmdpool)VUID-vkCmdDispatch-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT operations
- [](#VUID-vkCmdDispatch-videocoding)VUID-vkCmdDispatch-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Outside

VK\_QUEUE\_COMPUTE\_BIT

Action

Conditional Rendering

vkCmdDispatch is affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdDispatch).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0