# VK\_VERSION\_1\_2(3) Manual Page

## Name

VK\_VERSION\_1\_2 - Vulkan version 1.2



## [](#_description)Description

Vulkan Version 1.2 [promoted](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-promotion) a number of key extensions into the core API:

- [VK\_KHR\_8bit\_storage](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_8bit_storage.html)
- [VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html)
- [VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html)
- [VK\_KHR\_depth\_stencil\_resolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_depth_stencil_resolve.html)
- [VK\_KHR\_draw\_indirect\_count](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_draw_indirect_count.html)
- [VK\_KHR\_driver\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_driver_properties.html)
- [VK\_KHR\_image\_format\_list](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_image_format_list.html)
- [VK\_KHR\_imageless\_framebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_imageless_framebuffer.html)
- [VK\_KHR\_sampler\_mirror\_clamp\_to\_edge](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_sampler_mirror_clamp_to_edge.html)
- [VK\_KHR\_separate\_depth\_stencil\_layouts](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_separate_depth_stencil_layouts.html)
- [VK\_KHR\_shader\_atomic\_int64](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_atomic_int64.html)
- [VK\_KHR\_shader\_float16\_int8](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_float16_int8.html)
- [VK\_KHR\_shader\_float\_controls](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_float_controls.html)
- [VK\_KHR\_shader\_subgroup\_extended\_types](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_subgroup_extended_types.html)
- [VK\_KHR\_spirv\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_spirv_1_4.html)
- [VK\_KHR\_timeline\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_timeline_semaphore.html)
- [VK\_KHR\_uniform\_buffer\_standard\_layout](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_uniform_buffer_standard_layout.html)
- [VK\_KHR\_vulkan\_memory\_model](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_vulkan_memory_model.html)
- [VK\_EXT\_descriptor\_indexing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_indexing.html)
- [VK\_EXT\_host\_query\_reset](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_query_reset.html)
- [VK\_EXT\_sampler\_filter\_minmax](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sampler_filter_minmax.html)
- [VK\_EXT\_scalar\_block\_layout](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_scalar_block_layout.html)
- [VK\_EXT\_separate\_stencil\_usage](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_separate_stencil_usage.html)
- [VK\_EXT\_shader\_viewport\_index\_layer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_viewport_index_layer.html)

All differences in behavior between these extensions and the corresponding Vulkan 1.2 functionality are summarized below.

Differences Relative to `VK_KHR_8bit_storage`

If the `VK_KHR_8bit_storage` extension is not supported, support for the SPIR-V [`storageBuffer8BitAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-storageBuffer8BitAccess) capability in shader modules is optional. Support for this feature is defined by [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Features.html)::`storageBuffer8BitAccess` when queried via [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html).

Differences Relative to `VK_KHR_draw_indirect_count`

If the `VK_KHR_draw_indirect_count` extension is not supported, support for the commands [vkCmdDrawIndirectCount](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndirectCount.html) and [vkCmdDrawIndexedIndirectCount](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexedIndirectCount.html) is optional. Support for this feature is defined by [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Features.html)::`drawIndirectCount` when queried via [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html).

Differences Relative to `VK_KHR_sampler_mirror_clamp_to_edge`

If the `VK_KHR_sampler_mirror_clamp_to_edge` extension is not supported, support for the [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerAddressMode.html) `VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE` is optional. Support for this feature is defined by [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Features.html)::`samplerMirrorClampToEdge` when queried via [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html).

Differences Relative to `VK_EXT_descriptor_indexing`

If the `VK_EXT_descriptor_indexing` extension is not supported, support for the [`descriptorIndexing`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorIndexing) feature is optional. Support for this feature is defined by [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Features.html)::`descriptorIndexing` when queried via [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html).

Differences Relative to `VK_EXT_scalar_block_layout`

If the `VK_EXT_scalar_block_layout` extension is not supported, support for the [`scalarBlockLayout`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-scalarBlockLayout) feature is optional. Support for this feature is defined by [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Features.html)::`scalarBlockLayout` when queried via [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html).

Differences Relative to `VK_EXT_shader_viewport_index_layer`

The `ShaderViewportIndexLayerEXT` SPIR-V capability was replaced with the `ShaderViewportIndex` and `ShaderLayer` capabilities. Declaring both is equivalent to declaring `ShaderViewportIndexLayerEXT`. If the `VK_EXT_shader_viewport_index_layer` extension is not supported, support for the `ShaderViewportIndexLayerEXT` SPIR-V capability is optional. Support for this feature is defined by [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Features.html)::`shaderOutputViewportIndex` and [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Features.html)::`shaderOutputLayer` when queried via [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html).

Differences Relative to `VK_KHR_buffer_device_address`

If the `VK_KHR_buffer_device_address` extension is not supported, support for the [`bufferDeviceAddress`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddress) feature is optional. Support for this feature is defined by [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Features.html)::`bufferDeviceAddress` when queried via [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html).

Differences Relative to `VK_KHR_shader_atomic_int64`

If the `VK_KHR_shader_atomic_int64` extension is not supported, support for the [`shaderBufferInt64Atomics`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderBufferInt64Atomics) feature is optional. Support for this feature is defined by [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Features.html)::`shaderBufferInt64Atomics` when queried via [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html).

Differences Relative to `VK_KHR_shader_float16_int8`

If the `VK_KHR_shader_float16_int8` extension is not supported, support for the [`shaderFloat16`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderFloat16) and [`shaderInt8`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderInt8) features is optional. Support for these features are defined by [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Features.html)::`shaderFloat16` and [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Features.html)::`shaderInt8` when queried via [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html).

Differences Relative to `VK_KHR_vulkan_memory_model`

If the `VK_KHR_vulkan_memory_model` extension is not supported, support for the [`vulkanMemoryModel`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-vulkanMemoryModel) feature is optional. Support for this feature is defined by [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Features.html)::`vulkanMemoryModel` when queried via [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html).

Additional Vulkan 1.2 Feature Support

In addition to the promoted extensions described above, Vulkan 1.2 added support for:

- SPIR-V version 1.4.
- SPIR-V version 1.5.
- The [`samplerMirrorClampToEdge`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-samplerMirrorClampToEdge) feature which indicates whether the implementation supports the `VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE` sampler address mode.
- The [`ShaderNonUniform`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ShaderNonUniform) capability in SPIR-V version 1.5.
- The [`shaderOutputViewportIndex`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderOutputViewportIndex) feature which indicates that the [`ShaderViewportIndex`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ShaderViewportIndex) capability can be used.
- The [`shaderOutputLayer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderOutputLayer) feature which indicates that the [`ShaderLayer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ShaderLayer) capability can be used.
- The [`subgroupBroadcastDynamicId`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-subgroupBroadcastDynamicId) feature which allows the “Id” operand of `OpGroupNonUniformBroadcast` to be dynamically uniform within a subgroup, and the “Index” operand of `OpGroupNonUniformQuadBroadcast` to be dynamically uniform within a derivative group, in shader modules of version 1.5 or higher.
- The [`drawIndirectCount`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-drawIndirectCount) feature which indicates whether the [vkCmdDrawIndirectCount](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndirectCount.html) and [vkCmdDrawIndexedIndirectCount](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexedIndirectCount.html) functions can be used.
- The [`descriptorIndexing`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorIndexing) feature which indicates the implementation supports the minimum number of descriptor indexing features as defined in the [Feature Requirements](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-requirements) section.
- The [`samplerFilterMinmax`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-samplerFilterMinmax) feature which indicates whether the implementation supports the minimum number of image formats that support the `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT` feature bit as defined by the [`filterMinmaxSingleComponentFormats`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-filterMinmaxSingleComponentFormats-minimum-requirements) property minimum requirements.
- The [`framebufferIntegerColorSampleCounts`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-framebufferIntegerColorSampleCounts) limit which indicates the color sample counts that are supported for all framebuffer color attachments with integer formats.

### [](#_new_macros)New Macros

- [VK\_API\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_API_VERSION_1_2.html)

### [](#_new_commands)New Commands

- [vkCmdBeginRenderPass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRenderPass2.html)
- [vkCmdDrawIndexedIndirectCount](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexedIndirectCount.html)
- [vkCmdDrawIndirectCount](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndirectCount.html)
- [vkCmdEndRenderPass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndRenderPass2.html)
- [vkCmdNextSubpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdNextSubpass2.html)
- [vkCreateRenderPass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRenderPass2.html)
- [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html)
- [vkGetBufferOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferOpaqueCaptureAddress.html)
- [vkGetDeviceMemoryOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html)
- [vkGetSemaphoreCounterValue](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSemaphoreCounterValue.html)
- [vkResetQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetQueryPool.html)
- [vkSignalSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSignalSemaphore.html)
- [vkWaitSemaphores](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitSemaphores.html)

### [](#_new_structures)New Structures

- [VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription2.html)
- [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html)
- [VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressInfo.html)
- [VkConformanceVersion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConformanceVersion.html)
- [VkDeviceMemoryOpaqueCaptureAddressInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryOpaqueCaptureAddressInfo.html)
- [VkFramebufferAttachmentImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentImageInfo.html)
- [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html)
- [VkSemaphoreSignalInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSignalInfo.html)
- [VkSemaphoreWaitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreWaitInfo.html)
- [VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassBeginInfo.html)
- [VkSubpassDependency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency2.html)
- [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html)
- [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassEndInfo.html)
- Extending [VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription2.html):
  
  - [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayout.html)
- Extending [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html):
  
  - [VkAttachmentReferenceStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReferenceStencilLayout.html)
- Extending [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html):
  
  - [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html)
- Extending [VkDescriptorSetAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetAllocateInfo.html):
  
  - [VkDescriptorSetVariableDescriptorCountAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfo.html)
- Extending [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html):
  
  - [VkDescriptorSetLayoutBindingFlagsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBindingFlagsCreateInfo.html)
- Extending [VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutSupport.html):
  
  - [VkDescriptorSetVariableDescriptorCountLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetVariableDescriptorCountLayoutSupport.html)
- Extending [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html):
  
  - [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html)
- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html):
  
  - [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html)
- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html), [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html):
  
  - [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkMemoryOpaqueCaptureAddressAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryOpaqueCaptureAddressAllocateInfo.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevice8BitStorageFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice8BitStorageFeatures.html)
  - [VkPhysicalDeviceBufferDeviceAddressFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBufferDeviceAddressFeatures.html)
  - [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)
  - [VkPhysicalDeviceHostQueryResetFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceHostQueryResetFeatures.html)
  - [VkPhysicalDeviceImagelessFramebufferFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImagelessFramebufferFeatures.html)
  - [VkPhysicalDeviceScalarBlockLayoutFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceScalarBlockLayoutFeatures.html)
  - [VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures.html)
  - [VkPhysicalDeviceShaderAtomicInt64Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderAtomicInt64Features.html)
  - [VkPhysicalDeviceShaderFloat16Int8Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderFloat16Int8Features.html)
  - [VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures.html)
  - [VkPhysicalDeviceTimelineSemaphoreFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTimelineSemaphoreFeatures.html)
  - [VkPhysicalDeviceUniformBufferStandardLayoutFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceUniformBufferStandardLayoutFeatures.html)
  - [VkPhysicalDeviceVulkan11Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan11Features.html)
  - [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Features.html)
  - [VkPhysicalDeviceVulkanMemoryModelFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkanMemoryModelFeatures.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)
  - [VkPhysicalDeviceDescriptorIndexingProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingProperties.html)
  - [VkPhysicalDeviceDriverProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDriverProperties.html)
  - [VkPhysicalDeviceFloatControlsProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFloatControlsProperties.html)
  - [VkPhysicalDeviceSamplerFilterMinmaxProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSamplerFilterMinmaxProperties.html)
  - [VkPhysicalDeviceTimelineSemaphoreProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTimelineSemaphoreProperties.html)
  - [VkPhysicalDeviceVulkan11Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan11Properties.html)
  - [VkPhysicalDeviceVulkan12Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan12Properties.html)
- Extending [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html):
  
  - [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html)
- Extending [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html):
  
  - [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeCreateInfo.html)
- Extending [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html), [VkPhysicalDeviceExternalSemaphoreInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html):
  
  - [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreTypeCreateInfo.html)
- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html), [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html):
  
  - [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimelineSemaphoreSubmitInfo.html)
- Extending [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html):
  
  - [VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionDepthStencilResolve.html)

### [](#_new_enums)New Enums

- [VkDescriptorBindingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBindingFlagBits.html)
- [VkDriverId](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDriverId.html)
- [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBits.html)
- [VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionMode.html)
- [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html)
- [VkSemaphoreWaitFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreWaitFlagBits.html)
- [VkShaderFloatControlsIndependence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderFloatControlsIndependence.html)

### [](#_new_bitmasks)New Bitmasks

- [VkDescriptorBindingFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBindingFlags.html)
- [VkResolveModeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlags.html)
- [VkSemaphoreWaitFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreWaitFlags.html)

### [](#_new_enum_constants)New Enum Constants

- `VK_MAX_DRIVER_INFO_SIZE`
- `VK_MAX_DRIVER_NAME_SIZE`
- Extending [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateFlagBits.html):
  
  - `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT`
- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html):
  
  - `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT`
- Extending [VkDescriptorPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateFlagBits.html):
  
  - `VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT`
- Extending [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateFlagBits.html):
  
  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT`
- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html):
  
  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT`
- Extending [VkFramebufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateFlagBits.html):
  
  - `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`
- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html):
  
  - `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`
  - `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`
  - `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`
  - `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- Extending [VkMemoryAllocateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateFlagBits.html):
  
  - `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT`
  - `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_FRAGMENTATION`
  - `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS`
- Extending [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerAddressMode.html):
  
  - `VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2`
  - `VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_STENCIL_LAYOUT`
  - `VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2`
  - `VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT`
  - `VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO`
  - `VK_STRUCTURE_TYPE_BUFFER_OPAQUE_CAPTURE_ADDRESS_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO`
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT`
  - `VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO`
  - `VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO`
  - `VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_MEMORY_OPAQUE_CAPTURE_ADDRESS_ALLOCATE_INFO`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_EXTENDED_TYPES_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_1_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_1_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_FEATURES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_PROPERTIES`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2`
  - `VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO`
  - `VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO`
  - `VK_STRUCTURE_TYPE_SEMAPHORE_WAIT_INFO`
  - `VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO`
  - `VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2`
  - `VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2`
  - `VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE`
  - `VK_STRUCTURE_TYPE_SUBPASS_END_INFO`
  - `VK_STRUCTURE_TYPE_TIMELINE_SEMAPHORE_SUBMIT_INFO`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0