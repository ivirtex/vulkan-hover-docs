# vkCmdPreprocessGeneratedCommandsNV(3) Manual Page

## Name

vkCmdPreprocessGeneratedCommandsNV - Performs preprocessing for generated commands



## [](#_c_specification)C Specification

Commands **can** be preprocessed prior execution using the following command:

```c++
// Provided by VK_NV_device_generated_commands
void vkCmdPreprocessGeneratedCommandsNV(
    VkCommandBuffer                             commandBuffer,
    const VkGeneratedCommandsInfoNV*            pGeneratedCommandsInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer which does the preprocessing.
- `pGeneratedCommandsInfo` is a pointer to a [VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoNV.html) structure containing parameters affecting the preprocessing step.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-02974)VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-02974  
  `commandBuffer` **must** not be a protected command buffer
- [](#VUID-vkCmdPreprocessGeneratedCommandsNV-pGeneratedCommandsInfo-02927)VUID-vkCmdPreprocessGeneratedCommandsNV-pGeneratedCommandsInfo-02927  
  `pGeneratedCommandsInfo`\`s `indirectCommandsLayout` **must** have been created with the `VK_INDIRECT_COMMANDS_LAYOUT_USAGE_EXPLICIT_PREPROCESS_BIT_NV` bit set
- [](#VUID-vkCmdPreprocessGeneratedCommandsNV-deviceGeneratedCommands-02928)VUID-vkCmdPreprocessGeneratedCommandsNV-deviceGeneratedCommands-02928  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV`::`deviceGeneratedCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedCommandsNV) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-parameter)VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdPreprocessGeneratedCommandsNV-pGeneratedCommandsInfo-parameter)VUID-vkCmdPreprocessGeneratedCommandsNV-pGeneratedCommandsInfo-parameter  
  `pGeneratedCommandsInfo` **must** be a valid pointer to a valid [VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoNV.html) structure
- [](#VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-recording)VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-cmdpool)VUID-vkCmdPreprocessGeneratedCommandsNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdPreprocessGeneratedCommandsNV-renderpass)VUID-vkCmdPreprocessGeneratedCommandsNV-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdPreprocessGeneratedCommandsNV-videocoding)VUID-vkCmdPreprocessGeneratedCommandsNV-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Graphics  
Compute

Action

Conditional Rendering

vkCmdPreprocessGeneratedCommandsNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdPreprocessGeneratedCommandsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0