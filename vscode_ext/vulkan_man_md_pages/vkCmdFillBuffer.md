# vkCmdFillBuffer(3) Manual Page

## Name

vkCmdFillBuffer - Fill a region of a buffer with a fixed value



## <a href="#_c_specification" class="anchor"></a>C Specification

To clear buffer data, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdFillBuffer(
    VkCommandBuffer                             commandBuffer,
    VkBuffer                                    dstBuffer,
    VkDeviceSize                                dstOffset,
    VkDeviceSize                                size,
    uint32_t                                    data);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `dstBuffer` is the buffer to be filled.

- `dstOffset` is the byte offset into the buffer at which to start
  filling, and **must** be a multiple of 4.

- `size` is the number of bytes to fill, and **must** be either a
  multiple of 4, or `VK_WHOLE_SIZE` to fill the range from `offset` to
  the end of the buffer. If `VK_WHOLE_SIZE` is used and the remaining
  size of the buffer is not a multiple of 4, then the nearest smaller
  multiple is used.

- `data` is the 4-byte word written repeatedly to the buffer to fill
  `size` bytes of data. The data word is written to memory according to
  the host endianness.

## <a href="#_description" class="anchor"></a>Description

`vkCmdFillBuffer` is treated as a “transfer” operation for the purposes
of synchronization barriers. The `VK_BUFFER_USAGE_TRANSFER_DST_BIT`
**must** be specified in `usage` of
[VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) in order for the buffer to
be compatible with `vkCmdFillBuffer`.

Valid Usage

- <a href="#VUID-vkCmdFillBuffer-dstOffset-00024"
  id="VUID-vkCmdFillBuffer-dstOffset-00024"></a>
  VUID-vkCmdFillBuffer-dstOffset-00024  
  `dstOffset` **must** be less than the size of `dstBuffer`

- <a href="#VUID-vkCmdFillBuffer-dstOffset-00025"
  id="VUID-vkCmdFillBuffer-dstOffset-00025"></a>
  VUID-vkCmdFillBuffer-dstOffset-00025  
  `dstOffset` **must** be a multiple of `4`

- <a href="#VUID-vkCmdFillBuffer-size-00026"
  id="VUID-vkCmdFillBuffer-size-00026"></a>
  VUID-vkCmdFillBuffer-size-00026  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be greater
  than `0`

- <a href="#VUID-vkCmdFillBuffer-size-00027"
  id="VUID-vkCmdFillBuffer-size-00027"></a>
  VUID-vkCmdFillBuffer-size-00027  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be less
  than or equal to the size of `dstBuffer` minus `dstOffset`

- <a href="#VUID-vkCmdFillBuffer-size-00028"
  id="VUID-vkCmdFillBuffer-size-00028"></a>
  VUID-vkCmdFillBuffer-size-00028  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be a
  multiple of `4`

- <a href="#VUID-vkCmdFillBuffer-dstBuffer-00029"
  id="VUID-vkCmdFillBuffer-dstBuffer-00029"></a>
  VUID-vkCmdFillBuffer-dstBuffer-00029  
  `dstBuffer` **must** have been created with
  `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-vkCmdFillBuffer-apiVersion-07894"
  id="VUID-vkCmdFillBuffer-apiVersion-07894"></a>
  VUID-vkCmdFillBuffer-apiVersion-07894  
  If the [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html) extension is
  not enabled and
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, the [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) that
  `commandBuffer` was allocated from **must** support graphics or
  compute operations

- <a href="#VUID-vkCmdFillBuffer-dstBuffer-00031"
  id="VUID-vkCmdFillBuffer-dstBuffer-00031"></a>
  VUID-vkCmdFillBuffer-dstBuffer-00031  
  If `dstBuffer` is non-sparse then it **must** be bound completely and
  contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdFillBuffer-commandBuffer-01811"
  id="VUID-vkCmdFillBuffer-commandBuffer-01811"></a>
  VUID-vkCmdFillBuffer-commandBuffer-01811  
  If `commandBuffer` is an unprotected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, `dstBuffer` **must** not be a protected buffer

- <a href="#VUID-vkCmdFillBuffer-commandBuffer-01812"
  id="VUID-vkCmdFillBuffer-commandBuffer-01812"></a>
  VUID-vkCmdFillBuffer-commandBuffer-01812  
  If `commandBuffer` is a protected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, `dstBuffer` **must** not be an unprotected buffer

Valid Usage (Implicit)

- <a href="#VUID-vkCmdFillBuffer-commandBuffer-parameter"
  id="VUID-vkCmdFillBuffer-commandBuffer-parameter"></a>
  VUID-vkCmdFillBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdFillBuffer-dstBuffer-parameter"
  id="VUID-vkCmdFillBuffer-dstBuffer-parameter"></a>
  VUID-vkCmdFillBuffer-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-vkCmdFillBuffer-commandBuffer-recording"
  id="VUID-vkCmdFillBuffer-commandBuffer-recording"></a>
  VUID-vkCmdFillBuffer-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdFillBuffer-commandBuffer-cmdpool"
  id="VUID-vkCmdFillBuffer-commandBuffer-cmdpool"></a>
  VUID-vkCmdFillBuffer-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics or compute operations

- <a href="#VUID-vkCmdFillBuffer-renderpass"
  id="VUID-vkCmdFillBuffer-renderpass"></a>
  VUID-vkCmdFillBuffer-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdFillBuffer-videocoding"
  id="VUID-vkCmdFillBuffer-videocoding"></a>
  VUID-vkCmdFillBuffer-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdFillBuffer-commonparent"
  id="VUID-vkCmdFillBuffer-commonparent"></a>
  VUID-vkCmdFillBuffer-commonparent  
  Both of `commandBuffer`, and `dstBuffer` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdFillBuffer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
