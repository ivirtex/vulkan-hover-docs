# vkCmdSetExclusiveScissorEnableNV(3) Manual Page

## Name

vkCmdSetExclusiveScissorEnableNV - Dynamically enable each exclusive scissor for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) whether an exclusive scissor is enabled or not, call:

```c++
// Provided by VK_NV_scissor_exclusive
void vkCmdSetExclusiveScissorEnableNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstExclusiveScissor,
    uint32_t                                    exclusiveScissorCount,
    const VkBool32*                             pExclusiveScissorEnables);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `firstExclusiveScissor` is the index of the first exclusive scissor rectangle whose state is updated by the command.
- `exclusiveScissorCount` is the number of exclusive scissor rectangles updated by the command.
- `pExclusiveScissorEnables` is a pointer to an array of [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html) values defining whether the exclusive scissor is enabled.

## [](#_description)Description

The exclusive scissor enables taken from element i of `pExclusiveScissorEnables` replace the current state for the scissor index `firstExclusiveScissor` + i, for i in [0, `exclusiveScissorCount`).

This command sets the exclusive scissor enable for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_ENABLE_NV` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is implied by the [VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html)::`exclusiveScissorCount` value used to create the currently active pipeline, where all `exclusiveScissorCount` exclusive scissors are implicitly enabled and the remainder up to `VkPhysicalDeviceLimits`::`maxViewports` are implicitly disabled.

Valid Usage

- [](#VUID-vkCmdSetExclusiveScissorEnableNV-exclusiveScissor-07853)VUID-vkCmdSetExclusiveScissorEnableNV-exclusiveScissor-07853  
  The [`exclusiveScissor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-exclusiveScissor) feature **must** be enabled, and the implementation **must** support at least `specVersion` `2` of the `VK_NV_scissor_exclusive` extension

Valid Usage (Implicit)

- [](#VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-parameter)VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetExclusiveScissorEnableNV-pExclusiveScissorEnables-parameter)VUID-vkCmdSetExclusiveScissorEnableNV-pExclusiveScissorEnables-parameter  
  `pExclusiveScissorEnables` **must** be a valid pointer to an array of `exclusiveScissorCount` [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html) values
- [](#VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-recording)VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-cmdpool)VUID-vkCmdSetExclusiveScissorEnableNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetExclusiveScissorEnableNV-videocoding)VUID-vkCmdSetExclusiveScissorEnableNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetExclusiveScissorEnableNV-exclusiveScissorCount-arraylength)VUID-vkCmdSetExclusiveScissorEnableNV-exclusiveScissorCount-arraylength  
  `exclusiveScissorCount` **must** be greater than `0`

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

[VK\_NV\_scissor\_exclusive](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_scissor_exclusive.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetExclusiveScissorEnableNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0