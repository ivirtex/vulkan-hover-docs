# vkDestroyDescriptorPool(3) Manual Page

## Name

vkDestroyDescriptorPool - Destroy a descriptor pool object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a descriptor pool, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyDescriptorPool(
    VkDevice                                    device,
    VkDescriptorPool                            descriptorPool,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the descriptor pool.

- `descriptorPool` is the descriptor pool to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

When a pool is destroyed, all descriptor sets allocated from the pool
are implicitly freed and become invalid. Descriptor sets allocated from
a given pool do not need to be freed before destroying that descriptor
pool.

Valid Usage

- <a href="#VUID-vkDestroyDescriptorPool-descriptorPool-00303"
  id="VUID-vkDestroyDescriptorPool-descriptorPool-00303"></a>
  VUID-vkDestroyDescriptorPool-descriptorPool-00303  
  All submitted commands that refer to `descriptorPool` (via any
  allocated descriptor sets) **must** have completed execution

- <a href="#VUID-vkDestroyDescriptorPool-descriptorPool-00304"
  id="VUID-vkDestroyDescriptorPool-descriptorPool-00304"></a>
  VUID-vkDestroyDescriptorPool-descriptorPool-00304  
  If `VkAllocationCallbacks` were provided when `descriptorPool` was
  created, a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyDescriptorPool-descriptorPool-00305"
  id="VUID-vkDestroyDescriptorPool-descriptorPool-00305"></a>
  VUID-vkDestroyDescriptorPool-descriptorPool-00305  
  If no `VkAllocationCallbacks` were provided when `descriptorPool` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyDescriptorPool-device-parameter"
  id="VUID-vkDestroyDescriptorPool-device-parameter"></a>
  VUID-vkDestroyDescriptorPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyDescriptorPool-descriptorPool-parameter"
  id="VUID-vkDestroyDescriptorPool-descriptorPool-parameter"></a>
  VUID-vkDestroyDescriptorPool-descriptorPool-parameter  
  If `descriptorPool` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `descriptorPool` **must** be a valid
  [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html) handle

- <a href="#VUID-vkDestroyDescriptorPool-pAllocator-parameter"
  id="VUID-vkDestroyDescriptorPool-pAllocator-parameter"></a>
  VUID-vkDestroyDescriptorPool-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyDescriptorPool-descriptorPool-parent"
  id="VUID-vkDestroyDescriptorPool-descriptorPool-parent"></a>
  VUID-vkDestroyDescriptorPool-descriptorPool-parent  
  If `descriptorPool` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `descriptorPool` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyDescriptorPool"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
