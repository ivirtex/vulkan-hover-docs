# VK_KHR_video_encode_h264(3) Manual Page

## Name

VK_KHR_video_encode_h264 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

39

## <a href="#_revision" class="anchor"></a>Revision

14

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Ahmed Abdelkhalek <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_encode_h264%5D%20@aabdelkh%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_encode_h264%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>aabdelkh</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_video_encode_h264](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_encode_h264.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-12-05

**IP Status**  
No known IP claims.

**Contributors**  
- Ahmed Abdelkhalek, AMD

- George Hao, AMD

- Jake Beju, AMD

- Peter Fang, AMD

- Ping Liu, Intel

- Srinath Kumarapuram, NVIDIA

- Tony Zlatinski, NVIDIA

- Ravi Chaudhary, NVIDIA

- Yang Liu, AMD

- Daniel Rakos, RasterGrid

- Aidan Fabius, Core Avionics & Industrial Inc.

- Lynne Iribarren, Independent

## <a href="#_description" class="anchor"></a>Description

This extension builds upon the
[`VK_KHR_video_encode_queue`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html) extension
by adding support for encoding elementary video stream sequences
compliant with the H.264/AVC video compression standard.

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
extension <code>VK_EXT_video_encode_h264</code>.</p></td>
</tr>
</tbody>
</table>

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkVideoEncodeH264FrameSizeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264FrameSizeKHR.html)

- [VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)

- [VkVideoEncodeH264QpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264QpKHR.html)

- Extending [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html):

  - [VkVideoEncodeH264GopRemainingFrameInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264GopRemainingFrameInfoKHR.html)

- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html):

  - [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)

- Extending
  [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html),
  [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html):

  - [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlInfoKHR.html)

- Extending [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html):

  - [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)

- Extending
  [VkVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelPropertiesKHR.html):

  - [VkVideoEncodeH264QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264QualityLevelPropertiesKHR.html)

- Extending
  [VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlLayerInfoKHR.html):

  - [VkVideoEncodeH264RateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlLayerInfoKHR.html)

- Extending
  [VkVideoEncodeSessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersFeedbackInfoKHR.html):

  - [VkVideoEncodeH264SessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersFeedbackInfoKHR.html)

- Extending
  [VkVideoEncodeSessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersGetInfoKHR.html):

  - [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html)

- Extending [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html),
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html):

  - [VkVideoEncodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264ProfileInfoKHR.html)

- Extending
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html):

  - [VkVideoEncodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264DpbSlotInfoKHR.html)

- Extending
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html):

  - [VkVideoEncodeH264SessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionCreateInfoKHR.html)

- Extending
  [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html):

  - [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html)

- Extending
  [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersUpdateInfoKHR.html):

  - [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkVideoEncodeH264CapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilityFlagBitsKHR.html)

- [VkVideoEncodeH264RateControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlFlagBitsKHR.html)

- [VkVideoEncodeH264StdFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264StdFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkVideoEncodeH264CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilityFlagsKHR.html)

- [VkVideoEncodeH264RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlFlagsKHR.html)

- [VkVideoEncodeH264StdFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264StdFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_VIDEO_ENCODE_H264_EXTENSION_NAME`

- `VK_KHR_VIDEO_ENCODE_H264_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_CAPABILITIES_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_DPB_SLOT_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_GOP_REMAINING_FRAME_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_NALU_SLICE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_PICTURE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_PROFILE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_QUALITY_LEVEL_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_RATE_CONTROL_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_RATE_CONTROL_LAYER_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_SESSION_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_SESSION_PARAMETERS_ADD_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_SESSION_PARAMETERS_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_SESSION_PARAMETERS_FEEDBACK_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_SESSION_PARAMETERS_GET_INFO_KHR`

- Extending
  [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodecOperationFlagBitsKHR.html):

  - `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 0, 2018-7-23 (Ahmed Abdelkhalek)

  - Initial draft

- Revision 0.5, 2020-02-13 (Tony Zlatinski)

  - General Spec cleanup

  - Added DPB structures

  - Change the VCL frame encode structure

  - Added a common Non-VCL Picture Paramarameters structure

- Revision 1, 2021-03-29 (Tony Zlatinski)

  - Spec and API updates

- Revision 2, August 1 2021 (Srinath Kumarapuram)

  - Rename `VkVideoEncodeH264CapabilitiesFlagsEXT` to
    `VkVideoEncodeH264CapabilityFlagsEXT` and
    `VkVideoEncodeH264CapabilitiesFlagsEXT` to
    `VkVideoEncodeH264CapabilityFlagsEXT`, following Vulkan naming
    conventions.

- Revision 3, 2021-12-08 (Ahmed Abdelkhalek)

  - Rate control updates

- Revision 4, 2022-02-04 (Ahmed Abdelkhalek)

  - Align VkVideoEncodeH264VclFrameInfoEXT structure to similar one in
    VK_EXT_video_encode_h265 extension

- Revision 5, 2022-02-10 (Ahmed Abdelkhalek)

  - Updates to encode capability interface

- Revision 6, 2022-03-16 (Ahmed Abdelkhalek)

  - Relocate Std header version reporting/requesting from this extension
    to VK_KHR_video_queue extension.

  - Remove redundant maxPictureSizeInMbs from
    VkVideoEncodeH264SessionCreateInfoEXT.

  - Remove the now empty VkVideoEncodeH264SessionCreateInfoEXT.

- Revision 7, 2022-04-06 (Ahmed Abdelkhalek)

  - Add capability flag to report support to use B frame in L1 reference
    list.

  - Add capability flag to report support for disabling SPS
    direct_8x8_inference_flag.

- Revision 8, 2022-07-18 (Daniel Rakos)

  - Replace `VkVideoEncodeH264RateControlStructureFlagBitsEXT` bit enum
    with `VkVideoEncodeH264RateControlStructureEXT` enum

  - Rename `VkVideoEncodeH264ProfileEXT` to
    `VkVideoEncodeH264ProfileInfoEXT`

  - Rename `VkVideoEncodeH264ReferenceListsEXT` to
    `VkVideoEncodeH264ReferenceListsInfoEXT`

  - Rename `VkVideoEncodeH264EmitPictureParametersEXT` to
    `VkVideoEncodeH264EmitPictureParametersInfoEXT`

  - Rename `VkVideoEncodeH264NaluSliceEXT` to
    `VkVideoEncodeH264NaluSliceInfoEXT`

- Revision 9, 2022-09-18 (Daniel Rakos)

  - Rename `spsStdCount`, `pSpsStd`, `ppsStdCount`, and `pPpsStd` to
    `stdSPSCount`, `pStdSPSs`, `stdPPSCount`, and `pStdPPSs`,
    respectively, in `VkVideoEncodeH264SessionParametersAddInfoEXT`

  - Rename `maxSpsStdCount` and `maxPpsStdCount` to `maxStdSPSCount` and
    `maxStdPPSCount`, respectively, in
    `VkVideoEncodeH264SessionParametersCreateInfoEXT`

- Revision 10, 2023-03-06 (Daniel Rakos)

  - Removed `VkVideoEncodeH264EmitPictureParametersInfoEXT`

  - Changed member types in `VkVideoEncodeH264CapabilitiesEXT` and
    `VkVideoEncodeH264ReferenceListsInfoEXT` from `uint8_t` to
    `uint32_t`

  - Changed the type of
    `VkVideoEncodeH264RateControlInfoEXT::temporalLayerCount` and
    `VkVideoEncodeH264RateControlLayerInfoEXT::temporalLayerId` from
    `uint8_t` to `uint32_t`

  - Removed `VkVideoEncodeH264InputModeFlagsEXT` and
    `VkVideoEncodeH264OutputModeFlagsEXT` as we only support
    frame-in-frame-out mode for now

  - Rename `pCurrentPictureInfo` in `VkVideoEncodeH264VclFrameInfoEXT`
    to `pStdPictureInfo`

  - Rename `pSliceHeaderStd` in `VkVideoEncodeH264NaluSliceInfoEXT` to
    `pStdSliceHeader`

  - Rename `pReferenceFinalLists` in `VkVideoEncodeH264VclFrameInfoEXT`
    and `VkVideoEncodeH264NaluSliceInfoEXT` to `pStdReferenceFinalLists`

  - Removed the `slotIndex` member of `VkVideoEncodeH264DpbSlotInfoEXT`
    and changed it to be chained to `VkVideoReferenceSlotInfoKHR`

  - Replaced `VkVideoEncodeH264ReferenceListsInfoEXT` with the new Video
    Std header structure `StdVideoEncodeH264ReferenceLists` that also
    includes data previously part of the now removed
    `StdVideoEncodeH264RefMemMgmtCtrlOperations` structure

  - Added new capability flag
    `VK_VIDEO_ENCODE_H264_CAPABILITY_DIFFERENT_REFERENCE_FINAL_LISTS_BIT_EXT`

- Revision 11, 2023-05-22 (Daniel Rakos)

  - Renamed `VkVideoEncodeH264VclFrameInfoEXT` to
    `VkVideoEncodeH264PictureInfoEXT`

  - Added `VkVideoEncodeH264PictureInfoEXT::generatePrefixNalu` and
    `VK_VIDEO_ENCODE_H264_CAPABILITY_GENERATE_PREFIX_NALU_BIT_EXT` to
    enable the generation of H.264 prefix NALUs when supported by the
    implementation

  - Removed `VkVideoEncodeH264RateControlLayerInfoEXT::temporalLayerId`

  - Added `expectDyadicTemporalLayerPattern` capability

  - Added the `VkVideoEncodeH264SessionParametersGetInfoEXT` structure
    to identify the H.264 parameter sets to retrieve encoded parameter
    data for, and the
    `VkVideoEncodeH264SessionParametersFeedbackInfoEXT` structure to
    retrieve H.264 parameter set override information when using the new
    `vkGetEncodedVideoSessionParametersKHR` command

  - Added `VkVideoEncodeH264NaluSliceInfoEXT::constantQp` to specify
    per-slice constant QP when rate control mode is
    `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`

  - Added `VkVideoEncodeH264QualityLevelPropertiesEXT` for retrieving
    H.264 specific quality level recommendations

  - Replaced `VkVideoEncodeH264RateControlStructureEXT` enum with the
    flags type `VkVideoEncodeH264RateControlFlagsEXT` and bits defined
    in `VkVideoEncodeH264RateControlFlagBitsEXT` and added HRD
    compliance flag

  - Removed `useInitialRcQp` and `initialRcQp` members of
    `VkVideoEncodeH264RateControlLayerInfoEXT`

  - Added `prefersGopRemainingFrames` and `requiresGopRemainingFrames`,
    and the new `VkVideoEncodeH264GopRemainingFrameInfoEXT` structure to
    allow specifying remaining frames of each type in the rate control
    GOP

  - Added `maxTemporalLayers`, `maxQp`, and `minQp` capabilities

  - Added `maxLevelIdc` capability and new
    `VkVideoEncodeH264SessionCreateInfoEXT` structure to specify upper
    bounds on the H.264 level of the produced video bitstream

  - Moved capability flags specific to codec syntax restrictions from
    `VkVideoEncodeH264CapabilityFlagsEXT` to the new
    `VkVideoEncodeH264StdFlagsEXT` which is now included as a separate
    `stdSyntaxFlags` member in `VkVideoEncodeH264CapabilitiesEXT`

  - Removed codec syntax override values from
    `VkVideoEncodeH264CapabilitiesEXT`

  - Removed `VkVideoEncodeH264NaluSliceInfoEXT::mbCount` and
    `VK_VIDEO_ENCODE_H264_CAPABILITY_SLICE_MB_COUNT_BIT_EXT`

  - Replaced
    `VK_VIDEO_ENCODE_H264_CAPABILITY_MULTIPLE_SLICES_PER_FRAME_BIT_EXT`
    with the new `maxSliceCount` capability

  - Removed capability flag
    `VK_VIDEO_ENCODE_H264_CAPABILITY_DIFFERENT_REFERENCE_FINAL_LISTS_BIT_EXT`
    and removed `pStdReferenceFinalLists` members from the
    `VkVideoEncodeH264PictureInfoEXT` and
    `VkVideoEncodeH264NaluSliceInfoEXT` structures as reference lists
    info is now included in `pStdPictureInfo`

  - Added capability flag
    `VK_VIDEO_ENCODE_H264_CAPABILITY_B_FRAME_IN_L0_LIST_BIT_EXT`

- Revision 12, 2023-07-19 (Daniel Rakos)

  - Added video std capability flags
    `VK_VIDEO_ENCODE_H264_STD_SLICE_QP_DELTA_BIT_EXT` and
    `VK_VIDEO_ENCODE_H264_STD_DIFFERENT_SLICE_QP_DELTA_BIT_EXT`

  - Fixed optionality of the array members of
    `VkVideoEncodeH264SessionParametersAddInfoEXT`

  - Fixed optionality of `VkVideoEncodeH264RateControlInfoEXT::flags`

- Revision 13, 2023-09-04 (Daniel Rakos)

  - Change extension from `EXT` to `KHR`

  - Extension is no longer provisional

- Revision 14, 2023-12-05 (Daniel Rakos)

  - Condition reference picture setup based on the value of
    `StdVideoEncodeH264PictureInfo::flags.is_reference`

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_video_encode_h264"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
