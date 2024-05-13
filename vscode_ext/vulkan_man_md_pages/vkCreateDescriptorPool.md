# vkCreateDescriptorPool(3) Manual Page

## Name

vkCreateDescriptorPool - Creates a descriptor pool object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a descriptor pool object, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateDescriptorPool(
    VkDevice                                    device,
    const VkDescriptorPoolCreateInfo*           pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDescriptorPool*                           pDescriptorPool);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the descriptor pool.

- `pCreateInfo` is a pointer to a
  [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateInfo.html)
  structure specifying the state of the descriptor pool object.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pDescriptorPool` is a pointer to a
  [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html) handle in which the
  resulting descriptor pool object is returned.

## <a href="#_description" class="anchor"></a>Description

The created descriptor pool is returned in `pDescriptorPool`.

Valid Usage (Implicit)

- <a href="#VUID-vkCreateDescriptorPool-device-parameter"
  id="VUID-vkCreateDescriptorPool-device-parameter"></a>
  VUID-vkCreateDescriptorPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateDescriptorPool-pCreateInfo-parameter"
  id="VUID-vkCreateDescriptorPool-pCreateInfo-parameter"></a>
  VUID-vkCreateDescriptorPool-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateInfo.html)
  structure

- <a href="#VUID-vkCreateDescriptorPool-pAllocator-parameter"
  id="VUID-vkCreateDescriptorPool-pAllocator-parameter"></a>
  VUID-vkCreateDescriptorPool-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateDescriptorPool-pDescriptorPool-parameter"
  id="VUID-vkCreateDescriptorPool-pDescriptorPool-parameter"></a>
  VUID-vkCreateDescriptorPool-pDescriptorPool-parameter  
  `pDescriptorPool` **must** be a valid pointer to a
  [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_FRAGMENTATION_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html),
[VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateInfo.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateDescriptorPool"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
