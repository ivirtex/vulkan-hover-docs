# VK_NV_device_diagnostics_config(3) Manual Page

## Name

VK_NV_device_diagnostics_config - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

301

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Kedarnath Thangudu <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_device_diagnostics_config%5D%20@kthangudu%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_device_diagnostics_config%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>kthangudu</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-04-06

**Contributors**  
- Kedarnath Thangudu, NVIDIA

- Thomas Klein, NVIDIA

## <a href="#_description" class="anchor"></a>Description

Applications using Nvidia Nsight<sup>™</sup> Aftermath SDK for Vulkan to
integrate device crash dumps into their error reporting mechanisms,
**may** use this extension to configure options related to device crash
dump creation.

Version 2 of this extension adds
`VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_ERROR_REPORTING_BIT_NV`
which when set enables enhanced reporting of shader execution errors.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkDeviceDiagnosticsConfigCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDiagnosticsConfigCreateInfoNV.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceDiagnosticsConfigFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDiagnosticsConfigFeaturesNV.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkDeviceDiagnosticsConfigFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDiagnosticsConfigFlagBitsNV.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkDeviceDiagnosticsConfigFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDiagnosticsConfigFlagsNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_DEVICE_DIAGNOSTICS_CONFIG_EXTENSION_NAME`

- `VK_NV_DEVICE_DIAGNOSTICS_CONFIG_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DEVICE_DIAGNOSTICS_CONFIG_CREATE_INFO_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DIAGNOSTICS_CONFIG_FEATURES_NV`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-11-21 (Kedarnath Thangudu)

  - Internal revisions

- Revision 2, 2022-04-06 (Kedarnath Thangudu)

  - Added a config bit
    `VK_DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_ERROR_REPORTING_BIT_NV`

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_device_diagnostics_config"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
