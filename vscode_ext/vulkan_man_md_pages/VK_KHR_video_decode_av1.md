# VK_KHR_video_decode_av1(3) Manual Page

## Name

VK_KHR_video_decode_av1 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

513

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_video_decode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_queue.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Rakos <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_decode_av1%5D%20@aqnuep%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_decode_av1%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>aqnuep</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_video_decode_av1](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_decode_av1.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2024-01-02

**IP Status**  
No known IP claims.

**Contributors**  
- Ahmed Abdelkhalek, AMD

- Benjamin Cheng, AMD

- Ho Hin Lau, AMD

- Lynne Iribarren, Independent

- David Airlie, Red Hat, Inc.

- Ping Liu, Intel

- Srinath Kumarapuram, NVIDIA

- Vassili Nikolaev, NVIDIA

- Tony Zlatinski, NVIDIA

- Charlie Turner, Igalia

- Daniel Almeida, Collabora

- Nicolas Dufresne, Collabora

- Daniel Rakos, RasterGrid

## <a href="#_description" class="anchor"></a>Description

This extension builds upon the
[`VK_KHR_video_decode_queue`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_queue.html) extension
by adding support for decoding elementary video stream sequences
compliant with the AV1 video compression standard.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html):

  - [VkVideoDecodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1CapabilitiesKHR.html)

- Extending [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html):

  - [VkVideoDecodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1PictureInfoKHR.html)

- Extending [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html),
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html):

  - [VkVideoDecodeAV1ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1ProfileInfoKHR.html)

- Extending
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html):

  - [VkVideoDecodeAV1DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1DpbSlotInfoKHR.html)

- Extending
  [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html):

  - [VkVideoDecodeAV1SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1SessionParametersCreateInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_VIDEO_DECODE_AV1_EXTENSION_NAME`

- `VK_KHR_VIDEO_DECODE_AV1_SPEC_VERSION`

- `VK_MAX_VIDEO_AV1_REFERENCES_PER_FRAME_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_CAPABILITIES_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_DPB_SLOT_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_PICTURE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_PROFILE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_SESSION_PARAMETERS_CREATE_INFO_KHR`

- Extending
  [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodecOperationFlagBitsKHR.html):

  - `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2024-01-02 (Daniel Rakos)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_video_decode_av1"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
