# vkCmdEndRendering(3) Manual Page

## Name

vkCmdEndRendering - End a dynamic render pass instance



## <a href="#_c_specification" class="anchor"></a>C Specification

To end a render pass instance, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdEndRendering(
    VkCommandBuffer                             commandBuffer);
```

or the equivalent command

``` c
// Provided by VK_KHR_dynamic_rendering
void vkCmdEndRenderingKHR(
    VkCommandBuffer                             commandBuffer);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which to record the command.

## <a href="#_description" class="anchor"></a>Description

If the value of `pRenderingInfo->flags` used to begin this render pass
instance included `VK_RENDERING_SUSPENDING_BIT`, then this render pass
is suspended and will be resumed later in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>.

Valid Usage

- <a href="#VUID-vkCmdEndRendering-None-06161"
  id="VUID-vkCmdEndRendering-None-06161"></a>
  VUID-vkCmdEndRendering-None-06161  
  The current render pass instance **must** have been begun with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)

- <a href="#VUID-vkCmdEndRendering-commandBuffer-06162"
  id="VUID-vkCmdEndRendering-commandBuffer-06162"></a>
  VUID-vkCmdEndRendering-commandBuffer-06162  
  The current render pass instance **must** have been begun in
  `commandBuffer`

- <a href="#VUID-vkCmdEndRendering-None-06781"
  id="VUID-vkCmdEndRendering-None-06781"></a>
  VUID-vkCmdEndRendering-None-06781  
  This command **must** not be recorded when transform feedback is
  active

- <a href="#VUID-vkCmdEndRendering-None-06999"
  id="VUID-vkCmdEndRendering-None-06999"></a>
  VUID-vkCmdEndRendering-None-06999  
  If `vkCmdBeginQuery`\* was called within the render pass, the
  corresponding `vkCmdEndQuery`\* **must** have been called subsequently
  within the same subpass

Valid Usage (Implicit)

- <a href="#VUID-vkCmdEndRendering-commandBuffer-parameter"
  id="VUID-vkCmdEndRendering-commandBuffer-parameter"></a>
  VUID-vkCmdEndRendering-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdEndRendering-commandBuffer-recording"
  id="VUID-vkCmdEndRendering-commandBuffer-recording"></a>
  VUID-vkCmdEndRendering-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdEndRendering-commandBuffer-cmdpool"
  id="VUID-vkCmdEndRendering-commandBuffer-cmdpool"></a>
  VUID-vkCmdEndRendering-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdEndRendering-renderpass"
  id="VUID-vkCmdEndRendering-renderpass"></a>
  VUID-vkCmdEndRendering-renderpass  
  This command **must** only be called inside of a render pass instance

- <a href="#VUID-vkCmdEndRendering-videocoding"
  id="VUID-vkCmdEndRendering-videocoding"></a>
  VUID-vkCmdEndRendering-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Action<br />
State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdEndRendering"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
