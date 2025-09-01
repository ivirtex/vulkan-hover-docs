# VK\_VALVE\_fragment\_density\_map\_layered(3) Manual Page

## Name

VK\_VALVE\_fragment\_density\_map\_layered - device extension



## [](#_registered_extension_number)Registered Extension Number

612

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html)  
     or  
     [Vulkan Version 1.4](#versions-1.4)  
and  
[VK\_EXT\_fragment\_density\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map.html)

## [](#_contact)Contact

- Connor Abbott [\[GitHub\]cwabbott0](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_VALVE_fragment_density_map_layered%5D%20%40cwabbott0%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_VALVE_fragment_density_map_layered%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-06-06

**Interactions and External Dependencies**

- Interacts with Vulkan 1.1.
- Interacts with Vulkan 1.3.
- Interacts with Vulkan 1.4.
- Interacts with `VK_EXT_fragment_density_map`.

**IP Status**

No known IP claims.

**Contributors**

- Connor Abbott, VALVE
- Mike Blumenkrantz, VALVE

## [](#_description)Description

This extension enables the use of layered fragment density maps.

## [](#_new_structures)New Structures

- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html):
  
  - [VkPipelineFragmentDensityMapLayeredCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineFragmentDensityMapLayeredCreateInfoVALVE.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceFragmentDensityMapLayeredFeaturesVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentDensityMapLayeredFeaturesVALVE.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceFragmentDensityMapLayeredPropertiesVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentDensityMapLayeredPropertiesVALVE.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_VALVE_FRAGMENT_DENSITY_MAP_LAYERED_EXTENSION_NAME`
- `VK_VALVE_FRAGMENT_DENSITY_MAP_LAYERED_SPEC_VERSION`
- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_PER_LAYER_FRAGMENT_DENSITY_BIT_VALVE`
- Extending [VkRenderingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFlagBits.html):
  
  - `VK_RENDERING_PER_LAYER_FRAGMENT_DENSITY_BIT_VALVE`
- Extending [VkRenderPassCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateFlagBits.html):
  
  - `VK_RENDER_PASS_CREATE_PER_LAYER_FRAGMENT_DENSITY_BIT_VALVE`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_LAYERED_FEATURES_VALVE`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_LAYERED_PROPERTIES_VALVE`
  - `VK_STRUCTURE_TYPE_PIPELINE_FRAGMENT_DENSITY_MAP_LAYERED_CREATE_INFO_VALVE`

## [](#_version_history)Version History

- Revision 1, 2025-06-06 (Mike Blumenkrantz)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_VALVE_fragment_density_map_layered).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0