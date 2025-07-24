# vkCmdResetEvent(3) Manual Page

## Name

vkCmdResetEvent - Reset an event object to non-signaled state



## [](#_c_specification)C Specification

To set the state of an event to unsignaled from a device, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdResetEvent(
    VkCommandBuffer                             commandBuffer,
    VkEvent                                     event,
    VkPipelineStageFlags                        stageMask);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `event` is the event that will be unsignaled.
- `stageMask` is a bitmask of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) specifying the [source stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) used to determine when the `event` is unsignaled.

## [](#_description)Description

`vkCmdResetEvent` behaves identically to [vkCmdResetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResetEvent2.html).

Valid Usage

- [](#VUID-vkCmdResetEvent-stageMask-04090)VUID-vkCmdResetEvent-stageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`
- [](#VUID-vkCmdResetEvent-stageMask-04091)VUID-vkCmdResetEvent-stageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-vkCmdResetEvent-stageMask-04092)VUID-vkCmdResetEvent-stageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-vkCmdResetEvent-stageMask-04093)VUID-vkCmdResetEvent-stageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-vkCmdResetEvent-stageMask-04094)VUID-vkCmdResetEvent-stageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-vkCmdResetEvent-stageMask-04095)VUID-vkCmdResetEvent-stageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`
- [](#VUID-vkCmdResetEvent-stageMask-04096)VUID-vkCmdResetEvent-stageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`
- [](#VUID-vkCmdResetEvent-stageMask-07318)VUID-vkCmdResetEvent-stageMask-07318  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-vkCmdResetEvent-stageMask-03937)VUID-vkCmdResetEvent-stageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `stageMask` **must** not be `0`
- [](#VUID-vkCmdResetEvent-stageMask-07949)VUID-vkCmdResetEvent-stageMask-07949  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-vkCmdResetEvent-stageMask-10754)VUID-vkCmdResetEvent-stageMask-10754  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-vkCmdResetEvent-stageMask-06458)VUID-vkCmdResetEvent-stageMask-06458  
  Any pipeline stage included in `stageMask` **must** be supported by the capabilities of the queue family specified by the `queueFamilyIndex` member of the [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) structure that was used to create the `VkCommandPool` that `commandBuffer` was allocated from, as specified in the [table of supported pipeline stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-supported)
- [](#VUID-vkCmdResetEvent-stageMask-01153)VUID-vkCmdResetEvent-stageMask-01153  
  `stageMask` **must** not include `VK_PIPELINE_STAGE_HOST_BIT`
- [](#VUID-vkCmdResetEvent-event-03834)VUID-vkCmdResetEvent-event-03834  
  There **must** be an execution dependency between `vkCmdResetEvent` and the execution of any [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents.html) that includes `event` in its `pEvents` parameter
- [](#VUID-vkCmdResetEvent-event-03835)VUID-vkCmdResetEvent-event-03835  
  There **must** be an execution dependency between `vkCmdResetEvent` and the execution of any [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents2.html) that includes `event` in its `pEvents` parameter
- [](#VUID-vkCmdResetEvent-commandBuffer-01157)VUID-vkCmdResetEvent-commandBuffer-01157  
  `commandBuffer`’s current device mask **must** include exactly one physical device

Valid Usage (Implicit)

- [](#VUID-vkCmdResetEvent-commandBuffer-parameter)VUID-vkCmdResetEvent-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdResetEvent-event-parameter)VUID-vkCmdResetEvent-event-parameter  
  `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) handle
- [](#VUID-vkCmdResetEvent-stageMask-parameter)VUID-vkCmdResetEvent-stageMask-parameter  
  `stageMask` **must** be a valid combination of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) values
- [](#VUID-vkCmdResetEvent-commandBuffer-recording)VUID-vkCmdResetEvent-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdResetEvent-commandBuffer-cmdpool)VUID-vkCmdResetEvent-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, compute, decode, or encode operations
- [](#VUID-vkCmdResetEvent-renderpass)VUID-vkCmdResetEvent-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdResetEvent-commonparent)VUID-vkCmdResetEvent-commonparent  
  Both of `commandBuffer`, and `event` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Both

Graphics  
Compute  
Decode  
Encode

Synchronization

Conditional Rendering

vkCmdResetEvent is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html), [VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdResetEvent)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0