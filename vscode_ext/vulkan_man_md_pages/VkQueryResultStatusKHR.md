# VkQueryResultStatusKHR(3) Manual Page

## Name

VkQueryResultStatusKHR - Specific status codes for operations



## <a href="#_c_specification" class="anchor"></a>C Specification

Specific status codes that **can** be returned from a query are:

``` c
// Provided by VK_KHR_video_queue
typedef enum VkQueryResultStatusKHR {
    VK_QUERY_RESULT_STATUS_ERROR_KHR = -1,
    VK_QUERY_RESULT_STATUS_NOT_READY_KHR = 0,
    VK_QUERY_RESULT_STATUS_COMPLETE_KHR = 1,
  // Provided by VK_KHR_video_encode_queue
    VK_QUERY_RESULT_STATUS_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_KHR = -1000299000,
} VkQueryResultStatusKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_QUERY_RESULT_STATUS_NOT_READY_KHR` indicates that the query result
  is not yet available.

- `VK_QUERY_RESULT_STATUS_ERROR_KHR` indicates that operations did not
  complete successfully.

- `VK_QUERY_RESULT_STATUS_COMPLETE_KHR` indicates that operations
  completed successfully and the query result is available.

- `VK_QUERY_RESULT_STATUS_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_KHR`
  indicates that a video encode operation did not complete successfully
  due to the destination video bitstream buffer range not being
  sufficiently large to fit the encoded bitstream data.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueryResultStatusKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
