# vkCmdResetEvent2(3) Manual Page

## Name

vkCmdResetEvent2 - Reset an event object to non-signaled state



## <a href="#_c_specification" class="anchor"></a>C Specification

To unsignal the event from a device, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdResetEvent2(
    VkCommandBuffer                             commandBuffer,
    VkEvent                                     event,
    VkPipelineStageFlags2                       stageMask);
```

or the equivalent command

``` c
// Provided by VK_KHR_synchronization2
void vkCmdResetEvent2KHR(
    VkCommandBuffer                             commandBuffer,
    VkEvent                                     event,
    VkPipelineStageFlags2                       stageMask);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `event` is the event that will be unsignaled.

- `stageMask` is a [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2.html)
  mask of pipeline stages used to determine the first <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
  target="_blank" rel="noopener">synchronization scope</a>.

## <a href="#_description" class="anchor"></a>Description

When [vkCmdResetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResetEvent2.html) is submitted to a queue,
it defines an execution dependency on commands that were submitted
before it, and defines an event unsignal operation which resets the
event to the unsignaled state.

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands that occur earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>. The synchronization
scope is limited to operations by `stageMask` or stages that are <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-order"
target="_blank" rel="noopener">logically earlier</a> than `stageMask`.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes only
the event unsignal operation.

If `event` is already in the unsignaled state when
[vkCmdResetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResetEvent2.html) is executed on the device,
then this command has no effect, no event unsignal operation occurs, and
no execution dependency is generated.

Valid Usage

- <a href="#VUID-vkCmdResetEvent2-stageMask-03929"
  id="VUID-vkCmdResetEvent2-stageMask-03929"></a>
  VUID-vkCmdResetEvent2-stageMask-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`

- <a href="#VUID-vkCmdResetEvent2-stageMask-03930"
  id="VUID-vkCmdResetEvent2-stageMask-03930"></a>
  VUID-vkCmdResetEvent2-stageMask-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-vkCmdResetEvent2-stageMask-03931"
  id="VUID-vkCmdResetEvent2-stageMask-03931"></a>
  VUID-vkCmdResetEvent2-stageMask-03931  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-vkCmdResetEvent2-stageMask-03932"
  id="VUID-vkCmdResetEvent2-stageMask-03932"></a>
  VUID-vkCmdResetEvent2-stageMask-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-vkCmdResetEvent2-stageMask-03933"
  id="VUID-vkCmdResetEvent2-stageMask-03933"></a>
  VUID-vkCmdResetEvent2-stageMask-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-vkCmdResetEvent2-stageMask-03934"
  id="VUID-vkCmdResetEvent2-stageMask-03934"></a>
  VUID-vkCmdResetEvent2-stageMask-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdResetEvent2-stageMask-03935"
  id="VUID-vkCmdResetEvent2-stageMask-03935"></a>
  VUID-vkCmdResetEvent2-stageMask-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdResetEvent2-stageMask-07316"
  id="VUID-vkCmdResetEvent2-stageMask-07316"></a>
  VUID-vkCmdResetEvent2-stageMask-07316  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-vkCmdResetEvent2-stageMask-04957"
  id="VUID-vkCmdResetEvent2-stageMask-04957"></a>
  VUID-vkCmdResetEvent2-stageMask-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not
  enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`

- <a href="#VUID-vkCmdResetEvent2-stageMask-04995"
  id="VUID-vkCmdResetEvent2-stageMask-04995"></a>
  VUID-vkCmdResetEvent2-stageMask-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not
  enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

- <a href="#VUID-vkCmdResetEvent2-stageMask-07946"
  id="VUID-vkCmdResetEvent2-stageMask-07946"></a>
  VUID-vkCmdResetEvent2-stageMask-07946  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-vkCmdResetEvent2-synchronization2-03829"
  id="VUID-vkCmdResetEvent2-synchronization2-03829"></a>
  VUID-vkCmdResetEvent2-synchronization2-03829  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
  target="_blank" rel="noopener"><code>synchronization2</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdResetEvent2-stageMask-03830"
  id="VUID-vkCmdResetEvent2-stageMask-03830"></a>
  VUID-vkCmdResetEvent2-stageMask-03830  
  `stageMask` **must** not include `VK_PIPELINE_STAGE_2_HOST_BIT`

- <a href="#VUID-vkCmdResetEvent2-event-03831"
  id="VUID-vkCmdResetEvent2-event-03831"></a>
  VUID-vkCmdResetEvent2-event-03831  
  There **must** be an execution dependency between `vkCmdResetEvent2`
  and the execution of any [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html) that
  includes `event` in its `pEvents` parameter

- <a href="#VUID-vkCmdResetEvent2-event-03832"
  id="VUID-vkCmdResetEvent2-event-03832"></a>
  VUID-vkCmdResetEvent2-event-03832  
  There **must** be an execution dependency between `vkCmdResetEvent2`
  and the execution of any [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2.html)
  that includes `event` in its `pEvents` parameter

- <a href="#VUID-vkCmdResetEvent2-commandBuffer-03833"
  id="VUID-vkCmdResetEvent2-commandBuffer-03833"></a>
  VUID-vkCmdResetEvent2-commandBuffer-03833  
  `commandBuffer`’s current device mask **must** include exactly one
  physical device

Valid Usage (Implicit)

- <a href="#VUID-vkCmdResetEvent2-commandBuffer-parameter"
  id="VUID-vkCmdResetEvent2-commandBuffer-parameter"></a>
  VUID-vkCmdResetEvent2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdResetEvent2-event-parameter"
  id="VUID-vkCmdResetEvent2-event-parameter"></a>
  VUID-vkCmdResetEvent2-event-parameter  
  `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) handle

- <a href="#VUID-vkCmdResetEvent2-stageMask-parameter"
  id="VUID-vkCmdResetEvent2-stageMask-parameter"></a>
  VUID-vkCmdResetEvent2-stageMask-parameter  
  `stageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html) values

- <a href="#VUID-vkCmdResetEvent2-commandBuffer-recording"
  id="VUID-vkCmdResetEvent2-commandBuffer-recording"></a>
  VUID-vkCmdResetEvent2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdResetEvent2-commandBuffer-cmdpool"
  id="VUID-vkCmdResetEvent2-commandBuffer-cmdpool"></a>
  VUID-vkCmdResetEvent2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, decode, or encode operations

- <a href="#VUID-vkCmdResetEvent2-renderpass"
  id="VUID-vkCmdResetEvent2-renderpass"></a>
  VUID-vkCmdResetEvent2-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdResetEvent2-commonparent"
  id="VUID-vkCmdResetEvent2-commonparent"></a>
  VUID-vkCmdResetEvent2-commonparent  
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

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html),
[VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdResetEvent2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
