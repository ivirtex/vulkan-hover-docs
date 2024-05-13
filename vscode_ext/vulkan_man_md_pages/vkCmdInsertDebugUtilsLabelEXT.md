# vkCmdInsertDebugUtilsLabelEXT(3) Manual Page

## Name

vkCmdInsertDebugUtilsLabelEXT - Insert a label into a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

A single debug label can be inserted into a command buffer by calling:

``` c
// Provided by VK_EXT_debug_utils
void vkCmdInsertDebugUtilsLabelEXT(
    VkCommandBuffer                             commandBuffer,
    const VkDebugUtilsLabelEXT*                 pLabelInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `pInfo` is a pointer to a
  [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsLabelEXT.html) structure specifying
  parameters of the label to insert.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCmdInsertDebugUtilsLabelEXT-commandBuffer-parameter"
  id="VUID-vkCmdInsertDebugUtilsLabelEXT-commandBuffer-parameter"></a>
  VUID-vkCmdInsertDebugUtilsLabelEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdInsertDebugUtilsLabelEXT-pLabelInfo-parameter"
  id="VUID-vkCmdInsertDebugUtilsLabelEXT-pLabelInfo-parameter"></a>
  VUID-vkCmdInsertDebugUtilsLabelEXT-pLabelInfo-parameter  
  `pLabelInfo` **must** be a valid pointer to a valid
  [VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsLabelEXT.html) structure

- <a href="#VUID-vkCmdInsertDebugUtilsLabelEXT-commandBuffer-recording"
  id="VUID-vkCmdInsertDebugUtilsLabelEXT-commandBuffer-recording"></a>
  VUID-vkCmdInsertDebugUtilsLabelEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdInsertDebugUtilsLabelEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdInsertDebugUtilsLabelEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdInsertDebugUtilsLabelEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdInsertDebugUtilsLabelEXT-videocoding"
  id="VUID-vkCmdInsertDebugUtilsLabelEXT-videocoding"></a>
  VUID-vkCmdInsertDebugUtilsLabelEXT-videocoding  
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

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsLabelEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdInsertDebugUtilsLabelEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
