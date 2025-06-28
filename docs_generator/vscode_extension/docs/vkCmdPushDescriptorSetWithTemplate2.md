# vkCmdPushDescriptorSetWithTemplate2(3) Manual Page

## Name

vkCmdPushDescriptorSetWithTemplate2 - Pushes descriptor updates into a command buffer using a descriptor update template



## [](#_c_specification)C Specification

Alternatively, to use a descriptor update template to specify the push descriptors to update, call:

```c++
// Provided by VK_VERSION_1_4
void vkCmdPushDescriptorSetWithTemplate2(
    VkCommandBuffer                             commandBuffer,
    const VkPushDescriptorSetWithTemplateInfo*  pPushDescriptorSetWithTemplateInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_maintenance6 with VK_KHR_push_descriptor
void vkCmdPushDescriptorSetWithTemplate2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkPushDescriptorSetWithTemplateInfo*  pPushDescriptorSetWithTemplateInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that the descriptors will be recorded in.
- `pPushDescriptorSetWithTemplateInfo` is a pointer to a `VkPushDescriptorSetWithTemplateInfo` structure.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCmdPushDescriptorSetWithTemplate2-commandBuffer-parameter)VUID-vkCmdPushDescriptorSetWithTemplate2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdPushDescriptorSetWithTemplate2-pPushDescriptorSetWithTemplateInfo-parameter)VUID-vkCmdPushDescriptorSetWithTemplate2-pPushDescriptorSetWithTemplateInfo-parameter  
  `pPushDescriptorSetWithTemplateInfo` **must** be a valid pointer to a valid [VkPushDescriptorSetWithTemplateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushDescriptorSetWithTemplateInfo.html) structure
- [](#VUID-vkCmdPushDescriptorSetWithTemplate2-commandBuffer-recording)VUID-vkCmdPushDescriptorSetWithTemplate2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdPushDescriptorSetWithTemplate2-commandBuffer-cmdpool)VUID-vkCmdPushDescriptorSetWithTemplate2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdPushDescriptorSetWithTemplate2-videocoding)VUID-vkCmdPushDescriptorSetWithTemplate2-videocoding  
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

## [](#_see_also)See Also

[VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPushDescriptorSetWithTemplateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushDescriptorSetWithTemplateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdPushDescriptorSetWithTemplate2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0