# VK\_AMD\_mixed\_attachment\_samples(3) Manual Page

## Name

VK\_AMD\_mixed\_attachment\_samples - device extension



## [](#_registered_extension_number)Registered Extension Number

137

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_KHR\_dynamic\_rendering

## [](#_contact)Contact

- Matthaeus G. Chajdas [\[GitHub\]anteru](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_mixed_attachment_samples%5D%20%40anteru%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_mixed_attachment_samples%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-07-24

**Contributors**

- Mais Alnasser, AMD
- Matthaeus G. Chajdas, AMD
- Maciej Jesionowski, AMD
- Daniel Rakos, AMD

## [](#_description)Description

This extension enables applications to use multisampled rendering with a depth/stencil sample count that is larger than the color sample count. Having a depth/stencil sample count larger than the color sample count allows maintaining geometry and coverage information at a higher sample rate than color information. All samples are depth/stencil tested, but only the first color sample count number of samples get a corresponding color output.

## [](#_new_structures)New Structures

If [Vulkan Version 1.3](#versions-1.3) or [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html) is supported:

- Extending [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html):
  
  - [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentSampleCountInfoAMD.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_MIXED_ATTACHMENT_SAMPLES_EXTENSION_NAME`
- `VK_AMD_MIXED_ATTACHMENT_SAMPLES_SPEC_VERSION`

If [Vulkan Version 1.3](#versions-1.3) or [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ATTACHMENT_SAMPLE_COUNT_INFO_AMD`

## [](#_issues)Issues

None.

## [](#_version_history)Version History

- Revision 1, 2017-07-24 (Daniel Rakos)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_mixed_attachment_samples)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0