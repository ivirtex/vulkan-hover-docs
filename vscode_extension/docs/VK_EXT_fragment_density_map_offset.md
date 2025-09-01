# VK\_EXT\_fragment\_density\_map\_offset(3) Manual Page

## Name

VK\_EXT\_fragment\_density\_map\_offset - device extension



## [](#_registered_extension_number)Registered Extension Number

620

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)  
and  
[VK\_EXT\_fragment\_density\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map.html)  
and  
     [VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html)  
     or  
     [Vulkan Version 1.2](#versions-1.2)  
and  
     [Vulkan Version 1.3](#versions-1.3)  
     or  
     [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html)

## [](#_contact)Contact

- Connor Abbott [\[GitHub\]cwabbott0](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_fragment_density_map_offset%5D%20%40cwabbott0%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_fragment_density_map_offset%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_fragment\_density\_map\_offset](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_fragment_density_map_offset.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-02-14

**Contributors**

- Connor Abbott, Valve Corporation
- Matthew Netsch, Qualcomm Technologies, Inc.
- Jonathan Wicks, Qualcomm Technologies, Inc.
- Jonathan Tinkham, Qualcomm Technologies, Inc.
- Jeff Leger, Qualcomm Technologies, Inc.
- Manan Katwala, Qualcomm Technologies, Inc.
- Mike Blumenkrantz, Valve Corporation

## [](#_description)Description

This extension allows an application to specify offsets to a fragment density map attachment, changing the framebuffer location where density values are applied to without having to regenerate the fragment density map.

## [](#_new_commands)New Commands

- [vkCmdEndRendering2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndRendering2EXT.html)

## [](#_new_structures)New Structures

- [VkRenderingEndInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingEndInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceFragmentDensityMapOffsetFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentDensityMapOffsetFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceFragmentDensityMapOffsetPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentDensityMapOffsetPropertiesEXT.html)
- Extending [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassEndInfo.html), [VkRenderingEndInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingEndInfoEXT.html):
  
  - [VkRenderPassFragmentDensityMapOffsetEndInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassFragmentDensityMapOffsetEndInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_FRAGMENT_DENSITY_MAP_OFFSET_EXTENSION_NAME`
- `VK_EXT_FRAGMENT_DENSITY_MAP_OFFSET_SPEC_VERSION`
- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html):
  
  - `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_OFFSET_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_OFFSET_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_RENDERING_END_INFO_EXT`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_FRAGMENT_DENSITY_MAP_OFFSET_END_INFO_EXT`

## [](#_version_history)Version History

- Revision 1, 2025-02-14 (Connor Abbott)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_fragment_density_map_offset).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0