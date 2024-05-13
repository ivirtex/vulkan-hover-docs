# vkCmdWriteBufferMarker2AMD(3) Manual Page

## Name

vkCmdWriteBufferMarker2AMD - Execute a pipelined write of a marker value
into a buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To write a 32-bit marker value into a buffer as a pipelined operation,
call:

``` c
// Provided by VK_KHR_synchronization2 with VK_AMD_buffer_marker
void vkCmdWriteBufferMarker2AMD(
    VkCommandBuffer                             commandBuffer,
    VkPipelineStageFlags2                       stage,
    VkBuffer                                    dstBuffer,
    VkDeviceSize                                dstOffset,
    uint32_t                                    marker);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `stage` specifies the pipeline stage whose completion triggers the
  marker write.

- `dstBuffer` is the buffer where the marker will be written.

- `dstOffset` is the byte offset into the buffer where the marker will
  be written.

- `marker` is the 32-bit value of the marker.

## <a href="#_description" class="anchor"></a>Description

The command will write the 32-bit marker value into the buffer only
after all preceding commands have finished executing up to at least the
specified pipeline stage. This includes the completion of other
preceding `vkCmdWriteBufferMarker2AMD` commands so long as their
specified pipeline stages occur either at the same time or earlier than
this command’s specified `stage`.

While consecutive buffer marker writes with the same `stage` parameter
implicitly complete in submission order, memory and execution
dependencies between buffer marker writes and other operations **must**
still be explicitly ordered using synchronization commands. The access
scope for buffer marker writes falls under the
`VK_ACCESS_TRANSFER_WRITE_BIT`, and the pipeline stages for identifying
the synchronization scope **must** include both `stage` and
`VK_PIPELINE_STAGE_TRANSFER_BIT`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Similar to <code>vkCmdWriteTimestamp2</code>, if an implementation is
unable to write a marker at any specific pipeline stage, it
<strong>may</strong> instead do so at any logically later
stage.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Implementations <strong>may</strong> only support a limited number of
pipelined marker write operations in flight at a given time. Thus an
excessive number of marker write operations <strong>may</strong> degrade
command execution performance.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-03929"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-03929"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-03930"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-03930"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-03931"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-03931"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-03931  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-03932"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-03932"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-03933"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-03933"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-03934"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-03934"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `stage` **must** not contain `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-03935"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-03935"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `stage` **must** not contain `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-07316"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-07316"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-07316  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-04957"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-04957"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not
  enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-04995"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-04995"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not
  enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-07946"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-07946"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-07946  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-synchronization2-03893"
  id="VUID-vkCmdWriteBufferMarker2AMD-synchronization2-03893"></a>
  VUID-vkCmdWriteBufferMarker2AMD-synchronization2-03893  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
  target="_blank" rel="noopener"><code>synchronization2</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-03894"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-03894"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-03894  
  `stage` **must** include only a single pipeline stage

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-03895"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-03895"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-03895  
  `stage` **must** include only stages that are valid for the queue
  family that was used to create the command pool that `commandBuffer`
  was allocated from

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-dstOffset-03896"
  id="VUID-vkCmdWriteBufferMarker2AMD-dstOffset-03896"></a>
  VUID-vkCmdWriteBufferMarker2AMD-dstOffset-03896  
  `dstOffset` **must** be less than or equal to the size of `dstBuffer`
  minus `4`

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-03897"
  id="VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-03897"></a>
  VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-03897  
  `dstBuffer` **must** have been created with the
  `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-03898"
  id="VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-03898"></a>
  VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-03898  
  If `dstBuffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-dstOffset-03899"
  id="VUID-vkCmdWriteBufferMarker2AMD-dstOffset-03899"></a>
  VUID-vkCmdWriteBufferMarker2AMD-dstOffset-03899  
  `dstOffset` **must** be a multiple of `4`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-parameter"
  id="VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-parameter"></a>
  VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-stage-parameter"
  id="VUID-vkCmdWriteBufferMarker2AMD-stage-parameter"></a>
  VUID-vkCmdWriteBufferMarker2AMD-stage-parameter  
  `stage` **must** be a valid combination of
  [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html) values

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-parameter"
  id="VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-parameter"></a>
  VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-recording"
  id="VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-recording"></a>
  VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-cmdpool"
  id="VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-cmdpool"></a>
  VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-videocoding"
  id="VUID-vkCmdWriteBufferMarker2AMD-videocoding"></a>
  VUID-vkCmdWriteBufferMarker2AMD-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdWriteBufferMarker2AMD-commonparent"
  id="VUID-vkCmdWriteBufferMarker2AMD-commonparent"></a>
  VUID-vkCmdWriteBufferMarker2AMD-commonparent  
  Both of `commandBuffer`, and `dstBuffer` **must** have been created,
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
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Transfer<br />
Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_buffer_marker](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_buffer_marker.html),
[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdWriteBufferMarker2AMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
