# vkCreateDescriptorUpdateTemplate(3) Manual Page

## Name

vkCreateDescriptorUpdateTemplate - Create a new descriptor update template



## [](#_c_specification)C Specification

Updating a large `VkDescriptorSet` array **can** be an expensive operation since an application **must** specify one [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html) structure for each descriptor or descriptor array to update, each of which re-specifies the same state when updating the same descriptor in multiple descriptor sets. For cases when an application wishes to update the same set of descriptors in multiple descriptor sets allocated using the same `VkDescriptorSetLayout`, [vkUpdateDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSetWithTemplate.html) **can** be used as a replacement for [vkUpdateDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSets.html).

`VkDescriptorUpdateTemplate` allows implementations to convert a set of descriptor update operations on a single descriptor set to an internal format. In conjunction with [vkCmdPushDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSetWithTemplate.html) or [vkUpdateDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSetWithTemplate.html), this **can** be more efficient compared to calling [vkCmdPushDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSet.html) or [vkUpdateDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSets.html). The descriptors themselves are not specified in the `VkDescriptorUpdateTemplate`, rather, offsets into an application provided pointer to host memory are specified, which are combined with a pointer passed to [vkCmdPushDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSetWithTemplate.html) or [vkUpdateDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSetWithTemplate.html). This allows large batches of updates to be executed without having to convert application data structures into a strictly-defined Vulkan data structure.

To create a descriptor update template, call:

```c++
// Provided by VK_VERSION_1_1
VkResult vkCreateDescriptorUpdateTemplate(
    VkDevice                                    device,
    const VkDescriptorUpdateTemplateCreateInfo* pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDescriptorUpdateTemplate*                 pDescriptorUpdateTemplate);
```

or the equivalent command

```c++
// Provided by VK_KHR_descriptor_update_template
VkResult vkCreateDescriptorUpdateTemplateKHR(
    VkDevice                                    device,
    const VkDescriptorUpdateTemplateCreateInfo* pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDescriptorUpdateTemplate*                 pDescriptorUpdateTemplate);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the descriptor update template.
- `pCreateInfo` is a pointer to a [VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateCreateInfo.html) structure specifying the set of descriptors to update with a single call to [vkCmdPushDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSetWithTemplate.html) or [vkUpdateDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSetWithTemplate.html).
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pDescriptorUpdateTemplate` is a pointer to a `VkDescriptorUpdateTemplate` handle in which the resulting descriptor update template object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateDescriptorUpdateTemplate-device-parameter)VUID-vkCreateDescriptorUpdateTemplate-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateDescriptorUpdateTemplate-pCreateInfo-parameter)VUID-vkCreateDescriptorUpdateTemplate-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateCreateInfo.html) structure
- [](#VUID-vkCreateDescriptorUpdateTemplate-pAllocator-parameter)VUID-vkCreateDescriptorUpdateTemplate-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateDescriptorUpdateTemplate-pDescriptorUpdateTemplate-parameter)VUID-vkCreateDescriptorUpdateTemplate-pDescriptorUpdateTemplate-parameter  
  `pDescriptorUpdateTemplate` **must** be a valid pointer to a [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplate.html) handle
- [](#VUID-vkCreateDescriptorUpdateTemplate-device-queuecount)VUID-vkCreateDescriptorUpdateTemplate-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplate.html), [VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateCreateInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateDescriptorUpdateTemplate)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0