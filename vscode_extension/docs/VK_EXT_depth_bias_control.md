# VK\_EXT\_depth\_bias\_control(3) Manual Page

## Name

VK\_EXT\_depth\_bias\_control - device extension



## [](#_registered_extension_number)Registered Extension Number

284

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_use)Special Use

- [D3D support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Joshua Ashton [\[GitHub\]Joshua-Ashton](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_depth_bias_control%5D%20%40Joshua-Ashton%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_depth_bias_control%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_depth\_bias\_control](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_depth_bias_control.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-02-15

**IP Status**

No known IP claims.

**Contributors**

- Joshua Ashton, VALVE
- Hans-Kristian Arntzen, VALVE
- Mike Blumenkrantz, VALVE
- Georg Lehmann, VALVE
- Piers Daniell, NVIDIA
- Lionel Landwerlin, INTEL
- Tobias Hector, AMD
- Ricardo Garcia, IGALIA
- Jan-Harald Fredriksen, ARM
- Shahbaz Youssefi, GOOGLE
- Tom Olson, ARM

## [](#_description)Description

This extension adds a new structure, `VkDepthBiasRepresentationInfoEXT`, that can be added to a `pNext` chain of `VkPipelineRasterizationStateCreateInfo` and allows setting the scaling and representation of depth bias for a pipeline.

This state can also be set dynamically by using the new structure mentioned above in combination with the new `vkCmdSetDepthBias2EXT` command.

## [](#_new_commands)New Commands

- [vkCmdSetDepthBias2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthBias2EXT.html)

## [](#_new_structures)New Structures

- [VkDepthBiasInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasInfoEXT.html)
- Extending [VkDepthBiasInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasInfoEXT.html), [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html):
  
  - [VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasRepresentationInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDepthBiasControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthBiasControlFeaturesEXT.html)

## [](#_new_enums)New Enums

- [VkDepthBiasRepresentationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasRepresentationEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DEPTH_BIAS_CONTROL_EXTENSION_NAME`
- `VK_EXT_DEPTH_BIAS_CONTROL_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEPTH_BIAS_INFO_EXT`
  - `VK_STRUCTURE_TYPE_DEPTH_BIAS_REPRESENTATION_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_BIAS_CONTROL_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2022-09-22 (Joshua Ashton)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_depth_bias_control)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0