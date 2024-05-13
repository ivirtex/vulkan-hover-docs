# vkCmdCopyMemoryToAccelerationStructureKHR(3) Manual Page

## Name

vkCmdCopyMemoryToAccelerationStructureKHR - Copy device memory to an
acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy device memory to an acceleration structure call:

``` c
// Provided by VK_KHR_acceleration_structure
void vkCmdCopyMemoryToAccelerationStructureKHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyMemoryToAccelerationStructureInfoKHR* pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pInfo` is a pointer to a
  [VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html)
  structure defining the copy operation.

## <a href="#_description" class="anchor"></a>Description

Accesses to `pInfo->dst` **must** be <a
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
`VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`. Accesses to the buffer
indicated by `pInfo->src.deviceAddress` **must** be synchronized with
the `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> or the
`VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a>, and an access type of
`VK_ACCESS_TRANSFER_READ_BIT`.

This command can accept acceleration structures produced by either
[vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html)
or
[vkCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyAccelerationStructureToMemoryKHR.html).

The structure provided as input to deserialize is as described in
[vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html),
with any acceleration structure handles filled in with the newly-queried
handles to bottom level acceleration structures created before
deserialization. These do not need to be built at deserialize time, but
**must** be created.

Valid Usage

- <a
  href="#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-accelerationStructure-08927"
  id="VUID-vkCmdCopyMemoryToAccelerationStructureKHR-accelerationStructure-08927"></a>
  VUID-vkCmdCopyMemoryToAccelerationStructureKHR-accelerationStructure-08927  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructure"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructure</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03742"
  id="VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03742"></a>
  VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03742  
  `pInfo->src.deviceAddress` **must** be a valid device address for a
  buffer bound to device memory

- <a href="#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03743"
  id="VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03743"></a>
  VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03743  
  `pInfo->src.deviceAddress` **must** be aligned to `256` bytes

- <a href="#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03744"
  id="VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03744"></a>
  VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-03744  
  If the buffer pointed to by `pInfo->src.deviceAddress` is non-sparse
  then it **must** be bound completely and contiguously to a single
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-buffer-03745"
  id="VUID-vkCmdCopyMemoryToAccelerationStructureKHR-buffer-03745"></a>
  VUID-vkCmdCopyMemoryToAccelerationStructureKHR-buffer-03745  
  The `buffer` used to create `pInfo->dst` **must** be bound to device
  memory

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-parameter"
  id="VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-parameter"></a>
  VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-parameter"
  id="VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-parameter"></a>
  VUID-vkCmdCopyMemoryToAccelerationStructureKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html)
  structure

- <a
  href="#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-recording"
  id="VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-recording"></a>
  VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyMemoryToAccelerationStructureKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-renderpass"
  id="VUID-vkCmdCopyMemoryToAccelerationStructureKHR-renderpass"></a>
  VUID-vkCmdCopyMemoryToAccelerationStructureKHR-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyMemoryToAccelerationStructureKHR-videocoding"
  id="VUID-vkCmdCopyMemoryToAccelerationStructureKHR-videocoding"></a>
  VUID-vkCmdCopyMemoryToAccelerationStructureKHR-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyMemoryToAccelerationStructureKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
