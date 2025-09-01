# VK\_EXT\_separate\_stencil\_usage(3) Manual Page

## Name

VK\_EXT\_separate\_stencil\_usage - device extension



## [](#_registered_extension_number)Registered Extension Number

247

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Daniel Rakos [\[GitHub\]drakos-amd](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_separate_stencil_usage%5D%20%40drakos-amd%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_separate_stencil_usage%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-11-08

**IP Status**

No known IP claims.

**Contributors**

- Daniel Rakos, AMD
- Jordan Logan, AMD

## [](#_description)Description

This extension allows specifying separate usage flags for the stencil aspect of images with a depth-stencil format at image creation time.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the EXT suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_structures)New Structures

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html):
  
  - [VkImageStencilUsageCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SEPARATE_STENCIL_USAGE_EXTENSION_NAME`
- `VK_EXT_SEPARATE_STENCIL_USAGE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO_EXT`

## [](#_version_history)Version History

- Revision 1, 2018-11-08 (Daniel Rakos)
  
  - Internal revisions.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_separate_stencil_usage).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0