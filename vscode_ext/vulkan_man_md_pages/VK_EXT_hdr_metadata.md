# VK_EXT_hdr_metadata(3) Manual Page

## Name

VK_EXT_hdr_metadata - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

106

## <a href="#_revision" class="anchor"></a>Revision

2

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
2018-12-19

**IP Status**  
No known IP claims.

**Contributors**  
- Courtney Goeltzenleuchter, Google

## <a href="#_description" class="anchor"></a>Description

This extension defines two new structures and a function to assign SMPTE
(the Society of Motion Picture and Television Engineers) 2086 metadata
and CTA (Consumer Technology Association) 861.3 metadata to a swapchain.
The metadata includes the color primaries, white point, and luminance
range of the reference monitor, which all together define the color
volume containing all the possible colors the reference monitor can
produce. The reference monitor is the display where creative work is
done and creative intent is established. To preserve such creative
intent as much as possible and achieve consistent color reproduction on
different viewing displays, it is useful for the display pipeline to
know the color volume of the original reference monitor where content
was created or tuned. This avoids performing unnecessary mapping of
colors that are not displayable on the original reference monitor. The
metadata also includes the `maxContentLightLevel` and
`maxFrameAverageLightLevel` as defined by CTA 861.3.

While the intended purpose of the metadata is to assist in the
transformation between different color volumes of different displays and
help achieve better color reproduction, it is not in the scope of this
extension to define how exactly the metadata should be used in such a
process. It is up to the implementation to determine how to make use of
the metadata.

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

1\) Do we need a query function?

**PROPOSED**: No, Vulkan does not provide queries for state that the
application can track on its own.

2\) Should we specify default if not specified by the application?

**PROPOSED**: No, that leaves the default up to the display.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-12-27 (Courtney Goeltzenleuchter)

  - Initial version

- Revision 2, 2018-12-19 (Courtney Goeltzenleuchter)

  - Correct implicit validity for VkHdrMetadataEXT structure

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

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
