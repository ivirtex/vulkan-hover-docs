# VkCuModuleCreateInfoNVX(3) Manual Page

## Name

VkCuModuleCreateInfoNVX - Stub description of VkCuModuleCreateInfoNVX



## [](#_c_specification)C Specification

There is currently no specification language written for this type. This section acts only as placeholder and to avoid dead links in the specification and reference pages.

```c++
// Provided by VK_NVX_binary_import
typedef struct VkCuModuleCreateInfoNVX {
    VkStructureType    sType;
    const void*        pNext;
    size_t             dataSize;
    const void*        pData;
} VkCuModuleCreateInfoNVX;
```

## [](#_members)Members

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkCuModuleCreateInfoNVX-sType-sType)VUID-VkCuModuleCreateInfoNVX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CU_MODULE_CREATE_INFO_NVX`
- [](#VUID-VkCuModuleCreateInfoNVX-pNext-pNext)VUID-VkCuModuleCreateInfoNVX-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkCuModuleTexturingModeCreateInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleTexturingModeCreateInfoNVX.html)
- [](#VUID-VkCuModuleCreateInfoNVX-sType-unique)VUID-VkCuModuleCreateInfoNVX-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkCuModuleCreateInfoNVX-pData-parameter)VUID-VkCuModuleCreateInfoNVX-pData-parameter  
  If `dataSize` is not `0`, `pData` **must** be a valid pointer to an array of `dataSize` bytes

## [](#_see_also)See Also

[VK\_NVX\_binary\_import](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_binary_import.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateCuModuleNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCuModuleNVX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCuModuleCreateInfoNVX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0