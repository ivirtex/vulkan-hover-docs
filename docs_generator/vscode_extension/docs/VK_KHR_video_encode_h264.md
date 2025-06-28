# VK\_KHR\_video\_encode\_h264(3) Manual Page

## Name

VK\_KHR\_video\_encode\_h264 - device extension



## [](#_registered_extension_number)Registered Extension Number

39

## [](#_revision)Revision

14

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html)

## [](#_contact)Contact

- Ahmed Abdelkhalek [\[GitHub\]aabdelkh](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_encode_h264%5D%20%40aabdelkh%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_encode_h264%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_video\_encode\_h264](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_encode_h264.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

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
- Aidan Fabius, Core Avionics &amp; Industrial Inc.
- Lynne Iribarren, Independent

## [](#_description)Description

This extension builds upon the `VK_KHR_video_encode_queue` extension by adding support for encoding elementary video stream sequences compliant with the H.264/AVC video compression standard.

Note

This extension was promoted to `KHR` from the provisional extension `VK_EXT_video_encode_h264`.

## [](#_new_structures)New Structures

- [VkVideoEncodeH264FrameSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264FrameSizeKHR.html)
- [VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)
- [VkVideoEncodeH264QpKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264QpKHR.html)
- Extending [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoBeginCodingInfoKHR.html):
  
  - [VkVideoEncodeH264GopRemainingFrameInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264GopRemainingFrameInfoKHR.html)
- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html):
  
  - [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264CapabilitiesKHR.html)
- Extending [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlInfoKHR.html), [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoBeginCodingInfoKHR.html):
  
  - [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlInfoKHR.html)
- Extending [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html):
  
  - [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264PictureInfoKHR.html)
- Extending [VkVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelPropertiesKHR.html):
  
  - [VkVideoEncodeH264QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264QualityLevelPropertiesKHR.html)
- Extending [VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlLayerInfoKHR.html):
  
  - [VkVideoEncodeH264RateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlLayerInfoKHR.html)
- Extending [VkVideoEncodeSessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionParametersFeedbackInfoKHR.html):
  
  - [VkVideoEncodeH264SessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersFeedbackInfoKHR.html)
- Extending [VkVideoEncodeSessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionParametersGetInfoKHR.html):
  
  - [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html)
- Extending [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html), [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html):
  
  - [VkVideoEncodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264ProfileInfoKHR.html)
- Extending [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html):
  
  - [VkVideoEncodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264DpbSlotInfoKHR.html)
- Extending [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html):
  
  - [VkVideoEncodeH264SessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionCreateInfoKHR.html)
- Extending [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html):
  
  - [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html)
- Extending [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersUpdateInfoKHR.html):
  
  - [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)

## [](#_new_enums)New Enums

- [VkVideoEncodeH264CapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264CapabilityFlagBitsKHR.html)
- [VkVideoEncodeH264RateControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlFlagBitsKHR.html)
- [VkVideoEncodeH264StdFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264StdFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkVideoEncodeH264CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264CapabilityFlagsKHR.html)
- [VkVideoEncodeH264RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlFlagsKHR.html)
- [VkVideoEncodeH264StdFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264StdFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_VIDEO_ENCODE_H264_EXTENSION_NAME`
- `VK_KHR_VIDEO_ENCODE_H264_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
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
- Extending [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodecOperationFlagBitsKHR.html):
  
  - `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`

## [](#_version_history)Version History

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
  
  - Rename `VkVideoEncodeH264CapabilitiesFlagsEXT` to `VkVideoEncodeH264CapabilityFlagsEXT` and `VkVideoEncodeH264CapabilitiesFlagsEXT` to `VkVideoEncodeH264CapabilityFlagsEXT`, following Vulkan naming conventions.
- Revision 3, 2021-12-08 (Ahmed Abdelkhalek)
  
  - Rate control updates
- Revision 4, 2022-02-04 (Ahmed Abdelkhalek)
  
  - Align VkVideoEncodeH264VclFrameInfoEXT structure to similar one in VK\_EXT\_video\_encode\_h265 extension
- Revision 5, 2022-02-10 (Ahmed Abdelkhalek)
  
  - Updates to encode capability interface
- Revision 6, 2022-03-16 (Ahmed Abdelkhalek)
  
  - Relocate Std header version reporting/requesting from this extension to VK\_KHR\_video\_queue extension.
  - Remove redundant maxPictureSizeInMbs from VkVideoEncodeH264SessionCreateInfoEXT.
  - Remove the now empty VkVideoEncodeH264SessionCreateInfoEXT.
- Revision 7, 2022-04-06 (Ahmed Abdelkhalek)
  
  - Add capability flag to report support to use B frame in L1 reference list.
  - Add capability flag to report support for disabling SPS direct\_8x8\_inference\_flag.
- Revision 8, 2022-07-18 (Daniel Rakos)
  
  - Replace `VkVideoEncodeH264RateControlStructureFlagBitsEXT` bit enum with `VkVideoEncodeH264RateControlStructureEXT` enum
  - Rename `VkVideoEncodeH264ProfileEXT` to `VkVideoEncodeH264ProfileInfoEXT`
  - Rename `VkVideoEncodeH264ReferenceListsEXT` to `VkVideoEncodeH264ReferenceListsInfoEXT`
  - Rename `VkVideoEncodeH264EmitPictureParametersEXT` to `VkVideoEncodeH264EmitPictureParametersInfoEXT`
  - Rename `VkVideoEncodeH264NaluSliceEXT` to `VkVideoEncodeH264NaluSliceInfoEXT`
- Revision 9, 2022-09-18 (Daniel Rakos)
  
  - Rename `spsStdCount`, `pSpsStd`, `ppsStdCount`, and `pPpsStd` to `stdSPSCount`, `pStdSPSs`, `stdPPSCount`, and `pStdPPSs`, respectively, in `VkVideoEncodeH264SessionParametersAddInfoEXT`
  - Rename `maxSpsStdCount` and `maxPpsStdCount` to `maxStdSPSCount` and `maxStdPPSCount`, respectively, in `VkVideoEncodeH264SessionParametersCreateInfoEXT`
- Revision 10, 2023-03-06 (Daniel Rakos)
  
  - Removed `VkVideoEncodeH264EmitPictureParametersInfoEXT`
  - Changed member types in `VkVideoEncodeH264CapabilitiesEXT` and `VkVideoEncodeH264ReferenceListsInfoEXT` from `uint8_t` to `uint32_t`
  - Changed the type of `VkVideoEncodeH264RateControlInfoEXT::temporalLayerCount` and `VkVideoEncodeH264RateControlLayerInfoEXT::temporalLayerId` from `uint8_t` to `uint32_t`
  - Removed `VkVideoEncodeH264InputModeFlagsEXT` and `VkVideoEncodeH264OutputModeFlagsEXT` as we only support frame-in-frame-out mode for now
  - Rename `pCurrentPictureInfo` in `VkVideoEncodeH264VclFrameInfoEXT` to `pStdPictureInfo`
  - Rename `pSliceHeaderStd` in `VkVideoEncodeH264NaluSliceInfoEXT` to `pStdSliceHeader`
  - Rename `pReferenceFinalLists` in `VkVideoEncodeH264VclFrameInfoEXT` and `VkVideoEncodeH264NaluSliceInfoEXT` to `pStdReferenceFinalLists`
  - Removed the `slotIndex` member of `VkVideoEncodeH264DpbSlotInfoEXT` and changed it to be chained to `VkVideoReferenceSlotInfoKHR`
  - Replaced `VkVideoEncodeH264ReferenceListsInfoEXT` with the new Video Std header structure `StdVideoEncodeH264ReferenceLists` that also includes data previously part of the now removed `StdVideoEncodeH264RefMemMgmtCtrlOperations` structure
  - Added new capability flag `VK_VIDEO_ENCODE_H264_CAPABILITY_DIFFERENT_REFERENCE_FINAL_LISTS_BIT_EXT`
- Revision 11, 2023-05-22 (Daniel Rakos)
  
  - Renamed `VkVideoEncodeH264VclFrameInfoEXT` to `VkVideoEncodeH264PictureInfoEXT`
  - Added `VkVideoEncodeH264PictureInfoEXT::generatePrefixNalu` and `VK_VIDEO_ENCODE_H264_CAPABILITY_GENERATE_PREFIX_NALU_BIT_EXT` to enable the generation of H.264 prefix NALUs when supported by the implementation
  - Removed `VkVideoEncodeH264RateControlLayerInfoEXT::temporalLayerId`
  - Added `expectDyadicTemporalLayerPattern` capability
  - Added the `VkVideoEncodeH264SessionParametersGetInfoEXT` structure to identify the H.264 parameter sets to retrieve encoded parameter data for, and the `VkVideoEncodeH264SessionParametersFeedbackInfoEXT` structure to retrieve H.264 parameter set override information when using the new `vkGetEncodedVideoSessionParametersKHR` command
  - Added `VkVideoEncodeH264NaluSliceInfoEXT::constantQp` to specify per-slice constant QP when rate control mode is `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`
  - Added `VkVideoEncodeH264QualityLevelPropertiesEXT` for retrieving H.264 specific quality level recommendations
  - Replaced `VkVideoEncodeH264RateControlStructureEXT` enum with the flags type `VkVideoEncodeH264RateControlFlagsEXT` and bits defined in `VkVideoEncodeH264RateControlFlagBitsEXT` and added HRD compliance flag
  - Removed `useInitialRcQp` and `initialRcQp` members of `VkVideoEncodeH264RateControlLayerInfoEXT`
  - Added `prefersGopRemainingFrames` and `requiresGopRemainingFrames`, and the new `VkVideoEncodeH264GopRemainingFrameInfoEXT` structure to allow specifying remaining frames of each type in the rate control GOP
  - Added `maxTemporalLayers`, `maxQp`, and `minQp` capabilities
  - Added `maxLevelIdc` capability and new `VkVideoEncodeH264SessionCreateInfoEXT` structure to specify upper bounds on the H.264 level of the produced video bitstream
  - Moved capability flags specific to codec syntax restrictions from `VkVideoEncodeH264CapabilityFlagsEXT` to the new `VkVideoEncodeH264StdFlagsEXT` which is now included as a separate `stdSyntaxFlags` member in `VkVideoEncodeH264CapabilitiesEXT`
  - Removed codec syntax override values from `VkVideoEncodeH264CapabilitiesEXT`
  - Removed `VkVideoEncodeH264NaluSliceInfoEXT::mbCount` and `VK_VIDEO_ENCODE_H264_CAPABILITY_SLICE_MB_COUNT_BIT_EXT`
  - Replaced `VK_VIDEO_ENCODE_H264_CAPABILITY_MULTIPLE_SLICES_PER_FRAME_BIT_EXT` with the new `maxSliceCount` capability
  - Removed capability flag `VK_VIDEO_ENCODE_H264_CAPABILITY_DIFFERENT_REFERENCE_FINAL_LISTS_BIT_EXT` and removed `pStdReferenceFinalLists` members from the `VkVideoEncodeH264PictureInfoEXT` and `VkVideoEncodeH264NaluSliceInfoEXT` structures as reference lists info is now included in `pStdPictureInfo`
  - Added capability flag `VK_VIDEO_ENCODE_H264_CAPABILITY_B_FRAME_IN_L0_LIST_BIT_EXT`
- Revision 12, 2023-07-19 (Daniel Rakos)
  
  - Added video std capability flags `VK_VIDEO_ENCODE_H264_STD_SLICE_QP_DELTA_BIT_EXT` and `VK_VIDEO_ENCODE_H264_STD_DIFFERENT_SLICE_QP_DELTA_BIT_EXT`
  - Fixed optionality of the array members of `VkVideoEncodeH264SessionParametersAddInfoEXT`
  - Fixed optionality of `VkVideoEncodeH264RateControlInfoEXT::flags`
- Revision 13, 2023-09-04 (Daniel Rakos)
  
  - Change extension from `EXT` to `KHR`
  - Extension is no longer provisional
- Revision 14, 2023-12-05 (Daniel Rakos)
  
  - Condition reference picture setup based on the value of `StdVideoEncodeH264PictureInfo::flags.is_reference`

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_video_encode_h264)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0