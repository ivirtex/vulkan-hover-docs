# VK_NV_fill_rectangle(3) Manual Page

## Name

VK_NV_fill_rectangle - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

154

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_fill_rectangle%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_fill_rectangle%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-05-22

**Contributors**  
- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds a new [VkPolygonMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPolygonMode.html) `enum`
where a triangle is rasterized by computing and filling its axis-aligned
screen-space bounding box, disregarding the actual triangle edges. This
can be useful for drawing a rectangle without being split into two
triangles with an internal edge. It is also useful to minimize the
number of primitives that need to be drawn, particularly for a user
interface.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_FILL_RECTANGLE_EXTENSION_NAME`

- `VK_NV_FILL_RECTANGLE_SPEC_VERSION`

- Extending [VkPolygonMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPolygonMode.html):

  - `VK_POLYGON_MODE_FILL_RECTANGLE_NV`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-05-22 (Jeff Bolz)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_fill_rectangle"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
