# vkCmdEndQueryIndexedEXT(3) Manual Page

## Name

vkCmdEndQueryIndexedEXT - Ends a query



## <a href="#_c_specification" class="anchor"></a>C Specification

To end an indexed query after the set of desired drawing or dispatching
commands is recorded, call:

``` c
// Provided by VK_EXT_transform_feedback
void vkCmdEndQueryIndexedEXT(
    VkCommandBuffer                             commandBuffer,
    VkQueryPool                                 queryPool,
    uint32_t                                    query,
    uint32_t                                    index);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which this command will be
  recorded.

- `queryPool` is the query pool that is managing the results of the
  query.

- `query` is the query index within the query pool where the result is
  stored.

- `index` is the query type specific index.

## <a href="#_description" class="anchor"></a>Description

The command completes the query in `queryPool` identified by `query` and
`index`, and marks it as available.

The `vkCmdEndQueryIndexedEXT` command operates the same as the
[vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQuery.html) command, except that it also accepts
a query type specific `index` parameter.

This command defines an execution dependency between other query
commands that reference the same query index.

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands which reference the queries in `queryPool` indicated by `query`
that occur earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes only
the operation of this command.

Valid Usage

- <a href="#VUID-vkCmdEndQueryIndexedEXT-None-02342"
  id="VUID-vkCmdEndQueryIndexedEXT-None-02342"></a>
  VUID-vkCmdEndQueryIndexedEXT-None-02342  
  All queries used by the command **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a>

- <a href="#VUID-vkCmdEndQueryIndexedEXT-query-02343"
  id="VUID-vkCmdEndQueryIndexedEXT-query-02343"></a>
  VUID-vkCmdEndQueryIndexedEXT-query-02343  
  `query` **must** be less than the number of queries in `queryPool`

- <a href="#VUID-vkCmdEndQueryIndexedEXT-commandBuffer-02344"
  id="VUID-vkCmdEndQueryIndexedEXT-commandBuffer-02344"></a>
  VUID-vkCmdEndQueryIndexedEXT-commandBuffer-02344  
  `commandBuffer` **must** not be a protected command buffer

- <a href="#VUID-vkCmdEndQueryIndexedEXT-query-02345"
  id="VUID-vkCmdEndQueryIndexedEXT-query-02345"></a>
  VUID-vkCmdEndQueryIndexedEXT-query-02345  
  If `vkCmdEndQueryIndexedEXT` is called within a render pass instance,
  the sum of `query` and the number of bits set in the current subpass’s
  view mask **must** be less than or equal to the number of queries in
  `queryPool`

- <a href="#VUID-vkCmdEndQueryIndexedEXT-queryType-06694"
  id="VUID-vkCmdEndQueryIndexedEXT-queryType-06694"></a>
  VUID-vkCmdEndQueryIndexedEXT-queryType-06694  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` or
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT`, the `index` parameter
  **must** be less than
  `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackStreams`

- <a href="#VUID-vkCmdEndQueryIndexedEXT-queryType-06695"
  id="VUID-vkCmdEndQueryIndexedEXT-queryType-06695"></a>
  VUID-vkCmdEndQueryIndexedEXT-queryType-06695  
  If the `queryType` used to create `queryPool` was not
  `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` and not
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT`, the `index` **must** be zero

- <a href="#VUID-vkCmdEndQueryIndexedEXT-queryType-06696"
  id="VUID-vkCmdEndQueryIndexedEXT-queryType-06696"></a>
  VUID-vkCmdEndQueryIndexedEXT-queryType-06696  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` or
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT`, `index` **must** equal the
  `index` used to begin the query

<!-- -->

- <a href="#VUID-vkCmdEndQueryIndexedEXT-None-07007"
  id="VUID-vkCmdEndQueryIndexedEXT-None-07007"></a>
  VUID-vkCmdEndQueryIndexedEXT-None-07007  
  If called within a subpass of a render pass instance, the
  corresponding `vkCmdBeginQuery`\* command **must** have been called
  previously within the same subpass

Valid Usage (Implicit)

- <a href="#VUID-vkCmdEndQueryIndexedEXT-commandBuffer-parameter"
  id="VUID-vkCmdEndQueryIndexedEXT-commandBuffer-parameter"></a>
  VUID-vkCmdEndQueryIndexedEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdEndQueryIndexedEXT-queryPool-parameter"
  id="VUID-vkCmdEndQueryIndexedEXT-queryPool-parameter"></a>
  VUID-vkCmdEndQueryIndexedEXT-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

- <a href="#VUID-vkCmdEndQueryIndexedEXT-commandBuffer-recording"
  id="VUID-vkCmdEndQueryIndexedEXT-commandBuffer-recording"></a>
  VUID-vkCmdEndQueryIndexedEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdEndQueryIndexedEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdEndQueryIndexedEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdEndQueryIndexedEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, decode, or encode operations

- <a href="#VUID-vkCmdEndQueryIndexedEXT-videocoding"
  id="VUID-vkCmdEndQueryIndexedEXT-videocoding"></a>
  VUID-vkCmdEndQueryIndexedEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdEndQueryIndexedEXT-commonparent"
  id="VUID-vkCmdEndQueryIndexedEXT-commonparent"></a>
  VUID-vkCmdEndQueryIndexedEXT-commonparent  
  Both of `commandBuffer`, and `queryPool` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute<br />
Decode<br />
Encode</p></td>
<td class="tableblock halign-left valign-top"><p>Action<br />
State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_transform_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_transform_feedback.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html),
[vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQuery.html),
[vkCmdBeginQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQueryIndexedEXT.html),
[vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQuery.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdEndQueryIndexedEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
