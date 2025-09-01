# VK\_EXT\_hdr\_metadata(3) Manual Page

## Name

VK\_EXT\_hdr\_metadata - device extension



## [](#_registered_extension_number)Registered Extension Number

106

## [](#_revision)Revision

3

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)

## [](#_contact)Contact

- Courtney Goeltzenleuchter [\[GitHub\]courtney-g](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_hdr_metadata%5D%20%40courtney-g%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_hdr_metadata%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-03-26

**IP Status**

No known IP claims.

**Contributors**

- Courtney Goeltzenleuchter, Google
- Sebastian Wick, Red Hat Inc.
- Tobias Hector, AMD

## [](#_description)Description

This extension defines two new structures and a function to assign SMPTE (the Society of Motion Picture and Television Engineers) 2086 metadata and CTA (Consumer Technology Association) 861.3 metadata to a swapchain.

SMPTE 2086 metadata defines the color volume of the display on which the content was optimized for viewing and includes the color primaries, white point, and luminance range. When such content is reproduced on another display, this metadata can be used by the presentation engine to improve processing of images. For instance, values in the image can first be clamped to the color volume described in the metadata, and then what remains can be remapped to the color volume of the presentation engine.

CTA 861.3 metadata additionally includes the maximum intended luminance for the content and the maximum average light level across frames.

This extension does not define exactly how this metadata is used, however, it simply provides a mechanism to provide it to the presentation engine. Presentation engines may process the image based on the metadata before displaying it, resulting in the image being modified outside of Vulkan. For example, the clamping of colors in the image to the color volume may change those values in the image itself.

The metadata does not override or otherwise influence the color space and color encoding.

## [](#_new_commands)New Commands

- [vkSetHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetHdrMetadataEXT.html)

## [](#_new_structures)New Structures

- [VkHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHdrMetadataEXT.html)
- [VkXYColorEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkXYColorEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_HDR_METADATA_EXTENSION_NAME`
- `VK_EXT_HDR_METADATA_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_HDR_METADATA_EXT`

## [](#_issues)Issues

1\) Do we need a query function for the currently specified metadata?

No, Vulkan does not provide queries for state that the application can track on its own.

2\) Should we specify default metadata if not specified by the application?

No, the metadata is optional and the absence of the metadata is well-defined.

## [](#_version_history)Version History

- Revision 1, 2016-12-27 (Courtney Goeltzenleuchter)
  
  - Initial version
- Revision 2, 2018-12-19 (Courtney Goeltzenleuchter)
  
  - Correct implicit validity for VkHdrMetadataEXT structure
- Revision 3, 2024-03-26 (Tobias Hector &amp; Sebastian Wick)
  
  - Clarifications and removal of erroneous "reference monitor" term

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_hdr_metadata).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0