# VK\_KHR\_video\_maintenance1(3) Manual Page

## Name

VK\_KHR\_video\_maintenance1 - device extension



## [](#_registered_extension_number)Registered Extension Number

516

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html)

## [](#_contact)Contact

- Daniel Rakos [\[GitHub\]aqnuep](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_maintenance1%5D%20%40aqnuep%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_maintenance1%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_video\_maintenance1](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_maintenance1.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-07-27

**IP Status**

No known IP claims.

**Contributors**

- Ahmed Abdelkhalek, AMD
- Aidan Fabius, Core Avionics &amp; Industrial Inc.
- Ping Liu, Intel
- Lynne Iribarren, Independent
- Srinath Kumarapuram, NVIDIA
- Tony Zlatinski, NVIDIA
- Daniel Rakos, RasterGrid

## [](#_description)Description

`VK_KHR_video_maintenance1` adds a collection of minor video coding features, none of which would warrant an entire extension of their own.

The new features are as follows:

- Allow creating buffers that can be used in video coding operations, independent of the used video profile, using the new buffer creation flag `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`.
- Allow creating images that can be used as decode output or encode input pictures, independent of the used video profile, using the new image creation flag `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`.
- Allow specifying queries used by video coding operations as part of the video coding command parameters, instead of using begin/end query when the video session is created using the new video session creation flag `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceVideoMaintenance1FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoMaintenance1FeaturesKHR.html)
- Extending [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html), [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html):
  
  - [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoInlineQueryInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_VIDEO_MAINTENANCE_1_EXTENSION_NAME`
- `VK_KHR_VIDEO_MAINTENANCE_1_SPEC_VERSION`
- Extending [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateFlagBits.html):
  
  - `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`
- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html):
  
  - `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_MAINTENANCE_1_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_INLINE_QUERY_INFO_KHR`
- Extending [VkVideoSessionCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateFlagBitsKHR.html):
  
  - `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`

## [](#_version_history)Version History

- Revision 1, 2023-07-27 (Daniel Rakos)
  
  - internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_video_maintenance1)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0