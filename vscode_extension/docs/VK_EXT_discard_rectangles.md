# VK\_EXT\_discard\_rectangles(3) Manual Page

## Name

VK\_EXT\_discard\_rectangles - device extension



## [](#_registered_extension_number)Registered Extension Number

100

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_discard_rectangles%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_discard_rectangles%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-01-18

**Interactions and External Dependencies**

- Interacts with `VK_KHR_device_group`
- Interacts with Vulkan 1.1

**Contributors**

- Daniel Koch, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension provides additional orthogonally aligned “discard rectangles” specified in framebuffer-space coordinates that restrict rasterization of all points, lines and triangles.

From zero to an implementation-dependent limit (specified by `maxDiscardRectangles`) number of discard rectangles can be operational at once. When one or more discard rectangles are active, rasterized fragments can either survive if the fragment is within any of the operational discard rectangles (`VK_DISCARD_RECTANGLE_MODE_INCLUSIVE_EXT` mode) or be rejected if the fragment is within any of the operational discard rectangles (`VK_DISCARD_RECTANGLE_MODE_EXCLUSIVE_EXT` mode).

These discard rectangles operate orthogonally to the existing scissor test functionality. The discard rectangles can be different for each physical device in a device group by specifying the device mask and setting discard rectangle dynamic state.

Version 2 of this extension introduces new dynamic states `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT` and `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT`, and the corresponding functions [vkCmdSetDiscardRectangleEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDiscardRectangleEnableEXT.html) and [vkCmdSetDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDiscardRectangleModeEXT.html). Applications that use these dynamic states must ensure the implementation advertises at least `specVersion` `2` of this extension.

## [](#_new_commands)New Commands

- [vkCmdSetDiscardRectangleEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDiscardRectangleEXT.html)
- [vkCmdSetDiscardRectangleEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDiscardRectangleEnableEXT.html)
- [vkCmdSetDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDiscardRectangleModeEXT.html)

## [](#_new_structures)New Structures

- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html):
  
  - [VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceDiscardRectanglePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDiscardRectanglePropertiesEXT.html)

## [](#_new_enums)New Enums

- [VkDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDiscardRectangleModeEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkPipelineDiscardRectangleStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDiscardRectangleStateCreateFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DISCARD_RECTANGLES_EXTENSION_NAME`
- `VK_EXT_DISCARD_RECTANGLES_SPEC_VERSION`
- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT`
  - `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT`
  - `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DISCARD_RECTANGLE_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_DISCARD_RECTANGLE_STATE_CREATE_INFO_EXT`

## [](#_version_history)Version History

- Revision 2, 2023-01-18 (Piers Daniell)
  
  - Add dynamic states for discard rectangle enable/disable and mode.
- Revision 1, 2016-12-22 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_discard_rectangles).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0