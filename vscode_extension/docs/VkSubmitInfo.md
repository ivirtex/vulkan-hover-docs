# VkSubmitInfo(3) Manual Page

## Name

VkSubmitInfo - Structure specifying a queue submit operation



## [](#_c_specification)C Specification

The `VkSubmitInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkSubmitInfo {
    VkStructureType                sType;
    const void*                    pNext;
    uint32_t                       waitSemaphoreCount;
    const VkSemaphore*             pWaitSemaphores;
    const VkPipelineStageFlags*    pWaitDstStageMask;
    uint32_t                       commandBufferCount;
    const VkCommandBuffer*         pCommandBuffers;
    uint32_t                       signalSemaphoreCount;
    const VkSemaphore*             pSignalSemaphores;
} VkSubmitInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `waitSemaphoreCount` is the number of semaphores upon which to wait before executing the command buffers for the batch.
- `pWaitSemaphores` is a pointer to an array of [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handles upon which to wait before the command buffers for this batch begin execution. If semaphores to wait on are provided, they define a [semaphore wait operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-waiting).
- `pWaitDstStageMask` is a pointer to an array of pipeline stages at which each corresponding semaphore wait will occur.
- `commandBufferCount` is the number of command buffers to execute in the batch.
- `pCommandBuffers` is a pointer to an array of [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handles to execute in the batch.
- `signalSemaphoreCount` is the number of semaphores to be signaled once the commands specified in `pCommandBuffers` have completed execution.
- `pSignalSemaphores` is a pointer to an array of [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handles which will be signaled when the command buffers for this batch have completed execution. If semaphores to be signaled are provided, they define a [semaphore signal operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-signaling).

## [](#_description)Description

The order that command buffers appear in `pCommandBuffers` is used to determine [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order), and thus all the [implicit ordering guarantees](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-implicit) that respect it. Other than these implicit ordering guarantees and any [explicit synchronization primitives](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization), these command buffers **may** overlap or otherwise execute out of order.

Valid Usage

- [](#VUID-VkSubmitInfo-pWaitDstStageMask-04090)VUID-VkSubmitInfo-pWaitDstStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `pWaitDstStageMask` **must** not contain `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`
- [](#VUID-VkSubmitInfo-pWaitDstStageMask-04091)VUID-VkSubmitInfo-pWaitDstStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `pWaitDstStageMask` **must** not contain `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-VkSubmitInfo-pWaitDstStageMask-04092)VUID-VkSubmitInfo-pWaitDstStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `pWaitDstStageMask` **must** not contain `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-VkSubmitInfo-pWaitDstStageMask-04093)VUID-VkSubmitInfo-pWaitDstStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `pWaitDstStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-VkSubmitInfo-pWaitDstStageMask-04094)VUID-VkSubmitInfo-pWaitDstStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `pWaitDstStageMask` **must** not contain `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-VkSubmitInfo-pWaitDstStageMask-04095)VUID-VkSubmitInfo-pWaitDstStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `pWaitDstStageMask` **must** not contain `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`
- [](#VUID-VkSubmitInfo-pWaitDstStageMask-04096)VUID-VkSubmitInfo-pWaitDstStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `pWaitDstStageMask` **must** not contain `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`
- [](#VUID-VkSubmitInfo-pWaitDstStageMask-07318)VUID-VkSubmitInfo-pWaitDstStageMask-07318  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `pWaitDstStageMask` **must** not contain `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkSubmitInfo-pWaitDstStageMask-03937)VUID-VkSubmitInfo-pWaitDstStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `pWaitDstStageMask` **must** not be `0`
- [](#VUID-VkSubmitInfo-pWaitDstStageMask-07949)VUID-VkSubmitInfo-pWaitDstStageMask-07949  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `pWaitDstStageMask` **must** not contain `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkSubmitInfo-pWaitDstStageMask-10754)VUID-VkSubmitInfo-pWaitDstStageMask-10754  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `pWaitDstStageMask` **must** not contain `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-VkSubmitInfo-pCommandBuffers-00075)VUID-VkSubmitInfo-pCommandBuffers-00075  
  Each element of `pCommandBuffers` **must** not have been allocated with `VK_COMMAND_BUFFER_LEVEL_SECONDARY`
- [](#VUID-VkSubmitInfo-pWaitDstStageMask-00078)VUID-VkSubmitInfo-pWaitDstStageMask-00078  
  Each element of `pWaitDstStageMask` **must** not include `VK_PIPELINE_STAGE_HOST_BIT`
- [](#VUID-VkSubmitInfo-pWaitSemaphores-03239)VUID-VkSubmitInfo-pWaitSemaphores-03239  
  If any element of `pWaitSemaphores` or `pSignalSemaphores` was created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE`, then the `pNext` chain **must** include a [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimelineSemaphoreSubmitInfo.html) structure
- [](#VUID-VkSubmitInfo-pNext-03240)VUID-VkSubmitInfo-pNext-03240  
  If the `pNext` chain of this structure includes a [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimelineSemaphoreSubmitInfo.html) structure and any element of `pWaitSemaphores` was created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE`, then its `waitSemaphoreValueCount` member **must** equal `waitSemaphoreCount`
- [](#VUID-VkSubmitInfo-pNext-03241)VUID-VkSubmitInfo-pNext-03241  
  If the `pNext` chain of this structure includes a [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimelineSemaphoreSubmitInfo.html) structure and any element of `pSignalSemaphores` was created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE`, then its `signalSemaphoreValueCount` member **must** equal `signalSemaphoreCount`
- [](#VUID-VkSubmitInfo-pSignalSemaphores-03242)VUID-VkSubmitInfo-pSignalSemaphores-03242  
  For each element of `pSignalSemaphores` created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE` the corresponding element of [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimelineSemaphoreSubmitInfo.html)::`pSignalSemaphoreValues` **must** have a value greater than the current value of the semaphore when the [semaphore signal operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-signaling) is executed
- [](#VUID-VkSubmitInfo-pWaitSemaphores-03243)VUID-VkSubmitInfo-pWaitSemaphores-03243  
  For each element of `pWaitSemaphores` created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE` the corresponding element of [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimelineSemaphoreSubmitInfo.html)::`pWaitSemaphoreValues` **must** have a value which does not differ from the current value of the semaphore or the value of any outstanding semaphore wait or signal operation on that semaphore by more than [`maxTimelineSemaphoreValueDifference`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxTimelineSemaphoreValueDifference)
- [](#VUID-VkSubmitInfo-pSignalSemaphores-03244)VUID-VkSubmitInfo-pSignalSemaphores-03244  
  For each element of `pSignalSemaphores` created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE` the corresponding element of [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimelineSemaphoreSubmitInfo.html)::`pSignalSemaphoreValues` **must** have a value which does not differ from the current value of the semaphore or the value of any outstanding semaphore wait or signal operation on that semaphore by more than [`maxTimelineSemaphoreValueDifference`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxTimelineSemaphoreValueDifference)
- [](#VUID-VkSubmitInfo-pNext-04120)VUID-VkSubmitInfo-pNext-04120  
  If the `pNext` chain of this structure does not include a `VkProtectedSubmitInfo` structure with `protectedSubmit` set to `VK_TRUE`, then each element of the `pCommandBuffers` array **must** be an unprotected command buffer
- [](#VUID-VkSubmitInfo-pNext-04148)VUID-VkSubmitInfo-pNext-04148  
  If the `pNext` chain of this structure includes a `VkProtectedSubmitInfo` structure with `protectedSubmit` set to `VK_TRUE`, then each element of the `pCommandBuffers` array **must** be a protected command buffer
- [](#VUID-VkSubmitInfo-pCommandBuffers-06193)VUID-VkSubmitInfo-pCommandBuffers-06193  
  If `pCommandBuffers` contains any [resumed render pass instances](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-suspension), they **must** be suspended by a render pass instance earlier in submission order within `pCommandBuffers`
- [](#VUID-VkSubmitInfo-pCommandBuffers-06014)VUID-VkSubmitInfo-pCommandBuffers-06014  
  If `pCommandBuffers` contains any [suspended render pass instances](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-suspension), they **must** be resumed by a render pass instance later in submission order within `pCommandBuffers`
- [](#VUID-VkSubmitInfo-pCommandBuffers-06015)VUID-VkSubmitInfo-pCommandBuffers-06015  
  If `pCommandBuffers` contains any [suspended render pass instances](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-suspension), there **must** be no action or synchronization commands executed in a primary or [secondary](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-secondary) command buffer between that render pass instance and the render pass instance that resumes it
- [](#VUID-VkSubmitInfo-pCommandBuffers-06016)VUID-VkSubmitInfo-pCommandBuffers-06016  
  If `pCommandBuffers` contains any [suspended render pass instances](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-suspension), there **must** be no render pass instances between that render pass instance and the render pass instance that resumes it
- [](#VUID-VkSubmitInfo-variableSampleLocations-06017)VUID-VkSubmitInfo-variableSampleLocations-06017  
  If the [`variableSampleLocations`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-variableSampleLocations) limit is not supported, and any element of `pCommandBuffers` contains any [suspended render pass instances](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-suspension), where a graphics pipeline has been bound, any pipelines bound in the render pass instance that resumes it, or any subsequent render pass instances that resume from that one and so on, **must** use the same sample locations
- [](#VUID-VkSubmitInfo-pNext-09683)VUID-VkSubmitInfo-pNext-09683  
  If the `pNext` chain of this structure includes a [VkFrameBoundaryTensorsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryTensorsARM.html) structure then it **must** also include a [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryEXT.html) structure.
- [](#VUID-VkSubmitInfo-pCommandBufferInfos-09942)VUID-VkSubmitInfo-pCommandBufferInfos-09942  
  If at least one [VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferSubmitInfo.html) structure in `pCommandBufferInfos` references a `commandBuffer` allocated from a pool that was created with a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure in the `pNext` chain of [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) that included a foreign data graph processing engine in its `pProcessingEngines` member, then `pWaitSemaphoreInfos` and `pSignalSemaphoreInfos` **must** only reference `semaphore` objects that were created from external handle types reported as supported in a [VkQueueFamilyDataGraphProcessingEnginePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphProcessingEnginePropertiesARM.html)::`foreignSemaphoreHandleTypes` structure via [vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM.html) with a `queueFamilyIndex` matching the one the command pool was created for, for all the foreign data graph processing engines that were part of the [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) used to create the command pool

Valid Usage (Implicit)

- [](#VUID-VkSubmitInfo-sType-sType)VUID-VkSubmitInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBMIT_INFO`
- [](#VUID-VkSubmitInfo-pNext-pNext)VUID-VkSubmitInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkAmigoProfilingSubmitInfoSEC](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAmigoProfilingSubmitInfoSEC.html), [VkD3D12FenceSubmitInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkD3D12FenceSubmitInfoKHR.html), [VkDeviceGroupSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupSubmitInfo.html), [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryEXT.html), [VkFrameBoundaryTensorsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryTensorsARM.html), [VkLatencySubmissionPresentIdNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySubmissionPresentIdNV.html), [VkPerformanceQuerySubmitInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceQuerySubmitInfoKHR.html), [VkProtectedSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkProtectedSubmitInfo.html), [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimelineSemaphoreSubmitInfo.html), [VkWin32KeyedMutexAcquireReleaseInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWin32KeyedMutexAcquireReleaseInfoKHR.html), or [VkWin32KeyedMutexAcquireReleaseInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWin32KeyedMutexAcquireReleaseInfoNV.html)
- [](#VUID-VkSubmitInfo-sType-unique)VUID-VkSubmitInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkSubmitInfo-pWaitSemaphores-parameter)VUID-VkSubmitInfo-pWaitSemaphores-parameter  
  If `waitSemaphoreCount` is not `0`, `pWaitSemaphores` **must** be a valid pointer to an array of `waitSemaphoreCount` valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handles
- [](#VUID-VkSubmitInfo-pWaitDstStageMask-parameter)VUID-VkSubmitInfo-pWaitDstStageMask-parameter  
  If `waitSemaphoreCount` is not `0`, `pWaitDstStageMask` **must** be a valid pointer to an array of `waitSemaphoreCount` valid combinations of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) values
- [](#VUID-VkSubmitInfo-pCommandBuffers-parameter)VUID-VkSubmitInfo-pCommandBuffers-parameter  
  If `commandBufferCount` is not `0`, `pCommandBuffers` **must** be a valid pointer to an array of `commandBufferCount` valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handles
- [](#VUID-VkSubmitInfo-pSignalSemaphores-parameter)VUID-VkSubmitInfo-pSignalSemaphores-parameter  
  If `signalSemaphoreCount` is not `0`, `pSignalSemaphores` **must** be a valid pointer to an array of `signalSemaphoreCount` valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handles
- [](#VUID-VkSubmitInfo-commonparent)VUID-VkSubmitInfo-commonparent  
  Each of the elements of `pCommandBuffers`, the elements of `pSignalSemaphores`, and the elements of `pWaitSemaphores` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubmitInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0