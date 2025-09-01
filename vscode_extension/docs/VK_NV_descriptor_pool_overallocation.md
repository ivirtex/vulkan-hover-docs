# VK\_NV\_descriptor\_pool\_overallocation(3) Manual Page

## Name

VK\_NV\_descriptor\_pool\_overallocation - device extension



## [](#_registered_extension_number)Registered Extension Number

547

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_descriptor_pool_overallocation%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_descriptor_pool_overallocation%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-08-30

**Contributors**

- Jeff Bolz, NVIDIA

## [](#_description)Description

There are scenarios where the application does not know ahead of time how many descriptor sets it may need to allocate from a descriptor pool, or how many descriptors of any of the descriptor types it may need to allocate from the descriptor pool.

This extension gives applications the ability to request the implementation allow more sets or descriptors to be allocated than initially specified at descriptor pool creation time, subject to available resources.

The `VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_SETS_BIT_NV` flag lets the application allocate more than [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html)::`maxSets` descriptor sets, and the `VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_POOLS_BIT_NV` lets the application allocate more descriptors than initially specified by [VkDescriptorPoolSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolSize.html)::`descriptorCount` for any descriptor types.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_DESCRIPTOR_POOL_OVERALLOCATION_EXTENSION_NAME`
- `VK_NV_DESCRIPTOR_POOL_OVERALLOCATION_SPEC_VERSION`
- Extending [VkDescriptorPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateFlagBits.html):
  
  - `VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_POOLS_BIT_NV`
  - `VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_SETS_BIT_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_POOL_OVERALLOCATION_FEATURES_NV`

## [](#_version_history)Version History

- Revision 1, 2023-08-30 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_descriptor_pool_overallocation).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0