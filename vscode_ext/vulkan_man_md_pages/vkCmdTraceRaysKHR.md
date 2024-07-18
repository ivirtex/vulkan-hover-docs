# vkCmdTraceRaysKHR(3) Manual Page

## Name

vkCmdTraceRaysKHR - Initialize a ray tracing dispatch



## <a href="#_c_specification" class="anchor"></a>C Specification

To dispatch ray tracing use:

``` c
// Provided by VK_KHR_ray_tracing_pipeline
void vkCmdTraceRaysKHR(
    VkCommandBuffer                             commandBuffer,
    const VkStridedDeviceAddressRegionKHR*      pRaygenShaderBindingTable,
    const VkStridedDeviceAddressRegionKHR*      pMissShaderBindingTable,
    const VkStridedDeviceAddressRegionKHR*      pHitShaderBindingTable,
    const VkStridedDeviceAddressRegionKHR*      pCallableShaderBindingTable,
    uint32_t                                    width,
    uint32_t                                    height,
    uint32_t                                    depth);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pRaygenShaderBindingTable` is a
  [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html)
  that holds the shader binding table data for the ray generation shader
  stage.

- `pMissShaderBindingTable` is a
  [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html)
  that holds the shader binding table data for the miss shader stage.

- `pHitShaderBindingTable` is a
  [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html)
  that holds the shader binding table data for the hit shader stage.

- `pCallableShaderBindingTable` is a
  [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html)
  that holds the shader binding table data for the callable shader
  stage.

- `width` is the width of the ray trace query dimensions.

- `height` is height of the ray trace query dimensions.

- `depth` is depth of the ray trace query dimensions.

## <a href="#_description" class="anchor"></a>Description

When the command is executed, a ray generation group of `width` ×
`height` × `depth` rays is assembled.

Valid Usage

- <a href="#VUID-vkCmdTraceRaysKHR-magFilter-04553"
  id="VUID-vkCmdTraceRaysKHR-magFilter-04553"></a>
  VUID-vkCmdTraceRaysKHR-magFilter-04553  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with `magFilter` or
  `minFilter` equal to `VK_FILTER_LINEAR`, `reductionMode` equal to
  `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`, and `compareEnable`
  equal to `VK_FALSE` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`

- <a href="#VUID-vkCmdTraceRaysKHR-magFilter-09598"
  id="VUID-vkCmdTraceRaysKHR-magFilter-09598"></a>
  VUID-vkCmdTraceRaysKHR-magFilter-09598  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with `magFilter` or
  `minFilter` equal to `VK_FILTER_LINEAR` and `reductionMode` equal to
  either `VK_SAMPLER_REDUCTION_MODE_MIN` or
  `VK_SAMPLER_REDUCTION_MODE_MAX` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT`

- <a href="#VUID-vkCmdTraceRaysKHR-mipmapMode-04770"
  id="VUID-vkCmdTraceRaysKHR-mipmapMode-04770"></a>
  VUID-vkCmdTraceRaysKHR-mipmapMode-04770  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with `mipmapMode` equal to
  `VK_SAMPLER_MIPMAP_MODE_LINEAR`, `reductionMode` equal to
  `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`, and `compareEnable`
  equal to `VK_FALSE` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`

- <a href="#VUID-vkCmdTraceRaysKHR-mipmapMode-09599"
  id="VUID-vkCmdTraceRaysKHR-mipmapMode-09599"></a>
  VUID-vkCmdTraceRaysKHR-mipmapMode-09599  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with `mipmapMode` equal to
  `VK_SAMPLER_MIPMAP_MODE_LINEAR` and `reductionMode` equal to either
  `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` is
  used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this
  command, then the image view’s [format
  features](#resources-image-view-format-features) **must** contain
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT`

- <a href="#VUID-vkCmdTraceRaysKHR-unnormalizedCoordinates-09635"
  id="VUID-vkCmdTraceRaysKHR-unnormalizedCoordinates-09635"></a>
  VUID-vkCmdTraceRaysKHR-unnormalizedCoordinates-09635  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with
  `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s `levelCount` and `layerCount` **must** be 1

- <a href="#VUID-vkCmdTraceRaysKHR-unnormalizedCoordinates-09636"
  id="VUID-vkCmdTraceRaysKHR-unnormalizedCoordinates-09636"></a>
  VUID-vkCmdTraceRaysKHR-unnormalizedCoordinates-09636  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with
  `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s `viewType` **must** be `VK_IMAGE_VIEW_TYPE_1D` or
  `VK_IMAGE_VIEW_TYPE_2D`

- <a href="#VUID-vkCmdTraceRaysKHR-None-06479"
  id="VUID-vkCmdTraceRaysKHR-None-06479"></a>
  VUID-vkCmdTraceRaysKHR-None-06479  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is sampled with [depth
  comparison](#textures-depth-compare-operation), the image view’s
  [format features](#resources-image-view-format-features) **must**
  contain `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_DEPTH_COMPARISON_BIT`

- <a href="#VUID-vkCmdTraceRaysKHR-None-02691"
  id="VUID-vkCmdTraceRaysKHR-None-02691"></a>
  VUID-vkCmdTraceRaysKHR-None-02691  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is accessed using atomic
  operations as a result of this command, then the image view’s [format
  features](#resources-image-view-format-features) **must** contain
  `VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT`

- <a href="#VUID-vkCmdTraceRaysKHR-None-07888"
  id="VUID-vkCmdTraceRaysKHR-None-07888"></a>
  VUID-vkCmdTraceRaysKHR-None-07888  
  If a `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor is accessed
  using atomic operations as a result of this command, then the storage
  texel buffer’s [format
  features](#resources-buffer-view-format-features) **must** contain
  `VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT`

- <a href="#VUID-vkCmdTraceRaysKHR-None-02692"
  id="VUID-vkCmdTraceRaysKHR-None-02692"></a>
  VUID-vkCmdTraceRaysKHR-None-02692  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is sampled with
  `VK_FILTER_CUBIC_EXT` as a result of this command, then the image
  view’s [format features](#resources-image-view-format-features)
  **must** contain
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT`

- <a href="#VUID-vkCmdTraceRaysKHR-None-02693"
  id="VUID-vkCmdTraceRaysKHR-None-02693"></a>
  VUID-vkCmdTraceRaysKHR-None-02693  
  If the [VK_EXT_filter_cubic](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_filter_cubic.html) extension is
  not enabled and any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is sampled with
  `VK_FILTER_CUBIC_EXT` as a result of this command, it **must** not
  have a [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html) of
  `VK_IMAGE_VIEW_TYPE_3D`, `VK_IMAGE_VIEW_TYPE_CUBE`, or
  `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY`

- <a href="#VUID-vkCmdTraceRaysKHR-filterCubic-02694"
  id="VUID-vkCmdTraceRaysKHR-filterCubic-02694"></a>
  VUID-vkCmdTraceRaysKHR-filterCubic-02694  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being sampled with
  `VK_FILTER_CUBIC_EXT` as a result of this command **must** have a
  [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html) and format that supports cubic
  filtering, as specified by
  [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)::`filterCubic`
  returned by
  [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)

- <a href="#VUID-vkCmdTraceRaysKHR-filterCubicMinmax-02695"
  id="VUID-vkCmdTraceRaysKHR-filterCubicMinmax-02695"></a>
  VUID-vkCmdTraceRaysKHR-filterCubicMinmax-02695  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being sampled with
  `VK_FILTER_CUBIC_EXT` with a reduction mode of either
  `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` as
  a result of this command **must** have a
  [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html) and format that supports cubic
  filtering together with minmax filtering, as specified by
  [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)::`filterCubicMinmax`
  returned by
  [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)

- <a href="#VUID-vkCmdTraceRaysKHR-cubicRangeClamp-09212"
  id="VUID-vkCmdTraceRaysKHR-cubicRangeClamp-09212"></a>
  VUID-vkCmdTraceRaysKHR-cubicRangeClamp-09212  
  If the [`cubicRangeClamp`](#features-filter-cubic-range-clamp) feature
  is not enabled, then any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being sampled
  with `VK_FILTER_CUBIC_EXT` as a result of this command **must** not
  have a
  [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfo.html)::`reductionMode`
  equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM`

- <a href="#VUID-vkCmdTraceRaysKHR-reductionMode-09213"
  id="VUID-vkCmdTraceRaysKHR-reductionMode-09213"></a>
  VUID-vkCmdTraceRaysKHR-reductionMode-09213  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being sampled with a
  [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfo.html)::`reductionMode`
  equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM`
  as a result of this command **must** sample with `VK_FILTER_CUBIC_EXT`

- <a href="#VUID-vkCmdTraceRaysKHR-selectableCubicWeights-09214"
  id="VUID-vkCmdTraceRaysKHR-selectableCubicWeights-09214"></a>
  VUID-vkCmdTraceRaysKHR-selectableCubicWeights-09214  
  If the
  [`selectableCubicWeights`](#features-filter-cubic-weight-selection)
  feature is not enabled, then any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being
  sampled with `VK_FILTER_CUBIC_EXT` as a result of this command
  **must** have
  [VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html)::`cubicWeights`
  equal to `VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM`

- <a href="#VUID-vkCmdTraceRaysKHR-flags-02696"
  id="VUID-vkCmdTraceRaysKHR-flags-02696"></a>
  VUID-vkCmdTraceRaysKHR-flags-02696  
  Any [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) created with a
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags` containing
  `VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV` sampled as a result of this
  command **must** only be sampled using a
  [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html) of
  `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE`

- <a href="#VUID-vkCmdTraceRaysKHR-OpTypeImage-07027"
  id="VUID-vkCmdTraceRaysKHR-OpTypeImage-07027"></a>
  VUID-vkCmdTraceRaysKHR-OpTypeImage-07027  
  For any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being written as a storage
  image where the image format field of the `OpTypeImage` is `Unknown`,
  the view’s [format features](#resources-image-view-format-features)
  **must** contain
  `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT`

- <a href="#VUID-vkCmdTraceRaysKHR-OpTypeImage-07028"
  id="VUID-vkCmdTraceRaysKHR-OpTypeImage-07028"></a>
  VUID-vkCmdTraceRaysKHR-OpTypeImage-07028  
  For any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being read as a storage image
  where the image format field of the `OpTypeImage` is `Unknown`, the
  view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT`

- <a href="#VUID-vkCmdTraceRaysKHR-OpTypeImage-07029"
  id="VUID-vkCmdTraceRaysKHR-OpTypeImage-07029"></a>
  VUID-vkCmdTraceRaysKHR-OpTypeImage-07029  
  For any [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) being written as a storage
  texel buffer where the image format field of the `OpTypeImage` is
  `Unknown`, the view’s [buffer features](#VkFormatProperties3) **must**
  contain `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT`

- <a href="#VUID-vkCmdTraceRaysKHR-OpTypeImage-07030"
  id="VUID-vkCmdTraceRaysKHR-OpTypeImage-07030"></a>
  VUID-vkCmdTraceRaysKHR-OpTypeImage-07030  
  Any [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) being read as a storage texel
  buffer where the image format field of the `OpTypeImage` is `Unknown`
  then the view’s [buffer features](#VkFormatProperties3) **must**
  contain `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT`

- <a href="#VUID-vkCmdTraceRaysKHR-None-08600"
  id="VUID-vkCmdTraceRaysKHR-None-08600"></a>
  VUID-vkCmdTraceRaysKHR-None-08600  
  For each set *n* that is statically used by [a bound
  shader](#shaders-binding), a descriptor set **must** have been bound
  to *n* at the same pipeline bind point, with a
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) that is compatible for set
  *n*, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) used to create
  the current [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) or the
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) array used to
  create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) , as described in
  [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)

- <a href="#VUID-vkCmdTraceRaysKHR-None-08601"
  id="VUID-vkCmdTraceRaysKHR-None-08601"></a>
  VUID-vkCmdTraceRaysKHR-None-08601  
  For each push constant that is statically used by [a bound
  shader](#shaders-binding), a push constant value **must** have been
  set for the same pipeline bind point, with a
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) that is compatible for push
  constants, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) used to
  create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) or the
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) array used to
  create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) , as described in
  [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)

- <a href="#VUID-vkCmdTraceRaysKHR-maintenance4-08602"
  id="VUID-vkCmdTraceRaysKHR-maintenance4-08602"></a>
  VUID-vkCmdTraceRaysKHR-maintenance4-08602  
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

- <a href="#VUID-vkCmdTraceRaysKHR-None-08114"
  id="VUID-vkCmdTraceRaysKHR-None-08114"></a>
  VUID-vkCmdTraceRaysKHR-None-08114  
  Descriptors in each bound descriptor set, specified via
  [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets.html), **must** be
  valid as described by [descriptor validity](#descriptor-validity) if
  they are statically used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) bound to
  the pipeline bind point used by this command and the bound
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) was not created with
  `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdTraceRaysKHR-None-08115"
  id="VUID-vkCmdTraceRaysKHR-None-08115"></a>
  VUID-vkCmdTraceRaysKHR-None-08115  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) bound to
  the pipeline bind point were specified via
  [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets.html), the bound
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) **must** have been created without
  `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdTraceRaysKHR-None-08116"
  id="VUID-vkCmdTraceRaysKHR-None-08116"></a>
  VUID-vkCmdTraceRaysKHR-None-08116  
  Descriptors in bound descriptor buffers, specified via
  [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html),
  **must** be valid if they are dynamically used by the
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) bound to the pipeline bind point used by
  this command and the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) was created
  with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdTraceRaysKHR-None-08604"
  id="VUID-vkCmdTraceRaysKHR-None-08604"></a>
  VUID-vkCmdTraceRaysKHR-None-08604  
  Descriptors in bound descriptor buffers, specified via
  [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html),
  **must** be valid if they are dynamically used by any
  [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) bound to a stage corresponding to the
  pipeline bind point used by this command

- <a href="#VUID-vkCmdTraceRaysKHR-None-08117"
  id="VUID-vkCmdTraceRaysKHR-None-08117"></a>
  VUID-vkCmdTraceRaysKHR-None-08117  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) bound to
  the pipeline bind point were specified via
  [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html),
  the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) **must** have been created
  with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdTraceRaysKHR-None-08119"
  id="VUID-vkCmdTraceRaysKHR-None-08119"></a>
  VUID-vkCmdTraceRaysKHR-None-08119  
  If a descriptor is dynamically used with a
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) created with
  `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, the descriptor memory
  **must** be resident

- <a href="#VUID-vkCmdTraceRaysKHR-None-08605"
  id="VUID-vkCmdTraceRaysKHR-None-08605"></a>
  VUID-vkCmdTraceRaysKHR-None-08605  
  If a descriptor is dynamically used with a
  [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) created with a `VkDescriptorSetLayout`
  that was created with
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, the
  descriptor memory **must** be resident

- <a href="#VUID-vkCmdTraceRaysKHR-None-08606"
  id="VUID-vkCmdTraceRaysKHR-None-08606"></a>
  VUID-vkCmdTraceRaysKHR-None-08606  
  If the [`shaderObject`](#features-shaderObject) feature is not
  enabled, a valid pipeline **must** be bound to the pipeline bind point
  used by this command

- <a href="#VUID-vkCmdTraceRaysKHR-None-08608"
  id="VUID-vkCmdTraceRaysKHR-None-08608"></a>
  VUID-vkCmdTraceRaysKHR-None-08608  
  If a pipeline is bound to the pipeline bind point used by this
  command, there **must** not have been any calls to dynamic state
  setting commands for any state not specified as dynamic in the
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object bound to the pipeline bind point
  used by this command, since that pipeline was bound

- <a href="#VUID-vkCmdTraceRaysKHR-None-08609"
  id="VUID-vkCmdTraceRaysKHR-None-08609"></a>
  VUID-vkCmdTraceRaysKHR-None-08609  
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

- <a href="#VUID-vkCmdTraceRaysKHR-None-08610"
  id="VUID-vkCmdTraceRaysKHR-None-08610"></a>
  VUID-vkCmdTraceRaysKHR-None-08610  
  If the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object bound to the pipeline bind
  point used by this command or any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html)
  bound to a stage corresponding to the pipeline bind point used by this
  command accesses a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) object that uses
  unnormalized coordinates, that sampler **must** not be used with any
  of the SPIR-V `OpImageSample*` or `OpImageSparseSample*` instructions
  with `ImplicitLod`, `Dref` or `Proj` in their name, in any shader
  stage

- <a href="#VUID-vkCmdTraceRaysKHR-None-08611"
  id="VUID-vkCmdTraceRaysKHR-None-08611"></a>
  VUID-vkCmdTraceRaysKHR-None-08611  
  If the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object bound to the pipeline bind
  point used by this command or any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html)
  bound to a stage corresponding to the pipeline bind point used by this
  command accesses a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) object that uses
  unnormalized coordinates, that sampler **must** not be used with any
  of the SPIR-V `OpImageSample*` or `OpImageSparseSample*` instructions
  that includes a LOD bias or any offset values, in any shader stage

- <a href="#VUID-vkCmdTraceRaysKHR-None-08607"
  id="VUID-vkCmdTraceRaysKHR-None-08607"></a>
  VUID-vkCmdTraceRaysKHR-None-08607  
  If the [`shaderObject`](#features-shaderObject) is enabled, either a
  valid pipeline **must** be bound to the pipeline bind point used by
  this command, or a valid combination of valid and
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) shader objects **must** be bound
  to every supported shader stage corresponding to the pipeline bind
  point used by this command

- <a href="#VUID-vkCmdTraceRaysKHR-uniformBuffers-06935"
  id="VUID-vkCmdTraceRaysKHR-uniformBuffers-06935"></a>
  VUID-vkCmdTraceRaysKHR-uniformBuffers-06935  
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

- <a href="#VUID-vkCmdTraceRaysKHR-None-08612"
  id="VUID-vkCmdTraceRaysKHR-None-08612"></a>
  VUID-vkCmdTraceRaysKHR-None-08612  
  If the [`robustBufferAccess`](#features-robustBufferAccess) feature is
  not enabled, and any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) bound to a stage
  corresponding to the pipeline bind point used by this command accesses
  a uniform buffer, it **must** not access values outside of the range
  of the buffer as specified in the descriptor set bound to the same
  pipeline bind point

- <a href="#VUID-vkCmdTraceRaysKHR-storageBuffers-06936"
  id="VUID-vkCmdTraceRaysKHR-storageBuffers-06936"></a>
  VUID-vkCmdTraceRaysKHR-storageBuffers-06936  
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

- <a href="#VUID-vkCmdTraceRaysKHR-None-08613"
  id="VUID-vkCmdTraceRaysKHR-None-08613"></a>
  VUID-vkCmdTraceRaysKHR-None-08613  
  If the [`robustBufferAccess`](#features-robustBufferAccess) feature is
  not enabled, and any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) bound to a stage
  corresponding to the pipeline bind point used by this command accesses
  a storage buffer, it **must** not access values outside of the range
  of the buffer as specified in the descriptor set bound to the same
  pipeline bind point

- <a href="#VUID-vkCmdTraceRaysKHR-commandBuffer-02707"
  id="VUID-vkCmdTraceRaysKHR-commandBuffer-02707"></a>
  VUID-vkCmdTraceRaysKHR-commandBuffer-02707  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported, any
  resource accessed by [bound shaders](#shaders-binding) **must** not be
  a protected resource

- <a href="#VUID-vkCmdTraceRaysKHR-None-06550"
  id="VUID-vkCmdTraceRaysKHR-None-06550"></a>
  VUID-vkCmdTraceRaysKHR-None-06550  
  If [a bound shader](#shaders-binding) accesses a
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) or [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) object
  that enables [sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion](#samplers-YCbCr-conversion), that object **must** only be
  used with `OpImageSample*` or `OpImageSparseSample*` instructions

- <a href="#VUID-vkCmdTraceRaysKHR-ConstOffset-06551"
  id="VUID-vkCmdTraceRaysKHR-ConstOffset-06551"></a>
  VUID-vkCmdTraceRaysKHR-ConstOffset-06551  
  If [a bound shader](#shaders-binding) accesses a
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) or [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) object
  that enables [sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion](#samplers-YCbCr-conversion), that object **must** not use
  the `ConstOffset` and `Offset` operands

- <a href="#VUID-vkCmdTraceRaysKHR-viewType-07752"
  id="VUID-vkCmdTraceRaysKHR-viewType-07752"></a>
  VUID-vkCmdTraceRaysKHR-viewType-07752  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is accessed as a result of this
  command, then the image view’s `viewType` **must** match the `Dim`
  operand of the `OpTypeImage` as described in
  [\[textures-operation-validation\]](#textures-operation-validation)

- <a href="#VUID-vkCmdTraceRaysKHR-format-07753"
  id="VUID-vkCmdTraceRaysKHR-format-07753"></a>
  VUID-vkCmdTraceRaysKHR-format-07753  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is accessed as a result of this
  command, then the [numeric type](#formats-numericformat) of the image
  view’s `format` and the `Sampled` `Type` operand of the `OpTypeImage`
  **must** match

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageWrite-08795"
  id="VUID-vkCmdTraceRaysKHR-OpImageWrite-08795"></a>
  VUID-vkCmdTraceRaysKHR-OpImageWrite-08795  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) created with a format other than
  `VK_FORMAT_A8_UNORM_KHR` is accessed using `OpImageWrite` as a result
  of this command, then the `Type` of the `Texel` operand of that
  instruction **must** have at least as many components as the image
  view’s format

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageWrite-08796"
  id="VUID-vkCmdTraceRaysKHR-OpImageWrite-08796"></a>
  VUID-vkCmdTraceRaysKHR-OpImageWrite-08796  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) created with the format
  `VK_FORMAT_A8_UNORM_KHR` is accessed using `OpImageWrite` as a result
  of this command, then the `Type` of the `Texel` operand of that
  instruction **must** have four components

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageWrite-04469"
  id="VUID-vkCmdTraceRaysKHR-OpImageWrite-04469"></a>
  VUID-vkCmdTraceRaysKHR-OpImageWrite-04469  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) is accessed using
  `OpImageWrite` as a result of this command, then the `Type` of the
  `Texel` operand of that instruction **must** have at least as many
  components as the buffer view’s format

- <a href="#VUID-vkCmdTraceRaysKHR-SampledType-04470"
  id="VUID-vkCmdTraceRaysKHR-SampledType-04470"></a>
  VUID-vkCmdTraceRaysKHR-SampledType-04470  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)
  that has a 64-bit component width is accessed as a result of this
  command, the `SampledType` of the `OpTypeImage` operand of that
  instruction **must** have a `Width` of 64

- <a href="#VUID-vkCmdTraceRaysKHR-SampledType-04471"
  id="VUID-vkCmdTraceRaysKHR-SampledType-04471"></a>
  VUID-vkCmdTraceRaysKHR-SampledType-04471  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)
  that has a component width less than 64-bit is accessed as a result of
  this command, the `SampledType` of the `OpTypeImage` operand of that
  instruction **must** have a `Width` of 32

- <a href="#VUID-vkCmdTraceRaysKHR-SampledType-04472"
  id="VUID-vkCmdTraceRaysKHR-SampledType-04472"></a>
  VUID-vkCmdTraceRaysKHR-SampledType-04472  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) with a
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) that has a 64-bit component width is
  accessed as a result of this command, the `SampledType` of the
  `OpTypeImage` operand of that instruction **must** have a `Width` of
  64

- <a href="#VUID-vkCmdTraceRaysKHR-SampledType-04473"
  id="VUID-vkCmdTraceRaysKHR-SampledType-04473"></a>
  VUID-vkCmdTraceRaysKHR-SampledType-04473  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) with a
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) that has a component width less than 64-bit
  is accessed as a result of this command, the `SampledType` of the
  `OpTypeImage` operand of that instruction **must** have a `Width` of
  32

- <a href="#VUID-vkCmdTraceRaysKHR-sparseImageInt64Atomics-04474"
  id="VUID-vkCmdTraceRaysKHR-sparseImageInt64Atomics-04474"></a>
  VUID-vkCmdTraceRaysKHR-sparseImageInt64Atomics-04474  
  If the [`sparseImageInt64Atomics`](#features-sparseImageInt64Atomics)
  feature is not enabled, [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) objects created with
  the `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` flag **must** not be
  accessed by atomic instructions through an `OpTypeImage` with a
  `SampledType` with a `Width` of 64 by this command

- <a href="#VUID-vkCmdTraceRaysKHR-sparseImageInt64Atomics-04475"
  id="VUID-vkCmdTraceRaysKHR-sparseImageInt64Atomics-04475"></a>
  VUID-vkCmdTraceRaysKHR-sparseImageInt64Atomics-04475  
  If the [`sparseImageInt64Atomics`](#features-sparseImageInt64Atomics)
  feature is not enabled, [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) objects created with
  the `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT` flag **must** not be
  accessed by atomic instructions through an `OpTypeImage` with a
  `SampledType` with a `Width` of 64 by this command

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageWeightedSampleQCOM-06971"
  id="VUID-vkCmdTraceRaysKHR-OpImageWeightedSampleQCOM-06971"></a>
  VUID-vkCmdTraceRaysKHR-OpImageWeightedSampleQCOM-06971  
  If `OpImageWeightedSampleQCOM` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_WEIGHT_SAMPLED_IMAGE_BIT_QCOM`

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageWeightedSampleQCOM-06972"
  id="VUID-vkCmdTraceRaysKHR-OpImageWeightedSampleQCOM-06972"></a>
  VUID-vkCmdTraceRaysKHR-OpImageWeightedSampleQCOM-06972  
  If `OpImageWeightedSampleQCOM` uses a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html)
  as a sample weight image as a result of this command, then the image
  view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_WEIGHT_IMAGE_BIT_QCOM`

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageBoxFilterQCOM-06973"
  id="VUID-vkCmdTraceRaysKHR-OpImageBoxFilterQCOM-06973"></a>
  VUID-vkCmdTraceRaysKHR-OpImageBoxFilterQCOM-06973  
  If `OpImageBoxFilterQCOM` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_BOX_FILTER_SAMPLED_BIT_QCOM`

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageBlockMatchSSDQCOM-06974"
  id="VUID-vkCmdTraceRaysKHR-OpImageBlockMatchSSDQCOM-06974"></a>
  VUID-vkCmdTraceRaysKHR-OpImageBlockMatchSSDQCOM-06974  
  If `OpImageBlockMatchSSDQCOM` is used to read from an
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageBlockMatchSADQCOM-06975"
  id="VUID-vkCmdTraceRaysKHR-OpImageBlockMatchSADQCOM-06975"></a>
  VUID-vkCmdTraceRaysKHR-OpImageBlockMatchSADQCOM-06975  
  If `OpImageBlockMatchSADQCOM` is used to read from an
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageBlockMatchSADQCOM-06976"
  id="VUID-vkCmdTraceRaysKHR-OpImageBlockMatchSADQCOM-06976"></a>
  VUID-vkCmdTraceRaysKHR-OpImageBlockMatchSADQCOM-06976  
  If `OpImageBlockMatchSADQCOM` or OpImageBlockMatchSSDQCOM is used to
  read from a reference image as result of this command, then the
  specified reference coordinates **must** not fail [integer texel
  coordinate validation](#textures-integer-coordinate-validation)

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageWeightedSampleQCOM-06977"
  id="VUID-vkCmdTraceRaysKHR-OpImageWeightedSampleQCOM-06977"></a>
  VUID-vkCmdTraceRaysKHR-OpImageWeightedSampleQCOM-06977  
  If `OpImageWeightedSampleQCOM`, `OpImageBoxFilterQCOM`,
  `OpImageBlockMatchWindowSSDQCOM`, `OpImageBlockMatchWindowSADQCOM`,
  `OpImageBlockMatchGatherSSDQCOM`, `OpImageBlockMatchGatherSADQCOM`,
  `OpImageBlockMatchSSDQCOM`, or `OpImageBlockMatchSADQCOM` uses a
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) as a result of this command, then the
  sampler **must** have been created with
  `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageWeightedSampleQCOM-06978"
  id="VUID-vkCmdTraceRaysKHR-OpImageWeightedSampleQCOM-06978"></a>
  VUID-vkCmdTraceRaysKHR-OpImageWeightedSampleQCOM-06978  
  If any command other than `OpImageWeightedSampleQCOM`,
  `OpImageBoxFilterQCOM`, `OpImageBlockMatchWindowSSDQCOM`,
  `OpImageBlockMatchWindowSADQCOM`, `OpImageBlockMatchGatherSSDQCOM`,
  `OpImageBlockMatchGatherSADQCOM`, `OpImageBlockMatchSSDQCOM`, or
  `OpImageBlockMatchSADQCOM` uses a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) as a
  result of this command, then the sampler **must** not have been
  created with `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageBlockMatchWindow-09215"
  id="VUID-vkCmdTraceRaysKHR-OpImageBlockMatchWindow-09215"></a>
  VUID-vkCmdTraceRaysKHR-OpImageBlockMatchWindow-09215  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  instruction is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as
  a result of this command, then the image view’s [format
  features](#resources-image-view-format-features) **must** contain
  `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageBlockMatchWindow-09216"
  id="VUID-vkCmdTraceRaysKHR-OpImageBlockMatchWindow-09216"></a>
  VUID-vkCmdTraceRaysKHR-OpImageBlockMatchWindow-09216  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  instruction is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as
  a result of this command, then the image view’s format **must** be a
  single-component format

- <a href="#VUID-vkCmdTraceRaysKHR-OpImageBlockMatchWindow-09217"
  id="VUID-vkCmdTraceRaysKHR-OpImageBlockMatchWindow-09217"></a>
  VUID-vkCmdTraceRaysKHR-OpImageBlockMatchWindow-09217  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  read from a reference image as result of this command, then the
  specified reference coordinates **must** not fail [integer texel
  coordinate validation](#textures-integer-coordinate-validation)

- <a href="#VUID-vkCmdTraceRaysKHR-None-07288"
  id="VUID-vkCmdTraceRaysKHR-None-07288"></a>
  VUID-vkCmdTraceRaysKHR-None-07288  
  Any shader invocation executed by this command **must**
  [terminate](#shaders-termination)

- <a href="#VUID-vkCmdTraceRaysKHR-None-09600"
  id="VUID-vkCmdTraceRaysKHR-None-09600"></a>
  VUID-vkCmdTraceRaysKHR-None-09600  
  If a descriptor with type equal to any of
  `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM`,
  `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`,
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` is accessed as a result of this
  command, the image subresource identified by that descriptor **must**
  be in the image layout identified when the descriptor was written

- <a href="#VUID-vkCmdTraceRaysKHR-None-03429"
  id="VUID-vkCmdTraceRaysKHR-None-03429"></a>
  VUID-vkCmdTraceRaysKHR-None-03429  
  Any shader group handle referenced by this call **must** have been
  queried from the currently bound ray tracing pipeline

- <a href="#VUID-vkCmdTraceRaysKHR-None-09458"
  id="VUID-vkCmdTraceRaysKHR-None-09458"></a>
  VUID-vkCmdTraceRaysKHR-None-09458  
  If the bound ray tracing pipeline state was created with the
  `VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR` dynamic state
  enabled then
  [vkCmdSetRayTracingPipelineStackSizeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRayTracingPipelineStackSizeKHR.html)
  **must** have been called in the current command buffer prior to this
  trace command

<!-- -->

- <a href="#VUID-vkCmdTraceRaysKHR-maxPipelineRayRecursionDepth-03679"
  id="VUID-vkCmdTraceRaysKHR-maxPipelineRayRecursionDepth-03679"></a>
  VUID-vkCmdTraceRaysKHR-maxPipelineRayRecursionDepth-03679  
  This command **must** not cause a shader call instruction to be
  executed from a shader invocation with a [recursion
  depth](#ray-tracing-recursion-depth) greater than the value of
  `maxPipelineRayRecursionDepth` used to create the bound ray tracing
  pipeline

- <a href="#VUID-vkCmdTraceRaysKHR-commandBuffer-03635"
  id="VUID-vkCmdTraceRaysKHR-commandBuffer-03635"></a>
  VUID-vkCmdTraceRaysKHR-commandBuffer-03635  
  `commandBuffer` **must** not be a protected command buffer

<!-- -->

- <a href="#VUID-vkCmdTraceRaysKHR-size-04023"
  id="VUID-vkCmdTraceRaysKHR-size-04023"></a>
  VUID-vkCmdTraceRaysKHR-size-04023  
  The `size` member of `pRayGenShaderBindingTable` **must** be equal to
  its `stride` member

<!-- -->

- <a href="#VUID-vkCmdTraceRaysKHR-pRayGenShaderBindingTable-03680"
  id="VUID-vkCmdTraceRaysKHR-pRayGenShaderBindingTable-03680"></a>
  VUID-vkCmdTraceRaysKHR-pRayGenShaderBindingTable-03680  
  If the buffer from which `pRayGenShaderBindingTable->deviceAddress`
  was queried is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdTraceRaysKHR-pRayGenShaderBindingTable-03681"
  id="VUID-vkCmdTraceRaysKHR-pRayGenShaderBindingTable-03681"></a>
  VUID-vkCmdTraceRaysKHR-pRayGenShaderBindingTable-03681  
  The buffer from which the `pRayGenShaderBindingTable->deviceAddress`
  is queried **must** have been created with the
  `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag

- <a href="#VUID-vkCmdTraceRaysKHR-pRayGenShaderBindingTable-03682"
  id="VUID-vkCmdTraceRaysKHR-pRayGenShaderBindingTable-03682"></a>
  VUID-vkCmdTraceRaysKHR-pRayGenShaderBindingTable-03682  
  `pRayGenShaderBindingTable->deviceAddress` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`

- <a href="#VUID-vkCmdTraceRaysKHR-pMissShaderBindingTable-03683"
  id="VUID-vkCmdTraceRaysKHR-pMissShaderBindingTable-03683"></a>
  VUID-vkCmdTraceRaysKHR-pMissShaderBindingTable-03683  
  If the buffer from which `pMissShaderBindingTable->deviceAddress` was
  queried is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdTraceRaysKHR-pMissShaderBindingTable-03684"
  id="VUID-vkCmdTraceRaysKHR-pMissShaderBindingTable-03684"></a>
  VUID-vkCmdTraceRaysKHR-pMissShaderBindingTable-03684  
  The buffer from which the `pMissShaderBindingTable->deviceAddress` is
  queried **must** have been created with the
  `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag

- <a href="#VUID-vkCmdTraceRaysKHR-pMissShaderBindingTable-03685"
  id="VUID-vkCmdTraceRaysKHR-pMissShaderBindingTable-03685"></a>
  VUID-vkCmdTraceRaysKHR-pMissShaderBindingTable-03685  
  `pMissShaderBindingTable->deviceAddress` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`

- <a href="#VUID-vkCmdTraceRaysKHR-stride-03686"
  id="VUID-vkCmdTraceRaysKHR-stride-03686"></a>
  VUID-vkCmdTraceRaysKHR-stride-03686  
  `pMissShaderBindingTable->stride` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupHandleAlignment`

- <a href="#VUID-vkCmdTraceRaysKHR-stride-04029"
  id="VUID-vkCmdTraceRaysKHR-stride-04029"></a>
  VUID-vkCmdTraceRaysKHR-stride-04029  
  `pMissShaderBindingTable->stride` **must** be less than or equal to
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxShaderGroupStride`

- <a href="#VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-03687"
  id="VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-03687"></a>
  VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-03687  
  If the buffer from which `pHitShaderBindingTable->deviceAddress` was
  queried is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-03688"
  id="VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-03688"></a>
  VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-03688  
  The buffer from which the `pHitShaderBindingTable->deviceAddress` is
  queried **must** have been created with the
  `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag

- <a href="#VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-03689"
  id="VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-03689"></a>
  VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-03689  
  `pHitShaderBindingTable->deviceAddress` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`

- <a href="#VUID-vkCmdTraceRaysKHR-stride-03690"
  id="VUID-vkCmdTraceRaysKHR-stride-03690"></a>
  VUID-vkCmdTraceRaysKHR-stride-03690  
  `pHitShaderBindingTable->stride` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupHandleAlignment`

- <a href="#VUID-vkCmdTraceRaysKHR-stride-04035"
  id="VUID-vkCmdTraceRaysKHR-stride-04035"></a>
  VUID-vkCmdTraceRaysKHR-stride-04035  
  `pHitShaderBindingTable->stride` **must** be less than or equal to
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxShaderGroupStride`

- <a href="#VUID-vkCmdTraceRaysKHR-pCallableShaderBindingTable-03691"
  id="VUID-vkCmdTraceRaysKHR-pCallableShaderBindingTable-03691"></a>
  VUID-vkCmdTraceRaysKHR-pCallableShaderBindingTable-03691  
  If the buffer from which `pCallableShaderBindingTable->deviceAddress`
  was queried is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdTraceRaysKHR-pCallableShaderBindingTable-03692"
  id="VUID-vkCmdTraceRaysKHR-pCallableShaderBindingTable-03692"></a>
  VUID-vkCmdTraceRaysKHR-pCallableShaderBindingTable-03692  
  The buffer from which the `pCallableShaderBindingTable->deviceAddress`
  is queried **must** have been created with the
  `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` usage flag

- <a href="#VUID-vkCmdTraceRaysKHR-pCallableShaderBindingTable-03693"
  id="VUID-vkCmdTraceRaysKHR-pCallableShaderBindingTable-03693"></a>
  VUID-vkCmdTraceRaysKHR-pCallableShaderBindingTable-03693  
  `pCallableShaderBindingTable->deviceAddress` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupBaseAlignment`

- <a href="#VUID-vkCmdTraceRaysKHR-stride-03694"
  id="VUID-vkCmdTraceRaysKHR-stride-03694"></a>
  VUID-vkCmdTraceRaysKHR-stride-03694  
  `pCallableShaderBindingTable->stride` **must** be a multiple of
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`shaderGroupHandleAlignment`

- <a href="#VUID-vkCmdTraceRaysKHR-stride-04041"
  id="VUID-vkCmdTraceRaysKHR-stride-04041"></a>
  VUID-vkCmdTraceRaysKHR-stride-04041  
  `pCallableShaderBindingTable->stride` **must** be less than or equal
  to
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxShaderGroupStride`

- <a href="#VUID-vkCmdTraceRaysKHR-flags-03696"
  id="VUID-vkCmdTraceRaysKHR-flags-03696"></a>
  VUID-vkCmdTraceRaysKHR-flags-03696  
  If the currently bound ray tracing pipeline was created with `flags`
  that included
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`,
  `pHitShaderBindingTable->deviceAddress` **must** not be zero

- <a href="#VUID-vkCmdTraceRaysKHR-flags-03697"
  id="VUID-vkCmdTraceRaysKHR-flags-03697"></a>
  VUID-vkCmdTraceRaysKHR-flags-03697  
  If the currently bound ray tracing pipeline was created with `flags`
  that included
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`,
  `pHitShaderBindingTable->deviceAddress` **must** not be zero

- <a href="#VUID-vkCmdTraceRaysKHR-flags-03511"
  id="VUID-vkCmdTraceRaysKHR-flags-03511"></a>
  VUID-vkCmdTraceRaysKHR-flags-03511  
  If the currently bound ray tracing pipeline was created with `flags`
  that included
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`, the
  shader group handle identified by
  `pMissShaderBindingTable->deviceAddress` **must** not be set to zero

- <a href="#VUID-vkCmdTraceRaysKHR-flags-03512"
  id="VUID-vkCmdTraceRaysKHR-flags-03512"></a>
  VUID-vkCmdTraceRaysKHR-flags-03512  
  If the currently bound ray tracing pipeline was created with `flags`
  that included
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`,
  entries in the table identified by
  `pHitShaderBindingTable->deviceAddress` accessed as a result of this
  command in order to execute an any-hit shader **must** not be set to
  zero

- <a href="#VUID-vkCmdTraceRaysKHR-flags-03513"
  id="VUID-vkCmdTraceRaysKHR-flags-03513"></a>
  VUID-vkCmdTraceRaysKHR-flags-03513  
  If the currently bound ray tracing pipeline was created with `flags`
  that included
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`,
  entries in the table identified by
  `pHitShaderBindingTable->deviceAddress` accessed as a result of this
  command in order to execute a closest hit shader **must** not be set
  to zero

- <a href="#VUID-vkCmdTraceRaysKHR-flags-03514"
  id="VUID-vkCmdTraceRaysKHR-flags-03514"></a>
  VUID-vkCmdTraceRaysKHR-flags-03514  
  If the currently bound ray tracing pipeline was created with `flags`
  that included
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`,
  entries in the table identified by
  `pHitShaderBindingTable->deviceAddress` accessed as a result of this
  command in order to execute an intersection shader **must** not be set
  to zero

- <a href="#VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-04735"
  id="VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-04735"></a>
  VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-04735  
  Any non-zero hit shader group entries in the table identified by
  `pHitShaderBindingTable->deviceAddress` accessed by this call from a
  geometry with a `geometryType` of `VK_GEOMETRY_TYPE_TRIANGLES_KHR`
  **must** have been created with
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR`

- <a href="#VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-04736"
  id="VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-04736"></a>
  VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-04736  
  Any non-zero hit shader group entries in the table identified by
  `pHitShaderBindingTable->deviceAddress` accessed by this call from a
  geometry with a `geometryType` of `VK_GEOMETRY_TYPE_AABBS_KHR`
  **must** have been created with
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR`

<!-- -->

- <a href="#VUID-vkCmdTraceRaysKHR-width-03638"
  id="VUID-vkCmdTraceRaysKHR-width-03638"></a>
  VUID-vkCmdTraceRaysKHR-width-03638  
  `width` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[0\] ×
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[0\]

- <a href="#VUID-vkCmdTraceRaysKHR-height-03639"
  id="VUID-vkCmdTraceRaysKHR-height-03639"></a>
  VUID-vkCmdTraceRaysKHR-height-03639  
  `height` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[1\] ×
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[1\]

- <a href="#VUID-vkCmdTraceRaysKHR-depth-03640"
  id="VUID-vkCmdTraceRaysKHR-depth-03640"></a>
  VUID-vkCmdTraceRaysKHR-depth-03640  
  `depth` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[2\] ×
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[2\]

- <a href="#VUID-vkCmdTraceRaysKHR-width-03641"
  id="VUID-vkCmdTraceRaysKHR-width-03641"></a>
  VUID-vkCmdTraceRaysKHR-width-03641  
  `width` × `height` × `depth` **must** be less than or equal to
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxRayDispatchInvocationCount`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdTraceRaysKHR-commandBuffer-parameter"
  id="VUID-vkCmdTraceRaysKHR-commandBuffer-parameter"></a>
  VUID-vkCmdTraceRaysKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdTraceRaysKHR-pRaygenShaderBindingTable-parameter"
  id="VUID-vkCmdTraceRaysKHR-pRaygenShaderBindingTable-parameter"></a>
  VUID-vkCmdTraceRaysKHR-pRaygenShaderBindingTable-parameter  
  `pRaygenShaderBindingTable` **must** be a valid pointer to a valid
  [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html)
  structure

- <a href="#VUID-vkCmdTraceRaysKHR-pMissShaderBindingTable-parameter"
  id="VUID-vkCmdTraceRaysKHR-pMissShaderBindingTable-parameter"></a>
  VUID-vkCmdTraceRaysKHR-pMissShaderBindingTable-parameter  
  `pMissShaderBindingTable` **must** be a valid pointer to a valid
  [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html)
  structure

- <a href="#VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-parameter"
  id="VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-parameter"></a>
  VUID-vkCmdTraceRaysKHR-pHitShaderBindingTable-parameter  
  `pHitShaderBindingTable` **must** be a valid pointer to a valid
  [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html)
  structure

- <a href="#VUID-vkCmdTraceRaysKHR-pCallableShaderBindingTable-parameter"
  id="VUID-vkCmdTraceRaysKHR-pCallableShaderBindingTable-parameter"></a>
  VUID-vkCmdTraceRaysKHR-pCallableShaderBindingTable-parameter  
  `pCallableShaderBindingTable` **must** be a valid pointer to a valid
  [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html)
  structure

- <a href="#VUID-vkCmdTraceRaysKHR-commandBuffer-recording"
  id="VUID-vkCmdTraceRaysKHR-commandBuffer-recording"></a>
  VUID-vkCmdTraceRaysKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdTraceRaysKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdTraceRaysKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdTraceRaysKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdTraceRaysKHR-renderpass"
  id="VUID-vkCmdTraceRaysKHR-renderpass"></a>
  VUID-vkCmdTraceRaysKHR-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdTraceRaysKHR-videocoding"
  id="VUID-vkCmdTraceRaysKHR-videocoding"></a>
  VUID-vkCmdTraceRaysKHR-videocoding  
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
<tr>
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
<tr>
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdTraceRaysKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
