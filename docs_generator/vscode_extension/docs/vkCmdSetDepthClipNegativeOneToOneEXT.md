# vkCmdSetDepthClipNegativeOneToOneEXT(3) Manual Page

## Name

vkCmdSetDepthClipNegativeOneToOneEXT - Specify the negative one to one depth clip mode dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) `negativeOneToOne`, call:

```c++
// Provided by VK_EXT_depth_clip_control with VK_EXT_extended_dynamic_state3, VK_EXT_depth_clip_control with VK_EXT_shader_object
void vkCmdSetDepthClipNegativeOneToOneEXT(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    negativeOneToOne);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `negativeOneToOne` specifies the `negativeOneToOne` state.

## [](#_description)Description

This command sets the `negativeOneToOne` state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_DEPTH_CLIP_NEGATIVE_ONE_TO_ONE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineViewportDepthClipControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportDepthClipControlCreateInfoEXT.html)::`negativeOneToOne` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetDepthClipNegativeOneToOneEXT-None-09423)VUID-vkCmdSetDepthClipNegativeOneToOneEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3DepthClipNegativeOneToOne`](#features-extendedDynamicState3DepthClipNegativeOneToOne) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled
- [](#VUID-vkCmdSetDepthClipNegativeOneToOneEXT-depthClipControl-07453)VUID-vkCmdSetDepthClipNegativeOneToOneEXT-depthClipControl-07453  
  The [`depthClipControl`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-depthClipControl) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetDepthClipNegativeOneToOneEXT-commandBuffer-parameter)VUID-vkCmdSetDepthClipNegativeOneToOneEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetDepthClipNegativeOneToOneEXT-commandBuffer-recording)VUID-vkCmdSetDepthClipNegativeOneToOneEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetDepthClipNegativeOneToOneEXT-commandBuffer-cmdpool)VUID-vkCmdSetDepthClipNegativeOneToOneEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetDepthClipNegativeOneToOneEXT-videocoding)VUID-vkCmdSetDepthClipNegativeOneToOneEXT-videocoding  
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

## [](#_see_also)See Also

[VK\_EXT\_depth\_clip\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clip_control.html), [VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetDepthClipNegativeOneToOneEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0