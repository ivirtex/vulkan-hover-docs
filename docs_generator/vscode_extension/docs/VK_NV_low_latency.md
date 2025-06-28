# VK\_NV\_low\_latency(3) Manual Page

## Name

VK\_NV\_low\_latency - device extension



## [](#_registered_extension_number)Registered Extension Number

311

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Charles Hansen [\[GitHub\]cshansen](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_low_latency%5D%20%40cshansen%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_low_latency%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-02-10

**Contributors**

- Charles Hansen, NVIDIA

## [](#_description)Description

This extension adds the [VkQueryLowLatencySupportNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryLowLatencySupportNV.html) structure, a structure used to query support for NVIDIA Reflex.

## [](#_new_structures)New Structures

- Extending [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html):
  
  - [VkQueryLowLatencySupportNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryLowLatencySupportNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_LOW_LATENCY_EXTENSION_NAME`
- `VK_NV_LOW_LATENCY_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_QUERY_LOW_LATENCY_SUPPORT_NV`

## [](#_issues)Issues

1\) Why does `VkQueryLowLatencySupportNV` have output parameters in an input chain?

**RESOLVED**: We are stuck with this for legacy reasons - we are aware this is bad behavior and this should not be used as a precedent for future extensions.

## [](#_version_history)Version History

- Revision 1, 2023-02-10 (Charles Hansen)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_low_latency)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0