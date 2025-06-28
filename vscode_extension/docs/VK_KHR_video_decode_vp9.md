# VK\_KHR\_video\_decode\_vp9(3) Manual Page

## Name

VK\_KHR\_video\_decode\_vp9 - device extension



## [](#_registered_extension_number)Registered Extension Number

515

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_video\_decode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_queue.html)

## [](#_contact)Contact

- Ahmed Abdelkhalek [\[GitHub\]aabdelkh](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_decode_vp9%5D%20%40aabdelkh%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_decode_vp9%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_video\_decode\_vp9](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_decode_vp9.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-04-11

**IP Status**

No known IP claims.

**Contributors**

- Ahmed Abdelkhalek, AMD
- Benjamin Cheng, AMD
- Lynne Iribarren, Independent
- David Airlie, Red Hat, Inc.
- Ping Liu, Intel
- Srinath Kumarapuram, NVIDIA
- Vassili Nikolaev, NVIDIA
- Tony Zlatinski, NVIDIA
- Konda Raju, NVIDIA
- Daniel Almeida, Collabora
- Nicolas Dufresne, Collabora
- Daniel Rakos, RasterGrid

## [](#_description)Description

This extension builds upon the `VK_KHR_video_decode_queue` extension by adding support for decoding elementary video stream sequences compliant with the VP9 video compression standard.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceVideoDecodeVP9FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoDecodeVP9FeaturesKHR.html)
- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html):
  
  - [VkVideoDecodeVP9CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeVP9CapabilitiesKHR.html)
- Extending [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html):
  
  - [VkVideoDecodeVP9PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeVP9PictureInfoKHR.html)
- Extending [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html), [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html):
  
  - [VkVideoDecodeVP9ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeVP9ProfileInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_VIDEO_DECODE_VP9_EXTENSION_NAME`
- `VK_KHR_VIDEO_DECODE_VP9_SPEC_VERSION`
- `VK_MAX_VIDEO_VP9_REFERENCES_PER_FRAME_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_DECODE_VP9_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_VP9_CAPABILITIES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_VP9_PICTURE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_VP9_PROFILE_INFO_KHR`
- Extending [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodecOperationFlagBitsKHR.html):
  
  - `VK_VIDEO_CODEC_OPERATION_DECODE_VP9_BIT_KHR`

## [](#_version_history)Version History

- Revision 1, 2025-04-11 (Daniel Rakos)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_video_decode_vp9)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0