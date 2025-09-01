# VkCudaModuleNV(3) Manual Page

## Name

VkCudaModuleNV - Opaque handle to a CUDA module object



## [](#_c_specification)C Specification

CUDA modules **must** contain some kernel code and **must** expose at least one function entry point.

CUDA modules are represented by `VkCudaModuleNV` handles:

```c++
// Provided by VK_NV_cuda_kernel_launch
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkCudaModuleNV)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_NV\_cuda\_kernel\_launch](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cuda_kernel_launch.html), [VkCudaFunctionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaFunctionCreateInfoNV.html), [vkCreateCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCudaModuleNV.html), [vkDestroyCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyCudaModuleNV.html), [vkGetCudaModuleCacheNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetCudaModuleCacheNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCudaModuleNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0