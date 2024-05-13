# VkDescriptorPool(3) Manual Page

## Name

VkDescriptorPool - Opaque handle to a descriptor pool object



## <a href="#_c_specification" class="anchor"></a>C Specification

A *descriptor pool* maintains a pool of descriptors, from which
descriptor sets are allocated. Descriptor pools are externally
synchronized, meaning that the application **must** not allocate and/or
free descriptor sets from the same pool in multiple threads
simultaneously.

Descriptor pools are represented by `VkDescriptorPool` handles:

``` c
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkDescriptorPool)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDescriptorSetAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetAllocateInfo.html),
[vkCreateDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorPool.html),
[vkDestroyDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorPool.html),
[vkFreeDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkFreeDescriptorSets.html),
[vkResetDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetDescriptorPool.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorPool"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
