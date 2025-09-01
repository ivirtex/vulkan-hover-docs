# vkCmdPushConstants2(3) Manual Page

## Name

vkCmdPushConstants2 - Update the values of push constants



## [](#_c_specification)C Specification

Alternatively, to update push constants, call:

```c++
// Provided by VK_VERSION_1_4
void vkCmdPushConstants2(
    VkCommandBuffer                             commandBuffer,
    const VkPushConstantsInfo*                  pPushConstantsInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_maintenance6
void vkCmdPushConstants2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkPushConstantsInfo*                  pPushConstantsInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which the push constant update will be recorded.
- `pPushConstantsInfo` is a pointer to a [VkPushConstantsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantsInfo.html) structure.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCmdPushConstants2-commandBuffer-parameter)VUID-vkCmdPushConstants2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdPushConstants2-pPushConstantsInfo-parameter)VUID-vkCmdPushConstants2-pPushConstantsInfo-parameter  
  `pPushConstantsInfo` **must** be a valid pointer to a valid [VkPushConstantsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantsInfo.html) structure
- [](#VUID-vkCmdPushConstants2-commandBuffer-recording)VUID-vkCmdPushConstants2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdPushConstants2-commandBuffer-cmdpool)VUID-vkCmdPushConstants2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdPushConstants2-videocoding)VUID-vkCmdPushConstants2-videocoding  
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

vkCmdPushConstants2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPushConstantsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantsInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdPushConstants2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0