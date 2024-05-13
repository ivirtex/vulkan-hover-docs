# VkVideoEncodeCapabilityFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeCapabilityFlagBitsKHR - Video encode capability flags



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilitiesKHR.html)::`flags`,
indicating the encoding tools supported, are:

``` c
// Provided by VK_KHR_video_encode_queue
typedef enum VkVideoEncodeCapabilityFlagBitsKHR {
    VK_VIDEO_ENCODE_CAPABILITY_PRECEDING_EXTERNALLY_ENCODED_BYTES_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_CAPABILITY_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_DETECTION_BIT_KHR = 0x00000002,
} VkVideoEncodeCapabilityFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_ENCODE_CAPABILITY_PRECEDING_EXTERNALLY_ENCODED_BYTES_BIT_KHR`
  indicates that the implementation supports the use of
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`precedingExternallyEncodedBytes`.

- `VK_VIDEO_ENCODE_CAPABILITY_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_DETECTION_BIT_KHR`
  indicates that the implementation is able to detect and report when
  the destination video bitstream buffer range provided by the
  application is not sufficiently large to fit the encoded bitstream
  data produced by a video encode operation by reporting the
  `VK_QUERY_RESULT_STATUS_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_KHR` <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#query-result-status-codes"
  target="_blank" rel="noopener">query result status code</a>.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr class="odd">
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>Some implementations <strong>may</strong> not be able to reliably
  detect insufficient bitstream buffer range conditions in all situations.
  Such implementations will not report support for the
  <code>VK_VIDEO_ENCODE_CAPABILITY_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_DETECTION_BIT_KHR</code>
  encode capability flag for the video profile, but <strong>may</strong>
  still report the
  <code>VK_QUERY_RESULT_STATUS_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_KHR</code>
  query result status code in certain cases. Applications
  <strong>should</strong> always check for the specific query result
  status code
  <code>VK_QUERY_RESULT_STATUS_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_KHR</code>
  even when this encode capability flag is not supported by the
  implementation for the video profile in question. However, applications
  <strong>must</strong> not assume that a different negative query result
  status code indicating an unsuccessful completion of a video encode
  operation is not the result of an insufficient bitstream buffer
  condition unless this encode capability flag is supported.</p></td>
  </tr>
  </tbody>
  </table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkVideoEncodeCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilityFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeCapabilityFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
