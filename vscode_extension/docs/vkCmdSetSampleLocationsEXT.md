# vkCmdSetSampleLocationsEXT(3) Manual Page

## Name

vkCmdSetSampleLocationsEXT - Set sample locations dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the sample locations used for rasterization, call:

```c++
// Provided by VK_EXT_sample_locations
void vkCmdSetSampleLocationsEXT(
    VkCommandBuffer                             commandBuffer,
    const VkSampleLocationsInfoEXT*             pSampleLocationsInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pSampleLocationsInfo` is the sample locations state to set.

## [](#_description)Description

This command sets the custom sample locations for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`, and when the [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)::`sampleLocationsEnable` property of the bound graphics pipeline is `VK_TRUE`. Otherwise, this state is specified by the [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)::`sampleLocationsInfo` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetSampleLocationsEXT-variableSampleLocations-01530)VUID-vkCmdSetSampleLocationsEXT-variableSampleLocations-01530  
  If [VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html)::`variableSampleLocations` is `VK_FALSE` then the current render pass **must** have been begun by specifying a [VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html) structure whose `pPostSubpassSampleLocations` member contains an element with a `subpassIndex` matching the current subpass index and the `sampleLocationsInfo` member of that element **must** match the sample locations state pointed to by `pSampleLocationsInfo`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetSampleLocationsEXT-commandBuffer-parameter)VUID-vkCmdSetSampleLocationsEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetSampleLocationsEXT-pSampleLocationsInfo-parameter)VUID-vkCmdSetSampleLocationsEXT-pSampleLocationsInfo-parameter  
  `pSampleLocationsInfo` **must** be a valid pointer to a valid [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html) structure
- [](#VUID-vkCmdSetSampleLocationsEXT-commandBuffer-recording)VUID-vkCmdSetSampleLocationsEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetSampleLocationsEXT-commandBuffer-cmdpool)VUID-vkCmdSetSampleLocationsEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetSampleLocationsEXT-videocoding)VUID-vkCmdSetSampleLocationsEXT-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Outside

Graphics

State

Conditional Rendering

vkCmdSetSampleLocationsEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_sample\_locations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sample_locations.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetSampleLocationsEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0