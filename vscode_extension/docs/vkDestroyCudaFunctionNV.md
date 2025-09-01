# vkDestroyCudaFunctionNV(3) Manual Page

## Name

vkDestroyCudaFunctionNV - Destroy a CUDA function



## [](#_c_specification)C Specification

To destroy a CUDA function handle, call:

```c++
// Provided by VK_NV_cuda_kernel_launch
void vkDestroyCudaFunctionNV(
    VkDevice                                    device,
    VkCudaFunctionNV                            function,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the Function.
- `function` is the handle of the CUDA function to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkDestroyCudaFunctionNV-device-parameter)VUID-vkDestroyCudaFunctionNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyCudaFunctionNV-function-parameter)VUID-vkDestroyCudaFunctionNV-function-parameter  
  `function` **must** be a valid [VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaFunctionNV.html) handle
- [](#VUID-vkDestroyCudaFunctionNV-pAllocator-parameter)VUID-vkDestroyCudaFunctionNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyCudaFunctionNV-function-parent)VUID-vkDestroyCudaFunctionNV-function-parent  
  `function` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_NV\_cuda\_kernel\_launch](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cuda_kernel_launch.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaFunctionNV.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyCudaFunctionNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0