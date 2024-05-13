# vkDestroyDescriptorSetLayout(3) Manual Page

## Name

vkDestroyDescriptorSetLayout - Destroy a descriptor set layout object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a descriptor set layout, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyDescriptorSetLayout(
    VkDevice                                    device,
    VkDescriptorSetLayout                       descriptorSetLayout,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the descriptor set
  layout.

- `descriptorSetLayout` is the descriptor set layout to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-00284"
  id="VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-00284"></a>
  VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-00284  
  If `VkAllocationCallbacks` were provided when `descriptorSetLayout`
  was created, a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-00285"
  id="VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-00285"></a>
  VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-00285  
  If no `VkAllocationCallbacks` were provided when `descriptorSetLayout`
  was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyDescriptorSetLayout-device-parameter"
  id="VUID-vkDestroyDescriptorSetLayout-device-parameter"></a>
  VUID-vkDestroyDescriptorSetLayout-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-parameter"
  id="VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-parameter"></a>
  VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-parameter  
  If `descriptorSetLayout` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `descriptorSetLayout` **must** be a valid
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) handle

- <a href="#VUID-vkDestroyDescriptorSetLayout-pAllocator-parameter"
  id="VUID-vkDestroyDescriptorSetLayout-pAllocator-parameter"></a>
  VUID-vkDestroyDescriptorSetLayout-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-parent"
  id="VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-parent"></a>
  VUID-vkDestroyDescriptorSetLayout-descriptorSetLayout-parent  
  If `descriptorSetLayout` is a valid handle, it **must** have been
  created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `descriptorSetLayout` **must** be externally
  synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyDescriptorSetLayout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
