# VkDescriptorPool(3) Manual Page

## Name

VkDescriptorPool - Opaque handle to a descriptor pool object



## [](#_c_specification)C Specification

A *descriptor pool* maintains a pool of descriptors, from which descriptor sets are allocated. Descriptor pools are externally synchronized, meaning that the application **must** not allocate and/or free descriptor sets from the same pool in multiple threads simultaneously.

Descriptor pools are represented by `VkDescriptorPool` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkDescriptorPool)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorSetAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetAllocateInfo.html), [vkCreateDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDescriptorPool.html), [vkDestroyDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyDescriptorPool.html), [vkFreeDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkFreeDescriptorSets.html), [vkResetDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetDescriptorPool.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorPool).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0