# VkVideoEncodeFeedbackFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeFeedbackFlagBitsKHR - Bits specifying queried video encode
feedback values



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkQueryPoolVideoEncodeFeedbackCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolVideoEncodeFeedbackCreateInfoKHR.html)::`encodeFeedbackFlags`
for video encode feedback query pools are:

``` c
// Provided by VK_KHR_video_encode_queue
typedef enum VkVideoEncodeFeedbackFlagBitsKHR {
    VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_BUFFER_OFFSET_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_BYTES_WRITTEN_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_HAS_OVERRIDES_BIT_KHR = 0x00000004,
} VkVideoEncodeFeedbackFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_BUFFER_OFFSET_BIT_KHR` specifies
  that queries managed by the pool will capture the byte offset of the
  bitstream data written by the video encode operation to the bitstream
  buffer specified in
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`dstBuffer`
  relative to the offset specified in
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`dstBufferOffset`.
  For the first video encode operation issued by any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-encode-commands"
  target="_blank" rel="noopener">video encode command</a>, this value
  will always be zero, meaning that bitstream data is always written to
  the buffer specified in
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`dstBuffer`
  starting from the offset specified in
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`dstBufferOffset`.

- `VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_BYTES_WRITTEN_BIT_KHR` specifies
  that queries managed by the pool will capture the number of bytes
  written by the video encode operation to the bitstream buffer
  specified in
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`dstBuffer`.

- `VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_HAS_OVERRIDES_BIT_KHR` specifies
  that queries managed by the pool will capture a boolean value
  indicating that the data written to the bitstream buffer specified in
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`dstBuffer`
  contains <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-overrides"
  target="_blank" rel="noopener">overridden parameters</a>.

When retrieving the results of video encode feedback queries, the values
corresponding to each enabled video encode feedback are written in the
order of the bits defined above, followed by an optional value
indicating availability or result status if
`VK_QUERY_RESULT_WITH_AVAILABILITY_BIT` or
`VK_QUERY_RESULT_WITH_STATUS_BIT_KHR` is specified, respectively.

If the result status of a video encode feedback query is negative, then
the results of all enabled video encode feedback values will be
undefined.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Thus it is recommended that applications always specify
<code>VK_QUERY_RESULT_WITH_STATUS_BIT_KHR</code> when retrieving the
results of video encode feedback queries and ignore such undefined video
encode feedback values for any <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-unsuccessful"
target="_blank" rel="noopener">unsuccessfully</a> completed video encode
operations.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkVideoEncodeFeedbackFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeFeedbackFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeFeedbackFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
