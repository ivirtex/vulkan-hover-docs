# VK\_KHR\_calibrated\_timestamps(3) Manual Page

## Name

VK\_KHR\_calibrated\_timestamps - device extension



## [](#_registered_extension_number)Registered Extension Number

544

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Daniel Rakos [\[GitHub\]aqnuep](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_calibrated_timestamps%5D%20%40aqnuep%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_calibrated_timestamps%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_calibrated\_timestamps](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_calibrated_timestamps.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-07-12

**IP Status**

No known IP claims.

**Contributors**

- Matthaeus G. Chajdas, AMD
- Alan Harrison, AMD
- Derrick Owens, AMD
- Daniel Rakos, RasterGrid
- Faith Ekstrand, Intel
- Keith Packard, Valve

## [](#_description)Description

This extension provides an interface to query calibrated timestamps obtained quasi simultaneously from two time domains.

## [](#_new_commands)New Commands

- [vkGetCalibratedTimestampsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetCalibratedTimestampsKHR.html)
- [vkGetPhysicalDeviceCalibrateableTimeDomainsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCalibrateableTimeDomainsKHR.html)

## [](#_new_structures)New Structures

- [VkCalibratedTimestampInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCalibratedTimestampInfoKHR.html)

## [](#_new_enums)New Enums

- [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimeDomainKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_CALIBRATED_TIMESTAMPS_EXTENSION_NAME`
- `VK_KHR_CALIBRATED_TIMESTAMPS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_CALIBRATED_TIMESTAMP_INFO_KHR`

## [](#_version_history)Version History

- Revision 1, 2023-07-12 (Daniel Rakos)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_calibrated_timestamps).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0