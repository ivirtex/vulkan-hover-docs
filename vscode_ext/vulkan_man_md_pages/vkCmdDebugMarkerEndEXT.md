# vkCmdDebugMarkerEndEXT(3) Manual Page

## Name

vkCmdDebugMarkerEndEXT - Close a command buffer marker region



## <a href="#_c_specification" class="anchor"></a>C Specification

A marker region can be closed by calling:

``` c
// Provided by VK_EXT_debug_marker
void vkCmdDebugMarkerEndEXT(
    VkCommandBuffer                             commandBuffer);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

## <a href="#_description" class="anchor"></a>Description

An application **may** open a marker region in one command buffer and
close it in another, or otherwise split marker regions across multiple
command buffers or multiple queue submissions. When viewed from the
linear series of submissions to a single queue, the calls to
`vkCmdDebugMarkerBeginEXT` and `vkCmdDebugMarkerEndEXT` **must** be
matched and balanced.

Valid Usage

- <a href="#VUID-vkCmdDebugMarkerEndEXT-commandBuffer-01239"
  id="VUID-vkCmdDebugMarkerEndEXT-commandBuffer-01239"></a>
  VUID-vkCmdDebugMarkerEndEXT-commandBuffer-01239  
  There **must** be an outstanding
  [vkCmdDebugMarkerBeginEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDebugMarkerBeginEXT.html) command
  prior to the `vkCmdDebugMarkerEndEXT` on the queue that
  `commandBuffer` is submitted to

- <a href="#VUID-vkCmdDebugMarkerEndEXT-commandBuffer-01240"
  id="VUID-vkCmdDebugMarkerEndEXT-commandBuffer-01240"></a>
  VUID-vkCmdDebugMarkerEndEXT-commandBuffer-01240  
  If `commandBuffer` is a secondary command buffer, there **must** be an
  outstanding [vkCmdDebugMarkerBeginEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDebugMarkerBeginEXT.html)
  command recorded to `commandBuffer` that has not previously been ended
  by a call to [vkCmdDebugMarkerEndEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDebugMarkerEndEXT.html)

Valid Usage (Implicit)

- <a href="#VUID-vkCmdDebugMarkerEndEXT-commandBuffer-parameter"
  id="VUID-vkCmdDebugMarkerEndEXT-commandBuffer-parameter"></a>
  VUID-vkCmdDebugMarkerEndEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdDebugMarkerEndEXT-commandBuffer-recording"
  id="VUID-vkCmdDebugMarkerEndEXT-commandBuffer-recording"></a>
  VUID-vkCmdDebugMarkerEndEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdDebugMarkerEndEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdDebugMarkerEndEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdDebugMarkerEndEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdDebugMarkerEndEXT-videocoding"
  id="VUID-vkCmdDebugMarkerEndEXT-videocoding"></a>
  VUID-vkCmdDebugMarkerEndEXT-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_marker](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_marker.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdDebugMarkerEndEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
