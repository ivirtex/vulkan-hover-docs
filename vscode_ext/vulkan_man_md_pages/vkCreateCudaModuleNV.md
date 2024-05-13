# vkCreateCudaModuleNV(3) Manual Page

## Name

vkCreateCudaModuleNV - Creates a new CUDA module object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a CUDA module, call:

``` c
// Provided by VK_NV_cuda_kernel_launch
VkResult vkCreateCudaModuleNV(
    VkDevice                                    device,
    const VkCudaModuleCreateInfoNV*             pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkCudaModuleNV*                             pModule);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the shader module.

- `pCreateInfo` is a pointer to a
  [VkCudaModuleCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleCreateInfoNV.html) structure.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pModule` is a pointer to a [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleNV.html)
  handle in which the resulting CUDA module object is returned.

## <a href="#_description" class="anchor"></a>Description

Once a CUDA module has been created, the application **may** create the
function entry point, which **must** refer to one function in the
module.

Valid Usage (Implicit)

- <a href="#VUID-vkCreateCudaModuleNV-device-parameter"
  id="VUID-vkCreateCudaModuleNV-device-parameter"></a>
  VUID-vkCreateCudaModuleNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateCudaModuleNV-pCreateInfo-parameter"
  id="VUID-vkCreateCudaModuleNV-pCreateInfo-parameter"></a>
  VUID-vkCreateCudaModuleNV-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkCudaModuleCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleCreateInfoNV.html) structure

- <a href="#VUID-vkCreateCudaModuleNV-pAllocator-parameter"
  id="VUID-vkCreateCudaModuleNV-pAllocator-parameter"></a>
  VUID-vkCreateCudaModuleNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateCudaModuleNV-pModule-parameter"
  id="VUID-vkCreateCudaModuleNV-pModule-parameter"></a>
  VUID-vkCreateCudaModuleNV-pModule-parameter  
  `pModule` **must** be a valid pointer to a
  [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleNV.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_INITIALIZATION_FAILED`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cuda_kernel_launch](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cuda_kernel_launch.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkCudaModuleCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleCreateInfoNV.html),
[VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleNV.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateCudaModuleNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
