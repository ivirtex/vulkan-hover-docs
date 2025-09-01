# VK\_NV\_fill\_rectangle(3) Manual Page

## Name

VK\_NV\_fill\_rectangle - device extension



## [](#_registered_extension_number)Registered Extension Number

154

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_fill_rectangle%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_fill_rectangle%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-05-22

**Contributors**

- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension adds a new [VkPolygonMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPolygonMode.html) `enum` where a triangle is rasterized by computing and filling its axis-aligned screen-space bounding box, disregarding the actual triangle edges. This can be useful for drawing a rectangle without being split into two triangles with an internal edge. It is also useful to minimize the number of primitives that need to be drawn, particularly for a user interface.

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_FILL_RECTANGLE_EXTENSION_NAME`
- `VK_NV_FILL_RECTANGLE_SPEC_VERSION`
- Extending [VkPolygonMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPolygonMode.html):
  
  - `VK_POLYGON_MODE_FILL_RECTANGLE_NV`

## [](#_version_history)Version History

- Revision 1, 2017-05-22 (Jeff Bolz)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_fill_rectangle).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0