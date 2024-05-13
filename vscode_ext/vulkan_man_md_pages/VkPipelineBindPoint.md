# VkPipelineBindPoint(3) Manual Page

## Name

VkPipelineBindPoint - Specify the bind point of a pipeline object to a
command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[vkCmdBindPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindPipeline.html)::`pipelineBindPoint`,
specifying the bind point of a pipeline object, are:

``` c
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
  // Provided by VK_NV_ray_tracing
    VK_PIPELINE_BIND_POINT_RAY_TRACING_NV = VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR,
} VkPipelineBindPoint;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PIPELINE_BIND_POINT_COMPUTE` specifies binding as a compute
  pipeline.

- `VK_PIPELINE_BIND_POINT_GRAPHICS` specifies binding as a graphics
  pipeline.

- `VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR` specifies binding as a ray
  tracing pipeline.

- `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI` specifies binding as a
  subpass shading pipeline.

- `VK_PIPELINE_BIND_POINT_EXECUTION_GRAPH_AMDX` specifies binding as an
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#executiongraphs"
  target="_blank" rel="noopener">execution graph pipeline</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateInfo.html),
[VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeneratedCommandsInfoNV.html),
[VkGeneratedCommandsMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeneratedCommandsMemoryRequirementsInfoNV.html),
[VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutCreateInfoNV.html),
[VkPipelineIndirectDeviceAddressInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineIndirectDeviceAddressInfoNV.html),
[VkSubpassDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription.html),
[VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html),
[vkCmdBindDescriptorBufferEmbeddedSamplersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorBufferEmbeddedSamplersEXT.html),
[vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets.html),
[vkCmdBindPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindPipeline.html),
[vkCmdBindPipelineShaderGroupNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindPipelineShaderGroupNV.html),
[vkCmdPushDescriptorSetKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetKHR.html),
[vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html),
[vkCmdUpdatePipelineIndirectBufferNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdUpdatePipelineIndirectBufferNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineBindPoint"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
