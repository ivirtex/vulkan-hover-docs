# VK\_EXT\_attachment\_feedback\_loop\_dynamic\_state(3) Manual Page

## Name

VK\_EXT\_attachment\_feedback\_loop\_dynamic\_state - device extension



## [](#_registered_extension_number)Registered Extension Number

525

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)  
and  
[VK\_EXT\_attachment\_feedback\_loop\_layout](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_attachment_feedback_loop_layout.html)

## [](#_special_uses)Special Uses

- [OpenGL / ES support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)
- [D3D support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Mike Blumenkrantz [\[GitHub\]zmike](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_attachment_feedback_loop_dynamic_state%5D%20%40zmike%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_attachment_feedback_loop_dynamic_state%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_attachment\_feedback\_loop\_dynamic\_state](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_attachment_feedback_loop_dynamic_state.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-04-28

**IP Status**

No known IP claims.

**Contributors**

- Mike Blumenkrantz, Valve
- Daniel Story, Nintendo
- Stu Smith, AMD
- Samuel Pitoiset, Valve
- Ricardo Garcia, Igalia

## [](#_description)Description

This extension adds support for setting attachment feedback loops dynamically on command buffers.

## [](#_new_commands)New Commands

- [vkCmdSetAttachmentFeedbackLoopEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetAttachmentFeedbackLoopEnableEXT.html)

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceAttachmentFeedbackLoopDynamicStateFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceAttachmentFeedbackLoopDynamicStateFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_ATTACHMENT_FEEDBACK_LOOP_DYNAMIC_STATE_EXTENSION_NAME`
- `VK_EXT_ATTACHMENT_FEEDBACK_LOOP_DYNAMIC_STATE_SPEC_VERSION`
- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_ATTACHMENT_FEEDBACK_LOOP_ENABLE_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ATTACHMENT_FEEDBACK_LOOP_DYNAMIC_STATE_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2023-04-28 (Mike Blumenkrantz)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_attachment_feedback_loop_dynamic_state).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0