# VkVideoEncodeAV1PredictionModeKHR(3) Manual Page

## Name

VkVideoEncodeAV1PredictionModeKHR - AV1 encode prediction mode



## [](#_c_specification)C Specification

Possible AV1 encode prediction modes are as follows:

```c++
// Provided by VK_KHR_video_encode_av1
typedef enum VkVideoEncodeAV1PredictionModeKHR {
    VK_VIDEO_ENCODE_AV1_PREDICTION_MODE_INTRA_ONLY_KHR = 0,
    VK_VIDEO_ENCODE_AV1_PREDICTION_MODE_SINGLE_REFERENCE_KHR = 1,
    VK_VIDEO_ENCODE_AV1_PREDICTION_MODE_UNIDIRECTIONAL_COMPOUND_KHR = 2,
    VK_VIDEO_ENCODE_AV1_PREDICTION_MODE_BIDIRECTIONAL_COMPOUND_KHR = 3,
} VkVideoEncodeAV1PredictionModeKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_AV1_PREDICTION_MODE_INTRA_ONLY_KHR` specifies the use of *intra-only prediction mode*, used when encoding AV1 frames of type `STD_VIDEO_AV1_FRAME_TYPE_KEY` or `STD_VIDEO_AV1_FRAME_TYPE_INTRA_ONLY`.
- `VK_VIDEO_ENCODE_AV1_PREDICTION_MODE_SINGLE_REFERENCE_KHR` specifies the use of *single reference prediction mode*, used when encoding AV1 frames of type `STD_VIDEO_AV1_FRAME_TYPE_INTER` or `STD_VIDEO_AV1_FRAME_TYPE_SWITCH` with `reference_select`, as defined in section 6.8.23 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1), equal to 0. When using this prediction mode, the application **must** specify a reference picture for at least one [AV1 reference name](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-names) in [VkVideoEncodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1PictureInfoKHR.html)::`referenceNameSlotIndices` that is supported by the implementation, as reported in [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`singleReferenceNameMask`.
- `VK_VIDEO_ENCODE_AV1_PREDICTION_MODE_UNIDIRECTIONAL_COMPOUND_KHR` specifies the use of *unidirectional compound prediction mode*, used when encoding AV1 frames of type `STD_VIDEO_AV1_FRAME_TYPE_INTER` or `STD_VIDEO_AV1_FRAME_TYPE_SWITCH` with `reference_select`, as defined in section 6.8.23 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1), equal to 1, and both [reference names](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-names) used for prediction are from the same reference frame group, as defined in section 6.10.24 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1). When using this prediction mode, the application **must** specify a reference picture for at least two [AV1 reference names](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-names) in [VkVideoEncodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1PictureInfoKHR.html)::`referenceNameSlotIndices` that is supported by the implementation, as reported in [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`unidirectionalCompoundReferenceNameMask`, where those two reference names are one of the allowed pairs of reference names, as defined in section 5.11.25 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1), listed below:
  
  - `LAST_FRAME` and `LAST2_FRAME`,
  - `LAST_FRAME` and `LAST3_FRAME`,
  - `LAST_FRAME` and `GOLDEN_FRAME`, or
  - `BWDREF_FRAME` and `ALTREF_FRAME`.
- `VK_VIDEO_ENCODE_AV1_PREDICTION_MODE_BIDIRECTIONAL_COMPOUND_KHR` specifies the use of *bidirectional compound prediction mode*, used when encoding AV1 frames of type `STD_VIDEO_AV1_FRAME_TYPE_INTER` or `STD_VIDEO_AV1_FRAME_TYPE_SWITCH` with `reference_select`, as defined in section 6.8.23 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1), equal to 1, and the two [reference names](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-names) used for prediction are from different reference frame groups, as defined in section 6.10.24 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1). When using this prediction mode, the application **must** specify a reference picture for at least one [AV1 reference name](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-names) from each reference frame group in [VkVideoEncodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1PictureInfoKHR.html)::`referenceNameSlotIndices` that is supported by the implementation, as reported in [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`bidirectionalCompoundReferenceNameMask`.

The effective prediction mode used to encode individual AV1 mode info blocks **may** use simpler prediction modes than the one set by the application for the frame, as allowed by the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1), in particular:

- Frames encoded with single reference prediction mode **may** contain mode info blocks encoded with intra-only prediction mode.
- Frames encoded with unidirectional compound prediction mode **may** contain mode info blocks encoded with intra-only or single reference prediction mode.
- Frames encoded with bidirectional compound prediction mode **may** contain mode info blocks encoded with intra-only, single reference, or unidirectional compound prediction mode.

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkVideoEncodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1PictureInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1PredictionModeKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0