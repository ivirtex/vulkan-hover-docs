# VK\_KHR\_maintenance3(3) Manual Page

## Name

VK\_KHR\_maintenance3 - device extension



## [](#_registered_extension_number)Registered Extension Number

169

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance3%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance3%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-09-05

**Contributors**

- Jeff Bolz, NVIDIA

## [](#_description)Description

`VK_KHR_maintenance3` adds a collection of minor features that were intentionally left out or overlooked from the original Vulkan 1.0 release.

The new features are as follows:

- A limit on the maximum number of descriptors that are supported in a single descriptor set layout. Some implementations have a limit on the total size of descriptors in a set, which cannot be expressed in terms of the limits in Vulkan 1.0.
- A limit on the maximum size of a single memory allocation. Some platforms have kernel interfaces that limit the maximum size of an allocation.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_commands)New Commands

- [vkGetDescriptorSetLayoutSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetLayoutSupportKHR.html)

## [](#_new_structures)New Structures

- [VkDescriptorSetLayoutSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutSupportKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceMaintenance3PropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance3PropertiesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_MAINTENANCE3_EXTENSION_NAME`
- `VK_KHR_MAINTENANCE3_SPEC_VERSION`
- `VK_KHR_MAINTENANCE_3_EXTENSION_NAME`
- `VK_KHR_MAINTENANCE_3_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES_KHR`

## [](#_version_history)Version History

- Revision 1, 2017-08-22

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_maintenance3)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0