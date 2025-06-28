# VK\_HUAWEI\_hdr\_vivid(3) Manual Page

## Name

VK\_HUAWEI\_hdr\_vivid - device extension



## [](#_registered_extension_number)Registered Extension Number

591

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)  
and  
[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)  
and  
[VK\_EXT\_hdr\_metadata](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_hdr_metadata.html)

## [](#_contact)Contact

- Zehui Lin [\[GitHub\]bactlink](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_HUAWEI_hdr_vivid%5D%20%40bactlink%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_HUAWEI_hdr_vivid%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-10-08

**IP Status**

No known IP claims.

**Contributors**

- Juntao Li, Huawei
- Pan Gao, Huawei
- Xiufeng Zhang, Huawei
- Zehui Lin, Huawei

## [](#_description)Description

This extension allows applications to assign HDR Vivid (T/UWA 005.1-2022) metadata to swapchains.

HDR Vivid is an HDR standard released by UWA (UHD World Association). It defines tone mapping from the metadata to better preserve the creator’s intentions and achieve better consistency across devices with different display capabilities.

## [](#_new_structures)New Structures

- Extending [VkHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHdrMetadataEXT.html):
  
  - [VkHdrVividDynamicMetadataHUAWEI](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHdrVividDynamicMetadataHUAWEI.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceHdrVividFeaturesHUAWEI](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceHdrVividFeaturesHUAWEI.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_HUAWEI_HDR_VIVID_EXTENSION_NAME`
- `VK_HUAWEI_HDR_VIVID_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_HDR_VIVID_DYNAMIC_METADATA_HUAWEI`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HDR_VIVID_FEATURES_HUAWEI`

## [](#_version_history)Version History

- Revision 1, 2024-10-08 (Zehui Lin)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_HUAWEI_hdr_vivid)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0