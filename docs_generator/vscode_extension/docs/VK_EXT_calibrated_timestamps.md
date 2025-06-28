# VK\_EXT\_calibrated\_timestamps(3) Manual Page

## Name

VK\_EXT\_calibrated\_timestamps - device extension



## [](#_registered_extension_number)Registered Extension Number

185

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_calibrated\_timestamps](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_calibrated_timestamps.html) extension

## [](#_contact)Contact

- Daniel Rakos [\[GitHub\]drakos-amd](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_calibrated_timestamps%5D%20%40drakos-amd%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_calibrated_timestamps%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_calibrated\_timestamps](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_calibrated_timestamps.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-10-04

**IP Status**

No known IP claims.

**Contributors**

- Matthaeus G. Chajdas, AMD
- Alan Harrison, AMD
- Derrick Owens, AMD
- Daniel Rakos, AMD
- Faith Ekstrand, Intel
- Keith Packard, Valve

## [](#_description)Description

This extension provides an interface to query calibrated timestamps obtained quasi simultaneously from two time domains.

## [](#_promotion_to_vk_khr_calibrated_timestamps)Promotion to `VK_KHR_calibrated_timestamps`

All functionality in this extension is included in `VK_KHR_calibrated_timestamps`, with the suffix changed to KHR. The original enum names are still available as aliases of the KHR functionality.

## [](#_new_commands)New Commands

- [vkGetCalibratedTimestampsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetCalibratedTimestampsEXT.html)
- [vkGetPhysicalDeviceCalibrateableTimeDomainsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCalibrateableTimeDomainsEXT.html)

## [](#_new_structures)New Structures

- [VkCalibratedTimestampInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCalibratedTimestampInfoEXT.html)

## [](#_new_enums)New Enums

- [VkTimeDomainEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimeDomainEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_CALIBRATED_TIMESTAMPS_EXTENSION_NAME`
- `VK_EXT_CALIBRATED_TIMESTAMPS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_CALIBRATED_TIMESTAMP_INFO_EXT`
- Extending [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimeDomainKHR.html):
  
  - `VK_TIME_DOMAIN_CLOCK_MONOTONIC_EXT`
  - `VK_TIME_DOMAIN_CLOCK_MONOTONIC_RAW_EXT`
  - `VK_TIME_DOMAIN_DEVICE_EXT`
  - `VK_TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_EXT`

## [](#_version_history)Version History

- Revision 2, 2021-03-16 (Lionel Landwerlin)
  
  - Specify requirement on device timestamps
- Revision 1, 2018-10-04 (Daniel Rakos)
  
  - Internal revisions.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_calibrated_timestamps)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0