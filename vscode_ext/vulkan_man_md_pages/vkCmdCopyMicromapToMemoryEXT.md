# vkCmdCopyMicromapToMemoryEXT(3) Manual Page

## Name

vkCmdCopyMicromapToMemoryEXT - Copy a micromap to device memory



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy a micromap to device memory call:

``` c
// Provided by VK_EXT_opacity_micromap
void vkCmdCopyMicromapToMemoryEXT(
    VkCommandBuffer                             commandBuffer,
    const VkCopyMicromapToMemoryInfoEXT*        pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pInfo` is an a pointer to a
  [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapToMemoryInfoEXT.html)
  structure defining the copy operation.

## <a href="#_description" class="anchor"></a>Description

Accesses to `pInfo->src` **must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
`VK_ACCESS_2_MICROMAP_READ_BIT_EXT`. Accesses to the buffer indicated by
`pInfo->dst.deviceAddress` **must** be synchronized with the
`VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` pipeline stage and an
access type of `VK_ACCESS_TRANSFER_WRITE_BIT`.

This command produces the same results as
[vkCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMicromapToMemoryEXT.html), but writes
its result to a device address, and is executed on the device rather
than the host. The output **may** not necessarily be bit-for-bit
identical, but it can be equally used by either
[vkCmdCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToMicromapEXT.html) or
[vkCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMemoryToMicromapEXT.html).

The defined header structure for the serialized data consists of:

- `VK_UUID_SIZE` bytes of data matching
  `VkPhysicalDeviceIDProperties`::`driverUUID`

- `VK_UUID_SIZE` bytes of data identifying the compatibility for
  comparison using
  [vkGetDeviceMicromapCompatibilityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMicromapCompatibilityEXT.html)
  The serialized data is written to the buffer (or read from the buffer)
  according to the host endianness.

Valid Usage

- <a href="#VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07536"
  id="VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07536"></a>
  VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07536  
  `pInfo->dst.deviceAddress` **must** be a valid device address for a
  buffer bound to device memory

- <a href="#VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07537"
  id="VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07537"></a>
  VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07537  
  `pInfo->dst.deviceAddress` **must** be aligned to `256` bytes

- <a href="#VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07538"
  id="VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07538"></a>
  VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-07538  
  If the buffer pointed to by `pInfo->dst.deviceAddress` is non-sparse
  then it **must** be bound completely and contiguously to a single
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdCopyMicromapToMemoryEXT-buffer-07539"
  id="VUID-vkCmdCopyMicromapToMemoryEXT-buffer-07539"></a>
  VUID-vkCmdCopyMicromapToMemoryEXT-buffer-07539  
  The `buffer` used to create `pInfo->src` **must** be bound to device
  memory

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-parameter"
  id="VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-parameter"></a>
  VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-parameter"
  id="VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-parameter"></a>
  VUID-vkCmdCopyMicromapToMemoryEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapToMemoryInfoEXT.html)
  structure

- <a href="#VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-recording"
  id="VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-recording"></a>
  VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyMicromapToMemoryEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdCopyMicromapToMemoryEXT-renderpass"
  id="VUID-vkCmdCopyMicromapToMemoryEXT-renderpass"></a>
  VUID-vkCmdCopyMicromapToMemoryEXT-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyMicromapToMemoryEXT-videocoding"
  id="VUID-vkCmdCopyMicromapToMemoryEXT-videocoding"></a>
  VUID-vkCmdCopyMicromapToMemoryEXT-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapToMemoryInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyMicromapToMemoryEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
