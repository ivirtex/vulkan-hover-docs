# vkResetDescriptorPool(3) Manual Page

## Name

vkResetDescriptorPool - Resets a descriptor pool object



## [](#_c_specification)C Specification

To return all descriptor sets allocated from a given pool to the pool, rather than freeing individual descriptor sets, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkResetDescriptorPool(
    VkDevice                                    device,
    VkDescriptorPool                            descriptorPool,
    VkDescriptorPoolResetFlags                  flags);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the descriptor pool.
- `descriptorPool` is the descriptor pool to be reset.
- `flags` is reserved for future use.

## [](#_description)Description

Resetting a descriptor pool recycles all of the resources from all of the descriptor sets allocated from the descriptor pool back to the descriptor pool, and the descriptor sets are implicitly freed.

Valid Usage

- [](#VUID-vkResetDescriptorPool-descriptorPool-00313)VUID-vkResetDescriptorPool-descriptorPool-00313  
  All uses of `descriptorPool` (via any allocated descriptor sets) **must** have completed execution

Valid Usage (Implicit)

- [](#VUID-vkResetDescriptorPool-device-parameter)VUID-vkResetDescriptorPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkResetDescriptorPool-descriptorPool-parameter)VUID-vkResetDescriptorPool-descriptorPool-parameter  
  `descriptorPool` **must** be a valid [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPool.html) handle
- [](#VUID-vkResetDescriptorPool-flags-zerobitmask)VUID-vkResetDescriptorPool-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-vkResetDescriptorPool-descriptorPool-parent)VUID-vkResetDescriptorPool-descriptorPool-parent  
  `descriptorPool` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `descriptorPool` **must** be externally synchronized
- Host access to any `VkDescriptorSet` objects allocated from `descriptorPool` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

This command does not return any failure codes

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPool.html), [VkDescriptorPoolResetFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolResetFlags.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkResetDescriptorPool)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0