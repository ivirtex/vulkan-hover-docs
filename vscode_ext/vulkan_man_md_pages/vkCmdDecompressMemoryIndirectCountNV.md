# vkCmdDecompressMemoryIndirectCountNV(3) Manual Page

## Name

vkCmdDecompressMemoryIndirectCountNV - Indirect decompress data between
memory regions



## <a href="#_c_specification" class="anchor"></a>C Specification

To decompress data between one or more memory regions by specifying
decompression parameters indirectly in a buffer, call:

``` c
// Provided by VK_NV_memory_decompression
void vkCmdDecompressMemoryIndirectCountNV(
    VkCommandBuffer                             commandBuffer,
    VkDeviceAddress                             indirectCommandsAddress,
    VkDeviceAddress                             indirectCommandsCountAddress,
    uint32_t                                    stride);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `indirectCommandsAddress` is the device address containing
  decompression parameters laid out as an array of
  [VkDecompressMemoryRegionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDecompressMemoryRegionNV.html)
  structures.

- `indirectCommandsCountAddress` is the device address containing the
  decompression count.

- `stride` is the byte stride between successive sets of decompression
  parameters located starting from `indirectCommandsAddress`.

## <a href="#_description" class="anchor"></a>Description

Each region specified in `indirectCommandsAddress` is decompressed from
the source to destination region based on the specified decompression
method.

Valid Usage

- <a href="#VUID-vkCmdDecompressMemoryIndirectCountNV-None-07692"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-None-07692"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-None-07692  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-memoryDecompression"
  target="_blank" rel="noopener"><code>memoryDecompression</code></a>
  feature **must** be enabled

- <a
  href="#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsAddress-07693"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsAddress-07693"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsAddress-07693  
  If `indirectCommandsAddress` comes from a non-sparse buffer then it
  **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a
  href="#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsAddress-07694"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsAddress-07694"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsAddress-07694  
  The [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) that `indirectCommandsAddress` comes
  from **must** have been created with the
  `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` bit set

- <a href="#VUID-vkCmdDecompressMemoryIndirectCountNV-offset-07695"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-offset-07695"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-offset-07695  
  `offset` **must** be a multiple of `4`

- <a
  href="#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07696"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07696"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07696  
  If `indirectCommandsCountAddress` comes from a non-sparse buffer then
  it **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a
  href="#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07697"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07697"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07697  
  The [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) that `indirectCommandsCountAddress`
  comes from **must** have been created with the
  `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` bit set

- <a
  href="#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07698"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07698"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07698  
  `indirectCommandsCountAddress` **must** be a multiple of `4`

- <a
  href="#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07699"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07699"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07699  
  The count stored in `indirectCommandsCountAddress` **must** be less
  than or equal to
  `VkPhysicalDeviceMemoryDecompressionPropertiesNV`::`maxDecompressionIndirectCount`

- <a href="#VUID-vkCmdDecompressMemoryIndirectCountNV-stride-07700"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-stride-07700"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-stride-07700  
  `stride` **must** be a multiple of `4` and **must** be greater than or
  equal to sizeof(`VkDecompressMemoryRegionNV`)

- <a
  href="#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07701"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07701"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07701  
  If the count stored in `indirectCommandsCountAddress` is equal to `1`,
  (`offset` + sizeof(`VkDecompressMemoryRegionNV`)) **must** be less
  than or equal to the size of the [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) that
  `indirectCommandsAddress` comes from

- <a
  href="#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07702"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07702"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07702  
  If the count stored in `indirectCommandsCountAddress` is greater than
  `1`, `indirectCommandsAddress` +
  sizeof(`VkDecompressMemoryRegionNV`) + (`stride` × (count stored in
  `countBuffer` - 1)) **must** be less than or equal to the last valid
  address in the [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) that
  `indirectCommandsAddress` was created from

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-parameter"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-parameter"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-recording"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-recording"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-cmdpool"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdDecompressMemoryIndirectCountNV-renderpass"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-renderpass"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdDecompressMemoryIndirectCountNV-videocoding"
  id="VUID-vkCmdDecompressMemoryIndirectCountNV-videocoding"></a>
  VUID-vkCmdDecompressMemoryIndirectCountNV-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_memory_decompression](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_memory_decompression.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdDecompressMemoryIndirectCountNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
