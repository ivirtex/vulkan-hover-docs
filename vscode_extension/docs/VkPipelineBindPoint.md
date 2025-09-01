# VkPipelineBindPoint(3) Manual Page

## Name

VkPipelineBindPoint - Specify the bind point of a pipeline object to a command buffer



## [](#_c_specification)C Specification

Possible values of [vkCmdBindPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindPipeline.html)::`pipelineBindPoint`, specifying the bind point of a pipeline object, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkPipelineBindPoint {
    VK_PIPELINE_BIND_POINT_GRAPHICS = 0,
    VK_PIPELINE_BIND_POINT_COMPUTE = 1,
#ifdef VK_ENABLE_BETA_EXTENSIONS
  // Provided by VK_AMDX_shader_enqueue
    VK_PIPELINE_BIND_POINT_EXECUTION_GRAPH_AMDX = 1000134000,
#endif
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR = 1000165000,
  // Provided by VK_HUAWEI_subpass_shading
    VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI = 1000369003,
  // Provided by VK_ARM_data_graph
    VK_PIPELINE_BIND_POINT_DATA_GRAPH_ARM = 1000507000,
  // Provided by VK_NV_ray_tracing
    VK_PIPELINE_BIND_POINT_RAY_TRACING_NV = VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR,
} VkPipelineBindPoint;
```

## [](#_description)Description

- `VK_PIPELINE_BIND_POINT_COMPUTE` specifies binding as a compute pipeline.
- `VK_PIPELINE_BIND_POINT_GRAPHICS` specifies binding as a graphics pipeline.
- `VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR` specifies binding as a ray tracing pipeline.
- `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI` specifies binding as a subpass shading pipeline.
- `VK_PIPELINE_BIND_POINT_EXECUTION_GRAPH_AMDX` specifies binding as an [execution graph pipeline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#executiongraphs).

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateCreateInfo.html), [VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoNV.html), [VkGeneratedCommandsMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoNV.html), [VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoNV.html), [VkPipelineIndirectDeviceAddressInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineIndirectDeviceAddressInfoNV.html), [VkSubpassDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription.html), [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html), [vkCmdBindDescriptorBufferEmbeddedSamplersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorBufferEmbeddedSamplersEXT.html), [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets.html), [vkCmdBindPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindPipeline.html), [vkCmdBindPipelineShaderGroupNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindPipelineShaderGroupNV.html), [vkCmdPushDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSet.html), [vkCmdPushDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSet.html), [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), [vkCmdUpdatePipelineIndirectBufferNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdUpdatePipelineIndirectBufferNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineBindPoint).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0