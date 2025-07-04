# vkCmdDispatchGraphIndirectAMDX(3) Manual Page

## Name

vkCmdDispatchGraphIndirectAMDX - Dispatch an execution graph with node and payload parameters read on the device



## [](#_c_specification)C Specification

To record an execution graph dispatch with node and payload parameters read on device, call:

```c++
// Provided by VK_AMDX_shader_enqueue
void vkCmdDispatchGraphIndirectAMDX(
    VkCommandBuffer                             commandBuffer,
    VkDeviceAddress                             scratch,
    VkDeviceSize                                scratchSize,
    const VkDispatchGraphCountInfoAMDX*         pCountInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `scratch` is the address of scratch memory to be used.
- `scratchSize` is a range in bytes of scratch memory to be used.
- `pCountInfo` is a host pointer to a [VkDispatchGraphCountInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphCountInfoAMDX.html) structure defining the nodes which will be initially executed.

## [](#_description)Description

When this command is executed, the nodes specified in `pCountInfo` are executed. Nodes executed as part of this command are not implicitly synchronized in any way against each other once they are dispatched. There are no rasterization order guarantees between separately dispatched graphics nodes, though individual primitives within a single dispatch do adhere to rasterization order. Draw calls executed before or after the execution graph also execute relative to each graphics node with respect to rasterization order.

For this command, all device/host pointers in substructures are treated as device pointers and read during device execution of this command. The allocation and contents of these pointers only needs to be valid during device execution. All of these addresses will be read in the `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT` pipeline stage with the `VK_ACCESS_2_SHADER_STORAGE_READ_BIT` access flag.

Execution of this command **may** modify any memory locations in the range [`scratch`,`scratch` + `scratchSize`). Accesses to this memory range are performed in the `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT` pipeline stage with the `VK_ACCESS_2_SHADER_STORAGE_READ_BIT` and `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT` access flags.

This command [captures command buffer state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#executiongraphs-meshnodes-statecapture) for mesh nodes similarly to draw commands.

Valid Usage

- [](#VUID-vkCmdDispatchGraphIndirectAMDX-magFilter-04553)VUID-vkCmdDispatchGraphIndirectAMDX-magFilter-04553  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `magFilter` or `minFilter` equal to `VK_FILTER_LINEAR`, `reductionMode` equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`, and `compareEnable` equal to `VK_FALSE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-magFilter-09598)VUID-vkCmdDispatchGraphIndirectAMDX-magFilter-09598  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `magFilter` or `minFilter` equal to `VK_FILTER_LINEAR` and `reductionMode` equal to either `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-mipmapMode-04770)VUID-vkCmdDispatchGraphIndirectAMDX-mipmapMode-04770  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `mipmapMode` equal to `VK_SAMPLER_MIPMAP_MODE_LINEAR`, `reductionMode` equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`, and `compareEnable` equal to `VK_FALSE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-mipmapMode-09599)VUID-vkCmdDispatchGraphIndirectAMDX-mipmapMode-09599  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `mipmapMode` equal to `VK_SAMPLER_MIPMAP_MODE_LINEAR` and `reductionMode` equal to either `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-unnormalizedCoordinates-09635)VUID-vkCmdDispatchGraphIndirectAMDX-unnormalizedCoordinates-09635  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s `levelCount` and `layerCount` **must** be 1
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08609)VUID-vkCmdDispatchGraphIndirectAMDX-None-08609  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s `viewType` **must** be `VK_IMAGE_VIEW_TYPE_1D` or `VK_IMAGE_VIEW_TYPE_2D`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08610)VUID-vkCmdDispatchGraphIndirectAMDX-None-08610  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the sampler **must** not be used with any of the SPIR-V `OpImageSample*` or `OpImageSparseSample*` instructions with `ImplicitLod`, `Dref` or `Proj` in their name
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08611)VUID-vkCmdDispatchGraphIndirectAMDX-None-08611  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) created with `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the sampler **must** not be used with any of the SPIR-V `OpImageSample*` or `OpImageSparseSample*` instructions that includes a LOD bias or any offset values
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-06479)VUID-vkCmdDispatchGraphIndirectAMDX-None-06479  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is sampled with [depth comparison](#textures-depth-compare-operation), the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_DEPTH_COMPARISON_BIT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-02691)VUID-vkCmdDispatchGraphIndirectAMDX-None-02691  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is accessed using atomic operations as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-07888)VUID-vkCmdDispatchGraphIndirectAMDX-None-07888  
  If a `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor is accessed using atomic operations as a result of this command, then the storage texel buffer’s [format features](#resources-buffer-view-format-features) **must** contain `VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-02692)VUID-vkCmdDispatchGraphIndirectAMDX-None-02692  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is sampled with `VK_FILTER_CUBIC_EXT` as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-02693)VUID-vkCmdDispatchGraphIndirectAMDX-None-02693  
  If the [VK\_EXT\_filter\_cubic](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_filter_cubic.html) extension is not enabled and any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is sampled with `VK_FILTER_CUBIC_EXT` as a result of this command, it **must** not have a [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html) of `VK_IMAGE_VIEW_TYPE_3D`, `VK_IMAGE_VIEW_TYPE_CUBE`, or `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-filterCubic-02694)VUID-vkCmdDispatchGraphIndirectAMDX-filterCubic-02694  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with `VK_FILTER_CUBIC_EXT` as a result of this command **must** have a [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html) and format that supports cubic filtering, as specified by [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)::`filterCubic` returned by [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-filterCubicMinmax-02695)VUID-vkCmdDispatchGraphIndirectAMDX-filterCubicMinmax-02695  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with `VK_FILTER_CUBIC_EXT` with a reduction mode of either `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` as a result of this command **must** have a [VkImageViewType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewType.html) and format that supports cubic filtering together with minmax filtering, as specified by [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)::`filterCubicMinmax` returned by [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-cubicRangeClamp-09212)VUID-vkCmdDispatchGraphIndirectAMDX-cubicRangeClamp-09212  
  If the [`cubicRangeClamp`](#features-cubicRangeClamp) feature is not enabled, then any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with `VK_FILTER_CUBIC_EXT` as a result of this command **must** not have a [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeCreateInfo.html)::`reductionMode` equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-reductionMode-09213)VUID-vkCmdDispatchGraphIndirectAMDX-reductionMode-09213  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with a [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeCreateInfo.html)::`reductionMode` equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM` as a result of this command **must** sample with `VK_FILTER_CUBIC_EXT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-selectableCubicWeights-09214)VUID-vkCmdDispatchGraphIndirectAMDX-selectableCubicWeights-09214  
  If the [`selectableCubicWeights`](#features-selectableCubicWeights) feature is not enabled, then any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being sampled with `VK_FILTER_CUBIC_EXT` as a result of this command **must** have [VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html)::`cubicWeights` equal to `VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-flags-02696)VUID-vkCmdDispatchGraphIndirectAMDX-flags-02696  
  Any [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) created with a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`flags` containing `VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV` sampled as a result of this command **must** only be sampled using a [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerAddressMode.html) of `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpTypeImage-07027)VUID-vkCmdDispatchGraphIndirectAMDX-OpTypeImage-07027  
  For any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being written as a storage image where the image format field of the `OpTypeImage` is `Unknown`, the view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpTypeImage-07028)VUID-vkCmdDispatchGraphIndirectAMDX-OpTypeImage-07028  
  For any [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) being read as a storage image where the image format field of the `OpTypeImage` is `Unknown`, the view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpTypeImage-07029)VUID-vkCmdDispatchGraphIndirectAMDX-OpTypeImage-07029  
  For any [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) being written as a storage texel buffer where the image format field of the `OpTypeImage` is `Unknown`, the view’s [buffer features](#VkFormatProperties3) **must** contain `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpTypeImage-07030)VUID-vkCmdDispatchGraphIndirectAMDX-OpTypeImage-07030  
  Any [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) being read as a storage texel buffer where the image format field of the `OpTypeImage` is `Unknown` then the view’s [buffer features](#VkFormatProperties3) **must** contain `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08600)VUID-vkCmdDispatchGraphIndirectAMDX-None-08600  
  For each set *n* that is statically used by [a bound shader](#shaders-binding), a descriptor set **must** have been bound to *n* at the same pipeline bind point, with a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) that is compatible for set *n*, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) used to create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) or the [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) array used to create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) , as described in [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08601)VUID-vkCmdDispatchGraphIndirectAMDX-None-08601  
  For each push constant that is statically used by [a bound shader](#shaders-binding), a push constant value **must** have been set for the same pipeline bind point, with a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) that is compatible for push constants, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) used to create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) or the [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) array used to create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) , as described in [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-10068)VUID-vkCmdDispatchGraphIndirectAMDX-None-10068  
  For each array of resources that is used by [a bound shader](#shaders-binding), the indices used to access members of the array **must** be less than the descriptor count for the identified binding in the descriptor sets used by this command
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-maintenance4-08602)VUID-vkCmdDispatchGraphIndirectAMDX-maintenance4-08602  
  If the [`maintenance4`](#features-maintenance4) feature is not enabled, then for each push constant that is statically used by [a bound shader](#shaders-binding), a push constant value **must** have been set for the same pipeline bind point, with a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) that is compatible for push constants, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) used to create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) or the [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) and [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantRange.html) arrays used to create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) , as described in [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08114)VUID-vkCmdDispatchGraphIndirectAMDX-None-08114  
  Descriptors in each bound descriptor set, specified via [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets.html), **must** be valid as described by [descriptor validity](#descriptor-validity) if they are statically used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point used by this command and the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) was not created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08115)VUID-vkCmdDispatchGraphIndirectAMDX-None-08115  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point were specified via [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets.html), the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) **must** have been created without `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08116)VUID-vkCmdDispatchGraphIndirectAMDX-None-08116  
  Descriptors in bound descriptor buffers, specified via [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), **must** be valid if they are dynamically used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point used by this command and the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) was created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08604)VUID-vkCmdDispatchGraphIndirectAMDX-None-08604  
  Descriptors in bound descriptor buffers, specified via [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), **must** be valid if they are dynamically used by any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) bound to a stage corresponding to the pipeline bind point used by this command
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08117)VUID-vkCmdDispatchGraphIndirectAMDX-None-08117  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point were specified via [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) **must** have been created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08119)VUID-vkCmdDispatchGraphIndirectAMDX-None-08119  
  If a descriptor is dynamically used with a [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, the descriptor memory **must** be resident
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08605)VUID-vkCmdDispatchGraphIndirectAMDX-None-08605  
  If a descriptor is dynamically used with a [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) created with a `VkDescriptorSetLayout` that was created with `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, the descriptor memory **must** be resident
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08606)VUID-vkCmdDispatchGraphIndirectAMDX-None-08606  
  If the [`shaderObject`](#features-shaderObject) feature is not enabled, a valid pipeline **must** be bound to the pipeline bind point used by this command
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08608)VUID-vkCmdDispatchGraphIndirectAMDX-None-08608  
  If a pipeline is bound to the pipeline bind point used by this command, there **must** not have been any calls to dynamic state setting commands for any state specified statically in the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) object bound to the pipeline bind point used by this command, since that pipeline was bound
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-uniformBuffers-06935)VUID-vkCmdDispatchGraphIndirectAMDX-uniformBuffers-06935  
  If any stage of the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) object bound to the pipeline bind point used by this command accesses a uniform buffer, and that stage was created without enabling either `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS` or `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2` for `uniformBuffers`, and the [`robustBufferAccess`](#features-robustBufferAccess) feature is not enabled, that stage **must** not access values outside of the range of the buffer as specified in the descriptor set bound to the same pipeline bind point
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08612)VUID-vkCmdDispatchGraphIndirectAMDX-None-08612  
  If the [`robustBufferAccess`](#features-robustBufferAccess) feature is not enabled, and any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) bound to a stage corresponding to the pipeline bind point used by this command accesses a uniform buffer, it **must** not access values outside of the range of the buffer as specified in the descriptor set bound to the same pipeline bind point
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-storageBuffers-06936)VUID-vkCmdDispatchGraphIndirectAMDX-storageBuffers-06936  
  If any stage of the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) object bound to the pipeline bind point used by this command accesses a storage buffer, and that stage was created without enabling either `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS` or `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2` for `storageBuffers`, and the [`robustBufferAccess`](#features-robustBufferAccess) feature is not enabled, that stage **must** not access values outside of the range of the buffer as specified in the descriptor set bound to the same pipeline bind point
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-08613)VUID-vkCmdDispatchGraphIndirectAMDX-None-08613  
  If the [`robustBufferAccess`](#features-robustBufferAccess) feature is not enabled, and any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) bound to a stage corresponding to the pipeline bind point used by this command accesses a storage buffer, it **must** not access values outside of the range of the buffer as specified in the descriptor set bound to the same pipeline bind point
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-02707)VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-02707  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, any resource accessed by [bound shaders](#shaders-binding) **must** not be a protected resource
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-viewType-07752)VUID-vkCmdDispatchGraphIndirectAMDX-viewType-07752  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) is accessed as a result of this command, then the image view’s `viewType` **must** match the `Dim` operand of the `OpTypeImage` as described in [\[spirvenv-image-dimensions\]](#spirvenv-image-dimensions)
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-format-07753)VUID-vkCmdDispatchGraphIndirectAMDX-format-07753  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) or [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) is accessed as a result of this command, then the [numeric type](#formats-numericformat) of the view’s `format` and the `Sampled` `Type` operand of the `OpTypeImage` **must** match
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWrite-08795)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWrite-08795  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) created with a format other than `VK_FORMAT_A8_UNORM` is accessed using `OpImageWrite` as a result of this command, then the `Type` of the `Texel` operand of that instruction **must** have at least as many components as the image view’s format
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWrite-08796)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWrite-08796  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) created with the format `VK_FORMAT_A8_UNORM` is accessed using `OpImageWrite` as a result of this command, then the `Type` of the `Texel` operand of that instruction **must** have four components
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWrite-04469)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWrite-04469  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) is accessed using `OpImageWrite` as a result of this command, then the `Type` of the `Texel` operand of that instruction **must** have at least as many components as the buffer view’s format
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-SampledType-04470)VUID-vkCmdDispatchGraphIndirectAMDX-SampledType-04470  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that has a 64-bit component width is accessed as a result of this command, the `SampledType` of the `OpTypeImage` operand of that instruction **must** have a `Width` of 64
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-SampledType-04471)VUID-vkCmdDispatchGraphIndirectAMDX-SampledType-04471  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that has a component width less than 64-bit is accessed as a result of this command, the `SampledType` of the `OpTypeImage` operand of that instruction **must** have a `Width` of 32
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-SampledType-04472)VUID-vkCmdDispatchGraphIndirectAMDX-SampledType-04472  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that has a 64-bit component width is accessed as a result of this command, the `SampledType` of the `OpTypeImage` operand of that instruction **must** have a `Width` of 64
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-SampledType-04473)VUID-vkCmdDispatchGraphIndirectAMDX-SampledType-04473  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that has a component width less than 64-bit is accessed as a result of this command, the `SampledType` of the `OpTypeImage` operand of that instruction **must** have a `Width` of 32
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-sparseImageInt64Atomics-04474)VUID-vkCmdDispatchGraphIndirectAMDX-sparseImageInt64Atomics-04474  
  If the [`sparseImageInt64Atomics`](#features-sparseImageInt64Atomics) feature is not enabled, [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) objects created with the `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` flag **must** not be accessed by atomic instructions through an `OpTypeImage` with a `SampledType` with a `Width` of 64 by this command
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-sparseImageInt64Atomics-04475)VUID-vkCmdDispatchGraphIndirectAMDX-sparseImageInt64Atomics-04475  
  If the [`sparseImageInt64Atomics`](#features-sparseImageInt64Atomics) feature is not enabled, [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) objects created with the `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT` flag **must** not be accessed by atomic instructions through an `OpTypeImage` with a `SampledType` with a `Width` of 64 by this command
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWeightedSampleQCOM-06971)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWeightedSampleQCOM-06971  
  If `OpImageWeightedSampleQCOM` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_WEIGHT_SAMPLED_IMAGE_BIT_QCOM`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWeightedSampleQCOM-06972)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWeightedSampleQCOM-06972  
  If `OpImageWeightedSampleQCOM` uses a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a sample weight image as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_WEIGHT_IMAGE_BIT_QCOM`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBoxFilterQCOM-06973)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBoxFilterQCOM-06973  
  If `OpImageBoxFilterQCOM` is used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_BOX_FILTER_SAMPLED_BIT_QCOM`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBlockMatchSSDQCOM-06974)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBlockMatchSSDQCOM-06974  
  If `OpImageBlockMatchSSDQCOM` is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBlockMatchSADQCOM-06975)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBlockMatchSADQCOM-06975  
  If `OpImageBlockMatchSADQCOM` is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBlockMatchSADQCOM-06976)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBlockMatchSADQCOM-06976  
  If `OpImageBlockMatchSADQCOM` or OpImageBlockMatchSSDQCOM is used to read from a reference image as result of this command, then the specified reference coordinates **must** not fail [integer texel coordinate validation](#textures-integer-coordinate-validation)
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWeightedSampleQCOM-06977)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWeightedSampleQCOM-06977  
  If `OpImageWeightedSampleQCOM`, `OpImageBoxFilterQCOM`, `OpImageBlockMatchWindowSSDQCOM`, `OpImageBlockMatchWindowSADQCOM`, `OpImageBlockMatchGatherSSDQCOM`, `OpImageBlockMatchGatherSADQCOM`, `OpImageBlockMatchSSDQCOM`, or `OpImageBlockMatchSADQCOM` uses a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) as a result of this command, then the sampler **must** have been created with `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWeightedSampleQCOM-06978)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageWeightedSampleQCOM-06978  
  If any command other than `OpImageWeightedSampleQCOM`, `OpImageBoxFilterQCOM`, `OpImageBlockMatchWindowSSDQCOM`, `OpImageBlockMatchWindowSADQCOM`, `OpImageBlockMatchGatherSSDQCOM`, `OpImageBlockMatchGatherSADQCOM`, `OpImageBlockMatchSSDQCOM`, or `OpImageBlockMatchSADQCOM` uses a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) as a result of this command, then the sampler **must** not have been created with `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBlockMatchWindow-09215)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBlockMatchWindow-09215  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM` instruction is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s [format features](#resources-image-view-format-features) **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBlockMatchWindow-09216)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBlockMatchWindow-09216  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM` instruction is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) as a result of this command, then the image view’s format **must** be a single-component format
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBlockMatchWindow-09217)VUID-vkCmdDispatchGraphIndirectAMDX-OpImageBlockMatchWindow-09217  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM` read from a reference image as result of this command, then the specified reference coordinates **must** not fail [integer texel coordinate validation](#textures-integer-coordinate-validation)
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-07288)VUID-vkCmdDispatchGraphIndirectAMDX-None-07288  
  Any shader invocation executed by this command **must** [terminate](#shaders-termination)
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-09600)VUID-vkCmdDispatchGraphIndirectAMDX-None-09600  
  If a descriptor with type equal to any of `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM`, `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`, `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` is accessed as a result of this command, all image subresources identified by that descriptor **must** be in the image layout identified when the descriptor was written
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-10746)VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-10746  
  The `VkDeviceMemory` object allocated from a `VkMemoryHeap` with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property that is bound to a resource accessed as a result of this command **must** be the active bound [bound tile memory object](#memory-bind-tile-memory) in `commandBuffer`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-10678)VUID-vkCmdDispatchGraphIndirectAMDX-None-10678  
  If this command is recorded inside a [tile shading render pass](#renderpass-tile-shading) instance, the stages corresponding to the pipeline bind point used by this command **must** only include `VK_SHADER_STAGE_VERTEX_BIT`, `VK_SHADER_STAGE_FRAGMENT_BIT`, and/or `VK_SHADER_STAGE_COMPUTE_BIT`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-10679)VUID-vkCmdDispatchGraphIndirectAMDX-None-10679  
  If this command is recorded where [per-tile execution model](#renderpass-per-tile-execution-model) is enabled, there **must** be no access to any image while the image was be transitioned to the `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` layout
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-pDescription-09900)VUID-vkCmdDispatchGraphIndirectAMDX-pDescription-09900  
  If a `VK_DESCRIPTOR_TYPE_TENSOR_ARM` descriptor is accessed as a result of this command, then the underlying [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) object **must** have been created with a [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)::`pDescription` whose `usage` member contained `VK_TENSOR_USAGE_SHADER_BIT_ARM`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-dimensionCount-09905)VUID-vkCmdDispatchGraphIndirectAMDX-dimensionCount-09905  
  If a `VK_DESCRIPTOR_TYPE_TENSOR_ARM` descriptor is accessed as a result of this command, then the `Rank` of the `OpTypeTensorARM` of the tensor resource variable **must** be equal to the `dimensionCount` provided via [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)::`pDescription` when creating the underlying [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) object
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-OpTypeTensorARM-09906)VUID-vkCmdDispatchGraphIndirectAMDX-OpTypeTensorARM-09906  
  If a `VK_DESCRIPTOR_TYPE_TENSOR_ARM` descriptor is accessed as a result of this command, then the element type of the `OpTypeTensorARM` of the tensor resource variable **must** be [compatible](#spirvenv-tensor-formats) with the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of the [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html) used for the access
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-09181)VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-09181  
  `commandBuffer` **must** not be a protected command buffer
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-09182)VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-09182  
  `commandBuffer` **must** be a primary command buffer
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-scratch-10192)VUID-vkCmdDispatchGraphIndirectAMDX-scratch-10192  
  `scratch` **must** be the device address of an allocated memory range at least as large as `scratchSize`
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-scratchSize-10193)VUID-vkCmdDispatchGraphIndirectAMDX-scratchSize-10193  
  `scratchSize` **must** be greater than or equal to [VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html)::`minSize` returned by [vkGetExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExecutionGraphPipelineScratchSizeAMDX.html) for the bound execution graph pipeline
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-scratch-09184)VUID-vkCmdDispatchGraphIndirectAMDX-scratch-09184  
  `scratch` **must** be a device address within a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) created with the `VK_BUFFER_USAGE_EXECUTION_GRAPH_SCRATCH_BIT_AMDX` or `VK_BUFFER_USAGE_2_EXECUTION_GRAPH_SCRATCH_BIT_AMDX` flag
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-scratch-10194)VUID-vkCmdDispatchGraphIndirectAMDX-scratch-10194  
  The device memory range [`scratch`,`scratch`  
  `scratchSize`] **must** have been initialized with [vkCmdInitializeGraphScratchMemoryAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdInitializeGraphScratchMemoryAMDX.html) using the bound execution graph pipeline, and not modified after that by anything other than another execution graph dispatch command
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-maxComputeWorkGroupCount-09186)VUID-vkCmdDispatchGraphIndirectAMDX-maxComputeWorkGroupCount-09186  
  Execution of this command **must** not cause a node to be dispatched with a larger number of workgroups than that specified by either a `MaxNumWorkgroupsAMDX` decoration in the dispatched node or [`maxComputeWorkGroupCount`](#limits-maxComputeWorkGroupCount)
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-maxExecutionGraphShaderPayloadCount-09187)VUID-vkCmdDispatchGraphIndirectAMDX-maxExecutionGraphShaderPayloadCount-09187  
  Execution of this command **must** not cause any shader to initialize more than [`maxExecutionGraphShaderPayloadCount`](#limits-maxExecutionGraphShaderPayloadCount) output payloads
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-NodeMaxPayloadsAMDX-09188)VUID-vkCmdDispatchGraphIndirectAMDX-NodeMaxPayloadsAMDX-09188  
  Execution of this command **must** not cause any shader that declares `NodeMaxPayloadsAMDX` to initialize more output payloads than specified by the max number of payloads for that decoration. This requirement applies to each `NodeMaxPayloadsAMDX` decoration separately
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-None-10195)VUID-vkCmdDispatchGraphIndirectAMDX-None-10195  
  If the bound execution graph pipeline includes draw nodes, this command **must** be called within a render pass instance that is compatible with the graphics pipeline used to create each of those nodes
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09150)VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09150  
  `pCountInfo->infos` **must** be a device pointer to a memory allocation at least as large as the product of `count` and `stride` when this command is executed on the device
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09151)VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09151  
  `pCountInfo->infos` **must** be a device address within a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) created with the `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` flag
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09152)VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09152  
  `pCountInfo->infos` **must** be a multiple of [`executionGraphDispatchAddressAlignment`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-executionGraphDispatchAddressAlignment)
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-infos-09153)VUID-vkCmdDispatchGraphIndirectAMDX-infos-09153  
  Device memory locations at indexes in the range \[`infos`, `infos` + (`count`\*`stride`)), at a granularity of `stride` **must** contain valid [VkDispatchGraphInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphInfoAMDX.html) structures in the first 24 bytes when this command is executed on the device
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09154)VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09154  
  For each [VkDispatchGraphInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphInfoAMDX.html) structure in `pCountInfo->infos`, `payloads` **must** be a device pointer to a memory allocation at least as large as the product of `payloadCount` and `payloadStride` when this command is executed on the device
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09155)VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09155  
  For each [VkDispatchGraphInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphInfoAMDX.html) structure in `pCountInfo->infos`, `payloads` **must** be a device address within a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) created with the `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` flag
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09156)VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09156  
  For each [VkDispatchGraphInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphInfoAMDX.html) structure in `pCountInfo->infos`, `payloads` **must** be a multiple of [`executionGraphDispatchAddressAlignment`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-executionGraphDispatchAddressAlignment)
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09157)VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09157  
  For each [VkDispatchGraphInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphInfoAMDX.html) structure in `pCountInfo->infos`, `nodeIndex` **must** be a valid node index in the bound execution graph pipeline, as returned by [vkGetExecutionGraphPipelineNodeIndexAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExecutionGraphPipelineNodeIndexAMDX.html) when this command is executed on the device
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09158)VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-09158  
  For each [VkDispatchGraphInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphInfoAMDX.html) structure in `pCountInfo->infos`, device memory locations at indexes in the range [`payloads`, `payloads` + (`payloadCount` * `payloadStride`)), at a granularity of `payloadStride` **must** contain a payload matching the size of the input payload expected by the node in `nodeIndex` in the first bytes when this command is executed on the device

Valid Usage (Implicit)

- [](#VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-parameter)VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-parameter)VUID-vkCmdDispatchGraphIndirectAMDX-pCountInfo-parameter  
  `pCountInfo` **must** be a valid pointer to a valid [VkDispatchGraphCountInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphCountInfoAMDX.html) structure
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-recording)VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-cmdpool)VUID-vkCmdDispatchGraphIndirectAMDX-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-videocoding)VUID-vkCmdDispatchGraphIndirectAMDX-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdDispatchGraphIndirectAMDX-bufferlevel)VUID-vkCmdDispatchGraphIndirectAMDX-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`

Host Synchronization

- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary

Both

Outside

Graphics  
Compute

Action

## [](#_see_also)See Also

[VK\_AMDX\_shader\_enqueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_shader_enqueue.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkDispatchGraphCountInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphCountInfoAMDX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdDispatchGraphIndirectAMDX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0