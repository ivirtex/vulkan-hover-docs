# VkVideoEncodeIntraRefreshModeFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeIntraRefreshModeFlagBitsKHR - Video encode intra refresh modes



## [](#_c_specification)C Specification

The intra refresh modes are defined with the following enums:

```c++
// Provided by VK_KHR_video_encode_intra_refresh
typedef enum VkVideoEncodeIntraRefreshModeFlagBitsKHR {
    VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_NONE_KHR = 0,
    VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_PER_PICTURE_PARTITION_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_BLOCK_BASED_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_BLOCK_ROW_BASED_BIT_KHR = 0x00000004,
    VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_BLOCK_COLUMN_BASED_BIT_KHR = 0x00000008,
} VkVideoEncodeIntraRefreshModeFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_NONE_KHR` specifies that intra refresh **must** not be used.
- `VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_PER_PICTURE_PARTITION_BIT_KHR` specifies the use of *per picture partition intra refresh*. In this mode each intra refresh region i corresponds to the encoded picture partition i.
- `VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_BLOCK_BASED_BIT_KHR` specifies the use of any *block-based intra refresh*. In this mode each intra refresh region encompasses a set of coding blocks, independent of encoded picture partitions but without any additional guarantees on the granularity at which the picture is split into intra refresh regions. When using this mode, the set of coding blocks comprising the intra refresh regions and the direction of intra refresh are implementation-defined.
- `VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_BLOCK_ROW_BASED_BIT_KHR` specifies the use of *block-row-based intra refresh*. This mode is a block-based intra refresh mode where each intra refresh region encompasses a set of coding block rows.
- `VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_BLOCK_COLUMN_BASED_BIT_KHR` specifies the use of *block-column-based intra refresh*. This mode is a block-based intra refresh mode where each intra refresh region encompasses a set of coding block columns.

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_intra\_refresh](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_intra_refresh.html), [VkVideoEncodeIntraRefreshModeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshModeFlagsKHR.html), [VkVideoEncodeSessionIntraRefreshCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionIntraRefreshCreateInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeIntraRefreshModeFlagBitsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0