# vkDestroyDescriptorUpdateTemplate(3) Manual Page

## Name

vkDestroyDescriptorUpdateTemplate - Destroy a descriptor update template
object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a descriptor update template, call:

``` c
// Provided by VK_VERSION_1_1
void vkDestroyDescriptorUpdateTemplate(
    VkDevice                                    device,
    VkDescriptorUpdateTemplate                  descriptorUpdateTemplate,
    const VkAllocationCallbacks*                pAllocator);
```

or the equivalent command

``` c
// Provided by VK_KHR_descriptor_update_template
void vkDestroyDescriptorUpdateTemplateKHR(
    VkDevice                                    device,
    VkDescriptorUpdateTemplate                  descriptorUpdateTemplate,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that has been used to create the
  descriptor update template

- `descriptorUpdateTemplate` is the descriptor update template to
  destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkDestroyDescriptorUpdateTemplate-descriptorSetLayout-00356"
  id="VUID-vkDestroyDescriptorUpdateTemplate-descriptorSetLayout-00356"></a>
  VUID-vkDestroyDescriptorUpdateTemplate-descriptorSetLayout-00356  
  If `VkAllocationCallbacks` were provided when
  `descriptorUpdateTemplate` was created, a compatible set of callbacks
  **must** be provided here

- <a
  href="#VUID-vkDestroyDescriptorUpdateTemplate-descriptorSetLayout-00357"
  id="VUID-vkDestroyDescriptorUpdateTemplate-descriptorSetLayout-00357"></a>
  VUID-vkDestroyDescriptorUpdateTemplate-descriptorSetLayout-00357  
  If no `VkAllocationCallbacks` were provided when
  `descriptorUpdateTemplate` was created, `pAllocator` **must** be
  `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyDescriptorUpdateTemplate-device-parameter"
  id="VUID-vkDestroyDescriptorUpdateTemplate-device-parameter"></a>
  VUID-vkDestroyDescriptorUpdateTemplate-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkDestroyDescriptorUpdateTemplate-descriptorUpdateTemplate-parameter"
  id="VUID-vkDestroyDescriptorUpdateTemplate-descriptorUpdateTemplate-parameter"></a>
  VUID-vkDestroyDescriptorUpdateTemplate-descriptorUpdateTemplate-parameter  
  If `descriptorUpdateTemplate` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `descriptorUpdateTemplate`
  **must** be a valid
  [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html) handle

- <a href="#VUID-vkDestroyDescriptorUpdateTemplate-pAllocator-parameter"
  id="VUID-vkDestroyDescriptorUpdateTemplate-pAllocator-parameter"></a>
  VUID-vkDestroyDescriptorUpdateTemplate-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a
  href="#VUID-vkDestroyDescriptorUpdateTemplate-descriptorUpdateTemplate-parent"
  id="VUID-vkDestroyDescriptorUpdateTemplate-descriptorUpdateTemplate-parent"></a>
  VUID-vkDestroyDescriptorUpdateTemplate-descriptorUpdateTemplate-parent  
  If `descriptorUpdateTemplate` is a valid handle, it **must** have been
  created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `descriptorUpdateTemplate` **must** be externally
  synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyDescriptorUpdateTemplate"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
