# VkDescriptorUpdateTemplate(3) Manual Page

## Name

VkDescriptorUpdateTemplate - Opaque handle to a descriptor update
template



## <a href="#_c_specification" class="anchor"></a>C Specification

A descriptor update template specifies a mapping from descriptor update
information in host memory to descriptors in a descriptor set. It is
designed to avoid passing redundant information to the driver when
frequently updating the same set of descriptors in descriptor sets.

Descriptor update template objects are represented by
`VkDescriptorUpdateTemplate` handles:

``` c
// Provided by VK_VERSION_1_1
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkDescriptorUpdateTemplate)
```

or the equivalent

``` c
// Provided by VK_KHR_descriptor_update_template
typedef VkDescriptorUpdateTemplate VkDescriptorUpdateTemplateKHR;
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkPushDescriptorSetWithTemplateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushDescriptorSetWithTemplateInfoKHR.html),
[vkCmdPushDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html),
[vkCreateDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplate.html),
[vkCreateDescriptorUpdateTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplateKHR.html),
[vkDestroyDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorUpdateTemplate.html),
[vkDestroyDescriptorUpdateTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorUpdateTemplateKHR.html),
[vkUpdateDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplate.html),
[vkUpdateDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorUpdateTemplate"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
