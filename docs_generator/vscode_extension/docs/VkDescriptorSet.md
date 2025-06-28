# VkDescriptorSet(3) Manual Page

## Name

VkDescriptorSet - Opaque handle to a descriptor set object



## [](#_c_specification)C Specification

Descriptor sets are allocated from descriptor pool objects, and are represented by `VkDescriptorSet` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkDescriptorSet)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBindDescriptorSetsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDescriptorSetsInfo.html), [VkCopyDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyDescriptorSet.html), [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html), [vkAllocateDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateDescriptorSets.html), [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets.html), [vkFreeDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkFreeDescriptorSets.html), [vkGetDescriptorSetHostMappingVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetHostMappingVALVE.html), [vkUpdateDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSetWithTemplate.html), [vkUpdateDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSetWithTemplateKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorSet)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0