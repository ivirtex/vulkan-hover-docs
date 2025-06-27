# VK_EXT_attachment_feedback_loop_dynamic_state(3) Manual Page

## Name

VK_EXT_attachment_feedback_loop_dynamic_state - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

525

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Version 1.1](#versions-1.1)  
and  
[VK_EXT_attachment_feedback_loop_layout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_attachment_feedback_loop_layout.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Mike Blumenkrantz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_attachment_feedback_loop_dynamic_state%5D%20@zmike%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_attachment_feedback_loop_dynamic_state%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>zmike</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_attachment_feedback_loop_dynamic_state](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_attachment_feedback_loop_dynamic_state.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension adds support for setting attachment feedback loops
dynamically on command buffers.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetAttachmentFeedbackLoopEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAttachmentFeedbackLoopEnableEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceAttachmentFeedbackLoopDynamicStateFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAttachmentFeedbackLoopDynamicStateFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_ATTACHMENT_FEEDBACK_LOOP_DYNAMIC_STATE_EXTENSION_NAME`

- `VK_EXT_ATTACHMENT_FEEDBACK_LOOP_DYNAMIC_STATE_SPEC_VERSION`

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_ATTACHMENT_FEEDBACK_LOOP_ENABLE_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ATTACHMENT_FEEDBACK_LOOP_DYNAMIC_STATE_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-04-28 (Mike Blumenkrantz)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_attachment_feedback_loop_dynamic_state"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
