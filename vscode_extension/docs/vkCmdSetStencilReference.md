# vkCmdSetStencilReference(3) Manual Page

## Name

vkCmdSetStencilReference - Set stencil reference value dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the stencil reference value, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdSetStencilReference(
    VkCommandBuffer                             commandBuffer,
    VkStencilFaceFlags                          faceMask,
    uint32_t                                    reference);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `faceMask` is a bitmask of [VkStencilFaceFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilFaceFlagBits.html) specifying the set of stencil state for which to update the reference value, as described above for [vkCmdSetStencilCompareMask](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetStencilCompareMask.html).
- `reference` is the new value to use as the stencil reference value.

## [](#_description)Description

This command sets the stencil reference value for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_STENCIL_REFERENCE` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateInfo.html)::`reference` value used to create the currently active pipeline, for both front and back faces.

Valid Usage (Implicit)

- [](#VUID-vkCmdSetStencilReference-commandBuffer-parameter)VUID-vkCmdSetStencilReference-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetStencilReference-faceMask-parameter)VUID-vkCmdSetStencilReference-faceMask-parameter  
  `faceMask` **must** be a valid combination of [VkStencilFaceFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilFaceFlagBits.html) values
- [](#VUID-vkCmdSetStencilReference-faceMask-requiredbitmask)VUID-vkCmdSetStencilReference-faceMask-requiredbitmask  
  `faceMask` **must** not be `0`
- [](#VUID-vkCmdSetStencilReference-commandBuffer-recording)VUID-vkCmdSetStencilReference-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetStencilReference-commandBuffer-cmdpool)VUID-vkCmdSetStencilReference-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetStencilReference-videocoding)VUID-vkCmdSetStencilReference-videocoding  
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

vkCmdSetStencilReference is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkStencilFaceFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilFaceFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetStencilReference)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0