# vkDestroyCudaFunctionNV(3) Manual Page

## Name

vkDestroyCudaFunctionNV - Destroy a CUDA function



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a CUDA function handle, call:

``` c
// Provided by VK_NV_cuda_kernel_launch
void vkDestroyCudaFunctionNV(
    VkDevice                                    device,
    VkCudaFunctionNV                            function,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the Function.

- `function` is the handle of the CUDA function to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyCudaFunctionNV-device-parameter"
  id="VUID-vkDestroyCudaFunctionNV-device-parameter"></a>
  VUID-vkDestroyCudaFunctionNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyCudaFunctionNV-function-parameter"
  id="VUID-vkDestroyCudaFunctionNV-function-parameter"></a>
  VUID-vkDestroyCudaFunctionNV-function-parameter  
  `function` **must** be a valid
  [VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaFunctionNV.html) handle

- <a href="#VUID-vkDestroyCudaFunctionNV-pAllocator-parameter"
  id="VUID-vkDestroyCudaFunctionNV-pAllocator-parameter"></a>
  VUID-vkDestroyCudaFunctionNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyCudaFunctionNV-function-parent"
  id="VUID-vkDestroyCudaFunctionNV-function-parent"></a>
  VUID-vkDestroyCudaFunctionNV-function-parent  
  `function` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cuda_kernel_launch](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cuda_kernel_launch.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaFunctionNV.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyCudaFunctionNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
