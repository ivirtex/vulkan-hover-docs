# vkGetQueryPoolResults(3) Manual Page

## Name

vkGetQueryPoolResults - Copy results of queries in a query pool to a host memory region



## [](#_c_specification)C Specification

To retrieve status and results for a set of queries, call:

```c++
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

## [](#_parameters)Parameters

- `device` is the logical device that owns the query pool.
- `queryPool` is the query pool managing the queries containing the desired results.
- `firstQuery` is the initial query index.
- `queryCount` is the number of queries to read.
- `dataSize` is the size in bytes of the buffer pointed to by `pData`.
- `pData` is a pointer to an application-allocated buffer where the results will be written
- `stride` is the stride in bytes between results for individual queries within `pData`.
- `flags` is a bitmask of [VkQueryResultFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryResultFlagBits.html) specifying how and when results are returned.

## [](#_description)Description

Any results written for a query are written according to [a layout dependent on the query type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-operation-memorylayout).

If no bits are set in `flags`, and all requested queries are in the available state, results are written as an array of 32-bit unsigned integer values. Behavior when not all queries are available is described [below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-wait-bit-not-set).

If `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT` is set, results for all queries in `queryPool` identified by `firstQuery` and `queryCount` are copied to `pData`, along with an extra availability or status value written directly after the results of each query and interpreted as an unsigned integer. A value of zero indicates that the results are not yet available, otherwise the query is complete and results are available. The size of the availability or status values is 64 bits if `VK_QUERY_RESULT_64_BIT` is set in `flags`. Otherwise, it is 32 bits.

If `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR` is set, results for all queries in `queryPool` identified by `firstQuery` and `queryCount` are copied to `pData`, along with an extra status value written directly after the results of each query and interpreted as a signed integer. A value of zero indicates that the results are not yet available. Positive values indicate that the operations within the query completed successfully, and the query results are valid. Negative values indicate that the operations within the query completed unsuccessfully.

[VkQueryResultStatusKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryResultStatusKHR.html) defines specific meaning for values returned here, though implementations are free to return other values.

If the status value written is negative, indicating that the operations within the query completed unsuccessfully, then all other results written by this command are undefined unless otherwise specified for any of the results of the used query type.

Note

If `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT` or `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR` is set, the layout of data in the buffer is a *(result,availability)* or *(result,status)* pair for each query returned, and `stride` is the stride between each pair.

Results for any available query written by this command are final and represent the final result of the query. If `VK_QUERY_RESULT_PARTIAL_BIT` is set, then for any query that is unavailable, an intermediate result between zero and the final result value is written for that query. Otherwise, any result written by this command is undefined.

If `VK_QUERY_RESULT_64_BIT` is set, results and, if returned, availability or status values for all queries are written as an array of 64-bit values. If the `queryPool` was created with `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, results for each query are written as an array of the type indicated by [VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterKHR.html)::`storage` for the counter being queried. Otherwise, results and availability or status values are written as an array of 32-bit values. If an unsigned integer query’s value overflows the result type, the value **may** either wrap or saturate. If the [`maintenance7`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance7) feature is enabled, for an unsigned integer query, the 32-bit result value **must** be equal to the 32 least significant bits of the equivalent 64-bit result value. If a signed integer query’s value overflows the result type, the value is undefined. If a floating-point query’s value is not representable as the result type, the value is undefined.

If `VK_QUERY_RESULT_WAIT_BIT` is set, this command defines an execution dependency with any earlier commands that writes one of the identified queries. The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all instances of [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQuery.html), [vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQueryIndexedEXT.html), [vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html), [vkCmdWriteAccelerationStructuresPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesNV.html), [vkCmdWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteMicromapsPropertiesEXT.html), [vkCmdWriteTimestamp2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteTimestamp2.html), and [vkCmdWriteTimestamp](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteTimestamp.html) that reference any query in `queryPool` indicated by `firstQuery` and `queryCount`. The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes the host operations of this command.

If `VK_QUERY_RESULT_WAIT_BIT` is not set, `vkGetQueryPoolResults` **may** return `VK_NOT_READY` if there are queries in the unavailable state.

Note

Applications **must** take care to ensure that use of the `VK_QUERY_RESULT_WAIT_BIT` bit has the desired effect.

For example, if a query has been used previously and a command buffer records the commands `vkCmdResetQueryPool`, `vkCmdBeginQuery`, and `vkCmdEndQuery` for that query, then the query will remain in the available state until `vkResetQueryPool` is called or the `vkCmdResetQueryPool` command executes on a queue. Applications **can** use fences or events to ensure that a query has already been reset before checking for its results or availability status. Otherwise, a stale value could be returned from a previous use of the query.

The above also applies when `VK_QUERY_RESULT_WAIT_BIT` is used in combination with `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT`. In this case, the returned availability status **may** reflect the result of a previous use of the query unless `vkResetQueryPool` is called or the `vkCmdResetQueryPool` command has been executed since the last use of the query.

A similar situation can arise with the `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR` flag.

Note

Applications **can** double-buffer query pool usage, with a pool per frame, and reset queries at the end of the frame in which they are read.

Valid Usage

- [](#VUID-vkGetQueryPoolResults-firstQuery-09436)VUID-vkGetQueryPoolResults-firstQuery-09436  
  `firstQuery` **must** be less than the number of queries in `queryPool`
- [](#VUID-vkGetQueryPoolResults-firstQuery-09437)VUID-vkGetQueryPoolResults-firstQuery-09437  
  The sum of `firstQuery` and `queryCount` **must** be less than or equal to the number of queries in `queryPool`

<!--THE END-->

- [](#VUID-vkGetQueryPoolResults-queryCount-09438)VUID-vkGetQueryPoolResults-queryCount-09438  
  If `queryCount` is greater than 1, `stride` **must** not be zero
- [](#VUID-vkGetQueryPoolResults-queryType-09439)VUID-vkGetQueryPoolResults-queryType-09439  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_TIMESTAMP`, `flags` **must** not contain `VK_QUERY_RESULT_PARTIAL_BIT`
- [](#VUID-vkGetQueryPoolResults-queryType-09440)VUID-vkGetQueryPoolResults-queryType-09440  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, `flags` **must** not contain `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT`, `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`, `VK_QUERY_RESULT_PARTIAL_BIT`, or `VK_QUERY_RESULT_64_BIT`
- [](#VUID-vkGetQueryPoolResults-queryType-09441)VUID-vkGetQueryPoolResults-queryType-09441  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the `queryPool` **must** have been recorded once for each pass as retrieved via a call to [vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html)
- [](#VUID-vkGetQueryPoolResults-queryType-09442)VUID-vkGetQueryPoolResults-queryType-09442  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR`, then `flags` **must** include `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`
- [](#VUID-vkGetQueryPoolResults-flags-09443)VUID-vkGetQueryPoolResults-flags-09443  
  If `flags` includes `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`, then it **must** not include `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT`

<!--THE END-->

- [](#VUID-vkGetQueryPoolResults-None-09401)VUID-vkGetQueryPoolResults-None-09401  
  All queries used by the command **must** not be uninitialized
- [](#VUID-vkGetQueryPoolResults-flags-02828)VUID-vkGetQueryPoolResults-flags-02828  
  If `VK_QUERY_RESULT_64_BIT` is not set in `flags` and the `queryType` used to create `queryPool` was not `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, then `pData` and `stride` **must** be multiples of `4`
- [](#VUID-vkGetQueryPoolResults-flags-00815)VUID-vkGetQueryPoolResults-flags-00815  
  If `VK_QUERY_RESULT_64_BIT` is set in `flags` then `pData` and `stride` **must** be multiples of `8`
- [](#VUID-vkGetQueryPoolResults-stride-08993)VUID-vkGetQueryPoolResults-stride-08993  
  If `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT` is set, `stride` **must** be large enough to contain the unsigned integer representing availability or status in addition to the query result
- [](#VUID-vkGetQueryPoolResults-queryType-03229)VUID-vkGetQueryPoolResults-queryType-03229  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, then `pData` and `stride` **must** be multiples of the size of [VkPerformanceCounterResultKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterResultKHR.html)
- [](#VUID-vkGetQueryPoolResults-queryType-04519)VUID-vkGetQueryPoolResults-queryType-04519  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, then `stride` **must** be large enough to contain the [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolPerformanceCreateInfoKHR.html)::`counterIndexCount` used to create `queryPool` times the size of [VkPerformanceCounterResultKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterResultKHR.html)
- [](#VUID-vkGetQueryPoolResults-dataSize-00817)VUID-vkGetQueryPoolResults-dataSize-00817  
  `dataSize` **must** be large enough to contain the result of each query, as described [here](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-operation-memorylayout)

Valid Usage (Implicit)

- [](#VUID-vkGetQueryPoolResults-device-parameter)VUID-vkGetQueryPoolResults-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetQueryPoolResults-queryPool-parameter)VUID-vkGetQueryPoolResults-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html) handle
- [](#VUID-vkGetQueryPoolResults-pData-parameter)VUID-vkGetQueryPoolResults-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes
- [](#VUID-vkGetQueryPoolResults-flags-parameter)VUID-vkGetQueryPoolResults-flags-parameter  
  `flags` **must** be a valid combination of [VkQueryResultFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryResultFlagBits.html) values
- [](#VUID-vkGetQueryPoolResults-dataSize-arraylength)VUID-vkGetQueryPoolResults-dataSize-arraylength  
  `dataSize` **must** be greater than `0`
- [](#VUID-vkGetQueryPoolResults-queryPool-parent)VUID-vkGetQueryPoolResults-queryPool-parent  
  `queryPool` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_NOT_READY`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_DEVICE_LOST`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html), [VkQueryResultFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryResultFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetQueryPoolResults).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0