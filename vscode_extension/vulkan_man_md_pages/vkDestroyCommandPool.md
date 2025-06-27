# vkDestroyCommandPool(3) Manual Page

## Name

vkDestroyCommandPool - Destroy a command pool object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a command pool, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyCommandPool(
    VkDevice                                    device,
    VkCommandPool                               commandPool,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the command pool.

- `commandPool` is the handle of the command pool to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

When a pool is destroyed, all command buffers allocated from the pool
are <a href="vkFreeCommandBuffers.html" target="_blank"
rel="noopener">freed</a>.

Any primary command buffer allocated from another
[VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) that is in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">recording or executable state</a> and has
a secondary command buffer allocated from `commandPool` recorded into
it, becomes <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">invalid</a>.

Valid Usage

- <a href="#VUID-vkDestroyCommandPool-commandPool-00041"
  id="VUID-vkDestroyCommandPool-commandPool-00041"></a>
  VUID-vkDestroyCommandPool-commandPool-00041  
  All `VkCommandBuffer` objects allocated from `commandPool` **must**
  not be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending state</a>

- <a href="#VUID-vkDestroyCommandPool-commandPool-00042"
  id="VUID-vkDestroyCommandPool-commandPool-00042"></a>
  VUID-vkDestroyCommandPool-commandPool-00042  
  If `VkAllocationCallbacks` were provided when `commandPool` was
  created, a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyCommandPool-commandPool-00043"
  id="VUID-vkDestroyCommandPool-commandPool-00043"></a>
  VUID-vkDestroyCommandPool-commandPool-00043  
  If no `VkAllocationCallbacks` were provided when `commandPool` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyCommandPool-device-parameter"
  id="VUID-vkDestroyCommandPool-device-parameter"></a>
  VUID-vkDestroyCommandPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyCommandPool-commandPool-parameter"
  id="VUID-vkDestroyCommandPool-commandPool-parameter"></a>
  VUID-vkDestroyCommandPool-commandPool-parameter  
  If `commandPool` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `commandPool` **must** be a valid [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html)
  handle

- <a href="#VUID-vkDestroyCommandPool-pAllocator-parameter"
  id="VUID-vkDestroyCommandPool-pAllocator-parameter"></a>
  VUID-vkDestroyCommandPool-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyCommandPool-commandPool-parent"
  id="VUID-vkDestroyCommandPool-commandPool-parent"></a>
  VUID-vkDestroyCommandPool-commandPool-parent  
  If `commandPool` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `commandPool` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyCommandPool"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
