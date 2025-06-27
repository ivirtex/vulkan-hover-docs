# VkVideoDecodeAV1ProfileInfoKHR(3) Manual Page

## Name

VkVideoDecodeAV1ProfileInfoKHR - Structure specifying AV1 decode profile



## <a href="#_c_specification" class="anchor"></a>C Specification

A video profile supporting AV1 video decode operations is specified by
setting
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)::`videoCodecOperation`
to `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR` and adding a
`VkVideoDecodeAV1ProfileInfoKHR` structure to the
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)::`pNext` chain.

The `VkVideoDecodeAV1ProfileInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_av1
typedef struct VkVideoDecodeAV1ProfileInfoKHR {
    VkStructureType       sType;
    const void*           pNext;
    StdVideoAV1Profile    stdProfile;
    VkBool32              filmGrainSupport;
} VkVideoDecodeAV1ProfileInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stdProfile` is a `StdVideoAV1Profile` value specifying the AV1 codec
  profile, as defined in section A.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>.

- <span id="decode-av1-film-grain-support"></span> `filmGrainSupport`
  specifies whether AV1 film grain, as defined in section 7.8.3 of the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>, **can** be used
  with the video profile. When this member is set to `VK_TRUE`, video
  session objects created against the video profile will be able to
  decode pictures that have <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-film-grain"
  target="_blank" rel="noopener">film grain</a> enabled.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Enabling <code>filmGrainSupport</code> <strong>may</strong> increase
the memory requirements of video sessions and/or video picture resources
on some implementations.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeAV1ProfileInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeAV1ProfileInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeAV1ProfileInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_PROFILE_INFO_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_av1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_av1.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeAV1ProfileInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
