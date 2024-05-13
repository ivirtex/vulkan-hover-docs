# vkCmdDrawMeshTasksEXT(3) Manual Page

## Name

vkCmdDrawMeshTasksEXT - Draw mesh task work items



## <a href="#_c_specification" class="anchor"></a>C Specification

To record a mesh tasks drawing command, call:

``` c
// Provided by VK_EXT_mesh_shader
void vkCmdDrawMeshTasksEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    groupCountX,
    uint32_t                                    groupCountY,
    uint32_t                                    groupCountZ);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `groupCountX` is the number of local workgroups to dispatch in the X
  dimension.

- `groupCountY` is the number of local workgroups to dispatch in the Y
  dimension.

- `groupCountZ` is the number of local workgroups to dispatch in the Z
  dimension.

## <a href="#_description" class="anchor"></a>Description

When the command is executed, a global workgroup consisting of
`groupCountX` × `groupCountY` × `groupCountZ` local workgroups is
assembled.

Valid Usage

- <a href="#VUID-vkCmdDrawMeshTasksEXT-magFilter-04553"
  id="VUID-vkCmdDrawMeshTasksEXT-magFilter-04553"></a>
  VUID-vkCmdDrawMeshTasksEXT-magFilter-04553  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with `magFilter` or
  `minFilter` equal to `VK_FILTER_LINEAR`, `reductionMode` equal to
  `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`, and `compareEnable`
  equal to `VK_FALSE` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-magFilter-09598"
  id="VUID-vkCmdDrawMeshTasksEXT-magFilter-09598"></a>
  VUID-vkCmdDrawMeshTasksEXT-magFilter-09598  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with `magFilter` or
  `minFilter` equal to `VK_FILTER_LINEAR` and `reductionMode` equal to
  either `VK_SAMPLER_REDUCTION_MODE_MIN` or
  `VK_SAMPLER_REDUCTION_MODE_MAX` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-mipmapMode-04770"
  id="VUID-vkCmdDrawMeshTasksEXT-mipmapMode-04770"></a>
  VUID-vkCmdDrawMeshTasksEXT-mipmapMode-04770  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with `mipmapMode` equal to
  `VK_SAMPLER_MIPMAP_MODE_LINEAR`, `reductionMode` equal to
  `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`, and `compareEnable`
  equal to `VK_FALSE` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-mipmapMode-09599"
  id="VUID-vkCmdDrawMeshTasksEXT-mipmapMode-09599"></a>
  VUID-vkCmdDrawMeshTasksEXT-mipmapMode-09599  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with `mipmapMode` equal to
  `VK_SAMPLER_MIPMAP_MODE_LINEAR` and `reductionMode` equal to either
  `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` is
  used to sample a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this
  command, then the image view’s [format
  features](#resources-image-view-format-features) **must** contain
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-unnormalizedCoordinates-09635"
  id="VUID-vkCmdDrawMeshTasksEXT-unnormalizedCoordinates-09635"></a>
  VUID-vkCmdDrawMeshTasksEXT-unnormalizedCoordinates-09635  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with
  `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s `levelCount` and `layerCount` **must** be 1

- <a href="#VUID-vkCmdDrawMeshTasksEXT-unnormalizedCoordinates-09636"
  id="VUID-vkCmdDrawMeshTasksEXT-unnormalizedCoordinates-09636"></a>
  VUID-vkCmdDrawMeshTasksEXT-unnormalizedCoordinates-09636  
  If a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) created with
  `unnormalizedCoordinates` equal to `VK_TRUE` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s `viewType` **must** be `VK_IMAGE_VIEW_TYPE_1D` or
  `VK_IMAGE_VIEW_TYPE_2D`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-06479"
  id="VUID-vkCmdDrawMeshTasksEXT-None-06479"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-06479  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is sampled with [depth
  comparison](#textures-depth-compare-operation), the image view’s
  [format features](#resources-image-view-format-features) **must**
  contain `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_DEPTH_COMPARISON_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-02691"
  id="VUID-vkCmdDrawMeshTasksEXT-None-02691"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-02691  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is accessed using atomic
  operations as a result of this command, then the image view’s [format
  features](#resources-image-view-format-features) **must** contain
  `VK_FORMAT_FEATURE_STORAGE_IMAGE_ATOMIC_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07888"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07888"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07888  
  If a `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor is accessed
  using atomic operations as a result of this command, then the storage
  texel buffer’s [format
  features](#resources-buffer-view-format-features) **must** contain
  `VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_ATOMIC_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-02692"
  id="VUID-vkCmdDrawMeshTasksEXT-None-02692"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-02692  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is sampled with
  `VK_FILTER_CUBIC_EXT` as a result of this command, then the image
  view’s [format features](#resources-image-view-format-features)
  **must** contain
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-02693"
  id="VUID-vkCmdDrawMeshTasksEXT-None-02693"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-02693  
  If the [VK_EXT_filter_cubic](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_filter_cubic.html) extension is
  not enabled and any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is sampled with
  `VK_FILTER_CUBIC_EXT` as a result of this command, it **must** not
  have a [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html) of
  `VK_IMAGE_VIEW_TYPE_3D`, `VK_IMAGE_VIEW_TYPE_CUBE`, or
  `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-filterCubic-02694"
  id="VUID-vkCmdDrawMeshTasksEXT-filterCubic-02694"></a>
  VUID-vkCmdDrawMeshTasksEXT-filterCubic-02694  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being sampled with
  `VK_FILTER_CUBIC_EXT` as a result of this command **must** have a
  [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html) and format that supports cubic
  filtering, as specified by
  [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)::`filterCubic`
  returned by
  [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-filterCubicMinmax-02695"
  id="VUID-vkCmdDrawMeshTasksEXT-filterCubicMinmax-02695"></a>
  VUID-vkCmdDrawMeshTasksEXT-filterCubicMinmax-02695  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being sampled with
  `VK_FILTER_CUBIC_EXT` with a reduction mode of either
  `VK_SAMPLER_REDUCTION_MODE_MIN` or `VK_SAMPLER_REDUCTION_MODE_MAX` as
  a result of this command **must** have a
  [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html) and format that supports cubic
  filtering together with minmax filtering, as specified by
  [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html)::`filterCubicMinmax`
  returned by
  [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-cubicRangeClamp-09212"
  id="VUID-vkCmdDrawMeshTasksEXT-cubicRangeClamp-09212"></a>
  VUID-vkCmdDrawMeshTasksEXT-cubicRangeClamp-09212  
  If the [`cubicRangeClamp`](#features-filter-cubic-range-clamp) feature
  is not enabled, then any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being sampled
  with `VK_FILTER_CUBIC_EXT` as a result of this command **must** not
  have a
  [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfo.html)::`reductionMode`
  equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-reductionMode-09213"
  id="VUID-vkCmdDrawMeshTasksEXT-reductionMode-09213"></a>
  VUID-vkCmdDrawMeshTasksEXT-reductionMode-09213  
  Any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being sampled with a
  [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfo.html)::`reductionMode`
  equal to `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM`
  as a result of this command **must** sample with `VK_FILTER_CUBIC_EXT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-selectableCubicWeights-09214"
  id="VUID-vkCmdDrawMeshTasksEXT-selectableCubicWeights-09214"></a>
  VUID-vkCmdDrawMeshTasksEXT-selectableCubicWeights-09214  
  If the
  [`selectableCubicWeights`](#features-filter-cubic-weight-selection)
  feature is not enabled, then any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being
  sampled with `VK_FILTER_CUBIC_EXT` as a result of this command
  **must** have
  [VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html)::`cubicWeights`
  equal to `VK_CUBIC_FILTER_WEIGHTS_CATMULL_ROM_QCOM`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-flags-02696"
  id="VUID-vkCmdDrawMeshTasksEXT-flags-02696"></a>
  VUID-vkCmdDrawMeshTasksEXT-flags-02696  
  Any [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) created with a
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags` containing
  `VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV` sampled as a result of this
  command **must** only be sampled using a
  [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html) of
  `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07027"
  id="VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07027"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07027  
  For any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being written as a storage
  image where the image format field of the `OpTypeImage` is `Unknown`,
  the view’s [format features](#resources-image-view-format-features)
  **must** contain
  `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07028"
  id="VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07028"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07028  
  For any [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) being read as a storage image
  where the image format field of the `OpTypeImage` is `Unknown`, the
  view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07029"
  id="VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07029"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07029  
  For any [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) being written as a storage
  texel buffer where the image format field of the `OpTypeImage` is
  `Unknown`, the view’s [buffer features](#VkFormatProperties3) **must**
  contain `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07030"
  id="VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07030"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07030  
  Any [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) being read as a storage texel
  buffer where the image format field of the `OpTypeImage` is `Unknown`
  then the view’s [buffer features](#VkFormatProperties3) **must**
  contain `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08600"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08600"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08600  
  For each set *n* that is statically used by [a bound
  shader](#shaders-binding), a descriptor set **must** have been bound
  to *n* at the same pipeline bind point, with a
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) that is compatible for set
  *n*, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) used to create
  the current [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) or the
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) array used to
  create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) , as described in
  [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08601"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08601"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08601  
  For each push constant that is statically used by [a bound
  shader](#shaders-binding), a push constant value **must** have been
  set for the same pipeline bind point, with a
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) that is compatible for push
  constants, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) used to
  create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) or the
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) array used to
  create the current [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) , as described in
  [\[descriptorsets-compatibility\]](#descriptorsets-compatibility)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-maintenance4-08602"
  id="VUID-vkCmdDrawMeshTasksEXT-maintenance4-08602"></a>
  VUID-vkCmdDrawMeshTasksEXT-maintenance4-08602  
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

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08114"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08114"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08114  
  Descriptors in each bound descriptor set, specified via
  [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets.html), **must** be
  valid as described by [descriptor validity](#descriptor-validity) if
  they are statically used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) bound to
  the pipeline bind point used by this command and the bound
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) was not created with
  `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08115"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08115"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08115  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) bound to
  the pipeline bind point were specified via
  [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets.html), the bound
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) **must** have been created without
  `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08116"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08116"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08116  
  Descriptors in bound descriptor buffers, specified via
  [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html),
  **must** be valid if they are dynamically used by the
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) bound to the pipeline bind point used by
  this command and the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) was created
  with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08604"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08604"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08604  
  Descriptors in bound descriptor buffers, specified via
  [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html),
  **must** be valid if they are dynamically used by any
  [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) bound to a stage corresponding to the
  pipeline bind point used by this command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08117"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08117"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08117  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) bound to
  the pipeline bind point were specified via
  [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html),
  the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) **must** have been created
  with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08119"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08119"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08119  
  If a descriptor is dynamically used with a
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) created with
  `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, the descriptor memory
  **must** be resident

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08605"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08605"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08605  
  If a descriptor is dynamically used with a
  [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) created with a `VkDescriptorSetLayout`
  that was created with
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, the
  descriptor memory **must** be resident

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08606"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08606"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08606  
  If the [`shaderObject`](#features-shaderObject) feature is not
  enabled, a valid pipeline **must** be bound to the pipeline bind point
  used by this command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08608"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08608"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08608  
  If a pipeline is bound to the pipeline bind point used by this
  command, there **must** not have been any calls to dynamic state
  setting commands for any state not specified as dynamic in the
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object bound to the pipeline bind point
  used by this command, since that pipeline was bound

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08609"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08609"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08609  
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

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08610"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08610"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08610  
  If the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object bound to the pipeline bind
  point used by this command or any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html)
  bound to a stage corresponding to the pipeline bind point used by this
  command accesses a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) object that uses
  unnormalized coordinates, that sampler **must** not be used with any
  of the SPIR-V `OpImageSample*` or `OpImageSparseSample*` instructions
  with `ImplicitLod`, `Dref` or `Proj` in their name, in any shader
  stage

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08611"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08611"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08611  
  If the [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) object bound to the pipeline bind
  point used by this command or any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html)
  bound to a stage corresponding to the pipeline bind point used by this
  command accesses a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) object that uses
  unnormalized coordinates, that sampler **must** not be used with any
  of the SPIR-V `OpImageSample*` or `OpImageSparseSample*` instructions
  that includes a LOD bias or any offset values, in any shader stage

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08607"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08607"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08607  
  If the [`shaderObject`](#features-shaderObject) is enabled, either a
  valid pipeline **must** be bound to the pipeline bind point used by
  this command, or a valid combination of valid and
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) shader objects **must** be bound
  to every supported shader stage corresponding to the pipeline bind
  point used by this command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-uniformBuffers-06935"
  id="VUID-vkCmdDrawMeshTasksEXT-uniformBuffers-06935"></a>
  VUID-vkCmdDrawMeshTasksEXT-uniformBuffers-06935  
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

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08612"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08612"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08612  
  If the [`robustBufferAccess`](#features-robustBufferAccess) feature is
  not enabled, and any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) bound to a stage
  corresponding to the pipeline bind point used by this command accesses
  a uniform buffer, it **must** not access values outside of the range
  of the buffer as specified in the descriptor set bound to the same
  pipeline bind point

- <a href="#VUID-vkCmdDrawMeshTasksEXT-storageBuffers-06936"
  id="VUID-vkCmdDrawMeshTasksEXT-storageBuffers-06936"></a>
  VUID-vkCmdDrawMeshTasksEXT-storageBuffers-06936  
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

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08613"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08613"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08613  
  If the [`robustBufferAccess`](#features-robustBufferAccess) feature is
  not enabled, and any [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) bound to a stage
  corresponding to the pipeline bind point used by this command accesses
  a storage buffer, it **must** not access values outside of the range
  of the buffer as specified in the descriptor set bound to the same
  pipeline bind point

- <a href="#VUID-vkCmdDrawMeshTasksEXT-commandBuffer-02707"
  id="VUID-vkCmdDrawMeshTasksEXT-commandBuffer-02707"></a>
  VUID-vkCmdDrawMeshTasksEXT-commandBuffer-02707  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported, any
  resource accessed by [bound shaders](#shaders-binding) **must** not be
  a protected resource

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-06550"
  id="VUID-vkCmdDrawMeshTasksEXT-None-06550"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-06550  
  If [a bound shader](#shaders-binding) accesses a
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) or [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) object
  that enables [sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion](#samplers-YCbCr-conversion), that object **must** only be
  used with `OpImageSample*` or `OpImageSparseSample*` instructions

- <a href="#VUID-vkCmdDrawMeshTasksEXT-ConstOffset-06551"
  id="VUID-vkCmdDrawMeshTasksEXT-ConstOffset-06551"></a>
  VUID-vkCmdDrawMeshTasksEXT-ConstOffset-06551  
  If [a bound shader](#shaders-binding) accesses a
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) or [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) object
  that enables [sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion](#samplers-YCbCr-conversion), that object **must** not use
  the `ConstOffset` and `Offset` operands

- <a href="#VUID-vkCmdDrawMeshTasksEXT-viewType-07752"
  id="VUID-vkCmdDrawMeshTasksEXT-viewType-07752"></a>
  VUID-vkCmdDrawMeshTasksEXT-viewType-07752  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is accessed as a result of this
  command, then the image view’s `viewType` **must** match the `Dim`
  operand of the `OpTypeImage` as described in
  [\[textures-operation-validation\]](#textures-operation-validation)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-format-07753"
  id="VUID-vkCmdDrawMeshTasksEXT-format-07753"></a>
  VUID-vkCmdDrawMeshTasksEXT-format-07753  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) is accessed as a result of this
  command, then the [numeric type](#formats-numericformat) of the image
  view’s `format` and the `Sampled` `Type` operand of the `OpTypeImage`
  **must** match

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageWrite-08795"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageWrite-08795"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageWrite-08795  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) created with a format other than
  `VK_FORMAT_A8_UNORM_KHR` is accessed using `OpImageWrite` as a result
  of this command, then the `Type` of the `Texel` operand of that
  instruction **must** have at least as many components as the image
  view’s format

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageWrite-08796"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageWrite-08796"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageWrite-08796  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) created with the format
  `VK_FORMAT_A8_UNORM_KHR` is accessed using `OpImageWrite` as a result
  of this command, then the `Type` of the `Texel` operand of that
  instruction **must** have four components

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageWrite-04469"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageWrite-04469"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageWrite-04469  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) is accessed using
  `OpImageWrite` as a result of this command, then the `Type` of the
  `Texel` operand of that instruction **must** have at least as many
  components as the buffer view’s format

- <a href="#VUID-vkCmdDrawMeshTasksEXT-SampledType-04470"
  id="VUID-vkCmdDrawMeshTasksEXT-SampledType-04470"></a>
  VUID-vkCmdDrawMeshTasksEXT-SampledType-04470  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)
  that has a 64-bit component width is accessed as a result of this
  command, the `SampledType` of the `OpTypeImage` operand of that
  instruction **must** have a `Width` of 64

- <a href="#VUID-vkCmdDrawMeshTasksEXT-SampledType-04471"
  id="VUID-vkCmdDrawMeshTasksEXT-SampledType-04471"></a>
  VUID-vkCmdDrawMeshTasksEXT-SampledType-04471  
  If a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) with a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)
  that has a component width less than 64-bit is accessed as a result of
  this command, the `SampledType` of the `OpTypeImage` operand of that
  instruction **must** have a `Width` of 32

- <a href="#VUID-vkCmdDrawMeshTasksEXT-SampledType-04472"
  id="VUID-vkCmdDrawMeshTasksEXT-SampledType-04472"></a>
  VUID-vkCmdDrawMeshTasksEXT-SampledType-04472  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) with a
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) that has a 64-bit component width is
  accessed as a result of this command, the `SampledType` of the
  `OpTypeImage` operand of that instruction **must** have a `Width` of
  64

- <a href="#VUID-vkCmdDrawMeshTasksEXT-SampledType-04473"
  id="VUID-vkCmdDrawMeshTasksEXT-SampledType-04473"></a>
  VUID-vkCmdDrawMeshTasksEXT-SampledType-04473  
  If a [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) with a
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) that has a component width less than 64-bit
  is accessed as a result of this command, the `SampledType` of the
  `OpTypeImage` operand of that instruction **must** have a `Width` of
  32

- <a href="#VUID-vkCmdDrawMeshTasksEXT-sparseImageInt64Atomics-04474"
  id="VUID-vkCmdDrawMeshTasksEXT-sparseImageInt64Atomics-04474"></a>
  VUID-vkCmdDrawMeshTasksEXT-sparseImageInt64Atomics-04474  
  If the [`sparseImageInt64Atomics`](#features-sparseImageInt64Atomics)
  feature is not enabled, [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) objects created with
  the `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` flag **must** not be
  accessed by atomic instructions through an `OpTypeImage` with a
  `SampledType` with a `Width` of 64 by this command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-sparseImageInt64Atomics-04475"
  id="VUID-vkCmdDrawMeshTasksEXT-sparseImageInt64Atomics-04475"></a>
  VUID-vkCmdDrawMeshTasksEXT-sparseImageInt64Atomics-04475  
  If the [`sparseImageInt64Atomics`](#features-sparseImageInt64Atomics)
  feature is not enabled, [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) objects created with
  the `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT` flag **must** not be
  accessed by atomic instructions through an `OpTypeImage` with a
  `SampledType` with a `Width` of 64 by this command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageWeightedSampleQCOM-06971"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageWeightedSampleQCOM-06971"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageWeightedSampleQCOM-06971  
  If `OpImageWeightedSampleQCOM` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_WEIGHT_SAMPLED_IMAGE_BIT_QCOM`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageWeightedSampleQCOM-06972"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageWeightedSampleQCOM-06972"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageWeightedSampleQCOM-06972  
  If `OpImageWeightedSampleQCOM` uses a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html)
  as a sample weight image as a result of this command, then the image
  view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_WEIGHT_IMAGE_BIT_QCOM`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageBoxFilterQCOM-06973"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageBoxFilterQCOM-06973"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageBoxFilterQCOM-06973  
  If `OpImageBoxFilterQCOM` is used to sample a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_BOX_FILTER_SAMPLED_BIT_QCOM`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchSSDQCOM-06974"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchSSDQCOM-06974"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchSSDQCOM-06974  
  If `OpImageBlockMatchSSDQCOM` is used to read from an
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchSADQCOM-06975"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchSADQCOM-06975"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchSADQCOM-06975  
  If `OpImageBlockMatchSADQCOM` is used to read from an
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as a result of this command, then the
  image view’s [format features](#resources-image-view-format-features)
  **must** contain `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchSADQCOM-06976"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchSADQCOM-06976"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchSADQCOM-06976  
  If `OpImageBlockMatchSADQCOM` or OpImageBlockMatchSSDQCOM is used to
  read from a reference image as result of this command, then the
  specified reference coordinates **must** not fail [integer texel
  coordinate validation](#textures-integer-coordinate-validation)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageWeightedSampleQCOM-06977"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageWeightedSampleQCOM-06977"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageWeightedSampleQCOM-06977  
  If `OpImageWeightedSampleQCOM`, `OpImageBoxFilterQCOM`,
  `OpImageBlockMatchWindowSSDQCOM`, `OpImageBlockMatchWindowSADQCOM`,
  `OpImageBlockMatchGatherSSDQCOM`, `OpImageBlockMatchGatherSADQCOM`,
  `OpImageBlockMatchSSDQCOM`, or `OpImageBlockMatchSADQCOM` uses a
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) as a result of this command, then the
  sampler **must** have been created with
  `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageWeightedSampleQCOM-06978"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageWeightedSampleQCOM-06978"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageWeightedSampleQCOM-06978  
  If any command other than `OpImageWeightedSampleQCOM`,
  `OpImageBoxFilterQCOM`, `OpImageBlockMatchWindowSSDQCOM`,
  `OpImageBlockMatchWindowSADQCOM`, `OpImageBlockMatchGatherSSDQCOM`,
  `OpImageBlockMatchGatherSADQCOM`, `OpImageBlockMatchSSDQCOM`, or
  `OpImageBlockMatchSADQCOM` uses a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) as a
  result of this command, then the sampler **must** not have been
  created with `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchWindow-09215"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchWindow-09215"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchWindow-09215  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  instruction is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as
  a result of this command, then the image view’s [format
  features](#resources-image-view-format-features) **must** contain
  `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchWindow-09216"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchWindow-09216"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchWindow-09216  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  instruction is used to read from an [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) as
  a result of this command, then the image view’s format **must** be a
  single-component format

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchWindow-09217"
  id="VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchWindow-09217"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpImageBlockMatchWindow-09217  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  read from a reference image as result of this command, then the
  specified reference coordinates **must** not fail [integer texel
  coordinate validation](#textures-integer-coordinate-validation)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07288"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07288"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07288  
  Any shader invocation executed by this command **must**
  [terminate](#shaders-termination)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09600"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09600"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09600  
  If a descriptor with type equal to any of
  `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM`,
  `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`,
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` is accessed as a result of this
  command, the image subresource identified by that descriptor **must**
  be in the image layout identified when the descriptor was written

- <a href="#VUID-vkCmdDrawMeshTasksEXT-renderPass-02684"
  id="VUID-vkCmdDrawMeshTasksEXT-renderPass-02684"></a>
  VUID-vkCmdDrawMeshTasksEXT-renderPass-02684  
  The current render pass **must** be
  [compatible](#renderpass-compatibility) with the `renderPass` member
  of the `VkGraphicsPipelineCreateInfo` structure specified when
  creating the `VkPipeline` bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-subpass-02685"
  id="VUID-vkCmdDrawMeshTasksEXT-subpass-02685"></a>
  VUID-vkCmdDrawMeshTasksEXT-subpass-02685  
  The subpass index of the current render pass **must** be equal to the
  `subpass` member of the `VkGraphicsPipelineCreateInfo` structure
  specified when creating the `VkPipeline` bound to
  `VK_PIPELINE_BIND_POINT_GRAPHICS`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07748"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07748"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07748  
  If any shader statically accesses an input attachment, a valid
  descriptor **must** be bound to the pipeline via a descriptor set

- <a href="#VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07468"
  id="VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07468"></a>
  VUID-vkCmdDrawMeshTasksEXT-OpTypeImage-07468  
  If any shader executed by this pipeline accesses an `OpTypeImage`
  variable with a `Dim` operand of `SubpassData`, it **must** be
  decorated with an `InputAttachmentIndex` that corresponds to a valid
  input attachment in the current subpass

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07469"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07469"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07469  
  Input attachment views accessed in a subpass **must** be created with
  the same [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) as the corresponding subpass
  definition, and be created with a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) that
  is compatible with the attachment referenced by the subpass'
  `pInputAttachments`\[`InputAttachmentIndex`\] in the currently bound
  [VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html) as specified by [Fragment Input
  Attachment Compatibility](#compatibility-inputattachment)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pDepthInputAttachmentIndex-09595"
  id="VUID-vkCmdDrawMeshTasksEXT-pDepthInputAttachmentIndex-09595"></a>
  VUID-vkCmdDrawMeshTasksEXT-pDepthInputAttachmentIndex-09595  
  Input attachment views accessed in a dynamic render pass with a
  `InputAttachmentIndex` referenced by
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html),
  or no `InputAttachmentIndex` if
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html):`pDepthInputAttachmentIndex`
  or
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html):`pStencilInputAttachmentIndex`
  are `NULL`, **must** be created with a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html)
  that is compatible with the corresponding color, depth, or stencil
  attachment in [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pDepthInputAttachmentIndex-09596"
  id="VUID-vkCmdDrawMeshTasksEXT-pDepthInputAttachmentIndex-09596"></a>
  VUID-vkCmdDrawMeshTasksEXT-pDepthInputAttachmentIndex-09596  
  Input attachment views accessed in a dynamic render pass via a shader
  object **must** have an `InputAttachmentIndex` if both
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html):`pDepthInputAttachmentIndex`
  and
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html):`pStencilInputAttachmentIndex`
  are non-`NULL`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-InputAttachmentIndex-09597"
  id="VUID-vkCmdDrawMeshTasksEXT-InputAttachmentIndex-09597"></a>
  VUID-vkCmdDrawMeshTasksEXT-InputAttachmentIndex-09597  
  If an input attachment view accessed in a dynamic render pass via a
  shader object has an `InputAttachmentIndex`, the
  `InputAttachmentIndex` **must** match an index in
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-06537"
  id="VUID-vkCmdDrawMeshTasksEXT-None-06537"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-06537  
  Memory backing image subresources used as attachments in the current
  render pass **must** not be written in any way other than as an
  attachment by this command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09000"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09000"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09000  
  If a color attachment is written by any prior command in this subpass
  or by the load, store, or resolve operations for this subpass, it is
  not in the `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`
  image layout, and either:

  - the `VK_PIPELINE_CREATE_COLOR_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT` is
    set on the currently bound pipeline or

  - the last call to
    [vkCmdSetAttachmentFeedbackLoopEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAttachmentFeedbackLoopEnableEXT.html)
    included `VK_IMAGE_ASPECT_COLOR_BIT` and

    - there is no currently bound graphics pipeline or

    - the currently bound graphics pipeline was created with
      `VK_DYNAMIC_STATE_ATTACHMENT_FEEDBACK_LOOP_ENABLE_EXT` it **must**
      not be accessed in any way other than as an attachment by this
      command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09001"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09001"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09001  
  If a depth attachment is written by any prior command in this subpass
  or by the load, store, or resolve operations for this subpass, it is
  not in the `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`
  image layout, and either:

  - the
    `VK_PIPELINE_CREATE_DEPTH_STENCIL_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`
    is set on the currently bound pipeline or

  - the last call to
    [vkCmdSetAttachmentFeedbackLoopEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAttachmentFeedbackLoopEnableEXT.html)
    included `VK_IMAGE_ASPECT_DEPTH_BIT` and

    - there is no currently bound graphics pipeline or

    - the currently bound graphics pipeline was created with
      `VK_DYNAMIC_STATE_ATTACHMENT_FEEDBACK_LOOP_ENABLE_EXT` it **must**
      not be accessed in any way other than as an attachment by this
      command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09002"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09002"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09002  
  If a stencil attachment is written by any prior command in this
  subpass or by the load, store, or resolve operations for this subpass,
  it is not in the
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` image layout,
  and either:

  - the
    `VK_PIPELINE_CREATE_DEPTH_STENCIL_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`
    is set on the currently bound pipeline or

  - the last call to
    [vkCmdSetAttachmentFeedbackLoopEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAttachmentFeedbackLoopEnableEXT.html)
    included `VK_IMAGE_ASPECT_STENCIL_BIT` and

    - there is no currently bound graphics pipeline or

    - the currently bound graphics pipeline was created with
      `VK_DYNAMIC_STATE_ATTACHMENT_FEEDBACK_LOOP_ENABLE_EXT` it **must**
      not be accessed in any way other than as an attachment by this
      command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09003"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09003"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09003  
  If an attachment is written by any prior command in this subpass or by
  the load, store, or resolve operations for this subpass, it **must**
  not be accessed in any way other than as an attachment, storage image,
  or sampled image by this command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-06539"
  id="VUID-vkCmdDrawMeshTasksEXT-None-06539"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-06539  
  If any previously recorded command in the current subpass accessed an
  image subresource used as an attachment in this subpass in any way
  other than as an attachment, this command **must** not write to that
  image subresource as an attachment

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-06886"
  id="VUID-vkCmdDrawMeshTasksEXT-None-06886"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-06886  
  If the current render pass instance uses a depth/stencil attachment
  with a read-only layout for the depth aspect, [depth
  writes](#fragops-depth-write) **must** be disabled

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-06887"
  id="VUID-vkCmdDrawMeshTasksEXT-None-06887"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-06887  
  If the current render pass instance uses a depth/stencil attachment
  with a read-only layout for the stencil aspect, both front and back
  `writeMask` are not zero, and stencil test is enabled, [all stencil
  ops](#fragops-stencil) **must** be `VK_STENCIL_OP_KEEP`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07831"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07831"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07831  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VIEWPORT` dynamic state enabled then
  [vkCmdSetViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewport.html) **must** have been called in
  the current command buffer prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07832"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07832"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07832  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SCISSOR` dynamic state enabled then
  [vkCmdSetScissor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetScissor.html) **must** have been called in
  the current command buffer prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07833"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07833"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07833  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_LINE_WIDTH` dynamic state enabled then
  [vkCmdSetLineWidth](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineWidth.html) **must** have been called
  in the current command buffer prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08617"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08617"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08617  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetPolygonModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPolygonModeEXT.html) in the current
  command buffer set `polygonMode` to `VK_POLYGON_MODE_LINE`,
  [vkCmdSetLineWidth](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineWidth.html) **must** have been called
  in the current command buffer prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08618"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08618"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08618  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetPrimitiveTopology](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPrimitiveTopology.html) in the
  current command buffer set `primitiveTopology` to any line topology,
  [vkCmdSetLineWidth](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineWidth.html) **must** have been called
  in the current command buffer prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08619"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08619"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08619  
  If a shader object that outputs line primitives is bound to the
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT` or
  `VK_SHADER_STAGE_GEOMETRY_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, [vkCmdSetLineWidth](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineWidth.html) **must** have
  been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07834"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07834"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07834  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_DEPTH_BIAS` dynamic state enabled then
  [vkCmdSetDepthBias](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBias.html) or
  [vkCmdSetDepthBias2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBias2EXT.html) **must** have been
  called in the current command buffer prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08620"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08620"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08620  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetDepthBiasEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBiasEnable.html) in the current
  command buffer set `depthBiasEnable` to `VK_TRUE`,
  [vkCmdSetDepthBias](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBias.html) or
  [vkCmdSetDepthBias2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBias2EXT.html) **must** have been
  called in the current command buffer prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07835"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07835"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07835  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_BLEND_CONSTANTS` dynamic state enabled then
  [vkCmdSetBlendConstants](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetBlendConstants.html) **must** have
  been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08621"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08621"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08621  
  If a shader object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT`
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEnableEXT.html) in the
  current command buffer set any element of `pColorBlendEnables` to
  `VK_TRUE`, and the most recent call to
  [vkCmdSetColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEquationEXT.html) in
  the current command buffer set the same element of
  `pColorBlendEquations` to a `VkColorBlendEquationEXT` structure with
  any [VkBlendFactor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html) member with a value of
  `VK_BLEND_FACTOR_CONSTANT_COLOR`,
  `VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR`,
  `VK_BLEND_FACTOR_CONSTANT_ALPHA`, or
  `VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA`,
  [vkCmdSetBlendConstants](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetBlendConstants.html) **must** have
  been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07836"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07836"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07836  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_DEPTH_BOUNDS` dynamic state enabled, and if the
  current `depthBoundsTestEnable` state is `VK_TRUE`, then
  [vkCmdSetDepthBounds](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBounds.html) **must** have been
  called in the current command buffer prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08622"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08622"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08622  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetDepthBoundsTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBoundsTestEnable.html) in
  the current command buffer set `depthBoundsTestEnable` to `VK_TRUE`,
  then [vkCmdSetDepthBounds](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBounds.html) **must** have
  been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07837"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07837"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07837  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_STENCIL_COMPARE_MASK` dynamic state enabled, and if
  the current `stencilTestEnable` state is `VK_TRUE`, then
  [vkCmdSetStencilCompareMask](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilCompareMask.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08623"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08623"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08623  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetStencilTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilTestEnable.html) in the
  current command buffer set `stencilTestEnable` to `VK_TRUE`,
  [vkCmdSetStencilCompareMask](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilCompareMask.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07838"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07838"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07838  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_STENCIL_WRITE_MASK` dynamic state enabled, and if
  the current `stencilTestEnable` state is `VK_TRUE`, then
  [vkCmdSetStencilWriteMask](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilWriteMask.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08624"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08624"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08624  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetStencilTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilTestEnable.html) in the
  current command buffer set `stencilTestEnable` to `VK_TRUE`,
  [vkCmdSetStencilWriteMask](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilWriteMask.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07839"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07839"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07839  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_STENCIL_REFERENCE` dynamic state enabled, and if the
  current `stencilTestEnable` state is `VK_TRUE`, then
  [vkCmdSetStencilReference](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilReference.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08625"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08625"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08625  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetStencilTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilTestEnable.html) in the
  current command buffer set `stencilTestEnable` to `VK_TRUE`,
  [vkCmdSetStencilReference](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilReference.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-maxMultiviewInstanceIndex-02688"
  id="VUID-vkCmdDrawMeshTasksEXT-maxMultiviewInstanceIndex-02688"></a>
  VUID-vkCmdDrawMeshTasksEXT-maxMultiviewInstanceIndex-02688  
  If the draw is recorded in a render pass instance with multiview
  enabled, the maximum instance index **must** be less than or equal to
  [VkPhysicalDeviceMultiviewProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewProperties.html)::`maxMultiviewInstanceIndex`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-02689"
  id="VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-02689"></a>
  VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-02689  
  If the bound graphics pipeline was created with
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)::`sampleLocationsEnable`
  set to `VK_TRUE` and the current subpass has a depth/stencil
  attachment, then that attachment **must** have been created with the
  `VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT` bit set

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-06666"
  id="VUID-vkCmdDrawMeshTasksEXT-None-06666"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-06666  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT` dynamic state enabled then
  [vkCmdSetSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEXT.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08626"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08626"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08626  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetSampleLocationsEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEnableEXT.html)
  in the current command buffer set `sampleLocationsEnable` to
  `VK_TRUE`, then
  [vkCmdSetSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEXT.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07840"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07840"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07840  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_CULL_MODE` dynamic state enabled then
  [vkCmdSetCullMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCullMode.html) **must** have been called in
  the current command buffer prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08627"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08627"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08627  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then [vkCmdSetCullMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCullMode.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07841"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07841"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07841  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_FRONT_FACE` dynamic state enabled then
  [vkCmdSetFrontFace](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFrontFace.html) **must** have been called
  in the current command buffer prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08628"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08628"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08628  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then [vkCmdSetFrontFace](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFrontFace.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07843"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07843"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07843  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE` dynamic state enabled then
  [vkCmdSetDepthTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthTestEnable.html) **must** have
  been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08629"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08629"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08629  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetDepthTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthTestEnable.html) **must** have
  been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07844"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07844"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07844  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE` dynamic state enabled then
  [vkCmdSetDepthWriteEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthWriteEnable.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08630"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08630"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08630  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetDepthWriteEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthWriteEnable.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07845"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07845"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07845  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_DEPTH_COMPARE_OP` dynamic state enabled then
  [vkCmdSetDepthCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthCompareOp.html) **must** have
  been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08631"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08631"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08631  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetDepthTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthTestEnable.html) in the current
  command buffer set `depthTestEnable` to `VK_TRUE`, then
  [vkCmdSetDepthCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthCompareOp.html) **must** have
  been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07846"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07846"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07846  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE` dynamic state enabled then
  [vkCmdSetDepthBoundsTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBoundsTestEnable.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08632"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08632"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08632  
  If a shader object is bound to any graphics stage, and the
  [`depthBounds`](#features-depthBounds) feature is enabled, and the
  most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then the
  [vkCmdSetDepthBoundsTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBoundsTestEnable.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07847"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07847"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07847  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE` dynamic state enabled then
  [vkCmdSetStencilTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilTestEnable.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08633"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08633"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08633  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetStencilTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilTestEnable.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07848"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07848"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07848  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_STENCIL_OP` dynamic state enabled then
  [vkCmdSetStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilOp.html) **must** have been called
  in the current command buffer prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08634"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08634"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08634  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetStencilTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilTestEnable.html) in the
  current command buffer set `stencilTestEnable` to `VK_TRUE`, then
  [vkCmdSetStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilOp.html) **must** have been called
  in the current command buffer prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-viewportCount-03417"
  id="VUID-vkCmdDrawMeshTasksEXT-viewportCount-03417"></a>
  VUID-vkCmdDrawMeshTasksEXT-viewportCount-03417  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` dynamic state enabled, but not
  the `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT` dynamic state enabled, then
  [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html) **must**
  have been called in the current command buffer prior to this drawing
  command, and the `viewportCount` parameter of
  `vkCmdSetViewportWithCount` **must** match the
  `VkPipelineViewportStateCreateInfo`::`scissorCount` of the pipeline

- <a href="#VUID-vkCmdDrawMeshTasksEXT-scissorCount-03418"
  id="VUID-vkCmdDrawMeshTasksEXT-scissorCount-03418"></a>
  VUID-vkCmdDrawMeshTasksEXT-scissorCount-03418  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT` dynamic state enabled, but not
  the `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` dynamic state enabled, then
  [vkCmdSetScissorWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetScissorWithCount.html) **must**
  have been called in the current command buffer prior to this drawing
  command, and the `scissorCount` parameter of
  `vkCmdSetScissorWithCount` **must** match the
  `VkPipelineViewportStateCreateInfo`::`viewportCount` of the pipeline

- <a href="#VUID-vkCmdDrawMeshTasksEXT-viewportCount-03419"
  id="VUID-vkCmdDrawMeshTasksEXT-viewportCount-03419"></a>
  VUID-vkCmdDrawMeshTasksEXT-viewportCount-03419  
  If the bound graphics pipeline state was created with both the
  `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT` and
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` dynamic states enabled then
  both [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html) and
  [vkCmdSetScissorWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetScissorWithCount.html) **must**
  have been called in the current command buffer prior to this drawing
  command, and the `viewportCount` parameter of
  `vkCmdSetViewportWithCount` **must** match the `scissorCount`
  parameter of `vkCmdSetScissorWithCount`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08635"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08635"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08635  
  If a shader object is bound to any graphics stage, then both
  [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html) and
  [vkCmdSetScissorWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetScissorWithCount.html) **must**
  have been called in the current command buffer prior to this drawing
  command, and the `viewportCount` parameter of
  `vkCmdSetViewportWithCount` **must** match the `scissorCount`
  parameter of `vkCmdSetScissorWithCount`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-viewportCount-04137"
  id="VUID-vkCmdDrawMeshTasksEXT-viewportCount-04137"></a>
  VUID-vkCmdDrawMeshTasksEXT-viewportCount-04137  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` dynamic state enabled, but not
  the `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_NV` dynamic state enabled,
  then the bound graphics pipeline **must** have been created with
  [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)::`viewportCount`
  greater or equal to the `viewportCount` parameter in the last call to
  [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-viewportCount-04138"
  id="VUID-vkCmdDrawMeshTasksEXT-viewportCount-04138"></a>
  VUID-vkCmdDrawMeshTasksEXT-viewportCount-04138  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` and
  `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_NV` dynamic states enabled then
  the `viewportCount` parameter in the last call to
  [vkCmdSetViewportWScalingNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWScalingNV.html) **must**
  be greater than or equal to the `viewportCount` parameter in the last
  call to [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09232"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09232"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09232  
  If the [`VK_NV_clip_space_w_scaling`](VK_NV_clip_space_w_scaling.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetViewportWScalingEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWScalingEnableNV.html)
  in the current command buffer set `viewportWScalingEnable` to
  `VK_TRUE`, then
  [vkCmdSetViewportWScalingNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWScalingNV.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08636"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08636"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08636  
  If the [`VK_NV_clip_space_w_scaling`](VK_NV_clip_space_w_scaling.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetViewportWScalingEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWScalingEnableNV.html)
  in the current command buffer set `viewportWScalingEnable` to
  `VK_TRUE`, then the `viewportCount` parameter in the last call to
  [vkCmdSetViewportWScalingNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWScalingNV.html) **must**
  be greater than or equal to the `viewportCount` parameter in the last
  call to [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-viewportCount-04139"
  id="VUID-vkCmdDrawMeshTasksEXT-viewportCount-04139"></a>
  VUID-vkCmdDrawMeshTasksEXT-viewportCount-04139  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` dynamic state enabled, but not
  the `VK_DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV` dynamic state
  enabled, then the bound graphics pipeline **must** have been created
  with
  [VkPipelineViewportShadingRateImageStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportShadingRateImageStateCreateInfoNV.html)::`viewportCount`
  greater or equal to the `viewportCount` parameter in the last call to
  [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-viewportCount-04140"
  id="VUID-vkCmdDrawMeshTasksEXT-viewportCount-04140"></a>
  VUID-vkCmdDrawMeshTasksEXT-viewportCount-04140  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` and
  `VK_DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV` dynamic states
  enabled then the `viewportCount` parameter in the last call to
  [vkCmdSetViewportShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportShadingRatePaletteNV.html)
  **must** be greater than or equal to the `viewportCount` parameter in
  the last call to
  [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-shadingRateImage-09233"
  id="VUID-vkCmdDrawMeshTasksEXT-shadingRateImage-09233"></a>
  VUID-vkCmdDrawMeshTasksEXT-shadingRateImage-09233  
  If the [`shadingRateImage`](#features-shadingRateImage) feature is
  enabled, and a shader object is bound to any graphics stage, and the
  most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetCoarseSampleOrderNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoarseSampleOrderNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-shadingRateImage-09234"
  id="VUID-vkCmdDrawMeshTasksEXT-shadingRateImage-09234"></a>
  VUID-vkCmdDrawMeshTasksEXT-shadingRateImage-09234  
  If the [`shadingRateImage`](#features-shadingRateImage) feature is
  enabled, and a shader object is bound to any graphics stage, and the
  most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetShadingRateImageEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetShadingRateImageEnableNV.html)
  in the current command buffer set `shadingRateImageEnable` to
  `VK_TRUE`, then
  [vkCmdSetViewportShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportShadingRatePaletteNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08637"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08637"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08637  
  If the [`shadingRateImage`](#features-shadingRateImage) feature is
  enabled, and a shader object is bound to any graphics stage, and the
  most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetShadingRateImageEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetShadingRateImageEnableNV.html)
  in the current command buffer set `shadingRateImageEnable` to
  `VK_TRUE`, then the `viewportCount` parameter in the last call to
  [vkCmdSetViewportShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportShadingRatePaletteNV.html)
  **must** be greater than or equal to the `viewportCount` parameter in
  the last call to
  [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-VkPipelineVieportCreateInfo-04141"
  id="VUID-vkCmdDrawMeshTasksEXT-VkPipelineVieportCreateInfo-04141"></a>
  VUID-vkCmdDrawMeshTasksEXT-VkPipelineVieportCreateInfo-04141  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` dynamic state enabled and a
  [VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html)
  structure chained from
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html),
  then the bound graphics pipeline **must** have been created with
  [VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html)::`viewportCount`
  greater or equal to the `viewportCount` parameter in the last call to
  [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-VkPipelineVieportCreateInfo-04142"
  id="VUID-vkCmdDrawMeshTasksEXT-VkPipelineVieportCreateInfo-04142"></a>
  VUID-vkCmdDrawMeshTasksEXT-VkPipelineVieportCreateInfo-04142  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` dynamic state enabled and a
  [VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html)
  structure chained from
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html),
  then the bound graphics pipeline **must** have been created with
  [VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html)::`exclusiveScissorCount`
  greater or equal to the `viewportCount` parameter in the last call to
  [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07878"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07878"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07878  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_ENABLE_NV` dynamic state enabled
  then
  [vkCmdSetExclusiveScissorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExclusiveScissorEnableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07879"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07879"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07879  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV` dynamic state enabled then
  [vkCmdSetExclusiveScissorNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExclusiveScissorNV.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-exclusiveScissor-09235"
  id="VUID-vkCmdDrawMeshTasksEXT-exclusiveScissor-09235"></a>
  VUID-vkCmdDrawMeshTasksEXT-exclusiveScissor-09235  
  If the [`exclusiveScissor`](#features-exclusiveScissor) feature is
  enabled, and a shader object is bound to any graphics stage, then
  [vkCmdSetExclusiveScissorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExclusiveScissorEnableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08638"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08638"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08638  
  If the [`exclusiveScissor`](#features-exclusiveScissor) feature is
  enabled, and a shader object is bound to any graphics stage, and the
  most recent call to
  [vkCmdSetExclusiveScissorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExclusiveScissorEnableNV.html)
  in the current command buffer set any element of
  `pExclusiveScissorEnables` to `VK_TRUE`, then
  [vkCmdSetExclusiveScissorNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExclusiveScissorNV.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-04876"
  id="VUID-vkCmdDrawMeshTasksEXT-None-04876"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-04876  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE` dynamic state enabled
  then
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08639"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08639"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08639  
  If a shader object is bound to any graphics stage, then
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-04877"
  id="VUID-vkCmdDrawMeshTasksEXT-None-04877"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-04877  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE` dynamic state enabled then
  [vkCmdSetDepthBiasEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBiasEnable.html) **must** have
  been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08640"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08640"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08640  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetDepthBiasEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBiasEnable.html) **must** have
  been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-logicOp-04878"
  id="VUID-vkCmdDrawMeshTasksEXT-logicOp-04878"></a>
  VUID-vkCmdDrawMeshTasksEXT-logicOp-04878  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_LOGIC_OP_EXT` dynamic state enabled then
  [vkCmdSetLogicOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEXT.html) **must** have been
  called in the current command buffer prior to this drawing command and
  the `logicOp` **must** be a valid [VkLogicOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html) value

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08641"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08641"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08641  
  If a shader object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT`
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetLogicOpEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEnableEXT.html) set
  `logicOpEnable` to `VK_TRUE`, then
  [vkCmdSetLogicOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEXT.html) **must** have been
  called in the current command buffer prior to this drawing command and
  the `logicOp` **must** be a valid [VkLogicOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html) value

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-primitiveFragmentShadingRateWithMultipleViewports-04552"
  id="VUID-vkCmdDrawMeshTasksEXT-primitiveFragmentShadingRateWithMultipleViewports-04552"></a>
  VUID-vkCmdDrawMeshTasksEXT-primitiveFragmentShadingRateWithMultipleViewports-04552  
  If the
  [`primitiveFragmentShadingRateWithMultipleViewports`](#limits-primitiveFragmentShadingRateWithMultipleViewports)
  limit is not supported, the bound graphics pipeline was created with
  the `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` dynamic state enabled, and
  any of the shader stages of the bound graphics pipeline write to the
  `PrimitiveShadingRateKHR` built-in, then
  [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html) **must**
  have been called in the current command buffer prior to this drawing
  command, and the `viewportCount` parameter of
  `vkCmdSetViewportWithCount` **must** be `1`

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-primitiveFragmentShadingRateWithMultipleViewports-08642"
  id="VUID-vkCmdDrawMeshTasksEXT-primitiveFragmentShadingRateWithMultipleViewports-08642"></a>
  VUID-vkCmdDrawMeshTasksEXT-primitiveFragmentShadingRateWithMultipleViewports-08642  
  If the
  [`primitiveFragmentShadingRateWithMultipleViewports`](#limits-primitiveFragmentShadingRateWithMultipleViewports)
  limit is not supported, and any shader object bound to a graphics
  stage writes to the `PrimitiveShadingRateKHR` built-in, then
  [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html) **must**
  have been called in the current command buffer prior to this drawing
  command, and the `viewportCount` parameter of
  `vkCmdSetViewportWithCount` **must** be `1`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-blendEnable-04727"
  id="VUID-vkCmdDrawMeshTasksEXT-blendEnable-04727"></a>
  VUID-vkCmdDrawMeshTasksEXT-blendEnable-04727  
  If rasterization is not disabled in the bound graphics pipeline, then
  for each color attachment in the subpass, if the corresponding image
  view’s [format features](#resources-image-view-format-features) do not
  contain `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT`, then the
  `blendEnable` member of the corresponding element of the
  `pAttachments` member of `pColorBlendState` **must** be `VK_FALSE`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08643"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08643"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08643  
  If a shader object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT`
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then for each color attachment in the render pass, if the
  corresponding image view’s [format
  features](#resources-image-view-format-features) do not contain
  `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT`, then the corresponding
  member of `pColorBlendEnables` in the most recent call to
  `vkCmdSetColorBlendEnableEXT` in the current command buffer that
  affected that attachment index **must** have been `VK_FALSE`

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-multisampledRenderToSingleSampled-07284"
  id="VUID-vkCmdDrawMeshTasksEXT-multisampledRenderToSingleSampled-07284"></a>
  VUID-vkCmdDrawMeshTasksEXT-multisampledRenderToSingleSampled-07284  
  If rasterization is not disabled in the bound graphics pipeline, and
  none of the following is enabled:

  - the
    [`VK_AMD_mixed_attachment_samples`](VK_AMD_mixed_attachment_samples.html)
    extension

  - the
    [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
    extension

  - the
    [`multisampledRenderToSingleSampled`](#features-multisampledRenderToSingleSampled)
    feature

  then `rasterizationSamples` for the currently bound graphics pipeline
  **must** be the same as the current subpass color and/or depth/stencil
  attachments

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08644"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08644"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08644  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and none of the following is enabled:

  - the
    [`VK_AMD_mixed_attachment_samples`](VK_AMD_mixed_attachment_samples.html)
    extension

  - the
    [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
    extension

  - the
    [`multisampledRenderToSingleSampled`](#features-multisampledRenderToSingleSampled)
    feature

  then the most recent call to
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)
  in the current command buffer **must** have set `rasterizationSamples`
  to be the same as the number of samples for the current render pass
  color and/or depth/stencil attachments

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08876"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08876"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08876  
  If a shader object is bound to any graphics stage, the current render
  pass instance **must** have been begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-imageView-06172"
  id="VUID-vkCmdDrawMeshTasksEXT-imageView-06172"></a>
  VUID-vkCmdDrawMeshTasksEXT-imageView-06172  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the `imageView`
  member of `pDepthAttachment` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the `layout` member of
  `pDepthAttachment` is
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`, this command
  **must** not write any values to the depth attachment

- <a href="#VUID-vkCmdDrawMeshTasksEXT-imageView-06173"
  id="VUID-vkCmdDrawMeshTasksEXT-imageView-06173"></a>
  VUID-vkCmdDrawMeshTasksEXT-imageView-06173  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the `imageView`
  member of `pStencilAttachment` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the `layout` member of
  `pStencilAttachment` is
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`, this command
  **must** not write any values to the stencil attachment

- <a href="#VUID-vkCmdDrawMeshTasksEXT-imageView-06174"
  id="VUID-vkCmdDrawMeshTasksEXT-imageView-06174"></a>
  VUID-vkCmdDrawMeshTasksEXT-imageView-06174  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the `imageView`
  member of `pDepthAttachment` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the `layout` member of
  `pDepthAttachment` is
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`, this
  command **must** not write any values to the depth attachment

- <a href="#VUID-vkCmdDrawMeshTasksEXT-imageView-06175"
  id="VUID-vkCmdDrawMeshTasksEXT-imageView-06175"></a>
  VUID-vkCmdDrawMeshTasksEXT-imageView-06175  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the `imageView`
  member of `pStencilAttachment` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the `layout` member of
  `pStencilAttachment` is
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`, this
  command **must** not write any values to the stencil attachment

- <a href="#VUID-vkCmdDrawMeshTasksEXT-imageView-06176"
  id="VUID-vkCmdDrawMeshTasksEXT-imageView-06176"></a>
  VUID-vkCmdDrawMeshTasksEXT-imageView-06176  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the `imageView`
  member of `pDepthAttachment` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the `layout` member of
  `pDepthAttachment` is `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, this
  command **must** not write any values to the depth attachment

- <a href="#VUID-vkCmdDrawMeshTasksEXT-imageView-06177"
  id="VUID-vkCmdDrawMeshTasksEXT-imageView-06177"></a>
  VUID-vkCmdDrawMeshTasksEXT-imageView-06177  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the `imageView`
  member of `pStencilAttachment` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the `layout` member of
  `pStencilAttachment` is `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,
  this command **must** not write any values to the stencil attachment

- <a href="#VUID-vkCmdDrawMeshTasksEXT-viewMask-06178"
  id="VUID-vkCmdDrawMeshTasksEXT-viewMask-06178"></a>
  VUID-vkCmdDrawMeshTasksEXT-viewMask-06178  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the currently bound
  graphics pipeline **must** have been created with a
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`viewMask`
  equal to [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`viewMask`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-colorAttachmentCount-06179"
  id="VUID-vkCmdDrawMeshTasksEXT-colorAttachmentCount-06179"></a>
  VUID-vkCmdDrawMeshTasksEXT-colorAttachmentCount-06179  
  If the
  [`dynamicRenderingUnusedAttachments`](#features-dynamicRenderingUnusedAttachments)
  feature is not enabled and the current render pass instance was begun
  with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the currently
  bound graphics pipeline **must** have been created with a
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`colorAttachmentCount`
  equal to
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`colorAttachmentCount`

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08910"
  id="VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08910"></a>
  VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08910  
  If the
  [`dynamicRenderingUnusedAttachments`](#features-dynamicRenderingUnusedAttachments)
  feature is not enabled, and the current render pass instance was begun
  with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) and
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`colorAttachmentCount`
  greater than `0`, then each element of the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments` array
  with an `imageView` not equal to [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  **must** have been created with a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) equal to
  the corresponding element of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`pColorAttachmentFormats`
  used to create the currently bound graphics pipeline

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08912"
  id="VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08912"></a>
  VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08912  
  If the
  [`dynamicRenderingUnusedAttachments`](#features-dynamicRenderingUnusedAttachments)
  feature is not enabled, and the current render pass instance was begun
  with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) and
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`colorAttachmentCount`
  greater than `0`, then each element of the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments` array
  with an `imageView` equal to [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  **must** have the corresponding element of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`pColorAttachmentFormats`
  used to create the currently bound pipeline equal to
  `VK_FORMAT_UNDEFINED`

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08911"
  id="VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08911"></a>
  VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08911  
  If the
  [`dynamicRenderingUnusedAttachments`](#features-dynamicRenderingUnusedAttachments)
  feature is enabled, and the current render pass instance was begun
  with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) and
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`colorAttachmentCount`
  greater than `0`, then each element of the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments` array
  with an `imageView` not equal to [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  **must** have been created with a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) equal to
  the corresponding element of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`pColorAttachmentFormats`
  used to create the currently bound graphics pipeline, or the
  corresponding element of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`pColorAttachmentFormats`,
  if it exists, **must** be `VK_FORMAT_UNDEFINED`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-colorAttachmentCount-09362"
  id="VUID-vkCmdDrawMeshTasksEXT-colorAttachmentCount-09362"></a>
  VUID-vkCmdDrawMeshTasksEXT-colorAttachmentCount-09362  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), with a
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`colorAttachmentCount` equal
  to `1`, there is no shader object bound to any graphics stage, and a
  color attachment with a resolve mode of
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`, each element of
  the [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments` array
  with a `resolveImageView` not equal to
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) **must** have been created with
  an image created with a
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  value equal to the
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  value used to create the currently bound graphics pipeline

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09363"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09363"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09363  
  If there is no shader object bound to any graphics stage, the current
  render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) and a
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`colorAttachmentCount` equal
  to `1`, and a color attachment with a resolve mode of
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`, each element of
  the [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments` array
  with a `imageView` not equal to [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  **must** have been created with an image created with a
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  value equal to the
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  value used to create the currently bound graphics pipeline

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09364"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09364"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09364  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), there is no shader
  object bound to any graphics stage, and the currently bound graphics
  pipeline was created with a non-zero
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  value and with the `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT` dynamic
  state enabled, then
  [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEnableEXT.html)
  **must** have set the blend enable to `VK_FALSE` prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09365"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09365"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09365  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), there is no shader
  object bound to any graphics stage, and the currently bound graphics
  pipeline was created with a non-zero
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  value and with the `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT`
  dynamic state enabled, then
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)
  **must** have set `rasterizationSamples` to `VK_SAMPLE_COUNT_1_BIT`
  prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09366"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09366"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09366  
  If there is a shader object bound to any graphics stage, and the
  current render pass includes a color attachment that uses the
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID` resolve mode,
  then [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEnableEXT.html)
  **must** have set blend enable to `VK_FALSE` prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-09367"
  id="VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-09367"></a>
  VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-09367  
  If there is a shader object bound to any graphics stage, and the
  current render pass includes a color attachment that uses the
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID` resolve mode,
  then
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)
  **must** have set `rasterizationSamples` to `VK_SAMPLE_COUNT_1_BIT`
  prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09368"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09368"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09368  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), there is no shader
  object bound to any graphics stage, and the currently bound graphics
  pipeline was created with a non-zero
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  value and with the `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR`
  dynamic state enabled, then
  [vkCmdSetFragmentShadingRateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFragmentShadingRateKHR.html)
  **must** have set `pFragmentSize->width` to `1` prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09369"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09369"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09369  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), there is no shader
  object bound to any graphics stage, and the currently bound graphics
  pipeline was created with a non-zero
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  value and with the `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR`
  dynamic state enabled, then
  [vkCmdSetFragmentShadingRateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFragmentShadingRateKHR.html)
  **must** have set `pFragmentSize->height` to `1` prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pFragmentSize-09370"
  id="VUID-vkCmdDrawMeshTasksEXT-pFragmentSize-09370"></a>
  VUID-vkCmdDrawMeshTasksEXT-pFragmentSize-09370  
  If there is a shader object bound to any graphics stage, and the
  current render pass includes a color attachment that uses the
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID` resolve mode,
  then
  [vkCmdSetFragmentShadingRateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFragmentShadingRateKHR.html)
  **must** have set `pFragmentSize->width` to `1` prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pFragmentSize-09371"
  id="VUID-vkCmdDrawMeshTasksEXT-pFragmentSize-09371"></a>
  VUID-vkCmdDrawMeshTasksEXT-pFragmentSize-09371  
  If there is a shader object bound to any graphics stage, and the
  current render pass includes a color attachment that uses the
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID` resolve mode,
  then
  [vkCmdSetFragmentShadingRateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFragmentShadingRateKHR.html)
  **must** have set `pFragmentSize->height` to `1` prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07749"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07749"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07749  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COLOR_WRITE_ENABLE_EXT` dynamic state enabled then
  [vkCmdSetColorWriteEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorWriteEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08646"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08646"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08646  
  If the [`colorWriteEnable`](#features-colorWriteEnable) feature is
  enabled on the device, and a shader object is bound to the
  `VK_SHADER_STAGE_FRAGMENT_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetColorWriteEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorWriteEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-attachmentCount-07750"
  id="VUID-vkCmdDrawMeshTasksEXT-attachmentCount-07750"></a>
  VUID-vkCmdDrawMeshTasksEXT-attachmentCount-07750  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COLOR_WRITE_ENABLE_EXT` dynamic state enabled then
  the `attachmentCount` parameter of `vkCmdSetColorWriteEnableEXT`
  **must** be greater than or equal to the
  `VkPipelineColorBlendStateCreateInfo`::`attachmentCount` of the
  currently bound graphics pipeline

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08647"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08647"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08647  
  If the [`colorWriteEnable`](#features-colorWriteEnable) feature is
  enabled on the device, and a shader object is bound to the
  `VK_SHADER_STAGE_FRAGMENT_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then the `attachmentCount` parameter of most recent call
  to `vkCmdSetColorWriteEnableEXT` in the current command buffer
  **must** be greater than or equal to the number of color attachments
  in the current render pass instance

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07751"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07751"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07751  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT` dynamic state enabled then
  [vkCmdSetDiscardRectangleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command for each discard rectangle in
  [VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)::`discardRectangleCount`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07880"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07880"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07880  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT` dynamic state enabled
  then
  [vkCmdSetDiscardRectangleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09236"
  id="VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09236"></a>
  VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09236  
  If the [`VK_EXT_discard_rectangles`](VK_EXT_discard_rectangles.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetDiscardRectangleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleEnableEXT.html)
  in the current command buffer set `discardRectangleEnable` to
  `VK_TRUE`, then
  [vkCmdSetDiscardRectangleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08648"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08648"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08648  
  If the [`VK_EXT_discard_rectangles`](VK_EXT_discard_rectangles.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetDiscardRectangleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07881"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07881"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07881  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT` dynamic state enabled
  then
  [vkCmdSetDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleModeEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08649"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08649"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08649  
  If the [`VK_EXT_discard_rectangles`](VK_EXT_discard_rectangles.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetDiscardRectangleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleEnableEXT.html)
  in the current command buffer set `discardRectangleEnable` to
  `VK_TRUE`, then
  [vkCmdSetDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleModeEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08913"
  id="VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08913"></a>
  VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08913  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the
  [`dynamicRenderingUnusedAttachments`](#features-dynamicRenderingUnusedAttachments)
  feature is not enabled, and
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`
  was [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`depthAttachmentFormat`
  used to create the currently bound graphics pipeline **must** be equal
  to `VK_FORMAT_UNDEFINED`

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08914"
  id="VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08914"></a>
  VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08914  
  If current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the
  [`dynamicRenderingUnusedAttachments`](#features-dynamicRenderingUnusedAttachments)
  feature is not enabled, and
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`depthAttachmentFormat`
  used to create the currently bound graphics pipeline **must** be equal
  to the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) used to create
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08915"
  id="VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08915"></a>
  VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08915  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the
  [`dynamicRenderingUnusedAttachments`](#features-dynamicRenderingUnusedAttachments)
  feature is enabled,
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the value of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`depthAttachmentFormat`
  used to create the currently bound graphics pipeline was not equal to
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) used to create
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`,
  the value of the format **must** be `VK_FORMAT_UNDEFINED`

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08916"
  id="VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08916"></a>
  VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08916  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the
  [`dynamicRenderingUnusedAttachments`](#features-dynamicRenderingUnusedAttachments)
  feature is not enabled, and
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`
  was [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`stencilAttachmentFormat`
  used to create the currently bound graphics pipeline **must** be equal
  to `VK_FORMAT_UNDEFINED`

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08917"
  id="VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08917"></a>
  VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08917  
  If current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the
  [`dynamicRenderingUnusedAttachments`](#features-dynamicRenderingUnusedAttachments)
  feature is not enabled, and
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`stencilAttachmentFormat`
  used to create the currently bound graphics pipeline **must** be equal
  to the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) used to create
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08918"
  id="VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08918"></a>
  VUID-vkCmdDrawMeshTasksEXT-dynamicRenderingUnusedAttachments-08918  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the
  [`dynamicRenderingUnusedAttachments`](#features-dynamicRenderingUnusedAttachments)
  feature is enabled,
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the value of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`stencilAttachmentFormat`
  used to create the currently bound graphics pipeline was not equal to
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) used to create
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`,
  the value of the format **must** be `VK_FORMAT_UNDEFINED`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-imageView-06183"
  id="VUID-vkCmdDrawMeshTasksEXT-imageView-06183"></a>
  VUID-vkCmdDrawMeshTasksEXT-imageView-06183  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) and
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)::`imageView`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the currently bound
  graphics pipeline **must** have been created with
  `VK_PIPELINE_CREATE_RENDERING_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-imageView-06184"
  id="VUID-vkCmdDrawMeshTasksEXT-imageView-06184"></a>
  VUID-vkCmdDrawMeshTasksEXT-imageView-06184  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) and
  [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)::`imageView`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the currently bound
  graphics pipeline **must** have been created with
  `VK_PIPELINE_CREATE_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-colorAttachmentCount-06185"
  id="VUID-vkCmdDrawMeshTasksEXT-colorAttachmentCount-06185"></a>
  VUID-vkCmdDrawMeshTasksEXT-colorAttachmentCount-06185  
  If the currently bound pipeline was created with a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, and the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) with a
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`colorAttachmentCount`
  parameter greater than `0`, then each element of the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments` array
  with a `imageView` not equal to [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  **must** have been created with a sample count equal to the
  corresponding element of the `pColorAttachmentSamples` member of
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  used to create the currently bound graphics pipeline

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pDepthAttachment-06186"
  id="VUID-vkCmdDrawMeshTasksEXT-pDepthAttachment-06186"></a>
  VUID-vkCmdDrawMeshTasksEXT-pDepthAttachment-06186  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the currently bound
  pipeline was created with a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, and
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of the
  `depthStencilAttachmentSamples` member of
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  used to create the currently bound graphics pipeline **must** be equal
  to the sample count used to create
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pStencilAttachment-06187"
  id="VUID-vkCmdDrawMeshTasksEXT-pStencilAttachment-06187"></a>
  VUID-vkCmdDrawMeshTasksEXT-pStencilAttachment-06187  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the currently bound
  pipeline was created with a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, and
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of the
  `depthStencilAttachmentSamples` member of
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  used to create the currently bound graphics pipeline **must** be equal
  to the sample count used to create
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-multisampledRenderToSingleSampled-07285"
  id="VUID-vkCmdDrawMeshTasksEXT-multisampledRenderToSingleSampled-07285"></a>
  VUID-vkCmdDrawMeshTasksEXT-multisampledRenderToSingleSampled-07285  
  If the currently bound pipeline was created without a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, and the
  [`multisampledRenderToSingleSampled`](#features-multisampledRenderToSingleSampled)
  feature is not enabled, and the current render pass instance was begun
  with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) with a
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`colorAttachmentCount`
  parameter greater than `0`, then each element of the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments` array
  with a `imageView` not equal to [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  **must** have been created with a sample count equal to the value of
  `rasterizationSamples` for the currently bound graphics pipeline

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-multisampledRenderToSingleSampled-07286"
  id="VUID-vkCmdDrawMeshTasksEXT-multisampledRenderToSingleSampled-07286"></a>
  VUID-vkCmdDrawMeshTasksEXT-multisampledRenderToSingleSampled-07286  
  If the currently bound pipeline was created without a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, and the
  [`multisampledRenderToSingleSampled`](#features-multisampledRenderToSingleSampled)
  feature is not enabled, and
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of
  `rasterizationSamples` for the currently bound graphics pipeline
  **must** be equal to the sample count used to create
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-multisampledRenderToSingleSampled-07287"
  id="VUID-vkCmdDrawMeshTasksEXT-multisampledRenderToSingleSampled-07287"></a>
  VUID-vkCmdDrawMeshTasksEXT-multisampledRenderToSingleSampled-07287  
  If the currently bound pipeline was created without a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, and the
  [`multisampledRenderToSingleSampled`](#features-multisampledRenderToSingleSampled)
  feature is not enabled, and
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of
  `rasterizationSamples` for the currently bound graphics pipeline
  **must** be equal to the sample count used to create
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pNext-07935"
  id="VUID-vkCmdDrawMeshTasksEXT-pNext-07935"></a>
  VUID-vkCmdDrawMeshTasksEXT-pNext-07935  
  If this command has been called inside a render pass instance started
  with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), and the `pNext`
  chain of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) includes a
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)
  structure with `multisampledRenderToSingleSampledEnable` equal to
  `VK_TRUE`, then the value of `rasterizationSamples` for the currently
  bound graphics pipeline **must** be equal to
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)::`rasterizationSamples`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-renderPass-06198"
  id="VUID-vkCmdDrawMeshTasksEXT-renderPass-06198"></a>
  VUID-vkCmdDrawMeshTasksEXT-renderPass-06198  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), the currently bound
  pipeline **must** have been created with a
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`renderPass`
  equal to [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pColorAttachments-08963"
  id="VUID-vkCmdDrawMeshTasksEXT-pColorAttachments-08963"></a>
  VUID-vkCmdDrawMeshTasksEXT-pColorAttachments-08963  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), there is a graphics
  pipeline bound with a fragment shader that statically writes to a
  color attachment, the color write mask is not zero, color writes are
  enabled, and the corresponding element of the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments->imageView`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then the corresponding
  element of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`pColorAttachmentFormats`
  used to create the pipeline **must** not be `VK_FORMAT_UNDEFINED`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pDepthAttachment-08964"
  id="VUID-vkCmdDrawMeshTasksEXT-pDepthAttachment-08964"></a>
  VUID-vkCmdDrawMeshTasksEXT-pDepthAttachment-08964  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), there is a graphics
  pipeline bound, depth test is enabled, depth write is enabled, and the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then the
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`depthAttachmentFormat`
  used to create the pipeline **must** not be `VK_FORMAT_UNDEFINED`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pStencilAttachment-08965"
  id="VUID-vkCmdDrawMeshTasksEXT-pStencilAttachment-08965"></a>
  VUID-vkCmdDrawMeshTasksEXT-pStencilAttachment-08965  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), there is a graphics
  pipeline bound, stencil test is enabled and the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then the
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`stencilAttachmentFormat`
  used to create the pipeline **must** not be `VK_FORMAT_UNDEFINED`

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-primitivesGeneratedQueryWithRasterizerDiscard-06708"
  id="VUID-vkCmdDrawMeshTasksEXT-primitivesGeneratedQueryWithRasterizerDiscard-06708"></a>
  VUID-vkCmdDrawMeshTasksEXT-primitivesGeneratedQueryWithRasterizerDiscard-06708  
  If the
  [`primitivesGeneratedQueryWithRasterizerDiscard`](#features-primitivesGeneratedQueryWithRasterizerDiscard)
  feature is not enabled and the
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` query is active,
  [rasterization discard](#primsrast-discard) **must** not be enabled

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-primitivesGeneratedQueryWithNonZeroStreams-06709"
  id="VUID-vkCmdDrawMeshTasksEXT-primitivesGeneratedQueryWithNonZeroStreams-06709"></a>
  VUID-vkCmdDrawMeshTasksEXT-primitivesGeneratedQueryWithNonZeroStreams-06709  
  If the
  [`primitivesGeneratedQueryWithNonZeroStreams`](#features-primitivesGeneratedQueryWithNonZeroStreams)
  feature is not enabled and the
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` query is active, the bound
  graphics pipeline **must** not have been created with a non-zero value
  in
  `VkPipelineRasterizationStateStreamCreateInfoEXT`::`rasterizationStream`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07619"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07619"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07619  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_TESSELLATION_DOMAIN_ORIGIN_EXT` dynamic state
  enabled then
  [vkCmdSetTessellationDomainOriginEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetTessellationDomainOriginEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07620"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07620"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07620  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_DEPTH_CLAMP_ENABLE_EXT` dynamic state enabled then
  [vkCmdSetDepthClampEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClampEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09237"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09237"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09237  
  If a shader object is bound to the
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT` stage, then
  [vkCmdSetTessellationDomainOriginEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetTessellationDomainOriginEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08650"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08650"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08650  
  If the [`depthClamp`](#features-depthClamp) feature is enabled, and a
  shader object is bound to any graphics stage, and the most recent call
  to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetDepthClampEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClampEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07621"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07621"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07621  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_POLYGON_MODE_EXT` dynamic state enabled then
  [vkCmdSetPolygonModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPolygonModeEXT.html) **must** have
  been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08651"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08651"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08651  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then [vkCmdSetPolygonModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPolygonModeEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07622"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07622"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07622  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` dynamic state enabled
  then
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08652"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08652"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08652  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07623"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07623"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07623  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT` dynamic state enabled then
  [vkCmdSetSampleMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleMaskEXT.html) **must** have been
  called in the current command buffer prior to this drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08653"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08653"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08653  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then [vkCmdSetSampleMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleMaskEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07624"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07624"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07624  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_ALPHA_TO_COVERAGE_ENABLE_EXT` dynamic state enabled
  then
  [vkCmdSetAlphaToCoverageEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAlphaToCoverageEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-alphaToCoverageEnable-08919"
  id="VUID-vkCmdDrawMeshTasksEXT-alphaToCoverageEnable-08919"></a>
  VUID-vkCmdDrawMeshTasksEXT-alphaToCoverageEnable-08919  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_ALPHA_TO_COVERAGE_ENABLE_EXT` dynamic state enabled,
  and `alphaToCoverageEnable` was `VK_TRUE` in the last call to
  [vkCmdSetAlphaToCoverageEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAlphaToCoverageEnableEXT.html),
  then the [Fragment Output Interface](#interfaces-fragmentoutput)
  **must** contain a variable for the alpha `Component` word in
  `Location` 0 at `Index` 0

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08654"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08654"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08654  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetAlphaToCoverageEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAlphaToCoverageEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-alphaToCoverageEnable-08920"
  id="VUID-vkCmdDrawMeshTasksEXT-alphaToCoverageEnable-08920"></a>
  VUID-vkCmdDrawMeshTasksEXT-alphaToCoverageEnable-08920  
  If a shader object is bound to any graphics stage, and the most recent
  call to
  [vkCmdSetAlphaToCoverageEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAlphaToCoverageEnableEXT.html)
  in the current command buffer set `alphaToCoverageEnable` to
  `VK_TRUE`, then the [Fragment Output
  Interface](#interfaces-fragmentoutput) **must** contain a variable for
  the alpha `Component` word in `Location` 0 at `Index` 0

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07625"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07625"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07625  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_ALPHA_TO_ONE_ENABLE_EXT` dynamic state enabled then
  [vkCmdSetAlphaToOneEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAlphaToOneEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08655"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08655"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08655  
  If the [`alphaToOne`](#features-alphaToOne) feature is enabled, and a
  shader object is bound to any graphics stage, and the most recent call
  to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetAlphaToOneEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAlphaToOneEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07626"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07626"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07626  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_LOGIC_OP_ENABLE_EXT` dynamic state enabled then
  [vkCmdSetLogicOpEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEnableEXT.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08656"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08656"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08656  
  If the [`logicOp`](#features-logicOp) feature is enabled, and a shader
  object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT` stage, and the
  most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetLogicOpEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEnableEXT.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07627"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07627"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07627  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT` dynamic state enabled then
  [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08657"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08657"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08657  
  If a shader object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT`
  stage, and both the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE` and there are color attachments bound, then
  [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07628"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07628"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07628  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT` dynamic state enabled then
  [vkCmdSetColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEquationEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08658"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08658"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08658  
  If a shader object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT`
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEnableEXT.html) for
  any attachment set that attachment’s value in `pColorBlendEnables` to
  `VK_TRUE`, then
  [vkCmdSetColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEquationEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07629"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07629"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07629  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT` dynamic state enabled then
  [vkCmdSetColorWriteMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorWriteMaskEXT.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08659"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08659"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08659  
  If a shader object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT`
  stage, and both the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE` and there are color attachments bound, then
  [vkCmdSetColorWriteMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorWriteMaskEXT.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07630"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07630"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07630  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_RASTERIZATION_STREAM_EXT` dynamic state enabled then
  [vkCmdSetRasterizationStreamEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationStreamEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08660"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08660"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08660  
  If the [`geometryStreams`](#features-geometryStreams) feature is
  enabled, and a shader object is bound to the
  `VK_SHADER_STAGE_GEOMETRY_BIT` stage, then
  [vkCmdSetRasterizationStreamEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationStreamEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07631"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07631"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07631  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_CONSERVATIVE_RASTERIZATION_MODE_EXT` dynamic state
  enabled then
  [vkCmdSetConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetConservativeRasterizationModeEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08661"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08661"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08661  
  If the
  [`VK_EXT_conservative_rasterization`](VK_EXT_conservative_rasterization.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetConservativeRasterizationModeEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07632"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07632"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07632  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_EXTRA_PRIMITIVE_OVERESTIMATION_SIZE_EXT` dynamic
  state enabled then
  [vkCmdSetExtraPrimitiveOverestimationSizeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExtraPrimitiveOverestimationSizeEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08662"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08662"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08662  
  If the
  [`VK_EXT_conservative_rasterization`](VK_EXT_conservative_rasterization.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetConservativeRasterizationModeEXT.html)
  in the current command buffer set `conservativeRasterizationMode` to
  `VK_CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT`, then
  [vkCmdSetExtraPrimitiveOverestimationSizeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExtraPrimitiveOverestimationSizeEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07633"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07633"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07633  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_DEPTH_CLIP_ENABLE_EXT` dynamic state enabled then
  [vkCmdSetDepthClipEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClipEnableEXT.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08663"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08663"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08663  
  If the [`depthClipEnable`](#features-depthClipEnable) feature is
  enabled, and a shader object is bound to any graphics stage, then
  [vkCmdSetDepthClipEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClipEnableEXT.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07634"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07634"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07634  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT` dynamic state enabled
  then
  [vkCmdSetSampleLocationsEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08664"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08664"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08664  
  If the [`VK_EXT_sample_locations`](VK_EXT_sample_locations.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetSampleLocationsEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07635"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07635"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07635  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT` dynamic state enabled then
  [vkCmdSetColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendAdvancedEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09416"
  id="VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09416"></a>
  VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09416  
  If the
  [`VK_EXT_blend_operation_advanced`](VK_EXT_blend_operation_advanced.html)
  extension is enabled, and a shader object is bound to the
  `VK_SHADER_STAGE_FRAGMENT_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then at least one of
  [vkCmdSetColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEquationEXT.html)
  and
  [vkCmdSetColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendAdvancedEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07636"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07636"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07636  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_PROVOKING_VERTEX_MODE_EXT` dynamic state enabled
  then
  [vkCmdSetProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetProvokingVertexModeEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08665"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08665"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08665  
  If the [`VK_EXT_provoking_vertex`](VK_EXT_provoking_vertex.html)
  extension is enabled, and a shader object is bound to the
  `VK_SHADER_STAGE_VERTEX_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetProvokingVertexModeEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07637"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07637"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07637  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT` dynamic state enabled
  then
  [vkCmdSetLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineRasterizationModeEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08666"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08666"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08666  
  If the [`VK_KHR_line_rasterization`](VK_KHR_line_rasterization.html)
  or [`VK_EXT_line_rasterization`](VK_EXT_line_rasterization.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetPolygonModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPolygonModeEXT.html) in the current
  command buffer set `polygonMode` to `VK_POLYGON_MODE_LINE`, then
  [vkCmdSetLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineRasterizationModeEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08667"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08667"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08667  
  If the [`VK_KHR_line_rasterization`](VK_KHR_line_rasterization.html)
  or [`VK_EXT_line_rasterization`](VK_EXT_line_rasterization.html)
  extension is enabled, and a shader object is bound to the
  `VK_SHADER_STAGE_VERTEX_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetPrimitiveTopology](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPrimitiveTopology.html) in the
  current command buffer set `primitiveTopology` to any line topology,
  then
  [vkCmdSetLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineRasterizationModeEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08668"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08668"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08668  
  If the [`VK_KHR_line_rasterization`](VK_KHR_line_rasterization.html)
  or [`VK_EXT_line_rasterization`](VK_EXT_line_rasterization.html)
  extension is enabled, and a shader object that outputs line primitives
  is bound to the `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT` or
  `VK_SHADER_STAGE_GEOMETRY_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineRasterizationModeEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07638"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07638"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07638  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_LINE_STIPPLE_ENABLE_EXT` dynamic state enabled then
  [vkCmdSetLineStippleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineStippleEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08669"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08669"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08669  
  If the [`VK_KHR_line_rasterization`](VK_KHR_line_rasterization.html)
  or [`VK_EXT_line_rasterization`](VK_EXT_line_rasterization.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetPolygonModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPolygonModeEXT.html) in the current
  command buffer set `polygonMode` to `VK_POLYGON_MODE_LINE`, then
  [vkCmdSetLineStippleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineStippleEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08670"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08670"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08670  
  If the [`VK_KHR_line_rasterization`](VK_KHR_line_rasterization.html)
  or [`VK_EXT_line_rasterization`](VK_EXT_line_rasterization.html)
  extension is enabled, and a shader object is bound to the
  `VK_SHADER_STAGE_VERTEX_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetPrimitiveTopology](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPrimitiveTopology.html) in the
  current command buffer set `primitiveTopology` to any line topology,
  then [vkCmdSetLineStippleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineStippleEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08671"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08671"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08671  
  If the [`VK_KHR_line_rasterization`](VK_KHR_line_rasterization.html)
  or [`VK_EXT_line_rasterization`](VK_EXT_line_rasterization.html)
  extension is enabled, and a shader object that outputs line primitives
  is bound to the `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT` or
  `VK_SHADER_STAGE_GEOMETRY_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetLineStippleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineStippleEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07849"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07849"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07849  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_LINE_STIPPLE_KHR` dynamic state enabled then
  [vkCmdSetLineStippleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineStippleKHR.html) **must** have
  been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08672"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08672"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08672  
  If the [`VK_KHR_line_rasterization`](VK_KHR_line_rasterization.html)
  or [`VK_EXT_line_rasterization`](VK_EXT_line_rasterization.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetLineStippleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineStippleEnableEXT.html) in
  the current command buffer set `stippledLineEnable` to `VK_TRUE`, then
  [vkCmdSetLineStippleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineStippleEXT.html) **must** have
  been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07639"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07639"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07639  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_DEPTH_CLIP_NEGATIVE_ONE_TO_ONE_EXT` dynamic state
  enabled then
  [vkCmdSetDepthClipNegativeOneToOneEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClipNegativeOneToOneEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08673"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08673"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08673  
  If the [`depthClipControl`](#features-depthClipControl) feature is
  enabled, and a shader object is bound to any graphics stage, then
  [vkCmdSetDepthClipNegativeOneToOneEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClipNegativeOneToOneEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07640"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07640"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07640  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_ENABLE_NV` dynamic state enabled
  then
  [vkCmdSetViewportWScalingEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWScalingEnableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08674"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08674"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08674  
  If the [`VK_NV_clip_space_w_scaling`](VK_NV_clip_space_w_scaling.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, then
  [vkCmdSetViewportWScalingEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWScalingEnableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07641"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07641"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07641  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VIEWPORT_SWIZZLE_NV` dynamic state enabled then
  [vkCmdSetViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportSwizzleNV.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08675"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08675"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08675  
  If the [`VK_NV_viewport_swizzle`](VK_NV_viewport_swizzle.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, then
  [vkCmdSetViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportSwizzleNV.html) **must**
  have been called in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07642"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07642"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07642  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_ENABLE_NV` dynamic state enabled
  then
  [vkCmdSetCoverageToColorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorEnableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08676"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08676"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08676  
  If the
  [`VK_NV_fragment_coverage_to_color`](VK_NV_fragment_coverage_to_color.html)
  extension is enabled, and a shader object is bound to the
  `VK_SHADER_STAGE_FRAGMENT_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetCoverageToColorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorEnableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07643"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07643"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07643  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_LOCATION_NV` dynamic state enabled
  then
  [vkCmdSetCoverageToColorLocationNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorLocationNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08677"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08677"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08677  
  If the
  [`VK_NV_fragment_coverage_to_color`](VK_NV_fragment_coverage_to_color.html)
  extension is enabled, and a shader object is bound to the
  `VK_SHADER_STAGE_FRAGMENT_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetCoverageToColorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorEnableNV.html)
  in the current command buffer set `coverageToColorEnable` to
  `VK_TRUE`, then
  [vkCmdSetCoverageToColorLocationNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorLocationNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07644"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07644"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07644  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COVERAGE_MODULATION_MODE_NV` dynamic state enabled
  then
  [vkCmdSetCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationModeNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08678"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08678"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08678  
  If the
  [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationModeNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07645"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07645"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07645  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_ENABLE_NV` dynamic state
  enabled then
  [vkCmdSetCoverageModulationTableEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationTableEnableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08679"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08679"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08679  
  If the
  [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationModeNV.html)
  in the current command buffer set coverageModulationMode to any value
  other than `VK_COVERAGE_MODULATION_MODE_NONE_NV`, then
  [vkCmdSetCoverageModulationTableEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationTableEnableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07646"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07646"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07646  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_NV` dynamic state enabled
  then
  [vkCmdSetCoverageModulationTableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationTableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08680"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08680"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08680  
  If the
  [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the most recent call to
  [vkCmdSetCoverageModulationTableEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationTableEnableNV.html)
  in the current command buffer set `coverageModulationTableEnable` to
  `VK_TRUE`, then
  [vkCmdSetCoverageModulationTableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationTableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07647"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07647"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07647  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SHADING_RATE_IMAGE_ENABLE_NV` dynamic state enabled
  then
  [vkCmdSetShadingRateImageEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetShadingRateImageEnableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pipelineFragmentShadingRate-09238"
  id="VUID-vkCmdDrawMeshTasksEXT-pipelineFragmentShadingRate-09238"></a>
  VUID-vkCmdDrawMeshTasksEXT-pipelineFragmentShadingRate-09238  
  If the
  [`pipelineFragmentShadingRate`](#features-pipelineFragmentShadingRate)
  feature is enabled, and a shader object is bound to the
  `VK_SHADER_STAGE_FRAGMENT_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set rasterizerDiscardEnable to
  `VK_FALSE`, then
  [vkCmdSetFragmentShadingRateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFragmentShadingRateKHR.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08681"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08681"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08681  
  If the [`shadingRateImage`](#features-shadingRateImage) feature is
  enabled, and a shader object is bound to any graphics stage, and the
  most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetShadingRateImageEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetShadingRateImageEnableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07648"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07648"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07648  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_REPRESENTATIVE_FRAGMENT_TEST_ENABLE_NV` dynamic
  state enabled then
  [vkCmdSetRepresentativeFragmentTestEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRepresentativeFragmentTestEnableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08682"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08682"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08682  
  If the
  [`representativeFragmentTest`](#features-representativeFragmentTest)
  feature is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetRepresentativeFragmentTestEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRepresentativeFragmentTestEnableNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07649"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07649"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07649  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COVERAGE_REDUCTION_MODE_NV` dynamic state enabled
  then
  [vkCmdSetCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageReductionModeNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08683"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08683"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08683  
  If the [`coverageReductionMode`](#features-coverageReductionMode)
  feature is enabled, and a shader object is bound to any graphics
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageReductionModeNV.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pColorBlendEnables-07470"
  id="VUID-vkCmdDrawMeshTasksEXT-pColorBlendEnables-07470"></a>
  VUID-vkCmdDrawMeshTasksEXT-pColorBlendEnables-07470  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT` state enabled and the last
  call to
  [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEnableEXT.html) set
  `pColorBlendEnables` for any attachment to `VK_TRUE`, then for those
  attachments in the subpass the corresponding image view’s [format
  features](#resources-image-view-format-features) **must** contain
  `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-07471"
  id="VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-07471"></a>
  VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-07471  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` state enabled, and the
  current subpass does not use any color and/or depth/stencil
  attachments, then the `rasterizationSamples` in the last call to
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)
  **must** follow the rules for a [zero-attachment
  subpass](#renderpass-noattachments)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-samples-07472"
  id="VUID-vkCmdDrawMeshTasksEXT-samples-07472"></a>
  VUID-vkCmdDrawMeshTasksEXT-samples-07472  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT` state enabled and the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` state disabled, then the
  `samples` parameter in the last call to
  [vkCmdSetSampleMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleMaskEXT.html) **must** be
  greater or equal to the
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)::`rasterizationSamples`
  parameter used to create the bound graphics pipeline

- <a href="#VUID-vkCmdDrawMeshTasksEXT-samples-07473"
  id="VUID-vkCmdDrawMeshTasksEXT-samples-07473"></a>
  VUID-vkCmdDrawMeshTasksEXT-samples-07473  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT` state and
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` states enabled, then the
  `samples` parameter in the last call to
  [vkCmdSetSampleMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleMaskEXT.html) **must** be
  greater or equal to the `rasterizationSamples` parameter in the last
  call to
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-07474"
  id="VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-07474"></a>
  VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-07474  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` state enabled, and
  neither the
  [`VK_AMD_mixed_attachment_samples`](VK_AMD_mixed_attachment_samples.html)
  nor the
  [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
  extensions are enabled, then the `rasterizationSamples` in the last
  call to
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)
  **must** be the same as the current subpass color and/or depth/stencil
  attachments

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09211"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09211"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09211  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` state enabled, or a
  shader object is bound to any graphics stage, and the current render
  pass instance includes a
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)
  structure with `multisampledRenderToSingleSampledEnable` equal to
  `VK_TRUE`, then the `rasterizationSamples` in the last call to
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)
  **must** be the same as the `rasterizationSamples` member of that
  structure

- <a href="#VUID-vkCmdDrawMeshTasksEXT-firstAttachment-07476"
  id="VUID-vkCmdDrawMeshTasksEXT-firstAttachment-07476"></a>
  VUID-vkCmdDrawMeshTasksEXT-firstAttachment-07476  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT` dynamic state enabled then
  [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command, and the attachments specified by the
  `firstAttachment` and `attachmentCount` parameters of
  `vkCmdSetColorBlendEnableEXT` calls **must** specify an enable for all
  active color attachments in the current subpass

- <a href="#VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09417"
  id="VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09417"></a>
  VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09417  
  If a shader object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT`
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command, and the attachments specified by the
  `firstAttachment` and `attachmentCount` parameters of
  `vkCmdSetColorBlendEnableEXT` calls **must** specify an enable for all
  active color attachments in the current subpass

- <a href="#VUID-vkCmdDrawMeshTasksEXT-firstAttachment-07477"
  id="VUID-vkCmdDrawMeshTasksEXT-firstAttachment-07477"></a>
  VUID-vkCmdDrawMeshTasksEXT-firstAttachment-07477  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT` dynamic state enabled then
  [vkCmdSetColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEquationEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command, and the attachments specified by the
  `firstAttachment` and `attachmentCount` parameters of
  `vkCmdSetColorBlendEquationEXT` calls **must** specify the blend
  equations for all active color attachments in the current subpass
  where blending is enabled

- <a href="#VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09418"
  id="VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09418"></a>
  VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09418  
  If a shader object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT`
  stage, and both the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE` and there are color attachments bound, then
  [vkCmdSetColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEquationEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command, and the attachments specified by the
  `firstAttachment` and `attachmentCount` parameters of
  `vkCmdSetColorBlendEquationEXT` calls **must** specify the blend
  equations for all active color attachments in the current subpass
  where blending is enabled

- <a href="#VUID-vkCmdDrawMeshTasksEXT-firstAttachment-07478"
  id="VUID-vkCmdDrawMeshTasksEXT-firstAttachment-07478"></a>
  VUID-vkCmdDrawMeshTasksEXT-firstAttachment-07478  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT` dynamic state enabled then
  [vkCmdSetColorWriteMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorWriteMaskEXT.html) **must**
  have been called in the current command buffer prior to this drawing
  command, and the attachments specified by the `firstAttachment` and
  `attachmentCount` parameters of `vkCmdSetColorWriteMaskEXT` calls
  **must** specify the color write mask for all active color attachments
  in the current subpass

- <a href="#VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09419"
  id="VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09419"></a>
  VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09419  
  If a shader object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT`
  stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetColorWriteMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorWriteMaskEXT.html) **must**
  have been called in the current command buffer prior to this drawing
  command, and the attachments specified by the `firstAttachment` and
  `attachmentCount` parameters of `vkCmdSetColorWriteMaskEXT` calls
  **must** specify the color write mask for all active color attachments
  in the current subpass

- <a href="#VUID-vkCmdDrawMeshTasksEXT-firstAttachment-07479"
  id="VUID-vkCmdDrawMeshTasksEXT-firstAttachment-07479"></a>
  VUID-vkCmdDrawMeshTasksEXT-firstAttachment-07479  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT` dynamic state enabled then
  [vkCmdSetColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendAdvancedEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command, and the attachments specified by the
  `firstAttachment` and `attachmentCount` parameters of
  `vkCmdSetColorBlendAdvancedEXT` calls **must** specify the advanced
  blend equations for all active color attachments in the current
  subpass where blending is enabled

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-advancedBlendMaxColorAttachments-07480"
  id="VUID-vkCmdDrawMeshTasksEXT-advancedBlendMaxColorAttachments-07480"></a>
  VUID-vkCmdDrawMeshTasksEXT-advancedBlendMaxColorAttachments-07480  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT` and
  `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT` dynamic states enabled and
  the last calls to
  [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEnableEXT.html) and
  [vkCmdSetColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendAdvancedEXT.html)
  have enabled advanced blending, then the number of active color
  attachments in the current subpass **must** not exceed
  [`advancedBlendMaxColorAttachments`](#limits-advancedBlendMaxColorAttachments)

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-primitivesGeneratedQueryWithNonZeroStreams-07481"
  id="VUID-vkCmdDrawMeshTasksEXT-primitivesGeneratedQueryWithNonZeroStreams-07481"></a>
  VUID-vkCmdDrawMeshTasksEXT-primitivesGeneratedQueryWithNonZeroStreams-07481  
  If the
  [`primitivesGeneratedQueryWithNonZeroStreams`](#features-primitivesGeneratedQueryWithNonZeroStreams)
  feature is not enabled and the
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` query is active, and the
  bound graphics pipeline was created with
  `VK_DYNAMIC_STATE_RASTERIZATION_STREAM_EXT` state enabled, the last
  call to
  [vkCmdSetRasterizationStreamEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationStreamEXT.html)
  **must** have set the `rasterizationStream` to zero

- <a href="#VUID-vkCmdDrawMeshTasksEXT-sampleLocationsPerPixel-07482"
  id="VUID-vkCmdDrawMeshTasksEXT-sampleLocationsPerPixel-07482"></a>
  VUID-vkCmdDrawMeshTasksEXT-sampleLocationsPerPixel-07482  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT` state enabled and the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` state disabled, then the
  `sampleLocationsPerPixel` member of `pSampleLocationsInfo` in the last
  call to [vkCmdSetSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEXT.html)
  **must** equal the `rasterizationSamples` member of the
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
  structure the bound graphics pipeline has been created with

- <a href="#VUID-vkCmdDrawMeshTasksEXT-sampleLocationsPerPixel-07483"
  id="VUID-vkCmdDrawMeshTasksEXT-sampleLocationsPerPixel-07483"></a>
  VUID-vkCmdDrawMeshTasksEXT-sampleLocationsPerPixel-07483  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT` state enabled and the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` state enabled, then the
  `sampleLocationsPerPixel` member of `pSampleLocationsInfo` in the last
  call to [vkCmdSetSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEXT.html)
  **must** equal the `rasterizationSamples` parameter of the last call
  to
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07484"
  id="VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07484"></a>
  VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07484  
  If a shader object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT`
  stage, or the bound graphics pipeline was created with the
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT` state enabled, and
  `sampleLocationsEnable` was `VK_TRUE` in the last call to
  [vkCmdSetSampleLocationsEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEnableEXT.html),
  and the current subpass has a depth/stencil attachment, then that
  attachment **must** have been created with the
  `VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT` bit set

- <a href="#VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07485"
  id="VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07485"></a>
  VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07485  
  If a shader object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT`
  stage, or the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT` state enabled and the
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT` state enabled, and if
  `sampleLocationsEnable` was `VK_TRUE` in the last call to
  [vkCmdSetSampleLocationsEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEnableEXT.html),
  then the `sampleLocationsInfo.sampleLocationGridSize.width` in the
  last call to
  [vkCmdSetSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEXT.html) **must**
  evenly divide
  [VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisamplePropertiesEXT.html)::`sampleLocationGridSize.width`
  as returned by
  [vkGetPhysicalDeviceMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMultisamplePropertiesEXT.html)
  with a `samples` parameter equaling `rasterizationSamples`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07486"
  id="VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07486"></a>
  VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07486  
  If a shader object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT`
  stage, or the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT` state enabled and the
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT` state enabled, and if
  `sampleLocationsEnable` was `VK_TRUE` in the last call to
  [vkCmdSetSampleLocationsEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEnableEXT.html),
  then the `sampleLocationsInfo.sampleLocationGridSize.height` in the
  last call to
  [vkCmdSetSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEXT.html) **must**
  evenly divide
  [VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisamplePropertiesEXT.html)::`sampleLocationGridSize.height`
  as returned by
  [vkGetPhysicalDeviceMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMultisamplePropertiesEXT.html)
  with a `samples` parameter equaling `rasterizationSamples`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07487"
  id="VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07487"></a>
  VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07487  
  If a shader object is bound to the `VK_SHADER_STAGE_FRAGMENT_BIT`
  stage, or the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT` state enabled, and if
  `sampleLocationsEnable` was `VK_TRUE` in the last call to
  [vkCmdSetSampleLocationsEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEnableEXT.html),
  the fragment shader code **must** not statically use the extended
  instruction `InterpolateAtSample`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07936"
  id="VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07936"></a>
  VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07936  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT` state disabled and the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` state enabled, the
  `sampleLocationsEnable` member of a
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)::`sampleLocationsEnable`
  in the bound graphics pipeline is `VK_TRUE` or
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT` state enabled, then,
  `sampleLocationsInfo.sampleLocationGridSize.width` **must** evenly
  divide
  [VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisamplePropertiesEXT.html)::`sampleLocationGridSize.width`
  as returned by
  [vkGetPhysicalDeviceMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMultisamplePropertiesEXT.html)
  with a `samples` parameter equaling the value of
  `rasterizationSamples` in the last call to
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07937"
  id="VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07937"></a>
  VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07937  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT` state disabled and the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` state enabled, the
  `sampleLocationsEnable` member of a
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)::`sampleLocationsEnable`
  in the bound graphics pipeline is `VK_TRUE` or
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT` state enabled, then,
  `sampleLocationsInfo.sampleLocationGridSize.height` **must** evenly
  divide
  [VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisamplePropertiesEXT.html)::`sampleLocationGridSize.height`
  as returned by
  [vkGetPhysicalDeviceMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMultisamplePropertiesEXT.html)
  with a `samples` parameter equaling the value of
  `rasterizationSamples` in the last call to
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07938"
  id="VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07938"></a>
  VUID-vkCmdDrawMeshTasksEXT-sampleLocationsEnable-07938  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT` state disabled and the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` state enabled, the
  `sampleLocationsEnable` member of a
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)::`sampleLocationsEnable`
  in the bound graphics pipeline is `VK_TRUE` or
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT` state enabled, then,
  `sampleLocationsInfo.sampleLocationsPerPixel` **must** equal
  `rasterizationSamples` in the last call to
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-coverageModulationTableEnable-07488"
  id="VUID-vkCmdDrawMeshTasksEXT-coverageModulationTableEnable-07488"></a>
  VUID-vkCmdDrawMeshTasksEXT-coverageModulationTableEnable-07488  
  If a shader object is bound to any graphics stage or the bound
  graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_ENABLE_NV` state enabled,
  and the last call to
  [vkCmdSetCoverageModulationTableEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationTableEnableNV.html)
  set `coverageModulationTableEnable` to `VK_TRUE`, then the
  `coverageModulationTableCount` parameter in the last call to
  [vkCmdSetCoverageModulationTableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationTableNV.html)
  **must** equal the current `rasterizationSamples` divided by the
  number of color samples in the current subpass

- <a href="#VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-07489"
  id="VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-07489"></a>
  VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-07489  
  If the
  [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
  extension is enabled, and if current subpass has a depth/stencil
  attachment and depth test, stencil test, or depth bounds test are
  enabled in the currently bound pipeline state, then the current
  `rasterizationSamples` **must** be the same as the sample count of the
  depth/stencil attachment

- <a href="#VUID-vkCmdDrawMeshTasksEXT-coverageToColorEnable-07490"
  id="VUID-vkCmdDrawMeshTasksEXT-coverageToColorEnable-07490"></a>
  VUID-vkCmdDrawMeshTasksEXT-coverageToColorEnable-07490  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_ENABLE_NV` state enabled and the
  last call to
  [vkCmdSetCoverageToColorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorEnableNV.html)
  set the `coverageToColorEnable` to `VK_TRUE`, then the current subpass
  **must** have a color attachment at the location selected by the last
  call to
  [vkCmdSetCoverageToColorLocationNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorLocationNV.html)
  `coverageToColorLocation`, with a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of
  `VK_FORMAT_R8_UINT`, `VK_FORMAT_R8_SINT`, `VK_FORMAT_R16_UINT`,
  `VK_FORMAT_R16_SINT`, `VK_FORMAT_R32_UINT`, or `VK_FORMAT_R32_SINT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09420"
  id="VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09420"></a>
  VUID-vkCmdDrawMeshTasksEXT-rasterizerDiscardEnable-09420  
  If the
  [`VK_NV_fragment_coverage_to_color`](VK_NV_fragment_coverage_to_color.html)
  extension is enabled, and a shader object is bound to the
  `VK_SHADER_STAGE_FRAGMENT_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, and the last call to
  [vkCmdSetCoverageToColorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorEnableNV.html)
  set the `coverageToColorEnable` to `VK_TRUE`, then the current subpass
  **must** have a color attachment at the location selected by the last
  call to
  [vkCmdSetCoverageToColorLocationNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorLocationNV.html)
  `coverageToColorLocation`, with a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of
  `VK_FORMAT_R8_UINT`, `VK_FORMAT_R8_SINT`, `VK_FORMAT_R16_UINT`,
  `VK_FORMAT_R16_SINT`, `VK_FORMAT_R32_UINT`, or `VK_FORMAT_R32_SINT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-coverageReductionMode-07491"
  id="VUID-vkCmdDrawMeshTasksEXT-coverageReductionMode-07491"></a>
  VUID-vkCmdDrawMeshTasksEXT-coverageReductionMode-07491  
  If this
  [`VK_NV_coverage_reduction_mode`](VK_NV_coverage_reduction_mode.html)
  extension is enabled, the bound graphics pipeline state was created
  with the `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_ENABLE_NV` and
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` states enabled, the
  current coverage reduction mode `coverageReductionMode`, then the
  current `rasterizationSamples`, and the sample counts for the color
  and depth/stencil attachments (if the subpass has them) **must** be a
  valid combination returned by
  [vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-viewportCount-07492"
  id="VUID-vkCmdDrawMeshTasksEXT-viewportCount-07492"></a>
  VUID-vkCmdDrawMeshTasksEXT-viewportCount-07492  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` dynamic state enabled, but not
  the `VK_DYNAMIC_STATE_VIEWPORT_SWIZZLE_NV` dynamic state enabled, then
  the bound graphics pipeline **must** have been created with
  [VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html)::`viewportCount`
  greater or equal to the `viewportCount` parameter in the last call to
  [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-viewportCount-07493"
  id="VUID-vkCmdDrawMeshTasksEXT-viewportCount-07493"></a>
  VUID-vkCmdDrawMeshTasksEXT-viewportCount-07493  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` and
  `VK_DYNAMIC_STATE_VIEWPORT_SWIZZLE_NV` dynamic states enabled then the
  `viewportCount` parameter in the last call to
  [vkCmdSetViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportSwizzleNV.html) **must**
  be greater than or equal to the `viewportCount` parameter in the last
  call to [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-viewportCount-09421"
  id="VUID-vkCmdDrawMeshTasksEXT-viewportCount-09421"></a>
  VUID-vkCmdDrawMeshTasksEXT-viewportCount-09421  
  If the [`VK_NV_viewport_swizzle`](VK_NV_viewport_swizzle.html)
  extension is enabled, and a shader object is bound to any graphics
  stage, then the `viewportCount` parameter in the last call to
  [vkCmdSetViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportSwizzleNV.html) **must**
  be greater than or equal to the `viewportCount` parameter in the last
  call to [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-07494"
  id="VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-07494"></a>
  VUID-vkCmdDrawMeshTasksEXT-rasterizationSamples-07494  
  If the
  [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
  extension is enabled, and if the current subpass has any color
  attachments and `rasterizationSamples` of the last call to
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)
  is greater than the number of color samples, then the pipeline
  `sampleShadingEnable` **must** be `VK_FALSE`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-stippledLineEnable-07495"
  id="VUID-vkCmdDrawMeshTasksEXT-stippledLineEnable-07495"></a>
  VUID-vkCmdDrawMeshTasksEXT-stippledLineEnable-07495  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_LINE_STIPPLE_ENABLE_EXT` or
  `VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT` dynamic states enabled,
  and if the current `stippledLineEnable` state is `VK_TRUE` and the
  current `lineRasterizationMode` state is
  `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_KHR`, then the
  [`stippledRectangularLines`](#features-stippledRectangularLines)
  feature **must** be enabled

- <a href="#VUID-vkCmdDrawMeshTasksEXT-stippledLineEnable-07496"
  id="VUID-vkCmdDrawMeshTasksEXT-stippledLineEnable-07496"></a>
  VUID-vkCmdDrawMeshTasksEXT-stippledLineEnable-07496  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_LINE_STIPPLE_ENABLE_EXT` or
  `VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT` dynamic states enabled,
  and if the current `stippledLineEnable` state is `VK_TRUE` and the
  current `lineRasterizationMode` state is
  `VK_LINE_RASTERIZATION_MODE_BRESENHAM_KHR`, then the
  [`stippledBresenhamLines`](#features-stippledBresenhamLines) feature
  **must** be enabled

- <a href="#VUID-vkCmdDrawMeshTasksEXT-stippledLineEnable-07497"
  id="VUID-vkCmdDrawMeshTasksEXT-stippledLineEnable-07497"></a>
  VUID-vkCmdDrawMeshTasksEXT-stippledLineEnable-07497  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_LINE_STIPPLE_ENABLE_EXT` or
  `VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT` dynamic states enabled,
  and if the current `stippledLineEnable` state is `VK_TRUE` and the
  current `lineRasterizationMode` state is
  `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_KHR`, then the
  [`stippledSmoothLines`](#features-stippledSmoothLines) feature
  **must** be enabled

- <a href="#VUID-vkCmdDrawMeshTasksEXT-stippledLineEnable-07498"
  id="VUID-vkCmdDrawMeshTasksEXT-stippledLineEnable-07498"></a>
  VUID-vkCmdDrawMeshTasksEXT-stippledLineEnable-07498  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_LINE_STIPPLE_ENABLE_EXT` or
  `VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT` dynamic states enabled,
  and if the current `stippledLineEnable` state is `VK_TRUE` and the
  current `lineRasterizationMode` state is
  `VK_LINE_RASTERIZATION_MODE_DEFAULT_KHR`, then the
  [`stippledRectangularLines`](#features-stippledRectangularLines)
  feature **must** be enabled and
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`strictLines`
  **must** be `VK_TRUE`

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-conservativePointAndLineRasterization-07499"
  id="VUID-vkCmdDrawMeshTasksEXT-conservativePointAndLineRasterization-07499"></a>
  VUID-vkCmdDrawMeshTasksEXT-conservativePointAndLineRasterization-07499  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_CONSERVATIVE_RASTERIZATION_MODE_EXT` dynamic state
  enabled,
  [`conservativePointAndLineRasterization`](#limits-conservativePointAndLineRasterization)
  is not supported, and the effective primitive topology output by the
  last pre-rasterization shader stage is a line or point, then the
  `conservativeRasterizationMode` set by the last call to
  [vkCmdSetConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetConservativeRasterizationModeEXT.html)
  **must** be `VK_CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-stage-07073"
  id="VUID-vkCmdDrawMeshTasksEXT-stage-07073"></a>
  VUID-vkCmdDrawMeshTasksEXT-stage-07073  
  If the currently bound pipeline was created with the
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)::`stage`
  member of an element of
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`pStages`
  set to `VK_SHADER_STAGE_VERTEX_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT` or
  `VK_SHADER_STAGE_GEOMETRY_BIT`, then [Mesh Shader
  Queries](#queries-mesh-shader) **must** not be active

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08877"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08877"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08877  
  If the bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_ATTACHMENT_FEEDBACK_LOOP_ENABLE_EXT` dynamic state
  [vkCmdSetAttachmentFeedbackLoopEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAttachmentFeedbackLoopEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07850"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07850"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07850  
  If dynamic state was inherited from
  [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html),
  it **must** be set in the current command buffer prior to this drawing
  command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08684"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08684"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08684  
  If there is no bound graphics pipeline, `vkCmdBindShadersEXT` **must**
  have been called in the current command buffer with `pStages` with an
  element of `VK_SHADER_STAGE_VERTEX_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08685"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08685"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08685  
  If there is no bound graphics pipeline, and the
  [`tessellationShader`](#features-tessellationShader) feature is
  enabled, `vkCmdBindShadersEXT` **must** have been called in the
  current command buffer with `pStages` with an element of
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08686"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08686"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08686  
  If there is no bound graphics pipeline, and the
  [`tessellationShader`](#features-tessellationShader) feature is
  enabled, `vkCmdBindShadersEXT` **must** have been called in the
  current command buffer with `pStages` with an element of
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08687"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08687"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08687  
  If there is no bound graphics pipeline, and the
  [`geometryShader`](#features-geometryShader) feature is enabled,
  `vkCmdBindShadersEXT` **must** have been called in the current command
  buffer with `pStages` with an element of
  `VK_SHADER_STAGE_GEOMETRY_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08688"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08688"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08688  
  If there is no bound graphics pipeline, `vkCmdBindShadersEXT` **must**
  have been called in the current command buffer with `pStages` with an
  element of `VK_SHADER_STAGE_FRAGMENT_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08689"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08689"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08689  
  If there is no bound graphics pipeline, and the
  [`taskShader`](#features-taskShader) feature is enabled,
  `vkCmdBindShadersEXT` **must** have been called in the current command
  buffer with `pStages` with an element of
  `VK_SHADER_STAGE_TASK_BIT_EXT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08690"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08690"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08690  
  If there is no bound graphics pipeline, and the
  [`meshShader`](#features-meshShader) feature is enabled,
  `vkCmdBindShadersEXT` **must** have been called in the current command
  buffer with `pStages` with an element of
  `VK_SHADER_STAGE_MESH_BIT_EXT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08693"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08693"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08693  
  If there is no bound graphics pipeline, and at least one of the
  [`taskShader`](#features-taskShader) and
  [`meshShader`](#features-meshShader) features is enabled, one of the
  `VK_SHADER_STAGE_VERTEX_BIT` or `VK_SHADER_STAGE_MESH_BIT_EXT` stages
  **must** have a valid `VkShaderEXT` bound, and the other **must** have
  no `VkShaderEXT` bound

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08694"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08694"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08694  
  If there is no bound graphics pipeline, and both the
  [`taskShader`](#features-taskShader) and
  [`meshShader`](#features-meshShader) features are enabled, and a valid
  `VkShaderEXT` is bound the to the `VK_SHADER_STAGE_MESH_BIT_EXT`
  stage, and that `VkShaderEXT` was created without the
  `VK_SHADER_CREATE_NO_TASK_SHADER_BIT_EXT` flag, a valid `VkShaderEXT`
  **must** be bound to the `VK_SHADER_STAGE_TASK_BIT_EXT` stage

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08695"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08695"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08695  
  If there is no bound graphics pipeline, and both the
  [`taskShader`](#features-taskShader) and
  [`meshShader`](#features-meshShader) features are enabled, and a valid
  `VkShaderEXT` is bound the to the `VK_SHADER_STAGE_MESH_BIT_EXT`
  stage, and that `VkShaderEXT` was created with the
  `VK_SHADER_CREATE_NO_TASK_SHADER_BIT_EXT` flag, there **must** be no
  `VkShaderEXT` bound to the `VK_SHADER_STAGE_TASK_BIT_EXT` stage

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08696"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08696"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08696  
  If there is no bound graphics pipeline, and a valid `VkShaderEXT` is
  bound to the `VK_SHADER_STAGE_VERTEX_BIT` stage, there **must** be no
  `VkShaderEXT` bound to either the `VK_SHADER_STAGE_TASK_BIT_EXT` stage
  or the `VK_SHADER_STAGE_MESH_BIT_EXT` stage

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08698"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08698"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08698  
  If any graphics shader is bound which was created with the
  `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT` flag, then all shaders created
  with the `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT` flag in the same
  [vkCreateShadersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateShadersEXT.html) call **must** also be
  bound

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08699"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08699"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08699  
  If any graphics shader is bound which was created with the
  `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT` flag, any stages in between
  stages whose shaders which did not create a shader with the
  `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT` flag as part of the same
  [vkCreateShadersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateShadersEXT.html) call **must** not have
  any `VkShaderEXT` bound

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08878"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08878"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08878  
  All bound graphics shader objects **must** have been created with
  identical or identically defined push constant ranges

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08879"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08879"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08879  
  All bound graphics shader objects **must** have been created with
  identical or identically defined arrays of descriptor set layouts

- <a href="#VUID-vkCmdDrawMeshTasksEXT-colorAttachmentCount-09372"
  id="VUID-vkCmdDrawMeshTasksEXT-colorAttachmentCount-09372"></a>
  VUID-vkCmdDrawMeshTasksEXT-colorAttachmentCount-09372  
  If the current render pass instance was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) and a
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`colorAttachmentCount` equal
  to `1`, a color attachment with a resolve mode of
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`, and a fragment
  shader is bound, it **must** not declare the `DepthReplacing` or
  `StencilRefReplacingEXT` execution modes

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-08880"
  id="VUID-vkCmdDrawMeshTasksEXT-None-08880"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-08880  
  If the
  [attachmentFeedbackLoopDynamicState](#features-attachmentFeedbackLoopDynamicState)
  feature is enabled on the device, and a shader object is bound to the
  `VK_SHADER_STAGE_FRAGMENT_BIT` stage, and the most recent call to
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  in the current command buffer set `rasterizerDiscardEnable` to
  `VK_FALSE`, then
  [vkCmdSetAttachmentFeedbackLoopEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAttachmentFeedbackLoopEnableEXT.html)
  **must** have been called in the current command buffer prior to this
  drawing command

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pDynamicStates-08715"
  id="VUID-vkCmdDrawMeshTasksEXT-pDynamicStates-08715"></a>
  VUID-vkCmdDrawMeshTasksEXT-pDynamicStates-08715  
  If the bound graphics pipeline state includes a fragment shader stage,
  was created with `VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE` set in
  [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`,
  and the fragment shader declares the `EarlyFragmentTests` execution
  mode and uses `OpDepthAttachmentReadEXT`, the `depthWriteEnable`
  parameter in the last call to
  [vkCmdSetDepthWriteEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthWriteEnable.html) **must** be
  `VK_FALSE`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pDynamicStates-08716"
  id="VUID-vkCmdDrawMeshTasksEXT-pDynamicStates-08716"></a>
  VUID-vkCmdDrawMeshTasksEXT-pDynamicStates-08716  
  If the bound graphics pipeline state includes a fragment shader stage,
  was created with `VK_DYNAMIC_STATE_STENCIL_WRITE_MASK` set in
  [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`,
  and the fragment shader declares the `EarlyFragmentTests` execution
  mode and uses `OpStencilAttachmentReadEXT`, the `writeMask` parameter
  in the last call to
  [vkCmdSetStencilWriteMask](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilWriteMask.html) **must** be
  `0`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09116"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09116"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09116  
  If a shader object is bound to any graphics stage or the currently
  bound graphics pipeline was created with
  `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT`, and the format of any color
  attachment is `VK_FORMAT_E5B9G9R9_UFLOAT_PACK32`, the corresponding
  element of the `pColorWriteMasks` parameter of
  [vkCmdSetColorWriteMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorWriteMaskEXT.html) **must**
  either include all of `VK_COLOR_COMPONENT_R_BIT`,
  `VK_COLOR_COMPONENT_G_BIT`, and `VK_COLOR_COMPONENT_B_BIT`, or none of
  them

- <a
  href="#VUID-vkCmdDrawMeshTasksEXT-maxFragmentDualSrcAttachments-09239"
  id="VUID-vkCmdDrawMeshTasksEXT-maxFragmentDualSrcAttachments-09239"></a>
  VUID-vkCmdDrawMeshTasksEXT-maxFragmentDualSrcAttachments-09239  
  If [blending](#framebuffer-blending) is enabled for any attachment
  where either the source or destination blend factors for that
  attachment [use the secondary color input](#framebuffer-dsb), the
  maximum value of `Location` for any output attachment [statically
  used](#shaders-staticuse) in the `Fragment` `Execution` `Model`
  executed by this command **must** be less than
  [`maxFragmentDualSrcAttachments`](#limits-maxFragmentDualSrcAttachments)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09548"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09548"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09548  
  If the current render pass was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), and there is no
  shader object bound to any graphics stage, the value of each element
  of
  [VkRenderingAttachmentLocationInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentLocationInfoKHR.html)::`pColorAttachmentLocations`
  set by
  [vkCmdSetRenderingAttachmentLocationsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRenderingAttachmentLocationsKHR.html)
  **must** match the value set for the corresponding element in the
  currently bound pipeline

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09549"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09549"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09549  
  If the current render pass was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), and there is no
  shader object bound to any graphics stage, input attachment index
  mappings in the currently bound pipeline **must** match those set for
  the current render pass instance via
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09642"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09642"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09642  
  If the current render pass was begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) with the
  `VK_RENDERING_ENABLE_LEGACY_DITHERING_BIT_EXT` flag, the bound
  graphics pipeline **must** have been created with
  `VK_PIPELINE_CREATE_2_ENABLE_LEGACY_DITHERING_BIT_EXT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-09643"
  id="VUID-vkCmdDrawMeshTasksEXT-None-09643"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-09643  
  If the bound graphics pipeline was created with
  `VK_PIPELINE_CREATE_2_ENABLE_LEGACY_DITHERING_BIT_EXT`, the current
  render pass **must** have begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) with the
  `VK_RENDERING_ENABLE_LEGACY_DITHERING_BIT_EXT` flag

<!-- -->

- <a href="#VUID-vkCmdDrawMeshTasksEXT-stage-06480"
  id="VUID-vkCmdDrawMeshTasksEXT-stage-06480"></a>
  VUID-vkCmdDrawMeshTasksEXT-stage-06480  
  The bound graphics pipeline **must** not have been created with the
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)::`stage`
  member of an element of
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`pStages`
  set to `VK_SHADER_STAGE_VERTEX_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT` or
  `VK_SHADER_STAGE_GEOMETRY_BIT`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07074"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07074"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07074  
  [Transform Feedback Queries](#queries-transform-feedback) **must** not
  be active

- <a href="#VUID-vkCmdDrawMeshTasksEXT-None-07075"
  id="VUID-vkCmdDrawMeshTasksEXT-None-07075"></a>
  VUID-vkCmdDrawMeshTasksEXT-None-07075  
  [Primitives Generated Queries](#queries-primitives-generated) **must**
  not be active

- <a href="#VUID-vkCmdDrawMeshTasksEXT-pipelineStatistics-07076"
  id="VUID-vkCmdDrawMeshTasksEXT-pipelineStatistics-07076"></a>
  VUID-vkCmdDrawMeshTasksEXT-pipelineStatistics-07076  
  The `pipelineStatistics` member used to create any active [Pipeline
  Statistics Query](#queries-pipestats) **must** not contain
  `VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_VERTICES_BIT`,
  `VK_QUERY_PIPELINE_STATISTIC_INPUT_ASSEMBLY_PRIMITIVES_BIT`,
  `VK_QUERY_PIPELINE_STATISTIC_VERTEX_SHADER_INVOCATIONS_BIT`,
  `VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_INVOCATIONS_BIT`,
  `VK_QUERY_PIPELINE_STATISTIC_GEOMETRY_SHADER_PRIMITIVES_BIT`,
  `VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT`,
  `VK_QUERY_PIPELINE_STATISTIC_CLIPPING_PRIMITIVES_BIT`,
  `VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_CONTROL_SHADER_PATCHES_BIT`,
  or
  `VK_QUERY_PIPELINE_STATISTIC_TESSELLATION_EVALUATION_SHADER_INVOCATIONS_BIT`

<!-- -->

- <a href="#VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07322"
  id="VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07322"></a>
  VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07322  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  contains a shader using the `TaskEXT` `Execution` `Model`,
  `groupCountX` **must** be less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxTaskWorkGroupCount`\[0\]

- <a href="#VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07323"
  id="VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07323"></a>
  VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07323  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  contains a shader using the `TaskEXT` `Execution` `Model`,
  `groupCountY` **must** be less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxTaskWorkGroupCount`\[1\]

- <a href="#VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07324"
  id="VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07324"></a>
  VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07324  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  contains a shader using the `TaskEXT` `Execution` `Model`,
  `groupCountZ` **must** be less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxTaskWorkGroupCount`\[2\]

- <a href="#VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07325"
  id="VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07325"></a>
  VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07325  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  contains a shader using the `TaskEXT` `Execution` `Model`, The product
  of `groupCountX`, `groupCountY` and `groupCountZ` **must** be less
  than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxTaskWorkGroupTotalCount`

- <a href="#VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07326"
  id="VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07326"></a>
  VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07326  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  does not contain a shader using the `TaskEXT` `Execution` `Model`,
  `groupCountX` **must** be less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxMeshWorkGroupCount`\[0\]

- <a href="#VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07327"
  id="VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07327"></a>
  VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07327  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  does not contain a shader using the `TaskEXT` `Execution` `Model`,
  `groupCountY` **must** be less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxMeshWorkGroupCount`\[1\]

- <a href="#VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07328"
  id="VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07328"></a>
  VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07328  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  does not contain a shader using the `TaskEXT` `Execution` `Model`,
  `groupCountZ` **must** be less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxMeshWorkGroupCount`\[2\]

- <a href="#VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07329"
  id="VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07329"></a>
  VUID-vkCmdDrawMeshTasksEXT-TaskEXT-07329  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  does not contain a shader using the `TaskEXT` `Execution` `Model`, The
  product of `groupCountX`, `groupCountY` and `groupCountZ` **must** be
  less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxMeshWorkGroupTotalCount`

<!-- -->

- <a href="#VUID-vkCmdDrawMeshTasksEXT-MeshEXT-07087"
  id="VUID-vkCmdDrawMeshTasksEXT-MeshEXT-07087"></a>
  VUID-vkCmdDrawMeshTasksEXT-MeshEXT-07087  
  The current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  **must** contain a shader stage using the `MeshEXT` `Execution`
  `Model`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-commandBuffer-parameter"
  id="VUID-vkCmdDrawMeshTasksEXT-commandBuffer-parameter"></a>
  VUID-vkCmdDrawMeshTasksEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdDrawMeshTasksEXT-commandBuffer-recording"
  id="VUID-vkCmdDrawMeshTasksEXT-commandBuffer-recording"></a>
  VUID-vkCmdDrawMeshTasksEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdDrawMeshTasksEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdDrawMeshTasksEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdDrawMeshTasksEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdDrawMeshTasksEXT-renderpass"
  id="VUID-vkCmdDrawMeshTasksEXT-renderpass"></a>
  VUID-vkCmdDrawMeshTasksEXT-renderpass  
  This command **must** only be called inside of a render pass instance

- <a href="#VUID-vkCmdDrawMeshTasksEXT-videocoding"
  id="VUID-vkCmdDrawMeshTasksEXT-videocoding"></a>
  VUID-vkCmdDrawMeshTasksEXT-videocoding  
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

[VK_EXT_mesh_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_mesh_shader.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdDrawMeshTasksEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
