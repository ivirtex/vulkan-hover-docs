# VkDebugReportObjectTypeEXT(3) Manual Page

## Name

VkDebugReportObjectTypeEXT - Specify the type of an object handle



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values passed to the `objectType` parameter of the callback
function specified by
[VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html)::`pfnCallback`,
specifying the type of object handle being reported, are:

``` c
// Provided by VK_EXT_debug_marker, VK_EXT_debug_report
typedef enum VkDebugReportObjectTypeEXT {
    VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT = 0,
    VK_DEBUG_REPORT_OBJECT_TYPE_INSTANCE_EXT = 1,
    VK_DEBUG_REPORT_OBJECT_TYPE_PHYSICAL_DEVICE_EXT = 2,
    VK_DEBUG_REPORT_OBJECT_TYPE_DEVICE_EXT = 3,
    VK_DEBUG_REPORT_OBJECT_TYPE_QUEUE_EXT = 4,
    VK_DEBUG_REPORT_OBJECT_TYPE_SEMAPHORE_EXT = 5,
    VK_DEBUG_REPORT_OBJECT_TYPE_COMMAND_BUFFER_EXT = 6,
    VK_DEBUG_REPORT_OBJECT_TYPE_FENCE_EXT = 7,
    VK_DEBUG_REPORT_OBJECT_TYPE_DEVICE_MEMORY_EXT = 8,
    VK_DEBUG_REPORT_OBJECT_TYPE_BUFFER_EXT = 9,
    VK_DEBUG_REPORT_OBJECT_TYPE_IMAGE_EXT = 10,
    VK_DEBUG_REPORT_OBJECT_TYPE_EVENT_EXT = 11,
    VK_DEBUG_REPORT_OBJECT_TYPE_QUERY_POOL_EXT = 12,
    VK_DEBUG_REPORT_OBJECT_TYPE_BUFFER_VIEW_EXT = 13,
    VK_DEBUG_REPORT_OBJECT_TYPE_IMAGE_VIEW_EXT = 14,
    VK_DEBUG_REPORT_OBJECT_TYPE_SHADER_MODULE_EXT = 15,
    VK_DEBUG_REPORT_OBJECT_TYPE_PIPELINE_CACHE_EXT = 16,
    VK_DEBUG_REPORT_OBJECT_TYPE_PIPELINE_LAYOUT_EXT = 17,
    VK_DEBUG_REPORT_OBJECT_TYPE_RENDER_PASS_EXT = 18,
    VK_DEBUG_REPORT_OBJECT_TYPE_PIPELINE_EXT = 19,
    VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT_EXT = 20,
    VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_EXT = 21,
    VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_POOL_EXT = 22,
    VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_EXT = 23,
    VK_DEBUG_REPORT_OBJECT_TYPE_FRAMEBUFFER_EXT = 24,
    VK_DEBUG_REPORT_OBJECT_TYPE_COMMAND_POOL_EXT = 25,
    VK_DEBUG_REPORT_OBJECT_TYPE_SURFACE_KHR_EXT = 26,
    VK_DEBUG_REPORT_OBJECT_TYPE_SWAPCHAIN_KHR_EXT = 27,
    VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT = 28,
    VK_DEBUG_REPORT_OBJECT_TYPE_DISPLAY_KHR_EXT = 29,
    VK_DEBUG_REPORT_OBJECT_TYPE_DISPLAY_MODE_KHR_EXT = 30,
    VK_DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT = 33,
  // Provided by VK_VERSION_1_1 with VK_EXT_debug_report, VK_KHR_sampler_ycbcr_conversion with VK_EXT_debug_report
    VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT = 1000156000,
  // Provided by VK_VERSION_1_1 with VK_EXT_debug_report
    VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_EXT = 1000085000,
  // Provided by VK_EXT_debug_report with VK_NVX_binary_import
    VK_DEBUG_REPORT_OBJECT_TYPE_CU_MODULE_NVX_EXT = 1000029000,
  // Provided by VK_EXT_debug_report with VK_NVX_binary_import
    VK_DEBUG_REPORT_OBJECT_TYPE_CU_FUNCTION_NVX_EXT = 1000029001,
  // Provided by VK_KHR_acceleration_structure with VK_EXT_debug_report
    VK_DEBUG_REPORT_OBJECT_TYPE_ACCELERATION_STRUCTURE_KHR_EXT = 1000150000,
  // Provided by VK_EXT_debug_report with VK_NV_ray_tracing
    VK_DEBUG_REPORT_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV_EXT = 1000165000,
  // Provided by VK_EXT_debug_report with VK_NV_cuda_kernel_launch
    VK_DEBUG_REPORT_OBJECT_TYPE_CUDA_MODULE_NV_EXT = 1000307000,
  // Provided by VK_EXT_debug_report with VK_NV_cuda_kernel_launch
    VK_DEBUG_REPORT_OBJECT_TYPE_CUDA_FUNCTION_NV_EXT = 1000307001,
  // Provided by VK_EXT_debug_report with VK_FUCHSIA_buffer_collection
    VK_DEBUG_REPORT_OBJECT_TYPE_BUFFER_COLLECTION_FUCHSIA_EXT = 1000366000,
    VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_EXT = VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT,
    VK_DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT = VK_DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT,
  // Provided by VK_KHR_descriptor_update_template with VK_EXT_debug_report
    VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR_EXT = VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_EXT,
  // Provided by VK_KHR_sampler_ycbcr_conversion with VK_EXT_debug_report
    VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_KHR_EXT = VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT,
} VkDebugReportObjectTypeEXT;
```

## <a href="#_description" class="anchor"></a>Description

| [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html) | Vulkan Handle Type                                            |
|---------------------------------------------------------------|---------------------------------------------------------------|
| `VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT`                     | Unknown/Undefined Handle                                      |
| `VK_DEBUG_REPORT_OBJECT_TYPE_INSTANCE_EXT`                    | [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)                                 |
| `VK_DEBUG_REPORT_OBJECT_TYPE_PHYSICAL_DEVICE_EXT`             | [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)                     |
| `VK_DEBUG_REPORT_OBJECT_TYPE_DEVICE_EXT`                      | [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)                                     |
| `VK_DEBUG_REPORT_OBJECT_TYPE_QUEUE_EXT`                       | [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html)                                       |
| `VK_DEBUG_REPORT_OBJECT_TYPE_SEMAPHORE_EXT`                   | [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html)                               |
| `VK_DEBUG_REPORT_OBJECT_TYPE_COMMAND_BUFFER_EXT`              | [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)                       |
| `VK_DEBUG_REPORT_OBJECT_TYPE_FENCE_EXT`                       | [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html)                                       |
| `VK_DEBUG_REPORT_OBJECT_TYPE_DEVICE_MEMORY_EXT`               | [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)                         |
| `VK_DEBUG_REPORT_OBJECT_TYPE_BUFFER_EXT`                      | [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html)                                     |
| `VK_DEBUG_REPORT_OBJECT_TYPE_IMAGE_EXT`                       | [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html)                                       |
| `VK_DEBUG_REPORT_OBJECT_TYPE_EVENT_EXT`                       | [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html)                                       |
| `VK_DEBUG_REPORT_OBJECT_TYPE_QUERY_POOL_EXT`                  | [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html)                               |
| `VK_DEBUG_REPORT_OBJECT_TYPE_BUFFER_VIEW_EXT`                 | [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html)                             |
| `VK_DEBUG_REPORT_OBJECT_TYPE_IMAGE_VIEW_EXT`                  | [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html)                               |
| `VK_DEBUG_REPORT_OBJECT_TYPE_SHADER_MODULE_EXT`               | [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html)                         |
| `VK_DEBUG_REPORT_OBJECT_TYPE_PIPELINE_CACHE_EXT`              | [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html)                       |
| `VK_DEBUG_REPORT_OBJECT_TYPE_PIPELINE_LAYOUT_EXT`             | [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)                     |
| `VK_DEBUG_REPORT_OBJECT_TYPE_RENDER_PASS_EXT`                 | [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)                             |
| `VK_DEBUG_REPORT_OBJECT_TYPE_PIPELINE_EXT`                    | [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)                                 |
| `VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT_EXT`       | [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html)           |
| `VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_EXT`                     | [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html)                                   |
| `VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_POOL_EXT`             | [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html)                     |
| `VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_EXT`              | [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html)                       |
| `VK_DEBUG_REPORT_OBJECT_TYPE_FRAMEBUFFER_EXT`                 | [VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html)                           |
| `VK_DEBUG_REPORT_OBJECT_TYPE_COMMAND_POOL_EXT`                | [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html)                           |
| `VK_DEBUG_REPORT_OBJECT_TYPE_SURFACE_KHR_EXT`                 | [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)                             |
| `VK_DEBUG_REPORT_OBJECT_TYPE_SWAPCHAIN_KHR_EXT`               | [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)                         |
| `VK_DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT`   | [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackEXT.html)     |
| `VK_DEBUG_REPORT_OBJECT_TYPE_DISPLAY_KHR_EXT`                 | [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html)                             |
| `VK_DEBUG_REPORT_OBJECT_TYPE_DISPLAY_MODE_KHR_EXT`            | [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html)                     |
| `VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_EXT`  | [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html) |

Table 1. `VkDebugReportObjectTypeEXT` and Vulkan Handle Relationship
{#debug-report-object-types}

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The primary expected use of
<code>VK_ERROR_VALIDATION_FAILED_EXT</code> is for validation layer
testing. It is not expected that an application would see this error
code during normal use of the validation layers.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_marker](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_marker.html),
[VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html),
[VkDebugMarkerObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugMarkerObjectNameInfoEXT.html),
[VkDebugMarkerObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugMarkerObjectTagInfoEXT.html),
[vkDebugReportMessageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDebugReportMessageEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDebugReportObjectTypeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
