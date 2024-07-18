# vkCmdSetDiscardRectangleModeEXT(3) Manual Page

## Name

vkCmdSetDiscardRectangleModeEXT - Sets the discard rectangle mode
dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the discard rectangle
mode, call:

``` c
// Provided by VK_EXT_discard_rectangles
void vkCmdSetDiscardRectangleModeEXT(
    VkCommandBuffer                             commandBuffer,
    VkDiscardRectangleModeEXT                   discardRectangleMode);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `discardRectangleMode` specifies the discard rectangle mode for all
  discard rectangles, either inclusive or exclusive.

## <a href="#_description" class="anchor"></a>Description

This command sets the discard rectangle mode for subsequent drawing
commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT`
set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)::`discardRectangleMode`
value used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetDiscardRectangleModeEXT-specVersion-07852"
  id="VUID-vkCmdSetDiscardRectangleModeEXT-specVersion-07852"></a>
  VUID-vkCmdSetDiscardRectangleModeEXT-specVersion-07852  
  The [`VK_EXT_discard_rectangles`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_discard_rectangles.html)
  extension **must** be enabled, and the implementation **must** support
  at least `specVersion` `2` of this extension

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetDiscardRectangleModeEXT-discardRectangleMode-parameter"
  id="VUID-vkCmdSetDiscardRectangleModeEXT-discardRectangleMode-parameter"></a>
  VUID-vkCmdSetDiscardRectangleModeEXT-discardRectangleMode-parameter  
  `discardRectangleMode` **must** be a valid
  [VkDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDiscardRectangleModeEXT.html) value

- <a href="#VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-recording"
  id="VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetDiscardRectangleModeEXT-videocoding"
  id="VUID-vkCmdSetDiscardRectangleModeEXT-videocoding"></a>
  VUID-vkCmdSetDiscardRectangleModeEXT-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_discard_rectangles](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_discard_rectangles.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDiscardRectangleModeEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetDiscardRectangleModeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
