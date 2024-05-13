# VK_KHR_video_decode_queue(3) Manual Page

## Name

VK_KHR_video_decode_queue - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

25

## <a href="#_revision" class="anchor"></a>Revision

8

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

- <jake.beju@amd.com>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_video_decode_queue](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_decode_queue.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension builds upon the
[`VK_KHR_video_queue`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html) extension by adding
common APIs specific to video decoding and thus enabling implementations
to expose queue families supporting video decode operations.

More specifically, it adds video decode specific capabilities and a new
command buffer command that allows recording video decode operations
against a video session.

This extension is to be used in conjunction with other codec specific
video decode extensions that enable decoding video sequences of specific
video compression standards.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDecodeVideoKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)

- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html):

  - [VkVideoDecodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeCapabilitiesKHR.html)

- Extending [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html),
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html):

  - [VkVideoDecodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeUsageInfoKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkVideoDecodeCapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeCapabilityFlagBitsKHR.html)

- [VkVideoDecodeUsageFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeUsageFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkVideoDecodeCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeCapabilityFlagsKHR.html)

- [VkVideoDecodeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeFlagsKHR.html)

- [VkVideoDecodeUsageFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeUsageFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_VIDEO_DECODE_QUEUE_EXTENSION_NAME`

- `VK_KHR_VIDEO_DECODE_QUEUE_SPEC_VERSION`

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html):

  - `VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR`

  - `VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR`

- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html):

  - `VK_BUFFER_USAGE_VIDEO_DECODE_DST_BIT_KHR`

  - `VK_BUFFER_USAGE_VIDEO_DECODE_SRC_BIT_KHR`

- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html):

  - `VK_FORMAT_FEATURE_VIDEO_DECODE_DPB_BIT_KHR`

  - `VK_FORMAT_FEATURE_VIDEO_DECODE_OUTPUT_BIT_KHR`

- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html):

  - `VK_IMAGE_LAYOUT_VIDEO_DECODE_DPB_KHR`

  - `VK_IMAGE_LAYOUT_VIDEO_DECODE_DST_KHR`

  - `VK_IMAGE_LAYOUT_VIDEO_DECODE_SRC_KHR`

- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html):

  - `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`

  - `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`

  - `VK_IMAGE_USAGE_VIDEO_DECODE_SRC_BIT_KHR`

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`

- Extending [VkQueueFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFlagBits.html):

  - `VK_QUEUE_VIDEO_DECODE_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_CAPABILITIES_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_USAGE_INFO_KHR`

If [VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html) or
[Version 1.3](#versions-1.3) is supported:

- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2.html):

  - `VK_FORMAT_FEATURE_2_VIDEO_DECODE_DPB_BIT_KHR`

  - `VK_FORMAT_FEATURE_2_VIDEO_DECODE_OUTPUT_BIT_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-6-11 (Peter Fang)

  - Initial draft

- Revision 1.5, Nov 09 2018 (Tony Zlatinski)

  - API Updates

- Revision 1.6, Jan 08 2020 (Tony Zlatinski)

  - API unify with the video_encode_queue spec

- Revision 1.7, March 29 2021 (Tony Zlatinski)

  - Spec and API updates.

- Revision 2, September 30 2021 (Jon Leech)

  - Add interaction with
    [`VK_KHR_format_feature_flags2`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html)
    to `vk.xml`

- Revision 3, 2022-02-25 (Ahmed Abdelkhalek)

  - Add VkVideoDecodeCapabilitiesKHR with new flags to report support
    for decode DPB and output coinciding in the same image, or in
    distinct images.

- Revision 4, 2022-03-31 (Ahmed Abdelkhalek)

  - Remove redundant VkVideoDecodeInfoKHR.coded{Offset\|Extent}

- Revision 5, 2022-07-18 (Daniel Rakos)

  - Remove `VkVideoDecodeFlagBitsKHR` as it contains no defined flags
    for now

- Revision 6, 2022-08-12 (Daniel Rakos)

  - Add VkVideoDecodeUsageInfoKHR structure and related flags

- Revision 7, 2022-09-29 (Daniel Rakos)

  - Extension is no longer provisional

- Revision 8, 2023-12-05 (Daniel Rakos)

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
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_video_decode_queue"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
