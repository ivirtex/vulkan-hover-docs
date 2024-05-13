# VkVideoDecodeAV1SessionParametersCreateInfoKHR(3) Manual Page

## Name

VkVideoDecodeAV1SessionParametersCreateInfoKHR - Structure specifies AV1
decoder parameter set information



## <a href="#_c_specification" class="anchor"></a>C Specification

When a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-parameters"
target="_blank" rel="noopener">video session parameters</a> object is
created with the codec operation
`VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, the
[VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html)::`pNext`
chain **must** include a
`VkVideoDecodeAV1SessionParametersCreateInfoKHR` structure specifying
the contents of the object.

The `VkVideoDecodeAV1SessionParametersCreateInfoKHR` structure is
defined as:

``` c
// Provided by VK_KHR_video_decode_av1
typedef struct VkVideoDecodeAV1SessionParametersCreateInfoKHR {
    VkStructureType                     sType;
    const void*                         pNext;
    const StdVideoAV1SequenceHeader*    pStdSequenceHeader;
} VkVideoDecodeAV1SessionParametersCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pStdSequenceHeader` is a pointer to a `StdVideoAV1SequenceHeader`
  structure describing the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-sequence-header"
  target="_blank" rel="noopener">AV1 sequence header</a> entry to store
  in the created object.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>As AV1 video session parameters objects will only ever contain a
single AV1 sequence header, this has to be specified at object creation
time and such video session parameters objects cannot be updated using
the <a
href="vkUpdateVideoSessionParametersKHR.html">vkUpdateVideoSessionParametersKHR</a>
command. When a new AV1 sequence header is decoded from the input video
bitstream the application needs to create a new video session parameters
object to store it.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a
  href="#VUID-VkVideoDecodeAV1SessionParametersCreateInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeAV1SessionParametersCreateInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeAV1SessionParametersCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_SESSION_PARAMETERS_CREATE_INFO_KHR`

- <a
  href="#VUID-VkVideoDecodeAV1SessionParametersCreateInfoKHR-pStdSequenceHeader-parameter"
  id="VUID-VkVideoDecodeAV1SessionParametersCreateInfoKHR-pStdSequenceHeader-parameter"></a>
  VUID-VkVideoDecodeAV1SessionParametersCreateInfoKHR-pStdSequenceHeader-parameter  
  `pStdSequenceHeader` **must** be a valid pointer to a valid
  `StdVideoAV1SequenceHeader` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_av1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_av1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeAV1SessionParametersCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
