# VK\_KHR\_create\_renderpass2(3) Manual Page

## Name

VK\_KHR\_create\_renderpass2 - device extension



## [](#_registered_extension_number)Registered Extension Number

110

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_multiview](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_multiview.html)  
     and  
     [VK\_KHR\_maintenance2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobias](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_create_renderpass2%5D%20%40tobias%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_create_renderpass2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-02-07

**Contributors**

- Tobias Hector
- Jeff Bolz

## [](#_description)Description

This extension provides a new command to create render passes in a way that can be easily extended by other extensions through the substructures of render pass creation. The Vulkan 1.0 render pass creation sub-structures do not include `sType`/`pNext` members. Additionally, the render pass begin/next/end commands have been augmented with new extensible structures for passing additional subpass information.

The [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html) and [VkInputAttachmentAspectReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInputAttachmentAspectReference.html) structures that extended the original [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html) are not accepted into the new creation functions, and instead their parameters are folded into this extension as follows:

- Elements of [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html)::`pViewMasks` are now specified in [VkSubpassDescription2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2KHR.html)::`viewMask`.
- Elements of [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html)::`pViewOffsets` are now specified in [VkSubpassDependency2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency2KHR.html)::`viewOffset`.
- [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html)::`correlationMaskCount` and [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html)::`pCorrelationMasks` are directly specified in [VkRenderPassCreateInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2KHR.html).
- [VkInputAttachmentAspectReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInputAttachmentAspectReference.html)::`aspectMask` is now specified in the relevant input attachment reference in [VkAttachmentReference2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2KHR.html)::`aspectMask`

The details of these mappings are explained fully in the new structures.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_commands)New Commands

- [vkCmdBeginRenderPass2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRenderPass2KHR.html)
- [vkCmdEndRenderPass2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndRenderPass2KHR.html)
- [vkCmdNextSubpass2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdNextSubpass2KHR.html)
- [vkCreateRenderPass2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRenderPass2KHR.html)

## [](#_new_structures)New Structures

- [VkAttachmentDescription2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription2KHR.html)
- [VkAttachmentReference2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2KHR.html)
- [VkRenderPassCreateInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2KHR.html)
- [VkSubpassBeginInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassBeginInfoKHR.html)
- [VkSubpassDependency2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency2KHR.html)
- [VkSubpassDescription2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2KHR.html)
- [VkSubpassEndInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassEndInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_CREATE_RENDERPASS_2_EXTENSION_NAME`
- `VK_KHR_CREATE_RENDERPASS_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR`
  - `VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR`
  - `VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR`
  - `VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR`
  - `VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR`
  - `VK_STRUCTURE_TYPE_SUBPASS_END_INFO_KHR`

## [](#_version_history)Version History

- Revision 1, 2018-02-07 (Tobias Hector)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_create_renderpass2)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0