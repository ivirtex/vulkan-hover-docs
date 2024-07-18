# vkCmdCopyMemoryIndirectNV(3) Manual Page

## Name

vkCmdCopyMemoryIndirectNV - Copy data between memory regions



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy data between two memory regions by specifying copy parameters
indirectly in a buffer, call:

``` c
// Provided by VK_NV_copy_memory_indirect
void vkCmdCopyMemoryIndirectNV(
    VkCommandBuffer                             commandBuffer,
    VkDeviceAddress                             copyBufferAddress,
    uint32_t                                    copyCount,
    uint32_t                                    stride);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `copyBufferAddress` is the buffer address specifying the copy
  parameters. This buffer is laid out in memory as an array of
  [VkCopyMemoryIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryIndirectCommandNV.html)
  structures.

- `copyCount` is the number of copies to execute, and can be zero.

- `stride` is the stride in bytes between successive sets of copy
  parameters.

## <a href="#_description" class="anchor"></a>Description

Each region read from `copyBufferAddress` is copied from the source
region to the specified destination region. The results are undefined if
any of the source and destination regions overlap in memory.

Valid Usage

- <a href="#VUID-vkCmdCopyMemoryIndirectNV-None-07653"
  id="VUID-vkCmdCopyMemoryIndirectNV-None-07653"></a>
  VUID-vkCmdCopyMemoryIndirectNV-None-07653  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-indirectCopy"
  target="_blank" rel="noopener"><code>indirectCopy</code></a> feature
  **must** be enabled

- <a href="#VUID-vkCmdCopyMemoryIndirectNV-copyBufferAddress-07654"
  id="VUID-vkCmdCopyMemoryIndirectNV-copyBufferAddress-07654"></a>
  VUID-vkCmdCopyMemoryIndirectNV-copyBufferAddress-07654  
  `copyBufferAddress` **must** be 4 byte aligned

- <a href="#VUID-vkCmdCopyMemoryIndirectNV-stride-07655"
  id="VUID-vkCmdCopyMemoryIndirectNV-stride-07655"></a>
  VUID-vkCmdCopyMemoryIndirectNV-stride-07655  
  `stride` **must** be a multiple of `4` and **must** be greater than or
  equal to sizeof(`VkCopyMemoryIndirectCommandNV`)

- <a href="#VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-07656"
  id="VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-07656"></a>
  VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-07656  
  The [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) that `commandBuffer` was
  allocated from **must** support at least one of the
  [VkPhysicalDeviceCopyMemoryIndirectPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCopyMemoryIndirectPropertiesNV.html)::`supportedQueues`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-parameter"
  id="VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-parameter"></a>
  VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-recording"
  id="VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-recording"></a>
  VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdCopyMemoryIndirectNV-renderpass"
  id="VUID-vkCmdCopyMemoryIndirectNV-renderpass"></a>
  VUID-vkCmdCopyMemoryIndirectNV-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyMemoryIndirectNV-videocoding"
  id="VUID-vkCmdCopyMemoryIndirectNV-videocoding"></a>
  VUID-vkCmdCopyMemoryIndirectNV-videocoding  
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

[VK_NV_copy_memory_indirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_copy_memory_indirect.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyMemoryIndirectNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
