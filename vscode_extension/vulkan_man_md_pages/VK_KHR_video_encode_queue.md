# VK_KHR_video_encode_queue(3) Manual Page

## Name

VK_KHR_video_encode_queue - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

300

## <a href="#_revision" class="anchor"></a>Revision

12

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html)  
and  
     [VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html)  
     or  
     [Version 1.3](#versions-1.3)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_3

- Interacts with VK_KHR_format_feature_flags2

## <a href="#_contact" class="anchor"></a>Contact

- Ahmed Abdelkhalek <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_encode_queue%5D%20@aabdelkh%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_encode_queue%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>aabdelkh</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_video_encode_queue](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_encode_queue.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

- Aidan Fabius, Core Avionics & Industrial Inc.

- Lynne Iribarren, Independent

## <a href="#_description" class="anchor"></a>Description

This extension builds upon the
[`VK_KHR_video_queue`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html) extension by adding
common APIs specific to video encoding and thus enabling implementations
to expose queue families supporting video encode operations.

More specifically, it adds video encode specific capabilities and a new
command buffer command that allows recording video encode operations
against a video session.

This extension is to be used in conjunction with other codec specific
video encode extensions that enable encoding video sequences of specific
video compression standards.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdEncodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEncodeVideoKHR.html)

- [vkGetEncodedVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetEncodedVideoSessionParametersKHR.html)

- [vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR.html)

- [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)

- [VkVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelPropertiesKHR.html)

- [VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlLayerInfoKHR.html)

- [VkVideoEncodeSessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersFeedbackInfoKHR.html)

- [VkVideoEncodeSessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersGetInfoKHR.html)

- Extending [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html):

  - [VkQueryPoolVideoEncodeFeedbackCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolVideoEncodeFeedbackCreateInfoKHR.html)

- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html):

  - [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilitiesKHR.html)

- Extending
  [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html),
  [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html):

  - [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html)

- Extending
  [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html),
  [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html):

  - [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelInfoKHR.html)

- Extending [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html),
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html):

  - [VkVideoEncodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeUsageInfoKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkVideoEncodeCapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilityFlagBitsKHR.html)

- [VkVideoEncodeContentFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeContentFlagBitsKHR.html)

- [VkVideoEncodeFeedbackFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeFeedbackFlagBitsKHR.html)

- [VkVideoEncodeRateControlModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlModeFlagBitsKHR.html)

- [VkVideoEncodeTuningModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeTuningModeKHR.html)

- [VkVideoEncodeUsageFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeUsageFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkVideoEncodeCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilityFlagsKHR.html)

- [VkVideoEncodeContentFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeContentFlagsKHR.html)

- [VkVideoEncodeFeedbackFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeFeedbackFlagsKHR.html)

- [VkVideoEncodeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeFlagsKHR.html)

- [VkVideoEncodeRateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlFlagsKHR.html)

- [VkVideoEncodeRateControlModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlModeFlagsKHR.html)

- [VkVideoEncodeUsageFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeUsageFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_VIDEO_ENCODE_QUEUE_EXTENSION_NAME`

- `VK_KHR_VIDEO_ENCODE_QUEUE_SPEC_VERSION`

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html):

  - `VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR`

  - `VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR`

- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html):

  - `VK_BUFFER_USAGE_VIDEO_ENCODE_DST_BIT_KHR`

  - `VK_BUFFER_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`

- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html):

  - `VK_FORMAT_FEATURE_VIDEO_ENCODE_DPB_BIT_KHR`

  - `VK_FORMAT_FEATURE_VIDEO_ENCODE_INPUT_BIT_KHR`

- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html):

  - `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DPB_KHR`

  - `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DST_KHR`

  - `VK_IMAGE_LAYOUT_VIDEO_ENCODE_SRC_KHR`

- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html):

  - `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`

  - `VK_IMAGE_USAGE_VIDEO_ENCODE_DST_BIT_KHR`

  - `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`

- Extending [VkQueryResultStatusKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultStatusKHR.html):

  - `VK_QUERY_RESULT_STATUS_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_KHR`

- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html):

  - `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`

- Extending [VkQueueFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFlagBits.html):

  - `VK_QUEUE_VIDEO_ENCODE_BIT_KHR`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

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

- Extending
  [VkVideoCodingControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlFlagBitsKHR.html):

  - `VK_VIDEO_CODING_CONTROL_ENCODE_QUALITY_LEVEL_BIT_KHR`

  - `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR`

- Extending
  [VkVideoSessionCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateFlagBitsKHR.html):

  - `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_PARAMETER_OPTIMIZATIONS_BIT_KHR`

If [VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html) or
[Version 1.3](#versions-1.3) is supported:

- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2.html):

  - `VK_FORMAT_FEATURE_2_VIDEO_ENCODE_DPB_BIT_KHR`

  - `VK_FORMAT_FEATURE_2_VIDEO_ENCODE_INPUT_BIT_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-07-23 (Ahmed Abdelkhalek)

  - Initial draft

- Revision 1.1, 10/29/2019 (Tony Zlatinski)

  - Updated the reserved spec tokens and renamed VkVideoEncoderKHR to
    VkVideoSessionKHR

- Revision 1.6, Jan 08 2020 (Tony Zlatinski)

  - API unify with the video_decode_queue spec

- Revision 2, March 29 2021 (Tony Zlatinski)

  - Spec and API updates.

- Revision 3, 2021-09-30 (Jon Leech)

  - Add interaction with
    [`VK_KHR_format_feature_flags2`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html)
    to `vk.xml`

- Revision 4, 2022-02-10 (Ahmed Abdelkhalek)

  - Updates to encode capability interface

- Revision 5, 2022-03-31 (Ahmed Abdelkhalek)

  - Remove redundant VkVideoEncodeInfoKHR.codedExtent

- Revision 6, 2022-07-18 (Daniel Rakos)

  - Remove `VkVideoEncodeRateControlFlagBitsKHR` and
    `VkVideoEncodeFlagBitsKHR` as they contain no defined flags for now

  - Add `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR` and
    `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_LAYER_BIT_KHR` to
    indicate rate control and rate control layer change requests,
    respectively, in video coding control operations

- Revision 7, 2022-08-12 (Daniel Rakos)

  - Add VkVideoEncodeUsageInfoKHR structure and related flags

- Revision 8, 2023-03-06 (Daniel Rakos)

  - Replace `VK_QUERY_TYPE_VIDEO_ENCODE_BITSTREAM_BUFFER_RANGE_KHR`
    queries with more generic `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`
    queries that can be extended in the future with more feedback values

  - Rename `dstBitstreamBuffer`, `dstBitstreamBufferOffset`, and
    `dstBitstreamBufferMaxRange` in `VkVideoEncodeInfoKHR` to
    `dstBuffer`, `dstBufferOffset`, and `dstBufferRange`, respectively,
    for consistency with the naming convention in the video decode
    extensions

  - Change the type of `rateControlLayerCount` and `qualityLevelCount`
    in `VkVideoEncodeCapabilitiesKHR` from `uint8_t` to `uint32_t` and
    rename them to `maxRateControlLayers` and `maxQualityLevels`,
    respectively

  - Change the type of `averageBitrate` and `maxBitrate` in
    `` VkVideoEncodeRateControlLayerInfoKHR` `` from `uint32_t` to
    `uint64_t`

  - Fixed the definition of rate control flag bits and added the new
    `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DEFAULT_KHR` constant to indicate
    implementation-specific automatic rate control

  - Change the type of `VkVideoEncodeRateControlInfoKHR::layerCount`
    from `uint8_t` to `uint32_t`

  - Rename `pLayerConfigs` to `pLayers` in
    `VkVideoEncodeRateControlInfoKHR`

- Revision 9, 2023-03-28 (Daniel Rakos)

  - Removed `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_LAYER_BIT_KHR`
    and the ability to change the state of individual rate control
    layers

  - Added new `VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_HAS_OVERRIDES_BIT_KHR`
    flag to video encode feedback queries

  - Added new video session create flag
    `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_PARAMETER_OPTIMIZATIONS_BIT_KHR`
    to opt-in to video session and encoding parameter optimizations

  - Added the `vkGetEncodedVideoSessionParametersKHR` command to enable
    retrieving encoded video session parameter data

  - Moved `virtualBufferSizeInMs` and `initialVirtualBufferSizeInMs`
    from `VkVideoEncodeRateControlLayerInfoKHR` to
    `VkVideoEncodeRateControlInfoKHR`

  - Added `maxBitrate` capability

  - Renamed `inputImageDataFillAlignment` capability to
    `encodeInputPictureGranularity` to better reflect its purpose

  - Added new `vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR`
    command and related structures to enable querying recommended
    settings for video encode quality levels

  - Added `VK_VIDEO_CODING_CONTROL_ENCODE_QUALITY_LEVEL_BIT_KHR` flag
    and `VkVideoEncodeQualityLevelInfoKHR` structure to allow
    controlling video encode quality level and removed `qualityLevel`
    from the encode operation parameters

- Revision 10, 2023-07-19 (Daniel Rakos)

  - Added
    `VK_QUERY_RESULT_STATUS_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_KHR`
    query result status code and the related capability flag
    `VK_VIDEO_ENCODE_CAPABILITY_INSUFFICIENT_BITSTREAM_BUFFER_RANGE_DETECTION_BIT_KHR`

- Revision 11, 2023-09-04 (Daniel Rakos)

  - Extension is no longer provisional

- Revision 12, 2023-12-05 (Daniel Rakos)

  - Require the specification of a reconstructed picture in all cases,
    except when the video session was created with no DPB slots to match
    shipping implementations

  - Make DPB slot activation behavior codec-specific to continue
    allowing application control over reference picture setup now that a
    reconstructed picture is always mandatory

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_video_encode_queue"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
