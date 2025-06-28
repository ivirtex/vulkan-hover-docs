# VK\_KHR\_dynamic\_rendering\_local\_read(3) Manual Page

## Name

VK\_KHR\_dynamic\_rendering\_local\_read - device extension



## [](#_registered_extension_number)Registered Extension Number

233

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_dynamic_rendering_local_read%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_dynamic_rendering_local_read%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_dynamic\_rendering\_local\_read](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_dynamic_rendering_local_read.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-11-03

**Contributors**

- Tobias Hector, AMD
- Hans-Kristian Arntzen, Valve
- Connor Abbott, Valve
- Pan Gao, Huawei
- Lionel Landwerlin, Intel
- Shahbaz Youssefi, Google
- Alyssa Rosenzweig, Valve
- Jan-Harald Fredriksen, Arm
- Mike Blumenkrantz, Valve
- Graeme Leese, Broadcom
- Piers Daniell, Nvidia
- Stuart Smith, AMD
- Daniel Story, Nintendo
- James Fitzpatrick, Imagination
- Piotr Byszewski, Mobica
- Spencer Fricke, LunarG
- Tom Olson, Arm
- Michal Pietrasiuk, Intel
- Matthew Netsch, Qualcomm
- Marty Johnson, Khronos
- Wyvern Wang, Huawei
- Jeff Bolz, Nvidia
- Samuel (Sheng-Wen) Huang, MediaTek

## [](#_description)Description

This extension enables reads from attachments and resources written by previous fragment shaders within a dynamic render pass.

## [](#_new_commands)New Commands

- [vkCmdSetRenderingAttachmentLocationsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRenderingAttachmentLocationsKHR.html)
- [vkCmdSetRenderingInputAttachmentIndicesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRenderingInputAttachmentIndicesKHR.html)

## [](#_new_structures)New Structures

- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html):
  
  - [VkRenderingAttachmentLocationInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentLocationInfoKHR.html)
  - [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_DYNAMIC_RENDERING_LOCAL_READ_EXTENSION_NAME`
- `VK_KHR_DYNAMIC_RENDERING_LOCAL_READ_SPEC_VERSION`
- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html):
  
  - `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_LOCAL_READ_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_LOCATION_INFO_KHR`
  - `VK_STRUCTURE_TYPE_RENDERING_INPUT_ATTACHMENT_INDEX_INFO_KHR`

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4, with the KHR suffix omitted. However, Vulkan 1.4 implementations only have to support local read for storage resources and single sampled color attachments.

Support for reading depth/stencil attachments and multi-sampled attachments are respectively gated behind the new boolean `dynamicRenderingLocalReadDepthStencilAttachments` and `dynamicRenderingLocalReadMultisampledAttachments` properties, as described in the [Version 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4) appendix.

The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2023-11-03 (Tobias Hector)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_dynamic_rendering_local_read)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0