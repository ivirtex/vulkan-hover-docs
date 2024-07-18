# vkCmdCopyAccelerationStructureKHR(3) Manual Page

## Name

vkCmdCopyAccelerationStructureKHR - Copy an acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy an acceleration structure call:

``` c
// Provided by VK_KHR_acceleration_structure
void vkCmdCopyAccelerationStructureKHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyAccelerationStructureInfoKHR*   pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pInfo` is a pointer to a
  [VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureInfoKHR.html)
  structure defining the copy operation.

## <a href="#_description" class="anchor"></a>Description

This command copies the `pInfo->src` acceleration structure to the
`pInfo->dst` acceleration structure in the manner specified by
`pInfo->mode`.

Accesses to `pInfo->src` and `pInfo->dst` **must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> or the
`VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a>, and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
`VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` or
`VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR` as appropriate.

Valid Usage

- <a
  href="#VUID-vkCmdCopyAccelerationStructureKHR-accelerationStructure-08925"
  id="VUID-vkCmdCopyAccelerationStructureKHR-accelerationStructure-08925"></a>
  VUID-vkCmdCopyAccelerationStructureKHR-accelerationStructure-08925  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructure"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructure</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdCopyAccelerationStructureKHR-buffer-03737"
  id="VUID-vkCmdCopyAccelerationStructureKHR-buffer-03737"></a>
  VUID-vkCmdCopyAccelerationStructureKHR-buffer-03737  
  The `buffer` used to create `pInfo->src` **must** be bound to device
  memory

- <a href="#VUID-vkCmdCopyAccelerationStructureKHR-buffer-03738"
  id="VUID-vkCmdCopyAccelerationStructureKHR-buffer-03738"></a>
  VUID-vkCmdCopyAccelerationStructureKHR-buffer-03738  
  The `buffer` used to create `pInfo->dst` **must** be bound to device
  memory

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-parameter"
  id="VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-parameter"></a>
  VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyAccelerationStructureKHR-pInfo-parameter"
  id="VUID-vkCmdCopyAccelerationStructureKHR-pInfo-parameter"></a>
  VUID-vkCmdCopyAccelerationStructureKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureInfoKHR.html)
  structure

- <a
  href="#VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-recording"
  id="VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-recording"></a>
  VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyAccelerationStructureKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdCopyAccelerationStructureKHR-renderpass"
  id="VUID-vkCmdCopyAccelerationStructureKHR-renderpass"></a>
  VUID-vkCmdCopyAccelerationStructureKHR-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyAccelerationStructureKHR-videocoding"
  id="VUID-vkCmdCopyAccelerationStructureKHR-videocoding"></a>
  VUID-vkCmdCopyAccelerationStructureKHR-videocoding  
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

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyAccelerationStructureKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
