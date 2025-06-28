# VkVideoEncodeCapabilitiesKHR(3) Manual Page

## Name

VkVideoEncodeCapabilitiesKHR - Structure describing general video encode capabilities for a video profile



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) with `pVideoProfile->videoCodecOperation` specifying an encode operation, the [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilitiesKHR.html) structure **must** be included in the `pNext` chain of the [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html) structure to retrieve capabilities specific to video encoding.

The `VkVideoEncodeCapabilitiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeCapabilitiesKHR {
    VkStructureType                         sType;
    void*                                   pNext;
    VkVideoEncodeCapabilityFlagsKHR         flags;
    VkVideoEncodeRateControlModeFlagsKHR    rateControlModes;
    uint32_t                                maxRateControlLayers;
    uint64_t                                maxBitrate;
    uint32_t                                maxQualityLevels;
    VkExtent2D                              encodeInputPictureGranularity;
    VkVideoEncodeFeedbackFlagsKHR           supportedEncodeFeedbackFlags;
} VkVideoEncodeCapabilitiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkVideoEncodeCapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilityFlagBitsKHR.html) describing supported encoding features.
- `rateControlModes` is a bitmask of [VkVideoEncodeRateControlModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlModeFlagBitsKHR.html) indicating supported [rate control modes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control-modes).
- `maxRateControlLayers` indicates the maximum number of [rate control layers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control-layers) supported.
- `maxBitrate` indicates the maximum supported bitrate.
- `maxQualityLevels` indicates the number of discrete [video encode quality levels](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quality-level) supported. Implementations **must** support at least one quality level.
- `encodeInputPictureGranularity` indicates the granularity at which [encode input picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-input-picture) data is encoded and **may** indicate a texel granularity up to the size of the largest supported codec-specific coding block. This capability does not impose any valid usage constraints on the application, however, depending on the contents of the encode input picture, it **may** have effects on the encoded bitstream, as described in more detail below.
- `supportedEncodeFeedbackFlags` is a bitmask of [VkVideoEncodeFeedbackFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeFeedbackFlagBitsKHR.html) values specifying the supported flags for [video encode feedback queries](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-video-encode-feedback).

## [](#_description)Description

Implementations **must** include support for at least `VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_BUFFER_OFFSET_BIT_KHR` and `VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_BYTES_WRITTEN_BIT_KHR` in `supportedEncodeFeedbackFlags`.

`encodeInputPictureGranularity` provides information about the way [encode input picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-input-picture) data is used as input to video encode operations. In particular, some implementations **may** not be able to limit the set of texels used to encode the output video bitstream to the image subregion specified in the [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoPictureResourceInfoKHR.html) structure corresponding to the encode input picture (i.e. to the resolution of the image data to encode specified in its `codedExtent` member).

Note

For example, the application requests the coded extent to be 1920x1080, but the implementation is only able to source the encode input picture data at the granularity of the codec-specific coding block size which is 16x16 pixels (or as otherwise indicated in `encodeInputPictureGranularity`). In this example, the content is horizontally aligned with the coding block size, but not vertically aligned with it. Thus encoding of the last row of coding blocks will be impacted by the contents of the input image at texel rows 1080 to 1087 (the latter being the next row which is vertically aligned with the coding block size, assuming a zero-based texel row index).

If `codedExtent` rounded up to the next integer multiple of `encodeInputPictureGranularity` is greater than the extent of the image subresource specified for the [encode input picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-input-picture), then the texel values corresponding to texel coordinates outside of the bounds of the image subresource **may** be undefined. However, implementations **should** use well-defined default values for such texels in order to maximize the encoding efficiency for the last coding block row/column, and/or to ensure consistent encoding results across repeated encoding of the same input content. Nonetheless, the values used for such texels **must** not have an effect on whether the video encode operation produces a compliant bitstream, and **must** not have any other effects on the encoded picture data beyond what **may** otherwise result from using these texel values as input to any compression algorithm, as defined in the used video compression standard.

Note

While not required, it is generally a good practice for applications to make sure that the image subresource used for the encode input picture has an extent that is an integer multiple of the codec-specific coding block size (or at least `encodeInputPictureGranularity`) and that this padding area is filled with known values in order to improve encoding efficiency, portability, and reproducibility.

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeCapabilitiesKHR-sType-sType)VUID-VkVideoEncodeCapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_CAPABILITIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilityFlagsKHR.html), [VkVideoEncodeFeedbackFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeFeedbackFlagsKHR.html), [VkVideoEncodeRateControlModeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlModeFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeCapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0