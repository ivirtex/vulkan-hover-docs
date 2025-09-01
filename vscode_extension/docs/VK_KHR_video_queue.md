# VK\_KHR\_video\_queue(3) Manual Page

## Name

VK\_KHR\_video\_queue - device extension



## [](#_registered_extension_number)Registered Extension Number

24

## [](#_revision)Revision

8

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [Vulkan Version 1.1](#versions-1.1)  
     and  
     [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_contact)Contact

- Tony Zlatinski [\[GitHub\]tzlatinski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_queue%5D%20%40tzlatinski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_queue%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_video\_queue](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_queue.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-09-29

**IP Status**

No known IP claims.

**Contributors**

- Ahmed Abdelkhalek, AMD
- George Hao, AMD
- Jake Beju, AMD
- Piers Daniell, NVIDIA
- Srinath Kumarapuram, NVIDIA
- Tobias Hector, AMD
- Tony Zlatinski, NVIDIA
- Daniel Rakos, RasterGrid

## [](#_description)Description

This extension provides common APIs to enable exposing queue families with support for video codec operations by introducing the following new object types and related functionalities:

- Video session objects that represent and maintain the state needed to perform video codec operations.
- Video session parameters objects that act as a container for codec specific parameters.

In addition, it also introduces query commands that allow applications to determine video coding related capabilities, and command buffer commands that enable recording video coding operations against a video session.

This extension is to be used in conjunction with other extensions that enable specific video coding operations.

## [](#_new_object_types)New Object Types

- [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionKHR.html)
- [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html)

## [](#_new_commands)New Commands

- [vkBindVideoSessionMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindVideoSessionMemoryKHR.html)
- [vkCmdBeginVideoCodingKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginVideoCodingKHR.html)
- [vkCmdControlVideoCodingKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdControlVideoCodingKHR.html)
- [vkCmdEndVideoCodingKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndVideoCodingKHR.html)
- [vkCreateVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateVideoSessionKHR.html)
- [vkCreateVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateVideoSessionParametersKHR.html)
- [vkDestroyVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyVideoSessionKHR.html)
- [vkDestroyVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyVideoSessionParametersKHR.html)
- [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
- [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html)
- [vkGetVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetVideoSessionMemoryRequirementsKHR.html)
- [vkUpdateVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateVideoSessionParametersKHR.html)

## [](#_new_structures)New Structures

- [VkBindVideoSessionMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindVideoSessionMemoryInfoKHR.html)
- [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)
- [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoBeginCodingInfoKHR.html)
- [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)
- [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlInfoKHR.html)
- [VkVideoEndCodingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEndCodingInfoKHR.html)
- [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html)
- [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoPictureResourceInfoKHR.html)
- [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html)
- [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html)
- [VkVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionMemoryRequirementsKHR.html)
- [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html)
- [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersUpdateInfoKHR.html)
- Extending [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html), [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html), [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html):
  
  - [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html)
- Extending [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html):
  
  - [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)
- Extending [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2.html):
  
  - [VkQueueFamilyQueryResultStatusPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyQueryResultStatusPropertiesKHR.html)
  - [VkQueueFamilyVideoPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyVideoPropertiesKHR.html)

## [](#_new_enums)New Enums

- [VkQueryResultStatusKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryResultStatusKHR.html)
- [VkVideoCapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilityFlagBitsKHR.html)
- [VkVideoChromaSubsamplingFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoChromaSubsamplingFlagBitsKHR.html)
- [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodecOperationFlagBitsKHR.html)
- [VkVideoCodingControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlFlagBitsKHR.html)
- [VkVideoComponentBitDepthFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoComponentBitDepthFlagBitsKHR.html)
- [VkVideoSessionCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkVideoBeginCodingFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoBeginCodingFlagsKHR.html)
- [VkVideoCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilityFlagsKHR.html)
- [VkVideoChromaSubsamplingFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoChromaSubsamplingFlagsKHR.html)
- [VkVideoCodecOperationFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodecOperationFlagsKHR.html)
- [VkVideoCodingControlFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlFlagsKHR.html)
- [VkVideoComponentBitDepthFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoComponentBitDepthFlagsKHR.html)
- [VkVideoEndCodingFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEndCodingFlagsKHR.html)
- [VkVideoSessionCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateFlagsKHR.html)
- [VkVideoSessionParametersCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_VIDEO_QUEUE_EXTENSION_NAME`
- `VK_KHR_VIDEO_QUEUE_SPEC_VERSION`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_VIDEO_SESSION_KHR`
  - `VK_OBJECT_TYPE_VIDEO_SESSION_PARAMETERS_KHR`
- Extending [VkQueryResultFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryResultFlagBits.html):
  
  - `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`
- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html):
  
  - `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`
  - `VK_ERROR_VIDEO_PICTURE_LAYOUT_NOT_SUPPORTED_KHR`
  - `VK_ERROR_VIDEO_PROFILE_CODEC_NOT_SUPPORTED_KHR`
  - `VK_ERROR_VIDEO_PROFILE_FORMAT_NOT_SUPPORTED_KHR`
  - `VK_ERROR_VIDEO_PROFILE_OPERATION_NOT_SUPPORTED_KHR`
  - `VK_ERROR_VIDEO_STD_VERSION_NOT_SUPPORTED_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BIND_VIDEO_SESSION_MEMORY_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_FORMAT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_QUERY_RESULT_STATUS_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_VIDEO_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_BEGIN_CODING_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_CAPABILITIES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_CODING_CONTROL_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_END_CODING_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_FORMAT_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_PICTURE_RESOURCE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_PROFILE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_PROFILE_LIST_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_REFERENCE_SLOT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_SESSION_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_SESSION_MEMORY_REQUIREMENTS_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_SESSION_PARAMETERS_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_SESSION_PARAMETERS_UPDATE_INFO_KHR`

## [](#_version_history)Version History

- Revision 0.1, 2019-11-21 (Tony Zlatinski)
  
  - Initial draft
- Revision 0.2, 2019-11-27 (Tony Zlatinski)
  
  - Make vulkan video core common between decode and encode
- Revision 1, March 29 2021 (Tony Zlatinski)
  
  - Spec and API updates.
- Revision 2, August 1 2021 (Srinath Kumarapuram)
  
  - Rename `VkVideoCapabilitiesFlagBitsKHR` to `VkVideoCapabilityFlagBitsKHR` (along with the names of enumerants it defines) and `VkVideoCapabilitiesFlagsKHR` to `VkVideoCapabilityFlagsKHR`, following Vulkan naming conventions.
- Revision 3, 2022-03-16 (Ahmed Abdelkhalek)
  
  - Relocate Std header version reporting/requesting from codec-operation specific extensions to this extension.
  - Make Std header versions codec-operation specific instead of only codec-specific.
- Revision 4, 2022-05-30 (Daniel Rakos)
  
  - Refactor the video format query APIs and related language
  - Extend VkResult with video-specific error codes
- Revision 5, 2022-08-11 (Daniel Rakos)
  
  - Add `VkVideoSessionParametersCreateFlagsKHR`
  - Remove `VkVideoCodingQualityPresetFlagsKHR`
  - Rename `VkQueueFamilyQueryResultStatusProperties2KHR` to `VkQueueFamilyQueryResultStatusPropertiesKHR`
  - Rename `VkVideoQueueFamilyProperties2KHR` to `VkQueueFamilyVideoPropertiesKHR`
  - Rename `VkVideoProfileKHR` to `VkVideoProfileInfoKHR`
  - Rename `VkVideoProfilesKHR` to `VkVideoProfileListInfoKHR`
  - Rename `VkVideoGetMemoryPropertiesKHR` to `VkVideoSessionMemoryRequirementsKHR`
  - Rename `VkVideoBindMemoryKHR` to `VkBindVideoSessionMemoryInfoKHR`
  - Fix `pNext` constness of `VkPhysicalDeviceVideoFormatInfoKHR` and `VkVideoSessionMemoryRequirementsKHR`
  - Fix incorrectly named value enums in bit enum types `VkVideoCodecOperationFlagBitsKHR` and `VkVideoChromaSubsamplingFlagBitsKHR`
  - Remove unnecessary default values from `VkVideoSessionCreateFlagBitsKHR` and `VkVideoCodingControlFlagBitsKHR`
  - Eliminate nested pointer in `VkVideoSessionMemoryRequirementsKHR`
  - Rename `VkVideoPictureResourceKHR` to `VkVideoPictureResourceInfoKHR`
  - Rename `VkVideoReferenceSlotKHR` to `VkVideoReferenceSlotInfoKHR`
- Revision 6, 2022-09-18 (Daniel Rakos)
  
  - Rename the `maxReferencePicturesSlotsCount` and `maxReferencePicturesActiveCount` fields of `VkVideoCapabilitiesKHR` and `VkVideoSessionCreateInfoKHR` to `maxDpbSlots` and `maxActiveReferencePictures`, respectively, to clarify their meaning
  - Rename `capabilityFlags` to `flags` in `VkVideoCapabilitiesKHR`
  - Rename `videoPictureExtentGranularity` to `pictureAccessGranularity` in `VkVideoCapabilitiesKHR`
  - Rename `minExtent` and `maxExtent` to `minCodedExtent` and `maxCodedExtent`, respectively, in `VkVideoCapabilitiesKHR`
  - Rename `referencePicturesFormat` to `referencePictureFormat` in `VkVideoSessionCreateInfoKHR`
- Revision 7, 2022-09-26 (Daniel Rakos)
  
  - Change type of `VkVideoReferenceSlotInfoKHR::slotIndex` from `int8_t` to `int32_t`
- Revision 8, 2022-09-29 (Daniel Rakos)
  
  - Extension is no longer provisional

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_video_queue).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0