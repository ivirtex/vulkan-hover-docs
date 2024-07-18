# VkPipelineLayoutCreateInfo(3) Manual Page

## Name

VkPipelineLayoutCreateInfo - Structure specifying the parameters of a
newly created pipeline layout object



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)
structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkPipelineLayoutCreateInfo {
    VkStructureType                 sType;
    const void*                     pNext;
    VkPipelineLayoutCreateFlags     flags;
    uint32_t                        setLayoutCount;
    const VkDescriptorSetLayout*    pSetLayouts;
    uint32_t                        pushConstantRangeCount;
    const VkPushConstantRange*      pPushConstantRanges;
} VkPipelineLayoutCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkPipelineLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateFlagBits.html)
  specifying options for pipeline layout creation.

- `setLayoutCount` is the number of descriptor sets included in the
  pipeline layout.

- `pSetLayouts` is a pointer to an array of `VkDescriptorSetLayout`
  objects.

- `pushConstantRangeCount` is the number of push constant ranges
  included in the pipeline layout.

- `pPushConstantRanges` is a pointer to an array of
  `VkPushConstantRange` structures defining a set of push constant
  ranges for use in a single pipeline layout. In addition to descriptor
  set layouts, a pipeline layout also describes how many push constants
  **can** be accessed by each stage of the pipeline.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>Push constants represent a high speed path to modify constant data in
  pipelines that is expected to outperform memory-backed resource
  updates.</p></td>
  </tr>
  </tbody>
  </table>

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPipelineLayoutCreateInfo-setLayoutCount-00286"
  id="VUID-VkPipelineLayoutCreateInfo-setLayoutCount-00286"></a>
  VUID-VkPipelineLayoutCreateInfo-setLayoutCount-00286  
  `setLayoutCount` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxBoundDescriptorSets`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03016"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03016"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03016  
  The total number of descriptors in descriptor set layouts created
  without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_SAMPLER` and
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` accessible to any given
  shader stage across all elements of `pSetLayouts` **must** be less
  than or equal to
  `VkPhysicalDeviceLimits`::`maxPerStageDescriptorSamplers`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03017"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03017"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03017  
  The total number of descriptors in descriptor set layouts created
  without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` and
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` accessible to any given
  shader stage across all elements of `pSetLayouts` **must** be less
  than or equal to
  `VkPhysicalDeviceLimits`::`maxPerStageDescriptorUniformBuffers`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03018"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03018"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03018  
  The total number of descriptors in descriptor set layouts created
  without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` and
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` accessible to any given
  shader stage across all elements of `pSetLayouts` **must** be less
  than or equal to
  `VkPhysicalDeviceLimits`::`maxPerStageDescriptorStorageBuffers`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-06939"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-06939"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-06939  
  The total number of descriptors in descriptor set layouts created
  without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`,
  `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM`,
  `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`, and
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`, accessible to any given
  shader stage across all elements of `pSetLayouts` **must** be less
  than or equal to
  `VkPhysicalDeviceLimits`::`maxPerStageDescriptorSampledImages`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03020"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03020"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03020  
  The total number of descriptors in descriptor set layouts created
  without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, and
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` accessible to any given
  shader stage across all elements of `pSetLayouts` **must** be less
  than or equal to
  `VkPhysicalDeviceLimits`::`maxPerStageDescriptorStorageImages`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03021"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03021"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03021  
  The total number of descriptors in descriptor set layouts created
  without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`
  accessible to any given shader stage across all elements of
  `pSetLayouts` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxPerStageDescriptorInputAttachments`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-02214"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-02214"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-02214  
  The total number of bindings in descriptor set layouts created without
  the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit
  set and with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` accessible to any given
  shader stage across all elements of `pSetLayouts`, **must** be less
  than or equal to
  `VkPhysicalDeviceInlineUniformBlockProperties`::`maxPerStageDescriptorInlineUniformBlocks`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03022"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03022"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03022  
  The total number of descriptors with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_SAMPLER` and
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` accessible to any given
  shader stage across all elements of `pSetLayouts` **must** be less
  than or equal to
  `VkPhysicalDeviceDescriptorIndexingProperties`::`maxPerStageDescriptorUpdateAfterBindSamplers`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03023"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03023"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03023  
  The total number of descriptors with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` and
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` accessible to any given
  shader stage across all elements of `pSetLayouts` **must** be less
  than or equal to
  `VkPhysicalDeviceDescriptorIndexingProperties`::`maxPerStageDescriptorUpdateAfterBindUniformBuffers`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03024"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03024"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03024  
  The total number of descriptors with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` and
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` accessible to any given
  shader stage across all elements of `pSetLayouts` **must** be less
  than or equal to
  `VkPhysicalDeviceDescriptorIndexingProperties`::`maxPerStageDescriptorUpdateAfterBindStorageBuffers`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03025"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03025"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03025  
  The total number of descriptors with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, and
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` accessible to any given
  shader stage across all elements of `pSetLayouts` **must** be less
  than or equal to
  `VkPhysicalDeviceDescriptorIndexingProperties`::`maxPerStageDescriptorUpdateAfterBindSampledImages`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03026"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03026"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03026  
  The total number of descriptors with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, and
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` accessible to any given
  shader stage across all elements of `pSetLayouts` **must** be less
  than or equal to
  `VkPhysicalDeviceDescriptorIndexingProperties`::`maxPerStageDescriptorUpdateAfterBindStorageImages`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03027"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03027"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03027  
  The total number of descriptors with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` accessible to any given shader
  stage across all elements of `pSetLayouts` **must** be less than or
  equal to
  `VkPhysicalDeviceDescriptorIndexingProperties`::`maxPerStageDescriptorUpdateAfterBindInputAttachments`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-02215"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-02215"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-02215  
  The total number of bindings with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` accessible to any given
  shader stage across all elements of `pSetLayouts` **must** be less
  than or equal to
  `VkPhysicalDeviceInlineUniformBlockProperties`::`maxPerStageDescriptorUpdateAfterBindInlineUniformBlocks`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03028"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03028"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03028  
  The total number of descriptors in descriptor set layouts created
  without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_SAMPLER` and
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` accessible across all
  shader stages and across all elements of `pSetLayouts` **must** be
  less than or equal to
  `VkPhysicalDeviceLimits`::`maxDescriptorSetSamplers`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03029"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03029"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03029  
  The total number of descriptors in descriptor set layouts created
  without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`
  accessible across all shader stages and across all elements of
  `pSetLayouts` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxDescriptorSetUniformBuffers`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03030"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03030"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03030  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
  target="_blank" rel="noopener"><code>maintenance7</code></a> feature
  is not enabled, the total number of descriptors in descriptor set
  layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`
  accessible across all shader stages and across all elements of
  `pSetLayouts` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxDescriptorSetUniformBuffersDynamic`

- <a href="#VUID-VkPipelineLayoutCreateInfo-maintenance7-10003"
  id="VUID-VkPipelineLayoutCreateInfo-maintenance7-10003"></a>
  VUID-VkPipelineLayoutCreateInfo-maintenance7-10003  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
  target="_blank" rel="noopener"><code>maintenance7</code></a> feature
  is enabled, the total number of descriptors in descriptor set layouts
  created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`
  accessible across all shader stages and across all elements of
  `pSetLayouts` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxDescriptorSetTotalUniformBuffersDynamic"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceMaintenance7PropertiesKHR</code>::<code>maxDescriptorSetTotalUniformBuffersDynamic</code></a>

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03031"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03031"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03031  
  The total number of descriptors in descriptor set layouts created
  without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`
  accessible across all shader stages and across all elements of
  `pSetLayouts` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxDescriptorSetStorageBuffers`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03032"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03032"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03032  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
  target="_blank" rel="noopener"><code>maintenance7</code></a> feature
  is not enabled, the total number of descriptors in descriptor set
  layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`
  accessible across all shader stages and across all elements of
  `pSetLayouts` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxDescriptorSetStorageBuffersDynamic`

- <a href="#VUID-VkPipelineLayoutCreateInfo-maintenance7-10004"
  id="VUID-VkPipelineLayoutCreateInfo-maintenance7-10004"></a>
  VUID-VkPipelineLayoutCreateInfo-maintenance7-10004  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
  target="_blank" rel="noopener"><code>maintenance7</code></a> feature
  is enabled, the total number of descriptors in descriptor set layouts
  created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`
  accessible across all shader stages and across all elements of
  `pSetLayouts` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxDescriptorSetTotalStorageBuffersDynamic"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceMaintenance7PropertiesKHR</code>::<code>maxDescriptorSetTotalStorageBuffersDynamic</code></a>

- <a href="#VUID-VkPipelineLayoutCreateInfo-None-10005"
  id="VUID-VkPipelineLayoutCreateInfo-None-10005"></a>
  VUID-VkPipelineLayoutCreateInfo-None-10005  
  The total number of descriptors in descriptor set layouts created
  without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`
  or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` accessible across all
  shader stages and across all elements of `pSetLayouts` **must** be
  less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxDescriptorSetTotalBuffersDynamic"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceMaintenance7PropertiesKHR</code>::<code>maxDescriptorSetTotalBuffersDynamic</code></a>

- <a href="#VUID-VkPipelineLayoutCreateInfo-pSetLayouts-10006"
  id="VUID-VkPipelineLayoutCreateInfo-pSetLayouts-10006"></a>
  VUID-VkPipelineLayoutCreateInfo-pSetLayouts-10006  
  The total number of descriptors of the type
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` accessible across all
  shader stages and across all elements of `pSetLayouts` **must** be
  less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxDescriptorSetUpdateAfterBindTotalBuffersDynamic"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceMaintenance7PropertiesKHR</code>::<code>maxDescriptorSetUpdateAfterBindTotalBuffersDynamic</code></a>

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03033"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03033"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03033  
  The total number of descriptors in descriptor set layouts created
  without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, and
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` accessible across all shader
  stages and across all elements of `pSetLayouts` **must** be less than
  or equal to `VkPhysicalDeviceLimits`::`maxDescriptorSetSampledImages`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03034"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03034"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03034  
  The total number of descriptors in descriptor set layouts created
  without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, and
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` accessible across all shader
  stages and across all elements of `pSetLayouts` **must** be less than
  or equal to `VkPhysicalDeviceLimits`::`maxDescriptorSetStorageImages`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03035"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03035"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03035  
  The total number of descriptors in descriptor set layouts created
  without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  with a `descriptorType` of `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`
  accessible across all shader stages and across all elements of
  `pSetLayouts` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxDescriptorSetInputAttachments`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-02216"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-02216"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-02216  
  The total number of bindings in descriptor set layouts created without
  the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit
  set with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` accessible across all shader
  stages and across all elements of `pSetLayouts` **must** be less than
  or equal to
  `VkPhysicalDeviceInlineUniformBlockProperties`::`maxDescriptorSetInlineUniformBlocks`

- <a href="#VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03036"
  id="VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03036"></a>
  VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03036  
  The total number of descriptors of the type
  `VK_DESCRIPTOR_TYPE_SAMPLER` and
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` accessible across all
  shader stages and across all elements of `pSetLayouts` **must** be
  less than or equal to
  `VkPhysicalDeviceDescriptorIndexingProperties`::`maxDescriptorSetUpdateAfterBindSamplers`

- <a href="#VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03037"
  id="VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03037"></a>
  VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03037  
  The total number of descriptors of the type
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` accessible across all shader
  stages and across all elements of `pSetLayouts` **must** be less than
  or equal to
  `VkPhysicalDeviceDescriptorIndexingProperties`::`maxDescriptorSetUpdateAfterBindUniformBuffers`

- <a href="#VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03038"
  id="VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03038"></a>
  VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03038  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
  target="_blank" rel="noopener"><code>maintenance7</code></a> feature
  is not enabled, the total number of descriptors of the type
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` accessible across all
  shader stages and across all elements of `pSetLayouts` **must** be
  less than or equal to
  `VkPhysicalDeviceLimits`::`maxDescriptorSetUpdateAfterBindUniformBuffersDynamic`

- <a href="#VUID-VkPipelineLayoutCreateInfo-maintenance7-10007"
  id="VUID-VkPipelineLayoutCreateInfo-maintenance7-10007"></a>
  VUID-VkPipelineLayoutCreateInfo-maintenance7-10007  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
  target="_blank" rel="noopener"><code>maintenance7</code></a> feature
  is enabled, the total number of descriptors of the type
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` accessible across all
  shader stages and across all elements of `pSetLayouts` **must** be
  less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxDescriptorSetUpdateAfterBindTotalUniformBuffersDynamic"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceMaintenance7PropertiesKHR</code>::<code>maxDescriptorSetUpdateAfterBindTotalUniformBuffersDynamic</code></a>

- <a href="#VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03039"
  id="VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03039"></a>
  VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03039  
  The total number of descriptors of the type
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` accessible across all shader
  stages and across all elements of `pSetLayouts` **must** be less than
  or equal to
  `VkPhysicalDeviceDescriptorIndexingProperties`::`maxDescriptorSetUpdateAfterBindStorageBuffers`

- <a href="#VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03040"
  id="VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03040"></a>
  VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03040  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
  target="_blank" rel="noopener"><code>maintenance7</code></a> feature
  is not enabled, the total number of descriptors of the type
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` accessible across all
  shader stages and across all elements of `pSetLayouts` **must** be
  less than or equal to
  `VkPhysicalDeviceLimits`::`maxDescriptorSetUpdateAfterBindStorageBuffersDynamic`

- <a href="#VUID-VkPipelineLayoutCreateInfo-maintenance7-10008"
  id="VUID-VkPipelineLayoutCreateInfo-maintenance7-10008"></a>
  VUID-VkPipelineLayoutCreateInfo-maintenance7-10008  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
  target="_blank" rel="noopener"><code>maintenance7</code></a> feature
  is enabled, the total number of descriptors of the type
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` accessible across all
  shader stages and across all elements of `pSetLayouts` **must** be
  less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxDescriptorSetUpdateAfterBindTotalStorageBuffersDynamic"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceMaintenance7PropertiesKHR</code>::<code>maxDescriptorSetUpdateAfterBindTotalStorageBuffersDynamic</code></a>

- <a href="#VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03041"
  id="VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03041"></a>
  VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03041  
  The total number of descriptors of the type
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, and
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` accessible across all shader
  stages and across all elements of `pSetLayouts` **must** be less than
  or equal to
  `VkPhysicalDeviceDescriptorIndexingProperties`::`maxDescriptorSetUpdateAfterBindSampledImages`

- <a href="#VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03042"
  id="VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03042"></a>
  VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03042  
  The total number of descriptors of the type
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, and
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` accessible across all shader
  stages and across all elements of `pSetLayouts` **must** be less than
  or equal to
  `VkPhysicalDeviceDescriptorIndexingProperties`::`maxDescriptorSetUpdateAfterBindStorageImages`

- <a href="#VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03043"
  id="VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03043"></a>
  VUID-VkPipelineLayoutCreateInfo-pSetLayouts-03043  
  The total number of descriptors of the type
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` accessible across all shader
  stages and across all elements of `pSetLayouts` **must** be less than
  or equal to
  `VkPhysicalDeviceDescriptorIndexingProperties`::`maxDescriptorSetUpdateAfterBindInputAttachments`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-02217"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-02217"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-02217  
  The total number of bindings with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` accessible across all shader
  stages and across all elements of `pSetLayouts` **must** be less than
  or equal to
  `VkPhysicalDeviceInlineUniformBlockProperties`::`maxDescriptorSetUpdateAfterBindInlineUniformBlocks`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-06531"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-06531"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-06531  
  The total number of descriptors with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` accessible across all shader
  stages and across all elements of `pSetLayouts` **must** be less than
  or equal to
  `VkPhysicalDeviceVulkan13Properties`::`maxInlineUniformTotalSize`

- <a href="#VUID-VkPipelineLayoutCreateInfo-pPushConstantRanges-00292"
  id="VUID-VkPipelineLayoutCreateInfo-pPushConstantRanges-00292"></a>
  VUID-VkPipelineLayoutCreateInfo-pPushConstantRanges-00292  
  Any two elements of `pPushConstantRanges` **must** not include the
  same stage in `stageFlags`

- <a href="#VUID-VkPipelineLayoutCreateInfo-pSetLayouts-00293"
  id="VUID-VkPipelineLayoutCreateInfo-pSetLayouts-00293"></a>
  VUID-VkPipelineLayoutCreateInfo-pSetLayouts-00293  
  `pSetLayouts` **must** not contain more than one descriptor set layout
  that was created with
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR` set

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03571"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03571"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03571  
  The total number of bindings in descriptor set layouts created without
  the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit
  set with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR` accessible to any
  given shader stage across all elements of `pSetLayouts` **must** be
  less than or equal to
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxPerStageDescriptorAccelerationStructures`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03572"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03572"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03572  
  The total number of bindings with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR` accessible to any
  given shader stage across all elements of `pSetLayouts` **must** be
  less than or equal to
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxPerStageDescriptorUpdateAfterBindAccelerationStructures`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03573"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03573"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03573  
  The total number of bindings in descriptor set layouts created without
  the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit
  set with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR` accessible across all
  shader stages and across all elements of `pSetLayouts` **must** be
  less than or equal to
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxDescriptorSetAccelerationStructures`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-03574"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-03574"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-03574  
  The total number of bindings with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR` accessible across all
  shader stages and across all elements of `pSetLayouts` **must** be
  less than or equal to
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`maxDescriptorSetUpdateAfterBindAccelerationStructures`

- <a href="#VUID-VkPipelineLayoutCreateInfo-descriptorType-02381"
  id="VUID-VkPipelineLayoutCreateInfo-descriptorType-02381"></a>
  VUID-VkPipelineLayoutCreateInfo-descriptorType-02381  
  The total number of bindings with a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV` accessible across all
  shader stages and across all elements of `pSetLayouts` **must** be
  less than or equal to
  [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html)::`maxDescriptorSetAccelerationStructures`

- <a href="#VUID-VkPipelineLayoutCreateInfo-pImmutableSamplers-03566"
  id="VUID-VkPipelineLayoutCreateInfo-pImmutableSamplers-03566"></a>
  VUID-VkPipelineLayoutCreateInfo-pImmutableSamplers-03566  
  The total number of `pImmutableSamplers` created with `flags`
  containing `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT` or
  `VK_SAMPLER_CREATE_SUBSAMPLED_COARSE_RECONSTRUCTION_BIT_EXT` across
  all shader stages and across all elements of `pSetLayouts` **must** be
  less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxDescriptorSetSubsampledSamplers"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceFragmentDensityMap2PropertiesEXT</code>::<code>maxDescriptorSetSubsampledSamplers</code></a>

- <a href="#VUID-VkPipelineLayoutCreateInfo-pSetLayouts-04606"
  id="VUID-VkPipelineLayoutCreateInfo-pSetLayouts-04606"></a>
  VUID-VkPipelineLayoutCreateInfo-pSetLayouts-04606  
  Any element of `pSetLayouts` **must** not have been created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT` bit set

- <a href="#VUID-VkPipelineLayoutCreateInfo-graphicsPipelineLibrary-06753"
  id="VUID-VkPipelineLayoutCreateInfo-graphicsPipelineLibrary-06753"></a>
  VUID-VkPipelineLayoutCreateInfo-graphicsPipelineLibrary-06753  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-graphicsPipelineLibrary"
  target="_blank" rel="noopener"><code>graphicsPipelineLibrary</code></a>
  is not enabled, elements of `pSetLayouts` **must** be valid
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) objects

- <a href="#VUID-VkPipelineLayoutCreateInfo-pSetLayouts-08008"
  id="VUID-VkPipelineLayoutCreateInfo-pSetLayouts-08008"></a>
  VUID-VkPipelineLayoutCreateInfo-pSetLayouts-08008  
  If any element of `pSetLayouts` was created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT` bit set,
  all elements of `pSetLayouts` **must** have been created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT` bit set

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineLayoutCreateInfo-sType-sType"
  id="VUID-VkPipelineLayoutCreateInfo-sType-sType"></a>
  VUID-VkPipelineLayoutCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_LAYOUT_CREATE_INFO`

- <a href="#VUID-VkPipelineLayoutCreateInfo-flags-parameter"
  id="VUID-VkPipelineLayoutCreateInfo-flags-parameter"></a>
  VUID-VkPipelineLayoutCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkPipelineLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateFlagBits.html)
  values

- <a href="#VUID-VkPipelineLayoutCreateInfo-pSetLayouts-parameter"
  id="VUID-VkPipelineLayoutCreateInfo-pSetLayouts-parameter"></a>
  VUID-VkPipelineLayoutCreateInfo-pSetLayouts-parameter  
  If `setLayoutCount` is not `0`, `pSetLayouts` **must** be a valid
  pointer to an array of `setLayoutCount` valid or
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) handles

- <a href="#VUID-VkPipelineLayoutCreateInfo-pPushConstantRanges-parameter"
  id="VUID-VkPipelineLayoutCreateInfo-pPushConstantRanges-parameter"></a>
  VUID-VkPipelineLayoutCreateInfo-pPushConstantRanges-parameter  
  If `pushConstantRangeCount` is not `0`, `pPushConstantRanges` **must**
  be a valid pointer to an array of `pushConstantRangeCount` valid
  [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantRange.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html),
[VkPipelineLayoutCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateFlags.html),
[VkPushConstantRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantRange.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreatePipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreatePipelineLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineLayoutCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
