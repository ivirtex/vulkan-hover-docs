# vkCmdSetScissorWithCount(3) Manual Page

## Name

vkCmdSetScissorWithCount - Set the scissor count and scissor rectangular bounds dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the scissor count and scissor rectangular bounds, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdSetScissorWithCount(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    scissorCount,
    const VkRect2D*                             pScissors);
```

or the equivalent command

```c++
// Provided by VK_EXT_extended_dynamic_state, VK_EXT_shader_object
void vkCmdSetScissorWithCountEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    scissorCount,
    const VkRect2D*                             pScissors);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `scissorCount` specifies the scissor count.
- `pScissors` specifies the scissors to use for drawing.

## [](#_description)Description

This command sets the scissor count and scissor rectangular bounds state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the corresponding [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html)::`scissorCount` and `pScissors` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetScissorWithCount-None-08971)VUID-vkCmdSetScissorWithCount-None-08971  
  At least one of the following **must** be true:
  
  - the [`extendedDynamicState`](#features-extendedDynamicState) feature is enabled
  - the [`shaderObject`](#features-shaderObject) feature is enabled
  - the value of [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` used to create the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) parent of `commandBuffer` is greater than or equal to Version 1.3
- [](#VUID-vkCmdSetScissorWithCount-scissorCount-03397)VUID-vkCmdSetScissorWithCount-scissorCount-03397  
  `scissorCount` **must** be between `1` and `VkPhysicalDeviceLimits`::`maxViewports`, inclusive
- [](#VUID-vkCmdSetScissorWithCount-scissorCount-03398)VUID-vkCmdSetScissorWithCount-scissorCount-03398  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `scissorCount` **must** be `1`
- [](#VUID-vkCmdSetScissorWithCount-x-03399)VUID-vkCmdSetScissorWithCount-x-03399  
  The `x` and `y` members of `offset` member of any element of `pScissors` **must** be greater than or equal to `0`
- [](#VUID-vkCmdSetScissorWithCount-offset-03400)VUID-vkCmdSetScissorWithCount-offset-03400  
  Evaluation of (`offset.x` + `extent.width`) **must** not cause a signed integer addition overflow for any element of `pScissors`
- [](#VUID-vkCmdSetScissorWithCount-offset-03401)VUID-vkCmdSetScissorWithCount-offset-03401  
  Evaluation of (`offset.y` + `extent.height`) **must** not cause a signed integer addition overflow for any element of `pScissors`
- [](#VUID-vkCmdSetScissorWithCount-commandBuffer-04820)VUID-vkCmdSetScissorWithCount-commandBuffer-04820  
  `commandBuffer` **must** not have [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D` enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetScissorWithCount-commandBuffer-parameter)VUID-vkCmdSetScissorWithCount-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetScissorWithCount-pScissors-parameter)VUID-vkCmdSetScissorWithCount-pScissors-parameter  
  `pScissors` **must** be a valid pointer to an array of `scissorCount` [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) structures
- [](#VUID-vkCmdSetScissorWithCount-commandBuffer-recording)VUID-vkCmdSetScissorWithCount-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetScissorWithCount-commandBuffer-cmdpool)VUID-vkCmdSetScissorWithCount-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetScissorWithCount-videocoding)VUID-vkCmdSetScissorWithCount-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetScissorWithCount-scissorCount-arraylength)VUID-vkCmdSetScissorWithCount-scissorCount-arraylength  
  `scissorCount` **must** be greater than `0`

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

vkCmdSetScissorWithCount is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetScissorWithCount).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0