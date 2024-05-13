# VkCudaModuleCreateInfoNV(3) Manual Page

## Name

VkCudaModuleCreateInfoNV - Structure specifying the parameters to create
a CUDA Module



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCudaModuleCreateInfoNV` structure is defined as:

``` c
// Provided by VK_NV_cuda_kernel_launch
typedef struct VkCudaModuleCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    size_t             dataSize;
    const void*        pData;
} VkCudaModuleCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` **may** be `NULL` or **may** be a pointer to a structure
  extending this structure.

- `dataSize` is the length of the `pData` array.

- `pData` is a pointer to CUDA code

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCudaModuleCreateInfoNV-dataSize-09413"
  id="VUID-VkCudaModuleCreateInfoNV-dataSize-09413"></a>
  VUID-VkCudaModuleCreateInfoNV-dataSize-09413  
  `dataSize` **must** be the total size in bytes of the PTX files or
  binary cache passed to `pData`

Valid Usage (Implicit)

- <a href="#VUID-VkCudaModuleCreateInfoNV-sType-sType"
  id="VUID-VkCudaModuleCreateInfoNV-sType-sType"></a>
  VUID-VkCudaModuleCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CUDA_MODULE_CREATE_INFO_NV`

- <a href="#VUID-VkCudaModuleCreateInfoNV-pNext-pNext"
  id="VUID-VkCudaModuleCreateInfoNV-pNext-pNext"></a>
  VUID-VkCudaModuleCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCudaModuleCreateInfoNV-pData-parameter"
  id="VUID-VkCudaModuleCreateInfoNV-pData-parameter"></a>
  VUID-VkCudaModuleCreateInfoNV-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes

- <a href="#VUID-VkCudaModuleCreateInfoNV-dataSize-arraylength"
  id="VUID-VkCudaModuleCreateInfoNV-dataSize-arraylength"></a>
  VUID-VkCudaModuleCreateInfoNV-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cuda_kernel_launch](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cuda_kernel_launch.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCudaModuleNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCudaModuleCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
