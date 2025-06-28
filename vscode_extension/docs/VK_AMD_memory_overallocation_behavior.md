# VK\_AMD\_memory\_overallocation\_behavior(3) Manual Page

## Name

VK\_AMD\_memory\_overallocation\_behavior - device extension



## [](#_registered_extension_number)Registered Extension Number

190

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Martin Dinkov [\[GitHub\]mdinkov](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_memory_overallocation_behavior%5D%20%40mdinkov%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_memory_overallocation_behavior%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-09-19

**IP Status**

No known IP claims.

**Contributors**

- Martin Dinkov, AMD
- Matthaeus Chajdas, AMD
- Daniel Rakos, AMD
- Jon Campbell, AMD

## [](#_description)Description

This extension allows controlling whether explicit overallocation beyond the device memory heap sizes (reported by [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties.html)) is allowed or not. Overallocation may lead to performance loss and is not supported for all platforms.

## [](#_new_structures)New Structures

- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkDeviceMemoryOverallocationCreateInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryOverallocationCreateInfoAMD.html)

## [](#_new_enums)New Enums

- [VkMemoryOverallocationBehaviorAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryOverallocationBehaviorAMD.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_MEMORY_OVERALLOCATION_BEHAVIOR_EXTENSION_NAME`
- `VK_AMD_MEMORY_OVERALLOCATION_BEHAVIOR_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEVICE_MEMORY_OVERALLOCATION_CREATE_INFO_AMD`

## [](#_version_history)Version History

- Revision 1, 2018-09-19 (Martin Dinkov)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_memory_overallocation_behavior)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0