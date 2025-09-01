# vkCmdSetDepthBias2EXT(3) Manual Page

## Name

vkCmdSetDepthBias2EXT - Set depth bias factors and clamp dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the depth bias parameters, call:

```c++
// Provided by VK_EXT_depth_bias_control
void vkCmdSetDepthBias2EXT(
    VkCommandBuffer                             commandBuffer,
    const VkDepthBiasInfoEXT*                   pDepthBiasInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pDepthBiasInfo` is a pointer to a [VkDepthBiasInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasInfoEXT.html) structure specifying depth bias parameters.

## [](#_description)Description

This command is functionally identical to [vkCmdSetDepthBias](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthBias.html), but includes extensible sub-structures that include `sType` and `pNext` parameters, allowing them to be more easily extended.

Valid Usage (Implicit)

- [](#VUID-vkCmdSetDepthBias2EXT-commandBuffer-parameter)VUID-vkCmdSetDepthBias2EXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetDepthBias2EXT-pDepthBiasInfo-parameter)VUID-vkCmdSetDepthBias2EXT-pDepthBiasInfo-parameter  
  `pDepthBiasInfo` **must** be a valid pointer to a valid [VkDepthBiasInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasInfoEXT.html) structure
- [](#VUID-vkCmdSetDepthBias2EXT-commandBuffer-recording)VUID-vkCmdSetDepthBias2EXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetDepthBias2EXT-commandBuffer-cmdpool)VUID-vkCmdSetDepthBias2EXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetDepthBias2EXT-videocoding)VUID-vkCmdSetDepthBias2EXT-videocoding  
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

vkCmdSetDepthBias2EXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_depth\_bias\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_bias_control.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDepthBiasInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetDepthBias2EXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0