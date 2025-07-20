# vkCreateCommandPool(3) Manual Page

## Name

vkCreateCommandPool - Create a new command pool object



## [](#_c_specification)C Specification

To create a command pool, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkCreateCommandPool(
    VkDevice                                    device,
    const VkCommandPoolCreateInfo*              pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkCommandPool*                              pCommandPool);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the command pool.
- `pCreateInfo` is a pointer to a [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) structure specifying the state of the command pool object.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pCommandPool` is a pointer to a [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) handle in which the created pool is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCreateCommandPool-queueFamilyIndex-01937)VUID-vkCreateCommandPool-queueFamilyIndex-01937  
  `pCreateInfo->queueFamilyIndex` **must** be the index of a queue family available in the logical device `device`

Valid Usage (Implicit)

- [](#VUID-vkCreateCommandPool-device-parameter)VUID-vkCreateCommandPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateCommandPool-pCreateInfo-parameter)VUID-vkCreateCommandPool-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) structure
- [](#VUID-vkCreateCommandPool-pAllocator-parameter)VUID-vkCreateCommandPool-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateCommandPool-pCommandPool-parameter)VUID-vkCreateCommandPool-pCommandPool-parameter  
  `pCommandPool` **must** be a valid pointer to a [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) handle
- [](#VUID-vkCreateCommandPool-device-queuecount)VUID-vkCreateCommandPool-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html), [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateCommandPool)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0