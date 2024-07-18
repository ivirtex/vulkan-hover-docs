# VK_EXT_hdr_metadata(3) Manual Page

## Name

VK_EXT_hdr_metadata - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

106

## <a href="#_revision" class="anchor"></a>Revision

3

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Courtney Goeltzenleuchter <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_hdr_metadata%5D%20@courtney-g%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_hdr_metadata%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>courtney-g</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2024-03-26

**IP Status**  
No known IP claims.

**Contributors**  
- Courtney Goeltzenleuchter, Google

- Sebastian Wick, Red Hat Inc.

- Tobias Hector, AMD

## <a href="#_description" class="anchor"></a>Description

This extension defines two new structures and a function to assign SMPTE
(the Society of Motion Picture and Television Engineers) 2086 metadata
and CTA (Consumer Technology Association) 861.3 metadata to a swapchain.

SMPTE 2086 metadata defines the color volume of the display on which the
content was optimized for viewing and includes the color primaries,
white point, and luminance range. When such content is reproduced on
another display, this metadata can be used by the presentation engine to
improve processing of images. For instance, values in the image can
first be clamped to the color volume described in the metadata, and then
what remains can be remapped to the color volume of the presentation
engine.

CTA 861.3 metadata additionally includes the maximum intended luminance
for the content and the maximum average light level across frames.

This extension does not define exactly how this metadata is used,
however, it simply provides a mechanism to provide it to the
presentation engine. Presentation engines may process the image based on
the metadata before displaying it, resulting in the image being modified
outside of Vulkan. For example, the clamping of colors in the image to
the color volume may change those values in the image itself.

The metadata does not override or otherwise influence the color space
and color encoding.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkSetHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetHdrMetadataEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHdrMetadataEXT.html)

- [VkXYColorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXYColorEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_HDR_METADATA_EXTENSION_NAME`

- `VK_EXT_HDR_METADATA_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_HDR_METADATA_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) Do we need a query function for the currently specified metadata?

No, Vulkan does not provide queries for state that the application can
track on its own.

2\) Should we specify default metadata if not specified by the
application?

No, the metadata is optional and the absence of the metadata is
well-defined.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-12-27 (Courtney Goeltzenleuchter)

  - Initial version

- Revision 2, 2018-12-19 (Courtney Goeltzenleuchter)

  - Correct implicit validity for VkHdrMetadataEXT structure

- Revision 3, 2024-03-26 (Tobias Hector & Sebastian Wick)

  - Clarifications and removal of erroneous "reference monitor" term

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_hdr_metadata"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
