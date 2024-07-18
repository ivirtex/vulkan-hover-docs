# vkCmdCopyMemoryToMicromapEXT(3) Manual Page

## Name

vkCmdCopyMemoryToMicromapEXT - Copy device memory to a micromap



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy device memory to a micromap call:

``` c
// Provided by VK_EXT_opacity_micromap
void vkCmdCopyMemoryToMicromapEXT(
    VkCommandBuffer                             commandBuffer,
    const VkCopyMemoryToMicromapInfoEXT*        pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pInfo` is a pointer to a
  [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapToMemoryInfoEXT.html)
  structure defining the copy operation.

## <a href="#_description" class="anchor"></a>Description

Accesses to `pInfo->dst` **must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
`VK_ACCESS_2_MICROMAP_READ_BIT_EXT`. Accesses to the buffer indicated by
`pInfo->src.deviceAddress` **must** be synchronized with the
`VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` pipeline stage and an
access type of `VK_ACCESS_TRANSFER_READ_BIT`.

This command can accept micromaps produced by either
[vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMicromapToMemoryEXT.html) or
[vkCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMicromapToMemoryEXT.html).

Valid Usage

- <a href="#VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07543"
  id="VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07543"></a>
  VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07543  
  `pInfo->src.deviceAddress` **must** be a valid device address for a
  buffer bound to device memory

- <a href="#VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07544"
  id="VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07544"></a>
  VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07544  
  `pInfo->src.deviceAddress` **must** be aligned to `256` bytes

- <a href="#VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07545"
  id="VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07545"></a>
  VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-07545  
  If the buffer pointed to by `pInfo->src.deviceAddress` is non-sparse
  then it **must** be bound completely and contiguously to a single
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdCopyMemoryToMicromapEXT-buffer-07546"
  id="VUID-vkCmdCopyMemoryToMicromapEXT-buffer-07546"></a>
  VUID-vkCmdCopyMemoryToMicromapEXT-buffer-07546  
  The `buffer` used to create `pInfo->dst` **must** be bound to device
  memory

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-parameter"
  id="VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-parameter"></a>
  VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-parameter"
  id="VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-parameter"></a>
  VUID-vkCmdCopyMemoryToMicromapEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToMicromapInfoEXT.html)
  structure

- <a href="#VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-recording"
  id="VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-recording"></a>
  VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyMemoryToMicromapEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdCopyMemoryToMicromapEXT-renderpass"
  id="VUID-vkCmdCopyMemoryToMicromapEXT-renderpass"></a>
  VUID-vkCmdCopyMemoryToMicromapEXT-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyMemoryToMicromapEXT-videocoding"
  id="VUID-vkCmdCopyMemoryToMicromapEXT-videocoding"></a>
  VUID-vkCmdCopyMemoryToMicromapEXT-videocoding  
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
[VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToMicromapInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyMemoryToMicromapEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
