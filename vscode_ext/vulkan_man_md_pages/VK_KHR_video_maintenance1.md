# VK_KHR_video_maintenance1(3) Manual Page

## Name

VK_KHR_video_maintenance1 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

516

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Rakos <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_maintenance1%5D%20@aqnuep%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_maintenance1%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>aqnuep</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_video_maintenance1](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_maintenance1.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-07-27

**IP Status**  
No known IP claims.

**Contributors**  
- Ahmed Abdelkhalek, AMD

- Aidan Fabius, Core Avionics & Industrial Inc.

- Ping Liu, Intel

- Lynne Iribarren, Independent

- Srinath Kumarapuram, NVIDIA

- Tony Zlatinski, NVIDIA

- Daniel Rakos, RasterGrid

## <a href="#_description" class="anchor"></a>Description

`VK_KHR_video_maintenance1` adds a collection of minor video coding
features, none of which would warrant an entire extension of their own.

The new features are as follows:

- Allow creating buffers that can be used in video coding operations,
  independent of the used video profile, using the new buffer creation
  flag `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`.

- Allow creating images that can be used as decode output or encode
  input pictures, independent of the used video profile, using the new
  image creation flag
  `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`.

- Allow specifying queries used by video coding operations as part of
  the video coding command parameters, instead of using begin/end query
  when the video session is created using the new video session creation
  flag `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceVideoMaintenance1FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoMaintenance1FeaturesKHR.html)

- Extending [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html),
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html):

  - [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_VIDEO_MAINTENANCE_1_EXTENSION_NAME`

- `VK_KHR_VIDEO_MAINTENANCE_1_SPEC_VERSION`

- Extending [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlagBits.html):

  - `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`

- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html):

  - `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_MAINTENANCE_1_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_INLINE_QUERY_INFO_KHR`

- Extending
  [VkVideoSessionCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateFlagBitsKHR.html):

  - `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-07-27 (Daniel Rakos)

  - internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_video_maintenance1"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
