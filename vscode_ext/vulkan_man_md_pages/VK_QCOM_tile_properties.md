# VK_QCOM_tile_properties(3) Manual Page

## Name

VK_QCOM_tile_properties - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

485

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_3

- Interacts with VK_KHR_dynamic_rendering

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_tile_properties%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_tile_properties%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_QCOM_tile_properties](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_QCOM_tile_properties.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-07-11

**Interactions and External Dependencies**  
- This extension interacts with
  [`VK_EXT_subpass_merge_feedback`](VK_EXT_subpass_merge_feedback.html)

**Contributors**  
- Jonathan Wicks, Qualcomm Technologies, Inc.

- Jonathan Tinkham, Qualcomm Technologies, Inc.

- Arpit Agarwal, Qualcomm Technologies, Inc.

- Jeff Leger, Qualcomm Technologies, Inc.

## <a href="#_description" class="anchor"></a>Description

This extension allows an application to query the tile properties. This
extension supports both renderpasses and dynamic rendering.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetDynamicRenderingTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDynamicRenderingTilePropertiesQCOM.html)

- [vkGetFramebufferTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetFramebufferTilePropertiesQCOM.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTilePropertiesQCOM.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceTilePropertiesFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTilePropertiesFeaturesQCOM.html)

If [VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html) or [Version
1.3](#versions-1.3) is supported:

- [VkRenderingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_QCOM_TILE_PROPERTIES_EXTENSION_NAME`

- `VK_QCOM_TILE_PROPERTIES_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TILE_PROPERTIES_FEATURES_QCOM`

  - `VK_STRUCTURE_TYPE_TILE_PROPERTIES_QCOM`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-07-11 (Arpit Agarwal)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QCOM_tile_properties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
