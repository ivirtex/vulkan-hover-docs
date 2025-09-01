# VK\_NV\_present\_metering(3) Manual Page

## Name

VK\_NV\_present\_metering - device extension



## [](#_registered_extension_number)Registered Extension Number

614

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

- **This is a *provisional* extension and must be used with caution. See the [description](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#boilerplate-provisional-header) of provisional header files for enablement and stability details.**

## [](#_contact)Contact

- Charles Hansen [\[GitHub\]chansen](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_present_metering%5D%20%40chansen%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_present_metering%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-01-08

**Provisional**

\*This extension is *provisional* and **should** not be used in production applications. The functionality defined by this extension **may** change in ways that break backwards compatibility between revisions, and before the final release of the non-provisional version of this extension.

**Contributors**

- Charles Hansen, NVIDIA
- Lionel Duc, NVIDIA

## [](#_description)Description

This extension is used to evenly meter presents.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePresentMeteringFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePresentMeteringFeaturesNV.html)
- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html):
  
  - [VkSetPresentConfigNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSetPresentConfigNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_PRESENT_METERING_EXTENSION_NAME`
- `VK_NV_PRESENT_METERING_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRESENT_METERING_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_SET_PRESENT_CONFIG_NV`

## [](#_issues)Issues

## [](#_version_history)Version History

- Revision 1, 2025-01-08 (Charles Hansen)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_present_metering).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0