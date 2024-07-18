# vkGetQueryPoolResults(3) Manual Page

## Name

vkGetQueryPoolResults - Copy results of queries in a query pool to a
host memory region



## <a href="#_c_specification" class="anchor"></a>C Specification

To retrieve status and results for a set of queries, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkGetQueryPoolResults(
    VkDevice                                    device,
    VkQueryPool                                 queryPool,
    uint32_t                                    firstQuery,
    uint32_t                                    queryCount,
    size_t                                      dataSize,
    void*                                       pData,
    VkDeviceSize                                stride,
    VkQueryResultFlags                          flags);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the query pool.

- `queryPool` is the query pool managing the queries containing the
  desired results.

- `firstQuery` is the initial query index.

- `queryCount` is the number of queries to read.

- `dataSize` is the size in bytes of the buffer pointed to by `pData`.

- `pData` is a pointer to an application-allocated buffer where the
  results will be written

- `stride` is the stride in bytes between results for individual queries
  within `pData`.

- `flags` is a bitmask of
  [VkQueryResultFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlagBits.html) specifying how and
  when results are returned.

## <a href="#_description" class="anchor"></a>Description

Any results written for a query are written according to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-memorylayout"
target="_blank" rel="noopener">a layout dependent on the query type</a>.

If no bits are set in `flags`, and all requested queries are in the
available state, results are written as an array of 32-bit unsigned
integer values. Behavior when not all queries are available is described
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-wait-bit-not-set"
target="_blank" rel="noopener">below</a>.

If `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT` is set, results for all
queries in `queryPool` identified by `firstQuery` and `queryCount` are
copied to `pData`, along with an extra availability or status value
written directly after the results of each query and interpreted as an
unsigned integer. A value of zero indicates that the results are not yet
available, otherwise the query is complete and results are available.
The size of the availability or status values is 64 bits if
`VK_QUERY_RESULT_64_BIT` is set in `flags`. Otherwise, it is 32 bits.

If `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR` is set, results for all queries
in `queryPool` identified by `firstQuery` and `queryCount` are copied to
`pData`, along with an extra status value written directly after the
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

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>If <code>VK_QUERY_RESULT_WITH_AVAILABILITY_BIT</code> or
<code>VK_QUERY_RESULT_WITH_STATUS_BIT_KHR</code> is set, the layout of
data in the buffer is a <em>(result,availability)</em> or
<em>(result,status)</em> pair for each query returned, and
<code>stride</code> is the stride between each pair.</p></td>
</tr>
</tbody>
</table>

Results for any available query written by this command are final and
represent the final result of the query. If
`VK_QUERY_RESULT_PARTIAL_BIT` is set, then for any query that is
unavailable, an intermediate result between zero and the final result
value is written for that query. Otherwise, any result written by this
command is undefined.

If `VK_QUERY_RESULT_64_BIT` is set, results and, if returned,
availability or status values for all queries are written as an array of
64-bit values. If the `queryPool` was created with
`VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, results for each query are
written as an array of the type indicated by
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

If `VK_QUERY_RESULT_WAIT_BIT` is set, this command defines an execution
dependency with any earlier commands that writes one of the identified
queries. The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
instances of [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQuery.html),
[vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQueryIndexedEXT.html),
[vkCmdWriteTimestamp2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp2.html), and
[vkCmdWriteTimestamp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp.html) that reference any query
in `queryPool` indicated by `firstQuery` and `queryCount`. The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes the
host operations of this command.

If `VK_QUERY_RESULT_WAIT_BIT` is not set, `vkGetQueryPoolResults`
**may** return `VK_NOT_READY` if there are queries in the unavailable
state.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Applications <strong>must</strong> take care to ensure that use of
the <code>VK_QUERY_RESULT_WAIT_BIT</code> bit has the desired
effect.</p>
<p>For example, if a query has been used previously and a command buffer
records the commands <code>vkCmdResetQueryPool</code>,
<code>vkCmdBeginQuery</code>, and <code>vkCmdEndQuery</code> for that
query, then the query will remain in the available state until
<code>vkResetQueryPool</code> is called or the
<code>vkCmdResetQueryPool</code> command executes on a queue.
Applications <strong>can</strong> use fences or events to ensure that a
query has already been reset before checking for its results or
availability status. Otherwise, a stale value could be returned from a
previous use of the query.</p>
<p>The above also applies when <code>VK_QUERY_RESULT_WAIT_BIT</code> is
used in combination with
<code>VK_QUERY_RESULT_WITH_AVAILABILITY_BIT</code>. In this case, the
returned availability status <strong>may</strong> reflect the result of
a previous use of the query unless <code>vkResetQueryPool</code> is
called or the <code>vkCmdResetQueryPool</code> command has been executed
since the last use of the query.</p>
<p>A similar situation can arise with the
<code>VK_QUERY_RESULT_WITH_STATUS_BIT_KHR</code> flag.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Applications <strong>can</strong> double-buffer query pool usage,
with a pool per frame, and reset queries at the end of the frame in
which they are read.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkGetQueryPoolResults-firstQuery-09436"
  id="VUID-vkGetQueryPoolResults-firstQuery-09436"></a>
  VUID-vkGetQueryPoolResults-firstQuery-09436  
  `firstQuery` **must** be less than the number of queries in
  `queryPool`

- <a href="#VUID-vkGetQueryPoolResults-firstQuery-09437"
  id="VUID-vkGetQueryPoolResults-firstQuery-09437"></a>
  VUID-vkGetQueryPoolResults-firstQuery-09437  
  The sum of `firstQuery` and `queryCount` **must** be less than or
  equal to the number of queries in `queryPool`

<!-- -->

- <a href="#VUID-vkGetQueryPoolResults-queryCount-09438"
  id="VUID-vkGetQueryPoolResults-queryCount-09438"></a>
  VUID-vkGetQueryPoolResults-queryCount-09438  
  If `queryCount` is greater than 1, `stride` **must** not be zero

- <a href="#VUID-vkGetQueryPoolResults-queryType-09439"
  id="VUID-vkGetQueryPoolResults-queryType-09439"></a>
  VUID-vkGetQueryPoolResults-queryType-09439  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_TIMESTAMP`, `flags` **must** not contain
  `VK_QUERY_RESULT_PARTIAL_BIT`

- <a href="#VUID-vkGetQueryPoolResults-queryType-09440"
  id="VUID-vkGetQueryPoolResults-queryType-09440"></a>
  VUID-vkGetQueryPoolResults-queryType-09440  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, `flags` **must** not contain
  `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT`,
  `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`, `VK_QUERY_RESULT_PARTIAL_BIT`,
  or `VK_QUERY_RESULT_64_BIT`

- <a href="#VUID-vkGetQueryPoolResults-queryType-09441"
  id="VUID-vkGetQueryPoolResults-queryType-09441"></a>
  VUID-vkGetQueryPoolResults-queryType-09441  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the `queryPool` **must** have
  been recorded once for each pass as retrieved via a call to
  [vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html)

- <a href="#VUID-vkGetQueryPoolResults-queryType-09442"
  id="VUID-vkGetQueryPoolResults-queryType-09442"></a>
  VUID-vkGetQueryPoolResults-queryType-09442  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR`, then `flags` **must** include
  `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`

- <a href="#VUID-vkGetQueryPoolResults-flags-09443"
  id="VUID-vkGetQueryPoolResults-flags-09443"></a>
  VUID-vkGetQueryPoolResults-flags-09443  
  If `flags` includes `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`, then it
  **must** not include `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT`

<!-- -->

- <a href="#VUID-vkGetQueryPoolResults-None-09401"
  id="VUID-vkGetQueryPoolResults-None-09401"></a>
  VUID-vkGetQueryPoolResults-None-09401  
  All queries used by the command **must** not be uninitialized

- <a href="#VUID-vkGetQueryPoolResults-flags-02828"
  id="VUID-vkGetQueryPoolResults-flags-02828"></a>
  VUID-vkGetQueryPoolResults-flags-02828  
  If `VK_QUERY_RESULT_64_BIT` is not set in `flags` and the `queryType`
  used to create `queryPool` was not
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, then `pData` and `stride`
  **must** be multiples of `4`

- <a href="#VUID-vkGetQueryPoolResults-flags-00815"
  id="VUID-vkGetQueryPoolResults-flags-00815"></a>
  VUID-vkGetQueryPoolResults-flags-00815  
  If `VK_QUERY_RESULT_64_BIT` is set in `flags` then `pData` and
  `stride` **must** be multiples of `8`

- <a href="#VUID-vkGetQueryPoolResults-stride-08993"
  id="VUID-vkGetQueryPoolResults-stride-08993"></a>
  VUID-vkGetQueryPoolResults-stride-08993  
  If `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT` is set, `stride` **must**
  be large enough to contain the unsigned integer representing
  availability or status in addition to the query result

- <a href="#VUID-vkGetQueryPoolResults-queryType-03229"
  id="VUID-vkGetQueryPoolResults-queryType-03229"></a>
  VUID-vkGetQueryPoolResults-queryType-03229  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, then `pData` and `stride`
  **must** be multiples of the size of
  [VkPerformanceCounterResultKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterResultKHR.html)

- <a href="#VUID-vkGetQueryPoolResults-queryType-04519"
  id="VUID-vkGetQueryPoolResults-queryType-04519"></a>
  VUID-vkGetQueryPoolResults-queryType-04519  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, then `stride` **must** be large
  enough to contain the
  [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolPerformanceCreateInfoKHR.html)::`counterIndexCount`
  used to create `queryPool` times the size of
  [VkPerformanceCounterResultKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterResultKHR.html)

- <a href="#VUID-vkGetQueryPoolResults-dataSize-00817"
  id="VUID-vkGetQueryPoolResults-dataSize-00817"></a>
  VUID-vkGetQueryPoolResults-dataSize-00817  
  `dataSize` **must** be large enough to contain the result of each
  query, as described <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-memorylayout"
  target="_blank" rel="noopener">here</a>

Valid Usage (Implicit)

- <a href="#VUID-vkGetQueryPoolResults-device-parameter"
  id="VUID-vkGetQueryPoolResults-device-parameter"></a>
  VUID-vkGetQueryPoolResults-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetQueryPoolResults-queryPool-parameter"
  id="VUID-vkGetQueryPoolResults-queryPool-parameter"></a>
  VUID-vkGetQueryPoolResults-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

- <a href="#VUID-vkGetQueryPoolResults-pData-parameter"
  id="VUID-vkGetQueryPoolResults-pData-parameter"></a>
  VUID-vkGetQueryPoolResults-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes

- <a href="#VUID-vkGetQueryPoolResults-flags-parameter"
  id="VUID-vkGetQueryPoolResults-flags-parameter"></a>
  VUID-vkGetQueryPoolResults-flags-parameter  
  `flags` **must** be a valid combination of
  [VkQueryResultFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlagBits.html) values

- <a href="#VUID-vkGetQueryPoolResults-dataSize-arraylength"
  id="VUID-vkGetQueryPoolResults-dataSize-arraylength"></a>
  VUID-vkGetQueryPoolResults-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

- <a href="#VUID-vkGetQueryPoolResults-queryPool-parent"
  id="VUID-vkGetQueryPoolResults-queryPool-parent"></a>
  VUID-vkGetQueryPoolResults-queryPool-parent  
  `queryPool` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_NOT_READY`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html),
[VkQueryResultFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetQueryPoolResults"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
