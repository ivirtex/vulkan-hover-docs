# vkCmdSubpassShadingHUAWEI(3) Manual Page

## Name

vkCmdSubpassShadingHUAWEI - Dispatch compute work items



## <a href="#_c_specification" class="anchor"></a>C Specification

A subpass shading dispatches a compute pipeline work with the work
dimension of render area of the calling subpass and work groups are
partitioned by specified work group size. Subpass operations like
`subpassLoad` are allowed to be used.

To record a subpass shading, call:

``` c
// Provided by VK_HUAWEI_subpass_shading
void vkCmdSubpassShadingHUAWEI(
    VkCommandBuffer                             commandBuffer);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

## <a href="#_description" class="anchor"></a>Description

When the command is executed, a global workgroup consisting of ceil
(render area size / local workgroup size) local workgroups is assembled.

Valid Usage

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-magFilter-04553"
  id="VUID-vkCmdSubpassShadingHUAWEI-magFilter-04553"></a>
  VUID-vkCmdSubpassShadingHUAWEI-magFilter-04553  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with `magFilter` or
  `minFilter` equal to `VK_FILTER_LINEAR`, `reductionMode` equal to
  `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`, and `compareEnable`
  equal to `VK_FALSE` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-magFilter-09598"
  id="VUID-vkCmdSubpassShadingHUAWEI-magFilter-09598"></a>
  VUID-vkCmdSubpassShadingHUAWEI-magFilter-09598  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with `magFilter` or
  `minFilter` equal to `VK_FILTER_LINEAR` and `reductionMode` equal to
  either `VK_SAMPLER_REDUCTION_MODE_MIN` or
  `VK_SAMPLER_REDUCTION_MODE_MAX` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-mipmapMode-04770"
  id="VUID-vkCmdSubpassShadingHUAWEI-mipmapMode-04770"></a>
  VUID-vkCmdSubpassShadingHUAWEI-mipmapMode-04770  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with `mipmapMode` equal to
  `VK_SAMPLER_MIPMAP_MODE_LINEAR`, `reductionMode` equal to
  `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`, and `compareEnable`
  equal to `VK_FALSE` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-mipmapMode-09599"
  id="VUID-vkCmdSubpassShadingHUAWEI-mipmapMode-09599"></a>
  VUID-vkCmdSubpassShadingHUAWEI-mipmapMode-09599  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with `mipmapMode` equal to
  `VK_SAMPLER_MIPMAP_MODE_LINEAR` and `reductionMode` equal to either
  `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` is
  used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this
  command, then the image view’s [format
  features](#resources-image-view-format-features) **must** contain
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-unnormalizedCoordinates-09635"
  id="VUID-vkCmdSubpassShadingHUAWEI-unnormalizedCoordinates-09635"></a>
  VUID-vkCmdSubpassShadingHUAWEI-unnormalizedCoordinates-09635  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with
  `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s `levelCount` and `layerCount` **must** be 1

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-unnormalizedCoordinates-09636"
  id="VUID-vkCmdSubpassShadingHUAWEI-unnormalizedCoordinates-09636"></a>
  VUID-vkCmdSubpassShadingHUAWEI-unnormalizedCoordinates-09636  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with
  `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s `viewType` **must** be `VK_IMAGE_VIEW_TYPE_1D` or
  `VK_IMAGE_VIEW_TYPE_2D`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-06479"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-06479"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-06479  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is sampled with [depth
  comparison](#textures-depth-compare-operation), the image view’s
  [format features](#resources-image-view-format-features) **must**
  contain `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_DEPTH_COMPARISON_BIT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-02691"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-02691"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-02691  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is accessed using atomic
  operations as a result of this command, then the image view’s [format
  features](#resources-image-view-format-features) **must** contain
  `VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-07888"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-07888"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-07888  
  If a `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor is accessed
  using atomic operations as a result of this command, then the storage
  texel buffer’s [format
  features](#resources-buffer-view-format-features) **must** contain
  `VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-02692"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-02692"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-02692  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is sampled with
  `VK_FILTER_CUBIC_EXT` as a result of this command, then the image
  view’s [format features](#resources-image-view-format-features)
  **must** contain
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-02693"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-02693"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-02693  
  If the [VK_EXT_filter_cubic](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_filter_cubic.html) extension is
  not enabled and any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is sampled with
  `VK_FILTER_CUBIC_EXT` as a result of this command, it **must** not
  have a [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html) of
  `VK_IMAGE_VIEW_TYPE_3D`, `VK_IMAGE_VIEW_TYPE_CUBE`, or
  `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-filterCubic-02694"
  id="VUID-vkCmdSubpassShadingHUAWEI-filterCubic-02694"></a>
  VUID-vkCmdSubpassShadingHUAWEI-filterCubic-02694  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being sampled with
  `VK_FILTER_CUBIC_EXT` as a result of this command **must** have a
  [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html) and format that supports cubic
  filtering, as specified by
  [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)::`filterCubic`
  returned by
  [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-filterCubicMinmax-02695"
  id="VUID-vkCmdSubpassShadingHUAWEI-filterCubicMinmax-02695"></a>
  VUID-vkCmdSubpassShadingHUAWEI-filterCubicMinmax-02695  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being sampled with
  `VK_FILTER_CUBIC_EXT` with a reduction mode of either
  `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` as
  a result of this command **must** have a
  [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html) and format that supports cubic
  filtering together with minmax filtering, as specified by
  [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)::`filterCubicMinmax`
  returned by
  [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-cubicRangeClamp-09212"
  id="VUID-vkCmdSubpassShadingHUAWEI-cubicRangeClamp-09212"></a>
  VUID-vkCmdSubpassShadingHUAWEI-cubicRangeClamp-09212  
  If the [`cubicRangeClamp`](#features-filter-cubic-range-clamp) feature
  is not enabled, then any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being sampled
  with `VK_FILTER_CUBIC_EXT` as a result of this command **must** not
  have a
  [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfo.html)::`reductionMode`
  equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-reductionMode-09213"
  id="VUID-vkCmdSubpassShadingHUAWEI-reductionMode-09213"></a>
  VUID-vkCmdSubpassShadingHUAWEI-reductionMode-09213  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being sampled with a
  [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfo.html)::`reductionMode`
  equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM`
  as a result of this command **must** sample with `VK_FILTER_CUBIC_EXT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-selectableCubicWeights-09214"
  id="VUID-vkCmdSubpassShadingHUAWEI-selectableCubicWeights-09214"></a>
  VUID-vkCmdSubpassShadingHUAWEI-selectableCubicWeights-09214  
  If the
  [`selectableCubicWeights`](#features-filter-cubic-weight-selection)
  feature is not enabled, then any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being
  sampled with `VK_FILTER_CUBIC_EXT` as a result of this command
  **must** have
  [VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html)::`cubicWeights`
  equal to `VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-flags-02696"
  id="VUID-vkCmdSubpassShadingHUAWEI-flags-02696"></a>
  VUID-vkCmdSubpassShadingHUAWEI-flags-02696  
  Any [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) created with a
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags` containing
  `VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV` sampled as a result of this
  command **must** only be sampled using a
  [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html) of
  `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpTypeImage-07027"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpTypeImage-07027"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpTypeImage-07027  
  For any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being written as a storage
  image where the image format field of the `OpTypeImage` is `Unknown`,
  the view’s [format features](#resources-image-view-format-features)
  **must** contain
  `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpTypeImage-07028"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpTypeImage-07028"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpTypeImage-07028  
  For any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being read as a storage image
  where the image format field of the `OpTypeImage` is `Unknown`, the
  view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpTypeImage-07029"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpTypeImage-07029"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpTypeImage-07029  
  For any [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) being written as a storage
  texel buffer where the image format field of the `OpTypeImage` is
  `Unknown`, the view’s [buffer features](#VkFormatProperties3) **must**
  contain `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpTypeImage-07030"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpTypeImage-07030"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpTypeImage-07030  
  Any [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) being read as a storage texel
  buffer where the image format field of the `OpTypeImage` is `Unknown`
  then the view’s [buffer features](#VkFormatProperties3) **must**
  contain `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08600"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08600"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08600  
  For each set *n* that is statically used by [a bound
  shader](#shaders-binding), a descriptor set **must** have been bound
  to *n* at the same pipeline bind point, with a
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) that is compatible for set
  *n*, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) used to create
  the current [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) or the
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) array used to
  create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) , as described in
  [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08601"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08601"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08601  
  For each push constant that is statically used by [a bound
  shader](#shaders-binding), a push constant value **must** have been
  set for the same pipeline bind point, with a
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) that is compatible for push
  constants, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) used to
  create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) or the
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) array used to
  create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) , as described in
  [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-maintenance4-08602"
  id="VUID-vkCmdSubpassShadingHUAWEI-maintenance4-08602"></a>
  VUID-vkCmdSubpassShadingHUAWEI-maintenance4-08602  
  If the [`maintenance4`](#features-maintenance4) feature is not
  enabled, then for each push constant that is statically used by [a
  bound shader](#shaders-binding), a push constant value **must** have
  been set for the same pipeline bind point, with a
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) that is compatible for push
  constants, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) used to
  create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) or the
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) and
  [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantRange.html) arrays used to create
  the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) , as described in
  [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08114"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08114"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08114  
  Descriptors in each bound descriptor set, specified via
  [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets.html), **must** be
  valid as described by [descriptor validity](#descriptor-validity) if
  they are statically used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) bound to
  the pipeline bind point used by this command and the bound
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) was not created with
  `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08115"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08115"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08115  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) bound to
  the pipeline bind point were specified via
  [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets.html), the bound
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) **must** have been created without
  `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08116"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08116"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08116  
  Descriptors in bound descriptor buffers, specified via
  [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html),
  **must** be valid if they are dynamically used by the
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) bound to the pipeline bind point used by
  this command and the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) was created
  with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08604"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08604"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08604  
  Descriptors in bound descriptor buffers, specified via
  [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html),
  **must** be valid if they are dynamically used by any
  [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) bound to a stage corresponding to the
  pipeline bind point used by this command

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08117"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08117"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08117  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) bound to
  the pipeline bind point were specified via
  [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html),
  the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) **must** have been created
  with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08119"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08119"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08119  
  If a descriptor is dynamically used with a
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) created with
  `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, the descriptor memory
  **must** be resident

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08605"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08605"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08605  
  If a descriptor is dynamically used with a
  [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) created with a `VkDescriptorSetLayout`
  that was created with
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, the
  descriptor memory **must** be resident

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08606"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08606"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08606  
  If the [`shaderObject`](#features-shaderObject) feature is not
  enabled, a valid pipeline **must** be bound to the pipeline bind point
  used by this command

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08608"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08608"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08608  
  If a pipeline is bound to the pipeline bind point used by this
  command, there **must** not have been any calls to dynamic state
  setting commands for any state not specified as dynamic in the
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object bound to the pipeline bind point
  used by this command, since that pipeline was bound

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08609"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08609"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08609  
  If the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object bound to the pipeline bind
  point used by this command or any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html)
  bound to a stage corresponding to the pipeline bind point used by this
  command accesses a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) object that uses
  unnormalized coordinates, that sampler **must** not be used to sample
  from any [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) with a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) of the type `VK_IMAGE_VIEW_TYPE_3D`,
  `VK_IMAGE_VIEW_TYPE_CUBE`, `VK_IMAGE_VIEW_TYPE_1D_ARRAY`,
  `VK_IMAGE_VIEW_TYPE_2D_ARRAY` or `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY`, in
  any shader stage

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08610"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08610"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08610  
  If the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object bound to the pipeline bind
  point used by this command or any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html)
  bound to a stage corresponding to the pipeline bind point used by this
  command accesses a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) object that uses
  unnormalized coordinates, that sampler **must** not be used with any
  of the SPIR-V `OpImageSample*` or `OpImageSparseSample*` instructions
  with `ImplicitLod`, `Dref` or `Proj` in their name, in any shader
  stage

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08611"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08611"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08611  
  If the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object bound to the pipeline bind
  point used by this command or any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html)
  bound to a stage corresponding to the pipeline bind point used by this
  command accesses a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) object that uses
  unnormalized coordinates, that sampler **must** not be used with any
  of the SPIR-V `OpImageSample*` or `OpImageSparseSample*` instructions
  that includes a LOD bias or any offset values, in any shader stage

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08607"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08607"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08607  
  If the [`shaderObject`](#features-shaderObject) is enabled, either a
  valid pipeline **must** be bound to the pipeline bind point used by
  this command, or a valid combination of valid and
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) shader objects **must** be bound
  to every supported shader stage corresponding to the pipeline bind
  point used by this command

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-uniformBuffers-06935"
  id="VUID-vkCmdSubpassShadingHUAWEI-uniformBuffers-06935"></a>
  VUID-vkCmdSubpassShadingHUAWEI-uniformBuffers-06935  
  If any stage of the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object bound to the
  pipeline bind point used by this command accesses a uniform buffer,
  and that stage was created without enabling either
  `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_EXT` or
  `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2_EXT`
  for `uniformBuffers`, and the
  [`robustBufferAccess`](#features-robustBufferAccess) feature is not
  enabled, that stage **must** not access values outside of the range of
  the buffer as specified in the descriptor set bound to the same
  pipeline bind point

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08612"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08612"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08612  
  If the [`robustBufferAccess`](#features-robustBufferAccess) feature is
  not enabled, and any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) bound to a stage
  corresponding to the pipeline bind point used by this command accesses
  a uniform buffer, it **must** not access values outside of the range
  of the buffer as specified in the descriptor set bound to the same
  pipeline bind point

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-storageBuffers-06936"
  id="VUID-vkCmdSubpassShadingHUAWEI-storageBuffers-06936"></a>
  VUID-vkCmdSubpassShadingHUAWEI-storageBuffers-06936  
  If any stage of the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object bound to the
  pipeline bind point used by this command accesses a storage buffer,
  and that stage was created without enabling either
  `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_EXT` or
  `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2_EXT`
  for `storageBuffers`, and the
  [`robustBufferAccess`](#features-robustBufferAccess) feature is not
  enabled, that stage **must** not access values outside of the range of
  the buffer as specified in the descriptor set bound to the same
  pipeline bind point

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-08613"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-08613"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-08613  
  If the [`robustBufferAccess`](#features-robustBufferAccess) feature is
  not enabled, and any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) bound to a stage
  corresponding to the pipeline bind point used by this command accesses
  a storage buffer, it **must** not access values outside of the range
  of the buffer as specified in the descriptor set bound to the same
  pipeline bind point

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-commandBuffer-02707"
  id="VUID-vkCmdSubpassShadingHUAWEI-commandBuffer-02707"></a>
  VUID-vkCmdSubpassShadingHUAWEI-commandBuffer-02707  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported, any
  resource accessed by [bound shaders](#shaders-binding) **must** not be
  a protected resource

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-06550"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-06550"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-06550  
  If [a bound shader](#shaders-binding) accesses a
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) or [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) object
  that enables [sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion](#samplers-YCbCr-conversion), that object **must** only be
  used with `OpImageSample*` or `OpImageSparseSample*` instructions

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-ConstOffset-06551"
  id="VUID-vkCmdSubpassShadingHUAWEI-ConstOffset-06551"></a>
  VUID-vkCmdSubpassShadingHUAWEI-ConstOffset-06551  
  If [a bound shader](#shaders-binding) accesses a
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) or [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) object
  that enables [sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion](#samplers-YCbCr-conversion), that object **must** not use
  the `ConstOffset` and `Offset` operands

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-viewType-07752"
  id="VUID-vkCmdSubpassShadingHUAWEI-viewType-07752"></a>
  VUID-vkCmdSubpassShadingHUAWEI-viewType-07752  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is accessed as a result of this
  command, then the image view’s `viewType` **must** match the `Dim`
  operand of the `OpTypeImage` as described in
  [\[textures-operation-validation\]](#textures-operation-validation)

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-format-07753"
  id="VUID-vkCmdSubpassShadingHUAWEI-format-07753"></a>
  VUID-vkCmdSubpassShadingHUAWEI-format-07753  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is accessed as a result of this
  command, then the [numeric type](#formats-numericformat) of the image
  view’s `format` and the `Sampled` `Type` operand of the `OpTypeImage`
  **must** match

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageWrite-08795"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageWrite-08795"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageWrite-08795  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) created with a format other than
  `VK_FORMAT_A8_UNORM_KHR` is accessed using `OpImageWrite` as a result
  of this command, then the `Type` of the `Texel` operand of that
  instruction **must** have at least as many components as the image
  view’s format

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageWrite-08796"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageWrite-08796"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageWrite-08796  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) created with the format
  `VK_FORMAT_A8_UNORM_KHR` is accessed using `OpImageWrite` as a result
  of this command, then the `Type` of the `Texel` operand of that
  instruction **must** have four components

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageWrite-04469"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageWrite-04469"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageWrite-04469  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) is accessed using
  `OpImageWrite` as a result of this command, then the `Type` of the
  `Texel` operand of that instruction **must** have at least as many
  components as the buffer view’s format

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-SampledType-04470"
  id="VUID-vkCmdSubpassShadingHUAWEI-SampledType-04470"></a>
  VUID-vkCmdSubpassShadingHUAWEI-SampledType-04470  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)
  that has a 64-bit component width is accessed as a result of this
  command, the `SampledType` of the `OpTypeImage` operand of that
  instruction **must** have a `Width` of 64

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-SampledType-04471"
  id="VUID-vkCmdSubpassShadingHUAWEI-SampledType-04471"></a>
  VUID-vkCmdSubpassShadingHUAWEI-SampledType-04471  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)
  that has a component width less than 64-bit is accessed as a result of
  this command, the `SampledType` of the `OpTypeImage` operand of that
  instruction **must** have a `Width` of 32

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-SampledType-04472"
  id="VUID-vkCmdSubpassShadingHUAWEI-SampledType-04472"></a>
  VUID-vkCmdSubpassShadingHUAWEI-SampledType-04472  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) with a
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) that has a 64-bit component width is
  accessed as a result of this command, the `SampledType` of the
  `OpTypeImage` operand of that instruction **must** have a `Width` of
  64

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-SampledType-04473"
  id="VUID-vkCmdSubpassShadingHUAWEI-SampledType-04473"></a>
  VUID-vkCmdSubpassShadingHUAWEI-SampledType-04473  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) with a
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) that has a component width less than 64-bit
  is accessed as a result of this command, the `SampledType` of the
  `OpTypeImage` operand of that instruction **must** have a `Width` of
  32

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-sparseImageInt64Atomics-04474"
  id="VUID-vkCmdSubpassShadingHUAWEI-sparseImageInt64Atomics-04474"></a>
  VUID-vkCmdSubpassShadingHUAWEI-sparseImageInt64Atomics-04474  
  If the [`sparseImageInt64Atomics`](#features-sparseImageInt64Atomics)
  feature is not enabled, [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) objects created with
  the `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` flag **must** not be
  accessed by atomic instructions through an `OpTypeImage` with a
  `SampledType` with a `Width` of 64 by this command

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-sparseImageInt64Atomics-04475"
  id="VUID-vkCmdSubpassShadingHUAWEI-sparseImageInt64Atomics-04475"></a>
  VUID-vkCmdSubpassShadingHUAWEI-sparseImageInt64Atomics-04475  
  If the [`sparseImageInt64Atomics`](#features-sparseImageInt64Atomics)
  feature is not enabled, [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) objects created with
  the `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT` flag **must** not be
  accessed by atomic instructions through an `OpTypeImage` with a
  `SampledType` with a `Width` of 64 by this command

- <a
  href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageWeightedSampleQCOM-06971"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageWeightedSampleQCOM-06971"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageWeightedSampleQCOM-06971  
  If `OpImageWeightedSampleQCOM` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_WEIGHT_SAMPLED_IMAGE_BIT_QCOM`

- <a
  href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageWeightedSampleQCOM-06972"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageWeightedSampleQCOM-06972"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageWeightedSampleQCOM-06972  
  If `OpImageWeightedSampleQCOM` uses a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html)
  as a sample weight image as a result of this command, then the image
  view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_WEIGHT_IMAGE_BIT_QCOM`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageBoxFilterQCOM-06973"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageBoxFilterQCOM-06973"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageBoxFilterQCOM-06973  
  If `OpImageBoxFilterQCOM` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_BOX_FILTER_SAMPLED_BIT_QCOM`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchSSDQCOM-06974"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchSSDQCOM-06974"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchSSDQCOM-06974  
  If `OpImageBlockMatchSSDQCOM` is used to read from an
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchSADQCOM-06975"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchSADQCOM-06975"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchSADQCOM-06975  
  If `OpImageBlockMatchSADQCOM` is used to read from an
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchSADQCOM-06976"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchSADQCOM-06976"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchSADQCOM-06976  
  If `OpImageBlockMatchSADQCOM` or OpImageBlockMatchSSDQCOM is used to
  read from a reference image as result of this command, then the
  specified reference coordinates **must** not fail [integer texel
  coordinate validation](#textures-integer-coordinate-validation)

- <a
  href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageWeightedSampleQCOM-06977"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageWeightedSampleQCOM-06977"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageWeightedSampleQCOM-06977  
  If `OpImageWeightedSampleQCOM`, `OpImageBoxFilterQCOM`,
  `OpImageBlockMatchWindowSSDQCOM`, `OpImageBlockMatchWindowSADQCOM`,
  `OpImageBlockMatchGatherSSDQCOM`, `OpImageBlockMatchGatherSADQCOM`,
  `OpImageBlockMatchSSDQCOM`, or `OpImageBlockMatchSADQCOM` uses a
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) as a result of this command, then the
  sampler **must** have been created with
  `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`

- <a
  href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageWeightedSampleQCOM-06978"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageWeightedSampleQCOM-06978"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageWeightedSampleQCOM-06978  
  If any command other than `OpImageWeightedSampleQCOM`,
  `OpImageBoxFilterQCOM`, `OpImageBlockMatchWindowSSDQCOM`,
  `OpImageBlockMatchWindowSADQCOM`, `OpImageBlockMatchGatherSSDQCOM`,
  `OpImageBlockMatchGatherSADQCOM`, `OpImageBlockMatchSSDQCOM`, or
  `OpImageBlockMatchSADQCOM` uses a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) as a
  result of this command, then the sampler **must** not have been
  created with `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchWindow-09215"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchWindow-09215"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchWindow-09215  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  instruction is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as
  a result of this command, then the image view’s [format
  features](#resources-image-view-format-features) **must** contain
  `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchWindow-09216"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchWindow-09216"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchWindow-09216  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  instruction is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as
  a result of this command, then the image view’s format **must** be a
  single-component format

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchWindow-09217"
  id="VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchWindow-09217"></a>
  VUID-vkCmdSubpassShadingHUAWEI-OpImageBlockMatchWindow-09217  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  read from a reference image as result of this command, then the
  specified reference coordinates **must** not fail [integer texel
  coordinate validation](#textures-integer-coordinate-validation)

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-07288"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-07288"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-07288  
  Any shader invocation executed by this command **must**
  [terminate](#shaders-termination)

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-09600"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-09600"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-09600  
  If a descriptor with type equal to any of
  `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM`,
  `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`,
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` is accessed as a result of this
  command, the image subresource identified by that descriptor **must**
  be in the image layout identified when the descriptor was written

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-None-04931"
  id="VUID-vkCmdSubpassShadingHUAWEI-None-04931"></a>
  VUID-vkCmdSubpassShadingHUAWEI-None-04931  
  This command must be called in a subpass with bind point
  `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI`. No draw commands can
  be called in the same subpass. Only one
  [vkCmdSubpassShadingHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSubpassShadingHUAWEI.html) command
  can be called in a subpass

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-commandBuffer-parameter"
  id="VUID-vkCmdSubpassShadingHUAWEI-commandBuffer-parameter"></a>
  VUID-vkCmdSubpassShadingHUAWEI-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-commandBuffer-recording"
  id="VUID-vkCmdSubpassShadingHUAWEI-commandBuffer-recording"></a>
  VUID-vkCmdSubpassShadingHUAWEI-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-commandBuffer-cmdpool"
  id="VUID-vkCmdSubpassShadingHUAWEI-commandBuffer-cmdpool"></a>
  VUID-vkCmdSubpassShadingHUAWEI-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-renderpass"
  id="VUID-vkCmdSubpassShadingHUAWEI-renderpass"></a>
  VUID-vkCmdSubpassShadingHUAWEI-renderpass  
  This command **must** only be called inside of a render pass instance

- <a href="#VUID-vkCmdSubpassShadingHUAWEI-videocoding"
  id="VUID-vkCmdSubpassShadingHUAWEI-videocoding"></a>
  VUID-vkCmdSubpassShadingHUAWEI-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr class="header">
<th class="tableblock halign-left valign-top"><a
href="#VkCommandBufferLevel">Command Buffer Levels</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginRenderPass">Render Pass Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginVideoCodingKHR">Video Coding Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#VkQueueFlagBits">Supported Queue Types</a></th>
<th class="tableblock halign-left valign-top"><a
href="#fundamentals-queueoperation-command-types">Command Type</a></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Inside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_HUAWEI_subpass_shading](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_HUAWEI_subpass_shading.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSubpassShadingHUAWEI"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
