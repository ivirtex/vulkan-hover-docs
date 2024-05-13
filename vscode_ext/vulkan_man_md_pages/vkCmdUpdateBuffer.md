# vkCmdUpdateBuffer(3) Manual Page

## Name

vkCmdUpdateBuffer - Update a buffer's contents from host memory



## <a href="#_c_specification" class="anchor"></a>C Specification

To update buffer data inline in a command buffer, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdUpdateBuffer(
    VkCommandBuffer                             commandBuffer,
    VkBuffer                                    dstBuffer,
    VkDeviceSize                                dstOffset,
    VkDeviceSize                                dataSize,
    const void*                                 pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `dstBuffer` is a handle to the buffer to be updated.

- `dstOffset` is the byte offset into the buffer to start updating, and
  **must** be a multiple of 4.

- `dataSize` is the number of bytes to update, and **must** be a
  multiple of 4.

- `pData` is a pointer to the source data for the buffer update, and
  **must** be at least `dataSize` bytes in size.

## <a href="#_description" class="anchor"></a>Description

`dataSize` **must** be less than or equal to 65536 bytes. For larger
updates, applications **can** use buffer to buffer <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies-buffers"
target="_blank" rel="noopener">copies</a>.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Buffer updates performed with <code>vkCmdUpdateBuffer</code> first
copy the data into command buffer memory when the command is recorded
(which requires additional storage and may incur an additional
allocation), and then copy the data from the command buffer into
<code>dstBuffer</code> when the command is executed on a device.</p>
<p>The additional cost of this functionality compared to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies-buffers"
target="_blank" rel="noopener">buffer to buffer copies</a> means it is
only recommended for very small amounts of data, and is why it is
limited to only 65536 bytes.</p>
<p>Applications <strong>can</strong> work around this by issuing
multiple <code>vkCmdUpdateBuffer</code> commands to different ranges of
the same buffer, but it is strongly recommended that they
<strong>should</strong> not.</p></td>
</tr>
</tbody>
</table>

The source data is copied from the user pointer to the command buffer
when the command is called.

`vkCmdUpdateBuffer` is only allowed outside of a render pass. This
command is treated as a “transfer” operation for the purposes of
synchronization barriers. The `VK_BUFFER_USAGE_TRANSFER_DST_BIT`
**must** be specified in `usage` of
[VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) in order for the buffer to
be compatible with `vkCmdUpdateBuffer`.

Valid Usage

- <a href="#VUID-vkCmdUpdateBuffer-dstOffset-00032"
  id="VUID-vkCmdUpdateBuffer-dstOffset-00032"></a>
  VUID-vkCmdUpdateBuffer-dstOffset-00032  
  `dstOffset` **must** be less than the size of `dstBuffer`

- <a href="#VUID-vkCmdUpdateBuffer-dataSize-00033"
  id="VUID-vkCmdUpdateBuffer-dataSize-00033"></a>
  VUID-vkCmdUpdateBuffer-dataSize-00033  
  `dataSize` **must** be less than or equal to the size of `dstBuffer`
  minus `dstOffset`

- <a href="#VUID-vkCmdUpdateBuffer-dstBuffer-00034"
  id="VUID-vkCmdUpdateBuffer-dstBuffer-00034"></a>
  VUID-vkCmdUpdateBuffer-dstBuffer-00034  
  `dstBuffer` **must** have been created with
  `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-vkCmdUpdateBuffer-dstBuffer-00035"
  id="VUID-vkCmdUpdateBuffer-dstBuffer-00035"></a>
  VUID-vkCmdUpdateBuffer-dstBuffer-00035  
  If `dstBuffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdUpdateBuffer-dstOffset-00036"
  id="VUID-vkCmdUpdateBuffer-dstOffset-00036"></a>
  VUID-vkCmdUpdateBuffer-dstOffset-00036  
  `dstOffset` **must** be a multiple of `4`

- <a href="#VUID-vkCmdUpdateBuffer-dataSize-00037"
  id="VUID-vkCmdUpdateBuffer-dataSize-00037"></a>
  VUID-vkCmdUpdateBuffer-dataSize-00037  
  `dataSize` **must** be less than or equal to `65536`

- <a href="#VUID-vkCmdUpdateBuffer-dataSize-00038"
  id="VUID-vkCmdUpdateBuffer-dataSize-00038"></a>
  VUID-vkCmdUpdateBuffer-dataSize-00038  
  `dataSize` **must** be a multiple of `4`

- <a href="#VUID-vkCmdUpdateBuffer-commandBuffer-01813"
  id="VUID-vkCmdUpdateBuffer-commandBuffer-01813"></a>
  VUID-vkCmdUpdateBuffer-commandBuffer-01813  
  If `commandBuffer` is an unprotected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, `dstBuffer` **must** not be a protected buffer

- <a href="#VUID-vkCmdUpdateBuffer-commandBuffer-01814"
  id="VUID-vkCmdUpdateBuffer-commandBuffer-01814"></a>
  VUID-vkCmdUpdateBuffer-commandBuffer-01814  
  If `commandBuffer` is a protected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, `dstBuffer` **must** not be an unprotected buffer

Valid Usage (Implicit)

- <a href="#VUID-vkCmdUpdateBuffer-commandBuffer-parameter"
  id="VUID-vkCmdUpdateBuffer-commandBuffer-parameter"></a>
  VUID-vkCmdUpdateBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdUpdateBuffer-dstBuffer-parameter"
  id="VUID-vkCmdUpdateBuffer-dstBuffer-parameter"></a>
  VUID-vkCmdUpdateBuffer-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-vkCmdUpdateBuffer-pData-parameter"
  id="VUID-vkCmdUpdateBuffer-pData-parameter"></a>
  VUID-vkCmdUpdateBuffer-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes

- <a href="#VUID-vkCmdUpdateBuffer-commandBuffer-recording"
  id="VUID-vkCmdUpdateBuffer-commandBuffer-recording"></a>
  VUID-vkCmdUpdateBuffer-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdUpdateBuffer-commandBuffer-cmdpool"
  id="VUID-vkCmdUpdateBuffer-commandBuffer-cmdpool"></a>
  VUID-vkCmdUpdateBuffer-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdUpdateBuffer-renderpass"
  id="VUID-vkCmdUpdateBuffer-renderpass"></a>
  VUID-vkCmdUpdateBuffer-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdUpdateBuffer-videocoding"
  id="VUID-vkCmdUpdateBuffer-videocoding"></a>
  VUID-vkCmdUpdateBuffer-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdUpdateBuffer-dataSize-arraylength"
  id="VUID-vkCmdUpdateBuffer-dataSize-arraylength"></a>
  VUID-vkCmdUpdateBuffer-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

- <a href="#VUID-vkCmdUpdateBuffer-commonparent"
  id="VUID-vkCmdUpdateBuffer-commonparent"></a>
  VUID-vkCmdUpdateBuffer-commonparent  
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
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdUpdateBuffer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
