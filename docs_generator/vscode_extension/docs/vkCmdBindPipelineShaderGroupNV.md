# vkCmdBindPipelineShaderGroupNV(3) Manual Page

## Name

vkCmdBindPipelineShaderGroupNV - Bind a pipeline object



## [](#_c_specification)C Specification

For pipelines that were created with the support of multiple shader groups (see [Graphics Pipeline Shader Groups](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#graphics-shadergroups)), the regular `vkCmdBindPipeline` command will bind Shader Group `0`. To explicitly bind a shader group use:

```c++
// Provided by VK_NV_device_generated_commands
void vkCmdBindPipelineShaderGroupNV(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipeline                                  pipeline,
    uint32_t                                    groupIndex);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that the pipeline will be bound to.
- `pipelineBindPoint` is a [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value specifying the bind point to which the pipeline will be bound.
- `pipeline` is the pipeline to be bound.
- `groupIndex` is the shader group to be bound.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdBindPipelineShaderGroupNV-groupIndex-02893)VUID-vkCmdBindPipelineShaderGroupNV-groupIndex-02893  
  `groupIndex` **must** be `0` or less than the effective [VkGraphicsPipelineShaderGroupsCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineShaderGroupsCreateInfoNV.html)::`groupCount` including the referenced pipelines
- [](#VUID-vkCmdBindPipelineShaderGroupNV-pipelineBindPoint-02894)VUID-vkCmdBindPipelineShaderGroupNV-pipelineBindPoint-02894  
  The `pipelineBindPoint` **must** be `VK_PIPELINE_BIND_POINT_GRAPHICS`
- [](#VUID-vkCmdBindPipelineShaderGroupNV-groupIndex-02895)VUID-vkCmdBindPipelineShaderGroupNV-groupIndex-02895  
  The same restrictions as [vkCmdBindPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindPipeline.html) apply as if the bound pipeline was created only with the Shader Group from the `groupIndex` information
- [](#VUID-vkCmdBindPipelineShaderGroupNV-deviceGeneratedCommands-02896)VUID-vkCmdBindPipelineShaderGroupNV-deviceGeneratedCommands-02896  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV`::`deviceGeneratedCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedCommandsNV) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-parameter)VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindPipelineShaderGroupNV-pipelineBindPoint-parameter)VUID-vkCmdBindPipelineShaderGroupNV-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value
- [](#VUID-vkCmdBindPipelineShaderGroupNV-pipeline-parameter)VUID-vkCmdBindPipelineShaderGroupNV-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-recording)VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-cmdpool)VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdBindPipelineShaderGroupNV-videocoding)VUID-vkCmdBindPipelineShaderGroupNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBindPipelineShaderGroupNV-commonparent)VUID-vkCmdBindPipelineShaderGroupNV-commonparent  
  Both of `commandBuffer`, and `pipeline` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

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

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindPipelineShaderGroupNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0