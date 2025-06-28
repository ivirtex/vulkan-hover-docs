# VK\_NV\_device\_diagnostics\_config(3) Manual Page

## Name

VK\_NV\_device\_diagnostics\_config - device extension



## [](#_registered_extension_number)Registered Extension Number

301

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Kedarnath Thangudu [\[GitHub\]kthangudu](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_device_diagnostics_config%5D%20%40kthangudu%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_device_diagnostics_config%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-04-06

**Contributors**

- Kedarnath Thangudu, NVIDIA
- Thomas Klein, NVIDIA

## [](#_description)Description

Applications using Nvidia Nsight™ Aftermath SDK for Vulkan to integrate device crash dumps into their error reporting mechanisms, **may** use this extension to configure options related to device crash dump creation.

Version 2 of this extension adds `VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_ERROR_REPORTING_BIT_NV` which when set enables enhanced reporting of shader execution errors.

## [](#_new_structures)New Structures

- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkDeviceDiagnosticsConfigCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDiagnosticsConfigCreateInfoNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDiagnosticsConfigFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDiagnosticsConfigFeaturesNV.html)

## [](#_new_enums)New Enums

- [VkDeviceDiagnosticsConfigFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDiagnosticsConfigFlagBitsNV.html)

## [](#_new_bitmasks)New Bitmasks

- [VkDeviceDiagnosticsConfigFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDiagnosticsConfigFlagsNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_DEVICE_DIAGNOSTICS_CONFIG_EXTENSION_NAME`
- `VK_NV_DEVICE_DIAGNOSTICS_CONFIG_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEVICE_DIAGNOSTICS_CONFIG_CREATE_INFO_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DIAGNOSTICS_CONFIG_FEATURES_NV`

## [](#_version_history)Version History

- Revision 1, 2019-11-21 (Kedarnath Thangudu)
  
  - Internal revisions
- Revision 2, 2022-04-06 (Kedarnath Thangudu)
  
  - Added a config bit `VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_ERROR_REPORTING_BIT_NV`

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_device_diagnostics_config)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0