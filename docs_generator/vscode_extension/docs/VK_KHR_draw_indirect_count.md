# VK\_KHR\_draw\_indirect\_count(3) Manual Page

## Name

VK\_KHR\_draw\_indirect\_count - device extension



## [](#_registered_extension_number)Registered Extension Number

170

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_draw_indirect_count%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_draw_indirect_count%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-08-25

**IP Status**

No known IP claims.

**Contributors**

- Matthaeus G. Chajdas, AMD
- Derrick Owens, AMD
- Graham Sellers, AMD
- Daniel Rakos, AMD
- Dominik Witczak, AMD
- Piers Daniell, NVIDIA

## [](#_description)Description

This extension is based on the `VK_AMD_draw_indirect_count` extension. This extension allows an application to source the number of draws for indirect drawing calls from a buffer.

Applications might want to do culling on the GPU via a compute shader prior to drawing. This enables the application to generate an arbitrary number of drawing commands and execute them without host intervention.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the KHR suffix omitted. However, if Vulkan 1.2 is supported and this extension is not, the commands [vkCmdDrawIndirectCount](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndirectCount.html) and [vkCmdDrawIndexedIndirectCount](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexedIndirectCount.html) are optional. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_commands)New Commands

- [vkCmdDrawIndexedIndirectCountKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexedIndirectCountKHR.html)
- [vkCmdDrawIndirectCountKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndirectCountKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_DRAW_INDIRECT_COUNT_EXTENSION_NAME`
- `VK_KHR_DRAW_INDIRECT_COUNT_SPEC_VERSION`

## [](#_version_history)Version History

- Revision 1, 2017-08-25 (Piers Daniell)
  
  - Initial draft based on VK\_AMD\_draw\_indirect\_count

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_draw_indirect_count)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0