# VkVideoProfileInfoKHR(3) Manual Page

## Name

VkVideoProfileInfoKHR - Structure specifying a video profile



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoProfileInfoKHR` structure is defined as follows:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `videoCodecOperation` is a
  [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodecOperationFlagBitsKHR.html)
  value specifying a video codec operation.

- `chromaSubsampling` is a bitmask of
  [VkVideoChromaSubsamplingFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoChromaSubsamplingFlagBitsKHR.html)
  specifying video chroma subsampling information.

- `lumaBitDepth` is a bitmask of
  [VkVideoComponentBitDepthFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoComponentBitDepthFlagBitsKHR.html)
  specifying video luma bit depth information.

- `chromaBitDepth` is a bitmask of
  [VkVideoComponentBitDepthFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoComponentBitDepthFlagBitsKHR.html)
  specifying video chroma bit depth information.

## <a href="#_description" class="anchor"></a>Description

Video profiles are provided as input to video capability queries such as
[vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
or
[vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html),
as well as when creating resources to be used by video coding operations
such as images, buffers, query pools, and video sessions.

The full description of a video profile is specified by an instance of
this structure, and the codec-specific and auxiliary structures provided
in its `pNext` chain.

When this structure is specified as an input parameter to
[vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html),
or through the `pProfiles` member of a
[VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html) structure in
the `pNext` chain of the input parameter of a query command such as
[vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html)
or
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html),
the following error codes indicate specific causes of the failure of the
query operation:

- `VK_ERROR_VIDEO_PICTURE_LAYOUT_NOT_SUPPORTED_KHR` indicates that the
  requested video picture layout (e.g. through the `pictureLayout`
  member of a
  [VkVideoDecodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264ProfileInfoKHR.html)
  structure included in the `pNext` chain of `VkVideoProfileInfoKHR`) is
  not supported.

- `VK_ERROR_VIDEO_PROFILE_OPERATION_NOT_SUPPORTED_KHR` indicates that a
  video profile operation specified by `videoCodecOperation` is not
  supported.

- `VK_ERROR_VIDEO_PROFILE_FORMAT_NOT_SUPPORTED_KHR` indicates that video
  format parameters specified by `chromaSubsampling`, `lumaBitDepth`, or
  `chromaBitDepth` are not supported.

- `VK_ERROR_VIDEO_PROFILE_CODEC_NOT_SUPPORTED_KHR` indicates that the
  codec-specific parameters corresponding to the video codec operation
  are not supported.

Valid Usage

- <a href="#VUID-VkVideoProfileInfoKHR-chromaSubsampling-07013"
  id="VUID-VkVideoProfileInfoKHR-chromaSubsampling-07013"></a>
  VUID-VkVideoProfileInfoKHR-chromaSubsampling-07013  
  `chromaSubsampling` **must** have a single bit set

- <a href="#VUID-VkVideoProfileInfoKHR-lumaBitDepth-07014"
  id="VUID-VkVideoProfileInfoKHR-lumaBitDepth-07014"></a>
  VUID-VkVideoProfileInfoKHR-lumaBitDepth-07014  
  `lumaBitDepth` **must** have a single bit set

- <a href="#VUID-VkVideoProfileInfoKHR-chromaSubsampling-07015"
  id="VUID-VkVideoProfileInfoKHR-chromaSubsampling-07015"></a>
  VUID-VkVideoProfileInfoKHR-chromaSubsampling-07015  
  If `chromaSubsampling` is not
  `VK_VIDEO_CHROMA_SUBSAMPLING_MONOCHROME_BIT_KHR`, then
  `chromaBitDepth` **must** have a single bit set

- <a href="#VUID-VkVideoProfileInfoKHR-videoCodecOperation-07179"
  id="VUID-VkVideoProfileInfoKHR-videoCodecOperation-07179"></a>
  VUID-VkVideoProfileInfoKHR-videoCodecOperation-07179  
  If `videoCodecOperation` is
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the `pNext` chain
  **must** include a
  [VkVideoDecodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264ProfileInfoKHR.html)
  structure

- <a href="#VUID-VkVideoProfileInfoKHR-videoCodecOperation-07180"
  id="VUID-VkVideoProfileInfoKHR-videoCodecOperation-07180"></a>
  VUID-VkVideoProfileInfoKHR-videoCodecOperation-07180  
  If `videoCodecOperation` is
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the `pNext` chain
  **must** include a
  [VkVideoDecodeH265ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265ProfileInfoKHR.html)
  structure

- <a href="#VUID-VkVideoProfileInfoKHR-videoCodecOperation-09256"
  id="VUID-VkVideoProfileInfoKHR-videoCodecOperation-09256"></a>
  VUID-VkVideoProfileInfoKHR-videoCodecOperation-09256  
  If `videoCodecOperation` is
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then the `pNext` chain
  **must** include a
  [VkVideoDecodeAV1ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1ProfileInfoKHR.html)
  structure

- <a href="#VUID-VkVideoProfileInfoKHR-videoCodecOperation-07181"
  id="VUID-VkVideoProfileInfoKHR-videoCodecOperation-07181"></a>
  VUID-VkVideoProfileInfoKHR-videoCodecOperation-07181  
  If `videoCodecOperation` is
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the `pNext` chain
  **must** include a
  [VkVideoEncodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264ProfileInfoKHR.html)
  structure

- <a href="#VUID-VkVideoProfileInfoKHR-videoCodecOperation-07182"
  id="VUID-VkVideoProfileInfoKHR-videoCodecOperation-07182"></a>
  VUID-VkVideoProfileInfoKHR-videoCodecOperation-07182  
  If `videoCodecOperation` is
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the `pNext` chain
  **must** include a
  [VkVideoEncodeH265ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265ProfileInfoKHR.html)
  structure

Valid Usage (Implicit)

- <a href="#VUID-VkVideoProfileInfoKHR-sType-sType"
  id="VUID-VkVideoProfileInfoKHR-sType-sType"></a>
  VUID-VkVideoProfileInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_PROFILE_INFO_KHR`

- <a href="#VUID-VkVideoProfileInfoKHR-videoCodecOperation-parameter"
  id="VUID-VkVideoProfileInfoKHR-videoCodecOperation-parameter"></a>
  VUID-VkVideoProfileInfoKHR-videoCodecOperation-parameter  
  `videoCodecOperation` **must** be a valid
  [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodecOperationFlagBitsKHR.html)
  value

- <a href="#VUID-VkVideoProfileInfoKHR-chromaSubsampling-parameter"
  id="VUID-VkVideoProfileInfoKHR-chromaSubsampling-parameter"></a>
  VUID-VkVideoProfileInfoKHR-chromaSubsampling-parameter  
  `chromaSubsampling` **must** be a valid combination of
  [VkVideoChromaSubsamplingFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoChromaSubsamplingFlagBitsKHR.html)
  values

- <a href="#VUID-VkVideoProfileInfoKHR-chromaSubsampling-requiredbitmask"
  id="VUID-VkVideoProfileInfoKHR-chromaSubsampling-requiredbitmask"></a>
  VUID-VkVideoProfileInfoKHR-chromaSubsampling-requiredbitmask  
  `chromaSubsampling` **must** not be `0`

- <a href="#VUID-VkVideoProfileInfoKHR-lumaBitDepth-parameter"
  id="VUID-VkVideoProfileInfoKHR-lumaBitDepth-parameter"></a>
  VUID-VkVideoProfileInfoKHR-lumaBitDepth-parameter  
  `lumaBitDepth` **must** be a valid combination of
  [VkVideoComponentBitDepthFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoComponentBitDepthFlagBitsKHR.html)
  values

- <a href="#VUID-VkVideoProfileInfoKHR-lumaBitDepth-requiredbitmask"
  id="VUID-VkVideoProfileInfoKHR-lumaBitDepth-requiredbitmask"></a>
  VUID-VkVideoProfileInfoKHR-lumaBitDepth-requiredbitmask  
  `lumaBitDepth` **must** not be `0`

- <a href="#VUID-VkVideoProfileInfoKHR-chromaBitDepth-parameter"
  id="VUID-VkVideoProfileInfoKHR-chromaBitDepth-parameter"></a>
  VUID-VkVideoProfileInfoKHR-chromaBitDepth-parameter  
  `chromaBitDepth` **must** be a valid combination of
  [VkVideoComponentBitDepthFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoComponentBitDepthFlagBitsKHR.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoChromaSubsamplingFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoChromaSubsamplingFlagsKHR.html),
[VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodecOperationFlagBitsKHR.html),
[VkVideoComponentBitDepthFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoComponentBitDepthFlagsKHR.html),
[VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html),
[VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html),
[vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoProfileInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
