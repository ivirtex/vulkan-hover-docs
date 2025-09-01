# VK\_ARM\_rasterization\_order\_attachment\_access(3) Manual Page

## Name

VK\_ARM\_rasterization\_order\_attachment\_access - device extension



## [](#_registered_extension_number)Registered Extension Number

343

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_EXT\_rasterization\_order\_attachment\_access](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_rasterization_order_attachment_access.html) extension

## [](#_contact)Contact

- Jan-Harald Fredriksen [\[GitHub\]janharaldfredriksen-arm](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ARM_rasterization_order_attachment_access%5D%20%40janharaldfredriksen-arm%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ARM_rasterization_order_attachment_access%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-11-12

**IP Status**

No known IP claims.

**Contributors**

- Tobias Hector, AMD
- Jan-Harald Fredriksen, Arm

## [](#_description)Description

Render passes, and specifically subpass dependencies, enable much of the same functionality as the framebuffer fetch and pixel local storage extensions did for OpenGL ES. But certain techniques such as programmable blending are awkward or impractical to implement with these alone, in part because a self-dependency is required every time a fragment will read a value at a given sample coordinate.

This extension extends the mechanism of input attachments to allow access to framebuffer attachments when used as both input and color, or depth/stencil, attachments from one fragment to the next, in rasterization order, without explicit synchronization.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesARM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_ARM_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_EXTENSION_NAME`
- `VK_ARM_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_SPEC_VERSION`
- Extending [VkPipelineColorBlendStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendStateCreateFlagBits.html):
  
  - `VK_PIPELINE_COLOR_BLEND_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_BIT_ARM`
- Extending [VkPipelineDepthStencilStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateFlagBits.html):
  
  - `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_ARM`
  - `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_ARM`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_FEATURES_ARM`
- Extending [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionFlagBits.html):
  
  - `VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_COLOR_ACCESS_BIT_ARM`
  - `VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_ARM`
  - `VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_ARM`

## [](#_issues)Issues

1\) Is there any interaction with the `VK_KHR_dynamic_rendering` extension?

No. This extension only affects reads from input attachments. Render pass instances begun with [vkCmdBeginRenderingKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRenderingKHR.html) do not have input attachments and a different mechanism will be needed to provide similar functionality in this case.

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2021-11-12 (Jan-Harald Fredriksen)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_ARM_rasterization_order_attachment_access).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0