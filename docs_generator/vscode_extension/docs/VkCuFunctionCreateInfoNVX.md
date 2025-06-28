# VkCuFunctionCreateInfoNVX(3) Manual Page

## Name

VkCuFunctionCreateInfoNVX - Stub description of VkCuFunctionCreateInfoNVX



## [](#_c_specification)C Specification

There is currently no specification language written for this type. This section acts only as placeholder and to avoid dead links in the specification and reference pages.

```c++
// Provided by VK_NVX_binary_import
typedef struct VkCuFunctionCreateInfoNVX {
    VkStructureType    sType;
    const void*        pNext;
    VkCuModuleNVX      module;
    const char*        pName;
} VkCuFunctionCreateInfoNVX;
```

## [](#_members)Members

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkCuFunctionCreateInfoNVX-sType-sType)VUID-VkCuFunctionCreateInfoNVX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CU_FUNCTION_CREATE_INFO_NVX`
- [](#VUID-VkCuFunctionCreateInfoNVX-pNext-pNext)VUID-VkCuFunctionCreateInfoNVX-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCuFunctionCreateInfoNVX-module-parameter)VUID-VkCuFunctionCreateInfoNVX-module-parameter  
  `module` **must** be a valid [VkCuModuleNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleNVX.html) handle
- [](#VUID-VkCuFunctionCreateInfoNVX-pName-parameter)VUID-VkCuFunctionCreateInfoNVX-pName-parameter  
  `pName` **must** be a null-terminated UTF-8 string

## [](#_see_also)See Also

[VK\_NVX\_binary\_import](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_binary_import.html), [VkCuModuleNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleNVX.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateCuFunctionNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCuFunctionNVX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCuFunctionCreateInfoNVX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0