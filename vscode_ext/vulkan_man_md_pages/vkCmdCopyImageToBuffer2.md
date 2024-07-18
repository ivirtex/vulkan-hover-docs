# vkCmdCopyImageToBuffer2(3) Manual Page

## Name

vkCmdCopyImageToBuffer2 - Copy image data into a buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy data from an image object to a buffer object, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdCopyImageToBuffer2(
    VkCommandBuffer                             commandBuffer,
    const VkCopyImageToBufferInfo2*             pCopyImageToBufferInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_copy_commands2
void vkCmdCopyImageToBuffer2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyImageToBufferInfo2*             pCopyImageToBufferInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pCopyImageToBufferInfo` is a pointer to a
  [VkCopyImageToBufferInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToBufferInfo2.html) structure
  describing the copy parameters.

## <a href="#_description" class="anchor"></a>Description

This command is functionally identical to
[vkCmdCopyImageToBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImageToBuffer.html), but includes
extensible sub-structures that include `sType` and `pNext` parameters,
allowing them to be more easily extended.

Valid Usage

- <a href="#VUID-vkCmdCopyImageToBuffer2-commandBuffer-01831"
  id="VUID-vkCmdCopyImageToBuffer2-commandBuffer-01831"></a>
  VUID-vkCmdCopyImageToBuffer2-commandBuffer-01831  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `srcImage` **must** not be a protected image

- <a href="#VUID-vkCmdCopyImageToBuffer2-commandBuffer-01832"
  id="VUID-vkCmdCopyImageToBuffer2-commandBuffer-01832"></a>
  VUID-vkCmdCopyImageToBuffer2-commandBuffer-01832  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstBuffer` **must** not be a protected buffer

- <a href="#VUID-vkCmdCopyImageToBuffer2-commandBuffer-01833"
  id="VUID-vkCmdCopyImageToBuffer2-commandBuffer-01833"></a>
  VUID-vkCmdCopyImageToBuffer2-commandBuffer-01833  
  If `commandBuffer` is a protected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstBuffer` **must** not be an unprotected buffer

- <a href="#VUID-vkCmdCopyImageToBuffer2-commandBuffer-07746"
  id="VUID-vkCmdCopyImageToBuffer2-commandBuffer-07746"></a>
  VUID-vkCmdCopyImageToBuffer2-commandBuffer-07746  
  If the queue family used to create the
  [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) which `commandBuffer` was
  allocated from does not support `VK_QUEUE_GRAPHICS_BIT` or
  `VK_QUEUE_COMPUTE_BIT`, the `bufferOffset` member of any element of
  `pCopyImageToBufferInfo->pRegions` **must** be a multiple of `4`

- <a href="#VUID-vkCmdCopyImageToBuffer2-imageOffset-07747"
  id="VUID-vkCmdCopyImageToBuffer2-imageOffset-07747"></a>
  VUID-vkCmdCopyImageToBuffer2-imageOffset-07747  
  The `imageOffset` and `imageExtent` members of each element of
  `pCopyImageToBufferInfo->pRegions` **must** respect the image transfer
  granularity requirements of `commandBuffer`’s command pool’s queue
  family, as described in
  [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html)

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCopyImageToBuffer2-commandBuffer-parameter"
  id="VUID-vkCmdCopyImageToBuffer2-commandBuffer-parameter"></a>
  VUID-vkCmdCopyImageToBuffer2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyImageToBuffer2-pCopyImageToBufferInfo-parameter"
  id="VUID-vkCmdCopyImageToBuffer2-pCopyImageToBufferInfo-parameter"></a>
  VUID-vkCmdCopyImageToBuffer2-pCopyImageToBufferInfo-parameter  
  `pCopyImageToBufferInfo` **must** be a valid pointer to a valid
  [VkCopyImageToBufferInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToBufferInfo2.html) structure

- <a href="#VUID-vkCmdCopyImageToBuffer2-commandBuffer-recording"
  id="VUID-vkCmdCopyImageToBuffer2-commandBuffer-recording"></a>
  VUID-vkCmdCopyImageToBuffer2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyImageToBuffer2-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyImageToBuffer2-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyImageToBuffer2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdCopyImageToBuffer2-renderpass"
  id="VUID-vkCmdCopyImageToBuffer2-renderpass"></a>
  VUID-vkCmdCopyImageToBuffer2-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyImageToBuffer2-videocoding"
  id="VUID-vkCmdCopyImageToBuffer2-videocoding"></a>
  VUID-vkCmdCopyImageToBuffer2-videocoding  
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
[VkCopyImageToBufferInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToBufferInfo2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyImageToBuffer2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
