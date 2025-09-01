# vkCmdPushDescriptorSet2(3) Manual Page

## Name

vkCmdPushDescriptorSet2 - Pushes descriptor updates into a command buffer



## [](#_c_specification)C Specification

Alternatively, to push descriptor updates into a command buffer, call:

```c++
// Provided by VK_VERSION_1_4
void vkCmdPushDescriptorSet2(
    VkCommandBuffer                             commandBuffer,
    const VkPushDescriptorSetInfo*              pPushDescriptorSetInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_maintenance6 with VK_KHR_push_descriptor
void vkCmdPushDescriptorSet2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkPushDescriptorSetInfo*              pPushDescriptorSetInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that the descriptors will be recorded in.
- `pPushDescriptorSetInfo` is a pointer to a `VkPushDescriptorSetInfo` structure.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdPushDescriptorSet2-pPushDescriptorSetInfo-09468)VUID-vkCmdPushDescriptorSet2-pPushDescriptorSetInfo-09468  
  Each bit in `pPushDescriptorSetInfo->stageFlags` **must** be a stage supported by the `commandBuffer`’s parent `VkCommandPool`’s queue family
- [](#VUID-vkCmdPushDescriptorSet2-None-10357)VUID-vkCmdPushDescriptorSet2-None-10357  
  If the [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html) extension is not enabled, [`pushDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pushDescriptor) **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdPushDescriptorSet2-commandBuffer-parameter)VUID-vkCmdPushDescriptorSet2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdPushDescriptorSet2-pPushDescriptorSetInfo-parameter)VUID-vkCmdPushDescriptorSet2-pPushDescriptorSetInfo-parameter  
  `pPushDescriptorSetInfo` **must** be a valid pointer to a valid [VkPushDescriptorSetInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushDescriptorSetInfo.html) structure
- [](#VUID-vkCmdPushDescriptorSet2-commandBuffer-recording)VUID-vkCmdPushDescriptorSet2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdPushDescriptorSet2-commandBuffer-cmdpool)VUID-vkCmdPushDescriptorSet2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdPushDescriptorSet2-videocoding)VUID-vkCmdPushDescriptorSet2-videocoding  
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

vkCmdPushDescriptorSet2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPushDescriptorSetInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushDescriptorSetInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdPushDescriptorSet2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0