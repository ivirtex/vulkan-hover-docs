# vkCmdSetRenderingInputAttachmentIndicesKHR(3) Manual Page

## Name

vkCmdSetRenderingInputAttachmentIndicesKHR - Set input attachment index
mappings for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To set the input attachment index mappings during dynamic rendering,
call:

``` c
// Provided by VK_KHR_dynamic_rendering_local_read
void vkCmdSetRenderingInputAttachmentIndicesKHR(
    VkCommandBuffer                             commandBuffer,
    const VkRenderingInputAttachmentIndexInfoKHR* pInputAttachmentIndexInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pInputAttachmentIndexInfo` is a
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)
  structure indicating the new mappings.

## <a href="#_description" class="anchor"></a>Description

This command sets the input attachment index mappings for subsequent
drawing commands, and **must** match the mappings provided to the
currently bound pipeline, if one is bound, which **can** be set by
chaining
[VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)
to [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html).

Until this command is called, mappings in the command buffer state are
treated as each color attachment specified in
[vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) mapping to subpass
inputs with a `InputAttachmentIndex` equal to its index in
[VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments`, and
depth/stencil attachments mapping to input attachments without these
decorations. This state is reset whenever
[vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) is called.

Valid Usage

- <a
  href="#VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-dynamicRenderingLocalRead-09516"
  id="VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-dynamicRenderingLocalRead-09516"></a>
  VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-dynamicRenderingLocalRead-09516  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicRenderingLocalRead"
  target="_blank"
  rel="noopener"><code>dynamicRenderingLocalRead</code></a> **must** be
  enabled

- <a
  href="#VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-pInputAttachmentIndexInfo-09517"
  id="VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-pInputAttachmentIndexInfo-09517"></a>
  VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-pInputAttachmentIndexInfo-09517  
  `pInputAttachmentIndexInfo->colorAttachmentCount` **must** be equal to
  the value of
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`colorAttachmentCount` used
  to begin the current render pass instance

- <a
  href="#VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-commandBuffer-09518"
  id="VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-commandBuffer-09518"></a>
  VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-commandBuffer-09518  
  The current render pass instance **must** have been started or resumed
  by [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) in this
  `commandBuffer`

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-commandBuffer-parameter"
  id="VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-commandBuffer-parameter"></a>
  VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-pInputAttachmentIndexInfo-parameter"
  id="VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-pInputAttachmentIndexInfo-parameter"></a>
  VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-pInputAttachmentIndexInfo-parameter  
  `pInputAttachmentIndexInfo` **must** be a valid pointer to a valid
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)
  structure

- <a
  href="#VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-commandBuffer-recording"
  id="VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-commandBuffer-recording"></a>
  VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-renderpass"
  id="VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-renderpass"></a>
  VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-renderpass  
  This command **must** only be called inside of a render pass instance

- <a href="#VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-videocoding"
  id="VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-videocoding"></a>
  VUID-vkCmdSetRenderingInputAttachmentIndicesKHR-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Inside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_dynamic_rendering_local_read](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering_local_read.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetRenderingInputAttachmentIndicesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
