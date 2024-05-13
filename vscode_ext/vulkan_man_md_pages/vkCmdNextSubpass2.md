# vkCmdNextSubpass2(3) Manual Page

## Name

vkCmdNextSubpass2 - Transition to the next subpass of a render pass



## <a href="#_c_specification" class="anchor"></a>C Specification

To transition to the next subpass in the render pass instance after
recording the commands for a subpass, call:

``` c
// Provided by VK_VERSION_1_2
void vkCmdNextSubpass2(
    VkCommandBuffer                             commandBuffer,
    const VkSubpassBeginInfo*                   pSubpassBeginInfo,
    const VkSubpassEndInfo*                     pSubpassEndInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_create_renderpass2
void vkCmdNextSubpass2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkSubpassBeginInfo*                   pSubpassBeginInfo,
    const VkSubpassEndInfo*                     pSubpassEndInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which to record the command.

- `pSubpassBeginInfo` is a pointer to a
  [VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassBeginInfo.html) structure containing
  information about the subpass which is about to begin rendering.

- `pSubpassEndInfo` is a pointer to a
  [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassEndInfo.html) structure containing
  information about how the previous subpass will be ended.

## <a href="#_description" class="anchor"></a>Description

`vkCmdNextSubpass2` is semantically identical to
[vkCmdNextSubpass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass.html), except that it is extensible,
and that `contents` is provided as part of an extensible structure
instead of as a flat parameter.

Valid Usage

- <a href="#VUID-vkCmdNextSubpass2-None-03102"
  id="VUID-vkCmdNextSubpass2-None-03102"></a>
  VUID-vkCmdNextSubpass2-None-03102  
  The current subpass index **must** be less than the number of
  subpasses in the render pass minus one

- <a href="#VUID-vkCmdNextSubpass2-None-02350"
  id="VUID-vkCmdNextSubpass2-None-02350"></a>
  VUID-vkCmdNextSubpass2-None-02350  
  This command **must** not be recorded when transform feedback is
  active

Valid Usage (Implicit)

- <a href="#VUID-vkCmdNextSubpass2-commandBuffer-parameter"
  id="VUID-vkCmdNextSubpass2-commandBuffer-parameter"></a>
  VUID-vkCmdNextSubpass2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdNextSubpass2-pSubpassBeginInfo-parameter"
  id="VUID-vkCmdNextSubpass2-pSubpassBeginInfo-parameter"></a>
  VUID-vkCmdNextSubpass2-pSubpassBeginInfo-parameter  
  `pSubpassBeginInfo` **must** be a valid pointer to a valid
  [VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassBeginInfo.html) structure

- <a href="#VUID-vkCmdNextSubpass2-pSubpassEndInfo-parameter"
  id="VUID-vkCmdNextSubpass2-pSubpassEndInfo-parameter"></a>
  VUID-vkCmdNextSubpass2-pSubpassEndInfo-parameter  
  `pSubpassEndInfo` **must** be a valid pointer to a valid
  [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassEndInfo.html) structure

- <a href="#VUID-vkCmdNextSubpass2-commandBuffer-recording"
  id="VUID-vkCmdNextSubpass2-commandBuffer-recording"></a>
  VUID-vkCmdNextSubpass2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdNextSubpass2-commandBuffer-cmdpool"
  id="VUID-vkCmdNextSubpass2-commandBuffer-cmdpool"></a>
  VUID-vkCmdNextSubpass2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdNextSubpass2-renderpass"
  id="VUID-vkCmdNextSubpass2-renderpass"></a>
  VUID-vkCmdNextSubpass2-renderpass  
  This command **must** only be called inside of a render pass instance

- <a href="#VUID-vkCmdNextSubpass2-videocoding"
  id="VUID-vkCmdNextSubpass2-videocoding"></a>
  VUID-vkCmdNextSubpass2-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdNextSubpass2-bufferlevel"
  id="VUID-vkCmdNextSubpass2-bufferlevel"></a>
  VUID-vkCmdNextSubpass2-bufferlevel  
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

[VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassBeginInfo.html),
[VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassEndInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdNextSubpass2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
