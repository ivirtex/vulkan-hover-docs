# VkCudaFunctionNV(3) Manual Page

## Name

VkCudaFunctionNV - Opaque handle to a CUDA function object



## <a href="#_c_specification" class="anchor"></a>C Specification

CUDA functions are represented by `VkCudaFunctionNV` handles. Handles to
*`global`* functions **may** then be used to issue a kernel launch (i.e.
dispatch) from a commandbuffer. See <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#cudadispatch"
target="_blank" rel="noopener">Dispatching Command for CUDA PTX
kernel</a>.

``` c
// Provided by VK_NV_cuda_kernel_launch
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkCudaFunctionNV)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cuda_kernel_launch](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cuda_kernel_launch.html),
[VkCudaLaunchInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaLaunchInfoNV.html),
[vkCreateCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCudaFunctionNV.html),
[vkDestroyCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyCudaFunctionNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCudaFunctionNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
