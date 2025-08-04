# vkCmdCopyQueryPoolResults(3) Manual Page

## Name

vkCmdCopyQueryPoolResults - Copy the results of queries in a query pool to a buffer object



## [](#_c_specification)C Specification

To copy query statuses and numerical results directly to buffer memory, call:

```c++
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

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which this command will be recorded.
- `queryPool` is the query pool managing the queries containing the desired results.
- `firstQuery` is the initial query index.
- `queryCount` is the number of queries. `firstQuery` and `queryCount` together define a range of queries.
- `dstBuffer` is a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) object that will receive the results of the copy command.
- `dstOffset` is an offset into `dstBuffer`.
- `stride` is the stride in bytes between results for individual queries within `dstBuffer`. The required size of the backing memory for `dstBuffer` is determined as described above for [vkGetQueryPoolResults](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetQueryPoolResults.html).
- `flags` is a bitmask of [VkQueryResultFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryResultFlagBits.html) specifying how and when results are returned.

## [](#_description)Description

Any results written for a query are written according to [a layout dependent on the query type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-operation-memorylayout).

Results for any query in `queryPool` identified by `firstQuery` and `queryCount` that is available are copied to `dstBuffer`.

If `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT` is set, results for all queries in `queryPool` identified by `firstQuery` and `queryCount` are copied to `dstBuffer`, along with an extra availability value written directly after the results of each query and interpreted as an unsigned integer. A value of zero indicates that the results are not yet available, otherwise the query is complete and results are available.

If `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR` is set, results for all queries in `queryPool` identified by `firstQuery` and `queryCount` are copied to `dstBuffer`, along with an extra status value written directly after the results of each query and interpreted as a signed integer. A value of zero indicates that the results are not yet available. Positive values indicate that the operations within the query completed successfully, and the query results are valid. Negative values indicate that the operations within the query completed unsuccessfully.

[VkQueryResultStatusKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryResultStatusKHR.html) defines specific meaning for values returned here, though implementations are free to return other values.

If the status value written is negative, indicating that the operations within the query completed unsuccessfully, then all other results written by this command are undefined unless otherwise specified for any of the results of the used query type.

Results for any available query written by this command are final and represent the final result of the query. If `VK_QUERY_RESULT_PARTIAL_BIT` is set, then for any query that is unavailable, an intermediate result between zero and the final result value is written for that query. Otherwise, any result written by this command is undefined.

If `VK_QUERY_RESULT_64_BIT` is set, results and availability or status values for all queries are written as an array of 64-bit values. If the `queryPool` was created with `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, results for each query are written as an array of the type indicated by [VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterKHR.html)::`storage` for the counter being queried. Otherwise, results and availability or status values are written as an array of 32-bit values. If an unsigned integer query’s value overflows the result type, the value **may** either wrap or saturate. If the [`maintenance7`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance7) feature is enabled, for an unsigned integer query, the 32-bit result value **must** be equal to the 32 least significant bits of the equivalent 64-bit result value. If a signed integer query’s value overflows the result type, the value is undefined. If a floating-point query’s value is not representable as the result type, the value is undefined.

This command defines an execution dependency between other query commands that reference the same query.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands which reference the queries in `queryPool` indicated by `query` that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order). If `flags` does not include `VK_QUERY_RESULT_WAIT_BIT`, [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQuery.html), [vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQueryIndexedEXT.html), [vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html), [vkCmdWriteAccelerationStructuresPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesNV.html), [vkCmdWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteMicromapsPropertiesEXT.html), [vkCmdWriteTimestamp2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteTimestamp2.html), and [vkCmdWriteTimestamp](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteTimestamp.html) are excluded from this scope.

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands which reference the queries in `queryPool` indicated by `query` that occur later in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

The operation of this command happens after the first scope and happens before the second scope.

`vkCmdCopyQueryPoolResults` is considered to be a transfer operation, and its writes to buffer memory **must** be synchronized using `VK_PIPELINE_STAGE_TRANSFER_BIT` and `VK_ACCESS_TRANSFER_WRITE_BIT` before using the results.

Valid Usage

- [](#VUID-vkCmdCopyQueryPoolResults-firstQuery-09436)VUID-vkCmdCopyQueryPoolResults-firstQuery-09436  
  `firstQuery` **must** be less than the number of queries in `queryPool`
- [](#VUID-vkCmdCopyQueryPoolResults-firstQuery-09437)VUID-vkCmdCopyQueryPoolResults-firstQuery-09437  
  The sum of `firstQuery` and `queryCount` **must** be less than or equal to the number of queries in `queryPool`

<!--THE END-->

- [](#VUID-vkCmdCopyQueryPoolResults-queryCount-09438)VUID-vkCmdCopyQueryPoolResults-queryCount-09438  
  If `queryCount` is greater than 1, `stride` **must** not be zero
- [](#VUID-vkCmdCopyQueryPoolResults-queryType-09439)VUID-vkCmdCopyQueryPoolResults-queryType-09439  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_TIMESTAMP`, `flags` **must** not contain `VK_QUERY_RESULT_PARTIAL_BIT`
- [](#VUID-vkCmdCopyQueryPoolResults-queryType-09440)VUID-vkCmdCopyQueryPoolResults-queryType-09440  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, `flags` **must** not contain `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT`, `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`, `VK_QUERY_RESULT_PARTIAL_BIT`, or `VK_QUERY_RESULT_64_BIT`
- [](#VUID-vkCmdCopyQueryPoolResults-queryType-09441)VUID-vkCmdCopyQueryPoolResults-queryType-09441  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the `queryPool` **must** have been recorded once for each pass as retrieved via a call to [vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html)
- [](#VUID-vkCmdCopyQueryPoolResults-queryType-09442)VUID-vkCmdCopyQueryPoolResults-queryType-09442  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR`, then `flags` **must** include `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`
- [](#VUID-vkCmdCopyQueryPoolResults-flags-09443)VUID-vkCmdCopyQueryPoolResults-flags-09443  
  If `flags` includes `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`, then it **must** not include `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT`

<!--THE END-->

- [](#VUID-vkCmdCopyQueryPoolResults-None-09402)VUID-vkCmdCopyQueryPoolResults-None-09402  
  All queries used by the command **must** not be uninitialized when the command is executed
- [](#VUID-vkCmdCopyQueryPoolResults-dstOffset-00819)VUID-vkCmdCopyQueryPoolResults-dstOffset-00819  
  `dstOffset` **must** be less than the size of `dstBuffer`
- [](#VUID-vkCmdCopyQueryPoolResults-flags-00822)VUID-vkCmdCopyQueryPoolResults-flags-00822  
  If `VK_QUERY_RESULT_64_BIT` is not set in `flags` then `dstOffset` and `stride` **must** be multiples of `4`
- [](#VUID-vkCmdCopyQueryPoolResults-flags-00823)VUID-vkCmdCopyQueryPoolResults-flags-00823  
  If `VK_QUERY_RESULT_64_BIT` is set in `flags` then `dstOffset` and `stride` **must** be multiples of `8`
- [](#VUID-vkCmdCopyQueryPoolResults-dstBuffer-00824)VUID-vkCmdCopyQueryPoolResults-dstBuffer-00824  
  `dstBuffer` **must** have enough storage, from `dstOffset`, to contain the result of each query, as described [here](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-operation-memorylayout)
- [](#VUID-vkCmdCopyQueryPoolResults-dstBuffer-00825)VUID-vkCmdCopyQueryPoolResults-dstBuffer-00825  
  `dstBuffer` **must** have been created with `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-vkCmdCopyQueryPoolResults-dstBuffer-00826)VUID-vkCmdCopyQueryPoolResults-dstBuffer-00826  
  If `dstBuffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdCopyQueryPoolResults-queryType-03232)VUID-vkCmdCopyQueryPoolResults-queryType-03232  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, [VkPhysicalDevicePerformanceQueryPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePerformanceQueryPropertiesKHR.html)::`allowCommandBufferQueryCopies` **must** be `VK_TRUE`
- [](#VUID-vkCmdCopyQueryPoolResults-queryType-02734)VUID-vkCmdCopyQueryPoolResults-queryType-02734  
  [vkCmdCopyQueryPoolResults](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyQueryPoolResults.html) **must** not be called if the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PERFORMANCE_QUERY_INTEL`
- [](#VUID-vkCmdCopyQueryPoolResults-None-07429)VUID-vkCmdCopyQueryPoolResults-None-07429  
  All queries used by the command **must** not be active
- [](#VUID-vkCmdCopyQueryPoolResults-None-08752)VUID-vkCmdCopyQueryPoolResults-None-08752  
  All queries used by the command **must** have been made *available* by prior executed commands

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyQueryPoolResults-commandBuffer-parameter)VUID-vkCmdCopyQueryPoolResults-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyQueryPoolResults-queryPool-parameter)VUID-vkCmdCopyQueryPoolResults-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html) handle
- [](#VUID-vkCmdCopyQueryPoolResults-dstBuffer-parameter)VUID-vkCmdCopyQueryPoolResults-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-vkCmdCopyQueryPoolResults-flags-parameter)VUID-vkCmdCopyQueryPoolResults-flags-parameter  
  `flags` **must** be a valid combination of [VkQueryResultFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryResultFlagBits.html) values
- [](#VUID-vkCmdCopyQueryPoolResults-commandBuffer-recording)VUID-vkCmdCopyQueryPoolResults-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyQueryPoolResults-commandBuffer-cmdpool)VUID-vkCmdCopyQueryPoolResults-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdCopyQueryPoolResults-renderpass)VUID-vkCmdCopyQueryPoolResults-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyQueryPoolResults-videocoding)VUID-vkCmdCopyQueryPoolResults-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdCopyQueryPoolResults-commonparent)VUID-vkCmdCopyQueryPoolResults-commonparent  
  Each of `commandBuffer`, `dstBuffer`, and `queryPool` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Graphics  
Compute

Action

Conditional Rendering

vkCmdCopyQueryPoolResults is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html), [VkQueryResultFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryResultFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyQueryPoolResults)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0