# vkCmdResetQueryPool(3) Manual Page

## Name

vkCmdResetQueryPool - Reset queries in a query pool



## <a href="#_c_specification" class="anchor"></a>C Specification

To reset a range of queries in a query pool on a queue, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdResetQueryPool(
    VkCommandBuffer                             commandBuffer,
    VkQueryPool                                 queryPool,
    uint32_t                                    firstQuery,
    uint32_t                                    queryCount);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which this command will be
  recorded.

- `queryPool` is the handle of the query pool managing the queries being
  reset.

- `firstQuery` is the initial query index to reset.

- `queryCount` is the number of queries to reset.

## <a href="#_description" class="anchor"></a>Description

When executed on a queue, this command sets the status of query indices
\[`firstQuery`, `firstQuery` + `queryCount` - 1\] to unavailable.

This command defines an execution dependency between other query
commands that reference the same query.

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands which reference the queries in `queryPool` indicated by
`firstQuery` and `queryCount` that occur earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands which reference the queries in `queryPool` indicated by
`firstQuery` and `queryCount` that occur later in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>.

The operation of this command happens after the first scope and happens
before the second scope.

If the `queryType` used to create `queryPool` was
`VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, this command sets the status of
query indices \[`firstQuery`, `firstQuery` + `queryCount` - 1\] to
unavailable for each pass of `queryPool`, as indicated by a call to
[vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html).

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Because <code>vkCmdResetQueryPool</code> resets all the passes of the
indicated queries, applications must not record a
<code>vkCmdResetQueryPool</code> command for a <code>queryPool</code>
created with <code>VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR</code> in a
command buffer that needs to be submitted multiple times as indicated by
a call to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html">vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR</a>.
Otherwise applications will never be able to complete the recorded
queries.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkCmdResetQueryPool-firstQuery-09436"
  id="VUID-vkCmdResetQueryPool-firstQuery-09436"></a>
  VUID-vkCmdResetQueryPool-firstQuery-09436  
  `firstQuery` **must** be less than the number of queries in
  `queryPool`

- <a href="#VUID-vkCmdResetQueryPool-firstQuery-09437"
  id="VUID-vkCmdResetQueryPool-firstQuery-09437"></a>
  VUID-vkCmdResetQueryPool-firstQuery-09437  
  The sum of `firstQuery` and `queryCount` **must** be less than or
  equal to the number of queries in `queryPool`

<!-- -->

- <a href="#VUID-vkCmdResetQueryPool-None-02841"
  id="VUID-vkCmdResetQueryPool-None-02841"></a>
  VUID-vkCmdResetQueryPool-None-02841  
  All queries used by the command **must** not be active

- <a href="#VUID-vkCmdResetQueryPool-firstQuery-02862"
  id="VUID-vkCmdResetQueryPool-firstQuery-02862"></a>
  VUID-vkCmdResetQueryPool-firstQuery-02862  
  If `queryPool` was created with `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`,
  this command **must** not be recorded in a command buffer that, either
  directly or through secondary command buffers, also contains begin
  commands for a query from the set of queries \[`firstQuery`,
  `firstQuery` + `queryCount` - 1\]

Valid Usage (Implicit)

- <a href="#VUID-vkCmdResetQueryPool-commandBuffer-parameter"
  id="VUID-vkCmdResetQueryPool-commandBuffer-parameter"></a>
  VUID-vkCmdResetQueryPool-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdResetQueryPool-queryPool-parameter"
  id="VUID-vkCmdResetQueryPool-queryPool-parameter"></a>
  VUID-vkCmdResetQueryPool-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

- <a href="#VUID-vkCmdResetQueryPool-commandBuffer-recording"
  id="VUID-vkCmdResetQueryPool-commandBuffer-recording"></a>
  VUID-vkCmdResetQueryPool-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdResetQueryPool-commandBuffer-cmdpool"
  id="VUID-vkCmdResetQueryPool-commandBuffer-cmdpool"></a>
  VUID-vkCmdResetQueryPool-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, decode, encode, or optical flow operations

- <a href="#VUID-vkCmdResetQueryPool-renderpass"
  id="VUID-vkCmdResetQueryPool-renderpass"></a>
  VUID-vkCmdResetQueryPool-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdResetQueryPool-videocoding"
  id="VUID-vkCmdResetQueryPool-videocoding"></a>
  VUID-vkCmdResetQueryPool-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdResetQueryPool-commonparent"
  id="VUID-vkCmdResetQueryPool-commonparent"></a>
  VUID-vkCmdResetQueryPool-commonparent  
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
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute<br />
Decode<br />
Encode<br />
Opticalflow</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdResetQueryPool"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
