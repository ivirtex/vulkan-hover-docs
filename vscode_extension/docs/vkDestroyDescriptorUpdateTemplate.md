# vkDestroyDescriptorUpdateTemplate(3) Manual Page

## Name

vkDestroyDescriptorUpdateTemplate - Destroy a descriptor update template object



## [](#_c_specification)C Specification

To destroy a descriptor update template, call:

```c++
// Provided by VK_VERSION_1_1
void vkDestroyDescriptorUpdateTemplate(
    VkDevice                                    device,
    VkDescriptorUpdateTemplate                  descriptorUpdateTemplate,
    const VkAllocationCallbacks*                pAllocator);
```

or the equivalent command

```c++
// Provided by VK_KHR_descriptor_update_template
void vkDestroyDescriptorUpdateTemplateKHR(
    VkDevice                                    device,
    VkDescriptorUpdateTemplate                  descriptorUpdateTemplate,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that has been used to create the descriptor update template
- `descriptorUpdateTemplate` is the descriptor update template to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyDescriptorUpdateTemplate-descriptorSetLayout-00356)VUID-vkDestroyDescriptorUpdateTemplate-descriptorSetLayout-00356  
  If `VkAllocationCallbacks` were provided when `descriptorUpdateTemplate` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyDescriptorUpdateTemplate-descriptorSetLayout-00357)VUID-vkDestroyDescriptorUpdateTemplate-descriptorSetLayout-00357  
  If no `VkAllocationCallbacks` were provided when `descriptorUpdateTemplate` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyDescriptorUpdateTemplate-device-parameter)VUID-vkDestroyDescriptorUpdateTemplate-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyDescriptorUpdateTemplate-descriptorUpdateTemplate-parameter)VUID-vkDestroyDescriptorUpdateTemplate-descriptorUpdateTemplate-parameter  
  If `descriptorUpdateTemplate` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `descriptorUpdateTemplate` **must** be a valid [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplate.html) handle
- [](#VUID-vkDestroyDescriptorUpdateTemplate-pAllocator-parameter)VUID-vkDestroyDescriptorUpdateTemplate-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyDescriptorUpdateTemplate-descriptorUpdateTemplate-parent)VUID-vkDestroyDescriptorUpdateTemplate-descriptorUpdateTemplate-parent  
  If `descriptorUpdateTemplate` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `descriptorUpdateTemplate` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplate.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyDescriptorUpdateTemplate).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0