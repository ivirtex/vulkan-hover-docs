# VK\_NV\_linear\_color\_attachment(3) Manual Page

## Name

VK\_NV\_linear\_color\_attachment - device extension



## [](#_registered_extension_number)Registered Extension Number

431

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_KHR\_format\_feature\_flags2

## [](#_contact)Contact

- sourav parmar [\[GitHub\]souravpNV](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_linear_color_attachment%5D%20%40souravpNV%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_linear_color_attachment%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-12-02

**Interactions and External Dependencies**

- This extension requires `VK_KHR_format_feature_flags2`

**Contributors**

- Pat Brown, NVIDIA
- Piers Daniell, NVIDIA
- Sourav Parmar, NVIDIA

## [](#_description)Description

This extension expands support for using `VK_IMAGE_TILING_LINEAR` images as color attachments when all the color attachments in the render pass instance have `VK_IMAGE_TILING_LINEAR` tiling. This extension adds a new flag bit `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV` that extends the existing [VkFormatFeatureFlagBits2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2KHR.html) bits. This flag **can** be set for renderable color formats in the [VkFormatProperties3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties3KHR.html)::`linearTilingFeatures` format properties structure member. Formats with the `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV` flag **may** be used as color attachments as long as all the color attachments in the render pass instance have `VK_IMAGE_TILING_LINEAR` tiling, and the formats their images views are created with have [VkFormatProperties3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties3KHR.html)::`linearTilingFeatures` which include `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`. This extension supports both dynamic rendering and traditional render passes.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceLinearColorAttachmentFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLinearColorAttachmentFeaturesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_LINEAR_COLOR_ATTACHMENT_EXTENSION_NAME`
- `VK_NV_LINEAR_COLOR_ATTACHMENT_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINEAR_COLOR_ATTACHMENT_FEATURES_NV`

If [VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html) or [Vulkan Version 1.3](#versions-1.3) is supported:

- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`

## [](#_version_history)Version History

- Revision 1, 2021-11-29 (sourav parmar)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_linear_color_attachment).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0