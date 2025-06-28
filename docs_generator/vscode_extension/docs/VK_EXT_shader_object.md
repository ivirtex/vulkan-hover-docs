# VK\_EXT\_shader\_object(3) Manual Page

## Name

VK\_EXT\_shader\_object - device extension



## [](#_registered_extension_number)Registered Extension Number

483

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Vulkan Version 1.1](#versions-1.1)  
     and  
     [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_1
- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_EXT\_blend\_operation\_advanced
- Interacts with VK\_EXT\_conservative\_rasterization
- Interacts with VK\_EXT\_depth\_clamp\_control
- Interacts with VK\_EXT\_depth\_clip\_control
- Interacts with VK\_EXT\_depth\_clip\_enable
- Interacts with VK\_EXT\_fragment\_density\_map
- Interacts with VK\_EXT\_line\_rasterization
- Interacts with VK\_EXT\_mesh\_shader
- Interacts with VK\_EXT\_provoking\_vertex
- Interacts with VK\_EXT\_sample\_locations
- Interacts with VK\_EXT\_subgroup\_size\_control
- Interacts with VK\_EXT\_transform\_feedback
- Interacts with VK\_KHR\_device\_group
- Interacts with VK\_KHR\_fragment\_shading\_rate
- Interacts with VK\_NV\_clip\_space\_w\_scaling
- Interacts with VK\_NV\_coverage\_reduction\_mode
- Interacts with VK\_NV\_fragment\_coverage\_to\_color
- Interacts with VK\_NV\_framebuffer\_mixed\_samples
- Interacts with VK\_NV\_mesh\_shader
- Interacts with VK\_NV\_representative\_fragment\_test
- Interacts with VK\_NV\_shading\_rate\_image
- Interacts with VK\_NV\_viewport\_swizzle

## [](#_contact)Contact

- Daniel Story [\[GitHub\]daniel-story](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_object%5D%20%40daniel-story%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_object%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_shader\_object](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_shader_object.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-03-30

**Interactions and External Dependencies**

- Interacts with `VK_EXT_extended_dynamic_state`
- Interacts with `VK_EXT_extended_dynamic_state2`
- Interacts with `VK_EXT_extended_dynamic_state3`
- Interacts with `VK_EXT_vertex_input_dynamic_state`

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
- Caterina Shablia, Collabora
- Daniel Koch, NVIDIA
- Alyssa Rosenzweig, Collabora
- Mike Blumenkrantz, Valve
- Samuel Pitoiset, Valve
- Qun Lin, AMD
- Spencer Fricke, LunarG
- Soroush Faghihi Kashani, Imagination

## [](#_description)Description

This extension introduces a new [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) object type which represents a single compiled shader stage. Shader objects provide a more flexible alternative to [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) objects, which may be helpful in certain use cases.

## [](#_new_object_types)New Object Types

- [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html)

## [](#_new_commands)New Commands

- [vkCmdBindShadersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindShadersEXT.html)
- [vkCmdBindVertexBuffers2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindVertexBuffers2EXT.html)
- [vkCmdSetAlphaToCoverageEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetAlphaToCoverageEnableEXT.html)
- [vkCmdSetAlphaToOneEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetAlphaToOneEnableEXT.html)
- [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetColorBlendEnableEXT.html)
- [vkCmdSetColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetColorBlendEquationEXT.html)
- [vkCmdSetColorWriteMaskEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetColorWriteMaskEXT.html)
- [vkCmdSetCullModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCullModeEXT.html)
- [vkCmdSetDepthBiasEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthBiasEnableEXT.html)
- [vkCmdSetDepthBoundsTestEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthBoundsTestEnableEXT.html)
- [vkCmdSetDepthClampEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthClampEnableEXT.html)
- [vkCmdSetDepthCompareOpEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthCompareOpEXT.html)
- [vkCmdSetDepthTestEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthTestEnableEXT.html)
- [vkCmdSetDepthWriteEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthWriteEnableEXT.html)
- [vkCmdSetFrontFaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetFrontFaceEXT.html)
- [vkCmdSetLogicOpEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetLogicOpEXT.html)
- [vkCmdSetLogicOpEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetLogicOpEnableEXT.html)
- [vkCmdSetPatchControlPointsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPatchControlPointsEXT.html)
- [vkCmdSetPolygonModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPolygonModeEXT.html)
- [vkCmdSetPrimitiveRestartEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPrimitiveRestartEnableEXT.html)
- [vkCmdSetPrimitiveTopologyEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPrimitiveTopologyEXT.html)
- [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRasterizationSamplesEXT.html)
- [vkCmdSetRasterizerDiscardEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRasterizerDiscardEnableEXT.html)
- [vkCmdSetSampleMaskEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetSampleMaskEXT.html)
- [vkCmdSetScissorWithCountEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetScissorWithCountEXT.html)
- [vkCmdSetStencilOpEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetStencilOpEXT.html)
- [vkCmdSetStencilTestEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetStencilTestEnableEXT.html)
- [vkCmdSetTessellationDomainOriginEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetTessellationDomainOriginEXT.html)
- [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetVertexInputEXT.html)
- [vkCmdSetViewportWithCountEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetViewportWithCountEXT.html)
- [vkCreateShadersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateShadersEXT.html)
- [vkDestroyShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyShaderEXT.html)
- [vkGetShaderBinaryDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderBinaryDataEXT.html)

If [VK\_EXT\_blend\_operation\_advanced](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_blend_operation_advanced.html) is supported:

- [vkCmdSetColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetColorBlendAdvancedEXT.html)

If [VK\_EXT\_conservative\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_conservative_rasterization.html) is supported:

- [vkCmdSetConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetConservativeRasterizationModeEXT.html)
- [vkCmdSetExtraPrimitiveOverestimationSizeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetExtraPrimitiveOverestimationSizeEXT.html)

If [VK\_EXT\_depth\_clamp\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clamp_control.html) is supported:

- [vkCmdSetDepthClampRangeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthClampRangeEXT.html)

If [VK\_EXT\_depth\_clip\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clip_control.html) is supported:

- [vkCmdSetDepthClipNegativeOneToOneEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthClipNegativeOneToOneEXT.html)

If [VK\_EXT\_depth\_clip\_enable](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clip_enable.html) is supported:

- [vkCmdSetDepthClipEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthClipEnableEXT.html)

If [VK\_EXT\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_line_rasterization.html) is supported:

- [vkCmdSetLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetLineRasterizationModeEXT.html)
- [vkCmdSetLineStippleEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetLineStippleEnableEXT.html)

If [VK\_EXT\_provoking\_vertex](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_provoking_vertex.html) is supported:

- [vkCmdSetProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetProvokingVertexModeEXT.html)

If [VK\_EXT\_sample\_locations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sample_locations.html) is supported:

- [vkCmdSetSampleLocationsEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetSampleLocationsEnableEXT.html)

If [VK\_EXT\_transform\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_transform_feedback.html) is supported:

- [vkCmdSetRasterizationStreamEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRasterizationStreamEXT.html)

If [VK\_NV\_clip\_space\_w\_scaling](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_clip_space_w_scaling.html) is supported:

- [vkCmdSetViewportWScalingEnableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetViewportWScalingEnableNV.html)

If [VK\_NV\_coverage\_reduction\_mode](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_coverage_reduction_mode.html) is supported:

- [vkCmdSetCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageReductionModeNV.html)

If [VK\_NV\_fragment\_coverage\_to\_color](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_fragment_coverage_to_color.html) is supported:

- [vkCmdSetCoverageToColorEnableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageToColorEnableNV.html)
- [vkCmdSetCoverageToColorLocationNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageToColorLocationNV.html)

If [VK\_NV\_framebuffer\_mixed\_samples](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_framebuffer_mixed_samples.html) is supported:

- [vkCmdSetCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageModulationModeNV.html)
- [vkCmdSetCoverageModulationTableEnableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageModulationTableEnableNV.html)
- [vkCmdSetCoverageModulationTableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageModulationTableNV.html)

If [VK\_NV\_representative\_fragment\_test](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_representative_fragment_test.html) is supported:

- [vkCmdSetRepresentativeFragmentTestEnableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRepresentativeFragmentTestEnableNV.html)

If [VK\_NV\_shading\_rate\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shading_rate_image.html) is supported:

- [vkCmdSetShadingRateImageEnableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetShadingRateImageEnableNV.html)

If [VK\_NV\_viewport\_swizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_viewport_swizzle.html) is supported:

- [vkCmdSetViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetViewportSwizzleNV.html)

## [](#_new_structures)New Structures

- [VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorBlendAdvancedEXT.html)
- [VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorBlendEquationEXT.html)
- [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateInfoEXT.html)
- [VkVertexInputAttributeDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputAttributeDescription2EXT.html)
- [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription2EXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderObjectFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderObjectFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceShaderObjectPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderObjectPropertiesEXT.html)
- Extending [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateInfoEXT.html):
  
  - [VkShaderRequiredSubgroupSizeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderRequiredSubgroupSizeCreateInfoEXT.html)

## [](#_new_enums)New Enums

- [VkShaderCodeTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCodeTypeEXT.html)
- [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateFlagBitsEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkShaderCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SHADER_OBJECT_EXTENSION_NAME`
- `VK_EXT_SHADER_OBJECT_SPEC_VERSION`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_SHADER_EXT`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_INCOMPATIBLE_SHADER_BINARY_EXT`
  - `VK_INCOMPATIBLE_SHADER_BINARY_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_OBJECT_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_OBJECT_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_SHADER_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_SHADER_REQUIRED_SUBGROUP_SIZE_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_VERTEX_INPUT_ATTRIBUTE_DESCRIPTION_2_EXT`
  - `VK_STRUCTURE_TYPE_VERTEX_INPUT_BINDING_DESCRIPTION_2_EXT`

If [VK\_EXT\_fragment\_density\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map.html) is supported:

- Extending [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateFlagBitsEXT.html):
  
  - `VK_SHADER_CREATE_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`

If [VK\_EXT\_mesh\_shader](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_mesh_shader.html) or [VK\_NV\_mesh\_shader](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_mesh_shader.html) is supported:

- Extending [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateFlagBitsEXT.html):
  
  - `VK_SHADER_CREATE_NO_TASK_SHADER_BIT_EXT`

If [VK\_EXT\_subgroup\_size\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_subgroup_size_control.html) or [Vulkan Version 1.3](#versions-1.3) is supported:

- Extending [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateFlagBitsEXT.html):
  
  - `VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT`
  - `VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT`

If [VK\_KHR\_device\_group](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_device_group.html) or [Vulkan Version 1.1](#versions-1.1) is supported:

- Extending [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateFlagBitsEXT.html):
  
  - `VK_SHADER_CREATE_DISPATCH_BASE_BIT_EXT`

If [VK\_KHR\_fragment\_shading\_rate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_fragment_shading_rate.html) is supported:

- Extending [VkShaderCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateFlagBitsEXT.html):
  
  - `VK_SHADER_CREATE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_EXT`

## [](#_examples)Examples

**Example 1**

Create linked pair of vertex and fragment shaders.

```c++
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

Later, during command buffer recording, bind the linked shaders and draw.

```c++
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

```c++
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

Later, during command buffer recording, bind the linked shaders in different combinations and draw.

```c++
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

## [](#_version_history)Version History

- Revision 1, 2023-03-30 (Daniel Story)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_shader_object)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0