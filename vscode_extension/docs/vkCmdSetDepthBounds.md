# vkCmdSetDepthBounds(3) Manual Page

## Name

vkCmdSetDepthBounds - Set depth bounds range dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the depth bounds range, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdSetDepthBounds(
    VkCommandBuffer                             commandBuffer,
    float                                       minDepthBounds,
    float                                       maxDepthBounds);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `minDepthBounds` is the minimum depth bound.
- `maxDepthBounds` is the maximum depth bound.

## [](#_description)Description

This command sets the depth bounds range for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_DEPTH_BOUNDS` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateInfo.html)::`minDepthBounds` and [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateInfo.html)::`maxDepthBounds` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetDepthBounds-minDepthBounds-00600)VUID-vkCmdSetDepthBounds-minDepthBounds-00600  
  If the `VK_EXT_depth_range_unrestricted` extension is not enabled `minDepthBounds` **must** be between `0.0` and `1.0`, inclusive
- [](#VUID-vkCmdSetDepthBounds-maxDepthBounds-00601)VUID-vkCmdSetDepthBounds-maxDepthBounds-00601  
  If the `VK_EXT_depth_range_unrestricted` extension is not enabled `maxDepthBounds` **must** be between `0.0` and `1.0`, inclusive
- [](#VUID-vkCmdSetDepthBounds-minDepthBounds-10912)VUID-vkCmdSetDepthBounds-minDepthBounds-10912  
  `minDepthBounds` **must** be less than or equal to `maxDepthBounds`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetDepthBounds-commandBuffer-parameter)VUID-vkCmdSetDepthBounds-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetDepthBounds-commandBuffer-recording)VUID-vkCmdSetDepthBounds-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetDepthBounds-commandBuffer-cmdpool)VUID-vkCmdSetDepthBounds-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdSetDepthBounds-videocoding)VUID-vkCmdSetDepthBounds-videocoding  
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

VK\_QUEUE\_GRAPHICS\_BIT

State

Conditional Rendering

vkCmdSetDepthBounds is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetDepthBounds).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0