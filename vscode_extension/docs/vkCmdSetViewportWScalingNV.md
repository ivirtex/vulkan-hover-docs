# vkCmdSetViewportWScalingNV(3) Manual Page

## Name

vkCmdSetViewportWScalingNV - Set the viewport W scaling dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the viewport **W** scaling parameters, call:

```c++
// Provided by VK_NV_clip_space_w_scaling
void vkCmdSetViewportWScalingNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstViewport,
    uint32_t                                    viewportCount,
    const VkViewportWScalingNV*                 pViewportWScalings);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `firstViewport` is the index of the first viewport whose parameters are updated by the command.
- `viewportCount` is the number of viewports whose parameters are updated by the command.
- `pViewportWScalings` is a pointer to an array of [VkViewportWScalingNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportWScalingNV.html) structures specifying viewport parameters.

## [](#_description)Description

The viewport parameters taken from element i of `pViewportWScalings` replace the current state for the viewport index `firstViewport` + i, for i in [0, `viewportCount`).

This command sets the viewport **W** scaling for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_NV` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)::`pViewportWScalings` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetViewportWScalingNV-firstViewport-01324)VUID-vkCmdSetViewportWScalingNV-firstViewport-01324  
  The sum of `firstViewport` and `viewportCount` **must** be between `1` and [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`maxViewports`, inclusive

Valid Usage (Implicit)

- [](#VUID-vkCmdSetViewportWScalingNV-commandBuffer-parameter)VUID-vkCmdSetViewportWScalingNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetViewportWScalingNV-pViewportWScalings-parameter)VUID-vkCmdSetViewportWScalingNV-pViewportWScalings-parameter  
  `pViewportWScalings` **must** be a valid pointer to an array of `viewportCount` [VkViewportWScalingNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportWScalingNV.html) structures
- [](#VUID-vkCmdSetViewportWScalingNV-commandBuffer-recording)VUID-vkCmdSetViewportWScalingNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetViewportWScalingNV-commandBuffer-cmdpool)VUID-vkCmdSetViewportWScalingNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetViewportWScalingNV-videocoding)VUID-vkCmdSetViewportWScalingNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetViewportWScalingNV-viewportCount-arraylength)VUID-vkCmdSetViewportWScalingNV-viewportCount-arraylength  
  `viewportCount` **must** be greater than `0`

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

[VK\_NV\_clip\_space\_w\_scaling](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_clip_space_w_scaling.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkViewportWScalingNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportWScalingNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetViewportWScalingNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0