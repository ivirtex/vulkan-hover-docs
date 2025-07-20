# vkCreateCudaModuleNV(3) Manual Page

## Name

vkCreateCudaModuleNV - Creates a new CUDA module object



## [](#_c_specification)C Specification

To create a CUDA module, call:

```c++
// Provided by VK_NV_cuda_kernel_launch
VkResult vkCreateCudaModuleNV(
    VkDevice                                    device,
    const VkCudaModuleCreateInfoNV*             pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkCudaModuleNV*                             pModule);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the shader module.
- `pCreateInfo` is a pointer to a [VkCudaModuleCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleCreateInfoNV.html) structure.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pModule` is a pointer to a [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleNV.html) handle in which the resulting CUDA module object is returned.

## [](#_description)Description

Once a CUDA module has been created, the application **may** create the function entry point, which **must** refer to one function in the module.

Valid Usage (Implicit)

- [](#VUID-vkCreateCudaModuleNV-device-parameter)VUID-vkCreateCudaModuleNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateCudaModuleNV-pCreateInfo-parameter)VUID-vkCreateCudaModuleNV-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkCudaModuleCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleCreateInfoNV.html) structure
- [](#VUID-vkCreateCudaModuleNV-pAllocator-parameter)VUID-vkCreateCudaModuleNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateCudaModuleNV-pModule-parameter)VUID-vkCreateCudaModuleNV-pModule-parameter  
  `pModule` **must** be a valid pointer to a [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleNV.html) handle
- [](#VUID-vkCreateCudaModuleNV-device-queuecount)VUID-vkCreateCudaModuleNV-device-queuecount  
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

[VK\_NV\_cuda\_kernel\_launch](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cuda_kernel_launch.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkCudaModuleCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleCreateInfoNV.html), [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleNV.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateCudaModuleNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0