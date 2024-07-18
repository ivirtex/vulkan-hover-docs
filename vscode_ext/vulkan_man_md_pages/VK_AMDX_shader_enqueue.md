# VK_AMDX_shader_enqueue(3) Manual Page

## Name

VK_AMDX_shader_enqueue - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

135

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

            
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
             or  
             [Version 1.1](#versions-1.1)  
         and  
         [VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html)  
     or  
     [Version 1.3](#versions-1.3)  
and  
[VK_KHR_pipeline_library](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_pipeline_library.html)  
and  
[VK_KHR_spirv_1_4](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_spirv_1_4.html)  

- **This is a *provisional* extension and **must** be used with caution.
  See the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#boilerplate-provisional-header"
  target="_blank" rel="noopener">description</a> of provisional header
  files for enablement and stability details.**

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_KHR_maintenance5

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_AMDX_shader_enqueue](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/AMDX/SPV_AMDX_shader_enqueue.html)

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMDX_shader_enqueue%5D%20@tobski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMDX_shader_enqueue%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobski</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_AMDX_shader_enqueue](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_AMDX_shader_enqueue.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-07-22

**Provisional**  
**This extension is *provisional* and **should** not be used in
production applications. The functionality **may** change in ways that
break backwards compatibility between revisions, and before final
release.**

**Contributors**  
- Tobias Hector, AMD

- Matthaeus Chajdas, AMD

- Maciej Jesionowski, AMD

- Robert Martin, AMD

- Qun Lin, AMD

- Rex Xu, AMD

- Dominik Witczak, AMD

- Karthik Srinivasan, AMD

- Nicolai Haehnle, AMD

- Stuart Smith, AMD

## <a href="#_description" class="anchor"></a>Description

This extension adds the ability for developers to enqueue compute shader
workgroups from other compute shaders.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdDispatchGraphAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchGraphAMDX.html)

- [vkCmdDispatchGraphIndirectAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchGraphIndirectAMDX.html)

- [vkCmdDispatchGraphIndirectCountAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchGraphIndirectCountAMDX.html)

- [vkCmdInitializeGraphScratchMemoryAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdInitializeGraphScratchMemoryAMDX.html)

- [vkCreateExecutionGraphPipelinesAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateExecutionGraphPipelinesAMDX.html)

- [vkGetExecutionGraphPipelineNodeIndexAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetExecutionGraphPipelineNodeIndexAMDX.html)

- [vkGetExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetExecutionGraphPipelineScratchSizeAMDX.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDispatchGraphCountInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDispatchGraphCountInfoAMDX.html)

- [VkDispatchGraphInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDispatchGraphInfoAMDX.html)

- [VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html)

- [VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderEnqueueFeaturesAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderEnqueueFeaturesAMDX.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceShaderEnqueuePropertiesAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderEnqueuePropertiesAMDX.html)

- Extending
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html):

  - [VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html)

## <a href="#_new_unions" class="anchor"></a>New Unions

- [VkDeviceOrHostAddressConstAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstAMDX.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMDX_SHADER_ENQUEUE_EXTENSION_NAME`

- `VK_AMDX_SHADER_ENQUEUE_SPEC_VERSION`

- `VK_SHADER_INDEX_UNUSED_AMDX`

- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html):

  - `VK_BUFFER_USAGE_EXECUTION_GRAPH_SCRATCH_BIT_AMDX`

- Extending [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html):

  - `VK_PIPELINE_BIND_POINT_EXECUTION_GRAPH_AMDX`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_EXECUTION_GRAPH_PIPELINE_CREATE_INFO_AMDX`

  - `VK_STRUCTURE_TYPE_EXECUTION_GRAPH_PIPELINE_SCRATCH_SIZE_AMDX`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ENQUEUE_FEATURES_AMDX`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ENQUEUE_PROPERTIES_AMDX`

  - `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_NODE_CREATE_INFO_AMDX`

If [VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html) is supported:

- Extending [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html):

  - `VK_BUFFER_USAGE_2_EXECUTION_GRAPH_SCRATCH_BIT_AMDX`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-07-22 (Tobias Hector)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMDX_shader_enqueue"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
