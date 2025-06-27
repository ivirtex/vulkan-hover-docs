# vkCmdEndConditionalRenderingEXT(3) Manual Page

## Name

vkCmdEndConditionalRenderingEXT - Define the end of a conditional
rendering block



## <a href="#_c_specification" class="anchor"></a>C Specification

To end conditional rendering, call:

``` c
// Provided by VK_EXT_conditional_rendering
void vkCmdEndConditionalRenderingEXT(
    VkCommandBuffer                             commandBuffer);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which this command will be
  recorded.

## <a href="#_description" class="anchor"></a>Description

Once ended, conditional rendering becomes inactive.

Valid Usage

- <a href="#VUID-vkCmdEndConditionalRenderingEXT-None-01985"
  id="VUID-vkCmdEndConditionalRenderingEXT-None-01985"></a>
  VUID-vkCmdEndConditionalRenderingEXT-None-01985  
  Conditional rendering **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#active-conditional-rendering"
  target="_blank" rel="noopener">active</a>

- <a href="#VUID-vkCmdEndConditionalRenderingEXT-None-01986"
  id="VUID-vkCmdEndConditionalRenderingEXT-None-01986"></a>
  VUID-vkCmdEndConditionalRenderingEXT-None-01986  
  If conditional rendering was made <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#active-conditional-rendering"
  target="_blank" rel="noopener">active</a> outside of a render pass
  instance, it **must** not be ended inside a render pass instance

- <a href="#VUID-vkCmdEndConditionalRenderingEXT-None-01987"
  id="VUID-vkCmdEndConditionalRenderingEXT-None-01987"></a>
  VUID-vkCmdEndConditionalRenderingEXT-None-01987  
  If conditional rendering was made <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#active-conditional-rendering"
  target="_blank" rel="noopener">active</a> within a subpass it **must**
  be ended in the same subpass

Valid Usage (Implicit)

- <a href="#VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-parameter"
  id="VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-parameter"></a>
  VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-recording"
  id="VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-recording"></a>
  VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdEndConditionalRenderingEXT-videocoding"
  id="VUID-vkCmdEndConditionalRenderingEXT-videocoding"></a>
  VUID-vkCmdEndConditionalRenderingEXT-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action<br />
State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_conditional_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conditional_rendering.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdEndConditionalRenderingEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
