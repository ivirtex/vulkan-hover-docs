# VK_KHR_video_encode_h265(3) Manual Page

## Name

VK_KHR_video_encode_h265 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

40

## <a href="#_revision" class="anchor"></a>Revision

14

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Ahmed Abdelkhalek <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_encode_h265%5D%20@aabdelkh%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_encode_h265%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>aabdelkh</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_video_encode_h265](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_encode_h265.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-12-05

**IP Status**  
No known IP claims.

**Contributors**  
- Ahmed Abdelkhalek, AMD

- George Hao, AMD

- Jake Beju, AMD

- Chunbo Chen, Intel

- Ping Liu, Intel

- Srinath Kumarapuram, NVIDIA

- Tony Zlatinski, NVIDIA

- Ravi Chaudhary, NVIDIA

- Daniel Rakos, RasterGrid

- Aidan Fabius, Core Avionics & Industrial Inc.

- Lynne Iribarren, Independent

## <a href="#_description" class="anchor"></a>Description

This extension builds upon the
[`VK_KHR_video_encode_queue`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html) extension
by adding support for encoding elementary video stream sequences
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
extension <code>VK_EXT_video_encode_h265</code>.</p></td>
</tr>
</tbody>
</table>

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkVideoEncodeH265FrameSizeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265FrameSizeKHR.html)

- [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)

- [VkVideoEncodeH265QpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265QpKHR.html)

- Extending [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html):

  - [VkVideoEncodeH265GopRemainingFrameInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265GopRemainingFrameInfoKHR.html)

- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html):

  - [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)

- Extending
  [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html),
  [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html):

  - [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlInfoKHR.html)

- Extending [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html):

  - [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)

- Extending
  [VkVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelPropertiesKHR.html):

  - [VkVideoEncodeH265QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265QualityLevelPropertiesKHR.html)

- Extending
  [VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlLayerInfoKHR.html):

  - [VkVideoEncodeH265RateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlLayerInfoKHR.html)

- Extending
  [VkVideoEncodeSessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersFeedbackInfoKHR.html):

  - [VkVideoEncodeH265SessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersFeedbackInfoKHR.html)

- Extending
  [VkVideoEncodeSessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersGetInfoKHR.html):

  - [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)

- Extending [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html),
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html):

  - [VkVideoEncodeH265ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265ProfileInfoKHR.html)

- Extending
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html):

  - [VkVideoEncodeH265DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265DpbSlotInfoKHR.html)

- Extending
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html):

  - [VkVideoEncodeH265SessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionCreateInfoKHR.html)

- Extending
  [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html):

  - [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)

- Extending
  [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersUpdateInfoKHR.html):

  - [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkVideoEncodeH265CapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilityFlagBitsKHR.html)

- [VkVideoEncodeH265CtbSizeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CtbSizeFlagBitsKHR.html)

- [VkVideoEncodeH265RateControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlFlagBitsKHR.html)

- [VkVideoEncodeH265StdFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265StdFlagBitsKHR.html)

- [VkVideoEncodeH265TransformBlockSizeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265TransformBlockSizeFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkVideoEncodeH265CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilityFlagsKHR.html)

- [VkVideoEncodeH265CtbSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CtbSizeFlagsKHR.html)

- [VkVideoEncodeH265RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlFlagsKHR.html)

- [VkVideoEncodeH265StdFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265StdFlagsKHR.html)

- [VkVideoEncodeH265TransformBlockSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265TransformBlockSizeFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_VIDEO_ENCODE_H265_EXTENSION_NAME`

- `VK_KHR_VIDEO_ENCODE_H265_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_CAPABILITIES_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_DPB_SLOT_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_GOP_REMAINING_FRAME_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_NALU_SLICE_SEGMENT_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_PICTURE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_PROFILE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_QUALITY_LEVEL_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_RATE_CONTROL_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_RATE_CONTROL_LAYER_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_SESSION_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_SESSION_PARAMETERS_ADD_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_SESSION_PARAMETERS_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_SESSION_PARAMETERS_FEEDBACK_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_SESSION_PARAMETERS_GET_INFO_KHR`

- Extending
  [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodecOperationFlagBitsKHR.html):

  - `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 0, 2019-11-14 (Ahmed Abdelkhalek)

  - Initial draft

- Revision 0.5, 2020-02-13 (Tony Zlatinski)

  - General Spec cleanup

  - Added DPB structures

  - Change the VCL frame encode structure

  - Added a common Non-VCL Picture Paramarameters structure

- Revision 2, Oct 10 2021 (Srinath Kumarapuram)

  - Vulkan Video Encode h.265 update and spec edits

- Revision 3, 2021-12-08 (Ahmed Abdelkhalek)

  - Rate control updates

- Revision 4, 2022-01-11 (Ahmed Abdelkhalek)

  - Replace occurrences of “slice” by “slice segment” and rename
    structures/enums to reflect this.

- Revision 5, 2022-02-10 (Ahmed Abdelkhalek)

  - Updates to encode capability interface

- Revision 6, 2022-03-16 (Ahmed Abdelkhalek)

  - Relocate Std header version reporting/requesting from this extension
    to VK_KHR_video_queue extension.

  - Remove the now empty VkVideoEncodeH265SessionCreateInfoEXT.

- Revision 7, 2022-03-24 (Ahmed Abdelkhalek)

  - Add capability flags to report support to disable transform skip and
    support to use B frame in L1 reference list.

- Revision 8, 2022-07-18 (Daniel Rakos)

  - Replace `VkVideoEncodeH265RateControlStructureFlagBitsEXT` bit enum
    with `VkVideoEncodeH265RateControlStructureEXT` enum

  - Rename `VkVideoEncodeH265ProfileEXT` to
    `VkVideoEncodeH265ProfileInfoEXT`

  - Rename `VkVideoEncodeH265ReferenceListsEXT` to
    `VkVideoEncodeH265ReferenceListsInfoEXT`

  - Rename `VkVideoEncodeH265EmitPictureParametersEXT` to
    `VkVideoEncodeH265EmitPictureParametersInfoEXT`

  - Rename `VkVideoEncodeH265NaluSliceSegmentEXT` to
    `VkVideoEncodeH265NaluSliceSegmentInfoEXT`

- Revision 9, 2022-09-18 (Daniel Rakos)

  - Rename `vpsStdCount`, `pVpsStd`, `spsStdCount`, `pSpsStd`,
    `ppsStdCount`, and `pPpsStd` to `stdVPSCount`, `pStdVPSs`,
    `stdSPSCount`, `pStdSPSs`, `stdPPSCount`, and `pStdPPSs`,
    respectively, in `VkVideoEncodeH265SessionParametersAddInfoEXT`

  - Rename `maxVpsStdCount`, `maxSpsStdCount`, and `maxPpsStdCount` to
    `maxStdVPSCount`, `maxStdSPSCount` and `maxStdPPSCount`,
    respectively, in `VkVideoEncodeH265SessionParametersCreateInfoEXT`

- Revision 10, 2023-03-06 (Daniel Rakos)

  - Removed `VkVideoEncodeH265EmitPictureParametersInfoEXT`

  - Changed member types in `VkVideoEncodeH265CapabilitiesEXT` and
    `VkVideoEncodeH265ReferenceListsInfoEXT` from `uint8_t` to
    `uint32_t`

  - Changed the type of
    `VkVideoEncodeH265RateControlInfoEXT::subLayerCount` and
    `VkVideoEncodeH265RateControlLayerInfoEXT::temporalId` from
    `uint8_t` to `uint32_t`

  - Removed `VkVideoEncodeH265InputModeFlagsEXT` and
    `VkVideoEncodeH265OutputModeFlagsEXT` as we only support
    frame-in-frame-out mode for now

  - Rename `pCurrentPictureInfo` in `VkVideoEncodeH265VclFrameInfoEXT`
    to `pStdPictureInfo`

  - Rename `pSliceSegmentHeaderStd` in
    `VkVideoEncodeH265NaluSliceSegmentInfoEXT` to
    `pStdSliceSegmentHeader`

  - Rename `pReferenceFinalLists` in `VkVideoEncodeH265VclFrameInfoEXT`
    and `VkVideoEncodeH265NaluSliceSegmentInfoEXT` to
    `pStdReferenceFinalLists`

  - Removed the `slotIndex` member of `VkVideoEncodeH265DpbSlotInfoEXT`
    and changed it to be chained to `VkVideoReferenceSlotInfoKHR`

  - Replaced `VkVideoEncodeH265ReferenceListsInfoEXT` with the new Video
    Std header structure `StdVideoEncodeH265ReferenceLists`

  - Added new capability flag
    `VK_VIDEO_ENCODE_H265_CAPABILITY_DIFFERENT_REFERENCE_FINAL_LISTS_BIT_EXT`

- Revision 11, 2023-05-26 (Daniel Rakos)

  - Renamed `VkVideoEncodeH265VclFrameInfoEXT` to
    `VkVideoEncodeH265PictureInfoEXT`

  - Removed `VkVideoEncodeH265RateControlLayerInfoEXT::temporalId`

  - Added `expectDyadicTemporalSubLayerPattern` capability

  - Added the `VkVideoEncodeH265SessionParametersGetInfoEXT` structure
    to identify the H.265 parameter sets to retrieve encoded parameter
    data for, and the
    `VkVideoEncodeH265SessionParametersFeedbackInfoEXT` structure to
    retrieve H.265 parameter set override information when using the new
    `vkGetEncodedVideoSessionParametersKHR` command

  - Added `VkVideoEncodeH265NaluSliceSegmentInfoEXT::constantQp` to
    specify per-slice segment constant QP when rate control mode is
    `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`

  - Added `VkVideoEncodeH265QualityLevelPropertiesEXT` for retrieving
    H.265 specific quality level recommendations

  - Replaced `VkVideoEncodeH265RateControlStructureEXT` enum with the
    flags type `VkVideoEncodeH265RateControlFlagsEXT` and bits defined
    in `VkVideoEncodeH265RateControlFlagBitsEXT` and added HRD
    compliance flag

  - Removed `useInitialRcQp` and `initialRcQp` members of
    `VkVideoEncodeH265RateControlLayerInfoEXT`

  - Added `prefersGopRemainingFrames` and `requiresGopRemainingFrames`,
    and the new `VkVideoEncodeH265GopRemainingFrameInfoEXT` structure to
    allow specifying remaining frames of each type in the rate control
    GOP

  - Renamed `maxSubLayersCount` capability to `maxSubLayerCount`

  - Added `maxQp`, and `minQp` capabilities

  - Added `maxLevelIdc` capability and new
    `VkVideoEncodeH265SessionCreateInfoEXT` structure to specify upper
    bounds on the H.265 level of the produced video bitstream

  - Moved capability flags specific to codec syntax restrictions from
    `VkVideoEncodeH265CapabilityFlagsEXT` to the new
    `VkVideoEncodeH265StdFlagsEXT` which is now included as a separate
    `stdSyntaxFlags` member in `VkVideoEncodeH265CapabilitiesEXT`

  - Added `std` prefix to codec syntax capabilities in
    `VkVideoEncodeH265CapabilitiesEXT`

  - Removed `VkVideoEncodeH265NaluSliceSegmentInfoEXT::ctbCount` and
    `VK_VIDEO_ENCODE_H265_CAPABILITY_SLICE_SEGMENT_CTB_COUNT_BIT_EXT`

  - Replaced
    `VK_VIDEO_ENCODE_H265_CAPABILITY_MULTIPLE_SLICE_SEGMENTS_PER_FRAME_BIT_EXT`
    with the new `maxSliceSegmentCount` capability

  - Added `maxTiles` capability

  - Removed codec syntax min/max capabilities from
    `VkVideoEncodeH265CapabilitiesEXT`

  - Removed capability flag
    `VK_VIDEO_ENCODE_H265_CAPABILITY_DIFFERENT_REFERENCE_FINAL_LISTS_BIT_EXT`
    and removed `pStdReferenceFinalLists` members from the
    `VkVideoEncodeH265PictureInfoEXT` and
    `VkVideoEncodeH265NaluSliceSegmentInfoEXT` structures as reference
    lists info is now included in `pStdPictureInfo`

  - Added capability flag
    `VK_VIDEO_ENCODE_H265_CAPABILITY_B_FRAME_IN_L0_LIST_BIT_EXT`

- Revision 12, 2023-07-19 (Daniel Rakos)

  - Added video std capability flags
    `VK_VIDEO_ENCODE_H265_STD_SLICE_QP_DELTA_BIT_EXT` and
    `VK_VIDEO_ENCODE_H265_STD_DIFFERENT_SLICE_QP_DELTA_BIT_EXT`

  - Fixed optionality of the array members of
    `VkVideoEncodeH265SessionParametersAddInfoEXT`

  - Fixed optionality of `VkVideoEncodeH265RateControlInfoEXT::flags`

- Revision 13, 2023-09-04 (Daniel Rakos)

  - Change extension from `EXT` to `KHR`

  - Extension is no longer provisional

- Revision 14, 2023-12-05 (Daniel Rakos)

  - Condition reference picture setup based on the value of
    `StdVideoEncodeH265PictureInfo::flags.is_reference`

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_video_encode_h265"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
