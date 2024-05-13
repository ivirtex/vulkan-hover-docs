# vkCreateCommandPool(3) Manual Page

## Name

vkCreateCommandPool - Create a new command pool object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a command pool, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateCommandPool(
    VkDevice                                    device,
    const VkCommandPoolCreateInfo*              pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkCommandPool*                              pCommandPool);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the command pool.

- `pCreateInfo` is a pointer to a
  [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html) structure
  specifying the state of the command pool object.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pCommandPool` is a pointer to a [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html)
  handle in which the created pool is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCreateCommandPool-queueFamilyIndex-01937"
  id="VUID-vkCreateCommandPool-queueFamilyIndex-01937"></a>
  VUID-vkCreateCommandPool-queueFamilyIndex-01937  
  `pCreateInfo->queueFamilyIndex` **must** be the index of a queue
  family available in the logical device `device`

Valid Usage (Implicit)

- <a href="#VUID-vkCreateCommandPool-device-parameter"
  id="VUID-vkCreateCommandPool-device-parameter"></a>
  VUID-vkCreateCommandPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateCommandPool-pCreateInfo-parameter"
  id="VUID-vkCreateCommandPool-pCreateInfo-parameter"></a>
  VUID-vkCreateCommandPool-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html) structure

- <a href="#VUID-vkCreateCommandPool-pAllocator-parameter"
  id="VUID-vkCreateCommandPool-pAllocator-parameter"></a>
  VUID-vkCreateCommandPool-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateCommandPool-pCommandPool-parameter"
  id="VUID-vkCreateCommandPool-pCommandPool-parameter"></a>
  VUID-vkCreateCommandPool-pCommandPool-parameter  
  `pCommandPool` **must** be a valid pointer to a
  [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html),
[VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateCommandPool"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
