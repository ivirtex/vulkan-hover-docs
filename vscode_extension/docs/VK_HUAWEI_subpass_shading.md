# VK\_HUAWEI\_subpass\_shading(3) Manual Page

## Name

VK\_HUAWEI\_subpass\_shading - device extension



## [](#_registered_extension_number)Registered Extension Number

370

## [](#_revision)Revision

3

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html)  
         or  
         [Vulkan Version 1.2](#versions-1.2)  
     and  
     [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_HUAWEI\_subpass\_shading](https://github.khronos.org/SPIRV-Registry/extensions/HUAWEI/SPV_HUAWEI_subpass_shading.html)

## [](#_contact)Contact

- Pan Gao [\[GitHub\]PanGao-h](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_HUAWEI_subpass_shading%5D%20%40PanGao-h%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_HUAWEI_subpass_shading%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-06-01

**Interactions and External Dependencies**

- This extension provides API support for [`GL_HUAWEI_subpass_shading`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/huawei/GLSL_HUAWEI_subpass_shading.txt).

**Contributors**

- Hueilong Wang
- Juntao Li, Huawei
- Renmiao Lu, Huawei
- Pan Gao, Huawei

## [](#_description)Description

This extension allows applications to execute a subpass shading pipeline in a subpass of a render pass in order to save memory bandwidth for algorithms like tile-based deferred rendering and forward plus. A subpass shading pipeline is a pipeline with the compute pipeline ability, allowed to read values from input attachments, and only allowed to be dispatched inside a stand-alone subpass. Its work dimension is defined by the render pass’s render area size. Its workgroup size (width, height) shall be a power-of-two number in width or height, with minimum value from 8, and maximum value shall be decided from the render pass attachments and sample counts but depends on implementation.

The `GlobalInvocationId.xy` of a subpass shading pipeline is equal to the `FragCoord.xy` of a graphic pipeline in the same render pass subtracted the [`offset`](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) of the [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html)::`renderArea`. `GlobalInvocationId.z` is mapped to the Layer if `VK_EXT_shader_viewport_index_layer` is supported. The `GlobalInvocationId.xy` is equal to the index of the local workgroup multiplied by the size of the local workgroup plus the `LocalInvocationId` and the [`offset`](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) of the [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html)::`renderArea`.

This extension allows a subpass’s pipeline bind point to be `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI`.

## [](#_new_commands)New Commands

- [vkCmdSubpassShadingHUAWEI](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSubpassShadingHUAWEI.html)
- [vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI.html)

## [](#_new_structures)New Structures

- Extending [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html):
  
  - [VkSubpassShadingPipelineCreateInfoHUAWEI](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassShadingPipelineCreateInfoHUAWEI.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceSubpassShadingFeaturesHUAWEI](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubpassShadingFeaturesHUAWEI.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceSubpassShadingPropertiesHUAWEI](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubpassShadingPropertiesHUAWEI.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_HUAWEI_SUBPASS_SHADING_EXTENSION_NAME`
- `VK_HUAWEI_SUBPASS_SHADING_SPEC_VERSION`
- Extending [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html):
  
  - `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`
  - `VK_PIPELINE_STAGE_2_SUBPASS_SHADING_BIT_HUAWEI`
- Extending [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html):
  
  - `VK_SHADER_STAGE_SUBPASS_SHADING_BIT_HUAWEI`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBPASS_SHADING_FEATURES_HUAWEI`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBPASS_SHADING_PROPERTIES_HUAWEI`
  - `VK_STRUCTURE_TYPE_SUBPASS_SHADING_PIPELINE_CREATE_INFO_HUAWEI`

## [](#_sample_code)Sample Code

Example of subpass shading in a GLSL shader

```c
#extension GL_HUAWEI_subpass_shading: enable
#extension GL_KHR_shader_subgroup_arithmetic: enable

layout(constant_id = 0) const uint tileWidth = 8;
layout(constant_id = 1) const uint tileHeight = 8;
layout(local_size_x_id = 0, local_size_y_id = 1, local_size_z = 1) in;
layout(set=0, binding=0, input_attachment_index=0) uniform subpassInput depth;

void main()
{
  float d = subpassLoad(depth).x;
  float minD = subgroupMin(d);
  float maxD = subgroupMax(d);
}
```

Example of subpass shading dispatching in a subpass

```c
vkCmdNextSubpass(commandBuffer, VK_SUBPASS_CONTENTS_INLINE);
vkCmdBindPipeline(commandBuffer, VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI, subpassShadingPipeline);
vkCmdBindDescriptorSets(commandBuffer, VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI, subpassShadingPipelineLayout,
  firstSet, descriptorSetCount, pDescriptorSets, dynamicOffsetCount, pDynamicOffsets);
vkCmdSubpassShadingHUAWEI(commandBuffer)
vkCmdEndRenderPass(commandBuffer);
```

Example of subpass shading render pass creation

```c
VkAttachmentDescription2 attachments[] = {
  {
    VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2, NULL,
    0, VK_FORMAT_R8G8B8A8_UNORM, VK_SAMPLE_COUNT_1_BIT,
    VK_ATTACHMENT_LOAD_OP_CLEAR, VK_ATTACHMENT_STORE_OP_DONT_CARE,
    VK_ATTACHMENT_LOAD_OP_DONT_CARE, VK_ATTACHMENT_LOAD_OP_DONT_CARE,
    VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL, VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
  },
  {
    VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2, NULL,
    0, VK_FORMAT_R8G8B8A8_UNORM, VK_SAMPLE_COUNT_1_BIT,
    VK_ATTACHMENT_LOAD_OP_CLEAR, VK_ATTACHMENT_STORE_OP_DONT_CARE,
    VK_ATTACHMENT_LOAD_OP_DONT_CARE, VK_ATTACHMENT_LOAD_OP_DONT_CARE,
    VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL, VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
  },
  {
    VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2, NULL,
    0, VK_FORMAT_R8G8B8A8_UNORM, VK_SAMPLE_COUNT_1_BIT,
    VK_ATTACHMENT_LOAD_OP_CLEAR, VK_ATTACHMENT_STORE_OP_DONT_CARE,
    VK_ATTACHMENT_LOAD_OP_DONT_CARE, VK_ATTACHMENT_LOAD_OP_DONT_CARE,
    VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL, VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
  },
  {
    VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2, NULL,
    0, VK_FORMAT_D24_UNORM_S8_UINT, VK_SAMPLE_COUNT_1_BIT,
    VK_ATTACHMENT_LOAD_OP_CLEAR, VK_ATTACHMENT_STORE_OP_DONT_CARE,
    VK_ATTACHMENT_LOAD_OP_CLEAR, VK_ATTACHMENT_LOAD_OP_DONT_CARE,
    VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL, VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL
  },
  {
    VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2, NULL,
    0, VK_FORMAT_R8G8B8A8_UNORM, VK_SAMPLE_COUNT_1_BIT,
    VK_ATTACHMENT_LOAD_OP_CLEAR, VK_ATTACHMENT_STORE_OP_STORE,
    VK_ATTACHMENT_LOAD_OP_DONT_CARE, VK_ATTACHMENT_LOAD_OP_DONT_CARE,
    VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL, VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
  }
};

VkAttachmentReference2 gBufferAttachmentReferences[] = {
  { VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2, NULL, 0, VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL, VK_IMAGE_ASPECT_COLOR_BIT },
  { VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2, NULL, 1, VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL, VK_IMAGE_ASPECT_COLOR_BIT },
  { VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2, NULL, 2, VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL, VK_IMAGE_ASPECT_COLOR_BIT }
};
VkAttachmentReference2 gBufferDepthStencilAttachmentReferences =
  { VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2, NULL, 3, VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL, VK_IMAGE_ASPECT_DEPTH_BIT|VK_IMAGE_ASPECT_STENCIL_BIT };
VkAttachmentReference2 depthInputAttachmentReferences[] = {
  { VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2, NULL, 3, VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL, VK_IMAGE_ASPECT_DEPTH_BIT|VK_IMAGE_ASPECT_STENCIL_BIT };
};
VkAttachmentReference2 preserveAttachmentReferences[] = {
  { VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2, NULL, 0, VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL, VK_IMAGE_ASPECT_COLOR_BIT },
  { VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2, NULL, 1, VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL, VK_IMAGE_ASPECT_COLOR_BIT },
  { VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2, NULL, 2, VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL, VK_IMAGE_ASPECT_COLOR_BIT },
  { VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2, NULL, 3, VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL, VK_IMAGE_ASPECT_DEPTH_BIT|VK_IMAGE_ASPECT_STENCIL_BIT }
}; // G buffer including depth/stencil
VkAttachmentReference2 colorAttachmentReferences[] = {
  { VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2, NULL, 4, VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL, VK_IMAGE_ASPECT_COLOR_BIT }
};
VkAttachmentReference2 resolveAttachmentReference =
  { VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2, NULL, 4, VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL, VK_IMAGE_ASPECT_COLOR_BIT };

VkSubpassDescription2 subpasses[] = {
  {
    VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2, NULL, 0, VK_PIPELINE_BIND_POINT_GRAPHICS, 0,
    0, NULL, // input
    sizeof(gBufferAttachmentReferences)/sizeof(gBufferAttachmentReferences[0]), gBufferAttachmentReferences, // color
    NULL, &gBufferDepthStencilAttachmentReferences, // resolve & DS
    0, NULL
  },
  {
    VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2, NULL, 0, VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI , 0,
    sizeof(depthInputAttachmentReferences)/sizeof(depthInputAttachmentReferences[0]), depthInputAttachmentReferences, // input
    0, NULL, // color
    NULL, NULL, // resolve & DS
    sizeof(preserveAttachmentReferences)/sizeof(preserveAttachmentReferences[0]), preserveAttachmentReferences,
  },
  {
    VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2, NULL, 0, VK_PIPELINE_BIND_POINT_GRAPHICS, 0,
    sizeof(gBufferAttachmentReferences)/sizeof(gBufferAttachmentReferences[0]), gBufferAttachmentReferences, // input
    sizeof(colorAttachmentReferences)/sizeof(colorAttachmentReferences[0]), colorAttachmentReferences, // color
    &resolveAttachmentReference, &gBufferDepthStencilAttachmentReferences, // resolve & DS
    0, NULL
  },
};

VkMemoryBarrier2KHR fragmentToSubpassShading = {
  VK_STRUCTURE_TYPE_MEMORY_BARRIER_2_KHR, NULL,
  VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT_KHR, VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT|VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT,
  VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI, VK_ACCESS_INPUT_ATTACHMENT_READ_BIT
};

VkMemoryBarrier2KHR subpassShadingToFragment = {
  VK_STRUCTURE_TYPE_MEMORY_BARRIER_2_KHR, NULL,
  VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI, VK_ACCESS_SHADER_WRITE_BIT,
  VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT_KHR, VK_ACCESS_SHADER_READ_BIT
};

VkSubpassDependency2 dependencies[] = {
  {
    VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2, &fragmentToSubpassShading,
    0, 1,
    0, 0, 0, 0,
    0, 0
  },
  {
    VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2, &subpassShadingToFragment,
    1, 2,
    0, 0, 0, 0,
    0, 0
  },
};

VkRenderPassCreateInfo2 renderPassCreateInfo = {
  VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2, NULL, 0,
    sizeof(attachments)/sizeof(attachments[0]), attachments,
    sizeof(subpasses)/sizeof(subpasses[0]), subpasses,
    sizeof(dependencies)/sizeof(dependencies[0]), dependencies,
    0, NULL
};
VKRenderPass renderPass;
vkCreateRenderPass2(device, &renderPassCreateInfo, NULL, &renderPass);
```

Example of subpass shading pipeline creation

```c
VkExtent2D maxWorkgroupSize;

VkSpecializationMapEntry subpassShadingConstantMapEntries[] = {
  { 0, 0 * sizeof(uint32_t), sizeof(uint32_t) },
  { 1, 1 * sizeof(uint32_t), sizeof(uint32_t) }
};

VkSpecializationInfo subpassShadingConstants = {
  2, subpassShadingConstantMapEntries,
  sizeof(VkExtent2D), &maxWorkgroupSize
};

VkSubpassShadingPipelineCreateInfoHUAWEI subpassShadingPipelineCreateInfo {
  VK_STRUCTURE_TYPE_SUBPASSS_SHADING_PIPELINE_CREATE_INFO_HUAWEI, NULL,
  renderPass, 1
};

VkPipelineShaderStageCreateInfo subpassShadingPipelineStageCreateInfo {
  VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_CREATE_INFO, NULL,
  0, VK_SHADER_STAGE_SUBPASS_SHADING_BIT_HUAWEI,
  shaderModule, "main",
  &subpassShadingConstants
};

VkComputePipelineCreateInfo subpassShadingComputePipelineCreateInfo = {
  VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO, &subpassShadingPipelineCreateInfo,
  0, &subpassShadingPipelineStageCreateInfo,
  pipelineLayout, basePipelineHandle, basePipelineIndex
};

VKPipeline pipeline;

vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI(device, renderPass, &maxWorkgroupSize);
vkCreateComputePipelines(device, pipelineCache, 1, &subpassShadingComputePipelineCreateInfo, NULL, &pipeline);
```

## [](#_version_history)Version History

- Revision 3, 2023-06-19 (Pan Gao)
  
  - Rename `VK_PIPELINE_STAGE_2_SUBPASS_SHADING_BIT_HUAWEI` to `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI` to better aligned with naming of other pipeline stages
- Revision 2, 2021-06-28 (Hueilong Wang)
  
  - Change vkGetSubpassShadingMaxWorkgroupSizeHUAWEI to vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI to resolve issue [`pub1564`](https://github.com/KhronosGroup/Vulkan-Docs/issues/1564)
- Revision 1, 2020-12-15 (Hueilong Wang)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_HUAWEI_subpass_shading).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0