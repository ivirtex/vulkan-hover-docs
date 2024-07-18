# VK_NV_linear_color_attachment(3) Manual Page

## Name

VK_NV_linear_color_attachment - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

431

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_3

- Interacts with VK_KHR_format_feature_flags2

## <a href="#_contact" class="anchor"></a>Contact

- sourav parmar <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_linear_color_attachment%5D%20@souravpNV%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_linear_color_attachment%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>souravpNV</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-12-02

**Interactions and External Dependencies**  
- This extension requires
  [`VK_KHR_format_feature_flags2`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html)

**Contributors**  
- Pat Brown, NVIDIA

- Piers Daniell, NVIDIA

- Sourav Parmar, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension expands support for using `VK_IMAGE_TILING_LINEAR` images
as color attachments when all the color attachments in the render pass
instance have `VK_IMAGE_TILING_LINEAR` tiling. This extension adds a new
flag bit `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV` that
extends the existing
[VkFormatFeatureFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2KHR.html) bits.
This flag **can** be set for renderable color formats in the
[VkFormatProperties3KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties3KHR.html)::`linearTilingFeatures`
format properties structure member. Formats with the
`VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV` flag **may** be
used as color attachments as long as all the color attachments in the
render pass instance have `VK_IMAGE_TILING_LINEAR` tiling, and the
formats their images views are created with have
[VkFormatProperties3KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties3KHR.html)::`linearTilingFeatures`
which include `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`. This
extension supports both dynamic rendering and traditional render passes.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceLinearColorAttachmentFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLinearColorAttachmentFeaturesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_LINEAR_COLOR_ATTACHMENT_EXTENSION_NAME`

- `VK_NV_LINEAR_COLOR_ATTACHMENT_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINEAR_COLOR_ATTACHMENT_FEATURES_NV`

If [VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html) or
[Version 1.3](#versions-1.3) is supported:

- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2.html):

  - `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-11-29 (sourav parmar)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_linear_color_attachment"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
