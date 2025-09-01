# vkCreateCudaFunctionNV(3) Manual Page

## Name

vkCreateCudaFunctionNV - Creates a new CUDA function object



## [](#_c_specification)C Specification

To create a CUDA function, call:

```c++
// Provided by VK_NV_cuda_kernel_launch
VkResult vkCreateCudaFunctionNV(
    VkDevice                                    device,
    const VkCudaFunctionCreateInfoNV*           pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkCudaFunctionNV*                           pFunction);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the shader module.
- `pCreateInfo` is a pointer to a [VkCudaFunctionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaFunctionCreateInfoNV.html) structure.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pFunction` is a pointer to a [VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaFunctionNV.html) handle in which the resulting CUDA function object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateCudaFunctionNV-device-parameter)VUID-vkCreateCudaFunctionNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateCudaFunctionNV-pCreateInfo-parameter)VUID-vkCreateCudaFunctionNV-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkCudaFunctionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaFunctionCreateInfoNV.html) structure
- [](#VUID-vkCreateCudaFunctionNV-pAllocator-parameter)VUID-vkCreateCudaFunctionNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateCudaFunctionNV-pFunction-parameter)VUID-vkCreateCudaFunctionNV-pFunction-parameter  
  `pFunction` **must** be a valid pointer to a [VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaFunctionNV.html) handle
- [](#VUID-vkCreateCudaFunctionNV-device-queuecount)VUID-vkCreateCudaFunctionNV-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_cuda\_kernel\_launch](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cuda_kernel_launch.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkCudaFunctionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaFunctionCreateInfoNV.html), [VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaFunctionNV.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateCudaFunctionNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0