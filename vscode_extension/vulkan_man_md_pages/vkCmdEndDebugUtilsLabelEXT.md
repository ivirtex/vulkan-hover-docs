# vkCmdEndDebugUtilsLabelEXT(3) Manual Page

## Name

vkCmdEndDebugUtilsLabelEXT - Close a command buffer label region



## <a href="#_c_specification" class="anchor"></a>C Specification

A command buffer label region can be closed by calling:

``` c
// Provided by VK_EXT_debug_utils
void vkCmdEndDebugUtilsLabelEXT(
    VkCommandBuffer                             commandBuffer);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

## <a href="#_description" class="anchor"></a>Description

An application **may** open a debug label region in one command buffer
and close it in another, or otherwise split debug label regions across
multiple command buffers or multiple queue submissions. When viewed from
the linear series of submissions to a single queue, the calls to
[vkCmdBeginDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginDebugUtilsLabelEXT.html) and
[vkCmdEndDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndDebugUtilsLabelEXT.html) **must**
be matched and balanced.

There **can** be problems reporting command buffer debug labels during
the recording process because command buffers **may** be recorded out of
sequence with the resulting execution order. Since the recording order
**may** be different, a solitary command buffer **may** have an
inconsistent view of the debug label regions by itself. Therefore, if an
issue occurs during the recording of a command buffer, and the
environment requires returning debug labels, the implementation **may**
return only those labels it is aware of. This is true even if the
implementation is aware of only the debug labels within the command
buffer being actively recorded.

Valid Usage

- <a href="#VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-01912"
  id="VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-01912"></a>
  VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-01912  
  There **must** be an outstanding `vkCmdBeginDebugUtilsLabelEXT`
  command prior to the `vkCmdEndDebugUtilsLabelEXT` on the queue that
  `commandBuffer` is submitted to

- <a href="#VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-01913"
  id="VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-01913"></a>
  VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-01913  
  If `commandBuffer` is a secondary command buffer, there **must** be an
  outstanding `vkCmdBeginDebugUtilsLabelEXT` command recorded to
  `commandBuffer` that has not previously been ended by a call to
  `vkCmdEndDebugUtilsLabelEXT`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-parameter"
  id="VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-parameter"></a>
  VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-recording"
  id="VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-recording"></a>
  VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdEndDebugUtilsLabelEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdEndDebugUtilsLabelEXT-videocoding"
  id="VUID-vkCmdEndDebugUtilsLabelEXT-videocoding"></a>
  VUID-vkCmdEndDebugUtilsLabelEXT-videocoding  
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

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdEndDebugUtilsLabelEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
