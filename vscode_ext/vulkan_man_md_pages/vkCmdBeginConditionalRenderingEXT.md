# vkCmdBeginConditionalRenderingEXT(3) Manual Page

## Name

vkCmdBeginConditionalRenderingEXT - Define the beginning of a
conditional rendering block



## <a href="#_c_specification" class="anchor"></a>C Specification

To begin conditional rendering, call:

``` c
// Provided by VK_EXT_conditional_rendering
void vkCmdBeginConditionalRenderingEXT(
    VkCommandBuffer                             commandBuffer,
    const VkConditionalRenderingBeginInfoEXT*   pConditionalRenderingBegin);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which this command will be
  recorded.

- `pConditionalRenderingBegin` is a pointer to a
  [VkConditionalRenderingBeginInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConditionalRenderingBeginInfoEXT.html)
  structure specifying parameters of conditional rendering.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCmdBeginConditionalRenderingEXT-None-01980"
  id="VUID-vkCmdBeginConditionalRenderingEXT-None-01980"></a>
  VUID-vkCmdBeginConditionalRenderingEXT-None-01980  
  Conditional rendering **must** not already be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#active-conditional-rendering"
  target="_blank" rel="noopener">active</a>

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-parameter"
  id="VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-parameter"></a>
  VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdBeginConditionalRenderingEXT-pConditionalRenderingBegin-parameter"
  id="VUID-vkCmdBeginConditionalRenderingEXT-pConditionalRenderingBegin-parameter"></a>
  VUID-vkCmdBeginConditionalRenderingEXT-pConditionalRenderingBegin-parameter  
  `pConditionalRenderingBegin` **must** be a valid pointer to a valid
  [VkConditionalRenderingBeginInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConditionalRenderingBeginInfoEXT.html)
  structure

- <a
  href="#VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-recording"
  id="VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-recording"></a>
  VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdBeginConditionalRenderingEXT-videocoding"
  id="VUID-vkCmdBeginConditionalRenderingEXT-videocoding"></a>
  VUID-vkCmdBeginConditionalRenderingEXT-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action<br />
State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_conditional_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conditional_rendering.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkConditionalRenderingBeginInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConditionalRenderingBeginInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBeginConditionalRenderingEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
