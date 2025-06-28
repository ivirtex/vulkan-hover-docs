# VK\_KHR\_dynamic\_rendering(3) Manual Page

## Name

VK\_KHR\_dynamic\_rendering - device extension



## [](#_registered_extension_number)Registered Extension Number

45

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Vulkan Version 1.1](#versions-1.1)  
     and  
     [VK\_KHR\_depth\_stencil\_resolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_depth_stencil_resolve.html)  
or  
[Vulkan Version 1.2](#versions-1.2)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_dynamic_rendering%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_dynamic_rendering%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_dynamic\_rendering](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_dynamic_rendering.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-10-06

**Contributors**

- Tobias Hector, AMD
- Arseny Kapoulkine, Roblox
- François Duranleau, Gameloft
- Stuart Smith, AMD
- Hai Nguyen, Google
- Jean-François Roy, Google
- Jeff Leger, Qualcomm
- Jan-Harald Fredriksen, Arm
- Piers Daniell, Nvidia
- James Fitzpatrick, Imagination
- Piotr Byszewski, Mobica
- Jesse Hall, Google
- Mike Blumenkrantz, Valve

## [](#_description)Description

This extension allows applications to create single-pass render pass instances without needing to create render pass objects or framebuffers. Dynamic render passes can also span across multiple primary command buffers, rather than relying on secondary command buffers.

This extension also incorporates `VK_ATTACHMENT_STORE_OP_NONE_KHR` from `VK_QCOM_render_pass_store_ops`, enabling applications to avoid unnecessary synchronization when an attachment is not written during a render pass.

## [](#_new_commands)New Commands

- [vkCmdBeginRenderingKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRenderingKHR.html)
- [vkCmdEndRenderingKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndRenderingKHR.html)

## [](#_new_structures)New Structures

- [VkRenderingAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentInfoKHR.html)
- [VkRenderingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfoKHR.html)
- Extending [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html):
  
  - [VkCommandBufferInheritanceRenderingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceRenderingInfoKHR.html)
- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html):
  
  - [VkPipelineRenderingCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRenderingCreateInfoKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDynamicRenderingFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDynamicRenderingFeaturesKHR.html)

## [](#_new_enums)New Enums

- [VkRenderingFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkRenderingFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_DYNAMIC_RENDERING_EXTENSION_NAME`
- `VK_KHR_DYNAMIC_RENDERING_SPEC_VERSION`
- Extending [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html):
  
  - `VK_ATTACHMENT_STORE_OP_NONE_KHR`
- Extending [VkRenderingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFlagBits.html):
  
  - `VK_RENDERING_CONTENTS_SECONDARY_COMMAND_BUFFERS_BIT_KHR`
  - `VK_RENDERING_RESUMING_BIT_KHR`
  - `VK_RENDERING_SUSPENDING_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_RENDERING_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PIPELINE_RENDERING_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_RENDERING_INFO_KHR`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the KHR suffix omitted. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2021-10-06 (Tobias Hector)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_dynamic_rendering)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0