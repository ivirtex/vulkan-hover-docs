# VK\_EXT\_host\_query\_reset(3) Manual Page

## Name

VK\_EXT\_host\_query\_reset - device extension



## [](#_registered_extension_number)Registered Extension Number

262

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Bas Nieuwenhuizen [\[GitHub\]BNieuwenhuizen](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_host_query_reset%5D%20%40BNieuwenhuizen%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_host_query_reset%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-03-06

**IP Status**

No known IP claims.

**Contributors**

- Bas Nieuwenhuizen, Google
- Faith Ekstrand, Intel
- Jeff Bolz, NVIDIA
- Piers Daniell, NVIDIA

## [](#_description)Description

This extension adds a new function to reset queries from the host.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the EXT suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_commands)New Commands

- [vkResetQueryPoolEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetQueryPoolEXT.html)

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceHostQueryResetFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceHostQueryResetFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_HOST_QUERY_RESET_EXTENSION_NAME`
- `VK_EXT_HOST_QUERY_RESET_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2019-03-12 (Bas Nieuwenhuizen)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_host_query_reset)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0