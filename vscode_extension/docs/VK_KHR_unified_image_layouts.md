# VK\_KHR\_unified\_image\_layouts(3) Manual Page

## Name

VK\_KHR\_unified\_image\_layouts - device extension



## [](#_registered_extension_number)Registered Extension Number

528

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_EXT\_attachment\_feedback\_loop\_layout
- Interacts with VK\_KHR\_dynamic\_rendering

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_unified_image_layouts%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_unified_image_layouts%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_unified\_image\_layouts](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_unified_image_layouts.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-10-15

**Interactions and External Dependencies**

- This extension interacts with `VK_EXT_attachment_feedback_loop_layout`
- This extension interacts with `VK_KHR_video_decode_queue`
- This extension interacts with `VK_KHR_video_encode_queue`
- This extension interacts with `VK_KHR_video_encode_quantization_map`

**Contributors**

- Ahmed Abdelkhalek, AMD
- Tobias Hector, AMD
- Jan-Harald Fredriksen, ARM
- Ting Wei, ARM
- Faith Ekstrand, Collabora
- Lina Versace, Google
- Shahbaz Youssefi, Google
- James Fitzpatrick, Imagination
- Daniel Story, Nintendo
- James Jones, NVIDIA
- Jeff Juliano, NVIDIA
- Piers Daniell, NVIDIA
- Tony Zlatinski, NVIDIA
- Matthew Netsch, Qualcomm
- Patrick Boyle, Qualcomm
- Daniel Rakos, RasterGrid
- Ralph Potter, Samsung
- Hans-Kristian Arntzen, VALVE
- Samuel Pitoiset, VALVE

## [](#_description)Description

This extension significantly simplifies synchronization in Vulkan by removing the need for image layout transitions in most cases. In particular, it guarantees that using the `VK_IMAGE_LAYOUT_GENERAL` layout everywhere possible is just as efficient as using the other layouts.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceUnifiedImageLayoutsFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceUnifiedImageLayoutsFeaturesKHR.html)

If [VK\_EXT\_attachment\_feedback\_loop\_layout](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_attachment_feedback_loop_layout.html) and [Vulkan Version 1.3](#versions-1.3) or [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html) is supported:

- Extending [VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentInfo.html):
  
  - [VkAttachmentFeedbackLoopInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentFeedbackLoopInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_UNIFIED_IMAGE_LAYOUTS_EXTENSION_NAME`
- `VK_KHR_UNIFIED_IMAGE_LAYOUTS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFIED_IMAGE_LAYOUTS_FEATURES_KHR`

If [VK\_EXT\_attachment\_feedback\_loop\_layout](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_attachment_feedback_loop_layout.html) and [Vulkan Version 1.3](#versions-1.3) or [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ATTACHMENT_FEEDBACK_LOOP_INFO_EXT`

## [](#_version_history)Version History

- Revision 1, 2024-10-15 (Shahbaz Youssefi)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_unified_image_layouts).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0