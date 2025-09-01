# vkCmdPreprocessGeneratedCommandsEXT(3) Manual Page

## Name

vkCmdPreprocessGeneratedCommandsEXT - Performs preprocessing for generated commands



## [](#_c_specification)C Specification

Commands **can** be preprocessed prior execution using the following command:

```c++
// Provided by VK_EXT_device_generated_commands
void vkCmdPreprocessGeneratedCommandsEXT(
    VkCommandBuffer                             commandBuffer,
    const VkGeneratedCommandsInfoEXT*           pGeneratedCommandsInfo,
    VkCommandBuffer                             stateCommandBuffer);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer which does the preprocessing.
- `pGeneratedCommandsInfo` is a pointer to a [VkGeneratedCommandsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoEXT.html) structure containing parameters affecting the preprocessing step.
- `stateCommandBuffer` is a command buffer from which to snapshot current states affecting the preprocessing step. When a graphics command action token is used, graphics state is snapshotted. When a compute action command token is used, compute state is snapshotted. When a ray tracing action command token is used, ray tracing state is snapshotted. It can be deleted at any time after this command has been recorded.

## [](#_description)Description

Note

`stateCommandBuffer` access is not synchronized by the driver, meaning that this command buffer **must** not be modified between threads in an unsafe manner.

Valid Usage

- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-commandBuffer-11081)VUID-vkCmdPreprocessGeneratedCommandsEXT-commandBuffer-11081  
  `commandBuffer` **must** not be a protected command buffer
- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-pGeneratedCommandsInfo-11082)VUID-vkCmdPreprocessGeneratedCommandsEXT-pGeneratedCommandsInfo-11082  
  `pGeneratedCommandsInfo`’s `indirectCommandsLayout` **must** have been created with the `VK_INDIRECT_COMMANDS_LAYOUT_USAGE_EXPLICIT_PREPROCESS_BIT_EXT` bit set
- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-indirectCommandsLayout-11084)VUID-vkCmdPreprocessGeneratedCommandsEXT-indirectCommandsLayout-11084  
  If the token sequence of the passed [VkGeneratedCommandsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoEXT.html)::`indirectCommandsLayout` contains a `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT` token, the initial shader state of [VkGeneratedCommandsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoEXT.html)::`indirectExecutionSet` **must** be bound on `stateCommandBuffer`
- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-stateCommandBuffer-11138)VUID-vkCmdPreprocessGeneratedCommandsEXT-stateCommandBuffer-11138  
  `stateCommandBuffer` **must** be in the recording state
- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-deviceGeneratedCommands-11087)VUID-vkCmdPreprocessGeneratedCommandsEXT-deviceGeneratedCommands-11087  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsFeaturesEXT`::`deviceGeneratedCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedCommands) feature **must** be enabled
- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-supportedIndirectCommandsShaderStages-11088)VUID-vkCmdPreprocessGeneratedCommandsEXT-supportedIndirectCommandsShaderStages-11088  
  Only stages specified in [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-supportedIndirectCommandsShaderStages)[VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`supportedIndirectCommandsShaderStages` **can** be set in `pGeneratedCommandsInfo->shaderStages`

Valid Usage (Implicit)

- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-commandBuffer-parameter)VUID-vkCmdPreprocessGeneratedCommandsEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-pGeneratedCommandsInfo-parameter)VUID-vkCmdPreprocessGeneratedCommandsEXT-pGeneratedCommandsInfo-parameter  
  `pGeneratedCommandsInfo` **must** be a valid pointer to a valid [VkGeneratedCommandsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoEXT.html) structure
- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-stateCommandBuffer-parameter)VUID-vkCmdPreprocessGeneratedCommandsEXT-stateCommandBuffer-parameter  
  `stateCommandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-commandBuffer-recording)VUID-vkCmdPreprocessGeneratedCommandsEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-commandBuffer-cmdpool)VUID-vkCmdPreprocessGeneratedCommandsEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-renderpass)VUID-vkCmdPreprocessGeneratedCommandsEXT-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-videocoding)VUID-vkCmdPreprocessGeneratedCommandsEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-bufferlevel)VUID-vkCmdPreprocessGeneratedCommandsEXT-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`
- [](#VUID-vkCmdPreprocessGeneratedCommandsEXT-commonparent)VUID-vkCmdPreprocessGeneratedCommandsEXT-commonparent  
  Both of `commandBuffer`, and `stateCommandBuffer` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to `stateCommandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary

Outside

Outside

Graphics  
Compute

Action

Conditional Rendering

vkCmdPreprocessGeneratedCommandsEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkGeneratedCommandsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdPreprocessGeneratedCommandsEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0