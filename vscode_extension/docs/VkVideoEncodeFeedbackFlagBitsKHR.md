# VkVideoEncodeFeedbackFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeFeedbackFlagBitsKHR - Bits specifying queried video encode feedback values



## [](#_c_specification)C Specification

Bits which **can** be set in [VkQueryPoolVideoEncodeFeedbackCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolVideoEncodeFeedbackCreateInfoKHR.html)::`encodeFeedbackFlags` for video encode feedback query pools are:

```c++
// Provided by VK_KHR_video_encode_queue
typedef enum VkVideoEncodeFeedbackFlagBitsKHR {
    VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_BUFFER_OFFSET_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_BYTES_WRITTEN_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_HAS_OVERRIDES_BIT_KHR = 0x00000004,
} VkVideoEncodeFeedbackFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_BUFFER_OFFSET_BIT_KHR` specifies that queries managed by the pool will capture the byte offset of the bitstream data written by the video encode operation to the bitstream buffer specified in [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`dstBuffer` relative to the offset specified in [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`dstBufferOffset`. For the first video encode operation issued by any [video encode command](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-encode-commands), this value will always be zero, meaning that bitstream data is always written to the buffer specified in [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`dstBuffer` starting from the offset specified in [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`dstBufferOffset`.
- `VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_BYTES_WRITTEN_BIT_KHR` specifies that queries managed by the pool will capture the number of bytes written by the video encode operation to the bitstream buffer specified in [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`dstBuffer`.
- `VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_HAS_OVERRIDES_BIT_KHR` specifies that queries managed by the pool will capture a boolean value indicating that the data written to the bitstream buffer specified in [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`dstBuffer` contains [overridden parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides).

When retrieving the results of video encode feedback queries, the values corresponding to each enabled video encode feedback are written in the order of the bits defined above, followed by an optional value indicating availability or result status if `VK_QUERY_RESULT_WITH_AVAILABILITY_BIT` or `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR` is specified, respectively.

If the result status of a video encode feedback query is negative, then the results of all enabled video encode feedback values will be undefined.

Note

Applications should always specify `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR` when retrieving the results of video encode feedback queries and ignore such undefined video encode feedback values for any [unsuccessfully](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-unsuccessful) completed video encode operations.

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkVideoEncodeFeedbackFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeFeedbackFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeFeedbackFlagBitsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0