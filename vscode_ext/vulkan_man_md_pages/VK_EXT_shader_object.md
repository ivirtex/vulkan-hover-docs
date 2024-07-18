# VK_EXT_shader_object(3) Manual Page

## Name

VK_EXT_shader_object - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

483

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

        
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Version 1.1](#versions-1.1)  
     and  
     [VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html)  
or  
[Version 1.3](#versions-1.3)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_1

- Interacts with VK_VERSION_1_3

- Interacts with VK_EXT_blend_operation_advanced

- Interacts with VK_EXT_conservative_rasterization

- Interacts with VK_EXT_depth_clip_control

- Interacts with VK_EXT_depth_clip_enable

- Interacts with VK_EXT_fragment_density_map

- Interacts with VK_EXT_line_rasterization

- Interacts with VK_EXT_mesh_shader

- Interacts with VK_EXT_provoking_vertex

- Interacts with VK_EXT_sample_locations

- Interacts with VK_EXT_subgroup_size_control

- Interacts with VK_EXT_transform_feedback

- Interacts with VK_KHR_device_group

- Interacts with VK_KHR_fragment_shading_rate

- Interacts with VK_NV_clip_space_w_scaling

- Interacts with VK_NV_coverage_reduction_mode

- Interacts with VK_NV_fragment_coverage_to_color

- Interacts with VK_NV_framebuffer_mixed_samples

- Interacts with VK_NV_mesh_shader

- Interacts with VK_NV_representative_fragment_test

- Interacts with VK_NV_shading_rate_image

- Interacts with VK_NV_viewport_swizzle

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Story <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_object%5D%20@daniel-story%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_object%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>daniel-story</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_shader_object](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_shader_object.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-03-30

**Interactions and External Dependencies**  
- Interacts with
  [`VK_EXT_extended_dynamic_state`](VK_EXT_extended_dynamic_state.html)

- Interacts with
  [`VK_EXT_extended_dynamic_state2`](VK_EXT_extended_dynamic_state2.html)

- Interacts with
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)

- Interacts with
  [`VK_EXT_vertex_input_dynamic_state`](VK_EXT_vertex_input_dynamic_state.html)

**IP Status**  
No known IP claims.

**Contributors**  
- Piers Daniell, NVIDIA

- Sandy Jamieson, Nintendo

- Žiga Markuš, LunarG

- Tobias Hector, AMD

- Alex Walters, Imagination

- Shahbaz Youssefi, Google

- Ralph Potter, Samsung

- Jan-Harald Fredriksen, ARM

- Connor Abott, Valve

- Arseny Kapoulkine, Roblox

- Patrick Doane, Activision

- Jeff Leger, Qualcomm

- Stu Smith, AMD

- Chris Glover, Google

- Ricardo Garcia, Igalia

- Faith Ekstrand, Collabora

- Timur Kristóf, Valve

- Constantine Shablya, Collabora

- Daniel Koch, NVIDIA

- Alyssa Rosenzweig, Collabora

- Mike Blumenkrantz, Valve

- Samuel Pitoiset, Valve

- Qun Lin, AMD

- Spencer Fricke, LunarG

- Soroush Faghihi Kashani, Imagination

## <a href="#_description" class="anchor"></a>Description

This extension introduces a new [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) object
type which represents a single compiled shader stage. Shader objects
provide a more flexible alternative to [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)
objects, which may be helpful in certain use cases.

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdBindShadersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindShadersEXT.html)

- [vkCmdBindVertexBuffers2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindVertexBuffers2EXT.html)

- [vkCmdSetAlphaToCoverageEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAlphaToCoverageEnableEXT.html)

- [vkCmdSetAlphaToOneEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAlphaToOneEnableEXT.html)

- [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEnableEXT.html)

- [vkCmdSetColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEquationEXT.html)

- [vkCmdSetColorWriteMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorWriteMaskEXT.html)

- [vkCmdSetCullModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCullModeEXT.html)

- [vkCmdSetDepthBiasEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBiasEnableEXT.html)

- [vkCmdSetDepthBoundsTestEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBoundsTestEnableEXT.html)

- [vkCmdSetDepthClampEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClampEnableEXT.html)

- [vkCmdSetDepthCompareOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthCompareOpEXT.html)

- [vkCmdSetDepthTestEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthTestEnableEXT.html)

- [vkCmdSetDepthWriteEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthWriteEnableEXT.html)

- [vkCmdSetFrontFaceEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFrontFaceEXT.html)

- [vkCmdSetLogicOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEXT.html)

- [vkCmdSetLogicOpEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEnableEXT.html)

- [vkCmdSetPatchControlPointsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPatchControlPointsEXT.html)

- [vkCmdSetPolygonModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPolygonModeEXT.html)

- [vkCmdSetPrimitiveRestartEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPrimitiveRestartEnableEXT.html)

- [vkCmdSetPrimitiveTopologyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPrimitiveTopologyEXT.html)

- [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)

- [vkCmdSetRasterizerDiscardEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnableEXT.html)

- [vkCmdSetSampleMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleMaskEXT.html)

- [vkCmdSetScissorWithCountEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetScissorWithCountEXT.html)

- [vkCmdSetStencilOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilOpEXT.html)

- [vkCmdSetStencilTestEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilTestEnableEXT.html)

- [vkCmdSetTessellationDomainOriginEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetTessellationDomainOriginEXT.html)

- [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetVertexInputEXT.html)

- [vkCmdSetViewportWithCountEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCountEXT.html)

- [vkCreateShadersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateShadersEXT.html)

- [vkDestroyShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyShaderEXT.html)

- [vkGetShaderBinaryDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderBinaryDataEXT.html)

If
[VK_EXT_blend_operation_advanced](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_blend_operation_advanced.html)
is supported:

- [vkCmdSetColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendAdvancedEXT.html)

If
[VK_EXT_conservative_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conservative_rasterization.html)
is supported:

- [vkCmdSetConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetConservativeRasterizationModeEXT.html)

- [vkCmdSetExtraPrimitiveOverestimationSizeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExtraPrimitiveOverestimationSizeEXT.html)

If [VK_EXT_depth_clip_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_clip_control.html) is
supported:

- [vkCmdSetDepthClipNegativeOneToOneEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClipNegativeOneToOneEXT.html)

If [VK_EXT_depth_clip_enable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_clip_enable.html) is
supported:

- [vkCmdSetDepthClipEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClipEnableEXT.html)

If [VK_EXT_line_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_line_rasterization.html) is
supported:

- [vkCmdSetLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineRasterizationModeEXT.html)

- [vkCmdSetLineStippleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineStippleEnableEXT.html)

If [VK_EXT_provoking_vertex](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_provoking_vertex.html) is supported:

- [vkCmdSetProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetProvokingVertexModeEXT.html)

If [VK_EXT_sample_locations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sample_locations.html) is supported:

- [vkCmdSetSampleLocationsEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEnableEXT.html)

If [VK_EXT_transform_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_transform_feedback.html) is
supported:

- [vkCmdSetRasterizationStreamEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationStreamEXT.html)

If [VK_NV_clip_space_w_scaling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_clip_space_w_scaling.html) is
supported:

- [vkCmdSetViewportWScalingEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWScalingEnableNV.html)

If [VK_NV_coverage_reduction_mode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_coverage_reduction_mode.html)
is supported:

- [vkCmdSetCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageReductionModeNV.html)

If
[VK_NV_fragment_coverage_to_color](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_fragment_coverage_to_color.html)
is supported:

- [vkCmdSetCoverageToColorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorEnableNV.html)

- [vkCmdSetCoverageToColorLocationNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorLocationNV.html)

If
[VK_NV_framebuffer_mixed_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_framebuffer_mixed_samples.html)
is supported:

- [vkCmdSetCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationModeNV.html)

- [vkCmdSetCoverageModulationTableEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationTableEnableNV.html)

- [vkCmdSetCoverageModulationTableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationTableNV.html)

If
[VK_NV_representative_fragment_test](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_representative_fragment_test.html)
is supported:

- [vkCmdSetRepresentativeFragmentTestEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRepresentativeFragmentTestEnableNV.html)

If [VK_NV_shading_rate_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shading_rate_image.html) is
supported:

- [vkCmdSetShadingRateImageEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetShadingRateImageEnableNV.html)

If [VK_NV_viewport_swizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_viewport_swizzle.html) is supported:

- [vkCmdSetViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportSwizzleNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendAdvancedEXT.html)

- [VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendEquationEXT.html)

- [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html)

- [VkVertexInputAttributeDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription2EXT.html)

- [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription2EXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderObjectFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderObjectFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceShaderObjectPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderObjectPropertiesEXT.html)

- Extending
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
  [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html):

  - [VkShaderRequiredSubgroupSizeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderRequiredSubgroupSizeCreateInfoEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkShaderCodeTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCodeTypeEXT.html)

- [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateFlagBitsEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkShaderCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SHADER_OBJECT_EXTENSION_NAME`

- `VK_EXT_SHADER_OBJECT_SPEC_VERSION`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_SHADER_EXT`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_INCOMPATIBLE_SHADER_BINARY_EXT`

  - `VK_INCOMPATIBLE_SHADER_BINARY_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_OBJECT_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_OBJECT_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_SHADER_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_SHADER_REQUIRED_SUBGROUP_SIZE_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_VERTEX_INPUT_ATTRIBUTE_DESCRIPTION_2_EXT`

  - `VK_STRUCTURE_TYPE_VERTEX_INPUT_BINDING_DESCRIPTION_2_EXT`

If [VK_EXT_fragment_density_map](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_fragment_density_map.html) is
supported:

- Extending [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateFlagBitsEXT.html):

  - `VK_SHADER_CREATE_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`

If [VK_EXT_mesh_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_mesh_shader.html) or
[VK_NV_mesh_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_mesh_shader.html) is supported:

- Extending [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateFlagBitsEXT.html):

  - `VK_SHADER_CREATE_NO_TASK_SHADER_BIT_EXT`

If [VK_EXT_subgroup_size_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_subgroup_size_control.html) or
[Version 1.3](#versions-1.3) is supported:

- Extending [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateFlagBitsEXT.html):

  - `VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT`

  - `VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT`

If [VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html) or [Version
1.1](#versions-1.1) is supported:

- Extending [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateFlagBitsEXT.html):

  - `VK_SHADER_CREATE_DISPATCH_BASE_BIT_EXT`

If [VK_KHR_fragment_shading_rate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shading_rate.html) is
supported:

- Extending [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateFlagBitsEXT.html):

  - `VK_SHADER_CREATE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_EXT`

## <a href="#_examples" class="anchor"></a>Examples

**Example 1**

Create linked pair of vertex and fragment shaders.

``` c
// Logical device created with the shaderObject feature enabled
VkDevice device;

// SPIR-V shader code for a vertex shader, along with its size in bytes
void* pVertexSpirv;
size_t vertexSpirvSize;

// SPIR-V shader code for a fragment shader, along with its size in bytes
void* pFragmentSpirv;
size_t fragmentSpirvSize;

// Descriptor set layout compatible with the shaders
VkDescriptorSetLayout descriptorSetLayout;

VkShaderCreateInfoEXT shaderCreateInfos[2] =
{
    {
        .sType = VK_STRUCTURE_TYPE_SHADER_CREATE_INFO_EXT,
        .pNext = NULL,
        .flags = VK_SHADER_CREATE_LINK_STAGE_BIT_EXT,
        .stage = VK_SHADER_STAGE_VERTEX_BIT,
        .nextStage = VK_SHADER_STAGE_FRAGMENT_BIT,
        .codeType = VK_SHADER_CODE_TYPE_SPIRV_EXT,
        .codeSize = vertexSpirvSize,
        .pCode = pVertexSpirv,
        .pName = "main",
        .setLayoutCount = 1,
        .pSetLayouts = &descriptorSetLayout;
        .pushConstantRangeCount = 0,
        .pPushConstantRanges = NULL,
        .pSpecializationInfo = NULL
    },
    {
        .sType = VK_STRUCTURE_TYPE_SHADER_CREATE_INFO_EXT,
        .pNext = NULL,
        .flags = VK_SHADER_CREATE_LINK_STAGE_BIT_EXT,
        .stage = VK_SHADER_STAGE_FRAGMENT_BIT,
        .nextStage = 0,
        .codeType = VK_SHADER_CODE_TYPE_SPIRV_EXT,
        .codeSize = fragmentSpirvSize,
        .pCode = pFragmentSpirv,
        .pName = "main",
        .setLayoutCount = 1,
        .pSetLayouts = &descriptorSetLayout;
        .pushConstantRangeCount = 0,
        .pPushConstantRanges = NULL,
        .pSpecializationInfo = NULL
    }
};

VkResult result;
VkShaderEXT shaders[2];

result = vkCreateShadersEXT(device, 2, &shaderCreateInfos, NULL, shaders);
if (result != VK_SUCCESS)
{
    // Handle error
}
```

Later, during command buffer recording, bind the linked shaders and
draw.

``` c
// Command buffer in the recording state
VkCommandBuffer commandBuffer;

// Vertex and fragment shader objects created above
VkShaderEXT shaders[2];

// Assume vertex buffers, descriptor sets, etc. have been bound, and existing
// state setting commands have been called to set all required state

const VkShaderStageFlagBits stages[2] =
{
    VK_SHADER_STAGE_VERTEX_BIT,
    VK_SHADER_STAGE_FRAGMENT_BIT
};

// Bind linked shaders
vkCmdBindShadersEXT(commandBuffer, 2, stages, shaders);

// Equivalent to the previous line. Linked shaders can be bound one at a time,
// in any order:
// vkCmdBindShadersEXT(commandBuffer, 1, &stages[1], &shaders[1]);
// vkCmdBindShadersEXT(commandBuffer, 1, &stages[0], &shaders[0]);

// The above is sufficient to draw if the device was created with the
// tessellationShader and geometryShader features disabled. Otherwise, since
// those stages should not execute, vkCmdBindShadersEXT() must be called at
// least once with each of their stages in pStages before drawing:

const VkShaderStageFlagBits unusedStages[3] =
{
    VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT,
    VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT,
    VK_SHADER_STAGE_GEOMETRY_BIT
};

// NULL pShaders is equivalent to an array of stageCount VK_NULL_HANDLE values,
// meaning no shaders are bound to those stages, and that any previously bound
// shaders are unbound
vkCmdBindShadersEXT(commandBuffer, 3, unusedStages, NULL);

// Graphics shader objects may only be used to draw inside dynamic render pass
// instances begun with vkCmdBeginRendering(), assume one has already been begun

// Draw a triangle
vkCmdDraw(commandBuffer, 3, 1, 0, 0);
```

**Example 2**

Create unlinked vertex, geometry, and fragment shaders.

``` c
// Logical device created with the shaderObject feature enabled
VkDevice device;

// SPIR-V shader code for vertex shaders, along with their sizes in bytes
void* pVertexSpirv[2];
size_t vertexSpirvSize[2];

// SPIR-V shader code for a geometry shader, along with its size in bytes
void pGeometrySpirv;
size_t geometrySpirvSize;

// SPIR-V shader code for fragment shaders, along with their sizes in bytes
void* pFragmentSpirv[2];
size_t fragmentSpirvSize[2];

// Descriptor set layout compatible with the shaders
VkDescriptorSetLayout descriptorSetLayout;

VkShaderCreateInfoEXT shaderCreateInfos[5] =
{
    // Stage order does not matter
    {
        .sType = VK_STRUCTURE_TYPE_SHADER_CREATE_INFO_EXT,
        .pNext = NULL,
        .flags = 0,
        .stage = VK_SHADER_STAGE_GEOMETRY_BIT,
        .nextStage = VK_SHADER_STAGE_FRAGMENT_BIT,
        .codeType = VK_SHADER_CODE_TYPE_SPIRV_EXT,
        .codeSize = pGeometrySpirv,
        .pCode = geometrySpirvSize,
        .pName = "main",
        .setLayoutCount = 1,
        .pSetLayouts = &descriptorSetLayout;
        .pushConstantRangeCount = 0,
        .pPushConstantRanges = NULL,
        .pSpecializationInfo = NULL
    },
    {
        .sType = VK_STRUCTURE_TYPE_SHADER_CREATE_INFO_EXT,
        .pNext = NULL,
        .flags = 0,
        .stage = VK_SHADER_STAGE_VERTEX_BIT,
        .nextStage = VK_SHADER_STAGE_GEOMETRY_BIT,
        .codeType = VK_SHADER_CODE_TYPE_SPIRV_EXT,
        .codeSize = vertexSpirvSize[0],
        .pCode = pVertexSpirv[0],
        .pName = "main",
        .setLayoutCount = 1,
        .pSetLayouts = &descriptorSetLayout;
        .pushConstantRangeCount = 0,
        .pPushConstantRanges = NULL,
        .pSpecializationInfo = NULL
    },
    {
        .sType = VK_STRUCTURE_TYPE_SHADER_CREATE_INFO_EXT,
        .pNext = NULL,
        .flags = 0,
        .stage = VK_SHADER_STAGE_FRAGMENT_BIT,
        .nextStage = 0,
        .codeType = VK_SHADER_CODE_TYPE_SPIRV_EXT,
        .codeSize = fragmentSpirvSize[0],
        .pCode = pFragmentSpirv[0],
        .pName = "main",
        .setLayoutCount = 1,
        .pSetLayouts = &descriptorSetLayout;
        .pushConstantRangeCount = 0,
        .pPushConstantRanges = NULL,
        .pSpecializationInfo = NULL
    },
    {
        .sType = VK_STRUCTURE_TYPE_SHADER_CREATE_INFO_EXT,
        .pNext = NULL,
        .flags = 0,
        .stage = VK_SHADER_STAGE_FRAGMENT_BIT,
        .nextStage = 0,
        .codeType = VK_SHADER_CODE_TYPE_SPIRV_EXT,
        .codeSize = fragmentSpirvSize[1],
        .pCode = pFragmentSpirv[1],
        .pName = "main",
        .setLayoutCount = 1,
        .pSetLayouts = &descriptorSetLayout;
        .pushConstantRangeCount = 0,
        .pPushConstantRanges = NULL,
        .pSpecializationInfo = NULL
    },
    {
        .sType = VK_STRUCTURE_TYPE_SHADER_CREATE_INFO_EXT,
        .pNext = NULL,
        .flags = 0,
        .stage = VK_SHADER_STAGE_VERTEX_BIT,
        // Suppose we want this vertex shader to be able to be followed by
        // either a geometry shader or fragment shader:
        .nextStage = VK_SHADER_STAGE_GEOMETRY_BIT | VK_SHADER_STAGE_FRAGMENT_BIT,
        .codeType = VK_SHADER_CODE_TYPE_SPIRV_EXT,
        .codeSize = vertexSpirvSize[1],
        .pCode = pVertexSpirv[1],
        .pName = "main",
        .setLayoutCount = 1,
        .pSetLayouts = &descriptorSetLayout;
        .pushConstantRangeCount = 0,
        .pPushConstantRanges = NULL,
        .pSpecializationInfo = NULL
    }
};

VkResult result;
VkShaderEXT shaders[5];

result = vkCreateShadersEXT(device, 5, &shaderCreateInfos, NULL, shaders);
if (result != VK_SUCCESS)
{
    // Handle error
}
```

Later, during command buffer recording, bind the linked shaders in
different combinations and draw.

``` c
// Command buffer in the recording state
VkCommandBuffer commandBuffer;

// Vertex, geometry, and fragment shader objects created above
VkShaderEXT shaders[5];

// Assume vertex buffers, descriptor sets, etc. have been bound, and existing
// state setting commands have been called to set all required state

const VkShaderStageFlagBits stages[3] =
{
    // Any order is allowed
    VK_SHADER_STAGE_FRAGMENT_BIT,
    VK_SHADER_STAGE_VERTEX_BIT,
    VK_SHADER_STAGE_GEOMETRY_BIT,
};

VkShaderEXT bindShaders[3] =
{
    shaders[2], // FS
    shaders[1], // VS
    shaders[0]  // GS
};

// Bind unlinked shaders
vkCmdBindShadersEXT(commandBuffer, 3, stages, bindShaders);

// Assume the tessellationShader feature is disabled, so vkCmdBindShadersEXT()
// need not have been called with either tessellation stage

// Graphics shader objects may only be used to draw inside dynamic render pass
// instances begun with vkCmdBeginRendering(), assume one has already been begun

// Draw a triangle
vkCmdDraw(commandBuffer, 3, 1, 0, 0);

// Bind a different unlinked fragment shader
const VkShaderStageFlagBits fragmentStage = VK_SHADER_STAGE_FRAGMENT_BIT;
vkCmdBindShadersEXT(commandBuffer, 1, &fragmentStage, &shaders[3]);

// Draw another triangle
vkCmdDraw(commandBuffer, 3, 1, 0, 0);

// Bind a different unlinked vertex shader
const VkShaderStageFlagBits vertexStage = VK_SHADER_STAGE_VERTEX_BIT;
vkCmdBindShadersEXT(commandBuffer, 1, &vertexStage, &shaders[4]);

// Draw another triangle
vkCmdDraw(commandBuffer, 3, 1, 0, 0);
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-03-30 (Daniel Story)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_shader_object"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
