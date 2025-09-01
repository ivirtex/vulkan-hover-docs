# VK\_KHR\_video\_decode\_h265(3) Manual Page

## Name

VK\_KHR\_video\_decode\_h265 - device extension



## [](#_registered_extension_number)Registered Extension Number

188

## [](#_revision)Revision

8

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_video\_decode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_queue.html)

## [](#_contact)Contact

- [peter.fang@amd.com](mailto:peter.fang@amd.com)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_video\_decode\_h265](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_decode_h265.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

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

## [](#_description)Description

This extension builds upon the `VK_KHR_video_decode_queue` extension by adding support for decoding elementary video stream sequences compliant with the H.265/HEVC video compression standard.

Note

This extension was promoted to `KHR` from the provisional extension `VK_EXT_video_decode_h265`.

## [](#_new_structures)New Structures

- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html):
  
  - [VkVideoDecodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265CapabilitiesKHR.html)
- Extending [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html):
  
  - [VkVideoDecodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265PictureInfoKHR.html)
- Extending [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html), [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html):
  
  - [VkVideoDecodeH265ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265ProfileInfoKHR.html)
- Extending [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html):
  
  - [VkVideoDecodeH265DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265DpbSlotInfoKHR.html)
- Extending [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html):
  
  - [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)
- Extending [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersUpdateInfoKHR.html):
  
  - [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_VIDEO_DECODE_H265_EXTENSION_NAME`
- `VK_KHR_VIDEO_DECODE_H265_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_CAPABILITIES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_DPB_SLOT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_PICTURE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_PROFILE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_SESSION_PARAMETERS_ADD_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_SESSION_PARAMETERS_CREATE_INFO_KHR`
- Extending [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodecOperationFlagBitsKHR.html):
  
  - `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`

## [](#_version_history)Version History

- Revision 1, 2018-6-11 (Peter Fang)
  
  - Initial draft
- Revision 1.6, March 29 2021 (Tony Zlatinski)
  
  - Spec and API updates.
- Revision 2, 2022-03-16 (Ahmed Abdelkhalek)
  
  - Relocate Std header version reporting/requesting from this extension to VK\_KHR\_video\_queue extension.
  - Remove the now empty VkVideoDecodeH265SessionCreateInfoEXT.
- Revision 3, 2022-03-31 (Ahmed Abdelkhalek)
  
  - Use type StdVideoH265Level for VkVideoDecodeH265Capabilities.maxLevel
- Revision 4, 2022-08-09 (Daniel Rakos)
  
  - Rename `VkVideoDecodeH265ProfileEXT` to `VkVideoDecodeH265ProfileInfoEXT`
- Revision 5, 2022-09-18 (Daniel Rakos)
  
  - Rename `vpsStdCount`, `pVpsStd`, `spsStdCount`, `pSpsStd`, `ppsStdCount`, and `pPpsStd` to `stdVPSCount`, `pStdVPSs`, `stdSPSCount`, `pStdSPSs`, `stdPPSCount`, and `pStdPPSs`, respectively, in `VkVideoDecodeH265SessionParametersAddInfoEXT`
  - Rename `maxVpsStdCount`, `maxSpsStdCount`, and `maxPpsStdCount` to `maxStdVPSCount`, `maxStdSPSCount` and `maxStdPPSCount`, respectively, in `VkVideoDecodeH265SessionParametersCreateInfoEXT`
  - Rename `slicesCount` and `pSlicesDataOffsets` to `sliceCount` and `pSliceOffsets`, respectively, in `VkVideoDecodeH265PictureInfoEXT`
- Revision 6, 2022-11-14 (Daniel Rakos)
  
  - Rename `slice` to `sliceSegment` in the APIs for better clarity
- Revision 7, 2022-11-14 (Daniel Rakos)
  
  - Change extension from `EXT` to `KHR`
  - Extension is no longer provisional
- Revision 8, 2023-12-05 (Daniel Rakos)
  
  - Condition reference picture setup based on the value of `StdVideoDecodeH265PictureInfo::flags.IsReference`

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_video_decode_h265).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0