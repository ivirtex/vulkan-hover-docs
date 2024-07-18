# VkSubmitInfo(3) Manual Page

## Name

VkSubmitInfo - Structure specifying a queue submit operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSubmitInfo` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `waitSemaphoreCount` is the number of semaphores upon which to wait
  before executing the command buffers for the batch.

- `pWaitSemaphores` is a pointer to an array of
  [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handles upon which to wait before the
  command buffers for this batch begin execution. If semaphores to wait
  on are provided, they define a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-waiting"
  target="_blank" rel="noopener">semaphore wait operation</a>.

- `pWaitDstStageMask` is a pointer to an array of pipeline stages at
  which each corresponding semaphore wait will occur.

- `commandBufferCount` is the number of command buffers to execute in
  the batch.

- `pCommandBuffers` is a pointer to an array of
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handles to execute in the
  batch.

- `signalSemaphoreCount` is the number of semaphores to be signaled once
  the commands specified in `pCommandBuffers` have completed execution.

- `pSignalSemaphores` is a pointer to an array of
  [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handles which will be signaled when
  the command buffers for this batch have completed execution. If
  semaphores to be signaled are provided, they define a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
  target="_blank" rel="noopener">semaphore signal operation</a>.

## <a href="#_description" class="anchor"></a>Description

The order that command buffers appear in `pCommandBuffers` is used to
determine <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>, and thus all the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-implicit"
target="_blank" rel="noopener">implicit ordering guarantees</a> that
respect it. Other than these implicit ordering guarantees and any <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization"
target="_blank" rel="noopener">explicit synchronization primitives</a>,
these command buffers **may** overlap or otherwise execute out of order.

Valid Usage

- <a href="#VUID-VkSubmitInfo-pWaitDstStageMask-04090"
  id="VUID-VkSubmitInfo-pWaitDstStageMask-04090"></a>
  VUID-VkSubmitInfo-pWaitDstStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `pWaitDstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`

- <a href="#VUID-VkSubmitInfo-pWaitDstStageMask-04091"
  id="VUID-VkSubmitInfo-pWaitDstStageMask-04091"></a>
  VUID-VkSubmitInfo-pWaitDstStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `pWaitDstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-VkSubmitInfo-pWaitDstStageMask-04092"
  id="VUID-VkSubmitInfo-pWaitDstStageMask-04092"></a>
  VUID-VkSubmitInfo-pWaitDstStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `pWaitDstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-VkSubmitInfo-pWaitDstStageMask-04093"
  id="VUID-VkSubmitInfo-pWaitDstStageMask-04093"></a>
  VUID-VkSubmitInfo-pWaitDstStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `pWaitDstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-VkSubmitInfo-pWaitDstStageMask-04094"
  id="VUID-VkSubmitInfo-pWaitDstStageMask-04094"></a>
  VUID-VkSubmitInfo-pWaitDstStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `pWaitDstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-VkSubmitInfo-pWaitDstStageMask-04095"
  id="VUID-VkSubmitInfo-pWaitDstStageMask-04095"></a>
  VUID-VkSubmitInfo-pWaitDstStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `pWaitDstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

- <a href="#VUID-VkSubmitInfo-pWaitDstStageMask-04096"
  id="VUID-VkSubmitInfo-pWaitDstStageMask-04096"></a>
  VUID-VkSubmitInfo-pWaitDstStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `pWaitDstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

- <a href="#VUID-VkSubmitInfo-pWaitDstStageMask-07318"
  id="VUID-VkSubmitInfo-pWaitDstStageMask-07318"></a>
  VUID-VkSubmitInfo-pWaitDstStageMask-07318  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `pWaitDstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkSubmitInfo-pWaitDstStageMask-03937"
  id="VUID-VkSubmitInfo-pWaitDstStageMask-03937"></a>
  VUID-VkSubmitInfo-pWaitDstStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `pWaitDstStageMask` **must** not be `0`

- <a href="#VUID-VkSubmitInfo-pWaitDstStageMask-07949"
  id="VUID-VkSubmitInfo-pWaitDstStageMask-07949"></a>
  VUID-VkSubmitInfo-pWaitDstStageMask-07949  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `pWaitDstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-VkSubmitInfo-pCommandBuffers-00075"
  id="VUID-VkSubmitInfo-pCommandBuffers-00075"></a>
  VUID-VkSubmitInfo-pCommandBuffers-00075  
  Each element of `pCommandBuffers` **must** not have been allocated
  with `VK_COMMAND_BUFFER_LEVEL_SECONDARY`

- <a href="#VUID-VkSubmitInfo-pWaitDstStageMask-00078"
  id="VUID-VkSubmitInfo-pWaitDstStageMask-00078"></a>
  VUID-VkSubmitInfo-pWaitDstStageMask-00078  
  Each element of `pWaitDstStageMask` **must** not include
  `VK_PIPELINE_STAGE_HOST_BIT`

- <a href="#VUID-VkSubmitInfo-pWaitSemaphores-03239"
  id="VUID-VkSubmitInfo-pWaitSemaphores-03239"></a>
  VUID-VkSubmitInfo-pWaitSemaphores-03239  
  If any element of `pWaitSemaphores` or `pSignalSemaphores` was created
  with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE`, then the `pNext` chain **must** include
  a [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html)
  structure

- <a href="#VUID-VkSubmitInfo-pNext-03240"
  id="VUID-VkSubmitInfo-pNext-03240"></a>
  VUID-VkSubmitInfo-pNext-03240  
  If the `pNext` chain of this structure includes a
  [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html)
  structure and any element of `pWaitSemaphores` was created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE`, then its `waitSemaphoreValueCount`
  member **must** equal `waitSemaphoreCount`

- <a href="#VUID-VkSubmitInfo-pNext-03241"
  id="VUID-VkSubmitInfo-pNext-03241"></a>
  VUID-VkSubmitInfo-pNext-03241  
  If the `pNext` chain of this structure includes a
  [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html)
  structure and any element of `pSignalSemaphores` was created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE`, then its `signalSemaphoreValueCount`
  member **must** equal `signalSemaphoreCount`

- <a href="#VUID-VkSubmitInfo-pSignalSemaphores-03242"
  id="VUID-VkSubmitInfo-pSignalSemaphores-03242"></a>
  VUID-VkSubmitInfo-pSignalSemaphores-03242  
  For each element of `pSignalSemaphores` created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE` the corresponding element of
  [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html)::`pSignalSemaphoreValues`
  **must** have a value greater than the current value of the semaphore
  when the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
  target="_blank" rel="noopener">semaphore signal operation</a> is
  executed

- <a href="#VUID-VkSubmitInfo-pWaitSemaphores-03243"
  id="VUID-VkSubmitInfo-pWaitSemaphores-03243"></a>
  VUID-VkSubmitInfo-pWaitSemaphores-03243  
  For each element of `pWaitSemaphores` created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE` the corresponding element of
  [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html)::`pWaitSemaphoreValues`
  **must** have a value which does not differ from the current value of
  the semaphore or the value of any outstanding semaphore wait or signal
  operation on that semaphore by more than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxTimelineSemaphoreValueDifference"
  target="_blank"
  rel="noopener"><code>maxTimelineSemaphoreValueDifference</code></a>

- <a href="#VUID-VkSubmitInfo-pSignalSemaphores-03244"
  id="VUID-VkSubmitInfo-pSignalSemaphores-03244"></a>
  VUID-VkSubmitInfo-pSignalSemaphores-03244  
  For each element of `pSignalSemaphores` created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE` the corresponding element of
  [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html)::`pSignalSemaphoreValues`
  **must** have a value which does not differ from the current value of
  the semaphore or the value of any outstanding semaphore wait or signal
  operation on that semaphore by more than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxTimelineSemaphoreValueDifference"
  target="_blank"
  rel="noopener"><code>maxTimelineSemaphoreValueDifference</code></a>

- <a href="#VUID-VkSubmitInfo-pNext-04120"
  id="VUID-VkSubmitInfo-pNext-04120"></a>
  VUID-VkSubmitInfo-pNext-04120  
  If the `pNext` chain of this structure does not include a
  `VkProtectedSubmitInfo` structure with `protectedSubmit` set to
  `VK_TRUE`, then each element of the `pCommandBuffers` array **must**
  be an unprotected command buffer

- <a href="#VUID-VkSubmitInfo-pNext-04148"
  id="VUID-VkSubmitInfo-pNext-04148"></a>
  VUID-VkSubmitInfo-pNext-04148  
  If the `pNext` chain of this structure includes a
  `VkProtectedSubmitInfo` structure with `protectedSubmit` set to
  `VK_TRUE`, then each element of the `pCommandBuffers` array **must**
  be a protected command buffer

- <a href="#VUID-VkSubmitInfo-pCommandBuffers-06193"
  id="VUID-VkSubmitInfo-pCommandBuffers-06193"></a>
  VUID-VkSubmitInfo-pCommandBuffers-06193  
  If `pCommandBuffers` contains any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-suspension"
  target="_blank" rel="noopener">resumed render pass instances</a>, they
  **must** be suspended by a render pass instance earlier in submission
  order within `pCommandBuffers`

- <a href="#VUID-VkSubmitInfo-pCommandBuffers-06014"
  id="VUID-VkSubmitInfo-pCommandBuffers-06014"></a>
  VUID-VkSubmitInfo-pCommandBuffers-06014  
  If `pCommandBuffers` contains any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-suspension"
  target="_blank" rel="noopener">suspended render pass instances</a>,
  they **must** be resumed by a render pass instance later in submission
  order within `pCommandBuffers`

- <a href="#VUID-VkSubmitInfo-pCommandBuffers-06015"
  id="VUID-VkSubmitInfo-pCommandBuffers-06015"></a>
  VUID-VkSubmitInfo-pCommandBuffers-06015  
  If `pCommandBuffers` contains any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-suspension"
  target="_blank" rel="noopener">suspended render pass instances</a>,
  there **must** be no action or synchronization commands executed in a
  primary or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-secondary"
  target="_blank" rel="noopener">secondary</a> command buffer between
  that render pass instance and the render pass instance that resumes it

- <a href="#VUID-VkSubmitInfo-pCommandBuffers-06016"
  id="VUID-VkSubmitInfo-pCommandBuffers-06016"></a>
  VUID-VkSubmitInfo-pCommandBuffers-06016  
  If `pCommandBuffers` contains any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-suspension"
  target="_blank" rel="noopener">suspended render pass instances</a>,
  there **must** be no render pass instances between that render pass
  instance and the render pass instance that resumes it

- <a href="#VUID-VkSubmitInfo-variableSampleLocations-06017"
  id="VUID-VkSubmitInfo-variableSampleLocations-06017"></a>
  VUID-VkSubmitInfo-variableSampleLocations-06017  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-variableSampleLocations"
  target="_blank" rel="noopener"><code>variableSampleLocations</code></a>
  limit is not supported, and any element of `pCommandBuffers` contains
  any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-suspension"
  target="_blank" rel="noopener">suspended render pass instances</a>,
  where a graphics pipeline has been bound, any pipelines bound in the
  render pass instance that resumes it, or any subsequent render pass
  instances that resume from that one and so on, **must** use the same
  sample locations

Valid Usage (Implicit)

- <a href="#VUID-VkSubmitInfo-sType-sType"
  id="VUID-VkSubmitInfo-sType-sType"></a>
  VUID-VkSubmitInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBMIT_INFO`

- <a href="#VUID-VkSubmitInfo-pNext-pNext"
  id="VUID-VkSubmitInfo-pNext-pNext"></a>
  VUID-VkSubmitInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkAmigoProfilingSubmitInfoSEC](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAmigoProfilingSubmitInfoSEC.html),
  [VkD3D12FenceSubmitInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkD3D12FenceSubmitInfoKHR.html),
  [VkDeviceGroupSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupSubmitInfo.html),
  [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrameBoundaryEXT.html),
  [VkLatencySubmissionPresentIdNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySubmissionPresentIdNV.html),
  [VkPerformanceQuerySubmitInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceQuerySubmitInfoKHR.html),
  [VkProtectedSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkProtectedSubmitInfo.html),
  [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfo.html),
  [VkWin32KeyedMutexAcquireReleaseInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWin32KeyedMutexAcquireReleaseInfoKHR.html),
  or
  [VkWin32KeyedMutexAcquireReleaseInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWin32KeyedMutexAcquireReleaseInfoNV.html)

- <a href="#VUID-VkSubmitInfo-sType-unique"
  id="VUID-VkSubmitInfo-sType-unique"></a>
  VUID-VkSubmitInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkSubmitInfo-pWaitSemaphores-parameter"
  id="VUID-VkSubmitInfo-pWaitSemaphores-parameter"></a>
  VUID-VkSubmitInfo-pWaitSemaphores-parameter  
  If `waitSemaphoreCount` is not `0`, `pWaitSemaphores` **must** be a
  valid pointer to an array of `waitSemaphoreCount` valid
  [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handles

- <a href="#VUID-VkSubmitInfo-pWaitDstStageMask-parameter"
  id="VUID-VkSubmitInfo-pWaitDstStageMask-parameter"></a>
  VUID-VkSubmitInfo-pWaitDstStageMask-parameter  
  If `waitSemaphoreCount` is not `0`, `pWaitDstStageMask` **must** be a
  valid pointer to an array of `waitSemaphoreCount` valid combinations
  of [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) values

- <a href="#VUID-VkSubmitInfo-pCommandBuffers-parameter"
  id="VUID-VkSubmitInfo-pCommandBuffers-parameter"></a>
  VUID-VkSubmitInfo-pCommandBuffers-parameter  
  If `commandBufferCount` is not `0`, `pCommandBuffers` **must** be a
  valid pointer to an array of `commandBufferCount` valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handles

- <a href="#VUID-VkSubmitInfo-pSignalSemaphores-parameter"
  id="VUID-VkSubmitInfo-pSignalSemaphores-parameter"></a>
  VUID-VkSubmitInfo-pSignalSemaphores-parameter  
  If `signalSemaphoreCount` is not `0`, `pSignalSemaphores` **must** be
  a valid pointer to an array of `signalSemaphoreCount` valid
  [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handles

- <a href="#VUID-VkSubmitInfo-commonparent"
  id="VUID-VkSubmitInfo-commonparent"></a>
  VUID-VkSubmitInfo-commonparent  
  Each of the elements of `pCommandBuffers`, the elements of
  `pSignalSemaphores`, and the elements of `pWaitSemaphores` that are
  valid handles of non-ignored parameters **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags.html),
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkQueueSubmit](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubmitInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
