# vkCmdWriteAccelerationStructuresPropertiesKHR(3) Manual Page

## Name

vkCmdWriteAccelerationStructuresPropertiesKHR - Write acceleration
structure result parameters to query results.



## <a href="#_c_specification" class="anchor"></a>C Specification

To query acceleration structure size parameters call:

``` c
// Provided by VK_KHR_acceleration_structure
void vkCmdWriteAccelerationStructuresPropertiesKHR(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    accelerationStructureCount,
    const VkAccelerationStructureKHR*           pAccelerationStructures,
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
`VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> or the
`VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a>, and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
`VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR`.

- If `queryType` is
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR`, then the
  value written out is the number of bytes required by a compacted
  acceleration structure.

- If `queryType` is
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`, then
  the value written out is the number of bytes required by a serialized
  acceleration structure.

Valid Usage

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructure-08924"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructure-08924"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructure-08924  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructure"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructure</code></a>
  feature **must** be enabled

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-02493"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-02493"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-02493  
  `queryPool` **must** have been created with a `queryType` matching
  `queryType`

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-02494"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-02494"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-02494  
  The queries identified by `queryPool` and `firstQuery` **must** be
  *unavailable*

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-buffer-03736"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-buffer-03736"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-buffer-03736  
  The `buffer` used to create each acceleration structure in
  `pAccelerationStructures` **must** be bound to device memory

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-query-04880"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-query-04880"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-query-04880  
  The sum of `firstQuery` plus `accelerationStructureCount` **must** be
  less than or equal to the number of queries in `queryPool`

<!-- -->

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-04964"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-04964"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-04964  
  All acceleration structures in `pAccelerationStructures` **must** have
  been built prior to the execution of this command

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructures-03431"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructures-03431"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructures-03431  
  All acceleration structures in `pAccelerationStructures` **must** have
  been built with
  `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR` if
  `queryType` is
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR`

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryType-06742"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryType-06742"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryType-06742  
  `queryType` **must** be
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR`,
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`,
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR`, or
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-parameter"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-parameter"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-parameter"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-parameter"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-parameter  
  `pAccelerationStructures` **must** be a valid pointer to an array of
  `accelerationStructureCount` valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handles

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryType-parameter"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryType-parameter"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryType-parameter  
  `queryType` **must** be a valid [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html) value

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-parameter"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-parameter"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-recording"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-recording"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-renderpass"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-renderpass"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-renderpass  
  This command **must** only be called outside of a render pass instance

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-videocoding"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-videocoding"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructureCount-arraylength"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructureCount-arraylength"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-accelerationStructureCount-arraylength  
  `accelerationStructureCount` **must** be greater than `0`

- <a
  href="#VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commonparent"
  id="VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commonparent"></a>
  VUID-vkCmdWriteAccelerationStructuresPropertiesKHR-commonparent  
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

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html), [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdWriteAccelerationStructuresPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
