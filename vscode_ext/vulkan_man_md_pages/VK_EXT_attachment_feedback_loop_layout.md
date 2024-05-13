# VK_EXT_attachment_feedback_loop_layout(3) Manual Page

## Name

VK_EXT_attachment_feedback_loop_layout - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

340

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Joshua Ashton <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_attachment_feedback_loop_layout%5D%20@Joshua-Ashton%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_attachment_feedback_loop_layout%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>Joshua-Ashton</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_attachment_feedback_loop_layout](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_attachment_feedback_loop_layout.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-04-04

**IP Status**  
No known IP claims.

**Contributors**  
- Joshua Ashton, Valve

- Faith Ekstrand, Collabora

- Bas Nieuwenhuizen, Google

- Samuel Iglesias Gonsálvez, Igalia

- Ralph Potter, Samsung

- Jan-Harald Fredriksen, Arm

- Ricardo Garcia, Igalia

## <a href="#_description" class="anchor"></a>Description

This extension adds a new image layout,
`VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`, which allows
applications to have an image layout in which they are able to both
render to and sample/fetch from the same subresource of an image in a
given render pass.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_ATTACHMENT_FEEDBACK_LOOP_LAYOUT_EXTENSION_NAME`

- `VK_EXT_ATTACHMENT_FEEDBACK_LOOP_LAYOUT_SPEC_VERSION`

- Extending [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html):

  - `VK_DEPENDENCY_FEEDBACK_LOOP_BIT_EXT`

- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html):

  - `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`

- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html):

  - `VK_IMAGE_USAGE_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html):

  - `VK_PIPELINE_CREATE_COLOR_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`

  - `VK_PIPELINE_CREATE_DEPTH_STENCIL_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ATTACHMENT_FEEDBACK_LOOP_LAYOUT_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2022-04-04 (Joshua Ashton)

  - Renamed from VALVE to EXT.

- Revision 1, 2021-03-09 (Joshua Ashton)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_attachment_feedback_loop_layout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
