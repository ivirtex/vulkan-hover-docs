# vkCmdInitializeGraphScratchMemoryAMDX(3) Manual Page

## Name

vkCmdInitializeGraphScratchMemoryAMDX - Initialize scratch memory for an
execution graph



## <a href="#_c_specification" class="anchor"></a>C Specification

To initialize scratch memory for a particular execution graph, call:

``` c
// Provided by VK_AMDX_shader_enqueue
void vkCmdInitializeGraphScratchMemoryAMDX(
    VkCommandBuffer                             commandBuffer,
    VkDeviceAddress                             scratch);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `scratch` is a pointer to the scratch memory to be initialized.

## <a href="#_description" class="anchor"></a>Description

This command **must** be called before using `scratch` to dispatch the
currently bound execution graph pipeline.

Execution of this command **may** modify any memory locations in the
range \[`scratch`,`scratch` + `size`), where `size` is the value
returned in
[VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html)::`size`
by
[VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html)
for the currently bound execution graph pipeline. Accesses to this
memory range are performed in the
`VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT` pipeline stage with the
`VK_ACCESS_2_SHADER_STORAGE_READ_BIT` and
`VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT` access flags.

If any portion of `scratch` is modified by any command other than
[vkCmdDispatchGraphAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchGraphAMDX.html),
[vkCmdDispatchGraphIndirectAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchGraphIndirectAMDX.html),
[vkCmdDispatchGraphIndirectCountAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchGraphIndirectCountAMDX.html),
or `vkCmdInitializeGraphScratchMemoryAMDX` with the same execution
graph, it **must** be reinitialized for the execution graph again before
dispatching against it.

Valid Usage

- <a href="#VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratch-09143"
  id="VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratch-09143"></a>
  VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratch-09143  
  `scratch` **must** be the device address of an allocated memory range
  at least as large as the value of
  [VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html)::`size`
  returned by
  [VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html)
  for the currently bound execution graph pipeline.

- <a href="#VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratch-09144"
  id="VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratch-09144"></a>
  VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratch-09144  
  `scratch` **must** be a multiple of 64

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-parameter"
  id="VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-parameter"></a>
  VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-recording"
  id="VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-recording"></a>
  VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-cmdpool"
  id="VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-cmdpool"></a>
  VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdInitializeGraphScratchMemoryAMDX-renderpass"
  id="VUID-vkCmdInitializeGraphScratchMemoryAMDX-renderpass"></a>
  VUID-vkCmdInitializeGraphScratchMemoryAMDX-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdInitializeGraphScratchMemoryAMDX-videocoding"
  id="VUID-vkCmdInitializeGraphScratchMemoryAMDX-videocoding"></a>
  VUID-vkCmdInitializeGraphScratchMemoryAMDX-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdInitializeGraphScratchMemoryAMDX-bufferlevel"
  id="VUID-vkCmdInitializeGraphScratchMemoryAMDX-bufferlevel"></a>
  VUID-vkCmdInitializeGraphScratchMemoryAMDX-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`

Host Synchronization

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
<td class="tableblock halign-left valign-top"><p>Primary</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMDX_shader_enqueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMDX_shader_enqueue.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdInitializeGraphScratchMemoryAMDX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
