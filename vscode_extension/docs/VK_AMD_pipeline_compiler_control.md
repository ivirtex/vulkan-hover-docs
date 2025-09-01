# VK\_AMD\_pipeline\_compiler\_control(3) Manual Page

## Name

VK\_AMD\_pipeline\_compiler\_control - device extension



## [](#_registered_extension_number)Registered Extension Number

184

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Matthaeus G. Chajdas [\[GitHub\]anteru](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_pipeline_compiler_control%5D%20%40anteru%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_pipeline_compiler_control%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-07-26

**IP Status**

No known IP claims.

**Contributors**

- Matthaeus G. Chajdas, AMD
- Daniel Rakos, AMD
- Maciej Jesionowski, AMD
- Tobias Hector, AMD

## [](#_description)Description

This extension introduces [VkPipelineCompilerControlCreateInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCompilerControlCreateInfoAMD.html) structure that can be chained to a pipeline’s creation information to specify additional flags that affect pipeline compilation.

## [](#_new_structures)New Structures

- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), [VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html):
  
  - [VkPipelineCompilerControlCreateInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCompilerControlCreateInfoAMD.html)

## [](#_new_enums)New Enums

- [VkPipelineCompilerControlFlagBitsAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCompilerControlFlagBitsAMD.html)

## [](#_new_bitmasks)New Bitmasks

- [VkPipelineCompilerControlFlagsAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCompilerControlFlagsAMD.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_PIPELINE_COMPILER_CONTROL_EXTENSION_NAME`
- `VK_AMD_PIPELINE_COMPILER_CONTROL_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PIPELINE_COMPILER_CONTROL_CREATE_INFO_AMD`

## [](#_issues)Issues

None.

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2019-07-26 (Tobias Hector)
  
  - Initial revision.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_pipeline_compiler_control).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0