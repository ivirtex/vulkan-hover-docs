# VK\_KHR\_index\_type\_uint8(3) Manual Page

## Name

VK\_KHR\_index\_type\_uint8 - device extension



## [](#_registered_extension_number)Registered Extension Number

534

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_index_type_uint8%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_index_type_uint8%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-06-06

**IP Status**

No known IP claims.

**Contributors**

- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension allows `uint8_t` indices to be used with [vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer.html).

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceIndexTypeUint8FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceIndexTypeUint8FeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_INDEX_TYPE_UINT8_EXTENSION_NAME`
- `VK_KHR_INDEX_TYPE_UINT8_SPEC_VERSION`
- Extending [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html):
  
  - `VK_INDEX_TYPE_UINT8_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INDEX_TYPE_UINT8_FEATURES_KHR`

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4 with the KHR suffix omitted. The original type, enum and command names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2023-06-06 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_index_type_uint8)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0