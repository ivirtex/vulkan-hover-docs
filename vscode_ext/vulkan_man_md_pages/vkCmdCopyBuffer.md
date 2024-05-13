# vkCmdCopyBuffer(3) Manual Page

## Name

vkCmdCopyBuffer - Copy data between buffer regions



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy data between buffer objects, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdCopyBuffer(
    VkCommandBuffer                             commandBuffer,
    VkBuffer                                    srcBuffer,
    VkBuffer                                    dstBuffer,
    uint32_t                                    regionCount,
    const VkBufferCopy*                         pRegions);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `srcBuffer` is the source buffer.

- `dstBuffer` is the destination buffer.

- `regionCount` is the number of regions to copy.

- `pRegions` is a pointer to an array of
  [VkBufferCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCopy.html) structures specifying the regions to
  copy.

## <a href="#_description" class="anchor"></a>Description

Each source region specified by `pRegions` is copied from the source
buffer to the destination region of the destination buffer. If any of
the specified regions in `srcBuffer` overlaps in memory with any of the
specified regions in `dstBuffer`, values read from those overlapping
regions are undefined.

Valid Usage

- <a href="#VUID-vkCmdCopyBuffer-commandBuffer-01822"
  id="VUID-vkCmdCopyBuffer-commandBuffer-01822"></a>
  VUID-vkCmdCopyBuffer-commandBuffer-01822  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `srcBuffer` **must** not be a protected buffer

- <a href="#VUID-vkCmdCopyBuffer-commandBuffer-01823"
  id="VUID-vkCmdCopyBuffer-commandBuffer-01823"></a>
  VUID-vkCmdCopyBuffer-commandBuffer-01823  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstBuffer` **must** not be a protected buffer

- <a href="#VUID-vkCmdCopyBuffer-commandBuffer-01824"
  id="VUID-vkCmdCopyBuffer-commandBuffer-01824"></a>
  VUID-vkCmdCopyBuffer-commandBuffer-01824  
  If `commandBuffer` is a protected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstBuffer` **must** not be an unprotected buffer

<!-- -->

- <a href="#VUID-vkCmdCopyBuffer-srcOffset-00113"
  id="VUID-vkCmdCopyBuffer-srcOffset-00113"></a>
  VUID-vkCmdCopyBuffer-srcOffset-00113  
  The `srcOffset` member of each element of `pRegions` **must** be less
  than the size of `srcBuffer`

- <a href="#VUID-vkCmdCopyBuffer-dstOffset-00114"
  id="VUID-vkCmdCopyBuffer-dstOffset-00114"></a>
  VUID-vkCmdCopyBuffer-dstOffset-00114  
  The `dstOffset` member of each element of `pRegions` **must** be less
  than the size of `dstBuffer`

- <a href="#VUID-vkCmdCopyBuffer-size-00115"
  id="VUID-vkCmdCopyBuffer-size-00115"></a>
  VUID-vkCmdCopyBuffer-size-00115  
  The `size` member of each element of `pRegions` **must** be less than
  or equal to the size of `srcBuffer` minus `srcOffset`

- <a href="#VUID-vkCmdCopyBuffer-size-00116"
  id="VUID-vkCmdCopyBuffer-size-00116"></a>
  VUID-vkCmdCopyBuffer-size-00116  
  The `size` member of each element of `pRegions` **must** be less than
  or equal to the size of `dstBuffer` minus `dstOffset`

- <a href="#VUID-vkCmdCopyBuffer-pRegions-00117"
  id="VUID-vkCmdCopyBuffer-pRegions-00117"></a>
  VUID-vkCmdCopyBuffer-pRegions-00117  
  The union of the source regions, and the union of the destination
  regions, specified by the elements of `pRegions`, **must** not overlap
  in memory

- <a href="#VUID-vkCmdCopyBuffer-srcBuffer-00118"
  id="VUID-vkCmdCopyBuffer-srcBuffer-00118"></a>
  VUID-vkCmdCopyBuffer-srcBuffer-00118  
  `srcBuffer` **must** have been created with
  `VK_BUFFER_USAGE_TRANSFER_SRC_BIT` usage flag

- <a href="#VUID-vkCmdCopyBuffer-srcBuffer-00119"
  id="VUID-vkCmdCopyBuffer-srcBuffer-00119"></a>
  VUID-vkCmdCopyBuffer-srcBuffer-00119  
  If `srcBuffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdCopyBuffer-dstBuffer-00120"
  id="VUID-vkCmdCopyBuffer-dstBuffer-00120"></a>
  VUID-vkCmdCopyBuffer-dstBuffer-00120  
  `dstBuffer` **must** have been created with
  `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-vkCmdCopyBuffer-dstBuffer-00121"
  id="VUID-vkCmdCopyBuffer-dstBuffer-00121"></a>
  VUID-vkCmdCopyBuffer-dstBuffer-00121  
  If `dstBuffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCopyBuffer-commandBuffer-parameter"
  id="VUID-vkCmdCopyBuffer-commandBuffer-parameter"></a>
  VUID-vkCmdCopyBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyBuffer-srcBuffer-parameter"
  id="VUID-vkCmdCopyBuffer-srcBuffer-parameter"></a>
  VUID-vkCmdCopyBuffer-srcBuffer-parameter  
  `srcBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-vkCmdCopyBuffer-dstBuffer-parameter"
  id="VUID-vkCmdCopyBuffer-dstBuffer-parameter"></a>
  VUID-vkCmdCopyBuffer-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-vkCmdCopyBuffer-pRegions-parameter"
  id="VUID-vkCmdCopyBuffer-pRegions-parameter"></a>
  VUID-vkCmdCopyBuffer-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkBufferCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCopy.html) structures

- <a href="#VUID-vkCmdCopyBuffer-commandBuffer-recording"
  id="VUID-vkCmdCopyBuffer-commandBuffer-recording"></a>
  VUID-vkCmdCopyBuffer-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyBuffer-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyBuffer-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyBuffer-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdCopyBuffer-renderpass"
  id="VUID-vkCmdCopyBuffer-renderpass"></a>
  VUID-vkCmdCopyBuffer-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyBuffer-videocoding"
  id="VUID-vkCmdCopyBuffer-videocoding"></a>
  VUID-vkCmdCopyBuffer-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdCopyBuffer-regionCount-arraylength"
  id="VUID-vkCmdCopyBuffer-regionCount-arraylength"></a>
  VUID-vkCmdCopyBuffer-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

- <a href="#VUID-vkCmdCopyBuffer-commonparent"
  id="VUID-vkCmdCopyBuffer-commonparent"></a>
  VUID-vkCmdCopyBuffer-commonparent  
  Each of `commandBuffer`, `dstBuffer`, and `srcBuffer` **must** have
  been created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
[VkBufferCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCopy.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyBuffer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
