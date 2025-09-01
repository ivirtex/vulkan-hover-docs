# VkVideoEncodeH264NaluSliceInfoKHR(3) Manual Page

## Name

VkVideoEncodeH264NaluSliceInfoKHR - Structure specifies H.264 encode slice NALU parameters



## [](#_c_specification)C Specification

The [VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264NaluSliceInfoKHR {
    VkStructureType                         sType;
    const void*                             pNext;
    int32_t                                 constantQp;
    const StdVideoEncodeH264SliceHeader*    pStdSliceHeader;
} VkVideoEncodeH264NaluSliceInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `constantQp` is the QP to use for the slice if the current [rate control mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control-modes) configured for the video session is `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`.
- `pStdSliceHeader` is a pointer to a `StdVideoEncodeH264SliceHeader` structure specifying [H.264 slice header parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-slice-header-params) for the slice.

## [](#_description)Description

Std Slice Header Parameters

The members of the `StdVideoEncodeH264SliceHeader` structure pointed to by `pStdSliceHeader` are interpreted as follows:

- `flags.reserved` and `reserved1` are used only for padding purposes and are otherwise ignored;
- if `pWeightTable` is not `NULL`, then it is a pointer to a `StdVideoEncodeH264WeightTable` that is interpreted as follows:
  
  - `flags.reserved` is used only for padding purposes and is otherwise ignored;
  - all other members of `StdVideoEncodeH264WeightTable` are interpreted as defined in section 7.4.3.2 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264);
- all other members are interpreted as defined in section 7.4.3 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264).

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH264NaluSliceInfoKHR-sType-sType)VUID-VkVideoEncodeH264NaluSliceInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_NALU_SLICE_INFO_KHR`
- [](#VUID-VkVideoEncodeH264NaluSliceInfoKHR-pNext-pNext)VUID-VkVideoEncodeH264NaluSliceInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkVideoEncodeH264NaluSliceInfoKHR-pStdSliceHeader-parameter)VUID-VkVideoEncodeH264NaluSliceInfoKHR-pStdSliceHeader-parameter  
  `pStdSliceHeader` **must** be a valid pointer to a valid `StdVideoEncodeH264SliceHeader` value

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h264.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264PictureInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH264NaluSliceInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0