# VkVideoDecodeH265SessionParametersAddInfoKHR(3) Manual Page

## Name

VkVideoDecodeH265SessionParametersAddInfoKHR - Structure specifies H.265
decoder parameter set information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoDecodeH265SessionParametersAddInfoKHR` structure is defined
as:

``` c
// Provided by VK_KHR_video_decode_h265
typedef struct VkVideoDecodeH265SessionParametersAddInfoKHR {
    VkStructureType                            sType;
    const void*                                pNext;
    uint32_t                                   stdVPSCount;
    const StdVideoH265VideoParameterSet*       pStdVPSs;
    uint32_t                                   stdSPSCount;
    const StdVideoH265SequenceParameterSet*    pStdSPSs;
    uint32_t                                   stdPPSCount;
    const StdVideoH265PictureParameterSet*     pStdPPSs;
} VkVideoDecodeH265SessionParametersAddInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stdVPSCount` is the number of elements in the `pStdVPSs` array.

- `pStdVPSs` is a pointer to an array of `StdVideoH265VideoParameterSet`
  structures describing the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-vps"
  target="_blank" rel="noopener">H.265 VPS</a> entries to add.

- `stdSPSCount` is the number of elements in the `pStdSPSs` array.

- `pStdSPSs` is a pointer to an array of
  `StdVideoH265SequenceParameterSet` structures describing the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-sps"
  target="_blank" rel="noopener">H.265 SPS</a> entries to add.

- `stdPPSCount` is the number of elements in the `pStdPPSs` array.

- `pStdPPSs` is a pointer to an array of
  `StdVideoH265PictureParameterSet` structures describing the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-pps"
  target="_blank" rel="noopener">H.265 PPS</a> entries to add.

## <a href="#_description" class="anchor"></a>Description

This structure **can** be specified in the following places:

- In the `pParametersAddInfo` member of the
  [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)
  structure specified in the `pNext` chain of
  [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html)
  used to create a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-parameters"
  target="_blank" rel="noopener">video session parameters</a> object. In
  this case, if the video codec operation the video session parameters
  object is created with is
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then it defines the
  set of initial parameters to add to the created object (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#creating-video-session-parameters"
  target="_blank" rel="noopener">Creating Video Session Parameters</a>).

- In the `pNext` chain of
  [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersUpdateInfoKHR.html).
  In this case, if the video codec operation the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-parameters"
  target="_blank" rel="noopener">video session parameters</a> object to
  be updated was created with is
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then it defines the
  set of parameters to add to it (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-parameters-update"
  target="_blank" rel="noopener">Updating Video Session Parameters</a>).

Valid Usage

- <a href="#VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-None-04833"
  id="VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-None-04833"></a>
  VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-None-04833  
  The `vps_video_parameter_set_id` member of each
  `StdVideoH265VideoParameterSet` structure specified in the elements of
  `pStdVPSs` **must** be unique within `pStdVPSs`

- <a href="#VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-None-04834"
  id="VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-None-04834"></a>
  VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-None-04834  
  The pair constructed from the `sps_video_parameter_set_id` and
  `sps_seq_parameter_set_id` members of each
  `StdVideoH265SequenceParameterSet` structure specified in the elements
  of `pStdSPSs` **must** be unique within `pStdSPSs`

- <a href="#VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-None-04835"
  id="VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-None-04835"></a>
  VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-None-04835  
  The triplet constructed from the `sps_video_parameter_set_id`,
  `pps_seq_parameter_set_id`, and `pps_pic_parameter_set_id` members of
  each `StdVideoH265PictureParameterSet` structure specified in the
  elements of `pStdPPSs` **must** be unique within `pStdPPSs`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_SESSION_PARAMETERS_ADD_INFO_KHR`

- <a
  href="#VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-pStdVPSs-parameter"
  id="VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-pStdVPSs-parameter"></a>
  VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-pStdVPSs-parameter  
  If `stdVPSCount` is not `0`, `pStdVPSs` **must** be a valid pointer to
  an array of `stdVPSCount` `StdVideoH265VideoParameterSet` values

- <a
  href="#VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-pStdSPSs-parameter"
  id="VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-pStdSPSs-parameter"></a>
  VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-pStdSPSs-parameter  
  If `stdSPSCount` is not `0`, `pStdSPSs` **must** be a valid pointer to
  an array of `stdSPSCount` `StdVideoH265SequenceParameterSet` values

- <a
  href="#VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-pStdPPSs-parameter"
  id="VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-pStdPPSs-parameter"></a>
  VUID-VkVideoDecodeH265SessionParametersAddInfoKHR-pStdPPSs-parameter  
  If `stdPPSCount` is not `0`, `pStdPPSs` **must** be a valid pointer to
  an array of `stdPPSCount` `StdVideoH265PictureParameterSet` values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_h265.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeH265SessionParametersAddInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
