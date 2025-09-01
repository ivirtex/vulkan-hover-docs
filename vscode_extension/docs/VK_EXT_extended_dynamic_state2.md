# VK\_EXT\_extended\_dynamic\_state2(3) Manual Page

## Name

VK\_EXT\_extended\_dynamic\_state2 - device extension



## [](#_registered_extension_number)Registered Extension Number

378

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Vikram Kushwaha [\[GitHub\]vkushwaha-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_extended_dynamic_state2%5D%20%40vkushwaha-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_extended_dynamic_state2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-04-12

**IP Status**

No known IP claims.

**Contributors**

- Vikram Kushwaha, NVIDIA
- Piers Daniell, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension adds some more dynamic state to support applications that need to reduce the number of pipeline state objects they compile and bind.

## [](#_new_commands)New Commands

- [vkCmdSetDepthBiasEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthBiasEnableEXT.html)
- [vkCmdSetLogicOpEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetLogicOpEXT.html)
- [vkCmdSetPatchControlPointsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPatchControlPointsEXT.html)
- [vkCmdSetPrimitiveRestartEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPrimitiveRestartEnableEXT.html)
- [vkCmdSetRasterizerDiscardEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRasterizerDiscardEnableEXT.html)

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceExtendedDynamicState2FeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExtendedDynamicState2FeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_EXTENDED_DYNAMIC_STATE_2_EXTENSION_NAME`
- `VK_EXT_EXTENDED_DYNAMIC_STATE_2_SPEC_VERSION`
- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE_EXT`
  - `VK_DYNAMIC_STATE_LOGIC_OP_EXT`
  - `VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT`
  - `VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE_EXT`
  - `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_2_FEATURES_EXT`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

The dynamic state enumerants `VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE_EXT`, `VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE_EXT`, and `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE_EXT`; and the corresponding commands in this extension are included in core Vulkan 1.3, with the EXT suffix omitted. The enumerants and commands for dynamic logic operation and patch control points are not promoted, nor is the feature structure. Extension interfaces that were promoted remain available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2021-04-12 (Vikram Kushwaha)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_extended_dynamic_state2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0