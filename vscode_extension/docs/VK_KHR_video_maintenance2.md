# VK\_KHR\_video\_maintenance2(3) Manual Page

## Name

VK\_KHR\_video\_maintenance2 - device extension



## [](#_registered_extension_number)Registered Extension Number

587

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html)

## [](#_api_interactions)API Interactions

- Interacts with VK\_KHR\_video\_decode\_av1
- Interacts with VK\_KHR\_video\_decode\_h264
- Interacts with VK\_KHR\_video\_decode\_h265
- Interacts with VK\_KHR\_video\_decode\_queue

## [](#_contact)Contact

- Daniel Rakos [\[GitHub\]aqnuep](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_maintenance2%5D%20%40aqnuep%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_maintenance2%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_video\_maintenance2](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_maintenance2.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-10-14

**IP Status**

No known IP claims.

**Contributors**

- Ahmed Abdelkhalek, AMD
- Benjamin Cheng, AMD
- Aidan Fabius, Core Avionics &amp; Industrial Inc.
- Ping Liu, Intel
- Lynne Iribarren, Independent
- Srinath Kumarapuram, NVIDIA
- Tony Zlatinski, NVIDIA
- Daniel Rakos, RasterGrid

## [](#_description)Description

`VK_KHR_video_maintenance2` adds a collection of minor video coding features, none of which would warrant an entire extension of their own.

The new features are as follows:

- Allow video coding control commands (such as video session reset) to be issued without the need for a bound video session parameters object for video decode operations that would otherwise require the use of video session parameters objects.
- Allow applications to specify codec-specific parameter sets inline for each decode operation instead of having to construct video session parameters objects.
- Require support for `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR` in all applicable video encode profiles.
- Provide additional guarantees on Video Std parameters that the encoder implementation will not override.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceVideoMaintenance2FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoMaintenance2FeaturesKHR.html)

If [VK\_KHR\_video\_decode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_av1.html) is supported:

- Extending [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html):
  
  - [VkVideoDecodeAV1InlineSessionParametersInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeAV1InlineSessionParametersInfoKHR.html)

If [VK\_KHR\_video\_decode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_h264.html) is supported:

- Extending [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html):
  
  - [VkVideoDecodeH264InlineSessionParametersInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264InlineSessionParametersInfoKHR.html)

If [VK\_KHR\_video\_decode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_h265.html) is supported:

- Extending [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html):
  
  - [VkVideoDecodeH265InlineSessionParametersInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265InlineSessionParametersInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_VIDEO_MAINTENANCE_2_EXTENSION_NAME`
- `VK_KHR_VIDEO_MAINTENANCE_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_MAINTENANCE_2_FEATURES_KHR`

If [VK\_KHR\_video\_decode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_av1.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_INLINE_SESSION_PARAMETERS_INFO_KHR`

If [VK\_KHR\_video\_decode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_h264.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_INLINE_SESSION_PARAMETERS_INFO_KHR`

If [VK\_KHR\_video\_decode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_h265.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_INLINE_SESSION_PARAMETERS_INFO_KHR`

If [VK\_KHR\_video\_decode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_queue.html) is supported:

- Extending [VkVideoSessionCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateFlagBitsKHR.html):
  
  - `VK_VIDEO_SESSION_CREATE_INLINE_SESSION_PARAMETERS_BIT_KHR`

## [](#_version_history)Version History

- Revision 1, 2024-10-14 (Daniel Rakos)
  
  - internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_video_maintenance2)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0