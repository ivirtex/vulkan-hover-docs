# vkCmdCopyAccelerationStructureToMemoryKHR(3) Manual Page

## Name

vkCmdCopyAccelerationStructureToMemoryKHR - Copy an acceleration
structure to device memory



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy an acceleration structure to device memory call:

``` c
// Provided by VK_KHR_acceleration_structure
void vkCmdCopyAccelerationStructureToMemoryKHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyAccelerationStructureToMemoryInfoKHR* pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pInfo` is an a pointer to a
  [VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html)
  structure defining the copy operation.

## <a href="#_description" class="anchor"></a>Description

Accesses to `pInfo->src` **must** be <a
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
`VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR`. Accesses to the buffer
indicated by `pInfo->dst.deviceAddress` **must** be synchronized with
the `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> or the
`VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a>, and an and an access
type of `VK_ACCESS_TRANSFER_WRITE_BIT`.

This command produces the same results as
[vkCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyAccelerationStructureToMemoryKHR.html),
but writes its result to a device address, and is executed on the device
rather than the host. The output **may** not necessarily be bit-for-bit
identical, but it can be equally used by either
[vkCmdCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToAccelerationStructureKHR.html)
or
[vkCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMemoryToAccelerationStructureKHR.html).

The defined header structure for the serialized data consists of:

- `VK_UUID_SIZE` bytes of data matching
  `VkPhysicalDeviceIDProperties`::`driverUUID`

- `VK_UUID_SIZE` bytes of data identifying the compatibility for
  comparison using
  [vkGetDeviceAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceAccelerationStructureCompatibilityKHR.html)

- A 64-bit integer of the total size matching the value queried using
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`

- A 64-bit integer of the deserialized size to be passed in to
  `VkAccelerationStructureCreateInfoKHR`::`size`

- A 64-bit integer of the count of the number of acceleration structure
  handles following. This value matches the value queried using
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`.
  This will be zero for a bottom-level acceleration structure. For
  top-level acceleration structures this number is
  implementation-dependent; the number of and ordering of the handles
  may not match the instance descriptions which were used to build the
  acceleration structure.

The corresponding handles matching the values returned by
[vkGetAccelerationStructureDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureDeviceAddressKHR.html)
or
[vkGetAccelerationStructureHandleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureHandleNV.html)
are tightly packed in the buffer following the count. The application is
expected to store a mapping between those handles and the original
application-generated bottom-level acceleration structures to provide
when deserializing. The serialized data is written to the buffer (or
read from the buffer) according to the host endianness.

Valid Usage

- <a
  href="#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-accelerationStructure-08926"
  id="VUID-vkCmdCopyAccelerationStructureToMemoryKHR-accelerationStructure-08926"></a>
  VUID-vkCmdCopyAccelerationStructureToMemoryKHR-accelerationStructure-08926  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructure"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructure</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-03739"
  id="VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-03739"></a>
  VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-03739  
  `pInfo->dst.deviceAddress` **must** be a valid device address for a
  buffer bound to device memory

- <a href="#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-03740"
  id="VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-03740"></a>
  VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-03740  
  `pInfo->dst.deviceAddress` **must** be aligned to `256` bytes

- <a href="#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-03741"
  id="VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-03741"></a>
  VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-03741  
  If the buffer pointed to by `pInfo->dst.deviceAddress` is non-sparse
  then it **must** be bound completely and contiguously to a single
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-None-03559"
  id="VUID-vkCmdCopyAccelerationStructureToMemoryKHR-None-03559"></a>
  VUID-vkCmdCopyAccelerationStructureToMemoryKHR-None-03559  
  The `buffer` used to create `pInfo->src` **must** be bound to device
  memory

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-parameter"
  id="VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-parameter"></a>
  VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-parameter"
  id="VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-parameter"></a>
  VUID-vkCmdCopyAccelerationStructureToMemoryKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html)
  structure

- <a
  href="#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-recording"
  id="VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-recording"></a>
  VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyAccelerationStructureToMemoryKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-renderpass"
  id="VUID-vkCmdCopyAccelerationStructureToMemoryKHR-renderpass"></a>
  VUID-vkCmdCopyAccelerationStructureToMemoryKHR-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyAccelerationStructureToMemoryKHR-videocoding"
  id="VUID-vkCmdCopyAccelerationStructureToMemoryKHR-videocoding"></a>
  VUID-vkCmdCopyAccelerationStructureToMemoryKHR-videocoding  
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
[VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyAccelerationStructureToMemoryKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
