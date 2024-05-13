# VkCuFunctionCreateInfoNVX(3) Manual Page

## Name

VkCuFunctionCreateInfoNVX - Stub description of
VkCuFunctionCreateInfoNVX



## <a href="#_c_specification" class="anchor"></a>C Specification

There is currently no specification language written for this type. This
section acts only as placeholder and to avoid dead links in the
specification and reference pages.

``` c
// Provided by VK_NVX_binary_import
typedef struct VkCuFunctionCreateInfoNVX {
    VkStructureType    sType;
    const void*        pNext;
    VkCuModuleNVX      module;
    const char*        pName;
} VkCuFunctionCreateInfoNVX;
```

## <a href="#_members" class="anchor"></a>Members

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkCuFunctionCreateInfoNVX-sType-sType"
  id="VUID-VkCuFunctionCreateInfoNVX-sType-sType"></a>
  VUID-VkCuFunctionCreateInfoNVX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CU_FUNCTION_CREATE_INFO_NVX`

- <a href="#VUID-VkCuFunctionCreateInfoNVX-pNext-pNext"
  id="VUID-VkCuFunctionCreateInfoNVX-pNext-pNext"></a>
  VUID-VkCuFunctionCreateInfoNVX-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCuFunctionCreateInfoNVX-module-parameter"
  id="VUID-VkCuFunctionCreateInfoNVX-module-parameter"></a>
  VUID-VkCuFunctionCreateInfoNVX-module-parameter  
  `module` **must** be a valid [VkCuModuleNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuModuleNVX.html)
  handle

- <a href="#VUID-VkCuFunctionCreateInfoNVX-pName-parameter"
  id="VUID-VkCuFunctionCreateInfoNVX-pName-parameter"></a>
  VUID-VkCuFunctionCreateInfoNVX-pName-parameter  
  `pName` **must** be a null-terminated UTF-8 string

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NVX_binary_import](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NVX_binary_import.html),
[VkCuModuleNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuModuleNVX.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateCuFunctionNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCuFunctionNVX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCuFunctionCreateInfoNVX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
