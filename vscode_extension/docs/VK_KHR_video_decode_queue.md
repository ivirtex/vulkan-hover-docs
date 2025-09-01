# VK\_KHR\_video\_decode\_queue(3) Manual Page

## Name

VK\_KHR\_video\_decode\_queue - device extension



## [](#_registered_extension_number)Registered Extension Number

25

## [](#_revision)Revision

8

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

- [jake.beju@amd.com](mailto:jake.beju@amd.com)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_video\_decode\_queue](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_decode_queue.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-12-05

**IP Status**

No known IP claims.

**Contributors**

- Ahmed Abdelkhalek, AMD
- Jake Beju, AMD
- Olivier Lapicque, NVIDIA
- Peter Fang, AMD
- Piers Daniell, NVIDIA
- Srinath Kumarapuram, NVIDIA
- Tony Zlatinski, NVIDIA
- Daniel Rakos, RasterGrid

## [](#_description)Description

This extension builds upon the `VK_KHR_video_queue` extension by adding common APIs specific to video decoding and thus enabling implementations to expose queue families supporting video decode operations.

More specifically, it adds video decode specific capabilities and a new command buffer command that allows recording video decode operations against a video session.

This extension is to be used in conjunction with other codec specific video decode extensions that enable decoding video sequences of specific video compression standards.

## [](#_new_commands)New Commands

- [vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDecodeVideoKHR.html)

## [](#_new_structures)New Structures

- [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html)
- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html):
  
  - [VkVideoDecodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeCapabilitiesKHR.html)
- Extending [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html), [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html):
  
  - [VkVideoDecodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeUsageInfoKHR.html)

## [](#_new_enums)New Enums

- [VkVideoDecodeCapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeCapabilityFlagBitsKHR.html)
- [VkVideoDecodeUsageFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeUsageFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkVideoDecodeCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeCapabilityFlagsKHR.html)
- [VkVideoDecodeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeFlagsKHR.html)
- [VkVideoDecodeUsageFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeUsageFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_VIDEO_DECODE_QUEUE_EXTENSION_NAME`
- `VK_KHR_VIDEO_DECODE_QUEUE_SPEC_VERSION`
- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR`
  - `VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR`
- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html):
  
  - `VK_BUFFER_USAGE_VIDEO_DECODE_DST_BIT_KHR`
  - `VK_BUFFER_USAGE_VIDEO_DECODE_SRC_BIT_KHR`
- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html):
  
  - `VK_FORMAT_FEATURE_VIDEO_DECODE_DPB_BIT_KHR`
  - `VK_FORMAT_FEATURE_VIDEO_DECODE_OUTPUT_BIT_KHR`
- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html):
  
  - `VK_IMAGE_LAYOUT_VIDEO_DECODE_DPB_KHR`
  - `VK_IMAGE_LAYOUT_VIDEO_DECODE_DST_KHR`
  - `VK_IMAGE_LAYOUT_VIDEO_DECODE_SRC_KHR`
- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html):
  
  - `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`
  - `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`
  - `VK_IMAGE_USAGE_VIDEO_DECODE_SRC_BIT_KHR`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`
- Extending [VkQueueFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFlagBits.html):
  
  - `VK_QUEUE_VIDEO_DECODE_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_CAPABILITIES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_USAGE_INFO_KHR`

If [VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html) or [Vulkan Version 1.3](#versions-1.3) is supported:

- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_VIDEO_DECODE_DPB_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_VIDEO_DECODE_OUTPUT_BIT_KHR`

## [](#_version_history)Version History

- Revision 1, 2018-6-11 (Peter Fang)
  
  - Initial draft
- Revision 1.5, Nov 09 2018 (Tony Zlatinski)
  
  - API Updates
- Revision 1.6, Jan 08 2020 (Tony Zlatinski)
  
  - API unify with the video\_encode\_queue spec
- Revision 1.7, March 29 2021 (Tony Zlatinski)
  
  - Spec and API updates.
- Revision 2, September 30 2021 (Jon Leech)
  
  - Add interaction with `VK_KHR_format_feature_flags2` to `vk.xml`
- Revision 3, 2022-02-25 (Ahmed Abdelkhalek)
  
  - Add VkVideoDecodeCapabilitiesKHR with new flags to report support for decode DPB and output coinciding in the same image, or in distinct images.
- Revision 4, 2022-03-31 (Ahmed Abdelkhalek)
  
  - Remove redundant VkVideoDecodeInfoKHR.coded{Offset|Extent}
- Revision 5, 2022-07-18 (Daniel Rakos)
  
  - Remove `VkVideoDecodeFlagBitsKHR` as it contains no defined flags for now
- Revision 6, 2022-08-12 (Daniel Rakos)
  
  - Add VkVideoDecodeUsageInfoKHR structure and related flags
- Revision 7, 2022-09-29 (Daniel Rakos)
  
  - Extension is no longer provisional
- Revision 8, 2023-12-05 (Daniel Rakos)
  
  - Require the specification of a reconstructed picture in all cases, except when the video session was created with no DPB slots to match shipping implementations
  - Make DPB slot activation behavior codec-specific to continue allowing application control over reference picture setup now that a reconstructed picture is always mandatory

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_video_decode_queue).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0