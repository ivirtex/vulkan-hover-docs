# vkCreateQueryPool(3) Manual Page

## Name

vkCreateQueryPool - Create a new query pool object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a query pool, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateQueryPool(
    VkDevice                                    device,
    const VkQueryPoolCreateInfo*                pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkQueryPool*                                pQueryPool);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the query pool.

- `pCreateInfo` is a pointer to a
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html) structure
  containing the number and type of queries to be managed by the pool.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pQueryPool` is a pointer to a [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle
  in which the resulting query pool object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCreateQueryPool-device-09663"
  id="VUID-vkCreateQueryPool-device-09663"></a>
  VUID-vkCreateQueryPool-device-09663  
  `device` **must** support at least one queue family with one of the
  `VK_QUEUE_VIDEO_ENCODE_BIT_KHR`, `VK_QUEUE_VIDEO_DECODE_BIT_KHR`,
  `VK_QUEUE_COMPUTE_BIT`, or `VK_QUEUE_GRAPHICS_BIT` capabilities

Valid Usage (Implicit)

- <a href="#VUID-vkCreateQueryPool-device-parameter"
  id="VUID-vkCreateQueryPool-device-parameter"></a>
  VUID-vkCreateQueryPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateQueryPool-pCreateInfo-parameter"
  id="VUID-vkCreateQueryPool-pCreateInfo-parameter"></a>
  VUID-vkCreateQueryPool-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html) structure

- <a href="#VUID-vkCreateQueryPool-pAllocator-parameter"
  id="VUID-vkCreateQueryPool-pAllocator-parameter"></a>
  VUID-vkCreateQueryPool-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateQueryPool-pQueryPool-parameter"
  id="VUID-vkCreateQueryPool-pQueryPool-parameter"></a>
  VUID-vkCreateQueryPool-pQueryPool-parameter  
  `pQueryPool` **must** be a valid pointer to a
  [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html),
[VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateQueryPool"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
