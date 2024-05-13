# vkCmdDebugMarkerBeginEXT(3) Manual Page

## Name

vkCmdDebugMarkerBeginEXT - Open a command buffer marker region



## <a href="#_c_specification" class="anchor"></a>C Specification

A marker region can be opened by calling:

``` c
// Provided by VK_EXT_debug_marker
void vkCmdDebugMarkerBeginEXT(
    VkCommandBuffer                             commandBuffer,
    const VkDebugMarkerMarkerInfoEXT*           pMarkerInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `pMarkerInfo` is a pointer to a
  [VkDebugMarkerMarkerInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugMarkerMarkerInfoEXT.html)
  structure specifying the parameters of the marker region to open.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-parameter"
  id="VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-parameter"></a>
  VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdDebugMarkerBeginEXT-pMarkerInfo-parameter"
  id="VUID-vkCmdDebugMarkerBeginEXT-pMarkerInfo-parameter"></a>
  VUID-vkCmdDebugMarkerBeginEXT-pMarkerInfo-parameter  
  `pMarkerInfo` **must** be a valid pointer to a valid
  [VkDebugMarkerMarkerInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugMarkerMarkerInfoEXT.html)
  structure

- <a href="#VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-recording"
  id="VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-recording"></a>
  VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdDebugMarkerBeginEXT-videocoding"
  id="VUID-vkCmdDebugMarkerBeginEXT-videocoding"></a>
  VUID-vkCmdDebugMarkerBeginEXT-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_marker](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_marker.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDebugMarkerMarkerInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugMarkerMarkerInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdDebugMarkerBeginEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
