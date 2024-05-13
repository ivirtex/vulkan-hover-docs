# vkCmdNextSubpass(3) Manual Page

## Name

vkCmdNextSubpass - Transition to the next subpass of a render pass



## <a href="#_c_specification" class="anchor"></a>C Specification

To transition to the next subpass in the render pass instance after
recording the commands for a subpass, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdNextSubpass(
    VkCommandBuffer                             commandBuffer,
    VkSubpassContents                           contents);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which to record the command.

- `contents` specifies how the commands in the next subpass will be
  provided, in the same fashion as the corresponding parameter of
  [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html).

## <a href="#_description" class="anchor"></a>Description

The subpass index for a render pass begins at zero when
`vkCmdBeginRenderPass` is recorded, and increments each time
`vkCmdNextSubpass` is recorded.

After transitioning to the next subpass, the application **can** record
the commands for that subpass.

Valid Usage

- <a href="#VUID-vkCmdNextSubpass-None-00909"
  id="VUID-vkCmdNextSubpass-None-00909"></a>
  VUID-vkCmdNextSubpass-None-00909  
  The current subpass index **must** be less than the number of
  subpasses in the render pass minus one

- <a href="#VUID-vkCmdNextSubpass-None-02349"
  id="VUID-vkCmdNextSubpass-None-02349"></a>
  VUID-vkCmdNextSubpass-None-02349  
  This command **must** not be recorded when transform feedback is
  active

Valid Usage (Implicit)

- <a href="#VUID-vkCmdNextSubpass-commandBuffer-parameter"
  id="VUID-vkCmdNextSubpass-commandBuffer-parameter"></a>
  VUID-vkCmdNextSubpass-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdNextSubpass-contents-parameter"
  id="VUID-vkCmdNextSubpass-contents-parameter"></a>
  VUID-vkCmdNextSubpass-contents-parameter  
  `contents` **must** be a valid
  [VkSubpassContents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassContents.html) value

- <a href="#VUID-vkCmdNextSubpass-commandBuffer-recording"
  id="VUID-vkCmdNextSubpass-commandBuffer-recording"></a>
  VUID-vkCmdNextSubpass-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdNextSubpass-commandBuffer-cmdpool"
  id="VUID-vkCmdNextSubpass-commandBuffer-cmdpool"></a>
  VUID-vkCmdNextSubpass-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdNextSubpass-renderpass"
  id="VUID-vkCmdNextSubpass-renderpass"></a>
  VUID-vkCmdNextSubpass-renderpass  
  This command **must** only be called inside of a render pass instance

- <a href="#VUID-vkCmdNextSubpass-videocoding"
  id="VUID-vkCmdNextSubpass-videocoding"></a>
  VUID-vkCmdNextSubpass-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdNextSubpass-bufferlevel"
  id="VUID-vkCmdNextSubpass-bufferlevel"></a>
  VUID-vkCmdNextSubpass-bufferlevel  
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
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkSubpassContents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassContents.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdNextSubpass"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
