# VkVideoProfileInfoKHR(3) Manual Page

## Name

VkVideoProfileInfoKHR - Structure specifying a video profile



## [](#_c_specification)C Specification

The `VkVideoProfileInfoKHR` structure is defined as follows:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkVideoProfileInfoKHR {
    VkStructureType                     sType;
    const void*                         pNext;
    VkVideoCodecOperationFlagBitsKHR    videoCodecOperation;
    VkVideoChromaSubsamplingFlagsKHR    chromaSubsampling;
    VkVideoComponentBitDepthFlagsKHR    lumaBitDepth;
    VkVideoComponentBitDepthFlagsKHR    chromaBitDepth;
} VkVideoProfileInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `videoCodecOperation` is a [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodecOperationFlagBitsKHR.html) value specifying a video codec operation.
- `chromaSubsampling` is a bitmask of [VkVideoChromaSubsamplingFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoChromaSubsamplingFlagBitsKHR.html) specifying video chroma subsampling information.
- `lumaBitDepth` is a bitmask of [VkVideoComponentBitDepthFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoComponentBitDepthFlagBitsKHR.html) specifying video luma bit depth information.
- `chromaBitDepth` is a bitmask of [VkVideoComponentBitDepthFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoComponentBitDepthFlagBitsKHR.html) specifying video chroma bit depth information.

## [](#_description)Description

Video profiles are provided as input to video capability queries such as [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) or [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html), as well as when creating resources to be used by video coding operations such as images, buffers, query pools, and video sessions.

The full description of a video profile is specified by an instance of this structure, and the codec-specific and auxiliary structures provided in its `pNext` chain.

When this structure is specified as an input parameter to [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html), or through the `pProfiles` member of a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure in the `pNext` chain of the input parameter of a query command such as [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html) or [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html), the following error codes indicate specific causes of the failure of the query operation:

- `VK_ERROR_VIDEO_PICTURE_LAYOUT_NOT_SUPPORTED_KHR` specifies that the requested video picture layout (e.g. through the `pictureLayout` member of a [VkVideoDecodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264ProfileInfoKHR.html) structure included in the `pNext` chain of `VkVideoProfileInfoKHR`) is not supported.
- `VK_ERROR_VIDEO_PROFILE_OPERATION_NOT_SUPPORTED_KHR` specifies that a video profile operation specified by `videoCodecOperation` is not supported.
- `VK_ERROR_VIDEO_PROFILE_FORMAT_NOT_SUPPORTED_KHR` specifies that video format parameters specified by `chromaSubsampling`, `lumaBitDepth`, or `chromaBitDepth` are not supported.
- `VK_ERROR_VIDEO_PROFILE_CODEC_NOT_SUPPORTED_KHR` specifies that the codec-specific parameters corresponding to the video codec operation are not supported.

Valid Usage

- [](#VUID-VkVideoProfileInfoKHR-chromaSubsampling-07013)VUID-VkVideoProfileInfoKHR-chromaSubsampling-07013  
  `chromaSubsampling` **must** have a single bit set
- [](#VUID-VkVideoProfileInfoKHR-lumaBitDepth-07014)VUID-VkVideoProfileInfoKHR-lumaBitDepth-07014  
  `lumaBitDepth` **must** have a single bit set
- [](#VUID-VkVideoProfileInfoKHR-chromaSubsampling-07015)VUID-VkVideoProfileInfoKHR-chromaSubsampling-07015  
  If `chromaSubsampling` is not `VK_VIDEO_CHROMA_SUBSAMPLING_MONOCHROME_BIT_KHR`, then `chromaBitDepth` **must** have a single bit set
- [](#VUID-VkVideoProfileInfoKHR-videoCodecOperation-07179)VUID-VkVideoProfileInfoKHR-videoCodecOperation-07179  
  If `videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoDecodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264ProfileInfoKHR.html) structure
- [](#VUID-VkVideoProfileInfoKHR-videoCodecOperation-07180)VUID-VkVideoProfileInfoKHR-videoCodecOperation-07180  
  If `videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoDecodeH265ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265ProfileInfoKHR.html) structure
- [](#VUID-VkVideoProfileInfoKHR-videoCodecOperation-10791)VUID-VkVideoProfileInfoKHR-videoCodecOperation-10791  
  If `videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_DECODE_VP9_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoDecodeVP9ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeVP9ProfileInfoKHR.html) structure
- [](#VUID-VkVideoProfileInfoKHR-videoCodecOperation-09256)VUID-VkVideoProfileInfoKHR-videoCodecOperation-09256  
  If `videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoDecodeAV1ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeAV1ProfileInfoKHR.html) structure
- [](#VUID-VkVideoProfileInfoKHR-videoCodecOperation-07181)VUID-VkVideoProfileInfoKHR-videoCodecOperation-07181  
  If `videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoEncodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264ProfileInfoKHR.html) structure
- [](#VUID-VkVideoProfileInfoKHR-videoCodecOperation-07182)VUID-VkVideoProfileInfoKHR-videoCodecOperation-07182  
  If `videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoEncodeH265ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265ProfileInfoKHR.html) structure
- [](#VUID-VkVideoProfileInfoKHR-videoCodecOperation-10262)VUID-VkVideoProfileInfoKHR-videoCodecOperation-10262  
  If `videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoEncodeAV1ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1ProfileInfoKHR.html) structure

Valid Usage (Implicit)

- [](#VUID-VkVideoProfileInfoKHR-sType-sType)VUID-VkVideoProfileInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_PROFILE_INFO_KHR`
- [](#VUID-VkVideoProfileInfoKHR-videoCodecOperation-parameter)VUID-VkVideoProfileInfoKHR-videoCodecOperation-parameter  
  `videoCodecOperation` **must** be a valid [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodecOperationFlagBitsKHR.html) value
- [](#VUID-VkVideoProfileInfoKHR-chromaSubsampling-parameter)VUID-VkVideoProfileInfoKHR-chromaSubsampling-parameter  
  `chromaSubsampling` **must** be a valid combination of [VkVideoChromaSubsamplingFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoChromaSubsamplingFlagBitsKHR.html) values
- [](#VUID-VkVideoProfileInfoKHR-chromaSubsampling-requiredbitmask)VUID-VkVideoProfileInfoKHR-chromaSubsampling-requiredbitmask  
  `chromaSubsampling` **must** not be `0`
- [](#VUID-VkVideoProfileInfoKHR-lumaBitDepth-parameter)VUID-VkVideoProfileInfoKHR-lumaBitDepth-parameter  
  `lumaBitDepth` **must** be a valid combination of [VkVideoComponentBitDepthFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoComponentBitDepthFlagBitsKHR.html) values
- [](#VUID-VkVideoProfileInfoKHR-lumaBitDepth-requiredbitmask)VUID-VkVideoProfileInfoKHR-lumaBitDepth-requiredbitmask  
  `lumaBitDepth` **must** not be `0`
- [](#VUID-VkVideoProfileInfoKHR-chromaBitDepth-parameter)VUID-VkVideoProfileInfoKHR-chromaBitDepth-parameter  
  `chromaBitDepth` **must** be a valid combination of [VkVideoComponentBitDepthFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoComponentBitDepthFlagBitsKHR.html) values

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoChromaSubsamplingFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoChromaSubsamplingFlagsKHR.html), [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodecOperationFlagBitsKHR.html), [VkVideoComponentBitDepthFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoComponentBitDepthFlagsKHR.html), [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html), [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html), [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoProfileInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0