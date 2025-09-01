# VK\_NV\_extended\_sparse\_address\_space(3) Manual Page

## Name

VK\_NV\_extended\_sparse\_address\_space - device extension



## [](#_registered_extension_number)Registered Extension Number

493

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Russell Chou [\[GitHub\]russellcnv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_extended_sparse_address_space%5D%20%40russellcnv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_extended_sparse_address_space%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-10-03

**Contributors**

- Russell Chou, NVIDIA
- Christoph Kubisch, NVIDIA
- Eric Werness, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

Implementations may be able to support an extended address space for sparse memory resources, but only for a certain set of usages.

This extension adds a query for the extended limit, and the supported usages that are allowed for that limit. This limit is an increase to [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`sparseAddressSpaceSize` when the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) or [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) uses only usages that are supported.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceExtendedSparseAddressSpaceFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExtendedSparseAddressSpaceFeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_EXTENDED_SPARSE_ADDRESS_SPACE_EXTENSION_NAME`
- `VK_NV_EXTENDED_SPARSE_ADDRESS_SPACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_SPARSE_ADDRESS_SPACE_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_SPARSE_ADDRESS_SPACE_PROPERTIES_NV`

## [](#_version_history)Version History

- Revision 1, 2023-10-03 (Russell Chou)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_extended_sparse_address_space).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0