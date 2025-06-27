# vkCmdCopyQueryPoolResults(3) Manual Page

## Name

vkCmdCopyQueryPoolResults - Copy the results of queries in a query pool
to a buffer object



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy query statuses and numerical results directly to buffer memory,
call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdCopyQueryPoolResults(
    VkCommandBuffer                             commandBuffer,
    VkQueryPool                                 queryPool,
    uint32_t                                    firstQuery,
    uint32_t                                    queryCount,
    VkBuffer                                    dstBuffer,
    VkDeviceSize                                dstOffset,
    VkDeviceSize                                stride,
    VkQueryResultFlags                          flags);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which this command will be
  recorded.

- `queryPool` is the query pool managing the queries containing the
  desired results.

- `firstQuery` is the initial query index.

- `queryCount` is the number of queries. `firstQuery` and `queryCount`
  together define a range of queries.

- `dstBuffer` is a [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) object that will receive
  the results of the copy command.

- `dstOffset` is an offset into `dstBuffer`.

- `stride` is the stride in bytes between results for individual queries
  within `dstBuffer`. The required size of the backing memory for
  `dstBuffer` is determined as described above for
  [vkGetQueryPoolResults](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetQueryPoolResults.html).

- `flags` is a bitmask of
  [VkQueryResultFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlagBits.html) specifying how and
  when results are returned.

## <a href="#_description" class="anchor"></a>Description

Any results written for a query are written according to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-memorylayout"
target="_blank" rel="noopener">a layout dependent on the query type</a>.

Results for any query in `queryPool` identified by `firstQuery` and
`queryCount` that is available are copied to `dstBuffer`.

If `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT` is set, results for all
queries in `queryPool` identified by `firstQuery` and `queryCount` are
copied to `dstBuffer`, along with an extra availability value written
directly after the results of each query and interpreted as an unsigned
integer. A value of zero indicates that the results are not yet
available, otherwise the query is complete and results are available.

If `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR` is set, results for all queries
in `queryPool` identified by `firstQuery` and `queryCount` are copied to
`dstBuffer`, along with an extra status value written directly after the
results of each query and interpreted as a signed integer. A value of
zero indicates that the results are not yet available. Positive values
indicate that the operations within the query completed successfully,
and the query results are valid. Negative values indicate that the
operations within the query completed unsuccessfully.

[VkQueryResultStatusKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultStatusKHR.html) defines specific
meaning for values returned here, though implementations are free to
return other values.

If the status value written is negative, indicating that the operations
within the query completed unsuccessfully, then all other results
written by this command are undefined unless otherwise specified for any
of the results of the used query type.

Results for any available query written by this command are final and
represent the final result of the query. If
`VK_QUERY_RESULT_PARTIAL_BIT` is set, then for any query that is
unavailable, an intermediate result between zero and the final result
value is written for that query. Otherwise, any result written by this
command is undefined.

If `VK_QUERY_RESULT_64_BIT` is set, results and availability or status
values for all queries are written as an array of 64-bit values. If the
`queryPool` was created with `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`,
results for each query are written as an array of the type indicated by
[VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterKHR.html)::`storage` for
the counter being queried. Otherwise, results and availability or status
values are written as an array of 32-bit values. If an unsigned integer
query’s value overflows the result type, the value **may** either wrap
or saturate. If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
target="_blank" rel="noopener"><code>maintenance7</code></a> feature is
enabled, for an unsigned integer query, the 32-bit result value **must**
be equal to the 32 least significant bits of the equivalent 64-bit
result value. If a signed integer query’s value overflows the result
type, the value is undefined. If a floating-point query’s value is not
representable as the result type, the value is undefined.

This command defines an execution dependency between other query
commands that reference the same query.

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands which reference the queries in `queryPool` indicated by `query`
that occur earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>. If `flags` does not
include `VK_QUERY_RESULT_WAIT_BIT`,
[vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQueryIndexedEXT.html),
[vkCmdWriteTimestamp2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp2.html),
[vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQuery.html), and
[vkCmdWriteTimestamp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp.html) are excluded from this
scope.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands which reference the queries in `queryPool` indicated by `query`
that occur later in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>.

The operation of this command happens after the first scope and happens
before the second scope.

`vkCmdCopyQueryPoolResults` is considered to be a transfer operation,
and its writes to buffer memory **must** be synchronized using
`VK_PIPELINE_STAGE_TRANSFER_BIT` and `VK_ACCESS_TRANSFER_WRITE_BIT`
before using the results.

Valid Usage

- <a href="#VUID-vkCmdCopyQueryPoolResults-firstQuery-09436"
  id="VUID-vkCmdCopyQueryPoolResults-firstQuery-09436"></a>
  VUID-vkCmdCopyQueryPoolResults-firstQuery-09436  
  `firstQuery` **must** be less than the number of queries in
  `queryPool`

- <a href="#VUID-vkCmdCopyQueryPoolResults-firstQuery-09437"
  id="VUID-vkCmdCopyQueryPoolResults-firstQuery-09437"></a>
  VUID-vkCmdCopyQueryPoolResults-firstQuery-09437  
  The sum of `firstQuery` and `queryCount` **must** be less than or
  equal to the number of queries in `queryPool`

<!-- -->

- <a href="#VUID-vkCmdCopyQueryPoolResults-queryCount-09438"
  id="VUID-vkCmdCopyQueryPoolResults-queryCount-09438"></a>
  VUID-vkCmdCopyQueryPoolResults-queryCount-09438  
  If `queryCount` is greater than 1, `stride` **must** not be zero

- <a href="#VUID-vkCmdCopyQueryPoolResults-queryType-09439"
  id="VUID-vkCmdCopyQueryPoolResults-queryType-09439"></a>
  VUID-vkCmdCopyQueryPoolResults-queryType-09439  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_TIMESTAMP`, `flags` **must** not contain
  `VK_QUERY_RESULT_PARTIAL_BIT`

- <a href="#VUID-vkCmdCopyQueryPoolResults-queryType-09440"
  id="VUID-vkCmdCopyQueryPoolResults-queryType-09440"></a>
  VUID-vkCmdCopyQueryPoolResults-queryType-09440  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, `flags` **must** not contain
  `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT`,
  `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`, `VK_QUERY_RESULT_PARTIAL_BIT`,
  or `VK_QUERY_RESULT_64_BIT`

- <a href="#VUID-vkCmdCopyQueryPoolResults-queryType-09441"
  id="VUID-vkCmdCopyQueryPoolResults-queryType-09441"></a>
  VUID-vkCmdCopyQueryPoolResults-queryType-09441  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the `queryPool` **must** have
  been recorded once for each pass as retrieved via a call to
  [vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html)

- <a href="#VUID-vkCmdCopyQueryPoolResults-queryType-09442"
  id="VUID-vkCmdCopyQueryPoolResults-queryType-09442"></a>
  VUID-vkCmdCopyQueryPoolResults-queryType-09442  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR`, then `flags` **must** include
  `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`

- <a href="#VUID-vkCmdCopyQueryPoolResults-flags-09443"
  id="VUID-vkCmdCopyQueryPoolResults-flags-09443"></a>
  VUID-vkCmdCopyQueryPoolResults-flags-09443  
  If `flags` includes `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`, then it
  **must** not include `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT`

<!-- -->

- <a href="#VUID-vkCmdCopyQueryPoolResults-None-09402"
  id="VUID-vkCmdCopyQueryPoolResults-None-09402"></a>
  VUID-vkCmdCopyQueryPoolResults-None-09402  
  All queries used by the command **must** not be uninitialized when the
  command is executed

- <a href="#VUID-vkCmdCopyQueryPoolResults-dstOffset-00819"
  id="VUID-vkCmdCopyQueryPoolResults-dstOffset-00819"></a>
  VUID-vkCmdCopyQueryPoolResults-dstOffset-00819  
  `dstOffset` **must** be less than the size of `dstBuffer`

- <a href="#VUID-vkCmdCopyQueryPoolResults-flags-00822"
  id="VUID-vkCmdCopyQueryPoolResults-flags-00822"></a>
  VUID-vkCmdCopyQueryPoolResults-flags-00822  
  If `VK_QUERY_RESULT_64_BIT` is not set in `flags` then `dstOffset` and
  `stride` **must** be multiples of `4`

- <a href="#VUID-vkCmdCopyQueryPoolResults-flags-00823"
  id="VUID-vkCmdCopyQueryPoolResults-flags-00823"></a>
  VUID-vkCmdCopyQueryPoolResults-flags-00823  
  If `VK_QUERY_RESULT_64_BIT` is set in `flags` then `dstOffset` and
  `stride` **must** be multiples of `8`

- <a href="#VUID-vkCmdCopyQueryPoolResults-dstBuffer-00824"
  id="VUID-vkCmdCopyQueryPoolResults-dstBuffer-00824"></a>
  VUID-vkCmdCopyQueryPoolResults-dstBuffer-00824  
  `dstBuffer` **must** have enough storage, from `dstOffset`, to contain
  the result of each query, as described <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-memorylayout"
  target="_blank" rel="noopener">here</a>

- <a href="#VUID-vkCmdCopyQueryPoolResults-dstBuffer-00825"
  id="VUID-vkCmdCopyQueryPoolResults-dstBuffer-00825"></a>
  VUID-vkCmdCopyQueryPoolResults-dstBuffer-00825  
  `dstBuffer` **must** have been created with
  `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-vkCmdCopyQueryPoolResults-dstBuffer-00826"
  id="VUID-vkCmdCopyQueryPoolResults-dstBuffer-00826"></a>
  VUID-vkCmdCopyQueryPoolResults-dstBuffer-00826  
  If `dstBuffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdCopyQueryPoolResults-queryType-03232"
  id="VUID-vkCmdCopyQueryPoolResults-queryType-03232"></a>
  VUID-vkCmdCopyQueryPoolResults-queryType-03232  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`,
  [VkPhysicalDevicePerformanceQueryPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePerformanceQueryPropertiesKHR.html)::`allowCommandBufferQueryCopies`
  **must** be `VK_TRUE`

- <a href="#VUID-vkCmdCopyQueryPoolResults-queryType-02734"
  id="VUID-vkCmdCopyQueryPoolResults-queryType-02734"></a>
  VUID-vkCmdCopyQueryPoolResults-queryType-02734  
  [vkCmdCopyQueryPoolResults](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyQueryPoolResults.html) **must**
  not be called if the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_INTEL`

- <a href="#VUID-vkCmdCopyQueryPoolResults-None-07429"
  id="VUID-vkCmdCopyQueryPoolResults-None-07429"></a>
  VUID-vkCmdCopyQueryPoolResults-None-07429  
  All queries used by the command **must** not be active

- <a href="#VUID-vkCmdCopyQueryPoolResults-None-08752"
  id="VUID-vkCmdCopyQueryPoolResults-None-08752"></a>
  VUID-vkCmdCopyQueryPoolResults-None-08752  
  All queries used by the command **must** have been made *available* by
  prior executed commands

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCopyQueryPoolResults-commandBuffer-parameter"
  id="VUID-vkCmdCopyQueryPoolResults-commandBuffer-parameter"></a>
  VUID-vkCmdCopyQueryPoolResults-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyQueryPoolResults-queryPool-parameter"
  id="VUID-vkCmdCopyQueryPoolResults-queryPool-parameter"></a>
  VUID-vkCmdCopyQueryPoolResults-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

- <a href="#VUID-vkCmdCopyQueryPoolResults-dstBuffer-parameter"
  id="VUID-vkCmdCopyQueryPoolResults-dstBuffer-parameter"></a>
  VUID-vkCmdCopyQueryPoolResults-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-vkCmdCopyQueryPoolResults-flags-parameter"
  id="VUID-vkCmdCopyQueryPoolResults-flags-parameter"></a>
  VUID-vkCmdCopyQueryPoolResults-flags-parameter  
  `flags` **must** be a valid combination of
  [VkQueryResultFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlagBits.html) values

- <a href="#VUID-vkCmdCopyQueryPoolResults-commandBuffer-recording"
  id="VUID-vkCmdCopyQueryPoolResults-commandBuffer-recording"></a>
  VUID-vkCmdCopyQueryPoolResults-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyQueryPoolResults-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyQueryPoolResults-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyQueryPoolResults-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdCopyQueryPoolResults-renderpass"
  id="VUID-vkCmdCopyQueryPoolResults-renderpass"></a>
  VUID-vkCmdCopyQueryPoolResults-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyQueryPoolResults-videocoding"
  id="VUID-vkCmdCopyQueryPoolResults-videocoding"></a>
  VUID-vkCmdCopyQueryPoolResults-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdCopyQueryPoolResults-commonparent"
  id="VUID-vkCmdCopyQueryPoolResults-commonparent"></a>
  VUID-vkCmdCopyQueryPoolResults-commonparent  
  Each of `commandBuffer`, `dstBuffer`, and `queryPool` **must** have
  been created, allocated, or retrieved from the same
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
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html),
[VkQueryResultFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyQueryPoolResults"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
