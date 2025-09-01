# VK\_KHR\_relaxed\_block\_layout(3) Manual Page

## Name

VK\_KHR\_relaxed\_block\_layout - device extension



## [](#_registered_extension_number)Registered Extension Number

145

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- John Kessenich [\[GitHub\]johnkslang](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_relaxed_block_layout%5D%20%40johnkslang%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_relaxed_block_layout%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-03-26

**IP Status**

No known IP claims.

**Contributors**

- John Kessenich, Google

## [](#_description)Description

The `VK_KHR_relaxed_block_layout` extension allows implementations to indicate they can support more variation in block `Offset` decorations. For example, placing a vector of three floats at an offset of 16×N + 4.

See [Offset and Stride Assignment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-resources-layout) for details.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_RELAXED_BLOCK_LAYOUT_EXTENSION_NAME`
- `VK_KHR_RELAXED_BLOCK_LAYOUT_SPEC_VERSION`

## [](#_version_history)Version History

- Revision 1, 2017-03-26 (JohnK)

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_relaxed_block_layout).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0