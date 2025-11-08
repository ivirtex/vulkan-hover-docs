# vkCmdSetViewportWithCount(3) Manual Page

## Name

vkCmdSetViewportWithCount - Set the viewport count and viewports dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the viewport count and viewports, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdSetViewportWithCount(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    viewportCount,
    const VkViewport*                           pViewports);
```

or the equivalent command

```c++
// Provided by VK_EXT_extended_dynamic_state, VK_EXT_shader_object
void vkCmdSetViewportWithCountEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    viewportCount,
    const VkViewport*                           pViewports);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `viewportCount` specifies the viewport count.
- `pViewports` specifies the viewports to use for drawing.

## [](#_description)Description

This command sets the viewport count and viewports state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the corresponding [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html)::`viewportCount` and `pViewports` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetViewportWithCount-None-08971)VUID-vkCmdSetViewportWithCount-None-08971  
  At least one of the following **must** be true:
  
  - the [`extendedDynamicState`](#features-extendedDynamicState) feature is enabled
  - the [`shaderObject`](#features-shaderObject) feature is enabled
  - the value of [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` used to create the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) parent of `commandBuffer` is greater than or equal to Version 1.3
- [](#VUID-vkCmdSetViewportWithCount-viewportCount-03394)VUID-vkCmdSetViewportWithCount-viewportCount-03394  
  `viewportCount` **must** be between `1` and `VkPhysicalDeviceLimits`::`maxViewports`, inclusive
- [](#VUID-vkCmdSetViewportWithCount-viewportCount-03395)VUID-vkCmdSetViewportWithCount-viewportCount-03395  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `viewportCount` **must** be `1`
- [](#VUID-vkCmdSetViewportWithCount-commandBuffer-04819)VUID-vkCmdSetViewportWithCount-commandBuffer-04819  
  `commandBuffer` **must** not have [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D` enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetViewportWithCount-commandBuffer-parameter)VUID-vkCmdSetViewportWithCount-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetViewportWithCount-pViewports-parameter)VUID-vkCmdSetViewportWithCount-pViewports-parameter  
  `pViewports` **must** be a valid pointer to an array of `viewportCount` valid [VkViewport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewport.html) structures
- [](#VUID-vkCmdSetViewportWithCount-commandBuffer-recording)VUID-vkCmdSetViewportWithCount-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetViewportWithCount-commandBuffer-cmdpool)VUID-vkCmdSetViewportWithCount-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdSetViewportWithCount-videocoding)VUID-vkCmdSetViewportWithCount-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetViewportWithCount-viewportCount-arraylength)VUID-vkCmdSetViewportWithCount-viewportCount-arraylength  
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

VK\_QUEUE\_GRAPHICS\_BIT

State

Conditional Rendering

vkCmdSetViewportWithCount is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkViewport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewport.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetViewportWithCount).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0