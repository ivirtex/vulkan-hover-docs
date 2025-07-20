# vkCmdDispatchDataGraphARM(3) Manual Page

## Name

vkCmdDispatchDataGraphARM - Dispatch a data graph pipeline within a session



## [](#_c_specification)C Specification

To record a data graph pipeline dispatch, call:

```c++
// Provided by VK_ARM_data_graph
void vkCmdDispatchDataGraphARM(
    VkCommandBuffer                             commandBuffer,
    VkDataGraphPipelineSessionARM               session,
    const VkDataGraphPipelineDispatchInfoARM*   pInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `session` is the [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html) that data graph pipeline being dispatched will use.
- `pInfo` is `NULL` or a pointer to a [VkDataGraphPipelineDispatchInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineDispatchInfoARM.html) structure.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdDispatchDataGraphARM-session-09796)VUID-vkCmdDispatchDataGraphARM-session-09796  
  For each of the session bind point requirements returned by [vkGetDataGraphPipelineSessionBindPointRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelineSessionBindPointRequirementsARM.html) for `session`, [VkDataGraphPipelineSessionBindPointRequirementARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointRequirementARM.html)::`numObjects` objects **must** have been bound to `session`
- [](#VUID-vkCmdDispatchDataGraphARM-None-09797)VUID-vkCmdDispatchDataGraphARM-None-09797  
  For each set *n* that is statically used by a bound data graph pipeline, a descriptor set **must** have been bound to *n* at the same pipeline bind point, with a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) that is compatible for set *n*, with the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) used to create the current [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), as described in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-compatibility)
- [](#VUID-vkCmdDispatchDataGraphARM-None-09935)VUID-vkCmdDispatchDataGraphARM-None-09935  
  Descriptors in each bound descriptor set, specified via [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets.html), **must** be valid as described by [descriptor validity](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptor-validity) if they are statically used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point used by this command and the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) was not created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdDispatchDataGraphARM-None-09936)VUID-vkCmdDispatchDataGraphARM-None-09936  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point were specified via [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets.html), the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) **must** have been created without `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdDispatchDataGraphARM-None-09937)VUID-vkCmdDispatchDataGraphARM-None-09937  
  Descriptors in bound descriptor buffers, specified via [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), **must** be valid if they are dynamically used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point used by this command and the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) was created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdDispatchDataGraphARM-None-09938)VUID-vkCmdDispatchDataGraphARM-None-09938  
  If the descriptors used by the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) bound to the pipeline bind point were specified via [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), the bound [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) **must** have been created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdDispatchDataGraphARM-None-09939)VUID-vkCmdDispatchDataGraphARM-None-09939  
  If a descriptor is dynamically used with a [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) created with `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, the descriptor memory **must** be resident
- [](#VUID-vkCmdDispatchDataGraphARM-None-09799)VUID-vkCmdDispatchDataGraphARM-None-09799  
  A valid data graph pipeline **must** be bound to the `VK_PIPELINE_BIND_POINT_DATA_GRAPH_ARM` pipeline bind point used by this command
- [](#VUID-vkCmdDispatchDataGraphARM-pDescription-09930)VUID-vkCmdDispatchDataGraphARM-pDescription-09930  
  If a `VK_DESCRIPTOR_TYPE_TENSOR_ARM` descriptor is accessed as a result of this command, then the underlying [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) object **must** have been created with a [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)::`pDescription` whose `usage` member contained `VK_TENSOR_USAGE_DATA_GRAPH_BIT_ARM`
- [](#VUID-vkCmdDispatchDataGraphARM-commandBuffer-09940)VUID-vkCmdDispatchDataGraphARM-commandBuffer-09940  
  If `commandBuffer` was allocated from a pool that was created with a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure in the `pNext` chain of [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) that included a foreign data graph processing engine in its `pProcessingEngines` member, then all `VK_DESCRIPTOR_TYPE_TENSOR_ARM` descriptors accessed as a result of this command **must** be tied to [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) objects that have been bound to memory created from external handle types reported as supported in a [VkQueueFamilyDataGraphProcessingEnginePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphProcessingEnginePropertiesARM.html)::`foreignMemoryHandleTypes` structure via [vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM.html) with a `queueFamilyIndex` matching the one the command pool was created for, for all the foreign data graph processing engines that were part of the [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) used to create the command pool
- [](#VUID-vkCmdDispatchDataGraphARM-commandBuffer-09800)VUID-vkCmdDispatchDataGraphARM-commandBuffer-09800  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-protectedNoFault) is not supported, any resource accessed by bound data graph pipelines **must** not be a protected resource
- [](#VUID-vkCmdDispatchDataGraphARM-commandBuffer-09801)VUID-vkCmdDispatchDataGraphARM-commandBuffer-09801  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-protectedNoFault) is not supported, any resource written to by the `VkPipeline` object bound to the bind point used by this command **must** not be an unprotected resource
- [](#VUID-vkCmdDispatchDataGraphARM-commandBuffer-09941)VUID-vkCmdDispatchDataGraphARM-commandBuffer-09941  
  All the operations used by the bound data graph pipeline **must** be supported on the queue family for which the command pool out of which `commandBuffer` was allocated, as reported by [vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM.html)

Valid Usage (Implicit)

- [](#VUID-vkCmdDispatchDataGraphARM-commandBuffer-parameter)VUID-vkCmdDispatchDataGraphARM-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdDispatchDataGraphARM-session-parameter)VUID-vkCmdDispatchDataGraphARM-session-parameter  
  `session` **must** be a valid [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html) handle
- [](#VUID-vkCmdDispatchDataGraphARM-pInfo-parameter)VUID-vkCmdDispatchDataGraphARM-pInfo-parameter  
  If `pInfo` is not `NULL`, `pInfo` **must** be a valid pointer to a valid [VkDataGraphPipelineDispatchInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineDispatchInfoARM.html) structure
- [](#VUID-vkCmdDispatchDataGraphARM-commandBuffer-recording)VUID-vkCmdDispatchDataGraphARM-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdDispatchDataGraphARM-commandBuffer-cmdpool)VUID-vkCmdDispatchDataGraphARM-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support data\_graph operations
- [](#VUID-vkCmdDispatchDataGraphARM-renderpass)VUID-vkCmdDispatchDataGraphARM-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdDispatchDataGraphARM-videocoding)VUID-vkCmdDispatchDataGraphARM-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdDispatchDataGraphARM-commonparent)VUID-vkCmdDispatchDataGraphARM-commonparent  
  Both of `commandBuffer`, and `session` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Data\_Graph

Action

Conditional Rendering

vkCmdDispatchDataGraphARM is affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDataGraphPipelineDispatchInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineDispatchInfoARM.html), [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdDispatchDataGraphARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0