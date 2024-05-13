# VK_VERSION_1_1(3) Manual Page

## Name

VK_VERSION_1_1 - Vulkan version 1.1



## <a href="#_description" class="anchor"></a>Description

Vulkan Version 1.1 <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-promotion"
target="_blank" rel="noopener">promoted</a> a number of key extensions
into the core API:

- [VK_KHR_16bit_storage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_16bit_storage.html)

- [VK_KHR_bind_memory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_bind_memory2.html)

- [VK_KHR_dedicated_allocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dedicated_allocation.html)

- [VK_KHR_descriptor_update_template](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_descriptor_update_template.html)

- [VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html)

- [VK_KHR_device_group_creation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group_creation.html)

- [VK_KHR_external_fence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence.html)

- [VK_KHR_external_fence_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence_capabilities.html)

- [VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html)

- [VK_KHR_external_memory_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_capabilities.html)

- [VK_KHR_external_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore.html)

- [VK_KHR_external_semaphore_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore_capabilities.html)

- [VK_KHR_get_memory_requirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_memory_requirements2.html)

- [VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)

- [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html)

- [VK_KHR_maintenance2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance2.html)

- [VK_KHR_maintenance3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance3.html)

- [VK_KHR_multiview](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_multiview.html)

- [VK_KHR_relaxed_block_layout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_relaxed_block_layout.html)

- [VK_KHR_sampler_ycbcr_conversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_sampler_ycbcr_conversion.html)

- [VK_KHR_shader_draw_parameters](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_draw_parameters.html)

- [VK_KHR_storage_buffer_storage_class](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_storage_buffer_storage_class.html)

- [VK_KHR_variable_pointers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_variable_pointers.html)

All differences in behavior between these extensions and the
corresponding Vulkan 1.1 functionality are summarized in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
target="_blank" rel="noopener">Vulkan 1.1 specification appendix</a>.

### <a href="#_new_macros" class="anchor"></a>New Macros

- [VK_API_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_API_VERSION_1_1.html)

### <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html)

- [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html)

### <a href="#_new_commands" class="anchor"></a>New Commands

- [vkBindBufferMemory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory2.html)

- [vkBindImageMemory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory2.html)

- [vkCmdDispatchBase](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchBase.html)

- [vkCmdSetDeviceMask](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDeviceMask.html)

- [vkCreateDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplate.html)

- [vkCreateSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSamplerYcbcrConversion.html)

- [vkDestroyDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorUpdateTemplate.html)

- [vkDestroySamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySamplerYcbcrConversion.html)

- [vkEnumerateInstanceVersion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceVersion.html)

- [vkEnumeratePhysicalDeviceGroups](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceGroups.html)

- [vkGetBufferMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements2.html)

- [vkGetDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupport.html)

- [vkGetDeviceGroupPeerMemoryFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupPeerMemoryFeatures.html)

- [vkGetDeviceQueue2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceQueue2.html)

- [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html)

- [vkGetImageSparseMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSparseMemoryRequirements2.html)

- [vkGetPhysicalDeviceExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalBufferProperties.html)

- [vkGetPhysicalDeviceExternalFenceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalFenceProperties.html)

- [vkGetPhysicalDeviceExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html)

- [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html)

- [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html)

- [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)

- [vkGetPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties2.html)

- [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html)

- [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html)

- [vkGetPhysicalDeviceSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2.html)

- [vkTrimCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkTrimCommandPool.html)

- [vkUpdateDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplate.html)

### <a href="#_new_structures" class="anchor"></a>New Structures

- [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html)

- [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html)

- [VkBufferMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryRequirementsInfo2.html)

- [VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutSupport.html)

- [VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateInfo.html)

- [VkDescriptorUpdateTemplateEntry](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateEntry.html)

- [VkDeviceQueueInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueInfo2.html)

- [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalBufferProperties.html)

- [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceProperties.html)

- [VkExternalMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryProperties.html)

- [VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreProperties.html)

- [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html)

- [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html)

- [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryRequirementsInfo2.html)

- [VkImageSparseMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSparseMemoryRequirementsInfo2.html)

- [VkInputAttachmentAspectReference](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInputAttachmentAspectReference.html)

- [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html)

- [VkPhysicalDeviceExternalBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalBufferInfo.html)

- [VkPhysicalDeviceExternalFenceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFenceInfo.html)

- [VkPhysicalDeviceExternalSemaphoreInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html)

- [VkPhysicalDeviceGroupProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGroupProperties.html)

- [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)

- [VkPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties2.html)

- [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)

- [VkPhysicalDeviceSparseImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSparseImageFormatInfo2.html)

- [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html)

- [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)

- [VkSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties2.html)

- [VkSparseImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements2.html)

- Extending [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html):

  - [VkBindBufferMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryDeviceGroupInfo.html)

- Extending [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html):

  - [VkBindImageMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html)

  - [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html)

- Extending [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindSparseInfo.html):

  - [VkDeviceGroupBindSparseInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupBindSparseInfo.html)

- Extending [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html):

  - [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryBufferCreateInfo.html)

- Extending [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html):

  - [VkDeviceGroupCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupCommandBufferBeginInfo.html)

- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkDeviceGroupDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupDeviceCreateInfo.html)

  - [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html)

- Extending [VkFenceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceCreateInfo.html):

  - [VkExportFenceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportFenceCreateInfo.html)

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html):

  - [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)

- Extending [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html):

  - [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatProperties.html)

  - [VkSamplerYcbcrConversionImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionImageFormatProperties.html)

- Extending
  [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryRequirementsInfo2.html):

  - [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePlaneMemoryRequirementsInfo.html)

- Extending [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html):

  - [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewUsageCreateInfo.html)

- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html):

  - [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html)

  - [VkMemoryAllocateFlagsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagsInfo.html)

  - [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)

- Extending [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html):

  - [VkMemoryDedicatedRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedRequirements.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevice16BitStorageFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice16BitStorageFeatures.html)

  - [VkPhysicalDeviceMultiviewFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewFeatures.html)

  - [VkPhysicalDeviceProtectedMemoryFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProtectedMemoryFeatures.html)

  - [VkPhysicalDeviceSamplerYcbcrConversionFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSamplerYcbcrConversionFeatures.html)

  - [VkPhysicalDeviceShaderDrawParameterFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderDrawParameterFeatures.html)

  - [VkPhysicalDeviceShaderDrawParametersFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderDrawParametersFeatures.html)

  - [VkPhysicalDeviceVariablePointerFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVariablePointerFeatures.html)

  - [VkPhysicalDeviceVariablePointersFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVariablePointersFeatures.html)

- Extending
  [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html):

  - [VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalImageFormatInfo.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceIDProperties.html)

  - [VkPhysicalDeviceMaintenance3Properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance3Properties.html)

  - [VkPhysicalDeviceMultiviewProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewProperties.html)

  - [VkPhysicalDevicePointClippingProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePointClippingProperties.html)

  - [VkPhysicalDeviceProtectedMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProtectedMemoryProperties.html)

  - [VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubgroupProperties.html)

- Extending
  [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html):

  - [VkPipelineTessellationDomainOriginStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationDomainOriginStateCreateInfo.html)

- Extending [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html),
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html):

  - [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)

- Extending [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html):

  - [VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html)

  - [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)

- Extending [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html),
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html):

  - [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfo.html)

- Extending [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html):

  - [VkExportSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportSemaphoreCreateInfo.html)

- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html):

  - [VkDeviceGroupSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupSubmitInfo.html)

  - [VkProtectedSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkProtectedSubmitInfo.html)

### <a href="#_new_enums" class="anchor"></a>New Enums

- [VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html)

- [VkDescriptorUpdateTemplateType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateType.html)

- [VkDeviceQueueCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateFlagBits.html)

- [VkExternalFenceFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceFeatureFlagBits.html)

- [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)

- [VkExternalMemoryFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagBits.html)

- [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)

- [VkExternalSemaphoreFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlagBits.html)

- [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)

- [VkFenceImportFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlagBits.html)

- [VkMemoryAllocateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBits.html)

- [VkPeerMemoryFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagBits.html)

- [VkPointClippingBehavior](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPointClippingBehavior.html)

- [VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html)

- [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRange.html)

- [VkSemaphoreImportFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlagBits.html)

- [VkSubgroupFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlagBits.html)

- [VkTessellationDomainOrigin](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTessellationDomainOrigin.html)

### <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkCommandPoolTrimFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolTrimFlags.html)

- [VkDescriptorUpdateTemplateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateFlags.html)

- [VkExternalFenceFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceFeatureFlags.html)

- [VkExternalFenceHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlags.html)

- [VkExternalMemoryFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlags.html)

- [VkExternalMemoryHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlags.html)

- [VkExternalSemaphoreFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlags.html)

- [VkExternalSemaphoreHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlags.html)

- [VkFenceImportFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlags.html)

- [VkMemoryAllocateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlags.html)

- [VkPeerMemoryFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlags.html)

- [VkSemaphoreImportFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlags.html)

- [VkSubgroupFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlags.html)

### <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_LUID_SIZE`

- `VK_MAX_DEVICE_GROUP_SIZE`

- `VK_QUEUE_FAMILY_EXTERNAL`

- Extending [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlagBits.html):

  - `VK_BUFFER_CREATE_PROTECTED_BIT`

- Extending
  [VkCommandPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateFlagBits.html):

  - `VK_COMMAND_POOL_CREATE_PROTECTED_BIT`

- Extending [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html):

  - `VK_DEPENDENCY_DEVICE_GROUP_BIT`

  - `VK_DEPENDENCY_VIEW_LOCAL_BIT`

- Extending
  [VkDeviceQueueCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateFlagBits.html):

  - `VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT`

- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html):

  - `VK_FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16`

  - `VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16`

  - `VK_FORMAT_B16G16R16G16_422_UNORM`

  - `VK_FORMAT_B8G8R8G8_422_UNORM`

  - `VK_FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16`

  - `VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16`

  - `VK_FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16`

  - `VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16`

  - `VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16`

  - `VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16`

  - `VK_FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16`

  - `VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16`

  - `VK_FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16`

  - `VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16`

  - `VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16`

  - `VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16`

  - `VK_FORMAT_G16B16G16R16_422_UNORM`

  - `VK_FORMAT_G16_B16R16_2PLANE_420_UNORM`

  - `VK_FORMAT_G16_B16R16_2PLANE_422_UNORM`

  - `VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM`

  - `VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM`

  - `VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM`

  - `VK_FORMAT_G8B8G8R8_422_UNORM`

  - `VK_FORMAT_G8_B8R8_2PLANE_420_UNORM`

  - `VK_FORMAT_G8_B8R8_2PLANE_422_UNORM`

  - `VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM`

  - `VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM`

  - `VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM`

  - `VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16`

  - `VK_FORMAT_R10X6G10X6_UNORM_2PACK16`

  - `VK_FORMAT_R10X6_UNORM_PACK16`

  - `VK_FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16`

  - `VK_FORMAT_R12X4G12X4_UNORM_2PACK16`

  - `VK_FORMAT_R12X4_UNORM_PACK16`

- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html):

  - `VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT`

  - `VK_FORMAT_FEATURE_DISJOINT_BIT`

  - `VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT`

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT`

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT`

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT`

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT`

  - `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`

  - `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT`

- Extending [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html):

  - `VK_IMAGE_ASPECT_PLANE_0_BIT`

  - `VK_IMAGE_ASPECT_PLANE_1_BIT`

  - `VK_IMAGE_ASPECT_PLANE_2_BIT`

- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html):

  - `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT`

  - `VK_IMAGE_CREATE_ALIAS_BIT`

  - `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT`

  - `VK_IMAGE_CREATE_DISJOINT_BIT`

  - `VK_IMAGE_CREATE_EXTENDED_USAGE_BIT`

  - `VK_IMAGE_CREATE_PROTECTED_BIT`

  - `VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT`

- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html):

  - `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`

  - `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`

- Extending [VkMemoryHeapFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHeapFlagBits.html):

  - `VK_MEMORY_HEAP_MULTI_INSTANCE_BIT`

- Extending [VkMemoryPropertyFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryPropertyFlagBits.html):

  - `VK_MEMORY_PROPERTY_PROTECTED_BIT`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE`

  - `VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION`

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html):

  - `VK_PIPELINE_CREATE_DISPATCH_BASE`

  - `VK_PIPELINE_CREATE_DISPATCH_BASE_BIT`

  - `VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT`

- Extending [VkQueueFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFlagBits.html):

  - `VK_QUEUE_PROTECTED_BIT`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_INVALID_EXTERNAL_HANDLE`

  - `VK_ERROR_OUT_OF_POOL_MEMORY`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO`

  - `VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO`

  - `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO`

  - `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO`

  - `VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO`

  - `VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2`

  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT`

  - `VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO`

  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO`

  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO`

  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO`

  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO`

  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO`

  - `VK_STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2`

  - `VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO`

  - `VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO`

  - `VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO`

  - `VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES`

  - `VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES`

  - `VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES`

  - `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO`

  - `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO`

  - `VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES`

  - `VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2`

  - `VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2`

  - `VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2`

  - `VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO`

  - `VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2`

  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO`

  - `VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO`

  - `VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO`

  - `VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS`

  - `VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_FEATURES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETERS_FEATURES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DRAW_PARAMETER_FEATURES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES`

  - `VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO`

  - `VK_STRUCTURE_TYPE_PROTECTED_SUBMIT_INFO`

  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2`

  - `VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO`

  - `VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO`

  - `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO`

  - `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES`

  - `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO`

  - `VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2`

  - `VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
