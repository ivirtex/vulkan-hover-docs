# VkVideoEncodeH264SessionParametersCreateInfoKHR(3) Manual Page

## Name

VkVideoEncodeH264SessionParametersCreateInfoKHR - Structure specifies
H.264 encoder parameter set information



## <a href="#_c_specification" class="anchor"></a>C Specification

When a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-parameters"
target="_blank" rel="noopener">video session parameters</a> object is
created with the codec operation
`VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, the
[VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html)::`pNext`
chain **must** include a
`VkVideoEncodeH264SessionParametersCreateInfoKHR` structure specifying
the capacity and initial contents of the object.

The `VkVideoEncodeH264SessionParametersCreateInfoKHR` structure is
defined as:

``` c
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264SessionParametersCreateInfoKHR {
    VkStructureType                                        sType;
    const void*                                            pNext;
    uint32_t                                               maxStdSPSCount;
    uint32_t                                               maxStdPPSCount;
    const VkVideoEncodeH264SessionParametersAddInfoKHR*    pParametersAddInfo;
} VkVideoEncodeH264SessionParametersCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxStdSPSCount` is the maximum number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-sps"
  target="_blank" rel="noopener">H.264 SPS</a> entries the created
  `VkVideoSessionParametersKHR` **can** contain.

- `maxStdPPSCount` is the maximum number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-pps"
  target="_blank" rel="noopener">H.264 PPS</a> entries the created
  `VkVideoSessionParametersKHR` **can** contain.

- `pParametersAddInfo` is `NULL` or a pointer to a
  [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)
  structure specifying H.264 parameters to add upon object creation.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-VkVideoEncodeH264SessionParametersCreateInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH264SessionParametersCreateInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH264SessionParametersCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_SESSION_PARAMETERS_CREATE_INFO_KHR`

- <a
  href="#VUID-VkVideoEncodeH264SessionParametersCreateInfoKHR-pParametersAddInfo-parameter"
  id="VUID-VkVideoEncodeH264SessionParametersCreateInfoKHR-pParametersAddInfo-parameter"></a>
  VUID-VkVideoEncodeH264SessionParametersCreateInfoKHR-pParametersAddInfo-parameter  
  If `pParametersAddInfo` is not `NULL`, `pParametersAddInfo` **must**
  be a valid pointer to a valid
  [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h264.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH264SessionParametersCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
