# VkVideoEncodeH264RateControlInfoKHR(3) Manual Page

## Name

VkVideoEncodeH264RateControlInfoKHR - Structure describing H.264 stream
rate control parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeH264RateControlInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264RateControlInfoKHR {
    VkStructureType                         sType;
    const void*                             pNext;
    VkVideoEncodeH264RateControlFlagsKHR    flags;
    uint32_t                                gopFrameCount;
    uint32_t                                idrPeriod;
    uint32_t                                consecutiveBFrameCount;
    uint32_t                                temporalLayerCount;
} VkVideoEncodeH264RateControlInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkVideoEncodeH264RateControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlFlagBitsKHR.html)
  specifying H.264 rate control flags.

- `gopFrameCount` is the number of frames within a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-gop"
  target="_blank" rel="noopener">group of pictures (GOP)</a> intended to
  be used by the application. If it is set to 0, the rate control
  algorithm **may** assume an implementation-dependent GOP length. If it
  is set to `UINT32_MAX`, the GOP length is treated as infinite.

- `idrPeriod` is the interval, in terms of number of frames, between two
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-idr-pic"
  target="_blank" rel="noopener">IDR frames</a> (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-idr-period"
  target="_blank" rel="noopener">IDR period</a>). If it is set to 0, the
  rate control algorithm **may** assume an implementation-dependent IDR
  period. If it is set to `UINT32_MAX`, the IDR period is treated as
  infinite.

- `consecutiveBFrameCount` is the number of consecutive B frames between
  I and/or P frames within the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-gop"
  target="_blank" rel="noopener">GOP</a>.

- `temporalLayerCount` specifies the number of H.264 temporal layers
  that the application intends to use.

## <a href="#_description" class="anchor"></a>Description

When an instance of this structure is included in the `pNext` chain of
the [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html)
structure passed to the
[vkCmdControlVideoCodingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdControlVideoCodingKHR.html) command,
and
[VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html)::`flags`
includes `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR`, the
parameters in this structure are used as guidance for the
implementation’s rate control algorithm (see <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-coding-control"
target="_blank" rel="noopener">Video Coding Control</a>).

If `flags` includes
`VK_VIDEO_ENCODE_H264_RATE_CONTROL_ATTEMPT_HRD_COMPLIANCE_BIT_KHR`, then
the rate control state is reset to an initial state to meet HRD
compliance requirements. Otherwise the new rate control state **may** be
applied without a reset depending on the implementation and the
specified rate control parameters.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>It would be possible to infer the picture type to be used when
encoding a frame, on the basis of the values provided for
<code>consecutiveBFrameCount</code>, <code>idrPeriod</code>, and
<code>gopFrameCount</code>, but this inferred picture type will not be
used by implementations to override the picture type provided to the
video encode operation.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkVideoEncodeH264RateControlInfoKHR-flags-08280"
  id="VUID-VkVideoEncodeH264RateControlInfoKHR-flags-08280"></a>
  VUID-VkVideoEncodeH264RateControlInfoKHR-flags-08280  
  If
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`flags`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile, does not include
  `VK_VIDEO_ENCODE_H264_CAPABILITY_HRD_COMPLIANCE_BIT_KHR`, then `flags`
  **must** not contain
  `VK_VIDEO_ENCODE_H264_RATE_CONTROL_ATTEMPT_HRD_COMPLIANCE_BIT_KHR`

- <a href="#VUID-VkVideoEncodeH264RateControlInfoKHR-flags-08281"
  id="VUID-VkVideoEncodeH264RateControlInfoKHR-flags-08281"></a>
  VUID-VkVideoEncodeH264RateControlInfoKHR-flags-08281  
  If `flags` contains
  `VK_VIDEO_ENCODE_H264_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR` or
  `VK_VIDEO_ENCODE_H264_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR`,
  then it **must** also contain
  `VK_VIDEO_ENCODE_H264_RATE_CONTROL_REGULAR_GOP_BIT_KHR`

- <a href="#VUID-VkVideoEncodeH264RateControlInfoKHR-flags-08282"
  id="VUID-VkVideoEncodeH264RateControlInfoKHR-flags-08282"></a>
  VUID-VkVideoEncodeH264RateControlInfoKHR-flags-08282  
  If `flags` contains
  `VK_VIDEO_ENCODE_H264_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR`,
  then it **must** not also contain
  `VK_VIDEO_ENCODE_H264_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR`

- <a href="#VUID-VkVideoEncodeH264RateControlInfoKHR-flags-08283"
  id="VUID-VkVideoEncodeH264RateControlInfoKHR-flags-08283"></a>
  VUID-VkVideoEncodeH264RateControlInfoKHR-flags-08283  
  If `flags` contains
  `VK_VIDEO_ENCODE_H264_RATE_CONTROL_REGULAR_GOP_BIT_KHR`, then
  `gopFrameCount` **must** be greater than `0`

- <a href="#VUID-VkVideoEncodeH264RateControlInfoKHR-idrPeriod-08284"
  id="VUID-VkVideoEncodeH264RateControlInfoKHR-idrPeriod-08284"></a>
  VUID-VkVideoEncodeH264RateControlInfoKHR-idrPeriod-08284  
  If `idrPeriod` is not `0`, then it **must** be greater than or equal
  to `gopFrameCount`

- <a
  href="#VUID-VkVideoEncodeH264RateControlInfoKHR-consecutiveBFrameCount-08285"
  id="VUID-VkVideoEncodeH264RateControlInfoKHR-consecutiveBFrameCount-08285"></a>
  VUID-VkVideoEncodeH264RateControlInfoKHR-consecutiveBFrameCount-08285  
  If `consecutiveBFrameCount` is not `0`, then it **must** be less than
  `gopFrameCount`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH264RateControlInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH264RateControlInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH264RateControlInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_RATE_CONTROL_INFO_KHR`

- <a href="#VUID-VkVideoEncodeH264RateControlInfoKHR-flags-parameter"
  id="VUID-VkVideoEncodeH264RateControlInfoKHR-flags-parameter"></a>
  VUID-VkVideoEncodeH264RateControlInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkVideoEncodeH264RateControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlFlagBitsKHR.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h264.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeH264RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH264RateControlInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
