# VkCudaFunctionCreateInfoNV(3) Manual Page

## Name

VkCudaFunctionCreateInfoNV - Structure specifying the parameters to create a CUDA Function



## [](#_c_specification)C Specification

The `VkCudaFunctionCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_cuda_kernel_launch
typedef struct VkCudaFunctionCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkCudaModuleNV     module;
    const char*        pName;
} VkCudaFunctionCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `module` is the CUDA [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleNV.html) module in which the function resides.
- `pName` is a null-terminated UTF-8 string containing the name of the shader entry point for this stage.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkCudaFunctionCreateInfoNV-sType-sType)VUID-VkCudaFunctionCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CUDA_FUNCTION_CREATE_INFO_NV`
- [](#VUID-VkCudaFunctionCreateInfoNV-pNext-pNext)VUID-VkCudaFunctionCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCudaFunctionCreateInfoNV-module-parameter)VUID-VkCudaFunctionCreateInfoNV-module-parameter  
  `module` **must** be a valid [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleNV.html) handle
- [](#VUID-VkCudaFunctionCreateInfoNV-pName-parameter)VUID-VkCudaFunctionCreateInfoNV-pName-parameter  
  `pName` **must** be a null-terminated UTF-8 string

## [](#_see_also)See Also

[VK\_NV\_cuda\_kernel\_launch](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cuda_kernel_launch.html), [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCudaFunctionNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCudaFunctionCreateInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0