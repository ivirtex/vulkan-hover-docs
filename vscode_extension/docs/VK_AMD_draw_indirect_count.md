# VK\_AMD\_draw\_indirect\_count(3) Manual Page

## Name

VK\_AMD\_draw\_indirect\_count - device extension



## [](#_registered_extension_number)Registered Extension Number

34

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_draw\_indirect\_count](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_draw_indirect_count.html) extension
  
  - Which in turn was *promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Daniel Rakos [\[GitHub\]drakos-amd](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_draw_indirect_count%5D%20%40drakos-amd%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_draw_indirect_count%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-08-23

**IP Status**

No known IP claims.

**Contributors**

- Matthaeus G. Chajdas, AMD
- Derrick Owens, AMD
- Graham Sellers, AMD
- Daniel Rakos, AMD
- Dominik Witczak, AMD

## [](#_description)Description

This extension allows an application to source the number of draws for indirect drawing commands from a buffer. This enables applications to generate an arbitrary number of drawing commands and execute them without host intervention.

## [](#_promotion_to_vk_khr_draw_indirect_count)Promotion to `VK_KHR_draw_indirect_count`

All functionality in this extension is included in `VK_KHR_draw_indirect_count`, with the suffix changed to KHR. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_commands)New Commands

- [vkCmdDrawIndexedIndirectCountAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexedIndirectCountAMD.html)
- [vkCmdDrawIndirectCountAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndirectCountAMD.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_DRAW_INDIRECT_COUNT_EXTENSION_NAME`
- `VK_AMD_DRAW_INDIRECT_COUNT_SPEC_VERSION`

## [](#_version_history)Version History

- Revision 2, 2016-08-23 (Dominik Witczak)
  
  - Minor fixes
- Revision 1, 2016-07-21 (Matthaeus Chajdas)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_draw_indirect_count).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0