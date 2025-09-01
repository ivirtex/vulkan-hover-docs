# VkDescriptorPoolCreateFlagBits(3) Manual Page

## Name

VkDescriptorPoolCreateFlagBits - Bitmask specifying certain supported operations on a descriptor pool



## [](#_c_specification)C Specification

Bits which **can** be set in [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html)::`flags`, enabling operations on a descriptor pool, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkDescriptorPoolCreateFlagBits {
    VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT = 0x00000001,
  // Provided by VK_VERSION_1_2
    VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT = 0x00000002,
  // Provided by VK_EXT_mutable_descriptor_type
    VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT = 0x00000004,
  // Provided by VK_NV_descriptor_pool_overallocation
    VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_SETS_BIT_NV = 0x00000008,
  // Provided by VK_NV_descriptor_pool_overallocation
    VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_POOLS_BIT_NV = 0x00000010,
  // Provided by VK_EXT_descriptor_indexing
    VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT = VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT,
  // Provided by VK_VALVE_mutable_descriptor_type
    VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_VALVE = VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT,
} VkDescriptorPoolCreateFlagBits;
```

## [](#_description)Description

- `VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT` specifies that descriptor sets **can** return their individual allocations to the pool, i.e. all of [vkAllocateDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateDescriptorSets.html), [vkFreeDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkFreeDescriptorSets.html), and [vkResetDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetDescriptorPool.html) are allowed. Otherwise, descriptor sets allocated from the pool **must** not be individually freed back to the pool, i.e. only [vkAllocateDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateDescriptorSets.html) and [vkResetDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetDescriptorPool.html) are allowed.
- `VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT` specifies that descriptor sets allocated from this pool **can** include bindings with the `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` bit set. It is valid to allocate descriptor sets that have bindings that do not set the `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` bit from a pool that has `VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT` set.
- `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT` specifies that this descriptor pool and the descriptor sets allocated from it reside entirely in host memory and cannot be bound. Similar to descriptor sets allocated without this flag, applications **can** copy-from and copy-to descriptors sets allocated from this descriptor pool. Descriptor sets allocated from this pool are partially exempt from the external synchronization requirement in [vkUpdateDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSetWithTemplateKHR.html) and [vkUpdateDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSets.html). Descriptor sets and their descriptors can be updated concurrently in different threads, though the same descriptor **must** not be updated concurrently by two threads.
- `VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_SETS_BIT_NV` specifies that the implementation should allow the application to allocate more than [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html)::`maxSets` descriptor set objects from the descriptor pool as available resources allow. The implementation **may** use the `maxSets` value to allocate the initial available sets, but using zero is permitted.
- `VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_POOLS_BIT_NV` specifies that the implementation should allow the application to allocate more descriptors from the pool than was specified by the [VkDescriptorPoolSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolSize.html)::`descriptorCount` for any descriptor type as specified by [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html)::`poolSizeCount` and [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html)::`pPoolSizes`, as available resources allow. The implementation **may** use the `descriptorCount` for each descriptor type to allocate the initial pool, but the application is allowed to set the `poolSizeCount` to zero, or any of the `descriptorCount` values in the `pPoolSizes` array to zero.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorPoolCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorPoolCreateFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0