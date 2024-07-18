# vkCmdWriteAccelerationStructuresPropertiesNV(3) Manual Page

## Name

vkCmdWriteAccelerationStructuresPropertiesNV - Write acceleration
structure result parameters to query results.



## <a href="#_c_specification" class="anchor"></a>C Specification

To query acceleration structure size parameters call:

``` c
// Provided by VK_NV_ray_tracing
void vkCmdWriteAccelerationStructuresPropertiesNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    accelerationStructureCount,
    const VkAccelerationStructureNV*            pAccelerationStructures,
    VkQueryType                                 queryType,
    VkQueryPool                                 queryPool,
    uint32_t                                    firstQuery);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `accelerationStructureCount` is the count of acceleration structures
  for which to query the property.

- `pAccelerationStructures` is a pointer to an array of existing
  previously built acceleration structures.

- `queryType` is a [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html) value specifying the
  type of queries managed by the pool.

- `queryPool` is the query pool that will manage the results of the
  query.

- `firstQuery` is the first query index within the query pool that will
  contain the `accelerationStructureCount` number of results.

## <a href="#_description" class="anchor"></a>Description

Accesses to any of the acceleration structures listed in
`pAccelerationStructures` **must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
`VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR`.

Valid Usage

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-03755"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-03755"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-03755  
  `queryPool` **must** have been created with a `queryType` matching
  `queryType`

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-03756"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-03756"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-03756  
  The queries identified by `queryPool` and `firstQuery` **must** be
  *unavailable*

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-accelerationStructure-03757"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-accelerationStructure-03757"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-accelerationStructure-03757  
  `accelerationStructure` **must** be bound completely and contiguously
  to a single `VkDeviceMemory` object via
  [vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindAccelerationStructureMemoryNV.html)

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-04958"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-04958"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-04958  
  All acceleration structures in `pAccelerationStructures` **must** have
  been built prior to the execution of this command

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-06215"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-06215"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-06215  
  All acceleration structures in `pAccelerationStructures` **must** have
  been built with
  `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR` if
  `queryType` is
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV`

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryType-06216"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryType-06216"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryType-06216  
  `queryType` **must** be
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV`

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-parameter"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-parameter"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-parameter"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-parameter"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-pAccelerationStructures-parameter  
  `pAccelerationStructures` **must** be a valid pointer to an array of
  `accelerationStructureCount` valid
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) handles

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryType-parameter"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryType-parameter"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryType-parameter  
  `queryType` **must** be a valid [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html) value

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-parameter"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-parameter"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-recording"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-recording"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-cmdpool"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-renderpass"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-renderpass"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-videocoding"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-videocoding"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-accelerationStructureCount-arraylength"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-accelerationStructureCount-arraylength"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-accelerationStructureCount-arraylength  
  `accelerationStructureCount` **must** be greater than `0`

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commonparent"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commonparent"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesNV-commonparent  
  Each of `commandBuffer`, `queryPool`, and the elements of
  `pAccelerationStructures` **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html), [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdWriteAccelerationStructuresPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
