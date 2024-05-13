# VkCuModuleCreateInfoNVX(3) Manual Page

## Name

VkCuModuleCreateInfoNVX - Stub description of VkCuModuleCreateInfoNVX



## <a href="#_c_specification" class="anchor"></a>C Specification

There is currently no specification language written for this type. This
section acts only as placeholder and to avoid dead links in the
specification and reference pages.

``` c
// Provided by VK_NVX_binary_import
typedef struct VkCuModuleCreateInfoNVX {
    VkStructureType    sType;
    const void*        pNext;
    size_t             dataSize;
    const void*        pData;
} VkCuModuleCreateInfoNVX;
```

## <a href="#_members" class="anchor"></a>Members

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkCuModuleCreateInfoNVX-sType-sType"
  id="VUID-VkCuModuleCreateInfoNVX-sType-sType"></a>
  VUID-VkCuModuleCreateInfoNVX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CU_MODULE_CREATE_INFO_NVX`

- <a href="#VUID-VkCuModuleCreateInfoNVX-pNext-pNext"
  id="VUID-VkCuModuleCreateInfoNVX-pNext-pNext"></a>
  VUID-VkCuModuleCreateInfoNVX-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCuModuleCreateInfoNVX-pData-parameter"
  id="VUID-VkCuModuleCreateInfoNVX-pData-parameter"></a>
  VUID-VkCuModuleCreateInfoNVX-pData-parameter  
  If `dataSize` is not `0`, `pData` **must** be a valid pointer to an
  array of `dataSize` bytes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NVX_binary_import](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NVX_binary_import.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateCuModuleNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCuModuleNVX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCuModuleCreateInfoNVX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
