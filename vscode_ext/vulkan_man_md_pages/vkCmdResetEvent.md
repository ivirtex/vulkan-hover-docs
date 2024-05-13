# vkCmdResetEvent(3) Manual Page

## Name

vkCmdResetEvent - Reset an event object to non-signaled state



## <a href="#_c_specification" class="anchor"></a>C Specification

To set the state of an event to unsignaled from a device, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdResetEvent(
    VkCommandBuffer                             commandBuffer,
    VkEvent                                     event,
    VkPipelineStageFlags                        stageMask);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `event` is the event that will be unsignaled.

- `stageMask` is a bitmask of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) specifying the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
  target="_blank" rel="noopener">source stage mask</a> used to determine
  when the `event` is unsignaled.

## <a href="#_description" class="anchor"></a>Description

`vkCmdResetEvent` behaves identically to
[vkCmdResetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResetEvent2.html).

Valid Usage

- <a href="#VUID-vkCmdResetEvent-stageMask-04090"
  id="VUID-vkCmdResetEvent-stageMask-04090"></a>
  VUID-vkCmdResetEvent-stageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`

- <a href="#VUID-vkCmdResetEvent-stageMask-04091"
  id="VUID-vkCmdResetEvent-stageMask-04091"></a>
  VUID-vkCmdResetEvent-stageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-vkCmdResetEvent-stageMask-04092"
  id="VUID-vkCmdResetEvent-stageMask-04092"></a>
  VUID-vkCmdResetEvent-stageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-vkCmdResetEvent-stageMask-04093"
  id="VUID-vkCmdResetEvent-stageMask-04093"></a>
  VUID-vkCmdResetEvent-stageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-vkCmdResetEvent-stageMask-04094"
  id="VUID-vkCmdResetEvent-stageMask-04094"></a>
  VUID-vkCmdResetEvent-stageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-vkCmdResetEvent-stageMask-04095"
  id="VUID-vkCmdResetEvent-stageMask-04095"></a>
  VUID-vkCmdResetEvent-stageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdResetEvent-stageMask-04096"
  id="VUID-vkCmdResetEvent-stageMask-04096"></a>
  VUID-vkCmdResetEvent-stageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdResetEvent-stageMask-07318"
  id="VUID-vkCmdResetEvent-stageMask-07318"></a>
  VUID-vkCmdResetEvent-stageMask-07318  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-vkCmdResetEvent-stageMask-03937"
  id="VUID-vkCmdResetEvent-stageMask-03937"></a>
  VUID-vkCmdResetEvent-stageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `stageMask` **must** not be `0`

- <a href="#VUID-vkCmdResetEvent-stageMask-07949"
  id="VUID-vkCmdResetEvent-stageMask-07949"></a>
  VUID-vkCmdResetEvent-stageMask-07949  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-vkCmdResetEvent-stageMask-06458"
  id="VUID-vkCmdResetEvent-stageMask-06458"></a>
  VUID-vkCmdResetEvent-stageMask-06458  
  Any pipeline stage included in `stageMask` **must** be supported by
  the capabilities of the queue family specified by the
  `queueFamilyIndex` member of the
  [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html) structure that
  was used to create the `VkCommandPool` that `commandBuffer` was
  allocated from, as specified in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-supported"
  target="_blank" rel="noopener">table of supported pipeline stages</a>

- <a href="#VUID-vkCmdResetEvent-stageMask-01153"
  id="VUID-vkCmdResetEvent-stageMask-01153"></a>
  VUID-vkCmdResetEvent-stageMask-01153  
  `stageMask` **must** not include `VK_PIPELINE_STAGE_HOST_BIT`

- <a href="#VUID-vkCmdResetEvent-event-03834"
  id="VUID-vkCmdResetEvent-event-03834"></a>
  VUID-vkCmdResetEvent-event-03834  
  There **must** be an execution dependency between `vkCmdResetEvent`
  and the execution of any [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html) that
  includes `event` in its `pEvents` parameter

- <a href="#VUID-vkCmdResetEvent-event-03835"
  id="VUID-vkCmdResetEvent-event-03835"></a>
  VUID-vkCmdResetEvent-event-03835  
  There **must** be an execution dependency between `vkCmdResetEvent`
  and the execution of any [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2.html)
  that includes `event` in its `pEvents` parameter

- <a href="#VUID-vkCmdResetEvent-commandBuffer-01157"
  id="VUID-vkCmdResetEvent-commandBuffer-01157"></a>
  VUID-vkCmdResetEvent-commandBuffer-01157  
  `commandBuffer`’s current device mask **must** include exactly one
  physical device

Valid Usage (Implicit)

- <a href="#VUID-vkCmdResetEvent-commandBuffer-parameter"
  id="VUID-vkCmdResetEvent-commandBuffer-parameter"></a>
  VUID-vkCmdResetEvent-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdResetEvent-event-parameter"
  id="VUID-vkCmdResetEvent-event-parameter"></a>
  VUID-vkCmdResetEvent-event-parameter  
  `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) handle

- <a href="#VUID-vkCmdResetEvent-stageMask-parameter"
  id="VUID-vkCmdResetEvent-stageMask-parameter"></a>
  VUID-vkCmdResetEvent-stageMask-parameter  
  `stageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) values

- <a href="#VUID-vkCmdResetEvent-commandBuffer-recording"
  id="VUID-vkCmdResetEvent-commandBuffer-recording"></a>
  VUID-vkCmdResetEvent-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdResetEvent-commandBuffer-cmdpool"
  id="VUID-vkCmdResetEvent-commandBuffer-cmdpool"></a>
  VUID-vkCmdResetEvent-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, decode, or encode operations

- <a href="#VUID-vkCmdResetEvent-renderpass"
  id="VUID-vkCmdResetEvent-renderpass"></a>
  VUID-vkCmdResetEvent-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdResetEvent-commonparent"
  id="VUID-vkCmdResetEvent-commonparent"></a>
  VUID-vkCmdResetEvent-commonparent  
  Both of `commandBuffer`, and `event` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr class="header">
<th class="tableblock halign-left valign-top"><a
href="#VkCommandBufferLevel">Command Buffer Levels</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginRenderPass">Render Pass Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginVideoCodingKHR">Video Coding Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#VkQueueFlagBits">Supported Queue Types</a></th>
<th class="tableblock halign-left valign-top"><a
href="#fundamentals-queueoperation-command-types">Command Type</a></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute<br />
Decode<br />
Encode</p></td>
<td
class="tableblock halign-left valign-top"><p>Synchronization</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html),
[VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdResetEvent"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
