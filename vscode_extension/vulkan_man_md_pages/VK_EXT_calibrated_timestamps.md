# VK_EXT_calibrated_timestamps(3) Manual Page

## Name

VK_EXT_calibrated_timestamps - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

185

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to
  [VK_KHR_calibrated_timestamps](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_calibrated_timestamps.html)
  extension

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Rakos <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_calibrated_timestamps%5D%20@drakos-amd%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_calibrated_timestamps%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>drakos-amd</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_calibrated_timestamps](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_calibrated_timestamps.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension provides an interface to query calibrated timestamps
obtained quasi simultaneously from two time domains.

## <a href="#_promotion_to_vk_khr_calibrated_timestamps"
class="anchor"></a>Promotion to `VK_KHR_calibrated_timestamps`

All functionality in this extension is included in
[`VK_KHR_calibrated_timestamps`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_calibrated_timestamps.html),
with the suffix changed to KHR. The original enum names are still
available as aliases of the KHR functionality.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetCalibratedTimestampsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetCalibratedTimestampsEXT.html)

- [vkGetPhysicalDeviceCalibrateableTimeDomainsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceCalibrateableTimeDomainsEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkCalibratedTimestampInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCalibratedTimestampInfoEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkTimeDomainEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimeDomainEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_CALIBRATED_TIMESTAMPS_EXTENSION_NAME`

- `VK_EXT_CALIBRATED_TIMESTAMPS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_CALIBRATED_TIMESTAMP_INFO_EXT`

- Extending [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimeDomainKHR.html):

  - `VK_TIME_DOMAIN_CLOCK_MONOTONIC_EXT`

  - `VK_TIME_DOMAIN_CLOCK_MONOTONIC_RAW_EXT`

  - `VK_TIME_DOMAIN_DEVICE_EXT`

  - `VK_TIME_DOMAIN_QUERY_PERFORMANCE_COUNTER_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2021-03-16 (Lionel Landwerlin)

  - Specify requirement on device timestamps

- Revision 1, 2018-10-04 (Daniel Rakos)

  - Internal revisions.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_calibrated_timestamps"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
