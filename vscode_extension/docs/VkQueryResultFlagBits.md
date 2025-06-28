# VkQueryResultFlagBits(3) Manual Page

## Name

VkQueryResultFlagBits - Bitmask specifying how and when query results are returned



## [](#_c_specification)C Specification

Bits which **can** be set in [vkGetQueryPoolResults](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetQueryPoolResults.html)::`flags` and [vkCmdCopyQueryPoolResults](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyQueryPoolResults.html)::`flags`, specifying how and when results are returned, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkQueryResultFlagBits {
    VK_QUERY_RESULT_64_BIT = 0x00000001,
    VK_QUERY_RESULT_WAIT_BIT = 0x00000002,
    VK_QUERY_RESULT_WITH_AVAILABILITY_BIT = 0x00000004,
    VK_QUERY_RESULT_PARTIAL_BIT = 0x00000008,
  // Provided by VK_KHR_video_queue
    VK_QUERY_RESULT_WITH_STATUS_BIT_KHR = 0x00000010,
} VkQueryResultFlagBits;
```

## [](#_description)Description

- `VK_QUERY_RESULT_64_BIT` specifies the results will be written as an array of 64-bit unsigned integer values. If this bit is not set, the results will be written as an array of 32-bit unsigned integer values.
- `VK_QUERY_RESULT_WAIT_BIT` specifies that Vulkan will wait for each query’s status to become available before retrieving its results.
- `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT` specifies that the availability status accompanies the results.
- `VK_QUERY_RESULT_PARTIAL_BIT` specifies that returning partial results is acceptable.
- `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR` specifies that the last value returned in the query is a [VkQueryResultStatusKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryResultStatusKHR.html) value. See [result status query](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-result-status-only) for information on how an application can determine whether the use of this flag bit is supported.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkQueryResultFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryResultFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueryResultFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0