# VK_KHR_calibrated_timestamps(3) Manual Page

## Name

VK_KHR_calibrated_timestamps - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

544

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Rakos <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_calibrated_timestamps%5D%20@aqnuep%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_calibrated_timestamps%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>aqnuep</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_calibrated_timestamps](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_calibrated_timestamps.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension provides an interface to query calibrated timestamps
obtained quasi simultaneously from two time domains.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetCalibratedTimestampsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetCalibratedTimestampsKHR.html)

- [vkGetPhysicalDeviceCalibrateableTimeDomainsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceCalibrateableTimeDomainsKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkCalibratedTimestampInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCalibratedTimestampInfoKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkTimeDomainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimeDomainKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_CALIBRATED_TIMESTAMPS_EXTENSION_NAME`

- `VK_KHR_CALIBRATED_TIMESTAMPS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_CALIBRATED_TIMESTAMP_INFO_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-07-12 (Daniel Rakos)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_calibrated_timestamps"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
