# VK_EXT_conditional_rendering(3) Manual Page

## Name

VK_EXT_conditional_rendering - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

82

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Vikram Kushwaha <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_conditional_rendering%5D%20@vkushwaha%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_conditional_rendering%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>vkushwaha</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension allows the execution of one or more rendering commands to
be conditional on a value in buffer memory. This may help an application
reduce the latency by conditionally discarding rendering commands
without application intervention. The conditional rendering commands are
limited to draws, compute dispatches and clearing attachments within a
conditional rendering block.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdBeginConditionalRenderingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginConditionalRenderingEXT.html)

- [vkCmdEndConditionalRenderingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndConditionalRenderingEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkConditionalRenderingBeginInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConditionalRenderingBeginInfoEXT.html)

- Extending
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html):

  - [VkCommandBufferInheritanceConditionalRenderingInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceConditionalRenderingInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceConditionalRenderingFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceConditionalRenderingFeaturesEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkConditionalRenderingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConditionalRenderingFlagBitsEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkConditionalRenderingFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConditionalRenderingFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_CONDITIONAL_RENDERING_EXTENSION_NAME`

- `VK_EXT_CONDITIONAL_RENDERING_SPEC_VERSION`

- Extending [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html):

  - `VK_ACCESS_CONDITIONAL_RENDERING_READ_BIT_EXT`

- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html):

  - `VK_BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT`

- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html):

  - `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_CONDITIONAL_RENDERING_INFO_EXT`

  - `VK_STRUCTURE_TYPE_CONDITIONAL_RENDERING_BEGIN_INFO_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CONDITIONAL_RENDERING_FEATURES_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should conditional rendering affect copy and blit commands?

**RESOLVED**: Conditional rendering should not affect copies and blits.

2\) Should secondary command buffers be allowed to execute while
conditional rendering is active in the primary command buffer?

**RESOLVED**: The rendering commands in secondary command buffer will be
affected by an active conditional rendering in primary command buffer if
the `conditionalRenderingEnable` is set to `VK_TRUE`. Conditional
rendering **must** not be active in the primary command buffer if
`conditionalRenderingEnable` is `VK_FALSE`.

## <a href="#_examples" class="anchor"></a>Examples

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-04-19 (Vikram Kushwaha)

  - First Version

- Revision 2, 2018-05-21 (Vikram Kushwaha)

  - Add new pipeline stage, access flags and limit conditional rendering
    to a subpass or entire render pass.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_conditional_rendering"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
