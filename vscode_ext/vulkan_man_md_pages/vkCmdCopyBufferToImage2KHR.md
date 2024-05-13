# vkCmdCopyBufferToImage2(3) Manual Page

## Name

vkCmdCopyBufferToImage2 - Copy data from a buffer into an image



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy data from a buffer object to an image object, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdCopyBufferToImage2(
    VkCommandBuffer                             commandBuffer,
    const VkCopyBufferToImageInfo2*             pCopyBufferToImageInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_copy_commands2
void vkCmdCopyBufferToImage2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyBufferToImageInfo2*             pCopyBufferToImageInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pCopyBufferToImageInfo` is a pointer to a
  [VkCopyBufferToImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyBufferToImageInfo2.html) structure
  describing the copy parameters.

## <a href="#_description" class="anchor"></a>Description

This command is functionally identical to
[vkCmdCopyBufferToImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage.html), but includes
extensible sub-structures that include `sType` and `pNext` parameters,
allowing them to be more easily extended.

Valid Usage

- <a href="#VUID-vkCmdCopyBufferToImage2-commandBuffer-01828"
  id="VUID-vkCmdCopyBufferToImage2-commandBuffer-01828"></a>
  VUID-vkCmdCopyBufferToImage2-commandBuffer-01828  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `srcBuffer` **must** not be a protected buffer

- <a href="#VUID-vkCmdCopyBufferToImage2-commandBuffer-01829"
  id="VUID-vkCmdCopyBufferToImage2-commandBuffer-01829"></a>
  VUID-vkCmdCopyBufferToImage2-commandBuffer-01829  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be a protected image

- <a href="#VUID-vkCmdCopyBufferToImage2-commandBuffer-01830"
  id="VUID-vkCmdCopyBufferToImage2-commandBuffer-01830"></a>
  VUID-vkCmdCopyBufferToImage2-commandBuffer-01830  
  If `commandBuffer` is a protected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be an unprotected image

- <a href="#VUID-vkCmdCopyBufferToImage2-commandBuffer-07737"
  id="VUID-vkCmdCopyBufferToImage2-commandBuffer-07737"></a>
  VUID-vkCmdCopyBufferToImage2-commandBuffer-07737  
  If the queue family used to create the
  [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) which `commandBuffer` was
  allocated from does not support `VK_QUEUE_GRAPHICS_BIT` or
  `VK_QUEUE_COMPUTE_BIT`, the `bufferOffset` member of any element of
  `pCopyBufferToImageInfo->pRegions` **must** be a multiple of `4`

- <a href="#VUID-vkCmdCopyBufferToImage2-imageOffset-07738"
  id="VUID-vkCmdCopyBufferToImage2-imageOffset-07738"></a>
  VUID-vkCmdCopyBufferToImage2-imageOffset-07738  
  The `imageOffset` and `imageExtent` members of each element of
  `pCopyBufferToImageInfo->pRegions` **must** respect the image transfer
  granularity requirements of `commandBuffer`’s command pool’s queue
  family, as described in
  [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html)

- <a href="#VUID-vkCmdCopyBufferToImage2-commandBuffer-07739"
  id="VUID-vkCmdCopyBufferToImage2-commandBuffer-07739"></a>
  VUID-vkCmdCopyBufferToImage2-commandBuffer-07739  
  If the queue family used to create the
  [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) which `commandBuffer` was
  allocated from does not support `VK_QUEUE_GRAPHICS_BIT`, for each
  element of `pCopyBufferToImageInfo->pRegions`, the `aspectMask` member
  of `imageSubresource` **must** not be `VK_IMAGE_ASPECT_DEPTH_BIT` or
  `VK_IMAGE_ASPECT_STENCIL_BIT`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCopyBufferToImage2-commandBuffer-parameter"
  id="VUID-vkCmdCopyBufferToImage2-commandBuffer-parameter"></a>
  VUID-vkCmdCopyBufferToImage2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyBufferToImage2-pCopyBufferToImageInfo-parameter"
  id="VUID-vkCmdCopyBufferToImage2-pCopyBufferToImageInfo-parameter"></a>
  VUID-vkCmdCopyBufferToImage2-pCopyBufferToImageInfo-parameter  
  `pCopyBufferToImageInfo` **must** be a valid pointer to a valid
  [VkCopyBufferToImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyBufferToImageInfo2.html) structure

- <a href="#VUID-vkCmdCopyBufferToImage2-commandBuffer-recording"
  id="VUID-vkCmdCopyBufferToImage2-commandBuffer-recording"></a>
  VUID-vkCmdCopyBufferToImage2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyBufferToImage2-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyBufferToImage2-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyBufferToImage2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdCopyBufferToImage2-renderpass"
  id="VUID-vkCmdCopyBufferToImage2-renderpass"></a>
  VUID-vkCmdCopyBufferToImage2-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyBufferToImage2-videocoding"
  id="VUID-vkCmdCopyBufferToImage2-videocoding"></a>
  VUID-vkCmdCopyBufferToImage2-videocoding  
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
[VkCopyBufferToImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyBufferToImageInfo2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyBufferToImage2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
