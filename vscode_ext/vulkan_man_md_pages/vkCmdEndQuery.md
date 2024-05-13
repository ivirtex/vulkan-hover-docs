# vkCmdEndQuery(3) Manual Page

## Name

vkCmdEndQuery - Ends a query



## <a href="#_c_specification" class="anchor"></a>C Specification

To end a query after the set of desired drawing or dispatching commands
is executed, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdEndQuery(
    VkCommandBuffer                             commandBuffer,
    VkQueryPool                                 queryPool,
    uint32_t                                    query);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which this command will be
  recorded.

- `queryPool` is the query pool that is managing the results of the
  query.

- `query` is the query index within the query pool where the result is
  stored.

## <a href="#_description" class="anchor"></a>Description

The command completes the query in `queryPool` identified by `query`,
and marks it as available.

This command defines an execution dependency between other query
commands that reference the same query.

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

Calling `vkCmdEndQuery` is equivalent to calling
[vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQueryIndexedEXT.html) with the `index`
parameter set to zero.

Valid Usage

- <a href="#VUID-vkCmdEndQuery-None-01923"
  id="VUID-vkCmdEndQuery-None-01923"></a>
  VUID-vkCmdEndQuery-None-01923  
  All queries used by the command **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a>

- <a href="#VUID-vkCmdEndQuery-query-00810"
  id="VUID-vkCmdEndQuery-query-00810"></a>
  VUID-vkCmdEndQuery-query-00810  
  `query` **must** be less than the number of queries in `queryPool`

- <a href="#VUID-vkCmdEndQuery-commandBuffer-01886"
  id="VUID-vkCmdEndQuery-commandBuffer-01886"></a>
  VUID-vkCmdEndQuery-commandBuffer-01886  
  `commandBuffer` **must** not be a protected command buffer

- <a href="#VUID-vkCmdEndQuery-query-00812"
  id="VUID-vkCmdEndQuery-query-00812"></a>
  VUID-vkCmdEndQuery-query-00812  
  If `vkCmdEndQuery` is called within a render pass instance, the sum of
  `query` and the number of bits set in the current subpass’s view mask
  **must** be less than or equal to the number of queries in `queryPool`

- <a href="#VUID-vkCmdEndQuery-queryPool-03227"
  id="VUID-vkCmdEndQuery-queryPool-03227"></a>
  VUID-vkCmdEndQuery-queryPool-03227  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and one or more of the counters
  used to create `queryPool` was
  `VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_BUFFER_KHR`, the
  [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQuery.html) **must** be the last recorded
  command in `commandBuffer`

- <a href="#VUID-vkCmdEndQuery-queryPool-03228"
  id="VUID-vkCmdEndQuery-queryPool-03228"></a>
  VUID-vkCmdEndQuery-queryPool-03228  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and one or more of the counters
  used to create `queryPool` was
  `VK_PERFORMANCE_COUNTER_SCOPE_RENDER_PASS_KHR`, the
  [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQuery.html) **must** not be recorded within a
  render pass instance

<!-- -->

- <a href="#VUID-vkCmdEndQuery-None-07007"
  id="VUID-vkCmdEndQuery-None-07007"></a>
  VUID-vkCmdEndQuery-None-07007  
  If called within a subpass of a render pass instance, the
  corresponding `vkCmdBeginQuery`\* command **must** have been called
  previously within the same subpass

Valid Usage (Implicit)

- <a href="#VUID-vkCmdEndQuery-commandBuffer-parameter"
  id="VUID-vkCmdEndQuery-commandBuffer-parameter"></a>
  VUID-vkCmdEndQuery-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdEndQuery-queryPool-parameter"
  id="VUID-vkCmdEndQuery-queryPool-parameter"></a>
  VUID-vkCmdEndQuery-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

- <a href="#VUID-vkCmdEndQuery-commandBuffer-recording"
  id="VUID-vkCmdEndQuery-commandBuffer-recording"></a>
  VUID-vkCmdEndQuery-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdEndQuery-commandBuffer-cmdpool"
  id="VUID-vkCmdEndQuery-commandBuffer-cmdpool"></a>
  VUID-vkCmdEndQuery-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, decode, or encode operations

- <a href="#VUID-vkCmdEndQuery-commonparent"
  id="VUID-vkCmdEndQuery-commonparent"></a>
  VUID-vkCmdEndQuery-commonparent  
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
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Both</p></td>
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

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html),
[vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQuery.html),
[vkCmdBeginQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQueryIndexedEXT.html),
[vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQueryIndexedEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdEndQuery"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
