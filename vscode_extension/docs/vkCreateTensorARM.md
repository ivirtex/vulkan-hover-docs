# vkCreateTensorARM(3) Manual Page

## Name

vkCreateTensorARM - Create a new tensor object



## [](#_c_specification)C Specification

To create tensors, call:

```c++
// Provided by VK_ARM_tensors
VkResult vkCreateTensorARM(
    VkDevice                                    device,
    const VkTensorCreateInfoARM*                pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkTensorARM*                                pTensor);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the tensor.
- `pCreateInfo` is a pointer to a [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html) structure containing parameters to be used to create the tensor.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pTensor` is a pointer to a [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) handle in which the resulting tensor object is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCreateTensorARM-tensors-09832)VUID-vkCreateTensorARM-tensors-09832  
  The [`tensors`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tensors) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCreateTensorARM-device-parameter)VUID-vkCreateTensorARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateTensorARM-pCreateInfo-parameter)VUID-vkCreateTensorARM-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html) structure
- [](#VUID-vkCreateTensorARM-pAllocator-parameter)VUID-vkCreateTensorARM-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateTensorARM-pTensor-parameter)VUID-vkCreateTensorARM-pTensor-parameter  
  `pTensor` **must** be a valid pointer to a [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) handle
- [](#VUID-vkCreateTensorARM-device-queuecount)VUID-vkCreateTensorARM-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html), [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateTensorARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0