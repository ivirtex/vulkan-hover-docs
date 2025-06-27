# vkCmdWriteBufferMarkerAMD(3) Manual Page

## Name

vkCmdWriteBufferMarkerAMD - Execute a pipelined write of a marker value
into a buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To write a 32-bit marker value into a buffer as a pipelined operation,
call:

``` c
// Provided by VK_AMD_buffer_marker
void vkCmdWriteBufferMarkerAMD(
    VkCommandBuffer                             commandBuffer,
    VkPipelineStageFlagBits                     pipelineStage,
    VkBuffer                                    dstBuffer,
    VkDeviceSize                                dstOffset,
    uint32_t                                    marker);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pipelineStage` is a
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) value
  specifying the pipeline stage whose completion triggers the marker
  write.

- `dstBuffer` is the buffer where the marker will be written to.

- `dstOffset` is the byte offset into the buffer where the marker will
  be written to.

- `marker` is the 32-bit value of the marker.

## <a href="#_description" class="anchor"></a>Description

The command will write the 32-bit marker value into the buffer only
after all preceding commands have finished executing up to at least the
specified pipeline stage. This includes the completion of other
preceding `vkCmdWriteBufferMarkerAMD` commands so long as their
specified pipeline stages occur either at the same time or earlier than
this command’s specified `pipelineStage`.

While consecutive buffer marker writes with the same `pipelineStage`
parameter are implicitly complete in submission order, memory and
execution dependencies between buffer marker writes and other operations
must still be explicitly ordered using synchronization commands. The
access scope for buffer marker writes falls under the
`VK_ACCESS_TRANSFER_WRITE_BIT`, and the pipeline stages for identifying
the synchronization scope **must** include both `pipelineStage` and
`VK_PIPELINE_STAGE_TRANSFER_BIT`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Similar to <code>vkCmdWriteTimestamp</code>, if an implementation is
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
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Implementations <strong>may</strong> only support a limited number of
pipelined marker write operations in flight at a given time, thus
excessive number of marker write operations <strong>may</strong> degrade
command execution performance.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04074"
  id="VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04074"></a>
  VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04074  
  `pipelineStage` **must** be a [valid
  stage](#synchronization-pipeline-stages-supported) for the queue
  family that was used to create the command pool that `commandBuffer`
  was allocated from

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04075"
  id="VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04075"></a>
  VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04075  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `pipelineStage` **must** not be
  `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04076"
  id="VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04076"></a>
  VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04076  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `pipelineStage` **must** not be
  `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04077"
  id="VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04077"></a>
  VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04077  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `pipelineStage` **must** not be
  `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04078"
  id="VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04078"></a>
  VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04078  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `pipelineStage` **must** not be
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04079"
  id="VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04079"></a>
  VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04079  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `pipelineStage` **must** not be
  `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04080"
  id="VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04080"></a>
  VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04080  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `pipelineStage` **must** not be
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-07077"
  id="VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-07077"></a>
  VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-07077  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `pipelineStage` **must** not be
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-shadingRateImage-07314"
  id="VUID-vkCmdWriteBufferMarkerAMD-shadingRateImage-07314"></a>
  VUID-vkCmdWriteBufferMarkerAMD-shadingRateImage-07314  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `pipelineStage` **must** not be
  `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-synchronization2-06489"
  id="VUID-vkCmdWriteBufferMarkerAMD-synchronization2-06489"></a>
  VUID-vkCmdWriteBufferMarkerAMD-synchronization2-06489  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_NONE`

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-rayTracingPipeline-07943"
  id="VUID-vkCmdWriteBufferMarkerAMD-rayTracingPipeline-07943"></a>
  VUID-vkCmdWriteBufferMarkerAMD-rayTracingPipeline-07943  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `pipelineStage` **must** not be
  `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-dstOffset-01798"
  id="VUID-vkCmdWriteBufferMarkerAMD-dstOffset-01798"></a>
  VUID-vkCmdWriteBufferMarkerAMD-dstOffset-01798  
  `dstOffset` **must** be less than or equal to the size of `dstBuffer`
  minus `4`

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-01799"
  id="VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-01799"></a>
  VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-01799  
  `dstBuffer` **must** have been created with
  `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-01800"
  id="VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-01800"></a>
  VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-01800  
  If `dstBuffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-dstOffset-01801"
  id="VUID-vkCmdWriteBufferMarkerAMD-dstOffset-01801"></a>
  VUID-vkCmdWriteBufferMarkerAMD-dstOffset-01801  
  `dstOffset` **must** be a multiple of `4`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-parameter"
  id="VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-parameter"></a>
  VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-parameter"
  id="VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-parameter"></a>
  VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-parameter  
  If `pipelineStage` is not `0`, `pipelineStage` **must** be a valid
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) value

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-parameter"
  id="VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-parameter"></a>
  VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-recording"
  id="VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-recording"></a>
  VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-cmdpool"
  id="VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-cmdpool"></a>
  VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-videocoding"
  id="VUID-vkCmdWriteBufferMarkerAMD-videocoding"></a>
  VUID-vkCmdWriteBufferMarkerAMD-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdWriteBufferMarkerAMD-commonparent"
  id="VUID-vkCmdWriteBufferMarkerAMD-commonparent"></a>
  VUID-vkCmdWriteBufferMarkerAMD-commonparent  
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
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdWriteBufferMarkerAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
