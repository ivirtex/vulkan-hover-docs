# vkCmdUpdatePipelineIndirectBufferNV(3) Manual Page

## Name

vkCmdUpdatePipelineIndirectBufferNV - Update the indirect compute
pipeline's metadata



## <a href="#_c_specification" class="anchor"></a>C Specification

To save a compute pipeline’s metadata at a device address call:

``` c
// Provided by VK_NV_device_generated_commands_compute
void vkCmdUpdatePipelineIndirectBufferNV(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipeline                                  pipeline);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pipelineBindPoint` is a
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value specifying the
  type of pipeline whose metadata will be saved.

- `pipeline` is the pipeline whose metadata will be saved.

## <a href="#_description" class="anchor"></a>Description

`vkCmdUpdatePipelineIndirectBufferNV` is only allowed outside of a
render pass. This command is treated as a “transfer” operation for the
purposes of synchronization barriers. The writes to the address **must**
be synchronized using stages `VK_PIPELINE_STAGE_2_COPY_BIT` and
`VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_NV` and with access masks
`VK_ACCESS_MEMORY_WRITE_BIT` and
`VK_ACCESS_COMMAND_PREPROCESS_READ_BIT_NV` respectively before using the
results in preprocessing.

Valid Usage

- <a
  href="#VUID-vkCmdUpdatePipelineIndirectBufferNV-pipelineBindPoint-09018"
  id="VUID-vkCmdUpdatePipelineIndirectBufferNV-pipelineBindPoint-09018"></a>
  VUID-vkCmdUpdatePipelineIndirectBufferNV-pipelineBindPoint-09018  
  `pipelineBindPoint` **must** be `VK_PIPELINE_BIND_POINT_COMPUTE`

- <a href="#VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-09019"
  id="VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-09019"></a>
  VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-09019  
  `pipeline` **must** have been created with
  `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV` flag set

- <a href="#VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-09020"
  id="VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-09020"></a>
  VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-09020  
  `pipeline` **must** have been created with
  [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineIndirectBufferInfoNV.html)
  structure specifying a valid address where its metadata will be saved

- <a
  href="#VUID-vkCmdUpdatePipelineIndirectBufferNV-deviceGeneratedComputePipelines-09021"
  id="VUID-vkCmdUpdatePipelineIndirectBufferNV-deviceGeneratedComputePipelines-09021"></a>
  VUID-vkCmdUpdatePipelineIndirectBufferNV-deviceGeneratedComputePipelines-09021  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedComputePipelines"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV</code>::<code>deviceGeneratedComputePipelines</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-parameter"
  id="VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-parameter"></a>
  VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdUpdatePipelineIndirectBufferNV-pipelineBindPoint-parameter"
  id="VUID-vkCmdUpdatePipelineIndirectBufferNV-pipelineBindPoint-parameter"></a>
  VUID-vkCmdUpdatePipelineIndirectBufferNV-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value

- <a href="#VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-parameter"
  id="VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-parameter"></a>
  VUID-vkCmdUpdatePipelineIndirectBufferNV-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

- <a
  href="#VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-recording"
  id="VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-recording"></a>
  VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-cmdpool"
  id="VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdUpdatePipelineIndirectBufferNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdUpdatePipelineIndirectBufferNV-renderpass"
  id="VUID-vkCmdUpdatePipelineIndirectBufferNV-renderpass"></a>
  VUID-vkCmdUpdatePipelineIndirectBufferNV-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdUpdatePipelineIndirectBufferNV-videocoding"
  id="VUID-vkCmdUpdatePipelineIndirectBufferNV-videocoding"></a>
  VUID-vkCmdUpdatePipelineIndirectBufferNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdUpdatePipelineIndirectBufferNV-commonparent"
  id="VUID-vkCmdUpdatePipelineIndirectBufferNV-commonparent"></a>
  VUID-vkCmdUpdatePipelineIndirectBufferNV-commonparent  
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
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Transfer<br />
Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands_compute](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands_compute.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdUpdatePipelineIndirectBufferNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
