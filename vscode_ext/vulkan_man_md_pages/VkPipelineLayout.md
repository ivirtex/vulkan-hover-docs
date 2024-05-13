# VkPipelineLayout(3) Manual Page

## Name

VkPipelineLayout - Opaque handle to a pipeline layout object



## <a href="#_c_specification" class="anchor"></a>C Specification

Access to descriptor sets from a pipeline is accomplished through a
*pipeline layout*. Zero or more descriptor set layouts and zero or more
push constant ranges are combined to form a pipeline layout object
describing the complete set of resources that **can** be accessed by a
pipeline. The pipeline layout represents a sequence of descriptor sets
with each having a specific layout. This sequence of layouts is used to
determine the interface between shader stages and shader resources. Each
pipeline is created using a pipeline layout.

Pipeline layout objects are represented by `VkPipelineLayout` handles:

``` c
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkPipelineLayout)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBindDescriptorBufferEmbeddedSamplersInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindDescriptorBufferEmbeddedSamplersInfoEXT.html),
[VkBindDescriptorSetsInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindDescriptorSetsInfoKHR.html),
[VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html),
[VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateInfo.html),
[VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html),
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutTokenNV.html),
[VkPushConstantsInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantsInfoKHR.html),
[VkPushDescriptorSetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushDescriptorSetInfoKHR.html),
[VkPushDescriptorSetWithTemplateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushDescriptorSetWithTemplateInfoKHR.html),
[VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html),
[VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html),
[VkSetDescriptorBufferOffsetsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSetDescriptorBufferOffsetsInfoEXT.html),
[vkCmdBindDescriptorBufferEmbeddedSamplersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorBufferEmbeddedSamplersEXT.html),
[vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets.html),
[vkCmdPushConstants](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushConstants.html),
[vkCmdPushDescriptorSetKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetKHR.html),
[vkCmdPushDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html),
[vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html),
[vkCreatePipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreatePipelineLayout.html),
[vkDestroyPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyPipelineLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineLayout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
