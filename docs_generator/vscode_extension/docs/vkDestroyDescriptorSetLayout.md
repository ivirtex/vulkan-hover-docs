# vkDestroyDescriptorSetLayout(3) Manual Page

## Name

vkDestroyDescriptorSetLayout - Destroy a descriptor set layout object



## [](#_c_specification)C Specification

To destroy a descriptor set layout, call:

```c++
// Provided by VK_VERSION_1_0
void vkDestroyDescriptorSetLayout(
    VkDevice                                    device,
    VkDescriptorSetLayout                       descriptorSetLayout,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the descriptor set layout.
- `descriptorSetLayout` is the descriptor set layout to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-00284)VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-00284  
  If `VkAllocationCallbacks` were provided when `descriptorSetLayout` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-00285)VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-00285  
  If no `VkAllocationCallbacks` were provided when `descriptorSetLayout` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyDescriptorSetLayout-device-parameter)VUID-vkDestroyDescriptorSetLayout-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-parameter)VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-parameter  
  If `descriptorSetLayout` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `descriptorSetLayout` **must** be a valid [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) handle
- [](#VUID-vkDestroyDescriptorSetLayout-pAllocator-parameter)VUID-vkDestroyDescriptorSetLayout-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-parent)VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-parent  
  If `descriptorSetLayout` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `descriptorSetLayout` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyDescriptorSetLayout)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0