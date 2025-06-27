# vkCreateCudaFunctionNV(3) Manual Page

## Name

vkCreateCudaFunctionNV - Creates a new CUDA function object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a CUDA function, call:

``` c
// Provided by VK_NV_cuda_kernel_launch
VkResult vkCreateCudaFunctionNV(
    VkDevice                                    device,
    const VkCudaFunctionCreateInfoNV*           pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkCudaFunctionNV*                           pFunction);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the shader module.

- `pCreateInfo` is a pointer to a
  [VkCudaFunctionCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaFunctionCreateInfoNV.html)
  structure.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pFunction` is a pointer to a
  [VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaFunctionNV.html) handle in which the
  resulting CUDA function object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCreateCudaFunctionNV-device-parameter"
  id="VUID-vkCreateCudaFunctionNV-device-parameter"></a>
  VUID-vkCreateCudaFunctionNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateCudaFunctionNV-pCreateInfo-parameter"
  id="VUID-vkCreateCudaFunctionNV-pCreateInfo-parameter"></a>
  VUID-vkCreateCudaFunctionNV-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkCudaFunctionCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaFunctionCreateInfoNV.html)
  structure

- <a href="#VUID-vkCreateCudaFunctionNV-pAllocator-parameter"
  id="VUID-vkCreateCudaFunctionNV-pAllocator-parameter"></a>
  VUID-vkCreateCudaFunctionNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateCudaFunctionNV-pFunction-parameter"
  id="VUID-vkCreateCudaFunctionNV-pFunction-parameter"></a>
  VUID-vkCreateCudaFunctionNV-pFunction-parameter  
  `pFunction` **must** be a valid pointer to a
  [VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaFunctionNV.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_INITIALIZATION_FAILED`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cuda_kernel_launch](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cuda_kernel_launch.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkCudaFunctionCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaFunctionCreateInfoNV.html),
[VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaFunctionNV.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateCudaFunctionNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
