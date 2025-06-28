# VK\_EXT\_subpass\_merge\_feedback(3) Manual Page

## Name

VK\_EXT\_subpass\_merge\_feedback - device extension



## [](#_registered_extension_number)Registered Extension Number

459

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Ting Wei [\[GitHub\]catweiting](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_subpass_merge_feedback%5D%20%40catweiting%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_subpass_merge_feedback%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_subpass\_merge\_feedback](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_subpass_merge_feedback.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-05-24

**IP Status**

No known IP claims.

**Contributors**

- Jan-Harald Fredriksen, Arm
- Jorg Wagner, Arm
- Ting Wei, Arm

## [](#_description)Description

This extension adds a mechanism to provide feedback to an application about whether the subpasses specified on render pass creation are merged by the implementation. Additionally, it provides a control to enable or disable subpass merging in the render pass.

## [](#_new_structures)New Structures

- [VkRenderPassCreationFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreationFeedbackInfoEXT.html)
- [VkRenderPassSubpassFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSubpassFeedbackInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceSubpassMergeFeedbackFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubpassMergeFeedbackFeaturesEXT.html)
- Extending [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html):
  
  - [VkRenderPassCreationFeedbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreationFeedbackCreateInfoEXT.html)
- Extending [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html), [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html):
  
  - [VkRenderPassCreationControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreationControlEXT.html)
- Extending [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html):
  
  - [VkRenderPassSubpassFeedbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSubpassFeedbackCreateInfoEXT.html)

## [](#_new_enums)New Enums

- [VkSubpassMergeStatusEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassMergeStatusEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SUBPASS_MERGE_FEEDBACK_EXTENSION_NAME`
- `VK_EXT_SUBPASS_MERGE_FEEDBACK_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBPASS_MERGE_FEEDBACK_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_CREATION_CONTROL_EXT`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_CREATION_FEEDBACK_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_SUBPASS_FEEDBACK_CREATE_INFO_EXT`

## [](#_version_history)Version History

- Revision 1, 2022-03-10
  
  - Initial draft.
- Revision 2, 2022-05-24
  
  - Fix structextends and constness issues.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_subpass_merge_feedback)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0