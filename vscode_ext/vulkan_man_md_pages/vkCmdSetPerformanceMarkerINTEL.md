# vkCmdSetPerformanceMarkerINTEL(3) Manual Page

## Name

vkCmdSetPerformanceMarkerINTEL - Markers



## <a href="#_c_specification" class="anchor"></a>C Specification

To help associate query results with a particular point at which an
application emitted commands, markers can be set into the command
buffers with the call:

``` c
// Provided by VK_INTEL_performance_query
VkResult vkCmdSetPerformanceMarkerINTEL(
    VkCommandBuffer                             commandBuffer,
    const VkPerformanceMarkerInfoINTEL*         pMarkerInfo);
```

## <a href="#_description" class="anchor"></a>Description

The last marker set onto a command buffer before the end of a query will
be part of the query result.

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-parameter"
  id="VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-parameter"></a>
  VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetPerformanceMarkerINTEL-pMarkerInfo-parameter"
  id="VUID-vkCmdSetPerformanceMarkerINTEL-pMarkerInfo-parameter"></a>
  VUID-vkCmdSetPerformanceMarkerINTEL-pMarkerInfo-parameter  
  `pMarkerInfo` **must** be a valid pointer to a valid
  [VkPerformanceMarkerInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceMarkerInfoINTEL.html)
  structure

- <a href="#VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-recording"
  id="VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-recording"></a>
  VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-cmdpool"
  id="VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, or transfer operations

- <a href="#VUID-vkCmdSetPerformanceMarkerINTEL-videocoding"
  id="VUID-vkCmdSetPerformanceMarkerINTEL-videocoding"></a>
  VUID-vkCmdSetPerformanceMarkerINTEL-videocoding  
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
Compute<br />
Transfer</p></td>
<td class="tableblock halign-left valign-top"><p>Action<br />
State</p></td>
</tr>
</tbody>
</table>

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkPerformanceMarkerInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceMarkerInfoINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetPerformanceMarkerINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
