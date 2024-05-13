# vkCreateSemaphore(3) Manual Page

## Name

vkCreateSemaphore - Create a new queue semaphore object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a semaphore, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateSemaphore(
    VkDevice                                    device,
    const VkSemaphoreCreateInfo*                pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSemaphore*                                pSemaphore);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the semaphore.

- `pCreateInfo` is a pointer to a
  [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html) structure
  containing information about how the semaphore is to be created.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pSemaphore` is a pointer to a handle in which the resulting semaphore
  object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCreateSemaphore-device-parameter"
  id="VUID-vkCreateSemaphore-device-parameter"></a>
  VUID-vkCreateSemaphore-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateSemaphore-pCreateInfo-parameter"
  id="VUID-vkCreateSemaphore-pCreateInfo-parameter"></a>
  VUID-vkCreateSemaphore-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html) structure

- <a href="#VUID-vkCreateSemaphore-pAllocator-parameter"
  id="VUID-vkCreateSemaphore-pAllocator-parameter"></a>
  VUID-vkCreateSemaphore-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateSemaphore-pSemaphore-parameter"
  id="VUID-vkCreateSemaphore-pSemaphore-parameter"></a>
  VUID-vkCreateSemaphore-pSemaphore-parameter  
  `pSemaphore` **must** be a valid pointer to a
  [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateSemaphore"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
