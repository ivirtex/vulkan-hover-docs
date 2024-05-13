# VK_KHR_video_decode_h265(3) Manual Page

## Name

VK_KHR_video_decode_h265 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

188

## <a href="#_revision" class="anchor"></a>Revision

8

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_video_decode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_queue.html)  

## <a href="#_contact" class="anchor"></a>Contact

- <peter.fang@amd.com>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_video_decode_h265](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_decode_h265.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-12-05

**IP Status**  
No known IP claims.

**Contributors**  
- Ahmed Abdelkhalek, AMD

- HoHin Lau, AMD

- Jake Beju, AMD

- Peter Fang, AMD

- Ping Liu, Intel

- Srinath Kumarapuram, NVIDIA

- Tony Zlatinski, NVIDIA

- Daniel Rakos, RasterGrid

## <a href="#_description" class="anchor"></a>Description

This extension builds upon the
[`VK_KHR_video_decode_queue`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_queue.html) extension
by adding support for decoding elementary video stream sequences
compliant with the H.265/HEVC video compression standard.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>This extension was promoted to <code>KHR</code> from the provisional
extension <code>VK_EXT_video_decode_h265</code>.</p></td>
</tr>
</tbody>
</table>

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html):

  - [VkVideoDecodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265CapabilitiesKHR.html)

- Extending [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html):

  - [VkVideoDecodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265PictureInfoKHR.html)

- Extending [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html),
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html):

  - [VkVideoDecodeH265ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265ProfileInfoKHR.html)

- Extending
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html):

  - [VkVideoDecodeH265DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265DpbSlotInfoKHR.html)

- Extending
  [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html):

  - [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)

- Extending
  [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersUpdateInfoKHR.html):

  - [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_VIDEO_DECODE_H265_EXTENSION_NAME`

- `VK_KHR_VIDEO_DECODE_H265_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_CAPABILITIES_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_DPB_SLOT_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_PICTURE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_PROFILE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_SESSION_PARAMETERS_ADD_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_SESSION_PARAMETERS_CREATE_INFO_KHR`

- Extending
  [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodecOperationFlagBitsKHR.html):

  - `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-6-11 (Peter Fang)

  - Initial draft

- Revision 1.6, March 29 2021 (Tony Zlatinski)

  - Spec and API updates.

- Revision 2, 2022-03-16 (Ahmed Abdelkhalek)

  - Relocate Std header version reporting/requesting from this extension
    to VK_KHR_video_queue extension.

  - Remove the now empty VkVideoDecodeH265SessionCreateInfoEXT.

- Revision 3, 2022-03-31 (Ahmed Abdelkhalek)

  - Use type StdVideoH265Level for
    VkVideoDecodeH265Capabilities.maxLevel

- Revision 4, 2022-08-09 (Daniel Rakos)

  - Rename `VkVideoDecodeH265ProfileEXT` to
    `VkVideoDecodeH265ProfileInfoEXT`

- Revision 5, 2022-09-18 (Daniel Rakos)

  - Rename `vpsStdCount`, `pVpsStd`, `spsStdCount`, `pSpsStd`,
    `ppsStdCount`, and `pPpsStd` to `stdVPSCount`, `pStdVPSs`,
    `stdSPSCount`, `pStdSPSs`, `stdPPSCount`, and `pStdPPSs`,
    respectively, in `VkVideoDecodeH265SessionParametersAddInfoEXT`

  - Rename `maxVpsStdCount`, `maxSpsStdCount`, and `maxPpsStdCount` to
    `maxStdVPSCount`, `maxStdSPSCount` and `maxStdPPSCount`,
    respectively, in `VkVideoDecodeH265SessionParametersCreateInfoEXT`

  - Rename `slicesCount` and `pSlicesDataOffsets` to `sliceCount` and
    `pSliceOffsets`, respectively, in `VkVideoDecodeH265PictureInfoEXT`

- Revision 6, 2022-11-14 (Daniel Rakos)

  - Rename `slice` to `sliceSegment` in the APIs for better clarity

- Revision 7, 2022-11-14 (Daniel Rakos)

  - Change extension from `EXT` to `KHR`

  - Extension is no longer provisional

- Revision 8, 2023-12-05 (Daniel Rakos)

  - Condition reference picture setup based on the value of
    `StdVideoDecodeH265PictureInfo::flags.IsReference`

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_video_decode_h265"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
