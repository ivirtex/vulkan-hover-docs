# VkVideoDecodeH264SessionParametersAddInfoKHR(3) Manual Page

## Name

VkVideoDecodeH264SessionParametersAddInfoKHR - Structure specifies H.264
decoder parameter set information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoDecodeH264SessionParametersAddInfoKHR` structure is defined
as:

``` c
// Provided by VK_KHR_video_decode_h264
typedef struct VkVideoDecodeH264SessionParametersAddInfoKHR {
    VkStructureType                            sType;
    const void*                                pNext;
    uint32_t                                   stdSPSCount;
    const StdVideoH264SequenceParameterSet*    pStdSPSs;
    uint32_t                                   stdPPSCount;
    const StdVideoH264PictureParameterSet*     pStdPPSs;
} VkVideoDecodeH264SessionParametersAddInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stdSPSCount` is the number of elements in the `pStdSPSs` array.

- `pStdSPSs` is a pointer to an array of
  `StdVideoH264SequenceParameterSet` structures describing the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-sps"
  target="_blank" rel="noopener">H.264 SPS</a> entries to add.

- `stdPPSCount` is the number of elements in the `pStdPPSs` array.

- `pStdPPSs` is a pointer to an array of
  `StdVideoH264PictureParameterSet` structures describing the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-pps"
  target="_blank" rel="noopener">H.264 PPS</a> entries to add.

## <a href="#_description" class="anchor"></a>Description

This structure **can** be specified in the following places:

- In the `pParametersAddInfo` member of the
  [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)
  structure specified in the `pNext` chain of
  [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html)
  used to create a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-parameters"
  target="_blank" rel="noopener">video session parameters</a> object. In
  this case, if the video codec operation the video session parameters
  object is created with is
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then it defines the
  set of initial parameters to add to the created object (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#creating-video-session-parameters"
  target="_blank" rel="noopener">Creating Video Session Parameters</a>).

- In the `pNext` chain of
  [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersUpdateInfoKHR.html).
  In this case, if the video codec operation the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-parameters"
  target="_blank" rel="noopener">video session parameters</a> object to
  be updated was created with is
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then it defines the
  set of parameters to add to it (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-parameters-update"
  target="_blank" rel="noopener">Updating Video Session Parameters</a>).

Valid Usage

- <a href="#VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-None-04825"
  id="VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-None-04825"></a>
  VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-None-04825  
  The `seq_parameter_set_id` member of each
  `StdVideoH264SequenceParameterSet` structure specified in the elements
  of `pStdSPSs` **must** be unique within `pStdSPSs`

- <a href="#VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-None-04826"
  id="VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-None-04826"></a>
  VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-None-04826  
  The pair constructed from the `seq_parameter_set_id` and
  `pic_parameter_set_id` members of each
  `StdVideoH264PictureParameterSet` structure specified in the elements
  of `pStdPPSs` **must** be unique within `pStdPPSs`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_SESSION_PARAMETERS_ADD_INFO_KHR`

- <a
  href="#VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-pStdSPSs-parameter"
  id="VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-pStdSPSs-parameter"></a>
  VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-pStdSPSs-parameter  
  If `stdSPSCount` is not `0`, `pStdSPSs` **must** be a valid pointer to
  an array of `stdSPSCount` `StdVideoH264SequenceParameterSet` values

- <a
  href="#VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-pStdPPSs-parameter"
  id="VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-pStdPPSs-parameter"></a>
  VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-pStdPPSs-parameter  
  If `stdPPSCount` is not `0`, `pStdPPSs` **must** be a valid pointer to
  an array of `stdPPSCount` `StdVideoH264PictureParameterSet` values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_h264.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeH264SessionParametersAddInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
