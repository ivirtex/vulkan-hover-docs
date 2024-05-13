# VkVideoDecodeH265SessionParametersCreateInfoKHR(3) Manual Page

## Name

VkVideoDecodeH265SessionParametersCreateInfoKHR - Structure specifies
H.265 decoder parameter set information



## <a href="#_c_specification" class="anchor"></a>C Specification

When a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-parameters"
target="_blank" rel="noopener">video session parameters</a> object is
created with the codec operation
`VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, the
[VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html)::`pNext`
chain **must** include a
`VkVideoDecodeH265SessionParametersCreateInfoKHR` structure specifying
the capacity and initial contents of the object.

The `VkVideoDecodeH265SessionParametersCreateInfoKHR` structure is
defined as:

``` c
// Provided by VK_KHR_video_decode_h265
typedef struct VkVideoDecodeH265SessionParametersCreateInfoKHR {
    VkStructureType                                        sType;
    const void*                                            pNext;
    uint32_t                                               maxStdVPSCount;
    uint32_t                                               maxStdSPSCount;
    uint32_t                                               maxStdPPSCount;
    const VkVideoDecodeH265SessionParametersAddInfoKHR*    pParametersAddInfo;
} VkVideoDecodeH265SessionParametersCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxStdVPSCount` is the maximum number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-vps"
  target="_blank" rel="noopener">H.265 VPS</a> entries the created
  `VkVideoSessionParametersKHR` **can** contain.

- `maxStdSPSCount` is the maximum number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-sps"
  target="_blank" rel="noopener">H.265 SPS</a> entries the created
  `VkVideoSessionParametersKHR` **can** contain.

- `maxStdPPSCount` is the maximum number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-pps"
  target="_blank" rel="noopener">H.265 PPS</a> entries the created
  `VkVideoSessionParametersKHR` **can** contain.

- `pParametersAddInfo` is `NULL` or a pointer to a
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)
  structure specifying H.265 parameters to add upon object creation.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-VkVideoDecodeH265SessionParametersCreateInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeH265SessionParametersCreateInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeH265SessionParametersCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_SESSION_PARAMETERS_CREATE_INFO_KHR`

- <a
  href="#VUID-VkVideoDecodeH265SessionParametersCreateInfoKHR-pParametersAddInfo-parameter"
  id="VUID-VkVideoDecodeH265SessionParametersCreateInfoKHR-pParametersAddInfo-parameter"></a>
  VUID-VkVideoDecodeH265SessionParametersCreateInfoKHR-pParametersAddInfo-parameter  
  If `pParametersAddInfo` is not `NULL`, `pParametersAddInfo` **must**
  be a valid pointer to a valid
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_h265.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeH265SessionParametersCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
