# VK_NVX_image_view_handle(3) Manual Page

## Name

VK_NVX_image_view_handle - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

31

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Eric Werness <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NVX_image_view_handle%5D%20@ewerness-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NVX_image_view_handle%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ewerness-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-04-03

**Contributors**  
- Eric Werness, NVIDIA

- Jeff Bolz, NVIDIA

- Daniel Koch, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows applications to query an opaque handle from an
image view for use as a sampled image or storage image. This provides no
direct functionality itself.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetImageViewAddressNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageViewAddressNVX.html)

- [vkGetImageViewHandleNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageViewHandleNVX.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkImageViewAddressPropertiesNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewAddressPropertiesNVX.html)

- [VkImageViewHandleInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewHandleInfoNVX.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NVX_IMAGE_VIEW_HANDLE_EXTENSION_NAME`

- `VK_NVX_IMAGE_VIEW_HANDLE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_ADDRESS_PROPERTIES_NVX`

  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_HANDLE_INFO_NVX`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2020-04-03 (Piers Daniell)

  - Add [vkGetImageViewAddressNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageViewAddressNVX.html)

- Revision 1, 2018-12-07 (Eric Werness)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NVX_image_view_handle"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
