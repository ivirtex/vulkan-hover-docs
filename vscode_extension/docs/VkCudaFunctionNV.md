# VkCudaFunctionNV(3) Manual Page

## Name

VkCudaFunctionNV - Opaque handle to a CUDA function object



## [](#_c_specification)C Specification

CUDA functions are represented by `VkCudaFunctionNV` handles. Handles to `global` functions **may** then be used to issue a kernel launch (i.e. dispatch) from a commandbuffer. See [Dispatching Command for CUDA PTX kernel](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#cudadispatch).

```c++
// Provided by VK_NV_cuda_kernel_launch
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkCudaFunctionNV)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_NV\_cuda\_kernel\_launch](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cuda_kernel_launch.html), [VkCudaLaunchInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaLaunchInfoNV.html), [vkCreateCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCudaFunctionNV.html), [vkDestroyCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyCudaFunctionNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCudaFunctionNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0