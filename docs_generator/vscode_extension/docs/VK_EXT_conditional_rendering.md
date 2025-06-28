# VK\_EXT\_conditional\_rendering(3) Manual Page

## Name

VK\_EXT\_conditional\_rendering - device extension



## [](#_registered_extension_number)Registered Extension Number

82

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Vikram Kushwaha [\[GitHub\]vkushwaha](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_conditional_rendering%5D%20%40vkushwaha%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_conditional_rendering%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-05-21

**IP Status**

No known IP claims.

**Contributors**

- Vikram Kushwaha, NVIDIA
- Daniel Rakos, AMD
- Jesse Hall, Google
- Jeff Bolz, NVIDIA
- Piers Daniell, NVIDIA
- Stuart Smith, Imagination Technologies

## [](#_description)Description

This extension allows the execution of one or more rendering commands to be conditional on a value in buffer memory. This may help an application reduce the latency by conditionally discarding rendering commands without application intervention. The conditional rendering commands are limited to draws, compute dispatches and clearing attachments within a conditional rendering block.

## [](#_new_commands)New Commands

- [vkCmdBeginConditionalRenderingEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginConditionalRenderingEXT.html)
- [vkCmdEndConditionalRenderingEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndConditionalRenderingEXT.html)

## [](#_new_structures)New Structures

- [VkConditionalRenderingBeginInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConditionalRenderingBeginInfoEXT.html)
- Extending [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html):
  
  - [VkCommandBufferInheritanceConditionalRenderingInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceConditionalRenderingInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceConditionalRenderingFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceConditionalRenderingFeaturesEXT.html)

## [](#_new_enums)New Enums

- [VkConditionalRenderingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConditionalRenderingFlagBitsEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkConditionalRenderingFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConditionalRenderingFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_CONDITIONAL_RENDERING_EXTENSION_NAME`
- `VK_EXT_CONDITIONAL_RENDERING_SPEC_VERSION`
- Extending [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html):
  
  - `VK_ACCESS_CONDITIONAL_RENDERING_READ_BIT_EXT`
- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html):
  
  - `VK_BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT`
- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html):
  
  - `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_CONDITIONAL_RENDERING_INFO_EXT`
  - `VK_STRUCTURE_TYPE_CONDITIONAL_RENDERING_BEGIN_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CONDITIONAL_RENDERING_FEATURES_EXT`

## [](#_issues)Issues

1\) Should conditional rendering affect copy and blit commands?

**RESOLVED**: Conditional rendering should not affect copies and blits.

2\) Should secondary command buffers be allowed to execute while conditional rendering is active in the primary command buffer?

**RESOLVED**: The rendering commands in secondary command buffer will be affected by an active conditional rendering in primary command buffer if the `conditionalRenderingEnable` is set to `VK_TRUE`. Conditional rendering **must** not be active in the primary command buffer if `conditionalRenderingEnable` is `VK_FALSE`.

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2018-04-19 (Vikram Kushwaha)
  
  - First Version
- Revision 2, 2018-05-21 (Vikram Kushwaha)
  
  - Add new pipeline stage, access flags and limit conditional rendering to a subpass or entire render pass.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_conditional_rendering)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0