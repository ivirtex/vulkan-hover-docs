# VK_KHR_create_renderpass2(3) Manual Page

## Name

VK_KHR_create_renderpass2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

110

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

     [VK_KHR_multiview](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_multiview.html)  
     and  
     [VK_KHR_maintenance2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_create_renderpass2%5D%20@tobias%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_create_renderpass2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobias</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-02-07

**Contributors**  
- Tobias Hector

- Jeff Bolz

## <a href="#_description" class="anchor"></a>Description

This extension provides a new command to create render passes in a way
that can be easily extended by other extensions through the
substructures of render pass creation. The Vulkan 1.0 render pass
creation sub-structures do not include `sType`/`pNext` members.
Additionally, the render pass begin/next/end commands have been
augmented with new extensible structures for passing additional subpass
information.

The
[VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)
and
[VkInputAttachmentAspectReference](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInputAttachmentAspectReference.html)
structures that extended the original
[VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) are not accepted
into the new creation functions, and instead their parameters are folded
into this extension as follows:

- Elements of
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)::`pViewMasks`
  are now specified in
  [VkSubpassDescription2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2KHR.html)::`viewMask`.

- Elements of
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)::`pViewOffsets`
  are now specified in
  [VkSubpassDependency2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency2KHR.html)::`viewOffset`.

- [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)::`correlationMaskCount`
  and
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)::`pCorrelationMasks`
  are directly specified in
  [VkRenderPassCreateInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2KHR.html).

- [VkInputAttachmentAspectReference](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInputAttachmentAspectReference.html)::`aspectMask`
  is now specified in the relevant input attachment reference in
  [VkAttachmentReference2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2KHR.html)::`aspectMask`

The details of these mappings are explained fully in the new structures.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdBeginRenderPass2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass2KHR.html)

- [vkCmdEndRenderPass2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndRenderPass2KHR.html)

- [vkCmdNextSubpass2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass2KHR.html)

- [vkCreateRenderPass2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2KHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkAttachmentDescription2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription2KHR.html)

- [VkAttachmentReference2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2KHR.html)

- [VkRenderPassCreateInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2KHR.html)

- [VkSubpassBeginInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassBeginInfoKHR.html)

- [VkSubpassDependency2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency2KHR.html)

- [VkSubpassDescription2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2KHR.html)

- [VkSubpassEndInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassEndInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_CREATE_RENDERPASS_2_EXTENSION_NAME`

- `VK_KHR_CREATE_RENDERPASS_2_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR`

  - `VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR`

  - `VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR`

  - `VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR`

  - `VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR`

  - `VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR`

  - `VK_STRUCTURE_TYPE_SUBPASS_END_INFO_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-02-07 (Tobias Hector)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_create_renderpass2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
