# VK\_NVX\_image\_view\_handle(3) Manual Page

## Name

VK\_NVX\_image\_view\_handle - device extension



## [](#_registered_extension_number)Registered Extension Number

31

## [](#_revision)Revision

3

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Eric Werness [\[GitHub\]ewerness-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NVX_image_view_handle%5D%20%40ewerness-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NVX_image_view_handle%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-11-04

**Contributors**

- Eric Werness, NVIDIA
- Jeff Bolz, NVIDIA
- Daniel Koch, NVIDIA
- Liam Middlebrook, NVIDIA

## [](#_description)Description

This extension allows applications to query an opaque handle from an image view for use as a sampled image or storage image. This provides no direct functionality itself.

## [](#_new_commands)New Commands

- [vkGetImageViewAddressNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageViewAddressNVX.html)
- [vkGetImageViewHandle64NVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageViewHandle64NVX.html)
- [vkGetImageViewHandleNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageViewHandleNVX.html)

## [](#_new_structures)New Structures

- [VkImageViewAddressPropertiesNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewAddressPropertiesNVX.html)
- [VkImageViewHandleInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewHandleInfoNVX.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NVX_IMAGE_VIEW_HANDLE_EXTENSION_NAME`
- `VK_NVX_IMAGE_VIEW_HANDLE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_ADDRESS_PROPERTIES_NVX`
  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_HANDLE_INFO_NVX`

## [](#_version_history)Version History

- Revision 3, 2024-11-04 (Liam Middlebrook)
  
  - Add [vkGetImageViewHandle64NVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageViewHandle64NVX.html)
- Revision 2, 2020-04-03 (Piers Daniell)
  
  - Add [vkGetImageViewAddressNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageViewAddressNVX.html)
- Revision 1, 2018-12-07 (Eric Werness)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NVX_image_view_handle).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0