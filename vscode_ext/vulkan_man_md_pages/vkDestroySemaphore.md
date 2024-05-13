# vkDestroySemaphore(3) Manual Page

## Name

vkDestroySemaphore - Destroy a semaphore object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a semaphore, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroySemaphore(
    VkDevice                                    device,
    VkSemaphore                                 semaphore,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the semaphore.

- `semaphore` is the handle of the semaphore to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroySemaphore-semaphore-05149"
  id="VUID-vkDestroySemaphore-semaphore-05149"></a>
  VUID-vkDestroySemaphore-semaphore-05149  
  All submitted batches that refer to `semaphore` **must** have
  completed execution

- <a href="#VUID-vkDestroySemaphore-semaphore-01138"
  id="VUID-vkDestroySemaphore-semaphore-01138"></a>
  VUID-vkDestroySemaphore-semaphore-01138  
  If `VkAllocationCallbacks` were provided when `semaphore` was created,
  a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroySemaphore-semaphore-01139"
  id="VUID-vkDestroySemaphore-semaphore-01139"></a>
  VUID-vkDestroySemaphore-semaphore-01139  
  If no `VkAllocationCallbacks` were provided when `semaphore` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroySemaphore-device-parameter"
  id="VUID-vkDestroySemaphore-device-parameter"></a>
  VUID-vkDestroySemaphore-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroySemaphore-semaphore-parameter"
  id="VUID-vkDestroySemaphore-semaphore-parameter"></a>
  VUID-vkDestroySemaphore-semaphore-parameter  
  If `semaphore` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

- <a href="#VUID-vkDestroySemaphore-pAllocator-parameter"
  id="VUID-vkDestroySemaphore-pAllocator-parameter"></a>
  VUID-vkDestroySemaphore-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroySemaphore-semaphore-parent"
  id="VUID-vkDestroySemaphore-semaphore-parent"></a>
  VUID-vkDestroySemaphore-semaphore-parent  
  If `semaphore` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `semaphore` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroySemaphore"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
