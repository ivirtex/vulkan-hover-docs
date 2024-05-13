# vkCmdBindPipelineShaderGroupNV(3) Manual Page

## Name

vkCmdBindPipelineShaderGroupNV - Bind a pipeline object



## <a href="#_c_specification" class="anchor"></a>C Specification

For pipelines that were created with the support of multiple shader
groups (see <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#graphics-shadergroups"
target="_blank" rel="noopener">Graphics Pipeline Shader Groups</a>), the
regular `vkCmdBindPipeline` command will bind Shader Group `0`. To
explicitly bind a shader group use:

``` c
// Provided by VK_NV_device_generated_commands
void vkCmdBindPipelineShaderGroupNV(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipeline                                  pipeline,
    uint32_t                                    groupIndex);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer that the pipeline will be bound
  to.

- `pipelineBindPoint` is a
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value specifying the
  bind point to which the pipeline will be bound.

- `pipeline` is the pipeline to be bound.

- `groupIndex` is the shader group to be bound.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCmdBindPipelineShaderGroupNV-groupIndex-02893"
  id="VUID-vkCmdBindPipelineShaderGroupNV-groupIndex-02893"></a>
  VUID-vkCmdBindPipelineShaderGroupNV-groupIndex-02893  
  `groupIndex` **must** be `0` or less than the effective
  [VkGraphicsPipelineShaderGroupsCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineShaderGroupsCreateInfoNV.html)::`groupCount`
  including the referenced pipelines

- <a href="#VUID-vkCmdBindPipelineShaderGroupNV-pipelineBindPoint-02894"
  id="VUID-vkCmdBindPipelineShaderGroupNV-pipelineBindPoint-02894"></a>
  VUID-vkCmdBindPipelineShaderGroupNV-pipelineBindPoint-02894  
  The `pipelineBindPoint` **must** be `VK_PIPELINE_BIND_POINT_GRAPHICS`

- <a href="#VUID-vkCmdBindPipelineShaderGroupNV-groupIndex-02895"
  id="VUID-vkCmdBindPipelineShaderGroupNV-groupIndex-02895"></a>
  VUID-vkCmdBindPipelineShaderGroupNV-groupIndex-02895  
  The same restrictions as [vkCmdBindPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindPipeline.html)
  apply as if the bound pipeline was created only with the Shader Group
  from the `groupIndex` information

- <a
  href="#VUID-vkCmdBindPipelineShaderGroupNV-deviceGeneratedCommands-02896"
  id="VUID-vkCmdBindPipelineShaderGroupNV-deviceGeneratedCommands-02896"></a>
  VUID-vkCmdBindPipelineShaderGroupNV-deviceGeneratedCommands-02896  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedCommands"
  target="_blank" rel="noopener"><code>deviceGeneratedCommands</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-parameter"
  id="VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-parameter"></a>
  VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdBindPipelineShaderGroupNV-pipelineBindPoint-parameter"
  id="VUID-vkCmdBindPipelineShaderGroupNV-pipelineBindPoint-parameter"></a>
  VUID-vkCmdBindPipelineShaderGroupNV-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value

- <a href="#VUID-vkCmdBindPipelineShaderGroupNV-pipeline-parameter"
  id="VUID-vkCmdBindPipelineShaderGroupNV-pipeline-parameter"></a>
  VUID-vkCmdBindPipelineShaderGroupNV-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

- <a href="#VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-recording"
  id="VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-recording"></a>
  VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-cmdpool"
  id="VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdBindPipelineShaderGroupNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdBindPipelineShaderGroupNV-videocoding"
  id="VUID-vkCmdBindPipelineShaderGroupNV-videocoding"></a>
  VUID-vkCmdBindPipelineShaderGroupNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBindPipelineShaderGroupNV-commonparent"
  id="VUID-vkCmdBindPipelineShaderGroupNV-commonparent"></a>
  VUID-vkCmdBindPipelineShaderGroupNV-commonparent  
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

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBindPipelineShaderGroupNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
