# VK_EXT_rasterization_order_attachment_access(3) Manual Page

## Name

VK_EXT_rasterization_order_attachment_access - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

464

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Jan-Harald Fredriksen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_rasterization_order_attachment_access%5D%20@janharaldfredriksen-arm%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_rasterization_order_attachment_access%20extension*"
  target="_blank"
  rel="nofollow noopener"><em></em>janharaldfredriksen-arm</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_rasterization_order_attachment_access](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_rasterization_order_attachment_access.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-07-04

**IP Status**  
No known IP claims.

**Contributors**  
- Tobias Hector, AMD

- Jan-Harald Fredriksen, Arm

## <a href="#_description" class="anchor"></a>Description

This extension extends the mechanism of input attachments to allow
access to framebuffer attachments that are used both as input and as
color or depth/stencil attachments from one fragment to the next, in
rasterization order, without explicit synchronization.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkPipelineColorBlendStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateFlagBits.html)

- [VkPipelineDepthStencilStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateFlagBits.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_EXTENSION_NAME`

- `VK_EXT_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_SPEC_VERSION`

- Extending
  [VkPipelineColorBlendStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateFlagBits.html):

  - `VK_PIPELINE_COLOR_BLEND_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_BIT_EXT`

- Extending
  [VkPipelineDepthStencilStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateFlagBits.html):

  - `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_EXT`

  - `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_FEATURES_EXT`

- Extending
  [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionFlagBits.html):

  - `VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_COLOR_ACCESS_BIT_EXT`

  - `VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_EXT`

  - `VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_EXT`

## <a href="#_examples" class="anchor"></a>Examples

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-07-04 (Jan-Harald Fredriksen)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_rasterization_order_attachment_access"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
