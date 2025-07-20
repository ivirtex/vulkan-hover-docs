# vkFreeDescriptorSets(3) Manual Page

## Name

vkFreeDescriptorSets - Free one or more descriptor sets



## [](#_c_specification)C Specification

To free allocated descriptor sets, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkFreeDescriptorSets(
    VkDevice                                    device,
    VkDescriptorPool                            descriptorPool,
    uint32_t                                    descriptorSetCount,
    const VkDescriptorSet*                      pDescriptorSets);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the descriptor pool.
- `descriptorPool` is the descriptor pool from which the descriptor sets were allocated.
- `descriptorSetCount` is the number of elements in the `pDescriptorSets` array.
- `pDescriptorSets` is a pointer to an array of handles to [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html) objects.

## [](#_description)Description

After calling `vkFreeDescriptorSets`, all descriptor sets in `pDescriptorSets` are invalid.

Valid Usage

- [](#VUID-vkFreeDescriptorSets-pDescriptorSets-00309)VUID-vkFreeDescriptorSets-pDescriptorSets-00309  
  All submitted commands that refer to any element of `pDescriptorSets` **must** have completed execution
- [](#VUID-vkFreeDescriptorSets-pDescriptorSets-00310)VUID-vkFreeDescriptorSets-pDescriptorSets-00310  
  `pDescriptorSets` **must** be a valid pointer to an array of `descriptorSetCount` `VkDescriptorSet` handles, each element of which **must** either be a valid handle or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkFreeDescriptorSets-descriptorPool-00312)VUID-vkFreeDescriptorSets-descriptorPool-00312  
  `descriptorPool` **must** have been created with the `VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT` flag

Valid Usage (Implicit)

- [](#VUID-vkFreeDescriptorSets-device-parameter)VUID-vkFreeDescriptorSets-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkFreeDescriptorSets-descriptorPool-parameter)VUID-vkFreeDescriptorSets-descriptorPool-parameter  
  `descriptorPool` **must** be a valid [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPool.html) handle
- [](#VUID-vkFreeDescriptorSets-descriptorSetCount-arraylength)VUID-vkFreeDescriptorSets-descriptorSetCount-arraylength  
  `descriptorSetCount` **must** be greater than `0`
- [](#VUID-vkFreeDescriptorSets-descriptorPool-parent)VUID-vkFreeDescriptorSets-descriptorPool-parent  
  `descriptorPool` **must** have been created, allocated, or retrieved from `device`
- [](#VUID-vkFreeDescriptorSets-pDescriptorSets-parent)VUID-vkFreeDescriptorSets-pDescriptorSets-parent  
  Each element of `pDescriptorSets` that is a valid handle **must** have been created, allocated, or retrieved from `descriptorPool`

Host Synchronization

- Host access to `descriptorPool` **must** be externally synchronized
- Host access to each member of `pDescriptorSets` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPool.html), [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkFreeDescriptorSets)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0