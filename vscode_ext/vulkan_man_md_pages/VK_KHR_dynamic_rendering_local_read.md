# VK_KHR_dynamic_rendering_local_read(3) Manual Page

## Name

VK_KHR_dynamic_rendering_local_read - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

233

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html)  
or  
[Version 1.3](#versions-1.3)  

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_dynamic_rendering_local_read%5D%20@tobski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_dynamic_rendering_local_read%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobski</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_dynamic_rendering_local_read](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_dynamic_rendering_local_read.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension enables reads from attachments and resources written by
previous fragment shaders within a dynamic render pass.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetRenderingAttachmentLocationsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRenderingAttachmentLocationsKHR.html)

- [vkCmdSetRenderingInputAttachmentIndicesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRenderingInputAttachmentIndicesKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html):

  - [VkRenderingAttachmentLocationInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentLocationInfoKHR.html)

  - [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_DYNAMIC_RENDERING_LOCAL_READ_EXTENSION_NAME`

- `VK_KHR_DYNAMIC_RENDERING_LOCAL_READ_SPEC_VERSION`

- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html):

  - `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_LOCAL_READ_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_LOCATION_INFO_KHR`

  - `VK_STRUCTURE_TYPE_RENDERING_INPUT_ATTACHMENT_INDEX_INFO_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-11-03 (Tobias Hector)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_dynamic_rendering_local_read"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
