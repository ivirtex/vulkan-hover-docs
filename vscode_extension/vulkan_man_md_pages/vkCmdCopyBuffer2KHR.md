# vkCmdCopyBuffer2(3) Manual Page

## Name

vkCmdCopyBuffer2 - Copy data between buffer regions



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy data between buffer objects, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdCopyBuffer2(
    VkCommandBuffer                             commandBuffer,
    const VkCopyBufferInfo2*                    pCopyBufferInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_copy_commands2
void vkCmdCopyBuffer2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyBufferInfo2*                    pCopyBufferInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pCopyBufferInfo` is a pointer to a
  [VkCopyBufferInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyBufferInfo2.html) structure describing the
  copy parameters.

## <a href="#_description" class="anchor"></a>Description

Each source region specified by `pCopyBufferInfo->pRegions` is copied
from the source buffer to the destination region of the destination
buffer. If any of the specified regions in `pCopyBufferInfo->srcBuffer`
overlaps in memory with any of the specified regions in
`pCopyBufferInfo->dstBuffer`, values read from those overlapping regions
are undefined.

Valid Usage

- <a href="#VUID-vkCmdCopyBuffer2-commandBuffer-01822"
  id="VUID-vkCmdCopyBuffer2-commandBuffer-01822"></a>
  VUID-vkCmdCopyBuffer2-commandBuffer-01822  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `srcBuffer` **must** not be a protected buffer

- <a href="#VUID-vkCmdCopyBuffer2-commandBuffer-01823"
  id="VUID-vkCmdCopyBuffer2-commandBuffer-01823"></a>
  VUID-vkCmdCopyBuffer2-commandBuffer-01823  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstBuffer` **must** not be a protected buffer

- <a href="#VUID-vkCmdCopyBuffer2-commandBuffer-01824"
  id="VUID-vkCmdCopyBuffer2-commandBuffer-01824"></a>
  VUID-vkCmdCopyBuffer2-commandBuffer-01824  
  If `commandBuffer` is a protected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstBuffer` **must** not be an unprotected buffer

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCopyBuffer2-commandBuffer-parameter"
  id="VUID-vkCmdCopyBuffer2-commandBuffer-parameter"></a>
  VUID-vkCmdCopyBuffer2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyBuffer2-pCopyBufferInfo-parameter"
  id="VUID-vkCmdCopyBuffer2-pCopyBufferInfo-parameter"></a>
  VUID-vkCmdCopyBuffer2-pCopyBufferInfo-parameter  
  `pCopyBufferInfo` **must** be a valid pointer to a valid
  [VkCopyBufferInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyBufferInfo2.html) structure

- <a href="#VUID-vkCmdCopyBuffer2-commandBuffer-recording"
  id="VUID-vkCmdCopyBuffer2-commandBuffer-recording"></a>
  VUID-vkCmdCopyBuffer2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyBuffer2-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyBuffer2-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyBuffer2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdCopyBuffer2-renderpass"
  id="VUID-vkCmdCopyBuffer2-renderpass"></a>
  VUID-vkCmdCopyBuffer2-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyBuffer2-videocoding"
  id="VUID-vkCmdCopyBuffer2-videocoding"></a>
  VUID-vkCmdCopyBuffer2-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Transfer<br />
Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkCopyBufferInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyBufferInfo2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyBuffer2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
