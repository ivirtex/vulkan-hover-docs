# VK_EXT_depth_bias_control(3) Manual Page

## Name

VK_EXT_depth_bias_control - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

284

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">D3D support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Joshua Ashton <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_depth_bias_control%5D%20@Joshua-Ashton%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_depth_bias_control%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>Joshua-Ashton</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_depth_bias_control](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_depth_bias_control.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension adds a new structure, `VkDepthBiasRepresentationInfoEXT`,
that can be added to a `pNext` chain of
`VkPipelineRasterizationStateCreateInfo` and allows setting the scaling
and representation of depth bias for a pipeline.

This state can also be set dynamically by using the new structure
mentioned above in combination with the new `vkCmdSetDepthBias2EXT`
command.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetDepthBias2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBias2EXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDepthBiasInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasInfoEXT.html)

- Extending [VkDepthBiasInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasInfoEXT.html),
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html):

  - [VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasRepresentationInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceDepthBiasControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthBiasControlFeaturesEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkDepthBiasRepresentationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasRepresentationEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_DEPTH_BIAS_CONTROL_EXTENSION_NAME`

- `VK_EXT_DEPTH_BIAS_CONTROL_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DEPTH_BIAS_INFO_EXT`

  - `VK_STRUCTURE_TYPE_DEPTH_BIAS_REPRESENTATION_INFO_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_BIAS_CONTROL_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-09-22 (Joshua Ashton)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_depth_bias_control"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
