# vkCmdTraceRaysIndirectKHR(3) Manual Page

## Name

vkCmdTraceRaysIndirectKHR - Initialize an indirect ray tracing dispatch



## [](#_c_specification)C Specification

To dispatch ray tracing, with some parameters sourced on the device, use:

```c++
// Provided by VK_KHR_ray_tracing_pipeline
void vkCmdTraceRaysIndirectKHR(
    VkCommandBuffer                             commandBuffer,
    const VkStridedDeviceAddressRegionKHR*      pRaygenShaderBindingTable,
    const VkStridedDeviceAddressRegionKHR*      pMissShaderBindingTable,
    const VkStridedDeviceAddressRegionKHR*      pHitShaderBindingTable,
    const VkStridedDeviceAddressRegionKHR*      pCallableShaderBindingTable,
    VkDeviceAddress                             indirectDeviceAddress);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pRaygenShaderBindingTable` is a [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html) that holds the shader binding table data for the ray generation shader stage.
- `pMissShaderBindingTable` is a [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html) that holds the shader binding table data for the miss shader stage.
- `pHitShaderBindingTable` is a [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html) that holds the shader binding table data for the hit shader stage.
- `pCallableShaderBindingTable` is a [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html) that holds the shader binding table data for the callable shader stage.
- `indirectDeviceAddress` is a buffer device address which is a pointer to a [VkTraceRaysIndirectCommandKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTraceRaysIndirectCommandKHR.html) structure containing the trace ray parameters.

## [](#_description)Description

`vkCmdTraceRaysIndirectKHR` behaves similarly to [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysKHR.html) except that the ray trace query dimensions are read by the device from `indirectDeviceAddress` during execution.

Valid Usage

- [](#VUID-vkCmdTraceRaysIndirectKHR-magFilter-04553)VUID-vkCmdTraceRaysIndirectKHR-magFilter-04553  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `magFilter` or `minFilter` equal to `VK_FILTER_LINEAR`, `reductionMode` equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`, and `compareEnable` equal to `VK_FALSE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-magFilter-09598)VUID-vkCmdTraceRaysIndirectKHR-magFilter-09598  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `magFilter` or `minFilter` equal to `VK_FILTER_LINEAR` and `reductionMode` equal to either `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-mipmapMode-04770)VUID-vkCmdTraceRaysIndirectKHR-mipmapMode-04770  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `mipmapMode` equal to `VK_SAMPLER_MIPMAP_MODE_LINEAR`, `reductionMode` equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`, and `compareEnable` equal to `VK_FALSE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-mipmapMode-09599)VUID-vkCmdTraceRaysIndirectKHR-mipmapMode-09599  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `mipmapMode` equal to `VK_SAMPLER_MIPMAP_MODE_LINEAR` and `reductionMode` equal to either `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-unnormalizedCoordinates-09635)VUID-vkCmdTraceRaysIndirectKHR-unnormalizedCoordinates-09635  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s `levelCount` and `layerCount` **must** be 1
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08609)VUID-vkCmdTraceRaysIndirectKHR-None-08609  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s `viewType` **must** be `VK_IMAGE_VIEW_TYPE_1D` or `VK_IMAGE_VIEW_TYPE_2D`
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08610)VUID-vkCmdTraceRaysIndirectKHR-None-08610  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the sampler **must** not be used with any of the SPIR-V `OpImageSample*` or `OpImageSparseSample*` instructions with `ImplicitLod`, `Dref` or `Proj` in their name
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08611)VUID-vkCmdTraceRaysIndirectKHR-None-08611  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the sampler **must** not be used with any of the SPIR-V `OpImageSample*` or `OpImageSparseSample*` instructions that includes a LOD bias or any offset values
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-06479)VUID-vkCmdTraceRaysIndirectKHR-None-06479  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is sampled with [depth comparison](#textures-depth-compare-operation), the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_DEPTH_COMPARISON_BIT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-02691)VUID-vkCmdTraceRaysIndirectKHR-None-02691  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is accessed using atomic operations as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-07888)VUID-vkCmdTraceRaysIndirectKHR-None-07888  
  If a `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor is accessed using atomic operations as a result of this command, then the storage texel buffer’s [format features](#resources-buffer-view-format-features) **must** contain `VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-02692)VUID-vkCmdTraceRaysIndirectKHR-None-02692  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is sampled with `VK_FILTER_CUBIC_EXT` as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-02693)VUID-vkCmdTraceRaysIndirectKHR-None-02693  
  If the [VK\_EXT\_filter\_cubic](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_filter_cubic.html) extension is not enabled and any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is sampled with `VK_FILTER_CUBIC_EXT` as a result of this command, it **must** not have a [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html) of `VK_IMAGE_VIEW_TYPE_3D`, `VK_IMAGE_VIEW_TYPE_CUBE`, or `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY`
- [](#VUID-vkCmdTraceRaysIndirectKHR-filterCubic-02694)VUID-vkCmdTraceRaysIndirectKHR-filterCubic-02694  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with `VK_FILTER_CUBIC_EXT` as a result of this command **must** have a [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html) and format that supports cubic filtering, as specified by [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)::`filterCubic` returned by [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
- [](#VUID-vkCmdTraceRaysIndirectKHR-filterCubicMinmax-02695)VUID-vkCmdTraceRaysIndirectKHR-filterCubicMinmax-02695  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with `VK_FILTER_CUBIC_EXT` with a reduction mode of either `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` as a result of this command **must** have a [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html) and format that supports cubic filtering together with minmax filtering, as specified by [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)::`filterCubicMinmax` returned by [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
- [](#VUID-vkCmdTraceRaysIndirectKHR-cubicRangeClamp-09212)VUID-vkCmdTraceRaysIndirectKHR-cubicRangeClamp-09212  
  If the [`cubicRangeClamp`](#features-cubicRangeClamp) feature is not enabled, then any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with `VK_FILTER_CUBIC_EXT` as a result of this command **must** not have a [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeCreateInfo.html)::`reductionMode` equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM`
- [](#VUID-vkCmdTraceRaysIndirectKHR-reductionMode-09213)VUID-vkCmdTraceRaysIndirectKHR-reductionMode-09213  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with a [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeCreateInfo.html)::`reductionMode` equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM` as a result of this command **must** sample with `VK_FILTER_CUBIC_EXT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-selectableCubicWeights-09214)VUID-vkCmdTraceRaysIndirectKHR-selectableCubicWeights-09214  
  If the [`selectableCubicWeights`](#features-selectableCubicWeights) feature is not enabled, then any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with `VK_FILTER_CUBIC_EXT` as a result of this command **must** have [VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html)::`cubicWeights` equal to `VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM`
- [](#VUID-vkCmdTraceRaysIndirectKHR-flags-02696)VUID-vkCmdTraceRaysIndirectKHR-flags-02696  
  Any [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) created with a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`flags` containing `VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV` sampled as a result of this command **must** only be sampled using a [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerAddressMode.html) of `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE`
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpTypeImage-07027)VUID-vkCmdTraceRaysIndirectKHR-OpTypeImage-07027  
  For any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being written as a storage image where the image format field of the `OpTypeImage` is `Unknown`, the view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpTypeImage-07028)VUID-vkCmdTraceRaysIndirectKHR-OpTypeImage-07028  
  For any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being read as a storage image where the image format field of the `OpTypeImage` is `Unknown`, the view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpTypeImage-07029)VUID-vkCmdTraceRaysIndirectKHR-OpTypeImage-07029  
  For any [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) being written as a storage texel buffer where the image format field of the `OpTypeImage` is `Unknown`, the view’s [buffer features](#VkFormatProperties3) **must** contain `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpTypeImage-07030)VUID-vkCmdTraceRaysIndirectKHR-OpTypeImage-07030  
  Any [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) being read as a storage texel buffer where the image format field of the `OpTypeImage` is `Unknown` then the view’s [buffer features](#VkFormatProperties3) **must** contain `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08600)VUID-vkCmdTraceRaysIndirectKHR-None-08600  
  For each set *n* that is statically used by [a bound shader](#shaders-binding), a descriptor set **must** have been bound to *n* at the same pipeline bind point, with a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) that is compatible for set *n*, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) used to create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) or the [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) array used to create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) , as described in [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08601)VUID-vkCmdTraceRaysIndirectKHR-None-08601  
  For each push constant that is statically used by [a bound shader](#shaders-binding), a push constant value **must** have been set for the same pipeline bind point, with a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) that is compatible for push constants, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) used to create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) or the [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) array used to create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) , as described in [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-10068)VUID-vkCmdTraceRaysIndirectKHR-None-10068  
  For each array of resources that is used by [a bound shader](#shaders-binding), the indices used to access members of the array **must** be less than the descriptor count for the identified binding in the descriptor sets used by this command
- [](#VUID-vkCmdTraceRaysIndirectKHR-maintenance4-08602)VUID-vkCmdTraceRaysIndirectKHR-maintenance4-08602  
  If the [`maintenance4`](#features-maintenance4) feature is not enabled, then for each push constant that is statically used by [a bound shader](#shaders-binding), a push constant value **must** have been set for the same pipeline bind point, with a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) that is compatible for push constants, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) used to create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) or the [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) and [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantRange.html) arrays used to create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) , as described in [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08114)VUID-vkCmdTraceRaysIndirectKHR-None-08114  
  Descriptors in each bound descriptor set, specified via [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets.html), **must** be valid as described by [descriptor validity](#descriptor-validity) if they are statically used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point used by this command and the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) was not created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08115)VUID-vkCmdTraceRaysIndirectKHR-None-08115  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point were specified via [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets.html), the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) **must** have been created without `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08116)VUID-vkCmdTraceRaysIndirectKHR-None-08116  
  Descriptors in bound descriptor buffers, specified via [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), **must** be valid if they are dynamically used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point used by this command and the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) was created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08604)VUID-vkCmdTraceRaysIndirectKHR-None-08604  
  Descriptors in bound descriptor buffers, specified via [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), **must** be valid if they are dynamically used by any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) bound to a stage corresponding to the pipeline bind point used by this command
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08117)VUID-vkCmdTraceRaysIndirectKHR-None-08117  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point were specified via [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) **must** have been created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08119)VUID-vkCmdTraceRaysIndirectKHR-None-08119  
  If a descriptor is dynamically used with a [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, the descriptor memory **must** be resident
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08605)VUID-vkCmdTraceRaysIndirectKHR-None-08605  
  If a descriptor is dynamically used with a [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) created with a `VkDescriptorSetLayout` that was created with `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, the descriptor memory **must** be resident
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08606)VUID-vkCmdTraceRaysIndirectKHR-None-08606  
  If the [`shaderObject`](#features-shaderObject) feature is not enabled, a valid pipeline **must** be bound to the pipeline bind point used by this command
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08608)VUID-vkCmdTraceRaysIndirectKHR-None-08608  
  If a pipeline is bound to the pipeline bind point used by this command, there **must** not have been any calls to dynamic state setting commands for any state specified statically in the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) object bound to the pipeline bind point used by this command, since that pipeline was bound
- [](#VUID-vkCmdTraceRaysIndirectKHR-uniformBuffers-06935)VUID-vkCmdTraceRaysIndirectKHR-uniformBuffers-06935  
  If any stage of the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) object bound to the pipeline bind point used by this command accesses a uniform buffer, and that stage was created without enabling either `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS` or `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2` for `uniformBuffers`, and the [`robustBufferAccess`](#features-robustBufferAccess) feature is not enabled, that stage **must** not access values outside of the range of the buffer as specified in the descriptor set bound to the same pipeline bind point
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08612)VUID-vkCmdTraceRaysIndirectKHR-None-08612  
  If the [`robustBufferAccess`](#features-robustBufferAccess) feature is not enabled, and any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) bound to a stage corresponding to the pipeline bind point used by this command accesses a uniform buffer, it **must** not access values outside of the range of the buffer as specified in the descriptor set bound to the same pipeline bind point
- [](#VUID-vkCmdTraceRaysIndirectKHR-storageBuffers-06936)VUID-vkCmdTraceRaysIndirectKHR-storageBuffers-06936  
  If any stage of the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) object bound to the pipeline bind point used by this command accesses a storage buffer, and that stage was created without enabling either `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS` or `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2` for `storageBuffers`, and the [`robustBufferAccess`](#features-robustBufferAccess) feature is not enabled, that stage **must** not access values outside of the range of the buffer as specified in the descriptor set bound to the same pipeline bind point
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-08613)VUID-vkCmdTraceRaysIndirectKHR-None-08613  
  If the [`robustBufferAccess`](#features-robustBufferAccess) feature is not enabled, and any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) bound to a stage corresponding to the pipeline bind point used by this command accesses a storage buffer, it **must** not access values outside of the range of the buffer as specified in the descriptor set bound to the same pipeline bind point
- [](#VUID-vkCmdTraceRaysIndirectKHR-commandBuffer-02707)VUID-vkCmdTraceRaysIndirectKHR-commandBuffer-02707  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, any resource accessed by [bound shaders](#shaders-binding) **must** not be a protected resource
- [](#VUID-vkCmdTraceRaysIndirectKHR-viewType-07752)VUID-vkCmdTraceRaysIndirectKHR-viewType-07752  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is accessed as a result of this command, then the image view’s `viewType` **must** match the `Dim` operand of the `OpTypeImage` as described in [\[spirvenv-image-dimensions\]](#spirvenv-image-dimensions)
- [](#VUID-vkCmdTraceRaysIndirectKHR-format-07753)VUID-vkCmdTraceRaysIndirectKHR-format-07753  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) or [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) is accessed as a result of this command, then the [numeric type](#formats-numericformat) of the view’s `format` and the `Sampled` `Type` operand of the `OpTypeImage` **must** match
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageWrite-08795)VUID-vkCmdTraceRaysIndirectKHR-OpImageWrite-08795  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) created with a format other than `VK_FORMAT_A8_UNORM` is accessed using `OpImageWrite` as a result of this command, then the `Type` of the `Texel` operand of that instruction **must** have at least as many components as the image view’s format
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageWrite-08796)VUID-vkCmdTraceRaysIndirectKHR-OpImageWrite-08796  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) created with the format `VK_FORMAT_A8_UNORM` is accessed using `OpImageWrite` as a result of this command, then the `Type` of the `Texel` operand of that instruction **must** have four components
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageWrite-04469)VUID-vkCmdTraceRaysIndirectKHR-OpImageWrite-04469  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) is accessed using `OpImageWrite` as a result of this command, then the `Type` of the `Texel` operand of that instruction **must** have at least as many components as the buffer view’s format
- [](#VUID-vkCmdTraceRaysIndirectKHR-SampledType-04470)VUID-vkCmdTraceRaysIndirectKHR-SampledType-04470  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that has a 64-bit component width is accessed as a result of this command, the `SampledType` of the `OpTypeImage` operand of that instruction **must** have a `Width` of 64
- [](#VUID-vkCmdTraceRaysIndirectKHR-SampledType-04471)VUID-vkCmdTraceRaysIndirectKHR-SampledType-04471  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that has a component width less than 64-bit is accessed as a result of this command, the `SampledType` of the `OpTypeImage` operand of that instruction **must** have a `Width` of 32
- [](#VUID-vkCmdTraceRaysIndirectKHR-SampledType-04472)VUID-vkCmdTraceRaysIndirectKHR-SampledType-04472  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that has a 64-bit component width is accessed as a result of this command, the `SampledType` of the `OpTypeImage` operand of that instruction **must** have a `Width` of 64
- [](#VUID-vkCmdTraceRaysIndirectKHR-SampledType-04473)VUID-vkCmdTraceRaysIndirectKHR-SampledType-04473  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that has a component width less than 64-bit is accessed as a result of this command, the `SampledType` of the `OpTypeImage` operand of that instruction **must** have a `Width` of 32
- [](#VUID-vkCmdTraceRaysIndirectKHR-sparseImageInt64Atomics-04474)VUID-vkCmdTraceRaysIndirectKHR-sparseImageInt64Atomics-04474  
  If the [`sparseImageInt64Atomics`](#features-sparseImageInt64Atomics) feature is not enabled, [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) objects created with the `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` flag **must** not be accessed by atomic instructions through an `OpTypeImage` with a `SampledType` with a `Width` of 64 by this command
- [](#VUID-vkCmdTraceRaysIndirectKHR-sparseImageInt64Atomics-04475)VUID-vkCmdTraceRaysIndirectKHR-sparseImageInt64Atomics-04475  
  If the [`sparseImageInt64Atomics`](#features-sparseImageInt64Atomics) feature is not enabled, [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) objects created with the `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT` flag **must** not be accessed by atomic instructions through an `OpTypeImage` with a `SampledType` with a `Width` of 64 by this command
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageWeightedSampleQCOM-06971)VUID-vkCmdTraceRaysIndirectKHR-OpImageWeightedSampleQCOM-06971  
  If `OpImageWeightedSampleQCOM` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_WEIGHT_SAMPLED_IMAGE_BIT_QCOM`
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageWeightedSampleQCOM-06972)VUID-vkCmdTraceRaysIndirectKHR-OpImageWeightedSampleQCOM-06972  
  If `OpImageWeightedSampleQCOM` uses a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a sample weight image as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_WEIGHT_IMAGE_BIT_QCOM`
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageBoxFilterQCOM-06973)VUID-vkCmdTraceRaysIndirectKHR-OpImageBoxFilterQCOM-06973  
  If `OpImageBoxFilterQCOM` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_BOX_FILTER_SAMPLED_BIT_QCOM`
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageBlockMatchSSDQCOM-06974)VUID-vkCmdTraceRaysIndirectKHR-OpImageBlockMatchSSDQCOM-06974  
  If `OpImageBlockMatchSSDQCOM` is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageBlockMatchSADQCOM-06975)VUID-vkCmdTraceRaysIndirectKHR-OpImageBlockMatchSADQCOM-06975  
  If `OpImageBlockMatchSADQCOM` is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageBlockMatchSADQCOM-06976)VUID-vkCmdTraceRaysIndirectKHR-OpImageBlockMatchSADQCOM-06976  
  If `OpImageBlockMatchSADQCOM` or OpImageBlockMatchSSDQCOM is used to read from a reference image as result of this command, then the specified reference coordinates **must** not fail [integer texel coordinate validation](#textures-integer-coordinate-validation)
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageWeightedSampleQCOM-06977)VUID-vkCmdTraceRaysIndirectKHR-OpImageWeightedSampleQCOM-06977  
  If `OpImageWeightedSampleQCOM`, `OpImageBoxFilterQCOM`, `OpImageBlockMatchWindowSSDQCOM`, `OpImageBlockMatchWindowSADQCOM`, `OpImageBlockMatchGatherSSDQCOM`, `OpImageBlockMatchGatherSADQCOM`, `OpImageBlockMatchSSDQCOM`, or `OpImageBlockMatchSADQCOM` uses a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) as a result of this command, then the sampler **must** have been created with `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageWeightedSampleQCOM-06978)VUID-vkCmdTraceRaysIndirectKHR-OpImageWeightedSampleQCOM-06978  
  If any command other than `OpImageWeightedSampleQCOM`, `OpImageBoxFilterQCOM`, `OpImageBlockMatchWindowSSDQCOM`, `OpImageBlockMatchWindowSADQCOM`, `OpImageBlockMatchGatherSSDQCOM`, `OpImageBlockMatchGatherSADQCOM`, `OpImageBlockMatchSSDQCOM`, or `OpImageBlockMatchSADQCOM` uses a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) as a result of this command, then the sampler **must** not have been created with `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageBlockMatchWindow-09215)VUID-vkCmdTraceRaysIndirectKHR-OpImageBlockMatchWindow-09215  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM` instruction is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageBlockMatchWindow-09216)VUID-vkCmdTraceRaysIndirectKHR-OpImageBlockMatchWindow-09216  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM` instruction is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s format **must** be a single-component format
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpImageBlockMatchWindow-09217)VUID-vkCmdTraceRaysIndirectKHR-OpImageBlockMatchWindow-09217  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM` read from a reference image as result of this command, then the specified reference coordinates **must** not fail [integer texel coordinate validation](#textures-integer-coordinate-validation)
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-07288)VUID-vkCmdTraceRaysIndirectKHR-None-07288  
  Any shader invocation executed by this command **must** [terminate](#shaders-termination)
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-09600)VUID-vkCmdTraceRaysIndirectKHR-None-09600  
  If a descriptor with type equal to any of `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM`, `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`, `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` is accessed as a result of this command, all image subresources identified by that descriptor **must** be in the image layout identified when the descriptor was written
- [](#VUID-vkCmdTraceRaysIndirectKHR-commandBuffer-10746)VUID-vkCmdTraceRaysIndirectKHR-commandBuffer-10746  
  The `VkDeviceMemory` object allocated from a `VkMemoryHeap` with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property that is bound to a resource accessed as a result of this command **must** be the active bound [bound tile memory object](#memory-bind-tile-memory) in `commandBuffer`
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-10678)VUID-vkCmdTraceRaysIndirectKHR-None-10678  
  If this command is recorded inside a [tile shading render pass](#renderpass-tile-shading) instance, the stages corresponding to the pipeline bind point used by this command **must** only include `VK_SHADER_STAGE_VERTEX_BIT`, `VK_SHADER_STAGE_FRAGMENT_BIT`, and/or `VK_SHADER_STAGE_COMPUTE_BIT`
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-10679)VUID-vkCmdTraceRaysIndirectKHR-None-10679  
  If this command is recorded where [per-tile execution model](#renderpass-per-tile-execution-model) is enabled, there **must** be no access to any image while the image was be transitioned to the `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` layout
- [](#VUID-vkCmdTraceRaysIndirectKHR-pDescription-09900)VUID-vkCmdTraceRaysIndirectKHR-pDescription-09900  
  If a `VK_DESCRIPTOR_TYPE_TENSOR_ARM` descriptor is accessed as a result of this command, then the underlying [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) object **must** have been created with a [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)::`pDescription` whose `usage` member contained `VK_TENSOR_USAGE_SHADER_BIT_ARM`
- [](#VUID-vkCmdTraceRaysIndirectKHR-dimensionCount-09905)VUID-vkCmdTraceRaysIndirectKHR-dimensionCount-09905  
  If a `VK_DESCRIPTOR_TYPE_TENSOR_ARM` descriptor is accessed as a result of this command, then the `Rank` of the `OpTypeTensorARM` of the tensor resource variable **must** be equal to the `dimensionCount` provided via [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)::`pDescription` when creating the underlying [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) object
- [](#VUID-vkCmdTraceRaysIndirectKHR-OpTypeTensorARM-09906)VUID-vkCmdTraceRaysIndirectKHR-OpTypeTensorARM-09906  
  If a `VK_DESCRIPTOR_TYPE_TENSOR_ARM` descriptor is accessed as a result of this command, then the element type of the `OpTypeTensorARM` of the tensor resource variable **must** be [compatible](#spirvenv-tensor-formats) with the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of the [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html) used for the access
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-03429)VUID-vkCmdTraceRaysIndirectKHR-None-03429  
  Any shader group handle referenced by this call **must** have been queried from the bound ray tracing pipeline
- [](#VUID-vkCmdTraceRaysIndirectKHR-None-09458)VUID-vkCmdTraceRaysIndirectKHR-None-09458  
  If the bound ray tracing pipeline state was created with the `VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR` dynamic state enabled then [vkCmdSetRayTracingPipelineStackSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRayTracingPipelineStackSizeKHR.html) **must** have been called in the current command buffer prior to this trace command

<!--THE END-->

- [](#VUID-vkCmdTraceRaysIndirectKHR-maxPipelineRayRecursionDepth-03679)VUID-vkCmdTraceRaysIndirectKHR-maxPipelineRayRecursionDepth-03679  
  This command **must** not cause a shader call instruction to be executed from a shader invocation with a [recursion depth](#ray-tracing-recursion-depth) greater than the value of `maxPipelineRayRecursionDepth` used to create the bound ray tracing pipeline
- [](#VUID-vkCmdTraceRaysIndirectKHR-commandBuffer-03635)VUID-vkCmdTraceRaysIndirectKHR-commandBuffer-03635  
  `commandBuffer` **must** not be a protected command buffer

<!--THE END-->

- [](#VUID-vkCmdTraceRaysIndirectKHR-size-04023)VUID-vkCmdTraceRaysIndirectKHR-size-04023  
  The `size` member of `pRayGenShaderBindingTable` **must** be equal to its `stride` member

<!--THE END-->

- [](#VUID-vkCmdTraceRaysIndirectKHR-pRayGenShaderBindingTable-03680)VUID-vkCmdTraceRaysIndirectKHR-pRayGenShaderBindingTable-03680  
  If the buffer from which `pRayGenShaderBindingTable->deviceAddress` was queried is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdTraceRaysIndirectKHR-pRayGenShaderBindingTable-03681)VUID-vkCmdTraceRaysIndirectKHR-pRayGenShaderBindingTable-03681  
  The buffer from which the `pRayGenShaderBindingTable->deviceAddress` is queried **must** have been created with the `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag
- [](#VUID-vkCmdTraceRaysIndirectKHR-pRayGenShaderBindingTable-03682)VUID-vkCmdTraceRaysIndirectKHR-pRayGenShaderBindingTable-03682  
  `pRayGenShaderBindingTable->deviceAddress` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`
- [](#VUID-vkCmdTraceRaysIndirectKHR-pMissShaderBindingTable-03683)VUID-vkCmdTraceRaysIndirectKHR-pMissShaderBindingTable-03683  
  If the buffer from which `pMissShaderBindingTable->deviceAddress` was queried is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdTraceRaysIndirectKHR-pMissShaderBindingTable-03684)VUID-vkCmdTraceRaysIndirectKHR-pMissShaderBindingTable-03684  
  The buffer from which the `pMissShaderBindingTable->deviceAddress` is queried **must** have been created with the `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag
- [](#VUID-vkCmdTraceRaysIndirectKHR-pMissShaderBindingTable-03685)VUID-vkCmdTraceRaysIndirectKHR-pMissShaderBindingTable-03685  
  `pMissShaderBindingTable->deviceAddress` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`
- [](#VUID-vkCmdTraceRaysIndirectKHR-stride-03686)VUID-vkCmdTraceRaysIndirectKHR-stride-03686  
  `pMissShaderBindingTable->stride` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupHandleAlignment`
- [](#VUID-vkCmdTraceRaysIndirectKHR-stride-04029)VUID-vkCmdTraceRaysIndirectKHR-stride-04029  
  `pMissShaderBindingTable->stride` **must** be less than or equal to `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxShaderGroupStride`
- [](#VUID-vkCmdTraceRaysIndirectKHR-pHitShaderBindingTable-03687)VUID-vkCmdTraceRaysIndirectKHR-pHitShaderBindingTable-03687  
  If the buffer from which `pHitShaderBindingTable->deviceAddress` was queried is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdTraceRaysIndirectKHR-pHitShaderBindingTable-03688)VUID-vkCmdTraceRaysIndirectKHR-pHitShaderBindingTable-03688  
  The buffer from which the `pHitShaderBindingTable->deviceAddress` is queried **must** have been created with the `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag
- [](#VUID-vkCmdTraceRaysIndirectKHR-pHitShaderBindingTable-03689)VUID-vkCmdTraceRaysIndirectKHR-pHitShaderBindingTable-03689  
  `pHitShaderBindingTable->deviceAddress` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`
- [](#VUID-vkCmdTraceRaysIndirectKHR-stride-03690)VUID-vkCmdTraceRaysIndirectKHR-stride-03690  
  `pHitShaderBindingTable->stride` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupHandleAlignment`
- [](#VUID-vkCmdTraceRaysIndirectKHR-stride-04035)VUID-vkCmdTraceRaysIndirectKHR-stride-04035  
  `pHitShaderBindingTable->stride` **must** be less than or equal to `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxShaderGroupStride`
- [](#VUID-vkCmdTraceRaysIndirectKHR-pCallableShaderBindingTable-03691)VUID-vkCmdTraceRaysIndirectKHR-pCallableShaderBindingTable-03691  
  If the buffer from which `pCallableShaderBindingTable->deviceAddress` was queried is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdTraceRaysIndirectKHR-pCallableShaderBindingTable-03692)VUID-vkCmdTraceRaysIndirectKHR-pCallableShaderBindingTable-03692  
  The buffer from which the `pCallableShaderBindingTable->deviceAddress` is queried **must** have been created with the `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag
- [](#VUID-vkCmdTraceRaysIndirectKHR-pCallableShaderBindingTable-03693)VUID-vkCmdTraceRaysIndirectKHR-pCallableShaderBindingTable-03693  
  `pCallableShaderBindingTable->deviceAddress` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`
- [](#VUID-vkCmdTraceRaysIndirectKHR-stride-03694)VUID-vkCmdTraceRaysIndirectKHR-stride-03694  
  `pCallableShaderBindingTable->stride` **must** be a multiple of `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupHandleAlignment`
- [](#VUID-vkCmdTraceRaysIndirectKHR-stride-04041)VUID-vkCmdTraceRaysIndirectKHR-stride-04041  
  `pCallableShaderBindingTable->stride` **must** be less than or equal to `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxShaderGroupStride`
- [](#VUID-vkCmdTraceRaysIndirectKHR-flags-03696)VUID-vkCmdTraceRaysIndirectKHR-flags-03696  
  If the bound ray tracing pipeline was created with `flags` that included `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`, `pHitShaderBindingTable->deviceAddress` **must** not be zero
- [](#VUID-vkCmdTraceRaysIndirectKHR-flags-03697)VUID-vkCmdTraceRaysIndirectKHR-flags-03697  
  If the bound ray tracing pipeline was created with `flags` that included `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`, `pHitShaderBindingTable->deviceAddress` **must** not be zero
- [](#VUID-vkCmdTraceRaysIndirectKHR-flags-03511)VUID-vkCmdTraceRaysIndirectKHR-flags-03511  
  If the bound ray tracing pipeline was created with `flags` that included `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`, the shader group handle identified by `pMissShaderBindingTable->deviceAddress` **must** not be zero
- [](#VUID-vkCmdTraceRaysIndirectKHR-flags-03512)VUID-vkCmdTraceRaysIndirectKHR-flags-03512  
  If the bound ray tracing pipeline was created with `flags` that included `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`, entries in the table identified by `pHitShaderBindingTable->deviceAddress` accessed as a result of this command in order to execute an any-hit shader **must** not be zero
- [](#VUID-vkCmdTraceRaysIndirectKHR-flags-03513)VUID-vkCmdTraceRaysIndirectKHR-flags-03513  
  If the bound ray tracing pipeline was created with `flags` that included `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`, entries in the table identified by `pHitShaderBindingTable->deviceAddress` accessed as a result of this command in order to execute a closest hit shader **must** not be zero
- [](#VUID-vkCmdTraceRaysIndirectKHR-flags-03514)VUID-vkCmdTraceRaysIndirectKHR-flags-03514  
  If the bound ray tracing pipeline was created with `flags` that included `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`, entries in the table identified by `pHitShaderBindingTable->deviceAddress` accessed as a result of this command in order to execute an intersection shader **must** not be zero
- [](#VUID-vkCmdTraceRaysIndirectKHR-pHitShaderBindingTable-04735)VUID-vkCmdTraceRaysIndirectKHR-pHitShaderBindingTable-04735  
  Any non-zero hit shader group entries in the table identified by `pHitShaderBindingTable->deviceAddress` accessed by this call from a geometry with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR` **must** have been created with `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR`
- [](#VUID-vkCmdTraceRaysIndirectKHR-pHitShaderBindingTable-04736)VUID-vkCmdTraceRaysIndirectKHR-pHitShaderBindingTable-04736  
  Any non-zero hit shader group entries in the table identified by `pHitShaderBindingTable->deviceAddress` accessed by this call from a geometry with a `geometryType` of `VK_GEOMETRY_TYPE_AABBS_KHR` **must** have been created with `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR`

<!--THE END-->

- [](#VUID-vkCmdTraceRaysIndirectKHR-indirectDeviceAddress-03632)VUID-vkCmdTraceRaysIndirectKHR-indirectDeviceAddress-03632  
  If the buffer from which `indirectDeviceAddress` was queried is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdTraceRaysIndirectKHR-indirectDeviceAddress-03633)VUID-vkCmdTraceRaysIndirectKHR-indirectDeviceAddress-03633  
  The buffer from which `indirectDeviceAddress` was queried **must** have been created with the `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` bit set
- [](#VUID-vkCmdTraceRaysIndirectKHR-indirectDeviceAddress-03634)VUID-vkCmdTraceRaysIndirectKHR-indirectDeviceAddress-03634  
  `indirectDeviceAddress` **must** be a multiple of `4`
- [](#VUID-vkCmdTraceRaysIndirectKHR-indirectDeviceAddress-03636)VUID-vkCmdTraceRaysIndirectKHR-indirectDeviceAddress-03636  
  All device addresses between `indirectDeviceAddress` and `indirectDeviceAddress` + `sizeof`(`VkTraceRaysIndirectCommandKHR`) - 1 **must** be in the buffer device address range of the same buffer
- [](#VUID-vkCmdTraceRaysIndirectKHR-rayTracingPipelineTraceRaysIndirect-03637)VUID-vkCmdTraceRaysIndirectKHR-rayTracingPipelineTraceRaysIndirect-03637  
  The [`rayTracingPipelineTraceRaysIndirect`](#features-rayTracingPipelineTraceRaysIndirect) feature **must** be enabled
- [](#VUID-vkCmdTraceRaysIndirectKHR-rayTracingMotionBlurPipelineTraceRaysIndirect-04951)VUID-vkCmdTraceRaysIndirectKHR-rayTracingMotionBlurPipelineTraceRaysIndirect-04951  
  If the bound ray tracing pipeline was created with `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV` `VkPhysicalDeviceRayTracingMotionBlurFeaturesNV`::`rayTracingMotionBlurPipelineTraceRaysIndirect` feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdTraceRaysIndirectKHR-commandBuffer-parameter)VUID-vkCmdTraceRaysIndirectKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdTraceRaysIndirectKHR-pRaygenShaderBindingTable-parameter)VUID-vkCmdTraceRaysIndirectKHR-pRaygenShaderBindingTable-parameter  
  `pRaygenShaderBindingTable` **must** be a valid pointer to a valid [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html) structure
- [](#VUID-vkCmdTraceRaysIndirectKHR-pMissShaderBindingTable-parameter)VUID-vkCmdTraceRaysIndirectKHR-pMissShaderBindingTable-parameter  
  `pMissShaderBindingTable` **must** be a valid pointer to a valid [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html) structure
- [](#VUID-vkCmdTraceRaysIndirectKHR-pHitShaderBindingTable-parameter)VUID-vkCmdTraceRaysIndirectKHR-pHitShaderBindingTable-parameter  
  `pHitShaderBindingTable` **must** be a valid pointer to a valid [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html) structure
- [](#VUID-vkCmdTraceRaysIndirectKHR-pCallableShaderBindingTable-parameter)VUID-vkCmdTraceRaysIndirectKHR-pCallableShaderBindingTable-parameter  
  `pCallableShaderBindingTable` **must** be a valid pointer to a valid [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html) structure
- [](#VUID-vkCmdTraceRaysIndirectKHR-commandBuffer-recording)VUID-vkCmdTraceRaysIndirectKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdTraceRaysIndirectKHR-commandBuffer-cmdpool)VUID-vkCmdTraceRaysIndirectKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdTraceRaysIndirectKHR-renderpass)VUID-vkCmdTraceRaysIndirectKHR-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdTraceRaysIndirectKHR-videocoding)VUID-vkCmdTraceRaysIndirectKHR-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Compute

Action

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdTraceRaysIndirectKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0