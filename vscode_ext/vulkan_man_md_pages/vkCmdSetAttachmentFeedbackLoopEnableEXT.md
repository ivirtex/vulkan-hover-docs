# vkCmdSetAttachmentFeedbackLoopEnableEXT(3) Manual Page

## Name

vkCmdSetAttachmentFeedbackLoopEnableEXT - Specify whether attachment
feedback loops are enabled dynamically on a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> whether a pipeline
**can** access a resource as a non-attachment while it is also used as
an attachment that is written to, call:

``` c
// Provided by VK_EXT_attachment_feedback_loop_dynamic_state
void vkCmdSetAttachmentFeedbackLoopEnableEXT(
    VkCommandBuffer                             commandBuffer,
    VkImageAspectFlags                          aspectMask);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `aspectMask` specifies the types of attachments for which feedback
  loops will be enabled. Attachment types whose aspects are not included
  in `aspectMask` will have feedback loops disabled.

## <a href="#_description" class="anchor"></a>Description

For attachments that are written to in a render pass, only attachments
with the aspects specified in `aspectMask` **can** be accessed as
non-attachments by subsequent <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing"
target="_blank" rel="noopener">drawing commands</a>.

Valid Usage

- <a
  href="#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-attachmentFeedbackLoopDynamicState-08862"
  id="VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-attachmentFeedbackLoopDynamicState-08862"></a>
  VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-attachmentFeedbackLoopDynamicState-08862  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFeedbackLoopDynamicState"
  target="_blank"
  rel="noopener"><code>attachmentFeedbackLoopDynamicState</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-aspectMask-08863"
  id="VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-aspectMask-08863"></a>
  VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-aspectMask-08863  
  `aspectMask` **must** only include `VK_IMAGE_ASPECT_NONE`,
  `VK_IMAGE_ASPECT_COLOR_BIT`, `VK_IMAGE_ASPECT_DEPTH_BIT`, and
  `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a
  href="#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-attachmentFeedbackLoopLayout-08864"
  id="VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-attachmentFeedbackLoopLayout-08864"></a>
  VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-attachmentFeedbackLoopLayout-08864  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFeedbackLoopLayout"
  target="_blank"
  rel="noopener"><code>attachmentFeedbackLoopLayout</code></a> feature
  is not enabled, `aspectMask` **must** be `VK_IMAGE_ASPECT_NONE`

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-aspectMask-parameter"
  id="VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-aspectMask-parameter"></a>
  VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-aspectMask-parameter  
  `aspectMask` **must** be a valid combination of
  [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html) values

- <a
  href="#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-recording"
  id="VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-videocoding"
  id="VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-videocoding"></a>
  VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_attachment_feedback_loop_dynamic_state](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_attachment_feedback_loop_dynamic_state.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetAttachmentFeedbackLoopEnableEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
