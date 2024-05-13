# vkCmdEndRenderPass(3) Manual Page

## Name

vkCmdEndRenderPass - End the current render pass



## <a href="#_c_specification" class="anchor"></a>C Specification

To record a command to end a render pass instance after recording the
commands for the last subpass, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdEndRenderPass(
    VkCommandBuffer                             commandBuffer);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which to end the current
  render pass instance.

## <a href="#_description" class="anchor"></a>Description

Ending a render pass instance performs any multisample resolve
operations on the final subpass.

Valid Usage

- <a href="#VUID-vkCmdEndRenderPass-None-00910"
  id="VUID-vkCmdEndRenderPass-None-00910"></a>
  VUID-vkCmdEndRenderPass-None-00910  
  The current subpass index **must** be equal to the number of subpasses
  in the render pass minus one

- <a href="#VUID-vkCmdEndRenderPass-None-02351"
  id="VUID-vkCmdEndRenderPass-None-02351"></a>
  VUID-vkCmdEndRenderPass-None-02351  
  This command **must** not be recorded when transform feedback is
  active

- <a href="#VUID-vkCmdEndRenderPass-None-06170"
  id="VUID-vkCmdEndRenderPass-None-06170"></a>
  VUID-vkCmdEndRenderPass-None-06170  
  The current render pass instance **must** not have been begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)

- <a href="#VUID-vkCmdEndRenderPass-None-07004"
  id="VUID-vkCmdEndRenderPass-None-07004"></a>
  VUID-vkCmdEndRenderPass-None-07004  
  If `vkCmdBeginQuery`\* was called within a subpass of the render pass,
  the corresponding `vkCmdEndQuery`\* **must** have been called
  subsequently within the same subpass

Valid Usage (Implicit)

- <a href="#VUID-vkCmdEndRenderPass-commandBuffer-parameter"
  id="VUID-vkCmdEndRenderPass-commandBuffer-parameter"></a>
  VUID-vkCmdEndRenderPass-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdEndRenderPass-commandBuffer-recording"
  id="VUID-vkCmdEndRenderPass-commandBuffer-recording"></a>
  VUID-vkCmdEndRenderPass-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdEndRenderPass-commandBuffer-cmdpool"
  id="VUID-vkCmdEndRenderPass-commandBuffer-cmdpool"></a>
  VUID-vkCmdEndRenderPass-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdEndRenderPass-renderpass"
  id="VUID-vkCmdEndRenderPass-renderpass"></a>
  VUID-vkCmdEndRenderPass-renderpass  
  This command **must** only be called inside of a render pass instance

- <a href="#VUID-vkCmdEndRenderPass-videocoding"
  id="VUID-vkCmdEndRenderPass-videocoding"></a>
  VUID-vkCmdEndRenderPass-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdEndRenderPass-bufferlevel"
  id="VUID-vkCmdEndRenderPass-bufferlevel"></a>
  VUID-vkCmdEndRenderPass-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`

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
<td class="tableblock halign-left valign-top"><p>Primary</p></td>
<td class="tableblock halign-left valign-top"><p>Inside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>Action<br />
State<br />
Synchronization</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdEndRenderPass"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
