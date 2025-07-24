# vkCmdUpdatePipelineIndirectBufferNV(3) Manual Page

## Name

vkCmdUpdatePipelineIndirectBufferNV - Update the indirect compute pipeline's metadata



## [](#_c_specification)C Specification

To save a compute pipeline’s metadata at a device address call:

```c++
// Provided by VK_NV_device_generated_commands_compute
void vkCmdUpdatePipelineIndirectBufferNV(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipeline                                  pipeline);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pipelineBindPoint` is a [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value specifying the type of pipeline whose metadata will be saved.
- `pipeline` is the pipeline whose metadata will be saved.

## [](#_description)Description

`vkCmdUpdatePipelineIndirectBufferNV` is only allowed outside of a render pass. This command is treated as a “transfer” operation for the purposes of synchronization barriers. The writes to the address **must** be synchronized using stages `VK_PIPELINE_STAGE_2_COPY_BIT` and `VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_NV` and with access masks `VK_ACCESS_MEMORY_WRITE_BIT` and `VK_ACCESS_COMMAND_PREPROCESS_READ_BIT_NV` respectively before using the results in preprocessing.

Valid Usage

- [](#VUID-vkCmdUpdatePipelineIndirectBufferNV-pipelineBindPoint-09018)VUID-vkCmdUpdatePipelineIndirectBufferNV-pipelineBindPoint-09018  
  `pipelineBindPoint` **must** be `VK_PIPELINE_BIND_POINT_COMPUTE`
- [](#VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-09019)VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-09019  
  `pipeline` **must** have been created with `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV` flag set
- [](#VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-09020)VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-09020  
  `pipeline` **must** have been created with [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineIndirectBufferInfoNV.html) structure specifying a valid address where its metadata will be saved
- [](#VUID-vkCmdUpdatePipelineIndirectBufferNV-deviceGeneratedComputePipelines-09021)VUID-vkCmdUpdatePipelineIndirectBufferNV-deviceGeneratedComputePipelines-09021  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`::`deviceGeneratedComputePipelines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedComputePipelines) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-parameter)VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdUpdatePipelineIndirectBufferNV-pipelineBindPoint-parameter)VUID-vkCmdUpdatePipelineIndirectBufferNV-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value
- [](#VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-parameter)VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-recording)VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-cmdpool)VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support transfer, graphics, or compute operations
- [](#VUID-vkCmdUpdatePipelineIndirectBufferNV-renderpass)VUID-vkCmdUpdatePipelineIndirectBufferNV-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdUpdatePipelineIndirectBufferNV-videocoding)VUID-vkCmdUpdatePipelineIndirectBufferNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdUpdatePipelineIndirectBufferNV-commonparent)VUID-vkCmdUpdatePipelineIndirectBufferNV-commonparent  
  Both of `commandBuffer`, and `pipeline` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Transfer  
Graphics  
Compute

Action

Conditional Rendering

vkCmdUpdatePipelineIndirectBufferNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands\_compute](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands_compute.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdUpdatePipelineIndirectBufferNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0