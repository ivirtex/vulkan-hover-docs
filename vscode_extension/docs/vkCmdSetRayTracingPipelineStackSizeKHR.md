# vkCmdSetRayTracingPipelineStackSizeKHR(3) Manual Page

## Name

vkCmdSetRayTracingPipelineStackSizeKHR - Set the stack size dynamically for a ray tracing pipeline



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the stack size for a ray tracing pipeline, call:

```c++
// Provided by VK_KHR_ray_tracing_pipeline
void vkCmdSetRayTracingPipelineStackSizeKHR(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    pipelineStackSize);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pipelineStackSize` is the stack size to use for subsequent ray tracing trace commands.

## [](#_description)Description

This command sets the stack size for subsequent ray tracing commands when the ray tracing pipeline is created with `VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, the stack size is computed as described in [Ray Tracing Pipeline Stack](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-tracing-pipeline-stack).

Valid Usage

- [](#VUID-vkCmdSetRayTracingPipelineStackSizeKHR-pipelineStackSize-03610)VUID-vkCmdSetRayTracingPipelineStackSizeKHR-pipelineStackSize-03610  
  `pipelineStackSize` **must** be large enough for any dynamic execution through the shaders in the ray tracing pipeline used by a subsequent trace call

Valid Usage (Implicit)

- [](#VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-parameter)VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-recording)VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-cmdpool)VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT operations
- [](#VUID-vkCmdSetRayTracingPipelineStackSizeKHR-renderpass)VUID-vkCmdSetRayTracingPipelineStackSizeKHR-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdSetRayTracingPipelineStackSizeKHR-videocoding)VUID-vkCmdSetRayTracingPipelineStackSizeKHR-videocoding  
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

VK\_QUEUE\_COMPUTE\_BIT

State

Conditional Rendering

vkCmdSetRayTracingPipelineStackSizeKHR is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetRayTracingPipelineStackSizeKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0