# vkCmdSetEvent(3) Manual Page

## Name

vkCmdSetEvent - Set an event object to signaled state



## <a href="#_c_specification" class="anchor"></a>C Specification

To set the state of an event to signaled from a device, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdSetEvent(
    VkCommandBuffer                             commandBuffer,
    VkEvent                                     event,
    VkPipelineStageFlags                        stageMask);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `event` is the event that will be signaled.

- `stageMask` specifies the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
  target="_blank" rel="noopener">source stage mask</a> used to determine
  the first <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
  target="_blank" rel="noopener">synchronization scope</a>.

## <a href="#_description" class="anchor"></a>Description

`vkCmdSetEvent` behaves identically to
[vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2.html), except that it does not define an
access scope, and **must** only be used with
[vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html), not
[vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2.html).

Valid Usage

- <a href="#VUID-vkCmdSetEvent-stageMask-04090"
  id="VUID-vkCmdSetEvent-stageMask-04090"></a>
  VUID-vkCmdSetEvent-stageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`

- <a href="#VUID-vkCmdSetEvent-stageMask-04091"
  id="VUID-vkCmdSetEvent-stageMask-04091"></a>
  VUID-vkCmdSetEvent-stageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-vkCmdSetEvent-stageMask-04092"
  id="VUID-vkCmdSetEvent-stageMask-04092"></a>
  VUID-vkCmdSetEvent-stageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-vkCmdSetEvent-stageMask-04093"
  id="VUID-vkCmdSetEvent-stageMask-04093"></a>
  VUID-vkCmdSetEvent-stageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-vkCmdSetEvent-stageMask-04094"
  id="VUID-vkCmdSetEvent-stageMask-04094"></a>
  VUID-vkCmdSetEvent-stageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-vkCmdSetEvent-stageMask-04095"
  id="VUID-vkCmdSetEvent-stageMask-04095"></a>
  VUID-vkCmdSetEvent-stageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdSetEvent-stageMask-04096"
  id="VUID-vkCmdSetEvent-stageMask-04096"></a>
  VUID-vkCmdSetEvent-stageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdSetEvent-stageMask-07318"
  id="VUID-vkCmdSetEvent-stageMask-07318"></a>
  VUID-vkCmdSetEvent-stageMask-07318  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-vkCmdSetEvent-stageMask-03937"
  id="VUID-vkCmdSetEvent-stageMask-03937"></a>
  VUID-vkCmdSetEvent-stageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `stageMask` **must** not be `0`

- <a href="#VUID-vkCmdSetEvent-stageMask-07949"
  id="VUID-vkCmdSetEvent-stageMask-07949"></a>
  VUID-vkCmdSetEvent-stageMask-07949  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-vkCmdSetEvent-stageMask-06457"
  id="VUID-vkCmdSetEvent-stageMask-06457"></a>
  VUID-vkCmdSetEvent-stageMask-06457  
  Any pipeline stage included in `stageMask` **must** be supported by
  the capabilities of the queue family specified by the
  `queueFamilyIndex` member of the
  [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html) structure that
  was used to create the `VkCommandPool` that `commandBuffer` was
  allocated from, as specified in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-supported"
  target="_blank" rel="noopener">table of supported pipeline stages</a>

- <a href="#VUID-vkCmdSetEvent-stageMask-01149"
  id="VUID-vkCmdSetEvent-stageMask-01149"></a>
  VUID-vkCmdSetEvent-stageMask-01149  
  `stageMask` **must** not include `VK_PIPELINE_STAGE_HOST_BIT`

- <a href="#VUID-vkCmdSetEvent-commandBuffer-01152"
  id="VUID-vkCmdSetEvent-commandBuffer-01152"></a>
  VUID-vkCmdSetEvent-commandBuffer-01152  
  The current device mask of `commandBuffer` **must** include exactly
  one physical device

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetEvent-commandBuffer-parameter"
  id="VUID-vkCmdSetEvent-commandBuffer-parameter"></a>
  VUID-vkCmdSetEvent-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetEvent-event-parameter"
  id="VUID-vkCmdSetEvent-event-parameter"></a>
  VUID-vkCmdSetEvent-event-parameter  
  `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) handle

- <a href="#VUID-vkCmdSetEvent-stageMask-parameter"
  id="VUID-vkCmdSetEvent-stageMask-parameter"></a>
  VUID-vkCmdSetEvent-stageMask-parameter  
  `stageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) values

- <a href="#VUID-vkCmdSetEvent-commandBuffer-recording"
  id="VUID-vkCmdSetEvent-commandBuffer-recording"></a>
  VUID-vkCmdSetEvent-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetEvent-commandBuffer-cmdpool"
  id="VUID-vkCmdSetEvent-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetEvent-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, decode, or encode operations

- <a href="#VUID-vkCmdSetEvent-renderpass"
  id="VUID-vkCmdSetEvent-renderpass"></a>
  VUID-vkCmdSetEvent-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdSetEvent-commonparent"
  id="VUID-vkCmdSetEvent-commonparent"></a>
  VUID-vkCmdSetEvent-commonparent  
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
<tr>
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
<tr>
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
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetEvent"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
