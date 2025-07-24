# vkCmdSetViewport(3) Manual Page

## Name

vkCmdSetViewport - Set the viewport dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the viewport transformation parameters, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdSetViewport(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstViewport,
    uint32_t                                    viewportCount,
    const VkViewport*                           pViewports);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `firstViewport` is the index of the first viewport whose parameters are updated by the command.
- `viewportCount` is the number of viewports whose parameters are updated by the command.
- `pViewports` is a pointer to an array of [VkViewport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewport.html) structures specifying viewport parameters.

## [](#_description)Description

This command sets the viewport transformation parameters state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_VIEWPORT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the `VkPipelineViewportStateCreateInfo`::`pViewports` values used to create the currently active pipeline.

The viewport parameters taken from element i of `pViewports` replace the current state for the viewport index `firstViewport` + i, for i in [0, `viewportCount`).

Valid Usage

- [](#VUID-vkCmdSetViewport-firstViewport-01223)VUID-vkCmdSetViewport-firstViewport-01223  
  The sum of `firstViewport` and `viewportCount` **must** be between `1` and `VkPhysicalDeviceLimits`::`maxViewports`, inclusive
- [](#VUID-vkCmdSetViewport-firstViewport-01224)VUID-vkCmdSetViewport-firstViewport-01224  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `firstViewport` **must** be `0`
- [](#VUID-vkCmdSetViewport-viewportCount-01225)VUID-vkCmdSetViewport-viewportCount-01225  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `viewportCount` **must** be `1`
- [](#VUID-vkCmdSetViewport-commandBuffer-04821)VUID-vkCmdSetViewport-commandBuffer-04821  
  `commandBuffer` **must** not have [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D` enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetViewport-commandBuffer-parameter)VUID-vkCmdSetViewport-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetViewport-pViewports-parameter)VUID-vkCmdSetViewport-pViewports-parameter  
  `pViewports` **must** be a valid pointer to an array of `viewportCount` valid [VkViewport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewport.html) structures
- [](#VUID-vkCmdSetViewport-commandBuffer-recording)VUID-vkCmdSetViewport-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetViewport-commandBuffer-cmdpool)VUID-vkCmdSetViewport-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetViewport-videocoding)VUID-vkCmdSetViewport-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetViewport-viewportCount-arraylength)VUID-vkCmdSetViewport-viewportCount-arraylength  
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

Conditional Rendering

vkCmdSetViewport is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkViewport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewport.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetViewport)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0