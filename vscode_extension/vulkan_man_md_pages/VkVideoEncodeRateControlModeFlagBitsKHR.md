# VkVideoEncodeRateControlModeFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeRateControlModeFlagBitsKHR - Video encode rate control
modes



## <a href="#_c_specification" class="anchor"></a>C Specification

The rate control modes are defined with the following enums:

``` c
// Provided by VK_KHR_video_encode_queue
typedef enum VkVideoEncodeRateControlModeFlagBitsKHR {
    VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DEFAULT_KHR = 0,
    VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_RATE_CONTROL_MODE_CBR_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_RATE_CONTROL_MODE_VBR_BIT_KHR = 0x00000004,
} VkVideoEncodeRateControlModeFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DEFAULT_KHR` specifies the use of
  implementation-specific rate control.

- `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR` specifies that
  rate control is disabled and the application will specify
  per-operation rate control parameters controlling the encoding
  quality. In this mode implementations will encode pictures
  independently of the output bitrate of prior video encode operations.

  - When using an <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-profile"
    target="_blank" rel="noopener">H.264 encode profile</a>,
    implementations will use the QP value specified in
    [VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)::`constantQp`
    to control the quality of the encoded picture.

  - When using an <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-profile"
    target="_blank" rel="noopener">H.265 encode profile</a>,
    implementations will use the QP value specified in
    [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)::`constantQp`
    to control the quality of the encoded picture.

- `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_CBR_BIT_KHR` specifies the use of
  constant bitrate (CBR) rate control mode. In this mode the
  implementation will attempt to produce the encoded bitstream at a
  constant bitrate while conforming to the constraints of other rate
  control parameters.

- `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_VBR_BIT_KHR` specifies the use of
  variable bitrate (VBR) rate control mode. In this mode the
  implementation will produce the encoded bitstream at a variable
  bitrate according to the constraints of other rate control parameters.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelPropertiesKHR.html),
[VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html),
[VkVideoEncodeRateControlModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlModeFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeRateControlModeFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
