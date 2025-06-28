# VkPipelineLayout(3) Manual Page

## Name

VkPipelineLayout - Opaque handle to a pipeline layout object



## [](#_c_specification)C Specification

Access to descriptor sets from a pipeline is accomplished through a *pipeline layout*. Zero or more descriptor set layouts and zero or more push constant ranges are combined to form a pipeline layout object describing the complete set of resources that **can** be accessed by a pipeline. The pipeline layout represents a sequence of descriptor sets with each having a specific layout. This sequence of layouts is used to determine the interface between shader stages and shader resources. Each pipeline is created using a pipeline layout.

Pipeline layout objects are represented by `VkPipelineLayout` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkPipelineLayout)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBindDescriptorBufferEmbeddedSamplersInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDescriptorBufferEmbeddedSamplersInfoEXT.html), [VkBindDescriptorSetsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDescriptorSetsInfo.html), [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html), [VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateCreateInfo.html), [VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkIndirectCommandsLayoutCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoEXT.html), [VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenNV.html), [VkPushConstantsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantsInfo.html), [VkPushDescriptorSetInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushDescriptorSetInfo.html), [VkPushDescriptorSetWithTemplateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushDescriptorSetWithTemplateInfo.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html), [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html), [VkSetDescriptorBufferOffsetsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSetDescriptorBufferOffsetsInfoEXT.html), [vkCmdBindDescriptorBufferEmbeddedSamplersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorBufferEmbeddedSamplersEXT.html), [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets.html), [vkCmdPushConstants](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushConstants.html), [vkCmdPushDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSet.html), [vkCmdPushDescriptorSetKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSetKHR.html), [vkCmdPushDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSetWithTemplate.html), [vkCmdPushDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html), [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html), [vkCreatePipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePipelineLayout.html), [vkDestroyPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyPipelineLayout.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineLayout)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0