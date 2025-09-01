# VK\_QCOM\_tile\_properties(3) Manual Page

## Name

VK\_QCOM\_tile\_properties - device extension



## [](#_registered_extension_number)Registered Extension Number

485

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_KHR\_dynamic\_rendering

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_tile_properties%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_tile_properties%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_QCOM\_tile\_properties](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_QCOM_tile_properties.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-07-11

**Interactions and External Dependencies**

- This extension interacts with `VK_EXT_subpass_merge_feedback`

**Contributors**

- Jonathan Wicks, Qualcomm Technologies, Inc.
- Jonathan Tinkham, Qualcomm Technologies, Inc.
- Arpit Agarwal, Qualcomm Technologies, Inc.
- Jeff Leger, Qualcomm Technologies, Inc.

## [](#_description)Description

This extension allows an application to query the tile properties. This extension supports both renderpasses and dynamic rendering.

## [](#_new_commands)New Commands

- [vkGetDynamicRenderingTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDynamicRenderingTilePropertiesQCOM.html)
- [vkGetFramebufferTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetFramebufferTilePropertiesQCOM.html)

## [](#_new_structures)New Structures

- [VkTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTilePropertiesQCOM.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceTilePropertiesFeaturesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTilePropertiesFeaturesQCOM.html)

If [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html) or [Vulkan Version 1.3](#versions-1.3) is supported:

- [VkRenderingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_QCOM_TILE_PROPERTIES_EXTENSION_NAME`
- `VK_QCOM_TILE_PROPERTIES_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TILE_PROPERTIES_FEATURES_QCOM`
  - `VK_STRUCTURE_TYPE_TILE_PROPERTIES_QCOM`

## [](#_version_history)Version History

- Revision 1, 2022-07-11 (Arpit Agarwal)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QCOM_tile_properties).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0