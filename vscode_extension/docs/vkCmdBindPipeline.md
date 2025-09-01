# vkCmdBindPipeline(3) Manual Page

## Name

vkCmdBindPipeline - Bind a pipeline object to a command buffer



## [](#_c_specification)C Specification

Once a pipeline has been created, it **can** be bound to the command buffer using the command:

```c++
// Provided by VK_VERSION_1_0
void vkCmdBindPipeline(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipeline                                  pipeline);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that the pipeline will be bound to.
- `pipelineBindPoint` is a [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value specifying to which bind point the pipeline is bound. Binding one does not disturb the others.
- `pipeline` is the pipeline to be bound.

## [](#_description)Description

Once bound, a pipeline binding affects subsequent commands that interact with the given pipeline type in the command buffer until a different pipeline of the same type is bound to the bind point, or until the pipeline bind point is disturbed by binding a [shader object](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects) as described in [Interaction with Pipelines](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects-pipeline-interaction). Commands that do not interact with the [given pipeline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-binding) type **must** not be affected by the pipeline state.

Valid Usage

- [](#VUID-vkCmdBindPipeline-pipelineBindPoint-00777)VUID-vkCmdBindPipeline-pipelineBindPoint-00777  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_COMPUTE`, the `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdBindPipeline-pipelineBindPoint-00778)VUID-vkCmdBindPipeline-pipelineBindPoint-00778  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_GRAPHICS`, the `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBindPipeline-pipelineBindPoint-00779)VUID-vkCmdBindPipeline-pipelineBindPoint-00779  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_COMPUTE`, `pipeline` **must** be a compute pipeline
- [](#VUID-vkCmdBindPipeline-pipelineBindPoint-00780)VUID-vkCmdBindPipeline-pipelineBindPoint-00780  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_GRAPHICS`, `pipeline` **must** be a graphics pipeline
- [](#VUID-vkCmdBindPipeline-pipeline-00781)VUID-vkCmdBindPipeline-pipeline-00781  
  If the [`variableMultisampleRate`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-variableMultisampleRate) feature is not supported, `pipeline` is a graphics pipeline, the current subpass [uses no attachments](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-noattachments), and this is not the first call to this function with a graphics pipeline after transitioning to the current subpass, then the sample count specified by this pipeline **must** match that set in the previous pipeline
- [](#VUID-vkCmdBindPipeline-variableSampleLocations-01525)VUID-vkCmdBindPipeline-variableSampleLocations-01525  
  If [VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html)::`variableSampleLocations` is `VK_FALSE`, and `pipeline` is a graphics pipeline created with a `renderPass` that is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and with a [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html) structure having its `sampleLocationsEnable` member set to `VK_TRUE` but without `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT` enabled then the current render pass instance **must** have been begun by specifying a [VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html) structure whose `pPostSubpassSampleLocations` member contains an element with a `subpassIndex` matching the current subpass index and the `sampleLocationsInfo` member of that element **must** match the `sampleLocationsInfo` specified in [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html) when the pipeline was created
- [](#VUID-vkCmdBindPipeline-None-02323)VUID-vkCmdBindPipeline-None-02323  
  This command **must** not be recorded when transform feedback is active
- [](#VUID-vkCmdBindPipeline-pipelineBindPoint-02391)VUID-vkCmdBindPipeline-pipelineBindPoint-02391  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR`, the `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdBindPipeline-pipelineBindPoint-02392)VUID-vkCmdBindPipeline-pipelineBindPoint-02392  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR`, `pipeline` **must** be a ray tracing pipeline
- [](#VUID-vkCmdBindPipeline-pipelineBindPoint-06721)VUID-vkCmdBindPipeline-pipelineBindPoint-06721  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR`, `commandBuffer` **must** not be a protected command buffer
- [](#VUID-vkCmdBindPipeline-pipelineProtectedAccess-07408)VUID-vkCmdBindPipeline-pipelineProtectedAccess-07408  
  If the [`pipelineProtectedAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineProtectedAccess) feature is enabled, and `commandBuffer` is a protected command buffer, `pipeline` **must** have been created without `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT`
- [](#VUID-vkCmdBindPipeline-pipelineProtectedAccess-07409)VUID-vkCmdBindPipeline-pipelineProtectedAccess-07409  
  If the [`pipelineProtectedAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineProtectedAccess) feature is enabled, and `commandBuffer` is not a protected command buffer, `pipeline` **must** have been created without `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT`
- [](#VUID-vkCmdBindPipeline-pipeline-03382)VUID-vkCmdBindPipeline-pipeline-03382  
  `pipeline` **must** not have been created with `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` set
- [](#VUID-vkCmdBindPipeline-commandBuffer-04808)VUID-vkCmdBindPipeline-commandBuffer-04808  
  If `commandBuffer` is a secondary command buffer with [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D` enabled and `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_GRAPHICS`, then the `pipeline` **must** have been created with `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` or `VK_DYNAMIC_STATE_VIEWPORT`, and `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT` or `VK_DYNAMIC_STATE_SCISSOR` enabled
- [](#VUID-vkCmdBindPipeline-commandBuffer-04809)VUID-vkCmdBindPipeline-commandBuffer-04809  
  If `commandBuffer` is a secondary command buffer with [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D` enabled and `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_GRAPHICS` and `pipeline` was created with [VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html) structure and its `discardRectangleCount` member is not `0`, or the pipeline was created with `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT` enabled, then the pipeline **must** have been created with `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT` enabled
- [](#VUID-vkCmdBindPipeline-pipelineBindPoint-04881)VUID-vkCmdBindPipeline-pipelineBindPoint-04881  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_GRAPHICS` and the [`provokingVertexModePerPipeline`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-provokingVertexModePerPipeline) limit is `VK_FALSE`, then pipeline’s [VkPipelineRasterizationProvokingVertexStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationProvokingVertexStateCreateInfoEXT.html)::`provokingVertexMode` **must** be the same as that of any other pipelines previously bound to this bind point within the current render pass instance, including any pipeline already bound when beginning the render pass instance
- [](#VUID-vkCmdBindPipeline-pipelineBindPoint-04949)VUID-vkCmdBindPipeline-pipelineBindPoint-04949  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI`, the `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdBindPipeline-pipelineBindPoint-04950)VUID-vkCmdBindPipeline-pipelineBindPoint-04950  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI`, `pipeline` **must** be a subpass shading pipeline
- [](#VUID-vkCmdBindPipeline-pipelineBindPoint-09910)VUID-vkCmdBindPipeline-pipelineBindPoint-09910  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_DATA_GRAPH_ARM`, the `VkCommandPool` that `commandBuffer` was allocated from **must** have been created for a queue family that supports `VK_QUEUE_DATA_GRAPH_BIT_ARM`
- [](#VUID-vkCmdBindPipeline-pipelineBindPoint-09911)VUID-vkCmdBindPipeline-pipelineBindPoint-09911  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_DATA_GRAPH_ARM`, `pipeline` **must** be a [data graph pipeline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#graphs-pipelines)
- [](#VUID-vkCmdBindPipeline-pipeline-09912)VUID-vkCmdBindPipeline-pipeline-09912  
  If `pipeline` is a [data graph pipeline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#graphs-pipelines) and the [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html) structure used to create it had a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure in its `pNext` chain that specified any foreign data processing engines, then the command pool from which `commandBuffer` was allocated **must** have been created with a [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) structure that had a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure specifying a superset of the foreign data graph processing engines specified at pipeline creation time in its `pNext` chain
- [](#VUID-vkCmdBindPipeline-pipeline-09913)VUID-vkCmdBindPipeline-pipeline-09913  
  If `pipeline` is a [data graph pipeline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#graphs-pipelines) and the [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html) structure used to create it did not have a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure in its `pNext` chain, then the command pool from which `commandBuffer` was allocated **must** not have been created with a [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) that had a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure in its `pNext` chain

Valid Usage (Implicit)

- [](#VUID-vkCmdBindPipeline-commandBuffer-parameter)VUID-vkCmdBindPipeline-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindPipeline-pipelineBindPoint-parameter)VUID-vkCmdBindPipeline-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value
- [](#VUID-vkCmdBindPipeline-pipeline-parameter)VUID-vkCmdBindPipeline-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-vkCmdBindPipeline-commandBuffer-recording)VUID-vkCmdBindPipeline-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindPipeline-commandBuffer-cmdpool)VUID-vkCmdBindPipeline-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, compute, or data\_graph operations
- [](#VUID-vkCmdBindPipeline-videocoding)VUID-vkCmdBindPipeline-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBindPipeline-commonparent)VUID-vkCmdBindPipeline-commonparent  
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
Data\_Graph

State

Conditional Rendering

vkCmdBindPipeline is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindPipeline).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0