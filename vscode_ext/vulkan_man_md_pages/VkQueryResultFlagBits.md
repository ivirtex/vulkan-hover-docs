# VkQueryResultFlagBits(3) Manual Page

## Name

VkQueryResultFlagBits - Bitmask specifying how and when query results
are returned



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[vkGetQueryPoolResults](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetQueryPoolResults.html)::`flags` and
[vkCmdCopyQueryPoolResults](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyQueryPoolResults.html)::`flags`,
specifying how and when results are returned, are:

``` c
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

## <a href="#_description" class="anchor"></a>Description

- `VK_QUERY_RESULT_64_BIT` specifies the results will be written as an
  array of 64-bit unsigned integer values. If this bit is not set, the
  results will be written as an array of 32-bit unsigned integer values.

- `VK_QUERY_RESULT_WAIT_BIT` specifies that Vulkan will wait for each
  query’s status to become available before retrieving its results.

- `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT` specifies that the
  availability status accompanies the results.

- `VK_QUERY_RESULT_PARTIAL_BIT` specifies that returning partial results
  is acceptable.

- `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR` specifies that the last value
  returned in the query is a
  [VkQueryResultStatusKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultStatusKHR.html) value. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-result-status-only"
  target="_blank" rel="noopener">result status query</a> for information
  on how an application can determine whether the use of this flag bit
  is supported.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkQueryResultFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueryResultFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
