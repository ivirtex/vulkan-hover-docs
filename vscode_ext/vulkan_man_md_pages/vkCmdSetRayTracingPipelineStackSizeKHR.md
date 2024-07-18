# vkCmdSetRayTracingPipelineStackSizeKHR(3) Manual Page

## Name

vkCmdSetRayTracingPipelineStackSizeKHR - Set the stack size dynamically
for a ray tracing pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the stack size for a
ray tracing pipeline, call:

``` c
// Provided by VK_KHR_ray_tracing_pipeline
void vkCmdSetRayTracingPipelineStackSizeKHR(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    pipelineStackSize);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pipelineStackSize` is the stack size to use for subsequent ray
  tracing trace commands.

## <a href="#_description" class="anchor"></a>Description

This command sets the stack size for subsequent ray tracing commands
when the ray tracing pipeline is created with
`VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, the stack size is computed as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-tracing-pipeline-stack"
target="_blank" rel="noopener">Ray Tracing Pipeline Stack</a>.

Valid Usage

- <a
  href="#VUID-vkCmdSetRayTracingPipelineStackSizeKHR-pipelineStackSize-03610"
  id="VUID-vkCmdSetRayTracingPipelineStackSizeKHR-pipelineStackSize-03610"></a>
  VUID-vkCmdSetRayTracingPipelineStackSizeKHR-pipelineStackSize-03610  
  `pipelineStackSize` **must** be large enough for any dynamic execution
  through the shaders in the ray tracing pipeline used by a subsequent
  trace call

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-parameter"
  id="VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-parameter"></a>
  VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-recording"
  id="VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-recording"></a>
  VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetRayTracingPipelineStackSizeKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdSetRayTracingPipelineStackSizeKHR-renderpass"
  id="VUID-vkCmdSetRayTracingPipelineStackSizeKHR-renderpass"></a>
  VUID-vkCmdSetRayTracingPipelineStackSizeKHR-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdSetRayTracingPipelineStackSizeKHR-videocoding"
  id="VUID-vkCmdSetRayTracingPipelineStackSizeKHR-videocoding"></a>
  VUID-vkCmdSetRayTracingPipelineStackSizeKHR-videocoding  
  This command **must** only be called outside of a video coding scope

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
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Compute</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetRayTracingPipelineStackSizeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
