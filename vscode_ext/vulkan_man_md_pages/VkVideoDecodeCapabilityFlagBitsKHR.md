# VkVideoDecodeCapabilityFlagBitsKHR(3) Manual Page

## Name

VkVideoDecodeCapabilityFlagBitsKHR - Video decode capability flags



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkVideoDecodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeCapabilitiesKHR.html)::`flags`,
indicating the decoding capabilities supported, are:

``` c
// Provided by VK_KHR_video_decode_queue
typedef enum VkVideoDecodeCapabilityFlagBitsKHR {
    VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_COINCIDE_BIT_KHR = 0x00000001,
    VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR = 0x00000002,
} VkVideoDecodeCapabilityFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_COINCIDE_BIT_KHR` indicates
  support for using the same video picture resource as the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reconstructed-picture"
  target="_blank" rel="noopener">reconstructed picture</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-output-picture"
  target="_blank" rel="noopener">decode output picture</a> in a video
  decode operation.

- `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR` indicates
  support for using distinct video picture resources as the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reconstructed-picture"
  target="_blank" rel="noopener">reconstructed picture</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-output-picture"
  target="_blank" rel="noopener">decode output picture</a> in a video
  decode operation.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr class="odd">
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>Some video profiles allow using distinct video picture resources as
  the reconstructed picture and decode output picture in specific video
  decode operations even when the video decode profile does not support
  <code>VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR</code>.
  Even if the implementation only reports coincide, the decode output
  picture for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-film-grain"
  target="_blank" rel="noopener">film grain</a> enabled frames must be a
  different video picture resource from the reconstructed picture because
  film grain is applied outside of the coding loop.</p></td>
  </tr>
  </tbody>
  </table>

Implementations are only **required** to support one of
`VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_COINCIDE_BIT_KHR` and
`VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR`.
Accordingly, applications **should** handle both cases to maximize
portability.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>If both
<code>VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_COINCIDE_BIT_KHR</code>
and
<code>VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR</code>
are supported, an application can choose to create separate images for
decode DPB and decode output. E.g. in cases when linear tiling is
preferred (and supported) for the decode output picture and the DPB
requires optimal tiling, this avoids the need for a separate copy at the
expense of additional memory bandwidth requirements during
decoding.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_queue.html),
[VkVideoDecodeCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeCapabilityFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeCapabilityFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
