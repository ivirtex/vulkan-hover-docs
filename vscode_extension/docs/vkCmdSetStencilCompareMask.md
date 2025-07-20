# vkCmdSetStencilCompareMask(3) Manual Page

## Name

vkCmdSetStencilCompareMask - Set stencil compare mask dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the stencil compare mask, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdSetStencilCompareMask(
    VkCommandBuffer                             commandBuffer,
    VkStencilFaceFlags                          faceMask,
    uint32_t                                    compareMask);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `faceMask` is a bitmask of [VkStencilFaceFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilFaceFlagBits.html) specifying the set of stencil state for which to update the compare mask.
- `compareMask` is the new value to use as the stencil compare mask.

## [](#_description)Description

This command sets the stencil compare mask for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_STENCIL_COMPARE_MASK` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkStencilOpState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOpState.html)::`compareMask` value used to create the currently active pipeline, for both front and back faces.

Valid Usage (Implicit)

- [](#VUID-vkCmdSetStencilCompareMask-commandBuffer-parameter)VUID-vkCmdSetStencilCompareMask-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetStencilCompareMask-faceMask-parameter)VUID-vkCmdSetStencilCompareMask-faceMask-parameter  
  `faceMask` **must** be a valid combination of [VkStencilFaceFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilFaceFlagBits.html) values
- [](#VUID-vkCmdSetStencilCompareMask-faceMask-requiredbitmask)VUID-vkCmdSetStencilCompareMask-faceMask-requiredbitmask  
  `faceMask` **must** not be `0`
- [](#VUID-vkCmdSetStencilCompareMask-commandBuffer-recording)VUID-vkCmdSetStencilCompareMask-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetStencilCompareMask-commandBuffer-cmdpool)VUID-vkCmdSetStencilCompareMask-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetStencilCompareMask-videocoding)VUID-vkCmdSetStencilCompareMask-videocoding  
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

vkCmdSetStencilCompareMask is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkStencilFaceFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilFaceFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetStencilCompareMask)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0