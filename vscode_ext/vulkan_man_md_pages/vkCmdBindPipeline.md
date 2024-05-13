# vkCmdBindPipeline(3) Manual Page

## Name

vkCmdBindPipeline - Bind a pipeline object to a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

Once a pipeline has been created, it **can** be bound to the command
buffer using the command:

``` c
// Provided by VK_VERSION_1_0
void vkCmdBindPipeline(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipeline                                  pipeline);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer that the pipeline will be bound
  to.

- `pipelineBindPoint` is a
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value specifying to
  which bind point the pipeline is bound. Binding one does not disturb
  the others.

- `pipeline` is the pipeline to be bound.

## <a href="#_description" class="anchor"></a>Description

Once bound, a pipeline binding affects subsequent commands that interact
with the given pipeline type in the command buffer until a different
pipeline of the same type is bound to the bind point, or until the
pipeline bind point is disturbed by binding a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader object</a> as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects-pipeline-interaction"
target="_blank" rel="noopener">Interaction with Pipelines</a>. Commands
that do not interact with the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-binding"
target="_blank" rel="noopener">given pipeline</a> type **must** not be
affected by the pipeline state.

Valid Usage

- <a href="#VUID-vkCmdBindPipeline-pipelineBindPoint-00777"
  id="VUID-vkCmdBindPipeline-pipelineBindPoint-00777"></a>
  VUID-vkCmdBindPipeline-pipelineBindPoint-00777  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_COMPUTE`, the
  `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdBindPipeline-pipelineBindPoint-00778"
  id="VUID-vkCmdBindPipeline-pipelineBindPoint-00778"></a>
  VUID-vkCmdBindPipeline-pipelineBindPoint-00778  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_GRAPHICS`, the
  `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdBindPipeline-pipelineBindPoint-00779"
  id="VUID-vkCmdBindPipeline-pipelineBindPoint-00779"></a>
  VUID-vkCmdBindPipeline-pipelineBindPoint-00779  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_COMPUTE`, `pipeline`
  **must** be a compute pipeline

- <a href="#VUID-vkCmdBindPipeline-pipelineBindPoint-00780"
  id="VUID-vkCmdBindPipeline-pipelineBindPoint-00780"></a>
  VUID-vkCmdBindPipeline-pipelineBindPoint-00780  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_GRAPHICS`,
  `pipeline` **must** be a graphics pipeline

- <a href="#VUID-vkCmdBindPipeline-pipeline-00781"
  id="VUID-vkCmdBindPipeline-pipeline-00781"></a>
  VUID-vkCmdBindPipeline-pipeline-00781  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-variableMultisampleRate"
  target="_blank" rel="noopener"><code>variableMultisampleRate</code></a>
  feature is not supported, `pipeline` is a graphics pipeline, the
  current subpass <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-noattachments"
  target="_blank" rel="noopener">uses no attachments</a>, and this is
  not the first call to this function with a graphics pipeline after
  transitioning to the current subpass, then the sample count specified
  by this pipeline **must** match that set in the previous pipeline

- <a href="#VUID-vkCmdBindPipeline-variableSampleLocations-01525"
  id="VUID-vkCmdBindPipeline-variableSampleLocations-01525"></a>
  VUID-vkCmdBindPipeline-variableSampleLocations-01525  
  If
  [VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html)::`variableSampleLocations`
  is `VK_FALSE`, and `pipeline` is a graphics pipeline created with a
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)
  structure having its `sampleLocationsEnable` member set to `VK_TRUE`
  but without `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT` enabled then the
  current render pass instance **must** have been begun by specifying a
  [VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html)
  structure whose `pPostSubpassSampleLocations` member contains an
  element with a `subpassIndex` matching the current subpass index and
  the `sampleLocationsInfo` member of that element **must** match the
  `sampleLocationsInfo` specified in
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)
  when the pipeline was created

- <a href="#VUID-vkCmdBindPipeline-None-02323"
  id="VUID-vkCmdBindPipeline-None-02323"></a>
  VUID-vkCmdBindPipeline-None-02323  
  This command **must** not be recorded when transform feedback is
  active

- <a href="#VUID-vkCmdBindPipeline-pipelineBindPoint-02391"
  id="VUID-vkCmdBindPipeline-pipelineBindPoint-02391"></a>
  VUID-vkCmdBindPipeline-pipelineBindPoint-02391  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR`,
  the `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdBindPipeline-pipelineBindPoint-02392"
  id="VUID-vkCmdBindPipeline-pipelineBindPoint-02392"></a>
  VUID-vkCmdBindPipeline-pipelineBindPoint-02392  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR`,
  `pipeline` **must** be a ray tracing pipeline

- <a href="#VUID-vkCmdBindPipeline-pipelineBindPoint-06721"
  id="VUID-vkCmdBindPipeline-pipelineBindPoint-06721"></a>
  VUID-vkCmdBindPipeline-pipelineBindPoint-06721  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR`,
  `commandBuffer` **must** not be a protected command buffer

- <a href="#VUID-vkCmdBindPipeline-pipelineProtectedAccess-07408"
  id="VUID-vkCmdBindPipeline-pipelineProtectedAccess-07408"></a>
  VUID-vkCmdBindPipeline-pipelineProtectedAccess-07408  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineProtectedAccess"
  target="_blank" rel="noopener"><code>pipelineProtectedAccess</code></a>
  feature is enabled, and `commandBuffer` is a protected command buffer,
  `pipeline` **must** have been created without
  `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT_EXT`

- <a href="#VUID-vkCmdBindPipeline-pipelineProtectedAccess-07409"
  id="VUID-vkCmdBindPipeline-pipelineProtectedAccess-07409"></a>
  VUID-vkCmdBindPipeline-pipelineProtectedAccess-07409  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineProtectedAccess"
  target="_blank" rel="noopener"><code>pipelineProtectedAccess</code></a>
  feature is enabled, and `commandBuffer` is not a protected command
  buffer, `pipeline` **must** have been created without
  `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT_EXT`

- <a href="#VUID-vkCmdBindPipeline-pipeline-03382"
  id="VUID-vkCmdBindPipeline-pipeline-03382"></a>
  VUID-vkCmdBindPipeline-pipeline-03382  
  `pipeline` **must** not have been created with
  `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` set

- <a href="#VUID-vkCmdBindPipeline-commandBuffer-04808"
  id="VUID-vkCmdBindPipeline-commandBuffer-04808"></a>
  VUID-vkCmdBindPipeline-commandBuffer-04808  
  If `commandBuffer` is a secondary command buffer with
  [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D`
  enabled and `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_GRAPHICS`,
  then the `pipeline` **must** have been created with
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` or `VK_DYNAMIC_STATE_VIEWPORT`,
  and `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT` or
  `VK_DYNAMIC_STATE_SCISSOR` enabled

- <a href="#VUID-vkCmdBindPipeline-commandBuffer-04809"
  id="VUID-vkCmdBindPipeline-commandBuffer-04809"></a>
  VUID-vkCmdBindPipeline-commandBuffer-04809  
  If `commandBuffer` is a secondary command buffer with
  [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D`
  enabled and `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_GRAPHICS`
  and `pipeline` was created with
  [VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)
  structure and its `discardRectangleCount` member is not `0`, or the
  pipeline was created with
  `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT` enabled, then the
  pipeline **must** have been created with
  `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT` enabled

- <a href="#VUID-vkCmdBindPipeline-pipelineBindPoint-04881"
  id="VUID-vkCmdBindPipeline-pipelineBindPoint-04881"></a>
  VUID-vkCmdBindPipeline-pipelineBindPoint-04881  
  If `pipelineBindPoint` is `VK_PIPELINE_BIND_POINT_GRAPHICS` and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-provokingVertexModePerPipeline"
  target="_blank"
  rel="noopener"><code>provokingVertexModePerPipeline</code></a> limit
  is `VK_FALSE`, then pipeline’s
  [VkPipelineRasterizationProvokingVertexStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationProvokingVertexStateCreateInfoEXT.html)::`provokingVertexMode`
  **must** be the same as that of any other pipelines previously bound
  to this bind point within the current render pass instance, including
  any pipeline already bound when beginning the render pass instance

- <a href="#VUID-vkCmdBindPipeline-pipelineBindPoint-04949"
  id="VUID-vkCmdBindPipeline-pipelineBindPoint-04949"></a>
  VUID-vkCmdBindPipeline-pipelineBindPoint-04949  
  If `pipelineBindPoint` is
  `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI`, the `VkCommandPool`
  that `commandBuffer` was allocated from **must** support compute
  operations

- <a href="#VUID-vkCmdBindPipeline-pipelineBindPoint-04950"
  id="VUID-vkCmdBindPipeline-pipelineBindPoint-04950"></a>
  VUID-vkCmdBindPipeline-pipelineBindPoint-04950  
  If `pipelineBindPoint` is
  `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI`, `pipeline` **must**
  be a subpass shading pipeline

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBindPipeline-commandBuffer-parameter"
  id="VUID-vkCmdBindPipeline-commandBuffer-parameter"></a>
  VUID-vkCmdBindPipeline-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBindPipeline-pipelineBindPoint-parameter"
  id="VUID-vkCmdBindPipeline-pipelineBindPoint-parameter"></a>
  VUID-vkCmdBindPipeline-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value

- <a href="#VUID-vkCmdBindPipeline-pipeline-parameter"
  id="VUID-vkCmdBindPipeline-pipeline-parameter"></a>
  VUID-vkCmdBindPipeline-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

- <a href="#VUID-vkCmdBindPipeline-commandBuffer-recording"
  id="VUID-vkCmdBindPipeline-commandBuffer-recording"></a>
  VUID-vkCmdBindPipeline-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBindPipeline-commandBuffer-cmdpool"
  id="VUID-vkCmdBindPipeline-commandBuffer-cmdpool"></a>
  VUID-vkCmdBindPipeline-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdBindPipeline-videocoding"
  id="VUID-vkCmdBindPipeline-videocoding"></a>
  VUID-vkCmdBindPipeline-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBindPipeline-commonparent"
  id="VUID-vkCmdBindPipeline-commonparent"></a>
  VUID-vkCmdBindPipeline-commonparent  
  Both of `commandBuffer`, and `pipeline` **must** have been created,
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
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBindPipeline"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
