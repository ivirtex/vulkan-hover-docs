# vkCmdSetStencilWriteMask(3) Manual Page

## Name

vkCmdSetStencilWriteMask - Set stencil write mask dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the stencil write mask, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdSetStencilWriteMask(
    VkCommandBuffer                             commandBuffer,
    VkStencilFaceFlags                          faceMask,
    uint32_t                                    writeMask);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `faceMask` is a bitmask of [VkStencilFaceFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilFaceFlagBits.html) specifying the set of stencil state for which to update the write mask, as described above for [vkCmdSetStencilCompareMask](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetStencilCompareMask.html).
- `writeMask` is the new value to use as the stencil write mask.

## [](#_description)Description

This command sets the stencil write mask for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_STENCIL_WRITE_MASK` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the `writeMask` value used to create the currently active pipeline, for both [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateInfo.html)::`front` and [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateInfo.html)::`back` faces.

Valid Usage (Implicit)

- [](#VUID-vkCmdSetStencilWriteMask-commandBuffer-parameter)VUID-vkCmdSetStencilWriteMask-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetStencilWriteMask-faceMask-parameter)VUID-vkCmdSetStencilWriteMask-faceMask-parameter  
  `faceMask` **must** be a valid combination of [VkStencilFaceFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilFaceFlagBits.html) values
- [](#VUID-vkCmdSetStencilWriteMask-faceMask-requiredbitmask)VUID-vkCmdSetStencilWriteMask-faceMask-requiredbitmask  
  `faceMask` **must** not be `0`
- [](#VUID-vkCmdSetStencilWriteMask-commandBuffer-recording)VUID-vkCmdSetStencilWriteMask-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetStencilWriteMask-commandBuffer-cmdpool)VUID-vkCmdSetStencilWriteMask-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetStencilWriteMask-videocoding)VUID-vkCmdSetStencilWriteMask-videocoding  
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

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkStencilFaceFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilFaceFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetStencilWriteMask)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0