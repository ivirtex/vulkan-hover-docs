# vkCmdResetEvent2(3) Manual Page

## Name

vkCmdResetEvent2 - Reset an event object to non-signaled state



## [](#_c_specification)C Specification

To unsignal the event from a device, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdResetEvent2(
    VkCommandBuffer                             commandBuffer,
    VkEvent                                     event,
    VkPipelineStageFlags2                       stageMask);
```

or the equivalent command

```c++
// Provided by VK_KHR_synchronization2
void vkCmdResetEvent2KHR(
    VkCommandBuffer                             commandBuffer,
    VkEvent                                     event,
    VkPipelineStageFlags2                       stageMask);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `event` is the event that will be unsignaled.
- `stageMask` is a [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html) mask of pipeline stages used to determine the first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes).

## [](#_description)Description

When [vkCmdResetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResetEvent2.html) is submitted to a queue, it defines an execution dependency on commands that were submitted before it, and defines an event unsignal operation which resets the event to the unsignaled state.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order). The synchronization scope is limited to operations by `stageMask` or stages that are [logically earlier](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-order) than `stageMask`.

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes only the event unsignal operation.

If `event` is already in the unsignaled state when [vkCmdResetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResetEvent2.html) is executed on the device, then this command has no effect, no event unsignal operation occurs, and no execution dependency is generated.

Valid Usage

- [](#VUID-vkCmdResetEvent2-stageMask-03929)VUID-vkCmdResetEvent2-stageMask-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`
- [](#VUID-vkCmdResetEvent2-stageMask-03930)VUID-vkCmdResetEvent2-stageMask-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-vkCmdResetEvent2-stageMask-03931)VUID-vkCmdResetEvent2-stageMask-03931  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-vkCmdResetEvent2-stageMask-03932)VUID-vkCmdResetEvent2-stageMask-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-vkCmdResetEvent2-stageMask-03933)VUID-vkCmdResetEvent2-stageMask-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-vkCmdResetEvent2-stageMask-03934)VUID-vkCmdResetEvent2-stageMask-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`
- [](#VUID-vkCmdResetEvent2-stageMask-03935)VUID-vkCmdResetEvent2-stageMask-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`
- [](#VUID-vkCmdResetEvent2-stageMask-07316)VUID-vkCmdResetEvent2-stageMask-07316  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-vkCmdResetEvent2-stageMask-04957)VUID-vkCmdResetEvent2-stageMask-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`
- [](#VUID-vkCmdResetEvent2-stageMask-04995)VUID-vkCmdResetEvent2-stageMask-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`
- [](#VUID-vkCmdResetEvent2-stageMask-07946)VUID-vkCmdResetEvent2-stageMask-07946  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-vkCmdResetEvent2-stageMask-10751)VUID-vkCmdResetEvent2-stageMask-10751  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-vkCmdResetEvent2-stageMask-10752)VUID-vkCmdResetEvent2-stageMask-10752  
  If the [`rayTracingMaintenance1`](#features-rayTracingMaintenance1) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`
- [](#VUID-vkCmdResetEvent2-stageMask-10753)VUID-vkCmdResetEvent2-stageMask-10753  
  If the [`micromap`](#features-micromap) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`
- [](#VUID-vkCmdResetEvent2-synchronization2-03829)VUID-vkCmdResetEvent2-synchronization2-03829  
  The [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature **must** be enabled
- [](#VUID-vkCmdResetEvent2-stageMask-03830)VUID-vkCmdResetEvent2-stageMask-03830  
  `stageMask` **must** not include `VK_PIPELINE_STAGE_2_HOST_BIT`
- [](#VUID-vkCmdResetEvent2-event-03831)VUID-vkCmdResetEvent2-event-03831  
  There **must** be an execution dependency between `vkCmdResetEvent2` and the execution of any [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents.html) that includes `event` in its `pEvents` parameter
- [](#VUID-vkCmdResetEvent2-event-03832)VUID-vkCmdResetEvent2-event-03832  
  There **must** be an execution dependency between `vkCmdResetEvent2` and the execution of any [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents2.html) that includes `event` in its `pEvents` parameter
- [](#VUID-vkCmdResetEvent2-commandBuffer-03833)VUID-vkCmdResetEvent2-commandBuffer-03833  
  `commandBuffer`’s current device mask **must** include exactly one physical device

Valid Usage (Implicit)

- [](#VUID-vkCmdResetEvent2-commandBuffer-parameter)VUID-vkCmdResetEvent2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdResetEvent2-event-parameter)VUID-vkCmdResetEvent2-event-parameter  
  `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) handle
- [](#VUID-vkCmdResetEvent2-stageMask-parameter)VUID-vkCmdResetEvent2-stageMask-parameter  
  `stageMask` **must** be a valid combination of [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html) values
- [](#VUID-vkCmdResetEvent2-commandBuffer-recording)VUID-vkCmdResetEvent2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdResetEvent2-commandBuffer-cmdpool)VUID-vkCmdResetEvent2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, compute, decode, or encode operations
- [](#VUID-vkCmdResetEvent2-renderpass)VUID-vkCmdResetEvent2-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdResetEvent2-commonparent)VUID-vkCmdResetEvent2-commonparent  
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

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html), [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdResetEvent2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0