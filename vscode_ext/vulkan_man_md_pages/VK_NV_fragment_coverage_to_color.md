# VK_NV_fragment_coverage_to_color(3) Manual Page

## Name

VK_NV_fragment_coverage_to_color - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

150

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_fragment_coverage_to_color%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_fragment_coverage_to_color%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-05-21

**Contributors**  
- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows the fragment coverage value, represented as an
integer bitmask, to be substituted for a color output being written to a
single-component color attachment with integer components (e.g.
`VK_FORMAT_R8_UINT`). The functionality provided by this extension is
different from simply writing the `SampleMask` fragment shader output,
in that the coverage value written to the framebuffer is taken after
stencil test and depth test, as well as after fragment operations such
as alpha-to-coverage.

This functionality may be useful for deferred rendering algorithms,
where the second pass needs to know which samples belong to which
original fragments.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html):

  - [VkPipelineCoverageToColorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageToColorStateCreateInfoNV.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkPipelineCoverageToColorStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageToColorStateCreateFlagsNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_FRAGMENT_COVERAGE_TO_COLOR_EXTENSION_NAME`

- `VK_NV_FRAGMENT_COVERAGE_TO_COLOR_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_TO_COLOR_STATE_CREATE_INFO_NV`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-05-21 (Jeff Bolz)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_fragment_coverage_to_color"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
