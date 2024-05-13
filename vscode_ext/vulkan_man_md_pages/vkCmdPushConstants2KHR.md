# vkCmdPushConstants2KHR(3) Manual Page

## Name

vkCmdPushConstants2KHR - Update the values of push constants



## <a href="#_c_specification" class="anchor"></a>C Specification

Alternatively, to update push constants, call:

``` c
// Provided by VK_KHR_maintenance6
void vkCmdPushConstants2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkPushConstantsInfoKHR*               pPushConstantsInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which the push constant
  update will be recorded.

- `pPushConstantsInfo` is a pointer to a `VkPushConstantsInfoKHR`
  structure.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCmdPushConstants2KHR-commandBuffer-parameter"
  id="VUID-vkCmdPushConstants2KHR-commandBuffer-parameter"></a>
  VUID-vkCmdPushConstants2KHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdPushConstants2KHR-pPushConstantsInfo-parameter"
  id="VUID-vkCmdPushConstants2KHR-pPushConstantsInfo-parameter"></a>
  VUID-vkCmdPushConstants2KHR-pPushConstantsInfo-parameter  
  `pPushConstantsInfo` **must** be a valid pointer to a valid
  [VkPushConstantsInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantsInfoKHR.html) structure

- <a href="#VUID-vkCmdPushConstants2KHR-commandBuffer-recording"
  id="VUID-vkCmdPushConstants2KHR-commandBuffer-recording"></a>
  VUID-vkCmdPushConstants2KHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdPushConstants2KHR-commandBuffer-cmdpool"
  id="VUID-vkCmdPushConstants2KHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdPushConstants2KHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdPushConstants2KHR-videocoding"
  id="VUID-vkCmdPushConstants2KHR-videocoding"></a>
  VUID-vkCmdPushConstants2KHR-videocoding  
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
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance6](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance6.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkPushConstantsInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantsInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdPushConstants2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
