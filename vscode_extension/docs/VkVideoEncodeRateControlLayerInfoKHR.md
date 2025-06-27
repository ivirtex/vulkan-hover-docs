# VkVideoEncodeRateControlLayerInfoKHR(3) Manual Page

## Name

VkVideoEncodeRateControlLayerInfoKHR - Structure to set encode per-layer
rate control parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeRateControlLayerInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeRateControlLayerInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           averageBitrate;
    uint64_t           maxBitrate;
    uint32_t           frameRateNumerator;
    uint32_t           frameRateDenominator;
} VkVideoEncodeRateControlLayerInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is a pointer to a structure extending this structure.

- `averageBitrate` is the average <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-bitrate"
  target="_blank" rel="noopener">bitrate</a> to be targeted by the
  implementation’s rate control algorithm.

- `maxBitrate` is the peak <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-bitrate"
  target="_blank" rel="noopener">bitrate</a> to be targeted by the
  implementation’s rate control algorithm.

- `frameRateNumerator` is the numerator of the frame rate assumed by the
  implementation’s rate control algorithm.

- `frameRateDenominator` is the denominator of the frame rate assumed by
  the implementation’s rate control algorithm.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The ability of the implementation’s rate control algorithm to be able
to match the requested average and/or peak bitrates <strong>may</strong>
be limited by the set of other codec-independent and codec-specific rate
control parameters specified by the application, the input content, as
well as the application conforming to the rate control guidance provided
to the implementation, as described <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control"
target="_blank" rel="noopener">earlier</a>.</p></td>
</tr>
</tbody>
</table>

Additional structures providing codec-specific rate control parameters
**can** be included in the `pNext` chain of
`VkVideoEncodeRateControlLayerInfoKHR` depending on the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profiles"
target="_blank" rel="noopener">video profile</a> the bound video session
was created with. For further details see:

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-coding-control"
  target="_blank" rel="noopener">Video Coding Control</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-rate-control"
  target="_blank" rel="noopener">H.264 Encode Rate Control</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-rate-control"
  target="_blank" rel="noopener">H.265 Encode Rate Control</a>

Valid Usage

- <a
  href="#VUID-VkVideoEncodeRateControlLayerInfoKHR-frameRateNumerator-08350"
  id="VUID-VkVideoEncodeRateControlLayerInfoKHR-frameRateNumerator-08350"></a>
  VUID-VkVideoEncodeRateControlLayerInfoKHR-frameRateNumerator-08350  
  `frameRateNumerator` **must** be greater than zero

- <a
  href="#VUID-VkVideoEncodeRateControlLayerInfoKHR-frameRateDenominator-08351"
  id="VUID-VkVideoEncodeRateControlLayerInfoKHR-frameRateDenominator-08351"></a>
  VUID-VkVideoEncodeRateControlLayerInfoKHR-frameRateDenominator-08351  
  `frameRateDenominator` **must** be greater than zero

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeRateControlLayerInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeRateControlLayerInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeRateControlLayerInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_RATE_CONTROL_LAYER_INFO_KHR`

- <a href="#VUID-VkVideoEncodeRateControlLayerInfoKHR-pNext-pNext"
  id="VUID-VkVideoEncodeRateControlLayerInfoKHR-pNext-pNext"></a>
  VUID-VkVideoEncodeRateControlLayerInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkVideoEncodeH264RateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlLayerInfoKHR.html)
  or
  [VkVideoEncodeH265RateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlLayerInfoKHR.html)

- <a href="#VUID-VkVideoEncodeRateControlLayerInfoKHR-sType-unique"
  id="VUID-VkVideoEncodeRateControlLayerInfoKHR-sType-unique"></a>
  VUID-VkVideoEncodeRateControlLayerInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeRateControlLayerInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
