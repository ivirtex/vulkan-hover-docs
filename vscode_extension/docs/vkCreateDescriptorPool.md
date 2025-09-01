# vkCreateDescriptorPool(3) Manual Page

## Name

vkCreateDescriptorPool - Creates a descriptor pool object



## [](#_c_specification)C Specification

To create a descriptor pool object, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkCreateDescriptorPool(
    VkDevice                                    device,
    const VkDescriptorPoolCreateInfo*           pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDescriptorPool*                           pDescriptorPool);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the descriptor pool.
- `pCreateInfo` is a pointer to a [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html) structure specifying the state of the descriptor pool object.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pDescriptorPool` is a pointer to a [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPool.html) handle in which the resulting descriptor pool object is returned.

## [](#_description)Description

The created descriptor pool is returned in `pDescriptorPool`.

Valid Usage (Implicit)

- [](#VUID-vkCreateDescriptorPool-device-parameter)VUID-vkCreateDescriptorPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateDescriptorPool-pCreateInfo-parameter)VUID-vkCreateDescriptorPool-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html) structure
- [](#VUID-vkCreateDescriptorPool-pAllocator-parameter)VUID-vkCreateDescriptorPool-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateDescriptorPool-pDescriptorPool-parameter)VUID-vkCreateDescriptorPool-pDescriptorPool-parameter  
  `pDescriptorPool` **must** be a valid pointer to a [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPool.html) handle
- [](#VUID-vkCreateDescriptorPool-device-queuecount)VUID-vkCreateDescriptorPool-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_FRAGMENTATION_EXT`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPool.html), [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateDescriptorPool).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0