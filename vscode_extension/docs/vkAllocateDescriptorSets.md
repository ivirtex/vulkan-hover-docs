# vkAllocateDescriptorSets(3) Manual Page

## Name

vkAllocateDescriptorSets - Allocate one or more descriptor sets



## [](#_c_specification)C Specification

To allocate descriptor sets from a descriptor pool, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkAllocateDescriptorSets(
    VkDevice                                    device,
    const VkDescriptorSetAllocateInfo*          pAllocateInfo,
    VkDescriptorSet*                            pDescriptorSets);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the descriptor pool.
- `pAllocateInfo` is a pointer to a [VkDescriptorSetAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetAllocateInfo.html) structure describing parameters of the allocation.
- `pDescriptorSets` is a pointer to an array of [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html) handles in which the resulting descriptor set objects are returned.

## [](#_description)Description

The allocated descriptor sets are returned in `pDescriptorSets`.

When a descriptor set is allocated, the initial state is largely uninitialized and all descriptors are undefined, with the exception that samplers with a non-null `pImmutableSamplers` are initialized on allocation. Descriptors also become undefined if the underlying resource or view object is destroyed. Descriptor sets containing undefined descriptors **can** still be bound and used, subject to the following conditions:

- For descriptor set bindings created with the `VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT` bit set, all descriptors in that binding that are dynamically used **must** have been populated before the descriptor set is [consumed](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-binding).
- For descriptor set bindings created without the `VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT` bit set, all descriptors in that binding that are statically used **must** have been populated before the descriptor set is [consumed](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-binding).
- Descriptor bindings with descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` **can** be undefined when the descriptor set is [consumed](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-binding); though values in that block will be undefined.
- Entries that are not used by a pipeline **can** have undefined descriptors.

If a call to `vkAllocateDescriptorSets` would cause the total number of descriptor sets allocated from the pool to exceed the value of [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html)::`maxSets` used to create `pAllocateInfo->descriptorPool`, then the allocation **may** fail due to lack of space in the descriptor pool. Similarly, the allocation **may** fail due to lack of space if the call to `vkAllocateDescriptorSets` would cause the number of any given descriptor type to exceed the sum of all the `descriptorCount` members of each element of [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html)::`pPoolSizes` with a `type` equal to that type.

Additionally, the allocation **may** also fail if a call to `vkAllocateDescriptorSets` would cause the total number of inline uniform block bindings allocated from the pool to exceed the value of [VkDescriptorPoolInlineUniformBlockCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolInlineUniformBlockCreateInfo.html)::`maxInlineUniformBlockBindings` used to create the descriptor pool.

If the allocation fails due to no more space in the descriptor pool, and not because of system or device memory exhaustion, then `VK_ERROR_OUT_OF_POOL_MEMORY` **must** be returned.

`vkAllocateDescriptorSets` **can** be used to create multiple descriptor sets. If the creation of any of those descriptor sets fails, then the implementation **must** destroy all successfully created descriptor set objects from this command, set all entries of the `pDescriptorSets` array to [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and return the error.

Valid Usage (Implicit)

- [](#VUID-vkAllocateDescriptorSets-device-parameter)VUID-vkAllocateDescriptorSets-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkAllocateDescriptorSets-pAllocateInfo-parameter)VUID-vkAllocateDescriptorSets-pAllocateInfo-parameter  
  `pAllocateInfo` **must** be a valid pointer to a valid [VkDescriptorSetAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetAllocateInfo.html) structure
- [](#VUID-vkAllocateDescriptorSets-pDescriptorSets-parameter)VUID-vkAllocateDescriptorSets-pDescriptorSets-parameter  
  `pDescriptorSets` **must** be a valid pointer to an array of `pAllocateInfo->descriptorSetCount` [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html) handles
- [](#VUID-vkAllocateDescriptorSets-device-queuecount)VUID-vkAllocateDescriptorSets-device-queuecount  
  The device **must** have been created with at least `1` queue
- [](#VUID-vkAllocateDescriptorSets-pAllocateInfo::descriptorSetCount-arraylength)VUID-vkAllocateDescriptorSets-pAllocateInfo::descriptorSetCount-arraylength  
  `pAllocateInfo->descriptorSetCount` **must** be greater than `0`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_FRAGMENTED_POOL`
- `VK_ERROR_OUT_OF_POOL_MEMORY`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html), [VkDescriptorSetAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetAllocateInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkAllocateDescriptorSets)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0