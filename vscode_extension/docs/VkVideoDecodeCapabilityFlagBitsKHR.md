# VkVideoDecodeCapabilityFlagBitsKHR(3) Manual Page

## Name

VkVideoDecodeCapabilityFlagBitsKHR - Video decode capability flags



## [](#_c_specification)C Specification

Bits which **may** be set in [VkVideoDecodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeCapabilitiesKHR.html)::`flags`, indicating the decoding capabilities supported, are:

```c++
// Provided by VK_KHR_video_decode_queue
typedef enum VkVideoDecodeCapabilityFlagBitsKHR {
    VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_COINCIDE_BIT_KHR = 0x00000001,
    VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR = 0x00000002,
} VkVideoDecodeCapabilityFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_COINCIDE_BIT_KHR` specifies support for using the same video picture resource as the [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#reconstructed-picture) and [decode output picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-output-picture) in a video decode operation.
- `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR` specifies support for using distinct video picture resources as the [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#reconstructed-picture) and [decode output picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-output-picture) in a video decode operation.
  
  Note
  
  Some video profiles allow using distinct video picture resources as the reconstructed picture and decode output picture in specific video decode operations even when the video decode profile does not support `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR`. Even if the implementation only reports coincide, the decode output picture for [film grain](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-av1-film-grain) enabled frames must be a different video picture resource from the reconstructed picture because film grain is applied outside of the coding loop.

Implementations are only **required** to support one of `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_COINCIDE_BIT_KHR` and `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR`. Accordingly, applications **should** handle both cases to maximize portability.

Note

If both `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_COINCIDE_BIT_KHR` and `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR` are supported, an application can choose to create separate images for decode DPB and decode output. E.g. in cases when linear tiling is preferred (and supported) for the decode output picture and the DPB requires optimal tiling, this avoids the need for a separate copy at the expense of additional memory bandwidth requirements during decoding.

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_queue.html), [VkVideoDecodeCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeCapabilityFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeCapabilityFlagBitsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0