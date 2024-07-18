# VK_KHR_video_queue(3) Manual Page

## Name

VK_KHR_video_queue - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

24

## <a href="#_revision" class="anchor"></a>Revision

8

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

     [Version 1.1](#versions-1.1)  
     and  
     [VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html)  
or  
[Version 1.3](#versions-1.3)  

## <a href="#_contact" class="anchor"></a>Contact

- Tony Zlatinski <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_queue%5D%20@tzlatinski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_queue%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tzlatinski</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_video_queue](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_queue.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension provides common APIs to enable exposing queue families
with support for video codec operations by introducing the following new
object types and related functionalities:

- Video session objects that represent and maintain the state needed to
  perform video codec operations.

- Video session parameters objects that act as a container for codec
  specific parameters.

In addition, it also introduces query commands that allow applications
to determine video coding related capabilities, and command buffer
commands that enable recording video coding operations against a video
session.

This extension is to be used in conjunction with other extensions that
enable specific video coding operations.

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html)

- [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkBindVideoSessionMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindVideoSessionMemoryKHR.html)

- [vkCmdBeginVideoCodingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginVideoCodingKHR.html)

- [vkCmdControlVideoCodingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdControlVideoCodingKHR.html)

- [vkCmdEndVideoCodingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndVideoCodingKHR.html)

- [vkCreateVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateVideoSessionKHR.html)

- [vkCreateVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateVideoSessionParametersKHR.html)

- [vkDestroyVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyVideoSessionKHR.html)

- [vkDestroyVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyVideoSessionParametersKHR.html)

- [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)

- [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html)

- [vkGetVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetVideoSessionMemoryRequirementsKHR.html)

- [vkUpdateVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateVideoSessionParametersKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkBindVideoSessionMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindVideoSessionMemoryInfoKHR.html)

- [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)

- [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html)

- [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)

- [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html)

- [VkVideoEndCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEndCodingInfoKHR.html)

- [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)

- [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html)

- [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)

- [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)

- [VkVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionMemoryRequirementsKHR.html)

- [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html)

- [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersUpdateInfoKHR.html)

- Extending
  [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html),
  [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html),
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html):

  - [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)

- Extending [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html):

  - [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)

- Extending [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html):

  - [VkQueueFamilyQueryResultStatusPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyQueryResultStatusPropertiesKHR.html)

  - [VkQueueFamilyVideoPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyVideoPropertiesKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkQueryResultStatusKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultStatusKHR.html)

- [VkVideoCapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilityFlagBitsKHR.html)

- [VkVideoChromaSubsamplingFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoChromaSubsamplingFlagBitsKHR.html)

- [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodecOperationFlagBitsKHR.html)

- [VkVideoCodingControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlFlagBitsKHR.html)

- [VkVideoComponentBitDepthFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoComponentBitDepthFlagBitsKHR.html)

- [VkVideoSessionCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkVideoBeginCodingFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingFlagsKHR.html)

- [VkVideoCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilityFlagsKHR.html)

- [VkVideoChromaSubsamplingFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoChromaSubsamplingFlagsKHR.html)

- [VkVideoCodecOperationFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodecOperationFlagsKHR.html)

- [VkVideoCodingControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlFlagsKHR.html)

- [VkVideoComponentBitDepthFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoComponentBitDepthFlagsKHR.html)

- [VkVideoEndCodingFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEndCodingFlagsKHR.html)

- [VkVideoSessionCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateFlagsKHR.html)

- [VkVideoSessionParametersCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_VIDEO_QUEUE_EXTENSION_NAME`

- `VK_KHR_VIDEO_QUEUE_SPEC_VERSION`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_VIDEO_SESSION_KHR`

  - `VK_OBJECT_TYPE_VIDEO_SESSION_PARAMETERS_KHR`

- Extending [VkQueryResultFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlagBits.html):

  - `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR`

- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html):

  - `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`

  - `VK_ERROR_VIDEO_PICTURE_LAYOUT_NOT_SUPPORTED_KHR`

  - `VK_ERROR_VIDEO_PROFILE_CODEC_NOT_SUPPORTED_KHR`

  - `VK_ERROR_VIDEO_PROFILE_FORMAT_NOT_SUPPORTED_KHR`

  - `VK_ERROR_VIDEO_PROFILE_OPERATION_NOT_SUPPORTED_KHR`

  - `VK_ERROR_VIDEO_STD_VERSION_NOT_SUPPORTED_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

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

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 0.1, 2019-11-21 (Tony Zlatinski)

  - Initial draft

- Revision 0.2, 2019-11-27 (Tony Zlatinski)

  - Make vulkan video core common between decode and encode

- Revision 1, March 29 2021 (Tony Zlatinski)

  - Spec and API updates.

- Revision 2, August 1 2021 (Srinath Kumarapuram)

  - Rename `VkVideoCapabilitiesFlagBitsKHR` to
    `VkVideoCapabilityFlagBitsKHR` (along with the names of enumerants
    it defines) and `VkVideoCapabilitiesFlagsKHR` to
    `VkVideoCapabilityFlagsKHR`, following Vulkan naming conventions.

- Revision 3, 2022-03-16 (Ahmed Abdelkhalek)

  - Relocate Std header version reporting/requesting from
    codec-operation specific extensions to this extension.

  - Make Std header versions codec-operation specific instead of only
    codec-specific.

- Revision 4, 2022-05-30 (Daniel Rakos)

  - Refactor the video format query APIs and related language

  - Extend VkResult with video-specific error codes

- Revision 5, 2022-08-11 (Daniel Rakos)

  - Add `VkVideoSessionParametersCreateFlagsKHR`

  - Remove `VkVideoCodingQualityPresetFlagsKHR`

  - Rename `VkQueueFamilyQueryResultStatusProperties2KHR` to
    `VkQueueFamilyQueryResultStatusPropertiesKHR`

  - Rename `VkVideoQueueFamilyProperties2KHR` to
    `VkQueueFamilyVideoPropertiesKHR`

  - Rename `VkVideoProfileKHR` to `VkVideoProfileInfoKHR`

  - Rename `VkVideoProfilesKHR` to `VkVideoProfileListInfoKHR`

  - Rename `VkVideoGetMemoryPropertiesKHR` to
    `VkVideoSessionMemoryRequirementsKHR`

  - Rename `VkVideoBindMemoryKHR` to `VkBindVideoSessionMemoryInfoKHR`

  - Fix `pNext` constness of `VkPhysicalDeviceVideoFormatInfoKHR` and
    `VkVideoSessionMemoryRequirementsKHR`

  - Fix incorrectly named value enums in bit enum types
    `VkVideoCodecOperationFlagBitsKHR` and
    `VkVideoChromaSubsamplingFlagBitsKHR`

  - Remove unnecessary default values from
    `VkVideoSessionCreateFlagBitsKHR` and
    `VkVideoCodingControlFlagBitsKHR`

  - Eliminate nested pointer in `VkVideoSessionMemoryRequirementsKHR`

  - Rename `VkVideoPictureResourceKHR` to
    `VkVideoPictureResourceInfoKHR`

  - Rename `VkVideoReferenceSlotKHR` to `VkVideoReferenceSlotInfoKHR`

- Revision 6, 2022-09-18 (Daniel Rakos)

  - Rename the `maxReferencePicturesSlotsCount` and
    `maxReferencePicturesActiveCount` fields of `VkVideoCapabilitiesKHR`
    and `VkVideoSessionCreateInfoKHR` to `maxDpbSlots` and
    `maxActiveReferencePictures`, respectively, to clarify their meaning

  - Rename `capabilityFlags` to `flags` in `VkVideoCapabilitiesKHR`

  - Rename `videoPictureExtentGranularity` to `pictureAccessGranularity`
    in `VkVideoCapabilitiesKHR`

  - Rename `minExtent` and `maxExtent` to `minCodedExtent` and
    `maxCodedExtent`, respectively, in `VkVideoCapabilitiesKHR`

  - Rename `referencePicturesFormat` to `referencePictureFormat` in
    `VkVideoSessionCreateInfoKHR`

- Revision 7, 2022-09-26 (Daniel Rakos)

  - Change type of `VkVideoReferenceSlotInfoKHR::slotIndex` from
    `int8_t` to `int32_t`

- Revision 8, 2022-09-29 (Daniel Rakos)

  - Extension is no longer provisional

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_video_queue"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
