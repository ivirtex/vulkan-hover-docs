# vkCmdSetDiscardRectangleEnableEXT(3) Manual Page

## Name

vkCmdSetDiscardRectangleEnableEXT - Enable discard rectangles
dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> whether discard
rectangles are enabled, call:

``` c
// Provided by VK_EXT_discard_rectangles
void vkCmdSetDiscardRectangleEnableEXT(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    discardRectangleEnable);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `discardRectangleEnable` specifies whether discard rectangles are
  enabled or not.

## <a href="#_description" class="anchor"></a>Description

This command sets the discard rectangle enable for subsequent drawing
commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT`
set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is implied by the
[VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)::`discardRectangleCount`
value used to create the currently active pipeline, where a non-zero
`discardRectangleCount` implicitly enables discard rectangles, otherwise
they are disabled.

Valid Usage

- <a href="#VUID-vkCmdSetDiscardRectangleEnableEXT-specVersion-07851"
  id="VUID-vkCmdSetDiscardRectangleEnableEXT-specVersion-07851"></a>
  VUID-vkCmdSetDiscardRectangleEnableEXT-specVersion-07851  
  The [`VK_EXT_discard_rectangles`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_discard_rectangles.html)
  extension **must** be enabled, and the implementation **must** support
  at least `specVersion` `2` of this extension

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdSetDiscardRectangleEnableEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetDiscardRectangleEnableEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetDiscardRectangleEnableEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetDiscardRectangleEnableEXT-commandBuffer-recording"
  id="VUID-vkCmdSetDiscardRectangleEnableEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetDiscardRectangleEnableEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetDiscardRectangleEnableEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetDiscardRectangleEnableEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetDiscardRectangleEnableEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetDiscardRectangleEnableEXT-videocoding"
  id="VUID-vkCmdSetDiscardRectangleEnableEXT-videocoding"></a>
  VUID-vkCmdSetDiscardRectangleEnableEXT-videocoding  
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

[VK_EXT_discard_rectangles](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_discard_rectangles.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetDiscardRectangleEnableEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
