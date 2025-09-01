# VK\_KHR\_video\_decode\_h264(3) Manual Page

## Name

VK\_KHR\_video\_decode\_h264 - device extension



## [](#_registered_extension_number)Registered Extension Number

41

## [](#_revision)Revision

9

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_video\_decode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_queue.html)

## [](#_contact)Contact

- [peter.fang@amd.com](mailto:peter.fang@amd.com)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_video\_decode\_h264](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_decode_h264.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-12-05

**IP Status**

No known IP claims.

**Contributors**

- Ahmed Abdelkhalek, AMD
- Chunbo Chen, Intel
- HoHin Lau, AMD
- Jake Beju, AMD
- Peter Fang, AMD
- Ping Liu, Intel
- Srinath Kumarapuram, NVIDIA
- Tony Zlatinski, NVIDIA
- Daniel Rakos, RasterGrid

## [](#_description)Description

This extension builds upon the `VK_KHR_video_decode_queue` extension by adding support for decoding elementary video stream sequences compliant with the H.264/AVC video compression standard.

Note

This extension was promoted to `KHR` from the provisional extension `VK_EXT_video_decode_h264`.

## [](#_new_structures)New Structures

- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html):
  
  - [VkVideoDecodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264CapabilitiesKHR.html)
- Extending [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html):
  
  - [VkVideoDecodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264PictureInfoKHR.html)
- Extending [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html), [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html):
  
  - [VkVideoDecodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264ProfileInfoKHR.html)
- Extending [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html):
  
  - [VkVideoDecodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264DpbSlotInfoKHR.html)
- Extending [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html):
  
  - [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)
- Extending [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersUpdateInfoKHR.html):
  
  - [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)

## [](#_new_enums)New Enums

- [VkVideoDecodeH264PictureLayoutFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264PictureLayoutFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkVideoDecodeH264PictureLayoutFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264PictureLayoutFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_VIDEO_DECODE_H264_EXTENSION_NAME`
- `VK_KHR_VIDEO_DECODE_H264_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_CAPABILITIES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_DPB_SLOT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_PICTURE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_PROFILE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_SESSION_PARAMETERS_ADD_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_SESSION_PARAMETERS_CREATE_INFO_KHR`
- Extending [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodecOperationFlagBitsKHR.html):
  
  - `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`

## [](#_version_history)Version History

- Revision 1, 2018-6-11 (Peter Fang)
  
  - Initial draft
- Revision 2, March 29 2021 (Tony Zlatinski)
  
  - Spec and API Updates
- Revision 3, August 1 2021 (Srinath Kumarapuram)
  
  - Rename `VkVideoDecodeH264FieldLayoutFlagsEXT` to `VkVideoDecodeH264PictureLayoutFlagsEXT`, `VkVideoDecodeH264FieldLayoutFlagBitsEXT` to `VkVideoDecodeH264PictureLayoutFlagBitsEXT` (along with the names of enumerants it defines), and `VkVideoDecodeH264ProfileEXT.fieldLayout` to `VkVideoDecodeH264ProfileEXT.pictureLayout`, following Vulkan naming conventions.
- Revision 4, 2022-03-16 (Ahmed Abdelkhalek)
  
  - Relocate Std header version reporting/requesting from this extension to VK\_KHR\_video\_queue extension.
  - Remove the now empty VkVideoDecodeH264SessionCreateInfoEXT.
- Revision 5, 2022-03-31 (Ahmed Abdelkhalek)
  
  - Use type StdVideoH264Level for VkVideoDecodeH264Capabilities.maxLevel
- Revision 6, 2022-08-09 (Daniel Rakos)
  
  - Rename `VkVideoDecodeH264ProfileEXT` to `VkVideoDecodeH264ProfileInfoEXT`
  - Rename `VkVideoDecodeH264MvcEXT` to `VkVideoDecodeH264MvcInfoEXT`
- Revision 7, 2022-09-18 (Daniel Rakos)
  
  - Change type of `VkVideoDecodeH264ProfileInfoEXT::pictureLayout` to `VkVideoDecodeH264PictureLayoutFlagBitsEXT`
  - Remove MVC support and related `VkVideoDecodeH264MvcInfoEXT` structure
  - Rename `spsStdCount`, `pSpsStd`, `ppsStdCount`, and `pPpsStd` to `stdSPSCount`, `pStdSPSs`, `stdPPSCount`, and `pStdPPSs`, respectively, in `VkVideoDecodeH264SessionParametersAddInfoEXT`
  - Rename `maxSpsStdCount` and `maxPpsStdCount` to `maxStdSPSCount` and `maxStdPPSCount`, respectively, in `VkVideoDecodeH264SessionParametersCreateInfoEXT`
  - Rename `slicesCount` and `pSlicesDataOffsets` to `sliceCount` and `pSliceOffsets`, respectively, in `VkVideoDecodeH264PictureInfoEXT`
- Revision 8, 2022-09-29 (Daniel Rakos)
  
  - Change extension from `EXT` to `KHR`
  - Extension is no longer provisional
- Revision 9, 2023-12-05 (Daniel Rakos)
  
  - Condition reference picture setup based on the value of `StdVideoDecodeH264PictureInfo::flags.is_reference`

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_video_decode_h264).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0