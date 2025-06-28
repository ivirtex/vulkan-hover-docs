# VK\_NV\_fragment\_coverage\_to\_color(3) Manual Page

## Name

VK\_NV\_fragment\_coverage\_to\_color - device extension



## [](#_registered_extension_number)Registered Extension Number

150

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_fragment_coverage_to_color%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_fragment_coverage_to_color%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-05-21

**Contributors**

- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension allows the fragment coverage value, represented as an integer bitmask, to be substituted for a color output being written to a single-component color attachment with integer components (e.g. `VK_FORMAT_R8_UINT`). The functionality provided by this extension is different from simply writing the `SampleMask` fragment shader output, in that the coverage value written to the framebuffer is taken after stencil test and depth test, as well as after fragment operations such as alpha-to-coverage.

This functionality may be useful for deferred rendering algorithms, where the second pass needs to know which samples belong to which original fragments.

## [](#_new_structures)New Structures

- Extending [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineMultisampleStateCreateInfo.html):
  
  - [VkPipelineCoverageToColorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageToColorStateCreateInfoNV.html)

## [](#_new_bitmasks)New Bitmasks

- [VkPipelineCoverageToColorStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageToColorStateCreateFlagsNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_FRAGMENT_COVERAGE_TO_COLOR_EXTENSION_NAME`
- `VK_NV_FRAGMENT_COVERAGE_TO_COLOR_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_TO_COLOR_STATE_CREATE_INFO_NV`

## [](#_version_history)Version History

- Revision 1, 2017-05-21 (Jeff Bolz)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_fragment_coverage_to_color)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0