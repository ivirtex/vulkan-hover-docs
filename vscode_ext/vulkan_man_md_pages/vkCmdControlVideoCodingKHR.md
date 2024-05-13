# vkCmdControlVideoCodingKHR(3) Manual Page

## Name

vkCmdControlVideoCodingKHR - Control video coding parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

To apply dynamic controls to the currently bound video session object,
call:

``` c
// Provided by VK_KHR_video_queue
void vkCmdControlVideoCodingKHR(
    VkCommandBuffer                             commandBuffer,
    const VkVideoCodingControlInfoKHR*          pCodingControlInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which to record the command.

- `pCodingControlInfo` is a pointer to a
  [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html)
  structure specifying the control parameters.

## <a href="#_description" class="anchor"></a>Description

The control parameters provided in this call are applied to the video
session at the time the command executes on the device and are in effect
until a subsequent call to this command with the same video session
bound changes the corresponding control parameters.

A newly created video session **must** be reset before performing video
coding operations using it by including
`VK_VIDEO_CODING_CONTROL_RESET_BIT_KHR` in `pCodingControlInfo->flags`.
The reset operation also returns all DPB slots of the video session to
the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">inactive state</a>. Correspondingly, any
DPB slot index associated with the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#bound-reference-picture-resources"
target="_blank" rel="noopener">bound reference picture resources</a> is
removed.

For encode sessions, the reset operation returns <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control"
target="_blank" rel="noopener">rate control</a> configuration to
implementation default settings and sets the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-quality-level"
target="_blank" rel="noopener">video encode quality level</a> to zero.

After video coding operations are performed using a video session, the
reset operation **can** be used to return the video session to the same
*initial* state as after the reset of a newly created video session.
This **can** be used, for example, when different video sequences are
needed to be processed with the same video session object.

If `pCodingControlInfo->flags` includes
`VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR`, then the command
replaces the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control"
target="_blank" rel="noopener">rate control</a> configuration maintained
by the video session with the configuration specified in the
[VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html)
structure included in the `pCodingControlInfo->pNext` chain.

If `pCodingControlInfo->flags` includes
`VK_VIDEO_CODING_CONTROL_ENCODE_QUALITY_LEVEL_BIT_KHR`, then the command
changes the current <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-quality-level"
target="_blank" rel="noopener">video encode quality level</a> to the
value specified in the `qualityLevel` member of the
[VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelInfoKHR.html)
structure included in the `pCodingControlInfo->pNext` chain.

Valid Usage

- <a href="#VUID-vkCmdControlVideoCodingKHR-flags-07017"
  id="VUID-vkCmdControlVideoCodingKHR-flags-07017"></a>
  VUID-vkCmdControlVideoCodingKHR-flags-07017  
  If `pCodingControlInfo->flags` does not include
  `VK_VIDEO_CODING_CONTROL_RESET_BIT_KHR`, then the bound video session
  **must** not be in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-uninitialized"
  target="_blank" rel="noopener">uninitialized</a> state at the time the
  command is executed on the device

- <a href="#VUID-vkCmdControlVideoCodingKHR-pCodingControlInfo-08243"
  id="VUID-vkCmdControlVideoCodingKHR-pCodingControlInfo-08243"></a>
  VUID-vkCmdControlVideoCodingKHR-pCodingControlInfo-08243  
  If the bound video session was not created with an encode operation,
  then `pCodingControlInfo->flags` **must** not include
  `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR` or
  `VK_VIDEO_CODING_CONTROL_ENCODE_QUALITY_LEVEL_BIT_KHR`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdControlVideoCodingKHR-commandBuffer-parameter"
  id="VUID-vkCmdControlVideoCodingKHR-commandBuffer-parameter"></a>
  VUID-vkCmdControlVideoCodingKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdControlVideoCodingKHR-pCodingControlInfo-parameter"
  id="VUID-vkCmdControlVideoCodingKHR-pCodingControlInfo-parameter"></a>
  VUID-vkCmdControlVideoCodingKHR-pCodingControlInfo-parameter  
  `pCodingControlInfo` **must** be a valid pointer to a valid
  [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html)
  structure

- <a href="#VUID-vkCmdControlVideoCodingKHR-commandBuffer-recording"
  id="VUID-vkCmdControlVideoCodingKHR-commandBuffer-recording"></a>
  VUID-vkCmdControlVideoCodingKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdControlVideoCodingKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdControlVideoCodingKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdControlVideoCodingKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support decode, or encode operations

- <a href="#VUID-vkCmdControlVideoCodingKHR-renderpass"
  id="VUID-vkCmdControlVideoCodingKHR-renderpass"></a>
  VUID-vkCmdControlVideoCodingKHR-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdControlVideoCodingKHR-videocoding"
  id="VUID-vkCmdControlVideoCodingKHR-videocoding"></a>
  VUID-vkCmdControlVideoCodingKHR-videocoding  
  This command **must** only be called inside of a video coding scope

- <a href="#VUID-vkCmdControlVideoCodingKHR-bufferlevel"
  id="VUID-vkCmdControlVideoCodingKHR-bufferlevel"></a>
  VUID-vkCmdControlVideoCodingKHR-bufferlevel  
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
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Inside</p></td>
<td class="tableblock halign-left valign-top"><p>Decode<br />
Encode</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdControlVideoCodingKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
