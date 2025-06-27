# vkCreateDescriptorUpdateTemplate(3) Manual Page

## Name

vkCreateDescriptorUpdateTemplate - Create a new descriptor update
template



## <a href="#_c_specification" class="anchor"></a>C Specification

Updating a large `VkDescriptorSet` array **can** be an expensive
operation since an application **must** specify one
[VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html) structure for each
descriptor or descriptor array to update, each of which re-specifies the
same state when updating the same descriptor in multiple descriptor
sets. For cases when an application wishes to update the same set of
descriptors in multiple descriptor sets allocated using the same
`VkDescriptorSetLayout`,
[vkUpdateDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplate.html)
**can** be used as a replacement for
[vkUpdateDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSets.html).

`VkDescriptorUpdateTemplate` allows implementations to convert a set of
descriptor update operations on a single descriptor set to an internal
format that, in conjunction with
[vkUpdateDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplate.html)
or
[vkCmdPushDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html)
, **can** be more efficient compared to calling
[vkUpdateDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSets.html) or
[vkCmdPushDescriptorSetKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetKHR.html) . The
descriptors themselves are not specified in the
`VkDescriptorUpdateTemplate`, rather, offsets into an application
provided pointer to host memory are specified, which are combined with a
pointer passed to
[vkUpdateDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplate.html)
or
[vkCmdPushDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html)
. This allows large batches of updates to be executed without having to
convert application data structures into a strictly-defined Vulkan data
structure.

To create a descriptor update template, call:

``` c
// Provided by VK_VERSION_1_1
VkResult vkCreateDescriptorUpdateTemplate(
    VkDevice                                    device,
    const VkDescriptorUpdateTemplateCreateInfo* pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDescriptorUpdateTemplate*                 pDescriptorUpdateTemplate);
```

or the equivalent command

``` c
// Provided by VK_KHR_descriptor_update_template
VkResult vkCreateDescriptorUpdateTemplateKHR(
    VkDevice                                    device,
    const VkDescriptorUpdateTemplateCreateInfo* pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDescriptorUpdateTemplate*                 pDescriptorUpdateTemplate);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the descriptor update
  template.

- `pCreateInfo` is a pointer to a
  [VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateInfo.html)
  structure specifying the set of descriptors to update with a single
  call to
  [vkCmdPushDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html)
  or
  [vkUpdateDescriptorSetWithTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplate.html).

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pDescriptorUpdateTemplate` is a pointer to a
  `VkDescriptorUpdateTemplate` handle in which the resulting descriptor
  update template object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCreateDescriptorUpdateTemplate-device-parameter"
  id="VUID-vkCreateDescriptorUpdateTemplate-device-parameter"></a>
  VUID-vkCreateDescriptorUpdateTemplate-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateDescriptorUpdateTemplate-pCreateInfo-parameter"
  id="VUID-vkCreateDescriptorUpdateTemplate-pCreateInfo-parameter"></a>
  VUID-vkCreateDescriptorUpdateTemplate-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateInfo.html)
  structure

- <a href="#VUID-vkCreateDescriptorUpdateTemplate-pAllocator-parameter"
  id="VUID-vkCreateDescriptorUpdateTemplate-pAllocator-parameter"></a>
  VUID-vkCreateDescriptorUpdateTemplate-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a
  href="#VUID-vkCreateDescriptorUpdateTemplate-pDescriptorUpdateTemplate-parameter"
  id="VUID-vkCreateDescriptorUpdateTemplate-pDescriptorUpdateTemplate-parameter"></a>
  VUID-vkCreateDescriptorUpdateTemplate-pDescriptorUpdateTemplate-parameter  
  `pDescriptorUpdateTemplate` **must** be a valid pointer to a
  [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html),
[VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateInfo.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateDescriptorUpdateTemplate"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
