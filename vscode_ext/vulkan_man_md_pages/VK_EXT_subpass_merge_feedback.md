# VK_EXT_subpass_merge_feedback(3) Manual Page

## Name

VK_EXT_subpass_merge_feedback - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

459

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Ting Wei <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_subpass_merge_feedback%5D%20@catweiting%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_subpass_merge_feedback%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>catweiting</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_subpass_merge_feedback](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_subpass_merge_feedback.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-05-24

**IP Status**  
No known IP claims.

**Contributors**  
- Jan-Harald Fredriksen, Arm

- Jorg Wagner, Arm

- Ting Wei, Arm

## <a href="#_description" class="anchor"></a>Description

This extension adds a mechanism to provide feedback to an application
about whether the subpasses specified on render pass creation are merged
by the implementation. Additionally, it provides a control to enable or
disable subpass merging in the render pass.

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkRenderPassCreationFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreationFeedbackInfoEXT.html)

- [VkRenderPassSubpassFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSubpassFeedbackInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceSubpassMergeFeedbackFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubpassMergeFeedbackFeaturesEXT.html)

- Extending [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html):

  - [VkRenderPassCreationFeedbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreationFeedbackCreateInfoEXT.html)

- Extending [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html),
  [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html):

  - [VkRenderPassCreationControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreationControlEXT.html)

- Extending [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html):

  - [VkRenderPassSubpassFeedbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSubpassFeedbackCreateInfoEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkSubpassMergeStatusEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassMergeStatusEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SUBPASS_MERGE_FEEDBACK_EXTENSION_NAME`

- `VK_EXT_SUBPASS_MERGE_FEEDBACK_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBPASS_MERGE_FEEDBACK_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_RENDER_PASS_CREATION_CONTROL_EXT`

  - `VK_STRUCTURE_TYPE_RENDER_PASS_CREATION_FEEDBACK_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_RENDER_PASS_SUBPASS_FEEDBACK_CREATE_INFO_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-03-10

  - Initial draft.

- Revision 2, 2022-05-24

  - Fix structextends and constness issues.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_subpass_merge_feedback"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
