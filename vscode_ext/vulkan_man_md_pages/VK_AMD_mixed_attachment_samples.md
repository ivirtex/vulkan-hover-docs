# VK_AMD_mixed_attachment_samples(3) Manual Page

## Name

VK_AMD_mixed_attachment_samples - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

137

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Matthaeus G. Chajdas <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_mixed_attachment_samples%5D%20@anteru%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_mixed_attachment_samples%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>anteru</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-07-24

**Contributors**  
- Mais Alnasser, AMD

- Matthaeus G. Chajdas, AMD

- Maciej Jesionowski, AMD

- Daniel Rakos, AMD

## <a href="#_description" class="anchor"></a>Description

This extension enables applications to use multisampled rendering with a
depth/stencil sample count that is larger than the color sample count.
Having a depth/stencil sample count larger than the color sample count
allows maintaining geometry and coverage information at a higher sample
rate than color information. All samples are depth/stencil tested, but
only the first color sample count number of samples get a corresponding
color output.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_MIXED_ATTACHMENT_SAMPLES_EXTENSION_NAME`

- `VK_AMD_MIXED_ATTACHMENT_SAMPLES_SPEC_VERSION`

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-07-24 (Daniel Rakos)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_mixed_attachment_samples"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
