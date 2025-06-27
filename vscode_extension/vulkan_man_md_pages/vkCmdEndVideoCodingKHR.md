# vkCmdEndVideoCodingKHR(3) Manual Page

## Name

vkCmdEndVideoCodingKHR - End video coding scope



## <a href="#_c_specification" class="anchor"></a>C Specification

To end a video coding scope, call:

``` c
// Provided by VK_KHR_video_queue
void vkCmdEndVideoCodingKHR(
    VkCommandBuffer                             commandBuffer,
    const VkVideoEndCodingInfoKHR*              pEndCodingInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which to record the command.

- `pEndCodingInfo` is a pointer to a
  [VkVideoEndCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEndCodingInfoKHR.html) structure
  specifying the parameters for ending the video coding scope.

## <a href="#_description" class="anchor"></a>Description

After ending a video coding scope, the video session object, the
optional video session parameters object, and all <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#bound-reference-picture-resources"
target="_blank" rel="noopener">reference picture resources</a>
previously bound by the corresponding
[vkCmdBeginVideoCodingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginVideoCodingKHR.html) command are
*unbound*.

Valid Usage

- <a href="#VUID-vkCmdEndVideoCodingKHR-None-07251"
  id="VUID-vkCmdEndVideoCodingKHR-None-07251"></a>
  VUID-vkCmdEndVideoCodingKHR-None-07251  
  There **must** be no <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a> queries

Valid Usage (Implicit)

- <a href="#VUID-vkCmdEndVideoCodingKHR-commandBuffer-parameter"
  id="VUID-vkCmdEndVideoCodingKHR-commandBuffer-parameter"></a>
  VUID-vkCmdEndVideoCodingKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdEndVideoCodingKHR-pEndCodingInfo-parameter"
  id="VUID-vkCmdEndVideoCodingKHR-pEndCodingInfo-parameter"></a>
  VUID-vkCmdEndVideoCodingKHR-pEndCodingInfo-parameter  
  `pEndCodingInfo` **must** be a valid pointer to a valid
  [VkVideoEndCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEndCodingInfoKHR.html) structure

- <a href="#VUID-vkCmdEndVideoCodingKHR-commandBuffer-recording"
  id="VUID-vkCmdEndVideoCodingKHR-commandBuffer-recording"></a>
  VUID-vkCmdEndVideoCodingKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdEndVideoCodingKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdEndVideoCodingKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdEndVideoCodingKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support decode, or encode operations

- <a href="#VUID-vkCmdEndVideoCodingKHR-renderpass"
  id="VUID-vkCmdEndVideoCodingKHR-renderpass"></a>
  VUID-vkCmdEndVideoCodingKHR-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdEndVideoCodingKHR-videocoding"
  id="VUID-vkCmdEndVideoCodingKHR-videocoding"></a>
  VUID-vkCmdEndVideoCodingKHR-videocoding  
  This command **must** only be called inside of a video coding scope

- <a href="#VUID-vkCmdEndVideoCodingKHR-bufferlevel"
  id="VUID-vkCmdEndVideoCodingKHR-bufferlevel"></a>
  VUID-vkCmdEndVideoCodingKHR-bufferlevel  
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
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Inside</p></td>
<td class="tableblock halign-left valign-top"><p>Decode<br />
Encode</p></td>
<td class="tableblock halign-left valign-top"><p>Action<br />
State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkVideoEndCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEndCodingInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdEndVideoCodingKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
