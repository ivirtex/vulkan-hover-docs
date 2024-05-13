# vkCmdOpticalFlowExecuteNV(3) Manual Page

## Name

vkCmdOpticalFlowExecuteNV - Calculate optical flow vectors



## <a href="#_c_specification" class="anchor"></a>C Specification

Default direction of flow estimation is forward which calculates the
optical flow from input frame to reference frame. Optionally backward
flow estimation can be additionally calculated. An output flow vector
(Vx, Vy) means that current pixel (x, y) of input frame can be found at
location (x+Vx, y+Vy) in reference frame. A backward flow vector (Vx,
Vy) means that current pixel (x, y) of reference frame can be found at
location (x+Vx, y+Vy) in input frame.

To calculate optical flow vectors from two input frames, call:

``` c
// Provided by VK_NV_optical_flow
void vkCmdOpticalFlowExecuteNV(
    VkCommandBuffer                             commandBuffer,
    VkOpticalFlowSessionNV                      session,
    const VkOpticalFlowExecuteInfoNV*           pExecuteInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `session` is the optical flow session object on which this command is
  operating.

- `pExecuteInfo` Info is a pointer to a
  [VkOpticalFlowExecuteInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowExecuteInfoNV.html).

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-parameter"
  id="VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-parameter"></a>
  VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdOpticalFlowExecuteNV-session-parameter"
  id="VUID-vkCmdOpticalFlowExecuteNV-session-parameter"></a>
  VUID-vkCmdOpticalFlowExecuteNV-session-parameter  
  `session` **must** be a valid
  [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionNV.html) handle

- <a href="#VUID-vkCmdOpticalFlowExecuteNV-pExecuteInfo-parameter"
  id="VUID-vkCmdOpticalFlowExecuteNV-pExecuteInfo-parameter"></a>
  VUID-vkCmdOpticalFlowExecuteNV-pExecuteInfo-parameter  
  `pExecuteInfo` **must** be a valid pointer to a valid
  [VkOpticalFlowExecuteInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowExecuteInfoNV.html)
  structure

- <a href="#VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-recording"
  id="VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-recording"></a>
  VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-cmdpool"
  id="VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdOpticalFlowExecuteNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support optical flow operations

- <a href="#VUID-vkCmdOpticalFlowExecuteNV-renderpass"
  id="VUID-vkCmdOpticalFlowExecuteNV-renderpass"></a>
  VUID-vkCmdOpticalFlowExecuteNV-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdOpticalFlowExecuteNV-videocoding"
  id="VUID-vkCmdOpticalFlowExecuteNV-videocoding"></a>
  VUID-vkCmdOpticalFlowExecuteNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdOpticalFlowExecuteNV-commonparent"
  id="VUID-vkCmdOpticalFlowExecuteNV-commonparent"></a>
  VUID-vkCmdOpticalFlowExecuteNV-commonparent  
  Both of `commandBuffer`, and `session` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Opticalflow</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkOpticalFlowExecuteInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowExecuteInfoNV.html),
[VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdOpticalFlowExecuteNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
