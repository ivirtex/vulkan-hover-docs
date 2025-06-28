# VkQueryResultStatusKHR(3) Manual Page

## Name

VkQueryResultStatusKHR - Specific status codes for operations



## [](#_c_specification)C Specification

Specific status codes that **can** be returned from a query are:

```c++
// Provided by VK_KHR_video_queue
typedef enum VkQueryResultStatusKHR {
    VK_QUERY_RESULT_STATUS_ERROR_KHR = -1,
    VK_QUERY_RESULT_STATUS_NOT_READY_KHR = 0,
    VK_QUERY_RESULT_STATUS_COMPLETE_KHR = 1,
  // Provided by VK_KHR_video_encode_queue
    VK_QUERY_RESULT_STATUS_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_KHR = -1000299000,
} VkQueryResultStatusKHR;
```

## [](#_description)Description

- `VK_QUERY_RESULT_STATUS_NOT_READY_KHR` specifies that the query result is not yet available.
- `VK_QUERY_RESULT_STATUS_ERROR_KHR` specifies that operations did not complete successfully.
- `VK_QUERY_RESULT_STATUS_COMPLETE_KHR` specifies that operations completed successfully and the query result is available.
- `VK_QUERY_RESULT_STATUS_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_KHR` specifies that a video encode operation did not complete successfully due to the destination video bitstream buffer range not being sufficiently large to fit the encoded bitstream data.

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueryResultStatusKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0