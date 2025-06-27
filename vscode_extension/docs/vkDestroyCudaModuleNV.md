# vkDestroyCudaModuleNV(3) Manual Page

## Name

vkDestroyCudaModuleNV - Destroy a CUDA module



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a CUDA shader module, call:

``` c
// Provided by VK_NV_cuda_kernel_launch
void vkDestroyCudaModuleNV(
    VkDevice                                    device,
    VkCudaModuleNV                              module,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the shader module.

- `module` is the handle of the CUDA module to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyCudaModuleNV-device-parameter"
  id="VUID-vkDestroyCudaModuleNV-device-parameter"></a>
  VUID-vkDestroyCudaModuleNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyCudaModuleNV-module-parameter"
  id="VUID-vkDestroyCudaModuleNV-module-parameter"></a>
  VUID-vkDestroyCudaModuleNV-module-parameter  
  `module` **must** be a valid [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleNV.html)
  handle

- <a href="#VUID-vkDestroyCudaModuleNV-pAllocator-parameter"
  id="VUID-vkDestroyCudaModuleNV-pAllocator-parameter"></a>
  VUID-vkDestroyCudaModuleNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyCudaModuleNV-module-parent"
  id="VUID-vkDestroyCudaModuleNV-module-parent"></a>
  VUID-vkDestroyCudaModuleNV-module-parent  
  `module` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cuda_kernel_launch](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cuda_kernel_launch.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleNV.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyCudaModuleNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
