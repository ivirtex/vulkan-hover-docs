# VkVideoEncodeH264NaluSliceInfoKHR(3) Manual Page

## Name

VkVideoEncodeH264NaluSliceInfoKHR - Structure specifies H.264 encode
slice NALU parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)
structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264NaluSliceInfoKHR {
    VkStructureType                         sType;
    const void*                             pNext;
    int32_t                                 constantQp;
    const StdVideoEncodeH264SliceHeader*    pStdSliceHeader;
} VkVideoEncodeH264NaluSliceInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `constantQp` is the QP to use for the slice if the current <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
  target="_blank" rel="noopener">rate control mode</a> configured for
  the video session is
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`.

- `pStdSliceHeader` is a pointer to a `StdVideoEncodeH264SliceHeader`
  structure specifying <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-slice-header-params"
  target="_blank" rel="noopener">H.264 slice header parameters</a> for
  the slice.

## <a href="#_description" class="anchor"></a>Description

Std Slice Header Parameters  
The members of the `StdVideoEncodeH264SliceHeader` structure pointed to
by `pStdSliceHeader` are interpreted as follows:

- `flags.reserved` and `reserved1` are used only for padding purposes
  and are otherwise ignored;

- if `pWeightTable` is not `NULL`, then it is a pointer to a
  `StdVideoEncodeH264WeightTable` that is interpreted as follows:

  - `flags.reserved` is used only for padding purposes and is otherwise
    ignored;

  - all other members of `StdVideoEncodeH264WeightTable` are interpreted
    as defined in section 7.4.3.2 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
    target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- all other members are interpreted as defined in section 7.4.3 of the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>.

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH264NaluSliceInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH264NaluSliceInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH264NaluSliceInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_NALU_SLICE_INFO_KHR`

- <a href="#VUID-VkVideoEncodeH264NaluSliceInfoKHR-pNext-pNext"
  id="VUID-VkVideoEncodeH264NaluSliceInfoKHR-pNext-pNext"></a>
  VUID-VkVideoEncodeH264NaluSliceInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkVideoEncodeH264NaluSliceInfoKHR-pStdSliceHeader-parameter"
  id="VUID-VkVideoEncodeH264NaluSliceInfoKHR-pStdSliceHeader-parameter"></a>
  VUID-VkVideoEncodeH264NaluSliceInfoKHR-pStdSliceHeader-parameter  
  `pStdSliceHeader` **must** be a valid pointer to a valid
  `StdVideoEncodeH264SliceHeader` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h264.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH264NaluSliceInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
