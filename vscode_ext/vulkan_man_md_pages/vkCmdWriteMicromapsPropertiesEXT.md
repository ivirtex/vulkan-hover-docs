# vkCmdWriteMicromapsPropertiesEXT(3) Manual Page

## Name

vkCmdWriteMicromapsPropertiesEXT - Write micromap result parameters to
query results.



## <a href="#_c_specification" class="anchor"></a>C Specification

To query micromap size parameters call:

``` c
// Provided by VK_EXT_opacity_micromap
void vkCmdWriteMicromapsPropertiesEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    micromapCount,
    const VkMicromapEXT*                        pMicromaps,
    VkQueryType                                 queryType,
    VkQueryPool                                 queryPool,
    uint32_t                                    firstQuery);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `micromapCount` is the count of micromaps for which to query the
  property.

- `pMicromaps` is a pointer to an array of existing previously built
  micromaps.

- `queryType` is a [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html) value specifying the
  type of queries managed by the pool.

- `queryPool` is the query pool that will manage the results of the
  query.

- `firstQuery` is the first query index within the query pool that will
  contain the `micromapCount` number of results.

## <a href="#_description" class="anchor"></a>Description

Accesses to any of the micromaps listed in `pMicromaps` **must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">synchronized</a> with the
`VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
target="_blank" rel="noopener">pipeline stage</a> and an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types"
target="_blank" rel="noopener">access type</a> of
`VK_ACCESS_2_MICROMAP_READ_BIT_EXT`.

- If `queryType` is `VK_QUERY_TYPE_MICROMAP_SERIALIZATION_SIZE_EXT`,
  then the value written out is the number of bytes required by a
  serialized micromap.

- If `queryType` is `VK_QUERY_TYPE_MICROMAP_COMPACTED_SIZE_EXT`, then
  the value written out is the number of bytes required by a compacted
  micromap.

Valid Usage

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-queryPool-07525"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-queryPool-07525"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-queryPool-07525  
  `queryPool` **must** have been created with a `queryType` matching
  `queryType`

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-queryPool-07526"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-queryPool-07526"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-queryPool-07526  
  The queries identified by `queryPool` and `firstQuery` **must** be
  *unavailable*

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-buffer-07527"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-buffer-07527"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-buffer-07527  
  The `buffer` used to create each micromap in `pMicrmaps` **must** be
  bound to device memory

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-query-07528"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-query-07528"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-query-07528  
  The sum of `query` plus `micromapCount` **must** be less than or equal
  to the number of queries in `queryPool`

<!-- -->

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-pMicromaps-07501"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-pMicromaps-07501"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-pMicromaps-07501  
  All micromaps in `pMicromaps` **must** have been constructed prior to
  the execution of this command

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-pMicromaps-07502"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-pMicromaps-07502"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-pMicromaps-07502  
  All micromaps in `pMicromaps` **must** have been constructed with
  `VK_BUILD_MICROMAP_ALLOW_COMPACTION_BIT_EXT` if `queryType` is
  `VK_QUERY_TYPE_MICROMAP_COMPACTED_SIZE_EXT`

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-queryType-07503"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-queryType-07503"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-queryType-07503  
  `queryType` **must** be `VK_QUERY_TYPE_MICROMAP_COMPACTED_SIZE_EXT` or
  `VK_QUERY_TYPE_MICROMAP_SERIALIZATION_SIZE_EXT`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-commandBuffer-parameter"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-commandBuffer-parameter"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-pMicromaps-parameter"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-pMicromaps-parameter"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-pMicromaps-parameter  
  `pMicromaps` **must** be a valid pointer to an array of
  `micromapCount` valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html) handles

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-queryType-parameter"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-queryType-parameter"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-queryType-parameter  
  `queryType` **must** be a valid [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html) value

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-queryPool-parameter"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-queryPool-parameter"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-commandBuffer-recording"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-commandBuffer-recording"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-renderpass"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-renderpass"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-videocoding"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-videocoding"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdWriteMicromapsPropertiesEXT-micromapCount-arraylength"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-micromapCount-arraylength"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-micromapCount-arraylength  
  `micromapCount` **must** be greater than `0`

- <a href="#VUID-vkCmdWriteMicromapsPropertiesEXT-commonparent"
  id="VUID-vkCmdWriteMicromapsPropertiesEXT-commonparent"></a>
  VUID-vkCmdWriteMicromapsPropertiesEXT-commonparent  
  Each of `commandBuffer`, `queryPool`, and the elements of `pMicromaps`
  **must** have been created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
[VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html),
[VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdWriteMicromapsPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
