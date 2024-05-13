# VK_EXT_discard_rectangles(3) Manual Page

## Name

VK_EXT_discard_rectangles - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

100

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_discard_rectangles%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_discard_rectangles%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-01-18

**Interactions and External Dependencies**  
- Interacts with [`VK_KHR_device_group`](VK_KHR_device_group.html)

- Interacts with Vulkan 1.1

**Contributors**  
- Daniel Koch, NVIDIA

- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension provides additional orthogonally aligned “discard
rectangles” specified in framebuffer-space coordinates that restrict
rasterization of all points, lines and triangles.

From zero to an implementation-dependent limit (specified by
`maxDiscardRectangles`) number of discard rectangles can be operational
at once. When one or more discard rectangles are active, rasterized
fragments can either survive if the fragment is within any of the
operational discard rectangles
(`VK_DISCARD_RECTANGLE_MODE_INCLUSIVE_EXT` mode) or be rejected if the
fragment is within any of the operational discard rectangles
(`VK_DISCARD_RECTANGLE_MODE_EXCLUSIVE_EXT` mode).

These discard rectangles operate orthogonally to the existing scissor
test functionality. The discard rectangles can be different for each
physical device in a device group by specifying the device mask and
setting discard rectangle dynamic state.

Version 2 of this extension introduces new dynamic states
`VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT` and
`VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT`, and the corresponding
functions
[vkCmdSetDiscardRectangleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleEnableEXT.html)
and
[vkCmdSetDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleModeEXT.html).
Applications that use these dynamic states must ensure the
implementation advertises at least `specVersion` `2` of this extension.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetDiscardRectangleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleEXT.html)

- [vkCmdSetDiscardRectangleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleEnableEXT.html)

- [vkCmdSetDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleModeEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html):

  - [VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceDiscardRectanglePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDiscardRectanglePropertiesEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDiscardRectangleModeEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkPipelineDiscardRectangleStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_DISCARD_RECTANGLES_EXTENSION_NAME`

- `VK_EXT_DISCARD_RECTANGLES_SPEC_VERSION`

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT`

  - `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT`

  - `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DISCARD_RECTANGLE_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_PIPELINE_DISCARD_RECTANGLE_STATE_CREATE_INFO_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2023-01-18 (Piers Daniell)

  - Add dynamic states for discard rectangle enable/disable and mode.

- Revision 1, 2016-12-22 (Piers Daniell)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_discard_rectangles"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
