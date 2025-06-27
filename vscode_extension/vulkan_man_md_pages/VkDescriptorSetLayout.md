# VkDescriptorSetLayout(3) Manual Page

## Name

VkDescriptorSetLayout - Opaque handle to a descriptor set layout object



## <a href="#_c_specification" class="anchor"></a>C Specification

A descriptor set layout object is defined by an array of zero or more
descriptor bindings. Each individual descriptor binding is specified by
a descriptor type, a count (array size) of the number of descriptors in
the binding, a set of shader stages that **can** access the binding, and
(if using immutable samplers) an array of sampler descriptors.

Descriptor set layout objects are represented by `VkDescriptorSetLayout`
handles:

``` c
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkDescriptorSetLayout)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_DEFINE_NON_DISPATCHABLE_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html),
[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDescriptorSetAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetAllocateInfo.html),
[VkDescriptorSetBindingReferenceVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetBindingReferenceVALVE.html),
[VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateInfo.html),
[VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html),
[VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html),
[vkCreateDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorSetLayout.html),
[vkDestroyDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorSetLayout.html),
[vkGetDescriptorSetLayoutBindingOffsetEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutBindingOffsetEXT.html),
[vkGetDescriptorSetLayoutSizeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSizeEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorSetLayout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
