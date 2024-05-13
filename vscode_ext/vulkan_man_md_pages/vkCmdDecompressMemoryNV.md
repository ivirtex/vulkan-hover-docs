# vkCmdDecompressMemoryNV(3) Manual Page

## Name

vkCmdDecompressMemoryNV - Decompress data between memory regions



## <a href="#_c_specification" class="anchor"></a>C Specification

To decompress data between one or more memory regions call:

``` c
// Provided by VK_NV_memory_decompression
void vkCmdDecompressMemoryNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    decompressRegionCount,
    const VkDecompressMemoryRegionNV*           pDecompressMemoryRegions);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `decompressRegionCount` is the number of memory regions to decompress.

- `pDecompressMemoryRegions` is a pointer to an array of
  `decompressRegionCount`
  [VkDecompressMemoryRegionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDecompressMemoryRegionNV.html)
  structures specifying decompression parameters.

## <a href="#_description" class="anchor"></a>Description

Each region specified in `pDecompressMemoryRegions` is decompressed from
the source to destination region based on the specified decompression
method.

Valid Usage

- <a href="#VUID-vkCmdDecompressMemoryNV-None-07684"
  id="VUID-vkCmdDecompressMemoryNV-None-07684"></a>
  VUID-vkCmdDecompressMemoryNV-None-07684  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-memoryDecompression"
  target="_blank" rel="noopener"><code>memoryDecompression</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCmdDecompressMemoryNV-commandBuffer-parameter"
  id="VUID-vkCmdDecompressMemoryNV-commandBuffer-parameter"></a>
  VUID-vkCmdDecompressMemoryNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdDecompressMemoryNV-pDecompressMemoryRegions-parameter"
  id="VUID-vkCmdDecompressMemoryNV-pDecompressMemoryRegions-parameter"></a>
  VUID-vkCmdDecompressMemoryNV-pDecompressMemoryRegions-parameter  
  `pDecompressMemoryRegions` **must** be a valid pointer to an array of
  `decompressRegionCount` valid
  [VkDecompressMemoryRegionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDecompressMemoryRegionNV.html)
  structures

- <a href="#VUID-vkCmdDecompressMemoryNV-commandBuffer-recording"
  id="VUID-vkCmdDecompressMemoryNV-commandBuffer-recording"></a>
  VUID-vkCmdDecompressMemoryNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdDecompressMemoryNV-commandBuffer-cmdpool"
  id="VUID-vkCmdDecompressMemoryNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdDecompressMemoryNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdDecompressMemoryNV-renderpass"
  id="VUID-vkCmdDecompressMemoryNV-renderpass"></a>
  VUID-vkCmdDecompressMemoryNV-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdDecompressMemoryNV-videocoding"
  id="VUID-vkCmdDecompressMemoryNV-videocoding"></a>
  VUID-vkCmdDecompressMemoryNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdDecompressMemoryNV-decompressRegionCount-arraylength"
  id="VUID-vkCmdDecompressMemoryNV-decompressRegionCount-arraylength"></a>
  VUID-vkCmdDecompressMemoryNV-decompressRegionCount-arraylength  
  `decompressRegionCount` **must** be greater than `0`

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
[VkDecompressMemoryRegionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDecompressMemoryRegionNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdDecompressMemoryNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
