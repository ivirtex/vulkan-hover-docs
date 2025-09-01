# vkCmdBindDescriptorSets2(3) Manual Page

## Name

vkCmdBindDescriptorSets2 - Binds descriptor sets to a command buffer



## [](#_c_specification)C Specification

Alternatively, to bind one or more descriptor sets to a command buffer, call:

```c++
// Provided by VK_VERSION_1_4
void vkCmdBindDescriptorSets2(
    VkCommandBuffer                             commandBuffer,
    const VkBindDescriptorSetsInfo*             pBindDescriptorSetsInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_maintenance6
void vkCmdBindDescriptorSets2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkBindDescriptorSetsInfo*             pBindDescriptorSetsInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that the descriptor sets will be bound to.
- `pBindDescriptorSetsInfo` is a pointer to a `VkBindDescriptorSetsInfo` structure.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdBindDescriptorSets2-pBindDescriptorSetsInfo-09467)VUID-vkCmdBindDescriptorSets2-pBindDescriptorSetsInfo-09467  
  Each bit in `pBindDescriptorSetsInfo->stageFlags` **must** be a stage supported by the `commandBuffer`’s parent `VkCommandPool`’s queue family

Valid Usage (Implicit)

- [](#VUID-vkCmdBindDescriptorSets2-commandBuffer-parameter)VUID-vkCmdBindDescriptorSets2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindDescriptorSets2-pBindDescriptorSetsInfo-parameter)VUID-vkCmdBindDescriptorSets2-pBindDescriptorSetsInfo-parameter  
  `pBindDescriptorSetsInfo` **must** be a valid pointer to a valid [VkBindDescriptorSetsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDescriptorSetsInfo.html) structure
- [](#VUID-vkCmdBindDescriptorSets2-commandBuffer-recording)VUID-vkCmdBindDescriptorSets2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindDescriptorSets2-commandBuffer-cmdpool)VUID-vkCmdBindDescriptorSets2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdBindDescriptorSets2-videocoding)VUID-vkCmdBindDescriptorSets2-videocoding  
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
Compute

State

Conditional Rendering

vkCmdBindDescriptorSets2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBindDescriptorSetsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDescriptorSetsInfo.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindDescriptorSets2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0