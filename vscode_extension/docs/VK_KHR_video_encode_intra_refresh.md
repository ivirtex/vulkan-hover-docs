# VK\_KHR\_video\_encode\_intra\_refresh(3) Manual Page

## Name

VK\_KHR\_video\_encode\_intra\_refresh - device extension



## [](#_registered_extension_number)Registered Extension Number

553

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html)

## [](#_api_interactions)API Interactions

- Interacts with VK\_KHR\_video\_encode\_av1
- Interacts with VK\_KHR\_video\_encode\_h264
- Interacts with VK\_KHR\_video\_encode\_h265

## [](#_contact)Contact

- Ahmed Abdelkhalek [\[GitHub\]aabdelkh](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_encode_intra_refresh%5D%20%40aabdelkh%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_encode_intra_refresh%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_video\_encode\_intra\_refresh](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_encode_intra_refresh.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-03-28

**IP Status**

No known IP claims.

**Contributors**

- Ahmed Abdelkhalek, AMD
- Benjamin Cheng, AMD
- Srinath Kumarapuram, NVIDIA
- Tony Zlatinski, NVIDIA
- Ping Liu, Intel
- Daniel Rakos, RasterGrid
- Lynne Iribarren, Independent

## [](#_description)Description

This extension builds upon the `VK_KHR_video_encode_queue` extension by enabling the application to perform intra refresh in video encode operations.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceVideoEncodeIntraRefreshFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoEncodeIntraRefreshFeaturesKHR.html)
- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html):
  
  - [VkVideoEncodeIntraRefreshCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshCapabilitiesKHR.html)
- Extending [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html):
  
  - [VkVideoEncodeIntraRefreshInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshInfoKHR.html)
- Extending [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html):
  
  - [VkVideoReferenceIntraRefreshInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceIntraRefreshInfoKHR.html)
- Extending [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html):
  
  - [VkVideoEncodeSessionIntraRefreshCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionIntraRefreshCreateInfoKHR.html)

## [](#_new_enums)New Enums

- [VkVideoEncodeIntraRefreshModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshModeFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkVideoEncodeIntraRefreshModeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshModeFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_VIDEO_ENCODE_INTRA_REFRESH_EXTENSION_NAME`
- `VK_KHR_VIDEO_ENCODE_INTRA_REFRESH_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_ENCODE_INTRA_REFRESH_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_INTRA_REFRESH_CAPABILITIES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_INTRA_REFRESH_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_SESSION_INTRA_REFRESH_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_REFERENCE_INTRA_REFRESH_INFO_KHR`
- Extending [VkVideoEncodeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeFlagBitsKHR.html):
  
  - `VK_VIDEO_ENCODE_INTRA_REFRESH_BIT_KHR`

If [VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html) is supported:

- Extending [VkVideoEncodeAV1CapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilityFlagBitsKHR.html):
  
  - `VK_VIDEO_ENCODE_AV1_CAPABILITY_COMPOUND_PREDICTION_INTRA_REFRESH_BIT_KHR`

If [VK\_KHR\_video\_encode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h264.html) is supported:

- Extending [VkVideoEncodeH264CapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264CapabilityFlagBitsKHR.html):
  
  - `VK_VIDEO_ENCODE_H264_CAPABILITY_B_PICTURE_INTRA_REFRESH_BIT_KHR`

If [VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html) is supported:

- Extending [VkVideoEncodeH265CapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilityFlagBitsKHR.html):
  
  - `VK_VIDEO_ENCODE_H265_CAPABILITY_B_PICTURE_INTRA_REFRESH_BIT_KHR`

## [](#_version_history)Version History

- Revision 1, 2025-03-28 (Daniel Rakos)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_video_encode_intra_refresh).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0