# VK\_KHR\_video\_encode\_queue(3) Manual Page

## Name

VK\_KHR\_video\_encode\_queue - device extension



## [](#_registered_extension_number)Registered Extension Number

300

## [](#_revision)Revision

12

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html)  
and  
     [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html)  
     or  
     [Vulkan Version 1.3](#versions-1.3)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_KHR\_format\_feature\_flags2

## [](#_contact)Contact

- Ahmed Abdelkhalek [\[GitHub\]aabdelkh](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_encode_queue%5D%20%40aabdelkh%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_encode_queue%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_video\_encode\_queue](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_encode_queue.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-12-05

**IP Status**

No known IP claims.

**Contributors**

- Ahmed Abdelkhalek, AMD
- Damien Kessler, NVIDIA
- George Hao, AMD
- Jake Beju, AMD
- Peter Fang, AMD
- Piers Daniell, NVIDIA
- Srinath Kumarapuram, NVIDIA
- Thomas J. Meier, NVIDIA
- Tony Zlatinski, NVIDIA
- Ravi Chaudhary, NVIDIA
- Yang Liu, AMD
- Daniel Rakos, RasterGrid
- Ping Liu, Intel
- Aidan Fabius, Core Avionics &amp; Industrial Inc.
- Lynne Iribarren, Independent

## [](#_description)Description

This extension builds upon the `VK_KHR_video_queue` extension by adding common APIs specific to video encoding and thus enabling implementations to expose queue families supporting video encode operations.

More specifically, it adds video encode specific capabilities and a new command buffer command that allows recording video encode operations against a video session.

This extension is to be used in conjunction with other codec specific video encode extensions that enable encoding video sequences of specific video compression standards.

## [](#_new_commands)New Commands

- [vkCmdEncodeVideoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEncodeVideoKHR.html)
- [vkGetEncodedVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetEncodedVideoSessionParametersKHR.html)
- [vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR.html)

## [](#_new_structures)New Structures

- [VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR.html)
- [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)
- [VkVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelPropertiesKHR.html)
- [VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlLayerInfoKHR.html)
- [VkVideoEncodeSessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionParametersFeedbackInfoKHR.html)
- [VkVideoEncodeSessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionParametersGetInfoKHR.html)
- Extending [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html):
  
  - [VkQueryPoolVideoEncodeFeedbackCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolVideoEncodeFeedbackCreateInfoKHR.html)
- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html):
  
  - [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilitiesKHR.html)
- Extending [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlInfoKHR.html), [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoBeginCodingInfoKHR.html):
  
  - [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlInfoKHR.html)
- Extending [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlInfoKHR.html), [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html):
  
  - [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelInfoKHR.html)
- Extending [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html), [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html):
  
  - [VkVideoEncodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeUsageInfoKHR.html)

## [](#_new_enums)New Enums

- [VkVideoEncodeCapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilityFlagBitsKHR.html)
- [VkVideoEncodeContentFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeContentFlagBitsKHR.html)
- [VkVideoEncodeFeedbackFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeFeedbackFlagBitsKHR.html)
- [VkVideoEncodeRateControlModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlModeFlagBitsKHR.html)
- [VkVideoEncodeTuningModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeTuningModeKHR.html)
- [VkVideoEncodeUsageFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeUsageFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkVideoEncodeCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilityFlagsKHR.html)
- [VkVideoEncodeContentFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeContentFlagsKHR.html)
- [VkVideoEncodeFeedbackFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeFeedbackFlagsKHR.html)
- [VkVideoEncodeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeFlagsKHR.html)
- [VkVideoEncodeRateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlFlagsKHR.html)
- [VkVideoEncodeRateControlModeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlModeFlagsKHR.html)
- [VkVideoEncodeUsageFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeUsageFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_VIDEO_ENCODE_QUEUE_EXTENSION_NAME`
- `VK_KHR_VIDEO_ENCODE_QUEUE_SPEC_VERSION`
- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR`
  - `VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR`
- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html):
  
  - `VK_BUFFER_USAGE_VIDEO_ENCODE_DST_BIT_KHR`
  - `VK_BUFFER_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`
- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html):
  
  - `VK_FORMAT_FEATURE_VIDEO_ENCODE_DPB_BIT_KHR`
  - `VK_FORMAT_FEATURE_VIDEO_ENCODE_INPUT_BIT_KHR`
- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html):
  
  - `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DPB_KHR`
  - `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DST_KHR`
  - `VK_IMAGE_LAYOUT_VIDEO_ENCODE_SRC_KHR`
- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html):
  
  - `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`
  - `VK_IMAGE_USAGE_VIDEO_ENCODE_DST_BIT_KHR`
  - `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`
- Extending [VkQueryResultStatusKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryResultStatusKHR.html):
  
  - `VK_QUERY_RESULT_STATUS_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_KHR`
- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html):
  
  - `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`
- Extending [VkQueueFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFlagBits.html):
  
  - `VK_QUEUE_VIDEO_ENCODE_BIT_KHR`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_ENCODE_QUALITY_LEVEL_INFO_KHR`
  - `VK_STRUCTURE_TYPE_QUERY_POOL_VIDEO_ENCODE_FEEDBACK_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_CAPABILITIES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_QUALITY_LEVEL_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_QUALITY_LEVEL_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_RATE_CONTROL_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_RATE_CONTROL_LAYER_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_SESSION_PARAMETERS_FEEDBACK_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_SESSION_PARAMETERS_GET_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_USAGE_INFO_KHR`
- Extending [VkVideoCodingControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlFlagBitsKHR.html):
  
  - `VK_VIDEO_CODING_CONTROL_ENCODE_QUALITY_LEVEL_BIT_KHR`
  - `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR`
- Extending [VkVideoSessionCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateFlagBitsKHR.html):
  
  - `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_PARAMETER_OPTIMIZATIONS_BIT_KHR`

If [VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html) or [Vulkan Version 1.3](#versions-1.3) is supported:

- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_VIDEO_ENCODE_DPB_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_VIDEO_ENCODE_INPUT_BIT_KHR`

## [](#_version_history)Version History

- Revision 1, 2018-07-23 (Ahmed Abdelkhalek)
  
  - Initial draft
- Revision 1.1, 10/29/2019 (Tony Zlatinski)
  
  - Updated the reserved spec tokens and renamed VkVideoEncoderKHR to VkVideoSessionKHR
- Revision 1.6, Jan 08 2020 (Tony Zlatinski)
  
  - API unify with the video\_decode\_queue spec
- Revision 2, March 29 2021 (Tony Zlatinski)
  
  - Spec and API updates.
- Revision 3, 2021-09-30 (Jon Leech)
  
  - Add interaction with `VK_KHR_format_feature_flags2` to `vk.xml`
- Revision 4, 2022-02-10 (Ahmed Abdelkhalek)
  
  - Updates to encode capability interface
- Revision 5, 2022-03-31 (Ahmed Abdelkhalek)
  
  - Remove redundant VkVideoEncodeInfoKHR.codedExtent
- Revision 6, 2022-07-18 (Daniel Rakos)
  
  - Remove `VkVideoEncodeRateControlFlagBitsKHR` and `VkVideoEncodeFlagBitsKHR` as they contain no defined flags for now
  - Add `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR` and `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_LAYER_BIT_KHR` to indicate rate control and rate control layer change requests, respectively, in video coding control operations
- Revision 7, 2022-08-12 (Daniel Rakos)
  
  - Add VkVideoEncodeUsageInfoKHR structure and related flags
- Revision 8, 2023-03-06 (Daniel Rakos)
  
  - Replace `VK_QUERY_TYPE_VIDEO_ENCODE_BITSTREAM_BUFFER_RANGE_KHR` queries with more generic `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR` queries that can be extended in the future with more feedback values
  - Rename `dstBitstreamBuffer`, `dstBitstreamBufferOffset`, and `dstBitstreamBufferMaxRange` in `VkVideoEncodeInfoKHR` to `dstBuffer`, `dstBufferOffset`, and `dstBufferRange`, respectively, for consistency with the naming convention in the video decode extensions
  - Change the type of `rateControlLayerCount` and `qualityLevelCount` in `VkVideoEncodeCapabilitiesKHR` from `uint8_t` to `uint32_t` and rename them to `maxRateControlLayers` and `maxQualityLevels`, respectively
  - Change the type of `averageBitrate` and `maxBitrate` in `VkVideoEncodeRateControlLayerInfoKHR` from `uint32_t` to `uint64_t`
  - Fixed the definition of rate control flag bits and added the new `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DEFAULT_KHR` constant to indicate implementation-specific automatic rate control
  - Change the type of `VkVideoEncodeRateControlInfoKHR::layerCount` from `uint8_t` to `uint32_t`
  - Rename `pLayerConfigs` to `pLayers` in `VkVideoEncodeRateControlInfoKHR`
- Revision 9, 2023-03-28 (Daniel Rakos)
  
  - Removed `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_LAYER_BIT_KHR` and the ability to change the state of individual rate control layers
  - Added new `VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_HAS_OVERRIDES_BIT_KHR` flag to video encode feedback queries
  - Added new video session create flag `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_PARAMETER_OPTIMIZATIONS_BIT_KHR` to opt-in to video session and encoding parameter optimizations
  - Added the `vkGetEncodedVideoSessionParametersKHR` command to enable retrieving encoded video session parameter data
  - Moved `virtualBufferSizeInMs` and `initialVirtualBufferSizeInMs` from `VkVideoEncodeRateControlLayerInfoKHR` to `VkVideoEncodeRateControlInfoKHR`
  - Added `maxBitrate` capability
  - Renamed `inputImageDataFillAlignment` capability to `encodeInputPictureGranularity` to better reflect its purpose
  - Added new `vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR` command and related structures to enable querying recommended settings for video encode quality levels
  - Added `VK_VIDEO_CODING_CONTROL_ENCODE_QUALITY_LEVEL_BIT_KHR` flag and `VkVideoEncodeQualityLevelInfoKHR` structure to allow controlling video encode quality level and removed `qualityLevel` from the encode operation parameters
- Revision 10, 2023-07-19 (Daniel Rakos)
  
  - Added `VK_QUERY_RESULT_STATUS_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_KHR` query result status code and the related capability flag `VK_VIDEO_ENCODE_CAPABILITY_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_DETECTION_BIT_KHR`
- Revision 11, 2023-09-04 (Daniel Rakos)
  
  - Extension is no longer provisional
- Revision 12, 2023-12-05 (Daniel Rakos)
  
  - Require the specification of a reconstructed picture in all cases, except when the video session was created with no DPB slots to match shipping implementations
  - Make DPB slot activation behavior codec-specific to continue allowing application control over reference picture setup now that a reconstructed picture is always mandatory

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_video_encode_queue).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0