# VkVideoEncodeH265RateControlInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265RateControlInfoKHR - Structure describing H.265 stream
rate control parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeH265RateControlInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265RateControlInfoKHR {
    VkStructureType                         sType;
    const void*                             pNext;
    VkVideoEncodeH265RateControlFlagsKHR    flags;
    uint32_t                                gopFrameCount;
    uint32_t                                idrPeriod;
    uint32_t                                consecutiveBFrameCount;
    uint32_t                                subLayerCount;
} VkVideoEncodeH265RateControlInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkVideoEncodeH265RateControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlFlagBitsKHR.html)
  specifying H.265 rate control flags.

- `gopFrameCount` is the number of frames within a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-gop"
  target="_blank" rel="noopener">group of pictures (GOP)</a> intended to
  be used by the application. If it is set to 0, the rate control
  algorithm **may** assume an implementation-dependent GOP length. If it
  is set to `UINT32_MAX`, the GOP length is treated as infinite.

- `idrPeriod` is the interval, in terms of number of frames, between two
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-idr-pic"
  target="_blank" rel="noopener">IDR frames</a> (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-idr-period"
  target="_blank" rel="noopener">IDR period</a>). If it is set to 0, the
  rate control algorithm **may** assume an implementation-dependent IDR
  period. If it is set to `UINT32_MAX`, the IDR period is treated as
  infinite.

- `consecutiveBFrameCount` is the number of consecutive B frames between
  I and/or P frames within the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-gop"
  target="_blank" rel="noopener">GOP</a>.

- `temporalLayerCount` specifies the number of H.265 sub-layers that the
  application intends to use.

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
`VK_VIDEO_ENCODE_H265_RATE_CONTROL_ATTEMPT_HRD_COMPLIANCE_BIT_KHR`, then
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
<tr class="odd">
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

- <a href="#VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08291"
  id="VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08291"></a>
  VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08291  
  If
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile, does not include
  `VK_VIDEO_ENCODE_H265_CAPABILITY_HRD_COMPLIANCE_BIT_KHR`, then `flags`
  **must** not contain
  `VK_VIDEO_ENCODE_H265_RATE_CONTROL_ATTEMPT_HRD_COMPLIANCE_BIT_KHR`

- <a href="#VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08292"
  id="VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08292"></a>
  VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08292  
  If `flags` contains
  `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR` or
  `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR`,
  then it **must** also contain
  `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REGULAR_GOP_BIT_KHR`

- <a href="#VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08293"
  id="VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08293"></a>
  VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08293  
  If `flags` contains
  `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR`,
  then it **must** not also contain
  `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR`

- <a href="#VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08294"
  id="VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08294"></a>
  VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08294  
  If `flags` contains
  `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REGULAR_GOP_BIT_KHR`, then
  `gopFrameCount` **must** be greater than `0`

- <a href="#VUID-VkVideoEncodeH265RateControlInfoKHR-idrPeriod-08295"
  id="VUID-VkVideoEncodeH265RateControlInfoKHR-idrPeriod-08295"></a>
  VUID-VkVideoEncodeH265RateControlInfoKHR-idrPeriod-08295  
  If `idrPeriod` is not `0`, then it **must** be greater than or equal
  to `gopFrameCount`

- <a
  href="#VUID-VkVideoEncodeH265RateControlInfoKHR-consecutiveBFrameCount-08296"
  id="VUID-VkVideoEncodeH265RateControlInfoKHR-consecutiveBFrameCount-08296"></a>
  VUID-VkVideoEncodeH265RateControlInfoKHR-consecutiveBFrameCount-08296  
  If `consecutiveBFrameCount` is not `0`, then it **must** be less than
  `gopFrameCount`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH265RateControlInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH265RateControlInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH265RateControlInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_RATE_CONTROL_INFO_KHR`

- <a href="#VUID-VkVideoEncodeH265RateControlInfoKHR-flags-parameter"
  id="VUID-VkVideoEncodeH265RateControlInfoKHR-flags-parameter"></a>
  VUID-VkVideoEncodeH265RateControlInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkVideoEncodeH265RateControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlFlagBitsKHR.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeH265RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265RateControlInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
