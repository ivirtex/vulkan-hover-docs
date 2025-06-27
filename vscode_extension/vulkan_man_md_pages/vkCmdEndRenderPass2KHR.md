# vkCmdEndRenderPass2(3) Manual Page

## Name

vkCmdEndRenderPass2 - End the current render pass



## <a href="#_c_specification" class="anchor"></a>C Specification

To record a command to end a render pass instance after recording the
commands for the last subpass, call:

``` c
// Provided by VK_VERSION_1_2
void vkCmdEndRenderPass2(
    VkCommandBuffer                             commandBuffer,
    const VkSubpassEndInfo*                     pSubpassEndInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_create_renderpass2
void vkCmdEndRenderPass2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkSubpassEndInfo*                     pSubpassEndInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which to end the current
  render pass instance.

- `pSubpassEndInfo` is a pointer to a
  [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassEndInfo.html) structure containing
  information about how the last subpass will be ended.

## <a href="#_description" class="anchor"></a>Description

`vkCmdEndRenderPass2` is semantically identical to
[vkCmdEndRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndRenderPass.html), except that it is
extensible.

Valid Usage

- <a href="#VUID-vkCmdEndRenderPass2-None-03103"
  id="VUID-vkCmdEndRenderPass2-None-03103"></a>
  VUID-vkCmdEndRenderPass2-None-03103  
  The current subpass index **must** be equal to the number of subpasses
  in the render pass minus one

- <a href="#VUID-vkCmdEndRenderPass2-None-02352"
  id="VUID-vkCmdEndRenderPass2-None-02352"></a>
  VUID-vkCmdEndRenderPass2-None-02352  
  This command **must** not be recorded when transform feedback is
  active

- <a href="#VUID-vkCmdEndRenderPass2-None-06171"
  id="VUID-vkCmdEndRenderPass2-None-06171"></a>
  VUID-vkCmdEndRenderPass2-None-06171  
  The current render pass instance **must** not have been begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)

- <a href="#VUID-vkCmdEndRenderPass2-None-07005"
  id="VUID-vkCmdEndRenderPass2-None-07005"></a>
  VUID-vkCmdEndRenderPass2-None-07005  
  If `vkCmdBeginQuery`\* was called within a subpass of the render pass,
  the corresponding `vkCmdEndQuery`\* **must** have been called
  subsequently within the same subpass

Valid Usage (Implicit)

- <a href="#VUID-vkCmdEndRenderPass2-commandBuffer-parameter"
  id="VUID-vkCmdEndRenderPass2-commandBuffer-parameter"></a>
  VUID-vkCmdEndRenderPass2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdEndRenderPass2-pSubpassEndInfo-parameter"
  id="VUID-vkCmdEndRenderPass2-pSubpassEndInfo-parameter"></a>
  VUID-vkCmdEndRenderPass2-pSubpassEndInfo-parameter  
  `pSubpassEndInfo` **must** be a valid pointer to a valid
  [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassEndInfo.html) structure

- <a href="#VUID-vkCmdEndRenderPass2-commandBuffer-recording"
  id="VUID-vkCmdEndRenderPass2-commandBuffer-recording"></a>
  VUID-vkCmdEndRenderPass2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdEndRenderPass2-commandBuffer-cmdpool"
  id="VUID-vkCmdEndRenderPass2-commandBuffer-cmdpool"></a>
  VUID-vkCmdEndRenderPass2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdEndRenderPass2-renderpass"
  id="VUID-vkCmdEndRenderPass2-renderpass"></a>
  VUID-vkCmdEndRenderPass2-renderpass  
  This command **must** only be called inside of a render pass instance

- <a href="#VUID-vkCmdEndRenderPass2-videocoding"
  id="VUID-vkCmdEndRenderPass2-videocoding"></a>
  VUID-vkCmdEndRenderPass2-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdEndRenderPass2-bufferlevel"
  id="VUID-vkCmdEndRenderPass2-bufferlevel"></a>
  VUID-vkCmdEndRenderPass2-bufferlevel  
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

[VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassEndInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdEndRenderPass2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
