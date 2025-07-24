# vkCmdSetRasterizationStreamEXT(3) Manual Page

## Name

vkCmdSetRasterizationStreamEXT - Specify the rasterization stream dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the `rasterizationStream` state, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3 with VK_EXT_transform_feedback, VK_EXT_shader_object with VK_EXT_transform_feedback
void vkCmdSetRasterizationStreamEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    rasterizationStream);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `rasterizationStream` specifies the `rasterizationStream` state.

## [](#_description)Description

This command sets the `rasterizationStream` state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_RASTERIZATION_STREAM_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineRasterizationStateStreamCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateStreamCreateInfoEXT.html)::`rasterizationStream` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetRasterizationStreamEXT-None-09423)VUID-vkCmdSetRasterizationStreamEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3RasterizationStream`](#features-extendedDynamicState3RasterizationStream) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled
- [](#VUID-vkCmdSetRasterizationStreamEXT-transformFeedback-07411)VUID-vkCmdSetRasterizationStreamEXT-transformFeedback-07411  
  The [`transformFeedback`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-transformFeedback) feature **must** be enabled
- [](#VUID-vkCmdSetRasterizationStreamEXT-rasterizationStream-07412)VUID-vkCmdSetRasterizationStreamEXT-rasterizationStream-07412  
  `rasterizationStream` **must** be less than [VkPhysicalDeviceTransformFeedbackPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTransformFeedbackPropertiesEXT.html)::`maxTransformFeedbackStreams`
- [](#VUID-vkCmdSetRasterizationStreamEXT-rasterizationStream-07413)VUID-vkCmdSetRasterizationStreamEXT-rasterizationStream-07413  
  `rasterizationStream` **must** be zero if `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`transformFeedbackRasterizationStreamSelect` is `VK_FALSE`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetRasterizationStreamEXT-commandBuffer-parameter)VUID-vkCmdSetRasterizationStreamEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetRasterizationStreamEXT-commandBuffer-recording)VUID-vkCmdSetRasterizationStreamEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetRasterizationStreamEXT-commandBuffer-cmdpool)VUID-vkCmdSetRasterizationStreamEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetRasterizationStreamEXT-videocoding)VUID-vkCmdSetRasterizationStreamEXT-videocoding  
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

vkCmdSetRasterizationStreamEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_EXT\_transform\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_transform_feedback.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetRasterizationStreamEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0