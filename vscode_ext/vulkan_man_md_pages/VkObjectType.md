# VkObjectType(3) Manual Page

## Name

VkObjectType - Specify an enumeration to track object handle types



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html) enumeration defines values, each
of which corresponds to a specific Vulkan handle type. These values
**can** be used to associate debug information with a particular type of
object through one or more extensions.

``` c
// Provided by VK_VERSION_1_0
typedef enum VkObjectType {
    VK_OBJECT_TYPE_UNKNOWN = 0,
    VK_OBJECT_TYPE_INSTANCE = 1,
    VK_OBJECT_TYPE_PHYSICAL_DEVICE = 2,
    VK_OBJECT_TYPE_DEVICE = 3,
    VK_OBJECT_TYPE_QUEUE = 4,
    VK_OBJECT_TYPE_SEMAPHORE = 5,
    VK_OBJECT_TYPE_COMMAND_BUFFER = 6,
    VK_OBJECT_TYPE_FENCE = 7,
    VK_OBJECT_TYPE_DEVICE_MEMORY = 8,
    VK_OBJECT_TYPE_BUFFER = 9,
    VK_OBJECT_TYPE_IMAGE = 10,
    VK_OBJECT_TYPE_EVENT = 11,
    VK_OBJECT_TYPE_QUERY_POOL = 12,
    VK_OBJECT_TYPE_BUFFER_VIEW = 13,
    VK_OBJECT_TYPE_IMAGE_VIEW = 14,
    VK_OBJECT_TYPE_SHADER_MODULE = 15,
    VK_OBJECT_TYPE_PIPELINE_CACHE = 16,
    VK_OBJECT_TYPE_PIPELINE_LAYOUT = 17,
    VK_OBJECT_TYPE_RENDER_PASS = 18,
    VK_OBJECT_TYPE_PIPELINE = 19,
    VK_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT = 20,
    VK_OBJECT_TYPE_SAMPLER = 21,
    VK_OBJECT_TYPE_DESCRIPTOR_POOL = 22,
    VK_OBJECT_TYPE_DESCRIPTOR_SET = 23,
    VK_OBJECT_TYPE_FRAMEBUFFER = 24,
    VK_OBJECT_TYPE_COMMAND_POOL = 25,
  // Provided by VK_VERSION_1_1
    VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION = 1000156000,
  // Provided by VK_VERSION_1_1
    VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE = 1000085000,
  // Provided by VK_VERSION_1_3
    VK_OBJECT_TYPE_PRIVATE_DATA_SLOT = 1000295000,
  // Provided by VK_KHR_surface
    VK_OBJECT_TYPE_SURFACE_KHR = 1000000000,
  // Provided by VK_KHR_swapchain
    VK_OBJECT_TYPE_SWAPCHAIN_KHR = 1000001000,
  // Provided by VK_KHR_display
    VK_OBJECT_TYPE_DISPLAY_KHR = 1000002000,
  // Provided by VK_KHR_display
    VK_OBJECT_TYPE_DISPLAY_MODE_KHR = 1000002001,
  // Provided by VK_EXT_debug_report
    VK_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT = 1000011000,
  // Provided by VK_KHR_video_queue
    VK_OBJECT_TYPE_VIDEO_SESSION_KHR = 1000023000,
  // Provided by VK_KHR_video_queue
    VK_OBJECT_TYPE_VIDEO_SESSION_PARAMETERS_KHR = 1000023001,
  // Provided by VK_NVX_binary_import
    VK_OBJECT_TYPE_CU_MODULE_NVX = 1000029000,
  // Provided by VK_NVX_binary_import
    VK_OBJECT_TYPE_CU_FUNCTION_NVX = 1000029001,
  // Provided by VK_EXT_debug_utils
    VK_OBJECT_TYPE_DEBUG_UTILS_MESSENGER_EXT = 1000128000,
  // Provided by VK_KHR_acceleration_structure
    VK_OBJECT_TYPE_ACCELERATION_STRUCTURE_KHR = 1000150000,
  // Provided by VK_EXT_validation_cache
    VK_OBJECT_TYPE_VALIDATION_CACHE_EXT = 1000160000,
  // Provided by VK_NV_ray_tracing
    VK_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV = 1000165000,
  // Provided by VK_INTEL_performance_query
    VK_OBJECT_TYPE_PERFORMANCE_CONFIGURATION_INTEL = 1000210000,
  // Provided by VK_KHR_deferred_host_operations
    VK_OBJECT_TYPE_DEFERRED_OPERATION_KHR = 1000268000,
  // Provided by VK_NV_device_generated_commands
    VK_OBJECT_TYPE_INDIRECT_COMMANDS_LAYOUT_NV = 1000277000,
  // Provided by VK_NV_cuda_kernel_launch
    VK_OBJECT_TYPE_CUDA_MODULE_NV = 1000307000,
  // Provided by VK_NV_cuda_kernel_launch
    VK_OBJECT_TYPE_CUDA_FUNCTION_NV = 1000307001,
  // Provided by VK_FUCHSIA_buffer_collection
    VK_OBJECT_TYPE_BUFFER_COLLECTION_FUCHSIA = 1000366000,
  // Provided by VK_EXT_opacity_micromap
    VK_OBJECT_TYPE_MICROMAP_EXT = 1000396000,
  // Provided by VK_NV_optical_flow
    VK_OBJECT_TYPE_OPTICAL_FLOW_SESSION_NV = 1000464000,
  // Provided by VK_EXT_shader_object
    VK_OBJECT_TYPE_SHADER_EXT = 1000482000,
  // Provided by VK_KHR_descriptor_update_template
    VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR = VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_KHR = VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION,
  // Provided by VK_EXT_private_data
    VK_OBJECT_TYPE_PRIVATE_DATA_SLOT_EXT = VK_OBJECT_TYPE_PRIVATE_DATA_SLOT,
} VkObjectType;
```

## <a href="#_description" class="anchor"></a>Description

| [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html) | Vulkan Handle Type |
|----|----|
| `VK_OBJECT_TYPE_UNKNOWN` | Unknown/Undefined Handle |
| `VK_OBJECT_TYPE_INSTANCE` | [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) |
| `VK_OBJECT_TYPE_PHYSICAL_DEVICE` | [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) |
| `VK_OBJECT_TYPE_DEVICE` | [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) |
| `VK_OBJECT_TYPE_QUEUE` | [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) |
| `VK_OBJECT_TYPE_SEMAPHORE` | [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) |
| `VK_OBJECT_TYPE_COMMAND_BUFFER` | [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) |
| `VK_OBJECT_TYPE_FENCE` | [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) |
| `VK_OBJECT_TYPE_DEVICE_MEMORY` | [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) |
| `VK_OBJECT_TYPE_BUFFER` | [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) |
| `VK_OBJECT_TYPE_IMAGE` | [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) |
| `VK_OBJECT_TYPE_EVENT` | [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) |
| `VK_OBJECT_TYPE_QUERY_POOL` | [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) |
| `VK_OBJECT_TYPE_BUFFER_VIEW` | [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) |
| `VK_OBJECT_TYPE_IMAGE_VIEW` | [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) |
| `VK_OBJECT_TYPE_SHADER_MODULE` | [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html) |
| `VK_OBJECT_TYPE_PIPELINE_CACHE` | [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html) |
| `VK_OBJECT_TYPE_PIPELINE_LAYOUT` | [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) |
| `VK_OBJECT_TYPE_RENDER_PASS` | [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) |
| `VK_OBJECT_TYPE_PIPELINE` | [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) |
| `VK_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT` | [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) |
| `VK_OBJECT_TYPE_SAMPLER` | [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) |
| `VK_OBJECT_TYPE_DESCRIPTOR_POOL` | [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html) |
| `VK_OBJECT_TYPE_DESCRIPTOR_SET` | [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html) |
| `VK_OBJECT_TYPE_FRAMEBUFFER` | [VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html) |
| `VK_OBJECT_TYPE_COMMAND_POOL` | [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) |
| `VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION` | [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html) |
| `VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE` | [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html) |
| `VK_OBJECT_TYPE_PRIVATE_DATA_SLOT` | [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html) |
| `VK_OBJECT_TYPE_SURFACE_KHR` | [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) |
| `VK_OBJECT_TYPE_SWAPCHAIN_KHR` | [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) |
| `VK_OBJECT_TYPE_DISPLAY_KHR` | [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) |
| `VK_OBJECT_TYPE_DISPLAY_MODE_KHR` | [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html) |
| `VK_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT` | [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackEXT.html) |
| `VK_OBJECT_TYPE_VIDEO_SESSION_KHR` | [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html) |
| `VK_OBJECT_TYPE_VIDEO_SESSION_PARAMETERS_KHR` | [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html) |
| `VK_OBJECT_TYPE_DEBUG_UTILS_MESSENGER_EXT` | [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerEXT.html) |
| `VK_OBJECT_TYPE_ACCELERATION_STRUCTURE_KHR` | [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) |
| `VK_OBJECT_TYPE_VALIDATION_CACHE_EXT` | [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html) |
| `VK_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV` | [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) |
| `VK_OBJECT_TYPE_PERFORMANCE_CONFIGURATION_INTEL` | [VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationINTEL.html) |
| `VK_OBJECT_TYPE_DEFERRED_OPERATION_KHR` | [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) |
| `VK_OBJECT_TYPE_INDIRECT_COMMANDS_LAYOUT_NV` | [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutNV.html) |
| `VK_OBJECT_TYPE_BUFFER_COLLECTION_FUCHSIA` | [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) |
| `VK_OBJECT_TYPE_MICROMAP_EXT` | [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html) |
| `VK_OBJECT_TYPE_OPTICAL_FLOW_SESSION_NV` | [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionNV.html) |
| `VK_OBJECT_TYPE_SHADER_EXT` | [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) |

Table 1. `VkObjectType` and Vulkan Handle Relationship
{#debugging-object-types}

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDebugUtilsObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectNameInfoEXT.html),
[VkDebugUtilsObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectTagInfoEXT.html),
[VkDeviceMemoryReportCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportCallbackDataEXT.html),
[vkGetPrivateData](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPrivateData.html),
[vkGetPrivateDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPrivateDataEXT.html),
[vkSetPrivateData](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetPrivateData.html),
[vkSetPrivateDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetPrivateDataEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkObjectType"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
