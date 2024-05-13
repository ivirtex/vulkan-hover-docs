# VK_KHR_dynamic_rendering(3) Manual Page

## Name

VK_KHR_dynamic_rendering - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

45

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

        
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Version 1.1](#versions-1.1)  
     and  
     [VK_KHR_depth_stencil_resolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_depth_stencil_resolve.html)  
or  
[Version 1.2](#versions-1.2)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_AMD_mixed_attachment_samples

- Interacts with VK_EXT_fragment_density_map

- Interacts with VK_KHR_fragment_shading_rate

- Interacts with VK_NVX_multiview_per_view_attributes

- Interacts with VK_NV_framebuffer_mixed_samples

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_dynamic_rendering%5D%20@tobski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_dynamic_rendering%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobski</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_dynamic_rendering](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_dynamic_rendering.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension allows applications to create single-pass render pass
instances without needing to create render pass objects or framebuffers.
Dynamic render passes can also span across multiple primary command
buffers, rather than relying on secondary command buffers.

This extension also incorporates `VK_ATTACHMENT_STORE_OP_NONE_KHR` from
[`VK_QCOM_render_pass_store_ops`](VK_QCOM_render_pass_store_ops.html),
enabling applications to avoid unnecessary synchronization when an
attachment is not written during a render pass.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdBeginRenderingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderingKHR.html)

- [vkCmdEndRenderingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndRenderingKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkRenderingAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentInfoKHR.html)

- [VkRenderingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfoKHR.html)

- Extending
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html):

  - [VkCommandBufferInheritanceRenderingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfoKHR.html)

- Extending
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html):

  - [VkPipelineRenderingCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfoKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceDynamicRenderingFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDynamicRenderingFeaturesKHR.html)

If
[VK_AMD_mixed_attachment_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_mixed_attachment_samples.html)
is supported:

- Extending
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html),
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html):

  - [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)

If [VK_EXT_fragment_density_map](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_fragment_density_map.html) is
supported:

- Extending [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html):

  - [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)

If [VK_KHR_fragment_shading_rate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shading_rate.html) is
supported:

- Extending [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html):

  - [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)

If
[VK_NV_framebuffer_mixed_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_framebuffer_mixed_samples.html)
is supported:

- Extending
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html),
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html):

  - [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)

If
[VK_NVX_multiview_per_view_attributes](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NVX_multiview_per_view_attributes.html)
is supported:

- Extending
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html),
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html):

  - [VkMultiviewPerViewAttributesInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewAttributesInfoNVX.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkRenderingFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkRenderingFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_DYNAMIC_RENDERING_EXTENSION_NAME`

- `VK_KHR_DYNAMIC_RENDERING_SPEC_VERSION`

- Extending [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html):

  - `VK_ATTACHMENT_STORE_OP_NONE_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_RENDERING_INFO_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PIPELINE_RENDERING_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_INFO_KHR`

  - `VK_STRUCTURE_TYPE_RENDERING_INFO_KHR`

If
[VK_AMD_mixed_attachment_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_mixed_attachment_samples.html)
is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_ATTACHMENT_SAMPLE_COUNT_INFO_AMD`

If [VK_EXT_fragment_density_map](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_fragment_density_map.html) is
supported:

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html):

  - `VK_PIPELINE_CREATE_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`

  - `VK_PIPELINE_RASTERIZATION_STATE_CREATE_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_INFO_EXT`

If [VK_KHR_fragment_shading_rate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shading_rate.html) is
supported:

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html):

  - `VK_PIPELINE_CREATE_RENDERING_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

  - `VK_PIPELINE_RASTERIZATION_STATE_CREATE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_RENDERING_FRAGMENT_SHADING_RATE_ATTACHMENT_INFO_KHR`

If
[VK_NV_framebuffer_mixed_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_framebuffer_mixed_samples.html)
is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_ATTACHMENT_SAMPLE_COUNT_INFO_NV`

If
[VK_NVX_multiview_per_view_attributes](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NVX_multiview_per_view_attributes.html)
is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_MULTIVIEW_PER_VIEW_ATTRIBUTES_INFO_NVX`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

Functionality in this extension is included in core Vulkan 1.3, with the
KHR suffix omitted. The original type, enum and command names are still
available as aliases of the core functionality.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-10-06 (Tobias Hector)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_dynamic_rendering"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
