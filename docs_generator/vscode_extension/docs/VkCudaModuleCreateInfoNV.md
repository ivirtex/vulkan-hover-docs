# VkCudaModuleCreateInfoNV(3) Manual Page

## Name

VkCudaModuleCreateInfoNV - Structure specifying the parameters to create a CUDA Module



## [](#_c_specification)C Specification

The `VkCudaModuleCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_cuda_kernel_launch
typedef struct VkCudaModuleCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    size_t             dataSize;
    const void*        pData;
} VkCudaModuleCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` **may** be `NULL` or **may** be a pointer to a structure extending this structure.
- `dataSize` is the length of the `pData` array.
- `pData` is a pointer to CUDA code

## [](#_description)Description

Valid Usage

- [](#VUID-VkCudaModuleCreateInfoNV-dataSize-09413)VUID-VkCudaModuleCreateInfoNV-dataSize-09413  
  `dataSize` **must** be the total size in bytes of the PTX files or binary cache passed to `pData`

Valid Usage (Implicit)

- [](#VUID-VkCudaModuleCreateInfoNV-sType-sType)VUID-VkCudaModuleCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CUDA_MODULE_CREATE_INFO_NV`
- [](#VUID-VkCudaModuleCreateInfoNV-pNext-pNext)VUID-VkCudaModuleCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCudaModuleCreateInfoNV-pData-parameter)VUID-VkCudaModuleCreateInfoNV-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes
- [](#VUID-VkCudaModuleCreateInfoNV-dataSize-arraylength)VUID-VkCudaModuleCreateInfoNV-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_NV\_cuda\_kernel\_launch](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cuda_kernel_launch.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCudaModuleNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCudaModuleCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0