# VkCudaFunctionCreateInfoNV(3) Manual Page

## Name

VkCudaFunctionCreateInfoNV - Structure specifying the parameters to
create a CUDA Function



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCudaFunctionCreateInfoNV` structure is defined as:

``` c
// Provided by VK_NV_cuda_kernel_launch
typedef struct VkCudaFunctionCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkCudaModuleNV     module;
    const char*        pName;
} VkCudaFunctionCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `module` is the CUDA [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleNV.html) module in
  which the function resides.

- `pName` is a null-terminated UTF-8 string containing the name of the
  shader entry point for this stage.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkCudaFunctionCreateInfoNV-sType-sType"
  id="VUID-VkCudaFunctionCreateInfoNV-sType-sType"></a>
  VUID-VkCudaFunctionCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CUDA_FUNCTION_CREATE_INFO_NV`

- <a href="#VUID-VkCudaFunctionCreateInfoNV-pNext-pNext"
  id="VUID-VkCudaFunctionCreateInfoNV-pNext-pNext"></a>
  VUID-VkCudaFunctionCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCudaFunctionCreateInfoNV-module-parameter"
  id="VUID-VkCudaFunctionCreateInfoNV-module-parameter"></a>
  VUID-VkCudaFunctionCreateInfoNV-module-parameter  
  `module` **must** be a valid [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleNV.html)
  handle

- <a href="#VUID-VkCudaFunctionCreateInfoNV-pName-parameter"
  id="VUID-VkCudaFunctionCreateInfoNV-pName-parameter"></a>
  VUID-VkCudaFunctionCreateInfoNV-pName-parameter  
  `pName` **must** be a null-terminated UTF-8 string

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cuda_kernel_launch](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cuda_kernel_launch.html),
[VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCudaFunctionNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCudaFunctionCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
