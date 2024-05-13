# VkVideoEncodeUsageFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeUsageFlagBitsKHR - Video encode usage flags



## <a href="#_c_specification" class="anchor"></a>C Specification

The following bits **can** be specified in
[VkVideoEncodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeUsageInfoKHR.html)::`videoUsageHints`
as a hint about the video encode use case:

``` c
// Provided by VK_KHR_video_encode_queue
typedef enum VkVideoEncodeUsageFlagBitsKHR {
    VK_VIDEO_ENCODE_USAGE_DEFAULT_KHR = 0,
    VK_VIDEO_ENCODE_USAGE_TRANSCODING_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_USAGE_STREAMING_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_USAGE_RECORDING_BIT_KHR = 0x00000004,
    VK_VIDEO_ENCODE_USAGE_CONFERENCING_BIT_KHR = 0x00000008,
} VkVideoEncodeUsageFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_ENCODE_USAGE_TRANSCODING_BIT_KHR` specifies that video
  encoding is intended to be used in conjunction with video decoding to
  transcode a video bitstream with the same and/or different codecs.

- `VK_VIDEO_ENCODE_USAGE_STREAMING_BIT_KHR` specifies that video
  encoding is intended to be used to produce a video bitstream that is
  expected to be sent as a continuous flow over network.

- `VK_VIDEO_ENCODE_USAGE_RECORDING_BIT_KHR` specifies that video
  encoding is intended to be used for real-time recording for offline
  consumption.

- `VK_VIDEO_ENCODE_USAGE_CONFERENCING_BIT_KHR` specifies that video
  encoding is intended to be used in a video conferencing scenario.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>There are no restrictions on the combination of bits that
<strong>can</strong> be specified by the application. However,
applications <strong>should</strong> use reasonable combinations in
order for the implementation to be able to select the most appropriate
mode of operation for the particular use case.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkVideoEncodeUsageFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeUsageFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeUsageFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
