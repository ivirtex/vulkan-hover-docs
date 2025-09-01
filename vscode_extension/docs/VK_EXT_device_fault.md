# VK\_EXT\_device\_fault(3) Manual Page

## Name

VK\_EXT\_device\_fault - device extension



## [](#_registered_extension_number)Registered Extension Number

342

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Ralph Potter \[GitLab]r\_potter

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_device\_fault](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_device_fault.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-03-10

**IP Status**

No known IP claims.

**Contributors**

- Ralph Potter, Samsung
- Stuart Smith, AMD
- Jan-Harald Fredriksen, ARM
- Mark Bellamy, ARM
- Andrew Ellem, Google
- Alex Walters, IMG
- Jeff Bolz, NVIDIA
- Baldur Karlsson, Valve

## [](#_description)Description

Device loss can be triggered by a variety of issues, including invalid API usage, implementation errors, or hardware failures.

This extension introduces a new command: [vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceFaultInfoEXT.html), which may be called subsequent to a `VK_ERROR_DEVICE_LOST` error code having been returned by the implementation. This command allows developers to query for additional information on GPU faults which may have caused device loss, and to generate binary crash dumps, which may be loaded into external tools for further diagnosis.

## [](#_new_commands)New Commands

- [vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceFaultInfoEXT.html)

## [](#_new_structures)New Structures

- [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultAddressInfoEXT.html)
- [VkDeviceFaultCountsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultCountsEXT.html)
- [VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultInfoEXT.html)
- [VkDeviceFaultVendorBinaryHeaderVersionOneEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultVendorBinaryHeaderVersionOneEXT.html)
- [VkDeviceFaultVendorInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultVendorInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceFaultFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFaultFeaturesEXT.html)

## [](#_new_enums)New Enums

- [VkDeviceFaultAddressTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultAddressTypeEXT.html)
- [VkDeviceFaultVendorBinaryHeaderVersionEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultVendorBinaryHeaderVersionEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DEVICE_FAULT_EXTENSION_NAME`
- `VK_EXT_DEVICE_FAULT_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEVICE_FAULT_COUNTS_EXT`
  - `VK_STRUCTURE_TYPE_DEVICE_FAULT_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FAULT_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 2, 2023-04-05 (Ralph Potter)
  
  - Restored two missing members to the XML definition of VkDeviceFaultVendorBinaryHeaderVersionOneEXT. No functional change to the specification.
- Revision 1, 2020-10-19 (Ralph Potter)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_device_fault).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0